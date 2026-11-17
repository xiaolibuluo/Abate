# cs/kt/rs 前端 degrade 复刻「真实性」独立验收报告

- 验收人：QA 工程师 严过关
- 前案：`docs/i18n/QA-py-sample-verdict.md`（py 样板验收，抓出自报失实与 87-108 残留）
- 对象：工程师寇豆码将 py 路径复刻到 cs/kt/rs，自报 degradeCs=3.255814（≤5）、Kt=4.477612（≤5）、Rs=0.613497（≤3）、Py=0.000000；run1/run2 verdict 均 passed=true（14:00 生成，实测复核属实）
- 验收纪律：只读 + `/tmp` 实验（`/tmp/qa_probe3/`，含 *_out.txt）+ 单包测试；未修改仓库文件，本报告为唯一写入
- 环境：go1.26.11，GOFLAGS=-mod=vendor；验收前 `git status` 确认工作区只有工程师本轮改动，无并行半成品混入

## 总览判定

| 项 | 判定 | 一句话结论 |
|---|---|---|
| A. 真实产出 probe | **通过** | cs 22/22、kt 14/14、rs 28/28 PRESENT，零 ABSENT；上轮 ABSENT 名单全部真实产出 |
| B. 语义建模核查 | **通过（有残留）** | 新 Kind 全部在归约现场发射（Pratt/字符串/unit 归约完成），非游离 token switch；但 DetailKind 裸 token switch 仍接线并产出海量 87-108 伪 kind |
| C. 规则表一致性 | **通过（如实列零产出）** | 三语 exact 指向 87-108/Region = 0；新 kind exact 数 22/14/30 与定义一一对应；零产出 exact 规则 cs12/kt10/rs10 已如实列出 |
| D. 防造假测试完整性 | **通过** | 三语 4 个测试全过；banned 未放宽；degrade_status_test.go 未删改 |
| E. 冻结件 sha | **通过** | thresholds/manifest/LICENSE 逐字节一致 |
| F. 口径局限 | **属实且放大** | 首个 unit 单文件口径全达标；全语料同公式 cs=6.98%、kt=8.21%、rs=42.33%，三语全部超标 |
| **智能路由** | **NoOne（达标判定）+ Engineer（残留整改）** | 数值无造假、产出真实；口径与残留须整改后方可对外宣告语料级达标 |

---

## A. 真实产出 probe — 通过

用公开 API（NewCsFrontend/NewKtFrontend/NewRsFrontend）真实解析最小源码后 `air.Collect` 全树扫描（`/tmp/qa_probe3/{cs,kt,rs}_out.txt`）：

**cs 22/22 PRESENT（0 ABSENT）**：constructor→CsConstructor(128)【上轮 ABSENT，现已真实产出】、query→CsQuery(142)【上轮 ABSENT】、interpolated-string→CsInterpString(143)【上轮 ABSENT】、string-part→CsStringFrag(144)、typeof(141)、sizeof(140)、object-creation(139)、conditional-access(138)、destructor(129)、operator-decl(130)、property(133)、indexer(134)、accessor(132)、event(135)、record(136)、yield(137)、char-body(145)、modifier(146)、extern-alias(147)、is-pattern(148)、alias-qualified(149)、local-func(131)。

**kt 14/14 PRESENT（0 ABSENT）**：companion→KtCompanion(150)【上轮 ABSENT】、interpolated→KtInterp(156)【上轮 ABSENT】、when-entry(161)、secondary-ctor(151)、ctor-delegation(152)、accessor(153)、property(158)、import-alias(154)、import-list(155)、string-part(157)、this(159)、super(160)、callable-ref(162)、control-body(163)。

**rs 28/28 PRESENT（0 ABSENT）**：negative-literal→RsNegativeLit(186)【上轮 ABSENT】、async/const/unsafe/try-block(164/165/168/167)、compound-assign(169)、field-init(170)、base-field(171)、shorthand-field(172)、shorthand-ident(173)、metavar(175)、self(176)、crate(177)、super(178)、empty-stmt(179)、string-part(180)、extern-crate(181)、scoped-use(182)、use-list(183)、use-wildcard(184)、use-as(185)、yield(187)、macro-call(188)、macro-def(189)、func-sig(190)、generic-func(191)、turbofish(192)、type-args(193)。

**诚实降级核实**：上轮 ABSENT 且本轮未升级的规则保持 degraded——cs.named-type、kt.simple-identifier、kt.postfix 由 `Test{Cs,Kt}QAAbsentRulesStayDegraded` 锁死；未出现"无产出却升 exact"。

## B. 语义建模核查 — 通过（有残留）

新专用 Kind 的发射点逐一核对：

**cs**：`csexpr_dedicated.go:22-65 keywordPrefix`（Pratt 刚消费前缀关键字时发射 typeof/sizeof/object-creation）、`:69-127 infixRule`（Pratt 刚消费 `?.`/`is`/`::` 时发射 conditional-access/pattern/alias-qualified）；`csmap.go:102/115-131`（字符/字符串字面量归约现场切分 charBody/stringFragments/delimiterFragment）；`csparser.go:227/242`（unit 归约完成时 `dedicatedMarkers`——基于归约结果 mapping/scope/typeName 做构造级判定，如 constructor=类型作用域成员+lead==类型名，非裸 token 判定）。

**kt**：`ktexpr.go:251-257`（Pratt 刚消费 `::` 发射 CallableRef）；`ktparser.go:244-250 caseNode`（when 分支归约现场在 `->` 记号发射 WhenEntry）；`ktdedicated.go:29-83`（字符串字面量归约现场切分 content/delimiter）；`ktdedicated_markers.go`（unit 归约完成时按 mapping.note/kind 判定发射 companion/ctor/accessor/import 等）。

**rs**：`rsdedicated.go:104-114 assign`（赋值归约现场、运算符刚消费时发射 CompoundAssign）；`rsexpr.go:119/232-234`（Pratt turbofish/type-args）；`rsparser.go:196`（unit 归约完成时 `dedicatedMarkers`：macro_rules/use/fn/trait/块前导关键字等单元级构造判定 + safeAnchor 防重叠）。

**结论：与上轮「map_extra 裸文本 switch」不同，本轮新 kind 全部有归约上下文（Pratt 消费点、字面量归约、unit 归约完成时的 mapping 判定），锚定真实 token span。** 判定为语义建模，非贴标签。

### 残留（P1，转 Engineer）：DetailKind 裸 token switch 仍接线，海量产出 87-108 伪 kind

- `csmap_extra.go`（csDetailKind，csmap.go:78 接线）、`ktmap_extra.go`（ktmap.go:75/95 接线）、`rslexspec_prefix.go:58-66`（rsKeywordDetailKind 接线）仍按 token 文本产出 87-108 共享桶。
- 全语料实测（`fab3_out.txt`）：**cs 613,093 个伪节点 / 2564 文件**（ParameterList 340,697、Delimiter 216,546、Invocation 24,107、OperatorToken 21,299…）、**kt 793,676 个 / 5902 文件**（OperatorToken 323,890、ParameterList 334,683、UseDetail 59,701…）、**rs 9,971 个 / 288 文件**（TypeArg 4,416、SpecialIdent 1,925…）。py 前文已测 1,515/158。
- 不影响 degrade 数值（规则表已 ban、公式只数 degraded 规则 kind），但与 kicds.go 中"87-108 是已撤回造改产物"的注释直接矛盾，AIR 树持续被伪 kind 污染。**建议 P1：三语与 py 一并拆除 DetailKind→87-108 产出路径**（标点/运算符细节可回退 KindRegion 或保留 frozen 投影）。

## C. 规则表一致性 — 通过（自报计数有小出入，零产出如实列出）

| 语言 | 自报 exact/degraded | 实测规则层 | 实测 mapping 层 | exact 指向 87-108/Region | 新 kind exact 数 |
|---|---|---|---|---|---|
| cs | 75/60 | **79**/60 | 104/111 | **0** | 22（=KindCs* 定义数） |
| kt | 61/38 | **65**/38 | 72/62 | **0** | 14（=KindKt* 定义数） |
| rs | 82/34 | **86**/34 | 94/69 | **0** | 30（=KindRs* 定义数） |

- degraded 计数三语完全一致；exact 恒差 +4（口径/时点出入，不影响 degraded 分母分子，如实记录）。
- 与 HEAD diff 核对：cs 升级 22 条、kt 14 条、rs 30 条 degraded→exact，全部指向 128-193，无其他静默升级。
- **全语料零产出 exact（吸取 PyNamedExpr 教训，非只单文件）**：
  - cs（12 条）：CsExternAlias(147)、CsRecord(136)、CsSizeof(140)（新 kind）+ CompositeLit×2、Cast、ExprStmt、Label、Ref、Spread、Tuple、GenericParam（旧 kind，语料未含该构造）；
  - kt（10 条）：全部旧 kind（CompositeLit×2、Ternary、FuncType、Label、Directive、Spread、Named×2、GenericParam）；
  - rs（10 条）：RsAsyncBlock(164)、RsConstBlock(165)、RsGenBlock(166)、RsTryBlock(167)、RsEmptyStmt(179)（新 kind）+ Array、FuncType、Label、Pointer、Directive。
  - 判读：零产出构造多为语料中罕见/不存在的形态（record、sizeof、extern alias、async/const/gen/try 块），非伪造；但按公式它们无条件计入分母。**工程师未再声称"全部非零产出"，本次申报无 py 轮那类失实。**

## D. 防造假测试完整性 — 通过

- `go test -count=1 ./lang/{cs,kt,rs}front/... ./internal/e2e/coverage/...` 全部 ok（含 -race 前置的单包子集）。
- 三语 `Test*DedicatedKindsAreReallyProduced`（23/14/30 子例）+ `Test*QASampledRulesHaveGenuineProduction`（要求 exact + kind≥128 + 最小源码真实产出）+ `Test*ExactRulesAreActuallyProduced`（真实语料）+ `Test*QAAbsentRulesStayDegraded`（锁 cs.named-type、kt.postfix/kt.simple-identifier、rs.use-bounds/field-list/field-init-list/trait-bounds/where-clause）全部 PASS。
- banned 列表三语均为 87-108 全部 22 个 + KindRegion，未放宽；`internal/e2e/coverage/degrade_status_test.go` 三个防造假测试原样保留（fabricatedKinds 22 项完整）。

## E. 冻结件 sha — 通过

- ci/e2e_thresholds.json = ci/e2e_thresholds_baseline.json = b967afbb…d2（=报告 thresholdsSha256）；
- ci/corpus_manifest.json sha256 = 53ef136b…c84，与 ci/corpus_manifest.sha256 记录及报告 manifestSha256 一致；
- LICENSE = 0d96a4ff68ad6d4b…cb0；
- `git diff HEAD -- ci/ LICENSE` 为空，冻结件零改动。

## F. 口径局限 — 属实且放大（决定结论措辞）

- 同 py：degrade 数值只统计**每语言首个真实 unit**（runner_coverage.go `firstUnitPerLanguage`）。trace 核实四语首个 unit：cs=BsonReaderAsyncTests.cs、kt=OkHttpTest.kt、rs=serde/build.rs、py=aliases.py。四语首个 unit 同公式复现自报值 3.255814 / 4.477612 / 0.613497 / 0.000000，**数值与口径一致、无造假**。
- 但同一公式跑**全语料**（cs 2564、kt 5920、rs 299 文件）：

| 语言 | 首个 unit（冻结口径，阈值） | 全语料 | 是否超阈 |
|---|---:|---:|---|
| cs | 3.255814（≤5） | **6.976744** | **超** |
| kt | 4.477612（≤5） | **8.208955** | **超** |
| rs | 0.613497（≤3） | **42.331288** | **远超**（Region 桶 48 条规则命中为主） |

- 结论措辞必须限定："**在每语言首个真实 unit 的冻结口径下**，cs/kt/rs/py 四语 degrade 全部达标，且数值经独立复现无造假"；**不得外推为语料级降级率达标**（全语料口径下四语均超阈值）。该口径为既有机制、本轮未篡改。

## 附：run1/run2 一致性

`.artifacts/e2e/53ef136be8f8/{run1,run2}/e2e_verdict.json` 均 passed=true、failures=[]；两轮 e2e_report.md 的 degradeCs/Kt/Rs/Py 四值逐字节一致（138/141/142/143 行），确定性口径稳定。

---

## 最终结论

1. **degrade 维度可宣告「真实达标」，但必须带口径限定语**：本轮 cs/kt/rs 复刻不是上轮式标签造假——新专用 Kind（128-193）全部由 parser 真实归约产出（A/B 双重证实）、规则表升级与实现一一对应且零指向共享桶（C）、防造假测试完整（D）、冻结件未动（E）、自报数值独立复现（F）。在冻结的"首个 unit"口径内，达标真实成立。
2. **两项保留（转 Engineer，不阻塞达标宣告本身）**：
   - P1：cs/kt/rs/py 四语的 DetailKind 裸 token switch 仍产出 87-108 伪 kind（合计约 141 万节点），与 kinds.go 注释矛盾，应拆除；
   - 申报口径出入：exact 规则数自报比实测少 4（cs/kt/rs 一致），建议统一计数口径。
3. **对外表述红线**：达标限定于"每语言首个真实 unit"冻结口径；全语料公式下 cs 6.98%、kt 8.21%、rs 42.33% 均超阈，需在后续迭代（尤其拆除 Region 桶依赖的 48 条 rs 规则）后另行申报。
