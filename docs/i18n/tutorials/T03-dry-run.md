<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->

> 语言 / Language：[中文](T03-dry-run.md) | [English](T03-dry-run.en.md)

# T03 · 干跑校验 CES

## 1 目标读者

正在编写自己 CES 文件的用户：目标是只做解析+校验、不执行实验，快速把
CES 文件调到合法。

## 2 前置条件

- 已按 T01 完成构建；
- 仓库根目录执行；
- `ok.yaml` = `ctl/cli/testdata/ok.yaml` 的简写。

## 3 可复制命令

```bash
/tmp/abate_verify run ctl/cli/testdata/ok.yaml --dry-run
```

## 4 预期输出

报告头（实测，逐字）：

```
abate run report
schema: abate.report.v1
apiVersion: ces.tencent.com/v1
name: cli-e2e-ok
namespace: default
seed: 12345
dryRun: true
strict: false
status: VALIDATED
exitCode: 0
events: 0
observations: 0
```

**验证点**：`dryRun: true`、`status: VALIDATED`、`exitCode: 0`、
`events: 0`、`observations: 0` ——干跑只校验文档合法性，不执行实验，
因此 events/observations 均为 0。

## 5 失败排查

| 现象 | 原因 | 动作 |
|---|---|---|
| `exit 3` 且报 CES-VALID-* 违规 | 文档不合法（fail-fast，未知字段直接拒绝） | 逐条修复；速查表见 `docs/user/ces-guide.md` 最后一节 |
| `status: VALIDATED` 但想看执行效果 | 干跑不执行，属预期 | 去掉 `--dry-run` 再跑（见 T02） |
| YAML 语法错（`malformed.yaml` 类） | 缩进/格式坏 | 用 `ok.yaml` 作模板逐行对照 |
| 想把 warning 也当失败 | 默认 warning 不阻断 | 加 `--strict`，warning 级违规提升为退出码 3 |
