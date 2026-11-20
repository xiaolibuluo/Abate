/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package check

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"abate/lang/air"
	"abate/lang/rule"
)

// MultiChecker checks several languages in one invocation, dispatching each
// source file to the frontend whose extension list matches. It reuses the same
// deterministic aggregation, summary and fingerprint pipeline as Checker so a
// single-language run and a multi-language run share one result contract.
type MultiChecker struct {
	langs      map[string]LanguageSpec
	byExt      map[string]string
	readSource func(string) ([]byte, error)
}

// NewMultiChecker builds a dispatch table from the supplied language specs.
// Duplicate language tags or extensions are rejected with an error rather than
// silently overridden, keeping dispatch deterministic and auditable.
func NewMultiChecker(specs []LanguageSpec) (MultiChecker, error) {
	m := MultiChecker{
		langs:      make(map[string]LanguageSpec, len(specs)),
		byExt:      make(map[string]string, len(specs)*2),
		readSource: os.ReadFile,
	}
	for _, spec := range specs {
		lang := strings.ToLower(strings.TrimSpace(spec.Language))
		if lang == "" || spec.Frontend == nil {
			return MultiChecker{}, fmt.Errorf("%w: language %q has no frontend", ErrNoFrontend, lang)
		}
		if len(spec.Rules) == 0 {
			return MultiChecker{}, fmt.Errorf("%w: language %q has no rules", ErrNoRules, lang)
		}
		if _, exists := m.langs[lang]; exists {
			return MultiChecker{}, fmt.Errorf("%w: duplicate language %q", ErrInvalidOptions, lang)
		}
		m.langs[lang] = spec
		for _, ext := range spec.Extensions {
			key := strings.ToLower(strings.TrimSpace(ext))
			if !strings.HasPrefix(key, ".") {
				key = "." + key
			}
			if _, exists := m.byExt[key]; exists {
				return MultiChecker{}, fmt.Errorf("%w: duplicate extension %q", ErrInvalidOptions, key)
			}
			m.byExt[key] = lang
		}
	}
	return m, nil
}

// WithSourceReader replaces the source buffer reader, used by tests.
func (m MultiChecker) WithSourceReader(reader func(string) ([]byte, error)) MultiChecker {
	m.readSource = reader
	return m
}

// Languages returns the configured language tags in lexical order.
func (m MultiChecker) Languages() []string {
	out := make([]string, 0, len(m.langs))
	for lang := range m.langs {
		out = append(out, lang)
	}
	slices.Sort(out)
	return out
}

// Check expands paths, dispatches every recognized source file to its language
// frontend, then aggregates all findings into one deterministic result.
func (m MultiChecker) Check(paths []string, opts Options) (Result, error) {
	options := opts.Normalize()
	if err := options.Validate(); err != nil {
		return Result{}, err
	}
	files, err := m.expandPaths(paths)
	if err != nil {
		return Result{}, err
	}
	if len(files) == 0 {
		return Result{}, ErrNoPaths
	}
	raw := make([]rule.Finding, 0)
	totalRules := 0
	usedLangs := make(map[string]bool, len(m.langs))
	for _, entry := range files {
		lang := m.byExt[entry.ext]
		if lang == "" || !m.languageEnabled(lang, options) {
			continue
		}
		spec := m.langs[lang]
		active := filterRules(spec.Rules, options)
		if len(active) == 0 {
			continue
		}
		findings, checkErr := m.checkFile(entry.path, spec.Frontend, lang, active)
		if checkErr != nil {
			return Result{}, checkErr
		}
		raw = append(raw, findings...)
		usedLangs[lang] = true
		totalRules += len(active)
	}
	findings := aggregateFindings(raw)
	filesChecked := 0
	for _, entry := range files {
		lang := m.byExt[entry.ext]
		if usedLangs[lang] {
			filesChecked++
		}
	}
	if filesChecked == 0 {
		return Result{}, ErrNoPaths
	}
	summary := NewSummary(findings, filesChecked, totalRules)
	return Result{
		Findings: findings, Summary: summary, Format: options.Format,
		Fingerprint: computeFingerprint(findings, summary, options.Format),
	}, nil
}

// languageEnabled applies the optional --languages filter to a dispatch tag.
func (MultiChecker) languageEnabled(lang string, opts Options) bool {
	if len(opts.Languages) == 0 {
		return true
	}
	return slices.Contains(opts.Languages, lang)
}

// checkFile parses one buffer with the given frontend and runs active rules.
func (m MultiChecker) checkFile(path string, frontend air.Frontend,
	lang string, active []rule.Rule) ([]rule.Finding, error) {
	source, err := m.readSource(path) //nolint:gosec // paths come from the caller-supplied, expanded file list.
	if err != nil {
		return nil, fmt.Errorf("%w: %s: %w", ErrReadSource, path, err)
	}
	unit := air.SourceUnit{
		Language: lang,
		Files:    []air.SourceFile{air.NewSourceFile(path, source)},
	}
	built, diagnostics, err := frontend.Parse(unit)
	if err != nil {
		return nil, fmt.Errorf("%w: %s: %w", ErrFrontendFailed, path, err)
	}
	for _, diagnostic := range diagnostics {
		if diagnostic.IsError() {
			return nil, fmt.Errorf("%w: %s: %s", ErrParseFailed,
				path, diagnostic.String())
		}
	}
	out := make([]rule.Finding, 0)
	for _, current := range active {
		out = append(out, current.Check(built)...)
	}
	return out, nil
}

// dispatchFile pairs a path with its resolved extension for dispatch.
type dispatchFile struct {
	path string
	ext  string
}

// expandPaths collects recognized source files from files and directories in
// lexical order; unrecognized extensions are ignored.
func (m MultiChecker) expandPaths(paths []string) ([]dispatchFile, error) {
	out := make([]dispatchFile, 0, 16)
	for _, raw := range paths {
		if strings.TrimSpace(raw) == "" {
			continue
		}
		clean := filepath.Clean(raw)
		info, err := os.Stat(clean)
		if err != nil {
			return nil, fmt.Errorf("%w: %s: %w", ErrReadSource, clean, err)
		}
		if !info.IsDir() {
			if ext := strings.ToLower(filepath.Ext(clean)); m.knownExt(ext) {
				out = append(out, dispatchFile{clean, ext})
			}
			continue
		}
		if err := m.walkSourceDir(clean, &out); err != nil {
			return nil, err
		}
	}
	slices.SortFunc(out, func(a, b dispatchFile) int {
		return strings.Compare(a.path, b.path)
	})
	return dedupSourceFiles(out), nil
}

// walkSourceDir collects recognized source files below a directory.
func (m MultiChecker) walkSourceDir(dir string, out *[]dispatchFile) error {
	err := filepath.WalkDir(dir, func(path string, entry fs.DirEntry,
		walkErr error) error {
		if walkErr != nil {
			// A directory may be removed by a concurrent test while we scan the
			// repository (e.g. a t.Cleanup temp dir). Ignore entries that vanished.
			if os.IsNotExist(walkErr) {
				return nil
			}
			return walkErr
		}
		if entry.IsDir() {
			if skipDirectory(entry.Name()) {
				return filepath.SkipDir
			}
			return nil
		}
		if ext := strings.ToLower(filepath.Ext(path)); m.knownExt(ext) {
			*out = append(*out, dispatchFile{path, ext})
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("%w: %w", ErrReadSource, err)
	}
	return nil
}

// knownExt reports whether an extension routes to a configured language.
func (m MultiChecker) knownExt(ext string) bool {
	_, ok := m.byExt[strings.ToLower(ext)]
	return ok
}

// dedupSourceFiles removes repeated paths, keeping lexical order.
func dedupSourceFiles(in []dispatchFile) []dispatchFile {
	out := make([]dispatchFile, 0, len(in))
	for _, entry := range in {
		if len(out) > 0 && out[len(out)-1].path == entry.path {
			continue
		}
		out = append(out, entry)
	}
	return out
}
