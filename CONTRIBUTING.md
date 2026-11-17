<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com)
  Project: xiaolibuluo/Abate — https://github.com/xiaolibuluo/Abate
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->
> 语言 / Language：[中文](CONTRIBUTING.md) | [English](CONTRIBUTING.en.md)

# 贡献指南（根目录 CONTRIBUTING）

> 本文件是仓库根目录的贡献入口；`docs/dev/contributing.md` 是开发流程文档，二者路径不同、均保留。
> **本文件不构成法律意见 / This document does not constitute legal advice；需权利人或法务确认。**

## 1. 双许可与贡献

Abate 采用「AGPLv3（开源侧）+ 商业授权（闭源侧）」双许可，见 `LICENSING.md`。贡献前请阅读 `LICENSING.md`、`DCO.md`、`CLA.md`。所有法律判断**需权利人或法务确认**。

## 2. DCO 1.1 与签署（sign-off）

- **DCO 强制**：每次提交须含 `Signed-off-by` 行，声明你有权提交该贡献并同意 DCO 1.1。
- 文档示例（**仅示例，本轮不执行任何 git commit**）：

```bash
git commit -s -m "docs: add bilingual quickstart"
```

- `git commit -s` 会自动追加 `Signed-off-by: <Name> <email>`。DCO 1.1 官方全文见 `DCO.md`。

## 3. CLA 现状（草案、待确认）

- 现有治理为「**CLA 可选、DCO 强制**」，本文件**不改该现状、不把 CLA 改为必需**。
- `CLA.md` 为**草案**，**是否必需待权利人或法务确认**。
- 显式标注：**DCO 不授予项目方再许可（relicense）权；双许可下需 CLA 或明确再许可条款 —— 需权利人或法务确认**。贡献的版权权利如何被项目方在双许可下使用，本文件不代下结论。

## 4. 导入依赖白名单

- **允许**：MIT / Apache-2.0 / BSD / ISC / CC0 / Unlicense / Public Domain。
- **拒绝**：GPL / LGPL / AGPL 任何变体（与 AGPLv3 传染性冲突；项目自身双授权不在此限）。
- 审查流程见 `governance/THIRD-PARTY.md` 与 `third_party/LICENSE-MANIFEST.md`；逐项登记见 `THIRD-PARTY-NOTICES.md`；兼容性说明见 `LICENSE-COMPATIBILITY.md`。

## 5. 门禁（措辞依 §6b）

- `header-check.sh` **只扫 `*.go`、完全不扫 `.md`**；新建双语 `.md` 缺 6 行版权头**不会**触发 header-check FAIL，但依 `governance/HEADERS.md` §2 属**治理一致性要求**，应统一带上。
- `license-check.sh` **不扫仓库根 `LICENSE`**；落地 AGPLv3 全文 `LICENSE` **不会**误报 FAIL。
- `lint-gate.sh` 冻结九段（41–135 行）不改。

## 6. 声明

- 版权主体：**小黎的部落 (www.qingdujian.com)**；来源：`https://github.com/xiaolibuluo/Abate`；商业联系：`xiaolibuluo@163.com`（沿用真实署名，不扩写资质）。
- `competitorAbsolute = NOT_MEASURED`、`NOT_MEASURED` 原样保留。
- **本文件不构成法律意见 / This document does not constitute legal advice；需权利人或法务确认 / Subject to confirmation by the rights holder or legal counsel。**
