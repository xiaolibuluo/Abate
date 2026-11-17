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

// ErrBaselineMismatch reports that the local baseline port disagrees with the
// assertions fixed for the self-developed fixture.
var ErrBaselineMismatch = errors.New("pitparity: baseline port disagrees with the fixture assertions")

// witnessCase binds one self-developed case to the observations of the
// original and of the mutated baseline method.
//
// The witness is the independent oracle of the harness: it decides whether the
// suite can tell the mutant from the original without looking at the frozen
// expectation, so a mapping table cannot be tuned to the wanted outcome.
type witnessCase struct {
	// method is the local fixture method identifier of the case.
	method string

	// line is the local fixture line identifier.
	line int

	// mutator is the local fixture operator identifier of the case.
	mutator string

	// original observes the unmutated baseline method.
	original func(input int) observation

	// mutant observes the mutated baseline method.
	mutant func(input int) observation

	// inputs are the suite covered inputs of the mutated method.
	inputs []int

	// expected are the values the suite asserts for the inputs.
	expected []observation
}

// classifyWitness reports how the test suite sees one mutated site.
//
// The classification is derived from the local baseline port only: the
// original observations are first checked against the fixed assertions, so a
// porting mistake is reported instead of silently changing a verdict.
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
			return "", fmt.Errorf("%w: %s:%d input %d observes %s, the suite asserts %s",
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

// expectFromWitness maps one independent witness classification to the frozen
// expected classification stored in the self-developed golden.
func expectFromWitness(witness string) string {
	switch witness {
	case WitnessDetected:
		return ExpectKilled
	case WitnessEquivalent:
		return ExpectSurvived
	case WitnessUncovered:
		return ExpectNoCoverage
	default:
		return ""
	}
}

// witnessMatchesExpect reports whether the witness reproduces the frozen
// expectation.
//
// The mapping is the documented one: a detectable mutant is killed, a mutant
// the suite cannot distinguish survives, and an unreachable site has no cover.
func witnessMatchesExpect(witness, expect string) bool {
	return expectFromWitness(witness) == expect && expect != ""
}
