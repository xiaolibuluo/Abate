/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package derive

import (
	"testing"

	"abate/kernel/verify"
	"abate/lang/air"
	"abate/lang/rule"
	gobuiltin "abate/lang/rule/builtin/go"
)

// sampleSource is the in-memory Go source parsed into AIR for derivation tests.
const sampleSource = `package sample

func helper(value int) int { return value }

func compare(left int, right int) bool {
	if left < right {
		return true
	}
	if left == right {
		return false
	}
	return left >= 1
}

func loop(limit int) int {
	total := 0
	for i := 0; i < limit; i++ {
		if i == limit {
			break
		}
		total = total + 1
	}
	return helper(total)
}
`

// fixtureRule is a minimal rule.Rule used to exercise Derive in isolation.
type fixtureRule struct{ decl rule.Spec }

// Spec returns the declaration backing the fixture rule.
func (f fixtureRule) Spec() rule.Spec { return f.decl }

// Check returns no findings; the fixture exists only for derivation.
func (fixtureRule) Check(_ *air.Unit) []rule.Finding { return nil }

// rulesOf wraps declarations into the rule.Rule contract.
func rulesOf(specs ...rule.Spec) []rule.Rule {
	out := make([]rule.Rule, 0, len(specs))
	for _, spec := range specs {
		out = append(out, fixtureRule{decl: spec})
	}
	return out
}

// specFor builds one valid rule declaration with a chosen category and kind.
func specFor(id string, category rule.Category, kind string) rule.Spec {
	return rule.Spec{
		ID: rule.RuleID(id), Reference: "Effective Go, derivation fixture",
		Severity: rule.SeverityWarning, Category: category,
		Pattern: "Ident", Message: "derivation fixture",
		Mutation: rule.MutationDecl{Family: "M-DERIVED", Kind: kind},
	}
}

// sampleUnit parses the shared Go sample into a real AIR unit.
func sampleUnit(t *testing.T) *air.Unit {
	t.Helper()
	unit, err := gobuiltin.ParseUnit("derive/sample.go", []byte(sampleSource))
	if err != nil {
		t.Fatalf("parse sample: %v", err)
	}
	return unit
}

// builtinRules loads the 110 built-in Go rules in identifier order.
func builtinRules(t *testing.T) []rule.Rule {
	t.Helper()
	registry := rule.NewRegistry()
	if err := gobuiltin.LoadBuiltins(registry); err != nil {
		t.Fatalf("load built-ins: %v", err)
	}
	return registry.For("go")
}

// allocatedVerdict reports whether a verdict is one of the four kernel states.
func allocatedVerdict(value verify.MutationVerdict) bool {
	switch value {
	case verify.MutantKilled, verify.MutantSurvived,
		verify.MutantEquivalent, verify.MutantTimeout:
		return true
	default:
		return false
	}
}
