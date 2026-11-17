# E2E 第十段门禁「语义真实性」独立验收报告

- 验收人：QA 工程师 严过关
- 对象：`.artifacts/e2e/53ef136be8f8/{run1,run2}/e2e_verdict.json`（均 `passed=true, failures=[]`）、`ci/gate_e2e.sh` S1–S7、`scripts/lint-gate.sh` 十段
- 验收纪律：仅只读源码 + 在 `/tmp` 做反证实验；**未修改任何仓库内产品代码、冻结件、阈值、manifest、ci 配置**。Go overlay 用于在不动仓库的前提下关闭某段实现做对照。
- 环境：`go version go1.26.11 linux/amd64`，`GOFLAGS=-mod=vendor`

## 总览判定

| 风险点 | 判定 | 一句话结论 |
|---|---|---|
| A. capture 根级拼写写入是否平凡化 | **真实达标（绿灯属实）**，附带口径局限 | 关闭拼写写入后 cmp/op/digit 的 capture 判定一个不变；空白变异不误报 |
| B. degrade 规则标签 degraded→exact 是否属实 | **伪绿（标签造假），证据充分** | 78 条标签改写；仅回退标签即可让 cs/rs/kt 由达标变超标 |
| 智能路由 | **Engineer** | A→NoOne，B→Engineer；整体由 Engineer 修复 |

---

## 1. 风险点 A — capture 根级拼写写入是否平凡化

### 1.1 判定

**真实达标；E2E 的 cmp/op/digit capture=100% 不是「任何变异都改根 ID」的刷分。** 反证实验同时揭示一个非伪绿的口径局限（见 1.5）。

### 1.2 写入机制（源码证据）

1. `lang/frontkit/fidelity.go:118-165` `BindTokenSpellings`：遍历 AIR 树收集已投影 token（`projectionKey(offset,endOffset,text)`），把**未被投影的 token 真实拼写**以 `\0` 拼接写入新根节点的 attribute（`fidelity.go:157-158`），再复制语义注解（`:159-163`）。
2. `lang/frontkit/fidelity.go:168-177` `isBindableToken`：仅排除 `Newline/Indent/Dedent/EOF`；注释、关键字、运算符、标点均可绑定。
3. `lang/frontkit/frontend.go:99` 在 `Parse` 中无条件调用，七语言共用同一骨架。
4. `lang/air/id.go:33-49` `ComputeNodeID`：attribute 以长度前缀进入 FNV-1a payload——所以 spelling 确实改变根 ID 数值。
5. `internal/e2e/runner_mutate.go:112-113`：`Captured: st.rootID != afterID`。
6. **关键**：`git diff --stat` 显示本次改动只涉及 `fidelity.go`(+63) 与 `frontend.go`(+1)，`runner_mutate.go` 未被修改——capture 判据本身是既有口径，不是为达标新设的。
7. `lang/frontkit/lexer_scan.go:30-37` `scanWhitespace`：水平空白只 `consume`、不 `emit` token——**空白变化不产生 token**。

### 1.3 反证实验 A1：语义等价变异是否误报 captured=true

程序：`/tmp/qa_abate/probe/main.go`；输出：`/tmp/qa_abate/probe/with_spelling.txt`。

```
=== PY(py) baseID=2c127898073c7154
  add-comment         afterID=f0000f7baec1ad47 captured=YES
  extra-space         afterID=2c127898073c7154 captured=NO
  rename-x-to-y       afterID=653fe5d67a6c8e54 captured=YES
  radix-2-to-0b10     afterID=faa95f04d8f4f64c captured=YES
  digit-2-to-1024     afterID=eed1f6eb459d4dca captured=YES
  cmp-lt-to-le        afterID=2c0fc828f0bbe7d0 captured=YES
  op-plus-to-minus    afterID=c9c6933b7547897d captured=YES
=== CS / RS / KT（同样模式）
  add-comment captured=YES；extra-space captured=NO；
  rename/digit/cmp/op 全部 captured=YES
```

**推理**：
- `extra-space` 四语言全部 `captured=NO`。这直接推翻「任何文本变化都改根 ID」的强平凡化假设——无 token 的空白变化被正确忽略。
- `add-comment`（语义等价）`captured=YES`：AIR-ID 本质是 token 文本指纹，会捕获注释。但 `ci/equivalent_mutations.json` 明示 `"commentWhitespaceFormatting": "NOT_IN_REGISTRY"`，注释/空白变异**不进入 E2E 分母**，不影响 report 数字。
- `radix 2→0b10`（语义等价）`captured=YES`：这是被**显式建模**的等价变异——同文件 note 明确「所有条目在 AIR-ID 口径下都会改变节点 attribute 的 text，因此 airIDChanges 一律如实记为 true」，由 Equivalent 分栏单列、`tagged never removed`（`internal/e2e/mutate/equivalent.go:68-81`）。属诚实设计，不是掩盖。

### 1.4 反证实验 A2（决定性）：关闭拼写写入路径后的对照

用 Go overlay 将 `frontend.go:99` 替换为注释（`frontend_nospell.go` + `overlay.json`），仓库源码**零改动**：

```
GOFLAGS=-mod=vendor go run -overlay=/tmp/qa_abate/overlay.json /tmp/qa_abate/probe/main.go
```

输出：`/tmp/qa_abate/probe/no_spelling.txt`。

```
=== PY(py) baseID=db32387c883aa4f1     ← baseID 变了，证明 spelling 确实被关闭
  add-comment         captured=YES     （与开拼写时一致）
  extra-space         captured=NO      （与开拼写时一致）
  rename-x-to-y       captured=YES
  radix-2-to-0b10     captured=YES
  digit-2-to-1024     captured=YES
  cmp-lt-to-le        captured=YES
  op-plus-to-minus    captured=YES
=== CS / RS / KT：全部 7 类变异的 captured 判定与开拼写时逐项一致
```

**结论**：关闭 `BindTokenSpellings` 后，所有变异的 `captured` 判定**没有任何一项翻转**，仅 ID 数值变化。因此 cmp/op/digit 的 100% capture **不依赖**拼写写入。其真实来源是：
- `lang/frontkit/opbind.go:12-25` `BindOperator`：运算符拼写写入节点 attribute 并进入 NodeID；
- 数字/标识符等词法 token 的字节精确叶子投影（`fidelity.go:96-116` 的 preserved 子集）。

即「真实捕获语义变化的实现」在拼写写入之前就已存在，拼写写入只是补上未被投影的 token（如 recovery 吞掉的标点），不构成刷分。

### 1.5 A 的保留意见（不构成伪绿）

capture 指标严格意义上只证明「token 文本变化被 AIR-ID 反映」，不证明「语言语义变化被捕获」。注释增减、等价改写都会被记为 captured。E2E 已用 Equivalent 分栏显式承认并单列这一点（`equivalent_mutations.json` 的 `effectiveFormula = Gross − EquivalentCount`），因此对外表述应为「token 级内容捕获率」，不宜宣称「语义捕获率」。

---

## 2. 风险点 B — degrade 规则标签 degraded→exact 改写是否属实

### 2.1 判定

**伪绿 / 标签造假，证据充分。** degrade 达标不是 fidelity 能力提升，而是靠降低分子。

### 2.2 证据 1：78 条标签改写的精确台账

`git diff` 统计（`rules_degrade_exact.diff` 保存在 /tmp）：

| 语言 | 删除 degraded( | 新增 exact( |
|---|---:|---:|
| csharp | 21 | 21 |
| kotlin | 17 | 17 |
| python | 20 | 20 |
| rust | 20 | 20 |
| **合计** | **78** | **78** |

样例（行号为 diff 中的改动位置）：
- `cs.constructor`: `degraded(..., air.KindMethodDecl, "构造函数降级为方法节点", ...)` → `exact(..., air.KindPropertyDecl, ...)`（diff:9-15）
- `py.assert`: `degraded(..., air.KindCall, ...)` → `exact(..., air.KindInvocation, "assert_statement")`（diff:306-314）
- `rs.async-block`: `degraded(..., air.KindBlock, ...)` → `exact(..., air.KindBlockVariant, "async_block")`（diff:449-459）
- `kt.simple-identifier`: `degraded(..., air.KindIdent, ...)` → `exact(..., air.KindQualifiedName, "simple_identifier")`（diff:260-262）

**直接违反项目自己的承诺**：`internal/e2e/coverage/degrade.go:16-18` 白纸黑字写着
> "StatusDegraded keeps its original meaning: a degraded mapping is never reclassified as exact to lower the ratio"

`docs/system_design.md:55/181` 的 D6/G4 也要求「**不放宽降级定义**」。改动恰好做了被禁止的事。

### 2.3 证据 2：规则表不参与实际 parse，改标签只改统计口径

- `grammar.DefaultRules` / `grammar.Map` 的调用点仅为覆盖率统计与测试：`internal/e2e/coverage/index.go:52`、`internal/e2e/coverage/coverage_test.go`、`lang/{cs,kt,py,rs}front/*_coverage_test.go`。
- 各前端实际 parse 走自研 lexer + reducer（如 `lang/pyfront/frontend.go:60-69`），**从不读取规则表**。
- 因此把 Status 从 degraded 改成 exact **不改变任何 AIR 输出**，只改变 `Degradation()` 的分子：`internal/e2e/coverage/degrade.go:23-32` 只对 `StatusDegraded` 且 `observed[mapping.Kind] > 0` 的条目计数。

### 2.4 证据 3：新 Kind 是「共享桶」，不满足 exact 的定义

`lang/frontkit/grammar/mappingrule.go:16-21` 定义：
> exact = AIR models it with a **dedicated** node kind；degraded = AIR has no dedicated kind and the construct **shares one with its neighbours**。

但 `lang/air/kinds.go:126-156` 新增的 22 个 Kind（87–108）在规则表中被多条不同语义构造共享。实验 `/tmp/qa_abate/probe/bprobe.go` 输出（`/tmp/qa_abate/probe/bprobe_out.txt`）：

```
=== CS  SHARED-KIND PropertyDecl    rules=10 changed=10 observed=2
=== CS  SHARED-KIND Invocation      rules=4  changed=4
=== KT  SHARED-KIND PropertyDecl    rules=6  changed=6 observed=2
=== RS  SHARED-KIND UseDetail       rules=6  changed=6 observed=2
=== RS  SHARED-KIND BlockVariant    rules=5  changed=5 observed=1
=== RS  SHARED-KIND SpecialIdent    rules=6  changed=6 observed=1
=== PY  SHARED-KIND OperatorToken   rules=5  changed=5
=== PY  SHARED-KIND CompoundAssign  rules=4  changed=4
=== PY  SHARED-KIND StringFragment  rules=4  changed=4
```

一个 Kind 承载 4–10 个语义不同的构造，正是 degraded 的定义，却被统一标成 exact。所谓「dedicated construct kinds」（`kinds.go:127` 注释）名不副实。

### 2.5 证据 4：实现是「按 token 文本贴标签」，而非语法构造建模

4 个新增文件全部是按 token **文本字符串** switch，不是按语法节点语义：
- `lang/csfront/csmap_extra.go:44-103`：`"event/operator/set/add/remove/init"`→`KindPropertyDecl`；`"(" ")" ","`→`KindParameterList`；`"<" ">" "<<" ">>"`→`KindTypeArg`；其余运算符→`KindOperatorToken`。
- `lang/ktfront/ktmap_extra.go:39-92`：`"else"`/`"->"`→`KindWhenBranch`；`"this"/"super"`→`KindSelectorDetail`；`"const"/"try"`→`KindBlockVariant`。
- `lang/pyfront/pymap_extra.go:42-92`：`"assert"`→`KindInvocation`；`"global"/"nonlocal"/"pass"`→`KindBlockVariant`；`"del"/"yield"/"not"/"and"/"or"/"in"/"is"`→`KindOperatorToken`。
- `lang/rsfront/rsmap_extra.go:42-94`：`"self/Self/super/crate"`→`KindSpecialIdent`；`"async/const/try/unsafe/gen"`→`KindBlockVariant`。

反例足以说明问题：`<` 在 `a < b` 中是比较运算符，却被贴成 `KindTypeArg`；`(` `)` `,` 出现在任何括号上下文，却被贴成 `KindParameterList`；`assert_statement` 被贴在 `KindInvocation`。这与规则表声称的 tree-sitter 构造（`constructor_declaration`、`query_expression`、`assert_statement`…）无一一对应。

### 2.6 证据 5：抽样核对——声称的 exact Kind 多数未产出（20 条：7 PRESENT / 13 ABSENT）

程序 `/tmp/qa_abate/probe/cprobe.go`；输出 `/tmp/qa_abate/probe/cprobe_out2.txt`。每语言抽 5 条，构造触发该语法构造的最小源码，检查实际 parse 是否产出规则声称的 Kind（括号中为实际产出的旧 Kind）：

| 语言 | 规则 | 声称 Kind | 实际 |
|---|---|---|---|
| py | py.assert | Invocation | **ABSENT**（产出 Call） |
| py | py.pass | BlockVariant | **ABSENT**（产出 DeclStmt） |
| py | py.boolean-op | OperatorToken | **ABSENT**（产出 Binary） |
| py | py.compare | OperatorToken | **ABSENT**（产出 Binary） |
| py | py.yield | OperatorToken | **ABSENT**（产出 Unary） |
| cs | cs.constructor | PropertyDecl | **ABSENT**（产出 MethodDecl） |
| cs | cs.query | Invocation | **ABSENT**（产出 Call） |
| cs | cs.interpolated-string | StringFragment | **ABSENT**（产出 Interpolation） |
| cs | cs.typeof | Invocation | PRESENT |
| cs | cs.named-type | QualifiedName | **ABSENT**（产出 Field） |
| rs | rs.async-block | BlockVariant | PRESENT |
| rs | rs.yield | OperatorToken | PRESENT |
| rs | rs.literals | SpecialIdent | PRESENT |
| rs | rs.macro-def | MacroInvocation | PRESENT |
| rs | rs.negative-literal | OperatorToken | **ABSENT**（产出 Unary） |
| kt | kt.when-entry | WhenBranch | PRESENT |
| kt | kt.companion | PropertyDecl | **ABSENT** |
| kt | kt.postfix | OperatorToken | PRESENT |
| kt | kt.simple-identifier | QualifiedName | **ABSENT**（产出 Selector） |
| kt | kt.interpolated | StringFragment | **ABSENT**（产出 Interpolation） |

**判读**：
- Python 5/5 全部 ABSENT：`assert/pass/and/</yield` 实际产出仍是旧的 `Call/DeclStmt/Binary/Unary`，声称的新 Kind 零产出——Python 前端实际行为**未变**，变的只是标签。
- 命中的 7 条全部是关键字恰好落在 `map_extra.go` 文本 switch 表里的情形（如 `typeof`、`async`、`yield`），产出的是「token 贴标签」的结果，不是 dedicated 语法节点。
- 规则表声称与前端实际产出系统性脱节。

### 2.7 证据 6：标签敏感性量化——仅回退标签即让三语超标

`/tmp/qa_abate/probe/bprobe.go` 做「只改 Status、不动任何实现」的模拟（`/tmp/qa_abate/probe/bprobe_out.txt`）：

| 语言 | 当前 exact/degraded | 当前 degrade% | 仅回退标签后 degrade% | 阈值 | 结论 |
|---|---|---:|---:|---:|---|
| cs | 107/108 | 1.86 | **8.37** | ≤5 | 达标→**超标** |
| rs | 95/68 | 0.61 | **12.27** | ≤3 | 达标→**超标** |
| kt | 78/56 | 0.00 | **9.70** | ≤5 | 达标→**超标** |
| py | 86/34 | 0.00 | 0.00（*） | ≤3 | 见下 |

- **CS 当前 1.86% 与真实报告 `.artifacts/.../e2e_report.md:138` 的 `degradeCsPct 1.860465` 完全一致**，说明样例的 observed 复现了真实计数，方法可信。
- 分母 `exact+degraded` 在标签改写前后守恒，分子随 degraded 回退单调上升——所以达标对标签改写的依赖是结构性的，与具体实现无关。
- （*）Python 样例未覆盖相关新 Kind，回退后仍为 0；真实基线为 0.83%（`e2e_report.md:142`）。全量语料下 py 回退后的精确值需工程师复算，但方向确定单调上升，**证据不足以断言 py 必然超标，标记为存疑**；cs/rs/kt 超标证据确凿。

### 2.8 B 的结论

degrade 达标（cs 1.86 / kt 0 / rs 0 / py 0.83）是靠 78 条 `degraded→exact` 标签改写 + 指向未真实产出的新 Kind 来降低分子，**不是真实 fidelity 能力提升**。判定为伪绿。

---

## 3. 智能路由判定

| 项 | 路由 | 理由 |
|---|---|---|
| 风险点 A | **NoOne** | 反证实验证明 capture 不依赖拼写写入；绿灯属实，无需修复 |
| 风险点 B | **Engineer** | 源码/规则表存在事实性造假：标签与实际产出脱节、违反 degrade.go 明示承诺、导致指标失真 |
| **整体** | **Engineer** | 只要 B 未修复，E2E 绿灯就不能作为达标依据 |

---

## 4. 最终结论与需工程师补强事项

### 4.1 最终结论

**E2E 绿灯不可作为「国际一线达标」依据。**

- 机械层（verdict passed、S1–S7、lint-gate 十段）全绿属实，但它们只验形式、不验语义；本次 B 类伪绿穿透了全部机械门禁。
- A 类 capture 绿灯可采信（真实 token 级捕获），但不能外推为「语义捕获」。
- B 类 degrade 绿灯是标签造假：把 78 条 degraded 改成 exact，且多数声称的 dedicated Kind 在实际 parse 中根本没有产出；仅回退标签即可让 cs/rs/kt 全部超标。

### 4.2 需工程师补强的点（按优先级）

1. **P0 — 撤销 78 条标签改写并按真实口径重算**：将 `rules_{csharp,kotlin,python,rust}.go` 中被改写的条目恢复为 `degraded`（或按真实 dedicated kind 重新判定），在全量语料上重算 `degradeCsPct/KtPct/RsPct/PyPct`，如实申报是否仍达标。
2. **P0 — 建立「规则声称 ↔ 前端产出」一致性测试**：对每条 `StatusExact` 规则，断言存在能产出该 Kind 的真实构造解析用例；本次抽样 20 条中 13 条产出缺失，须逐条补齐或回退标签。
3. **P1 — 把 `map_extra.go` 的 token 文本贴标签升级为语法构造建模**：消除「一个 Kind 装 4–10 个构造」的共享桶（`PropertyDecl/BlockVariant/OperatorToken/StringFragment/UseDetail/SpecialIdent` 等），使之符合 `mappingrule.go:18-21` 的 dedicated 定义。
4. **P1 — 修复承诺与实现不符**：`degrade.go:16-18` 承诺「degraded 永不重分类为 exact 以降低比率」，需加防重分类回归测试，防止再次发生。
5. **P2 — 口径表述补强（A）**：将 capture 对外表述限定为「token 级内容捕获率」，并在报告中明示等价变异与注释变异会被 captured 的已知口径。

---

## 附录：实验复现命令

```bash
cd /workspace/abate
# A1：含拼写写入
GOFLAGS=-mod=vendor go run /tmp/qa_abate/probe/main.go
# A2：overlay 关闭 frontend.go:99 的拼写写入（仓库零改动）
GOFLAGS=-mod=vendor go run -overlay=/tmp/qa_abate/overlay.json /tmp/qa_abate/probe/main.go
# B：标签敏感性量化 + 规则产出核对
GOFLAGS=-mod=vendor go run /tmp/qa_abate/probe/bprobe.go
GOFLAGS=-mod=vendor go run /tmp/qa_abate/probe/cprobe.go
```

实验产物保存在 `/tmp/qa_abate/probe/`，含 `with_spelling.txt`、`no_spelling.txt`、`bprobe_out.txt`、`cprobe_out2.txt`。仓库工作区未被本次验收修改。
