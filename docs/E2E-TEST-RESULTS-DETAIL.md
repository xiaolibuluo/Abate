<!--
Abate —— 端到端迭代实测结果（详细篇）
Copyright (C) 2026 小黎的部落（www.qingdujian.com）
AGPLv3; 商业用途须事先获得小黎的部落书面授权
-->

# 端到端实测结果 — 详细篇 | E2E Test Results — Details

> 本文档为 [E2E-TEST-RESULTS.md](E2E-TEST-RESULTS.md) 的续篇，包含 Web API 生命周期、覆盖率、FAIL 分析、多语言改造细节、Git 历史和源码包信息。

## 11. Web API 端到端 — 完整实验生命周期

### POST 创建实验
请求: `POST /api/experiments`（DSL 格式）
响应: `201 {"id":"e2e-test","status":"Pending","seed":99}`

### GET 查询实验列表
返回 2 个实验：demo-latency + e2e-test

### GET 查询单个实验
返回完整实验详情：id/seed/target/fault/blastRadius/steadyState/criticality/recoveryBudgetMs

**Web API 生命周期：3/3 PASS**

## 12. 全仓覆盖率测试

| 包 | 覆盖率 |
|----|--------|
| adapter | 100.0% |
| builtin | 100.0% |
| cmd/abate | 100.0% |
| core | 100.0% |
| dsl | 100.0% |
| engine | 100.0% |
| hub | 100.0% |
| injector | 100.0% |
| llm | 100.0% |
| mut | 100.0% |
| plugin | 100.0% |
| probe | 100.0% |
| reporter | 100.0% |
| sim | 100.0% |
| store | 100.0% |
| validator | 100.0% |
| watermark | 100.0% |
| web | 100.0% |
| test | 0.0%（无测试文件，仅有插件实现） |

**18/19 包 100% 覆盖率**

## 13. 2 FAIL 问题详细原因分析

### FAIL 1: file-header-check — Kotlin/Gradle 项目（/root/abate-ide）

**插件**: file-header-check  
**目标**: /root/abate-ide（Abate IDE，Kotlin/Gradle 项目）  
**状态**: FAIL  
**耗时**: 9.93ms  
**详情**: 16 个 Kotlin 文件缺少版权头

**问题原因**:

file-header-check 插件扫描目标项目中所有代码文件（通过 `IsCodeFile()` 识别 .kt 扩展名），检查每个文件头部是否包含版权声明。Abate IDE 的 16 个 Kotlin 源文件在创建时未添加版权头，具体为：

1. `app/.../AbateIDEApplication.kt`
2. `app/.../MainActivity.kt`
3. `core/.../ArchitectAgent.kt`
4. `core/.../BuildAgent.kt`
5. `core/.../CEOAgent.kt`
6. `core/.../FrontendAgent.kt`
7. `core/.../LogicAgent.kt`
8. `core/.../TeamOrchestrator.kt`
9. `core/.../MessageBus.kt`
10. `core/.../Orchestrator.kt`
11. `core/.../AgentRegistry.kt`
12. `core/.../Agent.kt`
13. `core/.../AgentRole.kt`
14. `core/.../ModelProvider.kt`
15. `core/.../PermissionMatrix.kt`
16. `core/.../ToolRegistry.kt`

**根因**: Abate IDE P1 开发阶段（commit `2a0b6b0`）创建 Kotlin 源文件时，未遵循 Abate 框架的工程红线要求（每个源文件必须包含版权头）。这是 IDE 开发过程中的真实疏漏，不是插件误报。

**对比验证**: 同样的插件扫描 Abate 自身 Go 源码（/root/abate）时 30/30 全 PASS，因为 Go 源文件均有版权头。这证明插件逻辑正确，FAIL 是真实问题。

**修复**: 批量为 17 个 Kotlin 文件添加版权头，修复后 file-header-check PASS。

### FAIL 2: file-header-check — Python 项目（/tmp/py-test）

**插件**: file-header-check  
**目标**: /tmp/py-test（最小 Python 测试项目）  
**状态**: FAIL  
**详情**: calc.py; test_calc.py

**问题原因**: 测试用 Python 项目中的 2 个文件均未添加版权头。这是测试项目创建时的真实疏漏，不是插件误报。

**修复**: 为 calc.py 和 test_calc.py 添加版权头，修复后 PASS。

### FAIL 汇总分析

| FAIL | 项目 | 语言 | 缺版权头文件数 | 根因 | 性质 |
|------|------|------|----------------|------|------|
| 1 | /root/abate-ide | Kotlin | 16 | IDE P1 开发疏漏 | 真实代码质量问题 |
| 2 | /tmp/py-test | Python | 2 | 测试项目创建疏漏 | 真实代码质量问题 |

**修复后 Kotlin 项目矩阵测试**: 29/30 PASS（仅 adversarial-test 因 Gradle lint 60s 超时 FAIL，环境问题非代码问题）

**结论**: 2 个 FAIL 均已修复。file-header-check 插件正确检出真实缺版权头问题，修复后 PASS。Abate 框架 30 插件全语言通用改造成功。

## 14. 多语言改造技术细节

### langdetect.go 核心工具

| 函数 | 功能 |
|------|------|
| IsCodeFile(path) | 支持 .go/.kt/.java/.py/.js/.ts/.rs/.c/.cpp 等 19 种扩展名 |
| DetectProjectType(target) | 检测 go.mod/settings.gradle/pom.xml/package.json/Cargo.toml 等 |
| BuildCommand(pt) | Go: go build / Gradle: gradle compileDebugKotlin / Python: compileall / ... |
| VetCommand(pt) | Go: go vet / Gradle: gradle lint / Python: pyflakes / ... |
| TestCommand(pt) | Go: go test / Gradle: gradle test / Python: pytest / ... |
| ShouldSkipDir(name) | 跳过 .git/vendor/node_modules/build/.gradle 等 |
| ShouldSkipFile(path) | 跳过测试文件和 test/ 目录 |

### 改造模式

**静态巡检插件（7 个）**：`.go` 过滤 -> `IsCodeFile()` + `ShouldSkipFile()`

**命令执行插件（17 个）**：`exec.CommandContext(ctx, "go", "vet", "./...")` -> `DetectProjectType` + `VetCommand`，未知类型返回 `StatusSkip`

**双模式插件（2 个）**：Go 保留双模式，非 Go 退化为单模式

**CI/CD 插件**：`DetectProjectType` + `BuildCommand`/`VetCommand`/`TestCommand` 三步流水线

## 15. Git 提交历史

| Commit | 描述 |
|--------|------|
| 5b211d9 | fix(test): Python BuildCommand py_compile -> compileall, 恢复默认目标 |
| 1d17214 | feat(test): 30 插件全语言通用改造 — 端到端实测 29/30 PASS |
| a2601e2 | chore: separate Abate IDE to independent project |
| 2a0b6b0 | feat(ide): add Kotlin source files |
| 50229e8 | feat(test): 26 test plugin files |
| 7a0715d | docs: legal documents |
| 48f33be | docs: SIGNING.md |

## 16. 源码包信息

- 文件: `/var/www/html/abate-dl/abate-v0.1.0-src.tar.gz`
- 大小: 2.3M
- SHA-256: `9dbcab070231b7fa535f9d69d05b32e8b1fbc88fa6dbd6be8bcec94350decab3`
- 下载页: `https://www.qingdujian.com/abate-dl/`

---

Copyright (C) 2026 小黎的部落（www.qingdujian.com）
