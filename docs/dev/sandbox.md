<!--
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
-->

# 插件沙箱模型（plugin/sandbox）

> **如实声明：本轮未运行 runsc 进程，仅模型对接。**
> `plugin/sandbox` 只提供**纯函数**的配置模型与策略校验：不调用 `runsc` 二进制、
> 不 fork 任何进程、不做动态加载、不触碰网络。它把 gVisor runsc 的**配置语义**
> 抽象为 Abate 自有的、可确定性比较的纯数据文档。

## 1. 模型来源

| Abate 命名 | gVisor runsc 对应语义 | 上游来源（只读参考） |
|---|---|---|
| `Platform`（`ptrace` / `systrap` / `kvm`） | `--platform` 平台层选择 | `runsc/config/config.go`（`Platform` 字段）、`pkg/sentry/platform/{ptrace,systrap,kvm}` |
| `NetworkPolicy`（`none` / `sandbox` / `host`） | `--network` 网络栈选择 | `runsc/config/config.go`（`NetworkType`：`NetworkSandbox` / `NetworkHost` / `NetworkNone`） |
| `RootFS`（OCI 风格 rootfs 路径 + 摘要） | rootfs bundle 与 gofer 访问 | `runsc/config/config.go`（`FileAccess`、`Overlay`）、`pkg/sentry/fsgofer` |
| `FilesystemPolicy`（只读/可写挂载白名单） | bind mount 与 file-access 语义 | `runsc/config/config.go`（`FileAccessType`：`FileAccessExclusive` / `FileAccessShared`） |
| `MemoryLimitMB` | 沙箱内存上限 | `runsc/config/config.go`（`TotalMem` 等限额项的语义） |
| `Capability`（`fs` / `net` / `clock` / `env`） | runsc 的能力开关族 | `runsc/config/config.go`（`EnableRaw`、`PauseExternalNetworking` 等开关族） |

上游源码位于 `third_party/_sources/gvisor/`（`google/gvisor` 2026-11-11 快照，Apache-2.0），
**只读参考，不 import、不编译、不进 `go.mod`**。Abate 侧为**纯 Go 标准库自研重写**，
使用 Abate 自己的命名空间，不复制上游代码。

## 2. 裁剪边界（做了什么、没做什么）

**做了**：

- 平台层 / 网络姿态 / 文件系统白名单 / 内存上限 / 能力集合的**纯数据模型**；
- 从插件声明（`Manifest`）到沙箱配置（`Config`）的**确定性推导**（`Build`）；
- 能力清单与沙箱配置的一致性**校验**（`Validate`，默认全拒）；
- 确定性 JSON 序列化（`Marshal` / `Unmarshal` / `Config.Canonical`）。

**明确没做（如实登记）**：

| 未实现项 | 原因 |
|---|---|
| 不调用 `runsc` 进程、不 fork 任何进程 | 项目铁律：禁 `os/exec`、禁进程调用、禁动态加载外部二进制（`third_party/_references/MANIFEST.md` 登记裁定） |
| 不实现 runsc 的 checkpoint/restore、cgroup、metrics server | 超出「解析 + AST + 沙箱模型」范围，且属于 M5 前端蓝本范畴 |
| 不实现真实网络栈与真实文件系统挂载 | 本包只产出配置模型；真实隔离由宿主运行时负责 |
| 不实现 gVisor 的 sentry 平台层（KVM/ptrace/systrap） | 平台层是宿主能力，本包只声明「选哪一层」 |

## 3. 默认值与兼容性常量

| 常量 | 值 | 语义 |
|---|---|---|
| `DefaultPlatform` | `systrap` | 默认平台层（无 KVM 依赖，兼容性最好） |
| `DefaultNetwork` | `none` | 默认仅回环，零外网 |
| `DefaultMemoryLimitMB` | `256` | 默认内存上限（MiB） |
| `MaxMemoryLimitMB` | `4096` | 硬上限，超过即拒绝 |
| `DefaultRootFSPath` | `/var/lib/abate/sandbox/rootfs` | 默认 rootfs bundle 路径 |

默认姿态**默认全拒**：`DefaultConfig` 不请求任何能力、rootfs 只读、时钟虚拟化。

## 4. 校验口径（拒绝条件）

`Validate(config, granted)` 按固定顺序检查，命中即拒绝（错误用 `errors.Is` 判定）：

| 序号 | 条件 | 哨兵错误 |
|---|---|---|
| 1 | 插件 ID 为空 | `ErrInvalidManifest` |
| 2 | 平台层不在已知集合 | `ErrUnknownPlatform` |
| 3 | 网络姿态不在已知集合 | `ErrUnknownNetwork` |
| 4 | 内存上限 ≤ 0 | `ErrInvalidMemory` |
| 5 | 内存上限 > `MaxMemoryLimitMB` | `ErrMemoryLimit` |
| 6 | 请求了宿主未授予的能力 | `ErrCapabilityDenied` |
| 7 | rootfs 路径非绝对路径或含 `..` | `ErrInvalidPath` |
| 8 | 挂载点非绝对路径或含 `..` | `ErrInvalidPath` |
| 9 | 挂载点逃出 rootfs 根目录 | `ErrPathEscape` |
| 10 | 非 `none` 姿态但无 `net` 能力 | `ErrPolicyMismatch` |
| 11 | `none` 姿态却带网络白名单 | `ErrPolicyMismatch` |
| 12 | 需要写权限但无 `fs` 能力 | `ErrPolicyMismatch` |
| 13 | rootfs 可写但无 `fs` 能力 | `ErrPolicyMismatch` |
| 14 | 带环境变量白名单但无 `env` 能力 | `ErrPolicyMismatch` |
| 15 | 请求 `clock` 能力但未启用虚拟时钟 | `ErrPolicyMismatch` |

## 5. 确定性保证

- `Build` 对同一份 `Manifest` 恒等产出同一份 `Config`；所有白名单在推导时排序。
- `Config.Canonical()` 归一化白名单顺序，因此**语义相同的配置序列化后逐字节相同**。
- `Marshal` 使用标准库 `encoding/json`，字段顺序由结构体声明顺序固定，无 map 参与。
- 全包不使用 `time.Now`、全局 `math/rand`、map 迭代序依赖。

## 6. 测试与门禁

```bash
GOFLAGS=-mod=vendor go test -race -count=1 ./plugin/sandbox/...
GOFLAGS=-mod=vendor golangci-lint run ./plugin/sandbox/...
```

- 生态包覆盖率门槛 ≥85%（实测见收尾输出）。
- 单元测试覆盖：默认配置、非法配置拒绝、越权能力拒绝、越界路径拒绝、
  确定性（同输入两次序列化逐字节一致）、JSON 往返。
- `sandbox_chaos_test.go`：百次复跑同一份文档 + 并发构建/校验/序列化。
- `sandbox_mutation_test.go`：策略矩阵逐行断言哨兵错误。

## 7. 与后续里程碑的关系

- M4 `plugin/wasm`（wazero 宿主）与本包的关系：wasm 宿主把本包的 `Config`
  作为**声明式契约**消费（宿主负责真实隔离），本包不感知 wasm 运行时。
- M5 语言前端（libclang-min）复用同一「只建模、不执行」的边界：
  `scripts/build-libclang-min.sh` 同样不产出假产物，缺失前置时如实报错。
