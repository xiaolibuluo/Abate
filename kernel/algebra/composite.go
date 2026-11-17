/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package algebra

import (
	"crypto/sha256"
	"fmt"
	"slices"
)

// compositeOp is the lazy wrapper produced by Compose.
//
// It applies the inner operator first and then the outer operator. Its
// fingerprint flattens the application-order leaf sequence so that parenthesization
// is irrelevant (composition is associative).
type compositeOp struct {
	outer Operator
	inner Operator
}

// Compose assembles p1 ∘ p2, meaning p2 is applied first and then p1.
//
// The identity operator is a two-sided unit, so composing with it returns the
// other operand directly. Composing two non-invertible operators that share a
// target is rejected with ErrIncompatible, because neither can be rolled back.
func Compose(p1, p2 Operator) (Operator, error) {
	if p1 == nil || p2 == nil {
		return nil, ErrNilOperator
	}
	if p2.Family() == FamilyIdentity {
		return p1, nil
	}
	if p1.Family() == FamilyIdentity {
		return p2, nil
	}
	if depthOf(p1)+depthOf(p2)+1 > MaxDepth {
		return nil, ErrTooDeep
	}
	if intersect(targetsOf(p1), targetsOf(p2)) && !p1.Invertible() && !p2.Invertible() {
		return nil, ErrIncompatible
	}
	return &compositeOp{outer: p1, inner: p2}, nil
}

// ID returns the composite operator identifier.
func (*compositeOp) ID() OperatorID {
	return FormatOperatorID(FamilyMixed, "compose", "v1")
}

// Family derives the composite family from its leaves.
func (c *compositeOp) Family() Family {
	return deriveFamily(flattenFamilies(c))
}

// Params returns the canonical parameters of the composite.
func (c *compositeOp) Params() Params {
	return Params{
		"outer": fmt.Sprintf("%x", c.outer.Fingerprint()),
		"inner": fmt.Sprintf("%x", c.inner.Fingerprint()),
	}
}

// Targets returns the union of the operand target keys.
func (c *compositeOp) Targets() []TargetKey {
	return targetsUnion([]Operator{c.inner, c.outer})
}

// Invertible reports whether both operands are invertible.
func (c *compositeOp) Invertible() bool {
	return c.inner.Invertible() && c.outer.Invertible()
}

// flattenFamilies collects the family of every leaf operator exactly once.
//
// It recurses structurally rather than through Operator.Family, so it is linear
// in the tree size and never re-traverses subtrees.
func flattenFamilies(op Operator) []Family {
	switch n := op.(type) {
	case *compositeOp:
		return append(flattenFamilies(n.inner), flattenFamilies(n.outer)...)
	case *concurrentOp:
		return append(flattenFamilies(n.a), flattenFamilies(n.b)...)
	default:
		return []Family{op.Family()}
	}
}

// flattenTargets collects every target key of all leaves exactly once.
func flattenTargets(op Operator) []TargetKey {
	switch n := op.(type) {
	case *compositeOp:
		return append(flattenTargets(n.inner), flattenTargets(n.outer)...)
	case *concurrentOp:
		return append(flattenTargets(n.a), flattenTargets(n.b)...)
	default:
		return op.Targets()
	}
}

// unionTargets deduplicates and sorts a flat target-key list.
func unionTargets(flat []TargetKey) []TargetKey {
	set := make(map[TargetKey]struct{}, len(flat))
	for _, tk := range flat {
		set[tk] = struct{}{}
	}
	out := make([]TargetKey, 0, len(set))
	for tk := range set {
		out = append(out, tk)
	}
	slices.Sort(out)
	return out
}

// Inverse returns the dual composite (reverse order of duals).
func (c *compositeOp) Inverse() (Operator, error) {
	innerInv, err := Inverse(c.inner)
	if err != nil {
		return nil, err
	}
	outerInv, err := Inverse(c.outer)
	if err != nil {
		return nil, err
	}
	return Compose(outerInv, innerInv)
}

// Fingerprint returns the application-order flattened fingerprint.
func (c *compositeOp) Fingerprint() [32]byte {
	return seqFingerprint(flattenLeaves(c))
}

// flattenLeaves expands an operator into its application-order leaf sequence.
func flattenLeaves(op Operator) []Operator {
	switch n := op.(type) {
	case *compositeOp:
		return append(flattenLeaves(n.inner), flattenLeaves(n.outer)...)
	case *concurrentOp:
		return append(flattenLeaves(n.a), flattenLeaves(n.b)...)
	default:
		return []Operator{op}
	}
}

// seqFingerprint hashes an ordered leaf sequence into a single fingerprint.
func seqFingerprint(ops []Operator) [32]byte {
	h := sha256.New()
	for _, o := range ops {
		fp := o.Fingerprint()
		_, _ = h.Write(fp[:])
	}
	var out [32]byte
	copy(out[:], h.Sum(nil))
	return out
}

// deriveFamily collapses a leaf family list into a single family label.
func deriveFamily(fams []Family) Family {
	seen := make(map[Family]bool)
	for _, f := range fams {
		if f == FamilyIdentity {
			continue
		}
		seen[f] = true
	}
	if len(seen) == 0 {
		return FamilyIdentity
	}
	if len(seen) == 1 {
		for f := range seen {
			return f
		}
	}
	return FamilyMixed
}

// targetsOf returns the target set of an operator.
func targetsOf(op Operator) []TargetKey {
	return op.Targets()
}

// targetsUnion merges operator target sets, deduplicating and sorting.
//
// It uses the linear flatten so deep composite trees are traversed once.
func targetsUnion(ops []Operator) []TargetKey {
	var flat []TargetKey
	for _, o := range ops {
		flat = append(flat, flattenTargets(o)...)
	}
	return unionTargets(flat)
}

// intersect reports whether two target sets share any key.
func intersect(a, b []TargetKey) bool {
	set := make(map[TargetKey]struct{}, len(a))
	for _, tk := range a {
		set[tk] = struct{}{}
	}
	for _, tk := range b {
		if _, ok := set[tk]; ok {
			return true
		}
	}
	return false
}

// targetsEqual reports whether two target sets are identical.
func targetsEqual(a, b Operator) bool {
	ta := targetsUnion([]Operator{a})
	tb := targetsUnion([]Operator{b})
	if len(ta) != len(tb) {
		return false
	}
	for i := range ta {
		if ta[i] != tb[i] {
			return false
		}
	}
	return true
}

// isIdentity reports whether an operator is the identity.
func isIdentity(op Operator) bool {
	return op.Family() == FamilyIdentity
}

// depthOf computes the nested depth of an operator tree.
func depthOf(op Operator) int {
	switch n := op.(type) {
	case *compositeOp:
		return 1 + max(depthOf(n.inner), depthOf(n.outer))
	case *concurrentOp:
		return 1 + max(depthOf(n.a), depthOf(n.b))
	default:
		return 1
	}
}
