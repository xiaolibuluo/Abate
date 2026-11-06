# Abate IDE —— Agent 协作流程

> Copyright (C) 2026 小黎的部落（www.qingdujian.com）
> AGPLv3; 商业用途须事先获得小黎的部落书面授权

## 1. 全链路流程

```
用户输入需求
    |
    v
CEO Agent -- 解析需求，分配任务
    |
    v
CPO Agent -- 拆解产品需求文档
    |
    v
CTO Agent -- 技术方案设计
    |
    +-> 设计部: UX 流程 + UI 规范 + 原型
    +-> 工程部: 架构师拆模块 -> 前端/逻辑/数据/模型并行
    +-> QA 部: 测试设计 Agent 同步编写用例
    +-> 安全部: 审计 Agent 同步扫描
    |
    v
COO Agent -- 汇总 -> 构建 Agent 打包 -> QA 执行测试
    |
    v
混沌 Agent -- Abate 注入故障 -> 验证恢复
    |
    v
发布 Agent -- 版本检查 -> 交付用户
```

## 2. DAG 依赖编排

```
CEO (start)
  -> CPO (depends: CEO)
  -> CTO (depends: CEO)
      -> 架构师 (depends: CTO)
          -> 前端 (depends: 架构师)
          -> 逻辑 (depends: 架构师)
          -> 数据 (depends: 架构师)
          -> 模型 (depends: 架构师)
      -> UX (depends: CPO)
          -> UI (depends: UX)
              -> 原型 (depends: UI)
      -> 测试设计 (depends: 架构师)
      -> 审计 (depends: 架构师, 并行)
  -> COO (depends: 前端,逻辑,数据,模型,原型,测试设计,审计)
      -> 构建 (depends: COO)
          -> 测试执行 (depends: 构建)
              -> 混沌 (depends: 测试执行)
                  -> 性能 (depends: 混沌)
      -> 发布 (depends: 性能)
```

## 3. 消息总线

- Agent 间通过消息总线传递数据，不共享内存
- 消息类型：TASK_ASSIGN / TASK_RESULT / TASK_ERROR / TASK_CANCEL
- 消息格式：JSON（kotlinx.serialization）
- 消息持久化到 Room 数据库

## 4. 失败策略

| 策略 | 行为 |
|------|------|
| continue | 其他任务继续执行 |
| abort_dependents | 跳过依赖失败任务的后续任务 |
| abort_all | 中止所有待执行任务 |

默认策略：abort_dependents

## 5. 阶段定义

| 阶段 | 参与角色 | 产出 |
|------|----------|------|
| 需求解析 | CEO, CPO | 需求文档 |
| 技术设计 | CTO, 架构师 | 架构方案、接口定义 |
| UI/UX 设计 | UX, UI, 原型 | 设计规范、原型代码 |
| 并行开发 | 前端, 逻辑, 数据, 模型 | 功能代码 |
| 同步检查 | 测试设计, 审计 | 测试用例、安全报告 |
| 构建打包 | 构建 | APK/AAB |
| 测试验证 | 测试执行, 混沌, 性能 | 测试报告 |
| 发布交付 | 发布 | 版本产物 |
| 运行监控 | 监控 | 崩溃/ANR 报告 |

## 6. 用户交互点

- 需求输入：用户描述需求，CEO Agent 解析
- 方案确认：CTO 方案可选展示给用户确认
- 进度查看：控制台实时展示各 Agent 状态
| 模型配置：用户可为每个阶段指定模型
- 结果验收：用户查看最终产物和测试报告