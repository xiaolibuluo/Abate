<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->

> 语言 / Language：[中文](T06-exit-codes.md) | [English](T06-exit-codes.en.md)

# T06 · 退出码四态逐个验证

## 1 目标读者

要把 Abate 接进 CI/脚本的用户：目标是逐一实测退出码四态（0/1/2/3）中的
0/1/3，让脚本能可靠判灯。

## 2 前置条件

- 已按 T01 完成构建；
- 仓库根目录执行；
- 四份测试文档均为真实仓库文件：`ctl/cli/testdata/{ok,violated,invalid,malformed}.yaml`；
- 退出码 2（内部错误：调度或报告输出失败）为声明态，日常四份样例不触发，
  本教程不构造内部故障，如实说明即可。

## 3 可复制命令

```bash
/tmp/abate_verify run ctl/cli/testdata/ok.yaml        > /dev/null 2>&1; echo $?
/tmp/abate_verify run ctl/cli/testdata/violated.yaml  > /dev/null 2>&1; echo $?
/tmp/abate_verify run ctl/cli/testdata/invalid.yaml   > /dev/null 2>&1; echo $?
/tmp/abate_verify run ctl/cli/testdata/malformed.yaml > /dev/null 2>&1; echo $?
/tmp/abate_verify                                     > /dev/null 2>&1; echo $?
/tmp/abate_verify version                             > /dev/null 2>&1; echo $?
```

## 4 预期输出

六条 `echo $?` 依次输出（实测）：

```
0
1
3
3
3
0
```

**验证点**（与文档声明的 PLAN-07 §4 逐条对应）：

| 输入 | 退出码 | 含义 |
|---|---|---|
| `ok.yaml` | 0 | 合法 CES 且稳态 HELD |
| `violated.yaml` | 1 | 业务失败：执行中稳态 VIOLATED |
| `invalid.yaml` | 3 | 用法错误：CES 文档非法（CES-VALID 违规） |
| `malformed.yaml` | 3 | 用法错误：YAML 无法解析 |
| 无参数 | 3 | 用法错误：缺必选参数 |
| `version` | 0 | 成功 |

## 5 失败排查

| 现象 | 原因 | 动作 |
|---|---|---|
| 某行输出与本表不符 | 运行了旧二进制 | 重跑 T01 构建；如仍不符请保留现场并向维护者报告 |
| `violated.yaml` 让 CI 误判 | 退出码 1 是**业务失败**的预期教学态 | 在脚本里区分 1（稳态被打破）与 3（文档/参数非法） |
| 想复现退出码 2 | 需调度/报告输出内部错误，样例不触发 | 不建议构造；如实按 help 声明理解即可 |
| `echo $?` 输出顺序错乱 | 命令与 `echo $?` 之间插入了其他命令 | `$?` 只反映最近一条命令，逐行执行 |
