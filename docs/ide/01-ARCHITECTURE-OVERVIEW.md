# Abate IDE —— 架构总览

> Copyright (C) 2026 小黎的部落（www.qingdujian.com）
> AGPLv3; 商业用途须事先获得小黎的部落书面授权

## 1. 产品定位

| 维度 | 定义 |
|------|------|
| 产品形态 | Android 端原生应用，无服务器 |
| 核心理念 | 用户不写代码，Agent 团队全自动开发 |
| 模型支持 | 用户自定义 API 模型 + 用户本地模型上传加载 |
| 质量保障 | 内置 Abate 混沌引擎 + 30 个测试插件，dogfooding 自测 |
| 目标 | 国际超一线水准 |

## 2. 核心理念

### 2.1 全自动 Agent 团队

用户不写一行代码。用户描述需求，Agent 团队按公司制组织架构自动完成：
- CEO Agent 解析需求、分配资源
- CTO Agent 技术决策
- 工程部 6 个 Agent 并行开发
- QA 部 4 个 Agent 测试验证
- 安全部 3 个 Agent 安全审计
- 运维部 3 个 Agent 构建发布

共 23 个 Agent 角色，按部门层级协作。

### 2.2 双模型架构

- API 模型：OpenAI / Claude / Gemini / 自定义端点，适合复杂推理
- 本地模型：ONNX 量化 / GGUF / 用户上传，适合离线和隐私场景
- 智能路由：按任务类型、成本、网络状态自动选择最优模型
- 用户可为每个 Agent 角色单独指定模型

### 2.3 内置质量引擎

Abate 混沌测试系统直接内置到 IDE 中：
- 6 类故障注入（延迟/崩溃/分区/资源耗尽/数据损坏/时钟偏移）
- 30 个测试插件全覆盖（方法论 + 工程验证 + 静态巡检）
- 120 条 CES 规范作为代码质量基线
- TLA+ 形式化模型验证 Agent 编排器状态机
- IDE 开发的每个应用都经过混沌测试验证，确保质量

## 3. 系统全景图

用户交互层 (Jetpack Compose UI): 需求输入 | 团队面板 | 模型配置 | 项目管理 | 控制台
Agent 编排引擎: Agent Runtime (23角色) | Tool Registry + 权限矩阵 | Model Layer (API+本地+路由) | Abate 混沌引擎 (内置)
数据层 (Room + DataStore): 项目存储 | 会话历史 | 模型配置 | 审计日志 | 缓存
安全层 (Keystore + WASM Sandbox): API Key 加密 | 代码沙箱 | 权限强制 | 零遥测

## 4. 与 Abate 的关系

| 维度 | 说明 |
|------|------|
| 混沌引擎 | injector/probe/reporter 直接移植到 Android 端 |
| 测试体系 | 30 个测试插件全部内置，适配 Android 场景 |
| CES 规范 | 120 条规范作为 Agent 生成代码的质量基线 |
| TLA+ 模型 | Agent 编排器状态机可形式化验证 |
| Dogfooding | Abate IDE 用 Abate 引擎测自己，也测它产出的应用 |

## 5. 技术栈概要

| 层 | 技术 |
|----|------|
| 语言 | Kotlin |
| UI | Jetpack Compose + Material 3 |
| 本地推理 | ONNX Runtime + MLC-LLM |
| API 推理 | OkHttp + SSE 流式 |
| 代码沙箱 | WASM (wazero 移植) |
| 存储 | Room (SQLite) + DataStore |
| DI | Hilt |
| 并发 | Coroutines + Flow |
| 序列化 | kotlinx.serialization |