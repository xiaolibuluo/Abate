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

// chaosSrc is a local deterministic integer source used only by chaos tests.
type chaosSrc struct{ state uint64 }

// next advances the generator with integer arithmetic.
func (g *chaosSrc) next() uint64 {
	g.state = g.state*6364136223846793005 + 1442695040888963407
	return g.state
}

// index returns a deterministic slice index in [0, n).
func (g *chaosSrc) index(n int) int {
	return int(g.next() % uint64(n)) //nolint:gosec // the modulo result is below n and fits int.
}

// nanos returns a deterministic positive virtual-time step.
func (g *chaosSrc) nanos(n int64) int64 {
	return int64(g.next()%uint64(n)) + 1 //nolint:gosec // the modulo result is below n and fits int64.
}

// TestChaosVerifyIsStable replays a seeded observation storm 100 times.
func TestChaosVerifyIsStable(t *testing.T) {
	t.Parallel()
	g := chaosSrc{state: 0x9E3779B97F4A7C15}
	probes := []string{"p-a", "p-b", "p-c"}
	numbers := []float64{0, 10, 25, 30, 50, 75, 99, 100}
	stream := make([]observe.Observation, 0, 500)
	var seq uint64 = 1
	var at int64
	for i := 0; i < 500; i++ {
		seq++
		at += g.nanos(10)
		probe := probes[g.index(len(probes))]
		value := numbers[g.index(len(numbers))]
		stream = append(stream, sample(probe, seq, at, observe.NewNumber(value)))
	}
	states := []SteadyState{
		valueState("p-a", 50, Tolerance{Type: ToleranceAbsolute, Value: 25}),
		{ProbeID: "p-b", Expectation: Expectation{Type: ExpectPredicate, Predicate: ">=30"}},
		{
			ProbeID:     "p-c",
			Expectation: Expectation{Type: ExpectRange, Min: 20, Max: 80},
			Tolerance:   Tolerance{Type: ToleranceAbsolute, Value: 5},
		},
	}
	baseline, err := json.Marshal(Verify(stream, states))
	if err != nil {
		t.Fatalf("json.Marshal err = %v", err)
	}
	for i := 0; i < 100; i++ {
		got, err := json.Marshal(Verify(stream, states))
		if err != nil || string(got) != string(baseline) {
			t.Fatalf("chaos replay %d diverged", i)
		}
	}
}

// TestChaosToleranceCombinationsNeverPanics explores mixed value/tolerance pairs.
func TestChaosToleranceCombinationsNeverPanics(t *testing.T) {
	t.Parallel()
	g := chaosSrc{state: 0xD1B54A32D192ED03}
	tols := []Tolerance{
		{Type: ToleranceAbsolute, Value: 1},
		{Type: ToleranceRelative, Value: 0.1},
		{Type: TolerancePercent, Value: 5},
		{Type: "unknown", Value: 3},
	}
	expectations := []Expectation{
		{Type: ExpectValue, Value: 50},
		{Type: ExpectRange, Min: 20, Max: 80},
		{Type: ExpectPredicate, Predicate: ">=30"},
		{Type: "unknown"},
	}
	values := []observe.Value{
		observe.NewNumber(0),
		observe.NewNumber(100),
		observe.NewText("42"),
		observe.NewText("not-a-number"),
		observe.NewBoolean(true),
		observe.NewHistogram([]uint64{1, 2}),
	}
	for i := 0; i < 2000; i++ {
		st := SteadyState{
			ProbeID:     "p",
			Expectation: expectations[g.index(len(expectations))],
			Tolerance:   tols[g.index(len(tols))],
			HoldFor:     g.nanos(20),
		}
		stream := []observe.Observation{
			sample("p", 1, 0, values[g.index(len(values))]),
			sample("p", 2, g.nanos(20), values[g.index(len(values))]),
		}
		got := Verify(stream, []SteadyState{st})
		if len(got) != 1 {
			t.Fatalf("iteration %d produced %d verdicts", i, len(got))
		}
	}
}
