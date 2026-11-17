<!--
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
-->

# PIT 对拍报告（KRN-MUT 变异核 ↔ PITest 1.19.6）

> 本报告由 `scripts/pit-parity.sh` 离线复跑生成，内容全部来自入库工件与
> Abate 变异核的真实引擎输出，不含时间戳，逐字节可复现。

## 1. 结论

| 指标 | 数值 |
|---|---|
| PIT 版本 | 1.19.6 |
| PIT 变异总数 | 28 |
| 可比对点数 | 21 |
| 一致点数 | 21 |
| 一致率 | 100.00% |
| 排除点数（PIT 专属算子） | 7 |

PIT 参考产物生成命令：`mvn -q org.pitest:pitest-maven:1.19.6:mutationCoverage -Dmutators=DEFAULTS (cwd: kernel/mutation/testdata/pit/baseline)`

## 2. 对拍口径

三方独立判定逐点互校，任一方不一致即记为差异并定位根因：

1. **PIT 官方状态**：入库工件 `kernel/mutation/testdata/pit/pit-mutations.json` 中 PITest 1.19.6 的真实逐点结论。
2. **Abate 引擎判定**：`kernel/mutation` 的真实引擎 （`NewRegistryWith` + `NewEngineWithRegistry` + `Engine.Run`）在同构夹具上产出的判定。
3. **语义见证**：按 Java 基准工程语义独立判定该变异在套件内是否可区分，该判定不看 PIT 状态，用于反查映射表是否被结果反向拟合。

状态映射口径：

- PIT `KILLED` ↔ Abate `KILLED`；
- PIT `SURVIVED` 且语义见证判定为套件内不可区分 ↔ Abate `EQUIVALENT`（Abate 夹具用 `h-equivalent` 节点显式承载语义等价变异）；
- PIT `NO_COVERAGE` ↔ Abate 覆盖计数为 0 且未被杀死，即 `SURVIVED`（Abate 夹具用 `j-survived` 节点显式承载未覆盖变异）。

## 3. 逐点明细

表中 # 为参考报告中的行序号，与第 4 节的排除项编号一致。

| # | 方法 | 行 | PIT 算子 | PIT 状态 | Abate 节点 | Abate 算子类 | Abate 判定 | 语义见证 | 比对口径 | 一致 |
|---|---|---|---|---|---|---|---|---|---|---|
| 1 | accepts | 49 | ConditionalsBoundaryMutator | KILLED | b-bound | boundary-value | KILLED | detected | point | 是 |
| 3 | accepts | 49 | BooleanTrueReturnValsMutator | KILLED | e-return | return-value | KILLED | detected | point | 是 |
| 4 | add | 15 | MathMutator | KILLED | a-arith | arith-replace | KILLED | detected | point | 是 |
| 5 | add | 15 | PrimitiveReturnsMutator | KILLED | c-const | constant-perturb | KILLED | detected | point | 是 |
| 6 | audited | 63 | ConditionalsBoundaryMutator | KILLED | b-bound | boundary-value | KILLED | detected | point | 是 |
| 8 | audited | 64 | VoidMethodCallMutator | KILLED | g-call | remove-call | KILLED | detected | point | 是 |
| 9 | cap | 29 | ConditionalsBoundaryMutator | SURVIVED | h-equivalent | boundary-value | EQUIVALENT | equivalent | node | 是 |
| 11 | cap | 30 | PrimitiveReturnsMutator | KILLED | c-const | constant-perturb | KILLED | detected | point | 是 |
| 12 | cap | 32 | PrimitiveReturnsMutator | KILLED | c-const | constant-perturb | KILLED | detected | point | 是 |
| 13 | clamp | 20 | ConditionalsBoundaryMutator | SURVIVED | h-equivalent | boundary-value | EQUIVALENT | equivalent | node | 是 |
| 15 | clamp | 21 | PrimitiveReturnsMutator | KILLED | c-const | constant-perturb | KILLED | detected | point | 是 |
| 16 | clamp | 23 | PrimitiveReturnsMutator | KILLED | c-const | constant-perturb | KILLED | detected | point | 是 |
| 17 | label | 54 | RemoveConditionalMutator_EQUAL_ELSE | KILLED | f-negate | negate-condition | KILLED | detected | point | 是 |
| 20 | log | 74 | MathMutator | KILLED | a-arith | arith-replace | KILLED | detected | point | 是 |
| 21 | logCalls | 79 | PrimitiveReturnsMutator | KILLED | c-const | constant-perturb | KILLED | detected | point | 是 |
| 22 | sumUntil | 38 | ConditionalsBoundaryMutator | SURVIVED | h-equivalent | boundary-value | EQUIVALENT | equivalent | node | 是 |
| 23 | sumUntil | 42 | MathMutator | KILLED | a-arith | arith-replace | KILLED | detected | point | 是 |
| 24 | sumUntil | 39 | RemoveConditionalMutator_EQUAL_ELSE | KILLED | d-control | control-flow | KILLED | detected | point | 是 |
| 26 | sumUntil | 44 | PrimitiveReturnsMutator | KILLED | c-const | constant-perturb | KILLED | detected | point | 是 |
| 27 | uncovered | 70 | MathMutator | NO_COVERAGE | j-survived | arith-replace | SURVIVED | uncovered | point | 是 |
| 28 | uncovered | 70 | PrimitiveReturnsMutator | NO_COVERAGE | j-survived | constant-perturb | SURVIVED | uncovered | node | 是 |

映射说明与差异明细：

- **#1 accepts:49 ConditionalsBoundaryMutator** — score >= 5 的比较边界被平移，与 Abate b-bound 的边界平移算子同族；套件断言 accepts(5) 可区分该变异，故为可检出变异。
- **#3 accepts:49 BooleanTrueReturnValsMutator** — 布尔返回值被替换为 true，与 Abate e-return 的布尔返回值翻转同族。
- **#4 add:15 MathMutator** — 整数加法被替换为减法，与 Abate a-arith 的 + -> - 完全同构。
- **#5 add:15 PrimitiveReturnsMutator** — 整型返回值被置零，属数值结果扰动，与 Abate constant-perturb 同族。
- **#6 audited:63 ConditionalsBoundaryMutator** — 循环边界 i < n 被平移，与 Abate b-bound 的边界平移算子同族；日志计数器可区分该变异，故为可检出变异。
- **#8 audited:64 VoidMethodCallMutator** — void 调用 log() 被删除，与 Abate g-call 的调用删除完全同构。
- **#9 cap:29 ConditionalsBoundaryMutator** — 边界平移后 cap(9)/cap(99)/cap(6) 的返回值不变，套件内不可区分，对应 Abate h-equivalent 的语义等价变异节点。
- **#11 cap:30 PrimitiveReturnsMutator** — 返回常量 limit 的整型返回被置零，与 Abate constant-perturb 的数值扰动同族。
- **#12 cap:32 PrimitiveReturnsMutator** — 整型返回值被置零，与 Abate constant-perturb 的数值扰动同族。
- **#13 clamp:20 ConditionalsBoundaryMutator** — 边界平移后 clamp(10)/clamp(15)/clamp(4) 的返回值不变，套件内不可区分，对应 Abate h-equivalent 的语义等价变异节点。
- **#15 clamp:21 PrimitiveReturnsMutator** — 整型返回值被置零，与 Abate constant-perturb 的数值扰动同族。
- **#16 clamp:23 PrimitiveReturnsMutator** — 整型返回值被置零，与 Abate constant-perturb 的数值扰动同族。
- **#17 label:54 RemoveConditionalMutator_EQUAL_ELSE** — 相等判定被删除，与 Abate f-negate 的 == -> != 完全同构：label(1) 由 one 变为 other。
- **#20 log:74 MathMutator** — 计数器自增被替换为自减，与 Abate a-arith 的 + -> - 同族。
- **#21 logCalls:79 PrimitiveReturnsMutator** — 整型返回值被置零，与 Abate constant-perturb 的数值扰动同族。
- **#22 sumUntil:38 ConditionalsBoundaryMutator** — 循环边界平移后 sumUntil(3)/sumUntil(10) 的累加值不变，套件内不可区分，对应 Abate h-equivalent 的语义等价变异节点。
- **#23 sumUntil:42 MathMutator** — 累加被替换为累减，与 Abate a-arith 的 + -> - 同族。
- **#24 sumUntil:39 RemoveConditionalMutator_EQUAL_ELSE** — break 守卫 i == 3 被删除，等价于删除 break 语句本身，与 Abate d-control 的语句删除完全同构。
- **#26 sumUntil:44 PrimitiveReturnsMutator** — 整型返回值被置零，与 Abate constant-perturb 的数值扰动同族。
- **#27 uncovered:70 MathMutator** — x * 3 被替换为 x / 3，与 Abate j-survived 的 * -> / 完全同构；该方法无任何测试覆盖，按覆盖计数 0 判定为未被杀死。
- **#28 uncovered:70 PrimitiveReturnsMutator** — 整型返回值被置零；该方法无任何测试覆盖，按覆盖计数 0 判定为未被杀死。节点由覆盖规则选定，Abate 侧该节点只有 arith-replace 变异点，故本点按节点级判定比对，不参与算子类级比对。

## 4. 被排除的 PIT 专属算子

下列变异**不参与一致率计算**，逐条给出排除原因，不静默丢弃、不凑数。

| # | 方法 | 行 | PIT 算子 | PIT 状态 | 语义见证 | 排除原因 |
|---|---|---|---|---|---|---|
| 2 | accepts | 49 | RemoveConditionalMutator_ORDER_ELSE | KILLED | detected | PIT 专属算子：顺序比较条件被整体删除并以 false 取代。Abate 七类算子只覆盖边界平移（boundary-value）、等式取反（negate-condition）与语句删除（control-flow），没有「删除顺序判定」语义，故不参与对拍。 |
| 7 | audited | 63 | RemoveConditionalMutator_ORDER_ELSE | KILLED | detected | PIT 专属算子：顺序比较条件被整体删除并以 false 取代，Abate 没有对应的算子语义，故不参与对拍。 |
| 10 | cap | 29 | RemoveConditionalMutator_ORDER_ELSE | KILLED | detected | PIT 专属算子：顺序比较条件被整体删除并以 false 取代，Abate 没有对应的算子语义，故不参与对拍。 |
| 14 | clamp | 20 | RemoveConditionalMutator_ORDER_ELSE | KILLED | detected | PIT 专属算子：顺序比较条件被整体删除并以 false 取代，Abate 没有对应的算子语义，故不参与对拍。 |
| 18 | label | 55 | EmptyObjectReturnValsMutator | KILLED | detected | PIT 专属算子：字符串返回值被替换为空串；Abate return-value 只翻转布尔与 nil 返回值，没有字符串返回类，故不参与对拍。 |
| 19 | label | 57 | EmptyObjectReturnValsMutator | KILLED | detected | PIT 专属算子：字符串返回值被替换为空串；Abate return-value 只翻转布尔与 nil 返回值，没有字符串返回类，故不参与对拍。 |
| 25 | sumUntil | 38 | RemoveConditionalMutator_ORDER_ELSE | KILLED | detected | PIT 专属算子：顺序比较条件被整体删除并以 false 取代，Abate 没有对应的算子语义，故不参与对拍。 |

## 5. Abate 侧无 PIT 对应点的节点

| Abate 节点 | Abate 判定 | 说明 |
|---|---|---|
| i-timeout | TIMEOUT | 该 Abate 节点没有被任何 PIT 变异映射（PITest 1.19.6 没有对应的变异状态分类），不参与一致率计算。 |

## 6. 如实边界声明

1. **未启动 JVM**：本轮对拍不回放 PIT 进程，而是使用入库的官方实跑产物（`pit-mutations.json` / `pit-mutations.xml` / `pit-console-report.txt`）；`scripts/pit-parity.sh --rerun-pit` 才会真实重跑 PITest。
2. **基准工程覆盖面有限**：PIT 基准工程只有一个类、三个测试、十个方法，28 条变异覆盖七类算子族，不等于 PITest 全量算子族。
3. **映射表非一一覆盖**：PIT 的 7 条变异落在 Abate 七类算子没有语义的算子族上（顺序判定删除、字符串返回置空），已逐条登记为排除项。
4. **语义等价的口径**：PITest 不识别语义等价变异，只报告 `SURVIVED`；本报告的等价结论由语义见证给出，并在第 2 节写明映射口径。
5. **夹具同构**：Abate 侧夹具与 `kernel/mutation/helpers_test.go` 的夹具同构（十节点、七类算子、等价/超时/未覆盖三条支路），两侧夹具若发生漂移，两侧各自的单元测试会先行失败。
