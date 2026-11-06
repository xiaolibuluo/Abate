<!--
Abate —— 止息混沌 · 变异+混沌双结合的全量覆盖质量框架
Abate — where chaos abates. Mutation × Chaos unified total-coverage quality framework.
Copyright (C) 2026 小黎的部落（www.qingdujian.com）

本文件以 GNU Affero General Public License v3（AGPLv3）发布；商业用途须事先获得
小黎的部落书面授权，详见 COMMERCIAL.md。网络交互使用同样触发 AGPLv3 Copyleft。
This file is published under the GNU Affero General Public License v3 (AGPLv3).
Commercial use requires prior written authorization from 小黎的部落.
See COMMERCIAL.md for details. Network use triggers AGPLv3 copyleft.
-->

# Abate

> Abate —— 止息混沌。 | Abate — where chaos abates.
> 版本：v1.2.20（年度 LTS 轨） · 开源方：小黎的部落（www.qingdujian.com）
> Version: v1.2.20 (annual LTS track) · Open-sourced by 小黎的部落 (www.qingdujian.com)

---

## 中文

Abate 不是又一个混沌工具，也不是又一个变异测试工具——**Abate 是全球首个把变异测试与混沌测试统一到同一扰动代数之上的全量覆盖质量框架**。静态扰动（变异）证明你的测试有效，动态扰动（混沌）证明你的系统可靠；两者在数学上同构，Abate 用一套内核 HPOV 范式（假设 Hypothesis → 扰动 Perturbation → 观测 Observation → 判定 Verdict）驱动两者。

### 五大壁垒

| 壁垒 | 内容 | 竞品为何难追 |
|---|---|---|
| 双内核统一 | 变异与混沌统一于扰动代数 HPOV，同一引擎两种形态 | 需同时精通两个领域并打通底层，认知代差 |
| WASM 插件生态 | 变异算子/注入器/语言前端/适配器/报告器全部 WASM 化，任意语言、沙箱隔离、热加载 | 竞品改核心发版是季度级，社区 WASM 插件是小时级 |
| AI 自生长飞轮 | 存活变异体自动投喂 LLM 生成补测用例，测试集自我生长至杀死率 100%；MCP Server 内建，AI 助手原生调用 | 变异→补测→再变异闭环为全球首创 |
| 数据网络效应 | ChaosHub 故障模式库 + MutantHub 变异模式库，opt-in 遥测聚合全球经验，每个用户的使用增强所有用户 | 功能可抄，数据积累抄不走 |
| TQS 统一质量分 | Total Quality Score = 变异分 × 混沌分 × 覆盖分的几何聚合，单一数字准入准出，徽章 + GitOps 门禁嵌入流水线 | 一旦嵌入研发流程，迁移成本极高 |

### 零抖动底座（抖动/噪音≈0）

1. **虚拟时钟接口**：引擎、评分、模拟器中一切时间读取经由 `core.Clock` 接口，禁止直接调 `time.Now()`；生产注入物理时钟，测试与确定性模拟注入虚拟时钟——时间抖动在结构上被消除。
2. **seed 驱动全链路**：splitmix64 确定性 RNG（纯整数、标准库实现）驱动变异体选择、注入时序、采样顺序——同一输入永远同一输出，确定性重放为规范强制条款（CES-3.1.3）。
3. **定点整数运算**：评分与覆盖链路全部使用万分位（‱）定点整数（`core.BP`），禁止浮点——跨语言、跨平台零舍入漂移。
4. **WASM 沙箱指令计数**：变异体执行以 CPU 指令计数替代墙钟度量，性能噪音物理归零。

### 100% 全量覆盖的四层定义

| 层 | 定义 | 门禁 |
|---|---|---|
| 代码覆盖 | 语句覆盖率（go test -cover） | 100% |
| 变异杀死 | 非等价变异体杀死率（Mutation Score） | 100% |
| 混沌场景 | 服务 × 故障分类 × 依赖面覆盖矩阵 | ≥95%（CES-4.5 常量） |
| 统一判定 | TQS 几何聚合 + 三层门禁 | 任一不达标即红 |

100% 全量覆盖的数学定义：**Escaped = 0**——没有任何扰动逃逸出观测网（变异体全部被杀死，稳态破坏全部有预案）。

### 双轨发布（少维护 × 快迭代的统一）

- **内核轨**：年度 LTS，v1.2.20 起冻结 HPOV 语义、评分公式、五大变异算子类与故障六分类——守恒层十年不过时，企业敢用的信任来源。
- **生态轨**：WASM 插件、模式库、语言前端周级滚动发布，由社区与 AI 承载——速度的载体。竞品两难：追我们的稳定失去速度，追我们的速度失去信任。

### 许可（AGPLv3 + 商业双证）

- 本项目以 **GNU Affero General Public License v3（AGPLv3）** 发布；AGPLv3 下网络交互使用同样触发 Copyleft，合规运行（向服务用户提供对应源码）完全自由。
- 希望**免除 AGPLv3 义务**的商业用途（闭源集成、闭源 SaaS、免义务内部生产使用、商标商用），须事先获得小黎的部落书面商业授权。
- 授权联系：www.qingdujian.com，详见 COMMERCIAL.md。
- 第三方许可声明：wazero（Apache-2.0）、golang.org/x/sys（BSD 3-Clause），详见 THIRD-PARTY-NOTICES.md。
- 免责声明：无担保、混沌工程风险警告与支持政策（包C不包赢），详见 DISCLAIMER.md。

### 工程红线（不可协商）

1. 一类一个文件：每个类型/类独立成文件，禁止多类同文件。
2. 单文件 300 行上限：源文件不得超过 300 行（空行与注释计入）。例外：Markdown 文档与原样引用的法定全文文件（LICENSE）不受此限——文档按完整性优先，法定原文不得修改。
3. 文档中英双语：小型文档单文件双语对照；规范按语言分文件（`spec/CES-v0.1.md` 中文主版 + `spec/CES-v0.1.en.md` 英文版），CI 校验两版条款编号（CES-x.y.z）集合完全一致。
4. 代码覆盖率 100% + 变异杀死率 100%：CI 双门禁强制，禁止阈值豁免。区分：CES-4.5 的混沌场景覆盖率 95% 是规范评分常量，与本条代码红线无关。
5. 以上由 CI 静态检查强制，违例一律红。

### 仓库结构

```
abate/
├── README.md / COMMERCIAL.md / GOVERNANCE.md / CONTRIBUTING.md / SECURITY.md
├── LICENSE              # AGPLv3 官方全文（法定原文，不得修改）
├── spec/
│   ├── CES-v0.1.md      # 混沌实验规范（中文主版）
│   └── CES-v0.1.en.md   # CES specification (English)
├── core/                # 守恒内核：扰动代数、算子、BP 定点、RNG、虚拟时钟、TQS
├── engine/              # 实验状态机（CES 第 3 章）
├── dsl/                 # 实验 DSL 解析（零依赖 YAML 子集）
├── injector/            # 文件系统注入器（CES-7.1 Injector 角色）
├── adapter/             # 第三方平台适配器（CES-7.2，Chaos Mesh + Litmus 完整实现）
├── plugin/              # 插件注册与握手（CES-7.1）
├── probe/               # SLI 采样探针（CES-7.1 Probe 角色）
├── reporter/            # 报告产出（CES-7.1 Reporter 角色）
├── validator/           # 准入/准出校验（CES-7.1 Validator 角色）
├── sim/                 # 确定性模拟器
├── store/               # 实验记录存储
├── mut/                 # 变异测试生成器
├── llm/                 # LLM 客户端（OpenAI + Claude，实现 mut.LLMClient）
├── hub/                 # ChaosHub 子编码注册表、远程同步与公告
├── web/                 # Web 控制台 + MCP Server（纯标准库 net/http + embed）
├── cmd/abate/           # CLI 入口（serve / mcp / hub sync / validate / run / score / version）
├── tla/                 # TLA+ 模型与 TLC 配置（附录 A，CES-3.4）
├── examples/            # 示例实验文件
├── .githooks/           # pre-commit 红线自查
└── .github/workflows/   # CI 红线检查
```

### 核心功能

**MCP Server（AI 助手原生集成）**

```bash
abate mcp
```

通过 JSON-RPC 2.0 over stdio（协议版本 `2024-11-05`）暴露 4 个工具：`generate_mutants`（变异体生成）、`apply_mutant`（变异体应用）、`list_fault_classes`（故障分类列举）、`generate_test`（LLM 补测生成）。Claude Desktop、Cursor 等 AI 编码助手可原生调用（见 CES-7.3，示例见 examples/mcp-server.md）。

**LLM 客户端（AI 自生长飞轮动力源）**

`llm` 包提供 OpenAI（Chat Completions API）与 Anthropic Claude（Messages API）两种客户端，均实现 `mut.LLMClient` 接口，可直接注入 AI 自生长飞轮，为存活变异体生成补测用例。零外部依赖（仅标准库 net/http），支持自定义端点（兼容 Azure OpenAI / 代理网关），API Key 不落日志（见 CES-9.5.3/9.5.4）。CLI 用法：

```bash
abate llm generate -provider openai -api-key sk-... -source calc.go -mutant '{"file":"calc.go","line":1,"class":"M1","from":"+","to":"-"}'
```

**ChaosHub 子编码同步（数据网络效应）**

```bash
abate hub sync
```

从远程 ChaosHub 数据源拉取最新子编码表，对比本地种子表（附录 B 20 项）生成版本公告（added/updated/removed）。同步失败自动回退本地种子表，不阻断实验执行（见 CES-B.1/B.2/B.3，示例见 examples/hub-sync.md）。

**Web API**

- `GET /api/v1/hub/subcodes` — 暴露 ChaosHub 子编码注册表
- `GET /api/health`、`GET /api/version`、`GET/POST /api/experiments` — 原有端点

### 状态

Phase 1：内核重构与变异引擎建设进行中。18 个 Go 包全部 100% 覆盖；注入器矩阵 F1-F6 全部实现；TLA+ 模型检验通过（6 状态穷举，TypeOK + ResidueFree 不变量成立）；Chaos Mesh + Litmus + Gremlin 适配器完整实现；MCP Server CLI 集成完成；LLM 客户端（OpenAI + Claude）实现完成；ChaosHub 子编码同步与公告机制实现完成。代码级数字水印 + 运行时开机亮牌（watermark 包，AST 静态卡死零网络调用）；发布签名（SHA-256 + GPG）；CLI 一致性强制条款（CES-8.4）+ 自动化测试锁死；Debug + Release 双 8 验证通过（pass 数据完全一致）。规范 v1.0 冻结前，API 可能调整。

---

## English

Abate is neither another chaos tool nor another mutation testing tool — **Abate is the world's first total-coverage quality framework that unifies mutation testing and chaos engineering on a single perturbation algebra**. Static perturbation (mutation) proves your tests are effective; dynamic perturbation (chaos) proves your system is reliable. The two are mathematically isomorphic, and Abate drives both with one kernel paradigm, HPOV (Hypothesis → Perturbation → Observation → Verdict).

### Five Moats

| Moat | Content | Why competitors struggle to follow |
|---|---|---|
| Dual-kernel unification | Mutation and chaos unified on the HPOV perturbation algebra: one engine, two forms | Requires mastery of both disciplines and deep unification — a cognitive generational gap |
| WASM plugin ecosystem | Operators, injectors, language frontends, adapters, reporters — all WASM: any language, sandboxed, hot-loadable | Competitors ship core releases quarterly; community WASM plugins ship hourly |
| AI self-growing flywheel | Surviving mutants are automatically fed to an LLM to generate targeted tests; the test suite grows itself to a 100% kill ratio; built-in MCP Server for AI assistants | The mutate → supplement → re-mutate loop is a world first |
| Data network effects | ChaosHub fault-pattern library + MutantHub mutant-pattern library; opt-in telemetry aggregates global experience — every user's usage strengthens all users | Features can be copied; accumulated data cannot |
| TQS unified quality score | Total Quality Score = geometric aggregation of mutation × chaos × coverage; a single number for admission and exit, with badges and GitOps gates embedded in pipelines | Once embedded in the delivery pipeline, migration cost is extreme |

### Zero-Jitter Foundation (jitter/noise ≈ 0)

1. **Virtual clock interface**: every time read in the engine, scoring, and simulator goes through `core.Clock`; calling `time.Now()` directly is forbidden. Production injects a physical clock; tests and the deterministic simulator inject a virtual clock — time jitter is structurally eliminated.
2. **Seed-driven everything**: a splitmix64 deterministic RNG (pure integers, stdlib) drives mutant selection, injection timing, and sampling order — identical inputs always yield identical outputs; deterministic replay is a normative clause (CES-3.1.3).
3. **Fixed-point integer arithmetic**: the scoring and coverage paths use basis-point (‱) fixed-point integers (`core.BP`) exclusively; floating point is forbidden — zero rounding drift across languages and platforms.
4. **WASM sandbox instruction counting**: mutant execution is measured by CPU instruction count instead of wall clock, physically eliminating performance noise.

### Four Layers of 100% Total Coverage

| Layer | Definition | Gate |
|---|---|---|
| Code coverage | Statement coverage (go test -cover) | 100% |
| Mutation kill | Kill ratio over non-equivalent mutants (Mutation Score) | 100% |
| Chaos scenarios | Service × fault-class × dependency-surface coverage matrix | ≥95% (CES-4.5 constant) |
| Unified verdict | TQS geometric aggregation + three-layer gates | Any miss fails the build |

The mathematical definition of 100% total coverage: **Escaped = 0** — no perturbation escapes the observation net (all mutants killed, all steady-state violations pre-planned).

### Dual-Track Release (low maintenance × fast iteration, unified)

- **Kernel track**: annual LTS; from v1.2.20 the HPOV semantics, scoring formulas, five mutation operator classes, and six fault classes are frozen — a conserved layer that stays relevant for ten years, the trust foundation enterprises dare to build on.
- **Ecosystem track**: WASM plugins, pattern libraries, and language frontends roll weekly, carried by community and AI — the velocity carrier. Competitors face a dilemma: chase our stability and lose speed, chase our speed and lose trust.

### License (AGPLv3 + commercial dual-licensing)

- Published under the **GNU Affero General Public License v3 (AGPLv3)**; network interaction use triggers copyleft, so compliant operation (offering corresponding source to service users) is entirely free.
- Commercial use seeking **exemption from AGPLv3 obligations** (closed-source integration, closed-source SaaS, obligation-free internal production use, trademark use) requires prior written commercial authorization from 小黎的部落.
- Licensing contact: www.qingdujian.com — see COMMERCIAL.md.
- Third-party notices: wazero (Apache-2.0), golang.org/x/sys (BSD 3-Clause) — see THIRD-PARTY-NOTICES.md.
- Disclaimer: no warranty, chaos engineering risk warning, and support policy — see DISCLAIMER.md.

### Engineering Red Lines (non-negotiable)

1. One type per file: every type/class lives in its own file; multiple types per file are forbidden.
2. 300-line hard limit: no source file may exceed 300 lines (blank lines and comments included). Exceptions: Markdown documents and verbatim legal full-text files (LICENSE) are exempt — documents prioritize completeness; legal text must not be modified.
3. Bilingual documentation: small documents are bilingual within a single file; the specification is split by language (`spec/CES-v0.1.md` Chinese master + `spec/CES-v0.1.en.md` English); CI asserts identical clause ID sets (CES-x.y.z) across both versions.
4. 100% code coverage + 100% mutation kill ratio: dual CI gates, no threshold exemptions. Distinction: the 95% chaos-scenario coverage target in CES-4.5 is a specification scoring constant, unrelated to this code red line.
5. All of the above are enforced by CI checks; violations fail the build.

### Repository Structure

```
abate/
├── README.md / COMMERCIAL.md / GOVERNANCE.md / CONTRIBUTING.md / SECURITY.md
├── LICENSE              # AGPLv3 official full text (legal text, never modified)
├── spec/
│   ├── CES-v0.1.md      # Chaos Experiment Specification (Chinese master)
│   └── CES-v0.1.en.md   # CES specification (English)
├── core/                # Conserved kernel: perturbation algebra, operators, BP, RNG, virtual clock, TQS
├── engine/              # Experiment state machine (CES Chapter 3)
├── dsl/                 # Experiment DSL parsing (zero-dependency YAML subset)
├── injector/            # File-system injector (CES-7.1 Injector role)
├── adapter/             # Third-party platform adapters (CES-7.2, Chaos Mesh + Litmus full implementation)
├── plugin/              # Plugin registry and handshake (CES-7.1)
├── probe/               # SLI sampling probe (CES-7.1 Probe role)
├── reporter/            # Report output (CES-7.1 Reporter role)
├── validator/           # Pre/post-flight validation (CES-7.1 Validator role)
├── sim/                 # Deterministic simulator
├── store/               # Experiment record store
├── mut/                 # Mutation testing generator
├── llm/                 # LLM clients (OpenAI + Claude, implements mut.LLMClient)
├── hub/                 # ChaosHub sub-code registry, remote sync and announcement
├── web/                 # Web console + MCP Server (stdlib net/http + embed only)
├── cmd/abate/           # CLI entry (serve / mcp / hub sync / validate / run / score / version)
├── tla/                 # TLA+ model and TLC config (Appendix A, CES-3.4)
├── examples/            # Example experiment files
├── .githooks/           # pre-commit red-line self-check
└── .github/workflows/   # CI red-line checks
```

### Key Features

**MCP Server (native AI assistant integration)**

```bash
abate mcp
```

Exposes 4 tools via JSON-RPC 2.0 over stdio (protocol version `2024-11-05`): `generate_mutants`, `apply_mutant`, `list_fault_classes`, and `generate_test` (LLM test generation). AI coding assistants such as Claude Desktop and Cursor can invoke them natively (see CES-7.3; examples in examples/mcp-server.md).

**LLM Clients (powering the AI self-growth flywheel)**

The `llm` package provides OpenAI (Chat Completions API) and Anthropic Claude (Messages API) clients, both implementing the `mut.LLMClient` interface, ready to inject into the AI self-growth flywheel to generate tests for surviving mutants. Zero external dependencies (stdlib net/http only), custom endpoints supported (Azure OpenAI / proxy gateways compatible), API keys never logged (see CES-9.5.3/9.5.4). CLI usage:

```bash
abate llm generate -provider openai -api-key sk-... -source calc.go -mutant '{"file":"calc.go","line":1,"class":"M1","from":"+","to":"-"}'
```

**ChaosHub Sub-Code Sync (data network effects)**

```bash
abate hub sync
```

Fetches the latest sub-code table from the remote ChaosHub data source and generates a version announcement (added/updated/removed) against the local seed table (Appendix B, 20 items). On sync failure, it falls back to the local seed table without blocking experiments (see CES-B.1/B.2/B.3; examples in examples/hub-sync.md).

**Web API**

- `GET /api/v1/hub/subcodes` — exposes the ChaosHub sub-code registry
- `GET /api/health`, `GET /api/version`, `GET/POST /api/experiments` — existing endpoints

### Status

Phase 1: kernel refactoring and mutation engine construction in progress. 18 Go packages at 100% coverage; injector matrix F1-F6 fully implemented; TLA+ model checking passed (6 states exhausted, TypeOK + ResidueFree invariants hold); Chaos Mesh + Litmus + Gremlin adapters fully implemented; MCP Server CLI integration complete; LLM clients (OpenAI + Claude) implemented; ChaosHub sub-code sync and announcement mechanism implemented. Code-level digital watermark + runtime boot banner (watermark package, AST-enforced zero network calls); release signing (SHA-256 + GPG); CLI conformance mandatory clauses (CES-8.4) + automated test lockdown; Debug + Release dual-8 verification passed (pass data fully consistent). APIs may change before specification v1.0 freeze.

### Repository

- **GitHub**: https://github.com/xiaolibuluo/Abate
- **Clone**: `git clone https://github.com/xiaolibuluo/Abate.git`
- **Homepage**: https://www.qingdujian.com/abate
- **Download**: https://www.qingdujian.com/abate-dl/

---

Copyright (C) 2026 小黎的部落（www.qingdujian.com）
