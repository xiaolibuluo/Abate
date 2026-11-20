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

// Checker orchestrates one language-neutral checking pipeline.
//
// It holds only the AIR Frontend contract, the AIR-facing Rule contract and an
// injectable source reader. No language-specific syntax type is referenced, so
// adding a new frontend never changes the orchestration layer.
type Checker struct {
	frontend   air.Frontend
	rules      []rule.Rule
	readSource func(string) ([]byte, error)
	language   string
}

// NewChecker returns a checker for one frontend and its compiled rule set.
func NewChecker(frontend air.Frontend, rules []rule.Rule) Checker {
	return Checker{
		frontend: frontend, rules: slices.Clone(rules),
		readSource: os.ReadFile, language: "go",
	}
}

// WithSourceReader replaces the source buffer reader, used by tests to avoid
// touching the real filesystem.
func (c Checker) WithSourceReader(reader func(string) ([]byte, error)) Checker {
	c.readSource = reader
	return c
}

// WithLanguage sets the language tag written into each in-memory source unit.
func (c Checker) WithLanguage(language string) Checker {
	c.language = strings.ToLower(strings.TrimSpace(language))
	return c
}

// Check expands paths, parses each buffer, runs the active rules, aggregates
// the findings and returns a deterministic result.
func (c Checker) Check(paths []string, opts Options) (Result, error) {
	if c.frontend == nil {
		return Result{}, ErrNoFrontend
	}
	options := opts.Normalize()
	if err := options.Validate(); err != nil {
		return Result{}, err
	}
	files, err := expandPaths(paths)
	if err != nil {
		return Result{}, err
	}
	if len(files) == 0 {
		return Result{}, ErrNoPaths
	}
	active := filterRules(c.rules, options)
	if len(active) == 0 {
		return Result{}, ErrNoRules
	}
	raw := make([]rule.Finding, 0)
	for _, path := range files {
		findings, checkErr := c.checkFile(path, active)
		if checkErr != nil {
			return Result{}, checkErr
		}
		raw = append(raw, findings...)
	}
	findings := aggregateFindings(raw)
	summary := NewSummary(findings, len(files), len(active))
	return Result{
		Findings: findings, Summary: summary, Format: options.Format,
		Fingerprint: computeFingerprint(findings, summary, options.Format),
	}, nil
}

// checkFile parses one source buffer and runs every active rule on its unit.
func (c Checker) checkFile(path string, active []rule.Rule) ([]rule.Finding, error) {
	source, err := c.readSource(path) //nolint:gosec // paths come from the caller-supplied, expanded file list.
	if err != nil {
		return nil, fmt.Errorf("%w: %s: %w", ErrReadSource, path, err)
	}
	unit := air.SourceUnit{
		Language: c.language,
		Files:    []air.SourceFile{air.NewSourceFile(path, source)},
	}
	built, diagnostics, err := c.frontend.Parse(unit)
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

// filterRules selects the active rules and returns them in identifier order.
func filterRules(rules []rule.Rule, opts Options) []rule.Rule {
	out := make([]rule.Rule, 0, len(rules))
	for _, current := range rules {
		spec := current.Spec()
		if opts.Accepts(spec, spec.Severity) {
			out = append(out, current)
		}
	}
	slices.SortFunc(out, func(a, b rule.Rule) int {
		return strings.Compare(string(a.Spec().ID), string(b.Spec().ID))
	})
	return out
}

// expandPaths turns files and directories into a sorted .go file list.
func expandPaths(paths []string) ([]string, error) {
	out := make([]string, 0, 16)
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
			if strings.HasSuffix(clean, ".go") {
				out = append(out, clean)
			}
			continue
		}
		if err := walkSourceDir(clean, &out); err != nil {
			return nil, err
		}
	}
	slices.Sort(out)
	return slices.Compact(out), nil
}

// walkSourceDir collects .go files below a directory in lexical order.
func walkSourceDir(dir string, out *[]string) error {
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
		if strings.HasSuffix(path, ".go") {
			*out = append(*out, path)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("%w: %w", ErrReadSource, err)
	}
	return nil
}

// skipDirectory reports directories excluded from a self-check walk.
func skipDirectory(name string) bool {
	switch name {
	case "vendor", "third_party", ".git":
		return true
	default:
		return false
	}
}
