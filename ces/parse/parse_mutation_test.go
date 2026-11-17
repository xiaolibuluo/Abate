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
)

// TestMutationEnvelopeFields proves single-character mutations of the document
// envelope are detected with the correct rule id.
func TestMutationEnvelopeFields(t *testing.T) {
	t.Parallel()
	version := schema.APIVersion
	for i := 0; i < len(version); i++ {
		mut := version[:i] + "X" + version[i+1:]
		if mut == version {
			continue
		}
		doc := strings.Replace(baseJSON, version, mut, 1)
		_, err := Parse([]byte(doc), FormatJSON)
		requireRule(t, err, RuleUnsupportedVersion)
	}
	kind := schema.KindExperiment
	for i := 0; i < len(kind); i++ {
		mut := kind[:i] + "X" + kind[i+1:]
		if mut == kind {
			continue
		}
		doc := strings.Replace(baseJSON, `"kind": "`+kind+`"`, `"kind": "`+mut+`"`, 1)
		_, err := Parse([]byte(doc), FormatJSON)
		requireRule(t, err, RuleMalformedDocument)
	}
}

// TestMutationUnknownFieldInjection detects a tampered-in field for both codecs.
func TestMutationUnknownFieldInjection(t *testing.T) {
	t.Parallel()
	j := strings.Replace(baseJSON, `"seed": 12345`, `"seed": 12345, "injected": 1`, 1)
	_, err := Parse([]byte(j), FormatJSON)
	requireRule(t, err, RuleUnknownField)
	y := strings.Replace(baseYAML, "  seed: 12345", "  seed: 12345\n  injected: 1", 1)
	_, err = Parse([]byte(y), FormatYAML)
	requireRule(t, err, RuleUnknownField)
}

// TestMutationExpressionOperators verifies changing an operator symbol changes
// the compiled tree or is rejected, so the compiler cannot silently ignore it.
func TestMutationExpressionOperators(t *testing.T) {
	t.Parallel()
	compose, err := CompileExpr("a ∘ b", nil)
	requireOK(t, err)
	product, err := CompileExpr("a ⊗ b", nil)
	requireOK(t, err)
	if compose.Fingerprint() == product.Fingerprint() {
		t.Fatal("operator mutation ∘ -> ⊗ was not detected in the tree")
	}
	dual, err := CompileExpr("a⁻¹", nil)
	requireOK(t, err)
	baseOp, err := CompileExpr("a", nil)
	requireOK(t, err)
	if dual.Fingerprint() == baseOp.Fingerprint() {
		t.Fatal("operator mutation adding a dual was not detected")
	}
}

// TestMutationExpressionTokenRejection detects corrupted operator tokens.
func TestMutationExpressionTokenRejection(t *testing.T) {
	t.Parallel()
	for _, expr := range []string{"a ∘ ", "∘ a", "a ⊗ ⊗ b", "operator://", "a @ b"} {
		if _, err := CompileExpr(expr, nil); err == nil {
			t.Fatalf("mutated expression %q was not detected", expr)
		}
	}
}

// TestMutationByteFlipSensitivity counts how many single-byte mutations of a
// valid document are detected by the parser (rejected or content-changed).
func TestMutationByteFlipSensitivity(t *testing.T) {
	t.Parallel()
	base, err := Parse([]byte(baseJSON), FormatJSON)
	requireOK(t, err)
	baseDigest := digestOf(t, base)
	detected := 0
	total := 0
	for i := 0; i < len(baseJSON); i += 7 {
		if baseJSON[i] < 'a' || baseJSON[i] > 'z' {
			continue
		}
		mutated := baseJSON[:i] + "Z" + baseJSON[i+1:]
		if mutated == baseJSON {
			continue
		}
		total++
		got, err := Parse([]byte(mutated), FormatJSON)
		if err != nil {
			detected++
			continue
		}
		if digestOf(t, got) != baseDigest {
			detected++
		}
	}
	if total == 0 || detected*2 < total {
		t.Fatalf("mutation sensitivity too low: %d/%d detected", detected, total)
	}
}

// TestMutationNormalizeFields detects tampering with cross-field invariants.
func TestMutationNormalizeFields(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		name string
		mut  func(*schema.Experiment)
		want string
	}{
		{"dangling-ref", func(e *schema.Experiment) {
			e.Spec.SteadyState[0].ProbeRef = "ghost"
		}, RuleDanglingProbeRef},
		{"extra-schedule", func(e *schema.Experiment) {
			e.Spec.Schedule.Cron = &schema.CronSpec{Expression: "* * * * *"}
		}, RuleScheduleExclusive},
		{"duplicate-probe", func(e *schema.Experiment) {
			e.Spec.Probes = append(e.Spec.Probes, e.Spec.Probes[0])
		}, RuleMalformedDocument},
	} {
		exp := validExperiment(t)
		tc.mut(exp)
		_, err := Normalize(exp)
		requireRule(t, err, tc.want)
	}
}
