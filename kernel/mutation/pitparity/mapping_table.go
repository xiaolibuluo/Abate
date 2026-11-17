/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package pitparity

// mappingTable registers the semantic correspondence between every PIT
// mutation of the reference report and the Abate fixture.
//
// The entries follow the order of the reference report so the review document
// can be read side by side with the vendored JSON. Every entry states why the
// two mutations are the same mutation; the seven PIT specific operator
// families that Abate does not model are registered as excluded with their
// cause, so no mutation is silently dropped.
//
// The node of a comparable entry is not chosen by hand: the test suite
// verifies that it equals the node the semantic witness derives, which rules
// out a mapping table tuned to the wanted outcome.
var mappingTable = []Mapping{
	{
		Method: "accepts", Mutator: "ConditionalsBoundaryMutator", Line: 49,
		Node: "b-bound", Class: "boundary-value",
		Reason: "score >= 5 的比较边界被平移，与 Abate b-bound 的边界平移算子同族；" +
			"套件断言 accepts(5) 可区分该变异，故为可检出变异。",
	},
	{
		Method: "accepts", Mutator: "RemoveConditionalMutator_ORDER_ELSE", Line: 49,
		Excluded: true,
		Reason: "PIT 专属算子：顺序比较条件被整体删除并以 false 取代。" +
			"Abate 七类算子只覆盖边界平移（boundary-value）、等式取反（negate-condition）" +
			"与语句删除（control-flow），没有「删除顺序判定」语义，故不参与对拍。",
	},
	{
		Method: "accepts", Mutator: "BooleanTrueReturnValsMutator", Line: 49,
		Node: "e-return", Class: "return-value",
		Reason: "布尔返回值被替换为 true，与 Abate e-return 的布尔返回值翻转同族。",
	},
	{
		Method: "add", Mutator: "MathMutator", Line: 15,
		Node: "a-arith", Class: "arith-replace",
		Reason: "整数加法被替换为减法，与 Abate a-arith 的 + -> - 完全同构。",
	},
	{
		Method: "add", Mutator: "PrimitiveReturnsMutator", Line: 15,
		Node: "c-const", Class: "constant-perturb",
		Reason: "整型返回值被置零，属数值结果扰动，与 Abate constant-perturb 同族。",
	},
	{
		Method: "audited", Mutator: "ConditionalsBoundaryMutator", Line: 63,
		Node: "b-bound", Class: "boundary-value",
		Reason: "循环边界 i < n 被平移，与 Abate b-bound 的边界平移算子同族；" +
			"日志计数器可区分该变异，故为可检出变异。",
	},
	{
		Method: "audited", Mutator: "RemoveConditionalMutator_ORDER_ELSE", Line: 63,
		Excluded: true,
		Reason: "PIT 专属算子：顺序比较条件被整体删除并以 false 取代，" +
			"Abate 没有对应的算子语义，故不参与对拍。",
	},
	{
		Method: "audited", Mutator: "VoidMethodCallMutator", Line: 64,
		Node: "g-call", Class: "remove-call",
		Reason: "void 调用 log() 被删除，与 Abate g-call 的调用删除完全同构。",
	},
	{
		Method: "cap", Mutator: "ConditionalsBoundaryMutator", Line: 29,
		Node: "h-equivalent", Class: "boundary-value",
		Reason: "边界平移后 cap(9)/cap(99)/cap(6) 的返回值不变，套件内不可区分，" +
			"对应 Abate h-equivalent 的语义等价变异节点。",
	},
	{
		Method: "cap", Mutator: "RemoveConditionalMutator_ORDER_ELSE", Line: 29,
		Excluded: true,
		Reason: "PIT 专属算子：顺序比较条件被整体删除并以 false 取代，" +
			"Abate 没有对应的算子语义，故不参与对拍。",
	},
	{
		Method: "cap", Mutator: "PrimitiveReturnsMutator", Line: 30,
		Node: "c-const", Class: "constant-perturb",
		Reason: "返回常量 limit 的整型返回被置零，与 Abate constant-perturb 的数值扰动同族。",
	},
	{
		Method: "cap", Mutator: "PrimitiveReturnsMutator", Line: 32,
		Node: "c-const", Class: "constant-perturb",
		Reason: "整型返回值被置零，与 Abate constant-perturb 的数值扰动同族。",
	},
	{
		Method: "clamp", Mutator: "ConditionalsBoundaryMutator", Line: 20,
		Node: "h-equivalent", Class: "boundary-value",
		Reason: "边界平移后 clamp(10)/clamp(15)/clamp(4) 的返回值不变，套件内不可区分，" +
			"对应 Abate h-equivalent 的语义等价变异节点。",
	},
	{
		Method: "clamp", Mutator: "RemoveConditionalMutator_ORDER_ELSE", Line: 20,
		Excluded: true,
		Reason: "PIT 专属算子：顺序比较条件被整体删除并以 false 取代，" +
			"Abate 没有对应的算子语义，故不参与对拍。",
	},
	{
		Method: "clamp", Mutator: "PrimitiveReturnsMutator", Line: 21,
		Node: "c-const", Class: "constant-perturb",
		Reason: "整型返回值被置零，与 Abate constant-perturb 的数值扰动同族。",
	},
	{
		Method: "clamp", Mutator: "PrimitiveReturnsMutator", Line: 23,
		Node: "c-const", Class: "constant-perturb",
		Reason: "整型返回值被置零，与 Abate constant-perturb 的数值扰动同族。",
	},
	{
		Method: "label", Mutator: "RemoveConditionalMutator_EQUAL_ELSE", Line: 54,
		Node: "f-negate", Class: "negate-condition",
		Reason: "相等判定被删除，与 Abate f-negate 的 == -> != 完全同构：" +
			"label(1) 由 one 变为 other。",
	},
	{
		Method: "label", Mutator: "EmptyObjectReturnValsMutator", Line: 55,
		Excluded: true,
		Reason: "PIT 专属算子：字符串返回值被替换为空串；Abate return-value 只翻转" +
			"布尔与 nil 返回值，没有字符串返回类，故不参与对拍。",
	},
	{
		Method: "label", Mutator: "EmptyObjectReturnValsMutator", Line: 57,
		Excluded: true,
		Reason: "PIT 专属算子：字符串返回值被替换为空串；Abate return-value 只翻转" +
			"布尔与 nil 返回值，没有字符串返回类，故不参与对拍。",
	},
	{
		Method: "log", Mutator: "MathMutator", Line: 74,
		Node: "a-arith", Class: "arith-replace",
		Reason: "计数器自增被替换为自减，与 Abate a-arith 的 + -> - 同族。",
	},
	{
		Method: "logCalls", Mutator: "PrimitiveReturnsMutator", Line: 79,
		Node: "c-const", Class: "constant-perturb",
		Reason: "整型返回值被置零，与 Abate constant-perturb 的数值扰动同族。",
	},
	{
		Method: "sumUntil", Mutator: "ConditionalsBoundaryMutator", Line: 38,
		Node: "h-equivalent", Class: "boundary-value",
		Reason: "循环边界平移后 sumUntil(3)/sumUntil(10) 的累加值不变，套件内不可区分，" +
			"对应 Abate h-equivalent 的语义等价变异节点。",
	},
	{
		Method: "sumUntil", Mutator: "MathMutator", Line: 42,
		Node: "a-arith", Class: "arith-replace",
		Reason: "累加被替换为累减，与 Abate a-arith 的 + -> - 同族。",
	},
	{
		Method: "sumUntil", Mutator: "RemoveConditionalMutator_EQUAL_ELSE", Line: 39,
		Node: "d-control", Class: "control-flow",
		Reason: "break 守卫 i == 3 被删除，等价于删除 break 语句本身，" +
			"与 Abate d-control 的语句删除完全同构。",
	},
	{
		Method: "sumUntil", Mutator: "RemoveConditionalMutator_ORDER_ELSE", Line: 38,
		Excluded: true,
		Reason: "PIT 专属算子：顺序比较条件被整体删除并以 false 取代，" +
			"Abate 没有对应的算子语义，故不参与对拍。",
	},
	{
		Method: "sumUntil", Mutator: "PrimitiveReturnsMutator", Line: 44,
		Node: "c-const", Class: "constant-perturb",
		Reason: "整型返回值被置零，与 Abate constant-perturb 的数值扰动同族。",
	},
	{
		Method: "uncovered", Mutator: "MathMutator", Line: 70,
		Node: "j-survived", Class: "arith-replace",
		Reason: "x * 3 被替换为 x / 3，与 Abate j-survived 的 * -> / 完全同构；" +
			"该方法无任何测试覆盖，按覆盖计数 0 判定为未被杀死。",
	},
	{
		Method: "uncovered", Mutator: "PrimitiveReturnsMutator", Line: 70,
		Node: "j-survived", Class: "constant-perturb",
		Reason: "整型返回值被置零；该方法无任何测试覆盖，按覆盖计数 0 判定为未被杀死。" +
			"节点由覆盖规则选定，Abate 侧该节点只有 arith-replace 变异点，" +
			"故本点按节点级判定比对，不参与算子类级比对。",
	},
}
