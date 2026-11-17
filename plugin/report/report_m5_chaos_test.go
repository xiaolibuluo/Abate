/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package report

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"strings"
	"sync"
	"testing"

	"abate/internal/determinism"
	"abate/kernel/mutation"
	"abate/kernel/verify"
	"abate/lang/rule"
)

func TestM5ChaosShuffledFindings(t *testing.T) {
	findings := sampleFindings()
	rep := m5VerdictReport(verify.MutantKilled, verify.MutantSurvived, verify.MutantTimeout)
	baseHTML := m5MustHTML(t, findings, rep, DefaultHTMLOptions())
	baseJUnit := m5MustJUnit(t, findings, rep)
	baseOTel := m5MustOTel(t, findings, rep)
	for shift := 0; shift < len(findings); shift++ {
		rotated := make([]rule.Finding, 0, len(findings))
		for i := range findings {
			rotated = append(rotated, findings[(i+shift)%len(findings)])
		}
		if !bytes.Equal(baseHTML, m5MustHTML(t, rotated, rep, DefaultHTMLOptions())) ||
			!bytes.Equal(baseJUnit, m5MustJUnit(t, rotated, rep)) ||
			!bytes.Equal(baseOTel, m5MustOTel(t, rotated, rep)) {
			t.Fatalf("rotation %d changed the rendered bytes", shift)
		}
	}
}

func TestM5ChaosHugeMessage(t *testing.T) {
	huge := strings.Repeat(`<x>&"'`+"abcdef", 20000)
	findings := []rule.Finding{{
		RuleID: "GO-EG-001", Path: "huge.go", Line: 1, Column: 1,
		Severity: rule.SeverityError, Category: rule.CategoryStyle,
		Message: huge, Reference: "EG-001",
	}}
	rep := sampleReport()
	html := m5MustHTML(t, findings, rep, DefaultHTMLOptions())
	junit := m5MustJUnit(t, findings, rep)
	otel := m5MustOTel(t, findings, rep)
	if len(html) < len(huge) || len(junit) < len(huge) || len(otel) < len(huge) {
		t.Fatal("the huge message was silently dropped")
	}
	if !bytes.Equal(html, m5MustHTML(t, findings, rep, DefaultHTMLOptions())) ||
		!bytes.Equal(junit, m5MustJUnit(t, findings, rep)) ||
		!bytes.Equal(otel, m5MustOTel(t, findings, rep)) {
		t.Fatal("a huge message is not byte-stable")
	}
	determinism.MustMatch100(t, 9, func() []byte {
		return m5MustOTel(t, findings, rep)
	})
}

func TestM5ChaosUnicodePaths(t *testing.T) {
	paths := []string{"路径/日本語/Ω/🙂.go", "café/ naïve .go", "emoji-💥-file.go", `quote-"'<&.go`}
	for _, path := range paths {
		findings := []rule.Finding{{
			RuleID: "GO-EG-001", Path: path, Line: 7, Column: 3,
			Severity: rule.SeverityWarning, Category: rule.CategoryNaming,
			Message: "unicode ✓ path", Reference: "EG-001",
		}}
		rep := sampleReport()
		html := m5MustHTML(t, findings, rep, DefaultHTMLOptions())
		junit := m5MustJUnit(t, findings, rep)
		otel := m5MustOTel(t, findings, rep)
		if !bytes.Equal(html, m5MustHTML(t, findings, rep, DefaultHTMLOptions())) ||
			!bytes.Equal(junit, m5MustJUnit(t, findings, rep)) ||
			!bytes.Equal(otel, m5MustOTel(t, findings, rep)) {
			t.Fatalf("Unicode path %q is not byte-stable", path)
		}
		var suites junitDocument
		if err := xml.Unmarshal(junit, &suites); err != nil {
			t.Fatalf("Unicode path %q broke the XML document: %v", path, err)
		}
		var doc otelDocument
		if err := json.Unmarshal(otel, &doc); err != nil {
			t.Fatalf("Unicode path %q broke the JSON payload: %v", path, err)
		}
	}
}

func TestM5ChaosNilAndZeroInputs(t *testing.T) {
	var rep mutation.Report
	inputs := [][]rule.Finding{nil, {}, sampleFindings()}
	opts := []HTMLOptions{{}, DefaultHTMLOptions(), {MaxRows: -7, EmbedSource: true}}
	for i := range inputs {
		html := m5MustHTML(t, inputs[i], rep, opts[i])
		junit := m5MustJUnit(t, inputs[i], rep)
		otel := m5MustOTel(t, inputs[i], rep)
		if len(html) == 0 || len(junit) == 0 || len(otel) == 0 {
			t.Fatalf("input %d produced an empty document", i)
		}
		if !bytes.Equal(html, m5MustHTML(t, inputs[i], rep, opts[i])) ||
			!bytes.Equal(junit, m5MustJUnit(t, inputs[i], rep)) ||
			!bytes.Equal(otel, m5MustOTel(t, inputs[i], rep)) {
			t.Fatalf("input %d is not byte-stable", i)
		}
	}
	if !bytes.Equal(m5MustJUnit(t, nil, rep), m5MustJUnit(t, []rule.Finding{}, rep)) {
		t.Fatal("nil and empty findings must produce identical JUnit bytes")
	}
}

func TestM5ChaosConcurrentBuilders(t *testing.T) {
	const workers = 48
	findings := sampleFindings()
	rep := m5VerdictReport(verify.MutantKilled, verify.MutantSurvived)
	var wg sync.WaitGroup
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			if _, err := BuildHTML(findings, rep, DefaultHTMLOptions()); err != nil {
				t.Errorf("concurrent BuildHTML: %v", err)
			}
			if _, err := BuildJUnit(findings, rep); err != nil {
				t.Errorf("concurrent BuildJUnit: %v", err)
			}
			if _, err := BuildOTel(findings, rep); err != nil {
				t.Errorf("concurrent BuildOTel: %v", err)
			}
		}()
	}
	wg.Wait()
}

func TestM5ChaosMalformedVerdictsAndLengthMismatches(t *testing.T) {
	rep := m5VerdictReport(verify.MutationVerdict(""),
		verify.MutationVerdict("ODD"), verify.MutantKilled)
	rep.Points = rep.Points[:2]
	rep.KillingTests = nil
	rows := mutantRows(rep)
	if len(rows) != 2 {
		t.Fatalf("want two aligned rows, got %d", len(rows))
	}
	for _, r := range rows {
		if r.KillingTest != "" {
			t.Fatal("a missing killing test was not defaulted to empty")
		}
	}
	findings := sampleFindings()
	html := m5MustHTML(t, findings, rep, HTMLOptions{MaxRows: 1})
	junit := m5MustJUnit(t, findings, rep)
	otel := m5MustOTel(t, findings, rep)
	if len(html) == 0 || len(junit) == 0 || len(otel) == 0 {
		t.Fatal("malformed verdicts produced an empty document")
	}
	if !bytes.Equal(junit, m5MustJUnit(t, findings, rep)) ||
		!bytes.Equal(otel, m5MustOTel(t, findings, rep)) {
		t.Fatal("malformed verdicts are not byte-stable")
	}
}

func TestM5ChaosDeterministicHelperBranches(t *testing.T) {
	baseRow := mutantRow{Fingerprint: "a", Node: "n", Mutated: "m", Verdict: "KILLED", Index: 1}
	rowVariants := []mutantRow{
		{Fingerprint: "b"},
		{Fingerprint: "a", Node: "z"},
		{Fingerprint: "a", Node: "n", Mutated: "z"},
		{Fingerprint: "a", Node: "n", Mutated: "m", Verdict: "Z"},
		{Fingerprint: "a", Node: "n", Mutated: "m", Verdict: "KILLED", Index: 2},
	}
	for i, v := range rowVariants {
		if !mutantRowLess(baseRow, v) || mutantRowLess(v, baseRow) {
			t.Fatalf("mutantRowLess is not a total order at variant %d", i)
		}
	}
	if mutantRowLess(baseRow, baseRow) {
		t.Fatal("mutantRowLess must be irreflexive")
	}
	la := otelLogRecord{Body: "b", VirtualTime: "1", SeverityText: "INFO",
		Attributes: []otelAttribute{{Value: "verdict"}}}
	logVariants := []otelLogRecord{
		{Body: "c"},
		{Body: "b", VirtualTime: "2"},
		{Body: "b", VirtualTime: "1", SeverityText: "WARN"},
		{Body: "b", VirtualTime: "1", SeverityText: "INFO", Attributes: []otelAttribute{{Value: "z"}}},
	}
	for i, v := range logVariants {
		if !otelLogLess(la, v) || otelLogLess(v, la) {
			t.Fatalf("otelLogLess is not a total order at variant %d", i)
		}
	}
	if otelLogKind(otelLogRecord{}) != "" {
		t.Fatal("otelLogKind must return an empty sentinel for empty rows")
	}
	if !otelSpanLess(otelSpan{SpanID: "a"}, otelSpan{SpanID: "b"}) ||
		!junitCaseLess(junitCase{Name: "a"}, junitCase{Name: "b"}) {
		t.Fatal("span or case ordering is inverted")
	}
	verdictLevels := []struct {
		verdict string
		want    string
	}{
		{string(verify.MutantSurvived), "ERROR"},
		{string(verify.MutantEquivalent), "WARN"},
		{string(verify.MutantTimeout), "WARN"},
		{string(verify.MutantKilled), "INFO"},
		{"ODD", "WARN"},
	}
	for _, tc := range verdictLevels {
		if got := otelVerdictSeverity(tc.verdict); got != tc.want {
			t.Fatalf("otelVerdictSeverity(%q) = %q, want %q", tc.verdict, got, tc.want)
		}
	}
}

func TestM5ChaosDeterministicGaugesAndDigests(t *testing.T) {
	rep := sampleReport()
	severityLevels := []struct {
		severity rule.Severity
		want     string
	}{
		{rule.SeverityError, "ERROR"},
		{rule.SeverityWarning, "WARN"},
		{rule.SeverityInfo, "INFO"},
		{rule.Severity("odd"), "WARN"},
	}
	for _, tc := range severityLevels {
		if got := otelFindingSeverity(tc.severity); got != tc.want {
			t.Fatalf("otelFindingSeverity(%q) = %q, want %q", tc.severity, got, tc.want)
		}
	}
	if gaugeFromReport(rep).RatePermille != 777 {
		t.Fatal("per-mille kill rate is not integer-deterministic")
	}
	if gaugeFromReport(mutation.Report{}).RatePermille != 0 ||
		gaugeFromReport(mutation.Report{Score: verify.Score{Total: 1, Equivalent: 2}}).RatePermille != 0 {
		t.Fatal("a zero denominator must yield a zero rate")
	}
	caps := []struct{ total, max, want int }{{10, 0, 10}, {10, 3, 3}, {2, 3, 2}, {10, -1, 10}}
	for _, c := range caps {
		if got := rowCap(c.total, c.max); got != c.want {
			t.Fatalf("rowCap(%d,%d) = %d, want %d", c.total, c.max, got, c.want)
		}
	}
	if len(timelineRows(rep)) != 6 {
		t.Fatal("timeline must contain exactly six virtual phases")
	}
	sorted := sortedFindings(sampleFindings())
	digestA := findingsDigest(sorted)
	digestB := findingsDigest(sorted)
	if digestA != digestB {
		t.Fatal("findingsDigest is not stable")
	}
	mutated := append([]rule.Finding(nil), sorted...)
	mutated[0].Message = "changed"
	if findingsDigest(mutated) == findingsDigest(sorted) {
		t.Fatal("findingsDigest ignores a changed finding")
	}
	if len(otelTraceID(rep)) != 32 || len(otelSpanID(rep, "x")) != 16 ||
		otelSpanID(rep, "x") == otelSpanID(rep, "y") {
		t.Fatal("span identifiers are not deterministic digests")
	}
	shortA := findingShortDigest(sorted[0])
	shortB := findingShortDigest(sorted[0])
	if shortA != shortB {
		t.Fatal("findingShortDigest is not stable")
	}
}
