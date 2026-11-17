# 全项目反造假穷尽审计报告

- 审计人：QA 工程师 严过关
- 授权：用户最高优先级红线任务「彻底闭环根治全项目所有伪造/造假」
- 纪律：只读审计 + `/tmp` 实验（`/tmp/qa_audit/`，含 audit1_out.txt、fab3_current_out.txt、neutrality.json、source_list.txt）+ 单包测试；**未修改任何仓库产品代码/冻结件/阈值/门禁脚本**，本报告为唯一写入。未采信任何自报，全部独立验证。
- 时点说明：审计期间工程师并行拆除 kt/rs 的 `*map_extra.go`；审计结束时 `pymap_extra.go`/`csmap_extra.go` 已删除、`ktmap_extra.go`/`rsmap_extra.go` 仍在（已知在修，不作为造假判定）；`lang/{kt,rs}front/zz_probe_test.go` 为调试遗留（15:13/15:14 新建），记录待清理，不据以判定造假。
- 环境：go1.26.11，GOFLAGS=-mod=vendor；单包测试全部 ok。

## 0. 总判定表

| # | 维度 | 判定 | 一句话结论 |
|---|---|---|---|
| 1 | 标签造假（degrade） | **存疑→整改** | 无新增"未产出却升 exact"证据；但 7 语共 28 条零产出 exact 规则虚增分母，go 无规则表（NOT_MEASURED 如实标注） |
| 2 | 贴标签造假 | **造假（kt/rs 在修）+ 存疑（py/cs 残留 PartialNode）** | go/js/c 零 87-193 产出；py/cs 贴标签文件已删、仅剩恢复用 PartialNode；kt/rs 仍约 80.4 万伪 kind 节点（已知在修） |
| 3 | capture 造假 | **干净** | 反证实验复核：空白 NO、语义变异 YES；Equivalent 第四分栏如实，无静默丢弃 |
| 4 | 变异测试造假 | **干净** | 10003 条 drilldown 零 before==after、零 AIR-ID 未变；算子 before≠after 强制、min 30 校验、94 为实测 |
| 5 | 混沌测试造假 | **干净** | 1400 条 chaos 记录全部真实重解析（1398 captured，2 如实未捕获）；零 recover、零超时 |
| 6 | 语料造假 | **干净（附带说明）** | 20520 文件 sha 全匹配、零缺失、零排除；1142 重复 sha 均来自真实上游 fixtures，未合成 |
| 7 | 报告造假 | **存疑（口径披露不足）** | 数值由 drilldown 真重算、双跑逐字节一致；但"首个真实 unit"单文件口径未在产品文档披露，全语料公式下六语全部超阈 |
| 8 | 门禁判灯造假 | **干净（附带说明）** | S6 真 `jq -e .passed`、S3 真跑零豁免静态检查（96/96、findings=0）；静态扫描覆盖 96 个登记文件，非全仓 |
| 9 | 文档造假 | **基本干净（存疑 1 处）** | 无禁用夸大词，SPDX 附"需权利人确认"；`docs/dev/pending-work.md:23` 的「100% 达成」需补口径限定 |

**综合：项目检测数据尚不能宣告 100% 实打实。** 三处必须根治：①degrade 单文件冻结口径导致的系统性低报（7 项）；②kt/rs 贴标签残留（在修，1 项）；③零产出 exact 虚增分母（1 项）。其余 6 维度无造假证据。

---

## 1. 标签造假（degrade）— 存疑→整改

**证据 A（全语料 exact 产出审计）**：`/tmp/qa_audit/audit1_out.txt` + 前轮 `cs/kt/rs_out.txt`、`py_out.txt`。复现：

```
GOFLAGS=-mod=vendor go run /tmp/qa_audit/audit1.go        # go/js/c
GOFLAGS=-mod=vendor go run /tmp/qa_probe3/main.go cs|kt|rs
GOFLAGS=-mod=vendor go run /tmp/qa_py_probe/main.go       # py
```

| 语言 | exact 零产出条数 | 明细（节选） | 全语料 degrade |
|---|---:|---|---:|
| go | — | 无 rules_go.go；coverageStatus=NOT_MEASURED、`Degradation("go")=0`（index.go:105-107）诚实 | 不参与 |
| js | 2 | labeled_statement(Label36)、update_expression(IncDec58) | 49.12% |
| c | 4 | 数组/指针声明符（Array69/Pointer67） | 28.00% |
| py | 7 | PyNamedExpr113 + Await65/Case25/Arrow53/TypeDecl3/GenericParam76/Union77 | 26.89% |
| cs | 12 | CsExternAlias147/CsRecord136/CsSizeof140 + 9 旧 kind | 6.98% |
| kt | 10 | 全部旧 kind（CompositeLit/Ternary/FuncType/Label/Directive/Spread/Named×2/GenericParam） | 8.21% |
| rs | 10 | RsAsyncBlock164/Const165/Gen166/Try167/EmptyStmt179 + 5 旧 kind | 42.33% |

**判读**：无本轮新增「未产出却升 exact」证据（cs/kt/rs/py 的 21+22+14+30 条升级均有真实 reduce 实证，见 `QA-csktrs-sample-verdict.md`/`QA-py-sample-verdict.md`）；零产出多为语料中罕见构造（record/sizeof/async 块等）及历史遗留。但按 `degrade.go:23-32` 公式它们无条件进分母，属**分母虚增**，须如实登记。banned 87-108：`pyfront_rule_production_test.go:25-37`、`{cs,kt,rs}front_rule_production_test.go:25-37`、`degrade_status_test.go:22-31` 覆盖 22 个+Region；js/c 规则表实测 banned 引用=0；go 无规则表。无绕过证据。

## 2. 贴标签造假 — kt/rs 造假（在修），py/cs 仅恢复节点

**证据 B（全仓扫描 + 实测）**：

```
grep -rn "switch .*\.Text" lang/{go,js,c}front/*.go   # 全部为词法分类/Pratt 运算符，未产出 87-193
GOFLAGS=-mod=vendor go run /tmp/qa_audit/audit1.go     # go fabricated=0, js=0, c=0
GOFLAGS=-mod=vendor go run /tmp/qa_probe3/fab3.go      # 当前 cs/kt/rs 分布
GOFLAGS=-mod=vendor go run /tmp/qa_audit/pyfab.go      # 当前 py: kind=106 PartialNode n=1060
```

- **go/js/c 干净**：全语料解析 3376/4639/3336 文件，kinds 87-193 产出均为 0。jslexspec/jsmap/cexpr 的文本 switch 是词法/Pratt 运算语义（如 jsmap.go:64-80 lexKind、cexpr.go:176-183 unaryPrefix），非专用 Kind 冒充。
- **py/cs 已整改**：`pymap_extra.go`、`csmap_extra.go` 已删除；实测残留仅 `KindPartialNode(106)`（py 1060、cs 3008），来自 `pymap.go:74`/`pyparser.go:174`、`csmap.go:104/271` 的恢复路径——属错误恢复节点而非 token 文本贴标签；仍落在 banned 列表内，建议后续评估是否改 KindRegion 或显式豁免。
- **kt/rs 造假（已知在修）**：`ktmap.go:75/95` 仍调用 `ktKeywordDetailKind/ktDetailKind`，`rslexspec_prefix.go:58-66` 仍调用 `rsKeywordDetailKind`；`ktlexspec.go:282` 在字符串归约处直接 build `KindInterpolation(87)`（非 *map_extra.go 漏网点）。实测 87-108 节点：kt 793,676 个/5902 文件（ParameterList 334,683、OperatorToken 323,890、UseDetail 59,701…）、rs 9,971 个/288 文件（TypeArg 4,416、SpecialIdent 1,925…）。按用户指示记为「已知在修」，不据此判定工程师新造假，但**根治前不能宣告全项目闭环**。

## 3. capture 造假 — 干净

复现：`GOFLAGS=-mod=vendor go run /tmp/qa_abate/probe/main.go`（本轮实跑输出）：四语言 `extra-space captured=NO`（无 token 变化不误报）、`rename/digit/cmp/op captured=YES`；注释变异 captured=YES 但 `ci/equivalent_mutations.json` 明示 `"commentWhitespaceFormatting": "NOT_IN_REGISTRY"`，不进分母。`BindTokenSpellings`（fidelity.go:118-165）、`BindOperator`（opbind.go:12-25）机制未变；上轮 overlay 反证（关闭 spelling 后判定不翻转）仍适用。Equivalent 分栏：gross=10003、equivalent=1271、effective=8732，三栏与报告一致；`SilentDrops=0` 在 aggregate.go:76/142 是不变量声明，已用实测交叉验证（1271 条 equivalent 全数在 gross+captured 中，无静默删除）。判定：真实 token 级捕获，非语义捕获（口径限制延续上轮）。

## 4. 变异测试造假 — 干净

证据：`internal/e2e/mutate/spec.go:33-41`（forbiddenMarkers/Tokens 排除注释空白）、`:66-79`（MinValidPoints=30 注册校验）、`:154-163`（crossPoints 跳过 before==after）；`apply.go:20-42`（edit.Matches 校验+不重叠+内存改写）；`runner_mutate.go:113`（`Captured: st.rootID != afterID`）。drilldown 实测（10003 条）：`before==after` 记录 **0**、AIR-ID 未变记录 **0**、equivalent 记录 1271、各算子 94~1455 条，`validMutationPointsPerOperator=94` 来自 `recompute.go:56 minGrossPerOperator` 实测（call 算子），非写死。等价变异入 gross 并单列，未静默丢弃。单包测试 `go test ./internal/e2e/mutate/...` ok。

## 5. 混沌测试造假 — 干净

证据：`runner_chaos.go:33-107` 对每语言真实构造 damaged 输入（truncate一半 + G1 复现 + SplitMix64 splice，`:110-119`）并重新 `frontend.Parse`；drilldown 实测 chaos 1400 条（每语言 200），captured=1398、未捕获 2 条如实保留（报告 99.86%），outcome 全 OK、零超时。`recoverUsageCount=0` 由 neutrality 分析器 `analyzer.go:197-208` 对 `rec`+"over" 的 go/ast 检测得出（`runner_coverage.go:115`），非硬编码。单包测试 `go test ./internal/e2e/chaos/...` ok。

## 6. 语料造假 — 干净（附说明）

独立重算（Python sha256 遍历 `ci/corpus_manifest.json` 20520 条）：缺失 0、sha 不匹配 0；每语言 299~5920；`ci/corpus_manifest.sha256` 记录 53ef136b… 与实算一致；报告 registeredFiles=processedFiles=1540、`corpusExclusionCount=0`，抽样为 seed 分层（`internal/corpus/sample.go:83-143`），无按来源排除。重复 sha 178 组/1142 文件，抽样核查均为同一上游项目内 fixtures（prettier 138 组、kotlin 19 组、cJSON 8 组等），非合成填充；但计数未去重，`parseableFilesPerLanguage=215` 是 sampled 220 中的实测去重路径数（`recompute.go:160-169`），不是全语料可解析数——**口径说明项**。单包测试 `go test ./internal/corpus/...` ok。

## 7. 报告造假 — 存疑（口径披露不足）

链路真实：`cmd/e2e/main.go:85-101` 先 `report.Write`、再 `report.Recompute` 零差断言、再 `Evaluate`+`WriteVerdict`；`recompute.go:26-73` 从 drilldown/precision/facts 派生全部 baseline，`evaluate.go:98+` 执行 tighten-only 零容忍；S6 由 gate `jq -e .passed` 判灯。run1/run2 `e2e_report.json` 与 `e2e_drilldown.jsonl` 实测 `cmp -s` 逐字节一致。

**存疑点（重大）**：`runner_coverage.go:28-33,52-58 firstUnitPerLanguage` 只对每语言**第一个真实 unit（单文件）**算 degrade，该口径仅出现在 QA 报告（docs/i18n/QA-*.md），产品文档 `docs/system_design.md`、`docs/dev/frontend-coverage/*.md` 未披露；用同一公式跑全语料（本审计实测）js 49.12%、c 28.00%、py 26.89%、rs 42.33%、kt 8.21%、cs 6.98% —— 六语全部超阈，go 为 NOT_MEASURED。数值本身真实，但"达标"结论若脱离单文件口径即为**对外夸大外推**。另 `validMutationPointsPerOperator` 名称与实算（每算子最小 gross 变异记录数）略有语义偏差，非造假。

## 8. 门禁判灯造假 — 干净（附说明）

- `ci/gate_e2e.sh:173-181` S6：`jq -e '.passed' e2e_verdict.json` 真判灯，失败打印 failures；脚本全文无 skip/白名单/env 放宽开关。
- S2 `:55-106` 阈值 tighten-only jq 单调校验；S1 `:29-52` 全量 sha/SPDX 重算；S4 `:138-162` 双跑 `cmp`；S7 归档。
- S3 `:113-134` 真跑 `go run ./cmd/neutralitycheck`。独立复现：`GOFLAGS=-mod=vendor go run ./cmd/neutralitycheck -sources ci/e2e_sources.txt -out /tmp/qa_audit/neutrality.json` → `NEUTRALITY_OK registered=96 processed=96 findings=0`；`find` 反查 96=96 与 `ci/e2e_sources.txt` 逐字节一致。说明：96 是 `ci/e2e_new_roots.txt` 登记根下的 Go 文件，非全仓文件，文档宣称"零豁免 go/ast 静态检查"须限定此范围。单包测试 `go test ./internal/e2e/neutrality/...` ok。

## 9. 文档造假 — 基本干净（存疑 1 处）

- 全仓 `grep -rniE "国际一流|世界一流|行业领先|最先进|首创|100%达成"`（排除 third_party/.artifacts）未命中禁用营销词；禁用词表与例外规则见 `docs/i18n/PRD-dual-license-bilingual.md:425`。
- SPDX/许可：`LICENSE-COMPATIBILITY.md:4,31,34` 显式标注 "AGPL-3.0-or-later **需权利人确认**"、DCO 不授予再许可权、双许可缺口待确认；LICENSE sha256=0d96a4ff… 与冻结基线一致。无夸大。
- 存疑：`docs/dev/pending-work.md:23`「工程实现与全部门禁口径已 100% 达成」与同文件 :48/:151 的「功能级验证为已知边界、铁律口径未宣称 100%」张力，建议统一为「工程/门禁口径 100%，铁律口径另行说明」；:130 的 PIT 100% 有 `.artifacts/pit-parity.json` 实据。详见第 7 维度的口径披露缺口。

---

## 智能路由

| 路由 | 项 | 处理 |
|---|---|---|
| **Engineer** | ①kt/rs 贴标签残留（ktmap_extra.go、rslexspec_prefix.go、ktlexspec.go:282） | 沿 py/cs 已删路径根治；删后重跑全语料 87-108 产出须归零（恢复用 PartialNode 另行评估） |
| **Engineer** | ②`lang/{kt,rs}front/zz_probe_test.go` 调试遗留（绝对路径） | 清理出仓 |
| **Engineer** | ③零产出 exact 规则 28 条 | 逐条要么补真实 reduce（并纳入产出测试），要么回退 degraded 并从分母剔除 |
| **文档修订（主理人/Engineer）** | ④单文件口径披露 | 在 system_design.md / frontend-coverage 文档写明 degrade 按首 unit 计算及全语料差异；对外结论禁止外推 |
| **文档修订** | ⑤pending-work.md:23 「100% 达成」限定语 | 与 :48/:151 口径统一 |
| **NoOne** | 维度 3/4/5/6/8 及维度 1/2 的 go/js/c、py/cs 主体 | 无造假证据，保持现状；不因已知口径局限追责既有实现 |

## 最终结论

**项目检测数据目前不能宣告 100% 实打实：六语 degrade 数值在冻结单文件口径下真实、双跑确定、报告由 drilldown 真重算、门禁真判灯、capture/变异/混沌/语料四维度无造假；但"达标"结论依赖未披露的首 unit 单文件口径（全语料六语全部超阈），kt/rs 贴标签产出仍在进行中的根治，且有 28 条零产出 exact 规则虚增分母。**

**彻底根治完备清单（按优先级）**：

1. P0 — 将 degrade 统计改为全语料聚合（或明令阈值文档的口径键从 `degrade*Pct` 改为 `degradeFirstUnit*Pct` 并在所有对外文档披露），按真实口径重判六语是否达标；
2. P0 — kt/rs 拆除 `*map_extra.go` + `rslexspec_prefix.go` + `ktlexspec.go:282` 的 87-108 产出，重测全语料伪 kind=0（PartialNode 例外单独裁定）；
3. P1 — 28 条零产出 exact 规则逐条补产出或回退 degraded；
4. P1 — 清理 `zz_probe_test.go`；统一 docs 中 100% 表述与口径披露；
5. P2 — 评估语料 1142 重复 sha 是否影响"每语言 ≥200 parseable"的统计代表性，必要时对抽样去重。

完成 1–5 后方可对外宣告「检测数据 100% 实打实」。
