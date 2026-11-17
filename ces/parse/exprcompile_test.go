/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package parse

import (
	"strings"
	"testing"

	"abate/ces/schema"
	"abate/kernel/algebra"
)

// declaredOps is the operator table used to exercise reference resolution.
func declaredOps() []schema.OperatorDef {
	return []schema.OperatorDef{
		{ID: "c-delay-v1", Family: schema.FamilyChaos,
			Params: map[string]string{"ms": "100"}, Targets: []string{"net"}, Invertible: true},
		{ID: "m-repl-v1", Family: schema.FamilyMutation,
			Targets: []string{"src"}, Invertible: true},
		{ID: "c-sticky-v1", Family: schema.FamilyChaos,
			Targets: []string{"disk"}, Invertible: false},
		{ID: "c-weird-v1", Family: "unknown-family", Targets: []string{"x"}},
	}
}

// TestCompileExprSuccess covers every node kind and produces a stable tree.
func TestCompileExprSuccess(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		in      string
		wantFam algebra.Family
	}{
		{"ε", algebra.FamilyIdentity},
		{"operator://c-delay-v1", algebra.FamilyChaos},
		{"operator://m-repl-v1", algebra.FamilyMutation},
		{"m-inline-v1", algebra.FamilyMutation},
		{"c-inline-v1", algebra.FamilyChaos},
		{"a ∘ b", algebra.FamilyChaos},
		{"a ⊗ b", algebra.FamilyChaos},
		{"operator://c-delay-v1⁻¹", algebra.FamilyChaos},
	} {
		op, err := CompileExpr(tc.in, declaredOps())
		requireOK(t, err)
		if op.Family() != tc.wantFam {
			t.Fatalf("CompileExpr(%q) family = %q, want %q", tc.in, op.Family(), tc.wantFam)
		}
		if op.Fingerprint() == [32]byte{} {
			t.Fatalf("CompileExpr(%q) produced an empty fingerprint", tc.in)
		}
	}
}

// TestCompileExprLexAndParseErrors covers the two front-end failures.
func TestCompileExprLexAndParseErrors(t *testing.T) {
	t.Parallel()
	_, err := CompileExpr("a @ b", nil)
	requireRule(t, err, RuleExprSyntax)
	_, err = CompileExpr("", nil)
	requireRule(t, err, RuleExprSyntax)
	_, err = CompileExpr("a b", nil)
	requireRule(t, err, RuleExprSyntax)
}

// TestCompileExprUnknownRef rejects a reference absent from the table.
func TestCompileExprUnknownRef(t *testing.T) {
	t.Parallel()
	_, err := CompileExpr("operator://ghost-v1", declaredOps())
	requireRule(t, err, RuleExprUnknownOp)
}

// TestCompileExprUnknownFamily rejects an unsupported declared family.
func TestCompileExprUnknownFamily(t *testing.T) {
	t.Parallel()
	_, err := CompileExpr("operator://c-weird-v1", declaredOps())
	requireRule(t, err, RuleExprSyntax)
}

// TestCompileExprNotInvertible rejects a dual of a non-invertible operator.
func TestCompileExprNotInvertible(t *testing.T) {
	t.Parallel()
	_, err := CompileExpr("operator://c-sticky-v1⁻¹", declaredOps())
	requireRule(t, err, RuleExprNotInvertible)
}

// TestCompileExprConcurrentConflict maps a join failure to a syntax error.
func TestCompileExprConcurrentConflict(t *testing.T) {
	t.Parallel()
	// Two operands sharing a target are rejected by algebra.Concurrent.
	_, err := CompileExpr("a ⊗ a", nil)
	requireRule(t, err, RuleExprSyntax)
}

// TestCompileNodeInternalBranches drives defensive branches directly.
func TestCompileNodeInternalBranches(t *testing.T) {
	t.Parallel()
	_, err := compileNode(&exprNode{kind: nodeOperand, text: "a"}, nil, algebra.MaxDepth+1)
	requireRule(t, err, RuleExprDepth)
	_, err = compileNode(&exprNode{kind: nodeKind(99)}, nil, 0)
	requireRule(t, err, RuleExprSyntax)
}

// TestCompileNodeInverseChildFailure propagates an inner compile error.
func TestCompileNodeInverseChildFailure(t *testing.T) {
	t.Parallel()
	_, err := compileNode(&exprNode{
		kind: nodeInverse, left: &exprNode{kind: nodeKind(99)},
	}, nil, 0)
	if ruleOf(err) != RuleExprSyntax {
		t.Fatalf("inverse child failure not propagated: %v", err)
	}
}

// TestCombineErrorPropagation covers left and right child failures.
func TestCombineErrorPropagation(t *testing.T) {
	t.Parallel()
	bad := &exprNode{kind: nodeKind(99)}
	good := &exprNode{kind: nodeOperand, text: "a"}
	_, err := combine(&exprNode{kind: nodeCompose, left: bad, right: good}, nil, 0, algebra.Compose)
	if ruleOf(err) != RuleExprSyntax {
		t.Fatalf("left-child failure not propagated: %v", err)
	}
	_, err = combine(&exprNode{kind: nodeCompose, left: good, right: bad}, nil, 0, algebra.Compose)
	if ruleOf(err) != RuleExprSyntax {
		t.Fatalf("right-child failure not propagated: %v", err)
	}
}

// TestResolveOperandBranches covers declared, undeclared ref and inline ids.
func TestResolveOperandBranches(t *testing.T) {
	t.Parallel()
	op, err := resolveRefOperand("c-delay-v1", declaredOps(), 0)
	requireOK(t, err)
	if op.Family() != algebra.FamilyChaos || len(op.Targets()) != 1 {
		t.Fatalf("declared operand drift: %+v", op)
	}
	_, err = resolveRefOperand("ghost", declaredOps(), 0)
	requireRule(t, err, RuleExprUnknownOp)
	m, err := resolveOperand("m-x", nil, 0)
	requireOK(t, err)
	if m.Family() != algebra.FamilyMutation {
		t.Fatalf("inline m- id family = %q", m.Family())
	}
	c, err := resolveOperand("c-x", nil, 0)
	requireOK(t, err)
	if c.Family() != algebra.FamilyChaos {
		t.Fatalf("inline c- id family = %q", c.Family())
	}
}

// TestBuildOperatorBranches covers both families and the unknown default.
func TestBuildOperatorBranches(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		family string
		want   algebra.Family
	}{
		{schema.FamilyMutation, algebra.FamilyMutation},
		{schema.FamilyChaos, algebra.FamilyChaos},
	} {
		op, err := buildOperator(schema.OperatorDef{
			ID: "op", Family: tc.family, Params: map[string]string{"k": "v"},
			Targets: []string{"t1", "t2"}, Invertible: true,
		}, 0)
		requireOK(t, err)
		if op.Family() != tc.want || len(op.Targets()) != 2 || op.Params()["k"] != "v" {
			t.Fatalf("buildOperator drift: %+v", op)
		}
	}
	_, err := buildOperator(schema.OperatorDef{ID: "x", Family: "quantum"}, 0)
	requireRule(t, err, RuleExprSyntax)
}

// TestCompileExprDepthGuard verifies a left-deep chain is bounded.
func TestCompileExprDepthGuard(t *testing.T) {
	t.Parallel()
	long := strings.Repeat("a ∘ ", algebra.MaxDepth+2) + "a"
	_, err := CompileExpr(long, nil)
	if err == nil {
		t.Fatal("an over-deep operator chain must be rejected")
	}
}
