/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package verify

import (
	"fmt"
	"strconv"
	"strings"

	"abate/kernel/observe"
)

// Verify evaluates every steady-state hypothesis over an observation stream.
//
// The function is pure: the same stream and hypotheses always produce the
// same verdicts, in the same order, with byte-identical serialization.
func Verify(obs []observe.Observation, states []SteadyState) []Verdict {
	out := make([]Verdict, 0, len(states))
	for _, ss := range states {
		out = append(out, verifyOne(obs, ss))
	}
	return out
}

// Classify maps steady-state verdicts to a mutation verdict.
//
// A violating assertion means the test suite killed the mutant, an
// inconclusive assertion yields TIMEOUT, and an all-held run means the mutant
// survived. EQUIVALENT comes from the separate equivalence analysis and is
// passed straight into Summarize. A wall-clock timeout is reported separately
// by the caller as MutantTimeout.
func Classify(verdicts []Verdict) MutationVerdict {
	for _, v := range verdicts {
		switch v.Outcome {
		case OutcomeViolated:
			return MutantKilled
		case OutcomeInconclusive:
			return MutantTimeout
		}
	}
	return MutantSurvived
}

// verifyOne judges a single hypothesis against its probe stream.
func verifyOne(obs []observe.Observation, ss SteadyState) Verdict {
	v := Verdict{ProbeID: ss.ProbeID, Outcome: OutcomeHeld}
	samples := make([]observe.Observation, 0, len(obs))
	for _, o := range obs {
		if o.ProbeID == ss.ProbeID {
			samples = append(samples, o)
		}
	}
	if len(samples) == 0 {
		v.Outcome = OutcomeInconclusive
		v.Reason = "no observations for probe " + ss.ProbeID
		return v
	}

	violated := false
	runStart := int64(-1)
	var best int64
	for _, o := range samples {
		if !matches(o.Value, ss.Expectation, ss.Tolerance) {
			if !violated {
				violated = true
				v.FirstSeq = o.Seq
				v.Reason = fmt.Sprintf("seq=%d observed=%g expectation=%s tolerance=%s",
					o.Seq, numericOrZero(o.Value), ss.Expectation.Type, ss.Tolerance.Type)
			}
			v.Timeline = append(v.Timeline, o.Seq)
			runStart = -1
			continue
		}
		if runStart < 0 {
			runStart = o.At
		}
		if span := o.At - runStart; span > best {
			best = span
		}
	}
	if violated {
		v.Outcome = OutcomeViolated
		return v
	}
	if best < ss.HoldFor {
		v.Outcome = OutcomeInconclusive
		v.Reason = fmt.Sprintf("held span %d below holdFor %d", best, ss.HoldFor)
		return v
	}
	v.Reason = fmt.Sprintf("held span %d satisfies holdFor %d", best, ss.HoldFor)
	return v
}

// matches reports whether a value satisfies an expectation within tolerance.
func matches(v observe.Value, e Expectation, tol Tolerance) bool {
	switch e.Type {
	case ExpectValue:
		n, ok := numeric(v)
		if !ok {
			return false
		}
		return abs(n-e.Value) <= tol.margin(e.Value)
	case ExpectRange:
		n, ok := numeric(v)
		if !ok {
			return false
		}
		margin := tol.margin((e.Min + e.Max) / 2)
		return n >= e.Min-margin && n <= e.Max+margin
	case ExpectPredicate:
		return evalPredicate(v, e.Predicate)
	default:
		return false
	}
}

// numeric extracts a comparable number from numeric or textual values.
func numeric(v observe.Value) (float64, bool) {
	switch v.Kind {
	case observe.ValueNumber:
		return v.Number, true
	case observe.ValueText:
		n, err := strconv.ParseFloat(v.Text, 64)
		return n, err == nil
	default:
		return 0, false
	}
}

// numericOrZero extracts a number for diagnostics, defaulting to zero.
func numericOrZero(v observe.Value) float64 {
	n, ok := numeric(v)
	if !ok {
		return 0
	}
	return n
}

// evalPredicate evaluates a bounded comparison predicate against a value.
func evalPredicate(v observe.Value, expr string) bool {
	op, threshold, ok := parsePredicate(expr)
	if !ok {
		return false
	}
	n, ok := numeric(v)
	if !ok {
		return false
	}
	switch op {
	case ">":
		return n > threshold
	case ">=":
		return n >= threshold
	case "<":
		return n < threshold
	case "<=":
		return n <= threshold
	case "==":
		return n == threshold
	default:
		return n != threshold
	}
}

// parsePredicate splits a comparison operator from its numeric threshold.
func parsePredicate(expr string) (string, float64, bool) {
	ops := []string{">=", "<=", "==", "!=", ">", "<"}
	for _, op := range ops {
		if strings.HasPrefix(expr, op) {
			rest := strings.TrimSpace(strings.TrimPrefix(expr, op))
			n, err := strconv.ParseFloat(rest, 64)
			if err != nil {
				return "", 0, false
			}
			return op, n, true
		}
	}
	return "", 0, false
}
