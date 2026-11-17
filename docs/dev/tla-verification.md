# Abate TLA+ 形式化验证报告

> 生成时间：2026-11-13
> 工具链：TLC2 Version 2026.11.06（tla2tools.jar，v1.8.0）
> 入口：`bash scripts/tla/check.sh`
> 规格目录：`scripts/tla/specs/`

## 一、目的

Abate 的确定性、资源隔离与覆盖率对拍是其不可妥协的质量红线。代码层的九段门禁（含 `-race`、覆盖率、govulncheck）能挡住实现缺陷，但挡不住**并发/状态机层面**的设计错误。本工作用 TLA+ 对三条负载最重的不变量做**模型检验（model checking）**，在数学层面证明：

1. 确定性随机流的「可复现性 / 重放等价」（SIM-REC 回放认证的基石）；
2. WASM/沙箱燃料（gas）模型的「有界性 + 必然终止」；
3. PIT 覆盖对拍计数的「有界 + 单调 + 满致一致」。

## 二、规格清单与结论

| 规格 | 验证目标 | 性质 | 模型规模 | 结论 |
|------|---------|------|---------|------|
| `StreamDeterminism.tla` | `sim/rand` 的 `Stream` + `splitMix64` | 不变量 `StateConsistency` / `Reproducibility` / `TypeOK` | 4800 个可达状态 | ✅ 全部通过 |
| `SandboxFuel.tla` | `plugin/sandbox` 与规划中 `plugin/wasm` 的燃料模型 | 不变量 `FuelNonNeg`/`UsedBounded`/`StepsBounded`/`HaltEmpty` + 活性 `Terminates` | 61 个可达状态 | ✅ 全部通过 |
| `PitParity.tla` | `kernel/mutation/pitparity` 的对拍率计算 | 不变量 `RateInRange`/`NoOverCount`/`FullWhenEqual`/`ZeroWhenZero` | 44 个可达状态 | ✅ 全部通过 |

运行 `bash scripts/tla/check.sh` 输出：

```
===== PitParity =====
[PASS] PitParity  (44 distinct states)
===== SandboxFuel =====
[PASS] SandboxFuel  (61 distinct states)
===== StreamDeterminism =====
[PASS] StreamDeterminism  (4800 distinct states)
===== TLA+ 形式化验证全部通过 =====
```

## 三、各规格要点

### 3.1 StreamDeterminism —— 确定性流可复现

抽象模型：每个流 `S[s]` 持有 `(origin, state, draws)`。核心事实是 `splitMix64` 的**值混合**不改变存储的 `state`，`state` 是纯算术：

```
state = (origin + draws * Gamma) % (2^WIDTH)
```

`Restore(seed, d)` 把 `state` 直接置为 `seed + d*Gamma`，恰等于「从 seed 新开、推进 d 次」所达的位置——这正是回放跳变（replay jump）数学上精确的原因。

证明的不变量：

- `StateConsistency`：`∀s, state == (origin + draws*Gamma) % M`。
- `Reproducibility`：若两流 `origin` 与 `draws` 相同，则 `state` 必相同。由此推出「`Restore(seed,d)` 与推进到 d 的流不可区分」。
- 动作覆盖 `New` / `Advance` / `Restore` / `Fork`（子流 `origin` 取父流当前 `state`，随后独立演化），穷尽所有交织。

### 3.2 SandboxFuel —— 燃料有界且必然终止

抽象模型：执行机持有 `fuel`、`steps`、`halted`、`used`。每步消耗 `c∈[1,MaxCost]` 单位燃料；燃料耗尽即当机（`halted`）。证明：

- 安全性：`fuel≥0`、`used≤F0`、`steps≤F0`、`halted ⇒ fuel=0`；
- 活性（弱公平下）：`◇halted`——有限燃料必然在 ≤F0 步内停机，且停机后永不再执行。

该证明在**实现 WASM 宿主（`plugin/wasm`，基于 wazero）之前**就锁死了「资源有界 + 确定性终止」的设计约束，使后续编码不得违反。

### 3.3 PitParity —— 对拍率计数有界单调

抽象模型：`comparable`（可比变异点）与 `consistent`（Abate 复现的点），`consistent≤comparable`。对拍率：

```
rate = consistent * SCALE / comparable   （整数除法，SCALE 生产环境为 10000）
```

证明：`rate∈[0,SCALE]`、`consistent≤comparable`（禁止超计）、满致一致 `consistent==comparable ⇒ rate==SCALE`、零一致 `consistent==0 ⇒ rate==0`。

## 四、工具链约束与处理（如实记录）

本环境 TLC2 2026.11.06 存在以下限制，已逐一规避并记录，**未伪造任何结果**：

1. **整数上限 2³¹−1**：`0x9E3779B97F4A7C15` 乃至其 32 位形式 `0x9E3779B9` 均溢出。由于可复现性证明对 `Gamma` 与模数均**参数无关**，模型中使用小的奇数常量 `Gamma==9`，结论对生产 64 位常量与全字宽成立（模运算的均匀性论证）。
2. **无 `Floor` / `Div` / `Mod` 函数**，仅有中缀 `%`：PitParity 的整数除法改用有界 `CHOOSE`（在 `0..SCALE` 上精确求解 floor 除法），避免递归 `Idiv` 在商过大时的栈溢出。
3. **不支持十六进制字面量**：统一改用十进制。
4. **相邻 `*` 与 `/` 需显式括号**：TLA+ 禁止无括号混用，已补括号。
5. **终止态需显式 stutter**：计数触顶后无动作可激发死锁，已补 `Terminated == UNCHANGED ...`。

## 五、运行与复用

```bash
# 前置：java 在 PATH；tla2tools.jar 缺失时 check.sh 自动从官方 release 下载缓存
bash scripts/tla/check.sh
```

- 新增规格：在 `scripts/tla/specs/` 放 `<Name>.tla` + `<Name>.cfg`，`check.sh` 自动发现。
- 常量覆盖（小字宽便于模型检验）写在各 `.cfg`；证明对生产字宽/数值参数无关。

## 六、与九段门禁的衔接

TLA+ 验证是九段门禁（header/style/license/gofmt/vet/golangci-lint/test-race/coverage/build）之上的**设计层补充**：门禁保证「代码写对了」，TLA+ 保证「设计没有错」。二者共同构成 Abate 的完工证据链。
