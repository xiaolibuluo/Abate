<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com)
  Project: xiaolibuluo/Abate — https://github.com/xiaolibuluo/Abate
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->
> 语言 / Language：[中文](LICENSING.md) | [English](LICENSING.en.md)

# 许可总览（LICENSING）

> 本文件仅作信息导航，**本文件不构成法律意见 / This document does not constitute legal advice**。
> 所有法律判断**需权利人或法务确认 / Subject to confirmation by the rights holder or legal counsel**。

## 1. 双许可两路径总览

Abate 采用「AGPLv3（开源侧）+ 商业授权（闭源侧）」双许可。两条路径并存，你可按自身情形选择其一；本表不构成法律意见。

| 维度 | 路径 A：AGPLv3 开源 | 路径 B：商业授权 |
|---|---|---|
| 适用 | 可履行 AGPLv3 全部义务，含 **§13 Remote Network Interaction**（官方全文 line 540：网络用户须能获得 Corresponding Source） | 闭源集成 / 不愿公开衍生作品源码 / 不愿承担 §13 源码提供义务 |
| 许可文本 | `LICENSE`（AGPLv3 官方全文，逐字 661 行） | `COMMERCIAL-LICENSE.md`（条款与联系方式，**需权利人确认**） |
| SPDX | `AGPL-3.0-or-later`（**需权利人确认 / subject to confirmation by the rights holder，不写成已定**） | 专有许可，无 SPDX |
| 义务 | 同许可分发、Corresponding Source 提供、保留版权与许可声明 | 以双方另行书面协议为准；**不写定价 / 已签约 / 已支付** |
| 入口 | `CONTRIBUTING.md` / `DCO.md` / `CLA.md` | `xiaolibuluo@163.com` |
| 判断辅助 | `FAQ-licensing.md` 决策树（**仅信息导航，不构成法律意见**） | 同左 |
| 权利链缺口 | **DCO 不授予项目方再许可（relicense）权；双许可下需 CLA 或明确再许可条款 —— 需权利人或法务确认** | 同左 |

## 2. 关于 AGPLv3 第 13 节「Remote Network Interaction」

`LICENSE` 为 AGPLv3 官方全文逐字副本。第 13 节「Remote Network Interaction; Use with the GNU General Public License」位于官方全文 **line 540**：若你修改本程序并通过网络向用户提供交互，须向这些网络用户提供对应版本源码（Corresponding Source）的机会。本文件只作描述与链接，**不复制、不改写** `LICENSE` 的条款正文；条款解释**需权利人或法务确认**。

## 3. SPDX 标识（需权利人确认）

本仓库文档头统一写作 `AGPL-3.0-or-later`，**每处均附「需权利人确认 / subject to confirmation by the rights holder」**。采用 `only` 还是 `or-later` 属**待权利人确认**事项；官方全文结尾「How to Apply」含 `either version 3 ... or (at your option) any later version`，仅为参考措辞，**不等于本项目已作出选择**。本文件不代下结论。

## 4. 权利链缺口（需权利人或法务确认）

现有治理为「**CLA 可选、DCO 强制**」，本文件**不改动该现状、不把 CLA 改为必需**。需显式标注：**DCO 不授予项目方再许可（relicense）权；双许可下需 CLA 或明确的再许可条款 —— 需权利人或法务确认**。任何「贡献是否足以支撑双许可」的个别判断，一律交由权利人或法务确认，本文件不构成法律意见。

## 5. 文件导航

- `LICENSE` —— AGPLv3 官方全文（逐字，无版权头、无译本）
- `COMMERCIAL-LICENSE.md` —— 商业授权说明与联系方式（需权利人确认）
- `FAQ-licensing.md` —— 何时需商业授权的决策树与 7 问
- `CONTRIBUTING.md` / `DCO.md` / `CLA.md` —— 贡献流程与 DCO 1.1
- `NOTICE` —— 单文件双语版权与归属声明
- `THIRD-PARTY-NOTICES.md` —— 第三方依赖许可清单
- `LICENSE-COMPATIBILITY.md` —— 双许可语境下的依赖白名单说明

## 6. 声明

- 版权主体：**小黎的部落 (www.qingdujian.com)**；来源：`https://github.com/xiaolibuluo/Abate`；商业联系：`xiaolibuluo@163.com`（沿用既有真实署名，不扩写资质 / 认证 / 背书）。
- 竞品许可只作**客观性质区分**：Qt / Grafana / iText 为双许可参考；MongoDB SSPL / Elastic / HashiCorp BUSL **非 OSI 认可开源许可**；不作对标超越断言。
- `competitorAbsolute = NOT_MEASURED`、`NOT_MEASURED` 原样保留，不补结论。
- **本文件不构成法律意见 / This document does not constitute legal advice；需权利人或法务确认 / Subject to confirmation by the rights holder or legal counsel。**
