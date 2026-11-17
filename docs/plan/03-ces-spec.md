# Abate L2 CES 规范层规划 — 四模块极致细节

- 文档编号：PLAN-03　状态：生效　上游：PLAN-00
- 覆盖模块：CES-SCHEMA / CES-PARSE / CES-VALID / CES-TCK
- 通用约束（四模块共同遵守）：
  1. CES（Chaos Experiment Specification）是开放中立规范：规范文本独立于实现演进，语义化版本（semver 2.0.0）；任何第三方实现通过 CES-TCK 即可声明 CES 合规。
  2. L2 只依赖 L0 内核类型（KRN-ALG 算子类型、KRN-OBS 探针类型）与标准库；禁止依赖 L1/L3-L7。
  3. YAML 与 JSON 双格式一等公民：同一文档两种编码必须可无损互转（往返测试强制）。
  4. 覆盖率：语句 ≥95%、分支 ≥85%（CES 属内核层质量标准）。

---

## 模块 CES-SCHEMA：规范模式

### 1. 模块定位与职责

定义 CES 文档的完整数据模型（Go 类型）、JSON Schema（用于外部工具校验与 IDE 补全）、字段文档（每个字段的语义/类型/必填性/默认值/约束）。它是 CES 规范的单一事实来源（Single Source of Truth）：Go 类型、JSON Schema、字段文档三者由代码生成器保持同步，禁止手工漂移。

### 2. 包结构与文件布局

```
ces/schema/
  doc.go           // 包文档：CES 规范总览与版本策略
  experiment.go    // Experiment 顶层类型（完整字段）
  steadystate.go   // SteadyState 稳态假设类型
  method.go        // Method 扰动方法类型（算子表达式）
  probe.go         // Probe 探针类型（四类）
  rollback.go      // Rollback 回滚策略类型
  schedule.go      // Schedule 调度类型（一次性/cron/持续）
  target.go        // Target 目标类型（模拟器目标/代码目标）
  meta.go          // Metadata 元数据类型
  version.go       // APIVersion 常量与版本协商
  gen/             // 代码生成器：Go 类型 → JSON Schema + 字段文档
    main.go
  dist/            // 生成产物（入库，供外部工具使用）
    ces-v1.schema.json
    fields.md
```

### 3. 核心类型与接口（必须实现）

- `Experiment`（顶层，kind: Experiment）：
  - `APIVersion string`：必填，格式 `ces.tencent.com/v1`（v1 起步；major 变更需规范 RFC）。
  - `Kind string`：必填，固定 `Experiment`。
  - `Metadata Metadata`：`Name`（必填，DNS 子域名校验）、`Namespace`（默认 default）、`Labels map[string]string`（键值约束校验）、`Annotations map[string]string`。
  - `Spec ExperimentSpec`：必填，包含以下全部子结构。
- `ExperimentSpec`：
  - `Target Target`：必填。`Kind`（simulator 代码孪生 / source 代码目标 / external 外部环境——external 首版保留字段但校验拒绝，防半成品）、`Selector`（目标选择器：包/函数/节点/进程键值对）。
  - `SteadyState []SteadyState`：必填，至少 1 条。
  - `Method Method`：必填。
  - `Probes []Probe`：可选（稳态探针之外的观测探针）。
  - `Rollback Rollback`：必填（策略：auto 自动 / manual 手动 / timeout 超时触发；默认 auto）。
  - `Schedule Schedule`：必填（once 一次性 / cron 表达式 / continuous 持续）。
  - `Timeout VirtualDuration`：可选，实验虚拟时间预算（默认继承调度器 Budget）。
  - `Seed uint64`：可选，确定性种子（缺省由 CLI 生成并回写报告）。
- `SteadyState`：`ProbeRef string`（引用探针 ID）、`Expectation Expectation`（value 精确值 / range 区间 / predicate 谓词表达式）、`Tolerance Tolerance`（absolute / relative / percent 三型）、`HoldFor VirtualDuration`（稳态持续时长，>0）。
- `Method`：`Expression string`（算子代数表达式，支持 ∘ 复合、⊗ 并发积、p⁻¹ 对偶、ε 恒等、括号分组；操作数为算子引用 `operator://<id>` 或内联算子定义）、`Operators []OperatorDef`（内联算子定义表）。
- `Probe`：`ID string`（文档内唯一）、`Kind ProbeKind`（http / shell / metric / prom 四类）、`Spec ProbeSpec`（各 Kind 独立 schema：http 含 url/method/headers/expect-status；shell 含 command/timeout/expect-exit-code；metric 含 query/expect；prom 含 promql/range/expect）、`Interval VirtualDuration`（采集间隔，>0）。
- `Rollback`：`Strategy Strategy`（auto/manual/timeout）、`Timeout VirtualDuration`（strategy=timeout 时必填）、`VerifySteadyState bool`（回滚后是否复验稳态，默认 true）。
- `Schedule`：`Once *OnceSpec` / `Cron *CronSpec`（五段 cron，虚拟时区）/ `Continuous *ContinuousSpec`（Duration 必填）——三选一互斥。
- `Version`：`Current() string`、`Negotiate(requested string) (string, error)`（版本协商：支持 v1；未知版本返回 ErrUnsupportedVersion 并列出支持清单）。

### 4. 关键算法与流程

1. 代码生成器：解析 Go 类型注解（struct tag + 自定义 directive）→ 产出 JSON Schema（draft 2020-12）与字段文档（fields.md 表格：字段/类型/必填/默认/约束/示例）；CI 门禁校验生成产物与源码同步（漂移即 FAIL）。
2. 往返保证：Experiment → JSON → YAML → JSON → Experiment 全链路无损（未知字段保留策略：v1 拒绝未知字段，防拼写错误静默吞掉）。
3. 版本策略：APIVersion major 变更走规范 RFC；minor 变更只增不改不删；废弃字段保留 2 个 minor 版本并文档标注。

### 5. 依赖与约束

- 依赖：仅标准库（encoding/json、regexp、time 不用于运行时仅用于类型）+ YAML 解析器源码级并入（gopkg.in/yaml.v3，Apache-2.0，vendor 入 third_party/yaml）。
- 被依赖：CES-PARSE（类型）、CES-VALID（字段约束）、CES-TCK（Schema）、CTL-CLI（文档示例）。
- 禁止：运行时反射驱动的动态 schema（性能与确定性）；任何 I/O。

### 6. 测试要求

- 单元测试：每个类型字段约束（必填/格式/枚举/互斥）正反例全覆盖；版本协商（支持/不支持/空）。
- 往返测试：JSON↔YAML 互转无损（含中文/特殊字符/大数值）。
- 生成器测试：产物与手工黄金文件比对（golden file 测试）。
- 模糊测试：随机变异合法文档，解析器不得 panic（fuzz 入 CI，60s 快速档）。

### 7. 验收标准（可量化）

1. 全部类型字段有 godoc + JSON Schema + fields.md 三处一致（CI 漂移检查 0 差异）。
2. 字段约束正反例用例 100% 通过；未知字段拒绝用例通过。
3. JSON↔YAML 往返无损测试 ≥50 组样本（含边界字符）100% 通过。
4. 模糊测试 60s 无 panic（语料入库）。
5. 覆盖率 ≥95%/85%；门禁全绿。

---

## 模块 CES-PARSE：规范解析

### 1. 模块定位与职责

把 CES 文档（YAML/JSON 字节流）解析为强类型 Experiment，并完成规范化（默认值填充、引用解析、表达式编译）。解析产物直接可执行（无需二次处理）。解析是确定性的：同字节流同解析结果（含错误信息）。

### 2. 包结构与文件布局

```
ces/parse/
  parse.go        // Parse 主入口：字节流 → Experiment
  normalize.go    // 规范化：默认值填充/引用解析/互斥检查
  expr.go         // 算子表达式词法与语法分析（手写递归下降）
  exprcompile.go  // 表达式 → 算子树（KRN-ALG Operator 组装）
  errors.go       // 结构化错误：位置（行/列）+ 字段路径 + 规则 ID
  parse_test.go
  testdata/       // 合法/非法样本库（≥100 份）
```

### 3. 核心类型与接口（必须实现）

- `Parse(data []byte, format Format) (*Experiment, error)`：主入口；format 为 yaml/json/auto（auto 按首字节嗅探：`{` JSON、其余 YAML）。
- `Normalize(exp *Experiment) (*Experiment, error)`：默认值填充（Rollback.Strategy=auto 等）、ProbeRef 引用解析（悬空引用报错）、Schedule 三选一互斥检查、Seed 缺省生成（记录在返回值中供回写）。
- `CompileExpr(expr string, ops []OperatorDef) (algebra.Operator, error)`：表达式 → 算子树。文法（EBNF）：
  - `expr := term ("∘" term)*`
  - `term := factor ("⊗" factor)*`
  - `factor := "(" expr ")" | operand ["⁻¹"]`
  - `operand := "ε" | operatorRef | inlineDef`
  - `operatorRef := "operator://" id`
- `ParseError`：`Line`、`Column`、`FieldPath`（如 spec.steadyState[0].tolerance）、`RuleID`（如 CES-PARSE-001）、`Message`。错误信息确定性（禁止包含 map 迭代序与指针地址）。

### 4. 关键算法与流程

1. 表达式分析：手写递归下降（禁止 yacc 生成器依赖）；运算符优先级：∘ 最低、⊗ 次之、⁻¹ 后缀最高；左结合。
2. 算子树组装：operatorRef 查 Operators 定义表（未定义报错）；内联定义直接构造；ε 构造恒等算子；p⁻¹ 调用 KRN-ALG Inverse（不可逆算子在此报错，错误带算子 ID 与位置）。
3. 错误恢复：首个错误即返回（fail-fast，错误信息含完整上下文）；不做多错误聚合（v1 简化决策，记录于 ADR-001）。
4. 深度防护：表达式嵌套深度上限 64（与 KRN-ALG 一致），超限报错。

### 5. 依赖与约束

- 依赖：CES-SCHEMA（类型）、KRN-ALG（算子树组装）、third_party/yaml（源码级并入）。
- 被依赖：CES-VALID（规范化后校验）、CTL-CLI（入口）。
- 禁止：任何 I/O（只接受字节流）；运行时网络；全局状态（解析器可重入，并发安全）。

### 6. 测试要求

- 单元测试：文法全覆盖（四运算/括号/后缀/引用/内联/ε/非法 token/深度超限），每类 ≥3 用例。
- 样本库测试：testdata ≥100 份（合法 60 / 非法 40），全部断言通过/失败与错误 RuleID。
- 确定性测试：同字节流 100 次解析，产物序列化哈希一致（含错误信息字节一致）。
- 并发测试：100 goroutine 并发解析不同文档，零竞争（-race 门禁）。
- 模糊测试：随机字节流/截断文档/超深嵌套，不 panic、不 OOM（输入上限 10MB，超限拒绝）。

### 7. 验收标准（可量化）

1. 文法用例与样本库 100% 通过；错误 RuleID 全覆盖（每条错误路径有唯一 RuleID）。
2. 解析确定性 100 次哈希一致；并发 100 goroutine -race 零报告。
3. 模糊测试 60s 无 panic 无 OOM（语料入库）。
4. 解析性能：100KB 典型文档 P99 ≤10ms（基准报告入库）。
5. 覆盖率 ≥95%/85%；门禁全绿。

---

## 模块 CES-VALID：规范校验

### 1. 模块定位与职责

对规范化后的 Experiment 做语义级校验（超越字段格式）：稳态可满足性、方法可执行性、探针配置合法性、回滚完备性、调度合法性、目标与算子域匹配。校验通过 = 文档可执行（运行期不再因配置错误失败）。

### 2. 包结构与文件布局

```
ces/validate/
  validate.go     // Validate 主入口：规则链编排
  rules.go        // 规则注册表（每条规则唯一 ID）
  r_steady.go     // 稳态规则组
  r_method.go     // 方法规则组
  r_probe.go       // 探针规则组
  r_rollback.go    // 回滚规则组
  r_schedule.go    // 调度规则组
  r_target.go      // 目标域匹配规则组
  validate_test.go
```

### 3. 核心类型与接口（必须实现）

- `Validate(exp *Experiment) ([]Violation, error)`：返回全部违规（聚合模式，与 PARSE fail-fast 互补）；exp 自身非法（未规范化）返回 error。
- `Violation`：`RuleID string`（如 CES-VALID-STEADY-001）、`FieldPath string`、`Message string`、`Severity Severity`（error 必须修复 / warning 建议修复）。
- 规则注册表：`Register(rule)`；每条规则实现 `Check(exp *Experiment) []Violation`；规则 ID 命名空间：STEADY/METHOD/PROBE/ROLLBACK/SCHEDULE/TARGET 六组。

### 4. 规则清单（首版必须全量实现，禁止一语而过）

**稳态规则组（STEADY）**：
- CES-VALID-STEADY-001：SteadyState 至少 1 条。
- CES-VALID-STEADY-002：ProbeRef 必须指向存在的探针（悬空引用）。
- CES-VALID-STEADY-003：HoldFor > 0 且 ≥ 探针 Interval（否则稳态永不可满足）。
- CES-VALID-STEADY-004：Expectation 与 Tolerance 类型匹配（range 期望配 absolute 容差等，类型矩阵校验）。
- CES-VALID-STEADY-005：predicate 表达式语法合法（编译通过）。

**方法规则组（METHOD）**：
- CES-VALID-METHOD-001：表达式非空且编译通过（复用 PARSE CompileExpr）。
- CES-VALID-METHOD-002：算子树中不可逆算子必须显式声明补偿策略（查 OperatorDef.Compensation）。
- CES-VALID-METHOD-003：算子树深度 ≤64。
- CES-VALID-METHOD-004：并发积 ⊗ 两算子目标键集合不相交（复用 KRN-ALG 冲突检测）。

**探针规则组（PROBE）**：
- CES-VALID-PROBE-001：ID 文档内唯一。
- CES-VALID-PROBE-002：Interval > 0。
- CES-VALID-PROBE-003：http 类 url 必须为模拟器内可达目标（Target.Selector 域内；external URL 拒绝——模拟器铁律）。
- CES-VALID-PROBE-004：shell 类 command 禁止真实副作用命令（白名单校验：模拟器内 shell 子集）。
- CES-VALID-PROBE-005：expect-status/expect-exit-code/expect 取值类型合法（如 HTTP 状态 100-599）。

**回滚规则组（ROLLBACK）**：
- CES-VALID-ROLLBACK-001：Strategy 枚举合法（auto/manual/timeout）。
- CES-VALID-ROLLBACK-002：strategy=timeout 时 Timeout > 0。
- CES-VALID-ROLLBACK-003：VerifySteadyState=true 时稳态规则组必须全部通过（回滚复验依赖稳态定义合法）。

**调度规则组（SCHEDULE）**：
- CES-VALID-SCHEDULE-001：三选一互斥（once/cron/continuous 恰好一个非空）。
- CES-VALID-SCHEDULE-002：cron 五段语法合法且虚拟时区有效。
- CES-VALID-SCHEDULE-003：continuous.Duration > 0。
- CES-VALID-SCHEDULE-004：once.At 为未来虚拟时间（相对实验开始）。

**目标域匹配规则组（TARGET）**：
- CES-VALID-TARGET-001：Target.Kind 枚举合法（simulator/source；external 拒绝并提示未支持）。
- CES-VALID-TARGET-002：M 族算子目标必须为 source（代码目标）；C 族算子目标必须为 simulator（系统目标）——算子族与目标域匹配矩阵校验。
- CES-VALID-TARGET-003：Selector 非空且键值格式合法。
- CES-VALID-TARGET-004：混合表达式 M⊗C 要求 Target 同时提供 source 与 simulator 选择器（双目标声明）。

### 5. 依赖与约束

- 依赖：CES-SCHEMA（类型）、CES-PARSE（表达式编译复用）、KRN-ALG（算子树检查）。
- 被依赖：CTL-CLI（执行前校验）、CES-TCK（合规校验项）。
- 禁止：任何 I/O；规则执行顺序确定（按规则 ID 字典序，保证违规列表确定性）。

### 6. 测试要求

- 单元测试：每条规则正反例各 ≥1（22 条规则 ≥44 用例）；规则 ID 唯一性测试。
- 聚合测试：多违规文档返回全部违规（顺序确定）。
- 确定性测试：同文档 100 次校验，违规列表序列化哈希一致。
- 样本库复用：testdata 非法样本 40 份全部被正确拒绝且 RuleID 匹配。

### 7. 验收标准（可量化）

1. 22 条规则全量实现；正反例用例 100% 通过；规则 ID 零冲突。
2. 违规列表确定性 100 次哈希一致。
3. 校验性能：100KB 典型文档 P99 ≤5ms（基准报告入库）。
4. 覆盖率 ≥95%/85%；门禁全绿。

---

## 模块 CES-TCK：兼容套件

### 1. 模块定位与职责

CES 技术兼容性套件（Technology Compatibility Kit）：第三方实现通过 TCK 全部用例即可声明『CES 合规』。TCK 是开放规范的中立裁判：用例只依据规范文本，不依赖 Abate 实现细节。它是『开放规范 CES』价值支柱的落地机制。

### 2. 包结构与文件布局

```
ces/tck/
  tck.go          // TCK 运行器：加载用例→驱动被测实现→比对期望
  case.go         // 用例模型：输入文档/期望判定/期望违规 RuleID
  driver.go       // 被测实现驱动接口（Driver 接口，第三方适配）
  report.go       // 合规报告生成（通过率/失败明细/合规徽章判定）
  cases/          // 用例库（数据驱动，YAML）
    parse/       // 解析用例（≥60）
    validate/     // 校验用例（≥44，与规则一一对应）
    roundtrip/   // 往返用例（≥50）
  tck_test.go
```

### 3. 核心类型与接口（必须实现）

- `Driver` 接口（第三方实现适配点）：`Parse(data []byte, format string) (ParseResult, error)`、`Validate(doc []byte) ([]Violation, error)`——返回结构对齐 TCK 用例模型（JSON 中立格式，不强制 Go 类型）。
- `Runner`：`Run(driver Driver) (Report, error)`——加载全部用例→逐条驱动→比对期望。
- `Report`：`Total`、`Passed`、`Failed`、`Skipped`、`Details []CaseResult`、`Compliant bool`（Compliant = Failed == 0 且 Skipped == 0）。
- 用例模型（YAML）：`id`（TCK-PARSE-001 格式）、`input`（文档字节）、`expect`（accept/reject + RuleID 列表 + 规范条款引用）。

### 4. 关键算法与流程

1. 用例数据驱动：用例即数据（YAML），新增用例零代码；用例 ID 与规范条款号交叉引用（每条用例必须引用规范文本章节）。
2. 中立比对：期望输出为 JSON 中立描述（accept/reject + 违规 RuleID 集合），不比对实现内部结构。
3. 合规判定：全通过才 Compliant；部分通过输出失败明细（用例 ID + 期望 vs 实际 + 规范条款引用）。
4. 版本策略：TCK 随规范 minor 版本只增用例；major 版本独立 TCK（tck-v1/tck-v2 并存）。

### 5. 依赖与约束

- 依赖：CES-SCHEMA（用例引用的规范条款）、标准库。
- 被依赖：外部第三方实现（经 Driver 适配）；XCU-REL（TCK 随版本发布）。
- 禁止：用例依赖 Abate 私有行为（用例评审清单：每条用例必须能从规范文本推导出期望）。

### 6. 测试要求

- 自举测试：Abate 自身实现（CES-PARSE/CES-VALID）作为第一个 Driver 跑 TCK，必须 100% 通过（吃自己的狗粮）。
- 用例质量测试：每条用例有规范条款引用（扫描校验）；用例 ID 唯一性。
- 中立性测试：构造一个最小第三方参考实现（测试夹具），通过 TCK（证明 TCK 不绑定 Abate 私有行为）。

### 7. 验收标准（可量化）

1. 用例库 ≥154 条（parse ≥60 + validate ≥44 + roundtrip ≥50），全部有规范条款引用。
2. Abate 自身 Driver 通过率 100%（CI 常驻）。
3. 最小参考实现夹具通过 TCK（中立性证明，报告入库）。
4. TCK 运行器自身覆盖率 ≥95%/85%；门禁全绿。

---

## 附：L2 合入总验收（四模块齐备后）

1. 端到端：一份含稳态+混合扰动（M⊗C）+探针+回滚的完整 CES YAML，经 Parse→Validate→（CLI 执行）全链路零错误；同种子 100 次双跑报告逐字节一致。
2. 规范发布物：ces-v1.schema.json + fields.md + 规范文本（docs/spec/ces-v1.md）三件套同步入库，CI 漂移检查零差异。
3. TCK 自举 100% 通过；合规徽章判定流程演示（一个第三方 Driver 示例）。
