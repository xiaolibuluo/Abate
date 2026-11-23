<!--
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
-->

# 快速上手（User Quickstart）

本指南让你在 5 分钟内构建 Abate 并跑通第一个确定性混沌实验。所有命令均可
直接复制执行。

## 1. 构建

项目**零外部依赖**（`go.mod` 仅声明模块名与 Go 版本，`go.sum` 为空、无 `vendor/`），
Go 1.26.11+ 在离线环境即可构建：

```bash
export GOPROXY=off
go build -o abate .
./abate version
```

`abate version` 输出构建溯源（版本 / 规范 / 许可证 / 水印）：

```
abate 1.2.29
schema: ces.tencent.com/v1
build: deterministic
license: AGPLv3 or commercial
source: https://github.com/xiaolibuluo/Abate
watermark: 小黎的部落
```

## 2. 写一份 CES 实验

新建 `hello.yaml`：

```yaml
apiVersion: ces.tencent.com/v1
kind: Experiment
metadata:
  name: hello-ces
spec:
  target:
    kind: simulator
    selector: { node: n1 }
  steadyState:
    - probeRef: p1
      expectation: { type: value, value: ok }
      tolerance: { type: absolute, value: 1 }
      holdFor: 2s
  method: { expression: ε }
  probes:
    - id: p1
      kind: metric
      spec: { query: up, expect: "1" }
      interval: 1s
  rollback: { strategy: auto, verifySteadyState: false }
  schedule: { continuous: { duration: 1s } }
```

## 3. 运行

```bash
./abate run hello.yaml
```

加 `--format json --out report.json` 输出机器可读报告；加 `--record /tmp/run.arch`
写出确定性 SIM-REC 档案。

## 4. 验证确定性（核心卖点）

同种子同配置，任意次运行结果**逐字节一致**：

```bash
./abate run hello.yaml --seed 42 --format json --out a.json
./abate run hello.yaml --seed 42 --format json --out b.json
./abate verify a.json b.json      # 逐字节一致则退出 0
```

## 5. 回放证据

```bash
./abate run hello.yaml --seed 42 --record /tmp/run.arch
./abate replay /tmp/run.arch       # 校验帧哈希链 + 全档 SHA-256
```

篡改档案末字节后 `replay` 会拒绝（见安全威胁 T10）。

## 6. 下一步

- 用 7 语言检查源码（保姆式教程）：见 `docs/user/multilang-check-guide.md`
- 编写更复杂的实验：见 `docs/user/ces-guide.md`
- 全部命令与退出码：见 `docs/user/cli-reference.md`
- 规范原文：见 `docs/spec/ces-v1.md`

## 7. 已知边界（如实知悉）

| 事项 | 当前事实 |
|---|---|
| 依赖与构建 | 零外部依赖，`GOPROXY=off` 离线 `go build` 成功；产物为单个静态二进制 |
| 预编译二进制 | 不在仓库（`.gitignore`），从 [Release v1.2.29](https://github.com/xiaolibuluo/Abate/releases/tag/v1.2.29) 下载 5 平台版本 |
| `check` 自身源码 | 会对本项目源码报 info/error 级风格提示（规则引擎正常工作的证据），验收以 `scripts/lint-gate.sh` 为准 |
| 高级验证环境 | TLA+ 需 Java + `tla2tools.jar`；libclang C API 验证需对应工具链；全语料 E2E 需 `testdata/` 语料。核心 CLI 与常规门禁不需要这些 |
| 能力边界 | 见[能力与边界](capabilities-and-limits.md)：容错解析非语法守门、不做类型/CFG/数据流分析、竞品与全语料保真度有未测量项 |

更多命令见 [CLI 参考](cli-reference.md)，七语言逐个手把手见[七语言教程](multilang-check-guide.md)。
