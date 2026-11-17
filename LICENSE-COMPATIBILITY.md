<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com)
  Project: xiaolibuluo/Abate — https://github.com/xiaolibuluo/Abate
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->
> 语言 / Language：[中文](LICENSE-COMPATIBILITY.md) | [English](LICENSE-COMPATIBILITY.en.md)

# 许可证兼容性说明（LICENSE-COMPATIBILITY）

> 文件名由架构设计定案。本文件说明**双许可语境下的导入依赖白名单**，属信息导航；**不与 AGPLv3 条款冲突、不复制 / 改写 `LICENSE` 条款正文**。
> **本文件不构成法律意见 / This document does not constitute legal advice；需权利人或法务确认。**

## 1. 双许可语境

- Abate 采用「AGPLv3（开源侧）+ 商业授权（闭源侧）」双许可，见 `LICENSING.md`。
- 项目**自身**对外以 AGPLv3 / 商业授权分发；这一事实**不等于**允许把 GPL / LGPL / AGPL 变体的第三方代码并入内核：导入依赖仍须落在下方白名单内（`governance/DUAL-LICENSE.md` §3、`governance/THIRD-PARTY.md` §1）。

## 2. 导入依赖白名单

- **允许**：MIT / Apache-2.0 / BSD / ISC / CC0 / Unlicense / Public Domain。
- **拒绝**：GPL-2.0 / GPL-3.0 / LGPL / AGPL 任何变体 —— 与 AGPLv3 传染性冲突（项目自身双授权不在此限）。
- 任一具体依赖是否兼容，**以 `vendor/` 随附 `LICENSE` 文本与 `third_party/LICENSE-MANIFEST.md`（SSOT）为准，需权利人或法务确认**；本文件不代下逐项结论。

## 3. 与 AGPLv3 条款的关系

- `LICENSE` 为 AGPLv3 官方全文逐字副本；本文件**不改写**其任何条款，尤其 §13「Remote Network Interaction」（line 540）与 Corresponding Source 义务。
- 白名单是**项目治理约定**（`governance/HEADERS.md` / `governance/THIRD-PARTY.md`），不是对 AGPLv3 条款的重新解释；二者若有张力，**需权利人或法务确认**。

## 4. 权利链与再许可缺口

- 现有治理「CLA 可选、DCO 强制」不改；**DCO 不授予项目方再许可（relicense）权；双许可下需 CLA 或明确再许可条款 —— 需权利人或法务确认**；本文件**不把 CLA 改为必需**（`CLA.md` 为草案）。
- SPDX `AGPL-3.0-or-later`（**需权利人确认 / subject to confirmation by the rights holder**）；`only` / `or-later` 待权利人确认。

## 5. 交叉链接

- `THIRD-PARTY-NOTICES.md` —— wazero / x-sys / yaml.v3 逐项登记。
- `NOTICE` —— 版权与归属声明。
- `governance/THIRD-PARTY.md` —— 五步并入法与审查结论覆盖。
- `LICENSING.md` / `COMMERCIAL-LICENSE.md` / `FAQ-licensing.md`。

## 6. 声明

- 竞品许可只作**客观性质区分**：Qt / Grafana / iText 为双许可参考；MongoDB SSPL / Elastic / HashiCorp BUSL **非 OSI 认可开源许可**；不作对标超越断言。
- `competitorAbsolute = NOT_MEASURED`、`NOT_MEASURED` 原样保留。
- 版权主体：**小黎的部落 (www.qingdujian.com)**（沿用真实署名，不扩写资质）。
- **本文件不构成法律意见 / This document does not constitute legal advice；需权利人或法务确认 / Subject to confirmation by the rights holder or legal counsel。**
