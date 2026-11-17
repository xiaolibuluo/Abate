<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->

> 语言 / Language：[中文](T01-build.md) | [English](T01-build.en.md)

# T02 · 第一个混沌实验

## 1 目标读者

刚完成构建（见 T01）的新用户：目标是用仓库自带的示例 CES 跑通第一次实验，
读懂报告头。

## 2 前置条件

- 已按 T01 完成构建（本教程用 `/tmp/abate_verify` 指代构建产物）；
- 仓库根目录执行；
- 本教程的 `ok.yaml` 是 `ctl/cli/testdata/ok.yaml` 的简写（真实存在的仓库
  文件，36 行）。

## 3 可复制命令

```bash
/tmp/abate_verify run ctl/cli/testdata/ok.yaml
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
dryRun: false
strict: false
status: COMPLETED
exitCode: 0
events: 6
observations: 4
end: 10000000000
runDigest: 6444a24ed20f880d4b3b9c0815ecfc429afcc7b3793619ba5e5cc733058043ef
streamDigest: 20220908d7fcaeddb5eb3141230b7908e428b370bf4479362c6d0add98570e1d
violations: 0
verdicts: 1
  health | HELD | held span 3000000000 satisfies holdFor 1000000000
```

**验证点**：`status: COMPLETED`、`exitCode: 0`、`violations: 0`；判定行
`health | HELD | ...` 表示稳态断言在 holdFor 内成立（HELD 是状态枚举值，
原样保留不译）。`runDigest` 以 `6444a24ed2` 开头——同样的命令与文件，
digest 逐字节可复现。

## 5 失败排查

| 现象 | 原因 | 动作 |
|---|---|---|
| `exit 3` / usage error | 文件路径写错（不在仓库根执行） | `cd` 到仓库根，或用绝对路径 `ctl/cli/testdata/ok.yaml` |
| `status: VIOLATED` | 跑了 `violated.yaml` 而非 `ok.yaml`（退出码 1） | 确认文件名；`violated.yaml` 是故意让稳态失败的教学件 |
| 报告字段与本教程不符 | 运行了旧二进制 | 重跑 T01 构建步骤 |
| 想要机器可读输出 | console 是默认格式 | 加 `--format json`（见 T04） |
