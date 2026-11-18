<!--
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
-->

# Changelog

All notable changes to Abate are documented here. The format follows a
milestone-based record; released engine versions use `v<年轮>.<minor>.<patch>`
per the XCU-REL version policy (kernel year-ring + hourly ecosystem cadence).

Version-number consistency is enforced by CI: the `ldflags` version, this
CHANGELOG entry and the OpenAPI `info.version` must agree (drift → FAIL).

## [v1.2.28]

Patch release — 把 v1.2.27 之后源码侧的易用性修订随二进制一起带入（2026-11-19）。

- **带入 check 帮助文案修订**：v1.2.27 发布后源码已把 `check --help` 的退出码说明从共用的
  "稳态违反"四态语义改为 check 专用（0 干净 / 1 含 warning 的发现 / 2 内部错误 / 3 用法错误，
  `--severity error` 可过滤）；上一版二进制未带此文案，本版本重新构建 5 目标后**下载即见新文案**。
  报告输出、规则集、AIR 与确定性指纹均不受影响（帮助文本非报告产物）。
- **带入易用性文档修订**：README 快速开始七语言 check 复制即用示例、`GOFLAGS=-mod=vendor`→
  `GOPROXY=off` 纠正、T01 版本示例 `0.1.0`→`1.2.27` 等纯文档修订已在此前源码提交落地，本版本二进制与之配套发布。
- **版本升级**：`abate version` 输出 `abate 1.2.28`，`MetaVersion` 同步 `1.2.28`；用户文档
  （quickstart / cli-reference / multilang-check-guide / T01 双语）版本示例与 README 中英徽章同步。
  水印 `VersionLabel` / 三份 SHA 锚点保持不变；`SHA256SUMS` / `WATERMARK.sha256` 因 README 徽章变更按既有实践重生成。
- **重新发布**：5 二进制 + `manifest.txt` + SPDX SBOM + 7 份 detached GPG 签名，
  以独立签名 tag `v1.2.28` 发布为 Latest；v1.2.27 / v1.2.26 / v1.0.1 / v1.0.0 tag 与 Release 保留不改写。
  细节见发布就绪文档 §十六。
- **可信度质疑核查与诚信修复（发布后文档补充，不动二进制）**：对"七语言是正则匹配"与"Go 版本口径矛盾"
  两项外部质疑做实测核查。前者证伪——`lang/frontkit/parser.go` 为 Pratt 真解析器、规则在 AST 上做结构化
  匹配，实测嵌套结构产出多类规则命中；但承认**容错解析、非编译器级语义分析**的边界。后者成立——
  `go.mod`/本机均为 `1.26.11`，README/T01 原"历史值均为旧值"表述自曝矛盾，已改中性单一口径。
  新增 `docs/user/capabilities-and-limits.md` / `.en.md` 写清能力边界与复现实测；README 首屏加边界说明；
  GitHub description 改为 "structural linter ... not a compiler-grade analyzer"。细节见 §十七。
- **第二轮质疑处置（证据链/量化/复现，文档与测试，不动二进制）**：新增 `scripts/evidence-check.sh`
  （现场构建 + 7 项带断言复现，全 PASS 才退出 0）与 `scripts/language-coverage.sh`（实测七语言规则
  数与命中量化表 + `docs/user/language-coverage.md`/.en.md）；能力文档新增"社区验证状态"章节，
  如实说明 0 star/0 fork 且拒绝造假。全量 `go test ./...` 修复一处上轮引入的真实回归
  （`TestCheckHelpContent` 旧文案断言），并新增容错解析与真 AST 多类命中两个回归测试，最终 0 FAIL。
  细节见 §十八。

## [v1.2.27]

Patch release — 让已发布二进制真正包含七语言 check 能力（2026-11-19）。

- **修复发布产物与代码时序错位**：v1.2.26 的七语言接线源码/文档于 02:18 提交，
  但发布二进制早在 01:13 构建、01:14 上传，导致下载到的 5 个二进制仍是仅支持 Go
  的旧版（`check` 帮助为 "Run the built-in Go rule set"）。本版本从当前源码重新
  构建 5 目标二进制，`abate check` 实测按扩展名分派 7 语言（`.go/.py/.js/.c/.cs/.rs/.kt`），
  `--languages` 筛选、JSON/SARIF 输出、混合目录一次检查均验证通过。
- **版本升级**：`abate version` 输出 `abate 1.2.27`，`MetaVersion` 同步 `1.2.27`；
  用户文档（quickstart / cli-reference / multilang-check-guide 中英双语）版本示例同步。
  水印 `VersionLabel` / 三份 SHA 锚点保持不变；`SHA256SUMS` / `WATERMARK.sha256`
  因 README 首屏增强按既有实践重生成并同步，门禁全过。
- **重新发布**：5 二进制 + `manifest.txt` + SPDX SBOM + 7 份 detached GPG 签名，
  以独立签名 tag `v1.2.27` 发布为 Latest；v1.2.26 / v1.0.1 / v1.0.0 tag 与 Release 保留不改写。
- **仓库曝光度优化**：description 前置"七语言静态分析 / CLI"搜索意图；topics 由 10 个扩至 20 个
  （覆盖 python/javascript/golang/rust/kotlin/c-sharp/c 七语言 + static-analysis/code-analysis/
  linter/code-quality/ast/sarif/cli/developer-tools + 混沌工程/变异测试长尾词，按 API 实测仓库数选词）；
  README 中英首屏增版本/许可/七语言/Go 徽章与价值主张。细节见发布就绪文档 §十四。
- **发布后易用性修订（源码/文档，未随 v1.2.27 二进制）**：README 快速开始补七语言 check
  复制即用示例并前置退出码提示；README/T01 构建命令由 `GOFLAGS=-mod=vendor` 统一为与
  零依赖事实一致的 `GOPROXY=off`，T01 版本示例 `0.1.0`→`1.2.27`；`check --help` 退出码
  说明由共用的稳态语义改为 check 专用（明确 warning 也返回 1、`--severity error` 过滤）。
  这些为纯展示/文档改动，报告输出与确定性不变；`SHA256SUMS`/`WATERMARK.sha256` 同步重生成。
  注：已发布 v1.2.27 二进制的 `check --help` 仍为旧四态文案，将于下个版本带入。见 §十五。

## [v1.2.26]

Release version bump with rebuilt binaries and public repository publication.

- **版本升级**：面向用户的运行时版本常量统一升级为 `1.2.26`——
  `abate version` 输出 `abate 1.2.26`，健康与 OpenAPI 元数据 `MetaVersion`
  同步为 `1.2.26`。冻结件（水印 `VersionLabel` / 三份 SHA 锚点）保持不变，
  遵循既有冻结口径，不因版本号变动破坏哈希门禁。
- **二进制重新构建与发布**：5 目标交叉编译二进制（linux/darwin/windows ×
  amd64/arm64）+ `manifest.txt` + SPDX SBOM 重新构建，7 份 detached GPG
  签名 7/7 验证通过，`sha256sum -c` 清单校验通过，经 GitHub Release 上传。
- **仓库公开**：仓库可见性由私有改为 public，同步完善描述 / topics /
  主页等仓库设置；合并保护规则集保持生效。
- 本版本以独立签名 tag `v1.2.26` 标记；此前 v1.0.0 / v1.0.1 tag 保留不改写。
- **七语言 check 接线**：`abate check` 由仅支持 Go 扩展为按文件扩展名自动分派到
  7 个语言前端（`.go/.py/.js/.c/.cs/.rs/.kt`，含 `.jsx/.mjs/.cjs/.h/.kts`），
  运行对应语言自研规则集，输出统一的 console/JSON/SARIF 报告。新增 `--languages`
  筛选；新增通用 `check.MultiChecker`（语言无关的分派+聚合，复用既有确定性管线），
  7 语言分派/混合目录/确定性/筛选/扩展名路由守护测试齐全。保姆式教程见
  `docs/user/multilang-check-guide.md`。冻结件不动。

## [v1.0.1]

Patch release — GA 后真实质量巩固与正式签名发布收官（2026-11-19）。

- **全语言降级率收官**：自研语料全语料 20520 条逐条真实 parse，六语均降到
  5% 以下（py 3.36% / js 3.51% / c 4.80% / cs 4.65% / rs 4.91% / kt 4.48%，
  go `NOT_MEASURED`）。每个 exact 升级均有前端真实专用 kind 产出（最小样例 +
  真实语料双重守护），剩余 degraded 如实保留，不改标签。冻结单文件口径不变。
- **质量维度巩固**：为 js/c/cs/rs/kt 五前端补齐 parse 基准（确定性源生成 +
  真实 `Parse` + `NodeCount` 非零护栏）；cfront 补 Go 原生 fuzz 并修复一例
  经 fuzz 发现的真实 panic（未闭合 `template <` 负 cap 崩溃），统一加区间倒置
  防御并留回归锁。全量 `go test ./...`（73 包）通过、`golangci-lint` 0 issues。
- **libclang C API 功能级验证闭环**：`scripts/libclang-api-verify.sh` 真实调用
  自研 `libclang-18.so.18`（ldd 断言链接自研库），diagnostics=0、CXCursor 遍历
  回收 30 cursor / struct Point 1 / 函数 add+main 2，rc=0。铁律口径 100% 闭环。
- **SBOM / Release 流水线实跑**：修复 `sbom.sh` 对缺失清单的崩溃，SPDX 2.3
  如实产出；`release.sh --dry-run v1.0.0` 五目标交叉编译 + 清单 5/5 OK +
  双构建 SHA256 一致。
- **正式签名发布**：权利人授权后，7 个 GPG 签名提交推送至 `origin/main`，
  签名 tag `v1.0.0` 同步，7 份发布资产 detached 签名 7/7 验证通过；
  本版本 `v1.0.1` 以独立签名 tag 标记上述收官。详见
  `docs/i18n/RELEASE-READINESS-FINAL.md` §十。

## [v1.0.0]

GA release. Closes the M1–M6 program.

- **内核（L0）**：扰动代数四运算与代数律校验、确定性调度器、观测管道、
  验证器、变异核（PIT 算法移植）、注入核（chaosblade 源码二次并入）。
- **CES（L2）**：规范模式 / 解析 / 校验 / TCK 全量就位；TCK 164 用例
  （parse 68 / validate 46 / roundtrip 50），含示例第三方合规实现。
- **模拟器（L3）**：虚拟时钟 / 确定性 PRNG / 虚拟网络 / 虚拟文件系统 /
  虚拟调度 / 录制回放（帧哈希链 + 全档 SHA-256）。
- **语言前端（L4）**：AIR + Go/Python/JS/C/C#/Rust/Kotlin 七语言前端 +
  规则引擎 + 自研检查器（SARIF）+ 算子派生。
- **控制面（L7）**：CLI / REST / Web 控制台；REST 强制 Bearer 鉴权 +
  远程监听无 token 启动拒绝。
- **安全（XCU-SEC）**：威胁模型 T1–T10 全建立；T1/T6/T7/T8/T10 可执行
  攻击用例集（含 -race）入 CI；签名链 / 报告转义 / 档案防篡改已验证。
- **安全增量（T2/T3/T9 落地）**：新增 `security/redact` 确定性脱敏中间件并
  接入 `ctl/rest/middleware.go` 的 panic 与访问日志两处日志点；`security/tests`
  新增 `TestThreatT2CapabilityEscalation`、`TestThreatT3ResourceExhaustion`、
  `TestThreatT9SecretRedaction` 三条可执行攻击用例；威胁模型对齐校验更新为
  10/10。gVisor 平台级沙箱（`kernel/abi/`）与真 WASM 运行时
  （`plugin/wasm/` + wazero）已于 GA 后增量交付，详见下方「GA 后增量收官」。
  版本号未变更，单一事实来源保持一致。
- **治理（XCU-GOV）**：版权头门禁 + 水印 + 双授权 + 第三方源码并入登记。
- **发布（XCU-REL）**：`release.sh` 可复现构建矩阵 + 校验和清单（水印）+
  SBOM（SPDX 2.3）+ GPG 签名演练。

### GA 后增量收官（2026-11-13 ~ 2026-11-14）

以下为 v1.0.0 之后的增量交付与治理收口，**版本号不变**，单一事实来源保持一致。

- **真 WASM 运行时**：`plugin/wasm/` 32 个 .go（WAT 词法 / 解析 / 编码、模块加载、
  fuel/gas 限额、确定性时钟、宿主函数绑定），wazero v1.12.0 完整 vendoring
  （`go.mod` / `vendor/modules.txt` / `third_party/_sources/wazero/`）。
  修正此前「真 WASM 运行时未交付」的边界表述。
- **gVisor ABI 裁剪移植**：`kernel/abi/` syscall 表（amd64 352 / arm64 295）、
  errno、hostarch 原语，`syscall_test.go` 运行时直读上游逐项对拍；详见
  `docs/dev/abi-port-parity.md`。修正此前「gVisor 平台级沙箱未交付」表述。
- **lint 全量清零**：7 个高复杂度生产函数 + 6 个超限测试函数拆为职责单一的
  助手函数，消除 8 个连带命中，按 ≤300 行拆出 9 个伴生源文件；
  `golangci-lint` **0 issues**（revive enable-all-rules 全开）。
- **安全升级**：Go 1.26.10 → **1.26.11**，修复 GO-2026-6859
  （`net/http` MaxBytesReader header race）；`govulncheck` 实跑
  `No vulnerabilities found`。
- **规则精度提升**：`m5-reference-scan.sh` ③ 段改为以 `go.mod` 为权威的
  vendor/require 双向一致性校验；`license-check.sh` 增加 BSD 标准法律文本特征短语。
  均未放宽验收、未引入白名单或 skip。
- **门禁全绿**：九段 CI 门禁 `PASSED`；TLA+ 3 spec；determinism / drift /
  pit-parity(100%) / m5-reference-scan / `go mod verify` 全部通过。
- **仓库治理**：1013 个 `libclang-min/build/` CMake/Ninja 构建产物移出版本控制
  并纳入 `.gitignore`（工作区文件保留）；`docs/dev/pending-work.md` 按实况重写
  为完工状态归档。
- **全语言降级率收官（2026-11-18）**：自研语料全语料 20520 条逐条真实 parse，
  py 3.36% / js 3.51% / c 4.80% / cs 4.65% / rs 4.91% / kt 4.48%，
  go 无嵌入语法映射保持 `NOT_MEASURED`；六语言全部降到 5% 以下。
  每个 exact 映射均由前端真实专用 kind 产出（最小样例 + 真实语料双重守护），
  剩余 degraded 为共享 kind 连带或冻结语料零产出构造，如实保留，不改标签。
- **质量维度巩固（2026-11-18）**：
  - P8 `frontend parse throughput` 预算 owner 为 `LNG-*`，此前仅 go/py 前端有
    基准；为 js/c/cs/rs/kt 五个前端补齐 parse 性能基准（`*_bench_test.go`），
    确定性源生成 + 真实 `Parse` + `NodeCount` 非零护栏，报告 MB/s 与 nodes/op，
    无性能硬断言。实测 js 1.23 / c 1.19 / cs 0.73 / rs 2.85 / kt 0.80 MB/s。
  - cfront 补齐 Go 原生 fuzz（`cfront_fuzz_test.go`），补足固定 chaos 语料覆盖
    不到的输入空间。**fuzz 发现并修复一例真实 panic**：未闭合 `template <` 使
    `cAngleRange` 返回倒置区间，`cLeafNodes` 以负 cap 调用 `make` 崩溃；在
    cfront 的 `cLeafNodes`/`cExprReducer.reduce` 与 csfront/ktfront 的
    `leafNodes` 统一加区间倒置防御，并留 `TestCFuzzUnclosedTemplateRegression`
    回归锁。20s / 20 万次变异零崩溃。
  - 全量 `go test ./...`（73 包）通过、`golangci-lint` 0 issues、lint-gate 十段全绿。

## [M6] GA 收尾（v1.0.0 之前）

- CES-TCK 发布与示例第三方实现合规流程走通。
- 安全审计门禁（签名链 / 沙箱逃逸 / 依赖漏洞扫描）全绿。
- 全量文档（用户 / 开发 / 规范 / ADR）就位。
- 发布工程演练（tag + 资产 + 校验和 + SBOM）。

## [M5] 七语言 + 控制面

- LNG-PY/JS/C/RS/CS/KT 全部落地；CTL-REST 上线；HTML/JSON/JUnit/OTel 报告。

## [M4] WASM 运行时 + 生态最小闭环

- WAS-HOST/CAP/HOT/SIG 四模块；插件签名链 + 能力沙箱；SARIF 报告校验。

## [M3] Go 语言前端闭环（吃狗粮）

- LNG-AIR/GO/RULE/CHK/MUT 全量；Go 规则集 ≥100 条；仓库自检 0 违规。

## [M2] 模拟器完备 + 双核并入

- SIM-NET/FS/SCHED/REC + KRN-MUT/INJ；PIT/chaosblade 源码级并入。

## [M1] 内核最小闭环

- KRN-ALG/SCH/OBS/VER + CES-SCHEMA/PARSE/VALID + SIM-CLK/RNG + CTL-CLI。
- 同种子运行 100 次报告逐字节一致；扰动代数律全覆盖。
