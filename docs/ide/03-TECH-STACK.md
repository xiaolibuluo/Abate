# Abate IDE —— 技术栈选型

> Copyright (C) 2026 小黎的部落（www.qingdujian.com）
> AGPLv3; 商业用途须事先获得小黎的部落书面授权

## 1. 技术栈总览

| 层 | 技术 | 版本 | 选型理由 |
|----|------|------|----------|
| 语言 | Kotlin | 2.0+ | Android 一等公民，协程原生 |
| UI | Jetpack Compose | 1.7+ | 声明式 UI，动画流畅，Material 3 |
| 本地推理 | ONNX Runtime | 1.17+ | 端侧模型推理，支持量化 |
| 本地推理 | MLC-LLM | 0.1+ | GPU 加速推理 |
| API 推理 | OkHttp | 4.12+ | SSE 流式，兼容 OpenAI/Claude |
| 代码沙箱 | WASM (wazero) | v1.12 | 纯 Go 实现，无 CGO 依赖 |
| 数据库 | Room | 2.6+ | SQLite ORM，迁移支持 |
| KV 存储 | DataStore | 1.1+ | Preferences 替代 SharedPreferences |
| DI | Hilt | 2.51+ | 标准化依赖注入 |
| 并发 | Coroutines + Flow | 1.8+ | Agent 编排天然异步 |
| 序列化 | kotlinx.serialization | 1.6+ | 轻量高效 |
| Git | JGit | 6.9+ | 本地 Git 操作 |

## 2. 工具矩阵

### 2.1 文件操作工具

| 工具 | 功能 | 可用角色 |
|------|------|----------|
| read_file | 读取文件 | 全部只读角色 |
| write_file | 写入文件 | 架构师/前端/逻辑/数据/模型/构建 |
| search_code | 正则搜索 | 架构师/审计/逻辑 |
| glob_files | 通配符查找 | 架构师/审计 |
| dependency_tree | 依赖分析 | 架构师 |

### 2.2 构建工具

| 工具 | 功能 | 可用角色 |
|------|------|----------|
| build_run | Gradle 构建 | 构建/CI-CD |
| sign_config | 签名配置 | 构建/发布 |
| variant_manage | 构建变体 | 构建 |
| cache_manage | 构建缓存 | CI-CD |

### 2.3 测试工具

| 工具 | 功能 | 可用角色 |
|------|------|----------|
| test_run | 运行测试 | 测试执行 |
| test_report | 测试报告 | 测试执行/测试设计 |
| coverage_report | 覆盖率 | 测试设计 |
| abate_inject | 混沌注入 | 混沌 |
| abate_report | 混沌报告 | 混沌 |
| profile_run | 性能分析 | 性能 |
| trace_view | Trace 查看 | 性能 |

### 2.4 安全工具

| 工具 | 功能 | 可用角色 |
|------|------|----------|
| static_scan | 静态安全扫描 | 审计 |
| permission_audit | 权限审查 | 审计 |
| sandbox_test | 沙箱测试 | 沙箱 |
| isolation_verify | 隔离验证 | 沙箱 |
| compliance_check | 合规检查 | 隐私 |
| data_flow_audit | 数据流审计 | 隐私 |

### 2.5 设计工具

| 工具 | 功能 | 可用角色 |
|------|------|----------|
| preview_render | Compose 预览 | 前端/原型 |
| layout_inspector | 布局检查 | 前端 |
| flowchart_render | 流程图 | UX |
| color_palette | 色彩方案 | UI |
| theme_generate | 主题生成 | UI |
| spec_write | 规范文档 | UX |

### 2.6 运维工具

| 工具 | 功能 | 可用角色 |
|------|------|----------|
| pipeline_run | 流水线 | CI-CD |
| version_tag | 版本标记 | 发布 |
| release_check | 发布检查 | 发布 |
| crash_report | 崩溃报告 | 监控 |
| anr_detect | ANR 检测 | 监控 |

## 3. 工具隔离原则

- 每个 Agent 只能调用权限矩阵中允许的工具
- 写操作经过审计日志记录
- 构建和签名工具仅限构建 Agent 和发布 Agent
- 安全 Agent 只读不写，保证审计独立性
- Abate 混沌注入仅限混沌 Agent 调用
- 工具调用全部在 WASM 沙箱中执行
- 跨工具数据传递通过消息总线，不共享内存