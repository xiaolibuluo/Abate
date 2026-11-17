# kernel/abi 移植与对拍说明

`kernel/abi/` 是从已入库 gVisor 源码（`third_party/_sources/gvisor/`，快照
`v0.0.0-20261111045955-a42ad63bdbc8`，Apache-2.0）**裁剪移植**的最小 Linux
syscall ABI 与 host-arch 参数子集，按 Abate 风格重写，不 import 上游。

## 移植范围

- amd64 / arm64 syscall 号→名表（上游 `pkg/sentry/syscalls/linux/linux64.go`）。
- Linux errno 常量（上游 `pkg/abi/linux/errno/errno.go`）。
- 地址与页大小原语 `Addr` / `AddrRange` / 分页对齐 / `PageSize` 等
  （上游 `pkg/hostarch/*`，amd64 / arm64 构建标签）。

## 对拍方式

`kernel/abi/syscall_test.go` 在 `go test` 时**直接读取**上游
`linux64.go`，解析 AMD64 / ARM64 两张表，与本包 `AMD64Names` / `ARM64Names`
逐项比对（数量、编号、名称）。上游源文件若更新，测试会立即暴露偏差。

对拍结论（2026-11-13）：

- amd64：**352 条**（编号 0–441），与上游逐条一致。
- arm64：**295 条**（编号 0–441，与 amd64 编号体系不同），与上游逐条一致。
- errno：131 基础值 + 2 别名 + 4 内部重启码（`ERestartBlock` 对应上游
  `ERESTART_RESTARTBLOCK`，值 516 不变）。

## 维护

新增 syscall 时，从上游表取号→名补入对应 `*_low.go` / `*_high.go`；
`go test ./kernel/abi/...` 会自动校验与上游一致。
