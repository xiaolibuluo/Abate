# 深度集成算法蓝本 — 只读参考基线（Abate 供应链合规台账）

> 本目录是**只读算法参考**，用于「深度改进优化 + 自研彻底打磨」：吃透上游算法模型后
> 用纯 Go 标准库在 Abate 内核中自研重写，并做确定性化 / 代数化 / 对偶补偿等深度改进。
> **不 import、不 vendor 进 go.mod、不编译、不子进程调用**。文件存 `.go.txt` 且置于 `_` 前缀目录，
> Go 工具链与全部门禁均跳过；上游版权头原样保留。

## 锁定事实（2026-11-09 抓取；GitHub 匿名 REST 遇 403 限流，未取得项如实标注）

| 上游项目 | 许可证 | commit 快照 | 分支 |
|---|---|---|---|
| Shopify/toxiproxy | MIT | 49299d8838525126e09c85ebe0ba89940a3261e5 | main |
| go-gremlins/gremlins | Apache-2.0 | b48a4aad17eff33dc7262e0d1487b1a4d6eec322 | main |
| google/go-cmp | BSD-3-Clause(清单标Apache-2.0存疑,以LICENSE为准) | 5deca1c7dab7d33c102d8d7a26794da148485563 | master |
| panjf2005/ants | MIT | 未取得(API限流) | - |
| pingcap/failpoint | Apache-2.0 | 55ac33a48e3b10f21fdb30e9d29368e98bcd70fc | master |
| spf13/cobra | Apache-2.0 | adbc8813901bba65827259daa8e22ff94ec1f30e | main |

## 已抓取算法文件（文件级 SHA-256 前16位锁，可复核）

| 上游项目 | 文件 | 算法用途 | 字节 | SHA-256(前16) |
|---|---|---|---|---|
| Shopify/toxiproxy | `toxics__latency.go.txt` | 延迟/抖动 toxic 模型 | 1096 | `a85cd31f54de1d59` |
| Shopify/toxiproxy | `toxics__bandwidth.go.txt` | 令牌桶带宽限流 | 2701 | `c34d6fe8402d22c5` |
| Shopify/toxiproxy | `toxics__slicer.go.txt` | 包切片/限带宽调度 | 2559 | `ff993f521b6bd720` |
| google/go-cmp | `cmp__compare.go.txt` | 深度比较/差异报告算法 | 22812 | `547d779d6ecb7e9a` |
| go-gremlins/gremlins | `internal__mutator__mutator.go.txt` | Go 变异算子注册模型 | 4746 | `4cf48e79a490e12f` |
| spf13/cobra | `command.go.txt` | 命令树/参数解析设计 | 61142 | `f79613721ad04c9b` |

## 深度改进方向（自研重写时的优化点，不是照搬）

- **确定性化**：上游所有随机/时间/并发序改为 SIM-RNG 种子驱动 + 虚拟时钟 + 全序裁决。
- **代数化统一**：故障/变异统一实现 `kernel/algebra.Operator`，带对偶补偿与确定性指纹。
- **零侵入**：录制/观测只读订阅，不改动被观测事件流。
- **裁剪**：只取核心算法线，去掉上游 CLI/网络/报告/JVM/框架胶水。
- **边界硬化**：上游未覆盖的溢出/空输入/越权路径补显式分支与拒绝。
