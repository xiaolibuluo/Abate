/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package yamlmini

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

// A textDuration is a real text codec used to verify encoding precedence.
type textDuration struct {
	Nanos int64
}

// MarshalText encodes the duration as a canonical "Ns" string.
func (d textDuration) MarshalText() ([]byte, error) {
	if d.Nanos < 0 {
		return nil, errors.New("negative duration")
	}
	return []byte(strconv.FormatInt(d.Nanos, 10) + "ns"), nil
}

// UnmarshalText decodes a canonical "Ns" string into the duration.
func (d *textDuration) UnmarshalText(b []byte) error {
	if !strings.HasSuffix(string(b), "ns") {
		return errors.New("missing ns suffix")
	}
	v, err := strconv.ParseInt(strings.TrimSuffix(string(b), "ns"), 10, 64)
	if err != nil {
		return err
	}
	d.Nanos = v
	return nil
}

// TestDecodeEntryErrors covers invalid decode targets and empty streams.
func TestDecodeEntryErrors(t *testing.T) {
	t.Parallel()
	var target testDoc
	if err := Unmarshal([]byte("a: 1\n"), nil); err == nil {
		t.Fatal("nil target accepted")
	}
	var nilPtr *testDoc
	if err := Unmarshal([]byte("a: 1\n"), nilPtr); err == nil {
		t.Fatal("nil pointer target accepted")
	}
	if err := Unmarshal([]byte("a: 1\n"), target); err == nil {
		t.Fatal("non-pointer target accepted")
	}
	if err := (*Decoder)(nil).Decode(&target); err == nil {
		t.Fatal("nil decoder accepted")
	}
	if err := NewDecoder(nil).Decode(&target); err == nil {
		t.Fatal("nil reader accepted")
	}
	for _, raw := range []string{"", "  \n# only a comment\n", "\r\n"} {
		if err := NewDecoder(strings.NewReader(raw)).Decode(&target); err == nil {
			t.Fatalf("blank document %q accepted", raw)
		}
	}
}

// TestMarshalErrorPaths covers encode failures for unsupported inputs.
func TestMarshalErrorPaths(t *testing.T) {
	t.Parallel()
	for _, in := range []any{map[int]string{1: "a"}, make(chan int)} {
		if _, err := Marshal(in); err == nil {
			t.Fatalf("Marshal(%T) unexpectedly succeeded", in)
		}
	}
}

// TestTextCodecPrecedence verifies TextMarshaler and TextUnmarshaler handling.
func TestTextCodecPrecedence(t *testing.T) {
	t.Parallel()
	b, err := Marshal(textDuration{Nanos: 250})
	if err != nil || string(b) != "250ns\n" {
		t.Fatalf("Marshal(textDuration) = %q, %v", b, err)
	}
	var out struct {
		V textDuration `yaml:"v"`
	}
	if err := Unmarshal([]byte("v: 250ns\n"), &out); err != nil || out.V.Nanos != 250 {
		t.Fatalf("Unmarshal(textDuration) = %+v, %v", out, err)
	}
	if _, err := Marshal(textDuration{Nanos: -1}); err == nil {
		t.Fatal("negative duration MarshalText error ignored")
	}
	var bad struct {
		V textDuration `yaml:"v"`
	}
	if err := Unmarshal([]byte("v: bogus\n"), &bad); err == nil {
		t.Fatal("UnmarshalText error ignored")
	}
}

// TestBindingTypeErrors covers shape and scalar type mismatch diagnostics.
func TestBindingTypeErrors(t *testing.T) {
	t.Parallel()
	cases := []struct {
		raw  string
		dest any
	}{
		{"v: abc\n", &struct {
			V int `yaml:"v"`
		}{}},
		{"v: 1\n", &struct {
			V bool `yaml:"v"`
		}{}},
		{"v: -1\n", &struct {
			V uint `yaml:"v"`
		}{}},
		{"v: abc\n", &struct {
			V float64 `yaml:"v"`
		}{}},
		{"v: {a: b}\n", &struct {
			V string `yaml:"v"`
		}{}},
		{"v: 1\n", &struct {
			V [2]int `yaml:"v"`
		}{}},
		{"v: [1, 2]\n", &struct {
			V [3]int `yaml:"v"`
		}{}},
		{"v: [a]\n", &struct {
			V map[string]string `yaml:"v"`
		}{}},
		{"v: {a: b}\n", &struct {
			V []string `yaml:"v"`
		}{}},
		{"v: [a]\n", &struct {
			V testDoc `yaml:"v"`
		}{}},
		{"v: {a: b}\n", &struct {
			V map[int]string `yaml:"v"`
		}{}},
		{"v: 1\n", &struct {
			V fmt.Stringer `yaml:"v"`
		}{}},
		{"v: 999999999999999999999\n", &struct {
			V int8 `yaml:"v"`
		}{}},
		{"v: 1e300\n", &struct {
			V float32 `yaml:"v"`
		}{}},
	}
	for _, tc := range cases {
		if err := Unmarshal([]byte(tc.raw), tc.dest); err == nil {
			t.Fatalf("Unmarshal(%q) into %T unexpectedly succeeded", tc.raw, tc.dest)
		}
	}
}

// TestOmitEmptyZeroValues covers zero detection across reflect kinds.
func TestOmitEmptyZeroValues(t *testing.T) {
	t.Parallel()
	type inner struct {
		X int `yaml:"x"`
	}
	type doc struct {
		P *int           `yaml:"p,omitempty"`
		M map[string]int `yaml:"m,omitempty"`
		S []int          `yaml:"s,omitempty"`
		I any            `yaml:"i,omitempty"`
		A [2]int         `yaml:"a,omitempty"`
		B [2]string      `yaml:"b,omitempty"`
		N inner          `yaml:"n,omitempty"`
		V int            `yaml:"v,omitempty"`
	}
	if !isZero(reflect.Value{}) {
		t.Fatal("invalid value must be zero")
	}
	if b, err := Marshal(doc{}); err != nil || len(b) != 0 {
		t.Fatalf("Marshal all-zero doc = %q, %v", b, err)
	}
	b, err := Marshal(doc{A: [2]int{1, 0}, N: inner{X: 1}})
	if err != nil || !bytes.Contains(b, []byte("a:")) || !bytes.Contains(b, []byte("n:")) {
		t.Fatalf("Marshal non-zero doc = %q, %v", b, err)
	}
}

// TestScalarEmitHelpers covers canonical quoting and control-character escaping.
func TestScalarEmitHelpers(t *testing.T) {
	t.Parallel()
	if !plainSafe("abc") || plainSafe("") || plainSafe(" a") || plainSafe("a: b") ||
		plainSafe("a #b") || plainSafe("a:") || plainSafe("- x") {
		t.Fatal("plainSafe misclassified a scalar")
	}
	if !looksNonString("42") || !looksNonString("1.5") || !looksNonString("true") ||
		looksNonString("hello") {
		t.Fatal("looksNonString misclassified a scalar")
	}
	if got := quoteString("a\x01b\x7f"); got != `"a\u0001b\u007F"` {
		t.Fatalf("quoteString = %q", got)
	}
	if got := emitKey("a: b"); got != `"a: b"` {
		t.Fatalf("emitKey = %q", got)
	}
	if got := emitScalar(scalarNode("42", sString, 0)); got != `"42"` {
		t.Fatalf("emitScalar string = %q", got)
	}
	if got := emitScalar(nullNode(0)); got != "null" {
		t.Fatalf("emitScalar null = %q", got)
	}
	if got := emitScalar(emptySequence(0)); got != "[]" || emitScalar(emptyMapping(0)) != "{}" {
		t.Fatal("emitScalar empty collection mismatch")
	}
}

// TestMarshalMapOrderingAndNesting verifies sorted keys and nested collections.
func TestMarshalMapOrderingAndNesting(t *testing.T) {
	t.Parallel()
	b, err := Marshal(map[string]any{
		"z": 1,
		"a": map[string]any{"y": []int{1, 2}, "x": true},
	})
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	want := "a:\n  x: true\n  y:\n    - 1\n    - 2\nz: 1\n"
	if string(b) != want {
		t.Fatalf("Marshal = %q, want %q", b, want)
	}
}

// TestKindString covers every Kind label and the numeric fallback.
func TestKindString(t *testing.T) {
	t.Parallel()
	cases := []struct {
		k    Kind
		want string
	}{
		{Scalar, "scalar"}, {Mapping, "mapping"}, {Sequence, "sequence"},
		{Null, "null"}, {Kind(99), "kind(99)"},
	}
	for _, tc := range cases {
		if got := tc.k.String(); got != tc.want {
			t.Fatalf("Kind(%d).String() = %q, want %q", tc.k, got, tc.want)
		}
	}
}
