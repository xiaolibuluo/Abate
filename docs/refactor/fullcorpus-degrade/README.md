# 自研语料全语料降级率测量（诚实口径）

> 文档编号：ABATE-FULLCORPUS-DEGRADE-001
> 测量日期：2026-11-18
> 可复现命令：`bash scripts/fullcorpus-degrade.sh`（或直接 `go run ./cmd/degradecheck`）
> 原始数据：`fullcorpus-degrade.json`，SHA-256 见 `fullcorpus-degrade.json.sha256`。

## 1. 为什么做这次测量

`README.md` / `README.en.md` 的诚实口径此前写明：E2E degrade「达标」**仅限冻结的单文件（首 unit）口径**，并注明"自研语料的全语料降级率尚未重测，为 `NOT_MEASURED`"。

E2E 门禁（`internal/e2e/runner_coverage.go:firstUnitPerLanguage`）确实用每语言**第一个真实 unit** 计算降级率。`ci/corpus_manifest.json` 登记的 20520 条自研语料虽然全部被 E2E 采样链路真实 parse，但降级率聚合只取了每语言首 unit。因此"全语料降级率"是一个真实存在、此前未测的口径缺口。本次测量把它闭环为可复现的真实数据。

## 2. 测量方法（与 E2E 同公式，全量聚合）

1. 读取 `ci/corpus_manifest.json` 全部 20520 条自研语料条目。
2. 对**每条**条目用其真实前端（`internal/e2e.FrontendTable`）parse，得到 AIR unit。
3. 用 `lang/frontkit/coverage.Observe` 统计真实 parse 产出的 AIR 节点 kind；`Merge` 按语言合并全语料观测。
4. 对每种有嵌入语法映射的语言，用 `internal/e2e/coverage.Degradation` 计算降级率：
   `degraded 且被观测到的映射数 / (exact + degraded 映射数) × 100`。
5. Go 未嵌入语法映射（`mappings=0`），诚实保留 `NOT_MEASURED`。

**反造假约束**：degraded 映射绝不重分类为 exact（`degrade_status_test.go` 有专门回归测试守护）；观测全部来自真实 parse；`parseFail` 与 `emptyUnit` 如实记录、不跳过；双跑逐字节 `cmp` 通过才归档；报告附 SHA-256。

## 3. 实测结果（原始数据见 JSON）

| 语言 | 状态 | 降级率 | 文件数 | parse 失败 | 空 unit | 观测 kind 数 | 映射数 | 被观测 degraded |
|---|---|---:|---:|---:|---:|---:|---:|---:|
| go | NOT_MEASURED | — | 2932 | 0 | 0 | 38 | 0 | 0 |
| py | MEASURED | 3.36% | 2932 | 0 | 0 | 43 | 122 | 4 |
| js | MEASURED | 3.51% | 2932 | 0 | 0 | 58 | 114 | 4 |
| c | MEASURED | 4.80% | 2932 | 0 | 0 | 34 | 125 | 6 |
| cs | MEASURED | 4.65% | 2932 | 0 | 0 | 43 | 215 | 10 |
| rs | MEASURED | 4.91% | 2932 | 0 | 0 | 47 | 163 | 8 |
| kt | MEASURED | 4.48% | 2928 | 0 | 0 | 34 | 134 | 6 |

合计 20520 条，全部真实 parse 成功（parse 失败 0、空 unit 0）。六种有嵌入语法映射的语言全部降到 5% 以下。

## 4. js 高保真提升（真实产出，非改标签）

js 全语料降级率由 **46.49% 降至 21.93%**（被观测 degraded 映射 53→25，观测 kind 32→52）。做法是复制 py/cs/rs/kt 已验证的范式，且严格守住禁造假底线：

1. 在 `lang/air` 追加 `KindJs*`（194–222）专用节点类，只追加、不重编号，wire `SchemaVersion` 保持 2。
2. 在 `lang/jsfront/jsdedicated.go` 对语料中**真实存在**的 20 类构造（export、generator、let/const、this/super、class-body、class-heritage、switch-body/default、new、paren、sequence、incdec、yield、template-sub、string-part、property-ident、arguments、formal-params、declarator）于真实语法归约点产出**零宽专用节点**，锚定真实 token，不改任何既有投影，字节保真与 AIR-ID 不变。
3. 在 `lang/frontkit/grammar/rules_javascript.go` 把上述构造的映射从 degraded 改为 exact，**每处 exact 都有前端真实节点产出支撑**。`lang/jsfront/jsdedicated_test.go` 逐构造证明专用 kind 被真实 emit 且为零宽，是可执行的反造假守卫。

关键点：degrade 公式对 exact 映射不检查 `observed`——"只改标签"即可让数字下降，这正是历史严打的假 exact。本提升未采用该手法：每个 exact 映射均由前端真实产出对应节点，观测 kind 由 32 增至 52 即为证据。

**第二轮推进（js 21.93%→3.51%，c 10.40%→4.80%）**：在第一轮范式基础上继续以真实产出降级。

- js：新增 `KindJsDebugger(226)`、`KindJsMetaProperty(227)`、`KindJsComputedName(228)`、`KindJsPairPattern(229)`、`KindJsObjectAssignPat(230)`，并在 `jsdedicated.go` 补齐 debugger/class-expr/meta-property/computed-name/pair/pair-pattern/object-assign-default/rest-pattern 六类真实归约点 marker（含括号配对与模式上下文判定）；`jsx/regex/aug-assign/using/static-block` 直接升级到第一轮已产出的专用 kind。每处 exact 均由 `TestDedicatedKindsReallyEmitted` 最小样例 + `TestJsExactRulesAreActuallyProduced` 真实语料双重证明。
- c：新建 `lang/cfront/cdedicated.go` 专用 marker 基础设施（沿用 rsfront 的 safeAnchor 零宽范式），新增 `KindCUnion(231)`…`KindCFieldDecl(237)` 七个专用 kind，对 union/枚举成员/字段声明/解引用/初始化设计符/`defined`/语句标签七类构造在真实归约点产出零宽节点；`cfront_rule_production_test.go` 新建整套守护（专用 kind 禁区 + degraded pin + 真实产出验证）。
- 剩余 degraded（js 的 object/array/assign 三类解构模式与 empty-stmt，c 的 field/type/sized-type 标识符与 func/init-declarator、macro-type）均映射到 AIR 尚无独立节点的共享 kind，如实保留 degraded 并由 pin 测试锁定，不以凑数牺牲真实性。

## 5. 诚实结论与边界

- 本表是**全语料口径**的真实测量，与 E2E 门禁的**单文件（首 unit）口径**不同，两者不得混用或互相外推。
- js 全语料降级率 21.93% 是真实产出专用节点后的客观现状，**不掩盖、不美化、不通过改标签降低**；剩余部分为语料零产出构造，如实保留。
- 全语料口径数值**未达**单文件冻结阈值；本测量只作事实披露，**不构成**"高保真""达标""国际一线"等结论。
- Go 无嵌入语法映射，保持 `NOT_MEASURED`，不以 0 冒充达标。
- 历史第三方语料的六语数值（js 49.12% 等）见 `docs/i18n/QA-ANTI-FRAUD-AUDIT.md`，与本表自研语料口径不同，禁止混用。

## 6. 复现

```bash
cd <repo>
bash scripts/fullcorpus-degrade.sh
# 产物：docs/refactor/fullcorpus-degrade/fullcorpus-degrade.json
# 校验：sha256sum -c <(sed 's|$|  docs/refactor/fullcorpus-degrade/fullcorpus-degrade.json|' \
#         docs/refactor/fullcorpus-degrade/fullcorpus-degrade.json.sha256)
```

两份连续运行的报告逐字节一致，确定性成立。
