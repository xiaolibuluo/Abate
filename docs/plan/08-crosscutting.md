# Abate 横切体系规划 — 七体系极致细节

- 文档编号：PLAN-08　状态：生效　上游：PLAN-00
- 覆盖模块：XCU-SEC / XCU-GOV / XCU-CI / XCU-REL / XCU-TST / XCU-PRF / XCU-DOC
- 通用约束（七体系共同遵守）：
  1. 横切体系作用于全部 L0-L7 层，但不引入新的层级依赖（工具与流程，非运行时组件；唯一例外 XCU-SEC 的运行时部分已归属各层模块）。
  2. 横切体系自身同样受五段门禁约束（工具脚本也是代码：gofmt/vet/lint/test/build 全绿）。
  3. 每个体系有唯一负责人与季度评审节奏（登记于体系 README）。

---

## 体系 XCU-SEC：安全体系

### 1. 体系定位与职责

全引擎安全纵深：威胁建模、签名链、漏洞管理流程、依赖供应链安全、密钥管理。运行时安全机制（能力沙箱/插件验签）已归属 L1 各模块；本体系负责策略、流程与审计。

### 2. 文件布局

```
security/
  THREAT-MODEL.md     // 威胁模型文档（季度评审）
  policy/             // 安全策略文件
    key-policy.md     // 密钥管理策略
    disclosure.md     // 漏洞披露流程
  audit/              // 审计工具
    scan.sh           // 依赖漏洞扫描入口
  tests/              // 安全测试用例集
```

### 3. 威胁模型（首版必须完整建立）

| 编号 | 威胁 | 攻击面 | 缓解 | 归属模块 |
|---|---|---|---|---|
| T1 | 恶意插件加载 | 插件安装通道 | 强制签名验证（未签名拒绝） | WAS-SIG |
| T2 | 插件越权（fs/net/提权） | host function 调用 | 能力沙箱默认全拒+运行时门控+审计 | WAS-CAP |
| T3 | 插件资源耗尽（内存/燃料） | 插件执行 | 燃料+内存+结果长度三重上限 | WAS-HOST |
| T4 | 供应链投毒（依赖） | go.mod/third_party | 源码级并入+commit 哈希锁定+许可证审查 | XCU-GOV |
| T5 | 签名密钥泄露 | 发布通道 | 密钥离线保管+轮换流程+重叠期 | XCU-REL |
| T6 | CES 文档注入（恶意表达式/超深嵌套） | 文档解析 | 深度上限 64+输入上限 10MB+模糊测试 | CES-PARSE |
| T7 | 报告 XSS（HTML/消息注入） | 报告输出 | html/template 自动转义+模板变量白名单 | RPT-HTML/ADP-* |
| T8 | REST 未授权访问 | API 监听 | 强制 Bearer 鉴权+远程监听无 token 启动拒绝 | CTL-REST |
| T9 | 令牌/密钥泄漏入日志 | 全局日志 | 脱敏中间件+扫描门禁（正则扫描日志调用） | 全局 |
| T10 | 档案篡改（回放证据伪造） | 档案文件 | 帧哈希链+全档 SHA-256 | SIM-REC |

### 4. 关键流程（必须落地）

1. 漏洞披露流程：security@ 邮箱（xiaolibuluo@163.com，独立安全通道）→ 48h 确认 → 90 天修复窗口 → CVE 申请（严重级）→ 公告发布。披露策略文本（disclosure.md）随仓库公开。
2. 依赖漏洞扫描：CI 周期任务（govulncheck 源码级并入 + third_party 锁定版本人工审查清单）；新 CVE 命中 third_party 锁定版本 → 评审回移流程（对齐 KRN-INJ 上游回移机制）。
3. 密钥管理：发布签名密钥离线保管（硬件介质）；信任锚公钥随版本更新并公示指纹（发布公告+仓库 SECURITY.md）；轮换重叠期 2 个版本（新旧密钥并存验签）。
4. 安全审计节奏：每里程碑（M1-M6）结束跑全量安全测试用例集（T1-T10 各威胁至少 1 个攻击用例）；GA 前第三方安全审计（外部）。

### 5. 依赖与约束

- 依赖：各层模块（策略消费方）。
- 禁止：安全策略与实现脱节（威胁模型每条威胁必须指向具体归属模块与测试用例——扫描校验）。

### 6. 测试要求

- T1-T10 每威胁 ≥1 攻击用例（负向：攻击被拦截断言）；用例集入 security/tests/ 并入 CI。
- 日志脱敏扫描：全仓库日志调用正则扫描（token/密钥模式）零命中。

### 7. 验收标准（可量化）

1. 威胁模型 10 条威胁全部有归属模块+缓解措施+攻击用例（扫描校验 100% 对齐）。
2. 安全测试用例集 CI 全绿；日志脱敏扫描零命中。
3. 漏洞披露流程文档公开；密钥轮换流程演练一次（记录入库）。
4. GA 前外部安全审计完成（报告入库）。

---

## 体系 XCU-GOV：治理合规

### 1. 体系定位与职责

版权与授权治理：强制版权头、数字代码水印（小黎的部落）、AGPLv3+商业双授权、第三方源码并入合规（许可证审查+来源登记）。

### 2. 文件布局

```
governance/
  HEADERS.md          // 版权头规范与模板
  THIRD-PARTY.md     // 第三方源码并入登记册
  DUAL-LICENSE.md    // 双授权说明（仓库根）
scripts/
  header-check.sh     // 版权头检查门禁
  watermark.sh        // 构建水印注入
```

### 3. 版权头规范（强制，所有用户使用都必须强制版权头）

- 模板（Go 源文件）：

```go
// Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
// Abate is dual-licensed under AGPLv3 and a commercial license.
// For commercial licensing, contact xiaolibuluo@163.com.
// Source: https://github.com/xiaolibuluo/Abate
```

- 覆盖范围：全部 .go/.sh/.yml/.md（文档头部 YAML front-matter 或首行注释）；third_party/ 排除（保留上游版权头，登记册管理）。
- 门禁：header-check.sh 逐文件首部匹配（缺失/格式错误即 FAIL）；CI 常驻。

### 4. 数字代码水印（三重）

1. 版权头水印：如上模板（静态，随源码分发）。
2. 构建元数据水印：-ldflags 注入（版本/commit/构建时间/水印标识『小黎的部落』）；abate version 输出全量元数据。
3. 发布清单水印：发布资产 SHA256 清单（manifest.txt）含水印行；GPG 签名覆盖（防篡改）。

### 5. 第三方源码并入登记册（THIRD-PARTY.md，强制）

每条登记：库名/版本/上游 commit 哈希/许可证/审查结论/裁剪清单/回移策略。首版登记清单（对齐各分册）：

| 库 | 用途 | 许可证 | 并入位置 |
|---|---|---|---|
| PIT | 变异算法参考 | Apache-2.0 | third_party/pit（只读参考） |
| chaosblade | 注入核源码 | Apache-2.0 | third_party/chaosblade → kernel/inject |
| wazero | WASM 运行时 | Apache-2.0 | third_party/wazero → wasm/host |
| yaml.v3 | YAML 解析 | Apache-2.0 | third_party/yaml → ces/parse |
| openpgp (x/crypto) | GPG 验签 | BSD-3-Clause | third_party/openpgp → wasm/signature |
| sigstore/cosign 验证库 | cosign 验签 | Apache-2.0 | third_party/sigstore → wasm/signature |
| ruff parser crate | Python 解析 | MIT/Apache-2.0 | third_party/ruff-parser → lang/pyfront |
| tree-sitter-python | Python 容错解析 | MIT | third_party/tree-sitter-python → lang/pyfront |
| swc parser crate | JS/TS 解析 | Apache-2.0 | third_party/swc → lang/jsfront |
| LLVM/clang (libclang) | C/C++ 解析 | Apache-2.0+LLVM Exception | third_party/llvm → lang/cfront |
| syn | Rust 解析 | MIT/Apache-2.0 | third_party/syn → lang/rsfront |
| Roslyn | C# 解析 | MIT | third_party/roslyn → lang/csfront |
| kotlin-compiler 前端 | Kotlin 解析 | Apache-2.0 | third_party/kotlin-compiler → lang/ktfront |
| otel-go SDK | OTel 导出 | Apache-2.0 | third_party/otel → report/otel |
| SplitMix64 | PRNG 算法 | 公有领域 | sim/rand（算法移植） |

- 审查流程：新并入必须走登记册条目评审（许可证兼容性：AGPLv3 传染性审查——Apache-2.0/MIT/BSD 兼容；GPL-2-only 不兼容即拒绝）。

### 6. 双授权（DUAL-LICENSE.md）

- AGPLv3（开源侧）+ 商业授权（闭源侧，联系 xiaolibuluo@163.com）。
- 贡献者协议：CLA 可选+DCO 强制（Signed-off-by）；贡献即授予双授权权利（贡献者协议文本随仓库）。

### 7. 测试要求与验收标准

1. header-check.sh 全仓库扫描零缺失（CI 常驻；故意删头文件用例被拦截）。
2. 水印三重验证：version 元数据含水印；manifest 含水印行且签名覆盖；构建产物 ldflags 注入断言。
3. THIRD-PARTY.md 与 third_party/ 目录零漂移（扫描校验：目录内每个库有登记条目）。
4. 许可证兼容性审查结论 100% 覆盖登记清单。

---

## 体系 XCU-CI：CI/CD 门禁

### 1. 体系定位与职责

合入与发布之间的全量质量门禁：五段门禁链（已建成）+ 版权头门禁 + 基准回归门禁 + TCK 门禁 + 确定性门禁 + 漂移检查门禁。任何规则命中即 FAIL 并打印规则条例。

### 2. 文件布局

```
scripts/
  lint-gate.sh       // 五段门禁（已建成：gofmt→vet→golangci-lint→test -race→build）
  header-check.sh    // 版权头门禁
  drift-check.sh     // 漂移检查（生成物与源码同步）
  determinism-check.sh // 确定性门禁（双跑哈希比对）
.github/workflows/
  ci.yml             // 主门禁（已建成）
  benchmarks.yml     // 基准回归（M2 起）
  nightly.yml        // 夜间全量（模糊+对拍+大规模）
```

### 3. 门禁链全量定义（合入 main 的完整关卡）

1. 五段门禁（lint-gate.sh，已建成）：gofmt → go vet（36 分析器）→ golangci-lint（110+ linter，issues-exit-code=1）→ go test -race -count=1 → go build。
2. 版权头门禁（header-check.sh）：全文件首部模板匹配，缺失即 FAIL。
3. 确定性门禁（determinism-check.sh）：核心包双跑（同种子）→ SIM-REC Compare 零差异；命中包清单：kernel/*、sim/*、ces/*（M1 起逐步纳入）。
4. 漂移检查门禁（drift-check.sh）：四类生成物与源码零漂移——CES JSON Schema/fields.md（CES-SCHEMA 生成器）、OpenAPI（CTL-REST 路由）、前端 API 客户端（openapi.json 生成）、前端 dist（frontend 构建）。
5. 基准回归门禁（benchmarks.yml，M2 起）：基准套件运行与基线比对（性能预算表阈值），回退超阈值即 FAIL。
6. TCK 门禁：CES-TCK 自举（Abate Driver 100% 通过）+ 内核 API TCK 全绿。
7. 夜间全量（nightly.yml）：模糊测试长跑（各解析器 30 分钟档）+ PIT 对拍 + 规模测试（100 万帧档案/10 万包网络/10000 任务）。

### 4. 关键流程

1. 合入流程：本项目不接受 PR；权利人本地全部门禁通过后，直接 GPG 签名提交并推送 `main`（分支保护仅强制签名 + 禁 force push）；门禁失败输出规则条例（对齐『触发时明确打印规则条例』要求）。
2. 门禁分级：快速档（本地提交前，≤10 分钟：五段+版权头+确定性快速档）与全量档（夜间）；基准回归在专用 runner（性能抖动隔离）。
3. 门禁自身测试：门禁脚本有单元测试（注入违规样本断言 FAIL——已验证 godox/unused 拦截，扩展到全部门禁段）。

### 5. 测试要求与验收标准

1. 七类门禁全部建成且 CI 常驻；每类门禁有『注入违规被拦截』负向用例（对齐已验证的桩代码拦截测试）。
2. PR 快速档 P95 ≤10 分钟（CI 时长统计入库）。
3. 漂移检查四类生成物零漂移（CI 记录）。
4. 门禁脚本自身五段门禁全绿（工具也是代码）。

---

## 体系 XCU-REL：发布工程

### 1. 体系定位与职责

版本发布流程：tag 签名、资产构建、校验和清单、SBOM、公告。年轮内核版本+小时级生态发布的版本策略落地。

### 2. 文件布局

```
scripts/
  release.sh        // 发布流程编排（构建→校验→签名→清单）
  sbom.sh           // SBOM 生成（SPDX 格式）
SECURITY.md         // 安全策略与信任锚指纹公示
CHANGELOG.md        // 变更日志（每版本条目）
```

### 3. 版本策略

1. 内核年轮：每年一个 major（年轮号）；LTS 3 年维护窗口；内核 API semver（major ≥3 年间隔，对齐十年内核铁律）。
2. 生态小时级：插件仓库独立版本（semver）+ 兼容性矩阵（宿主版本区间）；插件 CI 按小时发布窗口。
3. 引擎版本：v<年轮>.<minor>.<patch>（如 v1.2.3）；预发布（alpha/beta/rc）不入 LTS。

### 4. 发布流程（release.sh，必须全自动化）

1. 前置校验：全量门禁绿（XCU-CI 全链）+ CHANGELOG 条目齐 + 版本号一致性（ldflags/CHANGELOG/OpenAPI 三处版本一致）。
2. 构建：目标矩阵（linux/amd64、linux/arm64、darwin/amd64、darwin/arm64、windows/amd64）；构建可复现（双构建字节一致——GOFLAGS=-trimpath+固定工具链）。
3. 资产：二进制 + 压缩包 + SHA256 清单（manifest.txt，含水印行）+ SBOM（SPDX 2.3，含 third_party 全登记）。
4. 签名：全部资产+清单 GPG 签名（.sig 文件）；tag 签名（annotated+signed）。
5. 公告：GitHub Release（资产+签名+SBOM）+ 信任锚指纹公示（SECURITY.md 更新，若轮换）。

### 5. 测试要求与验收标准

1. 发布流程演练（M6 验收）：全流程走通（校验→构建→签名→清单→Release 草稿）；构建可复现（双构建字节一致）。
2. SBOM 过 SPDX 校验器；签名链验证（公钥验签全部资产）。
3. 版本号三处一致性检查入 CI（漂移即 FAIL）。
4. 回滚流程：发布缺陷的回滚演练一次（记录入库）。

---

## 体系 XCU-TST：测试体系

### 1. 体系定位与职责

全引擎测试策略与公共设施：测试金字塔、对拍测试、模糊测试、混沌自测（吃自己的狗粮终极形态——Abate 对自身运行混沌实验）。

### 2. 测试金字塔（比例与归属）

1. 单元测试 70%：表驱动；正常/边界/错误三路径全覆盖；各模块分册已定义（合计 ≥2000 用例规模预估）。
2. 集成测试 20%：跨模块链路（分册『合入总验收』定义的端到端用例）。
3. 端到端测试 10%：CLI/REST/Web 三入口全流程。

### 3. 专项测试（四类）

1. 对拍测试：PIT 对拍（变异判定一致率 100%，KRN-MUT 验收项）；golangci-lint 对拍（Go 规则集差异报告，LNG-CHK 验收项）；上游引擎对拍（chaosblade 行为锁定，KRN-INJ TCK）。
2. 模糊测试：CES-PARSE（文档输入）、LNG-* 七前端（源码输入）、AIR 序列化（随机树）、WAS-HOST（wasm 字节）——四类目标各 30 分钟夜间档+60s PR 快速档；语料库入库持续演化。
3. 混沌自测（狗粮终极）：Abate 对自身代码库跑变异实验（M 族）+ 对自身测试环境跑混沌实验（C 族，模拟器内）；存活分数纳入质量看板（目标：核心包存活分数 ≥0.85；SURVIVED 变异体自动建改进 issue——ADP-ITS 联动）。
4. 确定性测试：全核心包双跑哈希一致（XCU-CI 确定性门禁）；SIM-REC 回放逐帧校验。

### 4. 公共设施

- testutil 包：确定性断言助手（双跑哈希/序列化比较）、夹具加载器（testdata）、基准语料管理。
- 语料库：testdata/ 统一布局（各模块分册定义）；语料来源与许可登记（开源快照的许可审查——对齐 XCU-GOV）。

### 5. 验收标准

1. 金字塔比例达标（用例数统计报告：单元/集成/端到端）；合计用例规模 ≥2000。
2. 四类专项测试全部建成入 CI；模糊语料库入库且持续增长（季度统计）。
3. 混沌自测：核心包存活分数 ≥0.85（看板入库）；SURVIVED 变异体 100% 建 issue。
4. 对拍报告（PIT/golangci-lint/chaosblade 三份）入库且里程碑更新。

---

## 体系 XCU-PRF：性能基准

### 1. 体系定位与职责

性能预算的制定、基准套件维护与回归门禁。『快』『优』两字的量化落点。

### 2. 文件布局

```
bench/
  suite_test.go     // 基准套件入口（go test -bench）
  baseline/         // 基线数据（按平台归档）
  budget.md         // 性能预算表（本表）
```

### 3. 性能预算表（全量，各分册验收项汇总）

| 编号 | 指标 | 预算 | 归属 |
|---|---|---|---|
| P1 | 调度器事件吞吐 | ≥10 万事件/秒 | KRN-SCH |
| P2 | PRNG 提取吞吐 | ≥10^8 次/秒 | SIM-RNG |
| P3 | 虚拟网络传播（10 万包） | P99 ≤1s | SIM-NET |
| P4 | 虚拟调度（1 万任务） | 吞吐 ≥10 万事件/秒 | SIM-SCHED |
| P5 | 档案 Load+Verify（100 万帧） | P99 ≤10s | SIM-REC |
| P6 | AIR 序列化（100 万节点） | P99 ≤500ms | LNG-AIR |
| P7 | AIR 反序列化（100 万节点） | P99 ≤700ms | LNG-AIR |
| P8 | 前端解析吞吐（六语言） | ≥1MB/s（C/C++ ≥0.5MB/s） | LNG-* |
| P9 | 规则匹配（100 规则/文件） | P99 ≤50ms | LNG-RULE |
| P10 | HTML 报告（10 万违规） | P99 ≤5s | RPT-HTML |
| P11 | 插件调用开销 | P99 ≤50μs | WAS-HOST |
| P12 | 插件验签 | P99 ≤10ms | WAS-SIG |
| P13 | CES 解析（100KB 文档） | P99 ≤10ms | CES-PARSE |
| P14 | CES 校验（100KB 文档） | P99 ≤5ms | CES-VALID |
| P15 | CI 快速档总时长 | P95 ≤10 分钟 | XCU-CI |

### 4. 关键流程

1. 基准运行：benchmarks.yml（专用 runner，CPU 型号固定）；每次 main 合入跑全量基准。
2. 回归判定：与基线比对（统计稳健：5 次中位数，benchstat 风格）；回退 >10% 即 FAIL（预算表指标）。
3. 基线更新：性能改进合入时基线随版本更新（记录于 CHANGELOG）。

### 5. 验收标准

1. 预算表 15 项指标全部有基准用例与基线数据（M2 起逐步纳入，M6 全齐）。
2. 回归门禁运行中（回退 >10% FAIL 用例验证）。
3. 基准报告按版本归档（趋势可追溯）。

---

## 体系 XCU-DOC：文档体系

### 1. 体系定位与职责

四类文档的完整体系：用户文档、开发文档、规范文本、架构决策记录（ADR）。

### 2. 文件布局

```
docs/
  ARCHITECTURE.md    // 架构总纲（已建成）
  plan/              // 规划文档体系（00-08，本体系）
  spec/              // 规范文本
    ces-v1.md        // CES 规范（CES-SCHEMA 同步）
  adr/               // 架构决策记录
    0001-record.md   // 模板与索引
    0002-fail-fast-parse.md
    0003-token-level-roundtrip.md
    ...（各分册 ADR 引用项全部落档）
  user/              // 用户文档
    quickstart.md
    cli-reference.md // 命令树全参考（--help 生成）
    ces-guide.md     // CES 编写指南
    rules/           // 规则集目录（七语言，含出处链接）
  dev/               // 开发文档
    contributing.md
    module-dev-guide.md // 模块开发指南（含验收 DoD）
    plugin-dev-guide.md // 插件开发指南（WIT 接口+脚手架）
```

### 3. 文档同步铁律

1. cli-reference.md 由 --help 输出生成（CI 漂移检查）；ces-v1.md 与 CES-SCHEMA 生成物同步（漂移检查）。
2. 规则集目录与 builtin/ 规则同步（每条规则：id/说明/出处/正反例——由规则 YAML 生成）。
3. ADR 索引与正文一致（编号连续；状态 proposed/accepted/superseded）。

### 4. 验收标准

1. 四类文档全量就位（M6 验收项）；两处生成同步零漂移（CI）。
2. quickstart 全命令可复制执行（文档测试）；规则目录覆盖全部内置规则（≥460 条）。
3. ADR 全部落档（各分册引用的 ADR-0001~000N 编号连续无缺）。

---

## 附：横切体系合入总验收（七体系齐备后）

1. 全链演练：一次提交 → 七类门禁全绿 → 版本发布流程走通（签名资产+SBOM+公告）→ 混沌自测报告产出 → 文档同步零漂移——完整闭环验证。
2. 安全审计：T1-T10 攻击用例全拦截；外部安全审计完成（GA 前提）。
3. 性能预算 15 项全达标且回归门禁运行中。
4. 治理合规：版权头零缺失、三重水印验证、THIRD-PARTY 零漂移、双授权文本就位。
5. 七体系覆盖率与用例统计汇总报告入库；五段门禁+全量门禁全绿。
