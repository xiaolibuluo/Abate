/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package wasm

import "testing"

// TestSkipSExprConsumesBalancedForms verifies nested forms and EOF rejection.
func TestSkipSExprConsumesBalancedForms(t *testing.T) {
	cases := []struct {
		name    string
		src     string
		wantErr bool
	}{
		{name: "flat form", src: "foo 1 2)", wantErr: false},
		{name: "nested forms", src: "foo (bar (baz)) tail)", wantErr: false},
		{name: "unterminated", src: "foo (bar", wantErr: true},
		{name: "extra closing", src: "foo) )", wantErr: false},
	}
	for _, tc := range cases {
		p := newWatParser([]byte(tc.src))
		err := p.skipSExpr()
		if (err != nil) != tc.wantErr {
			t.Errorf("%s: skipSExpr error = %v, wantErr = %v", tc.name, err, tc.wantErr)
		}
	}
}

// TestSkipSExprStopsAtMatchingParen verifies the parser resumes after the form.
func TestSkipSExprStopsAtMatchingParen(t *testing.T) {
	p := newWatParser([]byte("foo 1 2) tail"))
	if err := p.skipSExpr(); err != nil {
		t.Fatalf("skipSExpr: %v", err)
	}
	tok, err := p.peekTok()
	if err != nil {
		t.Fatalf("peekTok: %v", err)
	}
	if tok.kind != tokIdent || tok.text != "tail" {
		t.Errorf("next token after skipSExpr = %+v, want ident tail", tok)
	}
}

// TestNextTypeAcceptsSupportedValueTypes verifies i32/i64 and rejections.
func TestNextTypeAcceptsSupportedValueTypes(t *testing.T) {
	cases := []struct {
		name    string
		src     string
		want    valType
		wantErr bool
	}{
		{name: "i32", src: "i32", want: vtI32},
		{name: "i64", src: "i64", want: vtI64},
		{name: "unsupported type", src: "f32", wantErr: true},
		{name: "integer token", src: "7", wantErr: true},
	}
	for _, tc := range cases {
		p := newWatParser([]byte(tc.src))
		got, err := p.nextType()
		if tc.wantErr {
			if err == nil {
				t.Errorf("%s: nextType() = %d, want an error", tc.name, got)
			}
			continue
		}
		if err != nil {
			t.Errorf("%s: nextType() error = %v, want nil", tc.name, err)
			continue
		}
		if got != tc.want {
			t.Errorf("%s: nextType() = %d, want %d", tc.name, got, tc.want)
		}
	}
}

// TestParseTypeListReadsUntilClose verifies vector parsing and error propagation.
func TestParseTypeListReadsUntilClose(t *testing.T) {
	p := newWatParser([]byte("i32 i64)"))
	var dst []valType
	if err := p.parseTypeList(&dst); err != nil {
		t.Fatalf("parseTypeList: %v", err)
	}
	if len(dst) != 2 || dst[0] != vtI32 || dst[1] != vtI64 {
		t.Fatalf("parsed types = %v, want [i32 i64]", dst)
	}
	var ignored []valType
	badType := newWatParser([]byte("f32)"))
	if err := badType.parseTypeList(&ignored); err == nil {
		t.Fatal("parseTypeList accepted an unsupported type")
	}
	unterminated := newWatParser([]byte("i32"))
	if err := unterminated.parseTypeList(&ignored); err == nil {
		t.Fatal("parseTypeList accepted an unterminated list")
	}
}

// TestParseFuncSigReadsParamsAndResults verifies the directive loop and unknown forms.
func TestParseFuncSigReadsParamsAndResults(t *testing.T) {
	p := newWatParser([]byte("(param i32 i64) (result i32) (local i64))"))
	var params []valType
	var results []valType
	if err := p.parseFuncSig(&params, &results); err != nil {
		t.Fatalf("parseFuncSig: %v", err)
	}
	if len(params) != 2 || params[0] != vtI32 || params[1] != vtI64 {
		t.Errorf("params = %v, want [i32 i64]", params)
	}
	if len(results) != 1 || results[0] != vtI32 {
		t.Errorf("results = %v, want [i32]", results)
	}
	bad := newWatParser([]byte("(param f32))"))
	if err := bad.parseFuncSig(&params, &results); err == nil {
		t.Fatal("parseFuncSig accepted an unsupported param type")
	}
	noParen := newWatParser([]byte("i32)"))
	if err := noParen.parseFuncSig(&params, &results); err == nil {
		t.Fatal("parseFuncSig accepted a bare type")
	}
}

// TestParseFuncDirectiveHandlesEveryDirective verifies directive outcomes and errors.
func TestParseFuncDirectiveHandlesEveryDirective(t *testing.T) {
	cases := []struct {
		name        string
		directive   string
		src         string
		wantHandled bool
		wantErr     bool
		wantExport  string
		wantParams  int
		wantResults int
		wantLocals  int
	}{
		{name: "export", directive: "export", src: `"run")`, wantHandled: true, wantExport: "run"},
		{name: "param", directive: "param", src: "i32 i64)", wantHandled: true, wantParams: 2},
		{name: "result", directive: "result", src: "i32)", wantHandled: true, wantResults: 1},
		{name: "local", directive: "local", src: "i64)", wantHandled: true, wantLocals: 1},
		{name: "unknown sub-form", directive: "nop", src: "i32)", wantHandled: false},
		{name: "export without string", directive: "export", src: "i32)", wantErr: true},
		{name: "export without paren", directive: "export", src: `"run"`, wantErr: true, wantExport: "run"},
		{name: "param bad type", directive: "param", src: "f32)", wantErr: true},
		{name: "result bad type", directive: "result", src: "f32)", wantErr: true},
		{name: "local bad type", directive: "local", src: "f32)", wantErr: true},
	}
	for _, tc := range cases {
		p := newWatParser([]byte(tc.src))
		f := watFunc{}
		handled, err := p.parseFuncDirective(&f, tc.directive)
		if (err != nil) != tc.wantErr {
			t.Errorf("%s: error = %v, wantErr = %v", tc.name, err, tc.wantErr)
		}
		if handled != tc.wantHandled {
			t.Errorf("%s: handled = %v, want %v", tc.name, handled, tc.wantHandled)
		}
		if f.exportName != tc.wantExport || len(f.params) != tc.wantParams ||
			len(f.results) != tc.wantResults || len(f.locals) != tc.wantLocals {
			t.Errorf("%s: func = %+v, want export %q params %d results %d locals %d",
				tc.name, f, tc.wantExport, tc.wantParams, tc.wantResults, tc.wantLocals)
		}
	}
}
