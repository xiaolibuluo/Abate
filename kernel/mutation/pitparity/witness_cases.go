/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package pitparity

// witnessTable registers the semantic witness of every self-developed case.
//
// The covered inputs and the expected values are the fixed assertions of the
// baseline port, so the witness reproduces local semantics independently of
// the frozen expectation.
var witnessTable = []witnessCase{
	{
		method: "accepts", line: 49, mutator: "ConditionalsBoundaryMutator",
		original: baselineAccepts, mutant: acceptsBoundary,
		inputs: []int{5, 4}, expected: []observation{"true", "false"},
	},
	{
		method: "accepts", line: 49, mutator: "RemoveConditionalMutator_ORDER_ELSE",
		original: baselineAccepts, mutant: acceptsOrderElse,
		inputs: []int{5, 4}, expected: []observation{"true", "false"},
	},
	{
		method: "accepts", line: 49, mutator: "BooleanTrueReturnValsMutator",
		original: baselineAccepts, mutant: acceptsTrueReturn,
		inputs: []int{5, 4}, expected: []observation{"true", "false"},
	},
	{
		method: "add", line: 15, mutator: "MathMutator",
		original: baselineAdd, mutant: addSubtracted,
		inputs: []int{9}, expected: []observation{"14"},
	},
	{
		method: "add", line: 15, mutator: "PrimitiveReturnsMutator",
		original: baselineAdd, mutant: addZeroReturn,
		inputs: []int{9}, expected: []observation{"14"},
	},
	{
		method: "audited", line: 63, mutator: "ConditionalsBoundaryMutator",
		original: baselineAudited, mutant: auditedBoundary,
		inputs: []int{2}, expected: []observation{"2"},
	},
	{
		method: "audited", line: 63, mutator: "RemoveConditionalMutator_ORDER_ELSE",
		original: baselineAudited, mutant: auditedOrderElse,
		inputs: []int{2}, expected: []observation{"2"},
	},
	{
		method: "audited", line: 64, mutator: "VoidMethodCallMutator",
		original: baselineAudited, mutant: auditedNoCall,
		inputs: []int{2}, expected: []observation{"2"},
	},
	{
		method: "cap", line: 29, mutator: "ConditionalsBoundaryMutator",
		original: baselineCap, mutant: capBoundary,
		inputs: []int{9, 99, 6}, expected: []observation{"7", "7", "6"},
	},
	{
		method: "cap", line: 29, mutator: "RemoveConditionalMutator_ORDER_ELSE",
		original: baselineCap, mutant: capOrderElse,
		inputs: []int{9, 99, 6}, expected: []observation{"7", "7", "6"},
	},
	{
		method: "cap", line: 30, mutator: "PrimitiveReturnsMutator",
		original: baselineCap, mutant: capZeroReturn,
		inputs: []int{9, 99, 6}, expected: []observation{"7", "7", "6"},
	},
	{
		method: "cap", line: 32, mutator: "PrimitiveReturnsMutator",
		original: baselineCap, mutant: capZeroReturn,
		inputs: []int{9, 99, 6}, expected: []observation{"7", "7", "6"},
	},
	{
		method: "clamp", line: 20, mutator: "ConditionalsBoundaryMutator",
		original: baselineClamp, mutant: clampBoundary,
		inputs: []int{10, 15, 4}, expected: []observation{"10", "10", "4"},
	},
	{
		method: "clamp", line: 20, mutator: "RemoveConditionalMutator_ORDER_ELSE",
		original: baselineClamp, mutant: clampOrderElse,
		inputs: []int{10, 15, 4}, expected: []observation{"10", "10", "4"},
	},
	{
		method: "clamp", line: 21, mutator: "PrimitiveReturnsMutator",
		original: baselineClamp, mutant: clampZeroReturn,
		inputs: []int{10, 15, 4}, expected: []observation{"10", "10", "4"},
	},
	{
		method: "clamp", line: 23, mutator: "PrimitiveReturnsMutator",
		original: baselineClamp, mutant: clampZeroReturn,
		inputs: []int{10, 15, 4}, expected: []observation{"10", "10", "4"},
	},
	{
		method: "label", line: 54, mutator: "RemoveConditionalMutator_EQUAL_ELSE",
		original: baselineLabel, mutant: labelNoGuard,
		inputs: []int{1, 2}, expected: []observation{"one", "other"},
	},
	{
		method: "label", line: 55, mutator: "EmptyObjectReturnValsMutator",
		original: baselineLabel, mutant: labelEmptyReturn,
		inputs: []int{1, 2}, expected: []observation{"one", "other"},
	},
	{
		method: "label", line: 57, mutator: "EmptyObjectReturnValsMutator",
		original: baselineLabel, mutant: labelEmptyReturn,
		inputs: []int{1, 2}, expected: []observation{"one", "other"},
	},
	{
		method: "log", line: 74, mutator: "MathMutator",
		original: baselineAudited, mutant: logSubtracted,
		inputs: []int{2}, expected: []observation{"2"},
	},
	{
		method: "logCalls", line: 79, mutator: "PrimitiveReturnsMutator",
		original: baselineLogCalls, mutant: logCallsZeroReturn,
		inputs: []int{2}, expected: []observation{"2"},
	},
	{
		method: "sumUntil", line: 38, mutator: "ConditionalsBoundaryMutator",
		original: baselineSumUntil, mutant: sumUntilBoundary,
		inputs: []int{3, 10}, expected: []observation{"3", "3"},
	},
	{
		method: "sumUntil", line: 42, mutator: "MathMutator",
		original: baselineSumUntil, mutant: sumUntilSubtracted,
		inputs: []int{3, 10}, expected: []observation{"3", "3"},
	},
	{
		method: "sumUntil", line: 39, mutator: "RemoveConditionalMutator_EQUAL_ELSE",
		original: baselineSumUntil, mutant: sumUntilNoGuard,
		inputs: []int{3, 10}, expected: []observation{"3", "3"},
	},
	{
		method: "sumUntil", line: 38, mutator: "RemoveConditionalMutator_ORDER_ELSE",
		original: baselineSumUntil, mutant: sumUntilOrderElse,
		inputs: []int{3, 10}, expected: []observation{"3", "3"},
	},
	{
		method: "sumUntil", line: 44, mutator: "PrimitiveReturnsMutator",
		original: baselineSumUntil, mutant: sumUntilZeroReturn,
		inputs: []int{3, 10}, expected: []observation{"3", "3"},
	},
	{
		method: "uncovered", line: 70, mutator: "MathMutator",
		original: uncoveredOriginal, mutant: uncoveredDivided,
		inputs: nil, expected: nil,
	},
	{
		method: "uncovered", line: 70, mutator: "PrimitiveReturnsMutator",
		original: uncoveredOriginal, mutant: uncoveredZeroReturn,
		inputs: nil, expected: nil,
	},
}

// uncoveredOriginal observes uncovered, which no test of the suite reaches.
func uncoveredOriginal(x int) observation { return obsInt(x * 3) }

// witnessFor returns the witness case registered for one self-developed case.
func witnessFor(method, mutator string, line int) (witnessCase, bool) {
	for _, entry := range witnessTable {
		if entry.method == method && entry.mutator == mutator && entry.line == line {
			return entry, true
		}
	}
	return witnessCase{}, false
}
