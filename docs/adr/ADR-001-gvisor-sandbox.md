# ADR-001：采用 gVisor 作为能力沙箱（源码级并入）

- 状态：已采纳（Accepted）
- 日期：2026-11-08
- 提出：项目主理人（xiaolibuluo@163.com）
- 关联：ARCHITECTURE.md §5.2 安全模型、PLAN-00 模块地图 WAS-CAP、PLAN-08 横切安全体系

---

## 1. 背景与决策

项目要求：**所有接入的开源项目必须源码接入（禁止 API 接入）、必须是 Go 语言、禁止接入 GPL 许可的开源项目**（商业双授权前提下的供应链合规铁律）。

悬隔沙箱（sandbox）是 Abate 安全纵深的核心：L1 `WAS-CAP` 能力沙箱负责对 WASM 插件做 fs/net/clock/env 能力授予与默认全拒；M4 的 WASM 运行时需要在宿主机侧提供进程/命名空间级隔离，防止恶意或未签名插件逃逸。

经评估，采用 **gVisor**（`gvisor.dev/gvisor`）作为该沙箱底座。

---

## 2. 候选对比（事实）

| 候选 | 语言 | 许可证 | 是否 Go | 是否 GPL | 结论 |
|---|---|---|---|---|---|
| gVisor (runsc / pkg/sentry) | Go | Apache-2.0 | 是 | 否 | **采纳** |
| firecracker (vmm) | Rust | Apache-2.0 / BSD-3 | 否（Rust） | 否 | 排除：非 Go，违反「必须 Go」铁律 |
| containerd + runc (libcontainer) | Go | Apache-2.0 | 是 | 否 | 备选：需真实内核命名空间，嵌套沙箱不可用 |
| bubblewrap (bwrap) | C | LGPL-2.1 | 否（C） | **是（LGPL）** | **排除**：非 Go + GPL 系，双重违规 |

---

## 3. 许可证审查结论

- gVisor 主许可证：**Apache License 2.0**（与项目自带的 Apache-2.0 兼容条款一致，且为宽松许可，不构成对商业授权的传染性约束）。
- 不含 GPL/LGPL/AGPL 任何变体；无 copyleft 传染性风险。**符合「禁止接入 GPL」铁律。**
- 并入时以 `third_party/gvisor/` 路径 vendoring，并在 `third_party/LICENSE-MANIFEST.md` 记录来源 commit 哈希与许可证审查结论（DoD 第 6 条）。

---

## 4. 源码接入策略（五步并入法）

严格遵循 ARCHITECTURE.md §2.1 五步并入法，禁止任何 API/进程外/网络接入：

1. **vendor**：将 gVisor 相关源码（`pkg/sentry` 等隔离内核子集，非完整 runsc CLI）拉入 `third_party/gvisor/`，锁定来源 commit 哈希。
2. **二次更改**：确定性改造（对齐虚拟时钟/种子语义），去除与内核无关的 CLI/网络路径。
3. **优化**：裁剪非核心路径、收敛为库形态（沙箱原语），去除外部依赖漂移。
4. **深度集成**：将 gVisor 隔离原语封装为 `WAS-CAP` 能力授予/默认全拒的底层实现，不直接暴露 gVisor 专有类型。
5. **封装**：对外只暴露 Abate 能力沙箱接口（fs/net/clock/env port），gVisor 概念全部翻译为内核词汇。

---

## 5. 范围与里程碑归属（重要）

gVisor 沙箱归属 **M4（WASM 运行时 + 生态最小闭环）**，对应模块 **WAS-CAP / WAS-HOST**，**不在 M1 内核最小闭环范围内**。

理由（事实，非权宜）：

- 架构铁律：**L0/L3 内核必须纯软件数字孪生、零外部运行时依赖**（PLAN-04 通用约束第 1 条、PLAN-00 依赖铁律第 3 条）。在 M1 内核中直接链接 gVisor 会破坏该铁律。
- M1 的「沙箱」已由 L3 确定性模拟器（虚拟时钟/网络/FS/调度全虚拟化）与 L0 自包含内核承担，本就无需真实平台沙箱。
- 本 ADR 仅完成「决策 + 许可证审查 + 源码接入策略 + 门禁约束」四件可落地的当下工作；**具体源码 vendoring 与封装在 M4 执行**，届时受本 ADR 与 LICENSE-MANIFEST 约束。

---

## 6. 门禁约束（当下即可执行）

- `scripts/license-check.sh` 将 gVisor 列为已审查白名单（Apache-2.0）；任何新增依赖若被判定为 GPL 系，**门禁直接 FAIL**。
- 所有自研内核包不 import `third_party/gvisor` 之外的任何外部运行时；gVisor 仅在 L1 层经端口接入，L0 对其零感知（依赖铁律第 1、3 条）。

---

## 7. 正向后果

- 安全纵深获得经大规模生产验证的用户态内核隔离底座，且不引入 GPL 合规风险。
- 全链路源码自持，无供应链劫持面（XCU-SEC 目标）。

## 8. 反向后果 / 风险

- gVisor 体量较大，vendoring 后需维护上游安全补丁回移流程（与 chaosblade/PIT 同源管理）。
- 嵌套沙箱（在容器/沙箱中再跑 gVisor）在部分 CI 环境不可用；需在 M4 明确运行矩阵（ptrace 平台优先，KVM 平台按需）。

## 9. 备选方案

- 若 M4 运行环境无法承载 gVisor，回退到 containerd/runsc 同系的纯 Go 隔离原语子集，或项目自研轻量 syscall 过滤层；二者均为 Apache-2.0 / Go，不违反本 ADR 的许可证与语言铁律。
