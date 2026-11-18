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
abate 1.2.27
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

## 7. 开箱即用能力边界声明（Open-Box Readiness）

> 实测基准：2026-11-19，将源码备份解压至全新空目录（不含 `.git` / `dist` / `.artifacts`），
> 以"刚拿到项目的开发者"身份逐项执行。结论如实分层，不夸大、不外推。

### 7.1 核心链路：真正开箱即用

| 项 | 实测结果 |
|---|---|
| 依赖 | 零外部依赖，`GOPROXY=off` 离线 `go build` 一次成功 |
| 构建产物 | 单个约 16 MB 二进制，无需生成代码或下载工具链 |
| 版本 | `./abate version` → `abate 1.2.27` |
| 门禁 | header / style / hash / watermark / drift 五项全 PASS |
| 测试 | `go test ./...` 全量通过（73 包） |
| 功能命令 | `derive` 正常；`check` 按扩展名分派 7 语言（`.go/.py/.js/.c/.cs/.rs/.kt`）并检出规则违背；`run` 真实执行 CES 并产出报告 + 摘要；错误输入给出带规则号的精确解析错误 |

### 7.2 三项非开箱即用边界（必须知悉）

1. **`dist/` 预编译二进制不在仓库**。该目录被 `.gitignore` 排除，由 `release.sh --dry-run` 可重建；
   现成 5 平台二进制（linux / darwin / windows × amd64 / arm64）请从
   [GitHub Release v1.2.27](https://github.com/xiaolibuluo/Abate/releases/tag/v1.2.27) 下载。
   从源码 `go build` 不受影响。

2. **高级验证需额外环境，非 clone 即可全跑**：
   - TLA+ 形式化验证需 Java + `tla2tools.jar`；
   - libclang / C++ 功能级验证需对应工具链（`.artifacts/` 下约 776 MB 构建产物平台绑定、按
     `.gitignore` 不入库，可由脚本重建）；
   - E2E 全语料回归需约 264 MB `testdata/` 语料。
   核心 CLI 与五项门禁均不需要这些额外环境。

3. **`check` 对本项目自身源码会报 info / error 级风格提示**（实测 `check main.go` 检出 `GO-CRC-029`）。
   这是规则引擎正常工作的证据，不是故障；因此"check 自身源码全绿"不作为验收标准，
   验收以 `scripts/lint-gate.sh` 冻结的那套规则为准。

### 7.3 结论

作为 Go 命令行工具与确定性混沌实验引擎，Abate **开箱即用**：clone / 解压后
`go build && ./abate` 即可使用。作为包含 TLC / clang / 全语料 E2E 的完整质量体系，
需按 §7.2 准备额外工具链。

> 发布与验证的完整实证见 `docs/i18n/RELEASE-READINESS-FINAL.md` §十一、§十二。
