<!--
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
-->

# Abate 架构总览

> 本文描述当前代码的实际结构。所有组件都能在仓库对应目录找到源码，能力声明可用
> `scripts/evidence-check.sh` 一键复现。

## 1 定位与形态

Abate 是一个**单静态二进制**，包含两个协同子系统：

1. **确定性实验控制平面**：以 CES（`ces.tencent.com/v1`）描述实验，经校验、确定性调度、
   异常注入、SIM 模拟，产出可逐字节复现的报告与 SIM-REC 档案。
2. **七语言结构化静态分析器**：把源码解析为统一 AIR 语法树，做结构模式匹配，
   输出 console / JSON / SARIF 报告。

技术基座：**纯 Go 标准库自研，零外部依赖**（`go.mod` 无 `require`、无 `vendor/`、无 `third_party/`），
构建产物为单二进制，五平台交叉编译。

## 2 分层结构

代码按职责分层，依赖方向为单向（上层调用下层，下层不反向依赖）：

| 层 | 目录 | 职责 |
|---|---|---|
| 入口 / 控制 | `main.go`、`cmd/`、`ctl/cli/`、`ctl/rest/`、`ctl/web/` | CLI 命令、REST API、浏览器控制台 |
| CES 前端 | `ces/schema/`、`ces/parse/`、`ces/validate/`、`ces/tck/` | CES 文档解析、规范化、校验（CES-VALID / CES-PARSE 规则）、一致性测试套件 |
| 语言前端 | `lang/air/`、`lang/frontkit/`、`lang/{go,py,js,c,cs,rs,kt}front/`、`lang/rule/` | 七语言 lexer/parser → AIR AST → 结构规则匹配；`check` / `derive` 逻辑在 `lang/check/`、`lang/derive/` |
| 内核 | `kernel/algebra/`、`kernel/scheduler/`、`kernel/injection/`、`kernel/mutation/`、`kernel/observe/`、`kernel/verify/`、`kernel/abi/` | 扰动代数、确定性调度、注入算子、变异算子、观测、稳态判定、ABI 数据 |
| 模拟器 | `sim/clock/`、`sim/fs/`、`sim/net/`、`sim/rand/`、`sim/record/`、`sim/vsched/` | 虚拟时钟/文件系统/网络（延迟·分区·丢包·限带宽）/确定性随机/档案/调度 |
| 安全 | `security/policy/`、`security/redact/`、`security/audit/`、`security/tests/` | 能力策略、敏感信息脱敏、安全扫描 |
| 插件 | `plugin/host/`、`plugin/signature/`、`plugin/capability/`、`plugin/sandbox/`、`plugin/{adapter,reporter,sample,wasm,hotswap}/` | Ed25519 签名校验、能力模型、沙箱配置、适配器/报告器/示例、自研 WASM 宿主、热插拔 |
| 支撑 | `internal/determinism/`、`internal/e2e/`、`internal/corpus/`、`internal/watermark/`、`internal/yamlmini/`、`ctl/rest/` | 确定性工具、E2E 框架、语料、水印、自研最小 YAML 解析、REST |

## 3 核心链路

### 3.1 `check`：源码 → AIR → findings

```
<paths> → 扩展名分派 → <lang>front (lexer → Pratt parser) → air.Node AST
        → rule matcher (结构模式) → findings → console/json/sarif
```

- **AIR**（`lang/air/`）：统一 AST 节点类型（kind / SemKey / 属性 / 子节点），带 schema 版本。
- **规则**（`lang/rule/builtin/`）：共享 AIR 模板 72 条 + 每语言专属规则（Go 116，其余 12–14），
  用编译式结构模式（`$x` 捕获、`$..` 跨层、`!P` 否定、属性约束）匹配 AST 节点，不是文本正则。
- **确定性**：同输入输出逐字节一致，`summary.fingerprint` 可用于 `verify` 比对。
- **容错**：语法错误产出 partial AIR + warnings，不做语法守门。

### 3.2 `run`：CES → 实验 → 报告

```
CES YAML → parse → normalize → validate (CES-VALID/PARSE)
        → kernel scheduler (固定种子) + injection + sim → observe → verify (稳态)
        → report + 可选 SIM-REC archive
```

- **扰动代数**（`kernel/algebra/`）：结合律、单位元、对偶补偿等代数律。
- **确定性**：固定种子 + 虚拟时钟 + 全序事件，报告与档案逐字节可复现。
- **SIM-REC 档案**（`sim/record/`）：hash 链 + 全序序列 + 虚拟时间单调，`replay` 可校验。

### 3.3 `derive`：规则 → 变异算子

从内置 Go 规则派生 M 族变异算子，注册到 `kernel/mutation/`，跑一次确定性变异实验。
killing 为**静态 AIR 级近似**，不是真实 go test 执行。

## 4 插件与安全

- **签名**：`plugin/signature/` 用 Ed25519 校验 manifest 公钥，未签名/篡改的插件拒绝加载。
- **能力沙箱**：`plugin/capability/` + `plugin/sandbox/` 建模 runsc 配置语义（能力开关、
  rootfs 只读、时钟虚拟化），默认全拒、纯声明式；真实隔离由宿主负责。
- **WASM 宿主**：`plugin/wasm/` 为**自研** WAT lexer/parser/encoder 与运行时，零第三方依赖。
- **热插拔**：`plugin/hotswap/` 原子替换实例，不中断在途调用。

## 5 确定性体系

确定性是贯穿全仓的核心约束，配套可复现门禁：

| 机制 | 实现 / 脚本 |
|---|---|
| 固定种子 / 虚拟时钟 | `kernel/scheduler/`、`sim/clock/`、`sim/rand/`（SplitMix64） |
| 逐字节报告 | `abate verify`、`summary.fingerprint`、run digest / stream digest |
| 档案完整性 | `sim/record/` hash 链、`abate replay` |
| 确定性门禁 | `scripts/determinism-check.sh`（核心包同种子双跑比对） |
| 规格一致性 | `scripts/drift-check.sh` |
| 形式化验证 | `scripts/tla/`（TLA+：PitParity / SandboxFuel / StreamDeterminism） |
| 完整性水印 | `internal/watermark/`、`scripts/hash-check.sh`、`scripts/watermark-verify.sh` |

## 6 质量门禁

`scripts/lint-gate.sh` 串联：版权头 → 风格 → 许可证 → gofmt → go vet → golangci-lint →
魔法常量 → `go test -race` → 覆盖率 → 构建。零容忍项包括 TODO/FIXME（godox）、
未格式化、忽略错误、硬编码密钥、复杂度超标等。配套：`coverage-gate.sh`、`vuln-scan.sh`、
`m5-reference-scan.sh`、`audit-zero-third-party.sh`、`fullcorpus-degrade.sh`。

## 7 边界与诚实口径

- `check` 是结构化 linter，非编译器级语义分析（详见 [能力与边界](user/capabilities-and-limits.md)）。
- 全语料构造覆盖率六语实测见 `docs/dev/frontend-coverage/` 各语言报告（C 66.4% / C# 52.6% / JS 89.5% / Kotlin 47.0% / Python 91.8% / Rust 77.9%），Go `NOT_MEASURED`；E2E“达标”仅限冻结单文件口径，不构成“高保真”结论。
- 竞品绝对分数 `NOT_MEASURED`；derive killing 为静态近似；社区验证尚为空白。

## 8 导览

- CLI 与退出码：[CLI 参考](user/cli-reference.md)
- CES 规范：[CES v1 规范](spec/ces-v1.md)
- 开发约定：[贡献指南](dev/contributing.md) · [模块开发](dev/module-dev-guide.md) · [插件开发](dev/plugin-dev-guide.md) · [沙箱](dev/sandbox.md)
