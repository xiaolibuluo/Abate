<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->

> 语言 / Language：[中文](T09-license-decision.md) | [English](T09-license-decision.en.md)

# T09 · 判断 AGPLv3 还是商业授权

## 1 目标读者

评估是否在自有产品中使用 Abate 的技术决策者：目标是 10 分钟内完成许可
路径自查，明确该找哪份文件、该问谁。

## 2 前置条件

- 无需构建二进制，但跑一次 `abate version` 可以直观看到许可声明；
- 治理文件：`LICENSING.md` · `FAQ-licensing.md`（决策树）·
  `LICENSE`（AGPLv3 官方全文）· `COMMERCIAL-LICENSE.md`；
- **本教程不构成法律意见**；所有个别合法性判断需权利人或法务确认。

## 3 可复制命令

```bash
/tmp/abate_verify version
sed -n '536,544p' LICENSE
```

## 4 预期输出

`version` 输出中的许可行（6 行构建溯源之一，实测）：

```
license: AGPLv3 or commercial
```

`sed -n '536,544p' LICENSE` 落在 AGPLv3 官方全文的 **§13 Remote Network
Interaction; Use with the GNU General Public License**（该节标题位于
LICENSE 第 540 行）：

```
  13. Notwithstanding any other provision of this License, if you modify the
Program, your modified version must prominently offer all users interacting
with it remotely through a computer network (if your version supports such
interaction) an opportunity to receive the Corresponding Source of your
modified version by providing access to the Corresponding Source through a
network server at no charge, through some standard or customary means of
facilitating copying of software.
```

（节选自官方全文；完整义务以仓库根 `LICENSE` 逐字全文为准。）

**决策树（仅信息导航，不构成法律意见）**：

1. **能否履行 AGPLv3 全部义务**（含 §13：向通过网络远程交互的用户提供
   Corresponding Source）？
   - 能 → 走 AGPLv3（`LICENSE` / `CONTRIBUTING.md` / `DCO.md`）；
   - 否 → 进入第 2 步。
2. **是否闭源集成 / 不愿公开衍生作品源码 / 不愿承担 §13 源码提供义务**？
   - 是 → 需商业授权，联系 `xiaolibuluo@163.com`（沿用既有联系方式，
     是否适用需双方书面协议）。
3. **SPDX 用 only 还是 or-later** → 统一标注 `AGPL-3.0-or-later`
   （**需权利人确认，不写成已定**）。
4. **如何看待竞品许可** → 仅作客观性质区分：Qt / Grafana / iText 为双许可
   参考；MongoDB SSPL / Elastic / HashiCorp BUSL **非 OSI 认可开源许可**；
   不作对标超越断言。

**验证点**：能说出自己的场景属于第 1 步「能履行」还是第 2 步「需商业
授权」，并知道下一步该读哪份文件。

## 5 失败排查

| 现象 | 原因 | 动作 |
|---|---|---|
| 分不清「用 AGPLv3 不代表不能商业用」 | 双许可 = 两条可选路径，不是互斥门槛 | 读 `LICENSING.md` 两路径对比表 |
| §13 义务理解不确定 | 法律问题 | 需权利人或法务确认；只读官方全文原文，勿凭二手转述 |
| 商业询价无回音 | 本文档不含定价/已签约/用户案例/SLA | 直接邮件 `xiaolibuluo@163.com` 洽谈 |
| 任何「我这样做是否合法」 | 本教程只导航不裁决 | 统一答：需权利人或法务确认，本文档不构成法律意见 |
