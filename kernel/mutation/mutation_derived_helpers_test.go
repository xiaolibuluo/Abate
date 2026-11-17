/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package mutation

import (
	"testing"

	"abate/kernel/algebra"
)

// derivedOperator builds one well-formed derived M-family operator.
func derivedOperator(name, class, from, to string) algebra.Operator {
	return algebra.NewBasic(
		algebra.FormatOperatorID(algebra.FamilyMutation, "derived-"+name, "1"),
		algebra.FamilyMutation,
		algebra.Params{
			"class": class, "from": from, "to": to, "derived-from": name,
		},
		[]algebra.TargetKey{algebra.TargetKey(from)}, true)
}

// derivedBatch builds a deterministic batch of named derived operators.
func derivedBatch(count int, class, from, to string) []algebra.Operator {
	out := make([]algebra.Operator, 0, count)
	for i := 0; i < count; i++ {
		name := "rule-" + string(rune('a'+i%26)) + "-" + string(rune('a'+i/26))
		out = append(out, derivedOperator(name, class, from, to))
	}
	return out
}

// idsOf renders operator identifiers in their returned order.
func idsOf(ops []algebra.Operator) []string {
	out := make([]string, 0, len(ops))
	for _, op := range ops {
		out = append(out, string(op.ID()))
	}
	return out
}

// newRegistryWithNil builds the additive constructor with no extra operators.
func newRegistryWithNil(t *testing.T) *Registry {
	t.Helper()
	reg, err := NewRegistryWith(nil)
	if err != nil {
		t.Fatalf("NewRegistryWith(nil): %v", err)
	}
	return reg
}
