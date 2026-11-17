# 零第三方源码审计报告（骨架 · T05 定稿）

> 文档编号：ABATE-REFACTOR-ZERO-AUDIT-001　生成：2026-11-16（T01 骨架）
> 三口径依据：`docs/refactor/DESIGN.md` §1.1；审计脚本：`scripts/audit-zero-third-party.sh`（T01 新建）。
> **本报告在 T05 定稿，且必须附脚本原始输出；不以声明代替证据。**

---

## 1. 三口径与判据

| 口径 | 可验证命令 | 通过标准 | 当前状态 |
|---|---|---|---|
| C1 依赖面为零 | `GOFLAGS=-mod=mod go list -m all` | 仅 `abate` 自身 | 未达成（T04 清零） |
| C2 源码副本为零 | `bash scripts/audit-zero-third-party.sh` | 报告结论 0，附原始输出 | 未达成（T02/T04 清除） |
| C3 无隐藏运行依赖 | `go test -race -count=1 ./...` | 全绿，且不读已删路径 | 未达成（T01 护栏后 T02 验证） |

---

## 2. 清除前非空跑（证明审计脚本不是空跑）

T01 已完成审计脚本落地。清除前的实测输出将在本节原样附上：

- 执行时点：T01 出口（第三方尚未物理删除）
- 期望结论：`[VERDICT] 第三方源码副本 = N 项（未清空）`，N > 0
- 原始输出：`docs/refactor/audit-pre-cleanup.log`（82 行，`exit=1`，实测保存，未手工编辑）；
  结论行与交叉校验见「附录 A」。定稿时的清除后输出另行粘贴，禁止补造

## 3. 清除后终审（T05 回填）

- C1 原始输出：
- C2 原始输出：
- C3 原始输出：
- 结论：第三方源码副本 = **0**

## 4. 门禁联动状态（T05 回填）

| 门禁 | 改造后形态 | 状态 |
|---|---|---|
| `pit-parity.sh` | 自研变异语义 golden 回归（非官方对拍） | 待 T05 |
| `m5-reference-scan.sh` | 零蓝本 + 零依赖断言 | 待 T05 |
| `license-check.sh` | 历史台账存在且无当前接入条目 | 待 T05 |
| `kernel/abi` 测试 | 自研冻结 golden（T01 已完成） | 待 T02 删除后复验 |

## 附录 A — 原始输出

（T05 终审输出由定稿时粘贴实际终端输出，不手工编造。）

### A.0 T01 ABI golden 一次性交叉校验（上游文件尚存时执行）

删除前用 `third_party/_sources/gvisor/pkg/sentry/syscalls/linux/linux64.go` 对自研冻结 golden
做逐条交叉校验；原始输出：

```
upstream markers: 29 405 717
amd64: upstream=352 golden=352 identical=True
arm64: upstream=295 golden=295 identical=True
CROSSCHECK_RESULT= PASS
```

口径：golden 仅含 OS ABI 事实数据（number → name），不含任何上游源码片段；
校验完成后 `kernel/abi` 测试不再读取 `third_party/` 任何路径。
⚖️ syscall 表是否属不受版权保护的事实数据，待法务认定（TODO L05）。
