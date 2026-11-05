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

# 安全策略 | Security Policy

---

## 中文

### 1. 支持版本
仅最新的年度 LTS 内核版本接收安全修复；生态轨插件由各插件维护方负责，认证徽章随安全问题撤销。

### 2. 威胁模型（Abate 特有）
Abate 同时持有『修改代码』（变异引擎）与『破坏运行时』（混沌引擎）两种高危能力，威胁模型必须显式覆盖：
- 变异引擎将变异代码写入构建产物的供应链污染；
- 混沌注入越过爆炸半径造成真实生产损伤；
- WASM 插件沙箱逃逸（越权访问宿主文件系统/网络/进程）；
- AI 生成的补测用例被投毒（引入后门测试或弱化断言）；
- 遥测数据回传泄露源码或业务信息（违反 opt-in 脱敏约定）。

### 3. 报告漏洞
- 渠道：通过 www.qingdujian.com 的安全联系入口私密报告，**请勿公开 Issue**；
- 内容：影响版本、复现步骤、潜在影响、利用前提（Abate 本身具有破坏性，请说明触发所需的真实条件）。

### 4. 响应承诺
| 阶段 | 时限 |
|---|---|
| 确认收到 | 48 小时内 |
| 初步评估与分级 | 5 个工作日内 |
| 修复与发布 | 严重 14 天，一般随下一个补丁版本 |

### 5. 分级标准
- **严重**：非授权方在目标系统执行故障注入、绕过爆炸半径或急停机制失效、WASM 沙箱逃逸、变异引擎将变异产物混入正式发布包；
- **一般**：评分/TQS 计算错误、报告伪造、插件协议越权读取、遥测脱敏失效；
- **低**：文档误导、非安全相关的崩溃。

### 6. 披露流程
修复发布后 7 天公开细节；致谢报告者（可匿名）；严重问题同步发布 GitHub Security Advisory。

### 7. 供应链
发布产物附 SBOM 与 cosign 签名；依赖 CVE 由自动化扫描跟踪，严重项 72 小时内出修复版本；WASM 插件市场上架前必须通过一致性套件与安全扫描双认证。

---

## English

### 1. Supported Versions
Only the latest annual LTS kernel receives security fixes; ecosystem-track plugins are maintained by their respective owners, and conformance badges are revoked upon security issues.

### 2. Threat Model (Abate-specific)
Abate wields two high-risk capabilities at once — modifying code (mutation engine) and disrupting runtimes (chaos engine). The threat model explicitly covers:
- Supply-chain contamination from the mutation engine writing mutated code into build artifacts;
- Chaos injection exceeding blast radius and causing real production damage;
- WASM plugin sandbox escape (unauthorized host filesystem/network/process access);
- Poisoning of AI-generated supplementary test cases (backdoor tests or weakened assertions);
- Telemetry exfiltrating source code or business data (violating the opt-in de-identification covenant).

### 3. Reporting a Vulnerability
- Channel: report privately via the security contact at www.qingdujian.com — **do not open a public Issue**;
- Include: affected versions, reproduction steps, potential impact, and exploitation preconditions (Abate is destructive by design; state the real-world conditions required).

### 4. Response Commitments
| Stage | Deadline |
|---|---|
| Acknowledgement | within 48 hours |
| Initial assessment and triage | within 5 business days |
| Fix and release | 14 days for critical; next patch release otherwise |

### 5. Severity Criteria
- **Critical**: unauthorized fault injection on target systems, blast-radius bypass or abort-mechanism failure, WASM sandbox escape, mutation engine leaking mutated artifacts into official release builds;
- **Major**: scoring/TQS computation errors, report forgery, plugin-protocol privilege-escalation reads, telemetry de-identification failure;
- **Low**: misleading documentation, non-security crashes.

### 6. Disclosure Process
Details are published 7 days after the fix release; reporters are credited (anonymously on request); critical issues receive a GitHub Security Advisory.

### 7. Supply Chain
Release artifacts ship with SBOM and cosign signatures; dependency CVEs are tracked by automated scanning with critical items fixed within 72 hours; WASM plugins must pass both the conformance suite and security scanning before marketplace listing.

---

Copyright (C) 2026 小黎的部落（www.qingdujian.com）
