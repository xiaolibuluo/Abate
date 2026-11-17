/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package mutation

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"slices"
	"strings"
	"time"

	"abate/kernel/verify"
)

// syntheticTarget is a deterministic, hand-computable Target used as the M2
// stand-in for LNG-AIR. A mutant is killed by every test at or after killFrom.
type syntheticTarget struct {
	nodes      []string
	tokens     map[string]string
	killFrom   map[string]string
	equivalent map[string]bool
	timeout    map[string]bool
}

func (t *syntheticTarget) Fingerprint() [32]byte {
	var b strings.Builder
	for _, node := range t.nodes {
		b.WriteString(node)
		b.WriteByte('=')
		b.WriteString(t.tokens[node])
		b.WriteByte(';')
	}
	return sha256.Sum256([]byte(b.String()))
}

func (t *syntheticTarget) Nodes() []string {
	return append([]string(nil), t.nodes...)
}

func (t *syntheticTarget) Token(node string) string { return t.tokens[node] }

func (t *syntheticTarget) Equivalent(p MutationPoint) bool {
	return t.equivalent[p.Node] || p.Mutated == p.Original
}

func (t *syntheticTarget) Evaluate(p MutationPoint, test string) verify.Outcome {
	if t.timeout[p.Node] {
		return verify.OutcomeInconclusive
	}
	from := t.killFrom[p.Node]
	if from != "" && test >= from {
		return verify.OutcomeViolated
	}
	return verify.OutcomeHeld
}

// syntheticSuite is a deterministic test selector with integer cover counts.
type syntheticSuite struct {
	tests []string
	cover map[string]uint64
}

func (s *syntheticSuite) Fingerprint() [32]byte {
	return sha256.Sum256([]byte(strings.Join(s.tests, "|")))
}

func (s *syntheticSuite) Tests() []string {
	return append([]string(nil), s.tests...)
}

func (s *syntheticSuite) Coverage(node string) uint64 { return s.cover[node] }

// newFixture builds a target covering all seven operator classes plus the
// equivalent, timeout and no-kill paths, and a suite of three ordered tests.
func newFixture() (*syntheticTarget, *syntheticSuite, []string) {
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
	return &syntheticTarget{
		nodes: nodes, tokens: tokens, killFrom: killFrom,
		equivalent: equivalent, timeout: timeout,
	}, &syntheticSuite{tests: tests, cover: cover}, tests
}

// handVerdict is the independent oracle: it derives a verdict from the plain
// fixture rules without calling any engine, classifier or minimizer code.
func handVerdict(t *syntheticTarget, p MutationPoint, tests []string) verify.MutationVerdict {
	if t.equivalent[p.Node] || p.Mutated == p.Original {
		return verify.MutantEquivalent
	}
	if t.timeout[p.Node] {
		return verify.MutantTimeout
	}
	from := t.killFrom[p.Node]
	for _, test := range tests {
		if from != "" && test >= from {
			return verify.MutantKilled
		}
	}
	return verify.MutantSurvived
}

// shaHex renders a byte slice as lowercase hex.
func shaHex(data []byte) string {
	return hex.EncodeToString(data)
}

// reportSHA is the SHA-256 digest used by the determinism checks.
func reportSHA(rep Report) [32]byte {
	return sha256.Sum256(rep.Bytes())
}

// flipContext cancels the run after the first Err call, exercising the loop guard.
type flipContext struct {
	calls int
}

func (*flipContext) Deadline() (time.Time, bool) { return time.Time{}, false }
func (*flipContext) Done() <-chan struct{}       { return nil }
func (*flipContext) Value(_ any) any             { return nil }
func (c *flipContext) Err() error {
	c.calls++
	if c.calls > 1 {
		return context.Canceled
	}
	return nil
}
