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
)

// ParitySeed is the fixed seed of the Abate mutation run.
//
// The seed only drives the deterministic point selection of the engine, the
// verdicts do not depend on it, and a fixed seed keeps the golden byte stable.
const ParitySeed = 20261111

// rateScale is the permyriad scale used by formatRate: consistent * rateScale
// / comparable yields a value whose integer part is the percent and whose last
// two digits are the fractional part, rendered as "NN.NN%".
const rateScale = 10000

// buildCases derives the self-developed cases from the mapping and witness
// tables. The expected classification comes from the independent witness, not
// from the engine, so the golden is a genuine two-sided observation.
func buildCases() ([]Case, error) {
	cases := make([]Case, 0, len(mappingTable))
	for _, entry := range mappingTable {
		sample, ok := witnessFor(entry.Method, entry.Mutator, entry.Line)
		if !ok {
			return nil, fmt.Errorf("%w: %s:%d %s has no semantic witness",
				ErrUnmappedMutation, entry.Method, entry.Line, entry.Mutator)
		}
		witness, err := classifyWitness(sample)
		if err != nil {
			return nil, err
		}
		cases = append(cases, Case{
			Method:  entry.Method,
			Mutator: entry.Mutator,
			Line:    entry.Line,
			Expect:  expectFromWitness(witness),
		})
	}
	return cases, nil
}

// buildRegistry builds the seven class registry the regression run uses.
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

// formatRate renders the consistency rate with integer arithmetic only.
func formatRate(consistent, comparable int) string {
	if comparable == 0 {
		return "N/A"
	}
	basis := consistent * rateScale / comparable
	return strconv.Itoa(basis/100) + "." + fmt.Sprintf("%02d", basis%100) + "%"
}
