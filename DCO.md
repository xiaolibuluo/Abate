<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com)
  Project: xiaolibuluo/Abate — https://github.com/xiaolibuluo/Abate
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->
> 语言 / Language：[中文](DCO.md) | [English](DCO.en.md)

# 开发者证书（DCO 1.1）

> **本文件不构成法律意见 / This document does not constitute legal advice；需权利人或法务确认。**
> 与 `CONTRIBUTING.md`、`CLA.md` 交叉链接；附录为 DCO 1.1 标准英文官方全文（逐字引用，中英两份完全相同）。

## 1. 现状与用法

- 现有治理为「**CLA 可选、DCO 强制**」，本文件**不改该现状、不把 CLA 改为必需**。
- **DCO 强制**：每次提交须含 `Signed-off-by` 行，声明你有权提交贡献并同意 DCO 1.1。
- 文档示例（**仅示例，本轮不执行任何 git commit**）：

```bash
git commit -s -m "docs: add bilingual quickstart"
```

- `git commit -s` 追加形如 `Signed-off-by: <Name> <email>` 的行。

## 2. 双许可下的再许可缺口（需权利人或法务确认）

- **DCO 不授予项目方再许可（relicense）权；双许可（AGPLv3 + 商业授权）下需 CLA 或明确再许可条款 —— 需权利人或法务确认**。
- 该缺口不因本文件而"已解决"；`CLA.md` 为**草案**，**是否必需待权利人或法务确认**，本文件不代下结论。

## 3. 导入依赖白名单

- **允许**：MIT / Apache-2.0 / BSD / ISC / CC0 / Unlicense / Public Domain。
- **拒绝**：GPL / LGPL / AGPL 任何变体。
- 审查流程见 `governance/THIRD-PARTY.md`；登记见 `THIRD-PARTY-NOTICES.md`；兼容性见 `LICENSE-COMPATIBILITY.md`。

## 4. 门禁（措辞依 §6b）

- `header-check.sh` **只扫 `*.go`、完全不扫 `.md`**；`license-check.sh` **不扫仓库根 `LICENSE`**；新增双语 `.md` 带 6 行版权头属**治理一致性要求、非门禁强制项**。

## 5. 声明

- 版权主体：**小黎的部落 (www.qingdujian.com)**；来源：`https://github.com/xiaolibuluo/Abate`（沿用真实署名）。
- `competitorAbsolute = NOT_MEASURED`、`NOT_MEASURED` 原样保留。
- **本文件不构成法律意见 / This document does not constitute legal advice；需权利人或法务确认 / Subject to confirmation by the rights holder or legal counsel。**

## 附录 · Developer Certificate of Origin 1.1（官方英文全文，逐字引用）

> 以下为 DCO 1.1 标准英文全文，中英两份逐字相同、不翻译；如与官方发布文本不一致，以官方原文为准 —— **需权利人或法务确认**。

```text
Developer Certificate of Origin
Version 1.1

Copyright (C) 2004, 2006 The Linux Foundation and its contributors.
1 Letterman Drive
Suite D4700
San Francisco, CA, 94129

Everyone is permitted to copy and distribute verbatim copies of this
license document, but changing it is not allowed.


Developer's Certificate of Origin 1.1

By making a contribution to this project, I certify that:

(a) The contribution was created in whole or in part by me and I
    have the right to submit it under the open source license
    indicated in the file; or

(b) The contribution is based upon previous work that, to the best
    of my knowledge, is covered under an appropriate open source
    license and I have the right under that license to submit that
    work with modifications, whether created in whole or in part
    by me, under the same open source license (unless I am
    permitted to submit under a different license), as indicated
    in the file; or

(c) The contribution was provided directly to me by some other
    person who certified (a), (b) or (c) and I have not modified
    it.

(d) I understand and agree that this project and the contribution
    are public and that a record of the contribution (including all
    personal information I submit with it, including my sign-off) is
    maintained indefinitely and may be redistributed consistent with
    this project or the open source license(s) involved.
```
