/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package frontkit

import (
	"reflect"
	"strings"
	"testing"

	"abate/internal/determinism"
)

// assertFrozenOrder checks the length-descending then lexicographic sort key.
func assertFrozenOrder(t *testing.T, table []string) {
	t.Helper()
	for i := 1; i < len(table); i++ {
		prev, cur := table[i-1], table[i]
		if len(prev) < len(cur) {
			t.Fatalf("length key violated at %d: %q before %q", i, prev, cur)
		}
		if len(prev) == len(cur) && strings.Compare(prev, cur) > 0 {
			t.Fatalf("lexicographic key violated at %d: %q before %q", i, prev, cur)
		}
	}
}

func TestChaosNormalizeTablesUnordered(t *testing.T) {
	t.Parallel()
	keywords := []string{"a", "bbb", "cc", "dd", "e", "fracture", "fract", "frac"}
	operators := []string{"==", "!=", "===", ">", ">>="}
	reversedKeywords := append([]string(nil), keywords...)
	for left, right := 0, len(reversedKeywords)-1; left < right; left, right = left+1, right-1 {
		reversedKeywords[left], reversedKeywords[right] =
			reversedKeywords[right], reversedKeywords[left]
	}
	gotKeywords, gotOperators := NormalizeTables(keywords, operators)
	againKeywords, againOperators := NormalizeTables(reversedKeywords, operators)
	if !reflect.DeepEqual(gotKeywords, againKeywords) {
		t.Fatalf("output depends on input slice order: %v vs %v",
			gotKeywords, againKeywords)
	}
	if !reflect.DeepEqual(gotOperators, againOperators) {
		t.Fatalf("operator output differs: %v vs %v", gotOperators, againOperators)
	}
	assertFrozenOrder(t, gotKeywords)
	assertFrozenOrder(t, gotOperators)
}

func TestChaosNormalizeTablesDuplicatesAndLong(t *testing.T) {
	t.Parallel()
	duplicates := []string{"x", "x", "x", "yy", "yy", "z"}
	got, _ := NormalizeTables(duplicates, nil)
	want := []string{"yy", "yy", "x", "x", "x", "z"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("duplicate table = %v, want %v", got, want)
	}
	assertFrozenOrder(t, got)

	long := strings.Repeat("z", 4096)
	mixed := []string{"b", long, "aa", "ccc", long}
	gotLong, _ := NormalizeTables(mixed, nil)
	if len(gotLong) != len(mixed) {
		t.Fatalf("long input lost entries: %d != %d", len(gotLong), len(mixed))
	}
	if gotLong[0] != long || gotLong[1] != long {
		t.Fatal("longest entries must lead the normalized table")
	}
	assertFrozenOrder(t, gotLong)
}

func TestChaosNormalizeTablesNilEmptyAndNoMutation(t *testing.T) {
	t.Parallel()
	gotKeywords, gotOperators := NormalizeTables(nil, nil)
	if len(gotKeywords) != 0 || len(gotOperators) != 0 {
		t.Fatalf("nil input = (%v, %v), want two empty slices", gotKeywords, gotOperators)
	}
	gotKeywords, gotOperators = NormalizeTables([]string{}, []string{})
	if len(gotKeywords) != 0 || len(gotOperators) != 0 {
		t.Fatalf("empty input = (%v, %v), want two empty slices", gotKeywords, gotOperators)
	}
	input := []string{"b", "aaa", "cc"}
	got, _ := NormalizeTables(input, nil)
	if !reflect.DeepEqual(input, []string{"b", "aaa", "cc"}) {
		t.Fatalf("NormalizeTables mutated its input: %v", input)
	}
	if reflect.DeepEqual(got, input) && len(got) > 1 {
		t.Fatal("NormalizeTables returned the input slice instead of a fresh copy")
	}
}

func TestChaosCanonicalTagMalformed(t *testing.T) {
	t.Parallel()
	unknown := []string{
		strings.Repeat("p", 4096),
		"\x00\x01\x02",
		string([]byte{0xff, 0xfe, 0x80}),
		"py\x00",
		"\t\r\n  \x0b\x0c",
		"!!!###$$$",
		"python c++ c#",
	}
	for _, in := range unknown {
		got, ok := CanonicalTag(in)
		if ok || got != "" {
			t.Fatalf("CanonicalTag(%q) = (%q, %v), want a clean rejection", in[:min(len(in), 24)], got, ok)
		}
	}
	padded := strings.Repeat(" ", 2048) + "PyThOn" + strings.Repeat("\t", 2048)
	if got, ok := CanonicalTag(padded); !ok || got != "py" {
		t.Fatalf("padded alias = (%q, %v), want (\"py\", true)", got, ok)
	}
}

func TestChaosTagForPathAndExtensions(t *testing.T) {
	t.Parallel()
	paths := []string{
		strings.Repeat("a", 4096) + ".rb",
		"\x00",
		"a/b/c/\xff",
		strings.Repeat(".", 2048),
		"/",
	}
	for _, p := range paths {
		if got, ok := TagForPath(p); ok || got != "" {
			t.Fatalf("TagForPath(%q) = (%q, %v), want a clean rejection",
				p[:min(len(p), 24)], got, ok)
		}
	}
	for _, tag := range []string{"", "\x00", strings.Repeat("k", 4096), "not-a-language"} {
		if got := ExtensionsFor(tag); len(got) != 0 {
			t.Fatalf("ExtensionsFor(%q) = %v, want an empty slice",
				tag[:min(len(tag), 24)], got)
		}
	}
}

// chaosSnapshot folds chaotic inputs into one canonical byte string so the
// 100-run assertion proves every chaotic path is reproducible.
func chaosSnapshot() []byte {
	keywords := []string{"z", "fracture", "a", "cc", "zz", "fracture"}
	operators := []string{">", "===", "==", "!=", ">>="}
	kw, op := NormalizeTables(keywords, operators)
	parts := make([]string, 0, 64)
	parts = append(parts, strings.Join(kw, ","), strings.Join(op, ","))
	for _, in := range []string{"  C++ ", "GOLANG", "kotlin", "ruby", "\x00", ""} {
		tag, ok := CanonicalTag(in)
		okStr := "0"
		if ok {
			okStr = "1"
		}
		parts = append(parts, in+"="+tag+"="+okStr)
	}
	for _, p := range []string{"src/main.tsx", "src/x.hpp", "Makefile", "a.pyi", "a.rb"} {
		tag, ok := TagForPath(p)
		okStr := "0"
		if ok {
			okStr = "1"
		}
		parts = append(parts, p+"="+tag+"="+okStr)
	}
	return []byte(strings.Join(parts, "|"))
}

func TestChaosDeterminism100(t *testing.T) {
	t.Parallel()
	determinism.MustMatch100(t, 42, chaosSnapshot)
}
