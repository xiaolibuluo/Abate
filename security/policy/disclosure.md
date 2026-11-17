<!--
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
-->

# 漏洞披露流程（XCU-SEC / T4）

> 对应 PLAN-08 XCU-SEC §4.1。本策略随仓库公开，独立安全通道接收报告。

## 1. 报告通道

- 安全邮箱：**security@qingdujian.com**（独立通道，亦可用 xiaolibuluo@163.com）。
- **禁止**通过公开 issue 报告安全漏洞，避免 0-day 暴露。

## 2. 处理时间线

| 阶段 | 时限 | 动作 |
|---|---|---|
| 确认（Acknowledgement） | **48 小时** | 确认收到并分类（严重/高/中/低） |
| 修复窗口（Fix Window） | **90 天** | 严重级完成修复并准备公告 |
| CVE 申请 | 修复确认后 | 严重级向 MITRE 申请 CVE |
| 协调公告（Advisory） | 修复 + 验证后 | 发布安全公告（含影响范围与缓解） |

## 3. 严重级定义

- **严重**：远程未授权执行、签名链绕过、档案证据伪造。
- **高**：沙箱逃逸、密钥泄露、REST 鉴权绕过。
- **中**：信息泄漏、依赖已知 CVE。
- **低**：文档/配置缺陷。

## 4. 公告与信任锚

- 公告发布于 GitHub Security Advisory + 仓库 `SECURITY.md` 更新。
- 若涉及签名密钥轮换，同步更新信任锚指纹公示。

## 5. 安全港（Safe Harbor）

善意研究者按本流程报告、不利用漏洞、不泄露未公开信息，视作白帽行为，
不予追究。我们承诺不与研究者对立，优先修复而非追责。
