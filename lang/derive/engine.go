/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package derive

import (
	"context"
	"fmt"
	"slices"
	"strings"

	"abate/kernel/mutation"
	"abate/lang/air"
	"abate/lang/rule"
)

// experimentTests returns the fixed, ordered static assertion suite.
func experimentTests() []string {
	return []string{"t1", "t2", "t3"}
}

// RunExperiment wires derived operators into the mutation kernel and runs one
// deterministic experiment with an explicitly supplied target and suite.
func RunExperiment(seed uint64, target mutation.Target, suite mutation.Suite,
	specs []MutationSpec,
) (mutation.Report, error) {
	ops, err := Operators(specs)
	if err != nil {
		return mutation.Report{}, err
	}
	reg, err := mutation.NewRegistryWith(ops)
	if err != nil {
		return mutation.Report{}, fmt.Errorf("derive: registry: %w", err)
	}
	engine, err := mutation.NewEngineWithRegistry(seed, reg)
	if err != nil {
		return mutation.Report{}, err
	}
	rep, err := engine.Run(context.Background(), target, suite)
	if err != nil {
		return mutation.Report{}, err
	}
	if !hasDerivedPoints(rep) {
		return mutation.Report{}, ErrNoDerivedPoints
	}
	return rep, nil
}

// Experiment runs the full AIR -> rules -> derivation -> engine pipeline and
// returns the deterministic report together with the derived specifications.
func Experiment(seed uint64, unit *air.Unit, rules []rule.Rule) (mutation.Report, []MutationSpec, error) {
	specs, err := NewDeriver().Derive(rules)
	if err != nil {
		return mutation.Report{}, nil, err
	}
	tests := experimentTests()
	target, err := NewAIRTarget(unit, specs, tests)
	if err != nil {
		return mutation.Report{}, nil, err
	}
	suite, err := NewAIRSuite(target.Nodes(), tests)
	if err != nil {
		return mutation.Report{}, nil, err
	}
	rep, err := RunExperiment(seed, target, suite, specs)
	if err != nil {
		return mutation.Report{}, nil, err
	}
	return rep, specs, nil
}

// hasDerivedPoints reports whether the report carries at least one m-derived-*
// mutant with an allocated verdict.
func hasDerivedPoints(rep mutation.Report) bool {
	for _, point := range rep.Points {
		if point.Operator == nil {
			continue
		}
		if strings.HasPrefix(string(point.Operator.ID()), derivedPrefix) {
			return true
		}
	}
	return false
}

// DerivedIDs returns the sorted distinct derived operator identifiers present
// in a report, which is the enumerable evidence for the M3 acceptance item.
func DerivedIDs(rep mutation.Report) []string {
	seen := make(map[string]bool, len(rep.Points))
	for _, point := range rep.Points {
		if point.Operator == nil {
			continue
		}
		id := string(point.Operator.ID())
		if strings.HasPrefix(id, derivedPrefix) {
			seen[id] = true
		}
	}
	out := make([]string, 0, len(seen))
	for id := range seen {
		out = append(out, id)
	}
	slices.Sort(out)
	return out
}
