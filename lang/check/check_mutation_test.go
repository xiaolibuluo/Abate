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
	"testing"

	"abate/lang/air"
	"abate/lang/gofront"
	"abate/lang/rule"
)

// TestCheckerFailurePaths injects every contract, option and source failure.
func TestCheckerFailurePaths(t *testing.T) {
	t.Parallel()
	current := unitRule(t, "GO-EG-001", "Ident~foo", rule.SeverityWarning)
	path := sourceFile(t, "a.go", "package p\nvar foo int\n")
	noFrontend := NewChecker(nil, []rule.Rule{current})
	if _, err := noFrontend.Check([]string{path}, DefaultOptions()); !errors.Is(err, ErrNoFrontend) {
		t.Fatalf("nil frontend = %v", err)
	}
	baseline := NewChecker(stubFrontend{}, []rule.Rule{current})
	badOptions := Options{Format: "bogus", MinSeverity: rule.SeverityInfo}
	if _, err := baseline.Check([]string{path}, badOptions); !errors.Is(err, ErrInvalidFormat) {
		t.Fatalf("bad format = %v", err)
	}
	missing := filepath.Join(t.TempDir(), "absent.go")
	if _, err := baseline.Check([]string{missing}, DefaultOptions()); !errors.Is(err, ErrReadSource) {
		t.Fatalf("missing path = %v", err)
	}
	if _, err := baseline.Check([]string{t.TempDir()}, DefaultOptions()); !errors.Is(err, ErrNoPaths) {
		t.Fatalf("empty directory = %v", err)
	}
	noRules := NewChecker(stubFrontend{}, nil)
	if _, err := noRules.Check([]string{path}, DefaultOptions()); !errors.Is(err, ErrNoRules) {
		t.Fatalf("no rules = %v", err)
	}
	broken := NewChecker(stubFrontend{failure: errors.New("boom")},
		[]rule.Rule{current})
	if _, err := broken.Check([]string{path}, DefaultOptions()); !errors.Is(err, ErrFrontendFailed) {
		t.Fatalf("frontend failure = %v", err)
	}
	withDiagnostic := NewChecker(stubFrontend{
		diagnostics: []air.Diagnostic{{
			Severity: air.SeverityError, Path: "a.go", Message: "broken",
		}},
	}, []rule.Rule{current})
	if _, err := withDiagnostic.Check([]string{path}, DefaultOptions()); !errors.Is(err, ErrParseFailed) {
		t.Fatalf("error diagnostic = %v", err)
	}
	unreadable := baseline.WithSourceReader(
		func(string) ([]byte, error) { return nil, errors.New("denied") })
	if _, err := unreadable.Check([]string{path}, DefaultOptions()); !errors.Is(err, ErrReadSource) {
		t.Fatalf("reader failure = %v", err)
	}
	tagged := baseline.WithLanguage("  GO  ")
	if tagged.language != "go" {
		t.Fatalf("WithLanguage did not normalize: %q", tagged.language)
	}
}

// TestExpandPathsAndBoundaries covers directory walking and skipped trees.
func TestExpandPathsAndBoundaries(t *testing.T) {
	t.Parallel()
	root := t.TempDir()
	for _, dir := range []string{"vendor", "third_party", ".git", "nested"} {
		if err := os.MkdirAll(filepath.Join(root, dir), 0o750); err != nil {
			t.Fatalf("mkdir %s: %v", dir, err)
		}
	}
	for _, name := range []string{
		"vendor/skip.go", "third_party/skip.go", ".git/skip.go",
		"nested/deep.go", "ok.go", "notes.txt",
	} {
		if err := os.WriteFile(filepath.Join(root, name),
			[]byte("package p\n"), 0o600); err != nil {
			t.Fatalf("write %s: %v", name, err)
		}
	}
	files, err := expandPaths([]string{"", root, filepath.Join(root, "ok.go"),
		filepath.Join(root, "notes.txt")})
	if err != nil {
		t.Fatalf("expandPaths: %v", err)
	}
	if len(files) != 2 {
		t.Fatalf("expandPaths returned %d files: %v", len(files), files)
	}
	for _, name := range files {
		if filepath.Base(name) == "skip.go" {
			t.Fatal("expandPaths entered an excluded directory")
		}
	}
	for _, name := range []string{"vendor", "third_party", ".git"} {
		if !skipDirectory(name) {
			t.Fatalf("skipDirectory(%q) = false", name)
		}
	}
	if skipDirectory("nested") {
		t.Fatal("skipDirectory excluded a normal directory")
	}
	first := unitRule(t, "GO-VET-001", "Ident~x", rule.SeverityWarning)
	second := unitRule(t, "GO-EG-001", "Ident~x", rule.SeverityWarning)
	ordered := filterRules([]rule.Rule{first, second}, DefaultOptions())
	if len(ordered) != 2 || ordered[0].Spec().ID != "GO-EG-001" {
		t.Fatal("filterRules did not order rules by identifier")
	}
	filtered := DefaultOptions()
	filtered.Languages = []string{"python"}
	if got := filterRules([]rule.Rule{first}, filtered.Normalize()); len(got) != 0 {
		t.Fatal("filterRules ignored the language filter")
	}
}

// TestShuffledInputIsStable proves path order cannot change the result.
func TestShuffledInputIsStable(t *testing.T) {
	t.Parallel()
	first := sourceFile(t, "a.go", "package p\nvar foo_bar int\n")
	second := sourceFile(t, "b.go", "package p\nvar fooBar int\n")
	current := unitRule(t, "GO-EG-001",
		`Ident/^[A-Za-z][A-Za-z0-9]*_[A-Za-z0-9_]+$/`, rule.SeverityWarning)
	checker := NewChecker(gofront.NewGoFrontend().WithTypeCheck(false),
		[]rule.Rule{current})
	options := DefaultOptions()
	options.Format = FormatJSON
	left, err := checker.Check([]string{first, second}, options)
	if err != nil {
		t.Fatalf("left check: %v", err)
	}
	right, err := checker.Check([]string{second, first}, options)
	if err != nil {
		t.Fatalf("right check: %v", err)
	}
	if left.Fingerprint != right.Fingerprint ||
		left.FingerprintHex() != right.FingerprintHex() {
		t.Fatal("shuffled path order changed the result fingerprint")
	}
	if len(left.Findings) != 1 {
		t.Fatalf("shuffled check found %d violations, want 1", len(left.Findings))
	}
}
