/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package report

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"encoding/xml"
	"strconv"
	"strings"
	"testing"

	"abate/internal/determinism"
	"abate/kernel/mutation"
	"abate/kernel/verify"
	"abate/lang/rule"
)

// m5VerdictReport builds a report whose mutants carry the supplied verdicts.
func m5VerdictReport(verdicts ...verify.MutationVerdict) mutation.Report {
	rep := sampleReport()
	rep.Points = make([]mutation.MutationPoint, 0, len(verdicts))
	rep.Verdicts = append([]verify.MutationVerdict(nil), verdicts...)
	rep.KillingTests = make([]string, 0, len(verdicts))
	for i, v := range verdicts {
		fp := sha256.Sum256([]byte("mutant-" + strconv.Itoa(i)))
		rep.Points = append(rep.Points, mutation.MutationPoint{
			Node:        "node-" + strconv.Itoa(i),
			Original:    "a",
			Mutated:     "b",
			Priority:    1,
			Selector:    2,
			Fingerprint: fp,
		})
		rep.KillingTests = append(rep.KillingTests, string(v)+"-test")
	}
	return rep
}

func m5MustHTML(t *testing.T, findings []rule.Finding, rep mutation.Report, opts HTMLOptions) []byte {
	t.Helper()
	out, err := BuildHTML(findings, rep, opts)
	if err != nil {
		t.Fatalf("BuildHTML: %v", err)
	}
	return out
}

func m5MustJUnit(t *testing.T, findings []rule.Finding, rep mutation.Report) []byte {
	t.Helper()
	out, err := BuildJUnit(findings, rep)
	if err != nil {
		t.Fatalf("BuildJUnit: %v", err)
	}
	return out
}

func m5MustOTel(t *testing.T, findings []rule.Finding, rep mutation.Report) []byte {
	t.Helper()
	out, err := BuildOTel(findings, rep)
	if err != nil {
		t.Fatalf("BuildOTel: %v", err)
	}
	return out
}

func TestM5ThreeBuildersDeterministic(t *testing.T) {
	findings := sampleFindings()
	rep := m5VerdictReport(verify.MutantKilled, verify.MutantSurvived,
		verify.MutantEquivalent, verify.MutantTimeout)
	firstHTML := m5MustHTML(t, findings, rep, DefaultHTMLOptions())
	firstJUnit := m5MustJUnit(t, findings, rep)
	firstOTel := m5MustOTel(t, findings, rep)
	for i := 0; i < 20; i++ {
		if !bytes.Equal(firstHTML, m5MustHTML(t, findings, rep, DefaultHTMLOptions())) {
			t.Fatalf("HTML drifted at run %d", i)
		}
		if !bytes.Equal(firstJUnit, m5MustJUnit(t, findings, rep)) {
			t.Fatalf("JUnit drifted at run %d", i)
		}
		if !bytes.Equal(firstOTel, m5MustOTel(t, findings, rep)) {
			t.Fatalf("OTel drifted at run %d", i)
		}
	}
	determinism.MustMatch100(t, 42, func() []byte {
		return m5MustOTel(t, findings, rep)
	})
	determinism.MustMatch100(t, 43, func() []byte {
		return m5MustJUnit(t, findings, rep)
	})
	determinism.MustMatch100(t, 44, func() []byte {
		return m5MustHTML(t, findings, rep, DefaultHTMLOptions())
	})
}

func TestM5ThreeFormatsEscapeHostileText(t *testing.T) {
	hostile := rule.Finding{
		RuleID: "GO-EG-001", Path: `a/<b>&"'.go`, Line: 1, Column: 2,
		Severity: rule.SeverityError, Category: rule.CategoryStyle,
		Message:   `<script>alert("x")</script>&'"><`,
		Reference: `<ref & "quoted">`,
	}
	findings := []rule.Finding{hostile}
	rep := sampleReport()
	html := string(m5MustHTML(t, findings, rep, DefaultHTMLOptions()))
	junit := string(m5MustJUnit(t, findings, rep))
	otel := string(m5MustOTel(t, findings, rep))
	for _, item := range []struct {
		name string
		out  string
	}{{"html", html}, {"junit", junit}, {"otel", otel}} {
		if strings.Contains(item.out, `<script>alert`) {
			t.Fatalf("%s leaked a raw script payload", item.name)
		}
	}
	if !strings.Contains(html, `&lt;script&gt;`) || !strings.Contains(html, `&#39;`) {
		t.Fatalf("HTML did not escape angle brackets and quotes: %q", html)
	}
	if !strings.Contains(junit, `&lt;script&gt;`) || !strings.Contains(junit, `&amp;`) {
		t.Fatal("JUnit did not escape XML metacharacters")
	}
	if !strings.Contains(otel, `\u003c`) || !strings.Contains(otel, `\u0026`) {
		t.Fatal("OTel JSON did not escape HTML metacharacters")
	}
	var suites junitDocument
	if err := xml.Unmarshal([]byte(junit), &suites); err != nil {
		t.Fatalf("escaped JUnit is not valid XML: %v", err)
	}
	if len(suites.Suites) != 1 || len(suites.Suites[0].Cases) != 1 {
		t.Fatal("escaped JUnit lost its test cases")
	}
	body := suites.Suites[0].Cases[0].Failures[0].Body
	if !strings.Contains(body, `<script>alert`) {
		t.Fatalf("JUnit round trip lost the original message: %q", body)
	}
}

func TestM5HTMLSelfContainedAndFingerprintFooter(t *testing.T) {
	rep := sampleReport()
	out := string(m5MustHTML(t, sampleFindings(), rep, DefaultHTMLOptions()))
	if !strings.HasPrefix(out, "<!DOCTYPE html>") {
		t.Fatal("HTML document is missing its doctype")
	}
	for _, marker := range []string{"<style>", "<script>", "data-row", "data-step", "Score gauge", "Virtual time"} {
		if !strings.Contains(out, marker) {
			t.Fatalf("self-contained HTML is missing inline asset %q", marker)
		}
	}
	if !strings.Contains(out, hex.EncodeToString(rep.Fingerprint[:])) {
		t.Fatal("footer does not carry the report fingerprint")
	}
	if !strings.Contains(out, findingsDigest(sortedFindings(sampleFindings()))) {
		t.Fatal("footer does not carry the findings digest")
	}
	for _, forbidden := range []string{"http://", "https://", "fetch(", "XMLHttpRequest", "eval(", "<script src", "timestamp"} {
		if strings.Contains(out, forbidden) {
			t.Fatalf("HTML leaks forbidden external marker %q", forbidden)
		}
	}
}

func TestM5EmptyInputsProduceStableDocuments(t *testing.T) {
	var rep mutation.Report
	html := m5MustHTML(t, nil, rep, HTMLOptions{})
	junit := m5MustJUnit(t, nil, rep)
	otel := m5MustOTel(t, nil, rep)
	if len(html) == 0 || len(junit) == 0 || len(otel) == 0 {
		t.Fatal("an empty input produced an empty document")
	}
	if !bytes.Equal(html, m5MustHTML(t, nil, rep, HTMLOptions{})) ||
		!bytes.Equal(junit, m5MustJUnit(t, nil, rep)) ||
		!bytes.Equal(otel, m5MustOTel(t, nil, rep)) {
		t.Fatal("empty-input documents are not byte-stable")
	}
	var doc otelDocument
	if err := json.Unmarshal(otel, &doc); err != nil {
		t.Fatalf("empty OTel payload is not valid JSON: %v", err)
	}
	if len(doc.ResourceSpans) != 1 || len(doc.Metrics) != 6 || len(doc.Logs) != 0 {
		t.Fatal("empty OTel payload has unexpected stream sizes")
	}
	var suites junitDocument
	if err := xml.Unmarshal(junit, &suites); err != nil {
		t.Fatalf("empty JUnit document is not valid XML: %v", err)
	}
	if len(suites.Suites) != 1 || suites.Suites[0].Tests != 0 {
		t.Fatal("empty JUnit document has unexpected cases")
	}
}

func TestM5JUnitVerdictMapping(t *testing.T) {
	rep := m5VerdictReport(verify.MutantKilled, verify.MutantSurvived,
		verify.MutantEquivalent, verify.MutantTimeout, verify.MutationVerdict("ODD"))
	out := m5MustJUnit(t, nil, rep)
	var doc junitDocument
	if err := xml.Unmarshal(out, &doc); err != nil {
		t.Fatalf("JUnit document is not valid XML: %v", err)
	}
	if len(doc.Suites) != 1 {
		t.Fatalf("want one suite, got %d", len(doc.Suites))
	}
	suite := doc.Suites[0]
	if suite.Tests != len(suite.Cases) || suite.Tests != 5 {
		t.Fatalf("unexpected case count: %d", suite.Tests)
	}
	var failures, skips, passes int
	for _, c := range suite.Cases {
		switch {
		case len(c.Failures) > 0:
			failures++
		case c.Skipped != nil:
			skips++
		default:
			passes++
		}
	}
	if passes != 1 || skips != 2 || failures != 2 {
		t.Fatalf("mapping passes=%d skips=%d failures=%d, want 1/2/2", passes, skips, failures)
	}
	if suite.Failures != failures || suite.Skipped != skips || suite.Errors != 0 {
		t.Fatal("suite counters disagree with the case elements")
	}
	if suite.Time != junitTimeNotMeasured {
		t.Fatal("JUnit timing placeholder changed")
	}
	foundTimingMarker := false
	for _, prop := range suite.Properties.Properties {
		if prop.Name == "report.time" && prop.Value == junitTimeNotMeasuredMarker {
			foundTimingMarker = true
		}
	}
	if !foundTimingMarker {
		t.Fatal("JUnit document does not declare timing NOT_MEASURED")
	}
	if !strings.Contains(string(out), "<skipped") || !strings.Contains(string(out), "<failure") {
		t.Fatal("JUnit document is missing failure or skipped elements")
	}
}

func TestM5OTelAirGapPayload(t *testing.T) {
	rep := m5VerdictReport(verify.MutantKilled, verify.MutantSurvived,
		verify.MutantEquivalent, verify.MutantTimeout)
	out := string(m5MustOTel(t, sampleFindings(), rep))
	var doc otelDocument
	if err := json.Unmarshal([]byte(out), &doc); err != nil {
		t.Fatalf("OTel payload is not valid JSON: %v", err)
	}
	if len(doc.ResourceSpans) != 1 || len(doc.ResourceSpans[0].Spans) != 9 {
		t.Fatal("span tree does not contain one root, four phases and four mutants")
	}
	if len(doc.Metrics) != 6 || len(doc.Logs) != 4+len(sampleFindings()) {
		t.Fatal("metric or log stream has an unexpected size")
	}
	if !strings.Contains(out, `"virtualTime"`) {
		t.Fatal("OTel payload is missing virtual time")
	}
	for _, forbidden := range []string{"endpoint", "url", "host", "address", "http", "timeUnixNano", "startTimeUnixNano", "endTimeUnixNano", "timestamp"} {
		if strings.Contains(out, forbidden) {
			t.Fatalf("air-gap payload leaks forbidden marker %q", forbidden)
		}
	}
}

func TestM5LargeBatchIsStable(t *testing.T) {
	const total = 10000
	findings := make([]rule.Finding, 0, total)
	for i := 0; i < total; i++ {
		findings = append(findings, rule.Finding{
			RuleID: "GO-EG-001", Path: "pkg/file" + strconv.Itoa(i) + ".go",
			Line: i + 1, Column: 1, Severity: rule.SeverityWarning,
			Category: rule.CategoryStyle, Message: "finding " + strconv.Itoa(i),
			Reference: "EG-001",
		})
	}
	rep := m5VerdictReport(verify.MutantKilled, verify.MutantSurvived)
	html := m5MustHTML(t, findings, rep, DefaultHTMLOptions())
	junit := m5MustJUnit(t, findings, rep)
	otel := m5MustOTel(t, findings, rep)
	if len(html) == 0 || len(junit) == 0 || len(otel) == 0 {
		t.Fatal("a large batch produced an empty document")
	}
	if !bytes.Equal(html, m5MustHTML(t, findings, rep, DefaultHTMLOptions())) ||
		!bytes.Equal(junit, m5MustJUnit(t, findings, rep)) ||
		!bytes.Equal(otel, m5MustOTel(t, findings, rep)) {
		t.Fatal("a large batch is not byte-stable")
	}
	if !strings.Contains(string(html), "finding 9999") {
		t.Fatal("a large batch silently dropped its tail rows")
	}
}
