# Abate 彻底自研 + 清除全部开源接入 + 迭代整改 TODO/PRD

> 文档编号：ABATE-REFACTOR-TODO-001　版本：v1.0　状态：待执行（合规驱动·紧急）
> 负责人：许清楚（Xu，产品经理）　执行对象：架构师 / 工程师
> 语言：中文　适用仓库：`xiaolibuluo/Abate`　本地路径：`/workspace/abate`

---

## 0. 前置说明（本文件怎么用）

1. 本文件是**唯一执行清单**，架构师按编号推进，每完成一项在文末 DoD 勾选，并回写状态。
2. 本文件**只做规划与追踪**：此刻不改产品代码、不改门禁、不 commit、不 push。
3. 带 ⚖️ 标记的任务涉及法律判断，**AI 不出法律结论**，只列事实与待确认项，须权利人/法务书面确认后方可改动对外声明。
4. 每个阶段结束须通过 P5 阶段门（复绿）后再执行推送（见「推送策略」）。

---

## 1. 背景

用户收到合规警告信：**所有已接入的开源项目均"犯规"**。要求：

- **彻底自研**：产品代码不得依赖任何第三方源码。
- **清除所有接入开源组件源码**：含编译依赖、参考副本、benchmark 语料副本、"融入产品代码的二次并入"。
- **迭代整改**一系列"仔细检查才看得出来"的问题（保真度、derive、竞品、成熟度、许可、命名，见第 6 节专项）。
- 执行顺序：**先出本清单，再执行推进**。

这是一次合规驱动的紧急重构，优先级高于功能迭代；但**不得以合规为名伪造证据**（见诚实性总纲）。

---

## 2. 事实基线（已实地核对，可直接采信）

### 2.1 仓库与推送

| 项 | 事实 |
|---|---|
| 仓库可见性 | `xiaolibuluo/Abate` 现为**私有仓库**（`private=true`，API 已确认），用户紧急要求 |
| 推送通道 | SSH over 443 的部署密钥（deploy key） |
| 分支保护 | Owner-only Ruleset，**禁止强推**；强推需临时 disable → force push → enable |
| 本地状态 | `/workspace/abate`，`main` HEAD = `119cf2f5`；工作树有未跟踪目录 `docs/benchmark/`（上轮中断的 benchmark 残留） |

### 2.2 真正参与编译的第三方依赖（go.mod 权威，仅 3 个）

```
module abate
go 1.26.11
require (
    github.com/tetratelabs/wazero v1.12.0     // WASM 运行时
    gopkg.in/yaml.v3          v3.0.1          // YAML 解析
)
require golang.org/x/sys v0.44.0 // indirect  // wazero 间接引入
```

- 产品代码（`kernel/`、`lang/`、`internal/`、`plugin/`、`ctl/`、`cmd/`）自身 `import` 仅 **wazero** 与 **yaml.v3**。
- `vendor/` 目录存在，含 `github.com/tetratelabs/wazero`、`golang.org/x/sys`、`gopkg.in/yaml.v3` 三份源码副本。

### 2.3 不参与编译的纯参考源码副本（`_` 前缀，Go 工具链自动忽略）

- `third_party/_sources/`（约 **527MB**）：`gvisor`、`llvm-project`、`roslyn`、`kotlin-compiler`、`pit`、`swc_ecma_parser`、`syn`、`tree-sitter-*`（c / c-sharp / javascript / kotlin / python / rust / runtime）、`wazero`、`_validation` 等。
- `third_party/_references/`：`google__go-cmp`、`Shopify__toxiproxy`、`spf13__cobra`、`go-gremlins__gremlins`、`_m5_algorithms.go.txt`、`sarif-2.1.0.schema.json`、`MANIFEST.md`。
- 附：`third_party/LICENSE-MANIFEST.md`、`ci/corpus_manifest.json`（登记了 gin/requests/click/lodash/prettier/express/axios/cJSON/Newtonsoft.Json/serde/okhttp 等 11 个上游"只读源码落地"）。

### 2.4 benchmark 残留（属于"接入开源源码"，必须清除）

- `docs/benchmark/cases/js_lodash_4_17_21/src/`（lodash 4.17.21 完整源码 + `.github`/`dist`/`doc`/`fp`/`lib`/`perf`/`test`/`vendor`），`config/node_modules/`。
- `docs/benchmark/cases/c_cjson_v1_7_16/src/`（cJSON 1.7.16 源码 + `fuzzing`/`tests` 等）。
- `docs/benchmark/raw/`、`audit/`、`scripts/`、`scripts/corpus_fetch.sh`（下载逻辑）。

### 2.5 融入产品代码的"二次并入"（需逐一甄别自研 vs 开源改写）

| 位置 | 怀疑来源 | 需甄别 |
|---|---|---|
| `kernel/mutation/`（`derived.go`/`engine.go`/`minimize.go`/`discover.go` 等） | PIT 变异核二次开发 | 算法移植 vs 自研 |
| `kernel/injection/`（`ops_*.go`、`injector.go`、`world.go`）与 `plugin/` | chaosblade 注入核 + wazero 宿主 | 算子移植 vs 自研 |
| `kernel/abi/`（`amd64Names_*.go`、`arch_*.go`、`syscall_table.go`、`addr*.go`） | gVisor ABI 移植 | 移植 vs 独立实现 |
| `sim/rand`（SplitMix64）、`lang/*front`（疑似 tree-sitter 语法衍生）、`scripts/build-libclang-min.sh`（libclang） | 公有领域算法 / 其他 | 全仓扫描确认 |

### 2.6 门禁与治理体系（重构后必须全绿）

- `scripts/lint-gate.sh` 十段：`[0] 版权头`→`[1] 风格`→`[2] 许可证`→`[3] gofmt`→`[4] go vet`→`[5] golangci-lint`→`[5b] 魔法常量`→`[6] go test -race`→`[7] 覆盖率`，另 `coverage-gate.sh`。
- 配套：`tla/`（TLA+）、`watermark-verify.sh`、`hash-check.sh`、`SHA256SUMS`、`WATERMARK.sha256`、`m5-reference-scan.sh`、`pit-parity.sh`、`drift-check.sh`、`determinism-check.sh`。
- 许可/治理文档：双语 README / LICENSE（AGPLv3 全文）/ LICENSING / COMMERCIAL-LICENSE / DCO / CLA / TRADEMARK / NOTICE / THIRD-PARTY-NOTICES / FAQ-licensing / LICENSE-COMPATIBILITY 等，体量极大。

### 2.7 诚实性与成熟度现状

- README 第 131–137 行已诚实声明：`competitorAbsolute = NOT_MEASURED`；E2E degrade 达标**仅限冻结单文件（首 unit）口径**，全语料六语超阈（js 49.12% / c 28.00% / py 26.89% / rs 42.33% / kt 8.21% / cs 6.98%，go `NOT_MEASURED`）；derive killing 为静态 AIR 级近似，非真实执行 go test。
- 社区信号：0 star / 0 fork / 0 watcher / 3 commit；文档声称 GA `v1.0.0`。

---

## 3. 清除范围（本次 = 全部，含编译依赖）

> 判定口径：**"清除"= 产品仓库内不再存在任何第三方源码副本，且编译不依赖任何第三方模块。**

必须清除/替换的四类：

1. **编译依赖（全部）**：wazero、yaml.v3、x/sys → 全部自研替换，`go.mod` 第三方 require 清零。
2. **vendored 副本**：`vendor/` 下三份第三方源码 → 删除目录。
3. **参考副本**：`third_party/_sources/`、`_references/` 及全部 manifest/README 引用 → 删除并改写台账。
4. **融入代码的开源改写**：PIT / chaosblade / gVisor / tree-sitter 衍生等 → 甄别后重写为自研实现（或经法务确认属事实数据后保留并标注）。
5. **benchmark 语料副本**：lodash / cJSON 等 `src`、`node_modules` → 删除；测量结果数据若无源码片段可保留为"历史、未对拍"标注。

**不在清除范围**：项目自身按 AGPLv3 + 商业双许可对外授权（那是本项目对外授权，非导入依赖）；纯数学/事实数据（如 syscall 编号、SplitMix64 算法常量）——但须经第 8 节法务红线确认。

---

## 4. 推送策略（每阶段强推）

每个阶段完成 P5 阶段门复绿后，作为一个独立提交批次推送：

1. `scripts/ruleset-toggle.sh disable`（临时关闭 Owner-only Ruleset 的禁强推）。
2. `scripts/force-push.sh`（SSH over 443 部署密钥执行 force push）。
3. **回读校验**：远端 `git rev-parse main` == 本地 HEAD；抽查远端已无被删的第三方路径。
4. `scripts/ruleset-toggle.sh enable`（立即恢复 Ruleset）。
5. 记录每阶段推送时间、commit、校验结果到本文件日志。

> 强推属高风险操作：每阶段必须串行执行，禁止多阶段合并强推；任何一步失败立即停止并回滚 Ruleset 状态。

---

## 5. 诚实性总纲（全阶段强制）

1. **禁止夸大**：不得把"文档声称 GA v1.0.0"当作"社区验证的成熟度"；不得把静态近似说成真实执行。
2. **禁止伪造**：不得伪造 benchmark 分数、竞品对拍分数、哈希、SPDX 证据、commit 历史。
3. **禁止外推**：单文件（首 unit）口径的 degrade 结论不得外推到全语料/全语言。
4. **口径与实现一致**：README（中/英）、FAQ、CHANGELOG、SECURITY 等对外文案必须与代码真实能力同步；`NOT_MEASURED`/`RATE_LIMIT_PENDING` 未真实测量前保持原样。
5. **法律事实由权利人/法务定**：AI 只列待确认项，不写"已合规""无侵权"等法律结论。

---

## 6. 分期任务表

优先级说明：**P0 = 必须（合规红线/编译依赖）**，**P1 = 应该（整改与重写）**，**P2 = 可选增强（能力提升）**。（此 P0/P1/P2 为任务优先级，与阶段号 P0–P5 不同。）

### 阶段 P0　清除不参与编译的第三方源码（可立即执行，零编译影响）

| 编号 | 标题 | 具体动作 | 涉及文件/目录 | 依赖 | 验收标准 | 优先级 | 风险 |
|---|---|---|---|---|---|---|---|
| P0-01 | 建立第三方资产 SSOT 台账（清除前） | 逐目录记录路径/体积/来源 commit/许可证/sha256/处置决策（删 or 重写 or 保留待法务）；作为清除审计证据，删除后不可复原 | `third_party/`、新建 `docs/refactor/assets-inventory.md` | — | 台账覆盖 §2.3–2.5 全部项；每项有处置决策与责任人 | P0 | 漏登导致后续无法证明"清干净" |
| P0-02 | 删除 `third_party/_sources/` 全部源码副本 | 物理删除 527MB 全部子目录（gvisor/llvm/roslyn/kotlin/pit/swc/syn/tree-sitter-*/wazero/_validation） | `third_party/_sources/` | P0-01 | 目录不存在；`git status` 无残留；`go list ./...` 仍成功 | P0 | 误删产品代码（均 `_` 前缀，风险低） |
| P0-03 | 删除 `third_party/_references/` 全部参考副本 | 删除 go-cmp/toxiproxy/cobra/go-gremlins/`_m5_algorithms.go.txt`/sarif schema/MANIFEST | `third_party/_references/` | P0-01 | 目录不存在；编译无影响 | P0 | 低 |
| P0-04 | 清除 benchmark 残留开源源码 | 删除 `cases/*/src/`、`config/node_modules/`；`audit/raw` 中若含源码片段则剔除；停用 `scripts/corpus_fetch.sh` 下载逻辑并标注"已废弃" | `docs/benchmark/`、`scripts/corpus_fetch.sh` | P0-01 | `docs/benchmark/` 下无任何第三方源码文件（扩展名/版权头扫描为空）；结果数据保留需带"历史未对拍"标注 | P0 | benchmark 数据若含复制代码片段需人工甄别 |
| P0-05 | 改写第三方台账/清单 | 将 `LICENSE-MANIFEST.md`、`ci/corpus_manifest.json`、`THIRD-PARTY-NOTICES*`、`NOTICE` 中"已接入/只读落地"条目改为"已清除（日期+commit）"，保留历史审计但不声称仍接入 | `third_party/LICENSE-MANIFEST.md`、`ci/corpus_manifest.json`、`THIRD-PARTY-NOTICES*.md`、`NOTICE` | P0-02,P0-03,P0-04 | 无"当前接入"表述；每条注明清除状态 | P0 | 台账与事实不符 |

**阶段出口门**：`third_party/` 下仅剩（或不含）任何源码副本；benchmark 无第三方源码；编译仍通过；P0-01 台账归档。**通过后执行推送策略（第 4 节）。**

---

### 阶段 P1　自研替换核心编译依赖（wazero / yaml.v3 / x/sys）

| 编号 | 标题 | 具体动作 | 涉及文件/目录 | 依赖 | 验收标准 | 优先级 | 风险 |
|---|---|---|---|---|---|---|---|
| P1-01 | 决策并记录替换方案（ADR） | 为 wazero / yaml.v3 各出一份 ADR：选"最小自研实现"还是"改自研 ABI / 自研配置格式"；明确能力边界与测试策略 | `docs/adr/`（新增） | P0 出口 | ADR 经架构师签字；明确 do/don't | P0 | 方案选错导致返工 |
| P1-02 | 自研替换 wazero（WASM 运行时） | 移除全部 `import github.com/tetratelabs/wazero`；落地自研运行时或自研插件 ABI；改写 `plugin/wasm/`、`plugin/host/`、`plugin/sandbox/` 相关调用与测试 | `plugin/wasm/`、`plugin/host/`、`plugin/sandbox/`、`go.mod`、`go.sum` | P1-01 | 全仓无 wazero import；插件加载/调用/hotswap 测试通过；功能与性能基线记录在案 | P0 | **工作量最大**；插件语义兼容风险高 |
| P1-03 | 自研替换 yaml.v3 | 移除全部 `import gopkg.in/yaml.v3`；落地自研最小 YAML 子集解析器或改用项目自有配置格式；改写 `ctl/cli/`、`ci`、`config` 相关解析与 testdata | `ctl/cli/`、`ctl/rest/`、`cmd/`、`ci/`、`go.mod`、`go.sum` | P1-01 | 全仓无 yaml.v3 import；配置解析测试（含 invalid/malformed/ok/violated）通过 | P0 | YAML 子集覆盖不全导致配置回退 |
| P1-04 | 消除 x/sys 间接依赖 | 确认 x/sys 由 wazero 引入；P1-02 后若仍存在直接 syscall 使用，改用标准库 `syscall`/平台抽象重写 | `go.mod`、`go.sum`、涉及 syscall 的文件 | P1-02 | `go list -m all` 无 `golang.org/x/sys` | P0 | 低 |
| P1-05 | 清零 go.mod / go.sum 并删除 vendor | `go mod tidy`；确认 require 段无第三方；物理删除 `vendor/`（含 modules.txt） | `go.mod`、`go.sum`、`vendor/` | P1-02,P1-03,P1-04 | `require` 无第三方；`go.sum` 无第三方校验行；`vendor/` 不存在；`go build ./...` 与 `go test -race ./...` 通过；`license-check` 门禁 PASS | P0 | 构建中断（须先在分支验证） |

**阶段出口门**：`go.mod` 第三方依赖 = 0，`vendor/` 删除，全仓 import 扫描无第三方，构建/单测通过。**通过后推送。**

---

### 阶段 P2　甄别并重写"融入产品代码"的开源改写

| 编号 | 标题 | 具体动作 | 涉及文件/目录 | 依赖 | 验收标准 | 优先级 | 风险 |
|---|---|---|---|---|---|---|---|
| P2-01 | 甄别 `kernel/mutation`（PIT 疑似） | 逐文件判定"自研 vs PIT 算法移植"，产出逐文件结论；移植部分用自研变异算子/最小化算法等价重写 | `kernel/mutation/`（`derived.go`/`engine.go`/`minimize.go`/`discover.go` 等） | P0 出口 | 甄别台账每文件有结论；移植代码重写完成且有独立测试 | P0 | 算法等价性难证；须配 parity 测试 |
| P2-02 | 甄别 `kernel/injection` 与 `plugin`（chaosblade 疑似） | 甄别注入算子（`ops_*.go`、`injector.go`、`world.go`）是否 chaosblade 移植；移植部分重写；确认 wazero 宿主残留已在 P1 清除 | `kernel/injection/`、`plugin/` | P1 出口 | 无 chaosblade/wazero 源码痕迹；注入 e2e/chaos 测试通过 | P0 | 中 |
| P2-03 | 甄别 `kernel/abi`（gVisor ABI 疑似） | 判定 `amd64Names_*`/`arch_*`/`syscall_table.go`/`addr*.go` 是否为 gVisor 移植；属移植则改为自研生成的 syscall 表；属事实接口数据的走第 8 节法务确认后保留并标注 | `kernel/abi/` | P0 出口 | 甄别台账完整；移植代码重写或有法务确认标注 | P0 | 事实数据版权认定需法务介入 |
| P2-04 | 全仓扫描其他潜在并入 | 扫描 `sim/rand`（SplitMix64）、`lang/*front`（tree-sitter 语法衍生）、`scripts/build-libclang-min.sh`（libclang）、`docs/adr/` 引用；逐一判定 | `sim/`、`lang/`、`scripts/`、`docs/adr/` | P0 出口 | 扫描报告覆盖全部 Go/脚本/语法文件；无未判定项 | P1 | 隐藏衍生难发现 |
| P2-05 | 更新许可/通知文档为"零第三方源码" | 重写 `THIRD-PARTY-NOTICES*`、`NOTICE`、`LICENSE-COMPATIBILITY*`，移除"并入组件"清单，保留"已清除"历史 | `THIRD-PARTY-NOTICES*.md`、`NOTICE`、`LICENSE-COMPATIBILITY*.md` | P2-01..04 | 文档与甄别台账一致；无自相矛盾表述 | P1 | ⚖️ 涉及许可声明，须法务复核后再定稿 |

**阶段出口门**：甄别台账 100% 覆盖，所有判定为"移植"的代码完成重写，全仓无上游版权头/来源注释残留（header-check PASS）。**通过后推送。**

---

### 阶段 P3　保真度 / derive / 竞品 整改（专项 S01–S03 的技术执行）

| 编号 | 标题 | 具体动作 | 涉及文件/目录 | 依赖 | 验收标准 | 优先级 | 风险 |
|---|---|---|---|---|---|---|---|
| P3-01 | AIR 多语保真度处理（对应 S01） | 二选一：①自研提升单文件→全语料 fidelity（设阈值与回归基线）；②保持现状但锁死口径。见第 7 节 S01 | `lang/air/`、`lang/*front/`、`ci/e2e_thresholds*.json`、`README*.md` | P2 出口 | 有明确决策记录；实现与对外口径一致；未真实测量前 `NOT_MEASURED` 不改写 | P1 | 提升工作量大；口径外推风险 |
| P3-02 | derive 杀死率处理（对应 S02） | 二选一：①改为真实执行 go test（需沙箱+确定性设计）；②保持静态 AIR 近似并明确标注。见 S02 | `lang/derive/`、`ctl/cli/cmd_derive.go`、`README*.md` | P2 出口 | 决策记录到位；README 明确区分近似与真实执行 | P1 | 真实执行引入供应链/确定性复杂度 |
| P3-03 | 竞品对比处理（对应 S03） | 二选一：①合法真实对拍（不得引入开源源码）；②保持 `NOT_MEASURED`。见 S03 | `README*.md`、`docs/benchmark/`、`docs/reference/` | P2 出口 | 未对拍则 `competitorAbsolute=NOT_MEASURED` 不变；对拍则有方法与原始数据，且未引入第三方源码 | P1 | 对拍数据不可伪造 |

**阶段出口门**：三项决策各有书面记录，README 中/英口径同步且与实际一致。**通过后推送。**

---

### 阶段 P4　许可合规与命名澄清（专项 S04–S06 + 法务红线）

| 编号 | 标题 | 具体动作 | 涉及文件/目录 | 依赖 | 验收标准 | 优先级 | 风险 |
|---|---|---|---|---|---|---|---|
| P4-01 | 对外叙事校准（对应 S04） | 区分"文档声称 GA v1.0.0"与"社区可验证使用量"；补一段成熟度诚实说明，禁止把文档体量当成熟度 | `README*.md`、`CHANGELOG*.md`、`SECURITY*.md` | P3 出口 | 文档不再暗示社区成熟度；0 star/fork/watcher 事实如实呈现 | P1 | 措辞需权威把关 |
| P4-02 | 许可声明整改（对应 S05，⚖️） | 按第 8 节 L01–L04 逐项列出待法务确认事项；**AI 不改法律结论**，仅准备材料与占位标注 | `LICENSING*.md`、`FAQ-licensing*.md`、`LICENSE-COMPATIBILITY*.md`、`COMMERCIAL-LICENSE*.md`、`LICENSE`、`NOTICE` | P3 出口 | 每项标注"需权利人/法务确认"与责任人；无 AI 法律结论 | P0 | 法务结论不可由 AI 代替 |
| P4-03 | 命名澄清文档（对应 S06） | 新增命名说明：`ces.tencent.com/v1` 仅为 API 组名，项目与腾讯无关，版权主体"小黎的部落"；在 README/API 文档显著位置提示 | `README*.md`、`TRADEMARK*.md`、`ces/`、`docs/user/` | P3 出口 | 有独立命名澄清段落；中英一致；避免误导 | P1 | 低 |

**阶段出口门**：法务红线清单逐项有"待确认/已确认"状态与责任人；对外文案口径一致；无 AI 出具的法律结论。**通过后推送。**

---

### 阶段 P5　门禁复绿与 QA（每阶段出口 + 最终收尾）

| 编号 | 标题 | 具体动作 | 涉及文件/目录 | 依赖 | 验收标准 | 优先级 | 风险 |
|---|---|---|---|---|---|---|---|
| P5-01 | 重算水印与哈希 | 清除文件后文件集变化，重算 `WATERMARK.sha256`、`SHA256SUMS`；确认 `hash-check.sh`、`watermark-verify.sh` 通过（须走正式流水线，不手工伪造） | `WATERMARK.sha256`、`SHA256SUMS`、`scripts/hash-check.sh`、`scripts/watermark-verify.sh` | 各阶段出口 | 两脚本 PASS；哈希由工具生成且有日志 | P0 | 手工伪造哈希=红线违规 |
| P5-02 | 门禁十段全绿 | 跑 `scripts/lint-gate.sh`（0–7 + 5b）与 `coverage-gate.sh`；修复因删除/重写导致的失败 | `scripts/lint-gate.sh` 及被检文件 | P5-01 | 全段 PASS，零 FAIL | P0 | 单文件≤300行/一类一文件等风格约束可能因重写触发 |
| P5-03 | TLA+ 与一致性检查 | 跑 `scripts/tla/`、`drift-check.sh`、`determinism-check.sh`，确认规格与实现一致 | `scripts/tla/`、`scripts/drift-check.sh`、`scripts/determinism-check.sh` | P5-02 | 全部 PASS | P0 | 重写破坏规格一致性 |
| P5-04 | 全仓"零第三方源码"审计 | 跑 `m5-reference-scan.sh` + 自研全盘扫描（版权头/上游 URL/路径名/扩展名/manifest），产出审计报告 | `scripts/m5-reference-scan.sh`、新建 `docs/refactor/zero-third-party-audit.md` | P2 出口 | 报告结论：第三方源码 = 0；有扫描命令与原始输出 | P0 | 扫描规则不全导致漏检 |
| P5-05 | QA 终审 + 阶段推送 | 汇总 DoD 勾选；按第 4 节流程执行本阶段强推并回读校验；记录日志 | 本文件、`scripts/force-push.sh`、`scripts/ruleset-toggle.sh` | P5-01..04 全过 | DoD 全勾选；远端校验通过；Ruleset 已恢复 enable | P0 | 强推误操作 |

---

## 7. 专项章节：6 个"需特别注意的问题"（逐条：自研后如何处理 + 对外口径怎么改）

### S01　多语言 AIR 保真度真实短板

- **现状事实**：README L133–135 已诚实声明——degrade 达标**仅限冻结单文件（首 unit）**；同一公式跑全语料，六语全部超阈：js 49.12% / c 28.00% / py 26.89% / rs 42.33% / kt 8.21% / cs 6.98%，go `NOT_MEASURED`。多语言源码保真度实际不好。
- **自研后如何处理（技术，对应 P3-01）**：
  - 方案 A（能力优先，P2）：自研提升 AIR 保真度，把全语料 degrade 降到阈值内；必须先建立可复现基线（`ci/e2e_thresholds*.json` 冻结），再逐语优化，禁止改公式凑数。
  - 方案 B（诚实优先，P1）：不自研提升，维持现状，但把技术债与"仅限单文件口径"写进 README 与 ADR，列为后续路线图。
  - 推荐：**先 B 保合规诚实，A 作为独立增强里程碑评估**。
- **对外口径怎么改**：README（中/英）保留"单文件口径"限定语，禁止出现"全语料达标/多语言高保真"；对外结论附具体数字并注明 `go=NOT_MEASURED`；如未来真实提升，须用新测量数据替换后方可改口径。
- **验收**：口径与实现一致；无任何"全语料达标"表述；如提升则有新基线数据支撑。

### S02　derive 变异杀死率是静态近似，非真实执行

- **现状事实**：README L136–137 声明 killing 是 *static AIR-level approximation, not a real go test execution*。
- **自研后如何处理（对应 P3-02）**：
  - 方案 A（严谨优先，P2）：改为真实执行 go test——需自研沙箱执行器，保证确定性（固定种子）、隔离、可复现，并防止引入任何第三方执行引擎。
  - 方案 B（诚实优先，P1）：保持静态近似，但在 `ctl/cli/cmd_derive.go` 输出与 README 明确标注"静态近似，非真实执行"。
  - 推荐：**先 B；A 作为增强项，且不得借 A 引入开源测试框架**。
- **对外口径怎么改**：所有 killing/杀死率指标旁标注方法学（static approximation）；禁止表述为"真实测试杀死率/实测覆盖率"。
- **验收**：输出与文档口径一致；无"实测 go test"误导表述。

### S03　竞品对比分数一律 NOT_MEASURED

- **现状事实**：README L132 `competitorAbsolute = NOT_MEASURED`，从未真实对拍。
- **自研后如何处理（对应 P3-03）**：
  - 方案 A（P2）：合法获取竞品做真实对拍，公开方法、原始数据、版本与环境；**对拍过程不得把竞品源码落地到本仓库**（否则违反清除范围）。
  - 方案 B（P1，推荐）：保持 `NOT_MEASURED`，不写任何竞品绝对分数。
- **对外口径怎么改**：未真实对拍前 `NOT_MEASURED` 原样保留；竞品章节只做定性、可溯源描述，禁止编造或引用未经测量的分数象限图。
- **验收**：文档中竞品绝对分数字段 = `NOT_MEASURED`（或真实测量后有原始数据）；无伪造分数。

### S04　成熟度信号弱（文档与治理先行）

- **现状事实**：0 star / 0 fork / 0 watcher / 3 commit；双语 README/CLA/DCO/TRADEMARK/SHA256SUMS/WATERMARK 等治理体量极大；文档声称 GA `v1.0.0`。
- **自研后如何处理（对应 P4-01）**：叙事校准——明确区分"文档与治理声明"和"社区可验证的实际使用量"；重构后 commit 历史会被强推改写，更不得把 commit 数量/文档体量当成熟度证据。
- **对外口径怎么改**：README/CHANGELOG 增补"项目为个人或小团队项目、早期阶段、社区采用数据有限"的诚实说明；"GA v1.0.0"保留为版本标签但不得暗示广泛生产验证。
- **验收**：无"成熟/广泛采用/生产级验证"等无证据表述；成熟度表述与可验证数据一致。

### S05　许可风险（⚖️，见第 8 节法务红线）

- **现状事实**：双许可 AGPLv3 + 商业；缺口① AGPLv3 §13 远程网络交互义务；缺口② DCO 不授予再许可权，双许可需 CLA 或明确再许可条款；SPDX `AGPL-3.0-or-later` 标注"需权利人确认"；chaosblade（Apache-2.0）源码二次并入、PIT 算法移植待商用前复核。
- **自研后如何处理（对应 P4-02）**：清除开源源码后，**并入类许可风险（chaosblade/PIT/gVisor）随源码清除而消除**；剩余为项目自身对外授权（AGPL §13、DCO/CLA、SPDX）问题——须由权利人/法务确认，**AI 不改写法律结论**，仅准备待确认清单与占位标注。
- **对外口径怎么改**：在法务书面确认前，许可文档对争议点标注"待权利人/法务确认"，不擅自声称"已完全合规/无侵权"；SPDX 标注保持"需权利人确认"。
- **验收**：法务红线 L01–L04 逐项有状态与责任人；无 AI 法律结论。

### S06　命名易误读（`ces.tencent.com/v1`）

- **现状事实**：`ces.tencent.com/v1` 只是 API 组名，项目与腾讯无关；版权主体"小黎的部落"，属个人/小团队项目。
- **自研后如何处理（对应 P4-03）**：新增命名澄清文档与 README 显著提示；检查所有 API 组名、注释、文档中的 `tencent.com` 引用，确保不被误读为公司背书。
- **对外口径怎么改**：明确 "API group name only; not affiliated with Tencent; copyright © 小黎的部落"；中英一致。
- **验收**：命名澄清段落到位；无暗示企业关联的表述。

---

## 8. 法务红线（⚖️　AI 不出法律结论，须权利人/法务书面确认）

> 本节仅列**事实与待确认问题**，不构成法律意见。任何对外许可声明的修改，必须待权利人/法务书面确认后，由授权人执行。

| 编号 | 待确认事项 | 事实依据 | 需谁确认 | 当前标注 | 关联任务 |
|---|---|---|---|---|---|
| L01 | AGPLv3 §13 远程网络交互义务 | 若以 SaaS/网络服务形式提供，须向网络用户提供对应源码；项目是否有此场景、如何履行，待定 | 权利人 + 法务 | 待确认 | P4-02 / S05 |
| L02 | DCO 与双许可再许可权 | DCO 不构成再许可授权；AGPLv3 + 商业双许可需 CLA 或明确再许可条款支撑 | 权利人 + 法务 | 待确认 | P4-02 / S05 |
| L03 | SPDX `AGPL-3.0-or-later` 标注 | 现标注"需权利人确认"；确认前不得改为其他表达式 | 权利人 | 待确认（保持原标注） | P4-02 / S05 |
| L04 | chaosblade(Apache-2.0)/PIT 等并入兼容性 | 商用前的许可兼容性复核；**源码清除后并入风险应已消除，但须法务确认历史并入无遗留义务** | 法务 | 待确认 | P2-05 / P4-02 |
| L05 | gVisor ABI / 事实数据版权认定 | syscall 表等是否属不受版权保护的事实数据，保留需法务认定 | 法务 | 待确认 | P2-03 |

---

## 9. 完成定义（DoD — 全部勾选才算重构完成）

### 9.1 清除完整性（硬性 = 0）

- [ ] `go.mod` 第三方 require = **0**（`go list -m all` 仅 `abate` 自身）
- [ ] `go.sum` 第三方校验行 = **0**
- [ ] `vendor/` 目录 = **不存在**；vendor 第三方源码 = **0**
- [ ] `third_party/_sources/` = **不存在**；`third_party/_references/` = **不存在**；third_party 源码副本 = **0**
- [ ] `docs/benchmark/` 下第三方源码（lodash/cJSON/src/node_modules）= **0**
- [ ] 全仓 `import` 扫描无 `github.com/*`、`gopkg.in/*`、`golang.org/x/*`
- [ ] 融入代码甄别台账 100% 覆盖；判定为"移植"的代码全部重写为自研（或经法务确认标注）
- [ ] 全仓"零第三方源码"审计报告（`zero-third-party-audit.md`）结论 = 0，附扫描命令与原始输出

### 9.2 构建与门禁（全绿）

- [ ] `go build ./...` 通过；`go test -race -count=1 ./...` 通过
- [ ] `scripts/lint-gate.sh` 十段（0/1/2/3/4/5/5b/6/7）全部 PASS
- [ ] `scripts/coverage-gate.sh` PASS
- [ ] `scripts/tla/`、`drift-check.sh`、`determinism-check.sh` PASS
- [ ] `scripts/hash-check.sh`、`scripts/watermark-verify.sh` PASS（哈希由工具生成，无伪造）
- [ ] `scripts/m5-reference-scan.sh`、`scripts/pit-parity.sh` PASS

### 9.3 口径整改（诚实性）

- [ ] README（中/英）无夸大/伪造/外推表述
- [ ] S01：保真度口径与实现一致，未测量项保持 `NOT_MEASURED`
- [ ] S02：derive killing 明确标注"静态近似/真实执行"方法学
- [ ] S03：竞品绝对分数保持 `NOT_MEASURED`（或有真实对拍原始数据）
- [ ] S04：成熟度叙事区分"文档声明"与"可验证使用量"
- [ ] S06：命名澄清文档到位，无企业关联误导

### 9.4 法务与推送

- [ ] 法务红线 L01–L05 逐项有"待确认/已确认"状态与责任人；无 AI 法律结论
- [ ] 许可/通知文档与"零第三方源码"事实一致，争议点标注"待权利人/法务确认"
- [ ] 每阶段按推送策略强推并回读校验通过；Ruleset 已恢复 `enable`
- [ ] 本文件 DoD 全部勾选，重构状态标记为完成

---

## 10. 执行日志（每阶段回写）

| 阶段 | 完成日期 | commit | 强推时间 | 回读校验 | Ruleset 恢复 | 执行人 |
|---|---|---|---|---|---|---|
| P0 | | | | | | |
| P1 | | | | | | |
| P2 | | | | | | |
| P3 | | | | | | |
| P4 | | | | | | |
| P5（最终） | | | | | | |

---

> 结尾提醒：本清单是"先规划、后执行"的产物。真正的合规来自**证据**（零第三方源码扫描报告、全绿门禁、真实口径），而非声明。任何"看起来合规但无法验证"的改动，都违反第 5 节诚实性总纲。
