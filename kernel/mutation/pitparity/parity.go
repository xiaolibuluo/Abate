/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package pitparity

import (
	"context"
	"fmt"
	"strconv"

	"abate/kernel/mutation"
	"abate/kernel/verify"
)

// ParitySeed is the fixed seed of the Abate side mutation run.
//
// The seed only drives the deterministic point selection of the engine, the
// verdicts do not depend on it, and a fixed seed keeps the report byte stable.
const ParitySeed = 20261111

// The dispositions of one parity point.
const (
	// DispositionComparable marks a point that takes part in the rate.
	DispositionComparable = "comparable"

	// DispositionExcluded marks a point without an Abate counterpart.
	DispositionExcluded = "excluded"
)

// rateScale is the permyriad scale used by formatRate: consistent * rateScale
// / comparable yields a value whose integer part is the percent and whose last
// two digits are the fractional part, rendered as "NN.NN%".
const rateScale = 10000

// The verdict modes of one parity point.
const (
	// VerdictModePoint means the engine point matched both node and class.
	VerdictModePoint = "point"

	// VerdictModeNode means the node level verdict was used instead.
	VerdictModeNode = "node"
)

// VerdictNotApplicable marks a point without an Abate counterpart.
const VerdictNotApplicable = "N/A"

// Outcome is the parity result of one PIT mutation record.
type Outcome struct {
	// Method is the mutated method of the Java baseline project.
	Method string `json:"method"`

	// Mutator is the PIT mutator family of the mutation.
	Mutator string `json:"mutator"`

	// Line is the mutated source line of the Java baseline project.
	Line int `json:"line"`

	// PITStatus is the status reported by PITest.
	PITStatus string `json:"pitStatus"`

	// AbateVerdict is the verdict produced by the Abate mutation engine.
	AbateVerdict string `json:"abateVerdict"`

	// Node is the Abate fixture node of the comparison.
	Node string `json:"node"`

	// Class is the Abate operator class of the comparison.
	Class string `json:"class"`

	// Disposition is comparable or excluded.
	Disposition string `json:"disposition"`

	// Witness is the independent semantic classification of the mutant.
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

// Compare replays the reference PIT report against the Abate mutation engine.
//
// The Abate side is produced by the real KRN-MUT engine over the parity
// fixture, the PIT side is the vendored official report. Every record is
// cross checked twice: the semantic witness must reproduce the PIT status and
// the Abate verdict must agree with that status under the documented mapping.
func Compare(records []Mutation) (Report, error) {
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
		PITVersion: PITVersion,
		Command:    PITCommand,
		Total:      len(records),
	}
	used := make(map[string]bool, len(target.Nodes()))
	for _, record := range records {
		outcome, err := compareRecord(record, index, suite)
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

// buildRegistry builds the seven class registry the parity run uses.
func buildRegistry() (*mutation.Registry, error) {
	registry, err := mutation.NewRegistryWith(nil)
	if err != nil {
		return nil, fmt.Errorf("pitparity: build registry: %w", err)
	}
	return registry, nil
}

// runEngine runs the mutation engine over one target and suite.
//
// The registry, target and suite are parameters so the failure branches of the
// kernel API stay reachable from the unit tests.
func runEngine(registry *mutation.Registry, target mutation.Target,
	suite mutation.Suite,
) (mutation.Report, error) {
	engine, err := mutation.NewEngineWithRegistry(ParitySeed, registry)
	if err != nil {
		return mutation.Report{}, fmt.Errorf("pitparity: build engine: %w", err)
	}
	run, err := engine.Run(context.Background(), target, suite)
	if err != nil {
		return mutation.Report{}, fmt.Errorf("pitparity: run engine: %w", err)
	}
	return run, nil
}

// compareRecord compares one PIT record with the Abate verdict of its node.
func compareRecord(record Mutation, index verdictIndex, suite *fixtureSuite) (Outcome, error) {
	mapping, witnessCase, err := resolveSite(record.MutatedMethod, record.Mutator, record.LineNumber)
	if err != nil {
		return Outcome{}, err
	}
	witness, err := classifyWitness(witnessCase)
	if err != nil {
		return Outcome{}, err
	}
	outcome := Outcome{
		Method:    record.MutatedMethod,
		Mutator:   record.Mutator,
		Line:      record.LineNumber,
		PITStatus: record.Status,
		Node:      mapping.Node,
		Class:     mapping.Class,
		Witness:   witness,
		Reason:    mapping.Reason,
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
	if !witnessMatchesStatus(witness, record.Status) {
		outcome.Consistent = false
		outcome.Difference = "语义见证判定为 " + witness +
			"，与 PIT 状态 " + record.Status + " 不一致"
	}
	if reason := verdictDifference(verdict, record.Status, witness, suite, mapping); reason != "" {
		outcome.Consistent = false
		if outcome.Difference == "" {
			outcome.Difference = reason
		} else {
			outcome.Difference += "；" + reason
		}
	}
	return outcome, nil
}

// verdictDifference reports why an Abate verdict disagrees with a PIT status.
//
// An empty result means the point agrees under the documented mapping:
// KILLED maps to KILLED, a suite indistinguishable mutant maps to EQUIVALENT
// and an unreachable site maps to SURVIVED with a zero coverage count.
func verdictDifference(verdict verify.MutationVerdict, status, witness string,
	suite *fixtureSuite, mapping Mapping,
) string {
	switch status {
	case StatusKilled:
		if verdict == verify.MutantKilled {
			return ""
		}
		return "PIT 判定 KILLED，Abate 引擎判定 " + string(verdict)
	case StatusSurvived:
		if witness != WitnessEquivalent {
			return "PIT 判定 SURVIVED，但语义见证判定为 " + witness
		}
		if verdict == verify.MutantEquivalent {
			return ""
		}
		return "PIT 判定 SURVIVED（套件内不可区分），Abate 引擎判定 " + string(verdict)
	case StatusNoCoverage:
		if suite.Coverage(mapping.Node) != 0 {
			return "PIT 判定 NO_COVERAGE，Abate 侧节点 " +
				mapping.Node + " 的覆盖计数不为 0"
		}
		if verdict == verify.MutantSurvived {
			return ""
		}
		return "PIT 判定 NO_COVERAGE，Abate 引擎判定 " + string(verdict)
	default:
		return "未知的 PIT 状态 " + status
	}
}

// abateOnlyNodes lists the Abate nodes no PIT mutation maps to.
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
			Reason: "该 Abate 节点没有被任何 PIT 变异映射（PITest 1.19.6 没有对应的" +
				"变异状态分类），不参与一致率计算。",
		})
	}
	return out
}

// formatRate renders the consistency rate with integer arithmetic only.
func formatRate(consistent, comparable int) string {
	if comparable == 0 {
		return "N/A"
	}
	basis := consistent * rateScale / comparable
	return strconv.Itoa(basis/100) + "." + fmt.Sprintf("%02d", basis%100) + "%"
}
