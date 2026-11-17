# Abate L1 WASM 运行时规划 — 四模块极致细节

- 文档编号：PLAN-02　状态：生效　上游：PLAN-00
- 覆盖模块：WAS-HOST / WAS-CAP / WAS-HOT / WAS-SIG
- 通用约束（四模块共同遵守）：
  1. L1 是 L4-L7 插件进入内核的唯一通道（端口-适配器铁律）；L1 只依赖 L0 内核接口与标准库，禁止反向依赖 L4-L7。
  2. WASM 宿主运行时采用源码级接入（五步并入法），候选：wazero（Go 实现、零依赖、Apache-2.0、无 CGO、支持 WASI Preview2）——vendor 后二次更改（确定性化：时钟与随机源改道 SIM-CLK/SIM-RNG）、优化（裁剪不需要的特性）、深度集成（WIT 绑定生成器对齐 Abate 插件接口）、封装（对外仅暴露 Host 类型）。
  3. 插件永不信任：默认零能力，一切能力显式授予；未签名插件拒绝加载。
  4. 覆盖率：语句 ≥85%、分支 ≥75%（生态层标准）；安全测试用例为合入前置条件。

---

## 模块 WAS-HOST：WASM 宿主

### 1. 模块定位与职责

提供 WASM 模块的加载、实例化、调用与生命周期管理；定义五类插件（operator/injector/frontend/adapter/reporter）的统一 WIT（WebAssembly Interface Types）接口；把宿主能力（日志、时钟、随机、观测上报）以 host function 形式注入插件。它是『劲』字（生态小时级扩展）的运行时基座。

### 2. 包结构与文件布局

```
wasm/host/
  host.go          // Host：加载/实例化/调用入口
  wit.go           // 五类插件的 WIT 接口定义（版本化）
  instance.go      // Instance：单插件实例封装（内存表/调用句柄）
  hostfunc.go      // host function 实现：日志/时钟/随机/观测上报
  registry.go      // 插件注册表：类型→工厂
  host_test.go
```

### 3. 核心类型与接口（必须实现）

- `Host`：
  - `Load(ctx, path string, opts LoadOptions) (*Instance, error)`：加载 .wasm 文件（先经 WAS-SIG 验签，再实例化）。
  - `Registry() *Registry`：插件注册表访问。
- `LoadOptions`：`Capabilities Capabilities`（能力授予，默认全拒）、`Fuel uint64`（执行燃料上限，防死循环插件）、`MemoryMB uint32`（内存上限，默认 64MB）。
- `Instance`：
  - `Call(ctx, entry string, args []byte) ([]byte, error)`：调用插件导出函数（入口白名单校验）。
  - `Type() PluginType`：五类之一。
  - `Version() SemVer`：插件自报版本（与签名元数据交叉校验）。
  - `Close() error`：排空并释放（配合 WAS-HOT）。
- WIT 接口（五类插件各一组，版本化 `abate:plugin/operator@1.0.0` 等）：
  - operator：`describe() -> OperatorSpec`（算子声明：ID/族/参数 schema/可逆性）。
  - injector：`inject(spec) -> Result` / `rollback(spec) -> Result`。
  - frontend：`parse(source) -> AirBytes`（产出 AIR 序列化字节）。
  - adapter：`notify(event) -> Result`。
  - reporter：`report(report) -> Result`。
- host function（宿主注入，全部经能力门控）：`abate:log`（日志）、`abate:clock`（虚拟时钟读取）、`abate:rand`（确定性随机流）、`abate:observe`（观测上报）。

### 4. 关键算法与流程

1. 加载链：读文件 → WAS-SIG 验签（失败即拒）→ 能力声明解析 → WAS-CAP 授予 → 实例化（燃料+内存上限注入）→ 注册表登记。
2. 调用链：入口白名单校验 → 燃料检查 → 调用 → 结果校验（长度上限 1MB，防内存炸弹）→ 返回。
3. 燃料耗尽：调用立即终止，返回 ErrFuelExhausted，实例标记 degraded（不销毁，可复用）。

### 5. 依赖与约束

- 依赖：wazero 源码级并入（third_party/wazero，Apache-2.0）、KRN-ALG（算子类型）、SIM-CLK/SIM-RNG（host function 后端）。
- 被依赖：WAS-CAP/WAS-HOT/WAS-SIG、L4-L7 插件加载方。
- 禁止：CGO；真实网络监听；插件直接访问文件系统（必须经能力授予的 host function）。

### 6. 测试要求

- 单元测试：五类插件各一个最小 .wasm 测试夹具（Rust/TinyGo 编译产物入库），加载→调用→关闭全链路。
- 边界测试：燃料耗尽、内存超限、结果超长、入口白名单拒绝。
- 确定性测试：同种子下 host function（rand/clock）输出序列一致。

### 7. 验收标准（可量化）

1. 五类插件 WIT 接口全部定义并版本化；测试夹具全链路通过。
2. 燃料/内存/结果长度三上限各有触发用例且错误码正确。
3. 插件调用 P99 开销 ≤50μs（不含首次实例化；基准报告入库）。
4. 覆盖率 ≥85%/75%；门禁全绿。

---

## 模块 WAS-CAP：能力沙箱

### 1. 模块定位与职责

能力制安全模型：fs/net/clock/env/observe 五类能力，默认全拒、显式授予、按插件声明收敛。它是『强』字（安全纵深）的第一道墙。

### 2. 包结构与文件布局

```
wasm/capability/
  capability.go    // Capabilities 位集与五类能力定义
  grant.go        // 授予解析：声明→能力集（最小权限收敛）
  check.go        // 运行时检查：host function 调用前门控
  audit.go        // 能力审计日志（授予/拒绝/使用全记录）
  capability_test.go
```

### 3. 核心类型与接口（必须实现）

- `Capability` 枚举：`CapFS`（受限文件访问，仅插件数据目录）、`CapNet`（受限网络，仅声明的主机白名单）、`CapClock`（虚拟时钟读取）、`CapRand`（确定性随机）、`CapObserve`（观测上报）。
- `Capabilities`：位集类型；`Has(Capability) bool`、`Grant(Capability)`、`Revoke(Capability)`。
- `ParseGrant(decl string) (Capabilities, error)`：解析插件能力声明（manifest 格式：`capabilities: [clock, rand]`）；未知能力名拒绝加载（前向兼容白名单制）。
- `Check(cap Capabilities, c Capability) error`：host function 调用前门控；未授权返回 ErrCapabilityDenied 并记审计。
- `AuditLog`：授予/拒绝/使用三类事件全记录（插件 ID+能力+时间+结果），供安全审计与 XCU-SEC 消费。

### 4. 关键算法与流程

1. 最小权限收敛：插件声明的能力与宿主策略（策略文件可配置）取交集，交集为空即零能力。
2. 网络白名单：CapNet 必须伴随主机白名单（manifest `net-hosts`）；DNS 解析在宿主侧完成（防 DNS 重绑定）。
3. 文件沙箱：CapFS 仅映射插件专属数据目录（`/var/lib/abate/plugins/<id>/`），路径穿越（`..`）在宿主侧规范化后校验。

### 5. 依赖与约束

- 依赖：无外部依赖（纯标准库）。
- 被依赖：WAS-HOST（加载链与 host function 门控）。
- 禁止：任何『默认开启』能力；能力提升（提权）请求——插件运行期申请能力一律拒绝（能力只在加载时授予一次）。

### 6. 测试要求

- 单元测试：五类能力授予/拒绝矩阵（5×2 全组合）；声明解析（合法/非法/未知能力名）。
- 安全测试（合入前置）：越权调用全部被拦截（每类能力各一个越权用例）；路径穿越用例；DNS 白名单绕过用例。
- 审计测试：授予/拒绝/使用三类事件全部落审计日志。

### 7. 验收标准（可量化）

1. 5×2 能力矩阵用例 100% 通过；越权拦截用例 100% 通过。
 2. 能力声明解析：合法/非法/未知名三类各有用例，未知名 100% 拒绝加载。
3. 审计日志三类事件覆盖率 100%（用例断言日志条目）。
4. 覆盖率 ≥85%/75%；门禁全绿。

---

## 模块 WAS-HOT：热加载

### 1. 模块定位与职责

插件运行时升级：原子替换（新版本实例化→流量切换→旧实例排空→卸载），运行中实验不中断。它是『劲』字（小时级生态迭代无感升级）的直接落点。

### 2. 包结构与文件布局

```
wasm/hotswap/
  swap.go        // Swap：原子替换主流程（状态机）
  drain.go       // 排空：旧实例在途调用等待完成（带超时）
  compat.go     // 兼容性矩阵：semver 区间判定
  hotswap_test.go
```

### 3. 核心类型与接口（必须实现）

- `Swap(ctx, old, new *Instance) error`：原子替换状态机：`PREPARE`（新实例加载验签）→ `SWITCH`（注册表原子替换，单次指针写）→ `DRAIN`（旧实例在途调用等待）→ `UNLOAD`（旧实例卸载）。
- `DrainTimeout`：默认 30s（真实时间）；超时后旧实例强制 Close（在途调用返回 ErrDrainTimeout）。
- `CompatMatrix`：`Allowed(pluginType, from, to SemVer) bool`——按插件类型维护 semver 兼容区间（major 变更默认不兼容，需宿主策略显式放行）。

### 4. 关键算法与流程

1. 原子性：注册表替换为单次原子写（sync/atomic 或 RWMutex 写锁内单赋值）；读侧永不阻塞。
2. 排空：旧实例引用计数归零（在途调用持有引用）才卸载；引用计数泄漏检测（排空超时即报告泄漏插件 ID）。
3. 失败回滚：PREPARE 阶段失败（验签/实例化/兼容性不通过）不影响旧实例继续服务。

### 5. 依赖与约束

- 依赖：WAS-HOST（实例管理）、WAS-SIG（新版本验签）。
- 被依赖：CTL-CLI/CTL-REST（升级命令入口）。
- 禁止：替换期间任何注册表中间态可见（读侧要么全旧要么全新）；禁止卸载仍在途调用的实例。

### 6. 测试要求

- 并发测试：100 并发读 + 1 次替换，读侧零错误、零中间态（断言每次调用命中的实例版本单一）。
- 排空测试：在途慢调用（sleep 插件）+ 替换 → 调用完成 → 旧实例卸载；超时路径（强制关闭）。
- 失败回滚测试：新版本损坏（坏 wasm 字节）→ 旧实例不受影响继续服务。

### 7. 验收标准（可量化）

1. 并发替换测试 100 并发读零错误零中间态（-race 门禁下通过）。
2. 排空与超时路径各有用例；引用计数泄漏检测有用例。
3. 失败回滚：坏字节/验签失败/兼容性不通过三类，旧实例可用性 100% 保持。
4. 覆盖率 ≥85%/75%；门禁全绿。

---

## 模块 WAS-SIG：签名验证

### 1. 模块定位与职责

插件签名链验证：GPG/cosign 双通道验签、签名元数据（插件 ID/版本/能力声明哈希）交叉校验。它是『强』字（安全纵深）的第二道墙，也是供应链安全的入口。

### 2. 包结构与文件布局

```
wasm/signature/
  verify.go       // Verify：验签主入口（GPG 优先，cosign 备选）
  gpg.go          // GPG 通道：openpgp 源码级并入（golang.org/x/crypto/openpgp，BSD-3）
  cosign.go       // cosign 通道：sigstore 签名 blob 验证（源码级并入）
  manifest.go     // 插件 manifest 解析与哈希交叉校验
  truststore.go    // 信任锚：内置官方密钥 + 用户自增密钥
  signature_test.go
```

### 3. 核心类型与接口（必须实现）

- `Verify(path string) (*Manifest, error)`：验签主入口——签名文件（`<name>.wasm.sig`）+ 信任锚校验 + manifest 哈希交叉校验；任一失败返回 ErrSignatureInvalid 并拒绝加载。
- `Manifest`：`PluginID`、`Version SemVer`、`Type PluginType`、`CapabilitiesDecl string`、`SHA256 [32]byte`（wasm 字节哈希）、`SignedAt time.Time`。
- `TrustStore`：`Add(key)`（用户自增）、`Builtin()`（官方发布密钥，随版本更新）；密钥指纹去重。
- 交叉校验：manifest 声明的 SHA256 与实际 wasm 字节哈希一致；签名覆盖 manifest 全字段（防篡改）。

### 4. 关键算法与流程

1. 验签链：读 .wasm 与 .sig → 信任锚找公钥 → GPG/cosign 验签 → manifest 哈希交叉校验 → 通过才返回 Manifest 给 WAS-HOST。
2. 双通道策略：签名文件带通道标记；GPG 签名走 openpgp，cosign 签名走 sigstore blob 验证；两通道互不降级（不存在『验签失败退回无签名』路径）。
3. 密钥轮换：信任锚支持多密钥并存（新旧重叠期）；官方密钥随 Abate 版本更新并公示指纹。

### 5. 依赖与约束

- 依赖：openpgp 源码级并入（third_party/openpgp）、cosign 验证库源码级并入（third_party/sigstore，Apache-2.0）。
- 被依赖：WAS-HOST（加载链）、WAS-HOT（升级链）。
- 禁止：任何『跳过验签』的开关（测试用注入接口除外，生产构建禁用）；密钥私钥永不入库（仅公钥信任锚）。

### 6. 测试要求

- 单元测试：GPG/cosign 各一组（有效签名通过、篡改字节拒绝、错误密钥拒绝、manifest 哈希不符拒绝）。
- 信任锚测试：多密钥并存、用户自增、指纹去重。
- 安全测试（合入前置）：无签名插件加载 100% 被拒；篡改 manifest 100% 被拒。

### 7. 验收标准（可量化）

1. 双通道验签正反例全通过（每通道 ≥4 用例）。
2. 无签名/篡改/错钥/哈希不符四类拒绝用例 100% 通过。
3. 验签开销 P99 ≤10ms（基准报告入库）。
4. 覆盖率 ≥85%/75%；门禁全绿。

---

## 附：L1 合入总验收（四模块齐备后）

1. 端到端：一个签名插件从 Load（验签→能力授予→实例化）→ Call → Swap（热升级）→ Close 全链路演示，运行中实验不中断。
2. 安全审计：越权/无签名/篡改三类攻击用例全部拦截（安全测试报告入库）。
3. 插件全流程（提交→签名→发布→热加载）≤1 小时实测记录（M4 验收项）。
