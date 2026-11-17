/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package yamlmini

import (
	"reflect"
	"strings"
	"testing"
)

// TestDoubleQuotedDecodeSuccess covers every escape form of the quoted decoder.
func TestDoubleQuotedDecodeSuccess(t *testing.T) {
	t.Parallel()
	cases := []struct{ raw, want string }{
		{`"plain"`, "plain"},
		{`"\0\a\b\t\n\v\f\r\e"`, "\x00\a\b\t\n\v\f\r\x1b"},
		{"\"a\\tb\"", "a\tb"},
		{`"a\ b"`, "a b"},
		{`"a\"b\/c\\d"`, "a\"b/c\\d"},
		{`"\N\_\L\P"`, "\u0085\u00a0\u2028\u2029"},
		{`"\x41\u0042\U00000043"`, "ABC"},
		{`"\uD83D\uDE00"`, "\U0001F600"},
	}
	for _, tc := range cases {
		got, err := decodeDoubleQuoted(tc.raw, 7)
		if err != nil {
			t.Fatalf("decodeDoubleQuoted(%q): %v", tc.raw, err)
		}
		if got != tc.want {
			t.Fatalf("decodeDoubleQuoted(%q) = %q, want %q", tc.raw, got, tc.want)
		}
	}
}

// TestDoubleQuotedDecodeErrors verifies that every malformed escape is rejected.
func TestDoubleQuotedDecodeErrors(t *testing.T) {
	t.Parallel()
	cases := []struct{ raw, want string }{
		{`"unterminated`, "unterminated"},
		{"\"abc\\\"", "dangling escape"},
		{`"\x4"`, "truncated unicode escape"},
		{`"\xZZ"`, "invalid unicode escape"},
		{`"\U00110000"`, "invalid code point"},
		{`"\uD800"`, "lone surrogate"},
		{`"\q"`, "unknown escape"},
	}
	for _, tc := range cases {
		_, err := decodeDoubleQuoted(tc.raw, 3)
		if err == nil || !strings.Contains(err.Error(), tc.want) {
			t.Fatalf("decodeDoubleQuoted(%q) err = %v, want containing %q", tc.raw, err, tc.want)
		}
	}
}

// TestSingleQuotedDecode covers doubled quotes and malformed single-quoted scalars.
func TestSingleQuotedDecode(t *testing.T) {
	t.Parallel()
	got, err := decodeSingleQuoted(`'it''s ok'`, 1)
	if err != nil || got != "it's ok" {
		t.Fatalf("decodeSingleQuoted = %q, %v", got, err)
	}
	for _, raw := range []string{`'unterminated`, `'a'b'`} {
		if _, err := decodeSingleQuoted(raw, 2); err == nil {
			t.Fatalf("decodeSingleQuoted(%q) accepted malformed input", raw)
		}
	}
}

// TestPlainScalarClassification covers int/float/bool/null/string decoding.
func TestPlainScalarClassification(t *testing.T) {
	t.Parallel()
	cases := []struct {
		raw  string
		want any
	}{
		{"0", 0}, {"-7", -7}, {"+9", 9},
		{"18446744073709551615", uint64(18446744073709551615)},
		{"1.5", 1.5}, {"2e3", 2e3}, {"-0.5", -0.5},
		{"true", true}, {"True", true}, {"FALSE", false},
		{"null", nil}, {"Null", nil}, {"NULL", nil}, {"~", nil},
		{"yes", "yes"}, {"on", "on"},
	}
	for _, tc := range cases {
		var got any
		if err := Unmarshal([]byte("v: "+tc.raw+"\n"), &got); err != nil {
			t.Fatalf("Unmarshal(%q): %v", tc.raw, err)
		}
		m, ok := got.(map[string]any)
		if !ok {
			t.Fatalf("Unmarshal(%q) shape = %T", tc.raw, got)
		}
		if !reflect.DeepEqual(m["v"], tc.want) {
			t.Fatalf("Unmarshal(%q) = %#v, want %#v", tc.raw, m["v"], tc.want)
		}
	}
}

// TestFlowSequenceBranches covers controlled flow-sequence edge cases.
func TestFlowSequenceBranches(t *testing.T) {
	t.Parallel()
	var empty struct {
		V []string `yaml:"v"`
	}
	if err := Unmarshal([]byte("v: []\n"), &empty); err != nil {
		t.Fatalf("empty flow sequence: %v", err)
	}
	if empty.V == nil || len(empty.V) != 0 {
		t.Fatalf("empty flow sequence = %#v", empty.V)
	}
	var quoted struct {
		V []string `yaml:"v"`
	}
	if err := Unmarshal([]byte(`v: ["a,b", 'c']`+"\n"), &quoted); err != nil {
		t.Fatalf("quoted flow items: %v", err)
	}
	if len(quoted.V) != 2 || quoted.V[0] != "a,b" || quoted.V[1] != "c" {
		t.Fatalf("quoted flow items = %#v", quoted.V)
	}
	for _, raw := range []string{
		"v: [unterminated\n",
		"v: [a,,b]\n",
		"v: [[a]]\n",
		"v: [a, 'unterminated]\n",
		"v: {a: b}\n",
	} {
		var out any
		if err := Unmarshal([]byte(raw), &out); err == nil {
			t.Fatalf("flow input %q was accepted", raw)
		}
	}
}

// TestLexerAndGrammarRejections covers strict errors for unsupported syntax.
func TestLexerAndGrammarRejections(t *testing.T) {
	t.Parallel()
	cases := []struct{ raw, want string }{
		{"\ta: 1\n", "tab in indentation"},
		{"\x00", "NUL byte"},
		{"---\na: 1\n", "multi-document"},
		{"? a\n", "complex keys"},
		{"a: &x 1\n", "anchors"},
		{"a: *x\n", "aliases"},
		{"a: !!str v\n", "custom tags"},
		{"a: >\n  v\n", "folded block scalars"},
		{"a: 1\na: 2\n", "duplicate key"},
		{"a\n", "expected 'key: value'"},
		{"<<: x\n", "merge keys"},
		{"a: 1\n  b: 2\n", "indentation"},
	}
	for _, tc := range cases {
		var out any
		err := Unmarshal([]byte(tc.raw), &out)
		if err == nil || !strings.Contains(err.Error(), tc.want) {
			t.Fatalf("Unmarshal(%q) err = %v, want containing %q", tc.raw, err, tc.want)
		}
	}
}

// TestBlockScalarChomping covers literal block scalar clip and strip modes.
func TestBlockScalarChomping(t *testing.T) {
	t.Parallel()
	cases := []struct{ raw, want string }{
		{"a: |\n  x\n\n", "x\n"},
		{"a: |-\n  x\n\n", "x"},
	}
	for _, tc := range cases {
		var got struct {
			A string `yaml:"a"`
		}
		if err := Unmarshal([]byte(tc.raw), &got); err != nil {
			t.Fatalf("Unmarshal(%q): %v", tc.raw, err)
		}
		if got.A != tc.want {
			t.Fatalf("Unmarshal(%q) = %q, want %q", tc.raw, got.A, tc.want)
		}
	}
}

// TestQuotedMappingKeys covers quoted-key decoding in splitKey.
func TestQuotedMappingKeys(t *testing.T) {
	t.Parallel()
	var got map[string]string
	raw := "\"a: b\": 1\n'c''d': 2\n"
	if err := Unmarshal([]byte(raw), &got); err != nil {
		t.Fatalf("Unmarshal quoted keys: %v", err)
	}
	if got["a: b"] != "1" || got["c'd"] != "2" {
		t.Fatalf("quoted keys = %#v", got)
	}
}

// TestNodeLookup covers the ordered-key lookup helper.
func TestNodeLookup(t *testing.T) {
	t.Parallel()
	n := &Node{
		Kind: Mapping,
		Keys: []string{"a", "b"},
		Vals: []*Node{scalarNode("1", sInt, 1), scalarNode("2", sInt, 2)},
	}
	if _, ok := n.lookup("missing"); ok {
		t.Fatal("lookup(missing) unexpectedly succeeded")
	}
	v, ok := n.lookup("b")
	if !ok || v.Value != "2" {
		t.Fatalf("lookup(b) = %v, %v", v, ok)
	}
}
