# Abate 完工度终审报告 / Release Readiness Final Report

> 生成时间：2026-11-16；最近复核：2026-11-19（全语料降级口径更新 + SBOM/Release 实跑）
> 性质：发布前必检项（对应 `docs/dev/pending-work.md` §8）
> 原则：实事求是，所有结论均来自本轮独立实跑，不采信任何自报，禁止伪绿/夸大。

## 一、结论摘要

工程与门禁口径层面**已完工，具备发布工程条件**。对外发布前仍有一项须如实标注的边界（见 §四）：libclang-min 功能级为已知边界。全语料降级率已由历史超阈降至六语均 <5%（口径与阈值边界见 §四.1 更新）。版本化与 GPG 签名、GitHub 合并保护已完成配置；SBOM 与 Release 流水线已于 2026-11-19 离线实跑成功（见 §八）。

## 二、本轮独立实测门禁结果

| 门禁 | 命令 | 实测结果 |
|------|------|---------|
| 构建 | `GOFLAGS=-mod=vendor go build ./...` | PASS（rc=0） |
| 单测 | `GOFLAGS=-mod=vendor go test ./...` | PASS（FAIL=0） |
| E2E 第十段 | `bash ci/gate_e2e.sh` | `=== E2E GATE: PASSED ===`，verdict `passed=true`，failures=0 |
| 十段 lint | `bash scripts/lint-gate.sh` | `CI GATE: PASSED - 全部门禁规则通过`（rc=0） |
| TLA+ 形式化 | `bash scripts/tla/check.sh` | 3/3 PASS（真实 TLC，rc=0） |
| 三层水印 | `bash scripts/watermark-verify.sh` | PASS |
| 哈希校验 | `bash scripts/hash-check.sh verify` | PASS（11 文件一致） |

E2E 关键指标：`languages=7 records=10003 processed=1540 registered=1540 verdictPassed=true failures=0`。三项 sha 冻结未变（plan `d63b0004…`、manifest `53ef136b…`、thresholds `b967afbb…`）。

> 说明：曾有一次并行的 gate_e2e 与 lint-gate 因共用 `.artifacts/e2e/` 目录发生写冲突，导致 S4 双跑逐字节比对假阳性 FAIL；串行重跑后全部恢复 PASS。此为验证调度问题，非代码回归，已如实记录。

## 三、TLA+ 形式化验证（用户硬性要求：真实、禁造假）

工具链：`tla2tools.jar`（v1.8.0），TLC2 Version `2026.11.06.222357`，Java 21。三份规格的 `.cfg` 均真实声明 `SPECIFICATION` + `INVARIANT`（含活性 `PROPERTY Terminates`），非空跑。

| 规格 | 验证目标 | 模型规模 | 结果 |
|------|---------|---------|------|
| `PitParity.tla` | 对拍率有界/不重复/满致一致 | 44 distinct states | PASS |
| `SandboxFuel.tla` | 燃料有界/必然终止（含活性） | 61 distinct states | PASS |
| `StreamDeterminism.tla` | 随机流可复现/重放等价 | 4800 distinct states | PASS |

反证（证明非脚本伪造）：无管道退出码=0；直接调用 TLC 输出含 `Computing initial states` / `195 states generated, 61 distinct states found` / `Checking temporal properties` / `Model checking completed. No error has been found.`。

## 四、诚实边界（对外发布文案须遵守，禁止外推）

1. **降级率口径（2026-11-19 更新）**：冻结门禁的 `degrade*Pct` 按每语言首个真实 unit（单文件）计算，对应冻结阈值达标。用同一 Degradation 公式跑**全语料**（`scripts/fullcorpus-degrade.sh`，20520 条逐条真实 parse）：历史第三方语料数值为 js 49.12% / c 28.00% / py 26.89% / rs 42.33% / kt 8.21% / cs 6.98%（六语超阈，见 `docs/i18n/QA-ANTI-FRAUD-AUDIT.md` 第 7 维度，属语料自研化之前的**历史证据**，不代表当前状态）；**当前自研语料全语料**经两轮真实专用 kind 产出重测，已降至 py 3.36% / js 3.51% / c 4.80% / cs 4.65% / rs 4.91% / kt 4.48%，go=NOT_MEASURED（原始数据与 SHA-256 见 `docs/refactor/fullcorpus-degrade/`）。**关键边界**：六语全语料数值虽均 <5%，但按同一公式仍未达更严格的冻结单文件阈值，故只作事实披露，**不得据此宣称语料级/全语料「达标」或「高保真」**；冻结口径与全语料口径不得混用外推。每个 exact 升级均有前端真实产出（最小样例 + 真实语料双重守护测试）支撑，绝不靠改标签造假。
2. **libclang-min**：产物 `.so` 存在，功能级验证为已知边界，铁律口径未满足（见 `docs/dev/pending-work.md`）。对外如需声明，须标注为已知边界。
3. 反造假闭环已完成：九维审计中 capture/变异/混沌/语料/门禁判灯五维无造假；标签造假、贴标签伪 Kind、零产出 exact、临时探针均已清零；口径已透明披露。

## 五、本次新增交付

- **三层数字水印（归属：小黎的部落）**：源码层 `internal/watermark/`（规范串 `xiaolibuluo/Abate © 小黎的部落` → SHA256 `e92fa6a0a80b33498123d5545e07e5b2c8b1ce30f220e3a933644bd8118b81bc`）；构建层 `Makefile` ldflags 注入 + `cmd/watermark` 入口；完整性层 `scripts/watermark-verify.sh` + `WATERMARK.sha256`。
- **哈希校验**：`scripts/hash-check.sh`（generate/verify）+ `SHA256SUMS`（11 个关键文件）。
- **双语文档**：全仓 76 份 `.en.md`、T1–T10 保姆教程 20 份、双许可治理 22 份；命名方案 A、6 行版权头、顶部互链（译本→中文 76/76 完整；中文原文按 DESIGN B/C 类"不改"约束保持单向）。
- **GPG 签名**：RSA 4096，Key ID `5CF28F616D9888DB`，指纹 `2324 1C08 57FD 46BD 7D12 5CF2 8F61 6D98 88DB`；本地已启用 commit/tag 签名；私钥单独留存于交付区、不公开、不入库；公钥已上传 GitHub 供 Verified 验证。
- **GitHub 合并保护**：`Owner-only merge protection` Ruleset（active，默认分支）——强制 GPG 签名、1 审批 + 最近推送者审批 + 过期驳回 + 会话解决、禁 force push；仓库无外部协作者。

## 六、版本化状态（已完成）

- GPG 签名提交：`16f97d9b642f167154077bbd42d77473eb1f6d86`（签名状态 `G` / Good signature）。
- 首个签名 tag：`v1.0.0`（`git tag -v` 验证 Good signature）。
- 工作树干净（改动=0）。私钥不入库，单独留存于交付区。
- GitHub 合并保护 Ruleset 已 API 回读确认 active（强制签名 + 1 审批 + 最近推送者审批 + 过期驳回 + 会话解决 + 禁 force push）。

## 七、对外发布文案红线（发布时须遵守）

1. 仅用冻结口径数据 + 全语料数据如实并列，零夸大；libclang-min 标注为已知边界；全语料 <5% 不等于冻结口径达标（见 §四.1）。
2. 仓库尚未 push（本地版本化完成）。push 为对外发布行为，须由权利人明确授权后执行。
3. GitHub 规则集在首个 push 后于仓库设置页显示为 active（API 已回读确认生效）。

## 八、SBOM 与 Release 流水线实跑证据（2026-11-19）

对应 `docs/dev/pending-work.md` §7（此前标注「脚本就绪、无工件」）。本轮离线实跑并留存可复现证据（产物在 `dist/`，已纳入 `.gitignore`，不入库）。

| 项 | 命令 | 实测结果 |
|----|------|---------|
| SBOM 生成 | `bash scripts/sbom.sh v1.0.0 dist/` | PASS（rc=0），SPDX 2.3，1 个包（主模块 abate v1.0.0） |
| SBOM 可复现 | 双跑比对（剔除 `creationInfo.created`） | 逐字节一致 |
| Release 演练 | `bash scripts/release.sh --dry-run v1.0.0` | PASS（rc=0），5 目标交叉编译 + 校验和清单 + SBOM；跳过签名与 tag |
| 清单校验 | `sha256sum -c manifest.txt` | 5/5 OK |
| 二进制可复现 | linux/amd64 双构建 SHA256 | 一致：`763d7fd1…` |

**修复记录**：`scripts/sbom.sh` 原硬读 `third_party/LICENSE-MANIFEST.md` 而崩溃（rc=2）。项目已彻底自研化——`go.mod` 仅 `module abate`+`go 1.26.11`，零第三方依赖、无 `vendor/`、无 `third_party/` 目录—故 SBOM 的两类来源均允许为空。已修脚本对缺失清单优雅降级为空数组，如实产出「仅含主模块」的 SBOM，不改依赖事实、不造假。

**产物边界**：`--dry-run` 仅做前置校验 + 构建 + 清单 + SBOM，**未签名、未建 tag**；正式签名发布仍须权利人授权后按脚本非 dry-run 路径执行。交叉编译二进制与平台绑定，属构建产物，按 `.gitignore` 不入库。
