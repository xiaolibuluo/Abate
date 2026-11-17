<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com)
  Project: xiaolibuluo/Abate — https://github.com/xiaolibuluo/Abate
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->
> 语言 / Language：[中文](THIRD-PARTY-NOTICES.md) | [English](THIRD-PARTY-NOTICES.en.md)

# 第三方依赖许可声明（THIRD-PARTY-NOTICES）

> 本文件为**逐项许可清单**；治理视角的并入总册与审查流程见 `governance/THIRD-PARTY.md`（二者不冲突：前者管流程、后者管逐项归属）。
> **本文件不构成法律意见 / This document does not constitute legal advice；需权利人或法务确认。**

## 1. 登记范围

- 仅登记交接件已核实、随仓库 `vendor/` 落地的第三方 Go 依赖；**不虚构未核实的组件 / 版本 / 许可文本**。
- 单一事实来源（SSOT）为各依赖随附的 `LICENSE` 文本与 `third_party/LICENSE-MANIFEST.md`；本表许可列为 `scripts/license-check.sh` 的既有判定（ALLOW / MIT），具体许可文本**以 SSOT 为准**，**需权利人或法务确认**。

## 2. 逐项登记

| 组件 | 版本 | 用途 | 许可（license-check 判定） | 接入方式 |
|---|---|---|---|---|
| wazero | v1.12.0 | 真 WASM 运行时 | ALLOW（在白名单内；具体许可文本以 SSOT 为准） | vendoring / 模块源码并入 |
| golang.org/x/sys | v0.44.0 | 系统调用原语 | ALLOW（在白名单内；具体许可文本以 SSOT 为准） | Go module（源码可读） |
| gopkg.in/yaml.v3 | v3.0.1 | YAML 解析 | MIT | Go module（源码可读） |

## 3. 导入依赖白名单

- **允许**：MIT / Apache-2.0 / BSD / ISC / CC0 / Unlicense / Public Domain。
- **拒绝**：GPL / LGPL / AGPL 任何变体（项目自身对外双授权 AGPLv3 + 商业不在此限）。
- 审查结论覆盖：每条登记须有明确「允许 / 拒绝 / 待评审」；GPL 系命中即拒绝；未知许可默认拒绝，不默许。

## 4. 交叉链接

- `governance/THIRD-PARTY.md` —— 五步并入法、铁律、治理索引。
- `NOTICE` —— 版权与归属声明（单文件双语）。
- `LICENSE-COMPATIBILITY.md` —— 双许可语境下的依赖白名单说明。
- `LICENSING.md` —— 双许可两路径总览。

## 5. 声明

- 版权主体：**小黎的部落 (www.qingdujian.com)**；来源：`https://github.com/xiaolibuluo/Abate`（沿用真实署名）。
- 竞品许可只作**客观性质区分**，不作对标超越断言；`competitorAbsolute = NOT_MEASURED`、`NOT_MEASURED` 原样保留。
- SPDX `AGPL-3.0-or-later`（**需权利人确认 / subject to confirmation by the rights holder**）。
- **本文件不构成法律意见 / This document does not constitute legal advice；需权利人或法务确认 / Subject to confirmation by the rights holder or legal counsel。**
