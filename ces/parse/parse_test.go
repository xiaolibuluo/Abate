/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package parse

import (
	"errors"
	"strings"
	"testing"

	"abate/ces/schema"
)

// TestParseValidDocuments decodes the canonical YAML and JSON twins.
func TestParseValidDocuments(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		name   string
		data   string
		format Format
	}{
		{"yaml-auto", baseYAML, FormatAuto},
		{"yaml-explicit", baseYAML, FormatYAML},
		{"json-auto", baseJSON, FormatAuto},
		{"json-explicit", baseJSON, FormatJSON},
	} {
		exp, err := Parse([]byte(tc.data), tc.format)
		requireOK(t, err)
		if exp.Metadata.Name != "sample" || exp.Spec.Seed != 12345 {
			t.Fatalf("%s: decoded fields drift: %+v", tc.name, exp.Spec)
		}
		if _, err := CompileExpr(exp.Spec.Method.Expression, exp.Spec.Method.Operators); err != nil {
			t.Fatalf("%s: expression not compilable: %v", tc.name, err)
		}
	}
}

// TestSniffFormat verifies first-non-space-byte detection.
func TestSniffFormat(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		in   string
		want Format
	}{
		{"{}", FormatJSON},
		{"  \n\t {\"a\":1}", FormatJSON},
		{"apiVersion: x", FormatYAML},
		{"   ", FormatYAML},
		{"", FormatYAML},
	} {
		if got := sniffFormat([]byte(tc.in)); got != tc.want {
			t.Fatalf("sniffFormat(%q) = %q, want %q", tc.in, got, tc.want)
		}
	}
}

// TestParseEmptyAndOversize covers the two size guards.
func TestParseEmptyAndOversize(t *testing.T) {
	t.Parallel()
	_, err := Parse(nil, FormatAuto)
	requireRule(t, err, RuleEmptyDocument)
	_, err = Parse([]byte{}, FormatJSON)
	requireRule(t, err, RuleEmptyDocument)
	_, err = Parse(make([]byte, MaxDocumentBytes+1), FormatAuto)
	requireRule(t, err, RuleInputTooLarge)
}

// TestParseExplicitFormatRejection covers the default switch branch.
func TestParseExplicitFormatRejection(t *testing.T) {
	t.Parallel()
	_, err := Parse([]byte(baseYAML), Format("xml"))
	requireRule(t, err, RuleInvalidFormat)
}

// TestParseStrictUnknownFields proves both strict decoders reject typos.
func TestParseStrictUnknownFields(t *testing.T) {
	t.Parallel()
	j := strings.Replace(baseJSON, `"kind": "Experiment"`,
		`"kind": "Experiment", "unexpected": 1`, 1)
	_, err := Parse([]byte(j), FormatJSON)
	requireRule(t, err, RuleUnknownField)
	y := baseYAML + "unexpectedTop: 1\n"
	_, err = Parse([]byte(y), FormatYAML)
	requireRule(t, err, RuleUnknownField)
}

// TestParseMalformedDocuments covers the generic decode-failure path.
func TestParseMalformedDocuments(t *testing.T) {
	t.Parallel()
	_, err := Parse([]byte(`{"apiVersion":`), FormatJSON)
	requireRule(t, err, RuleMalformedDocument)
	_, err = Parse([]byte("spec: [unterminated\n"), FormatYAML)
	requireRule(t, err, RuleMalformedDocument)
}

// TestDecodeFailureClassification drives the message classifier directly.
func TestDecodeFailureClassification(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		msg  string
		want string
	}{
		{`json: unknown field "foo"`, RuleUnknownField},
		{"yaml: field foo not found in type schema.Experiment", RuleUnknownField},
		{"unexpected end of JSON input", RuleMalformedDocument},
	} {
		if got := decodeFailure(errors.New(tc.msg)).RuleID; got != tc.want {
			t.Fatalf("decodeFailure(%q) = %q, want %q", tc.msg, got, tc.want)
		}
	}
}

// TestCheckEnvelopeBranches drives every envelope validation branch.
func TestCheckEnvelopeBranches(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		name string
		mut  func(*schema.Experiment)
		want string
	}{
		{"missing-version", func(e *schema.Experiment) { e.APIVersion = "" }, RuleMissingField},
		{"unsupported-version", func(e *schema.Experiment) { e.APIVersion = "ces/v9" }, RuleUnsupportedVersion},
		{"wrong-kind", func(e *schema.Experiment) { e.Kind = "Pod" }, RuleMalformedDocument},
		{"missing-name", func(e *schema.Experiment) { e.Metadata.Name = "" }, RuleMissingField},
	} {
		exp := validExperiment(t)
		tc.mut(exp)
		err := checkEnvelope(exp)
		requireRule(t, err, tc.want)
	}
	if err := checkEnvelope(validExperiment(t)); err != nil {
		t.Fatalf("valid envelope rejected: %v", err)
	}
}

// TestParseErrorRendering checks deterministic, address-free diagnostics.
func TestParseErrorRendering(t *testing.T) {
	t.Parallel()
	doc := newError(RuleMissingField, "$.metadata.name", "name is required")
	if doc.Line != 1 || doc.Column != 1 || doc.FieldPath != "$.metadata.name" {
		t.Fatalf("document error position drift: %+v", doc)
	}
	want := "CES-PARSE-006 $.metadata.name:1:1 name is required"
	if doc.Error() != want {
		t.Fatalf("Error() = %q, want %q", doc.Error(), want)
	}
	expr := newExprError(RuleExprSyntax, "bad token", 7)
	if expr.Column != 8 || expr.FieldPath != "$.spec.method.expression" {
		t.Fatalf("expression error position drift: %+v", expr)
	}
	wantExpr := "CES-PARSE-009 $.spec.method.expression:1:8 bad token"
	if expr.Error() != wantExpr {
		t.Fatalf("Error() = %q, want %q", expr.Error(), wantExpr)
	}
}
