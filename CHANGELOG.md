<!--
Abate —— 止息混沌 · 变异+混沌双结合的全量覆盖质量框架
Abate — where chaos abates. Mutation × Chaos unified total-coverage quality framework.
Copyright (C) 2026 小黎的部落（www.qingdujian.com）

本文件以 GNU Affero General Public License v3（AGPLv3）发布；商业用途须事先获得
小黎的部落书面授权，详见 COMMERCIAL.md。
This file is published under the GNU Affero General Public License v3 (AGPLv3).
Commercial use requires prior written authorization from 小黎的部落.
See COMMERCIAL.md for details.
-->

# 变更日志 | Changelog

---

## [Unreleased] — 2026-11-06

### 新增 | Added

#### 测试插件生态（第 18 个 Go 包 test/）

**第一阶段：30 插件创建**
- `test/plugin.go` — TestPlugin 接口 + TestResult + TestStatus（PASS/FAIL/SKIP）
- `test/registry.go` — 线程安全插件注册表，init() 自动注册
- `test/runner.go` — 矩阵编排器，串行执行 + 逐插件增量输出 + 超时控制
- 23 个原始测试插件：unit/property/integration/concurrency/compatibility/perf_baseline/backpressure/stress/soak/fuzz/conflict/defect/crash/vuln/security/adversarial/fault_tolerance/lint_intensive/static_analysis/dynamic_analysis/cicd_check/dual_build_verify/dual_mode_run
- 7 个配套检查插件：placeholder_check/stub_check/fake_impl_check/logic_check/file_header_check/function_check/vuln_check
- `test/matrix_test.go` — 全矩阵测试入口（build tag `matrix`）

**第二阶段：全语言通用改造**
- `test/langdetect.go`（新建）— 通用语言检测工具
  - `IsCodeFile()`: 支持 .go/.kt/.java/.py/.js/.ts/.rs/.c/.cpp 等 19 种扩展名
  - `DetectProjectType()`: 检测 go.mod/settings.gradle/pom.xml/package.json/Cargo.toml 等 6 种项目类型
  - `BuildCommand/VetCommand/TestCommand`: 按项目类型分发命令
  - `ShouldSkipDir/ShouldSkipFile`: 跳过 .git/vendor/node_modules/build/.gradle 等
- 7 个静态巡检插件: `.go` 过滤 -> `IsCodeFile()` + `ShouldSkipFile()`
- 17 个 go vet 插件: 硬编码 `go vet ./...` -> `DetectProjectType` + `VetCommand`，未知类型返回 `StatusSkip`
- `lint_intensive/static_analysis`: go vet -> `VetCommand` 分发
- `dual_build_verify`: Go 双模式保留，非 Go 单模式 `BuildCommand`
- `dual_mode_run`: Go 双模式保留，非 Go 单模式 `VetCommand`
- `cicd_check/compatibility`: `DetectProjectType` + `Build/Vet/TestCommand`

**第三阶段：端到端迭代实测**

| 轮次 | 目标 | Pass | Fail | Skip | 说明 |
|------|------|------|------|------|------|
| 第一轮（改造前） | Kotlin/Gradle | 7 | 23 | 0 | Go 专属插件在 Kotlin 上全部失败 |
| 第二轮（改造后未预热） | Kotlin/Gradle | 6 | 24 | 0 | Gradle 首次启动超时 |
| 第三轮（改造后+预热） | Kotlin/Gradle | 29 | 1 | 0 | 多语言改造成功 |
| 自托管验证 | Go | 30 | 0 | 0 | Go 兼容性完全保留 |
| 跨语言验证 | Python | 29 | 1 | 0 | 第三语言验证通过 |

**全链路实测汇总（96 项测试）**
- 插件矩阵: Go 30/30, Kotlin 29/30, Python 29/30
- CLI: 7/7 validate OK, 2/2 run OK
- Web API: 5/5 端点正常
- Engine: 4/4 PASS
- Sim: 19/19 PASS
- 2 FAIL 均为真实代码问题（Kotlin/Python 文件缺版权头）

#### 归属与完整性凭据（水印 / 亮牌 / 签名）
- **`watermark/` 包**（第 18 个 Go 包）— 代码级数字水印与运行时开机亮牌
  - `watermark/watermark.go`（55 行）— 水印常量（项目 / 许可证 / 小黎的部落 / www.qingdujian.com / xiaolibuluo@163.com）与 Banner 亮牌函数，纯内存计算
  - `watermark/no_network_test.go`（91 行）— AST 静态检查：零网络调用、零文件 I/O，禁止联网心跳校验
  - `watermark/watermark_test.go`（61 行）— 水印常量与亮牌横幅测试
- `abate serve` 启动亮牌输出到 stdout；`abate mcp` 启动亮牌输出到 stderr（stdout 完整保留给 JSON-RPC 2.0 流）
- `scripts/release.sh`（82 行）— 构建发布产物并生成 SHA-256 校验值文件与 GPG 分离签名（退出码 0/1/2，gpg 或私钥缺失时明确提示）
- `docs/SIGNING.md`（90 行）—《数字签名与验证指南》：水印、亮牌、SHA-256、GPG 验证步骤、发布流程与联系方式

#### CLI 一致性（CES-8.4，强制）
- `cmd/abate/cli_conformance_test.go`（163 行）— 锁死子命令集合、退出码 0/1/2 语义、flags 表面与亮牌输出流（CES-8.4.5）
- CES-8.4.1 — 子命令集合固定：serve / mcp / hub sync / llm generate / run / score / validate / version
- CES-8.4.2 — 退出码语义固定：0 成功 / 1 运行失败 / 2 用法错误
- CES-8.4.3 — serve 亮牌走 stdout、mcp 亮牌走 stderr
- CES-8.4.4 — flags 命名与默认值固定
- CES-8.4.5 — CLI 一致性由自动化测试锁死

#### 新 Go 包
- **`llm/` 包**（第 16 个 Go 包）— LLM 客户端实现，对接 OpenAI Chat Completions API 与 Anthropic Claude Messages API，均实现 `mut.LLMClient` 接口，零外部依赖（仅标准库 `net/http`）
  - `llm/client.go`（80 行）— Config 配置、BaseClient 公共 HTTP 能力、buildPrompt 提示词构造
  - `llm/openai.go`（92 行）— OpenAIClient，默认模型 `gpt-4o`
  - `llm/claude.go`（94 行）— ClaudeClient，默认模型 `claude-sonnet-4-20250514`
  - 17 个测试用例，100% 覆盖率

- **`hub/` 包**（第 17 个 Go 包）— ChaosHub 子编码注册表、远程同步与公告机制
  - `hub/registry.go`（162 行）— 线程安全 Registry、附录 B v0.1 种子表 20 项、Merge 合并
  - `hub/syncer.go`（82 行）— HTTP 远程同步，支持 API Key 认证
  - `hub/announce.go`（91 行）— Diff 差异计算、Announcement 公告、Format 格式化
  - 21 个测试用例 + 6 个基准测试，100% 覆盖率

#### CLI 子命令
- `abate mcp` — 启动 MCP Server（JSON-RPC 2.0 over stdio，协议版本 `2024-11-05`）
- `abate hub sync` — 同步 ChaosHub 子编码表并输出版本公告
- `abate llm generate` — 从命令行调用 LLM 生成补测代码

#### MCP Server 工具
- `generate_test` — 通过 LLM（OpenAI/Claude）生成补测代码杀死存活变异体
- MCP Server 现暴露 4 个工具：`generate_mutants` / `apply_mutant` / `list_fault_classes` / `generate_test`

#### Web API 端点
- `GET /api/v1/hub/subcodes` — 暴露 ChaosHub 子编码注册表（CES-B.1）

#### CES 规范条款（9 条新增）
- CES-7.3.1 — MCP Server 协议版本与 JSON-RPC 2.0 传输
- CES-7.3.2 — MCP Server 工具集要求
- CES-7.3.3 — generate_test provider 支持规范
- CES-9.5.3 — LLMClient 接口规范
- CES-9.5.4 — LLM 客户端自定义端点与安全要求
- CES-B.1 — 子编码注册表为唯一权威来源
- CES-B.2 — 远程同步与回退机制
- CES-B.3 — 版本公告要求

#### 实验示例（F1-F6 全矩阵）
- `examples/latency-injection.yaml` — F1 延迟注入（已有）
- `examples/pod-crash-f2.yaml` — F2 进程崩溃（新增）
- `examples/network-partition.yaml` — F3 网络分区（新增）
- `examples/disk-exhaustion.yaml` — F4 磁盘耗尽（新增）
- `examples/data-corruption.yaml` — F5 数据损坏（新增）
- `examples/clock-skew.yaml` — F6 时钟偏移（新增）

#### 文档示例
- `examples/mcp-server.md` — MCP Server 使用示例（中英双语）
- `examples/hub-sync.md` — hub sync 使用示例（中英双语）

#### 基准测试
- `hub/benchmark_test.go` — 6 个基准：AllCodes/AllCodeStrings/Lookup/Register/ComputeDiff/AnnouncementFormat
- `llm/benchmark_test.go` — 5 个基准：buildPrompt/NewOpenAIClient/NewClaudeClient/OpenAIGenerateTest/ClaudeGenerateTest

#### LLM 集成测试
- `mut/llm_integration_test.go` — 验证 OpenAIClient/ClaudeClient 实现 mut.LLMClient 接口并注入 Flywheel

#### 适配器集成测试（真实 k3s 集群）
- `adapter/integration_test.go`（build tag `integration`）— ChaosMesh/Litmus 在真实 k3s 集群上的 Apply→验证 CR 存在→Cleanup→验证删除→二次 Cleanup 幂等全链路测试；Gremlin 用 CLI shim 验证命令行契约
- 运行方式：`go test -tags integration ./adapter/ -v -count=1`
- 集成环境：k3s v1.37.2+k3s1（单机 control-plane）

#### 便捷脚本
- `scripts/hub-sync.sh`（74 行）— ChaosHub 子编码同步脚本（curl+jq，支持 -u/-k 参数与 ABATE_HUB_URL/ABATE_HUB_API_KEY 环境变量）
- `scripts/mcp-demo.sh`（71 行）— MCP Server 演示脚本（initialize / tools/list / list_fault_classes / generate_mutants 四个真实 JSON-RPC 请求）

#### Web 前端
- `web/static/index.html` — 新增「子编码 Sub-Codes」视图（207→226 行）：导航按钮、注册表面板（编码/名称/参数三列 + 版本徽标）、subcodes 加载器调用 `GET /api/v1/hub/subcodes`

#### 文档
- `README.md` — 中英文「核心功能 / Key Features」章节（+62 行：MCP Server、LLM 客户端、ChaosHub 同步、Web API）
- `COMMERCIAL.md` — 双轨许可边界新增两条：LLM 第三方服务（上游费用/数据政策与部落无关，API Key 不落日志）、ChaosHub 远程数据源（默认端点归部落运营，-url 自建端点责任归运营方）

#### ChaosHub 真实数据源验证
- 同版本同步：`ChaosHub v0.1 → v0.1: no changes` ✅（abate serve 作为数据源）
- 变更同步：`ChaosHub v0.1 → v0.2: + added: F1.4, ~ updated: F1.1` ✅（自定义数据源，Diff 公告机制完整工作）

### 修复 | Fixed
- `adapter/chaosmesh.go` — `cmDeleteFn` 对 kubectl NotFound 幂等返回成功（CES-3.1.2 Cleanup 可重复调用合规，集成测试发现的真实缺陷）
- `adapter/litmus.go` — `litDeleteFn` 同样修复 NotFound 幂等（CES-3.1.2 合规）
- `adapter/chaosmesh.go` / `adapter/litmus.go` — 清理 `_ = applyResult` 死代码
- `plugin/sandbox.go` — 清理未使用的 `before` 变量（设计残留）

### 变更 | Changed
- `adapter/mapping.go` — `allChaosHubSubCodes()` 从硬编码改为委托 `hub.NewSeedRegistry().AllCodeStrings()`
- `web/mcp.go` — 工具列表新增 `generate_test`，工具数量 3→4
- `.github/workflows/redline.yml` — Go 版本 1.23 → 1.25 匹配 go.mod
- `README.md` — 包数量 15→17，仓库结构添加 llm/hub，CLI 命令添加 mcp/hub sync/llm generate
- `CONTRIBUTING.md` — PR scope 添加 llm/hub，新增第 8 节 LLM/Hub 贡献规则
- `GOVERNANCE.md` — 生态轨添加 LLM 客户端与 ChaosHub 同步，数据飞轮治理补充 CES-B.1/B.2/B.3

### 统计 | Stats
- Go 包：15 → 18（新增 test 包）
- 测试插件：30 个（3 框架 + 30 插件 = 33 文件）
- 端到端实测：96 项测试 94 PASS 2 FAIL（真实问题）
- 测试用例：新增 60+ 个
- CES 条款：新增 9 条（中英文同步）
- 全项目覆盖率：100%（18/18 包，含新增 watermark 包）
- 非测试文件最大行数：291 行（dsl/parse.go，≤300 红线 ✅）
- 测试文件最大行数：299 行（plugin/registry_test.go，≤300 红线 ✅）
- 集成测试：3 个（真实 k3s 集群，`-tags integration`）
- 零外部依赖（仅 wazero v1.12.0 + golang.org/x/sys）
- 代码级数字水印：watermark 包，纯内存计算，AST 静态卡死零网络调用
- 运行时开机亮牌：serve→stdout / mcp→stderr，MCP stdout 零字节污染
- 发布签名：SHA-256 校验值 + GPG 分离签名（scripts/release.sh）
- CLI 一致性：CES-8.4.1~8.4.5 强制条款 + 自动化测试锁死
- Debug + Release 双 8 验证：vet/race/cover/count/红线/文件头/条款号/build 全部 PASS，pass 数据完全一致

---

Copyright (C) 2026 小黎的部落（www.qingdujian.com）
