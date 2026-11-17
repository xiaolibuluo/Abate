/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package gofront

import (
	"crypto/sha256"
	"slices"
	"testing"

	"abate/internal/determinism"
	"abate/lang/air"
)

// TestDeterminismHashState100 requires 100 identical HashState(Unit) digests.
func TestDeterminismHashState100(t *testing.T) {
	t.Parallel()
	unit := mustParse(t, sampleGo)
	digests := make(map[[32]byte]struct{})
	determinism.MustMatch100(t, 0xAB47E001, func() []byte {
		sum := determinism.HashState(unit)
		digests[sum] = struct{}{}
		return sum[:]
	})
	if len(digests) != 1 {
		t.Fatalf("HashState produced %d distinct digests", len(digests))
	}
}

// TestDeterminismFingerprint100 requires 100 identical AIR fingerprints.
func TestDeterminismFingerprint100(t *testing.T) {
	t.Parallel()
	frontend := NewGoFrontend().WithTypeCheck(false)
	digests := make(map[[32]byte]struct{})
	for i := 0; i < 100; i++ {
		unit, diagnostics, err := frontend.Parse(chaosSourceUnit())
		if err != nil || len(diagnostics) != 0 {
			t.Fatalf("run %d: %v %v", i, err, diagnostics)
		}
		archive, err := air.NewEncoder().Encode(unit)
		if err != nil {
			t.Fatalf("run %d encode: %v", i, err)
		}
		digests[sha256.Sum256(archive)] = struct{}{}
	}
	if len(digests) != 1 {
		t.Fatalf("expected exactly one unique SHA-256, got %d", len(digests))
	}
}

// TestDeterminismFileOrderIndependence reorders the same source buffers.
func TestDeterminismFileOrderIndependence(t *testing.T) {
	t.Parallel()
	frontend := NewGoFrontend().WithTypeCheck(false)
	files := []air.SourceFile{
		air.NewSourceFile("b.go", []byte("package p\n\nvar B = 2\n")),
		air.NewSourceFile("a.go", []byte("package p\n\nvar A = 1\n")),
		air.NewSourceFile("c.go", []byte("package p\n\nvar C = 3\n")),
	}
	first, _, err := frontend.Parse(air.SourceUnit{Language: "go", Files: files})
	if err != nil {
		t.Fatal(err)
	}
	reversed := slices.Clone(files)
	slices.Reverse(reversed)
	second, _, err := frontend.Parse(air.SourceUnit{Language: "go", Files: reversed})
	if err != nil {
		t.Fatal(err)
	}
	if first.Fingerprint() != second.Fingerprint() {
		t.Fatal("file order leaked into the AIR fingerprint")
	}
	if err := air.CompareUnits(first, second); err != nil {
		t.Fatalf("reordered units differ: %v", err)
	}
}

// TestDeterminismAcrossFrontends proves the type-check option only adds
// annotations and never changes the syntactic tree shape.
func TestDeterminismAcrossFrontends(t *testing.T) {
	t.Parallel()
	plain, _, err := NewGoFrontend().WithTypeCheck(false).Parse(chaosSourceUnit())
	if err != nil {
		t.Fatal(err)
	}
	enriched, _, err := NewGoFrontend().Parse(chaosSourceUnit())
	if err != nil {
		t.Fatal(err)
	}
	if plain.NodeCount() != enriched.NodeCount() {
		t.Fatal("semantic annotation changed the node count")
	}
	plainNodes := air.Collect(plain)
	enrichedNodes := air.Collect(enriched)
	for i := range plainNodes {
		if plainNodes[i].Kind() != enrichedNodes[i].Kind() ||
			plainNodes[i].Pos().Canonical() != enrichedNodes[i].Pos().Canonical() ||
			air.AttributeOf(plainNodes[i]) != air.AttributeOf(enrichedNodes[i]) {
			t.Fatalf("node %d changed shape under type checking", i)
		}
	}
}
