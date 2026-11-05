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

# 治理章程 | Governance

---

## 中文

### 1. 项目定位
Abate 是变异测试与混沌测试统一于扰动代数（HPOV）的全量覆盖质量框架。本章程治理的对象分双轨：**内核轨**追求十年守恒，**生态轨**追求迭代速度——治理的核心是守住两条轨的边界。

### 2. 双轨治理边界
| 轨道 | 内容 | 变更门槛 |
|---|---|---|
| 内核轨（冻结层） | HPOV 语义、评分公式（CES 第 4 章）、五大变异算子类（M1-M5）、故障六分类（F1-F6）、TQS 聚合规则、插件协议主版本 | RFC + 两个独立实现佐证 + 30 天公示 + 委员会三分之二多数 |
| 生态轨（活跃层） | WASM 插件、语言前端、平台适配器、ChaosHub/MutantHub 模式库、AI 提示词、LLM 客户端实现、ChaosHub 子编码同步 | 两名 Maintainer 评审，插件市场自动认证 |

边界仲裁：一项变更属于哪条轨由规范委员会裁定，裁定记录公开。

### 3. 角色与权责
| 角色 | 权责 |
|---|---|
| 规范委员会 | 小黎的部落指派的初始成员组成；裁决 RFC、守护冻结条款、签发认证徽章、仲裁双轨边界 |
| Maintainer | 代码评审与合并；无权单方面修改内核轨内容 |
| Contributor | 通过 PR 与 RFC 参与；签署 DCO 后贡献被接受 |

### 4. 决策机制
- **规范变更**：RFC 流程（动机、条款草案、影响分析），硬性门槛为上表内核轨要求；条款编号永久有效，只可标记废弃，不得回收复用。
- **代码变更**：两名 Maintainer 评审通过即可合并；触及 CES 条款引用（CES-x.y.z）的代码必须链接对应条款。
- **门禁豁免**：变异杀死率/代码覆盖率 100% 门禁不设豁免开关；确属等价变异体的，以 RFC 形式申报并由委员会确认后加入等价变异体登记册（等价变异体不计入分母）。
- **本章程变更**：视同规范变更，走 RFC 流程。

### 5. AI 生成代码的贡献规则
- AI 辅助贡献允许且鼓励，但必须：①在 PR 描述中披露 AI 使用范围；②由人类贡献者对每一行负责（理解与测试义务不转移给 AI）；③AI 生成的测试用例必须能杀死目标变异体并在 CI 中复现；④禁止提交未人工核验的 AI 批量产出。

### 6. 数据飞轮治理
- ChaosHub/MutantHub 遥测一律 opt-in，默认关闭；
- 上报内容限于脱敏的模式签名（代码模式哈希、故障分类、判定结果），禁止包含源码原文、凭据、业务数据；
- 遥测数据集的可审计摘要随年度 LTS 发布公开。
- ChaosHub 子编码注册表（CES-B.1）是子编码数据唯一权威来源；远程同步失败时回退至本地种子表，不得阻断实验执行（CES-B.2）；同步完成后必须生成版本公告（CES-B.3）。
- LLM 客户端属生态轨，新增 provider 须实现 `mut.LLMClient` 接口并支持自定义端点（CES-9.5.3/9.5.4）；API Key 不得硬编码或写入日志。

### 7. 版本与兼容承诺
- 内核轨年度 LTS；v1.2.20 起语义化版本严格生效；
- v1.0 规范冻结后：任何破坏向后兼容的行为变更须升 MAJOR 版本，且不得早于发布后 24 个月；
- WASM 插件协议主版本内永远向后兼容，旧插件在新内核上必须可加载。

### 8. 利益冲突与商业边界
- 商业授权收入与治理决策隔离：委员会成员不得因商业授权压力放宽规范合规判定或门禁标准；
- 一致性认证结果只对规范负责，不受任何商业关系影响。

### 9. 行为准则
采用 Contributor Covenant v2.1；违规由委员会处置。

---

## English

### 1. Project Positioning
Abate is a total-coverage quality framework unifying mutation testing and chaos engineering on the perturbation algebra (HPOV). This charter governs two tracks: the **kernel track** pursues ten-year conservation, the **ecosystem track** pursues iteration velocity — the core of governance is guarding the boundary between them.

### 2. Dual-Track Governance Boundary
| Track | Content | Change Threshold |
|---|---|---|
| Kernel track (frozen layer) | HPOV semantics, scoring formulas (CES Chapter 4), five mutation operator classes (M1-M5), six fault classes (F1-F6), TQS aggregation rules, plugin protocol major versions | RFC + feasibility proof from two independent implementations + 30-day public comment + two-thirds committee majority |
| Ecosystem track (active layer) | WASM plugins, language frontends, platform adapters, ChaosHub/MutantHub pattern libraries, AI prompts, LLM client implementations, ChaosHub sub-code sync | Two Maintainer reviews; plugin-market automatic certification |

Boundary arbitration: the specification committee decides which track a change belongs to; rulings are public.

### 3. Roles and Responsibilities
| Role | Responsibilities |
|---|---|
| Specification Committee | Initial members appointed by 小黎的部落; adjudicates RFCs, guards freeze clauses, issues conformance badges, arbitrates the dual-track boundary |
| Maintainer | Code review and merge; no unilateral power over kernel-track content |
| Contributor | Participates via PRs and RFCs; contributions accepted after DCO sign-off |

### 4. Decision Making
- **Specification changes**: RFC process (motivation, clause draft, impact analysis) with the kernel-track gates above; clause IDs are permanent — deprecatable, never recycled.
- **Code changes**: merged after two Maintainer reviews; code referencing CES clauses (CES-x.y.z) must link the corresponding clause.
- **Gate exemptions**: the 100% mutation-kill / code-coverage gates have no bypass switch; genuinely equivalent mutants are declared via RFC and, upon committee confirmation, entered into the equivalent-mutant registry (excluded from the denominator).
- **Changes to this charter**: treated as specification changes and follow the RFC process.

### 5. Rules for AI-Generated Contributions
- AI-assisted contributions are allowed and encouraged, provided that: (1) the scope of AI use is disclosed in the PR description; (2) a human contributor remains responsible for every line (the duties of understanding and testing do not transfer to the AI); (3) AI-generated test cases must actually kill the target mutants and reproduce in CI; (4) unverified bulk AI output is forbidden.

### 6. Data Flywheel Governance
- ChaosHub/MutantHub telemetry is strictly opt-in and off by default;
- Reports are limited to de-identified pattern signatures (hashed code patterns, fault classes, verdicts) — never source code, credentials, or business data;
- An auditable summary of the telemetry dataset is published with each annual LTS.
- The ChaosHub sub-code registry (CES-B.1) is the sole authoritative source of sub-code data; on remote sync failure, implementations must fall back to the local seed table and must not block experiment execution (CES-B.2); a version announcement must be generated after synchronization (CES-B.3).
- LLM clients belong to the ecosystem track; new providers must implement the `mut.LLMClient` interface and support custom endpoints (CES-9.5.3/9.5.4); API keys must not be hardcoded or logged.

### 7. Versioning and Compatibility Commitment
- Annual LTS for the kernel track; strict semantic versioning from v1.2.20 onward;
- After the v1.0 specification freeze: any backward-incompatible behavioral change requires a MAJOR bump, no earlier than 24 months after release;
- The WASM plugin protocol is always backward compatible within a major version: old plugins must load on new kernels.

### 8. Conflicts of Interest and Commercial Boundary
- Commercial licensing revenue is isolated from governance decisions: committee members must not relax conformance judgments or gate standards under licensing pressure;
- Conformance results answer only to the specification, independent of any commercial relationship.

### 9. Code of Conduct
Contributor Covenant v2.1 applies; violations are handled by the committee.

---

Copyright (C) 2026 小黎的部落（www.qingdujian.com）
