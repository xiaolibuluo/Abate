<!--
Abate —— 止息混沌 · 混沌工程开放规范（CES）与参考实现
Abate — where chaos abates. Open Chaos Experiment Specification (CES) & reference implementation.
Copyright (C) 2026 小黎的部落（www.qingdujian.com）
本文件以 GNU Affero General Public License v3（AGPLv3）发布；商业用途须事先获得
小黎的部落书面授权，详见 COMMERCIAL.md。
This file is published under the GNU Affero General Public License v3 (AGPLv3).
Commercial use requires prior written authorization from 小黎的部落.
See COMMERCIAL.md for details.
-->

# 代码质量文档 | Code Quality Document

> Abate —— 止息混沌 · 混沌工程开放规范（CES）与参考实现
> Copyright (C) 2026 小黎的部落（www.qingdujian.com）· AGPLv3 + 商业书面授权
> 联系邮箱：xiaolibuluo@163.com

---

## 1. 质量指标总览

| 指标 | 数值 | 状态 |
|---|---|---|
| Go 包数量 | 18 | ✅ |
| 全项目测试覆盖率 | 100.0%（18/18 包） | ✅ |
| go vet | 通过，零告警 | ✅ |
| go test -race | 18 包通过，零竞态 | ✅ |
| 非测试文件最大行数 | 291 行（dsl/parse.go） | ✅ ≤300 |
| 测试文件最大行数 | 299 行（plugin/registry_test.go） | ✅ ≤300 |
| CES 规范条款数 | 120 条（中英文完全一致） | ✅ |
| 零外部依赖 | 仅 wazero v1.12.0 + golang.org/x/sys | ✅ |
| 集成测试 | 3 个（真实 k3s 集群，-tags integration） | ✅ |

## 2. 质量门禁机制

### 2.1 pre-commit hook（scripts/pre-commit）

每次 `git commit` 自动执行：
1. **300 行红线**：扫描暂存的 .go 文件，超过 300 行即拒绝提交
2. **go vet**：全仓静态分析
3. **go test**：全仓测试通过

### 2.2 CI 红线检查（.github/workflows/redline.yml）

push / pull_request 触发 6 项检查：
1. **300 行红线**：全仓 .go/.md/.sh/.py/.rs/.ts/.js/.yml/.yaml 文件
2. **开源文件头**：所有源文件前 20 行须含『小黎的部落』版权声明
3. **go vet**：全仓静态分析
4. **go test -cover**：全仓测试 + 覆盖率
5. **中英规范条款号一致性**：diff CES-v0.1.md 与 CES-v0.1.en.md 的条款 ID 集合
6. **DCO sign-off**：PR 提交须含 Signed-off-by

## 3. 归属与完整性凭据

### 3.1 代码级数字水印（watermark 包）

- 编译进二进制的常量：项目名、CES v0.1、AGPL-3.0-only + commercial、小黎的部落、www.qingdujian.com、xiaolibuluo@163.com
- 纯内存计算：不读写文件、不发起网络调用
- AST 静态卡死：no_network_test.go 禁止导入 net/http、net、os、io、crypto/tls，禁止 http.Get、net.Dial、os.ReadFile 等调用

### 3.2 运行时开机亮牌

- `abate serve`：横幅输出到 stdout（随后输出监听地址）
- `abate mcp`：横幅输出到 stderr，stdout 完整保留给 JSON-RPC 2.0 流（MCP stdout = 0 字节）
- 横幅内容：项目名、版权声明、许可证、联系邮箱、站点、版本号

### 3.3 发布签名

- scripts/release.sh：构建二进制 + SHA-256 校验值文件 + GPG 分离签名（armor）
- 退出码：0 成功 / 1 失败（缺 go、gpg 或私钥）/ 2 用法错误
- 验证方法见 docs/SIGNING.md

### 3.4 禁止联网心跳校验

- 水印与亮牌为纯内存计算，不依赖任何网络请求
- watermark 包由 AST 测试静态卡死零网络调用、零文件 I/O
- 许可证执法依赖法律途径（AGPLv3 传染条款 + 商业书面授权），不依赖技术遥测

## 4. CLI 一致性（CES-8.4，强制）

| 条款 | 内容 |
|---|---|
| CES-8.4.1 | 子命令集合固定：serve / mcp / hub sync / llm generate / run / score / validate / version |
| CES-8.4.2 | 退出码语义固定：0 成功 / 1 运行失败 / 2 用法错误 |
| CES-8.4.3 | serve 亮牌→stdout；mcp 亮牌→stderr，stdout 保留给 JSON-RPC 2.0 |
| CES-8.4.4 | flags 命名与默认值固定：serve -addr(:8080)、hub sync -url/-api-key、score -coverage-total(100)/-coverage-exercised(0) |
| CES-8.4.5 | CLI 一致性由自动化测试锁死（cmd/abate/cli_conformance_test.go） |

## 5. Debug + Release 双 8 验证

### 5.1 验证项

| # | 检查项 | Debug | Release |
|---|---|---|---|
| 1 | go vet ./... | PASS | PASS |
| 2 | go test -race ./... | 18 包 ok | 18 包 ok |
| 3 | go test -cover ./... | 18 包 100% | 18 包 100% |
| 4 | go test -count=1 ./... | 18 包 ok | 18 包 ok |
| 5 | 300 行红线 | PASS | PASS |
| 6 | 开源文件头 | PASS | PASS |
| 7 | 中英条款号一致性 | 120 条 | 120 条 |
| 8 | build + smoke test | PASS (dev) | PASS (v0.1.0) |

### 5.2 一致性比对

| 比对项 | 结果 |
|---|---|
| 覆盖率（18 包逐包 %） | ✅ COVER_CONSISTENT |
| pass 包集合（count=1） | ✅ COUNT_CONSISTENT |
| race pass 包集合 | ✅ RACE_CONSISTENT |
| MCP stdout 字节数 | ✅ D=0, R=0 一致 |
| 亮牌输出流 | ✅ serve→stdout / mcp→stderr 两模式一致 |

### 5.3 预期差异（非数据不一致）

| 项 | Debug | Release | 说明 |
|---|---|---|---|
| 二进制大小 | 11,500,274 字节 | 7,831,815 字节 | release 用 -s -w 剥离符号表与 DWARF |
| 版本号 | dev | v0.1.0 | release 用 -X main.version=v0.1.0 注入 |
| 构建参数 | 无优化 | -trimpath -ldflags '-s -w' | release 优化预期行为 |

## 6. 包清单与覆盖率

| 包 | 覆盖率 | 说明 |
|---|---|---|
| adapter | 100.0% | Chaos Mesh + Litmus + Gremlin 适配器 |
| builtin | 100.0% | 内置注入器 |
| cmd/abate | 100.0% | CLI 入口 |
| core | 100.0% | 核心类型与评分引擎 |
| dsl | 100.0% | DSL 解析器 |
| engine | 100.0% | 实验引擎 |
| hub | 100.0% | ChaosHub 注册表与同步 |
| injector | 100.0% | 注入器路由 |
| llm | 100.0% | LLM 客户端（OpenAI + Claude） |
| mut | 100.0% | 变异引擎与飞轮 |
| plugin | 100.0% | WASM 插件注册表 |
| probe | 100.0% | 探针 |
| reporter | 100.0% | 报告生成 |
| sim | 100.0% | 确定性模拟器 |
| store | 100.0% | 文件存储 |
| validator | 100.0% | DSL 校验器 |
| watermark | 100.0% | 数字水印与开机亮牌 |
| web | 100.0% | Web 控制台与 MCP Server |

## 7. 脚本清单

| 脚本 | 行数 | 功能 |
|---|---|---|
| scripts/pre-commit | 41 | pre-commit hook：300 行红线 + go vet + go test |
| scripts/hub-sync.sh | 82 | ChaosHub 子编码同步（curl + jq） |
| scripts/mcp-demo.sh | 100 | MCP Server 演示（JSON-RPC 2.0 over stdio） |
| scripts/release.sh | 82 | 发布构建 + SHA-256 + GPG 签名 |

## 8. 联系方式

- 邮箱：xiaolibuluo@163.com
- 站点：www.qingdujian.com

商业授权、公钥获取与安全问题报告均通过上述邮箱联系。
