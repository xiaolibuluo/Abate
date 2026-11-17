/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package frontkit

import (
	"strings"
	"testing"

	"abate/internal/determinism"
	"abate/lang/air"
)

var fuzzSeeds = [][]byte{
	{},
	[]byte("let a = 1\n"),
	[]byte(lexerFixture),
	{0xEF, 0xBB, 0xBF, 'l', 'e', 't', '\r', '\n'},
	{0xff, 0xfe, 0x00, 0x80, 0x41},
	[]byte("\"unterminated /* block */ r\"raw\" 'c'"),
	[]byte("fn f\n  x\n  y\nz\n"),
	[]byte("===> == = + - * / ! -> ====="),
	[]byte(strings.Repeat("!", 200)),
	[]byte(strings.Repeat(";", 300)),
}

// FuzzLexer feeds arbitrary bytes to both the plain and the indent-sensitive
// scanner. Seed corpora run as ordinary tests; no panic, no invalid token and
// no drift between two scans of the same buffer are tolerated.
func FuzzLexer(f *testing.F) {
	for _, seed := range fuzzSeeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, data []byte) {
		for _, sensitive := range []bool{false, true} {
			spec := newFakeSpec()
			spec.indentSensitive = sensitive
			lexer, err := NewLexer(spec, data)
			if err != nil {
				t.Fatalf("NewLexer: %v", err)
			}
			tokens, err := lexer.Tokenize()
			if err != nil {
				t.Fatalf("Tokenize: %v", err)
			}
			assertLexicalInvariants(t, string(data), tokens)
			again, err := lexer.Tokenize()
			if err != nil {
				t.Fatalf("second Tokenize: %v", err)
			}
			if string(tokenSnapshot(tokens)) != string(tokenSnapshot(again)) {
				t.Fatal("lexer output drifted between identical inputs")
			}
		}
	})
}

// FuzzParse drives the whole B02 pipeline from raw bytes: full parse, token
// fidelity, Pratt reduction with recovery, incremental reparse and canonical
// hashing. Seed corpora run as ordinary tests and must never panic.
func FuzzParse(f *testing.F) {
	for _, seed := range fuzzSeeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, data []byte) {
		spec := newFakeSpec()
		frontend := mustFrontend(t, spec)
		file := air.NewSourceFile("a.fk", data)
		unit, _, err := frontend.Parse(fakeSourceUnit(file))
		if err != nil {
			t.Fatalf("Parse: %v", err)
		}
		if err := unit.Validate(); err != nil {
			t.Fatalf("unit validation: %v", err)
		}
		digest := unit.Fingerprint()
		_ = digest
		state := determinism.HashState(unit)
		_ = state

		lexer, err := NewLexer(spec, data)
		if err != nil {
			t.Fatal(err)
		}
		tokens, err := lexer.Tokenize()
		if err != nil {
			t.Fatal(err)
		}
		if err := VerifyTokenFidelity(tokens, unit); err != nil {
			t.Fatalf("token fidelity: %v", err)
		}
		parser := newExpressionParser(t, tokens)
		if _, err := parser.ParseExpression(0); err != nil {
			parser.Recover()
		}
		if err := air.CheckIntervals(unit.Roots[0], len(data)); err != nil {
			t.Fatalf("AIR intervals: %v", err)
		}
		edits := []air.Edit{{Path: "a.fk", Offset: 0, EndOffset: 0,
			Text: []byte("x")}}
		next, diff, err := frontend.ParseIncremental(
			fakeSourceUnit(file), unit, edits)
		if err != nil {
			t.Fatalf("ParseIncremental: %v", err)
		}
		if next == nil || next.Fingerprint() == ([32]byte{}) || diff.Reused < 0 {
			t.Fatal("incremental unit is malformed")
		}
	})
}
