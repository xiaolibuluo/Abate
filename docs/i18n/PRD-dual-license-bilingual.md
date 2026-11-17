# PRD：Abate「AGPLv3 + 商业授权双许可」与「全量中英双语文档」整理

> 阶段 1/4 产物 · Product Manager（接替者）· 交给阶段 2 架构师的**唯一输入**
> 本文件是 PRD 文档，不含、也不修改任何 `.go/.sh/.json/.yml`；不含联网调研结论。
> 全部事实采信自主理人交接件 `HANDOFF_docs.md`（下称「交接件」，§n 指其章节）与 GNU 官方 `agpl-3.0.txt`。凡交接件未固化的信息，一律标注「待确认 / 由架构师定案」，不虚构。

---

## 0. 结论先行 · 给架构师的关键约束摘要（≤200 字）

> 推荐**同目录双文件**（`X.md` 中文 / `X.en.md` 英文）：相对链接改动最小、冻结文件只需新建译本不动原文、中英逐节对应可机械核验；README 主语言待主理人定，**最终由架构师定案**。
> **P0 九项**：AGPLv3 全文 `LICENSE`（逐字 + sha256）、商业授权说明、CLA/再许可缺口标注待确认、`LICENSING` 总览、README 双语重构、用户向 3 份双语、≥8 条保姆式教程、governance 三件套双语、夸大与版本清理。
> **诚实红线**：`competitorAbsolute = NOT_MEASURED`，零竞品绝对断言、零夸大虚构，版本统一 Go 1.26.11，规划/事实分离，`derive` 保留「非真实执行 go test」措辞，教程命令逐字对齐交接件 §6d。
> **待确认**：CLA 是否改必需、商业邮箱、SSPL 类表述、only/or-later、商标政策、命名方案。

---

## 1. 项目信息

| 项 | 内容 |
|---|---|
| Language | 简体中文（与用户诉求一致）；交付物另含英文译本 |
| Programming Language | **不适用**：本轮为纯「文档 + 许可」工程，不写应用代码、不改产品代码。关联仓库技术栈为 Go 1.26.11（`GOFLAGS=-mod=vendor` 离线构建）。默认 Web 模板（Vite+React+MUI+Tailwind）与本任务无关，不套用 |
| Project Name | `abate_dual_license_bilingual_docs`（snake_case） |
| 项目实体 | Abate：`module abate`，七语言（go/py/js/c/cs/rs/kt）源码保真度 + 变异 + 混沌测试工具；仓库 `/workspace/abate`，**无 git remote** |
| 原始需求复述 | 「严格按照国际开源 AGPLv3 和商业授权双许可标准进行整理所有全量中英双语文档。必须实事求是，禁止夸大其词。保姆式使用教程文档等等必须整理好。」 |
| 四条硬要求 | ①严格 AGPLv3 + 商业双许可国际标准；②全量文档中英双语；③实事求是禁止夸大（**最高红线**）；④保姆式 step-by-step 教程齐备可照做 |
| 本 PRD 边界 | 只产出需求与设计约束；不落任何许可/文档正文，不执行 `git add/commit/push`，不联网 |

---

## 2. 采信事实基线（交接件已核实，本轮不再调研）

| 维度 | 已核实事实 | 来源 |
|---|---|---|
| 版本 | 真实 `go 1.26.11`；`GOFLAGS=-mod=vendor`；依赖 wazero v1.12.0 / yaml.v3 v3.0.1 / x-sys v0.44.0 | §2 |
| **版本不实** | README 写 `go1.27`、quickstart 写 `Go 1.24+`，**均须纠正为 1.26.11**；README 的 `golangci-lint v2.15.0` 未核实，**禁照抄断言** | §2 / §6c |
| 许可现状 | 仓库根**无 `LICENSE`/`COPYING`**；仅 `governance/{DUAL-LICENSE,HEADERS,THIRD-PARTY}.md` 三份纯中文（均带 6 行 HTML 注释版权头） | §2 |
| 版权主体 | 「小黎的部落 (www.qingdujian.com)」；商业联系 `xiaolibuluo@163.com`；来源 `https://github.com/xiaolibuluo/Abate`（既有真实署名，**沿用、勿扩写为不存在的资质**） | §2 |
| **法律缺口** | 治理把「CLA 标可选、DCO 强制」；双许可下 DCO **不授予项目方再许可（relicense）权**，需 CLA 或明确再许可条款。标注「**需权利人或法务确认**」，**不擅自下法律结论** | §2 |
| 文档总量 | 排除 `.git/vendor/third_party/.artifacts` 后 **56 份 `.md`**，**全中文、0 英文、0 双语** | §2 |
| README | 仅 CI/CD 门禁说明，无项目介绍/功能/安装/许可/教程入口 | §2 / §6c |
| 夸大分布 | 全仓 129 处「100%/唯一/绝对」类命中，**绝大多数在 `docs/plan/` 且为「验收目标 100% 通过」规划表述，非既成事实** | §3 |
| **决定性红线** | E2E 报告 `competitorAbsolute = NOT_MEASURED`；冻结 `docs/system_design.md` 规定：未真实对拍 tree-sitter/PIT/Stryker/cargo-mutants/mutmut/libFuzzer 前，**不写任何竞品绝对分数、更不宣称「超过竞品 X%」** | §3 |
| 竞品许可（客观陈述，勿写对标超越） | Qt(GPL/LGPL/商业)、Grafana(AGPLv3+商业)、iText(AGPL/商业) 为双许可参考；MongoDB(SSPL)、Elastic(SSPL/Elastic License)、HashiCorp(BUSL) **均非 OSI 认可开源许可** | §4 |
| AGPLv3 黄金基准 | `agpl-3.0.txt`：GNU/FSF 官方，661 行 / 34523 字节，`sha256 = 0d96a4ff68ad6d4b6f1f30f713b18d5184912ba8dd389f86aa7710db079abcb0`；第 13 节「Remote Network Interaction」在 line 540；结尾「How to Apply」含 `either version 3 ... or (at your option) any later version` | §4 |
| SPDX | `AGPL-3.0-only` / `AGPL-3.0-or-later`（**选哪个待确认**） | 用户诉求 |
| 双许可标准文件集 | `LICENSE`(AGPLv3 全文)、`LICENSING`/`COMMERCIAL-LICENSE`、`CONTRIBUTING`、`CLA`、`DCO`、`NOTICE`、`TRADEMARK`、`THIRD-PARTY-NOTICES`、`SECURITY`、`CODE_OF_CONDUCT`、许可证兼容性说明、`FAQ`(何时需商业授权) | 用户诉求 |

### 2.1 门禁纠偏（§6b，措辞必须精确，勿沿用旧 brief 的错误说法）

- `scripts/header-check.sh` **只扫 `*.go`**（`find . -name '*.go'`），**完全不扫 `.md`** → 新建双语 `.md` 缺版权头**不会**触发 header-check FAIL。
- 但 `governance/HEADERS.md` §2 明确「覆盖范围含 `.md`（文档头部 HTML 注释 / YAML front-matter）」，这是**治理约定**，非脚本门禁。
- 正解：新建双语 `.md` **应**带 6 行 HTML 注释版权头（含 `Copyright` + `xiaolibuluo/Abate` + `小黎的部落`），以符合治理约定与既有 56 份文档惯例；架构师须把它列为**治理一致性要求**，而非门禁强制项。
- `scripts/license-check.sh` 的 GPL 扫描只针对 `vendor/modules.txt` 依赖 LICENSE 文本 + `third_party/LICENSE-MANIFEST.md`（后者命中 GPL 字样仅 `[WARN]` 不 FAIL），**不扫仓库根 `LICENSE`** → 根落地 AGPLv3 全文 `LICENSE`（含 "GNU AFFERO GENERAL PUBLIC LICENSE"）**不会**触发 license-check FAIL。
- 改动前实测基线：`bash scripts/header-check.sh` = PASSED、`bash scripts/license-check.sh` = PASSED（wazero/x-sys → ALLOW、yaml.v3 → MIT；LICENSE-MANIFEST 的 GPL 字样为预期 WARN）。`lint-gate.sh` 第 [0] 段调 header-check、第 [2] 段调 license-check，均在冻结九段内，不改。

### 2.2 保姆式教程真实 CLI 证据（§6d，教程须逐字引用，禁照抄旧 cli-reference.md）

> 用源码实时构建 `GOFLAGS=-mod=vendor go build -o /tmp/abate_verify .`（**非 11-08 旧二进制**）真跑所得。为与证据逐字一致，下文以 `abate` 指代该二进制（可 `alias abate=/tmp/abate_verify` 或复制入 PATH）；`ok.yaml` 为 `ctl/cli/testdata/ok.yaml` 的简写，教程须写明此等价关系。

- `abate version` 真实输出（6 行）：
  ```
  abate 0.1.0
  schema: ces.tencent.com/v1
  build: deterministic
  license: AGPLv3 or commercial
  source: https://github.com/xiaolibuluo/Abate
  watermark: 小黎的部落
  ```
- 示例 CES 文件真实存在：`ctl/cli/testdata/{ok,violated,invalid,malformed}.yaml`（ok.yaml 36 行）。
- `abate run ctl/cli/testdata/ok.yaml` 报告头：`schema: abate.report.v1` / `apiVersion: ces.tencent.com/v1` / `name: cli-e2e-ok` / `seed: 12345` / `status: COMPLETED` / `exitCode: 0` / `events: 6` / `observations: 4` / `runDigest: 6444a24ed2…` / `health | HELD | held span 3000000000 satisfies holdFor 1000000000`。
- `abate run --dry-run ok.yaml` → `status: VALIDATED` / `exitCode: 0` / `events: 0` / `observations: 0`。
- `abate run ok.yaml --seed 42 --format json` → 标准 JSON，`"seed": 42` / `"status": "COMPLETED"` / `"runDigest": "32dbea3d9…"`（**seed 覆盖会改变 digest，若宣称确定性须固定 seed**）。
- `--help` 命令树：`run/verify/replay/check/derive/plugin/web/version`；退出码四态 `0/1/2/3` 声明。
- 退出码干净实测：`ok.yaml → 0`、`violated.yaml → 1`、`invalid.yaml → 3`、`malformed.yaml → 3`、`无参数 → 3`、`version → 0`。
- `derive --help`：注明「Killing is a static AIR-level approximation, **not a real go test execution**」→ **不得**把 derive 说成「真实执行 go test 的变异测试」，此诚实措辞须保留。
- 可用 flag：`--seed uint` / `--format console|json` / `--out path` / `--record path` / `--dry-run` / `--strict`。

---

## 3. 产品目标（3 个正交目标）

- **G1 合规完整（Must）**：按 AGPLv3 + 商业授权双许可国际标准补齐仓库许可文件集，使「开源侧许可」「商业侧入口」「再许可权利链」三件事**文件齐备、可机械核验、法律缺口显式标注**。可测量：§7 验收清单 8 项全部通过。
- **G2 全量双语可达（Must）**：56 份 `.md` **100% 覆盖**（含「全量双语 / 中文+英文摘要 / 内部规划降级」三类处置，无遗漏），英文读者无需读中文即可完成**安装 → 第一个实验 → CLI 全命令 → 写 CES → 贡献 → 判断授权 → 跑门禁**。可测量：分组处置表 14 组合计 56 份、保姆式教程 ≥8 条逐条可照做。
- **G3 诚实可信（最高红线）**：所有能力/版本/竞品/许可描述**只基于可验证命令、真实文件、实测数据**；零夸大词、零虚构、零竞品绝对断言，「设计目标」与「已实测事实」严格分离。可测量：夸大词扫描命中归零（规划性表述除外且须保留目标语气）、版本统一 1.26.11、竞品绝对分数出现次数 = 0。

---

## 4. 用户故事（≥5 类）

| # | 角色 | 用户故事 |
|---|---|---|
| US1 | **开源自托管用户** | As an open-source self-hosting user, I want 一份英文 README + 从零离线构建/第一个实验的保姆式教程, so that 我能在无外网、Go 1.26.11 + vendor 环境下确认 AGPLv3 义务并跑起来，而不必读中文或猜命令。 |
| US2 | **闭源商业集成决策者** | As a decision-maker integrating Abate into a closed-source product, I want 清晰的双许可总览、商业授权说明、FAQ 与「何时需商业授权」判断路径（含 AGPLv3 §13 网络交互义务）, so that 我能在不下法律结论的前提下判断走 AGPLv3 还是联系商业授权。 |
| US3 | **贡献者** | As a contributor, I want 中英双语 CONTRIBUTING / DCO / CLA 现状与「再许可权缺口待确认」的诚实标注, so that 我知道 `git commit -s`（DCO 1.1）怎么做、以及贡献的版权权利如何被项目方使用。 |
| US4 | **运维 / CI 管理员** | As an operator, I want 双语的门禁运行说明（header-check / license-check / lint-gate）与「根 LICENSE 不触发 FAIL」「header-check 只扫 *.go」的准确说明, so that 我能正确复跑门禁、不误判告警、不误改冻结段。 |
| US5 | **七语言前端开发者**（go/py/js/c/cs/rs/kt） | As a multi-language frontend developer, I want 双语的 CES 指南、CLI 参考与 frontend-coverage 文档, so that 我能按真实命令写 CES、理解 `derive` 是静态 AIR 近似而非真实 go test 变异测试，避免被能力措辞误导。 |
| US6 | 法务 / 合规审阅者（附加） | As a compliance reviewer, I want SPDX 标识、第三方许可清单、OSI 认可与否的严谨区分、以及所有待确认法律点的显式列表, so that 我能定位哪些结论需权利人/法务拍板而非由文档代断。 |

---

## 5. 需求池（P0 Must / P1 Should / P2 Could）

### 5.1 P0 —— 必须（用户硬要求，逐项带可测量判据）

| ID | 需求 | 语言 | 可测量判据 |
|---|---|---|---|
| P0-1 | **AGPLv3 全文 `LICENSE`** 落地仓库根 | MUST | 与 `agpl-3.0.txt` **逐字 diff 无差异**；661 行 / 34523 字节；`sha256 = 0d96a4ff68ad6d4b6f1f30f713b18d5184912ba8dd389f86aa7710db079abcb0` |
| P0-2 | **商业授权说明**（`COMMERCIAL-LICENSE` / 商业授权条款与联系方式）中英双语 | MUST | 含版权主体、商业联系 `xiaolibuluo@163.com`（沿用待确认）、「本文件不构成法律意见」声明；不含任何定价/已签约/已支付等虚构内容 |
| P0-3 | **CLA / 再许可缺口处置** | MUST | 显式标注「DCO 不授予再许可权；双许可下需 CLA 或明确再许可条款；**需权利人或法务确认**」；**不得**写成已解决的法律结论，不得擅自把 CLA 从「可选」改为「必需」 |
| P0-4 | **`LICENSING` 总览**（双许可入口）中英双语 | MUST | 一张表说清两条许可路径、SPDX 标识（only/or-later 待确认）、指向 `LICENSE`/`COMMERCIAL-LICENSE`/`FAQ`/`CONTRIBUTING` |
| P0-5 | **README 双语重构** | MUST | 补齐项目介绍 / 功能 / 安装 / 许可 / 教程入口（现状仅 CI/CD 门禁说明）；纠正 `go1.27 → 1.26.11`；删去未经核实的 `golangci-lint v2.15.0` 断言；中文与英文两份逐节对应 |
| P0-6 | **3 份用户向文档双语**：`docs/user/{quickstart,cli-reference,ces-guide}.md` | MUST | 全量双语；纠正 quickstart `Go 1.24+ → Go 1.26.11`；命令/输出逐字对齐 §6d，禁照抄旧二进制结论 |
| P0-7 | **保姆式教程 ≥8 条** | MUST | 见 §6；每条含「目标读者 / 前置条件 / 可复制命令 / 预期输出 / 失败排查」；命令与预期输出与 §6d 一致；可照做复现 |
| P0-8 | **governance 三件套双语**：`{DUAL-LICENSE,HEADERS,THIRD-PARTY}.md` | MUST | 全量双语、逐节对应；保留 6 行 HTML 注释版权头；DUAL-LICENSE 内「CLA 可选 / DCO 强制」须与 P0-3 缺口标注一致 |
| P0-9 | **夸大措辞清理 + 版本纠正** | MUST | 129 处命中逐处判定：`docs/plan/` 规划性表述保留「验收目标/设计目标」语气、不得译成既成事实；禁新增「国际超一线/世界级/超越竞品/行业领先/最先进/首创/第一」；竞品绝对分数 = 0；全仓版本统一 1.26.11 |

> P0 附带合规基线（不单列新 P0，随上述交付一并满足）：新建双语 `.md` 统一带 6 行 HTML 注释版权头（**治理一致性要求，非门禁强制**，措辞依 §6b）；所有法律不确定点标注待确认。

### 5.2 P1 —— 应该（双许可标准文件集余项 + 门面周边 + 大部头双语）

| 需求 | 语言 | 说明 / 判据 |
|---|---|---|
| `CONTRIBUTING`（双语） | SHOULD | DCO 1.1 + `git commit -s`、CLA 现状、导入依赖白名单（MIT/Apache-2.0/BSD/ISC/CC0/Unlicense/PD；拒 GPL/LGPL/AGPL 变体） |
| `CLA`（草案） | SHOULD | 先备草案并标注「是否必需待权利人/法务确认」，不擅自定性 |
| `DCO`（DCO 1.1 全文 + 用法） | SHOULD | 与 CONTRIBUTING 交叉链接 |
| `NOTICE` / `THIRD-PARTY-NOTICES` | SHOULD | 第三方清单对齐 vendor：wazero / x-sys / yaml.v3，许可逐项如实 |
| `SECURITY`（双语）/ `CODE_OF_CONDUCT` | SHOULD | 安全披露流程，不含 SLA / 虚构响应时限 |
| `TRADEMARK`（先标注） | SHOULD | **待确认是否存在**；存在则补名称/标识使用规则，不存在则不虚构 |
| 许可证兼容性说明 | SHOULD | 解释双许可语境下的依赖白名单，不与 AGPLv3 条款冲突 |
| `FAQ`（何时需商业授权，双语） | SHOULD | 覆盖 AGPLv3 §13 网络交互义务、闭源集成、SPDX、「国际开源标准」措辞区分 |
| SPDX 标识落地 | SHOULD | `LICENSING` 与文件头统一；only/or-later **待确认** |
| 冻结文件**独立译本** | SHOULD | `docs/system_design.md` → 仅新建 `system_design.en.md`（中文亦可另建译本），**原文一行不改**；两份 `.mermaid` 同理（见 §6.5） |
| `CHANGELOG` / `docs/ARCHITECTURE` / `docs/dev/contributing` / `SECURITY.md` 双语 | SHOULD | 交接件 §6c 列为门面优先 |
| 其余大部头技术文档双语 | SHOULD | `docs/plan/*`(9 份 ~3085 行)、`docs/dev/frontend-coverage/*`(6 份 ~1385 行)、`docs/spec/ces-v1.md`、`docs/dev/tla-verification.md`、`docs/procurement.md`、`docs/dev/pending-work.md`、`plugin/CI-EVIDENCE.md` |
| 工作树边界遵守 | SHOULD | 5 份 M 状态 frontend-coverage（`{csharp,javascript,kotlin,python,rust}.md`）**基于工作树当前内容**双语化，**不得 `git checkout HEAD` 回退**、不得基于 HEAD 旧版翻译（§6e） |

### 5.3 P2 —— 可以（降级处置 + 增值项）

| 需求 | 语言 | 说明 |
|---|---|---|
| `docs/reference/` 12 份（18–46 行，外部规范链接/规则清单） | COULD | **中文说明 + 英文摘要**，不全译；保留原始外部链接，不改写外部规范内容 |
| `ces/schema/dist/fields.md`（70 行，生成物） | COULD | **先确认是否纳入**；纳入则降级为摘要，避免与生成器输出冲突 |
| `docs/plan` 内部规划语气保留 | COULD | 双语化但统一标注 "Design target / 验收目标"，不升级为既成事实、不删除 |
| 中英术语对照表 glossary | COULD | 统一 CES/AIR/HELD/derive 等术语译法，支撑逐节对应 |
| 每份文档顶部 zh↔en 互链 | COULD | 语言切换锚点，便于核验 |
| 中英逐节对应 / 夸大词自动校验脚本 | COULD | 交给阶段 4 QA 复用的机械核验工具（本轮不改 `scripts/` 冻结段，独立新增或交 QA 临时执行） |

---

## 6. 全量双语范围清单（56 份 `.md`，逐一给处置策略）

### 6.1 覆盖口径与诚实声明（先读）

- 交接件 §6c **固化了 14 组分组计数（合计 56）与部分文件名**，并注明「完整逐份清单（路径/行数/标题）已由主理人生成，可向主理人索取或复跑盘点脚本」。
- 本 PRD **只使用交接件已固化的事实**：对**已具名文件逐一给处置**（§6.3），对**未具名份按所属分组给统一处置规则**（§6.2）。**不虚构任何文件名、行数或标题**。
- 因此「56 份逐一处置」= 14 组分组规则（覆盖 100% 份数，合计 56）+ 已具名文件逐份表；架构师阶段 2 以主理人盘点脚本输出为准，把未具名份对号入座即可，**不需重新全仓扫描本 PRD 之外的内容**。
- 三类处置级别定义：
  - **A 全量双语**：新建完整译本，中英逐节对应（章节标题/数量/代码块一一对应）。
  - **B 中文 + 英文摘要**：保留中文全文，新增英文摘要（不逐节全译）。
  - **C 内部规划降级**：双语化但统一保留「设计目标 / 验收目标」语气，明确「非已实测达成」，禁译成既成事实或营销语言。

### 6.2 分组处置表（14 组，合计 56 份，§6c 计数）

| # | 分组（§6c） | 份数 | 处置 | 优先级 | 要点 |
|---|---|---:|---|---|---|
| 1 | `root` | 4 | A 全量双语 | P0/P1 | README 重构为中英门面（P0-5）；CHANGELOG 双语；其余按具名逐一；统一 6 行版权头 |
| 2 | `ces/schema/dist` | 1 | B（待确认纳入） | P2 | `fields.md` 为生成物，先确认是否纳入；纳入则中文说明 + 英文摘要 |
| 3 | `docs` | 3 | A 全量双语 | P1 | 含冻结 `system_design.md`：**仅新建译本、原文不动**；ARCHITECTURE、procurement 全量双语 |
| 4 | `docs/adr` | 1 | A 全量双语 | P1 | 保留 ADR 决策语境，不改结论 |
| 5 | `docs/dev` | 9 | A 全量双语 | P0/P1 | contributing 优先；tla-verification、pending-work 等全量双语 |
| 6 | `docs/dev/frontend-coverage` | 6 | A 全量双语 | P1 | **基于工作树当前内容**；其中 5 份为上一任务 M 状态（覆盖数字已腰斩），**禁回退、禁基于 HEAD 旧版翻译**（§6e） |
| 7 | `docs/plan` | 9 | **C 内部规划降级** | P1 + 诚实约束 | ~3085 行含大量「验收目标 100%」；双语但保留目标语气，禁译成既成事实、禁营销化 |
| 8 | `docs/reference` | 12 | B 中文 + 英文摘要 | P2 | 外部规范链接/规则清单，不全译、不改写外部规范 |
| 9 | `docs/spec` | 1 | A 全量双语 | P1 | `ces-v1.md` 规格逐节对应 |
| 10 | `docs/user` | 3 | A 全量双语 | **P0** | quickstart / cli-reference / ces-guide；纠正版本；命令逐字对齐 §6d |
| 11 | `governance` | 3 | A 全量双语 | **P0** | DUAL-LICENSE / HEADERS / THIRD-PARTY；保留版权头；与 P0-3 缺口标注一致 |
| 12 | `plugin` | 1 | A 全量双语 / B | P1/P2 | `CI-EVIDENCE.md`，含门禁证据须与 §6b/§6d 一致 |
| 13 | `security` | 1 | A 全量双语 | P1 | 安全披露，无 SLA 虚构 |
| 14 | `security/policy` | 2 | A 全量双语 | P1 | 策略类逐节对应 |
| | **合计** | **56** | | | 4+1+3+1+9+6+9+12+1+3+3+1+1+2 = 56 ✓ |

### 6.3 已具名文件逐份处置（交接件明确点名者，不臆造其余）

| 文件（交接件已具名） | 行数/标记 | 处置 | 优先级 |
|---|---|---|---|
| `README.md` | 47, hdr=N, 仅 CI 门禁说明 | A：双语**重构**（补介绍/安装/许可/教程入口，纠正 go1.27） | P0 |
| `CHANGELOG.md` | 97 | A 全量双语 | P0 门面优先 / P1 |
| `SECURITY.md` | 72 | A 全量双语 | P0 门面优先 / P1 |
| `docs/user/quickstart.md` | 93/94, hdr=Y | A 全量双语；纠正 `Go 1.24+ → 1.26.11` | **P0** |
| `docs/user/cli-reference.md` | 79/80, hdr=Y | A 全量双语；命令逐字对齐 §6d | **P0** |
| `docs/user/ces-guide.md` | 99/100, hdr=Y | A 全量双语 | **P0** |
| `governance/DUAL-LICENSE.md` | 39, hdr=Y | A 全量双语；CLA/DCO 缺口标注待确认 | **P0** |
| `governance/HEADERS.md` | 45, hdr=Y | A 全量双语；§2 治理约定 vs §6b 门禁纠偏措辞精确 | **P0** |
| `governance/THIRD-PARTY.md` | 49, hdr=Y | A 全量双语 | **P0** |
| `docs/dev/contributing.md` | 57 | A 全量双语；DCO/CLA、依赖白名单 | P0 门面优先 / P1 |
| `docs/system_design.md` | 1366, **冻结** | **仅新建独立译本**（如 `system_design.en.md`），**原文一行不改** | P1 |
| `docs/class-diagram.mermaid` | **冻结** | 仅新建译本，原文不改（见 §6.5：非 `.md`，不计入 56） | P1 |
| `docs/sequence-diagram.mermaid` | **冻结** | 同上 | P1 |
| `docs/ARCHITECTURE.md` | 258, hdr=N | A 全量双语 | P1 |
| `docs/procurement.md` | 84 | A 全量双语 | P1 |
| `docs/spec/ces-v1.md` | 82 | A 全量双语 | P1 |
| `docs/dev/tla-verification.md` | 96 | A 全量双语 | P1 |
| `docs/dev/pending-work.md` | 164 | A 全量双语（保留「待办/规划」语气） | P1/P2 |
| `plugin/CI-EVIDENCE.md` | 170 | A 全量双语 / B 摘要 | P1/P2 |
| `ces/schema/dist/fields.md` | 70, 生成物 | B：先确认是否纳入 | P2 |
| `docs/plan/*` | 9 份, ~3085 行 | **C 内部规划降级** | P1 + 诚实约束 |
| `docs/dev/frontend-coverage/*` | 6 份, ~1385 行 | A 全量双语；其中 `csharp/javascript/kotlin/python/rust.md` 5 份为 M 状态，基于工作树当前内容、禁回退 | P1 |
| `docs/reference/*` | 12 份, 18–46 行 | B 中文 + 英文摘要 | P2 |
| 其余未具名份（`root` / `docs/adr` / `docs/dev` / `security` / `security/policy` 等组内） | 见 §6.2 计数 | 按所属分组规则 A 处置 | 依组 |

> 版权头现状（§6c）：56 份中约 20 份 hdr=Y（治理/用户向/安全政策等），plan / reference / frontend-coverage / ARCHITECTURE 等 hdr=N。架构师须决定双语新文档**统一带头**的落地方式；依 §6b 属治理一致性，非门禁强制。

### 6.4 双语命名与目录方案（3 个 + 推荐 + 理由；**最终由架构师定案**）

| 方案 | 结构 | 优点 | 缺点 / 风险 |
|---|---|---|---|
| **A（推荐）同目录双文件** | `README.md` 中文 + `README.en.md` 英文；`docs/user/quickstart.md` + `quickstart.en.md` | ① 现有**相对链接/图片路径不因目录搬迁失效**，改动最小；② 冻结 `system_design.md` 只需并列新建 `system_design.en.md`，**原文零改动**，不与「禁改」冲突；③ 同名同目录 →「中英逐节对应」脚本最易机械核验；④ 6 行版权头可批量统一校验；⑤ 符合多数开源仓惯例 | 同目录文件数翻倍；README 主语言（中文 or 英文）需主理人拍板 |
| B. `i18n/zh` + `i18n/en` 双树 | 镜像两棵目录树 | 语言边界最清晰，适合多语言扩展 | ① 所有相对链接需重写，56 份迁移成本高、易断链；② 冻结文件要么复制进双树（视觉/语义上像「改动」冻结件）要么遗漏；③ 与「工作树已有 ~40 份 M 状态」叠加，冲突风险高 |
| C. 单文件内中英对照段落 | 一份文件内逐段 `中文 / English` | 逐节对应天然可见、无需配对文件 | ① 对 1366 行冻结 `system_design.md` **必然要改原文 → 直接违反冻结约束**；② 文件体积翻倍、中文读者阅读体验差；③ `docs/plan` 9 份 ~3085 行维护困难 |

**推荐 A**，理由：唯一同时满足「冻结文件零改动」「链接不返工」「逐节对应可机械核验」「版权头可批量校验」的方案，风险最低。
**变体**：README 可采用「`README.md` 主语言 + `README.<lang>.md` 并列 + 顶部互链」；主语言选中文还是英文由**主理人拍板**。
**定案权**：用户明确「最终由架构师定案」→ 本 PRD 给推荐与判据，**不越权拍板**；架构师在阶段 2 输出最终目录规范并同步给阶段 3/4。

### 6.5 冻结文件特殊处置（**禁改，违反作废**）

- 交接件 §5 冻结：`docs/system_design.md`、`docs/class-diagram.mermaid`、`docs/sequence-diagram.mermaid`、`ci/e2e_thresholds.json` / `_baseline.json`、`ci/e2e_plan.json`、`scripts/lint-gate.sh`（尤其 41–135 行九段）、`ci/gate_e2e.sh`、`.golangci.yml`、若干 `internal/e2e/**` 与 `lang/frontkit/**` 的 `.go`、`go.mod`/`go.sum`/`vendor/`、`third_party/LICENSE-MANIFEST.md` 及既有 E2E 报告产物、所有 `.go` 源码。
- **精确区分**：两份 `.mermaid` **不属于 56 份 `.md` 计数**，但同属冻结文件；双语化同样**只能新建译本、不得改原文**。
- `docs/system_design.md`：**只能新建独立译本，不得改原文**（用户硬约束 + §6c）。其「`competitorAbsolute = NOT_MEASURED`、未对拍前不写竞品绝对分数」的诚实留白，译本须**原样保留语义**，不得在翻译中「顺手补上」结论。

---

## 7. 保姆式教程清单（10 条，每条 5 要素；命令与预期输出逐字对齐 §6d）

> **教程编写纪律（阶段 3 工程师必遵守）**
> ① 所有命令/输出/退出码**逐字引用 §6d 源码实时构建证据**，禁用 11-08 旧二进制或旧 `cli-reference.md` 照抄；
> ② §6d 未给实测输出的子命令（`verify`/`replay`/`check`/`plugin`/`web`）与 CES 文件逐字段内容，**以 `abate <cmd> --help` 与 `ctl/cli/testdata/ok.yaml` 真实文件为准现场补写，PRD 与教程均不臆造**；
> ③ `ok.yaml` 须写明为 `ctl/cli/testdata/ok.yaml` 的简写；
> ④ 凡宣称确定性须固定 `--seed`；
> ⑤ `derive` 保留「static AIR-level approximation, not a real go test execution」诚实措辞；
> ⑥ 每份教程同样带 6 行 HTML 注释版权头（治理一致性）；
> ⑦ 阶段 4 QA 须**重跑**每条命令核对输出。
> ⑧ 本清单中的 `git commit -s` 是**教最终贡献者**的文档内容示例，本轮工作**不执行任何 git commit**。

---

**T1 · 从零离线安装与构建（真实 Go 1.26.11 + vendor）**
- 目标读者：开源自托管用户 / 运维（US1、US4）
- 前置条件：仓库 `/workspace/abate` 完整（含 `vendor/`）；本机 Go 1.26.11；**离线无外网**（无 git remote）
- 可复制命令：
  ```bash
  go version
  cd /workspace/abate
  GOFLAGS=-mod=vendor go build -o /tmp/abate_verify .
  alias abate=/tmp/abate_verify
  abate version
  ```
- 预期输出：
  ```
  abate 0.1.0
  schema: ces.tencent.com/v1
  build: deterministic
  license: AGPLv3 or commercial
  source: https://github.com/xiaolibuluo/Abate
  watermark: 小黎的部落
  ```
- 失败排查：`go version` 非 `go1.26.11` → 版本不符，先对齐工具链（**勿**沿用 README 的 go1.27 / quickstart 的 Go 1.24+）；未设 `GOFLAGS=-mod=vendor` → 离线环境尝试联网拉依赖失败，须显式带上；`vendor/` 缺失 → 需完整仓库快照；`license:` 行是核验双许可标识的最快方式

---

**T2 · 第一个混沌实验（run ok.yaml）**
- 目标读者：开源自托管用户（US1）
- 前置条件：T1 完成
- 可复制命令：
  ```bash
  abate run ctl/cli/testdata/ok.yaml
  ```
- 预期输出（报告头）：
  ```
  schema: abate.report.v1
  apiVersion: ces.tencent.com/v1
  name: cli-e2e-ok
  seed: 12345
  status: COMPLETED
  exitCode: 0
  events: 6
  observations: 4
  runDigest: 6444a24ed2…
  health | HELD | held span 3000000000 satisfies holdFor 1000000000
  ```
- 失败排查：`exitCode` 非 0 → 参照 T6 退出码表定位；`ctl/cli/testdata/ok.yaml` 找不到 → 确认在仓库根执行、路径未拼错；`status` 非 `COMPLETED` → 检查 CES 文件是否被改动

---

**T3 · 干跑校验 CES（`--dry-run`）**
- 目标读者：七语言前端开发者 / 运维（US5、US4）
- 前置条件：T1 完成；示例 CES 文件存在
- 可复制命令：
  ```bash
  abate run --dry-run ok.yaml     # ok.yaml = ctl/cli/testdata/ok.yaml
  ```
- 预期输出：`status: VALIDATED` / `exitCode: 0` / `events: 0` / `observations: 0`
- 失败排查：出现 `events`/`observations` 非 0 → 说明未进入纯校验态，检查是否漏写 `--dry-run`；`exitCode: 3` → CES 格式/校验失败，对照 T6 的 `invalid`/`malformed` 样例

---

**T4 · 固定种子 + JSON 输出（`--seed` / `--format json`）**
- 目标读者：运维 / 集成开发者（US4、US5）
- 前置条件：T1 完成
- 可复制命令：
  ```bash
  abate run ok.yaml --seed 42 --format json
  ```
- 预期输出：标准 JSON，含 `"seed": 42` / `"status": "COMPLETED"` / `"runDigest": "32dbea3d9…"`
- 失败排查：**seed 覆盖会改变 digest** —— 若教程宣称输出确定性，必须固定 `--seed`，不得用变化 seed 复现同一 `runDigest`；`--format` 仅接受 `console|json`

---

**T5 · CLI 全命令逐个手把手（`--help` 命令树）**
- 目标读者：全体用户（US1/US4/US5）
- 前置条件：T1 完成
- 可复制命令：
  ```bash
  abate --help
  abate run --help
  abate verify --help
  abate replay --help
  abate check --help
  abate derive --help
  abate plugin --help
  abate web --help
  abate version --help
  ```
- 预期输出：`--help` 命令树列出 `run/verify/replay/check/derive/plugin/web/version`，并声明退出码四态 `0/1/2/3`；可用 flag：`--seed uint` / `--format console|json` / `--out path` / `--record path` / `--dry-run` / `--strict`。
  `derive --help` **必须**原样保留：「Killing is a static AIR-level approximation, **not a real go test execution**」。
- 失败排查：子命令 `--help` 的具体参数以**实时输出**为准逐条补写（§6d 未固化者不臆造）；**不得**把 `derive` 描述成「真实执行 go test 的变异测试」；退出码 `2` 目前仅有 `--help` 声明、§6d 未提供独立实测样例，教程如实转述声明、不虚构样例

---

**T6 · 退出码四态逐个验证（可照做核对）**
- 目标读者：运维 / 集成开发者（US4、US5）
- 前置条件：T1 完成；`ctl/cli/testdata/{ok,violated,invalid,malformed}.yaml` 存在
- 可复制命令：
  ```bash
  abate run ctl/cli/testdata/ok.yaml;         echo $?   # 0
  abate run ctl/cli/testdata/violated.yaml;   echo $?   # 1
  abate run ctl/cli/testdata/invalid.yaml;    echo $?   # 3
  abate run ctl/cli/testdata/malformed.yaml;  echo $?   # 3
  abate run;                                   echo $?   # 3（无参数）
  abate version;                               echo $?   # 0
  ```
- 预期输出：依次 `0 / 1 / 3 / 3 / 3 / 0`（与 §6d 干净实测逐条一致）
- 失败排查：任一不符 → 确认用的是源码实时构建二进制而非旧版；`violated` 为 `1`（策略违反，非崩溃）；`invalid`/`malformed`/无参数均为 `3`（输入/用法错误）

---

**T7 · CES 文件从零写（以真实模板为基线 + 干跑校验）**
- 目标读者：七语言前端开发者（US5）
- 前置条件：T1、T3 完成；`ctl/cli/testdata/ok.yaml`（36 行）可读
- 可复制命令：
  ```bash
  cat ctl/cli/testdata/ok.yaml
  cp ctl/cli/testdata/ok.yaml my-first.ces.yaml
  # 按文件逐字段修改（字段名/取值以真实 ok.yaml 为准，教程逐行讲解，不臆造）
  abate run --dry-run my-first.ces.yaml
  abate run ctl/cli/testdata/ok.yaml
  ```
- 预期输出：`--dry-run` → `status: VALIDATED` / `exitCode: 0` / `events: 0` / `observations: 0`；正式 run → T2 报告头（`schema: abate.report.v1`、`apiVersion: ces.tencent.com/v1`、`status: COMPLETED`、`health | HELD | …`）
- 失败排查：**逐字段讲解必须以 36 行真实 `ok.yaml` 为准**（交接件未固化其逐行内容，PRD 不虚构字段名/值）；改坏后对照 `invalid.yaml`（`3`）/ `malformed.yaml`（`3`）；每改一步先 `--dry-run` 再正式 run

---

**T8 · 如何贡献（DCO / CLA）**
- 目标读者：贡献者（US3）
- 前置条件：已阅读双语 `CONTRIBUTING` / `DCO` / `LICENSING`
- 可复制命令（文档示例，本轮不执行）：
  ```bash
  git commit -s -m "docs: add bilingual quickstart"   # DCO 1.1 sign-off
  ```
- 预期输出：提交含 `Signed-off-by:` 行，满足 DCO 1.1；导入依赖须在白名单内（MIT/Apache-2.0/BSD/ISC/CC0/Unlicense/PD），拒 GPL/LGPL/AGPL 变体
- 失败排查：治理现为「**CLA 可选、DCO 强制**」，但双许可下 **DCO 不授予项目方再许可（relicense）权** → 教程须**显式标注「此缺口需权利人或法务确认」**，不得下「已合规」结论，也不得擅自把 CLA 改成必需

---

**T9 · 如何判断 AGPLv3 还是商业授权（决策路径）**
- 目标读者：闭源商业集成决策者 / 法务（US2、US6）
- 前置条件：已读双语 `LICENSING` / `COMMERCIAL-LICENSE` / `FAQ`
- 可复制命令：
  ```bash
  abate version          # 输出含 license: AGPLv3 or commercial，核验双许可标识
  ```
- 预期输出：决策树（**仅作信息导航，不构成法律意见**）：
  1. 能否履行 AGPLv3 义务（含**第 13 节 Remote Network Interaction**，官方全文 line 540：网络交互修改版须向用户提供 Corresponding Source）？能 → 走 AGPLv3；
  2. 闭源集成 / 不愿公开衍生作品源码 / 不愿承担 §13 源码提供义务 → **需商业授权**，联系 `xiaolibuluo@163.com`（沿用待确认）；
  3. SPDX 采用 `only` 还是 `or-later` **待权利人确认**；
  4. 竞品许可只作**性质区分**陈述（Qt/Grafana/iText 为双许可参考；MongoDB SSPL / Elastic / HashiCorp BUSL **非 OSI 认可开源许可**），不作对标超越断言。
- 失败排查：任何「这样是否合法」的个别判断 → 教程统一答「需权利人或法务确认，本文档不构成法律意见」；不得出现「国际开源标准认证」「已获授权」等虚构表述

---

**T10 · 双许可 FAQ 导航 + 如何跑门禁**
- 目标读者：运维 / 合规 / 全体（US4、US6）
- 前置条件：T1 完成；仓库脚本可读
- 可复制命令（门禁，改动前基线为 PASSED）：
  ```bash
  bash scripts/header-check.sh
  bash scripts/license-check.sh
  bash scripts/lint-gate.sh
  go build ./...
  go test ./...
  ```
- 预期输出：`header-check.sh` = PASSED、`license-check.sh` = PASSED（wazero/x-sys → ALLOW、yaml.v3 → MIT；`LICENSE-MANIFEST` 的 GPL 字样为**预期 WARN**，不是 FAIL）；`lint-gate.sh` 九段（冻结 41–135 行）不被破坏；build/test 不回归
- 失败排查（**§6b 纠偏，措辞须精确**）：`header-check.sh` **只扫 `*.go`、不扫 `.md`** → 新建双语 `.md` 缺版权头**不会**触发 FAIL，但属治理一致性问题应补 6 行头；`license-check.sh` **不扫仓库根 `LICENSE`** → 落地 AGPLv3 全文 `LICENSE`（含 "GNU AFFERO GENERAL PUBLIC LICENSE"）**不会**误报 FAIL；FAQ 同时回答「何时需商业授权」「SPDX 是什么」「为何要 THIRD-PARTY-NOTICES」

---

## 8. 诚实性验收清单（可机械核验，交阶段 4 QA）

> 下表中的「禁用词表」仅作**扫描规则**，不是项目描述；引用它是为定义验收，不构成任何能力宣称。

| # | 验收项 | 机械核验方法 | 通过判据 |
|---|---|---|---|
| V1 | **无夸大词** | 全仓扫描禁用词表：`国际超一线 / 世界级 / 超越竞品 / 超过竞品X% / 行业领先 / 最先进 / 首创 / 第一 / 100%达成 / 已达到…水平` | 命中 = 0；`docs/plan/` 129 处规划性「验收目标 100% 通过」**允许保留但必须带「设计目标/验收目标」限定语气**，不得升级为既成事实或营销语言 |
| V2 | **版本一致** | 扫描 `go1.27` / `Go 1.24+` / `golangci-lint v2.15.0` 等 | 全仓 Go 版本唯一为 `1.26.11`；两处旧版本号已纠正；`golangci-lint v2.15.0` 未核实则删除断言或标注「未核实」，不照抄 |
| V3 | **LICENSE 与官方逐字一致** | `diff` 根 `LICENSE` vs `agpl-3.0.txt`；`sha256sum` | diff 无差异；661 行 / 34523 字节；`sha256 = 0d96a4ff68ad6d4b6f1f30f713b18d5184912ba8dd389f86aa7710db079abcb0` |
| V4 | **竞品零绝对断言** | 扫描竞品名 + 数字/百分比/绝对优势词；核对 E2E 报告 | `competitorAbsolute = NOT_MEASURED` 语义保留；任何竞品绝对分数、「超过竞品 X%」、绝对优势断言 = 0；竞品许可仅作**客观性质区分**（含「非 OSI 认可开源许可」的严谨措辞） |
| V5 | **版权头完整** | 检查所有**新建双语 `.md`** 头部 6 行 HTML 注释 | 含 `Copyright` + `xiaolibuluo/Abate` + `小黎的部落`；依 §6b 属**治理一致性要求**（header-check 只扫 `*.go`，缺失不触发门禁 FAIL，但应补齐） |
| V6 | **中英逐节对应** | 按方案 A 逐对文件比对章节标题序列 / 数量 / 代码块数 | 逐节一一对应；冻结 `system_design.md` 原文 sha256 不变、译本独立存在；两份 `.mermaid` 原文不变 |
| V7 | **命令/输出可复现** | 阶段 4 重跑 T1–T10 全部命令 | 输出与 §6d 逐字一致；`derive` 保留「not a real go test execution」；确定性宣称处均固定 `--seed`；无虚构子命令参数/CES 字段 |
| V8 | **零虚构 + 门禁不破坏** | 扫描 `认证 / 背书 / 线上地址 / 已发布版本 / 已支付 / 用户案例 / 性能数字 / SLA`；复跑门禁 | 虚构项 = 0；`header-check.sh`/`license-check.sh` 仍 PASSED（根 LICENSE 不误报）；`go build ./...`/`go test ./...`/`lint-gate.sh` 未破坏；冻结文件 sha256 一致 |

---

## 9. 待确认问题（**标注需权利人 / 主理人拍板**，本 PRD 不代断）

| # | 问题 | 现状 / 依据 | 需谁确认 |
|---|---|---|---|
| Q1 | **CLA 是否改必需？** | 治理现为「CLA 可选、DCO 强制」；双许可下 DCO **不授予再许可（relicense）权**，构成法律缺口。PRD 只标注缺口，**不改现状、不下结论** | **权利人或法务**确认；主理人拍板是否在文档中调整 |
| Q2 | **商业邮箱是否沿用 `xiaolibuluo@163.com`？** | 交接件 §2 既有真实联系；沿用待确认，勿新增未核实联系方式 | 权利人 / 主理人 |
| Q3 | **SSPL 类如何表述？** | MongoDB(SSPL)、Elastic(SSPL/Elastic License)、HashiCorp(BUSL) **均非 OSI 认可开源许可**；「国际开源标准」措辞须严谨区分，避免与 AGPLv3 混为一谈 | 主理人 / 法务确认措辞 |
| Q4 | **`only` 还是 `or-later`？** | SPDX：`AGPL-3.0-only` vs `AGPL-3.0-or-later`；官方「How to Apply」含 `or (at your option) any later version` 可作参考但**不等于本项目已选** | **权利人**确认 |
| Q5 | **商标政策是否存在？** | `TRADEMARK` 是否在标准文件集中落地、名称/标识使用规则为何，交接件未固化，不虚构 | 权利人 / 主理人 |
| Q6 | **双语命名与目录方案定案** | 本 PRD 推荐方案 A（同目录 `X.md` / `X.en.md`）；用户明确「最终由架构师定案」 | **架构师**（阶段 2）定案；README 主语言由主理人拍板 |

**附带待确认（不阻塞 P0，交架构师/主理人）**：
- 完整逐份清单（路径/行数/标题）由主理人盘点脚本提供，架构师据此把 §6.2 未具名份对号入座（§6c 说明）。
- `ces/schema/dist/fields.md`（生成物）是否纳入双语范围。
- 新建双语 `.md` 版权头的统一落地方式（治理一致性，非门禁强制，§6b）。
- `README.md` 主语言（中文 or 英文）。

---

## 10. 硬约束与工作树边界（转交阶段 2/3/4，违反作废）

1. **只做文档与许可工程**：不改任何 `.go/.sh/.json/.yml`；不改产品代码。
2. **禁改冻结文件**（§5 全表，含 `docs/system_design.md`、两份 `.mermaid`、`ci/e2e_thresholds*.json`、`ci/e2e_plan.json`、`scripts/lint-gate.sh` 41–135 行、`ci/gate_e2e.sh`、`.golangci.yml`、指定 `.go`、`go.mod/go.sum/vendor`、`third_party/LICENSE-MANIFEST.md`、既有 E2E 报告）。冻结文档只能**新建译本**。
3. **不 push / 不 commit / 不联网**：仓库无 git remote；最终本地提交由主理人统一执行；成员不得自行 `git add/commit`。
4. **工作树边界（§6e）**：当前 ~40 份 M 状态文件是**上一任务 E2E 修复遗留**，非本轮脏改动；其中 5 份 `docs/dev/frontend-coverage/*.md` 双语化须**基于工作树当前内容**，**不得 `git checkout HEAD` 回退**、不得基于 HEAD 旧版翻译；最终提交须区分「文档任务产出」与「E2E 遗留 M 改动」，避免混提。用户未授权回退或丢弃任何既有改动。
5. **不确定的法律点一律标注待确认、不编造**；竞品/许可只作客观性质陈述；「设计目标」与「已实测事实」严格分离。
6. **教程命令与输出逐字对齐 §6d**，阶段 3 复现、阶段 4 重跑核对。
7. **门禁纠偏（§6b）措辞精确**：header-check 只扫 `*.go`、license-check 不扫根 `LICENSE`；带版权头 = 治理一致性要求而非门禁强制。

---

## 11. 文档信息架构（IA）草案（替代 UI Design Draft）

> 本轮为纯文档/许可工程、无应用界面，故以「文档信息架构」替代 UI 布局草案（不硬套前端模板，属诚实处理）。

```
/workspace/abate/
├── LICENSE                      # AGPLv3 官方全文（逐字，sha256 校验）      [P0-1]
├── LICENSING.md / LICENSING.en.md        # 双许可总览 + SPDX + 路径导航      [P0-4]
├── COMMERCIAL-LICENSE.md / .en.md        # 商业授权说明 + 联系方式（待确认） [P0-2]
├── README.md / README.en.md              # 双语门面：介绍→安装→许可→教程入口 [P0-5]
├── CONTRIBUTING / CLA / DCO / NOTICE / TRADEMARK / THIRD-PARTY-NOTICES
├── SECURITY.md / CODE_OF_CONDUCT / FAQ（何时需商业授权）/ 许可证兼容性说明
├── docs/
│   ├── i18n/                             # 本 PRD 所在目录（本轮唯一新建目录）
│   │   ├── PRD-dual-license-bilingual.md
│   │   └── tutorials/                   # 保姆式教程 T1–T10（中英逐节对应）
│   ├── user/{quickstart,cli-reference,ces-guide}.md(+ .en.md)   [P0-6]
│   ├── system_design.md                 # 【冻结·原文不动】
│   ├── system_design.en.md              # 【新建译本】
│   ├── class-diagram.mermaid / sequence-diagram.mermaid  # 【冻结】
│   ├── plan/*（9）· reference/*（12）· dev/*（9）· dev/frontend-coverage/*（6）
│   ├── spec/ · adr/ · ARCHITECTURE.md · procurement.md
├── governance/{DUAL-LICENSE,HEADERS,THIRD-PARTY}.md(+ .en.md)    [P0-8]
├── security/ · security/policy/ · plugin/ · ces/schema/dist/
└── CHANGELOG.md / CHANGELOG.en.md
```

**门面章节骨架（README 双语重构）**：项目一句话介绍（基于真实能力，无夸大）→ 许可（AGPLv3 / 商业双许可 + `LICENSING` 链接）→ 环境要求（Go 1.26.11 + vendor 离线）→ 快速开始（指向 T1/T2）→ 教程索引（T1–T10）→ 贡献（`CONTRIBUTING`/DCO/CLA 缺口待确认）→ 门禁 → 中英互链。

---

## 12. Open Questions（汇总，等价 §9）

Q1 CLA 是否改必需（权利人/法务）· Q2 商业邮箱沿用（权利人/主理人）· Q3 SSPL 类表述（主理人/法务）· Q4 only / or-later（权利人）· Q5 商标政策是否存在（权利人/主理人）· Q6 双语命名与目录方案（**架构师定案**）+ README 主语言（主理人）。附带：完整逐份清单索取、`fields.md` 是否纳入、版权头统一落地方式。

---

> **交付声明**：本 PRD 为阶段 1 唯一产出，自包含、可直接转交阶段 2 架构师；不含联网调研、不含全仓扫描结果、不含任何法律结论代断、不含虚构文件名/命令/输出。所有事实来源为交接件 `HANDOFF_docs.md` 与 GNU 官方 `agpl-3.0.txt`。
