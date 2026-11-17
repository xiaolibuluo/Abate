/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package wasm

import "testing"

// lexAll drains a lexer into a token slice.
func lexAll(t *testing.T, src string) []watToken {
	t.Helper()
	l := newWatLexer([]byte(src))
	var toks []watToken
	for {
		tok, err := l.next()
		if err != nil {
			t.Fatalf("lexer error on %q: %v", src, err)
		}
		toks = append(toks, *tok)
		if tok.kind == tokEOF {
			return toks
		}
	}
}

// TestHexDigitDecodesEveryHexCharacter verifies all digit classes and rejections.
func TestHexDigitDecodesEveryHexCharacter(t *testing.T) {
	cases := []struct {
		in   rune
		want int
	}{
		{in: '0', want: 0}, {in: '1', want: 1}, {in: '2', want: 2},
		{in: '3', want: 3}, {in: '4', want: 4}, {in: '5', want: 5},
		{in: '6', want: 6}, {in: '7', want: 7}, {in: '8', want: 8}, {in: '9', want: 9},
		{in: 'a', want: 10}, {in: 'b', want: 11}, {in: 'c', want: 12},
		{in: 'd', want: 13}, {in: 'e', want: 14}, {in: 'f', want: 15},
		{in: 'A', want: 10}, {in: 'B', want: 11}, {in: 'C', want: 12},
		{in: 'D', want: 13}, {in: 'E', want: 14}, {in: 'F', want: 15},
		{in: 'g', want: -1}, {in: 'G', want: -1}, {in: 'z', want: -1},
		{in: 'Z', want: -1}, {in: ' ', want: -1}, {in: '/', want: -1},
		{in: ':', want: -1}, {in: '@', want: -1}, {in: 0, want: -1},
		{in: 'é', want: -1},
	}
	for _, tc := range cases {
		if got := hexDigit(tc.in); got != tc.want {
			t.Errorf("hexDigit(%q) = %d, want %d", tc.in, got, tc.want)
		}
	}
}

// TestParseWatIntHandlesDecimalAndHex verifies literal parsing and rejection.
func TestParseWatIntHandlesDecimalAndHex(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want int64
		ok   bool
	}{
		{name: "empty", in: "", want: 0, ok: false},
		{name: "zero", in: "0", want: 0, ok: true},
		{name: "decimal", in: "12345", want: 12345, ok: true},
		{name: "negative decimal", in: "-42", want: -42, ok: true},
		{name: "hex lower prefix", in: "0x2a", want: 42, ok: true},
		{name: "hex upper prefix", in: "0XFF", want: 255, ok: true},
		{name: "hex mixed case", in: "0xAbCd", want: 43981, ok: true},
		{name: "negative hex", in: "-0x10", want: -16, ok: true},
		{name: "hex zero", in: "0x0", want: 0, ok: true},
		{name: "invalid suffix", in: "12x", want: 0, ok: false},
		{name: "invalid hex digit", in: "0x1g", want: 0, ok: false},
		{name: "plus sign", in: "+1", want: 0, ok: false},
		{name: "embedded space", in: "1 2", want: 0, ok: false},
		{name: "decimal point", in: "1.5", want: 0, ok: false},
	}
	for _, tc := range cases {
		got, ok := parseWatInt(tc.in)
		if ok != tc.ok {
			t.Errorf("%s: parseWatInt(%q) ok = %v, want %v", tc.name, tc.in, ok, tc.ok)
			continue
		}
		if got != tc.want {
			t.Errorf("%s: parseWatInt(%q) = %d, want %d", tc.name, tc.in, got, tc.want)
		}
	}
}

// TestWatLexerTokenStream verifies token kinds, values and line tracking.
func TestWatLexerTokenStream(t *testing.T) {
	src := "(module \"hi\" 42 0x2A -7\n nop)"
	toks := lexAll(t, src)
	wantKinds := []watTokKind{tokLParen, tokIdent, tokString, tokInt, tokInt, tokInt, tokIdent, tokRParen, tokEOF}
	wantTexts := []string{"", "module", "hi", "", "", "", "nop", "", ""}
	wantValues := []int64{0, 0, 0, 42, 42, -7, 0, 0, 0}
	wantLines := []int{1, 1, 1, 1, 1, 1, 2, 2, 2}
	if len(toks) != len(wantKinds) {
		t.Fatalf("token count = %d, want %d", len(toks), len(wantKinds))
	}
	for i := range toks {
		if toks[i].kind != wantKinds[i] || toks[i].text != wantTexts[i] ||
			toks[i].value != wantValues[i] || toks[i].line != wantLines[i] {
			t.Errorf("token %d = %+v, want kind/text/value/line %v/%q/%d/%d",
				i, toks[i], wantKinds[i], wantTexts[i], wantValues[i], wantLines[i])
		}
	}
}

// TestWatLexerSkipsCommentsAndTracksLines verifies comment and whitespace handling.
func TestWatLexerSkipsCommentsAndTracksLines(t *testing.T) {
	toks := lexAll(t, "(module ; note\r\n\tnop)")
	wantKinds := []watTokKind{tokLParen, tokIdent, tokIdent, tokRParen, tokEOF}
	wantLines := []int{1, 1, 2, 2, 2}
	if len(toks) != len(wantKinds) {
		t.Fatalf("token count = %d, want %d", len(toks), len(wantKinds))
	}
	for i := range toks {
		if toks[i].kind != wantKinds[i] || toks[i].line != wantLines[i] {
			t.Errorf("token %d = %+v, want kind/line %v/%d", i, toks[i], wantKinds[i], wantLines[i])
		}
	}
	commentOnly := lexAll(t, "; trailing comment")
	if len(commentOnly) != 1 || commentOnly[0].kind != tokEOF {
		t.Errorf("comment-only input produced %v, want a single EOF token", commentOnly)
	}
}

// TestWatLexerStringEscapes verifies quoted strings and backslash escapes.
func TestWatLexerStringEscapes(t *testing.T) {
	toks := lexAll(t, `"a\"b" "plain"`)
	if len(toks) != 3 {
		t.Fatalf("token count = %d, want 3 (two strings and EOF)", len(toks))
	}
	if toks[0].kind != tokString || toks[0].text != `a\"b` {
		t.Errorf("first token = %+v, want string a\\\"b", toks[0])
	}
	if toks[1].kind != tokString || toks[1].text != "plain" {
		t.Errorf("second token = %+v, want string plain", toks[1])
	}
}

// TestWatLexerRejectsUnterminatedString verifies the unterminated string error.
func TestWatLexerRejectsUnterminatedString(t *testing.T) {
	l := newWatLexer([]byte(`"open`))
	if _, err := l.next(); err == nil {
		t.Fatal("expected an unterminated string error")
	}
}

// TestLexAtomRejectsDelimiterAtStart verifies the empty-atom error branch.
func TestLexAtomRejectsDelimiterAtStart(t *testing.T) {
	l := newWatLexer([]byte("(abc"))
	if _, err := l.lexAtom(); err == nil {
		t.Fatal("expected an empty-atom error")
	}
}
