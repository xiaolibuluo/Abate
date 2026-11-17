/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package wasm

import "testing"

// TestParseFuncRecordsAliasAndDirectives verifies function parsing and name binding.
func TestParseFuncRecordsAliasAndDirectives(t *testing.T) {
	p := newWatParser([]byte(`$f (export "run") (param i32) i32.const 1)`))
	mod := &watModule{}
	if err := p.parseFunc(mod); err != nil {
		t.Fatalf("parseFunc: %v", err)
	}
	if len(mod.funcs) != 1 {
		t.Fatalf("funcs = %d, want 1", len(mod.funcs))
	}
	f := mod.funcs[0]
	if f.exportName != "run" || len(f.params) != 1 || f.params[0] != vtI32 {
		t.Errorf("func = %+v, want export run and one i32 param", f)
	}
	if len(f.body) != 1 || f.body[0].op != "i32.const" || f.body[0].imm != 1 {
		t.Errorf("func body = %+v, want one i32.const 1", f.body)
	}
	if idx, ok := p.names["$f"]; !ok || idx != 0 {
		t.Errorf("alias $f = %d/%v, want 0/true", idx, ok)
	}
}

// TestParseFuncRejectsMalformedBodies verifies failure paths.
func TestParseFuncRejectsMalformedBodies(t *testing.T) {
	cases := []struct {
		name string
		src  string
	}{
		{name: "bad directive", src: `(export 7))`},
		{name: "bad instruction", src: "7)"},
		{name: "missing closing paren on sub-form", src: "(nop 7)"},
		{name: "unterminated", src: "nop"},
	}
	for _, tc := range cases {
		p := newWatParser([]byte(tc.src))
		if err := p.parseFunc(&watModule{}); err == nil {
			t.Errorf("%s: expected an error", tc.name)
		}
	}
}

// TestRecordFuncNameIgnoresNonAliases verifies alias detection and lexer errors.
func TestRecordFuncNameIgnoresNonAliases(t *testing.T) {
	p := newWatParser([]byte("nop"))
	if err := p.recordFuncName(&watModule{}); err != nil {
		t.Fatalf("recordFuncName: %v", err)
	}
	if len(p.names) != 0 {
		t.Errorf("names = %v, want empty", p.names)
	}
	bad := newWatParser([]byte(`"open`))
	if err := bad.recordFuncName(&watModule{}); err == nil {
		t.Fatal("expected a lexer error from recordFuncName")
	}
}
