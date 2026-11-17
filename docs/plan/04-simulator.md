# Abate L3 确定性模拟器规划 — 六模块极致细节

- 文档编号：PLAN-04　状态：生效　上游：PLAN-00
- 覆盖模块：SIM-CLK / SIM-RNG / SIM-NET / SIM-FS / SIM-SCHED / SIM-REC
- 通用约束（六模块共同遵守）：
  1. 模拟器是纯软件数字孪生：禁止任何真实系统调用（网络/进程/文件系统写操作全部虚拟化）；断网环境全量测试通过（验收项）。
  2. 确定性铁律：同种子同配置，任意次运行事件流逐字节一致；所有随机性经 SIM-RNG，所有时间经 SIM-CLK。
  3. L3 只依赖 L0 内核类型与标准库；禁止依赖 L1/L2/L4-L7（L2 的 CES 类型不进入 L3——模拟器只认内核算子类型）。
  4. 覆盖率：语句 ≥95%、分支 ≥85%（内核层标准）；确定性回放测试为合入前置条件。

---

## 模块 SIM-CLK：虚拟时钟

### 1. 模块定位与职责

全引擎唯一时间源：整数纳秒虚拟时钟，支持快进/慢放/冻结/偏移四态操作。真实时间在模拟器内不存在；一切调度、超时、间隔、持续时间都以虚拟时间计。它是确定性（『稳』）与混沌时钟扰动（时钟偏移注入）的共同地基。

### 2. 包结构与文件布局

```
sim/clock/
  clock.go        // Clock 接口与 VirtualClock 实现
  instant.go     // Instant/Duration 类型（整数纳秒）
  state.go       // 四态：RUNNING/FAST/FREEZE/OFFSET
  clock_test.go
```

### 3. 核心类型与接口（必须实现）

- `Instant`：int64 纳秒（禁止浮点）；`Add(Duration) Instant`、`Sub(Instant) Duration`、`Before/After(Instant) bool`、`UnixNano() int64`。
- `Duration`：int64 纨秒；常量 `Nanosecond/Microsecond/Millisecond/Second/Minute/Hour`；`Abs()`。
- `Clock` 接口：
  - `Now() Instant`：当前虚拟时刻。
  - `Advance(d Duration) Instant`：推进时钟（仅调度器可调用；包级封装防误用——导出受控句柄 `SchedulerHandle`）。
  - `Freeze() / Unfreeze()`：冻结/解冻（冻结期间 Advance 无效，返回错误）。
  - `Offset(d Duration)`：时钟偏移注入（混沌时钟扰动的实现原语：Now() 返回值叠加偏移，底层事件时间不变）。
  - `Rate(multiplier int64, divisor int64) error`：快进/慢放（虚拟时间流速倍率；仅影响 Advance 的语义换算，事件排序不变）。
- `VirtualClock`：实现 Clock；内部状态（当前时刻/冻结标志/偏移量/倍率分子分母）全部整数运算，无锁（单线程虚拟时钟驱动铁律——并发由调度器边界保证）。

### 4. 关键算法与流程

1. 偏移语义：Offset 只影响 Now() 读数，不影响事件队列排序时间（被观测者看到的时间被扰动，调度秩序不被破坏——这是混沌时钟注入的正确语义）。
2. 倍率语义：Rate(2,1) 表示虚拟时间流速 ×2（快进）；Advance(d) 实际推进 d×multiplier/divisor；倍率变更记录为时钟事件（可回放）。
3. 溢出防护：int64 纳秒上限约 292 年；Advance 越界返回 ErrClockOverflow（拒绝而非回绕）。

### 5. 依赖与约束

- 依赖：仅标准库（math/bits 用于溢出检查）。
- 被依赖：KRN-SCH（时钟推进）、SIM-*（全部模块时间源）、KRN-INJ（时钟偏移注入）、WAS-HOST（host function abate:clock）。
- 禁止：time.Now/time.Sleep（vet+门禁扫描禁用清单）；浮点时间；真实定时器。

### 6. 测试要求

- 单元测试：四态操作全覆盖（运行/冻结/偏移/倍率）；偏移不影响事件排序的语义用例；溢出拒绝用例。
- 确定性测试：同操作序列 100 次执行，时刻序列哈希一致。
- 边界测试：零时长/负时长拒绝（Advance 参数校验）；最大值边界。

### 7. 验收标准（可量化）

1. 四态操作与组合（冻结+偏移、倍率+偏移）用例 100% 通过。
2. 偏移不影响事件排序的语义测试通过（混沌时钟注入正确性）。
3. 溢出防护用例通过（拒绝而非回绕）。
4. 确定性 100 次哈希一致；覆盖率 ≥95%/85%；门禁全绿。

---

## 模块 SIM-RNG：确定性 PRNG

### 1. 模块定位与职责

全引擎唯一随机源：种子驱动的确定性伪随机数流。同种子必然产生同序列；随机流消费全部记账（哪个算子/探针在哪个虚拟时刻取了多少随机数），保证回放可精确复现。它是确定性（『稳』）的另一半地基。

### 2. 包结构与文件布局

```
sim/rand/
  rand.go        // Stream：确定性随机流
  splitmix.go    // SplitMix64 核心算法（公有领域，源码级并入）
  accounting.go  // 消费记账：随机数提取全记录
  rand_test.go
```

### 3. 核心类型与接口（必须实现）

- `Stream`：
  - `New(seed uint64) *Stream`：从种子构造（SplitMix64 初始化）。
  - `Uint64() uint64`、`Int63() int64`、`Float64() float64`（[0,1)）、`IntN(n int) int`（[0,n) 拒绝采样无偏）、`Bool() bool`、`Bytes(n int) []byte`。
  - `Fork() *Stream`：派生子流（子流种子由父流下一随机数派生；用于并行实验的独立随机流——子流互不干扰）。
  - `State() (seed uint64, draws uint64)`：当前状态（种子+已提取次数）。
  - `Restore(seed, draws uint64) error`：状态恢复（回放核心：跳过 draws 次提取，O(1) 或 O(log n) 快进）。
- `Accounting`：每次提取记录（虚拟时刻/调用方 ID/方法/次数）；`Snapshot() []DrawRecord` 供 SIM-REC 录制。

### 4. 关键算法与流程

1. 核心算法 SplitMix64（公有领域，无许可证负担）：64 位状态、每步输出经黄金比例混合；周期 2^64；通过 TestU01 BigCrush（文献结论，验收时以 PractRand 1TB 快速档复测）。
2. 无偏 IntN：拒绝采样（拒绝区域重取，不取模偏差）。
3. 状态恢复：draws 计数 + 重放提取（O(draws)）或预计算跳跃表（O(log n)，v1 采用重放法，性能预算内）。
4. Fork 语义：子流种子 = 父流 Uint64() 一次；Fork 后父流与子流独立演进；Fork 记账（保证回放时 Fork 序列一致）。

### 5. 依赖与约束

- 依赖：仅标准库。
- 被依赖：KRN-SCH（调度随机）、KRN-MUT（变异点选择）、KRN-INJ（故障参数随机）、SIM-NET（丢包/延迟抖动）、WAS-HOST（host function abate:rand）。
- 禁止：math/rand 全局源（全局状态不可回放）；crypto/rand（非确定性；安全场景由 XCU-SEC 单独处理，不进模拟器）；time 作种子。

### 6. 测试要求

- 单元测试：各方法边界（IntN(0)/IntN(1) 拒绝或退化、Float64 ∈ [0,1)）；无偏性统计测试（卡方检验，10^6 样本）。
- 确定性测试：同种子 100 次提取序列哈希一致；Fork 后父子流独立（同种子同 Fork 点，子流序列一致）。
- 回放测试：State/Restore 往返（提取 10^6 次后恢复，继续提取序列与连续提取一致）。
- 统计质量：PractRand 快速档（128MB）无异常告警（工具输出入库）。

### 7. 验收标准（可量化）

1. 同种子序列确定性 100 次哈希一致；Fork 独立性用例通过。
2. State/Restore 往返一致性用例通过（10^6 提取规模）。
3. 卡方无偏性测试通过（显著性水平 0.01）；PractRand 128MB 无告警。
4. 提取性能：≥10^8 次/秒（基准报告入库）；覆盖率 ≥95%/85%；门禁全绿。

---

## 模块 SIM-NET：虚拟网络

### 1. 模块定位与职责

全软件虚拟网络栈：节点/链路/路由的内存模型，支持四类故障注入（延迟/分区/丢包/限带宽）与确定性传播仿真。网络包在虚拟链路上的传播、排队、丢弃全部事件化（进调度器队列），回放可逐包复现。它是 C 族网络算子的执行地基。

### 2. 包结构与文件布局

```
sim/net/
  net.go         // Network：拓扑管理（节点/链路/路由）
  node.go        // Node：虚拟端点（含协议栈桩：TCP/UDP 语义子集）
  link.go        // Link：虚拟链路（带宽/延迟/丢包率/分区状态）
  packet.go     // Packet：虚拟包（载荷/序号/时间戳）
  transport.go  // Transport 接口：模拟器内 http/shell 探针的虚拟通道
  fault.go      // 故障注入点：四类故障的链路级实现
  net_test.go
```

### 3. 核心类型与接口（必须实现）

- `Network`：`AddNode(id string) *Node`、`Link(a, b string, cfg LinkConfig) *Link`、`Route(from, to string) []string`（最短路径，确定性平局裁决：按节点 ID 字典序）、`Send(from, to string, payload []byte) (PacketID, error)`。
- `LinkConfig`：`BaseLatency Duration`（基础延迟）、`Jitter Duration`（抖动上限，经 SIM-RNG）、`BandwidthBps uint64`（带宽，字节/秒）、`LossRate float64`（丢包率 [0,1]，经 SIM-RNG）。
- `Link`：`Partition() error / Heal() error`（分区/恢复——分区期间 Send 返回 ErrPartitioned）、`SetLatency(d Duration)`、`SetLossRate(r float64)`、`SetBandwidth(bps uint64)`——四类故障注入的链路级原语（KRN-INJ network.go 调用）。
- `Packet`：`ID PacketID`（全局序号）、`From/To string`、`Payload []byte`（上限 1MB）、`SentAt/DeliverAt Instant`（虚拟时间）。
- `Transport` 接口：`Dial(ctx, addr string) (Conn, error)`——模拟器内 HTTP 探针与 shell 探针的虚拟连接（Conn 语义子集：Read/Write/Close，全部虚拟化）。

### 4. 关键算法与流程

1. 传播仿真：Send → 路由计算 → 逐跳入队（每跳：丢包判定→带宽排队→延迟传播，全部生成调度器事件）→ 送达事件触发接收方回调。
2. 丢包判定：每包独立伯努利试验（SIM-RNG Bool 加权）；丢弃生成 drop 事件（可回放）。
3. 带宽排队：令牌桶模型（字节令牌，虚拟时间补充）；排队包按 FIFO，同刻按包序号（确定性平局裁决）。
4. 分区语义：Partition 后链路双向不可达（Send 立即返回 ErrPartitioned，不产生传播事件）；Heal 恢复。分区状态变更生成链路事件。

### 5. 依赖与约束

- 依赖：SIM-CLK（虚拟时间）、SIM-RNG（丢包/抖动）、KRN-SCH（事件入队）。
- 被依赖：KRN-INJ（网络故障算子）、KRN-OBS（http 探针经 Transport）。
- 禁止：真实 socket（net.Dial 禁用清单）；真实 DNS；系统调用。

### 6. 测试要求

- 单元测试：四类故障各自注入→观测→恢复全链路用例；路由确定性（同拓扑 100 次路由结果一致）。
- 确定性测试：同种子同流量 100 次双跑，包事件流（发送/送达/丢弃全序）哈希一致。
- 压力测试：1000 节点/10000 链路拓扑，10 万包传播，事件吞吐达标（性能预算项）。
- 语义测试：分区期间 Send 报错、Heal 后恢复；带宽限制下大包传输时间符合令牌桶计算值。

### 7. 验收标准（可量化）

1. 四类故障注入/恢复用例 100% 通过；对偶回滚（Heal/SetLatency 复原）后基线恢复率 100%。
2. 包事件流确定性 100 次哈希一致。
3. 压力测试：10 万包传播 P99 ≤1s（基准报告入库）。
4. 覆盖率 ≥95%/85%；门禁全绿。

---

## 模块 SIM-FS：虚拟文件系统

### 1. 模块定位与职责

内存虚拟文件系统：目录树/文件/权限的内存模型，支持 IO 延迟与磁盘耗尽故障注入。模拟器内一切文件操作（探针 shell 的文件访问、被测目标的文件读写）都落在虚拟 FS。它是 C 族资源算子（IO/磁盘）的执行地基。

### 2. 包结构与文件布局

```
sim/fs/
  fs.go          // FS：挂载点与目录树管理
  file.go        // File：虚拟文件（内容/权限/时间戳）
  dir.go         // Dir：虚拟目录
  handle.go      // Handle：打开文件句柄（读写游标）
  fault.go       // 故障注入点：IO 延迟/磁盘耗尽/坏块
  fs_test.go
```

### 3. 核心类型与接口（必须实现）

- `FS`：`MkdirAll(path string) error`、`Create(path string) (*Handle, error)`、`Open(path string) (*Handle, error)`、`Remove(path string) error`、`Stat(path string) (FileInfo, error)`、`Walk(root string, fn func(path string) error) error`（字典序遍历，确定性）。
- `Handle`：`Read(p []byte) (int, error)`、`Write(p []byte) (int, error)`、`Seek(off int64) (int64, error)`、`Close() error`——语义对齐 os.File 子集（错误码语义对齐：ErrNotExist/ErrPermission/ErrNoSpace）。
- `FileInfo`：`Size int64`、`Mode FileMode`、`ModTime Instant`（虚拟时间）。
- 故障注入原语（KRN-INJ resource.go 调用）：`SetIODelay(d Duration)`（每次读写附加虚拟延迟）、`SetQuota(bytes int64)`（磁盘配额；写超配额返回 ErrNoSpace——磁盘耗尽故障）、`SetBadBlocks(rate float64)`（坏块率：读返回 IO 错误，经 SIM-RNG）。

### 4. 关键算法与流程

1. 路径规范化：`/` 分隔、`.` 与 `..` 解析、拒绝穿越根（对齐 path.Clean 语义）；路径校验在 FS 边界（防探针 shell 逃逸）。
2. IO 延迟：每次 Read/Write 生成延迟事件（经调度器），操作在虚拟时间推进后完成——被观测者感知真实 IO 节奏。
3. 配额判定：写前检查（当前用量+写入量 > 配额即 ErrNoSpace）；配额变更生成 FS 事件。
4. 内容存储：[]byte 全内存（v1 上限单文件 100MB/总量 1GB，超限拒绝——防内存炸弹；大文件场景属生态层扩展）。

### 5. 依赖与约束

- 依赖：SIM-CLK（时间戳/延迟）、SIM-RNG（坏块）、KRN-SCH（延迟事件）。
- 被依赖：KRN-INJ（IO/磁盘算子）、KRN-OBS（shell 探针文件访问）、LNG 前端（源码目标读取，经端口注入）。
- 禁止：真实文件系统读写（os.Create/os.Open 禁用清单，测试夹具除外）；真实权限系统调用。

### 6. 测试要求

- 单元测试：CRUD 全路径（含错误码语义）；路径规范化（穿越/空路径/根逃逸拒绝）；IO 延迟/配额/坏块三类故障注入→恢复用例。
- 确定性测试：同操作序列 100 次双跑，FS 事件流与最终目录树快照哈希一致。
- 边界测试：配额恰好写满（边界值）；单文件/总量上限拒绝。

### 7. 验收标准（可量化）

1. CRUD 与错误码语义用例 100% 通过；路径逃逸拒绝用例通过。
2. 三类故障注入/恢复用例通过；回滚后基线恢复率 100%。
3. FS 事件流与目录树快照确定性 100 次哈希一致。
4. 覆盖率 ≥95%/85%；门禁全绿。

---

## 模块 SIM-SCHED：虚拟调度

### 1. 模块定位与职责

虚拟进程/协程模型：虚拟任务（VTASK）的创建/终止/挂起/恢复与 CPU 配额仿真。被测目标在模拟器内以虚拟任务形态运行；进程终止/挂起故障（C 族 process 算子）作用于虚拟任务。它是『进程/协程级故障注入』的执行地基。

### 2. 包结构与文件布局

```
sim/vsched/
  vsched.go     // VScheduler：虚拟任务表与 CPU 配额
  vtask.go      // VTask：虚拟任务（状态机）
  cgroup.go     // 虚拟 cgroup：CPU/内存配额与压力
  fault.go      // 故障注入点：终止/挂起/CPU 压力/内存压力
  vsched_test.go
```

### 3. 核心类型与接口（必须实现）

- `VTask`：`ID VTaskID`、`Name string`、`State VTaskState`（READY/RUNNING/BLOCKED/SUSPENDED/TERMINATED）、`Entry func(ctx VContext)`（任务入口——在虚拟时间片内执行，遇阻塞原语让出）。
- `VScheduler`：`Spawn(name string, entry func(VContext)) VTaskID`、`Kill(id VTaskID) error`（终止——进程终止故障原语）、`Suspend(id VTaskID) error / Resume(id VTaskID) error`（挂起/恢复）、`Tasks() []VTaskInfo`（字典序）。
- `VContext`：任务内上下文：`Sleep(d Duration)`（虚拟睡眠——让出 CPU，生成唤醒事件）、`Yield()`、`Done() <-chan struct{}`（取消信号，虚拟化 channel）。
- `VCgroup`：`SetCPUShare(permille int)`（CPU 配额，千分比）、`SetMemoryLimit(bytes int64)`（内存配额；超限触发 OOM-Kill 事件——内存压力故障原语）、`PressureCPU(permille int)`（CPU 压力注入：配额压缩模拟）。

### 4. 关键算法与流程

1. 调度模型：协作式调度（任务在阻塞原语/显式 Yield 让出）；就绪队列按（优先级, 创建序号）排序——确定性平局裁决；每个虚拟时间片执行就绪任务至让出。
2. CPU 配额：每虚拟时间片按 cgroup share 加权分配执行预算（整数切片，无浮点）；超预算任务让出，下片继续。
3. OOM 语义：内存用量超限 → 生成 OOM 事件 → 任务状态 TERMINATED（原因 OOM）——与 Kill 区分（可观测区分故障类型）。
4. Kill/Suspend 语义：Kill 立即终止（在途虚拟操作作废）；Suspend 冻结（不参与调度，状态保持）；Resume 回到 READY。

### 5. 依赖与约束

- 依赖：SIM-CLK、SIM-RNG（压力抖动）、KRN-SCH（事件入队）。
- 被依赖：KRN-INJ（process/resource 算子）、SIM-NET（协议栈桩任务化）。
- 禁止：真实 goroutine 直接作为被测载体（goroutine 只作宿主机制，虚拟任务语义独立）；真实进程操作（os/exec 禁用清单）；真实信号（SIGKILL 等）。

### 6. 测试要求

- 单元测试：任务状态机全转移（READY→RUNNING→BLOCKED→…→TERMINATED 合法路径 + 非法转移拒绝）；Kill/Suspend/Resume 语义用例；CPU 配额加权分配正确性（整数切片校验）。
- 确定性测试：同种子同任务集 100 次双跑，调度事件流哈希一致。
- 压力测试：10000 虚拟任务并发调度，事件吞吐达标（性能预算项）。
- OOM 测试：内存超限触发 OOM 事件且任务终止、原因正确。

### 7. 验收标准（可量化）

1. 状态机全转移用例 100% 通过（含非法转移拒绝）。
2. 调度事件流确定性 100 次哈希一致。
3. 压力测试：10000 任务 P99 调度延迟（虚拟时间）达标；事件吞吐 ≥10 万/秒（基准报告入库）。
4. 覆盖率 ≥95%/85%；门禁全绿。

---

## 模块 SIM-REC：录制回放

### 1. 模块定位与职责

观测与事件的全程录制（WAL 式追加日志）与逐帧回放：把一次实验的全部事件流（调度事件/观测事件/随机提取记账/时钟操作/故障注入）按全序录制为可回放档案；回放时逐帧重演并逐字节校验。它是确定性验收（『同种子双跑哈希一致』）的裁判与证据链。

### 2. 包结构与文件布局

```
sim/record/
  record.go     // Recorder：录制器（多源订阅，追加写入）
  frame.go      // Frame：帧模型（序号/虚拟时间/类型/载荷）
  archive.go   // Archive：档案格式（版本化二进制，含索引与校验和）
  replay.go     // Replayer：回放器（逐帧重演+校验）
  verify.go    // Verify：双跑比对（哈希链校验）
  record_test.go
```

### 3. 核心类型与接口（必须实现）

- `Frame`：`Seq uint64`（全序，与调度器 Seq 同源）、`At Instant`、`Kind FrameKind`（SCHED/OBS/RNG/CLK/FAULT/NET/FS/VTASK）、`Payload []byte`（类型化载荷，版本化编码）。
- `Recorder`：`Subscribe(sources ...Source)`（订阅调度器/观测管道/RNG 记账/时钟/注入器）、`Start(ctx)` / `Stop() (*Archive, error)`。
- `Archive`：版本化二进制格式：头部（magic+版本+实验指纹）+ 帧块（追加写）+ 索引（按 Seq）+ 尾部 SHA-256 全档校验和；`Save(w io.Writer) error`、`Load(r io.Reader) (*Archive, error)`（加载即校验完整性）。
- `Replayer`：`Replay(arch *Archive, hooks ReplayHooks) error`——逐帧回调（供重演引擎消费）；`Verify(arch *Archive) error`——档案自校验（校验和+序号连续性+时间单调性）。
- `Compare(a, b *Archive) (Diff, error)`：双跑比对：逐帧哈希比对，输出首个差异点（Seq/Kind/期望/实际）。

### 4. 关键算法与流程

1. 录制链：各源事件 → 帧编码（类型化+版本化）→ 追加写入（内存缓冲，Stop 时落 Archive）→ 尾部校验和。
2. 回放链：Load（校验）→ 逐帧解码 → hooks 回调（重演引擎按帧重放：恢复 RNG 状态/时钟状态/重放事件）→ 逐帧校验（重放产生的帧与档案帧哈希一致）。
3. 哈希链：每帧含前帧哈希（链式）——任何一帧篡改导致后续全部校验失败（防篡改证据链）。
4. 双跑比对：Compare 输出结构化 Diff（首个差异+统计摘要），供 CI 确定性验收与故障诊断。

### 5. 依赖与约束

- 依赖：KRN-SCH（Seq 同源）、KRN-OBS（观测订阅）、SIM-CLK/SIM-RNG（状态记账）、标准库（hash/sha256、encoding/binary）。
- 被依赖：XCU-TST（确定性验收裁判）、CTL-CLI（replay 命令）、RPT-*（回放报告）。
- 禁止：录制过程影响被录事件流（零侵入：只读订阅）；真实时间戳入档（档案内只有虚拟时间+序号）。

### 6. 测试要求

- 单元测试：帧编码往返（各 Kind 载荷无损）；档案 Save/Load 往返；校验和篡改检测（翻转任一字节，Load 拒绝）。
- 回放测试：录制一个完整实验（含 M⊗C 混合扰动）→ 回放 → 逐帧哈希一致；RNG/时钟状态恢复正确性（回放后继续运行与原运行一致）。
- 哈希链测试：篡改中间帧，后续校验全部失败。
- 规模测试：100 万帧档案（约 1GB）Load+Verify P99 ≤10s（性能预算项）。

### 7. 验收标准（可量化）

1. 帧编码往返与 Save/Load 往返 100% 无损；篡改检测用例 100% 通过。
2. 完整实验回放逐帧哈希一致（含混合扰动场景）。
3. 哈希链防篡改：单帧篡改导致后续全失败用例通过。
4. 规模性能：100 万帧 Load+Verify P99 ≤10s（基准报告入库）。
5. 覆盖率 ≥95%/85%；门禁全绿。

---

## 附：L3 合入总验收（六模块齐备后）

1. 端到端确定性：一个含网络故障（延迟+丢包）+ 资源压力（CPU 配额）+ 时钟偏移的混合实验，同种子 100 次双跑，SIM-REC 档案逐字节一致（Compare 零差异）。
2. 断网验证：容器内禁网环境，L3 全量测试套件通过（零真实网络依赖铁律）。
3. 禁用清单扫描：time.Now/math/rand 全局源/net.Dial/os.Create/os.Exec 在 sim/ 树内零出现（扫描脚本入库，CI 常驻）。
4. 六模块覆盖率汇总报告入库；五段门禁+基准回归全绿。
