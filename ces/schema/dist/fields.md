# CES v1 字段参考

本文件由 `ces/schema/gen` 从 Go 数据模型自动生成，请勿手工编辑。

| 类型 | 字段 | Go 类型 | 必填 | 默认 | 约束 | 示例 |
| --- | --- | --- | --- | --- | --- | --- |
| Experiment | `apiVersion` | `string` | 是 | 必填 | 固定 ces.tencent.com/v1 | ces.tencent.com/v1 |
| Experiment | `kind` | `string` | 是 | 必填 | 固定 Experiment | Experiment |
| Metadata | `name` | `string` | 是 | — | 见 CES-VALID 规则清单 | — |
| Metadata | `namespace` | `string` | 否 | default | 见 CES-VALID 规则清单 | default |
| Metadata | `labels` | `map[string]string` | 否 | 省略 | 见 CES-VALID 规则清单 | — |
| Metadata | `annotations` | `map[string]string` | 否 | 省略 | 见 CES-VALID 规则清单 | — |
| Experiment | `metadata` | `schema.Metadata` | 是 | — | 见 CES-VALID 规则清单 | — |
| Target | `kind` | `string` | 是 | 必填 | 固定 Experiment | Experiment |
| Target | `selector` | `map[string]string` | 是 | — | 见 CES-VALID 规则清单 | — |
| ExperimentSpec | `target` | `schema.Target` | 是 | — | 见 CES-VALID 规则清单 | — |
| SteadyState | `probeRef` | `string` | 是 | — | 见 CES-VALID 规则清单 | — |
| Expectation | `type` | `string` | 是 | — | 枚举见对应类型说明 | value |
| Expectation | `value` | `string` | 否 | 省略 | 见 CES-VALID 规则清单 | — |
| Expectation | `min` | `string` | 否 | 省略 | 见 CES-VALID 规则清单 | — |
| Expectation | `max` | `string` | 否 | 省略 | 见 CES-VALID 规则清单 | — |
| Expectation | `predicate` | `string` | 否 | 省略 | 见 CES-VALID 规则清单 | — |
| SteadyState | `expectation` | `schema.Expectation` | 是 | — | 见 CES-VALID 规则清单 | — |
| Tolerance | `type` | `string` | 是 | — | 枚举见对应类型说明 | value |
| Tolerance | `value` | `float64` | 是 | — | 见 CES-VALID 规则清单 | — |
| SteadyState | `tolerance` | `schema.Tolerance` | 是 | — | 见 CES-VALID 规则清单 | — |
| SteadyState | `holdFor` | `schema.VirtualDuration` | 是 | — | 整数纳秒时长，须 > 0 | 5s |
| ExperimentSpec | `steadyState` | `[]schema.SteadyState` | 是 | — | 见 CES-VALID 规则清单 | — |
| Method | `expression` | `string` | 是 | — | 见 CES-VALID 规则清单 | — |
| OperatorDef | `id` | `string` | 是 | — | 见 CES-VALID 规则清单 | — |
| OperatorDef | `family` | `string` | 是 | — | 见 CES-VALID 规则清单 | — |
| OperatorDef | `params` | `map[string]string` | 否 | 省略 | 见 CES-VALID 规则清单 | — |
| OperatorDef | `targets` | `[]string` | 是 | — | 见 CES-VALID 规则清单 | — |
| OperatorDef | `invertible` | `bool` | 是 | — | 见 CES-VALID 规则清单 | — |
| OperatorDef | `compensation` | `string` | 否 | 省略 | 见 CES-VALID 规则清单 | — |
| OperatorDef | `targetKind` | `string` | 否 | 省略 | 见 CES-VALID 规则清单 | — |
| Method | `operators` | `[]schema.OperatorDef` | 否 | 省略 | 见 CES-VALID 规则清单 | — |
| ExperimentSpec | `method` | `schema.Method` | 是 | — | 见 CES-VALID 规则清单 | — |
| Probe | `id` | `string` | 是 | — | 见 CES-VALID 规则清单 | — |
| Probe | `kind` | `schema.ProbeKind` | 是 | 必填 | 固定 Experiment | Experiment |
| ProbeSpec | `url` | `string` | 否 | 省略 | 见 CES-VALID 规则清单 | — |
| ProbeSpec | `method` | `string` | 否 | 省略 | 见 CES-VALID 规则清单 | — |
| ProbeSpec | `headers` | `map[string]string` | 否 | 省略 | 见 CES-VALID 规则清单 | — |
| ProbeSpec | `expectStatus` | `int` | 否 | 省略 | 见 CES-VALID 规则清单 | — |
| ProbeSpec | `command` | `string` | 否 | 省略 | 见 CES-VALID 规则清单 | — |
| ProbeSpec | `timeout` | `schema.VirtualDuration` | 否 | 省略 | 整数纳秒时长，须 > 0 | 5s |
| ProbeSpec | `expectExitCode` | `int` | 否 | 省略 | 见 CES-VALID 规则清单 | — |
| ProbeSpec | `query` | `string` | 否 | 省略 | 见 CES-VALID 规则清单 | — |
| ProbeSpec | `expect` | `string` | 否 | 省略 | 见 CES-VALID 规则清单 | — |
| ProbeSpec | `promql` | `string` | 否 | 省略 | 见 CES-VALID 规则清单 | — |
| ProbeSpec | `range` | `schema.VirtualDuration` | 否 | 省略 | 见 CES-VALID 规则清单 | — |
| Probe | `spec` | `schema.ProbeSpec` | 是 | — | 见 CES-VALID 规则清单 | — |
| Probe | `interval` | `schema.VirtualDuration` | 是 | — | 整数纳秒时长，须 > 0 | 5s |
| ExperimentSpec | `probes` | `[]schema.Probe` | 否 | 省略 | 见 CES-VALID 规则清单 | — |
| Rollback | `strategy` | `schema.Strategy` | 否 | auto | auto/manual/timeout | auto |
| Rollback | `timeout` | `schema.VirtualDuration` | 否 | 省略 | 整数纳秒时长，须 > 0 | 5s |
| Rollback | `verifySteadyState` | `bool` | 是 | true | 见 CES-VALID 规则清单 | true |
| ExperimentSpec | `rollback` | `schema.Rollback` | 是 | — | 见 CES-VALID 规则清单 | — |
| OnceSpec | `at` | `schema.VirtualDuration` | 是 | — | 整数纳秒时长，须 > 0 | 5s |
| Schedule | `once` | `*schema.OnceSpec` | 否 | 省略 | 见 CES-VALID 规则清单 | — |
| CronSpec | `expression` | `string` | 是 | — | 见 CES-VALID 规则清单 | — |
| CronSpec | `timezone` | `string` | 否 | 省略 | 见 CES-VALID 规则清单 | — |
| Schedule | `cron` | `*schema.CronSpec` | 否 | 省略 | 见 CES-VALID 规则清单 | — |
| ContinuousSpec | `duration` | `schema.VirtualDuration` | 是 | — | 整数纳秒时长，须 > 0 | 5s |
| Schedule | `continuous` | `*schema.ContinuousSpec` | 否 | 省略 | 见 CES-VALID 规则清单 | — |
| ExperimentSpec | `schedule` | `schema.Schedule` | 是 | — | 见 CES-VALID 规则清单 | — |
| ExperimentSpec | `timeout` | `schema.VirtualDuration` | 否 | 省略 | 整数纳秒时长，须 > 0 | 5s |
| ExperimentSpec | `seed` | `uint64` | 否 | CLI 生成 | uint64 确定性种子 | 12345 |
| Experiment | `spec` | `schema.ExperimentSpec` | 是 | — | 见 CES-VALID 规则清单 | — |
