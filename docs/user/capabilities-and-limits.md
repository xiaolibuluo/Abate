<!--
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
-->

# 能做什么、不能做什么（能力与边界）

> 本文用**可复现的实测证据**说明 `abate check` 七语言支持的真实边界，
> 避免"多语言 linter"被误读为"编译器级语义分析"。所有结论都能用文末命令在本机复现。

## 一句话定位

Abate 是**结构化静态 linter**：把 7 种语言源码**真正解析成统一的 AIR 抽象语法树**，
在树上做**结构模式匹配**。它不是编译器，不做类型检查、不建调用图、不做跨过程数据流分析。

## 真解析——证据

外界常怀疑"零依赖、不用 tree-sitter / libclang，七语言只能是正则匹配"。
这个怀疑对了一半：**确实零依赖、不引第三方解析库，但不是正则匹配**。解析链路完整且自研：

| 阶段 | 实现 | 代码位置 |
|------|------|----------|
| 词法 | 手写 lexer + token 规格（每语言一份） | `lang/<lang>front/*lexspec.go`、`*tokens.go` |
| 语法 | Pratt 算子优先级爬升解析器，绑定幂为整数数据 | `lang/frontkit/parser.go` |
| AST | 解析产出 `air.Node` 结构化节点（kind / SemKey / 属性 / 子节点） | `lang/air/node.go`、`lang/<lang>front/*map*.go` |
| 规则 | 编译式结构模式：`$_` 任意、`$x` 捕获、`$..` 跨层、`!P` 否定、`$x*` 重复、属性约束 | `lang/rule/pattern.go`、`lang/rule/matcher.go` |

规则匹配的是 **AST 节点结构**，不是文本行。例如一条 Python 规则：

```go
// lang/rule/builtin/py/rules.go
ID:      "PY-PEP8-002",
Pattern: `ClassDecl~_`,   // 匹配 kind=ClassDecl、属性名含下划线的节点
```

`Ident/[A-Z]/`、`ClassDecl~_` 这类模式里的 `Ident` / `ClassDecl` 是 **AST 节点类型**，
正则只是对节点**属性文本**的约束——和"对整个文件跑正则"是两回事。

**实测佐证**：一个含 `if → for → 赋值 → return` 四层嵌套的 Python 文件产出 `total=28`，
命中 `PY-AIR-005/007/009/010/050` 五类结构规则——说明解析器建立了嵌套结构、规则在树上匹配。

## 不是什么——边界

诚实说明边界，比夸大能力更值得信任：

| 边界 | 说明 |
|------|------|
| **容错解析，非严格语法校验** | 语法错误输入不会让命令报错退出，而是产出**部分 AIR + warnings** 继续检查。**不做"语法是否正确"的守门。** |
| **不做类型检查** | 不推断类型、不校验类型匹配；不是 mypy / pyright / rustc 的类型系统。 |
| **不做控制流 / 数据流语义分析** | 不建 CFG、不做跨过程分析、不追踪变量在分支间的传播；发现是**局部结构模式**命中。 |
| **不是 go vet 级深度** | `go vet` 是构建门禁的一部分（`scripts/lint-gate.sh`），但 `abate check` 的七语言规则是**统一 AIR 模板 + 每语言专属规则**，深度与 go vet / eslint / clippy 不可等同。 |
| **规则数量有限** | 共享 AIR 结构模板 72 条 + 每语言专属 12–14 条（Go 侧另有 vet/EG/CRC 系列共 116 条）。覆盖命名/风格/常见结构坏味道，**不是全量语言规范**。 |

## 七语言各自覆盖

按扩展名分派 `.go` `.py` `.js` `.c` `.cs` `.rs` `.kt`（别名 `.jsx/.mjs/.cjs/.h/.kts`）：
每语言在共享 72 条 AIR 模板之上叠加专属规则。实测细节见 [七语言覆盖量化](language-coverage.md)。

| 语言 | 专属规则数 | 覆盖侧重 |
|------|-----------|----------|
| Go | 116 | vet / Effective Go / CodeReviewComments 系列（命中需特定语义上下文） |
| Python | 12 | PEP 8 命名 / 风格 / 惯用法 |
| JavaScript | 12 | 命名 / 风格 / 常见结构 |
| C | 14 | 命名 / 结构 / CERT 安全编码 |
| C# | 12 | 命名 / 风格 / .NET 设计指南 |
| Rust | 14 | 命名 / 风格 / 惯用法 |
| Kotlin | 12 | 命名 / 风格 / 设计约定 |

## 适用场景与不适用场景

**适合**：多语言仓库的统一结构风格巡检（一份报告覆盖 7 种语言）；CI 中确定性的坏味道门禁；需要 JSON/SARIF 机器可读报告的场景。

**不适合**（请改用专用工具）：完整语法错误诊断（用对应编译器/parser）；类型安全/空指针/借用检查（用 mypy / rustc / Roslyn / go vet）；深度语义/跨过程分析（用专用静态分析平台）。

## 复现命令

```bash
bash scripts/evidence-check.sh     # 现场构建 + 7 项带断言复现，全 PASS 才退出 0
bash scripts/language-coverage.sh  # 七语言规则数与实测命中量化表
```

手工拆解见脚本内部：`abate check /tmp/nested.py --format json` 观察嵌套命中；
`abate check /tmp/broken.py` 观察容错（破损输入仍产出报告）；`grep -c require go.mod` 观察零依赖。

## 诚实口径（勿外推）

- **全语料构造覆盖率**：六语实测见 `docs/dev/frontend-coverage/` 各语言报告（C 66.4% / C# 52.6% / JS 89.5% / Kotlin 47.0% / Python 91.8% / Rust 77.9%，均为构造级精确+降级命中口径），
  Go 无嵌入语法映射为 `NOT_MEASURED`；E2E 门禁"达标"仅限冻结单文件（首 unit）口径，
  全语料口径未达该阈值，**不构成"高保真"结论**。
- **竞品对比**：`competitorAbsolute = NOT_MEASURED`，从未真实对拍。
- **derive killing**：static AIR-level approximation，**非真实 go test 执行**。
- **社区验证**：0 star / 0 fork / 0 watcher、个人维护——不刷 star、不买 fork、不伪造 benchmark。
  降低试用门槛靠一键复现 + 公开量化 + 机器可证伪的确定性，欢迎真实失败用例。

一句话：这是一个**能力可自证、但社区尚未验证**的早期项目，请按早期项目标准评估。

## 版本口径

唯一权威构建版本是 `go.mod` 的 `go 1.26.11`，与本机 `go version` 一致。仓库内其他曾出现的 go 版本号均为历史记录，不代表当前构建版本。
