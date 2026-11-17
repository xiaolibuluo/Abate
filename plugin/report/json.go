/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package report

import (
	"encoding/hex"
	"encoding/json"
	"sort"

	"abate/kernel/mutation"
	"abate/lang/rule"
)

// jsonFinding is the deterministic JSON form of a single finding. The field order
// is fixed so encoding/json always emits identical bytes for identical findings.
type jsonFinding struct {
	RuleID    string `json:"ruleId"`
	Path      string `json:"path"`
	Line      int    `json:"line"`
	Column    int    `json:"column"`
	Severity  string `json:"severity"`
	Category  string `json:"category"`
	Message   string `json:"message"`
	Reference string `json:"reference"`
}

// jsonScore is the deterministic JSON form of a mutation score.
type jsonScore struct {
	Total      uint64 `json:"total"`
	Killed     uint64 `json:"killed"`
	Survived   uint64 `json:"survived"`
	Equivalent uint64 `json:"equivalent"`
	Timeout    uint64 `json:"timeout"`
}

// jsonReport is the deterministic JSON form of a mutation report summary.
type jsonReport struct {
	Seed              uint64    `json:"seed"`
	TargetFingerprint string    `json:"targetFingerprint"`
	SuiteFingerprint  string    `json:"suiteFingerprint"`
	Fingerprint       string    `json:"fingerprint"`
	Points            int       `json:"points"`
	Verdicts          int       `json:"verdicts"`
	KillingTests      int       `json:"killingTests"`
	Score             jsonScore `json:"score"`
}

// jsonDocument is the top-level compact JSON report. It concatenates the findings
// and the mutation report summary in a fixed field order.
type jsonDocument struct {
	Findings []jsonFinding `json:"findings"`
	Report   jsonReport    `json:"report"`
}

// BuildJSON renders findings and a mutation report into a compact, deterministic
// JSON document. Output is a pure function of the inputs: findings are sorted by a
// canonical key and the struct field order is fixed, so identical inputs always
// yield byte-identical output.
func BuildJSON(findings []rule.Finding, rep mutation.Report) ([]byte, error) {
	sorted := append([]rule.Finding(nil), findings...)
	sort.Slice(sorted, func(i, j int) bool {
		return findingLess(sorted[i], sorted[j])
	})
	out := make([]jsonFinding, 0, len(sorted))
	for _, f := range sorted {
		out = append(out, jsonFinding{
			RuleID:    string(f.RuleID),
			Path:      f.Path,
			Line:      f.Line,
			Column:    f.Column,
			Severity:  string(f.Severity),
			Category:  string(f.Category),
			Message:   f.Message,
			Reference: f.Reference,
		})
	}
	doc := jsonDocument{
		Findings: out,
		Report: jsonReport{
			Seed:              rep.Seed,
			TargetFingerprint: hex.EncodeToString(rep.TargetFingerprint[:]),
			SuiteFingerprint:  hex.EncodeToString(rep.SuiteFingerprint[:]),
			Fingerprint:       hex.EncodeToString(rep.Fingerprint[:]),
			Points:            len(rep.Points),
			Verdicts:          len(rep.Verdicts),
			KillingTests:      len(rep.KillingTests),
			Score: jsonScore{
				Total:      rep.Score.Total,
				Killed:     rep.Score.Killed,
				Survived:   rep.Score.Survived,
				Equivalent: rep.Score.Equivalent,
				Timeout:    rep.Score.Timeout,
			},
		},
	}
	return json.Marshal(doc)
}
