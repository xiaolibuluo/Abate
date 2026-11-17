/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package rand

import (
	"errors"
	"testing"
)

// TestMutationRestoreRoundTrip verifies that jumping to a draw ordinal via
// Restore produces the same continuation as a continuously drawn stream.
func TestMutationRestoreRoundTrip(t *testing.T) {
	t.Parallel()
	const jump = 1000
	continuous := New(31337)
	reference := make([]uint64, 0, 500)
	for i := 0; i < jump+500; i++ {
		reference = append(reference, continuous.Uint64())
	}

	jumped := New(31337)
	if err := jumped.Restore(31337, jump); err != nil {
		t.Fatalf("Restore err = %v", err)
	}
	for i := 0; i < 500; i++ {
		if jumped.Uint64() != reference[jump+i] {
			t.Fatalf("restore divergence at offset %d", i)
		}
	}
}

// TestMutationRestoreRejectsAlienSeed checks that Restore refuses a seed that
// does not match the stream origin, protecting replay reproducibility.
func TestMutationRestoreRejectsAlienSeed(t *testing.T) {
	t.Parallel()
	s := New(1)
	if err := s.Restore(2, 0); !errors.Is(err, ErrInvalidRestore) {
		t.Fatalf("Restore alien seed err = %v, want ErrInvalidRestore", err)
	}
}

// TestMutationForkDeterminism asserts a forked child is reproducible across two
// independently seeded parents sharing the fork seed.
func TestMutationForkDeterminism(t *testing.T) {
	t.Parallel()
	build := func() uint64 {
		p := New(909)
		c := p.Fork()
		return c.Uint64()
	}
	first := build()
	for i := 0; i < 100; i++ {
		if got := build(); got != first {
			t.Fatalf("forked child divergent at round %d", i)
		}
	}
}
