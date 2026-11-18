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
