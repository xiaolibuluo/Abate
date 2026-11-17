# ADR-WASM-002：移除 wazero，收敛到原生 Go 插件宿主

- 状态：已采纳（Accepted）
- 日期：2026-11-16
- 决策来源：高见远（架构师）`docs/refactor/DESIGN.md` §2.1
- 落盘：寇豆码（工程师）
- 关联：TODO P1-01/P1-02/P1-04、WAS-HOST 冻结决策（`plugin/host/doc.go`）、ADR-001

> 本文记录架构师已作出的技术决策，不含手写签名；任何法律结论均不在本文范围。

---

## 1. 背景

原方案以 `github.com/tetratelabs/wazero v1.12.0` 作为 L1 插件运行时。合规重构要求
产品仓库零第三方源码副本、编译零第三方依赖，故必须处置该运行时。

## 2. 读码事实（非推测）

`plugin/wasm/` 共 32 个 `.go` / 4440 行，但 wazero 只出现在 5 个文件的薄层：

| 能力 | 具体 API | 出处 | 是否 WASI/全规范 |
|---|---|---|---|
| 运行时构建 | `NewRuntimeConfig().WithMemoryLimitPages`、`NewRuntimeWithConfig`、`Close` | `runtime.go:43,46,127` | 否，仅 1 个配置项 |
| 模块编译 | `rt.CompileModule` | `runtime.go:75`、`loader.go:25,34` | WAT 由自研 `watToWasm` 先转换（`wat_encode.go:15`） |
| 实例化 | `InstantiateModule(..., NewModuleConfig())` | `runtime.go:79` | 默认配置，无 FS/网络 |
| 宿主函数 | `NewHostModuleBuilder` → `WithGoModuleFunction` → `Export` → `Instantiate` | `hostfunc.go:56-64` | 仅 Go 闭包，无 WASI 导入 |
| 调用与内存 | `ExportedFunction(name).Call`、`api.ValueType`、`api.GoModuleFunction` | `instance.go:35,42,56` | 仅 i64 时钟读写 |

三个决定性事实：

1. **孤岛包**：排除 `vendor/`、`third_party/` 后，import `abate/plugin/wasm` 的产品代码 = **0 个文件**；
   `plugin/host` 则有 25+ 引用文件（CLI、hotswap、adapter、reporter、sample、security 测试）。
2. **已有替代**：`plugin/host/doc.go:10-16` 记载冻结决策 WAS-HOST —— 原生 Go 宿主，不 vendor 第三方 WASM 运行时。
   确定性依赖已由 `sim/clock`（SIM-CLK）与 `sim/rand`（SIM-RNG）注入，无需 WASM host function。
3. **fuel 本就是近似**：`fuel.go:19-21` 自述无 per-opcode fuel hook，仅在宿主调用边界精确，纯 guest 循环为过近似。

## 3. 候选方案

| 方案 | 工作量 | 零第三方 | 能力保留 | 评价 |
|---|---|---|---|---|
| A 自研极简 WASM 解释器 | 数千~上万行 / 数周~数月 | ✅ | 可加载子集，无 WASI/AOT | ❌ 无消费者，重写规范是负收益工程 |
| **B 废弃 WASM 执行路径，收敛 `plugin/host`** | 百行级语义迁移 + 目录删除 | ✅ | 保留全部**实际在用**能力：算子/注入器/前端/适配器/报告器插件、能力门控、验签、热替换、虚拟时钟/随机 | ✅ 符合冻结决策、风险最低、诚实 |
| C 自研插件二进制 ABI（非 WASM） | 中等 | ✅ | 仍非跨语言 | ⚠️ 远期增强，本次不做 |
| D 标准库 `plugin`（`.so`） | 小 | ✅（标准库） | 动态加载 | ❌ CGO/平台依赖，破坏气隙与确定性铁律 |

## 4. 决策

**采纳方案 B**。保留自研 WAT 编解码工具，但不作为插件加载路径。

分步实施：

| 步骤 | 动作 | 完成判据 |
|---|---|---|
| B1 | 冻结依赖面证据（本 ADR + §2 清单 + 审计脚本） | ADR 落盘；清单可复跑 |
| B2 | `FuelMeter`/`VirtualClock`/`Config` 确定性语义并入宿主层（`plugin/host/` 或 `internal/determinism/`，实现阶段定目录）；`SandboxFuel.tla` 状态机不变 | 行为与 T01 迁移前 golden 逐点一致 |
| B3 | `wat_*.go` 保留为离线工具或迁 `internal/wat/`（二选一，须记录） | 包内无第三方 import；`go test` 通过 |
| B4 | 删 `runtime.go`/`loader.go`/`hostfunc.go`/`instance.go`/`doc.go` 及相关测试 | 全仓产品代码 `wazero` 命中 = 0 |
| B5 | `plugin/host` 能力闭环确认（Load/capability/signature/hotswap） | 四类测试全绿，基线记入 `docs/refactor/plugin-baseline.md` |
| B6 | `go.mod` 清零（与 T04 合并） | 见 ADR-XSYS-001 / DESIGN §2.4 |

## 5. 门禁与叙事影响

- `SandboxFuel.tla`：抽象状态机不变；注释中的「WASM guest」改为「guest plugin」后重跑。
- `m5-reference-scan.sh` ① 段禁用 import 列表新增 `wazero` 关键词（零依赖断言）。
- `docs/plan/02-wasm-runtime.md`：改写为「原生插件宿主机（WAS-HOST 决策实录）」。
- 价值支柱：「快」改「确定性调度 + 轻量原生调用」；「劲」改「Go 插件 registry 热替换」。

## 6. 诚实边界（必须写进对外文档）

- ❌ 不得再声明「WASM 插件生态」「WASI Preview2」「WASM AOT 近原生」「任意语言编写插件」。
- ❌ 不得声称替换后性能与 wazero 等价；无对拍数据，保持 `NOT_MEASURED`。
- ✅ 可如实声明「原生 Go 插件宿主，能力门控 + 确定性时钟/随机 + 验签 + 原子热替换；插件以 Go 接口实现」。
- ⚠️ WAT 工具若保留，只是确定性字节码编解码工具，**不是插件运行时**。
