/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package record

import (
	"bytes"
	"context"
	"crypto/sha256"
	"errors"
	"sync"
	"testing"

	"abate/sim/clock"
)

// staticSource is a read-only source used to exercise recorder subscriptions.
type staticSource struct {
	frames []Frame
}

// Collect returns the source frames in their declared deterministic order.
func (s *staticSource) Collect() []Frame { return s.frames }

// recordOnce records a fixed scenario through both Record and a subscribed
// source, then returns the sealed archive bytes.
func recordOnce() []byte {
	rec := NewRecorder(testFingerprint())
	rec.Start(context.Background())
	rec.Subscribe(&staticSource{frames: []Frame{
		{At: 9, Kind: KindClk, Payload: []byte("9 RUNNING")},
	}})
	for _, f := range sampleFrames() {
		rec.Record(f)
	}
	arch, err := rec.Stop()
	if err != nil {
		panic(err)
	}
	var buf bytes.Buffer
	if err := arch.Save(&buf); err != nil {
		panic(err)
	}
	return buf.Bytes()
}

// TestChaosArchiveUniqueSHA256 certifies one hundred same-seed recordings.
func TestChaosArchiveUniqueSHA256(t *testing.T) {
	t.Parallel()
	seen := make(map[[32]byte]int)
	var first [32]byte
	for i := 0; i < 100; i++ {
		h := sha256.Sum256(recordOnce())
		seen[h]++
		if i == 0 {
			first = h
			continue
		}
		if h != first {
			t.Fatalf("archive diverged on run %d", i)
		}
	}
	if len(seen) != 1 {
		t.Fatalf("expected exactly one unique SHA-256, got %d", len(seen))
	}
	t.Logf("REC unique SHA-256 over 100 runs = %d (hash=%x)", len(seen), first)
}

// TestChaosConcurrentRecording proves independent recorders agree and that one
// shared recorder stays race free under concurrent Record calls.
func TestChaosConcurrentRecording(t *testing.T) {
	t.Parallel()
	const runners = 32
	digests := make([][32]byte, runners)
	var wg sync.WaitGroup
	for i := 0; i < runners; i++ {
		wg.Add(1)
		go func(k int) {
			defer wg.Done()
			digests[k] = sha256.Sum256(recordOnce())
		}(i)
	}
	wg.Wait()
	for i := range digests {
		if digests[i] != digests[0] {
			t.Fatalf("concurrent recorder %d diverged", i)
		}
	}

	shared := NewRecorder(testFingerprint())
	const writers, perWriter = 16, 8
	for w := 0; w < writers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < perWriter; j++ {
				shared.Record(Frame{
					At:      clock.Instant(1),
					Kind:    KindSched,
					Payload: []byte("x"),
				})
			}
		}()
	}
	wg.Wait()
	arch, err := shared.Stop()
	if err != nil {
		t.Fatalf("stop shared recorder: %v", err)
	}
	if arch.Len() != writers*perWriter {
		t.Fatalf("shared recorder lost frames: %d", arch.Len())
	}
	if err := (Replayer{}).Verify(arch); err != nil {
		t.Fatalf("shared recorder produced an invalid archive: %v", err)
	}
}

// TestChaosHashChainTamperBreaksSuffix proves a middle-frame edit invalidates
// every frame after it and that Load rejects the rewritten archive.
func TestChaosHashChainTamperBreaksSuffix(t *testing.T) {
	t.Parallel()
	frames := make([]Frame, 0, 8)
	for i := 0; i < 8; i++ {
		frames = append(frames, Frame{
			At:      clock.Instant(i + 1),
			Kind:    KindSched,
			Payload: []byte{byte(i), byte(i * 3), 0x7E},
		})
	}
	good := makeArchive(t, frames)
	arch := makeArchive(t, frames)
	const k = 3
	arch.entries[k].frame.Payload = append([]byte(nil), arch.entries[k].frame.Payload...)
	arch.entries[k].frame.Payload[1] ^= 0xFF
	// Without recomputing the chain, the edit breaks verification and every
	// downstream consumer of the archive refuses it.
	if err := arch.Verify(); !errors.Is(err, ErrCorruptChain) {
		t.Fatalf("tampered chain was accepted: %v", err)
	}
	if err := (Replayer{}).Replay(arch, &frameCollector{}); !errors.Is(err, ErrCorruptChain) {
		t.Fatalf("Replay accepted a broken chain: %v", err)
	}
	if _, err := Compare(good, arch); !errors.Is(err, ErrCorruptChain) {
		t.Fatalf("Compare accepted a broken right archive: %v", err)
	}
	if _, err := Compare(arch, good); !errors.Is(err, ErrCorruptChain) {
		t.Fatalf("Compare accepted a broken left archive: %v", err)
	}
	if _, err := Load(bytes.NewReader(archiveBytes(t, arch))); !errors.Is(err, ErrInvalidFormat) {
		t.Fatalf("Load accepted a broken chain: %v", err)
	}
	// Recomputing the whole suffix makes the archive internally consistent, but
	// every hash from the edited frame onwards now differs from the original.
	reseled := makeArchive(t, frames)
	reseled.entries[k].frame.Payload = append([]byte(nil), reseled.entries[k].frame.Payload...)
	reseled.entries[k].frame.Payload[1] ^= 0xFF
	reseled.entries[k].hash = chainHash(reseled.entries[k].prev, reseled.entries[k].frame)
	for i := k + 1; i < reseled.Len(); i++ {
		reseled.entries[i].prev = reseled.entries[i-1].hash
		reseled.entries[i].hash = chainHash(reseled.entries[i].prev, reseled.entries[i].frame)
	}
	if err := (Replayer{}).Verify(reseled); err != nil {
		t.Fatalf("reseled chain failed verification: %v", err)
	}
	for i := k; i < reseled.Len(); i++ {
		if reseled.Hash(i) == good.Hash(i) {
			t.Fatalf("frame %d hash did not change after a middle-frame edit", i)
		}
	}
	diff, err := Compare(good, reseled)
	if err != nil || diff.Equal || diff.Summary.FieldsDifferent < reseled.Len()-k {
		t.Fatalf("reseled chain comparison = equal=%v different=%d err=%v",
			diff.Equal, diff.Summary.FieldsDifferent, err)
	}
}
