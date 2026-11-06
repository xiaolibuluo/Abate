# Abate IDE —— 模块设计

> Copyright (C) 2026 小黎的部落（www.qingdujian.com）
> AGPLv3; 商业用途须事先获得小黎的部落书面授权

## 1. 目录结构

```
abate-ide/
├── app/                          # 应用入口
│   ├── MainActivity.kt
│   └── AbateIDEApplication.kt
├── core/                         # 核心引擎
│   ├── agent/                    # Agent 运行时
│   │   ├── runtime/              # Agent 执行引擎
│   │   ├── orchestrator/         # 多 Agent 编排器
│   │   ├── registry/             # Agent 角色注册
│   │   └── message/              # Agent 间消息总线
│   ├── model/                    # 模型抽象层
│   │   ├── api/                  # API 模型适配
│   │   ├── local/                # 本地模型推理
│   │   ├── router/               # 智能路由
│   │   └── config/               # 用户模型配置
│   ├── tool/                     # 工具系统
│   │   ├── registry/             # 工具注册中心
│   │   ├── permission/           # 权限矩阵
│   │   ├── sandbox/              # WASM 沙箱
│   │   └── audit/                # 操作审计
│   ├── project/                  # 项目管理
│   │   ├── workspace/            # 工作空间
│   │   ├── git/                  # 本地 Git
│   │   └── build/                # 构建系统
│   └── chaos/                    # Abate 混沌引擎（内置移植）
│       ├── injector/             # 故障注入器
│       ├── probe/                # 探针
│       ├── reporter/             # 报告
│       └── plugins/              # 30 个测试插件
├── feature/                      # 功能模块
│   ├── editor/                   # 代码编辑器（只读展示）
│   ├── console/                  # Agent 工作台
│   ├── team/                     # 团队管理面板
│   ├── model/                    # 模型配置界面
│   ├── project/                  # 项目管理界面
│   └── settings/                 # 设置
├── data/                         # 数据层
│   ├── db/                       # Room 数据库
│   ├── store/                    # DataStore
│   └── cache/                    # 缓存
└── design/                       # 设计系统
    ├── theme/                    # 主题
    ├── components/               # 组件库
    └── icons/                    # 图标
```

## 2. 核心模块职责

### 2.1 Agent Runtime

- 加载 23 个 Agent 角色定义
- 管理 Agent 生命周期（创建/执行/销毁）
- 每个 Agent 绑定独立模型配置
- 执行 Setup -> Run -> Teardown 全生命周期

### 2.2 Orchestrator

- DAG 依赖编排（CEO -> CTO -> 各部门并行 -> COO 汇总）
- 消息总线传递跨 Agent 数据
- 失败策略：continue / abort_dependents / abort_all
- 支持子 Agent 任务委派

### 2.3 Model Router

- 按任务类型匹配模型（架构决策->API大模型，格式化->本地小模型）
- 网络状态感知（离线降级到本地模型）
- 成本控制（Token 计数 + 预算限制）
- 用户可手动覆盖路由策略

### 2.4 Tool Registry + Permission

- 工具注册中心管理全部 30+ 工具
- 权限矩阵强制校验每个 Agent 的工具调用
- 审计日志记录所有写操作
- WASM 沙箱隔离工具执行

### 2.5 Chaos Engine（内置 Abate）

- 6 类故障注入器：Latency/Crash/Partition/Exhaustion/Corruption/ClockSkew
- 30 个测试插件全部内置
- 120 条 CES 规范作为质量基线
- Agent 生成的应用自动经过混沌测试验证

## 3. 数据模型

| 实体 | 存储 | 说明 |
|------|------|------|
| Project | Room | 项目元数据、文件树 |
| Session | Room | Agent 会话历史 |
| AgentConfig | DataStore | Agent 角色配置 |
| ModelConfig | DataStore | 模型配置（API Key 加密） |
| AuditLog | Room | 操作审计日志 |
| TestResult | Room | 混沌测试结果 |
| BuildCache | 文件系统 | 构建缓存 |

## 4. 依赖关系

```
app -> core, feature, data, design
feature -> core, data, design
core.agent -> core.model, core.tool, core.chaos
core.model -> data
core.tool -> core.tool.sandbox
core.chaos -> core (复用 Abate injector/probe/reporter)
data -> design (无)
design -> (无依赖，最底层)
```

所有依赖单向流动，无循环依赖。core.chaos 是 Abate 引擎的 Android 移植版。