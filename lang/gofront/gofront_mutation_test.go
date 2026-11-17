/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package gofront

import (
	"errors"
	"strings"
	"testing"

	"abate/lang/air"
)

// hasErrorDiagnostic reports whether a parse produced a syntax error finding.
func hasErrorDiagnostic(diagnostics []air.Diagnostic) bool {
	for _, diagnostic := range diagnostics {
		if diagnostic.IsError() {
			return true
		}
	}
	return false
}

// parseVariant parses one mutated source buffer.
func parseVariant(t *testing.T, frontend GoFrontend, src string) (*air.Unit, []air.Diagnostic) {
	t.Helper()
	unit, diagnostics, err := frontend.Parse(air.SourceUnit{
		Language: "go",
		Files:    []air.SourceFile{air.NewSourceFile("mut.go", []byte(src))},
	})
	if err != nil {
		t.Fatalf("mutation caused a contract failure: %v", err)
	}
	if unit == nil || len(unit.Files) != 1 || unit.Roots[0] == nil {
		t.Fatal("mutation produced an incomplete unit")
	}
	return unit, diagnostics
}

// TestMutationSourceCharacters replaces every character and requires detection.
func TestMutationSourceCharacters(t *testing.T) {
	t.Parallel()
	frontend := NewGoFrontend().WithTypeCheck(false)
	base, diagnostics := parseVariant(t, frontend, sampleGo)
	if hasErrorDiagnostic(diagnostics) {
		t.Fatal("the baseline source has a syntax error")
	}
	baseDigest := base.Fingerprint()
	detected, rejected := 0, 0
	for i := range sampleGo {
		mutated := sampleGo[:i] + "X" + sampleGo[i+1:]
		if mutated == sampleGo {
			continue
		}
		unit, findings := parseVariant(t, frontend, mutated)
		if hasErrorDiagnostic(findings) {
			rejected++
			continue
		}
		if unit.Fingerprint() != baseDigest {
			detected++
			continue
		}
		t.Fatalf("character %d (%q) mutation was not detected",
			i, sampleGo[i])
	}
	t.Logf("character mutations: %d changed AIR, %d produced syntax diagnostics",
		detected, rejected)
}

// TestMutationTokens replaces every lexical token and requires detection.
func TestMutationTokens(t *testing.T) {
	t.Parallel()
	frontend := NewGoFrontend().WithTypeCheck(false)
	base, _ := parseVariant(t, frontend, sampleGo)
	baseDigest := base.Fingerprint()
	tokens, err := ScanTokens([]byte(sampleGo))
	if err != nil {
		t.Fatalf("ScanTokens: %v", err)
	}
	checked := 0
	for i, item := range tokens {
		if !isLexicalToken(item.Kind) || item.EndOffset > len(sampleGo) ||
			item.Offset < 0 || item.Offset >= item.EndOffset {
			continue
		}
		replacement := "Q"
		switch item.Kind {
		case "STRING", "CHAR":
			replacement = `"z"`
		case "COMMENT":
			replacement = "// mutated\n"
		}
		mutated := sampleGo[:item.Offset] + replacement + sampleGo[item.EndOffset:]
		unit, findings := parseVariant(t, frontend, mutated)
		if hasErrorDiagnostic(findings) {
			checked++
			continue
		}
		if unit.Fingerprint() == baseDigest {
			t.Fatalf("token %d (%s %q) mutation was not detected",
				i, item.Kind, item.Text)
		}
		checked++
	}
	if checked < 50 {
		t.Fatalf("only %d tokens were mutated", checked)
	}
}

// TestMutationSyntaxDiagnostics checks structured errors and the hard cap.
func TestMutationSyntaxDiagnostics(t *testing.T) {
	t.Parallel()
	const broken = "package p\nfunc f( {\nvar = = =\nswitch {\n"
	frontend := NewGoFrontend().WithTypeCheck(false)
	_, diagnostics := parseVariant(t, frontend, broken)
	if len(diagnostics) == 0 || !hasErrorDiagnostic(diagnostics) {
		t.Fatal("a broken source produced no error diagnostic")
	}
	for _, diagnostic := range diagnostics {
		if !diagnostic.IsError() ||
			!strings.HasPrefix(diagnostic.Code, "LANG-GO-E-") ||
			diagnostic.Line < 1 || diagnostic.Column < 1 ||
			diagnostic.Message == "" || diagnostic.Path != "mut.go" {
			t.Fatalf("unstructured diagnostic: %+v", diagnostic)
		}
	}
	_, capped := parseVariant(t, frontend.WithMaxDiagnostics(3), broken)
	if len(capped) > 3 {
		t.Fatalf("diagnostic cap ignored: %d", len(capped))
	}
}

// TestMutationTokenFidelityRejectsDrift checks that AIR from one source cannot
// pass the token equivalence check against a different source.
func TestMutationTokenFidelityRejectsDrift(t *testing.T) {
	t.Parallel()
	const src = "package p\n\nvar A = 1\n"
	unit := mustParse(t, src)
	if err := VerifyTokenFidelity([]byte("package p\n\nvar A = 2\n"), unit); !errors.Is(err, ErrTokenFidelity) {
		t.Fatalf("drifted source accepted: %v", err)
	}
	if err := VerifyTokenFidelity([]byte(src), unit); err != nil {
		t.Fatalf("clean source rejected: %v", err)
	}
}

// TestMutationBoundarySources covers empty, comment-only, BOM, and huge inputs.
func TestMutationBoundarySources(t *testing.T) {
	t.Parallel()
	frontend := NewGoFrontend().WithTypeCheck(false)
	sources := []string{
		"",
		"// only a comment\n",
		"\xef\xbb\xbfpackage p\n",
		"package p\n\nvar A = 1\n\x00",
		"package p\n\nfunc F() { " + repeatText("_ = 1; ", 500) + "}\n",
	}
	for i, src := range sources {
		unit, diagnostics := parseVariant(t, frontend, src)
		if i == 0 || i == 1 || i == 3 {
			if !hasErrorDiagnostic(diagnostics) && unit.NodeCount() == 0 {
				t.Fatalf("boundary source %d produced an empty AIR", i)
			}
			continue
		}
		if hasErrorDiagnostic(diagnostics) {
			t.Fatalf("valid boundary source %d was rejected: %v", i, diagnostics)
		}
		if err := VerifyTokenFidelity([]byte(src), unit); err != nil {
			t.Fatalf("boundary source %d token fidelity: %v", i, err)
		}
	}
}
