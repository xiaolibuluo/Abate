# Abate IDE —— Agent 团队组织架构

> Copyright (C) 2026 小黎的部落（www.qingdujian.com）
> AGPLv3; 商业用途须事先获得小黎的部落书面授权

## 1. 组织架构

CEO Agent (战略决策) -> CTO/CPO/COO -> 5 部门 19 角色，共 23 人。

## 2. 决策层（4人）

| 角色 | 职责 |
|------|------|
| CEO | 需求解析、任务分解、资源分配 |
| CTO | 技术选型、架构决策、风险评估 |
| CPO | 产品需求拆解、用户体验定义 |
| COO | 流程编排、进度追踪、汇总交付 |

## 3. 工程部（6人）

| 角色 | 职责 | 工具 |
|------|------|------|
| 架构师 | 模块拆分、接口定义、依赖管理 | read/write/search/dependency_tree |
| 前端 | Compose UI、页面交互、动画 | write/preview/layout_inspector |
| 逻辑 | 业务逻辑、状态管理、数据流 | write/read/unit_test |
| 数据 | Room 数据库、DAO、迁移 | write/schema_diff/migration_check |
| 模型 | 模型适配、API 对接、本地推理 | write/model_test/inference_benchmark |
| 构建 | Gradle、ProGuard、签名、打包 | build_run/sign_config/variant |

## 4. 设计部（3人）

| 角色 | 职责 | 工具 |
|------|------|------|
| UX | 用户流程、交互逻辑、信息架构 | write/flowchart/spec_write |
| UI | 视觉规范、色彩字体间距、暗色模式 | write/color_palette/theme_generate |
| 原型 | 可交互原型、Compose 预览 | write/preview_render |

## 5. QA 部（4人）

| 角色 | 职责 | 工具 |
|------|------|------|
| 测试设计 | 测试计划、用例编写、覆盖率 | write/coverage_report |
| 测试执行 | 运行单元/集成/UI 测试 | test_run/test_report |
| 混沌 | Abate 引擎注入故障、验证恢复 | abate_inject/abate_report |
| 性能 | 内存/CPU/启动/帧率分析 | profile_run/trace_view |

## 6. 安全部（3人）

| 角色 | 职责 | 工具 |
|------|------|------|
| 审计 | 代码安全扫描、权限审查 | static_scan/permission_audit |
| 沙箱 | WASM 沙箱验证、逃逸检测 | sandbox_test/isolation_verify |
| 隐私 | GDPR/CCPA 合规、数据审计 | compliance_check/data_flow_audit |

## 7. 运维部（3人）

| 角色 | 职责 | 工具 |
|------|------|------|
| CI/CD | 流水线编排、增量构建、缓存 | pipeline_run/cache_manage |
| 发布 | 版本管理、changelog、发布检查 | version_tag/release_check |
| 监控 | 崩溃捕获、ANR 检测、性能埋点 | crash_report/anr_detect |

## 8. 技能矩阵

| 角色 | 代码生成 | 代码审查 | 测试 | 安全 | 构建 | 混沌 | 性能 | UI |
|------|:------:|:------:|:----:|:----:|:----:|:----:|:----:|:---:|
| CEO | - | - | - | - | - | - | - | - |
| CTO | - | 高级 | - | - | - | - | - | - |
| 架构师 | 高级 | 高级 | - | - | - | - | - | - |
| 前端 | 高级 | 中级 | 初级 | - | - | - | - | 中级 |
| 逻辑 | 高级 | 中级 | 初级 | - | - | - | - | - |
| 数据 | 高级 | 中级 | 初级 | - | - | - | - | - |
| 模型 | 高级 | 中级 | 初级 | - | - | - | 中级 | - |
| 构建 | 中级 | - | - | - | 高级 | - | - | - |
| UX | 初级 | - | - | - | - | - | - | 高级 |
| UI | 初级 | - | - | - | - | - | - | 高级 |
| 原型 | 中级 | - | - | - | - | - | - | 高级 |
| 测试设计 | - | 初级 | 高级 | - | - | - | - | - |
| 测试执行 | - | - | 高级 | - | - | - | - | - |
| 混沌 | - | - | 中级 | - | - | 高级 | - | - |
| 性能 | - | - | 中级 | - | - | - | 高级 | - |
| 审计 | - | 高级 | - | 高级 | - | - | - | - |
| 沙箱 | - | 中级 | - | 高级 | - | - | - | - |
| 隐私 | - | 中级 | - | 高级 | - | - | - | - |
| CI/CD | 初级 | - | - | - | 高级 | - | - | - |
| 发布 | 初级 | - | - | - | 高级 | - | - | - |
| 监控 | 初级 | - | - | - | - | - | 高级 | - |

## 9. 协作规则

- 决策层通过消息总线下发任务，部门内并行执行
- 跨部门依赖通过 COO Agent 编排
- 写操作需经权限矩阵校验 + 审计日志
- 安全部只读不写，保证审计独立性
- 混沌 Agent 仅在 QA 阶段激活
- 每个 Agent 可绑定独立模型，支持按角色配置