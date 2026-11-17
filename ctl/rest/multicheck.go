/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package rest

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"abate/lang/air"
	"abate/lang/cfront"
	"abate/lang/check"
	"abate/lang/csfront"
	"abate/lang/gofront"
	"abate/lang/jsfront"
	"abate/lang/ktfront"
	"abate/lang/pyfront"
	"abate/lang/rsfront"
	"abate/lang/rule"
	cbuiltin "abate/lang/rule/builtin/c"
	csbuiltin "abate/lang/rule/builtin/cs"
	gobuiltin "abate/lang/rule/builtin/go"
	jsbuiltin "abate/lang/rule/builtin/js"
	ktbuiltin "abate/lang/rule/builtin/kt"
	pybuiltin "abate/lang/rule/builtin/py"
	rsbuiltin "abate/lang/rule/builtin/rs"
)

// langSpec binds one language to its AIR frontend, built-in rule loader and the
// file extensions a check walk should consider.
type langSpec struct {
	tag          string
	extensions   []string
	newFrontend  func() (air.Frontend, error)
	loadBuiltins func(*rule.Registry) error
}

// languageSpecs is the frozen language registry of the multi-checker. Each
// entry names a real, vendored frontend and its matching built-in rule set, so
// the control plane never embeds language-specific logic.
var languageSpecs = []langSpec{
	{
		tag: "go", extensions: []string{".go"},
		newFrontend: func() (air.Frontend, error) {
			return gofront.NewGoFrontend().WithTypeCheck(false), nil
		},
		loadBuiltins: gobuiltin.LoadBuiltins,
	},
	{
		tag: "python", extensions: []string{".py"},
		newFrontend:  pyfront.NewPyFrontend,
		loadBuiltins: pybuiltin.LoadBuiltins,
	},
	{
		tag: "js", extensions: []string{".js", ".ts"},
		newFrontend:  jsfront.NewJsFrontend,
		loadBuiltins: jsbuiltin.LoadBuiltins,
	},
	{
		tag: "c", extensions: []string{".c", ".h"},
		newFrontend:  cfront.NewCFrontend,
		loadBuiltins: cbuiltin.LoadBuiltins,
	},
	{
		tag: "rust", extensions: []string{".rs"},
		newFrontend:  rsfront.NewRsFrontend,
		loadBuiltins: rsbuiltin.LoadBuiltins,
	},
	{
		tag: "csharp", extensions: []string{".cs"},
		newFrontend:  csfront.NewCsFrontend,
		loadBuiltins: csbuiltin.LoadBuiltins,
	},
	{
		tag: "kotlin", extensions: []string{".kt"},
		newFrontend:  ktfront.NewKtFrontend,
		loadBuiltins: ktbuiltin.LoadBuiltins,
	},
}

// languageByTag returns the spec for a lowercased language tag, or nil.
func languageByTag(tag string) *langSpec {
	for i := range languageSpecs {
		if languageSpecs[i].tag == tag {
			return &languageSpecs[i]
		}
	}
	return nil
}

// SupportedLanguages returns the lowercased tags accepted by the multi-checker.
func SupportedLanguages() []string {
	out := make([]string, 0, len(languageSpecs))
	for _, spec := range languageSpecs {
		out = append(out, spec.tag)
	}
	slices.Sort(out)
	return out
}

// multiCheckRequest is the body of a language-neutral check task.
type multiCheckRequest struct {
	Language string   `json:"language"`
	Paths    []string `json:"paths"`
	Format   string   `json:"format"`
	Rules    []string `json:"rules"`
	Severity string   `json:"severity"`
}

// checkJobResult carries the deterministic outcome of one check task.
type checkJobResult struct {
	Result  check.Result  `json:"result"`
	Summary check.Summary `json:"summary"`
}

// runMultiCheck executes a language-neutral check against the filesystem.
//
// It is deterministic: file lists are sorted before parsing, and findings are
// canonically ordered before aggregation. The function performs no network I/O.
func runMultiCheck(req multiCheckRequest) (checkJobResult, error) {
	spec := languageByTag(strings.ToLower(strings.TrimSpace(req.Language)))
	if spec == nil {
		return checkJobResult{}, fmt.Errorf("%w: %s", errUnsupportedLanguage, req.Language)
	}
	files, err := expandLanguagesPaths(req.Paths, spec.extensions)
	if err != nil {
		return checkJobResult{}, err
	}
	if len(files) == 0 {
		return checkJobResult{}, errNoPaths
	}
	frontend, ferr := spec.newFrontend()
	if ferr != nil {
		return checkJobResult{}, fmt.Errorf("%w: %w", errFrontendFailed, ferr)
	}
	registry := rule.NewRegistry()
	if lerr := spec.loadBuiltins(registry); lerr != nil {
		return checkJobResult{}, fmt.Errorf("%w: %w", errRuleLoadFailed, lerr)
	}
	active := filterByLanguageAndRules(registry.For(spec.tag), req.Rules)
	if len(active) == 0 {
		return checkJobResult{}, errNoRules
	}
	sev := rule.Severity(req.Severity)
	if !sev.Valid() {
		sev = rule.SeverityInfo
	}
	raw := make([]rule.Finding, 0)
	for _, path := range files {
		// path is normalized by expandLanguagesPaths (filepath.Clean + extension
		// allow-list) from the caller's explicit list; reading user-nominated
		// source files is the checker's intended behavior. #nosec G304
		source, rerr := os.ReadFile(path) // #nosec G304
		if rerr != nil {
			return checkJobResult{}, fmt.Errorf("%w: %s: %w", errReadSource, path, rerr)
		}
		unit := air.SourceUnit{
			Language: spec.tag,
			Files:    []air.SourceFile{air.NewSourceFile(path, source)},
		}
		built, diagnostics, perr := frontend.Parse(unit)
		if perr != nil {
			return checkJobResult{}, fmt.Errorf("%w: %s: %w", errFrontendFailed, path, perr)
		}
		for _, diag := range diagnostics {
			if diag.IsError() {
				return checkJobResult{}, fmt.Errorf("%w: %s: %s",
					errParseFailed, path, diag.String())
			}
		}
		for _, current := range active {
			if sev.Rank() <= current.Spec().Severity.Rank() {
				raw = append(raw, current.Check(built)...)
			}
		}
	}
	findings := dedupeFindings(raw)
	summary := check.NewSummary(findings, len(files), len(active))
	format := req.Format
	if !validCheckFormat(format) {
		format = check.FormatConsole
	}
	return checkJobResult{
		Result:  check.Result{Findings: findings, Summary: summary, Format: format},
		Summary: summary,
	}, nil
}

// filterByLanguageAndRules selects rules of one language matching any prefix.
func filterByLanguageAndRules(rules []rule.Rule, prefixes []string) []rule.Rule {
	clean := make([]string, 0, len(prefixes))
	for _, p := range prefixes {
		if t := strings.TrimSpace(p); t != "" {
			clean = append(clean, t)
		}
	}
	out := make([]rule.Rule, 0, len(rules))
	for _, current := range rules {
		if len(clean) == 0 {
			out = append(out, current)
			continue
		}
		id := string(current.Spec().ID)
		for _, prefix := range clean {
			if strings.HasPrefix(id, prefix) {
				out = append(out, current)
				break
			}
		}
	}
	slices.SortFunc(out, func(a, b rule.Rule) int {
		return strings.Compare(string(a.Spec().ID), string(b.Spec().ID))
	})
	return out
}

// dedupeFindings removes byte-identical findings and orders them canonically so
// the result is reproducible regardless of walk or rule order.
func dedupeFindings(in []rule.Finding) []rule.Finding {
	seen := make(map[string]struct{}, len(in))
	out := make([]rule.Finding, 0, len(in))
	for _, f := range in {
		key := f.Canonical()
		if _, ok := seen[key]; ok {
			continue
		}
		seen[key] = struct{}{}
		out = append(out, f)
	}
	slices.SortFunc(out, func(a, b rule.Finding) int {
		if c := strings.Compare(a.Path, b.Path); c != 0 {
			return c
		}
		if a.Line != b.Line {
			return a.Line - b.Line
		}
		return strings.Compare(string(a.RuleID), string(b.RuleID))
	})
	return out
}
