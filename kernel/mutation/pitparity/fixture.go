/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package pitparity

import (
	"crypto/sha256"
	"slices"
	"strings"

	"abate/kernel/mutation"
	"abate/kernel/verify"
)

// fixtureTarget is the Abate side synthetic target of the parity harness.
//
// It mirrors the KRN-MUT fixture declared in kernel/mutation/helpers_test.go:
// the same ten nodes, tokens, kill thresholds, equivalent flag, timeout flag
// and coverage counters, so the parity verdicts are the verdicts the kernel
// itself produces for that fixture.
type fixtureTarget struct {
	nodes      []string
	tokens     map[string]string
	killFrom   map[string]string
	equivalent map[string]bool
	timeout    map[string]bool
}

// Fingerprint returns the stable digest of the fixture target.
func (t *fixtureTarget) Fingerprint() [32]byte {
	var b strings.Builder
	for _, node := range t.nodes {
		_, _ = b.WriteString(node)
		_ = b.WriteByte('=')
		_, _ = b.WriteString(t.tokens[node])
		_ = b.WriteByte(';')
	}
	return sha256.Sum256([]byte(b.String()))
}

// Nodes returns the mutable node references in deterministic order.
func (t *fixtureTarget) Nodes() []string {
	return append([]string(nil), t.nodes...)
}

// Token returns the original lexical token of one node.
func (t *fixtureTarget) Token(node string) string { return t.tokens[node] }

// Equivalent reports whether one mutant is semantically equivalent.
func (t *fixtureTarget) Equivalent(p mutation.MutationPoint) bool {
	return t.equivalent[p.Node] || p.Mutated == p.Original
}

// Evaluate reports the steady state outcome of one test on one mutant.
func (t *fixtureTarget) Evaluate(p mutation.MutationPoint, test string) verify.Outcome {
	if t.timeout[p.Node] {
		return verify.OutcomeInconclusive
	}
	from := t.killFrom[p.Node]
	if from != "" && test >= from {
		return verify.OutcomeViolated
	}
	return verify.OutcomeHeld
}

// fixtureSuite is the Abate side test suite of the regression harness.
//
// The three tests carry the t1/t2/t3 semantics of the self-developed baseline
// port declared in baseline.go.
type fixtureSuite struct {
	tests []string
	cover map[string]uint64
}

// Fingerprint returns the stable digest of the fixture suite.
func (s *fixtureSuite) Fingerprint() [32]byte {
	return sha256.Sum256([]byte(strings.Join(s.tests, "|")))
}

// Tests returns the test identifiers in deterministic order.
func (s *fixtureSuite) Tests() []string {
	return append([]string(nil), s.tests...)
}

// Coverage returns the integer number of tests covering one node.
func (s *fixtureSuite) Coverage(node string) uint64 { return s.cover[node] }

// newFixture builds the regression fixture: ten nodes covering the seven
// operator classes plus the equivalent, timeout and no kill paths.
//
// The node set, tokens, kill thresholds and coverage counters are identical to
// the kernel fixture, which is what makes the Abate verdicts of this harness
// the verdicts of the real mutation kernel.
func newFixture() (*fixtureTarget, *fixtureSuite) {
	nodes := []string{
		"a-arith", "b-bound", "c-const", "d-control", "e-return",
		"f-negate", "g-call", "h-equivalent", "i-timeout", "j-survived",
	}
	tokens := map[string]string{
		"a-arith": "+", "b-bound": "<", "c-const": "1", "d-control": "break",
		"e-return": "true", "f-negate": "==", "g-call": "call",
		"h-equivalent": "+", "i-timeout": "-", "j-survived": "*",
	}
	killFrom := map[string]string{
		"a-arith": "t2", "b-bound": "t3", "c-const": "t1", "d-control": "t2",
		"e-return": "t1", "f-negate": "t3", "g-call": "t2", "j-survived": "",
	}
	equivalent := map[string]bool{"h-equivalent": true}
	timeout := map[string]bool{"i-timeout": true}
	cover := map[string]uint64{
		"a-arith": 3, "b-bound": 2, "c-const": 1, "d-control": 4,
		"e-return": 5, "f-negate": 2, "g-call": 3, "h-equivalent": 1,
		"i-timeout": 1, "j-survived": 0,
	}
	slices.Sort(nodes)
	tests := []string{"t1", "t2", "t3"}
	return &fixtureTarget{
		nodes: nodes, tokens: tokens, killFrom: killFrom,
		equivalent: equivalent, timeout: timeout,
	}, &fixtureSuite{tests: tests, cover: cover}
}
