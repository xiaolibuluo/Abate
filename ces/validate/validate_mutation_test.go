/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package validate

import (
	"errors"
	"slices"
	"strings"
	"testing"

	"abate/ces/parse"
	"abate/ces/schema"
	"abate/kernel/algebra"
)

// TestMutationFieldFlips applies one single-value field mutation at a time to
// the legal baseline and requires the corresponding rule to fire through both
// the isolated rule and the aggregate validator.
func TestMutationFieldFlips(t *testing.T) {
	t.Parallel()
	for _, c := range []struct {
		name string
		mut  func(*schema.Experiment)
		want string
	}{
		{"method-expression", func(e *schema.Experiment) {
			e.Spec.Method.Expression = "@"
		}, RuleMethod001},
		{"method-depth", func(e *schema.Experiment) {
			e.Spec.Method.Expression = strings.Repeat("(", 70) + "ε" + strings.Repeat(")", 70)
		}, RuleMethod003},
		{"method-conflict", func(e *schema.Experiment) {
			e.Spec.Method.Expression = "a ⊗ a"
		}, RuleMethod004},
		{"steady-holdfor", func(e *schema.Experiment) {
			e.Spec.SteadyState[0].HoldFor = 0
		}, RuleSteady003},
		{"probe-interval", func(e *schema.Experiment) {
			e.Spec.Probes[0].Interval = 0
		}, RuleProbe002},
		{"probe-url", func(e *schema.Experiment) {
			e.Spec.Probes[0].Spec.URL = "https://evil.example/x"
		}, RuleProbe003},
		{"probe-status", func(e *schema.Experiment) {
			e.Spec.Probes[0].Spec.ExpectStatus = 700
		}, RuleProbe005},
		{"rollback-strategy", func(e *schema.Experiment) {
			e.Spec.Rollback.Strategy = "quantum"
		}, RuleRollback001},
		{"target-kind", func(e *schema.Experiment) {
			e.Spec.Target.Kind = "quantum"
		}, RuleTarget001},
		{"target-selector", func(e *schema.Experiment) {
			e.Spec.Target.Selector = nil
		}, RuleTarget003},
		{"schedule-once", func(e *schema.Experiment) {
			e.Spec.Schedule.Once.At = 0
		}, RuleSchedule004},
	} {
		t.Run(c.name, func(t *testing.T) {
			exp := fresh(t)
			c.mut(exp)
			assertRule(t, ruleViolations(t, c.want, exp), c.want)
			vs := mustValidate(t, exp)
			assertRule(t, vs, c.want)
		})
	}
}

// TestMutationSingleCharacterFlips flips one character of the baseline's
// expression, strategy enum and probe URL, requiring the rules to detect them.
func TestMutationSingleCharacterFlips(t *testing.T) {
	t.Parallel()
	for _, bad := range []string{"@", "a @ b", "(", "a⁻", "operator://"} {
		exp := fresh(t)
		exp.Spec.Method.Expression = bad
		assertRule(t, ruleViolations(t, RuleMethod001, exp), RuleMethod001)
	}
	const strategy = "auto"
	found := 0
	for i := 0; i < len(strategy); i++ {
		exp := fresh(t)
		exp.Spec.Rollback.Strategy = schema.Strategy(strategy[:i] + "X" + strategy[i+1:])
		if hasRule(ruleViolations(t, RuleRollback001, exp), RuleRollback001) {
			found++
		}
	}
	if found != len(strategy) {
		t.Fatalf("strategy character flips detected %d/%d", found, len(strategy))
	}
	const url = "http://sim/health"
	urlFound := 0
	for i := 0; i < len(url); i++ {
		exp := fresh(t)
		exp.Spec.Probes[0].Spec.URL = url[:i] + "X" + url[i+1:]
		if hasRule(ruleViolations(t, RuleProbe003, exp), RuleProbe003) {
			urlFound++
		}
	}
	if urlFound == 0 {
		t.Fatal("no single-character URL flip was detected")
	}
}

// TestMutationReferences exercises operator:// reference extraction.
func TestMutationReferences(t *testing.T) {
	t.Parallel()
	for _, c := range []struct {
		in   string
		want []string
	}{
		{"", nil},
		{"ε", nil},
		{"operator://a", []string{"a"}},
		{"operator://a ∘ operator://b", []string{"a", "b"}},
		{"operator://a⊗operator://b", []string{"a", "b"}},
		{"operator://m-1_v2.x", []string{"m-1_v2.x"}},
		{"operator://", nil},
		{"operator://a/operator://b", []string{"a", "b"}},
		{"operator://a ∘ operator://a", []string{"a", "a"}},
	} {
		if got := referencedIDs(c.in); !slices.Equal(got, c.want) {
			t.Fatalf("referencedIDs(%q) = %v, want %v", c.in, got, c.want)
		}
	}
}

// TestMutationExpressionCompile covers the compiler helper and every
// deterministic error classification branch of compile.go.
func TestMutationExpressionCompile(t *testing.T) {
	t.Parallel()
	exp := fresh(t)
	op, err := compileMethod(exp)
	ok(t, err)
	if op == nil {
		t.Fatal("compileMethod returned a nil operator")
	}
	if !expressionHasProduct("a ⊗ b") || expressionHasProduct("a ∘ b") {
		t.Fatal("concurrent product detection disagrees with the expression")
	}
	for _, c := range []struct {
		name     string
		err      error
		depth    bool
		conflict bool
		msg      string
	}{
		{"nil-error", nil, false, false, ""},
		{"plain-error", errors.New("boom"), false, false, "boom"},
		{"depth-rule", &parse.ParseError{
			RuleID: parse.RuleExprDepth, Message: "too deep",
		}, true, false, "too deep"},
		{"depth-text", &parse.ParseError{
			Message: algebra.ErrTooDeep.Error(),
		}, true, false, algebra.ErrTooDeep.Error()},
		{"target-conflict", &parse.ParseError{
			Message: algebra.ErrTargetConflict.Error(),
		}, false, true, algebra.ErrTargetConflict.Error()},
		{"incompatible", &parse.ParseError{
			Message: algebra.ErrIncompatible.Error(),
		}, false, true, algebra.ErrIncompatible.Error()},
		{"generic-parse", &parse.ParseError{
			RuleID: parse.RuleExprSyntax, Message: "syntax",
		}, false, false, "syntax"},
	} {
		depth, conflict, msg := compileFailure(c.err)
		if depth != c.depth || conflict != c.conflict || msg != c.msg {
			t.Fatalf("compileFailure(%s) = (%v,%v,%q), want (%v,%v,%q)",
				c.name, depth, conflict, msg, c.depth, c.conflict, c.msg)
		}
	}
}

// TestMutationToleranceMatrix covers the expectation/tolerance type matrix
// and the known-type predicates.
func TestMutationToleranceMatrix(t *testing.T) {
	t.Parallel()
	for _, c := range []struct {
		exp  string
		tol  string
		want bool
	}{
		{schema.ExpectValue, schema.ToleranceAbsolute, true},
		{schema.ExpectValue, schema.ToleranceRelative, true},
		{schema.ExpectValue, schema.TolerancePercent, true},
		{schema.ExpectValue, "", false},
		{schema.ExpectValue, "quantum", false},
		{schema.ExpectRange, schema.ToleranceAbsolute, true},
		{schema.ExpectRange, schema.ToleranceRelative, false},
		{schema.ExpectRange, "", false},
		{schema.ExpectPredicate, "", true},
		{schema.ExpectPredicate, schema.ToleranceAbsolute, false},
		{"other", schema.ToleranceAbsolute, false},
	} {
		if got := toleranceAllowed(c.exp, c.tol); got != c.want {
			t.Fatalf("toleranceAllowed(%q,%q) = %v, want %v",
				c.exp, c.tol, got, c.want)
		}
	}
	for _, e := range []string{
		schema.ExpectValue, schema.ExpectRange, schema.ExpectPredicate,
	} {
		if !knownExpectation(e) {
			t.Fatalf("knownExpectation(%q) = false", e)
		}
	}
	if knownExpectation("other") {
		t.Fatal("unknown expectation accepted")
	}
	for _, tol := range []string{
		schema.ToleranceAbsolute, schema.ToleranceRelative, schema.TolerancePercent,
	} {
		if !knownTolerance(tol) {
			t.Fatalf("knownTolerance(%q) = false", tol)
		}
	}
	if knownTolerance("") || knownTolerance("other") {
		t.Fatal("unknown tolerance accepted")
	}
}

// TestMutationRegistryBranches covers the nil, duplicate and successful
// registration paths plus deterministic id ordering and aggregation.
func TestMutationRegistryBranches(t *testing.T) {
	t.Parallel()
	reg := NewRegistry()
	if err := reg.Register(nil); !errors.Is(err, ErrNilRule) {
		t.Fatalf("Register(nil) = %v, want ErrNilRule", err)
	}
	rule := newRule("X", func(*schema.Experiment) []Violation { return nil })
	ok(t, reg.Register(rule))
	if err := reg.Register(rule); !errors.Is(err, ErrDuplicateRule) {
		t.Fatalf("duplicate Register = %v, want ErrDuplicateRule", err)
	}
	if !slices.Equal(reg.IDs(), []string{"X"}) {
		t.Fatalf("single registry ids = %v", reg.IDs())
	}
	ordered := NewRegistry()
	ok(t, ordered.Register(newRule("B", func(*schema.Experiment) []Violation {
		return []Violation{errorAt("B", "p", "m")}
	})))
	ok(t, ordered.Register(newRule("A", func(*schema.Experiment) []Violation {
		return []Violation{warningAt("A", "p", "m")}
	})))
	if !slices.Equal(ordered.IDs(), []string{"A", "B"}) {
		t.Fatalf("registry ids not sorted: %v", ordered.IDs())
	}
	vs := ordered.run(fresh(t))
	if len(vs) != 2 || vs[0].RuleID != "A" || vs[0].Severity != SeverityWarning ||
		vs[1].RuleID != "B" || vs[1].Severity != SeverityError {
		t.Fatalf("registry aggregate = %+v, want A(warning) then B(error)", vs)
	}
}
