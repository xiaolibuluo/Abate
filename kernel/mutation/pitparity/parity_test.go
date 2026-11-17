/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package pitparity

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"testing"

	"abate/kernel/mutation"
)

// TestCompareReachesFullParity asserts the acceptance criterion of M2.
func TestCompareReachesFullParity(t *testing.T) {
	t.Parallel()
	records := loadReference(t)
	report, err := Compare(records)
	if err != nil {
		t.Fatalf("compare: %v", err)
	}
	if report.Total != 28 {
		t.Fatalf("total = %d, want 28", report.Total)
	}
	if report.Comparable != 21 {
		t.Fatalf("comparable = %d, want 21", report.Comparable)
	}
	if report.Consistent != report.Comparable {
		t.Fatalf("consistent = %d, comparable = %d", report.Consistent, report.Comparable)
	}
	if report.Rate != "100.00%" {
		t.Fatalf("rate = %q, want 100.00%%", report.Rate)
	}
	for _, outcome := range report.Outcomes {
		if outcome.Disposition != DispositionComparable {
			continue
		}
		if !outcome.Consistent {
			t.Fatalf("point %s:%d %s is inconsistent: %s",
				outcome.Method, outcome.Line, outcome.Mutator, outcome.Difference)
		}
		if outcome.AbateVerdict == VerdictNotApplicable {
			t.Fatalf("comparable point %s:%d has no Abate verdict",
				outcome.Method, outcome.Line)
		}
	}
}

// TestExcludedPointsAreRegistered checks that every excluded point has a cause.
func TestExcludedPointsAreRegistered(t *testing.T) {
	t.Parallel()
	records := loadReference(t)
	report, err := Compare(records)
	if err != nil {
		t.Fatalf("compare: %v", err)
	}
	excluded := 0
	for _, outcome := range report.Outcomes {
		if outcome.Disposition != DispositionExcluded {
			continue
		}
		excluded++
		if strings.TrimSpace(outcome.Reason) == "" {
			t.Fatalf("excluded point %s:%d %s has no reason",
				outcome.Method, outcome.Line, outcome.Mutator)
		}
		if outcome.AbateVerdict != VerdictNotApplicable {
			t.Fatalf("excluded point %s:%d %s carries an Abate verdict",
				outcome.Method, outcome.Line, outcome.Mutator)
		}
	}
	if excluded != 7 {
		t.Fatalf("excluded points = %d, want 7", excluded)
	}
	if len(report.AbateOnly) != 1 || report.AbateOnly[0].Node != "i-timeout" {
		t.Fatalf("Abate only nodes = %v, want the i-timeout node", report.AbateOnly)
	}
}

// TestCompareIsDeterministic checks that the report is byte stable.
func TestCompareIsDeterministic(t *testing.T) {
	t.Parallel()
	records := loadReference(t)
	first, err := Compare(records)
	if err != nil {
		t.Fatalf("first compare: %v", err)
	}
	second, err := Compare(records)
	if err != nil {
		t.Fatalf("second compare: %v", err)
	}
	left, err := first.JSON()
	if err != nil {
		t.Fatalf("first JSON: %v", err)
	}
	right, err := second.JSON()
	if err != nil {
		t.Fatalf("second JSON: %v", err)
	}
	if string(left) != string(right) {
		t.Fatal("two comparisons produced different JSON reports")
	}
}

// TestJSONRoundTripKeepsEveryField covers the report encoding.
func TestJSONRoundTripKeepsEveryField(t *testing.T) {
	t.Parallel()
	records := loadReference(t)
	report, err := Compare(records)
	if err != nil {
		t.Fatalf("compare: %v", err)
	}
	payload, err := report.JSON()
	if err != nil {
		t.Fatalf("JSON: %v", err)
	}
	var decoded Report
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if decoded.Total != report.Total || decoded.Consistent != report.Consistent {
		t.Fatalf("round trip lost the counters: %+v", decoded)
	}
	if len(decoded.Outcomes) != len(report.Outcomes) {
		t.Fatalf("round trip lost outcomes: %d vs %d",
			len(decoded.Outcomes), len(report.Outcomes))
	}
	if !strings.Contains(string(payload), `"rate": "100.00%"`) {
		t.Fatal("JSON report does not carry the consistency rate")
	}
}

// TestMarkdownCarriesTheEvidence checks the review document content.
func TestMarkdownCarriesTheEvidence(t *testing.T) {
	t.Parallel()
	records := loadReference(t)
	report, err := Compare(records)
	if err != nil {
		t.Fatalf("compare: %v", err)
	}
	document := string(report.Markdown())
	for _, needle := range []string{
		"一致率 | 100.00%",
		"golden 判例总数 | 28",
		"可比对点数 | 21",
		"RemoveConditionalMutator_ORDER_ELSE",
		"EmptyObjectReturnValsMutator",
		"i-timeout",
		"如实边界声明",
	} {
		if !strings.Contains(document, needle) {
			t.Fatalf("the review document does not mention %q", needle)
		}
	}
}

// TestMarkdownHandlesAnEmptyAbateSide covers the empty branch of the report.
func TestMarkdownHandlesAnEmptyAbateSide(t *testing.T) {
	t.Parallel()
	empty := Report{Rate: "N/A"}
	document := string(empty.Markdown())
	if !strings.Contains(document, "无。") {
		t.Fatal("the report does not state that no Abate only node exists")
	}
}

// itoa renders an int for the assertion messages.
func itoa(value int) string { return strconv.Itoa(value) }

// mutationReport runs the real mutation engine over the parity fixture.
func mutationReport(t *testing.T, target *fixtureTarget, suite *fixtureSuite) mutation.Report {
	t.Helper()
	registry, err := mutation.NewRegistryWith(nil)
	if err != nil {
		t.Fatalf("registry: %v", err)
	}
	engine, err := mutation.NewEngineWithRegistry(ParitySeed, registry)
	if err != nil {
		t.Fatalf("engine: %v", err)
	}
	report, err := engine.Run(context.Background(), target, suite)
	if err != nil {
		t.Fatalf("run: %v", err)
	}
	return report
}

// TestFixtureMirrorsTheKernelFixture pins the parity fixture shape.
//
// TestWitnessRejectsInconsistentCases covers the witness error branches.
func TestWitnessRejectsInconsistentCases(t *testing.T) {
	t.Parallel()
	broken := witnessCase{
		method: "x", line: 1, mutator: "m",
		original: baselineAdd, mutant: addSubtracted,
		inputs: []int{9, 8}, expected: []observation{"14"},
	}
	if _, err := classifyWitness(broken); !errors.Is(err, ErrBaselineMismatch) {
		t.Fatalf("error = %v, want ErrBaselineMismatch", err)
	}
	wrongValue := witnessCase{
		method: "x", line: 1, mutator: "m",
		original: baselineAdd, mutant: addSubtracted,
		inputs: []int{9}, expected: []observation{"15"},
	}
	if _, err := classifyWitness(wrongValue); !errors.Is(err, ErrBaselineMismatch) {
		t.Fatalf("error = %v, want ErrBaselineMismatch", err)
	}
}

// TestFormatRateCoversTheEmptyCase covers the rate formatting branches.
func TestFormatRateCoversTheEmptyCase(t *testing.T) {
	t.Parallel()
	if got := formatRate(0, 0); got != "N/A" {
		t.Fatalf("rate = %q, want N/A", got)
	}
	if got := formatRate(2, 3); got != "66.66%" {
		t.Fatalf("rate = %q, want 66.66%%", got)
	}
	if got := formatRate(21, 21); got != "100.00%" {
		t.Fatalf("rate = %q, want 100.00%%", got)
	}
}

// TestCompareRejectsUnmappedCase covers the regression failure branch.
func TestCompareRejectsUnmappedCase(t *testing.T) {
	t.Parallel()
	cases := []Case{{
		Method: "nowhere", Mutator: "NoMutator", Line: 42, Expect: ExpectKilled,
	}}
	if _, err := Compare(cases); !errors.Is(err, ErrUnmappedMutation) {
		t.Fatalf("error = %v, want ErrUnmappedMutation", err)
	}
}

// TestCompareCaseReportsMismatch covers the cross-check difference path.
func TestCompareCaseReportsMismatch(t *testing.T) {
	t.Parallel()
	target, suite := newFixture()
	index, err := newVerdictIndex(mutationReport(t, target, suite))
	if err != nil {
		t.Fatalf("index: %v", err)
	}
	wrong := Case{
		Method: "add", Mutator: "MathMutator", Line: 15, Expect: ExpectSurvived,
	}
	outcome, err := compareCase(wrong, index, suite)
	if err != nil {
		t.Fatalf("compare: %v", err)
	}
	if outcome.Consistent {
		t.Fatal("a mismatched expectation was accepted")
	}
	if !strings.Contains(outcome.Difference, "语义见证") {
		t.Fatalf("difference = %q", outcome.Difference)
	}
	if _, err := compareCase(wrong, verdictIndex{}, suite); !errors.Is(err, ErrMissingVerdict) {
		t.Fatalf("error = %v, want ErrMissingVerdict", err)
	}
}

// TestMarkdownCarriesDifference covers the difference rendering branch.
func TestMarkdownCarriesDifference(t *testing.T) {
	t.Parallel()
	report := Report{
		Rate: "50.00%",
		Outcomes: []Outcome{
			{
				Method: "add", Line: 15, Mutator: "MathMutator", Expect: ExpectKilled,
				Node: "a-arith", Class: "arith-replace", AbateVerdict: "SURVIVED",
				Witness: WitnessDetected, VerdictMode: VerdictModePoint,
				Disposition: DispositionComparable, Reason: "reason", Difference: "差异文本",
			},
			{
				Method: "cap", Line: 29, Mutator: "Bucket", Expect: ExpectKilled,
				Disposition: DispositionExcluded, Reason: "excluded reason",
			},
		},
		AbateOnly: []Outcome{{
			Node: "i-timeout", AbateVerdict: "TIMEOUT", Reason: "orphan",
		}},
	}
	document := string(report.Markdown())
	if !strings.Contains(document, "差异：差异文本") {
		t.Fatal("the review document does not carry the mismatch")
	}
	if !strings.Contains(document, "excluded reason") {
		t.Fatal("the review document does not carry the exclusion reason")
	}
}
