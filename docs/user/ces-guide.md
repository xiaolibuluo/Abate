<!--
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
-->

# CES 编写指南（User CES Guide）

> 配套规范：`docs/spec/ces-v1.md`。所有示例均可通过 `abate run --dry-run` 校验。

## 信封铁律

- `apiVersion` 必须是 `ces.tencent.com/v1`，`kind` 必须是 `Experiment`。
- 未知字段**直接拒绝**（fail-fast），不要依赖默认值兜底。
- `metadata.name` 为 DNS 子域名；`namespace` 缺省 `default`。

## 目标（target）

| kind | 含义 | 状态 |
|---|---|---|
| `simulator` | 模拟器内目标（默认，确定性） | ✅ 支持 |
| `source` | 源码级目标（语言前端） | ✅ 支持 |
| `external` | 外部系统 | ❌ 保留，校验拒绝 |

`selector` 为稳定键值对：`node` / `process` / `package` / `function`。

## 稳态（steadyState）

至少一项假设；每项结构：

```yaml
- probeRef: p1                       # 引用 probes[].id
  expectation: { type: value, value: ok }   # value | range | predicate
  tolerance:   { type: absolute, value: 1 } # absolute | relative | percent
  holdFor: 2s                        # 虚拟时长，须为正
```

`expectation.type` 取值：`value`（用 `value`）、`range`（用 `min`/`max`）、
`predicate`（用 `predicate`）。`tolerance.type` 取值：`absolute`/`relative`/
`percent`，`value` 非负。

## 方法（method.expression）

扰动代数（详见规范 §3）：`∘` 复合、`⊗` 并发积、`a⁻¹` 对偶、`ε` 恒等、括号分组。
嵌套深度上限 **64**。非可逆算子须声明 `compensation`。

```yaml
method:
  expression: "kill ∘ restart⁻¹"
  operators:
    - id: kill
      family: chaos
      targets: [n1]
      invertible: false
      compensation: restart
```

## 探针（probes）

`id` 文档内唯一；`kind` ∈ {`http`,`shell`,`metric`,`prom`}，各自字段不同：

```yaml
- id: p1
  kind: metric
  spec: { query: up, expect: "1" }
  interval: 1s
- id: p2
  kind: http
  spec: { url: http://svc/health, method: GET, expectStatus: 200 }
  interval: 5s
```

采样 `interval` 须为正；`shell` 命令经安全白名单校验（只读/安全命令）。

## 回滚（rollback）

`strategy` ∈ {`auto`,`manual`,`timeout`}；`timeout` 策略必须填 `timeout`。
`verifySteadyState: true` 会在对偶算子应用后重跑稳态断言。

## 调度（schedule）

三选一（互斥）：

```yaml
schedule: { continuous: { duration: 30s } }   # 持续扰动
schedule: { once: { at: 2026-01-01T00:00:00Z } }  # 单次
schedule: { cron: { expr: "0 * * * *" } }     # 周期
```

## 常见校验错误速查

| 现象 | 规则 | 修复 |
|---|---|---|
| 表达式语法错 | CES-VALID-METHOD-001 | 检查算子/括号/符号 |
| 非可逆算子无补偿 | CES-VALID-METHOD-002 | 加 `compensation` 或 `invertible: true` |
| 表达式超深嵌套 | CES-VALID-METHOD-003 | 深度 ≤ 64 |
| 调度多模式并存 | CES-VALID-SCHEDULE-001 | 只保留一个 |
| external 目标 | CES-VALID-TARGET-003 | 改用 simulator/source |
