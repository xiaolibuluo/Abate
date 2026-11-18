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

	"abate/lang/gofront"
	"abate/lang/rule"
)

// validSpec returns a MultiChecker language spec backed by a real Go frontend
// and one compiled rule, so every test path exercises real parse + check.
func validSpec(t *testing.T, tag string, exts ...string) LanguageSpec {
	t.Helper()
	return LanguageSpec{
		Language: tag,
		Frontend: gofront.NewGoFrontend().WithTypeCheck(false),
		Rules:    []rule.Rule{unitRule(t, "GO-EG-001", "Ident~foo", rule.SeverityWarning)},
		Extensions: exts,
	}
}

// TestNewMultiCheckerRejectsDuplicates proves ambiguous dispatch tables fail
// closed instead of silently overriding a language or extension.
func TestNewMultiCheckerRejectsDuplicates(t *testing.T) {
	t.Parallel()
	if _, err := NewMultiChecker([]LanguageSpec{
		validSpec(t, "go", ".go"),
		validSpec(t, "go", ".other"),
	}); !errors.Is(err, ErrInvalidOptions) {
		t.Fatalf("duplicate language error = %v, want ErrInvalidOptions", err)
	}
	if _, err := NewMultiChecker([]LanguageSpec{
		validSpec(t, "go", ".same"),
		validSpec(t, "py", ".same"),
	}); !errors.Is(err, ErrInvalidOptions) {
		t.Fatalf("duplicate extension error = %v, want ErrInvalidOptions", err)
	}
}

// TestNewMultiCheckerRejectsEmptySpecs proves missing frontend/rules are
// classified with the existing sentinels.
func TestNewMultiCheckerRejectsEmptySpecs(t *testing.T) {
	t.Parallel()
	if _, err := NewMultiChecker([]LanguageSpec{{Language: "go", Extensions: []string{".go"}}});
		!errors.Is(err, ErrNoFrontend) {
		t.Fatalf("nil frontend error = %v, want ErrNoFrontend", err)
	}
	if _, err := NewMultiChecker([]LanguageSpec{{
		Language: "go", Extensions: []string{".go"},
		Frontend: gofront.NewGoFrontend().WithTypeCheck(false),
	}}); !errors.Is(err, ErrNoRules) {
		t.Fatalf("empty rules error = %v, want ErrNoRules", err)
	}
}

// TestMultiCheckerLanguagesSorted proves the configured language list is
// lexical, keeping any consumer output deterministic.
func TestMultiCheckerLanguagesSorted(t *testing.T) {
	t.Parallel()
	m, err := NewMultiChecker([]LanguageSpec{
		validSpec(t, "rs", ".rs"),
		validSpec(t, "go", ".go"),
		validSpec(t, "js", ".js"),
	})
	if err != nil {
		t.Fatalf("NewMultiChecker: %v", err)
	}
	got := m.Languages()
	want := []string{"go", "js", "rs"}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("languages = %v, want %v", got, want)
		}
	}
}

// TestMultiCheckerIgnoresUnknownExtensions proves unrecognized files are
// skipped rather than failing the whole run.
func TestMultiCheckerIgnoresUnknownExtensions(t *testing.T) {
	t.Parallel()
	m, err := NewMultiChecker([]LanguageSpec{validSpec(t, "go", ".go")})
	if err != nil {
		t.Fatalf("NewMultiChecker: %v", err)
	}
	dir := t.TempDir()
	goPath := filepath.Join(dir, "a.go")
	if err := os.WriteFile(goPath, []byte("package p\nvar foo_bar int\n"), 0o600); err != nil {
		t.Fatalf("write go: %v", err)
	}
	txtPath := filepath.Join(dir, "notes.txt")
	if err := os.WriteFile(txtPath, []byte("ignore me"), 0o600); err != nil {
		t.Fatalf("write txt: %v", err)
	}
	result, err := m.Check([]string{dir}, DefaultOptions())
	if err != nil {
		t.Fatalf("Check: %v", err)
	}
	if result.Summary.Files != 1 {
		t.Fatalf("files = %d, want 1 (unknown extension should be ignored)",
			result.Summary.Files)
	}
}
