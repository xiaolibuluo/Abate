# 历史台账：清除前的第三方许可登记（仅供溯源）

> 归档来源：`third_party/LICENSE-MANIFEST.md`（git HEAD @ `3d822901`，T02 清除前原样取出）。
> 归档时间：2026-11-17（T02）。归档人：寇豆码（Kou，工程师）。
>
> **状态声明**：这是**清除前**的历史审计快照，**不代表当前状态**。
> 彻底自研重构执行后，产品运行时的第三方依赖 = **0**（`go.mod` 清零在 T04 完成，
> `vendor/` 在 T04 最后删除）。本文件仅用于溯源历史登记与许可证审查过程，
> **不参与任何门禁判定**，不作为"当前接入"的证据。
>
> 原表中 `gVisor / yaml.v3 / SplitMix64` 及"E2E 验证语料"各条均为历史规划或
> 已迁移的测试夹具来源登记：E2E 语料已整体搬迁到 `testdata/corpus/`，仅作解析器
> 测试输入，不参与产品编译、不随产品二进制分发（详见 `testdata/corpus/README.md`）。
>
> ⚖️ 历史并入/许可兼容性的遗留义务须由权利人/法务确认；本文件只陈述事实，不出法律结论。

---

以下为原文件全文（未改一字，保持历史审计真实性）：

# 第三方开源接入清单与许可证审查（third_party）

> 本清单是 Abate 供应链合规的**单一事实来源（SSOT）**。
> 任何新增开源依赖（无论直接 import 还是 vendoring 到本目录）**必须**先在此登记并通过许可证审查，否则 `scripts/license-check.sh` 门禁 FAIL。

## 接入铁律（项目红线）

1. **源码接入，禁止 API 接入**：一律 vendoring / 模块源码并入，禁止进程外调用、gRPC、网络依赖、运行时动态拉取外部引擎。
2. **必须是 Go 语言**：非 Go 实现的开源（如 Rust/C 系）不予接入（gVisor 为 Go 实现，例外仅限已被 ADR 明确批准的 Go 子集）。
3. **禁止 GPL 系许可证**：GPL-2.0 / GPL-3.0 / LGPL / AGPL 任何变体**一律拒绝**（项目自身双授权 AGPLv3+商业 不在此限，那是本项目对外授权，不是导入依赖）。允许：MIT / Apache-2.0 / BSD 系列 / ISC / 公有领域（CC0 / Unlicense / public-domain）。
4. **登记来源 commit 哈希**：vendoring 必须锁定上游 commit，便于安全补丁回移与可溯。

## 已审查登记

| 组件 | 来源 | 版本/Commit | 许可证 | 语言 | 接入方式 | 里程碑 | 状态 | 审查结论 |
|---|---|---|---|---|---|---|---|---|
| gVisor（隔离内核子集） | gvisor.dev/gvisor | _待 M4 锁定_ | Apache-2.0 | Go | vendor → 五步并入 | M4 | 决策已采纳，未 vendoring | ✅ 无 GPL，纯 Go，符合铁律（见 ADR-001） |
| yaml.v3 | gopkg.in/yaml.v3 | v3.0.1 | MIT | Go | Go module（源码可读） | M1 | 规划中 | ✅ 无 GPL，纯 Go |
| SplitMix64 | 公有领域算法（数学文献） | _算法常量_ | Public Domain | Go（自实现） | 算法源码并入 sim/rand | M1 | 规划中 | ✅ 零许可证负担 |

## 许可证判定规则（scripts/license-check.sh 执行）

- 扫描 `go list -m all` 的全部导入模块 + `third_party/` 登记项。
- 命中以下关键词即判定为 **GPL 系 → FAIL**：`GNU GENERAL PUBLIC LICENSE`、`GNU LESSER GENERAL PUBLIC`、`GNU AFFERO`、`GPL`、`LGPL`、`AGPL`（大小写不敏感，但仅针对**导入依赖**，不针对本仓库自身 LICENSE）。
- 命中以下关键词判定为 **允许**：`MIT`、`Apache`、`BSD`、`ISC`、`CC0`、`Unlicense`、`public domain`。
- 未知/无法判定的许可证 → 人工评审 FAIL（默认拒绝，不默许）。

## 审计记录

- 2026-11-08：建立本清单；gVisor 经 ADR-001 审查采纳（Apache-2.0，无 GPL，纯 Go）。

## E2E 验证语料 SPDX 审查登记（T01）

> 11 个 pinned 上游项目的 SPDX 表达式**逐字登记**，全部落在允许集
> `MIT` / `Apache-2.0` / `BSD-3-Clause` / `MIT OR Apache-2.0`，未命中任何拒绝项。
> `ci/corpus_manifest.json` 每条 entry 携带 `spdx` 与 `spdxEvidence`
> （LICENSE 文件相对路径 + 该文件原始字节 sha256），`corpusctl verify` 对全量
> 20520 条重算并分类，实测 `SPDXViolations=0`、`BoundarySPDXCount=3302`
> （`Apache-2.0 + LLVM exception` 单列，如实登记、不删除、不豁免）。
> 4 项 commit 为 `RATE_LIMIT_PENDING`，以 tag + tarball_sha256 锁定版本，不伪造哈希。

| 组件 | 来源 | 版本/tag | 许可证 | 语言 | 接入方式 | 审查结论 |
|---|---|---|---|---|---|---|
| gin | github.com/gin-gonic/gin | v1.9.1 | MIT | Go | 只读源码落地 | ✅ 允许集内 |
| requests | github.com/psf/requests | v2.31.0 | Apache-2.0 | Python | 只读源码落地 | ✅ 允许集内 |
| click | github.com/pallets/click | 8.1.7 | BSD-3-Clause | Python | 只读源码落地 | ✅ 允许集内 |
| lodash | github.com/lodash/lodash | 4.17.21 | MIT | JavaScript | 只读源码落地 | ✅ 允许集内 |
| prettier | github.com/prettier/prettier | 3.1.1 | MIT | JavaScript | 只读源码落地 | ✅ 允许集内 |
| express | github.com/expressjs/express | 4.18.2 | MIT | JavaScript | 只读源码落地 | ✅ 允许集内 |
| axios | github.com/axios/axios | v1.6.2 | MIT | JavaScript | 只读源码落地 | ✅ 允许集内 |
| cJSON | github.com/DaveGamble/cJSON | v1.7.16 | MIT | C | 只读源码落地 | ✅ 允许集内 |
| Newtonsoft.Json | github.com/JamesNK/Newtonsoft.Json | 13.0.3 | MIT | C# | 只读源码落地 | ✅ 允许集内 |
| serde | github.com/serde-rs/serde | v1.0.193 | MIT OR Apache-2.0 | Rust | 只读源码落地 | ✅ 双许可均在允许集内 |
| okhttp | github.com/square/okhttp | parent-4.12.0 | Apache-2.0 | Kotlin | 只读源码落地 | ✅ 允许集内 |

> 审计补充（T01）：以上源码全部位于 `_` 前缀目录，不进入 `go list ./...`、
> 不参与编译与覆盖率、不新增 `go.mod` 依赖面；`go.mod` / `go.sum` /
> `vendor/modules.txt` 与基线逐字节一致。逐包 tarball sha256 见
> `scripts/corpus_fetch.sh` 与 `third_party/_sources/README.md`。
