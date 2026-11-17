<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->

> 语言 / Language：[中文](T04-seed-json.md) | [English](T04-seed-json.en.md)

# T04 · 固定种子 + JSON 输出

## 1 目标读者

要把实验结果接入流水线/做逐字节比对复核的用户：目标是掌握 `--seed` 与
`--format json`，并理解「确定性必须固定种子」这一口径。

## 2 前置条件

- 已按 T01 完成构建；
- 仓库根目录执行；
- `ok.yaml` = `ctl/cli/testdata/ok.yaml` 的简写。

## 3 可复制命令

```bash
/tmp/abate_verify run ctl/cli/testdata/ok.yaml --seed 42 --format json > /tmp/report-42.json
/tmp/abate_verify run ctl/cli/testdata/ok.yaml --seed 42 --format json > /tmp/report-42-again.json
/tmp/abate_verify verify /tmp/report-42.json /tmp/report-42-again.json
```

## 4 预期输出

JSON 报告中的关键字段（实测）：

```json
{
  "seed": 42,
  "status": "COMPLETED",
  "runDigest": "32dbea3d9acc1c154a49a869dda11243db1a2e24cbeb64d5248619b3cb318339"
}
```

（节选：完整 JSON 报告包含更多字段，此处仅列出教程锚定的三个键。）

`verify` 退出码 `0`（两份报告逐字节一致）。**验证点**：固定 `--seed 42`
后，`runDigest` 每次运行均为同一值（以 `32dbea3d9` 开头）。

**诚实口径**：确定性宣称**必须固定 `--seed`**。种子覆盖会改变 digest——
同一文档不指定 `--seed` 时用文档种子 12345，digest 以 `6444a24ed2` 开头
（见 T02）；`--seed 42` 时为 `32dbea3d9…`。两个 digest 不同正是种子生效的
证据，不是不确定性。

## 5 失败排查

| 现象 | 原因 | 动作 |
|---|---|---|
| `verify` 退出码 1 | 两份报告有差异（`verify` 打印首个差异偏移与 SHA-256） | 确认两次命令的 `--seed` 与文件完全一致 |
| digest 与本教程不同 | 用了不同的 seed 或不同的 CES 文件 | 精确照抄第 3 节命令 |
| JSON 解析失败 | 管道里混入了非 JSON 输出 | 确认 `--format json`；报告只写 stdout，无其他前缀 |
| 想留档复核 | 内存报告易丢 | 加 `--out /tmp/report.json` 同时落盘（报告字节与 stdout 一致） |
