# Abate 完工度终审报告 / Release Readiness Final Report

> 生成时间：2026-11-16；最近复核：2026-11-19（全语料降级口径更新 + SBOM/Release 实跑 + libclang C API 功能级验证 + 正式签名发布已执行）
> 性质：发布前必检项（对应 `docs/dev/pending-work.md` §8）
> 原则：实事求是，所有结论均来自本轮独立实跑，不采信任何自报，禁止伪绿/夸大。

## 一、结论摘要

工程与门禁口径层面**已完工，具备发布工程条件**；铁律口径（真实产物 + 功能级验证）已于 2026-11-19 100% 闭环。全语料降级率已由历史超阈降至六语均 <5%（口径与阈值边界见 §四.1 更新）。版本化与 GPG 签名、GitHub 合并保护已完成配置；SBOM 与 Release 流水线已于 2026-11-19 离线实跑成功（见 §八）；libclang C API 功能级验证已真实通过（见 §九）。**正式签名发布已于 2026-11-19 在权利人明确授权后执行完成**（见 §十）：本地 4 个收尾提交已 GPG 签名推送至 `origin/main`（`783dc7ea..0ca21c37`），`v1.0.0` 签名 tag 已同步，发布资产 detached 签名 7/7 验证通过。对外发布授权边界已关闭。

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
2. **libclang-min（2026-11-19 更新）**：产物 `.so` 存在（76MB / 410 C API 符号），功能级验证已真实通过（`scripts/libclang-api-verify.sh`：自研库 CXCursor 遍历，diagnostics=0、回收 30 cursor / 1 struct / 2 函数，rc=0，ldd 断言链接自研库），铁律口径 100% 闭环（见 §九）。验证仅限「解析 + AST 访问」能力，不含 driver / CodeGen / LLVM 后端（构建已裁剪）；AIR 公开契约不含 clang 类型，功能级验证在 .so 层独立进行。
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

1. 仅用冻结口径数据 + 全语料数据如实并列，零夸大；libclang-min 可如实声明「解析 + AST 访问功能级已验证」，不得外推为含 driver/CodeGen 的完整 clang；全语料 <5% 不等于冻结口径达标（见 §四.1）。
2. 仓库已于 2026-11-19 在权利人明确授权后完成 push（`origin/main` `783dc7ea..0ca21c37`，签名 tag `v1.0.0` 同步，见 §十）；此前"尚未 push / 须授权"边界已关闭。
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

**产物边界**：`--dry-run` 仅做前置校验 + 构建 + 清单 + SBOM，**未签名、未建 tag**；正式签名发布已于 2026-11-19 在权利人授权后执行完成（独立 detached 签名，见 §十；未走脚本内 `git tag -s` 路径以避免重写已发布 tag）。交叉编译二进制与平台绑定，属构建产物，按 `.gitignore` 不入库。

## 九、libclang C API 功能级验证证据（2026-11-19）

对应 `docs/dev/pending-work.md` §1 铁律口径最后一项（此前为「功能级未闭环」）。本轮真实调用自构建库完成验证。

| 项 | 命令 / 断言 | 实测结果 |
|----|------|---------|
| 功能级验证 | `bash scripts/libclang-api-verify.sh` | PASS（rc=0），四段全过 |
| 链接来源 | `ldd` 断言 | 链接自研 `.artifacts/.../libclang-18.so.18`，非系统库 |
| 解析诊断 | `clang_getNumDiagnostics` | diagnostics=0 |
| AST 遍历 | `clang_visitChildren`（CXCursor） | total_cursors=30 / struct_Point=1 / funcs(add,main)=2 |
| 产物级 | `nm -D` + `ldd` | 76MB，410 C API 符号（与系统库同数），依赖全解析 |

**真实性保障**：验证脚本自包含——C 探针与样例源码由脚本内联生成到临时目录（`mktemp -d` + `trap` 清理），不依赖仓库外文件；clang 头文件路径从 `CMakeCache.txt` 的 `CMAKE_HOME_DIRECTORY` 动态推导（不硬编码 `/tmp`）；三重断言（链接来源 / diagnostics / 声明回收计数）任一不满足即非零退出。**AIR 架构边界不变**：`lang/rule/builtin/c/doc.go` 明确公开契约不含 clang/libclang 类型，本验证在 `.so` C API 层独立进行，不构成对 AIR 契约的改动。

**报告更新**：`.artifacts/libclang-min/build-report.txt` 已据实重写——旧文件是 2026-11-11 编译 TableGen OOM 的失败快照，与 2026-11-13 18:55 成功构建的 `.so` 矛盾；新报告如实记录产物级 + 功能级双达成与遗留边界（不含 driver/CodeGen、平台绑定、正式签名需授权）。

**铁律口径判定**：**100% 闭环**。对外仅可声明「解析 + AST 访问」能力已验证，不得外推为完整 clang 工具链。

---

## 十、正式签名发布执行记录（2026-11-19）

权利人明确授权"全权执行"后，按最小不可逆原则完成对外发布，全部经真实命令回读确认：

| 步骤 | 命令 / 断言 | 实测结果 |
|------|-------------|----------|
| 前置核验 | `git log --format='%G?' v1.0.0..HEAD` | 7/7 提交签名有效（`G`）；工作树干净；SSH `git@github.com` 认证成功 |
| 推送分支 | `git push origin main` | `783dc7ea..0ca21c37  main -> main`（fast-forward，4 个收尾提交） |
| 同步标签 | `git push origin refs/tags/v1.0.0` | `Everything up-to-date`（本地与远程一致，未移动/未重写已发布 tag） |
| 跟踪上游 | `git branch --set-upstream-to=origin/main main` | `branch 'main' set up to track 'origin/main'` |
| 资产签名 | `gpg --detach-sign --armor`（密钥 `5CF28F616D9888DB`） | 5 个交叉二进制 + `manifest.txt` + `sbom.spdx.json` 共 7 份 `.asc` |
| 签名验证 | `gpg --verify` 逐份 | 7/7 `OK` |
| 清单校验 | `sha256sum -c manifest.txt` | 5/5 `OK` |
| 远程回读 | `git ls-remote origin` | `refs/heads/main = 0ca21c37…`（与本地一致）；`refs/tags/v1.0.0^{} = a1b0796d…` |

**边界说明**：
1. `v1.0.0` 注解 tag 指向 `a1b0796d`（历史重写后重签的既定发布点），本轮 4 个收尾提交以 fast-forward 进入 `main`，**未移动/重写已发布 tag**——重写已发布 tag 属破坏性操作。
2. 发布资产及其 detached 签名位于 `dist/v1.0.0/`，受 `.gitignore` 排除，不入库；签名私钥不入库、不对外。
3. 本记录仅陈述真实执行的本地 + 远程状态，不外推为"冻结口径语料达标"或"完整 clang 工具链"等未验证结论（见 §四）。

### 十.1 后续授权执行（2026-11-19，添附）

权利人就"新建 v1.0.1 签名 tag / 创建 GitHub Release / git history 瘦身"三项再次授权后，现场核验与执行结果如下：

| 事项 | 决定 | 依据 / 实测 |
|------|------|-------------|
| **v1.0.1 签名 tag** | ✅ 已执行 | CHANGELOG 中英文新增 `[v1.0.1]`；header/style/drift 门禁全过、`go build` + cli/rest 测试通过；签名提交 `26084cf3`（`G`）→ 注解签名 tag `v1.0.1`（`728be101`，Good signature）已推送；`refs/heads/main = 26084cf3`，`refs/tags/v1.0.1^{} = 26084cf3`。遵循既有实践：tag 名与源码版本常量（`MetaVersion=0.1.0` / `versionText=1.0.0`）解耦，不拖动运行时版本，不动冻结件 |
| **GitHub Release + 资产上传** | ⬜ 未执行（凭据阻塞） | `gh` 未登录、无 `GH_TOKEN`/`GITHUB_TOKEN`、无 GitHub 连接器、`api.github.com` HTTP 403；SSH 仅能推 git 代码，建 Release/传资产必须 API token。需权利人本人完成 `gh auth login`（设备码/浏览器），Agent 不代登、不接受粘贴 token（凭据隔离） |
| **git history 瘦身（filter-repo）** | ⬜ 未执行（目标已不存在） | 历史仅 9 提交，`a1b0796d` 为重写后 root；`git rev-list --objects --all` 全量扫描确认历史中**无任何** libclang `.so`/`.a`/build 产物 blob（仅 2 个脚本共 14KB），76MB 产物早于 `a1b0796d` 清除；最大 blob 为 `ci/corpus_manifest.json`（12MB，正常源文件）。`git-filter-repo` 未安装。再执行需强推重写全部 hash、撞仓库禁-force-push 保护规则，且会作废 `v1.0.0`/`v1.0.1` tag——纯破坏零收益，故不执行 |

**对外发布状态**：**授权边界已关闭，正式签名发布完成，v1.0.1 patch tag 已发布**。剩余两项分别为凭据阻塞（Release，需权利人 OAuth）与目标已达成（history 瘦身，无需再做）；均非开发缺陷。
