/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package frontkit

import (
	"reflect"
	"slices"
	"strings"
	"testing"
)

// lexicalOnly is a mutated NormalizeTables that drops the length-descending
// key and keeps lexicographic order alone.
func lexicalOnly(in []string) []string {
	out := append([]string(nil), in...)
	slices.SortStableFunc(out, func(a, b string) int {
		return strings.Compare(a, b)
	})
	return out
}

// ascendingLength is a mutated NormalizeTables that reverses the length key.
func ascendingLength(in []string) []string {
	out := append([]string(nil), in...)
	slices.SortStableFunc(out, func(a, b string) int {
		if len(a) != len(b) {
			if len(a) < len(b) {
				return -1
			}
			return 1
		}
		return strings.Compare(a, b)
	})
	return out
}

func TestMutationNormalizeTablesSortKey(t *testing.T) {
	t.Parallel()
	input := []string{"a", "bbb", "cc", "dd"}
	got, _ := NormalizeTables(input, nil)
	want := []string{"bbb", "cc", "dd", "a"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("NormalizeTables(%v) = %v, want %v", input, got, want)
	}
	if reflect.DeepEqual(got, lexicalOnly(input)) {
		t.Fatal("mutation not detected: dropping the length key changes nothing")
	}
	if reflect.DeepEqual(got, ascendingLength(input)) {
		t.Fatal("mutation not detected: reversing the length key changes nothing")
	}
	if !reflect.DeepEqual(lexicalOnly(input), []string{"a", "bbb", "cc", "dd"}) {
		t.Fatal("the lexical-only mutant is not the intended counter-example")
	}
	if !reflect.DeepEqual(ascendingLength(input), []string{"a", "cc", "dd", "bbb"}) {
		t.Fatal("the ascending-length mutant is not the intended counter-example")
	}
}

func TestMutationAliasTableEntry(t *testing.T) {
	t.Parallel()
	mutated := make([]languageAlias, len(languageAliases))
	copy(mutated, languageAliases)
	changed := 0
	for i := range mutated {
		if mutated[i].alias == "python" {
			mutated[i].canonical = "go"
			changed++
		}
	}
	if changed != 1 {
		t.Fatalf("expected exactly one python alias entry, found %d", changed)
	}
	mutatedLookup := func(s string) (string, bool) {
		for i := range mutated {
			if mutated[i].alias == s {
				return mutated[i].canonical, true
			}
		}
		return "", false
	}
	got, ok := CanonicalTag("python")
	if !ok || got != "py" {
		t.Fatalf("CanonicalTag(\"python\") = (%q, %v), want (\"py\", true)", got, ok)
	}
	mutatedTag, mutatedOK := mutatedLookup("python")
	if !mutatedOK || mutatedTag == got {
		t.Fatal("mutation not detected: a changed alias entry must change the output")
	}
	if mutatedTag != "go" {
		t.Fatalf("mutated alias resolves to %q, want \"go\"", mutatedTag)
	}
}

func TestMutationCanonicalTagSet(t *testing.T) {
	t.Parallel()
	if canonicalTags[1] != "py" {
		t.Fatalf("canonicalTags[1] = %q, want \"py\"", canonicalTags[1])
	}
	mutatedTags := append([]string(nil), canonicalTags[:]...)
	mutatedTags[1] = "ruby"
	if reflect.DeepEqual(mutatedTags, canonicalTags[:]) {
		t.Fatal("the canonical tag mutant is identical to the production table")
	}
	if got, ok := CanonicalTag("py"); !ok || got != "py" {
		t.Fatalf("CanonicalTag(\"py\") = (%q, %v), want (\"py\", true)", got, ok)
	}
	if got, ok := CanonicalTag("ruby"); ok || got != "" {
		t.Fatalf("CanonicalTag(\"ruby\") = (%q, %v), want a clean rejection", got, ok)
	}
}

func TestMutationExtensionsTableEntry(t *testing.T) {
	t.Parallel()
	mutated := make([]languageExtension, len(languageExtensions))
	copy(mutated, languageExtensions)
	changed := 0
	for i := range mutated {
		if mutated[i].extension == ".pyi" {
			mutated[i].tag = "js"
			changed++
		}
	}
	if changed != 1 {
		t.Fatalf("expected exactly one .pyi entry, found %d", changed)
	}
	mutatedLookup := func(ext string) string {
		for i := range mutated {
			if mutated[i].extension == ext {
				return mutated[i].tag
			}
		}
		return ""
	}
	got, ok := TagForPath("src/stub.pyi")
	if !ok || got != "py" {
		t.Fatalf("TagForPath(\"src/stub.pyi\") = (%q, %v), want (\"py\", true)", got, ok)
	}
	if mutatedLookup(".pyi") == got {
		t.Fatal("mutation not detected: a changed extension entry must change the output")
	}
}

// mutatedKindString reads the supplied name table instead of the frozen one.
func mutatedKindString(k TokenKind, names []string) string {
	if int(k) < len(names) {
		return names[k]
	}
	return "TokenKind(out-of-range)"
}

func TestMutationTokenKindNameTable(t *testing.T) {
	t.Parallel()
	mutated := make([]string, len(tokenKindNames))
	copy(mutated, tokenKindNames[:])
	mutated[TokenIdent] = "Identifier"
	if TokenIdent.String() != "Ident" {
		t.Fatalf("TokenIdent.String() = %q, want \"Ident\"", TokenIdent.String())
	}
	if TokenIdent.String() == mutated[TokenIdent] {
		t.Fatal("mutation not detected: String() must be driven by the name table")
	}
	if mutatedKindString(TokenIdent, mutated) == TokenIdent.String() {
		t.Fatal("mutation not detected: the mutated name table did not change the output")
	}
	for _, tc := range kindCases {
		if tc.kind.String() != tokenKindNames[tc.kind] {
			t.Fatalf("TokenKind(%d) does not read its frozen name", uint16(tc.kind))
		}
	}
}
