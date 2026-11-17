/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package check

import (
	"fmt"
	"slices"
	"strings"

	"abate/lang/rule"
)

// Summary holds the deterministic statistics of one check run.
type Summary struct {
	// Files is the number of source buffers parsed.
	Files int
	// Rules is the number of active rules after filtering.
	Rules int
	// Total is the number of deduplicated findings.
	Total int
	// BySeverity counts findings per severity wire name.
	BySeverity map[string]int
	// ByRule counts findings per rule identifier.
	ByRule map[string]int
}

// NewSummary derives statistics from a deduplicated finding list.
func NewSummary(findings []rule.Finding, files, rules int) Summary {
	out := Summary{
		Files: files, Rules: rules, Total: len(findings),
		BySeverity: make(map[string]int, 3), ByRule: make(map[string]int, 8),
	}
	for _, finding := range findings {
		out.BySeverity[finding.Severity.String()]++
		out.ByRule[string(finding.RuleID)]++
	}
	return out
}

// SortedSeverities returns the severity keys in ascending order.
func (s Summary) SortedSeverities() []string {
	return sortedKeys(s.BySeverity)
}

// SortedRules returns the rule identifier keys in ascending order.
func (s Summary) SortedRules() []string {
	return sortedKeys(s.ByRule)
}

// Canonical renders the summary in a byte-stable form.
func (s Summary) Canonical() string {
	parts := []string{
		fmt.Sprintf("files=%d", s.Files),
		fmt.Sprintf("rules=%d", s.Rules),
		fmt.Sprintf("total=%d", s.Total),
	}
	for _, key := range s.SortedSeverities() {
		parts = append(parts, fmt.Sprintf("sev:%s=%d", key, s.BySeverity[key]))
	}
	for _, key := range s.SortedRules() {
		parts = append(parts, fmt.Sprintf("rule:%s=%d", key, s.ByRule[key]))
	}
	return strings.Join(parts, "\x00")
}

// String renders the summary as stable one-line text.
func (s Summary) String() string {
	return fmt.Sprintf("files=%d rules=%d findings=%d", s.Files, s.Rules, s.Total)
}

// sortedKeys returns the keys of a count map in ascending order.
func sortedKeys(counts map[string]int) []string {
	keys := make([]string, 0, len(counts))
	for key := range counts {
		keys = append(keys, key)
	}
	slices.Sort(keys)
	return keys
}
