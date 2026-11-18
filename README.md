<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->

> 语言 / Language：[中文](README.md) | [English](README.en.md)

# Abate

[![version](https://img.shields.io/badge/version-1.2.27-blue)](https://github.com/xiaolibuluo/Abate/releases/latest)
[![license](https://img.shields.io/badge/license-AGPLv3%20%7C%20commercial-green)](LICENSING.md)
[![checks](https://img.shields.io/badge/checks-7%20languages-orange)](docs/user/multilang-check-guide.md)
[![Go](https://img.shields.io/badge/Go-1.26.11-00ADD8?logo=go)](go.mod)

**确定性多语言静态分析 / linter** —— `abate check` 按文件扩展名自动分派
**Go、Python、JavaScript、C、C#、Rust、Kotlin** 七种语言，经统一 AIR 中间表示
运行各语言自研规则集，输出逐字节可复现的 console / JSON / SARIF 报告，可直接入 CI。

同时也是一个确定性实验控制平面：以 CES（`ces.tencent.com/v1`）描述混沌实验与变异测试，
经由 AIR 前端、确定性内核与 SIM-REC 档案，产出可逐字节复现（固定种子下）的运行报告与证据链。

## 1 许可 / License

Abate 采用 **AGPLv3（开源）+ 商业授权（闭源）双许可**：

- 开源路径：`LICENSE`（AGPLv3 官方全文），SPDX 标识 `AGPL-3.0-or-later`
  （**需权利人确认，不构成法律意见**）；
- 商业路径：`COMMERCIAL-LICENSE.md`（联系 `xiaolibuluo@163.com`）；
- 双许可总览与两路径对比表：`LICENSING.md`；
- 判断辅助：`FAQ-licensing.md`（仅信息导航）。

## 2 环境要求 / Requirements

- Go **1.26.11**（真实构建版本；历史文档中的 `go1.27`、`Go 1.24+` 均为旧值）
- 离线构建：`GOFLAGS=-mod=vendor`（依赖已 vendor，无 git remote）
- 可选：`golangci-lint`（版本以本机安装为准，本文档不写未核实的版本号）

## 3 快速开始 / Quick start

```bash
export GOFLAGS=-mod=vendor
go build -o abate .
./abate version
```

保姆级教程：

- T01 从零离线安装与构建：`docs/i18n/tutorials/T01-build.md`
- T02 第一个混沌实验：`docs/i18n/tutorials/T02-first-run.md`

## 4 教程索引 / Tutorials

| 教程 | 中文 | English |
|---|---|---|
| T01 从零离线安装与构建 | [T01-build.md](docs/i18n/tutorials/T01-build.md) | [T01-build.en.md](docs/i18n/tutorials/T01-build.en.md) |
| T02 第一个混沌实验 | [T02-first-run.md](docs/i18n/tutorials/T02-first-run.md) | [T02-first-run.en.md](docs/i18n/tutorials/T02-first-run.en.md) |
| T03 干跑校验 CES | [T03-dry-run.md](docs/i18n/tutorials/T03-dry-run.md) | [T03-dry-run.en.md](docs/i18n/tutorials/T03-dry-run.en.md) |
| T04 固定种子 + JSON 输出 | [T04-seed-json.md](docs/i18n/tutorials/T04-seed-json.md) | [T04-seed-json.en.md](docs/i18n/tutorials/T04-seed-json.en.md) |
| T05 CLI 全命令逐个手把手 | [T05-cli-help.md](docs/i18n/tutorials/T05-cli-help.md) | [T05-cli-help.en.md](docs/i18n/tutorials/T05-cli-help.en.md) |
| T06 退出码四态逐个验证 | [T06-exit-codes.md](docs/i18n/tutorials/T06-exit-codes.md) | [T06-exit-codes.en.md](docs/i18n/tutorials/T06-exit-codes.en.md) |
| T07 CES 文件从零写 | [T07-write-ces.md](docs/i18n/tutorials/T07-write-ces.md) | [T07-write-ces.en.md](docs/i18n/tutorials/T07-write-ces.en.md) |
| T08 如何贡献（DCO/CLA） | [T08-contribute.md](docs/i18n/tutorials/T08-contribute.md) | [T08-contribute.en.md](docs/i18n/tutorials/T08-contribute.en.md) |
| T09 判断 AGPLv3 还是商业授权 | [T09-license-decision.md](docs/i18n/tutorials/T09-license-decision.md) | [T09-license-decision.en.md](docs/i18n/tutorials/T09-license-decision.en.md) |
| T10 FAQ 导航 + 如何跑门禁 | [T10-faq-gates.md](docs/i18n/tutorials/T10-faq-gates.md) | [T10-faq-gates.en.md](docs/i18n/tutorials/T10-faq-gates.en.md) |

## 5 文档导航 / Documentation

- 许可与合规：`LICENSING.md` · `COMMERCIAL-LICENSE.md` · `FAQ-licensing.md` · `LICENSE-COMPATIBILITY.md`
- 贡献与署名：`CONTRIBUTING.md` · `DCO.md` · `CLA.md`（草案） · `CODE_OF_CONDUCT.md`
- 安全与声明：`SECURITY.md` · `NOTICE` · `THIRD-PARTY-NOTICES.md` · `TRADEMARK.md`
- 治理细则：`governance/DUAL-LICENSE.md` · `governance/HEADERS.md` · `governance/THIRD-PARTY.md`
- 用户文档：`docs/user/quickstart.md` · `docs/user/cli-reference.md` · `docs/user/ces-guide.md`

## 6 贡献 / Contributing

- 每次提交须带 `Signed-off-by`（`git commit -s`，DCO 1.1）；
- CLA 当前为**可选**；注意 **DCO 不授予项目方再许可（relicense）权，双许可下
  需要 CLA 或明确再许可条款——需权利人或法务确认**；
- 依赖白名单：MIT / Apache-2.0 / BSD / ISC / CC0 / Unlicense / Public Domain；
  拒绝 GPL 系（详见 `LICENSE-COMPATIBILITY.md`）。

## 7 门禁 / Gates

| 阶段 | 工具 | 规则来源 | 失败条件 |
|------|------|----------|----------|
| 1 | gofmt | CodeReviewComments · Gofmt | 存在未格式化文件 |
| 2 | go vet | 官方 36 条分析器（printf/copylocks/loopclosure/lostcancel 等全量） | 任一分析器报错 |
| 3 | golangci-lint | `.golangci.yml` 全量 linter（110+ 条） | 任一 linter 命中（--issues-exit-code=1） |
| 4 | go test -race -count=1 | 全部测试 + 竞态检测 | 测试失败或存在数据竞态 |
| 5 | go build | 编译验证 | 编译失败 |

### 杜绝项（零容忍）

- **桩/占位**：`TODO` / `FIXME` / `BUG` / `HACK` 注释 → `godox` 拦截
- **伪示例**：Example 无 `// Output:` → `testableexamples` 拦截
- **忽略错误**：未检查返回值 → `errcheck`(含 blank/类型断言)、`wrapcheck`、`err113`、`errorlint`
- **安全隐患**：硬编码密钥、弱随机数、注入面 → `gosec`、`bidichk`
- **资源泄漏**：HTTP body / sql.Rows 未关闭 → `bodyclose`、`rowserrcheck`、`sqlclosecheck`
- **复杂度超标**：圈复杂度 >15（`gocyclo`/`cyclop`）、认知复杂度 >20（`gocognit`）、嵌套 >4（`nestif`）、函数 >60 行（`funlen`）、重复代码 >100 token（`dupl`）
- **规范缺失**：命名/注释/导出规范 → `revive`(全规则)、`stylecheck`、`godot`（注释句号结尾）
- **并发/测试纪律**：缺 `t.Parallel()` → `paralleltest`；helper 缺 `t.Helper()` → `thelper`；竞态 → `-race`
- **半成品**：不可编译、未使用代码 → `go build`、`unused`、`deadcode` 类规则

### 使用方式

```bash
# 本地执行完整门禁链
make gate
# 或直接调用脚本
bash scripts/lint-gate.sh
```

任一阶段失败，脚本立即以非零码退出，并打印命中的规则条例与位置，CI 同步 FAIL。

### 规则基准

- [CodeReviewComments](https://go.dev/wiki/CodeReviewComments)
- [Effective Go](https://go.dev/doc/effective_go)
- [go vet 分析器清单](https://pkg.go.dev/cmd/vet)（36 条全量启用）
- [Go 语言规范](https://go.dev/ref/spec)（Go 1.26.11）
- golangci-lint：`default: none` 起步，启用全部非废弃 linter；废弃项（exhaustruct/gomodguard/wsl）已按官方指引迁移至 v5/v2 替代或剔除

### 版权头与许可证脚本的精确边界（勿误读）

- `scripts/header-check.sh` **只扫 `*.go`，完全不扫 `.md`**：新建双语 `.md`
  缺版权头**不会**触发该门禁 FAIL；统一带 6 行 HTML 版权头属**治理一致性要求**
  （见 `governance/HEADERS.md`），非门禁强制项。
- `scripts/license-check.sh` 的 GPL 扫描针对 `vendor/modules.txt` 依赖许可文本与
  `third_party/LICENSE-MANIFEST.md`，**不扫仓库根 `LICENSE`**：落地 AGPLv3 全文
  不会误报 FAIL。
- `scripts/lint-gate.sh` 第 41–135 行为冻结九段，不改。

## 8 语言与声明 / Language & Status

- 本文件为中文主门面，英文逐节对应译本：`README.en.md`；
- 本文件不构成法律意见；`AGPL-3.0-or-later` 需权利人确认；
- 诚实口径（如实限定，勿外推）：
  - 竞品绝对分数 `competitorAbsolute = NOT_MEASURED`（未真实对拍前不写竞品绝对分数）；
  - E2E degrade「达标」**仅限冻结的单文件（首 unit）口径**。早期第三方语料上按同一
    公式测得六语均超阈（js 49.12% / c 28.00% / py 26.89% / rs 42.33% / kt 8.21% /
    cs 6.98%），go 为 `NOT_MEASURED`——该组数值是**语料全面自研化之前的历史测量**，
    仅见 `docs/i18n/QA-ANTI-FRAUD-AUDIT.md` 等历史审计证据，**不代表当前语料状态**。
    当前测试语料已 100% 替换为自研语料（见 `docs/refactor/`）；自研语料的**全语料口径**
    降级率已用 `scripts/fullcorpus-degrade.sh` 对全部 20520 条逐条真实 parse 重测，
    实测为 py 3.36% / js 3.51% / c 4.80% / cs 4.65% / rs 4.91% / kt 4.48%，
    go 无嵌入语法映射仍为 `NOT_MEASURED`（原始数据与 SHA-256 见
    `docs/refactor/fullcorpus-degrade/`）。js 由 46.49% 降至 3.51%，源于 jsfront
    为语料中真实存在的 20 类构造（export/generator/lexical/this-super/class-body/
    class-heritage/switch/new/paren/sequence/incdec/yield/template-sub/string-part/
    property-ident/arguments/formal-params/declarator 及 debugger/meta-property/
    computed-name/pair/pair-pattern/object-assign/rest-pattern 等）在真实语法归约点
    产出专用 `KindJs*` 零宽节点、并由规则标为 exact——每个 exact 都有前端真实产出
    支撑（最小样例 + 真实语料双重守护测试），绝不靠改标签造假；c 亦以同范式新建
    `cdedicated.go` 专用 marker 基础设施（union/枚举成员/字段声明/解引用/初始化设计符/
    defined/语句标签七类）由 10.40% 降至 4.80%。剩余降级构造（解构模式/empty-stmt 等
    映射到 AIR 尚无独立节点的共享 kind）如实保留 degraded 并由 pin 测试锁定。
    该全语料数值只作诚实披露，单文件口径与全语料口径不得混用外推，
    也禁止把上述历史第三方数值当作当前自研语料的实测值；
  - `derive` 的 killing 是 **static AIR-level approximation, not a real go test
    execution**（静态 AIR 级近似，非真实执行 go test）。
