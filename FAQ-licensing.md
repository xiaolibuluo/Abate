<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com)
  Project: xiaolibuluo/Abate — https://github.com/xiaolibuluo/Abate
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->
> 语言 / Language：[中文](FAQ-licensing.md) | [English](FAQ-licensing.en.md)

# 许可常见问题（FAQ-licensing）

> 本文件仅作信息导航，**本文件不构成法律意见 / This document does not constitute legal advice**；任何「这样是否合法」的个别判断统一答「**需权利人或法务确认**」。

## 1. 决策树（何时需商业授权）

```
① 能否履行 AGPLv3 全部义务（含 §13 Remote Network Interaction 向网络用户提供 Corresponding Source）？
    能 ──► 走 AGPLv3（LICENSE / CONTRIBUTING / DCO）
    否 ──► ② 闭源集成 / 不愿公开衍生作品源码 / 不愿承担 §13 义务？
              是 ──► 需商业授权，联系 xiaolibuluo@163.com（沿用待确认）
    ③ SPDX 采用 only 还是 or-later ──► 待权利人确认
    ④ 竞品许可只作性质区分陈述，不作对标超越断言
    ⑤ 任何「这样是否合法」的个别判断 ──► 统一答「需权利人或法务确认，本文档不构成法律意见」
```

## 2. 七问

### Q1 · 选 AGPLv3 还是商业授权？
两条**可选路径**并存（见 `LICENSING.md`）。能履行 AGPLv3 全部义务（含 §13）→ 开源侧；闭源集成 / 不愿公开衍生源码 / 不愿承担 §13 源码提供义务 → 商业侧，联系 `xiaolibuluo@163.com`（**沿用待确认**）。本 FAQ 不代下结论。

### Q2 · AGPLv3 §13 是什么？
第 13 节「Remote Network Interaction; Use with the GNU General Public License」位于 `LICENSE` 官方全文 **line 540**：若修改本程序并通过网络向用户提供交互，须向网络用户提供 Corresponding Source 的机会。本 FAQ 只描述与链接，**不复制、不改写**条款正文；解释**需权利人或法务确认**。

### Q3 · SPDX 是什么？本项目用哪个？
SPDX 是许可证标识。本项目文档头统一写 `AGPL-3.0-or-later`（**需权利人确认 / subject to confirmation by the rights holder，不写成已定**）；`AGPL-3.0-only` 与 `or-later` 的选择**待权利人确认**。官方全文「How to Apply」的 `or (at your option) any later version` 仅为参考措辞，不等于本项目已选。

### Q4 · 为何要 THIRD-PARTY-NOTICES 与 NOTICE？
`THIRD-PARTY-NOTICES.md` 逐项登记 wazero / x-sys / yaml.v3 等依赖（以 `vendor/` 与 `third_party/LICENSE-MANIFEST.md` 为 SSOT）；`NOTICE` 为单文件双语的版权与归属声明。二者用于如实披露第三方许可与归属，不作能力宣称。

### Q5 · 有了 DCO 是否足够（再许可缺口）？
现有治理「**CLA 可选、DCO 强制**」不改，**不把 CLA 改为必需**；`CLA.md` 是否必需**待权利人或法务确认**。显式标注：**DCO 不授予项目方再许可（relicense）权；双许可下需 CLA 或明确再许可条款 —— 需权利人或法务确认**。本 FAQ 不写「已合规」结论。

### Q6 · 竞品许可如何看？
仅作**客观性质区分**：Qt / Grafana / iText 为双许可参考；MongoDB SSPL / Elastic / HashiCorp BUSL **非 OSI 认可开源许可**。不与 AGPLv3 混为一谈，不作对标超越断言；`competitorAbsolute = NOT_MEASURED`、`NOT_MEASURED` 原样保留。

### Q7 · 根 `LICENSE` 与新建 `.md` 会不会触发门禁 FAIL？
依 §6b 纠偏：`header-check.sh` **只扫 `*.go`、完全不扫 `.md`**，新建双语 `.md` 缺 6 行版权头**不会**触发 FAIL，但依 `governance/HEADERS.md` §2 属治理一致性要求应带上；`license-check.sh` **不扫仓库根 `LICENSE`**，落地 AGPLv3 全文 `LICENSE` **不会**误报 FAIL；`lint-gate.sh` 冻结九段（41–135 行）不改。改动前两脚本基线均为 PASSED。

## 3. 交叉链接

`LICENSING.md` / `COMMERCIAL-LICENSE.md` / `LICENSE` / `CONTRIBUTING.md` / `DCO.md` / `CLA.md` / `NOTICE` / `THIRD-PARTY-NOTICES.md` / `LICENSE-COMPATIBILITY.md` / `TRADEMARK.md` / `SECURITY.md`。

## 4. 声明

- 版权主体：**小黎的部落 (www.qingdujian.com)**；来源：`https://github.com/xiaolibuluo/Abate`；商业联系：`xiaolibuluo@163.com`（沿用真实署名，不扩写资质）。
- 教程版决策路径见 `docs/i18n/tutorials/T09-license-decision.md`（中文）/ `.en.md`（英文）。
- **本文件不构成法律意见 / This document does not constitute legal advice；需权利人或法务确认 / Subject to confirmation by the rights holder or legal counsel。**
