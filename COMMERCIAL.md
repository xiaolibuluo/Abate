<!--
Abate —— 止息混沌 · 混沌工程开放规范（CES）与参考实现
Abate — where chaos abates. Open Chaos Experiment Specification (CES) & reference implementation.
Copyright (C) 2026 小黎的部落（www.qingdujian.com）

本文件以 GNU Affero General Public License v3（AGPLv3）发布；商业用途须事先获得
小黎的部落书面授权，详见本文件正文。
This file is published under the GNU Affero General Public License v3 (AGPLv3).
Commercial use requires prior written authorization from 小黎的部落.
See the body of this document for details.
-->

# Abate 商业授权说明 | Abate Commercial Licensing

---

## 中文

Abate 以 GNU Affero General Public License v3（AGPLv3）发布。AGPLv3 授予你运行、研究、修改和再分发的自由，且**网络交互使用同样触发 Copyleft**：以 AGPLv3 合规方式（向用户提供完整对应源码）运行服务是自由的，无需额外授权。

以下希望**免除 AGPLv3 义务**的情形，**须事先获得小黎的部落书面商业授权**：

1. 将 Abate（或其修改版）闭源集成进商业产品并对外销售或许可；
2. 以闭源 SaaS、PaaS 或托管服务形式，基于 Abate 向第三方提供混沌工程/可靠性测试服务（即不向服务用户提供 AGPLv3 要求的对应源码）；
3. 在企业内部将 Abate 用于生产环境的商业运营支撑，且不希望履行 AGPLv3 义务；
4. 对 Abate 的商标、品牌进行商业化使用；
5. 基于 ChaosHub/MutantHub 聚合数据集提供商业数据服务或行业基准报告。

### 双轨许可边界

- **内核轨**（core/engine/dsl/web/CLI 与官方 WASM 插件）：一律 AGPLv3，适用上述商业授权规则；
- **第三方 WASM 插件**：通过插件协议以独立进程/沙箱加载的社区插件，许可由插件作者自选（建议 AGPLv3 兼容），不因运行于 Abate 内核而被传染——但插件如对内核代码做静态链接或派生修改，则落入 AGPLv3 范围；
- **AI 补测飞轮产物**：由 AI 生成并经 CI 验证的测试用例，其权属归运行方项目所有，Abate 不主张权利。
- **LLM 第三方服务**：Abate 内置的 LLM 客户端（`llm` 包，OpenAI/Claude，CES-9.5.3）仅是调用第三方 API 的技术通道；对 OpenAI、Anthropic 等上游模型服务的调用费用、数据政策与使用条款，由使用者与该上游服务方自行约定，与小黎的部落无关。经 Abate 发送的源码与变异体数据请自行评估上游服务的保密承诺（CES-9.5.4 要求 API Key 不落日志，但传输内容仍受上游服务条款约束）。
- **ChaosHub 远程数据源**：Abate 内置的子编码同步器（`hub` 包，CES-B.2）默认端点指向小黎的部落运营的 ChaosHub 服务。该服务以何种条款提供（含可用性承诺与频控），以服务页面公示为准；使用自建或第三方端点（`-url` 参数）时，数据准确性由端点运营方负责。

### 授权方式

- 授权方：小黎的部落
- 官网：www.qingdujian.com
- 流程：官网联系 → 场景说明 → 书面授权协议 → 授权生效。未经授权的商业使用不享有任何权利。

### 常见问题

**Q：个人学习、研究、非商业内部测试需要授权吗？**
A：不需要。AGPLv3 已覆盖这些自由。

**Q：修改 Abate 后只在自己公司内部使用，不分发，算商业用途吗？**
A：若该使用支撑商业运营且你希望免除 AGPLv3 义务，则需要授权；纯 AGPLv3 合规的内部使用不需要额外授权。

**Q：商业授权是排他的吗？**
A：默认非排他。排他授权另行商谈。

---

## English

Abate is published under the GNU Affero General Public License v3 (AGPLv3). AGPLv3 grants you the freedom to run, study, modify, and redistribute the software, and **network interaction use triggers copyleft as well**: operating a service in an AGPLv3-compliant manner (offering complete corresponding source to its users) is free and needs no extra license.

The following scenarios that seek **exemption from AGPLv3 obligations** **require prior written commercial authorization from 小黎的部落**:

1. Integrating Abate (or a modified version) closed-source into a commercial product offered for sale or license;
2. Providing chaos engineering / reliability testing services to third parties based on Abate as closed-source SaaS, PaaS, or managed services (i.e., without offering users the corresponding source AGPLv3 requires);
3. Using Abate within an organization to support commercial production operations without fulfilling AGPLv3 obligations;
4. Commercial use of the Abate trademark or brand;
5. Providing commercial data services or industry benchmark reports based on the aggregated ChaosHub/MutantHub datasets.

### Dual-Track Licensing Boundary

- **Kernel track** (core/engine/dsl/web/CLI and official WASM plugins): strictly AGPLv3, subject to the commercial licensing rules above;
- **Third-party WASM plugins**: community plugins loaded as isolated processes/sandboxes via the plugin protocol are licensed at their authors' choice (AGPLv3-compatible recommended) and are not contaminated by running on the Abate kernel — however, static linking against or derivative modification of kernel code falls under AGPLv3;
- **AI flywheel output**: AI-generated, CI-verified test cases belong to the operating project; Abate claims no rights over them.
- **Third-party LLM services**: the built-in LLM clients (the `llm` package, OpenAI/Claude, CES-9.5.3) are merely technical channels for calling third-party APIs; call costs, data policies, and terms of use for upstream model services such as OpenAI and Anthropic are agreed between the user and those providers — 小黎的部落 is not a party. Evaluate the upstream services' confidentiality commitments before sending source code and mutant data through Abate (CES-9.5.4 requires API keys never be logged, but transmitted content remains bound by upstream terms).
- **ChaosHub remote data source**: the built-in sub-code syncer (the `hub` package, CES-B.2) defaults to the ChaosHub service operated by 小黎的部落. The terms under which that service is offered (including availability commitments and rate limits) are as published on the service page; when using a self-hosted or third-party endpoint (the `-url` flag), data accuracy is the responsibility of that endpoint's operator.

### How to Obtain a License

- Licensor: 小黎的部落
- Website: www.qingdujian.com
- Process: contact via website → describe your scenario → sign a written license agreement → license takes effect. Unauthorized commercial use confers no rights.

### FAQ

**Q: Do I need a license for personal learning, research, or non-commercial internal testing?**
A: No. AGPLv3 already covers these freedoms.

**Q: Is modified Abate used only internally, without distribution, considered commercial use?**
A: If the usage supports commercial operations and you seek exemption from AGPLv3 obligations, a license is required. Internal use fully compliant with AGPLv3 requires no additional license.

**Q: Is the commercial license exclusive?**
A: Non-exclusive by default. Exclusive licensing is negotiable.

---

Copyright (C) 2026 小黎的部落（www.qingdujian.com）
