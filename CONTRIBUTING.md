<!--
Abate —— 止息混沌 · 变异+混沌双结合的全量覆盖质量框架
Abate — where chaos abates. Mutation × Chaos unified total-coverage quality framework.
Copyright (C) 2026 小黎的部落（www.qingdujian.com）

本文件以 GNU Affero General Public License v3（AGPLv3）发布；商业用途须事先获得
小黎的部落书面授权，详见 COMMERCIAL.md。
This file is published under the GNU Affero General Public License v3 (AGPLv3).
Commercial use requires prior written authorization from 小黎的部落.
See COMMERCIAL.md for details.
-->

# 贡献指南 | Contributing

---

## 中文

### 1. 欢迎贡献
规范意见、内核代码、WASM 插件、变异算子、文档翻译、ChaosHub/MutantHub 模式、AI 提示词——全部欢迎。开始大型改动前请先开 Issue 或 RFC 讨论方向，避免返工。

### 2. 双轨贡献通道
| 通道 | 适用 | 流程 |
|---|---|---|
| 内核轨 | core/engine/dsl 等冻结层代码与规范 | RFC 先行，见 GOVERNANCE 第 4 节 |
| 生态轨 | WASM 插件、适配器、模式库 | 常规 PR + 插件市场自动认证 |

### 3. DCO 签署
所有提交必须带 `Signed-off-by: 姓名 <邮箱>`（`git commit -s`），声明你有权以 AGPLv3 贡献该代码。

### 4. 工程红线自查（提交前必做）
1. 一类一个文件；
2. 源文件 ≤300 行（Markdown 文档与法定全文文件除外）；
3. 触及规范行为的代码必须在注释中引用条款编号（如 `// CES-4.1.2 定点整数运算`）；
4. 中英双语文档同步修改；规范双文件条款编号集合保持一致；
5. 时间读取必须经由 `core.Clock`，禁止直接调用 `time.Now()`（零抖动底座）。

### 5. 双门禁测试证据
PR 描述必填：动机、关联 Issue/条款、**覆盖率 100% 证据**与**变异杀死率 100% 证据**（对改动包运行 `go test -cover` 与变异测试的输出）。等价变异体不设豁免开关——确属等价的，按 GOVERNANCE 第 4 节 RFC 申报入登记册。

### 6. 变异测试贡献特别规则
- 新变异算子必须归入五大算子类（M1-M5）之一；不属于任何一类的，意味着新增算子类——属内核轨变更，须 RFC；
- 算子实现必须确定性：同一源位置 + 同一 seed 产出同一变异体；
- 算子自带杀死用例：必须同时提交能杀死该算子产出变异体的测试。

### 7. AI 辅助贡献规则
允许且鼓励，但必须披露 AI 使用范围；人类贡献者对每一行负责；AI 生成测试必须真实杀死目标变异体并在 CI 复现；禁止未人工核验的批量产出（GOVERNANCE 第 5 节）。

### 8. LLM 客户端与 ChaosHub 贡献规则
- LLM 客户端必须实现 `mut.LLMClient` 接口；新增 provider 必须支持自定义 BaseURL，API Key 不得硬编码或写入日志（CES-9.5.3/9.5.4）；
- ChaosHub 子编码注册表是子编码数据唯一权威来源，适配器不得硬编码子编码列表（CES-B.1）；
- 远程同步失败时必须回退至本地种子表，不得阻断实验执行（CES-B.2）；
- MCP Server 新增工具必须遵循 JSON-RPC 2.0 错误码标准，不支持的能力必须显式报错（CES-7.3.1/7.3.3）。

### 9. PR 流程
- 标题格式：`[scope] 简述`，scope ∈ {spec, core, engine, dsl, mut, llm, hub, wasm, web, adapter, docs, ci}；
- 两名 Maintainer 评审通过后合并；CI 全绿是合并前提；
- 规范章节只接受 RFC 合并：直接修改 `spec/` 条款文本的 PR 一律拒绝，除非附带已通过的 RFC 编号。

---

## English

### 1. Contributions Welcome
Specification feedback, kernel code, WASM plugins, mutation operators, documentation translation, ChaosHub/MutantHub patterns, AI prompts — all welcome. For large changes, open an Issue or RFC first to align on direction and avoid rework.

### 2. Dual-Track Contribution Channels
| Channel | Applies to | Process |
|---|---|---|
| Kernel track | Frozen-layer code and specification (core/engine/dsl) | RFC first — see GOVERNANCE §4 |
| Ecosystem track | WASM plugins, adapters, pattern libraries | Regular PR + plugin-market automatic certification |

### 3. DCO Sign-off
Every commit must carry `Signed-off-by: Name <email>` (`git commit -s`), declaring your right to contribute the code under AGPLv3.

### 4. Red-Line Self-Check (required before submitting)
1. One type per file;
2. Source files ≤300 lines (Markdown documents and verbatim legal files excepted);
3. Code touching specification behavior must reference the clause ID in a comment (e.g., `// CES-4.1.2 fixed-point integer arithmetic`);
4. Bilingual documents updated in both languages; the two specification files keep identical clause ID sets;
5. All time reads must go through `core.Clock`; calling `time.Now()` directly is forbidden (zero-jitter foundation).

### 5. Dual-Gate Test Evidence
The PR description must include: motivation, linked Issue/clause, **100% coverage evidence** and **100% mutation-kill evidence** (output of `go test -cover` and mutation testing on the changed packages). There is no bypass switch for equivalent mutants — genuine equivalents are declared via RFC into the registry per GOVERNANCE §4.

### 6. Special Rules for Mutation Contributions
- A new mutation operator must belong to one of the five operator classes (M1-M5); if it belongs to none, it implies a new operator class — a kernel-track change requiring an RFC;
- Operator implementations must be deterministic: the same source location + the same seed produce the same mutant;
- Operators ship with kill cases: tests that kill the mutants produced by the operator must be included.

### 7. AI-Assisted Contribution Rules
Allowed and encouraged, with disclosure of AI usage scope; a human contributor remains responsible for every line; AI-generated tests must genuinely kill the target mutants and reproduce in CI; unverified bulk AI output is forbidden (GOVERNANCE §5).

### 8. LLM Client and ChaosHub Contribution Rules
- LLM clients must implement the `mut.LLMClient` interface; new providers must support custom BaseURL, and API keys must not be hardcoded or logged (CES-9.5.3/9.5.4);
- The ChaosHub sub-code registry is the sole authoritative source of sub-code data; adapters must not hardcode sub-code lists (CES-B.1);
- On remote sync failure, implementations must fall back to the local seed table and must not block experiment execution (CES-B.2);
- New MCP Server tools must follow JSON-RPC 2.0 error code standards; unsupported capabilities must explicitly error (CES-7.3.1/7.3.3).

### 9. PR Process
- Title format: `[scope] summary`, scope ∈ {spec, core, engine, dsl, mut, llm, hub, wasm, web, adapter, docs, ci};
- Merged after two Maintainer approvals; all-green CI is a precondition;
- Specification chapters only accept RFC-merged changes: PRs directly modifying clause text under `spec/` are rejected unless accompanied by an approved RFC number.

---

Copyright (C) 2026 小黎的部落（www.qingdujian.com）
