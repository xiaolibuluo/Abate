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

// TestMutationObservationChangesVerdict proves stream tampering is detected.
func TestMutationObservationChangesVerdict(t *testing.T) {
	t.Parallel()
	base := []observe.Observation{
		sample("p", 1, 0, observe.NewNumber(42)),
		sample("p", 2, 10, observe.NewNumber(42)),
	}
	st := valueState("p", 42, Tolerance{Type: ToleranceAbsolute, Value: 1})
	st.HoldFor = 10
	baseline, err := json.Marshal(Verify(base, []SteadyState{st}))
	if err != nil {
		t.Fatalf("json.Marshal err = %v", err)
	}
	mutations := [][]observe.Observation{
		{sample("p", 1, 0, observe.NewNumber(999)), sample("p", 2, 10, observe.NewNumber(42))},
		{sample("p", 1, 0, observe.NewNumber(42)), sample("p", 2, 10, observe.NewNumber(44))},
		{sample("p", 1, 0, observe.NewNumber(42)), sample("p", 2, 9, observe.NewNumber(42))},
		{sample("p-x", 1, 0, observe.NewNumber(42)), sample("p-x", 2, 10, observe.NewNumber(42))},
	}
	for i, m := range mutations {
		got, err := json.Marshal(Verify(m, []SteadyState{st}))
		if err != nil {
			t.Fatalf("mutation %d json.Marshal err = %v", i, err)
		}
		if string(got) == string(baseline) {
			t.Fatalf("mutation %d was not detected", i)
		}
	}
}

// TestMutationScoreChanges proves a changed verdict changes the score.
func TestMutationScoreChanges(t *testing.T) {
	t.Parallel()
	verdicts := []MutationVerdict{
		MutantKilled, MutantSurvived, MutantEquivalent, MutantTimeout,
	}
	base := Summarize(verdicts)
	mutated := append([]MutationVerdict(nil), verdicts...)
	mutated[1] = MutantKilled
	if Summarize(mutated).Value() == base.Value() {
		t.Fatalf("verdict mutation did not change the score")
	}
	equivalent := append([]MutationVerdict(nil), verdicts...)
	equivalent[1] = MutantEquivalent
	if Summarize(equivalent).Value() <= base.Value() {
		t.Fatalf("an equivalent mutant must leave the denominator")
	}
}

// TestMutationReplayRejudgement checks replayed streams judge identically.
func TestMutationReplayRejudgement(t *testing.T) {
	t.Parallel()
	stream := []observe.Observation{
		sample("p", 1, 0, observe.NewNumber(41)),
		sample("p", 2, 5, observe.NewNumber(42)),
		sample("p", 3, 10, observe.NewNumber(43)),
		sample("p", 4, 15, observe.NewNumber(900)),
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
			t.Fatalf("re-judgement %d diverged", i)
		}
	}
}
