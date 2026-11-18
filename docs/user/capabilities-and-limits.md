<!--
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
-->

# Abate 能做什么、不能做什么（能力与边界说明）

> 本文目的：用**可复现的实测证据**说清楚 `abate check` 七语言支持的真实能力边界，
> 避免"多语言 linter"被误读为"编译器级语义分析 / go vet 那种吃透七种语言"。
> 所有结论都能用文末命令在本机复现。

## 一句话定位

Abate 是一个**结构化静态检查器（structural linter）**：它把 7 种语言的源码
**真正解析成统一的 AIR 抽象语法树**，然后在 AST 上做**结构模式匹配**来发现问题。
它不是编译器，不做类型检查、不建调用图、不做跨过程数据流分析。

## 它"真解析"——证据

外界常怀疑"零依赖、不用 tree-sitter / libclang，七语言只能是正则匹配"。
这个怀疑对了一半：**确实零依赖、不引第三方解析库，但不是正则匹配**。

解析链路是完整的、自研的：

| 阶段 | 实现 | 代码位置 |
|------|------|----------|
| 词法 | 手写 lexer + token 规格（每种语言一份） | `lang/<lang>front/*lexspec.go`、`*tokens.go` |
| 语法 | Pratt 算子优先级爬升解析器，绑定幂为整数数据、无浮点/时钟/随机态 | `lang/frontkit/parser.go` |
| AST | 解析产出 `air.Node` 结构化节点（kind / SemKey / 属性 / 子节点） | `lang/air/node.go`、`lang/<lang>front/*map*.go` |
| 规则 | 编译式结构模式：`$_` 任意节点、`$x` 捕获、`$..` 跨层、`!P` 否定、`$x*` 重复、属性约束 `= != ~ ^ + /regex/` | `lang/rule/pattern.go`、`lang/rule/matcher.go` |

规则匹配的是 **AST 节点结构**，不是文本行。例如一条 Python 规则：

```go
// lang/rule/builtin/py/rules.go
ID:      "PY-PEP8-002",
Pattern: `ClassDecl~_`,   // 匹配 kind=ClassDecl、属性名含下划线的节点
```

`Ident/[A-Z]/`、`ClassDecl~_`、`NumLit/^[0-9]{4,}$/` 这类模式里的 kind
（`Ident` / `ClassDecl` / `NumLit`）是 **AST 节点类型**，正则只是对节点**属性文本**的约束——
和"对整个文件跑正则"是两回事。

**实测佐证**：一个含 `if → for → 赋值 → return` 四层嵌套的 Python 文件，
产出 `files=1, total=28`，命中 `PY-AIR-005/007/009/010/050` 五类结构规则——
说明解析器建立了嵌套结构、规则在树上匹配，而非逐行文本扫描。

## 它"不是什么"——边界（同样重要）

诚实说明边界，比夸大能力更值得信任：

| 边界 | 说明 |
|------|------|
| **容错解析，非严格语法校验** | 语法错误输入不会让命令报错退出，而是产出**部分 AIR + diagnostics** 继续检查（`lang/frontkit/frontend.go` 明确设计如此）。实测 `def f(:` / `x =` 这类破损输入仍产出 12 条 warning。**它不做"语法是否正确"的守门。** |
| **不做类型检查** | 不推断类型、不校验类型匹配。它不是 mypy / pyright / rustc 的类型系统。 |
| **不做控制流 / 数据流语义分析** | 不建 CFG、不做跨过程（interprocedural）分析、不追踪变量在分支间的传播。发现是**局部结构模式**命中。 |
| **不是 go vet 级深度** | Go 语言侧的 `go vet` 分析器是构建门禁的一部分（`scripts/lint-gate.sh`），但 `abate check` 的七语言规则是**统一 AIR 模板 + 每语言十余条专属规则**，深度与专用工具（go vet / eslint / clippy / Pylint）不可等同。 |
| **规则数量有限** | 共享 AIR 结构模板 72 条 + 每语言专属 12–14 条（Go 侧另有 vet/effector/CRC 系列）。覆盖的是命名/风格/常见结构坏味道，**不是全量语言规范**。 |

## 七语言各自的实际覆盖

`abate check` 按扩展名分派：`.go` `.py` `.js` `.c` `.cs` `.rs` `.kt`
（别名 `.jsx/.mjs/.cjs/.h/.kts`）。每语言在共享 72 条 AIR 模板之上叠加专属规则：

| 语言 | 专属规则数 | 覆盖侧重 |
|------|-----------|----------|
| Go | 116（vet / effector / CRC 系列） | 最接近 go vet 风格，含多组 vet 规则 |
| Python | 12 | PEP 8 命名 / 风格 / 惯用法 |
| JavaScript | 12 | 命名 / 风格 / 常见结构 |
| C | 14 | 命名 / 结构 / 惯用法 |
| C# | 12 | 命名 / 风格 / 惯用法 |
| Rust | 14 | 命名 / 风格 / 惯用法 |
| Kotlin | 12 | 命名 / 风格 / 惯用法 |

> 规则条数按 `lang/rule/builtin/` 各目录 `ID:` 声明统计。专属规则 + 共享模板共同构成
> `check --format json` 报告里 `summary.rules` 字段（示例文件显示为 72，即共享模板数）。

## 适用场景与不适用场景

**适合**：
- 多语言仓库的统一结构风格巡检，一份报告覆盖 7 种语言；
- CI 中做确定性的坏味道门禁（固定输入输出逐字节可复现，可用 `verify` 比对）；
- 需要机器可读报告（JSON / SARIF）并做二次处理的场景。

**不适合**（请改用专用工具）：
- 要完整语法错误诊断 → 用对应语言的编译器 / parser；
- 要类型安全、空指针、借用检查 → 用 mypy / rustc / Roslyn / go vet 等；
- 要深度语义、跨过程分析 → 用专用静态分析平台。

## 复现命令

**首选：一条命令复现全部 7 项声明**（现场从源码构建二进制 + 带断言）：

```bash
bash scripts/evidence-check.sh     # 全部 PASS 才退出 0；含真解析/容错/确定性/七语言/帮助文案
bash scripts/language-coverage.sh  # 产出七语言规则数与实测命中量化表
```

以下为逐条手工拆解（脚本内部即这些步骤）：

```bash
# 1. 真解析证据：嵌套结构产出多条结构规则命中
printf 'def outer(x):\n    if x:\n        for i in range(3):\n            y = i\n            return y\n    return x\n' > /tmp/nested.py
abate check /tmp/nested.py --format json --out /tmp/nested.json
# 观察 summary.total > 0 且 byRule 含 PY-AIR-* 多类

# 2. 容错边界：语法错误输入不报错退出，产出 partial 诊断
printf 'def f(:\n  x =\n' > /tmp/broken.py
abate check /tmp/broken.py --format json --out /tmp/broken.json
# 观察：仍产出报告（非语法守门）

# 3. 零依赖证据：无 vendor/、go.mod 无第三方 require
grep -c 'require' go.mod   # 输出 0
ls vendor 2>/dev/null      # 不存在

# 4. 结构化匹配证据：看规则 DSL 匹配的是 AST 节点 kind
sed -n '20,40p' lang/rule/builtin/py/rules.go
```

## 版本口径

唯一权威构建版本是 **`go.mod` 的 `go 1.26.11`**，与本机 `go version` 输出一致。
仓库内其他曾出现的 go 版本号（如设计草案中的旧值）均为历史记录，不代表当前构建版本。

## 社区验证状态（如实说明）

截至本文更新，仓库为 **0 star / 0 fork / 0 watcher、个人维护**。这是客观事实，
不回避、也不粉饰。需要说清楚三点：

1. **它意味着什么**：缺少外部背书，意味着"社区验证"这一维度目前是空白——
   没有人用真实项目大规模跑过、没有第三方复现、没有 issue 驱动的缺陷发现。
   对"能不能放心用于生产"这个问题，star 数本身不是能力证明，但它确实反映
   "被多少人检验过"。**当前答案：检验很少，请勿据此认为已久经考验。**
2. **我们不会做什么**：不会刷 star、不会买 fork、不会伪造用户或 benchmark。
   任何声称的"社区认可"若非真实，都是对这个项目可信度的二次伤害——而可信度
   正是本文要修复的东西。
3. **降低试用门槛的实际做法**（不靠造假，靠可验证）：
   - 全部能力声明附 `scripts/evidence-check.sh` 一键复现，任何人可自建证据、不必信文档；
   - `scripts/language-coverage.sh` 公开每语言真实规则数，不藏拙；
   - 确定性可用 `abate verify` 逐字节比对，Claim 可被机器证伪；
   - 欢迎用真实仓库试跑并提 issue——**具体的失败用例比 star 更有价值**。

一句话：这是一个**能力可自证、但社区尚未验证**的早期项目。请按早期项目的标准评估它。
