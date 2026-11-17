<!--
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
-->

# Abate 威胁模型（XCU-SEC）

> 本文档是 XCU-SEC 的首版完整威胁模型，对应 PLAN-08 §3。每条威胁必须指向
> 具体归属模块与可执行测试用例（扫描校验 100% 对齐）。测试集位于
> `security/tests/`，CI 常驻全绿。

## 威胁覆盖矩阵

| 编号 | 威胁 | 攻击面 | 缓解 | 归属模块 | 可执行用例 |
|---|---|---|---|---|---|
| T1 | 恶意插件加载 | 插件安装通道 | 强制签名验证（未签名拒绝） | `plugin/signature` | ✅ `TestThreatT1SignatureRejection` |
| T2 | 插件越权（fs/net/提权） | host function 调用 | 能力沙箱默认全拒 + 运行时门控 + 审计 | `plugin/capability` + `plugin/host` | ✅ `TestThreatT2CapabilityEscalation` |
| T3 | 插件资源耗尽 | 插件执行 | 原生 Go host 的 1 MiB 结果上限 | `plugin/host`（WAS-HOST） | ✅ `TestThreatT3ResourceExhaustion` |
| T4 | 供应链投毒 | go.mod / third_party | 源码级并入 + commit 哈希锁定 + 许可证审查 | XCU-GOV | ✅ `security/audit/scan.sh`（锁版本 + 许可证审查） |
| T5 | 签名密钥泄露 | 发布通道 | 密钥离线保管 + 轮换流程 + 重叠期 | XCU-REL | ✅ `security/policy/key-policy.md`（流程演练） |
| T6 | CES 文档注入 | 文档解析 | 深度上限 64 + 输入上限 + 模糊测试 | `ces/parse` | ✅ `TestThreatT6ParseInjection` |
| T7 | 报告 XSS | 报告输出 | html/template 自动转义 + 模板变量白名单 | `report` / `ctl/web` | ✅ `TestThreatT7ReportXSS` |
| T8 | REST 未授权访问 | API 监听 | 强制 Bearer 鉴权 + 远程监听无 token 拒绝 | `ctl/rest` | ✅ `TestThreatT8UnauthorizedAccess` |
| T9 | 令牌/密钥泄漏入日志 | 全局日志 | 脱敏中间件 + REST 日志点接入 + 扫描门禁 | `security/redact` + `ctl/rest` | ✅ `TestThreatT9SecretRedaction` |
| T10 | 档案篡改 | 档案文件 | 帧哈希链 + 全档 SHA-256 | `sim/record` | ✅ `TestThreatT10ArchiveTamper` |

## 对齐校验

- 已交付可执行攻击用例：T1 / T2 / T3 / T6 / T7 / T8 / T9 / T10，均含
  `-race` 通过；T2/T3 走 `plugin/host` 原生 Go 调用链，T9 走
  `security/redact` 纯函数与 `ctl/rest` 两处日志点双通道验证。
- 流程/文档控制：T4（供应链扫描脚本）、T5（密钥策略）已落地为可执行/文档控制。
- 威胁覆盖对齐：10/10 条威胁均有对应的可执行用例或已落地的流程/文档控制，
  `security/tests/doc.go` 的覆盖矩阵与本节一致。
- 交付边界（如实保留）：gVisor 平台级沙箱（ADR-001）与真正的 WASM 运行时
  在气隙离线环境未 vendoring，不属于本里程碑交付范围。当前 T2/T3 的缓解由
  原生 Go 实现承载（`plugin/capability` 默认全拒 + `AuditLog` 审计门控，
  `plugin/host` 能力门控 + `MaxResultBytes` 1 MiB 结果上限；
  `plugin/host/hostfunc.go` 明确以 SIM-CLK/SIM-RNG 实现 host-function 而非
  WASM 运行时）；T9 由 `security/redact` 脱敏中间件实现并接入
  `ctl/rest/middleware.go` 的 panic 与访问日志两处 `log.Printf`。上述平台级
  隔离与真 WASM 运行时仍为未交付边界，不得写成已交付。

## 设计原则

1. 纵深防御：签名链 + 能力沙箱 + 输入上限 + 输出转义 + 鉴权 + 证据防篡改，分层阻断。
2. 默认拒绝：能力沙箱默认全拒；远程监听默认拒绝无 token 启动。
3. 证据不可抵赖：录制回放帧哈希链 + 全档 SHA-256，篡改即拒。
4. 离线优先：内核零网络依赖，发布签名密钥离线保管，供应链源码级并入。

## 审计节奏

每里程碑（M1–M6）结束跑全量安全测试用例集（T1–T10 各 ≥1 攻击用例）；GA 前
完成第三方外部安全审计（报告入库）。
