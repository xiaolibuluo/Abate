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

// assertStreamState checks the origin and draw ordinal reported by State.
func assertStreamState(t *testing.T, s *Stream, wantSeed, wantDraws uint64) {
	t.Helper()
	gotSeed, gotDraws := s.State()
	if gotSeed != wantSeed || gotDraws != wantDraws {
		t.Fatalf("State() = (%d, %d), want (%d, %d)",
			gotSeed, gotDraws, wantSeed, wantDraws)
	}
}

// TestStreamStateTracksOriginAndDraws proves State is a pure observer of the
// stream origin and the number of draws consumed, including after Restore and
// after a Fork.
func TestStreamStateTracksOriginAndDraws(t *testing.T) {
	t.Parallel()
	s := New(42)
	assertStreamState(t, s, 42, 0)
	s.Uint64()
	s.Int63()
	s.Float64()
	assertStreamState(t, s, 42, 3)
	_, before := s.State()
	_, after := s.State()
	if before != after {
		t.Fatalf("State consumed a draw: %d -> %d", before, after)
	}
	if err := s.Restore(42, 17); err != nil {
		t.Fatalf("Restore(42, 17): %v", err)
	}
	assertStreamState(t, s, 42, 17)
	if err := s.Restore(99, 5); !errors.Is(err, ErrInvalidRestore) {
		t.Fatalf("Restore with a foreign seed = %v, want ErrInvalidRestore", err)
	}
	assertStreamState(t, s, 42, 17)
	fresh := New(7)
	child := fresh.Fork()
	childSeed, childDraws := child.State()
	if childDraws != 0 {
		t.Fatalf("forked stream starts at draw %d, want 0", childDraws)
	}
	if childSeed != New(7).Uint64() {
		t.Fatalf("forked origin = %d, want the first draw of seed 7", childSeed)
	}
	_, parentDraws := fresh.State()
	if parentDraws != 1 {
		t.Fatalf("Fork consumed %d parent draws, want 1", parentDraws)
	}
	p1, p2 := New(123), New(123)
	c1, c2 := p1.Fork(), p2.Fork()
	s1, _ := c1.State()
	s2, _ := c2.State()
	if s1 != s2 {
		t.Fatalf("identical parents forked to origins %d and %d", s1, s2)
	}
}
