/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/
package gomutate

import (
	"abate/kernel/mutation"
	"abate/kernel/verify"
)

// targetAdapter projects the session core onto mutation.Target. It is a
// value type carrying the shared core pointer: the kernel stores the
// adapter while the session owns all mutable state behind it.
type targetAdapter struct {
	core *core
}

// compile-time proof of the port implementation.
var _ mutation.Target = targetAdapter{}

// Fingerprint returns the digest of the scanned token set.
func (a targetAdapter) Fingerprint() [32]byte {
	return a.core.tokenFP
}

// Nodes returns the kernel-facing node references in (path, offset) order.
func (a targetAdapter) Nodes() []string {
	out := make([]string, 0, len(a.core.tokens))
	for _, tok := range a.core.tokens {
		out = append(out, tok.Node())
	}
	return out
}

// Token returns the original lexical text of one node reference.
func (a targetAdapter) Token(node string) string {
	if tok, ok := a.core.byNode[node]; ok {
		return tok.Text
	}
	return ""
}

// Equivalent reports whether a mutant rewrites the token to its original
// text. That is the only statically decidable equivalence here; every other
// mutant must face the real test suite.
func (a targetAdapter) Equivalent(point mutation.MutationPoint) bool {
	return point.Mutated == point.Original
}

// Evaluate returns the cached outcome of one (point, test) pair, running
// the real toolchain on first use. A mutant that fails to compile or whose
// test run fails yields OutcomeViolated, which the kernel classifies as a
// kill. An unknown node or test yields OutcomeInconclusive instead of a
// fabricated verdict.
func (a targetAdapter) Evaluate(point mutation.MutationPoint, test string) verify.Outcome {
	return a.core.evaluate(point, test)
}

// suiteAdapter projects the session core onto mutation.Suite.
type suiteAdapter struct {
	core *core
}

// compile-time proof of the port implementation.
var _ mutation.Suite = suiteAdapter{}

// Fingerprint returns the digest of the discovered tests plus their
// measured coverage file set.
func (a suiteAdapter) Fingerprint() [32]byte {
	return a.core.testFP
}

// Tests returns the sorted fully qualified test identifiers.
func (a suiteAdapter) Tests() []string {
	out := make([]string, 0, len(a.core.tests))
	for _, test := range a.core.tests {
		out = append(out, test.ID)
	}
	return out
}

// Coverage returns how many discovered tests execute the line of one node,
// as measured by the per-test coverprofiles of the pristine workspace.
func (a suiteAdapter) Coverage(node string) uint64 {
	tok, ok := a.core.byNode[node]
	if !ok {
		return 0
	}
	return a.core.matrix.covering(tok.Path, tok.Line)
}
