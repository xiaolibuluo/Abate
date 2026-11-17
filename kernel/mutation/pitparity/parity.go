/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package pitparity

import (
	"fmt"

	"abate/kernel/verify"
)

// The dispositions of one regression point.
const (
	// DispositionComparable marks a point that takes part in the rate.
	DispositionComparable = "comparable"

	// DispositionExcluded marks a point without an Abate counterpart.
	DispositionExcluded = "excluded"
)

// The verdict modes of one regression point.
const (
	// VerdictModePoint means the engine point matched both node and class.
	VerdictModePoint = "point"

	// VerdictModeNode means the node level verdict was used instead.
	VerdictModeNode = "node"
)

// VerdictNotApplicable marks a point without an Abate counterpart.
const VerdictNotApplicable = "N/A"

// goldenSource identifies the self-developed source of a report.
const goldenSource = "self-developed-golden"

// Outcome is the regression result of one self-developed golden case.
type Outcome struct {
	// Method is the local fixture method identifier of the case.
	Method string `json:"method"`

	// Mutator is the local fixture operator identifier of the case.
	Mutator string `json:"mutator"`

	// Line is the local fixture line identifier of the case.
	Line int `json:"line"`

	// Expect is the expected classification frozen in the golden.
	Expect string `json:"expect"`

	// AbateVerdict is the verdict produced by the Abate mutation engine.
	AbateVerdict string `json:"abateVerdict"`

	// Node is the Abate fixture node of the comparison.
	Node string `json:"node"`

	// Class is the Abate operator class of the comparison.
	Class string `json:"class"`

	// Disposition is comparable or excluded.
	Disposition string `json:"disposition"`

	// Witness is the independent semantic classification of the case.
	Witness string `json:"witness"`

	// VerdictMode is point when the engine point matched the class and node
	// when the node level verdict was used instead.
	VerdictMode string `json:"verdictMode"`

	// Consistent reports whether both cross checks agreed.
	Consistent bool `json:"consistent"`

	// Reason states the semantic correspondence of the mapping.
	Reason string `json:"reason"`

	// Difference describes the mismatch when the cross checks disagree.
	Difference string `json:"difference,omitempty"`
}

// Compare replays the frozen self-developed golden against the Abate engine.
//
// The expected side is the classification frozen in the golden, the actual
// side is produced by the real KRN-MUT engine over the parity fixture. Every
// case is cross checked twice: the independent semantic witness must reproduce
// the frozen expectation and the engine verdict must agree with it under the
// documented mapping.
func Compare(cases []Case) (Report, error) {
	target, suite := newFixture()
	registry, err := buildRegistry()
	if err != nil {
		return Report{}, err
	}
	run, err := runEngine(registry, target, suite)
	if err != nil {
		return Report{}, err
	}
	index, err := newVerdictIndex(run)
	if err != nil {
		return Report{}, err
	}
	report := Report{
		Source: goldenSource,
		Seed:   ParitySeed,
		Total:  len(cases),
	}
	used := make(map[string]bool, len(target.Nodes()))
	for _, entry := range cases {
		outcome, err := compareCase(entry, index, suite)
		if err != nil {
			return Report{}, err
		}
		report.Outcomes = append(report.Outcomes, outcome)
		if outcome.Disposition != DispositionComparable {
			continue
		}
		report.Comparable++
		if outcome.Consistent {
			report.Consistent++
		}
		used[outcome.Node] = true
	}
	report.Rate = formatRate(report.Consistent, report.Comparable)
	report.AbateOnly = abateOnlyNodes(target, index, used)
	return report, nil
}

// compareCase compares one golden case with the Abate verdict of its node.
func compareCase(entry Case, index verdictIndex, suite *fixtureSuite) (Outcome, error) {
	mapping, witnessCase, err := resolveSite(entry.Method, entry.Mutator, entry.Line)
	if err != nil {
		return Outcome{}, err
	}
	witness, err := classifyWitness(witnessCase)
	if err != nil {
		return Outcome{}, err
	}
	outcome := Outcome{
		Method:  entry.Method,
		Mutator: entry.Mutator,
		Line:    entry.Line,
		Expect:  entry.Expect,
		Node:    mapping.Node,
		Class:   mapping.Class,
		Witness: witness,
		Reason:  mapping.Reason,
	}
	if mapping.Excluded {
		outcome.Disposition = DispositionExcluded
		outcome.AbateVerdict = VerdictNotApplicable
		return outcome, nil
	}
	verdict, mode, ok := index.lookup(mapping)
	if !ok {
		return Outcome{}, fmt.Errorf("%w: node %q class %q",
			ErrMissingVerdict, mapping.Node, mapping.Class)
	}
	outcome.Disposition = DispositionComparable
	outcome.AbateVerdict = string(verdict)
	outcome.VerdictMode = mode
	outcome.Consistent = true
	if !witnessMatchesExpect(witness, entry.Expect) {
		outcome.Consistent = false
		outcome.Difference = "语义见证判定为 " + witness +
			"，与冻结期望分类 " + entry.Expect + " 不一致"
	}
	if reason := verdictDifference(verdict, entry.Expect, witness, suite, mapping); reason != "" {
		outcome.Consistent = false
		if outcome.Difference == "" {
			outcome.Difference = reason
		} else {
			outcome.Difference += "；" + reason
		}
	}
	return outcome, nil
}

// verdictDifference reports why an Abate verdict disagrees with the frozen
// expectation.
//
// An empty result means the point agrees under the documented mapping:
// KILLED maps to KILLED, a suite indistinguishable mutant maps to EQUIVALENT
// and an unreachable site maps to SURVIVED with a zero coverage count.
func verdictDifference(verdict verify.MutationVerdict, expect, witness string,
	suite *fixtureSuite, mapping Mapping,
) string {
	switch expect {
	case ExpectKilled:
		if verdict == verify.MutantKilled {
			return ""
		}
		return "冻结期望 KILLED，Abate 引擎判定 " + string(verdict)
	case ExpectSurvived:
		if witness != WitnessEquivalent {
			return "冻结期望 SURVIVED，但语义见证判定为 " + witness
		}
		if verdict == verify.MutantEquivalent {
			return ""
		}
		return "冻结期望 SURVIVED（套件内不可区分），Abate 引擎判定 " + string(verdict)
	case ExpectNoCoverage:
		if suite.Coverage(mapping.Node) != 0 {
			return "冻结期望 NO_COVERAGE，Abate 侧节点 " +
				mapping.Node + " 的覆盖计数不为 0"
		}
		if verdict == verify.MutantSurvived {
			return ""
		}
		return "冻结期望 NO_COVERAGE，Abate 引擎判定 " + string(verdict)
	default:
		return "未知的期望分类 " + expect
	}
}

// abateOnlyNodes lists the Abate nodes no golden case maps to.
func abateOnlyNodes(target *fixtureTarget, index verdictIndex, used map[string]bool) []Outcome {
	out := make([]Outcome, 0, len(target.Nodes()))
	for _, node := range target.Nodes() {
		if used[node] {
			continue
		}
		verdict := VerdictNotApplicable
		if value, ok := index.nodes[node]; ok {
			verdict = string(value)
		}
		out = append(out, Outcome{
			Node:         node,
			AbateVerdict: verdict,
			Disposition:  DispositionExcluded,
			Witness:      WitnessNotApplicable,
			Reason:       "该 Abate 节点没有被任何自研 golden 判例映射，不参与一致率计算。",
		})
	}
	return out
}
