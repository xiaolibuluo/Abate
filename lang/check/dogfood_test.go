/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package check

import (
	"path/filepath"
	"slices"
	"testing"

	"abate/lang/gofront"
	"abate/lang/rule"
	gobuiltin "abate/lang/rule/builtin/go"
)

// TestDogfoodSelfCheck runs the self-developed checker over every Abate Go
// source file (excluding vendor and third_party) and verifies that after the
// registered exemptions are deducted there are zero unexplained violations.
//
// No M1/M2/lang production source is modified to pass this test, no rule is
// removed, and every exemption carries a documented basis in the built-in
// table. The counts are reported honestly through t.Logf.
func TestDogfoodSelfCheck(t *testing.T) {
	root, err := filepath.Abs("../..")
	if err != nil {
		t.Fatalf("resolve repository root: %v", err)
	}
	registry := rule.NewRegistry()
	if err := gobuiltin.LoadBuiltins(registry); err != nil {
		t.Fatalf("load built-in rules: %v", err)
	}
	checker := NewChecker(gofront.NewGoFrontend().WithTypeCheck(false),
		registry.For("go"))
	result, err := checker.Check([]string{root}, DefaultOptions())
	if err != nil {
		t.Fatalf("self-check: %v", err)
	}
	exemptions := gobuiltin.Exemptions()
	exemptCount := 0
	unexplained := make(map[string]int)
	for _, finding := range result.Findings {
		if _, ok := exemptions[string(finding.RuleID)]; ok {
			exemptCount++
			continue
		}
		unexplained[string(finding.RuleID)]++
	}
	ids := make([]string, 0, len(unexplained))
	for id := range unexplained {
		ids = append(ids, id)
	}
	slices.Sort(ids)
	t.Logf("dogfood: files=%d activeRules=%d totalHits=%d exemptHits=%d unexplainedRules=%d",
		result.Summary.Files, result.Summary.Rules,
		len(result.Findings), exemptCount, len(ids))
	for _, id := range ids {
		t.Logf("unexplained rule %s: %d hits", id, unexplained[id])
	}
	if len(ids) != 0 {
		t.Fatalf("self-check has %d unexplained rule(s): %v", len(ids), ids)
	}
	if len(result.Findings) != exemptCount {
		t.Fatal("exemption accounting does not match the total hit count")
	}
}
