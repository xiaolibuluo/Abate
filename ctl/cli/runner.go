/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package cli

import (
	"context"
	"fmt"
	"strconv"

	"abate/ces/schema"
	"abate/kernel/observe"
	"abate/kernel/scheduler"
	"abate/kernel/verify"
	"abate/sim/rand"
	"abate/sim/record"
)

// probeSamples is the fixed number of deterministic probe collections.
const probeSamples = 4

// execution is the deterministic result of one scheduler-driven experiment.
type execution struct {
	result       scheduler.Result
	verdicts     []verify.Verdict
	streamDigest [32]byte
	observations int
	archive      *record.Archive
}

// runExperiment drives the CES document through KRN-SCH, KRN-OBS and KRN-VER.
func runExperiment(exp *schema.Experiment) (execution, error) {
	times := probeSchedule(exp)
	pipe, err := buildPipeline(exp)
	if err != nil {
		return execution{}, err
	}
	obs := collectObservations(pipe, times)
	h := newExperimentHandler(exp, times, obs)
	result := scheduler.New(h.handler(), scheduler.DefaultBudget()).Run(context.Background(),
		scheduler.Experiment{Seed: exp.Spec.Seed, Initial: buildInitial(exp)})
	if result.Err != nil {
		return execution{}, fmt.Errorf("scheduler: %w", result.Err)
	}
	arch, err := h.stop()
	if err != nil {
		return execution{}, fmt.Errorf("record: %w", err)
	}
	return execution{
		result:       result,
		verdicts:     verify.Verify(obs, buildSteadyStates(exp)),
		streamDigest: observe.StreamDigest(obs),
		observations: len(obs),
		archive:      arch,
	}, nil
}

// buildInitial creates the single apply event that seeds the event cascade.
func buildInitial(exp *schema.Experiment) []scheduler.Event {
	rng := rand.New(exp.Spec.Seed)
	return []scheduler.Event{{
		At:      perturbationTime(exp),
		Kind:    scheduler.EventApply,
		Payload: []byte(fmt.Sprintf("%s/%016x", exp.Spec.Method.Expression, rng.Uint64())),
	}}
}

// cascadeHandler derives probe and rollback events from one apply event.
func cascadeHandler(times []scheduler.VirtualTime,
	interval scheduler.VirtualTime,
) scheduler.Handler {
	return func(e scheduler.Event) ([]scheduler.Event, error) {
		if e.Kind != scheduler.EventApply {
			return nil, nil
		}
		children := make([]scheduler.Event, 0, len(times)+1)
		for i, at := range times {
			children = append(children, scheduler.Event{
				At:      at,
				Kind:    scheduler.EventProbe,
				Payload: []byte(fmt.Sprintf("probe/%d", i+1)),
			})
		}
		children = append(children, scheduler.Event{
			At:      times[len(times)-1].Add(interval),
			Kind:    scheduler.EventRollback,
			Payload: []byte("rollback"),
		})
		return children, nil
	}
}

// probeSchedule derives the fixed virtual times at which probes are sampled.
func probeSchedule(exp *schema.Experiment) []scheduler.VirtualTime {
	interval := probeInterval(exp)
	at := perturbationTime(exp)
	times := make([]scheduler.VirtualTime, 0, probeSamples)
	for i := 0; i < probeSamples; i++ {
		at = at.Add(interval)
		times = append(times, at)
	}
	return times
}

// perturbationTime resolves the virtual time of the perturbation event.
func perturbationTime(exp *schema.Experiment) scheduler.VirtualTime {
	if exp.Spec.Schedule.Once != nil {
		return scheduler.VirtualTime(exp.Spec.Schedule.Once.At)
	}
	if exp.Spec.Schedule.Continuous != nil {
		return scheduler.VirtualTime(exp.Spec.Schedule.Continuous.Duration)
	}
	return scheduler.Second
}

// probeInterval resolves the virtual sampling interval of the first probe.
func probeInterval(exp *schema.Experiment) scheduler.VirtualTime {
	if len(exp.Spec.Probes) > 0 && exp.Spec.Probes[0].Interval > 0 {
		return scheduler.VirtualTime(exp.Spec.Probes[0].Interval)
	}
	return scheduler.Second
}

// buildPipeline registers every CES probe as a deterministic source.
func buildPipeline(exp *schema.Experiment) (*observe.Pipeline, error) {
	pipe := observe.NewPipeline()
	for i := range exp.Spec.Probes {
		probe, err := newCESProbe(&exp.Spec.Probes[i])
		if err != nil {
			return nil, fmt.Errorf("probe %q: %w", exp.Spec.Probes[i].ID, err)
		}
		if err := pipe.Register(probe); err != nil {
			return nil, fmt.Errorf("probe registration: %w", err)
		}
	}
	return pipe, nil
}

// collectObservations samples every probe at each scheduled virtual time.
func collectObservations(pipe *observe.Pipeline,
	times []scheduler.VirtualTime,
) []observe.Observation {
	seq := uint64(1)
	all := make([]observe.Observation, 0, len(times))
	for _, at := range times {
		batch := pipe.Collect(at.UnixNano(), seq)
		all = append(all, batch...)
		seq += uint64(len(batch)) //nolint:gosec // len is non-negative and bounded by the probe count.
	}
	return all
}

// buildSteadyStates projects CES hypotheses onto the kernel verifier model.
func buildSteadyStates(exp *schema.Experiment) []verify.SteadyState {
	states := make([]verify.SteadyState, 0, len(exp.Spec.SteadyState))
	for _, ss := range exp.Spec.SteadyState {
		states = append(states, verify.SteadyState{
			ProbeID: ss.ProbeRef,
			Expectation: verify.Expectation{
				Type:      ss.Expectation.Type,
				Value:     numericLiteral(ss.Expectation.Value),
				Min:       numericLiteral(ss.Expectation.Min),
				Max:       numericLiteral(ss.Expectation.Max),
				Predicate: ss.Expectation.Predicate,
			},
			Tolerance: verify.Tolerance{
				Type: ss.Tolerance.Type, Value: ss.Tolerance.Value,
			},
			HoldFor: int64(ss.HoldFor),
		})
	}
	return states
}

// numericLiteral parses a canonical CES numeric string, defaulting to zero.
func numericLiteral(s string) float64 {
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}
	return v
}

// steadyViolated reads kernel verifier outcomes; the CLI never judges a value.
func steadyViolated(verdicts []verify.Verdict) bool {
	for _, v := range verdicts {
		if v.Outcome == verify.OutcomeViolated {
			return true
		}
	}
	return false
}
