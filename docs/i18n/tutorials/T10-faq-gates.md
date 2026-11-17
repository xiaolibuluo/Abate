<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->

> 语言 / Language：[中文](T10-faq-gates.md) | [English](T10-faq-gates.en.md)

# T10 · FAQ 导航 + 如何跑门禁

## 1 目标读者

新加入的维护者与本地开发环境的搭建者：目标是知道常见问题的答案放在哪里
（FAQ 导航），并能本地跑通版权头与许可证两道脚本门禁。

## 2 前置条件

- 已按 T01 完成构建（`make test` 等需要完整工具链）；
- 仓库根目录执行；
- FAQ 入口：`FAQ-licensing.md`（许可决策树 + 7 问）；门禁细节：
  `README.md` §7 与 `governance/HEADERS.md`。

## 3 可复制命令

```bash
bash scripts/header-check.sh
bash scripts/license-check.sh
make gate
make test
make build
```

## 4 预期输出

`bash scripts/header-check.sh`（实测，逐字）：

```
[PASS] 全部源文件版权头齐全（含小黎的部落水印）

== HEADER GATE: PASSED ==
```

`bash scripts/license-check.sh`（实测，逐字；WARN 属预期，不 FAIL）：

```
[PASS] 全部导入依赖许可证合规（无 GPL）
[WARN] LICENSE-MANIFEST 含 GPL 字样，请确认仅指代项目自身双授权而非导入依赖。
[PASS] third_party/LICENSE-MANIFEST.md 存在且已登记

== LICENSE GATE: PASSED - 无 GPL 系依赖 ==
```

**验证点**：两脚本退出码均为 0、末行均为 `PASSED`（改动前基线即为
PASSED，本教程写作时实测复验一致）。

**门禁边界的精确措辞（勿误读）**：

- `header-check.sh` **只扫 `*.go`、完全不扫 `.md`**：新建双语 `.md` 缺
  版权头**不会**触发该门禁 FAIL；统一带头属治理一致性要求
  （`governance/HEADERS.md` §2），**非门禁强制项**；
- `license-check.sh` 的 GPL 扫描只针对 `vendor/modules.txt` 依赖许可文本与
  `third_party/LICENSE-MANIFEST.md`，**不扫仓库根 `LICENSE`**：仓库根落地
  AGPLv3 全文不会误报 FAIL（LICENSE-MANIFEST 中的 GPL 字样为项目自身双授权
  的预期 WARN）；
- `make gate`（`scripts/lint-gate.sh`）是完整五阶段门禁链（gofmt → go vet
  → golangci-lint → go test -race → go build），其中第 41–135 行为冻结
  九段，不改；本地需装好 `golangci-lint`（版本以本机安装为准）；
- `make test` / `make build` 分别对应 `go test -race -count=1 ./...` 与
  `go build ./...`，输出随运行环境而异，以「退出码 0 = 通过」判读。

**FAQ 导航（7 问入口）**：选 AGPLv3 还是商业授权 / §13 是什么 / SPDX 是
什么且本项目用哪个 / 为何要 THIRD-PARTY-NOTICES 与 NOTICE / DCO 够不够
（再许可缺口待确认）/ 竞品许可如何看 / 根 LICENSE 与新 .md 会不会触发
门禁 FAIL —— 全部见 `FAQ-licensing.md`。

## 5 失败排查

| 现象 | 原因 | 动作 |
|---|---|---|
| HEADER GATE: FAILED | 新增/改动的 `.go` 文件版权头缺失或不在前 8 行 | 按 `governance/HEADERS.md` §1 模板逐字节补齐 |
| LICENSE GATE: FAILED | `vendor/modules.txt` 依赖许可文本命中 GPL 系 | 不引入该依赖；流程见 `governance/THIRD-PARTY.md` |
| `.md` 缺头但门禁 PASSED | 属预期：header-check 不扫 `.md` | 想统一治理一致性则按 6 行 HTML 模板补头，非强制 |
| `make gate` 报 golangci-lint 不存在 | 本机未安装 | `make lint` 前先安装（路径取 `go env GOPATH`/bin），版本以本机安装为准 |
| WARN 让人误以为 FAIL | WARN 不改变退出码 | 看末行 `PASSED/FAILED` 与退出码判读 |
