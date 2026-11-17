<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->

> 语言 / Language：[中文](T08-contribute.md) | [English](T08-contribute.en.md)

# T08 · 如何贡献（DCO/CLA）

## 1 目标读者

想向 Abate 提交补丁的贡献者：目标是走通 DCO 签名的提交格式，并理解依赖
白名单与再许可缺口这两条治理现状。

## 2 前置条件

- 已安装 `git`；
- 已 fork/克隆仓库；
- 治理文件：`CONTRIBUTING.md` · `DCO.md` · `CLA.md`（草案）·
  `governance/DUAL-LICENSE.md`；
- **本教程中的 `git commit` 仅为格式示例，本轮不执行任何 git 写操作**。

## 3 可复制命令

```bash
git clone https://github.com/xiaolibuluo/Abate.git
cd Abate
# （修改文件后）
git add <your-files>
git commit -s -m "docs: add bilingual quickstart"
git log -1 --format=%B
```

## 4 预期输出

`git commit -s` 会把 `Signed-off-by` 行追加到提交信息末尾，`git log -1
--format=%B` 显示的提交信息形如：

```
docs: add bilingual quickstart

Signed-off-by: Your Name <your.email@example.com>
```

**验证点**：提交信息末行是 `Signed-off-by: 真实姓名 <真实邮箱>`（DCO 1.1
声明：你有权提交该补丁并同意 DCO 1.1）。**没有该行 = 提交不合规范**。

**治理现状（如实陈述，不代下结论）**：

- **DCO 强制**：每次提交必须 `git commit -s`；
- **CLA 可选**：当前为可选，签署后明确授予项目使用与再许可权利；
- **再许可缺口**：**DCO 不授予项目方再许可（relicense）权；双许可下需要
  CLA 或明确再许可条款——需权利人或法务确认**；
- **依赖白名单**：仅允许 MIT / Apache-2.0 / BSD / ISC / CC0 / Unlicense /
  Public Domain；拒绝 GPL 系（GPL-2.0 / GPL-3.0 / LGPL / AGPL 任何变体）。
  引入依赖前先查 `third_party/LICENSE-MANIFEST.md` 与
  `LICENSE-COMPATIBILITY.md`。

## 5 失败排查

| 现象 | 原因 | 动作 |
|---|---|---|
| 提交缺 `Signed-off-by` | 忘了 `-s` | 补签：`git commit --amend -s`（未推送前） |
| `-s` 报「 unable to auto-detect email address」 | git 未配置身份 | `git config user.name` / `git config user.email` 后重试 |
| 想引入一个 LGPL 库 | 白名单拒绝 GPL 系 | 不引入；如确有必要走 ADR 评估，默认拒绝 |
| 不确定商业许可下能否用你的贡献 | 再许可权属法律问题 | 需权利人或法务确认；本教程不构成法律意见 |
