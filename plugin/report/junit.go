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
	"encoding/xml"
	"fmt"
	"sort"
	"strconv"

	"abate/kernel/mutation"
	"abate/kernel/verify"
	"abate/lang/rule"
)

// junitMutantCase maps one mutant row to a JUnit test case: a killed mutant
// passes, a surviving mutant fails and equivalent or timed-out mutants skip.
func junitMutantCase(row mutantRow) junitCase {
	base := junitCase{
		Name:      "mutant-" + strconv.Itoa(row.Index) + "-" + row.Fingerprint,
		Classname: "abate.mutation",
		Time:      junitTimeNotMeasured,
	}
	switch verify.MutationVerdict(row.Verdict) {
	case verify.MutantKilled:
		return base
	case verify.MutantSurvived:
		base.Failures = []junitFailure{{
			Message: "mutant survived the suite",
			Type:    row.Verdict,
			Body:    "node=" + row.Node + " original=" + row.Original + " mutated=" + row.Mutated,
		}}
	case verify.MutantEquivalent, verify.MutantTimeout:
		base.Skipped = &junitSkipped{Message: "mutant " + row.Verdict}
	default:
		base.Failures = []junitFailure{{
			Message: "unknown mutant verdict",
			Type:    "UNKNOWN",
			Body:    row.Verdict,
		}}
	}
	return base
}

// findingShortDigest returns a short, stable identity for a finding test case.
func findingShortDigest(f rule.Finding) string {
	sum := sha256.Sum256([]byte(f.Canonical()))
	return hex.EncodeToString(sum[:6])
}

// junitCaseLess is the canonical order over generated test cases.
func junitCaseLess(a, b junitCase) bool {
	return a.Name < b.Name
}

// BuildJUnit renders findings and a mutation report into a deterministic JUnit
// XML document. Every mutant becomes a test case and every finding becomes a
// failing test case; encoding/xml escapes all text and no wall-clock or duration
// is ever measured. Timing is reported as the numeric placeholder "0" together
// with an explicit report.time=NOT_MEASURED property so it is never mistaken for
// a measured elapsed time.
//
// 参考 JUnit XML 的 testsuite/testcase 语义，改进点=四类变异判定显式映射且
// 明确标注耗时未测量（NOT_MEASURED），输出与输入顺序无关。
// The returned bytes are a pure function of the two inputs.
func BuildJUnit(findings []rule.Finding, rep mutation.Report) ([]byte, error) {
	sorted := sortedFindings(findings)
	rows := mutantRows(rep)
	cases := make([]junitCase, 0, len(sorted)+len(rows))
	for _, row := range rows {
		cases = append(cases, junitMutantCase(row))
	}
	for i, f := range sorted {
		cases = append(cases, junitCase{
			Name:      "finding-" + strconv.Itoa(i+1) + "-" + findingShortDigest(f),
			Classname: "abate.rule." + string(f.RuleID),
			Time:      junitTimeNotMeasured,
			Failures: []junitFailure{{
				Message: f.Message,
				Type:    f.Severity.String(),
				Body:    f.String(),
			}},
		})
	}
	sort.Slice(cases, func(i, j int) bool {
		return junitCaseLess(cases[i], cases[j])
	})
	suite := junitSuite{
		Name:  "abate",
		Tests: len(cases),
		Time:  junitTimeNotMeasured,
		Properties: junitProperties{Properties: []junitProperty{
			{Name: "report.seed", Value: strconv.FormatUint(rep.Seed, 10)},
			{Name: "report.fingerprint", Value: hex.EncodeToString(rep.Fingerprint[:])},
			{Name: "findings.digest", Value: findingsDigest(sorted)},
			{Name: "report.time", Value: junitTimeNotMeasuredMarker},
		}},
		Cases: cases,
	}
	for _, c := range cases {
		if len(c.Failures) > 0 {
			suite.Failures++
		}
		if c.Skipped != nil {
			suite.Skipped++
		}
	}
	body, err := xml.MarshalIndent(junitDocument{Suites: []junitSuite{suite}}, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("report: marshal junit document: %w", err)
	}
	out := make([]byte, 0, len(xml.Header)+len(body)+1)
	out = append(out, xml.Header...)
	out = append(out, body...)
	out = append(out, '\n')
	return out, nil
}
