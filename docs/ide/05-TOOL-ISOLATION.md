# Abate IDE —— 工具隔离体系

> Copyright (C) 2026 小黎的部落（www.qingdujian.com）
> AGPLv3; 商业用途须事先获得小黎的部落书面授权

## 1. 架构

```
Tool Registry (注册中心)
  |
  +-- read_file / write_file / search_code / glob_files / dependency_tree
  +-- build_run / sign_config / variant_manage / cache_manage
  +-- test_run / test_report / coverage_report / abate_inject / abate_report
  +-- profile_run / trace_view
  +-- static_scan / permission_audit / sandbox_test / isolation_verify
  +-- compliance_check / data_flow_audit
  +-- preview_render / layout_inspector / flowchart_render
  +-- color_palette / theme_generate / spec_write
  +-- pipeline_run / version_tag / release_check / crash_report / anr_detect
  |
  v
Permission Matrix (权限矩阵) -- 强制校验
  |
  v
WASM Sandbox (执行沙箱) -- 隔离执行
  |
  v
Audit Log (审计日志) -- 全记录
```

## 2. 权限矩阵

| 角色 | read | write | build | test | abate | security | design | ops |
|------|:----:|:-----:|:-----:|:----:|:-----:|:--------:|:------:|:---:|
| CEO | R | - | - | - | - | - | - | - |
| CTO | R | - | - | - | - | R | - | - |
| CPO | R | - | - | - | - | - | R | - |
| COO | R | - | - | R | - | - | - | R |
| 架构师 | R | W | - | - | - | - | - | - |
| 前端 | R | W | - | - | - | - | W | - |
| 逻辑 | R | W | - | - | - | - | - | - |
| 数据 | R | W | - | - | - | - | - | - |
| 模型 | R | W | - | - | - | - | - | - |
| 构建 | R | W | W | - | - | - | - | - |
| UX | R | W | - | - | - | - | W | - |
| UI | R | W | - | - | - | - | W | - |
| 原型 | R | W | - | - | - | - | W | - |
| 测试设计 | R | W | - | W | - | - | - | - |
| 测试执行 | R | - | - | W | - | - | - | - |
| 混沌 | R | - | - | - | W | - | - | - |
| 性能 | R | - | - | W | - | - | - | - |
| 审计 | R | - | - | - | - | W | - | - |
| 沙箱 | R | - | - | - | - | W | - | - |
| 隐私 | R | - | - | - | - | W | - | - |
| CI/CD | R | W | W | - | - | - | - | W |
| 发布 | R | W | W | - | - | - | - | W |
| 监控 | R | - | - | - | - | - | - | W |

R=只读 W=可写

## 3. WASM 沙箱

- 所有工具在 WASM 沙箱中执行（wazero 移植）
- 禁止文件系统越权访问
- 限制内存和 CPU 配额
- 网络访问仅限模型 API 端点白名单
- 沙箱逃逸检测由安全部沙箱 Agent 负责

## 4. 审计日志

- 所有写操作记录：Agent ID、工具名、参数、时间戳、结果
- 审计日志存储在 Room 数据库，不可篡改
- 安全部审计 Agent 可读取全部审计日志
- 用户可在控制台查看操作历史

## 5. 隔离原则

1. 每个 Agent 只能调用权限矩阵中允许的工具
2. 写操作经过审计日志记录
3. 构建和签名工具仅限构建 Agent 和发布 Agent
4. 安全 Agent 只读不写，保证审计独立性
5. Abate 混沌注入仅限混沌 Agent 调用
6. 工具调用全部在 WASM 沙箱中执行
7. 跨工具数据传递通过消息总线，不共享内存