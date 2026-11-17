# ADR-XSYS-001：消除 x/sys 间接依赖

- 状态：已采纳（Accepted）
- 日期：2026-11-16
- 决策来源：高见远（架构师）`docs/refactor/DESIGN.md` §2.3
- 落盘：寇豆码（工程师）
- 关联：TODO P1-04 / P1-05、ADR-WASM-002

> 本文记录架构师已作出的技术决策，不含手写签名；任何法律结论均不在本文范围。

---

## 1. 事实

| 项 | 实测 |
|---|---|
| `go.mod` 声明 | `golang.org/x/sys v0.44.0 // indirect` |
| 引入路径 | 由 wazero v1.12.0 间接引入 |
| 产品代码直接 import | **0 命中**（扫描 `cmd/ plugin/ kernel/ lang/ internal/ ctl/ ces/ sim/ security/`） |
| vendored 副本 | `vendor/golang.org/x/sys/`，随 vendor 目录处置 |

## 2. 决策

**不写替代代码。** 删除 wazero 后执行 `go mod tidy`，由工具链自动清除该 indirect 条目。

## 3. 残留兜底（实现阶段核实）

若 tidy 后仍存在：

1. 逐一检查使用点，优先改用标准库 `os` / `syscall` / `os/user`；
2. 确缺的常量以本仓自有常量定义（syscall 编号属 OS ABI 事实数据，⚖️ L05 需法务确认）；
3. **禁止从 `third_party/` 拷贝任何片段**。

## 4. 验收

```bash
GOFLAGS=-mod=mod go list -m all    # 只应输出 abate
```

配合 DESIGN §2.4 的三依赖清零顺序：产品代码 import 扫描 → `go mod tidy` → 确认 `go.mod` 无 require →
最后删 `vendor/` → 改 `-mod=mod` → `go build ./... && go test -race -count=1 ./...`。

## 5. 后果

- 正向：依赖面从 3 个模块降为 0；供应链审计面归零。
- 反向：若未来引入需要平台常量的功能，须以本仓自有定义实现并再次走 ADR 流程。
