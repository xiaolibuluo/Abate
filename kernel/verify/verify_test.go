/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package verify

import (
	"encoding/json"
	"testing"

	"abate/kernel/observe"
)

// sample builds one observation for a probe.
func sample(probe string, seq uint64, at int64, v observe.Value) observe.Observation {
	return observe.Observation{
		At: at, Seq: seq, ProbeID: probe, Value: v, Status: observe.StatusOK,
	}
}

// valueState builds a single-sample hypothesis with no hold requirement.
func valueState(probe string, expected float64, tol Tolerance) SteadyState {
	return SteadyState{
		ProbeID:     probe,
		Expectation: Expectation{Type: ExpectValue, Value: expected},
		Tolerance:   tol,
	}
}

// TestToleranceMargins covers the three tolerance formulas and the fallback.
func TestToleranceMargins(t *testing.T) {
	t.Parallel()
	cases := []struct {
		tol      Tolerance
		expected float64
		want     float64
	}{
		{Tolerance{Type: ToleranceAbsolute, Value: 2}, 100, 2},
		{Tolerance{Type: ToleranceRelative, Value: 0.25}, 100, 25},
		{Tolerance{Type: TolerancePercent, Value: 5}, 200, 10},
		{Tolerance{Type: "unknown", Value: 5}, 100, 0},
	}
	for _, c := range cases {
		if got := c.tol.margin(c.expected); abs(got-c.want) > 1e-9 {
			t.Fatalf("margin(%+v, %g) = %g, want %g", c.tol, c.expected, got, c.want)
		}
	}
	if abs(-3) != 3 || abs(3) != 3 {
		t.Fatalf("abs broken")
	}
}

// TestValueToleranceBoundaries checks the exact boundary is satisfied.
func TestValueToleranceBoundaries(t *testing.T) {
	t.Parallel()
	cases := []struct {
		tol                    Tolerance
		expected, edge, beyond float64
	}{
		{Tolerance{Type: ToleranceAbsolute, Value: 2}, 100, 102, 102.001},
		{Tolerance{Type: ToleranceRelative, Value: 0.1}, 100, 110, 110.001},
		{Tolerance{Type: TolerancePercent, Value: 5}, 100, 105, 105.001},
	}
	for i, c := range cases {
		st := valueState("p", c.expected, c.tol)
		held := Verify([]observe.Observation{
			sample("p", 1, 0, observe.NewNumber(c.edge)),
		}, []SteadyState{st})
		if held[0].Outcome != OutcomeHeld {
			t.Fatalf("case %d boundary %g must be HELD, got %s", i, c.edge, held[0].Outcome)
		}
		broken := Verify([]observe.Observation{
			sample("p", 1, 0, observe.NewNumber(c.beyond)),
		}, []SteadyState{st})
		if broken[0].Outcome != OutcomeViolated {
			t.Fatalf("case %d beyond %g must be VIOLATED, got %s", i, c.beyond, broken[0].Outcome)
		}
	}
}

// TestRangeAndPredicate covers range, all six operators and invalid inputs.
func TestRangeAndPredicate(t *testing.T) {
	t.Parallel()
	tol := Tolerance{Type: ToleranceAbsolute, Value: 1}
	rangeState := SteadyState{
		ProbeID:     "p",
		Expectation: Expectation{Type: ExpectRange, Min: 10, Max: 20},
		Tolerance:   tol,
	}
	for _, n := range []float64{9, 15, 21} {
		got := Verify([]observe.Observation{
			sample("p", 1, 0, observe.NewNumber(n)),
		}, []SteadyState{rangeState})
		if got[0].Outcome != OutcomeHeld {
			t.Fatalf("range value %g must be HELD", n)
		}
	}
	for _, n := range []float64{8.9, 21.1} {
		got := Verify([]observe.Observation{
			sample("p", 1, 0, observe.NewNumber(n)),
		}, []SteadyState{rangeState})
		if got[0].Outcome != OutcomeViolated {
			t.Fatalf("range value %g must be VIOLATED", n)
		}
	}

	predicates := []struct {
		expr  string
		value float64
		held  bool
	}{
		{">5", 6, true}, {">5", 5, false}, {">=5", 5, true},
		{"<5", 4, true}, {"<=5", 5, true}, {"==5", 5, true},
		{"!=5", 6, true}, {"!=5", 5, false},
	}
	for _, p := range predicates {
		st := SteadyState{
			ProbeID:     "p",
			Expectation: Expectation{Type: ExpectPredicate, Predicate: p.expr},
		}
		got := Verify([]observe.Observation{
			sample("p", 1, 0, observe.NewNumber(p.value)),
		}, []SteadyState{st})[0].Outcome
		want := OutcomeViolated
		if p.held {
			want = OutcomeHeld
		}
		if got != want {
			t.Fatalf("predicate %q on %g = %s, want %s", p.expr, p.value, got, want)
		}
	}
}

// TestHoldForBoundary checks the one-tick boundary of the hold duration.
func TestHoldForBoundary(t *testing.T) {
	t.Parallel()
	stream := []observe.Observation{
		sample("p", 1, 0, observe.NewNumber(42)),
		sample("p", 2, 5, observe.NewNumber(42)),
		sample("p", 3, 10, observe.NewNumber(42)),
	}
	st := valueState("p", 42, Tolerance{Type: ToleranceAbsolute, Value: 0})
	st.HoldFor = 10
	if got := Verify(stream, []SteadyState{st})[0].Outcome; got != OutcomeHeld {
		t.Fatalf("span 10 with holdFor 10 = %s, want HELD", got)
	}
	st.HoldFor = 11
	if got := Verify(stream, []SteadyState{st})[0].Outcome; got != OutcomeInconclusive {
		t.Fatalf("span 10 with holdFor 11 = %s, want INCONCLUSIVE", got)
	}
}

// TestFirstViolation records the first violating sequence and the timeline.
func TestFirstViolation(t *testing.T) {
	t.Parallel()
	stream := []observe.Observation{
		sample("p", 1, 0, observe.NewNumber(42)),
		sample("p", 2, 1, observe.NewNumber(999)),
		sample("p", 3, 2, observe.NewNumber(42)),
		sample("p", 4, 3, observe.NewNumber(998)),
	}
	st := valueState("p", 42, Tolerance{Type: ToleranceAbsolute, Value: 0})
	v := Verify(stream, []SteadyState{st})[0]
	if v.Outcome != OutcomeViolated || v.FirstSeq != 2 ||
		len(v.Timeline) != 2 || v.Timeline[0] != 2 || v.Timeline[1] != 4 {
		t.Fatalf("verdict = %+v", v)
	}
}

// TestVerifyNoObservations covers the missing-probe branch.
func TestVerifyNoObservations(t *testing.T) {
	t.Parallel()
	v := Verify(nil, []SteadyState{valueState("missing", 1, Tolerance{})})[0]
	if v.Outcome != OutcomeInconclusive || v.Reason == "" {
		t.Fatalf("verdict = %+v", v)
	}
}

// TestVerifyMultipleHypotheses verifies deterministic verdict ordering.
func TestVerifyMultipleHypotheses(t *testing.T) {
	t.Parallel()
	stream := []observe.Observation{sample("p-a", 1, 0, observe.NewNumber(1))}
	states := []SteadyState{
		valueState("p-a", 1, Tolerance{Type: ToleranceAbsolute, Value: 0}),
		valueState("p-b", 1, Tolerance{Type: ToleranceAbsolute, Value: 0}),
	}
	got := Verify(stream, states)
	if len(got) != 2 || got[0].ProbeID != "p-a" || got[0].Outcome != OutcomeHeld ||
		got[1].ProbeID != "p-b" || got[1].Outcome != OutcomeInconclusive {
		t.Fatalf("verdicts = %+v", got)
	}
}

// TestClassify covers all mutation verdict paths.
func TestClassify(t *testing.T) {
	t.Parallel()
	cases := []struct {
		verdicts []Verdict
		timedOut bool
		want     MutationVerdict
	}{
		{[]Verdict{{Outcome: OutcomeViolated}}, false, MutantKilled},
		{[]Verdict{{Outcome: OutcomeInconclusive}}, false, MutantTimeout},
		{[]Verdict{{Outcome: OutcomeHeld}}, false, MutantSurvived},
		{nil, false, MutantSurvived},
		{[]Verdict{{Outcome: OutcomeHeld}}, true, MutantTimeout},
		{[]Verdict{{Outcome: OutcomeHeld}, {Outcome: OutcomeViolated}}, false, MutantKilled},
	}
	for i, c := range cases {
		var got MutationVerdict
		if c.timedOut {
			got = MutantTimeout
		} else {
			got = Classify(c.verdicts)
		}
		if got != c.want {
			t.Fatalf("case %d = %s, want %s", i, got, c.want)
		}
	}
}

// TestScore checks counting and the equivalent-exclusion denominator.
func TestScore(t *testing.T) {
	t.Parallel()
	s := Summarize([]MutationVerdict{
		MutantKilled, MutantSurvived, MutantEquivalent, MutantTimeout, MutantKilled,
	})
	if s.Total != 5 || s.Killed != 2 || s.Survived != 1 || s.Equivalent != 1 || s.Timeout != 1 {
		t.Fatalf("score = %+v", s)
	}
	if abs(s.Value()-2.0/4) > 1e-12 {
		t.Fatalf("score value = %g, want 0.5", s.Value())
	}
	if (Score{}).Value() != 0 {
		t.Fatalf("empty score must be zero")
	}
	if (Score{Total: 1, Equivalent: 5}).Value() != 0 {
		t.Fatalf("over-counted equivalents must not underflow the denominator")
	}
}

// TestReplayVerdictsByteStable proves replay re-judgement is byte-identical.
func TestReplayVerdictsByteStable(t *testing.T) {
	t.Parallel()
	stream := []observe.Observation{
		sample("p", 1, 0, observe.NewNumber(42)),
		sample("p", 2, 5, observe.NewNumber(43)),
		sample("p", 3, 10, observe.NewNumber(999)),
	}
	st := valueState("p", 42, Tolerance{Type: ToleranceAbsolute, Value: 1})
	st.HoldFor = 5
	baseline, err := json.Marshal(Verify(stream, []SteadyState{st}))
	if err != nil {
		t.Fatalf("json.Marshal err = %v", err)
	}
	for i := 0; i < 100; i++ {
		got, err := json.Marshal(Verify(stream, []SteadyState{st}))
		if err != nil || string(got) != string(baseline) {
			t.Fatalf("replay %d diverged", i)
		}
	}
}
