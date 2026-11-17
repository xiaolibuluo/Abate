# py 前端 degrade 达标样板「真实性」独立验收报告

- 验收人：QA 工程师 严过关
- 对象：工程师寇豆码自报的 py 前端 G4 改造（19 个 KindPy* 109–127 真实接线、21 条规则升级 exact、degradePyPct 20.833333→0.000000）
- 验收纪律：只读审查 + 单包测试 + `/tmp` 实验程序；未修改仓库任何产品代码/冻结件/测试。本文件是本次验收唯一的仓库写入。
- 环境：go1.26.11，GOFLAGS=-mod=vendor。并行复刻 cs/kt/rs 期间，未跑全量 build/test/lint-gate。
- 实验产物：`/tmp/qa_py_probe/`（main.go、fabprobe.go、firstunit.go、zerorule.go 及对应 *_out.txt）

## 总览判定

| 项 | 判定 | 一句话结论 |
|---|---|---|
| A. 产出真实性 | **通过** | 独立 probe 22/22 PRESENT，工程师点名的 10 条全部真实产出 |
| B. 语义建模而非贴标签 | **通过（有残留）** | 19 个 dedicated() 全部发生在真实 reduce 现场；但上轮造假签名 pymap_extra.go 仍部分接线（见残留 1） |
| C. 口径与规则表一致性 | **基本通过（自报有失实）** | 21/21 全部 KindPy*、零指向 87-108；但"19 Kind 全语料全非零"不实：PyNamedExpr(113)=0 |
| D. 防造假测试完整性 | **通过** | 4 个测试 21 子例全过；banned 未放宽；12 条锁 degraded 与实况一一对应 |
| E. 冻结件 sha | **通过** | thresholds/manifest/LICENSE 与基线逐字节一致 |
| F. 口径局限 | **属实且重大** | degradePyPct 只统计每语言**首个真实 unit（1 个文件）**；全 366 文件同公式 = 26.67% |
| **智能路由** | **NoOne（主体）+ Engineer（P2 残留）** | py 样板绿灯可采信（限口径）；两处残留与自报失实转 Engineer |

---

## A. 产出真实性 — 通过

独立复刻工程师的 21 个最小触发源（另加 `or` 一例，共 22 例），用公开 API `pyfront.NewPyFrontend().Parse` 真实解析后 `air.Collect` 全树扫描（`/tmp/qa_py_probe/probe_out.txt`）：

| 构造 | 声称 Kind | 结果 |
|---|---|---|
| assert x | PyAssert(109) n=1 | **PRESENT** |
| pass | PyPass(125) n=1 | **PRESENT** |
| global g / nonlocal n | PyScopeDecl(124) n=1 | **PRESENT** |
| del xs[0] | PyDelete(121) n=1 | **PRESENT** |
| yield 1 | PyYield(123) n=1 | **PRESENT** |
| not y | PyNotOp(122) n=1 | **PRESENT** |
| x and y / x or y | PyBoolOp(117) n=1 | **PRESENT** |
| x < y | PyCompare(118) n=1 | **PRESENT** |
| x += 1 | PyAugAssign(112) n=1 | **PRESENT** |
| n := f() | PyNamedExpr(113) n=1 | **PRESENT** |
| f-string 内容 s="abc" | PyStringFragment(111) n=2 | **PRESENT** |
| f"{a!r}" | PyTypeConversion(127) n=1 | **PRESENT** |
| with … as fh | PyWithItem(126) n=1 | **PRESENT** |
| 推导式子句 | PyCompClause(116) n=2 | **PRESENT** |
| (x for x in xs) | PyGenerator(119) n=1 | **PRESENT** |
| {"a": 1} | PyDictPair(115) n=1 | **PRESENT** |
| f(x=1) | PyKeywordArg(114) n=1 | **PRESENT** |
| from . import | PyImportPrefix(120) n=1 | **PRESENT** |
| import a.b.c / x: os.PathLike | PyQualifiedName(110) n=1 | **PRESENT** |

工程师点名的 10 条（assert→PyAssert、and/or→PyBoolOp、+=→PyAugAssign、f-string 内容→PyStringFragment、!r→PyTypeConversion、with 项→PyWithItem、:=→PyNamedExpr、推导式→PyCompClause、yield→PyYield、global→PyScopeDecl）**全部覆盖且全部 PRESENT，零 ABSENT**。与上轮"Python 5/5 全 ABSENT"的伪绿形成根本区别：本轮前端行为真实改变。

## B. 语义建模而非贴标签 — 通过（有残留）

全部 `dedicated()` 调用点逐个核对（文件:行号）：

| Kind | 调用现场 | 构造级上下文 |
|---|---|---|
| PyBoolOp / PyCompare | pymap_operator.go:42-58 `binary()` | 由 Pratt 循环真实调用（pyexpr.go:72 `left = e.mapper.binary(left, tok, e.parse(next))`），`pyOperatorKind` 只对**已消费的运算符 token** 分类，锚定运算符真实 span |
| PyStringFragment | pylexspec.go:187 `parseString` reduce | 字符串字面量 reduce 时按真实引号定界切分 content/delimiter（pydedicated.go:85-118），非游离 token |
| PyAssert/PyPass/PyDelete/PyScopeDecl/PyYield | pysimple.go:81-95 `statementMarker` | 关键字简单语句 reduce 现场发射 marker，保留原外层 Kind |
| PyAugAssign / PyNamedExpr | pysimple.go:114-121（赋值 reduce）；pymap_operator.go:65-70 | 解析到 `+=`/`:=` 完成赋值归约时发射 |
| PyWithItem | pywith.go:53 | with 头部逐项 reduce |
| PyCompClause | pycomprehension.go:43 `clauseParts` | for/if 子句 reduce，锚定子句全 span |
| PyNotOp | pyexpr.go parseKeywordPrefix "not" | 前缀表达式 reduce |
| PyTypeConversion | pymap.go:117 `fstringPart` | f-string 插值 reduce，在插值文本内定位 `!r` 真实偏移 |
| PyDictPair | pymap.go:142 `pair` | 键值对 reduce |
| PyKeywordArg | pyexpr_list.go:114 | 实参列表 reduce 中识别 `name=` |
| PyGenerator | pylexspec.go:154 `finishComprehension` | 圆括号推导式 reduce |
| PyImportPrefix | pysimple.go:37 | from 语句 reduce 中数相对点 |
| PyQualifiedName | pydedicated.go:43-66 | import/类型注解 reduce，锚定点号真实 span |

**结论：全部发生在"解析器刚归约完具体构造"处，有构造级上下文与真实源码 span，不是上一轮那种对游离 token 文本的 switch。**（`pyOperatorKind` 虽按文本判别，但输入是 Pratt 循环刚消费的运算符 token，属构造上下文内的合法分类。）

### 残留 1（P2，转 Engineer）：上轮「token 文本贴标签」代码未拆除

- `lang/pyfront/pymap_extra.go`（上轮造假签名文件）仍存在且部分接线：`partialLeaf`→`pyPartialLeafKind`→`pyDetailKind`（pymap.go:69-71，pymap_extra.go:20-37），以及 pyexpr.go:160 在表达式前缀 default 分支调用 `pyKeywordDetailKind`（global/nonlocal/pass→BlockVariant(99)、assert→Invocation(101) 等假 kind）。
- `pymap.go:127` `fstringPart` 直接 `build(air.KindInterpolation=87)`。
- 实测（`fabprobe_out.txt`）：**366 文件全语料产出 fabricated 87-108 节点 1515 个、出现在 158 个文件，且多数在 diags=0 的"干净"解析中**（87 Interpolation×736、93 ParameterList×270、104 OperatorToken×218、103 Delimiter×157、105 SyntaxDetail×108、96 UseDetail×19、100 TypeArg×5、101/102 各 1）。
- 影响评估：**不影响 degradePyPct**（规则表已 ban 87-108，公式只统计 degraded 规则指向的 kind），但 AIR 树仍有伪 kind 污染，"贴标签"叙事未根除。建议：清理 pymap_extra.go 的 87-108 产出路径，f-string 插值改用真实语义或入 degraded 口径。

## C. 口径与规则表一致性 — 基本通过（自报有一处失实）

- 规则层统计（probe C1）：rules=85，exact=71，degraded=12；**exact 且 Kind∈[109,127] 恰好 21 条**；**exact 指向 87-108/Region = 0 条**。mapping 层（e2e 实际口径）exact=87 / degraded=33，与自报"87 exact/33 degraded"一致。
- 与 HEAD 对比：rules_python.go 恰好 21 条 degraded→exact 升级，全部指向新 KindPy*，对应 21 条真实 reduce（见 A/B）；无其他静默升级。
- e2e 实际首个 py unit 经 trace（seq=2202）核实为 `_sources/_validation/_py/_click_8_1_7/examples/aliases/aliases.py`，同公式复现 **degradePct=0.000000**（firstunit_out.txt），与权威报告 e2e_report.md:142 一致。
- **自报失实 1**：自报"全语料 366 文件 probe 显示 19 Kind 全部非零产出"——实测 `KindPyNamedExpr(113) 全语料 observed=0`（366 文件零产出；walrus 在语料中不存在）。另 6 个旧 kind 的 exact 规则也零产出：Await(65)、Case(25)、Arrow(53)、TypeDecl(3)、GenericParam(76)、Union(77)（zerorule_out.txt）。**零产出 exact 规则按公式无条件计入分母（稀释分母、但不进分子）**，这是 freeze 公式的固有行为，非 py 特有，但申报数据必须修正。
- 分子明细（全语料 32 条 degraded 命中，主要 Region(86) 桶的 parameters/patterns/interpolation/paren-expr/argument-list 与 Tuple(79)/Star(59)）——这些 degraded 标签与真实产出一致，属诚实降级。

## D. 防造假测试完整性 — 通过

- `go test -count=1 ./lang/pyfront/... ./internal/e2e/coverage/...` → **全部 ok**；`TestPyDedicatedKindsAreReallyProduced` 21 子例全 PASS，`TestPyQASampledRulesHaveGenuineProduction`、`TestPyExactRulesAreActuallyProduced`、`TestPyStructuralDegradationsStayDegraded` 全 PASS（test_out.txt）。
- banned 列表 = 87-108 全部 22 个 + KindRegion（pyfront_rule_production_test.go:25-37），未放宽；`fabricatedKinds`（degrade_status_test.go:22-31）同口径；另含防重分类测试（TestDegradationNeverReclassifiesDegradedByKind）。
- 12 条锁 degraded 的名单与当前规则表 12 条 degraded 规则一一对应（type-constrained/splat-type/expr-list/list-splat/paren-expr/parameters/patterns/separators/argument-list/interpolation/ellipsis/line-continuation），无放宽。
- 如实记录：①`internal/e2e/coverage/` 整个目录为本轮新增、HEAD 中不存在（`git ls-tree HEAD` 无此目录），"未删改"无 git 基线可比，只能确认当前内容完整；②`TestPyExactRulesAreActuallyProduced` 断言较弱（仅要求 `seenProduced>0`），不能单靠它发现个别 exact kind 零产出——本次由我的全语料 probe 补位。

## E. 冻结件 sha — 通过

| 文件 | sha256（实测） | 与基线 |
|---|---|---|
| ci/e2e_thresholds.json | b967afbba18efdfc62d5d5e4f79e7e802468130f10788603f71d9938a39312d2 | 与 e2e_thresholds_baseline.json **完全一致**，且 = 权威报告 thresholdsSha256 |
| ci/corpus_manifest.sha256 | 记录 53ef136be8f8…；实测 manifest 内容 sha256 = 53ef136be8f8…289c84 | 一致，且 = 报告 manifestSha256 |
| LICENSE | 0d96a4ff68ad6d4b6f1f30f713b18d5184912ba8dd389f86aa7710db079abcb0 | 与要求值一致 |
| git 状态 | `git diff HEAD --stat -- ci/ LICENSE` 为空 | 冻结件零改动 |

## F. 口径局限 — 属实，必须如实记录

- `internal/e2e/coverage` 的 degrade 数值来自 `runner_coverage.go:28-33` + `firstUnitPerLanguage`（runner_coverage.go:52-58）：**每语言只取"首个真实解析 unit"，py 即 aliases.py 一个文件**。degradePyPct=0.000000 仅证明该单文件内 33 条 degraded 规则的 kind 零命中。
- 用同一公式对**全部 366 个 py 语料文件**聚合计算：degraded 命中 32/120 → **26.67%**（probe C2）。即全语料口径下 py 远超 ≤3 阈值；"达标"仅在冻结的单文件口径内成立。
- 对外表述必须限定为："**在每语言首个真实 unit 的冻结口径下** degradePyPct=0（该文件无降级构造产出）"，不得表述为语料级降级率达标。该口径并非本轮新设（firstUnitPerLanguage 为既有机制），本轮 py 改造未篡改口径本身——这一点与上轮改标签不同。

## 附：ineffassign 专项

`golangci-lint run -E ineffassign`（最小配置，仅 ./lang/pyfront/... ./internal/e2e/coverage/...）→ **0 issues**，自报"修 pydedicated.go ineffassign"与实况一致。

---

## 最终结论

1. **py 样板绿灯可采信**：19 个 KindPy* 由 parser 真实 reduce 产出（A/B 双重证实）、规则表升级与实现一一对应（C）、防造假测试完整且全过（D）、冻结件未动（E）。与上轮"78 条标签造假、Python 5/5 ABSENT"的伪绿有本质区别。
2. **两点保留**（不影响绿灯，但须整改/修正申报）：
   - P2 残留：pymap_extra.go 的 token 贴标签路径与 fstringPart 的 KindInterpolation(87) 仍产出伪 kind（1515 节点/158 文件），建议拆除；
   - 自报失实："19 Kind 全语料全非零"不实（PyNamedExpr=0），另 6 条旧 exact 规则全语料零产出，申报数据须修正。
3. **degradePyPct=0 的适用范围**：仅"每语言首个真实 unit"冻结口径（py=aliases.py 单文件）；全 366 文件同公式为 26.67%，**不得对外宣称语料级达标**。
4. **cs/kt/rs 复刻可放行**，附三个前置条件：①逐条复用 pyfront 同款「声称↔产出」真实性测试（含全语料 observed 核查，不能只做最小源码子例）；②拆除各语言等价的 map_extra.go 贴标签路径，不让 87-108 进入解析输出；③申报数据（如"全部非零产出"）先经全语料实测再落纸。
