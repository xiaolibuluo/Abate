/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package chaos

import (
	"crypto/sha256"
	"encoding/hex"
	"reflect"
	"testing"
)

// TestTruncateHalfReproducesG1 proves the two G1 minimal reproductions are
// exactly the half truncations of complete synthetic sources.
func TestTruncateHalfReproducesG1(t *testing.T) {
	t.Parallel()
	repros := G1Reproductions()
	if len(repros) != 2 {
		t.Fatalf("expected two reproductions, got %d", len(repros))
	}
	if string(repros[0]) != string(g1Minimal[0]) ||
		string(repros[1]) != string(g1Minimal[1]) {
		t.Fatalf("G1 reproductions wrong: %q / %q", repros[0], repros[1])
	}
	for i, minimal := range g1Minimal {
		full := g1FullSource(minimal)
		if len(full) != 2*len(minimal) ||
			!reflect.DeepEqual(TruncateHalf(full), minimal) {
			t.Fatalf("reproduction %d is not a half truncation", i)
		}
	}
	if len(TruncateHalf([]byte("abc"))) != 1 || len(TruncateHalf(nil)) != 0 {
		t.Fatal("odd-length or empty truncation wrong")
	}
}

// TestMergeDeduplicatesBySHA256 proves identical inputs collapse and the
// survivors are ordered by ascending digest, never by map iteration.
func TestMergeDeduplicatesBySHA256(t *testing.T) {
	t.Parallel()
	inputs := [][]byte{
		[]byte("alpha"), []byte("alpha"), []byte("beta"), []byte("alpha"),
	}
	merged := Merge(inputs)
	if len(merged) != 2 {
		t.Fatalf("expected two survivors, got %d", len(merged))
	}
	previous := ""
	for _, item := range merged {
		sum := sha256.Sum256(item)
		digest := hex.EncodeToString(sum[:])
		if previous != "" && previous >= digest {
			t.Fatal("merge output is not in ascending digest order")
		}
		previous = digest
	}
	if len(Merge(nil)) != 0 {
		t.Fatal("empty input set must merge to an empty set")
	}
}

// TestSpliceAndDropToken proves byte insertion and token removal are bounded,
// deterministic and produce independent buffers.
func TestSpliceAndDropToken(t *testing.T) {
	t.Parallel()
	content := []byte("abcdef")
	if string(Splice(content, 3, []byte("XY"))) != "abcXYdef" {
		t.Fatal("middle splice wrong")
	}
	if string(Splice(content, -5, []byte("Z"))) != "Zabcdef" {
		t.Fatal("negative offset was not clamped")
	}
	if string(Splice(content, 99, []byte("Z"))) != "abcdefZ" {
		t.Fatal("oversized offset was not clamped")
	}
	if string(DropToken(content, []byte("cd"))) != "abef" {
		t.Fatal("token drop wrong")
	}
	if string(DropToken(content, nil)) != "abcdef" {
		t.Fatal("empty token must be a no-op")
	}
	if string(content) != "abcdef" {
		t.Fatal("damage mutated the source buffer")
	}
}

// TestGeneratorIsDeterministic proves the same seed and manifest digest always
// produce the same merged damage set.
func TestGeneratorIsDeterministic(t *testing.T) {
	t.Parallel()
	content := []byte("package sample\n\nfunc Add(a int, b int) int { return a + b }\n")
	generator := NewGenerator(20261114, "manifest-digest", 8)
	first := generator.Generate(content, "go:add")
	second := generator.Generate(content, "go:add")
	if len(first) == 0 || len(first) > 9 || !reflect.DeepEqual(first, second) {
		t.Fatalf("generator is not deterministic: %d variants", len(first))
	}
	truncated := string(TruncateHalf(content))
	found := false
	for _, item := range first {
		if string(item) == truncated {
			found = true
		}
	}
	if !found {
		t.Fatal("the truncated body is not part of the merged set")
	}
	zero := NewGenerator(20261114, "manifest-digest", 0)
	if len(zero.Generate(content, "go:add")) != 0 {
		t.Fatal("a zero budget must produce no variants")
	}
}

// TestDamageKindValid proves the four declared shapes and rejects unknown ones.
func TestDamageKindValid(t *testing.T) {
	t.Parallel()
	for _, kind := range []DamageKind{DamageTruncate, DamageSplice,
		DamageDrop, DamageMerge} {
		if !kind.Valid() {
			t.Fatalf("declared kind %q was rejected", kind)
		}
	}
	if DamageKind("unknown").Valid() {
		t.Fatal("unknown damage kind was accepted")
	}
}
