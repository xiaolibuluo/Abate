/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package mutation

import (
	"bytes"
	"context"
	"testing"

	"abate/kernel/verify"
)

// This is a deterministic-judgement parity check, not a PIT official CLI
// cross-check: the M2 sandbox has no JVM, so the oracle below is a set of
// hand-derived verdicts over synthetic targets. The nightly PIT CLI comparison
// remains an explicitly deferred M2 item and is never claimed here.

// handKillingTest independently derives the first killing test of a mutant.
func handKillingTest(t *syntheticTarget, p MutationPoint, tests []string) string {
	if t.equivalent[p.Node] || p.Mutated == p.Original || t.timeout[p.Node] {
		return ""
	}
	for _, test := range tests {
		from := t.killFrom[p.Node]
		if from != "" && test >= from {
			return test
		}
	}
	return ""
}

// handScore independently tallies hand-derived verdicts without Summarize.
func handScore(verdicts []verify.MutationVerdict) verify.Score {
	var score verify.Score
	for _, v := range verdicts {
		score.Total++
		switch v {
		case verify.MutantKilled:
			score.Killed++
		case verify.MutantSurvived:
			score.Survived++
		case verify.MutantEquivalent:
			score.Equivalent++
		case verify.MutantTimeout:
			score.Timeout++
		}
	}
	return score
}

func subsetTarget(base *syntheticTarget, from, to int) *syntheticTarget {
	return &syntheticTarget{
		nodes:      append([]string(nil), base.nodes[from:to]...),
		tokens:     base.tokens,
		killFrom:   base.killFrom,
		equivalent: base.equivalent,
		timeout:    base.timeout,
	}
}

func TestDeterministicOracleParity100Percent(t *testing.T) {
	t.Parallel()
	base, suite, tests := newFixture()
	targets := []*syntheticTarget{
		base,
		subsetTarget(base, 0, 5),
		subsetTarget(base, 5, len(base.nodes)),
	}
	seeds := []uint64{1, 42, 20261109}
	matched, total := 0, 0
	for _, target := range targets {
		for _, seed := range seeds {
			rep, err := NewEngine(seed).Run(context.Background(), target, suite)
			if err != nil {
				t.Fatalf("seed %d: %v", seed, err)
			}
			expected := make([]verify.MutationVerdict, 0, len(rep.Points))
			for i, point := range rep.Points {
				want := handVerdict(target, point, tests)
				expected = append(expected, want)
				total++
				if rep.Verdicts[i] == want {
					matched++
				} else {
					t.Errorf("seed=%d node=%s %q->%q: got %s, hand-derived %s",
						seed, point.Node, point.Original, point.Mutated, rep.Verdicts[i], want)
				}
				if got := handKillingTest(target, point, tests); rep.KillingTests[i] != got {
					t.Errorf("seed=%d node=%s: killing test %q, hand-derived %q",
						seed, point.Node, rep.KillingTests[i], got)
				}
			}
			if rep.Score != handScore(expected) {
				t.Fatalf("seed=%d score = %+v, hand-derived %+v", seed, rep.Score, handScore(expected))
			}
			if rep.Fingerprint != ReportFingerprint(rep) {
				t.Fatalf("seed=%d report fingerprint disagrees with canonical body", seed)
			}
			again, err := NewEngine(seed).Run(context.Background(), target, suite)
			if err != nil {
				t.Fatalf("second run seed %d: %v", seed, err)
			}
			if !bytes.Equal(rep.Bytes(), again.Bytes()) || reportSHA(rep) != reportSHA(again) {
				t.Fatalf("seed=%d double-run bytes differ", seed)
			}
		}
	}
	if matched != total || total == 0 {
		t.Fatalf("deterministic parity = %d/%d, want 100%%", matched, total)
	}
	t.Logf("deterministic judgement parity = %d/%d = 100%% (synthetic targets + hand-derived verdicts)",
		matched, total)
}
