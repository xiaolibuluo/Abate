<!--
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
-->

# 模块开发指南（Module Dev Guide）

> 模块地图与依赖图见 `docs/plan/00-MASTER-PLAN.md`。本指南补充开发 DoD 与分层铁律。

## 1. 分层（L0–L7）

| 层 | 模块举例 | 约束 |
|---|---|---|
| L0 内核 | KRN-ALG/SCH/OBS/VER/MUT/INJ | 仅依赖 Go 标准库 + 源码级并入的自有代码；**十年不大更新** |
| L1 WASM | WAS-HOST/CAP/HOT/SIG | 插件进入内核的唯一通道（端口-适配器） |
| L2 CES | CES-SCHEMA/PARSE/VALID/TCK | 规范层，与 JSON Schema / fields.md 同源 |
| L3 模拟器 | SIM-CLK/RNG/NET/FS/SCHED/REC | 确定性虚拟环境 |
| L4 语言前端 | LNG-AIR/GO/PY/JS/.../RULE/CHK/MUT | 解析复用，规则共享 AIR 层 |
| L5/L6 适配器/报告 | ADP-*/RPT-* | 经 WAS-HOST 加载 |
| L7 控制面 | CTL-CLI/REST/WEB | 入口 |

## 2. 依赖铁律

- 只允许自上而下、自外而内；**L0 不 import L1–L7**。
- 禁止循环依赖（CI depguard 强制）。
- L0 仅依赖标准库与自有源码（十年内核铁律）。

## 3. 模块 DoD（Definition of Done）

1. 合入 `main` 且五段门禁全绿；无 TODO/FIXME；版权头 + 水印齐全。
2. 测试：单元（表驱动，三路径）+ 集成 + 确定性回放（同种子双跑哈希一致）。
   覆盖率：内核 ≥95%/85%，生态 ≥85%/75%。
3. 接口：对外 API 有 Go doc 注释（godot 强制句号结尾）；破坏性变更走 ADR + 2 版本废弃窗口。
4. 文档：模块 README（职责/接口/验收清单）+ CHANGELOG 条目。
5. 性能：热路径附基准测试且不超性能预算（见 `bench/budget.md`，15 项指标）。
6. 安全：新依赖源码级并入并记录来源 commit 与许可证审查结论。

## 4. 确定性开发清单

- 不使用 `time.Now` / `time.Date` 读取真实时钟；用 `sim/clock` 虚拟时钟。
- 不使用全局 `math/rand`；用 `sim/rand` 种子驱动 PRNG。
- 不依赖 map 迭代顺序；排序后再比较/输出。
- 不读取 `os.Environ` 决定行为；配置显式传入。
- 内核单线程虚拟时钟驱动；并发只存在于模拟器边界，事件序号全序化。

## 5. 文件与命名

- 每个文件 ≤300 行；超长则拆包。
- 禁止 TODO/FIXME/godox 命中。
- 新包首部逐字节复用版权头（4 行注释 + 空行 + package 注释）。
- 验证只跑自己包：`go test -race -count=1 ./<pkg>`（离线 `-mod=vendor`）。
