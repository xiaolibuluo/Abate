/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package air

import (
	"crypto/sha256"
	"errors"
	"sync"
	"testing"

	"abate/internal/determinism"
)

// TestChaosEncodeDeterminism100 requires one unique SHA-256 across 100 runs.
func TestChaosEncodeDeterminism100(t *testing.T) {
	t.Parallel()
	u := mustUnit(t)
	fn := func() []byte {
		out, err := NewEncoder().Encode(u)
		if err != nil {
			t.Fatalf("Encode under determinism run: %v", err)
		}
		return out
	}
	determinism.MustMatch100(t, 0xA17C0DE5, fn)
	unique := make(map[[32]byte]struct{})
	for i := 0; i < 100; i++ {
		unique[sha256.Sum256(fn())] = struct{}{}
	}
	if len(unique) != 1 {
		t.Fatalf("expected exactly one unique archive hash, got %d", len(unique))
	}
	if digestOfUnit(u) != u.Fingerprint() {
		t.Fatal("unit fingerprint helper drifted")
	}
}

// TestChaosHashStateStable checks the canonical JSON projection of a unit.
func TestChaosHashStateStable(t *testing.T) {
	t.Parallel()
	u := mustUnit(t)
	determinism.MustMatch100(t, 0x5EED1234, func() []byte {
		sum := determinism.HashState(u)
		return sum[:]
	})
}

// TestChaosConcurrentEncodeDecode runs the codec under the race detector.
func TestChaosConcurrentEncodeDecode(t *testing.T) {
	t.Parallel()
	u := mustUnit(t)
	archive := mustEncode(t, NewEncoder(), u)
	want := u.Fingerprint()
	failures := make(chan error, 64)
	var wg sync.WaitGroup
	for worker := 0; worker < 16; worker++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 20; i++ {
				out, err := NewEncoder().Encode(u)
				if err != nil {
					failures <- err
					return
				}
				back, err := NewDecoder().Decode(out)
				if err != nil {
					failures <- err
					return
				}
				if back.Fingerprint() != want {
					failures <- ErrUnitMismatch
					return
				}
				if _, err := NewDecoder().Decode(archive); err != nil {
					failures <- err
					return
				}
			}
		}()
	}
	wg.Wait()
	close(failures)
	for err := range failures {
		t.Fatalf("concurrent codec: %v", err)
	}
}

// TestChaosLargeAndEmptyUnits stresses wide, deep, over-deep, and empty units.
func TestChaosLargeAndEmptyUnits(t *testing.T) {
	t.Parallel()
	wide := NewUnit(NewPackageRef("wide", nil))
	children := make([]Node, 0, 4000)
	for i := 0; i < 4000; i++ {
		children = append(children,
			NewNode(KindIdent, NewPosition(1, i+1, i, i+1), "x"))
	}
	root := NewNode(KindBlock, NewPosition(1, 1, 0, 4000), "", children...)
	if err := wide.AddFile(NewFileRef("wide.go", make([]byte, 4000)), root); err != nil {
		t.Fatalf("wide AddFile: %v", err)
	}
	back := mustDecode(t, NewDecoder(), mustEncode(t, NewEncoder(), wide))
	if back.NodeCount() != 4001 {
		t.Fatalf("wide unit lost nodes: %d", back.NodeCount())
	}

	chain := NewNode(KindIdent, NewPosition(1, 1, 0, 1), "leaf")
	for i := 0; i < 400; i++ {
		chain = NewNode(KindBlock, NewPosition(1, 1, 0, i+2), "", chain)
	}
	deep := NewUnit(NewPackageRef("deep", nil))
	if err := deep.AddFile(NewFileRef("deep.go", make([]byte, 402)), chain); err != nil {
		t.Fatalf("deep AddFile: %v", err)
	}
	if got := mustDecode(t, NewDecoder(), mustEncode(t, NewEncoder(), deep)).NodeCount(); got != 401 {
		t.Fatalf("deep unit lost nodes: %d", got)
	}

	for i := 0; i < 200; i++ {
		chain = NewNode(KindBlock, NewPosition(1, 1, 0, i+403), "", chain)
	}
	tooDeep := NewUnit(NewPackageRef("tooDeep", nil))
	if err := tooDeep.AddFile(NewFileRef("deep.go", make([]byte, 700)), chain); err != nil {
		t.Fatalf("over-deep AddFile: %v", err)
	}
	if _, err := NewDecoder().Decode(mustEncode(t, NewEncoder(), tooDeep)); !errors.Is(err, ErrTreeTooDeep) {
		t.Fatalf("over-deep tree accepted: %v", err)
	}

	empty := NewUnit(NewPackageRef("empty", nil))
	if got := mustDecode(t, NewDecoder(), mustEncode(t, NewEncoder(), empty)).NodeCount(); got != 0 {
		t.Fatalf("empty unit grew %d nodes", got)
	}
}

// TestChaosDecodeNeverPanics feeds deterministic random byte strings to Load.
func TestChaosDecodeNeverPanics(t *testing.T) {
	t.Parallel()
	state := uint64(0x9E3779B97F4A7C15)
	next := func() uint64 {
		state = state*6364136223846793005 + 1442695040888963407
		return state
	}
	for i := 0; i < 2000; i++ {
		size := int(next() % 200) //nolint:gosec // the modulo bounds the value, so the conversion is safe.
		buf := make([]byte, size)
		for j := range buf {
			buf[j] = byte(next() >> 33)
		}
		unit, err := NewDecoder().Load(buf)
		if err == nil && unit == nil {
			t.Fatalf("iteration %d returned no unit and no error", i)
		}
		if err != nil && !isAIRSentinel(err) {
			t.Fatalf("iteration %d: unclassified error %v", i, err)
		}
	}
}
