# Abate L0 内核规划 — HPOV 内核六模块极致细节

- 文档编号：PLAN-01　状态：生效　上游：PLAN-00
- 覆盖模块：KRN-ALG / KRN-SCH / KRN-OBS / KRN-VER / KRN-MUT / KRN-INJ
- 通用约束（六模块共同遵守）：
  1. 仅依赖 Go 标准库与内核自有包；禁止 import L1-L7 任何包（depguard 门禁强制）。
  2. 全部公开类型与函数必须有 Go doc 注释（godot 句号结尾门禁）。
  3. 全部随机性必须经 SIM-RNG 注入；禁止直接使用 math/rand 全局源与时间读取（虚拟时钟门禁）。
  4. 覆盖率：语句 ≥95%、分支 ≥85%；确定性回放测试（同种子双跑哈希一致）为合入前置条件。

---

## 模块 KRN-ALG：扰动代数

### 1. 模块定位与职责

定义扰动算子的类型系统与四种代数运算（复合 ∘、并发积 ⊗、对偶 p⁻¹、恒等 ε），并提供代数律性质校验器。它是整个引擎的数学地基：变异（M 族）与混沌（C 族）算子都实现同一 Operator 接口，调度器与验证器只面向代数结构编程，不感知具体算子语义。

### 2. 包结构与文件布局

```
kernel/algebra/
  operator.go        // Operator 接口、Family 枚举、运算错误类型
  composite.go       // 复合运算 Compose
  concurrent.go      // 并发积运算 Concurrent
  dual.go            // 对偶（回滚）运算 Inverse
  identity.go        // 恒等算子 Identity
  laws.go            // 代数律性质校验器（结合律/单位元/对偶补偿）
  laws_test.go       // 性质测试（1000 组随机组合）
  doc.go             // 包文档：代数定义与记号表
```

### 3. 核心类型与接口（必须实现）

- `Operator` 接口：
  - `ID() OperatorID`：全局唯一算子标识（字符串，命名规范：族前缀-名称-版本，如 m-arith-replace-v1、c-network-delay-v1）。
  - `Family() Family`：返回 M 族或 C 族。
  - `Params() Params`：算子参数（有序键值，可序列化，参与确定性哈希）。
  - `Inverse() (Operator, error)`：返回对偶算子；不可逆算子返回 ErrNotInvertible（仅允许出现在显式声明不可逆的 C 族算子，且必须同时提供补偿策略）。
  - `Fingerprint() [32]byte`：确定性指纹（ID+族+参数规范序列化后 SHA-256），用于回放哈希与去重。
- `Compose(p1, p2 Operator) (Operator, error)`：复合，语义为先施加 p2 再施加 p1；校验两算子目标域兼容（不兼容返回 ErrIncompatible）。
- `Concurrent(p1, p2 Operator) (Operator, error)`：并发积；要求两算子无目标冲突（冲突检测：目标键集合相交即拒绝，返回 ErrTargetConflict）。
- `Identity(target TargetKey) Operator`：恒等算子。
- 代数律校验器 `Laws`：
  - `CheckAssociativity(p1, p2, p3) bool`：Compose(p1, Compose(p2, p3)) ≡ Compose(Compose(p1, p2), p3)。
  - `CheckIdentity(p) bool`：Compose(p, ε) ≡ p ≡ Compose(ε, p)。
  - `CheckDual(p) bool`：施加 p 后施加 p⁻¹，目标状态回到基线（容差内）。

### 4. 关键算法与流程

1. 指纹计算：参数按键名排序后规范序列化（禁止 map 迭代序泄漏），保证同参数同指纹（确定性铁律）。
2. 复合展开：Compose 结果是惰性包装节点，Apply 时按右结合顺序展开；深度上限 64（防算子炸弹），超限返回 ErrTooDeep。
3. 并发积展开：Apply 时两算子在同一虚拟时间片并发施加，事件序号由调度器全序化。

### 5. 依赖与约束

- 依赖：仅标准库（crypto/sha256、sort、errors、fmt）。
- 被依赖：KRN-SCH（调度算子）、KRN-VER（验证对偶补偿）、KRN-MUT/KRN-INJ（算子实现方）。
- 禁止：任何 I/O、任何时间读取、任何全局状态。

### 6. 测试要求

- 单元测试：表驱动，覆盖正常/边界（空参数、最大深度、不可逆算子）/错误（目标冲突、不兼容）三类路径。
- 性质测试：随机生成 1000 组算子组合，验证三条代数律全部成立。
- 确定性测试：同种子构造同算子树 100 次，指纹逐字节一致。

### 7. 验收标准（可量化）

1. Operator 接口六方法全部实现并有 godoc；四运算+三律校验器全部就位。
2. 性质测试 1000 组随机组合 100% 通过。
3. 覆盖率语句 ≥95%、分支 ≥85%。
4. 指纹确定性测试 100 次双跑哈希一致。
5. 无 TODO/FIXME（godox 门禁）；五段门禁全绿。

---

## 模块 KRN-SCH：确定性调度器

### 1. 模块定位与职责

以虚拟时钟为唯一时间源、以全序事件队列为核心的离散事件仿真（DES）调度器。负责实验的确定性执行：同种子同配置必然产生逐字节一致的事件流。它是『稳』字的直接落点。

### 2. 包结构与文件布局

```
kernel/scheduler/
  event.go          // Event 类型：虚拟时间+全序序号+载荷
  queue.go          // 优先队列：按（虚拟时间, 序号）字典序出队
  scheduler.go      // 调度循环：取事件→执行→派生事件入队
  budget.go         // 执行预算：事件数/虚拟时长/算子成本三重上限
  scheduler_test.go // 确定性与吞吐测试
```

### 3. 核心类型与接口（必须实现）

- `Event`：`At VirtualTime`（纳秒整数，禁止浮点）、`Seq uint64`（全序序号，单调递增）、`Kind EventKind`（tick/probe/apply/rollback/timeout）、`Payload`。
- `Queue`：`Push(Event)`、`Pop() (Event, bool)`；排序键（At, Seq）严格全序，同键禁止（入队时校验，冲突返回 ErrSeqCollision）。
- `Scheduler`：
  - `Run(ctx, Experiment) Result`：主循环。
  - `Now() VirtualTime`：当前虚拟时间。
  - `Emit(Event)`：派生事件入队（供算子与探针回调使用）。
- `Budget`：`MaxEvents`（默认 1,000,000）、`MaxVirtualDuration`（默认 24h 虚拟时间）、`MaxCost`（算子成本和上限）；任一超限即终止实验并标记 BUDGET_EXCEEDED（防失控实验）。

### 4. 关键算法与流程

1. 主循环：Pop 事件 → 推进虚拟时钟到 At → 按 Kind 分发执行 → 执行产物（观测事件/派生事件）入队 → 循环直至队列空或预算耗尽或稳态判定完成。
2. 全序化铁律：所有并发（goroutine）只允许出现在模拟器边界；进入队列前必须携带调度器分配的 Seq，保证事件流可全序回放。
3. 时钟推进：单调不回退；回滚事件也只允许在未来虚拟时间生效（对偶补偿的最终一致语义）。

### 5. 依赖与约束

- 依赖：SIM-CLK（时钟接口）、SIM-RNG（随机流）、KRN-ALG（算子类型）。
- 被依赖：KRN-VER（消费观测事件）、CTL-CLI（驱动实验）。
- 禁止：真实时间读取（time.Now 禁用，vet+门禁扫描）、浮点时间、map 迭代序参与调度决策。

### 6. 测试要求

- 单元测试：队列全序性（10 万事件随机入队，出队严格有序）、预算三上限各自触发。
- 确定性测试：同种子同实验双跑，事件流 SHA-256 一致；100 次循环。
- 吞吐基准：单进程 ≥10 万事件/秒（M2 性能预算首测项）。

### 7. 验收标准（可量化）

1. 事件流确定性：100 次双跑哈希一致率 100%。
2. 队列全序测试 10 万事件 0 冲突（Seq 冲突被拒绝的负例单独覆盖）。
3. 预算三上限各有触发用例且终止码正确。
4. 吞吐基准 ≥10 万事件/秒（基准报告入库）。
5. 覆盖率 ≥95%/85%；门禁全绿。

---

## 模块 KRN-OBS：观测管道

### 1. 模块定位与职责

把探针数据、指标、追踪三类观测源归一化为带全序序号的观测事件流，供验证器消费与 SIM-REC 录制。它是『准』字（观测零漏报）的直接落点。

### 2. 包结构与文件布局

```
kernel/observe/
  observation.go    // Observation 类型：虚拟时间+序号+探针ID+值+状态
  probe.go          // Probe 接口与注册表
  sink.go           // Sink 接口：验证器/录制器/报告器统一消费口
  pipeline.go       // 管道：归一化→序号分配→扇出
  pipeline_test.go
```

### 3. 核心类型与接口（必须实现）

- `Observation`：`At VirtualTime`、`Seq uint64`（调度器分配）、`ProbeID string`、`Value Value`（数值/布尔/字符串/直方图四型）、`Status ProbeStatus`（ok/degraded/failed/timeout）。
- `Probe` 接口：`ID() string`、`Kind() ProbeKind`（http/shell/metric/prom）、`Collect(ctx, at VirtualTime) (Value, Status)`——在模拟器内执行，禁止真实网络。
- `Sink` 接口：`Observe(Observation) error`；实现方：KRN-VER（断言）、SIM-REC（录制）、RPT-*（报告，经端口注入）。
- `Pipeline`：`Register(Probe)`、`Attach(Sink)`、`Run(ctx)`——每轮 tick 按注册序（稳定排序）采集全部探针，归一化后扇出到全部 Sink。

### 4. 关键算法与流程

1. 采集顺序稳定：探针按 ID 字典序采集，消除注册顺序不确定性。
2. 扇出语义：Sink 消费失败不阻断其他 Sink（错误记录进观测流自身，作为 meta 观测）。
3. 超时：探针采集超过虚拟时间预算即标记 timeout 状态（不阻塞调度循环）。

### 5. 依赖与约束

- 依赖：SIM-CLK、KRN-SCH（Seq 分配）。
- 被依赖：KRN-VER、SIM-REC、RPT-*（经端口）。
- 禁止：真实 I/O；探针实现一律面向模拟器（真实环境探针属于生态层 injector 插件职责）。

### 6. 测试要求

- 单元测试：四型值归一化、超时路径、Sink 失败隔离。
- 确定性测试：同种子双跑观测流哈希一致。
- 集成测试：与 KRN-SCH 联调，探针采集序与事件序一致。

### 7. 验收标准（可量化）

1. 四型 Value 全部有编解码往返测试（无损）。
2. 观测流确定性 100 次双跑一致。
3. Sink 失败隔离用例：单 Sink 报错不影响其余 Sink 与调度循环。
4. 覆盖率 ≥95%/85%；门禁全绿。

---

## 模块 KRN-VER：验证器

### 1. 模块定位与职责

消费观测流，输出实验判定：稳态断言（HELD/VIOLATED）、变异判定（KILLED/SURVIVED/EQUIVALENT）、存活分数（mutation score）。它是 HPOV 的 V，也是变异测试的核心价值出口。

### 2. 包结构与文件布局

```
kernel/verify/
  steady.go         // 稳态断言：探针+容差+持续时间
  verdict.go        // Verdict 类型与判定状态机
  score.go          // 存活分数与等价性分析
  verify.go         // Verify 主入口：观测流→判定
  verify_test.go
```

### 3. 核心类型与接口（必须实现）

- `SteadyState`：`ProbeID`、`Expectation`（值/区间/谓词）、`Tolerance`（绝对/相对/百分比）、`HoldFor VirtualTime`（稳态必须持续的最短虚拟时长）。
- `Verdict`：`Outcome`（HELD/VIOLATED/INCONCLUSIVE）、`Reason`（首个违规观测的 Seq 与期望/实际值）、`Timeline`（违规时间线）。
- `MutationVerdict`：`KILLED`（测试击杀变异体）/`SURVIVED`（测试未击杀）/`EQUIVALENT`（等价变异体，经等价性分析判定）/`TIMEOUT`（超预算）。
- `Score`：`Total`、`Killed`、`Survived`、`Equivalent`、`Timeout`、`Score = Killed / (Total - Equivalent)`（等价体不计分母，业界标准口径）。
- `Verify(obs []Observation, ss []SteadyState) []Verdict`：纯函数，无状态。

### 4. 关键算法与流程

1. 稳态判定：观测值在容差内且持续满 HoldFor 才判 HELD；任一观测越限即 VIOLATED 并记录首个违规点。
2. 变异判定：变异体运行测试套件（模拟器内）→ 测试失败即 KILLED；全过即 SURVIVED。
3. 等价性分析（三段法）：静态等价（AIR 结构等价检查）→ 动态等价（同输入观测流逐字节比对）→ 标记 EQUIVALENT（存疑等价体进入人工复核队列，不自动计分）。

### 5. 依赖与约束

- 依赖：KRN-OBS（观测类型）、LNG-AIR（等价性静态检查，经端口注入，内核只定义接口）。
- 被依赖：RPT-*（报告判定结果）、CTL-*（展示）。
- 禁止：Verify 内任何 I/O 与时间读取；判定必须纯函数化（可回放重判）。

### 6. 测试要求

- 单元测试：容差三型（绝对/相对/百分比）× 边界（恰好等于容差边界）全组合；HoldFor 边界（差一 tick）。
- 判定测试：四类 MutationVerdict 各有构造用例；分数计算含等价体剔除口径。
- 回放重判测试：录制观测流重放重判，Verdict 逐字节一致。

### 7. 验收标准（可量化）

1. 容差×边界全组合表驱动用例 100% 通过。
2. 四类变异判定与分数口径与 PIT 参考实现对拍一致（对拍报告入库）。
3. 回放重判 100 次逐字节一致。
4. 覆盖率 ≥95%/85%；门禁全绿。

---

## 模块 KRN-MUT：变异核（PIT 算法源码级移植）

### 1. 模块定位与职责

M 族算子的生成与调度引擎：变异点发现、变异体生成、覆盖率引导调度、最小化（定位首个击杀测试）。算法源自 PITest，按五步并入法移植为 Go 内核自有代码。

### 2. 五步并入法执行细节（逐项）

1. **源码 vendor**：PIT 源码固定版本入 `third_party/pit/`（只读参考基线，记录上游 commit 哈希与许可证审查结论：PIT 为 Apache-2.0，可并入）。
2. **二次更改**：核心算法 Go 重写并确定性化——变异点选择由 SIM-RNG 种子驱动（PIT 原生非确定点全部替换）；调度顺序全序化。
3. **优化**：只移植四条核心算法线（变异算子族/覆盖率引导/最小化/优先级排序），裁剪报告、CLI、Maven 集成等非核心路径。
4. **深度集成**：算子族实现 KRN-ALG 的 Operator 接口（M 族）；变异点发现面向 AIR（LNG-AIR）而非 JVM 字节码——这是与 PIT 的架构差异点：Abate 变异作用在源码 AST 层，天然多语言。
5. **封装**：对外仅暴露 `Engine` 类型与 `Run(ctx, Target, Suite) Report`；PIT 概念（MutationDetails 等）不得出现在内核 API。

### 3. 包结构与文件布局

```
kernel/mutation/
  engine.go        // Engine：编排变异实验
  operators.go     // M 族算子注册表（算符替换/边界值/常量扰动/控制流/返回值/否定条件/删除调用）
  discover.go      // 变异点发现：AIR 遍历→候选点
  schedule.go      // 覆盖率引导调度：按测试覆盖排序变异点（PIT 线程优先级算法移植）
  minimize.go      // 最小化：二分定位首个击杀测试（PIT 最小化算法移植）
  engine_test.go
```

### 4. 核心类型与接口（必须实现）

- `Engine`：`Run(ctx, Target, Suite) (Report, error)`——Target 为 AIR 单元（函数/方法），Suite 为测试选择器。
- `MutationPoint`：`AIRNodeRef`（节点引用）、`Operator OperatorID`、`Priority`（覆盖率引导优先级）。
- `Report`：逐变异体 `MutationVerdict` 列表 + 汇总 `Score` + 确定性指纹（种子+目标指纹+套件指纹）。
- M 族七类基础算子（首版必须全量）：算符替换（+↔-、*↔/、&&↔||）、边界值扰动（<↔<=、>↔>=）、常量扰动（+1/-1/置零）、控制流扰动（删除 break/continue）、返回值扰动（true↔false、nil↔非 nil）、否定条件、删除调用。

### 5. 依赖与约束

- 依赖：KRN-ALG、KRN-SCH、KRN-VER、LNG-AIR（经端口接口）。
- 禁止：JVM 运行时依赖（架构上不需要）；PIT 专有类型泄漏（API 审查项）。

### 6. 测试要求

- 单元测试：七类算子各有正/反例；调度优先级排序确定性。
- 对拍测试（关键验收）：构造 Go 基准目标集，Abate 移植算子与 PIT 参考实现在同一变异点集合上判定结果一致率 100%（PIT 侧经其官方 CLI 产出，对拍脚本与报告入库）。
- 确定性测试：同种子双跑 Report 哈希一致。

### 7. 验收标准（可量化）

1. 七类基础算子全量落地且注册表可枚举。
2. PIT 对拍一致率 100%（对拍报告入库，差异逐条说明）。
3. 覆盖率引导调度与最小化各有端到端用例。
4. 同种子 100 次双跑 Report 哈希一致。
5. 覆盖率 ≥95%/85%；门禁全绿；无 PIT 专有类型出现在内核 API（审查清单勾选）。

---

## 模块 KRN-INJ：注入核（chaosblade 源码二次开发）

### 1. 模块定位与职责

C 族算子的执行引擎：把混沌故障（网络/资源/时钟/进程/依赖）以确定性方式注入模拟器目标，并提供对偶回滚。基于 chaosblade 源码二次开发（同为 Go，天然亲和）。

### 2. 五步并入法执行细节（逐项）

1. **源码 vendor**：chaosblade 固定版本源码入 `third_party/chaosblade/`（记录上游 commit 哈希；许可证 Apache-2.0，可并入）。
2. **二次更改**：执行模型确定性化——注入动作改为面向 SIM-NET/SIM-FS/SIM-SCHED 的模拟器调用（原真实系统调用路径全部改道）；随机参数改由 SIM-RNG 种子驱动。
3. **优化**：裁剪 CLI 外壳、渠道/云厂商集成、exec 拉起路径；收敛为纯库形态。
4. **深度集成**：注入动作直译 C 族 Operator（实现 KRN-ALG 接口）；回滚实现 Inverse（对偶运算）；场景模型映射为 CES 规范词汇。
5. **封装**：对外仅暴露 `Injector` 类型与 `Inject(ctx, op Operator) error` / `Rollback(ctx, op Operator) error`；chaosblade 专有概念（Model/ExpModel 等）不得泄漏。

### 3. 包结构与文件布局

```
kernel/inject/
  injector.go      // Injector：注入/回滚入口
  network.go       // 网络类：延迟/分区/丢包/限带宽（落 SIM-NET）
  resource.go      // 资源类：CPU/内存/IO/磁盘压力（落 SIM-SCHED/SIM-FS）
  clock.go         // 时钟类：偏移/抖动（落 SIM-CLK）
  process.go       // 进程类：终止/挂起（落 SIM-SCHED）
  dependency.go    // 依赖类：依赖故障（错误注入/超时）
  injector_test.go
```

### 4. 核心类型与接口（必须实现）

- `Injector`：`Inject(ctx, op Operator) error`、`Rollback(ctx, op Operator) error`——全部面向模拟器，禁止真实系统调用。
- C 族五类首版算子（必须全量）：网络延迟、网络分区、丢包、限带宽、CPU 压力、内存压力、IO 延迟、磁盘耗尽、时钟偏移、进程终止、依赖错误注入、依赖超时。
- 每个算子必须实现 `Inverse`（对偶回滚）；不可逆算子必须提供补偿策略并在 CES 校验时显式声明。

### 5. 依赖与约束

- 依赖：KRN-ALG、SIM-NET/SIM-FS/SIM-SCHED/SIM-CLK、SIM-RNG。
- 禁止：真实网络操作、真实进程操作、真实文件系统写（全部改道模拟器）；chaosblade 专有类型泄漏。

### 6. 测试要求

- 单元测试：五类算子各自注入→观测→回滚全链路用例；回滚后目标状态回到基线（对偶补偿律）。
- 确定性测试：同种子双跑注入事件流哈希一致。
- 上游回移演练：模拟一个上游安全补丁，走完评审→回移→TCK 回归全流程（M2 验收项）。

### 7. 验收标准（可量化）

1. C 族五类 12 个首版算子全量落地，全部实现 Inverse 或显式补偿。
2. 注入→回滚后基线恢复率 100%（对偶补偿测试）。
3. 同种子 100 次双跑事件流哈希一致。
4. chaosblade CLI 外壳移除确认（仓库内无 exec 拉起路径残留，扫描报告入库）。
5. 覆盖率 ≥95%/85%；门禁全绿。

---
## 附：L0 内核合入总验收（六模块齐备后）

1. 端到端：一份 CES YAML（稳态+M⊗C 混合扰动+探针+回滚）经 CTL-CLI 执行，产出确定性报告；同种子 100 次双跑逐字节一致。
2. 断网测试：拔网线环境（容器内禁网模拟）全量内核测试套件通过（零外部依赖铁律）。
3. 全模块覆盖率汇总报告入库；五段门禁+基准回归全绿。
