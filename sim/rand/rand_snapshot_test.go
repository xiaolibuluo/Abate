/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package rand

import (
	"testing"

	"abate/internal/determinism"
)

// TestSnapshotDeterminism asserts that two independently seeded streams drawing
// the same sequence produce byte-identical DrawRecord snapshots.
func TestSnapshotDeterminism(t *testing.T) {
	t.Parallel()
	a := New(2026)
	b := New(2026)
	for i := 0; i < 50; i++ {
		a.Uint64()
		a.IntN(17)
		a.Bool()
		a.Bytes(3)
		b.Uint64()
		b.IntN(17)
		b.Bool()
		b.Bytes(3)
	}
	sa := determinism.HashState(a.Snapshot())
	sb := determinism.HashState(b.Snapshot())
	if sa != sb {
		t.Fatalf("snapshot non-deterministic across equal seeds")
	}
}

// TestSnapshotDefensiveCopy confirms Snapshot returns a copy, not a live view.
func TestSnapshotDefensiveCopy(t *testing.T) {
	t.Parallel()
	s := New(7)
	for i := 0; i < 10; i++ {
		s.Uint64()
	}
	snap := s.Snapshot()
	snap[0].Result = 0
	if s.Snapshot()[0].Result == 0 {
		t.Fatalf("snapshot is not a defensive copy")
	}
}

// TestSnapshotRestoreConsistency checks that a stream restored to draw ordinal d
// continues exactly where a continuously drawn stream is at d.
func TestSnapshotRestoreConsistency(t *testing.T) {
	t.Parallel()
	const jump = 1000
	continuous := New(31337)
	reference := make([]uint64, 0, jump+200)
	for i := 0; i < jump+200; i++ {
		reference = append(reference, continuous.Uint64())
	}

	jumped := New(31337)
	if err := jumped.Restore(31337, jump); err != nil {
		t.Fatalf("Restore err = %v", err)
	}
	for i := 0; i < 200; i++ {
		if jumped.Uint64() != reference[jump+i] {
			t.Fatalf("restore continuation diverges at offset %d", i)
		}
	}
}

// TestSnapshotAfterRestoreIsConsistent asserts the values drawn after a Restore
// exactly match the corresponding ordinals of a continuously drawn stream, which
// is the determinism contract that makes replay reproducible.
func TestSnapshotAfterRestoreIsConsistent(t *testing.T) {
	t.Parallel()
	const jump = 500
	const after = 40

	continuous := New(99)
	reference := make([]uint64, 0, jump+after)
	for i := 0; i < jump+after; i++ {
		reference = append(reference, continuous.Uint64())
	}

	restored := New(99)
	if err := restored.Restore(99, jump); err != nil {
		t.Fatalf("Restore err = %v", err)
	}
	snap := restored.Snapshot()
	if len(snap) != 0 {
		t.Fatalf("ledger not reset after Restore: len = %d", len(snap))
	}
	for i := 0; i < after; i++ {
		restored.Uint64()
	}
	snap = restored.Snapshot()
	if len(snap) != after {
		t.Fatalf("ledger len = %d, want %d", len(snap), after)
	}
	for i := 0; i < after; i++ {
		if snap[i].Result != reference[jump+i] {
			t.Fatalf("post-restore draw %d = %d, want continuous ordinal %d = %d",
				i, snap[i].Result, jump+i, reference[jump+i])
		}
	}
}
