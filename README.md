<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->

> 语言 / Language：[中文](README.md) | [English](README.en.md)

# Abate

[![version](https://img.shields.io/badge/version-1.2.28-blue)](https://github.com/xiaolibuluo/Abate/releases/latest)
[![license](https://img.shields.io/badge/license-AGPLv3%20%7C%20commercial-green)](LICENSING.md)
[![checks](https://img.shields.io/badge/checks-7%20languages-orange)](docs/user/multilang-check-guide.md)
[![Go](https://img.shields.io/badge/Go-1.26.11-00ADD8?logo=go)](go.mod)

**确定性多语言静态分析器 + 实验控制平面**。单个静态二进制提供两类能力：

- **`abate check`**：按扩展名自动分派 **Go / Python / JavaScript / C / C# / Rust / Kotlin** 七种语言，
  用自研 lexer + Pratt 解析器把源码解析成统一 AIR 语法树，在树上做结构模式匹配，
  输出逐字节可复现的 console / JSON / SARIF 报告。
- **`abate run`**：以 CES（`ces.tencent.com/v1`）描述确定性实验，用固定种子内核与 SIM-REC 档案，
  产出可逐字节比对（配合 `abate verify` / `abate replay`）的运行报告。

> **能力边界（先说清楚）**：`abate check` 是**结构化静态 linter**，不是编译器级语义分析。
> 它不做类型检查、控制流/数据流分析，也不等同 go vet / eslint / clippy 的深度；
> 采用容错解析，语法错误不会让命令报错退出，因此**不用于语法错误守门**。
> 所有声明都能用一条命令自证——见[证据与边界](#7-证据与边界自证)。

## 1 目录 / Contents

| 章节 | 内容 |
|---|---|
| [2 环境要求](#2-环境要求) | Go 版本、零依赖、安装 |
| [3 快速开始](#3-快速开始) | 构建 + 两个真实例子 |
| [4 命令一览](#4-命令一览) | 9 个命令与退出码 |
| [5 能力边界与诚实口径](#5-能力边界与诚实口径) | 真解析/容错/规则深度/未测量项 |
| [6 文档导航](#6-文档导航) | 用户/开发/规范/法律文档索引 |
| [7 证据与边界（自证）](#7-证据与边界自证) | 一键复现脚本与门禁 |
| [8 许可与社区状态](#8-许可与社区状态) | 双许可、贡献、社区事实 |

## 2 环境要求

- **Go 1.26.11**（唯一权威版本，以 `go.mod` 的 `go` 指令为准）。
- **零外部依赖**：`go.mod` 无 `require`、`go.sum` 为空、无 `vendor/`、无 `third_party/`；
  离线可直接 `GOPROXY=off` 构建。
- 可选：`golangci-lint`（版本以本机安装为准）。

发布二进制为五平台静态构建：`linux/amd64` `linux/arm64` `darwin/amd64` `darwin/arm64` `windows/amd64`，
每个都带 GPG detached 签名（密钥 `5CF28F616D9888DB`）。见 [Releases](https://github.com/xiaolibuluo/Abate/releases/latest)。

## 3 快速开始

```bash
export GOPROXY=off
go build -o abate .
./abate version          # abate 1.2.28 / schema ces.tencent.com/v1 / ...
```

**例子一：七语言源码检查**（按扩展名自动分派，无需指定语言）：

```bash
./abate check ./path/to/your/code        # 混合目录一次检查 .go/.py/.js/.c/.cs/.rs/.kt
./abate check app.py --format json       # 机器可读报告
./abate check ./src --languages py,js --rules PY-PEP8
./abate check ./src --format sarif --out report.sarif
```

> `check` 只要检出任意 finding（含 warning）即返回退出码 **1**；
> 只想让 error 触发失败，加 `--severity error`。详见 [CLI 参考](docs/user/cli-reference.md)。

**例子二：跑一个 CES 实验**（仓库自带合法样例）：

```bash
./abate run ctl/cli/testdata/ok.yaml --seed 42 --format json --out report.json
./abate verify report.json report.json     # 逐字节比对，diff=0 退出 0
./abate run ctl/cli/testdata/ok.yaml --seed 42 --record /tmp/run.arch
./abate replay /tmp/run.arch               # 校验档案 hash 链/全序/单调性
```

## 4 命令一览

| 命令 | 作用 |
|---|---|
| `run <ces-file>` | 解析/校验 CES，跑确定性实验出报告；`--dry-run` 只校验、`--record` 写 SIM-REC 档案 |
| `verify <a> <b>` | 逐字节比较两份报告，打印 SHA256、大小、首个差异偏移 |
| `replay <archive> [other]` | 校验档案完整性，可选比较两份并输出结构化 Diff |
| `check <paths...>` | 七语言结构检查；`--languages` 筛选、`--rules` 前缀筛选、`--format console\|json\|sarif` |
| `derive` | 从内置 Go 规则派生 M 族变异算子，跑一次确定性变异实验（killing 为静态近似） |
| `plugin` | 签名插件：`load` / `list` / `call` / `hotswap`（Ed25519 签名校验） |
| `web` | 本地控制 API + 内嵌浏览器控制台（同源；非 loopback 绑定必须带 `--token`） |
| `version` | 打印版本与 provenance |
| `--help` | 顶层帮助 |

**退出码**：`0` 成功；`1` 业务失败（稳态 VIOLATED / 有 finding）；`2` 内部错误；`3` 用法错误。
`check` 和 `replay` 有各自的专用语义，详见 [CLI 参考](docs/user/cli-reference.md) 与 [退出码说明](docs/user/cli-reference.md#退出码)。

## 5 能力边界与诚实口径

Abate 只写能被当前代码与脚本复现的事实，以下边界如实声明：

| 维度 | 当前事实 |
|---|---|
| 解析方式 | 自研 lexer + Pratt parser → `air.Node` AST → 结构模式匹配（**不是正则扫描文本**） |
| 语法守门 | 容错解析：破损输入产出 partial AIR + warnings，**不做语法正确性守门** |
| 语义深度 | 不做类型检查 / CFG / 数据流 / 跨过程分析；不等同专用工具深度 |
| 规则深度 | 共享 AIR 模板 72 条 + 每语言专属规则（Go 116，其余 12–14）；Go 命中需特定语义上下文 |
| 全语料保真度 | 六语 3%–5%（见 [七语言覆盖](docs/user/language-coverage.md)）；Go 无嵌入语法映射为 `NOT_MEASURED`；不构成"高保真"结论 |
| 竞品对比 | `competitorAbsolute = NOT_MEASURED`（从未真实对拍，不写绝对分数） |
| derive killing | 静态 AIR 级近似，**非真实 go test 执行** |
| 社区验证 | 0 star / 0 fork / 0 watcher、个人维护；不刷 star、不伪造 benchmark |

详见 [能力与边界](docs/user/capabilities-and-limits.md)。

## 6 文档导航

**用户文档**（中英双语）
- [快速开始](docs/user/quickstart.md) · [CLI 参考](docs/user/cli-reference.md) · [CES 指南](docs/user/ces-guide.md)
- [七语言 check 指南](docs/user/multilang-check-guide.md) · [能力与边界](docs/user/capabilities-and-limits.md) · [七语言覆盖量化](docs/user/language-coverage.md)

**架构与规范**
- [架构总览](docs/ARCHITECTURE.md) · [CES v1 规范](docs/spec/ces-v1.md)

**开发**
- [开发指南](docs/dev/contributing.md) · [插件开发](docs/dev/plugin-dev-guide.md) · [模块开发](docs/dev/module-dev-guide.md) · [沙箱](docs/dev/sandbox.md)

**法律与治理**（不构成法律意见）
- [LICENSING.md](LICENSING.md) · [COMMERCIAL-LICENSE.md](COMMERCIAL-LICENSE.md) · [FAQ-licensing.md](FAQ-licensing.md) · [LICENSE-COMPATIBILITY.md](LICENSE-COMPATIBILITY.md)
- [CONTRIBUTING.md](CONTRIBUTING.md) · [DCO.md](DCO.md) · [CLA.md](CLA.md) · [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md)
- [SECURITY.md](SECURITY.md) · [NOTICE](NOTICE) · [THIRD-PARTY-NOTICES.md](THIRD-PARTY-NOTICES.md) · [TRADEMARK.md](TRADEMARK.md)
- 治理细则：`governance/`

## 7 证据与边界（自证）

一切能力声明都能从源码现场复现，不必信文档：

```bash
bash scripts/evidence-check.sh     # 现场构建 + 7 项能力断言，全 PASS 才退出 0
bash scripts/language-coverage.sh  # 七语言规则数与实测命中量化表
bash scripts/fullcorpus-degrade.sh # 全语料降级率（诚实口径）
bash scripts/lint-gate.sh          # 门禁：header/style/license/vet/golangci-lint/test-race/build
```

此外：`hash-check.sh` / `watermark-verify.sh`（完整性水印）、`determinism-check.sh`（确定性）、
`drift-check.sh`（规格一致性）、`audit-zero-third-party.sh`（零第三方源码）、
`m5-reference-scan.sh`、`vuln-scan.sh`、`tla/`（TLA+ 形式化）。

## 8 许可与社区状态

Abate 采用 **AGPLv3（开源）+ 商业授权** 双许可：开源路径见 `LICENSE`（AGPLv3 全文），
SPDX 标识 `AGPL-3.0-or-later`（**需权利人确认，不构成法律意见**）；商业路径见 `COMMERCIAL-LICENSE.md`。

贡献：提交须带 `Signed-off-by`（`git commit -s`，DCO 1.1）；CLA 当前可选
（注意 **DCO 不授予再许可权，双许可下需 CLA 或明确再许可条款，待权利人/法务确认**）。

社区状态：截至本文更新，仓库为 **0 star / 0 fork / 0 watcher、个人维护**——早期项目的客观事实，
如实呈现，不粉饰。项目靠一键复现 + 公开量化 + 机器可证伪的确定性降低试用门槛，欢迎真实失败用例。
