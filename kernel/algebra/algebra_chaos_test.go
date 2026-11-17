/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package algebra

import (
	"testing"

	"abate/sim/rand"
)

// TestChaosAlgebraAssociativity storms 1000 random operator triples and asserts
// the associativity law holds for all of them (flatten-based fingerprints make
// composition associative regardless of family mix).
func TestChaosAlgebraAssociativity(t *testing.T) {
	t.Parallel()
	const rounds = 1000
	l := NewLaws()
	s := rand.New(0xC0FFEE)
	fams := []Family{FamilyMutation, FamilyChaos, FamilyIdentity, FamilyMixed}
	pool := make([]*Basic, 0, 16)
	for i := 0; i < 16; i++ {
		f := fams[s.IntN(len(fams))]
		pool = append(pool, NewBasic(
			OperatorID("op-"+itoa(i)),
			f,
			Params{"i": itoa(i)},
			[]TargetKey{TargetKey("t" + itoa(i%4))},
			true,
		))
	}
	for r := 0; r < rounds; r++ {
		p1 := pool[s.IntN(len(pool))]
		p2 := pool[s.IntN(len(pool))]
		p3 := pool[s.IntN(len(pool))]
		if !l.CheckAssociativity(p1, p2, p3) {
			t.Fatalf("associativity failed for round %d: (%s ∘ (%s ∘ %s))", r, p1.ID(), p2.ID(), p3.ID())
		}
	}
}

// TestChaosAlgebraDeterminism asserts the same random triple sequence is
// reproducible from the same seed.
func TestChaosAlgebraDeterminism(t *testing.T) {
	t.Parallel()
	first := algebraHash(4242)
	for i := 0; i < 50; i++ {
		if got := algebraHash(4242); got != first {
			t.Fatalf("algebra hash diverged at round %d", i)
		}
	}
}

func algebraHash(seed uint64) uint64 {
	l := NewLaws()
	s := rand.New(seed)
	fams := []Family{FamilyMutation, FamilyChaos, FamilyIdentity}
	var h uint64
	for i := 0; i < 500; i++ {
		f := fams[s.IntN(len(fams))]
		op := NewBasic(OperatorID("o"+itoa(i)), f, Params{}, []TargetKey{TargetKey("t" + itoa(i%3))}, true)
		if l.CheckDual(op) {
			h ^= uint64(op.Fingerprint()[0])
			h = h<<13 | h>>51
		}
	}
	return h
}

func itoa(i int) string {
	if i == 0 {
		return "0"
	}
	neg := i < 0
	if neg {
		i = -i
	}
	buf := make([]byte, 0, 12)
	for i > 0 {
		buf = append([]byte{byte('0' + i%10)}, buf...)
		i /= 10
	}
	if neg {
		buf = append([]byte{'-'}, buf...)
	}
	return string(buf)
}
