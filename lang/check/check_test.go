/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package check

import (
	"errors"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"

	"abate/lang/air"
	"abate/lang/rule"
)

// stubFrontend is a deterministic AIR frontend used to drive error paths.
type stubFrontend struct {
	unit        *air.Unit
	diagnostics []air.Diagnostic
	failure     error
}

// Parse returns the configured unit, diagnostics and failure.
func (s stubFrontend) Parse(air.SourceUnit) (*air.Unit, []air.Diagnostic, error) {
	return s.unit, s.diagnostics, s.failure
}

// ParseIncremental satisfies the rest of the AIR Frontend contract.
func (s stubFrontend) ParseIncremental(air.SourceUnit, *air.Unit,
	[]air.Edit) (*air.Unit, air.Diff, error) {
	return s.unit, air.Diff{}, s.failure
}

// unitRule compiles a tiny rule for checker tests.
func unitRule(t *testing.T, id, pattern string, severity rule.Severity) rule.Rule {
	t.Helper()
	compiled, err := rule.Compile(rule.Spec{
		ID: rule.RuleID(id), Reference: "Effective Go, Names",
		Severity: severity, Category: rule.CategoryStyle,
		Pattern: pattern, Message: "bad identifier",
	})
	if err != nil {
		t.Fatalf("compile %s: %v", id, err)
	}
	return compiled
}

// sourceFile writes a Go source file into a fresh temporary directory.
func sourceFile(t *testing.T, name, content string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), name)
	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatalf("write source: %v", err)
	}
	return path
}

// TestOptions covers defaults, normalization, validation and filters.
func TestOptions(t *testing.T) {
	t.Parallel()
	defaults := DefaultOptions()
	if defaults.Format != FormatConsole || defaults.MinSeverity != rule.SeverityInfo {
		t.Fatalf("DefaultOptions = %+v", defaults)
	}
	normalized := Options{
		Format: "  SARIF  ", Languages: []string{" GO ", "go", ""},
		Rules: []string{" GO-EG ", "GO-EG", ""}, MinSeverity: "bogus",
	}.Normalize()
	if normalized.Format != FormatSARIF ||
		normalized.MinSeverity != rule.SeverityInfo ||
		!slices.Equal(normalized.Languages, []string{"go"}) ||
		!slices.Equal(normalized.Rules, []string{"go-eg"}) {
		t.Fatalf("Normalize = %+v", normalized)
	}
	if err := normalized.Validate(); err != nil {
		t.Fatalf("normalized options rejected: %v", err)
	}
	badFormat := Options{Format: "bogus", MinSeverity: rule.SeverityInfo}
	if err := badFormat.Validate(); !errors.Is(err, ErrInvalidFormat) {
		t.Fatalf("bad format = %v", err)
	}
	badSeverity := Options{Format: FormatConsole, MinSeverity: "bogus"}
	if err := badSeverity.Validate(); !errors.Is(err, ErrInvalidOptions) {
		t.Fatalf("bad severity = %v", err)
	}
	badLanguage := Options{
		Format: FormatConsole, MinSeverity: rule.SeverityInfo,
		Languages: []string{" "},
	}
	if err := badLanguage.Validate(); !errors.Is(err, ErrInvalidOptions) {
		t.Fatalf("blank language = %v", err)
	}
	badRule := Options{
		Format: FormatConsole, MinSeverity: rule.SeverityInfo,
		Rules: []string{" "},
	}
	if err := badRule.Validate(); !errors.Is(err, ErrInvalidOptions) {
		t.Fatalf("blank rule filter = %v", err)
	}
	filters := Options{
		Format: FormatConsole, MinSeverity: rule.SeverityWarning,
		Languages: []string{"go"}, Rules: []string{"go-eg"},
	}.Normalize()
	eg := rule.Spec{ID: "GO-EG-001"}
	vet := rule.Spec{ID: "GO-VET-001"}
	py := rule.Spec{ID: "PY-PEP8-001"}
	if !filters.Accepts(eg, rule.SeverityError) ||
		filters.Accepts(eg, rule.SeverityInfo) ||
		filters.Accepts(vet, rule.SeverityWarning) ||
		filters.Accepts(py, rule.SeverityError) {
		t.Fatal("Accepts did not apply all four filters")
	}
	open := Options{Format: FormatConsole, MinSeverity: rule.SeverityInfo}
	if !open.Accepts(py, rule.SeverityInfo) {
		t.Fatal("unfiltered options must accept every rule")
	}
}

// TestSummaryAndAggregate covers counting, sorted keys and dedup ordering.
func TestSummaryAndAggregate(t *testing.T) {
	t.Parallel()
	findings := []rule.Finding{
		{RuleID: "B", Severity: rule.SeverityError, Path: "b.go", Line: 1, Column: 1},
		{RuleID: "A", Severity: rule.SeverityWarning, Path: "b.go", Line: 1, Column: 1},
		{RuleID: "A", Severity: rule.SeverityInfo, Path: "a.go", Line: 2, Column: 1},
		{RuleID: "A", Severity: rule.SeverityInfo, Path: "a.go", Line: 2, Column: 1},
		{RuleID: "A", Severity: rule.SeverityInfo, Path: "a.go", Line: 1, Column: 9},
	}
	ordered := aggregateFindings(findings)
	if len(ordered) != 4 {
		t.Fatalf("aggregateFindings returned %d findings, want 4", len(ordered))
	}
	if ordered[0].Path != "a.go" || ordered[0].Line != 1 ||
		ordered[3].RuleID != "B" {
		t.Fatalf("aggregate ordering is wrong: %v", ordered)
	}
	summary := NewSummary(ordered, 3, 5)
	if summary.Total != 4 || summary.Files != 3 || summary.Rules != 5 ||
		summary.ByRule["A"] != 3 || summary.BySeverity["error"] != 1 {
		t.Fatalf("summary = %+v", summary)
	}
	if !slices.Equal(summary.SortedRules(), []string{"A", "B"}) ||
		!slices.Equal(summary.SortedSeverities(),
			[]string{"error", "info", "warning"}) {
		t.Fatal("summary keys are not sorted")
	}
	if !strings.Contains(summary.Canonical(), "sev:error=1") ||
		!strings.Contains(summary.String(), "findings=4") {
		t.Fatal("summary rendering lost a field")
	}
	if sortedKeys(nil) == nil {
		t.Fatal("sortedKeys must return a slice for a nil map")
	}
	if (Summary{}).Canonical() == "" {
		t.Fatal("zero summary must still render")
	}
}

// TestResultAndRenderers covers fingerprints, all formats and SARIF levels.
func TestResultAndRenderers(t *testing.T) {
	t.Parallel()
	findings := []rule.Finding{
		{
			RuleID: "GO-EG-001", Path: "p.go", Line: 2, Column: 3,
			Severity: rule.SeverityError, Category: rule.CategoryNaming,
			Message: "bad", Reference: "Effective Go",
			Captures: []rule.Capture{{Name: "x", Text: "foo"}},
		},
		{
			RuleID: "GO-EG-002", Path: "p.go", Line: 4, Column: 1,
			Severity: rule.SeverityInfo, Category: rule.CategoryStyle,
			Message: "info", Reference: "Effective Go",
		},
	}
	summary := NewSummary(findings, 1, 2)
	result := Result{
		Findings: findings, Summary: summary, Format: FormatConsole,
		Fingerprint: computeFingerprint(findings, summary, FormatConsole),
	}
	if !result.HasViolations() || len(result.FingerprintHex()) != 64 ||
		result.Canonical() == "" {
		t.Fatal("result accessors are incomplete")
	}
	if computeFingerprint(findings, summary, FormatConsole) != result.Fingerprint {
		t.Fatal("fingerprint is not reproducible")
	}
	if computeFingerprint(findings, summary, FormatSARIF) == result.Fingerprint {
		t.Fatal("format must participate in the fingerprint")
	}
	for _, format := range []string{
		FormatConsole, FormatJSON, FormatSARIF, "unknown",
	} {
		if len(result.Render(format)) == 0 {
			t.Fatalf("Render(%q) produced no bytes", format)
		}
	}
	console := string(result.Render(FormatConsole))
	if !strings.Contains(console, "p.go:2:3") ||
		!strings.Contains(console, "fingerprint=") {
		t.Fatalf("console report lost a field: %q", console)
	}
	sarif := string(result.Render(FormatSARIF))
	for _, want := range []string{
		`"version": "2.1.0"`, `"ruleId"`, `"partialFingerprints"`,
		`"abateFinding/v1"`, `"level": "error"`, `"level": "note"`,
	} {
		if !strings.Contains(sarif, want) {
			t.Fatalf("SARIF lost %q", want)
		}
	}
	if !strings.Contains(string(result.Render(FormatJSON)), `"schema": "abate.check.v1"`) {
		t.Fatal("JSON report lost its schema field")
	}
	empty := Result{
		Findings: nil, Summary: NewSummary(nil, 0, 0), Format: FormatSARIF,
		Fingerprint: computeFingerprint(nil, NewSummary(nil, 0, 0), FormatSARIF),
	}
	if empty.HasViolations() || !strings.Contains(string(empty.Render(FormatSARIF)),
		`"results": []`) {
		t.Fatal("empty SARIF report is malformed")
	}
	if sarifLevel(rule.SeverityError) != "error" ||
		sarifLevel(rule.SeverityWarning) != "warning" ||
		sarifLevel(rule.Severity("bogus")) != "note" {
		t.Fatal("SARIF level mapping is wrong")
	}
	orphan := summary
	orphan.ByRule["GO-EG-777"] = 1
	if descriptor := ruleDescriptor(findings, "GO-EG-777"); descriptor.ID != "GO-EG-777" || descriptor.ShortDescription.Text != "" {
		t.Fatal("ruleDescriptor invented metadata for an unknown rule")
	}
	synthetic := sarifFinding(rule.Finding{RuleID: "X"}, 0)
	if synthetic.Locations[0].PhysicalLocation.Region.StartLine != 1 ||
		synthetic.Locations[0].PhysicalLocation.Region.StartColumn != 1 {
		t.Fatal("SARIF must clamp unknown positions to one")
	}
}

// TestExitCode covers the three check pipeline states.
func TestExitCode(t *testing.T) {
	t.Parallel()
	if ExitCode(Result{}, nil) != ExitClean {
		t.Fatal("clean result must exit zero")
	}
	violated := Result{Findings: []rule.Finding{{RuleID: "X"}}}
	if ExitCode(violated, nil) != ExitViolations {
		t.Fatal("violated result must exit one")
	}
	if ExitCode(violated, ErrNoFrontend) != ExitInternal {
		t.Fatal("an error must override violations with exit two")
	}
	if ExitClean != 0 || ExitViolations != 1 || ExitInternal != 2 {
		t.Fatal("check exit codes drifted from PLAN-07 section 4")
	}
}
