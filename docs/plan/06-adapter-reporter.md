# Abate L5/L6 适配器与报告器规划 — 八模块极致细节

- 文档编号：PLAN-06　状态：生效　上游：PLAN-00
- 覆盖模块：ADP-CI / ADP-ITS / ADP-MSG / RPT-SARIF / RPT-HTML / RPT-JSON / RPT-JUNIT / RPT-OTEL
- 通用约束（八模块共同遵守）：
  1. L5/L6 全部以 WASM 插件形态交付（WIT 接口 adapter/reporter），同时提供官方内置实现（Go 原生，作为默认插件随引擎分发）——内置实现同样实现 WIT 接口语义，保证插件可替换内置（生态开放铁律）。
  2. 适配器与报告器是纯消费者：只读实验结果与违规数据，禁止反向影响内核执行。
  3. 输出确定性：同输入（实验报告/违规清单）产出字节确定（HTML 含生成时间戳除外——时间戳置于页脚且可关闭，CI 模式强制关闭）。
  4. 覆盖率：语句 ≥85%、分支 ≥75%（生态层标准）。
  5. 网络边界：适配器的外部网络调用（CI 平台 API、issue 系统、消息通道）只允许出现在适配器内部；报告器除 RPT-OTEL 外零网络。

---

## 模块 ADP-CI：CI 适配器

### 1. 模块定位与职责

把 Abate 门禁接入 CI 平台：GitHub Actions / GitLab CI 两种首版平台。职责：在 CI 流水线中驱动 abate check 与 abate run（实验），把退出码与结果工件（SARIF/JSON 报告）回传平台（PR 注释、检查运行、工件上传）。

### 2. 包结构与文件布局

```
adapter/ci/
  ci.go           // CIAdapter 接口与注册表
  github/        // GitHub Actions 适配
    actions.go   // 工作流动作入口（composite action）
    api.go       // GitHub API 客户端（检查运行/PR 注释/工件上传）
    sarif_upload.go // code scanning SARIF 上传
  gitlab/
    gitlab.go    // GitLab CI 作业入口
    api.go       // GitLab API 客户端（MR 讨论/工件）
  config.go      // 适配配置（平台检测/令牌来源/工件命名）
  ci_test.go
```

### 3. 核心类型与接口（必须实现）

- `CIAdapter` 接口：
  - `Detect(env Env) (Platform, bool)`：平台检测（GITHUB_ACTIONS/GITLAB_CI 环境变量嗅探）。
  - `ReportCheck(ctx, result CheckResult) error`：回传检查结果（GitHub：check run + 结论；GitLab：MR 报告 + 结论）。
  - `UploadArtifact(ctx, files []string) error`：上传工件（SARIF/JSON 报告归档）。
  - `Comment(ctx, summary Summary) error`：PR/MR 摘要评论（违规统计/存活分数/门禁结论）。
- `CheckResult`：ExitCode（0/1/2 三态语义对齐 LNG-CHK）/ ViolationSummary / MutationScore / GateVerdict。
- `Summary`：Markdown 文本（模板化：结论徽章 + 分级统计表 + Top 违规规则 + 报告工件链接）。
- 配置：令牌从环境变量注入（GITHUB_TOKEN/GITLAB_TOKEN），禁止配置文件存令牌。

### 4. 关键算法与流程

1. 平台检测：环境变量嗅探（确定性：固定检测顺序 GitHub → GitLab；双平台环境共存时显式配置优先）。
2. 摘要生成：模板渲染（Go text/template；数据为结构化 Summary，禁止字符串拼接 HTML——防注入）。
3. 失败语义：门禁 FAIL（exit 1）时 CI 作业 FAIL 且评论含失败规则条例；exit 2（内部错误）与 FAIL 严格区分（评论标注 internal error，不误导为代码违规）。
4. SARIF 上传：GitHub code scanning API（SARIF 2020-21 格式，由 RPT-SARIF 产出）。

### 5. 依赖与约束

- 依赖：RPT-SARIF/RPT-JSON（工件）、LNG-CHK（退出码语义）。
- 网络边界：GitHub/GitLab API 调用仅在此模块；令牌仅环境变量。
- 禁止：适配器内执行检查逻辑（纯回传层）；令牌落盘或入日志（脱敏门禁）。

### 6. 测试要求

- 单元：平台检测矩阵（有/无环境变量、双平台共存）；摘要模板渲染（含特殊字符转义用例）。
- 集成：GitHub Actions 实跑工作流（测试仓库，PR 触发 → 检查运行 + 评论 + 工件断言）；GitLab CI 同。
- 失败语义：exit 1/exit 2 两路径的评论内容断言。

### 7. 验收标准（可量化）

1. 平台检测矩阵用例 100% 通过；双平台共存显式配置用例通过。
2. GitHub Actions 与 GitLab CI 实跑集成测试通过（检查运行/评论/工件三项断言）。
3. 摘要模板特殊字符转义用例通过（XSS/注入防护）。
4. 覆盖率 ≥85%/75%；门禁全绿。

---

## 模块 ADP-ITS：跟踪适配器

### 1. 模块定位与职责

把检查违规与变异结果接入 issue 跟踪系统（首版：GitHub Issues）：违规自动建 issue（去重）、SURVIVED 变异体建改进 issue、等价变异体建复核 issue。

### 2. 包结构与文件布局

```
adapter/its/
  its.go          // ITSAdapter 接口
  github/
    issues.go     // GitHub Issues 客户端（建/查/关）
    dedupe.go     // 去重：指纹匹配已有 issue
  fingerprint.go  // 违规指纹（规则+文件+函数级归一）
  its_test.go
```

### 3. 核心类型与接口（必须实现）

- `ITSAdapter` 接口：`CreateIssue(ctx, issue Issue) (IssueRef, error)`、`FindOpen(ctx, fp Fingerprint) (IssueRef, bool)`、`CloseIssue(ctx, ref IssueRef, reason string) error`。
- `Issue`：Title/Body（Markdown 模板）/Labels（规则类别/严重级/语言）/Fingerprint。
- `Fingerprint`：SHA-256(规则 id + 文件路径 + 函数名 + 违规归一化摘要)——行号不参与指纹（行漂移不产生重复 issue）。
- 同步策略：全量对账——每次运行：新建（新指纹）/关闭（已修复：上次开 issue 本次无对应违规）/保留（仍存在）。

### 4. 关键算法与流程

1. 指纹归一：违规位置归一到函数级（AIR NodeID 定位函数）；同函数同规则多违规合并为一 issue（正文列全位置）。
2. 对账流程：拉取全部 open issue（标签过滤）→ 与本次结果指纹集合比对 → 三向操作（新建/关闭/保留）。
3. 限流：API 令牌桶（GitHub 5000/h）；对账报告输出操作统计。

### 5. 依赖与约束

- 依赖：LNG-CHK（违规数据）、KRN-MUT（变异判定）。
- 网络边界：GitHub API 仅此模块。
- 禁止：issue 正文含完整源码（只含违规片段 ±3 行上下文）。

### 6. 测试要求

- 单元：指纹归一（行漂移同指纹/函数重命名不同指纹）；三向对账逻辑。
- 集成：测试仓库实跑（建 3 issue → 修复 1 → 再跑 → 断言 1 关 2 留）。

### 7. 验收标准（可量化）

1. 指纹归一用例（行漂移/函数重命名/同函数合并）100% 通过。
2. 三向对账集成测试通过（操作统计断言）。
3. 覆盖率 ≥85%/75%；门禁全绿。

---

## 模块 ADP-MSG：消息适配器

### 1. 模块定位与职责

实验与门禁结果的消息通道通知（首版：Webhook 通用 + 飞书 + 企业微信）：门禁 FAIL、实验完成、存活分数低于阈值三类事件通知。

### 2. 包结构与文件布局

```
adapter/msg/
  msg.go         // MsgAdapter 接口与事件模型
  webhook/       // 通用 Webhook（POST JSON）
  feishu/        // 飞书机器人
  wecom/         // 企业微信机器人
  template.go    // 消息模板（Markdown/卡片）
  msg_test.go
```

### 3. 核心类型与接口（必须实现）

- `MsgAdapter` 接口：`Send(ctx, event Event) error`。
- `Event`：Kind（GATE_FAIL/EXPERIMENT_DONE/SCORE_BELOW）/Payload（结构化摘要）/Level（info/warn/critical）。
- 配置：通道列表（类型 + endpoint + 密钥）；密钥仅环境变量/受控配置，日志脱敏。
- 阈值：SCORE_BELOW 携带阈值（如存活分数 < 0.7 触发）。

### 4. 关键算法与流程

1. 事件路由：事件 → 订阅通道列表（配置驱动）→ 并行发送（单通道失败不阻断其他，错误聚合上报）。
2. 重试：5xx/超时重试 3 次（指数退避 1s/2s/4s，真实时间）；4xx 不重试（配置错误）。
3. 模板：三通道各自模板（Webhook JSON 结构化；飞书/企微 Markdown 卡片）；模板变量白名单（禁止任意 HTML 注入）。

### 5. 依赖与约束

- 依赖：KRN-VER（分数）、LNG-CHK（门禁结果）。
- 网络边界：三通道 endpoint 仅此模块。
- 禁止：消息含源码片段（只含统计与链接）。

### 6. 测试要求

- 单元：事件路由矩阵；重试策略（5xx 重试/4xx 不重试）；模板变量白名单。
- 集成：本地 mock server（Webhook）实跑断言载荷；飞书/企微以官方测试机器人实发（人工验收记录）。

### 7. 验收标准（可量化）

1. 路由矩阵与重试用例 100% 通过。
2. Webhook mock 集成断言载荷结构通过。
3. 覆盖率 ≥85%/75%；门禁全绿。

---

## 模块 RPT-SARIF：SARIF 报告

### 1. 模块定位与职责

检查违规与变异结果的 SARIF 2.1.02 标准输出：供 GitHub code scanning、IDE、CI 平台消费的标准格式。

### 2. 包结构与文件布局

```
report/sarif/
  sarif.go       // SARIF 文档模型（2.1.02 子集）
  build.go       // 违规/变异结果 → SARIF 构建
  validate.go    // 自校验（schema + 语义）
  sarif_test.go
```

### 3. 核心类型与接口（必须实现）

- `Build(violations []Violation, mutations []MutationResult) (*SarifDoc, error)`：主入口。
- SarifDoc 模型（2.1.02 必备子集）：version/$schema/runs[]（tool.driver{name,version,rules[]}、results[]{ruleId,level,message,locations[]}、artifacts[]、columnKind: unicodeCodePoints）。
- 规则元数据：每条规则入 rules[]（id/shortDescription/helpUri——指向官方规范出处）。
- 变异结果映射：SURVIVED → level note + 属性区（score 明细）；KILLED 不入 results（汇总入属性区）。

### 4. 关键算法与流程

1. 构建管线：违规分组（按规则）→ rules[] 生成 → results[] 生成（位置：文件 URI + 行列，1 起，unicodeCodePoints）→ artifacts 注册。
2. 自校验：构建后内置 schema 校验（JSON Schema draft 2020-12）+ 语义校验（ruleId 必须存在于 rules[]；位置不越界）——校验失败返回错误而非产出坏文档。
3. 确定性：同输入产出字节一致（键序固定：结构体序，非 map 序）。

### 5. 依赖与约束

- 依赖：LNG-CHK/KRN-MUT（数据类型）、标准库 encoding/json。
- 禁止：网络；非标准扩展字段（保持最大兼容）。

### 6. 测试要求

- 单元：构建管线（违规/变异/混合三场景）；确定性（双跑字节一致）。
- 合规：输出过 SARIF 官方 JSON Schema 校验 + Microsoft sarif-sdk 校验器（双校验）。
- 边界：空违规（合法空文档）、超长消息、Unicode 文件名/路径。

### 7. 验收标准（可量化）

1. 双官方校验器校验通过（报告入库）。
2. 确定性双跑字节一致；空/边界用例通过。
3. 覆盖率 ≥85%/75%；门禁全绿。

---

## 模块 RPT-HTML：HTML 报告

### 1. 模块定位与职责

人读交互式报告：违规清单（可筛选/排序）、变异矩阵（逐变异体判定）、实验时间线（HPOV 状态流转）、存活分数仪表。单文件自包含（内联 CSS/JS，零外部依赖，可离线打开）。

### 2. 包结构与文件布局

```
report/html/
  html.go        // HTMLDoc 构建主入口
  layout.go      // 页面骨架（模板）
  section_violation.go // 违规区
  section_mutation.go  // 变异矩阵区
  section_timeline.go  // 实验时间线区
  section_score.go     // 分数仪表区
  assets/       // 内联资源（CSS/JS 源文件）
  html_test.go
```

### 3. 核心类型与接口（必须实现）

- `Build(report FullReport, opts HTMLOptions) ([]byte, error)`：主入口；FullReport = 违规 + 变异 + 实验档案元数据 + 分数。
- HTMLOptions：EmbedSource（违规片段嵌入开关，默认开 ±3 行）/Timestamp（页脚时间戳，CI 模式强制关）。
- 交互能力（内联 JS，无框架依赖）：违规表筛选（语言/严重级/规则）与排序；变异矩阵展开（逐变异体：位置/算子/判定/击杀测试）；时间线折叠展开。
- 安全：全部动态内容 HTML 转义（html/template 自动转义，禁用 text/template 渲染用户数据）；源码片段语法高亮为纯 CSS 类标注（无 eval）。

### 4. 关键算法与流程

1. 渲染管线：数据模型 → 模板树（html/template，编译一次）→ 字节输出；模板全部自动转义（XSS 防护）。
2. 变异矩阵：行=变异体，列=测试（判定矩阵 KILLED 标记击杀测试）；大数据量分页（>1000 变异体分页渲染，单页 100）。
3. 确定性：除页脚时间戳外字节确定；CI 模式（Timestamp 关）双跑字节一致。

### 5. 依赖与约束

- 依赖：LNG-CHK/KRN-MUT/KRN-VER/SIM-REC（档案元数据）。
- 禁止：外部 CDN/字体/脚本（完全自包含）；网络。

### 6. 测试要求

- 单元：渲染管线（四区各自数据模型 → 输出断言）；转义用例（脚本注入载荷转义断言）。
- 确定性：CI 模式双跑字节一致。
- 规模：10 万违规/1 万变异体渲染 P99 ≤5s（基准入库）。

### 7. 验收标准（可量化）

1. 四区渲染与转义用例 100% 通过。
2. CI 模式确定性双跑一致。
3. 规模性能达标；覆盖率 ≥85%/75%；门禁全绿。

---

## 模块 RPT-JSON：JSON 报告

### 1. 模块定位与职责

机器可读规范报告：全量结果（违规/变异/实验元数据/分数/档案引用）的稳定 JSON Schema（版本化），供下游工具与 CI 消费。

### 2. 包结构与文件布局

```
report/json/
  json.go        // Build 主入口
  schema.go      // JSON Schema 定义与版本
  json_test.go
```

### 3. 核心类型与接口（必须实现）

- `Build(report FullReport) ([]byte, error)`。
- 文档结构：schemaVersion（semver）/generated（确定性元数据：种子/指纹/工具版本）/violations[]/mutations[]/experiments[]/scores/summary。
- Schema：随版本发布 report-schema-v1.json；字段废弃走 2 个 minor 窗口（对齐内核 API 政策）。
- 确定性：同输入字节一致（结构体序，无 map 序；时间戳仅 generated.stable 元数据，不含真实时间）。

### 4. 关键算法与流程

1. 构建管线：FullReport → 结构化模型 → json.Marshal（结构体字段序固定）。
2. 自校验：构建后过自带 schema 校验（防模型与 schema 漂移；CI 漂移检查）。

### 5. 依赖与约束

- 依赖：数据类型同 RPT-HTML。禁止：网络；schema 漂移（CI 强制）。

### 6. 测试要求

- 单元：构建管线；schema 自校验；确定性双跑。
- 兼容：v1.x 旧文档新代码可读（向后兼容用例）。

### 7. 验收标准（可量化）

1. schema 自校验与 CI 漂移检查零差异。
2. 确定性双跑字节一致；向后兼容用例通过。
3. 覆盖率 ≥85%/75%；门禁全绿。

---

## 模块 RPT-JUNIT：JUnit 报告

### 1. 模块定位与职责

JUnit XML 输出（CI 平台测试报告消费标准）：把变异体判定映射为测试用例（SURVIVED=失败用例，KILLED=通过用例），门禁结论映射为 testsuite。

### 2. 包结构与文件布局

```
report/junit/
  junit.go       // Build 主入口
  junit_test.go
```

### 3. 核心类型与接口（必须实现）

- `Build(report MutationReport) ([]byte, error)`。
- 映射：testsuite name=abate-mutation，tests=变异体总数，failures=SURVIVED 数，skipped=EQUIVALENT+TIMEOUT 数；testcase 逐变异体（name=算子:位置；SURVIVED 生成 failure 节点含消息与击杀尝试摘要）。
- 遵循 JUnit XML 常见方言（GitHub/GitLab 消费子集：testsuite/testcase/failure/skipped 属性集）。

### 4. 关键算法与流程

1. 映射管线：变异结果 → testcase 列表（确定性排序：位置 → 算子）→ XML 序列化（encoding/xml，转义正确）。
2. 确定性：同输入字节一致。

### 5. 依赖与约束

- 依赖：KRN-MUT 数据类型。禁止：网络；方言扩展属性（保持通用兼容）。

### 6. 测试要求

- 单元：映射（四类判定各用例）；XML 转义（消息含 <>&" 用例）。
- 消费验证：GitHub Actions 测试报告注解与 GitLab 单元测试报告实跑解析通过。

### 7. 验收标准（可量化）

1. 四类判定映射与转义用例 100% 通过。
2. 双平台实跑解析通过（记录入库）。
3. 确定性双跑一致；覆盖率 ≥85%/75%；门禁全绿。

---

## 模块 RPT-OTEL：OTel 报告

### 1. 模块定位与职责

OpenTelemetry 导出：实验执行 trace（HPOV 状态 span）、观测指标（探针值 metric）、判定事件（log）。供可观测平台（Jaeger/Prometheus/Grafana）消费。

### 2. 包结构与文件布局

```
report/otel/
  otel.go        // Exporter 构建（OTLP 导出器封装）
  trace.go       // HPOV 状态 → span 树
  metrics.go     // 观测流 → metric 序列
  logs.go        // 判定/违规 → log 记录
  otel_test.go
```

### 3. 核心类型与接口（必须实现）

- `NewExporter(ctx, cfg OTLPConfig) (*Exporter, error)`：OTLP gRPC/HTTP 导出器（endpoint/协议/头/批量参数）。
- `Exporter.Export(ctx, report FullReport, archive *sim/record.Archive) error`：三流导出。
- trace 映射：实验根 span（实验指纹为 trace 属性）→ HPOV 状态 span（状态名/虚拟时间区间）→ 扰动 span（算子 id/参数哈希）→ 探针 span。
- metric 映射：探针观测 → gauge（虚拟时间为时间轴属性）；存活分数 → gauge。
- log 映射：判定事件（VERDICT）与违规（VIOLATION）→ log record（severity 映射：VIOLATED→WARN，SURVIVED→ERROR）。

### 4. 关键算法与流程

1. 批量导出：span/metric/log 三流独立批量（默认 512 条/批或 5s 真实时间冲刷）；导出失败重试 3 次后丢弃并计数（不阻断实验）。
2. 虚拟时间语义：OTel 时间戳为真实导出时间；虚拟时间作为 span/log 属性（virtualTime）传递——保真不混淆。

### 5. 依赖与约束

- 依赖：OTel Go SDK 源码级并入（third_party/otel，Apache-2.0，五步并入法：剥离 collector，只留 SDK 导出路径）。
- 网络边界：OTLP endpoint 仅此模块；导出失败不影响实验结果（旁路语义）。
- 禁止：导出阻塞实验主流程（异步队列上限 10 万条，满则丢弃计数）。

### 6. 测试要求

- 单元：三流映射（各 ≥3 用例）；批量与重试策略；队列满丢弃计数。
- 集成：本地 OTLP collector（容器外进程，测试夹具）接收断言（span 树结构/metric 值/log 记录）。
- 旁路语义：collector 不可达时实验正常完成（断言结果不受影响）。

### 7. 验收标准（可量化）

1. 三流映射与批量/重试用例 100% 通过。
2. 本地 collector 集成断言通过（结构化校验）。
3. 旁路语义用例通过（不可达不影响实验）。
4. 覆盖率 ≥85%/75%；门禁全绿。

---

## 附：L5/L6 合入总验收（八模块齐备后）

1. 端到端：一次完整实验（检查+变异+混沌）产出五格式报告（SARIF/HTML/JSON/JUnit/OTel）+ CI 回传（GitHub Actions 实跑：检查运行+评论+工件+code scanning）全链路演示。
2. 标准合规：SARIF 过双官方校验器；JUnit 过双平台解析；JSON schema 零漂移。
3. 确定性：SARIF/JSON/JUnit 双跑字节一致；HTML CI 模式双跑一致。
4. 八模块覆盖率汇总报告入库；五段门禁+基准回归全绿。
