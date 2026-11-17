/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package report

import (
	"crypto/sha256"
	"encoding/hex"
	"sort"
	"strconv"

	"abate/kernel/mutation"
	"abate/lang/rule"
)

// HTMLModel is the fully deterministic four-section data model of the page.
//
// 参考 PITest 的 HTML 报告四区布局，改进点=违规区 / 变异矩阵区 / 时间线区 /
// 分数仪表区全部由 findings 与 mutation.Report 的确定性字段投影，零时间戳。
// Every field is a pure function of the two inputs and the rendering options.
type HTMLModel struct {
	Title          string
	Violations     []htmlViolation
	Mutants        []mutantRow
	Timeline       []htmlTimeline
	Gauge          htmlGauge
	Footer         string
	FindingsDigest string
	ViolationTotal int
	MutantTotal    int
	Truncated      bool
	ShowSource     bool
}

// htmlViolation is one rule match rendered in the violations section.
type htmlViolation struct {
	Index     int
	RuleID    string
	Path      string
	Line      int
	Column    int
	Severity  string
	Category  string
	Message   string
	Reference string
	Evidence  string
}

// mutantRow is the shared canonical projection of one aligned point, verdict
// and killing test, consumed by all three deterministic renderers.
type mutantRow struct {
	Index       int
	Node        string
	Original    string
	Mutated     string
	Priority    uint64
	Selector    uint64
	Verdict     string
	KillingTest string
	Fingerprint string
}

// htmlTimeline is one logical phase of the run, ordered by a virtual integer.
type htmlTimeline struct {
	VirtualTime int
	Phase       string
	Value       string
}

// htmlGauge is the integer-only score gauge; no floating point value is emitted.
type htmlGauge struct {
	Seed         uint64
	Total        uint64
	Killed       uint64
	Survived     uint64
	Equivalent   uint64
	Timeout      uint64
	RatePermille uint64
	KillingTests int
}

// sortedFindings returns a copy of the findings in canonical order so every
// renderer is independent of the caller's input order.
func sortedFindings(findings []rule.Finding) []rule.Finding {
	sorted := append([]rule.Finding(nil), findings...)
	sort.Slice(sorted, func(i, j int) bool {
		return findingLess(sorted[i], sorted[j])
	})
	return sorted
}

// findingsDigest returns the leading hex digits of the digest over the canonical
// form of every sorted finding; it replaces a wall-clock timestamp in the footer.
func findingsDigest(sorted []rule.Finding) string {
	h := sha256.New()
	for _, f := range sorted {
		_, _ = h.Write([]byte(f.Canonical()))
	}
	sum := h.Sum(nil)
	return hex.EncodeToString(sum[:8])
}

// mutantRowLess is the canonical total order over mutant rows.
func mutantRowLess(a, b mutantRow) bool {
	if a.Fingerprint != b.Fingerprint {
		return a.Fingerprint < b.Fingerprint
	}
	if a.Node != b.Node {
		return a.Node < b.Node
	}
	if a.Mutated != b.Mutated {
		return a.Mutated < b.Mutated
	}
	if a.Verdict != b.Verdict {
		return a.Verdict < b.Verdict
	}
	return a.Index < b.Index
}

// mutantRows projects the aligned point, verdict and killing-test triples into
// a canonical sorted list; length mismatches are handled deterministically.
func mutantRows(rep mutation.Report) []mutantRow {
	n := len(rep.Verdicts)
	if len(rep.Points) < n {
		n = len(rep.Points)
	}
	rows := make([]mutantRow, 0, n)
	for i := 0; i < n; i++ {
		point := rep.Points[i]
		killing := ""
		if i < len(rep.KillingTests) {
			killing = rep.KillingTests[i]
		}
		rows = append(rows, mutantRow{
			Index:       i,
			Node:        point.Node,
			Original:    point.Original,
			Mutated:     point.Mutated,
			Priority:    point.Priority,
			Selector:    point.Selector,
			Verdict:     string(rep.Verdicts[i]),
			KillingTest: killing,
			Fingerprint: hex.EncodeToString(point.Fingerprint[:]),
		})
	}
	sort.Slice(rows, func(i, j int) bool {
		return mutantRowLess(rows[i], rows[j])
	})
	for i := range rows {
		rows[i].Index = i + 1
	}
	return rows
}

// timelineRows builds the six fixed logical phases with integer virtual time.
func timelineRows(rep mutation.Report) []htmlTimeline {
	names := [6]string{"seed", "target", "suite", "schedule", "verify", "score"}
	values := [6]string{
		strconv.FormatUint(rep.Seed, 10),
		hex.EncodeToString(rep.TargetFingerprint[:8]),
		hex.EncodeToString(rep.SuiteFingerprint[:8]),
		strconv.FormatUint(rep.Score.Total, 10),
		strconv.FormatUint(rep.Score.Killed, 10),
		hex.EncodeToString(rep.Fingerprint[:8]),
	}
	rows := make([]htmlTimeline, 0, len(names))
	for i, name := range names {
		rows = append(rows, htmlTimeline{
			VirtualTime: i,
			Phase:       name,
			Value:       values[i],
		})
	}
	return rows
}

// gaugeFromReport derives an integer-only score gauge, including a per-mille
// kill rate computed with integer arithmetic.
func gaugeFromReport(rep mutation.Report) htmlGauge {
	denom := rep.Score.Total
	if rep.Score.Equivalent <= denom {
		denom -= rep.Score.Equivalent
	} else {
		denom = 0
	}
	rate := uint64(0)
	if denom > 0 {
		rate = rep.Score.Killed * 1000 / denom
	}
	return htmlGauge{
		Seed:         rep.Seed,
		Total:        rep.Score.Total,
		Killed:       rep.Score.Killed,
		Survived:     rep.Score.Survived,
		Equivalent:   rep.Score.Equivalent,
		Timeout:      rep.Score.Timeout,
		RatePermille: rate,
		KillingTests: len(rep.KillingTests),
	}
}

// rowCap returns the number of rows to render under the configured cap.
func rowCap(total, limit int) int {
	if limit > 0 && total > limit {
		return limit
	}
	return total
}

// buildHTMLModel assembles the four deterministic sections from findings and rep.
func buildHTMLModel(findings []rule.Finding, rep mutation.Report, opts HTMLOptions) HTMLModel {
	norm := normalizeHTMLOptions(opts)
	sorted := sortedFindings(findings)
	rows := mutantRows(rep)
	model := HTMLModel{
		Title:          norm.Title,
		ShowSource:     norm.EmbedSource,
		Footer:         hex.EncodeToString(rep.Fingerprint[:]),
		FindingsDigest: findingsDigest(sorted),
		ViolationTotal: len(sorted),
		MutantTotal:    len(rows),
		Gauge:          gaugeFromReport(rep),
		Timeline:       timelineRows(rep),
	}
	violationLimit := rowCap(len(sorted), norm.MaxRows)
	mutantLimit := rowCap(len(rows), norm.MaxRows)
	model.Truncated = violationLimit < len(sorted) || mutantLimit < len(rows)
	model.Violations = make([]htmlViolation, 0, violationLimit)
	for i := 0; i < violationLimit; i++ {
		f := sorted[i]
		evidence := ""
		if norm.EmbedSource {
			evidence = f.String()
		}
		model.Violations = append(model.Violations, htmlViolation{
			Index:     i + 1,
			RuleID:    string(f.RuleID),
			Path:      f.Path,
			Line:      f.Line,
			Column:    f.Column,
			Severity:  f.Severity.String(),
			Category:  f.Category.String(),
			Message:   f.Message,
			Reference: f.Reference,
			Evidence:  evidence,
		})
	}
	model.Mutants = rows[:mutantLimit]
	return model
}
