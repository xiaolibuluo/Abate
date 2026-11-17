/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package gofront

import (
	"errors"
	"testing"

	"abate/lang/air"
)

// fidelityFixtures are syntactically valid sources used for token equivalence.
var fidelityFixtures = []string{
	"package p\n",
	"package p\n\n// a comment\nvar A = 1\n",
	"package p\n\nfunc F(a int) int { return a + 1 }\n",
	"package p\n\ntype S struct {\n\tName string `json:\"name\"`\n}\n",
	"package p\n\nconst (\n\tA = iota\n\tB\n)\n",
	"package p\n\nfunc G() {\n\tdefer func() {}()\n\tgo func() { println(\"x\") }()\n}\n",
	"package p\n\nfunc H[T any](v T) T { return v }\n\nfunc use() { _ = H[int](3) }\n",
	"package p\n\nfunc R(ch chan int) {\n\tselect {\n\tcase n := <-ch:\n\t\t_ = n\n\tcase ch <- 1:\n\t}\n}\n",
	"package p\n\nfunc L() {\n\tvar s = []int{1, 2, 3}\n\t_ = s[0:2]\n\t_ = map[string]int{\"a\": 1}\n}\n",
	sampleGo,
}

// TestTokenFidelity checks bidirectional token equivalence on every fixture.
func TestTokenFidelity(t *testing.T) {
	t.Parallel()
	for i, src := range fidelityFixtures {
		unit, diagnostics := parseSource(t, src)
		for _, diagnostic := range diagnostics {
			if diagnostic.IsError() {
				t.Fatalf("fixture %d has a syntax error: %s", i, diagnostic.String())
			}
		}
		if err := VerifyTokenFidelity([]byte(src), unit); err != nil {
			t.Fatalf("fixture %d token fidelity: %v", i, err)
		}
	}
}

// TestScanTokensAndRegenerate checks the scanner and AIR projection agree.
func TestScanTokensAndRegenerate(t *testing.T) {
	t.Parallel()
	tokens, err := ScanTokens([]byte(sampleGo))
	if err != nil || len(tokens) == 0 {
		t.Fatalf("ScanTokens: %d tokens, %v", len(tokens), err)
	}
	comments := 0
	for _, item := range tokens {
		if item.Kind == "COMMENT" {
			comments++
		}
		if item.EndOffset < item.Offset {
			t.Fatalf("negative token span: %+v", item)
		}
	}
	if comments < 3 {
		t.Fatalf("expected at least three comments, got %d", comments)
	}
	unit := mustParse(t, sampleGo)
	regenerated := Regenerate(unit)
	if len(regenerated) == 0 || len(regenerated) > len(tokens) {
		t.Fatalf("implausible regeneration size: %d of %d",
			len(regenerated), len(tokens))
	}
	for i := 1; i < len(regenerated); i++ {
		if regenerated[i].Offset < regenerated[i-1].Offset {
			t.Fatal("regenerated tokens are not offset-ordered")
		}
	}
	if err := VerifyTokenFidelity([]byte(sampleGo), air.NewUnit(air.NewPackageRef("x", nil))); !errors.Is(err, ErrTokenFidelity) {
		t.Fatalf("an empty unit must fail fidelity: %v", err)
	}
}

// TestTokenFidelityRejectsInventedTokens checks the anti-fabrication side.
func TestTokenFidelityRejectsInventedTokens(t *testing.T) {
	t.Parallel()
	const src = "package p\n\nvar A = 1\n"
	fake := air.NewUnit(air.NewPackageRef("p", nil))
	invented := air.NewNode(air.KindIdent, air.NewPosition(99, 99, 9999, 10003), "ghost")
	if err := fake.AddFile(air.NewFileRef("x.go", []byte(src)), invented); err != nil {
		t.Fatal(err)
	}
	if err := VerifyTokenFidelity([]byte(src), fake); !errors.Is(err, ErrTokenFidelity) {
		t.Fatalf("invented token accepted: %v", err)
	}
}

// TestIntervalCoverageOnCleanFile checks replayed source intervals.
func TestIntervalCoverageOnCleanFile(t *testing.T) {
	t.Parallel()
	const src = "package p\n\nfunc F(a int) int { return a + 1 }\n"
	unit := mustParse(t, src)
	if err := air.CheckIntervals(unit.Roots[0], len(src)); err != nil {
		t.Fatalf("interval coverage: %v", err)
	}
	start, end := air.Span(unit.Roots[0])
	if start != 0 || end < len(src)-1 {
		t.Fatalf("root span [%d,%d) does not cover %d bytes", start, end, len(src))
	}
}

// TestTokenKindHelpers covers lexical classification tables.
func TestTokenKindHelpers(t *testing.T) {
	t.Parallel()
	for _, kind := range []string{"IDENT", "INT", "FLOAT", "IMAG", "CHAR", "STRING", "COMMENT"} {
		if !isLexicalToken(kind) {
			t.Fatalf("%s must be lexical", kind)
		}
	}
	for _, kind := range []string{"FUNC", "SEMICOLON", "ILLEGAL", ""} {
		if isLexicalToken(kind) {
			t.Fatalf("%s must not be lexical", kind)
		}
	}
	if !isLexicalNode(air.KindDirective) || isLexicalNode(air.KindCall) {
		t.Fatal("AIR lexical classification drifted")
	}
}
