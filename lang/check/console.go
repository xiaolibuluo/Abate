/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package check

import (
	"encoding/json"
	"strings"

	"abate/lang/rule"
)

// jsonCount is one sorted count entry in a JSON summary.
type jsonCount struct {
	Key   string `json:"key"`
	Count int    `json:"count"`
}

// jsonSummary mirrors Summary with ordered arrays instead of maps.
type jsonSummary struct {
	Files      int         `json:"files"`
	Rules      int         `json:"rules"`
	Total      int         `json:"total"`
	BySeverity []jsonCount `json:"bySeverity"`
	ByRule     []jsonCount `json:"byRule"`
	Canonical  string      `json:"canonical"`
}

// jsonFinding is the stable JSON projection of one finding.
type jsonFinding struct {
	RuleID    string   `json:"ruleId"`
	Path      string   `json:"path"`
	Line      int      `json:"line"`
	Column    int      `json:"column"`
	Severity  string   `json:"severity"`
	Category  string   `json:"category"`
	Message   string   `json:"message"`
	Reference string   `json:"reference"`
	Captures  []string `json:"captures"`
}

// jsonReport is the deterministic intermediate JSON report layout.
type jsonReport struct {
	Schema      string        `json:"schema"`
	Fingerprint string        `json:"fingerprint"`
	Summary     jsonSummary   `json:"summary"`
	Findings    []jsonFinding `json:"findings"`
}

// renderConsole produces the human-readable, line-oriented report.
func renderConsole(r Result) []byte {
	var builder strings.Builder
	for _, finding := range r.Findings {
		_, _ = builder.WriteString(finding.String())
		_ = builder.WriteByte('\n')
	}
	_, _ = builder.WriteString(r.Summary.String())
	_ = builder.WriteByte('\n')
	_, _ = builder.WriteString("fingerprint=")
	_, _ = builder.WriteString(r.FingerprintHex())
	_ = builder.WriteByte('\n')
	return []byte(builder.String())
}

// renderJSON produces the deterministic intermediate JSON report.
func renderJSON(r Result) []byte {
	report := jsonReport{
		Schema: "abate.check.v1", Fingerprint: r.FingerprintHex(),
		Summary:  buildJSONSummary(r.Summary),
		Findings: buildJSONFindings(r.Findings),
	}
	data, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		return renderConsole(r)
	}
	return append(data, '\n')
}

// buildJSONSummary turns count maps into sorted arrays.
func buildJSONSummary(summary Summary) jsonSummary {
	out := jsonSummary{
		Files: summary.Files, Rules: summary.Rules, Total: summary.Total,
		Canonical:  summary.Canonical(),
		BySeverity: make([]jsonCount, 0, len(summary.BySeverity)),
		ByRule:     make([]jsonCount, 0, len(summary.ByRule)),
	}
	for _, key := range summary.SortedSeverities() {
		out.BySeverity = append(out.BySeverity,
			jsonCount{Key: key, Count: summary.BySeverity[key]})
	}
	for _, key := range summary.SortedRules() {
		out.ByRule = append(out.ByRule,
			jsonCount{Key: key, Count: summary.ByRule[key]})
	}
	return out
}

// buildJSONFindings projects findings into the stable JSON layout.
func buildJSONFindings(findings []rule.Finding) []jsonFinding {
	out := make([]jsonFinding, 0, len(findings))
	for _, finding := range findings {
		captures := make([]string, 0, len(finding.Captures))
		for _, capture := range finding.Captures {
			captures = append(captures, capture.String())
		}
		out = append(out, jsonFinding{
			RuleID: string(finding.RuleID), Path: finding.Path,
			Line: finding.Line, Column: finding.Column,
			Severity: finding.Severity.String(),
			Category: finding.Category.String(),
			Message:  finding.Message, Reference: finding.Reference,
			Captures: captures,
		})
	}
	return out
}
