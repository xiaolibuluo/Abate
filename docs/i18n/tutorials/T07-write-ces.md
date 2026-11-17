<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->

> 语言 / Language：[中文](T07-write-ces.md) | [English](T07-write-ces.en.md)

# T07 · CES 文件从零写

## 1 目标读者

要编写自己第一个 CES 实验的用户：目标是从仓库现成的 36 行样例出发，看懂
每个字段，复制修改并通过干跑校验。

## 2 前置条件

- 已按 T01 完成构建；
- 仓库根目录执行；
- `ok.yaml` = `ctl/cli/testdata/ok.yaml` 的简写（下文全文即该真实文件的
  36 行，**字段名一律以真实文件为准，不臆造**）。

## 3 可复制命令

先看真实样例全文，再复制改造：

```bash
cat ctl/cli/testdata/ok.yaml
cp ctl/cli/testdata/ok.yaml /tmp/my-exp.yaml
/tmp/abate_verify run /tmp/my-exp.yaml --dry-run
/tmp/abate_verify run /tmp/my-exp.yaml
```

## 4 预期输出

`cat ctl/cli/testdata/ok.yaml`（真实文件全文，36 行）：

```yaml
apiVersion: ces.tencent.com/v1
kind: Experiment
metadata:
  name: cli-e2e-ok
  namespace: default
spec:
  target:
    kind: simulator
    selector:
      node: n1
  steadyState:
    - probeRef: health
      expectation:
        type: value
        value: "200"
      tolerance:
        type: absolute
        value: 1
      holdFor: 1s
  method:
    expression: "ε"
  probes:
    - id: health
      kind: http
      spec:
        url: http://sim/health
        method: GET
        expectStatus: 200
      interval: 1s
  rollback:
    strategy: auto
    verifySteadyState: true
  schedule:
    once:
      at: 5s
  seed: 12345
```

`--dry-run` 报告头（与 T03 相同的实测锚点）：

```
status: VALIDATED
exitCode: 0
events: 0
observations: 0
```

完整 run 的关键行（与 T02 相同的实测锚点）：

```
status: COMPLETED
exitCode: 0
  health | HELD | held span 3000000000 satisfies holdFor 1000000000
```

**验证点**：干跑 `VALIDATED`；完整 run `COMPLETED` 且判定行 HELD。

**逐字段速读**：`apiVersion/kind` 信封铁律（见 `docs/user/ces-guide.md`）；
`target.selector.node` 指向模拟器内节点；`steadyState` 引用探针 `health`
并要求其值在容差 1 内保持 1s；`method.expression` 取恒等扰动 `ε`（本例
不注入真实故障，是最小合法实验）；`probes[].id` 文档内唯一；
`rollback.strategy: auto`；`schedule` 三选一（本例 `once`）；`seed` 为
文档种子（12345，可被 `--seed` 覆盖）。

## 5 失败排查

| 现象 | 规则 | 修复 |
|---|---|---|
| 表达式语法错 | CES-VALID-METHOD-001 | 检查算子/括号/符号；恒等扰动写 `ε` |
| 非可逆算子无补偿 | CES-VALID-METHOD-002 | 加 `compensation` 或 `invertible: true` |
| 嵌套超深 | CES-VALID-METHOD-003 | 深度 ≤ 64 |
| 调度多模式并存 | CES-VALID-SCHEDULE-001 | `continuous/once/cron` 只留一个 |
| `external` 目标 | CES-VALID-TARGET-003 | 改用 `simulator/source` |
| 改了字段名报未知字段 | 信封 fail-fast，未知字段直接拒绝 | 对照上方 36 行真实文件逐字段核对 |
