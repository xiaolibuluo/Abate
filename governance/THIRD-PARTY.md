<!--
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
-->

# 第三方源码并入登记册（XCU-GOV）

> 治理视角的第三方并入总册。逐条登记的**单一事实来源（SSOT）**为
> `third_party/LICENSE-MANIFEST.md`（运营清单，CI 由 `scripts/license-check.sh`
> 与 `security/audit/scan.sh` 校验）。本文档是其治理索引与流程说明。

## 1. 并入铁律（项目红线）

1. **源码接入，禁止 API 接入**：一律 vendoring / 模块源码并入，禁止进程外调用、
   gRPC、网络依赖、运行时动态拉取外部引擎。
2. **必须是 Go 语言**：非 Go 实现（Rust/C 系）不予接入；gVisor 为 Go 实现，
   例外仅限已被 ADR 明确批准的 Go 子集。
3. **禁止 GPL 系许可证**：GPL-2.0 / GPL-3.0 / LGPL / AGPL 任何变体一律拒绝
   （项目自身双授权 AGPLv3+商业不在此限）。允许：MIT / Apache-2.0 / BSD /
   ISC / CC0 / Unlicense / Public Domain。
4. **登记来源 commit 哈希**：vendoring 必须锁定上游 commit，便于安全补丁回移与可溯。

## 2. 五步并入法

1. 评估：许可证兼容性 + 语言 + 功能必要性（ADR 记录决策）。
2. 登记：`third_party/LICENSE-MANIFEST.md` 新增条目（组件/来源/版本或 commit/
   许可证/语言/接入方式/里程碑/状态/审查结论）。
3. 并入：vendoring 或模块源码并入，锁定 commit；保留上游版权头。
4. 裁剪：移除无关文件，记录裁剪清单；不引入非 Go 依赖。
5. 回移：上游安全补丁经评审回移，记录于审计（对齐 KRN-INJ 上游回移机制）。

## 3. 当前登记索引（指向 SSOT）

| 组件 | 用途 | 许可证 | 接入位置 | 状态 |
|---|---|---|---|---|
| gVisor（隔离内核子集） | 沙箱隔离（M4） | Apache-2.0 | vendor → 五步并入 | 决策已采纳，未 vendoring |
| yaml.v3 | YAML 解析 | MIT | Go module（源码可读） | 规划中 |
| SplitMix64 | PRNG 算法 | Public Domain | 算法源码并入 sim/rand | 规划中 |

> 实际登记与状态以 `third_party/LICENSE-MANIFEST.md` 为准；本索引仅作治理概览。

## 4. 审查结论覆盖

- 每条登记必须有明确许可证审查结论（允许 / 拒绝 / 待评审）。
- GPL 系命中即拒绝，零例外（除本项目对外双授权）。
- 未知许可证默认拒绝，不默许。
