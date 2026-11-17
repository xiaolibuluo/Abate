<!--
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
-->

# CES v1 规范文本（CES-SCHEMA）

> 对应 PLAN-03 / PLAN-08 XCU-DOC。CES（Chaos Experiment Specification）是
> Abate 的混沌实验声明规范，版本 `ces.tencent.com/v1`。本文本与
> `ces/schema/dist/ces-v1.schema.json`（JSON Schema）和 `fields.md` 为同源
> 单一事实来源，CI 由 `scripts/drift-check.sh` 校验零漂移。

## 1. 文档信封（Envelope）

每个 CES 文档是顶层对象，含固定信封 + `spec` 可执行体。未知字段**一律拒绝**
（fail-fast，避免静默忽略拼写错误）。

| 字段 | 类型 | 必填 | 说明 |
|---|---|---|---|
| `apiVersion` | string | ✅ | 固定 `ces.tencent.com/v1` |
| `kind` | string | ✅ | 固定 `Experiment` |
| `metadata.name` | string | ✅ | DNS 子域名；命名空间默认 `default` |
| `metadata.namespace` | string | ⬜ | 默认 `default` |
| `spec` | object | ✅ | 可执行体（见下） |

## 2. spec 可执行体

| 字段 | 必填 | 说明 |
|---|---|---|
| `target` | ✅ | 扰动目标：`kind` ∈ {`simulator`,`source`,`external`}；`external` 为保留字段，运行时未支持，**校验拒绝**。`selector` 为稳定 key/value（package/function/node/process）。 |
| `steadyState` | ✅ | 至少一个稳态假设（hypothesis）。每项引用一个 probe，含 `expectation`/`tolerance`/`holdFor`。 |
| `method` | ✅ | 扰动代数表达式（见 §3）。 |
| `probes` | ⬜ | 观测源列表；`id` 在文档内唯一，按 ID 顺序注册（顺序不影响观测流）。 |
| `rollback` | ✅ | 补偿策略：`strategy` ∈ {`auto`,`manual`,`timeout`}（归一化默认 `auto`）；`timeout` 策略须填 `timeout`。 |
| `schedule` | ✅ | 互斥三选一：`once` / `cron` / `continuous`（指针互斥，见 CES-VALID-SCHEDULE-001）。 |
| `timeout` | ⬜ | 可选；归一化填默认。 |
| `seed` | ⬜ | 可选；归一化派生确定性种子。 |

## 3. 扰动代数（Method Expression）

`method.expression` 使用 CES 算子微积分记号：

| 符号 | 含义 |
|---|---|
| `∘` | 复合（composition） |
| `⊗` | 并发积（concurrent product） |
| `a⁻¹` | 对偶（dual / 逆补偿） |
| `ε` | 恒等（identity） |
| `( )` | 分组 |

- 操作数为 `operator://` 引用（解析 `method.operators` 内联定义或引用）。
- 深度上限 **64**（CES-PARSE/METHOD-003 攻击面，防超深嵌套注入）。
- 非可逆算子必须声明 `compensation`，否则 CES-VALID-METHOD-002 拒绝。

## 4. 校验规则目录（CES-VALID-*，共 31 条）

按族分组；每条规则可报告 `error` 或 `warning` 严重级（`--strict` 将 warning 提升为失败）。

| 族 | 规则 | 关注点 |
|---|---|---|
| STEADY | 001–005 | 稳态假设完整性、探针引用存在、期望值/容差类型矩阵、holdFor 合法性 |
| ROLLBACK | 001–003 | 策略合法性、timeout 策略须填 timeout、补偿一致性 |
| SCHEDULE | 001–004 | 三模式互斥、cron 表达式合法、once/continuous 字段合法性 |
| METHOD | 001–004 | 表达式语法、对偶补偿（可逆性）、深度上限(003)、算子解析 |
| PROBE | 001–005 | 探针 id 唯一、kind 与字段匹配、采样间隔为正、命令/URL 安全 |
| TARGET | 001–004 | 目标 kind 合法、selector 键合法、`external` 拒绝、域匹配 |

## 5. 解析规则目录（CES-PARSE-*，共 13 条）

解析/归一化阶段的结构与信封校验（如信封必填、版本协商、YAML/JSON 解码、
未知字段拒绝、字段类型合法性）。归一化负责填默认（namespace/strategy/seed/
timeout）并派生确定性种子。

## 6. 一致性（TCK）

`ces/tck` 提供数据驱动的兼容测试套件（164 用例：parse 68 / validate 46 /
roundtrip 50）。`ces/tck/example` 是参考第三方实现，证明其消费中性用例格式
并全量合规。第三方实现只需实现 `tck.Driver` 接口（Parse/Validate），即可用
同一套用例认证合规。
