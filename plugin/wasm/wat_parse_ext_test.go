/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package wasm

import (
	"bytes"
	"testing"
)

// TestParserTokenHelpersAcceptMatchingKinds verifies lookahead and typed accessors.
func TestParserTokenHelpersAcceptMatchingKinds(t *testing.T) {
	p := newWatParser([]byte(`("s" 7 word)`))
	if err := p.expect(tokLParen); err != nil {
		t.Fatalf("expect lparen: %v", err)
	}
	peeked, err := p.peekTok()
	if err != nil {
		t.Fatalf("peekTok: %v", err)
	}
	if peeked.kind != tokString || peeked.text != "s" {
		t.Errorf("peeked token = %+v, want string s", peeked)
	}
	s, err := p.nextString()
	if err != nil {
		t.Fatalf("nextString: %v", err)
	}
	if s != "s" {
		t.Errorf("nextString() = %q, want s", s)
	}
	v, err := p.nextInt()
	if err != nil {
		t.Fatalf("nextInt: %v", err)
	}
	if v != 7 {
		t.Errorf("nextInt() = %d, want 7", v)
	}
	id, err := p.nextIdent()
	if err != nil {
		t.Fatalf("nextIdent: %v", err)
	}
	if id != "word" {
		t.Errorf("nextIdent() = %q, want word", id)
	}
	last, err := p.peekTok()
	if err != nil {
		t.Fatalf("peekTok last: %v", err)
	}
	if last.kind != tokRParen {
		t.Errorf("last token = %+v, want right paren", last)
	}
	if err := p.expect(tokRParen); err != nil {
		t.Fatalf("expect rparen: %v", err)
	}
}

// TestParserTokenHelpersRejectWrongKinds verifies each typed accessor fails closed.
func TestParserTokenHelpersRejectWrongKinds(t *testing.T) {
	checks := []struct {
		name string
		src  string
		call func(*watParser) error
	}{
		{name: "expect on ident", src: "word", call: func(p *watParser) error { return p.expect(tokLParen) }},
		{name: "nextIdent on integer", src: "7", call: func(p *watParser) error { _, err := p.nextIdent(); return err }},
		{name: "nextString on ident", src: "word", call: func(p *watParser) error { _, err := p.nextString(); return err }},
		{name: "nextInt on ident", src: "word", call: func(p *watParser) error { _, err := p.nextInt(); return err }},
		{name: "expect on lexer error", src: `"open`, call: func(p *watParser) error { return p.expect(tokLParen) }},
		{name: "nextIdent on lexer error", src: `"open`, call: func(p *watParser) error { _, err := p.nextIdent(); return err }},
		{name: "nextString on lexer error", src: `"open`, call: func(p *watParser) error { _, err := p.nextString(); return err }},
		{name: "nextInt on lexer error", src: `"open`, call: func(p *watParser) error { _, err := p.nextInt(); return err }},
	}
	for _, tc := range checks {
		p := newWatParser([]byte(tc.src))
		if err := tc.call(p); err == nil {
			t.Errorf("%s: expected an error", tc.name)
		}
	}
}

// TestPeekTokPropagatesLexerErrors verifies lookahead surfaces lexer failures.
func TestPeekTokPropagatesLexerErrors(t *testing.T) {
	p := newWatParser([]byte(`"open`))
	if _, err := p.peekTok(); err == nil {
		t.Fatal("expected a lexer error from peekTok")
	}
}

// TestParseModuleParsesTopLevelForms verifies module boundaries and rejection.
func TestParseModuleParsesTopLevelForms(t *testing.T) {
	cases := []struct {
		name    string
		src     string
		wantErr bool
	}{
		{name: "empty module", src: "(module)", wantErr: false},
		{name: "unknown top-level form", src: "(module (unknown (nested 1) 2))", wantErr: false},
		{name: "missing lparen", src: "module)", wantErr: true},
		{name: "wrong keyword", src: "(other)", wantErr: true},
		{name: "unbalanced", src: "(module (func", wantErr: true},
		{name: "bare top-level token", src: "(module junk)", wantErr: true},
		{name: "non-ident header", src: "(module (7))", wantErr: true},
		{name: "lexer error", src: `(module (func (export "x))`, wantErr: true},
	}
	for _, tc := range cases {
		p := newWatParser([]byte(tc.src))
		mod, err := p.parseModule()
		if (err != nil) != tc.wantErr {
			t.Errorf("%s: error = %v, wantErr = %v", tc.name, err, tc.wantErr)
		}
		if !tc.wantErr && mod == nil {
			t.Errorf("%s: parseModule returned a nil module", tc.name)
		}
	}
}

// TestParseMemoryReadsLimitsAndExport verifies memory forms and failures.
func TestParseMemoryReadsLimitsAndExport(t *testing.T) {
	cases := []struct {
		name       string
		src        string
		wantErr    bool
		wantMin    uint32
		wantMax    uint32
		wantExport string
	}{
		{name: "min and max", src: "1 2)", wantMin: 1, wantMax: 2},
		{name: "min only", src: "1)", wantMin: 1},
		{name: "empty", src: ")"},
		{name: "export then min", src: `(export "mem") 1)`, wantMin: 1, wantExport: "mem"},
		{name: "unknown nested form", src: "(data 0) 3)", wantMin: 3},
		{name: "negative min", src: "-1)", wantErr: true},
		{name: "negative max", src: "1 -2)", wantErr: true},
		{name: "non-integer min", src: "word)", wantErr: true},
		{name: "bad export", src: "(export 7) 1)", wantErr: true},
		{name: "unterminated export", src: `(export "mem" 1)`, wantErr: true},
		{name: "lexer error", src: `"open`, wantErr: true},
	}
	for _, tc := range cases {
		p := newWatParser([]byte(tc.src))
		mod := &watModule{}
		err := p.parseMemory(mod)
		if (err != nil) != tc.wantErr {
			t.Errorf("%s: error = %v, wantErr = %v", tc.name, err, tc.wantErr)
			continue
		}
		if tc.wantErr {
			continue
		}
		if mod.memMin != tc.wantMin || mod.memMax != tc.wantMax || mod.memExport != tc.wantExport {
			t.Errorf("%s: memory = min %d max %d export %q, want %d %d %q",
				tc.name, mod.memMin, mod.memMax, mod.memExport, tc.wantMin, tc.wantMax, tc.wantExport)
		}
	}
}

// TestParseImportReadsFuncAndMemoryForms verifies import parsing and skipping.
func TestParseImportReadsFuncAndMemoryForms(t *testing.T) {
	funcCase := newWatParser([]byte(`"env" "f" (func $f (param i32) (result i32)) )`))
	mod := &watModule{}
	if err := funcCase.parseImport(mod); err != nil {
		t.Fatalf("parseImport func: %v", err)
	}
	if len(mod.imports) != 1 {
		t.Fatalf("imports = %d, want 1", len(mod.imports))
	}
	fi := mod.imports[0]
	if fi.module != "env" || fi.name != "f" {
		t.Errorf("import = %q/%q, want env/f", fi.module, fi.name)
	}
	if len(fi.params) != 1 || fi.params[0] != vtI32 || len(fi.results) != 1 || fi.results[0] != vtI32 {
		t.Errorf("import signature = %+v, want [i32] -> [i32]", fi)
	}
	if idx, ok := funcCase.names["$f"]; !ok || idx != 0 {
		t.Errorf("alias $f = %d/%v, want 0/true", idx, ok)
	}
	if funcCase.impN != 1 {
		t.Errorf("impN = %d, want 1", funcCase.impN)
	}

	memCase := newWatParser([]byte(`"env" "mem" (memory 1 2) )`))
	memMod := &watModule{}
	if err := memCase.parseImport(memMod); err != nil {
		t.Fatalf("parseImport memory: %v", err)
	}
	if memMod.memImport == nil || memMod.memImport.module != "env" || memMod.memImport.name != "mem" {
		t.Errorf("memImport = %+v, want env/mem", memMod.memImport)
	}
	if memMod.memMin != 1 || memMod.memMax != 2 {
		t.Errorf("memory limits = %d/%d, want 1/2", memMod.memMin, memMod.memMax)
	}

	unknown := newWatParser([]byte(`"env" "t" (table 1 funcref) )`))
	unknownMod := &watModule{}
	if err := unknown.parseImport(unknownMod); err != nil {
		t.Fatalf("parseImport unknown: %v", err)
	}
	if len(unknownMod.imports) != 0 || unknownMod.memImport != nil {
		t.Error("unknown import kind must be skipped without recording an import")
	}
}

// TestParseImportRejectsMalformedForms verifies all failure paths.
func TestParseImportRejectsMalformedForms(t *testing.T) {
	cases := []struct {
		name string
		src  string
	}{
		{name: "missing module string", src: `env "f" (func))`},
		{name: "missing name string", src: `"env" f (func))`},
		{name: "missing lparen", src: `"env" "f" func)`},
		{name: "bad func signature", src: `"env" "f" (func (param f32)) )`},
		{name: "negative memory min", src: `"env" "mem" (memory -1) )`},
		{name: "negative memory max", src: `"env" "mem" (memory 1 -2) )`},
		{name: "missing memory close", src: `"env" "mem" (memory 1`},
	}
	for _, tc := range cases {
		p := newWatParser([]byte(tc.src))
		if err := p.parseImport(&watModule{}); err == nil {
			t.Errorf("%s: expected an error", tc.name)
		}
	}
}

// TestBindImportAliasIgnoresPlainIdentifiers verifies only $-prefixed names bind.
func TestBindImportAliasIgnoresPlainIdentifiers(t *testing.T) {
	p := newWatParser([]byte("plain"))
	if err := p.bindImportAlias(&watModule{}); err != nil {
		t.Fatalf("bindImportAlias: %v", err)
	}
	if len(p.names) != 0 {
		t.Errorf("names = %v, want empty", p.names)
	}
	p2 := newWatParser([]byte("$alias"))
	if err := p2.bindImportAlias(&watModule{}); err != nil {
		t.Fatalf("bindImportAlias alias: %v", err)
	}
	if idx, ok := p2.names["$alias"]; !ok || idx != 0 {
		t.Errorf("alias index = %d/%v, want 0/true", idx, ok)
	}
	bad := newWatParser([]byte(`"open`))
	if err := bad.bindImportAlias(&watModule{}); err == nil {
		t.Fatal("expected a lexer error from bindImportAlias")
	}
}

// TestWatToWasmRichModuleIsDeterministic verifies the full translator pipeline.
func TestWatToWasmRichModuleIsDeterministic(t *testing.T) {
	src := `(module
  (memory (export "mem") 1 2)
  (func (export "calc") (param i32) (result i32)
    (local i32)
    block (result i32)
      i32.const 10
      i32.const 20
      i32.add
    end
    drop
    i32.const 7
    i32.const 3
    i32.add
    i32.const 5
    i32.mul
  )
)`
	first, err := watToWasm([]byte(src))
	if err != nil {
		t.Fatalf("watToWasm: %v", err)
	}
	second, err := watToWasm([]byte(src))
	if err != nil {
		t.Fatalf("watToWasm second pass: %v", err)
	}
	if !bytes.Equal(first, second) {
		t.Fatal("watToWasm is not deterministic")
	}
	if len(first) < 8 || !bytes.Equal(first[:4], wasmMagic) {
		t.Fatalf("watToWasm lost the wasm magic header: %v", first)
	}
}
