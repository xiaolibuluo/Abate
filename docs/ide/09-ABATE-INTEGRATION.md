# Abate IDE —— Abate 集成方案

> Copyright (C) 2026 小黎的部落（www.qingdujian.com）
> AGPLv3; 商业用途须事先获得小黎的部落书面授权

## 1. 集成策略

Abate 混沌测试系统直接内置到 IDE 中，不作为外部依赖。

| 维度 | 说明 |
|------|------|
| 引擎移植 | injector/probe/reporter 从 Go 移植到 Kotlin |
| 插件内置 | 30 个测试插件全部内置，适配 Android 场景 |
| CES 规范 | 120 条规范作为 Agent 生成代码的质量基线 |
| TLA+ 模型 | Agent 编排器状态机可形式化验证 |
| Dogfooding | IDE 用 Abate 测自己，也测它产出的应用 |

## 2. 故障注入器（6 类）

| 类型 | 原始 | Android 适配 |
|------|------|-------------|
| F1 Latency | 网络延迟 | 主线程延迟、ANR 触发 |
| F2 Crash | 进程崩溃 | Activity 崩溃、Service 崩溃 |
| F3 Partition | 网络分区 | 模型 API 断连、离线模拟 |
| F4 Exhaustion | 资源耗尽 | 内存耗尽(OOM)、磁盘满 |
| F5 Corruption | 数据损坏 | Room 数据损坏、文件损坏 |
| F6 ClockSkew | 时钟偏移 | 系统时间偏移、定时器错乱 |

## 3. 30 个测试插件映射

### 3.1 测试方法论（5）

| 插件 | 原始功能 | Android 适配 |
|------|----------|-------------|
| compatibility | go vet 兼容性 | API 级别兼容性检查 |
| perf_baseline | 性能基线 | 启动时间/帧率基线 |
| backpressure | 背压测试 | 快速切换 Activity |
| stress | 压力测试 | 并发操作压力 |
| soak | 长时间运行 | 长时间稳定性 |

### 3.2 健壮安全（8）

| 插件 | 原始功能 | Android 适配 |
|------|----------|-------------|
| fuzz | 模糊测试 | 输入模糊测试 |
| conflict | 冲突检测 | 并发冲突检测 |
| defect | 缺陷检测 | 已知缺陷模式 |
| crash | 崩溃测试 | 崩溃恢复验证 |
| vuln | CVE 扫描 | 依赖漏洞扫描 |
| security | 安全测试 | 权限安全验证 |
| adversarial | 对抗测试 | 恶意输入测试 |
| fault_tolerance | 容错测试 | 故障容错验证 |

### 3.3 工程验证（6）

| 插件 | 原始功能 | Android 适配 |
|------|----------|-------------|
| lint_intensive | Lint 检查 | Android Lint 强化 |
| static_analysis | 静态分析 | 静态代码分析 |
| dynamic_analysis | 动态分析 | 运行时分析 |
| cicd_check | CI/CD 检查 | 本地流水线检查 |
| dual_build_verify | 双构建验证 | Debug+Release 构建 |
| dual_mode_run | 双模式运行 | 多设备运行验证 |

### 3.4 测试类型（4）

| 插件 | 原始功能 | Android 适配 |
|------|----------|-------------|
| unit | 单元测试 | JUnit 单元测试 |
| integration | 集成测试 | 模块集成测试 |
| property | 属性测试 | 属性随机测试 |
| concurrency | 并发测试 | Coroutines 并发测试 |

### 3.5 静态巡检（7）

| 插件 | 原始功能 | Android 适配 |
|------|----------|-------------|
| placeholder_check | TODO/FIXME 检测 | 同 |
| stub_check | 桩函数检测 | 同 |
| fake_impl_check | 假实现检测 | 同 |
| logic_check | 逻辑缺陷检测 | 同 |
| file_header_check | 文件头检查 | 同 |
| function_check | 函数复杂度 | 同 |
| vuln_check | 静态安全扫描 | 硬编码密钥/SQL 注入检测 |

## 4. 质量基线

- Agent 生成的每份代码自动经过静态巡检 7 插件检查
- 构建后自动经过工程验证 6 插件检查
- 发布前强制经过健壮安全 8 插件 + 混沌注入验证
- 120 条 CES 规范作为代码审查基线
- 不达标的代码自动退回工程部修复

## 5. Dogfooding

- Abate IDE 自身开发过程使用 Abate 引擎测试
- IDE 产出的每个应用都经过混沌测试
- 测试结果反馈给 Agent 团队持续改进