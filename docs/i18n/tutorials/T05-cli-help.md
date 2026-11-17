<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->

> 语言 / Language：[中文](T05-cli-help.md) | [English](T05-cli-help.en.md)

# T05 · CLI 全命令逐个手把手

## 1 目标读者

想把 8 个子命令全部过一遍的用户与运维：目标是会用 `--help` 自助查命令，
并能读懂命令树、退出码四态与 6 个常用 flag。

## 2 前置条件

- 已按 T01 完成构建；
- 本教程只跑 `--help`，不会执行实验、不写任何文件；
- 命令树：`run / verify / replay / check / derive / plugin / web / version`
  （外加 `help` / `--help` / `-h`）。

## 3 可复制命令

```bash
/tmp/abate_verify --help
/tmp/abate_verify run --help
/tmp/abate_verify verify --help
/tmp/abate_verify replay --help
/tmp/abate_verify check --help
/tmp/abate_verify derive --help
/tmp/abate_verify plugin --help
/tmp/abate_verify web --help
/tmp/abate_verify version --help
```

## 4 预期输出

`--help`（顶层，实测节选——命令树与退出码段为逐字引用）：

```
abate - deterministic experiment control plane

Usage:
  abate <command> [flags]

Commands:
  run <ces-file>     Execute a CES experiment and emit a deterministic report.
                     --record <path> also writes a SIM-REC archive.
  verify <a> <b>     Compare two run reports byte for byte.
  replay <archive> [other]  Verify an archive, optionally compare two.
  check <paths...>   Run the built-in Go rule set; emit console/JSON/SARIF.
  derive             Derive M-family operators from rules; run a mutation
                     experiment (static AIR-level killing approximation).
  plugin             Verify, load, call, list and hot-swap signed plugins.
  web                Serve the control API with a self-contained browser console.
  version            Print build version and provenance.

Exit codes (PLAN-07 §4):
  0  success (legal CES and steady state HELD).
  1  business failure (steady state VIOLATED during execution).
  2  internal error (scheduler or report output failure).
  3  usage error (invalid arguments, invalid CES document, or a CES-VALID
     error violation; --strict promotes warnings to this code too).
```

`run --help`（6 个 flag，逐字）：

```
  --seed uint        Deterministic seed override (0 uses the document seed).
  --format console|json  Report format (default console).
  --out path          Also write the report bytes to a file.
  --record path       Write the deterministic SIM-REC archive to a file.
  --dry-run           Parse and validate only; do not execute.
  --strict            Promote warning-severity violations to failures.
```

`derive --help`（关键句**原样保留**，不得改写为「真实执行 go test 的变异
测试」）：

```
Killing is a static AIR-level approximation, not a real go test execution.
```

**验证点**：9 条 `--help` 全部退出码 0；命令树 8 个子命令与上文一致；
退出码四态 0/1/2/3 在 `run`/`replay`/`check` 的 help 中均有声明。

## 5 失败排查

| 现象 | 原因 | 动作 |
|---|---|---|
| `--help` 报 exit 3 | 子命令名拼错 | 先跑 `abate --help` 看命令树 |
| help 文本与本教程不符 | 运行了旧二进制 | 重跑 T01 构建 |
| 分不清 `check` 与 `derive` | 两者都涉及「规则」 | `check` 只读源码出报告，**从不改动源码**；`derive` 派生 M 族算子并跑变异实验（静态 AIR 级近似） |
| `web` 起服务但打不开控制台 | 绑定了非回环地址且未给 token | 默认绑 `127.0.0.1:8430`；非回环绑定必须配 `--token`，否则拒绝启动 |
