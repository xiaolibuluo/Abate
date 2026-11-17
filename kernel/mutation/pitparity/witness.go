/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package pitparity

import (
	"errors"
	"fmt"
)

// The classifications the semantic witness can reach.
const (
	// WitnessDetected means the suite distinguishes the mutant from the original.
	WitnessDetected = "detected"

	// WitnessEquivalent means the suite cannot distinguish the mutant.
	WitnessEquivalent = "equivalent"

	// WitnessUncovered means no test reaches the mutated site.
	WitnessUncovered = "uncovered"

	// WitnessNotApplicable marks a point without a semantic witness.
	WitnessNotApplicable = "not-applicable"
)

// ErrBaselineMismatch reports that the Go port of the Java baseline disagrees
// with the assertions of the Java suite.
var ErrBaselineMismatch = errors.New("pitparity: baseline port disagrees with the Java suite")

// witnessCase binds one PIT mutation site to the observations of the original
// and of the mutated baseline method.
//
// The witness is the independent oracle of the harness: it decides whether the
// suite can tell the mutant from the original without looking at the PIT
// status, so a mapping table cannot be tuned to the wanted outcome.
type witnessCase struct {
	// method is the mutated method of the Java baseline project.
	method string

	// line is the mutated source line.
	line int

	// mutator is the PIT mutator family of the mutation.
	mutator string

	// original observes the unmutated baseline method.
	original func(input int) observation

	// mutant observes the mutated baseline method.
	mutant func(input int) observation

	// inputs are the suite covered inputs of the mutated method.
	inputs []int

	// expected are the values the Java suite asserts for the inputs.
	expected []observation
}

// classifyWitness reports how the test suite sees one mutated site.
//
// The classification is derived from the ported baseline only: the original
// observations are first checked against the assertions of the Java suite, so
// a porting mistake is reported instead of silently changing a verdict.
func classifyWitness(c witnessCase) (string, error) {
	if len(c.inputs) != len(c.expected) {
		return "", fmt.Errorf("%w: %s:%d has %d inputs but %d expected values",
			ErrBaselineMismatch, c.method, c.line, len(c.inputs), len(c.expected))
	}
	if len(c.inputs) == 0 {
		return WitnessUncovered, nil
	}
	for i, input := range c.inputs {
		if got := c.original(input); got != c.expected[i] {
			return "", fmt.Errorf("%w: %s:%d input %d observes %s, the Java suite asserts %s",
				ErrBaselineMismatch, c.method, c.line, input, got, c.expected[i])
		}
	}
	for _, input := range c.inputs {
		if c.mutant(input) != c.original(input) {
			return WitnessDetected, nil
		}
	}
	return WitnessEquivalent, nil
}

// witnessMatchesStatus reports whether the witness reproduces the PIT status.
//
// The mapping is the documented one: a detectable mutant is killed, a mutant
// the suite cannot distinguish survives, and an unreachable site has no cover.
func witnessMatchesStatus(witness, status string) bool {
	switch witness {
	case WitnessDetected:
		return status == StatusKilled
	case WitnessEquivalent:
		return status == StatusSurvived
	case WitnessUncovered:
		return status == StatusNoCoverage
	default:
		return false
	}
}
