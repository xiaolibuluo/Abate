/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package report

import (
	"bytes"
	"encoding/hex"
	"encoding/xml"
	"strings"
	"testing"

	"abate/kernel/verify"
	"abate/lang/rule"
)

func TestM5MutationOneFindingChangesAllFormats(t *testing.T) {
	base := sampleFindings()
	changed := append([]rule.Finding(nil), base...)
	changed[1].Message = "a single byte changed"
	rep := m5VerdictReport(verify.MutantKilled, verify.MutantSurvived)
	if bytes.Equal(m5MustHTML(t, base, rep, DefaultHTMLOptions()),
		m5MustHTML(t, changed, rep, DefaultHTMLOptions())) {
		t.Fatal("one changed finding did not change the HTML bytes")
	}
	if bytes.Equal(m5MustJUnit(t, base, rep), m5MustJUnit(t, changed, rep)) {
		t.Fatal("one changed finding did not change the JUnit bytes")
	}
	if bytes.Equal(m5MustOTel(t, base, rep), m5MustOTel(t, changed, rep)) {
		t.Fatal("one changed finding did not change the OTel bytes")
	}
}

func TestM5MutationOneVerdictChangesAllFormats(t *testing.T) {
	base := m5VerdictReport(verify.MutantKilled, verify.MutantKilled, verify.MutantSurvived)
	changed := m5VerdictReport(verify.MutantKilled, verify.MutantSurvived, verify.MutantSurvived)
	findings := sampleFindings()
	aHTML := m5MustHTML(t, findings, base, DefaultHTMLOptions())
	bHTML := m5MustHTML(t, findings, changed, DefaultHTMLOptions())
	aJUnit := m5MustJUnit(t, findings, base)
	bJUnit := m5MustJUnit(t, findings, changed)
	aOTel := m5MustOTel(t, findings, base)
	bOTel := m5MustOTel(t, findings, changed)
	if bytes.Equal(aHTML, bHTML) || bytes.Equal(aJUnit, bJUnit) || bytes.Equal(aOTel, bOTel) {
		t.Fatal("one changed verdict did not change all three formats")
	}
	if strings.Count(string(bJUnit), "<failure") != strings.Count(string(aJUnit), "<failure")+1 {
		t.Fatal("a killed-to-survived verdict did not add one JUnit failure")
	}
}

func TestM5MutationOptionsChangeHTML(t *testing.T) {
	findings := sampleFindings()
	rep := m5VerdictReport(verify.MutantKilled, verify.MutantSurvived)
	base := m5MustHTML(t, findings, rep, HTMLOptions{})
	if !bytes.Equal(base, m5MustHTML(t, findings, rep, DefaultHTMLOptions())) {
		t.Fatal("a zero options value must equal the defaults")
	}
	if !bytes.Equal(base, m5MustHTML(t, findings, rep, HTMLOptions{MaxRows: -3})) {
		t.Fatal("a negative row cap must equal the defaults")
	}
	custom := m5MustHTML(t, findings, rep, HTMLOptions{Title: "自定义 & <标题>"})
	if bytes.Equal(base, custom) {
		t.Fatal("a custom title did not change the HTML bytes")
	}
	if !strings.Contains(string(custom), "&amp;") || !strings.Contains(string(custom), "&lt;") {
		t.Fatal("a custom title was not escaped")
	}
	embed := m5MustHTML(t, findings, rep, HTMLOptions{EmbedSource: true})
	if bytes.Equal(base, embed) || !strings.Contains(string(embed), "Evidence") {
		t.Fatal("EmbedSource did not add the evidence column")
	}
}

func TestM5MutationReportIdentityChangesAllFormats(t *testing.T) {
	base := sampleReport()
	changed := sampleReport()
	changed.Fingerprint[0] ^= 0xFF
	changed.Seed++
	findings := sampleFindings()
	if bytes.Equal(m5MustHTML(t, findings, base, DefaultHTMLOptions()),
		m5MustHTML(t, findings, changed, DefaultHTMLOptions())) {
		t.Fatal("a changed fingerprint did not change the HTML bytes")
	}
	aJUnit := m5MustJUnit(t, findings, base)
	bJUnit := m5MustJUnit(t, findings, changed)
	if bytes.Equal(aJUnit, bJUnit) {
		t.Fatal("a changed fingerprint did not change the JUnit bytes")
	}
	if !strings.Contains(string(aJUnit), hex.EncodeToString(base.Fingerprint[:])) {
		t.Fatal("JUnit properties do not carry the report fingerprint")
	}
	if bytes.Equal(m5MustOTel(t, findings, base), m5MustOTel(t, findings, changed)) {
		t.Fatal("a changed fingerprint did not change the OTel bytes")
	}
}

func TestM5MutationMaxRowsTruncatesOnlyHTML(t *testing.T) {
	findings := sampleFindings()
	rep := m5VerdictReport(verify.MutantKilled, verify.MutantSurvived, verify.MutantTimeout)
	full := m5MustHTML(t, findings, rep, HTMLOptions{})
	one := m5MustHTML(t, findings, rep, HTMLOptions{MaxRows: 1})
	two := m5MustHTML(t, findings, rep, HTMLOptions{MaxRows: 2})
	if len(one) >= len(two) || len(two) >= len(full) {
		t.Fatal("MaxRows does not shrink the rendered document")
	}
	if !strings.Contains(string(one), "rows truncated by MaxRows") ||
		strings.Contains(string(full), "rows truncated by MaxRows") {
		t.Fatal("the truncation marker is inconsistent")
	}
	var suites junitDocument
	if err := xml.Unmarshal(m5MustJUnit(t, findings, rep), &suites); err != nil {
		t.Fatalf("JUnit document is not valid XML: %v", err)
	}
	if suites.Suites[0].Tests != len(findings)+3 {
		t.Fatalf("JUnit case count = %d, want %d", suites.Suites[0].Tests, len(findings)+3)
	}
	if suites.Suites[0].Properties.Properties[0].Name != "report.seed" {
		t.Fatal("JUnit properties lost their deterministic report identity")
	}
}
