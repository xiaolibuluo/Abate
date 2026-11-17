/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package check

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"abate/internal/determinism"
	"abate/lang/gofront"
	"abate/lang/rule"
)

// TestCheckEndToEnd drives a two-file fixture through the full pipeline and
// asserts the violation list, SARIF field order, exit states and reproducibility.
func TestCheckEndToEnd(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()
	bad := filepath.Join(dir, "bad.go")
	clean := filepath.Join(dir, "clean.go")
	if err := os.WriteFile(bad,
		[]byte("package p\n\nvar foo_bar int\n"), 0o600); err != nil {
		t.Fatalf("write bad fixture: %v", err)
	}
	if err := os.WriteFile(clean,
		[]byte("package p\n\nvar fooBar int\n"), 0o600); err != nil {
		t.Fatalf("write clean fixture: %v", err)
	}
	current := unitRule(t, "GO-EG-001",
		`Ident/^[A-Za-z][A-Za-z0-9]*_[A-Za-z0-9_]+$/`, rule.SeverityWarning)
	checker := NewChecker(gofront.NewGoFrontend().WithTypeCheck(false),
		[]rule.Rule{current})
	options := DefaultOptions()
	options.Format = FormatSARIF
	result, err := checker.Check([]string{dir}, options)
	if err != nil {
		t.Fatalf("end-to-end check: %v", err)
	}
	if len(result.Findings) != 1 {
		t.Fatalf("findings = %d, want 1: %v", len(result.Findings), result.Findings)
	}
	if result.Findings[0].Path != bad || result.Findings[0].Line != 3 {
		t.Fatalf("finding location is wrong: %+v", result.Findings[0])
	}
	if result.Summary.Files != 2 || result.Summary.Rules != 1 ||
		result.Summary.Total != 1 {
		t.Fatalf("summary = %+v", result.Summary)
	}
	sarif := string(result.Render(FormatSARIF))
	ordered := []string{
		`"$schema"`, `"version"`, `"runs"`, `"tool"`, `"driver"`,
		`"rules"`, `"results"`, `"ruleId"`, `"level"`, `"message"`,
		`"locations"`, `"partialFingerprints"`,
	}
	previous := -1
	for _, field := range ordered {
		at := strings.Index(sarif, field)
		if at < 0 || at < previous {
			t.Fatalf("SARIF field %q is missing or out of order", field)
		}
		previous = at
	}
	if _, err := os.Stat(filepath.Join(dir, "bad.go")); err != nil {
		t.Fatalf("fixture disappeared: %v", err)
	}
	determinism.MustMatch100(t, 42, func() []byte {
		got, checkErr := checker.Check([]string{dir}, options)
		if checkErr != nil {
			t.Fatalf("repeated check: %v", checkErr)
		}
		return got.Fingerprint[:]
	})
	if ExitCode(result, nil) != ExitViolations {
		t.Fatal("a violated result must exit one")
	}
	if ExitCode(Result{}, nil) != ExitClean {
		t.Fatal("a clean result must exit zero")
	}
	if ExitCode(Result{}, ErrNoFrontend) != ExitInternal {
		t.Fatal("an internal error must exit two")
	}
}

// TestCheckCleanFixture proves a compliant fixture yields exit zero.
func TestCheckCleanFixture(t *testing.T) {
	t.Parallel()
	path := sourceFile(t, "clean.go", "package p\n\nvar fooBar int\n")
	current := unitRule(t, "GO-EG-001",
		`Ident/^[A-Za-z][A-Za-z0-9]*_[A-Za-z0-9_]+$/`, rule.SeverityWarning)
	checker := NewChecker(gofront.NewGoFrontend().WithTypeCheck(false),
		[]rule.Rule{current})
	result, err := checker.Check([]string{path}, DefaultOptions())
	if err != nil {
		t.Fatalf("clean check: %v", err)
	}
	if result.HasViolations() || ExitCode(result, nil) != ExitClean {
		t.Fatalf("clean fixture produced %d findings", len(result.Findings))
	}
}
