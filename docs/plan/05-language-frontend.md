# Abate L4 语言前端规划 — 十一模块极致细节

- 文档编号：PLAN-05　状态：生效　上游：PLAN-00
- 覆盖模块：LNG-AIR / LNG-GO / LNG-PY / LNG-JS / LNG-C / LNG-RS / LNG-CS / LNG-KT / LNG-RULE / LNG-CHK / LNG-MUT
- 通用约束（十一模块共同遵守）：
  1. 一切前端产出 AIR：语言差异止步于前端；规则、检查、变异派生只面向 AIR——一次实现，七语言受益。
  2. 源码级接入铁律：第三方解析库一律五步并入法；禁止 API 接入（不进程调用、不 gRPC、不网络、不运行时动态加载）。
  3. 前端确定性：同源字节流解析 100 次，AIR 序列化哈希一致；位置信息字节精确。
  4. 分层质量标准：LNG-AIR 为全引擎共享地基，执行内核层标准（覆盖率 ≥95%/85%）；其余 L4 模块执行生态层标准（≥85%/75%）。
  5. 增量性：AIR NodeID 跨解析稳定（同构节点同 ID），支撑增量解析与增量检查。

---

## 模块 LNG-AIR：统一 IR

### 1. 模块定位与职责

定义 Abate 中间表示：覆盖七语言公共构造超集的节点模型、语义注记体系、稳定节点 ID、版本化二进制序列化。AIR 是规则、检查、变异三条管线的汇合点：语言前端只产出 AIR，下游只消费 AIR。其设计质量直接决定全引擎上限（十年架构）。

### 2. 包结构与文件布局

```
lang/air/
  node.go        // Node 接口、节点注册表
  kinds.go       // NodeKind 全枚举（五大类）
  pos.go         // Position：行/列/字节偏移
  sem.go         // 语义注记：类型/作用域/CFG/纯度
  id.go          // NodeID：内容哈希（同构感知）
  unit.go        // Unit/File/Package：编译单元模型
  encode.go      // 版本化二进制序列化
  decode.go      // 反序列化 + 版本协商
  roundtrip.go   // 往返校验辅助
  air_test.go
```

### 3. 核心类型与接口（必须实现）

- NodeKind 全枚举（首版必须齐备，五大类）：
  - 声明类（15）：Package/Import/TypeDecl/ConstDecl/VarDecl/FuncDecl/MethodDecl/StructDecl/InterfaceDecl/EnumDecl/ClassDecl/TraitDecl/ModuleDecl/AnnotationDecl/UsingDecl
  - 语句类（22）：Block/ExprStmt/DeclStmt/If/For/While/DoWhile/RangeFor/Switch/Case/Select/Defer/GoStmt/Return/Branch/Throw/Try/Catch/Finally/With/Label/Match
  - 表达式类（26）：Ident/NumLit/StrLit/CharLit/BoolLit/NilLit/Binary/Unary/Call/Index/Slice/Field/Selector/CompositeLit/FuncLit/Arrow/Lambda/Cast/TypeAssert/Assign/IncDec/Star/Addr/Ternary/Spread/ChanOp/MatchExpr/Await
  - 类型类（15）：Named/Pointer/Slice/Array/Map/Chan/FuncType/StructType/InterfaceType/GenericInst/GenericParam/Union/Optional/Tuple/Ref
  - 其他类（6）：Comment/Doc/Attribute/Decorator/Directive/Region
- Node 接口：Kind() NodeKind / Pos() Position / End() Position / Children() []Node / Sem(key SemKey) any
- Position：Line（1 起）/Column（1 起）/Offset（0 起字节偏移）/EndOffset；文件归属经 Unit
- 语义注记 SemInfo：TypeRef（类型引用：已解析指向 NodeID，未解析保留名）/ScopeID（作用域链）/CFGEdges（控制流边：Reachable/Fallthrough/LoopBack/Unreachable）/Mutability（immutable/mutable/unknown）/Purity（pure/side-effect/unknown）
- NodeID：uint64 内容哈希 = FNV-1a(kind + 规范化自身属性 + 有序子 ID)；同构子树必同 ID
- Unit：Files（Path/ContentHash/LineEnding/Encoding）+ Package（Name/Imports）；Unit 指纹参与实验确定性哈希
- 序列化：二进制格式 = magic(4B) + version(2B) + unit 头 + 节点表 + 字符串表 + 语义表 + SHA-256 校验和；版本协商：解码端声明支持区间 [min, current]，超区间拒绝并报双方版本

### 4. 关键算法与流程

1. NodeID 稳定算法：自底向上递归哈希（叶节点先哈希，内部节点组合有序子 ID）；增量解析只需重哈希受影响路径，成本 O(编辑距离) 而非 O(文件大小)。
2. 字符串表：标识符/字面量/类型名去重，索引引用；基准验收：典型项目序列化体积缩减 ≥60%。
3. 语义惰性附加：解析产出纯语法 AIR（零语义开销）；类型/作用域/CFG 注记由语言前端第二遍按需附加。
4. 往返校验：源码 → AIR → 位置区间回放 → 区间覆盖检查（无缝隙、无越界重叠）；逐语言再生策略由各前端声明（见公共契约 F3）。

### 5. 依赖与约束

- 依赖：仅标准库（hash/fnv、encoding/binary、sort）。
- 被依赖：全部 LNG-*、KRN-MUT（变异点面向 AIR 节点）、KRN-VER（等价性静态检查）。
- 禁止：任何语言专有类型泄漏进 AIR（超集中立模型）；任何 I/O（只接受内存字节）；map 迭代序参与任何哈希（全部排序规范化）。

### 6. 测试要求

- 单元：每 NodeKind 构造/遍历/注记读写；ID 稳定性（同构同 ID、一字节差异必不同 ID）；序列化往返（各 kind 载荷无损）。
- 版本兼容：旧版编码新版解码（区间边界）；新版编码旧版解码（拒绝并报版本）。
- 模糊：随机 AIR 树序列化/反序列化往返无损、不 panic（60s 快速档，语料入库）。
- 性能：100 万节点 Unit 序列化 P99 ≤500ms、反序列化 P99 ≤700ms（基准报告入库）。

### 7. 验收标准（可量化）

1. NodeKind 全集覆盖七语言公共构造（七语言构造覆盖清单入库；未覆盖构造必须列明理由与降级映射）。
2. ID 稳定性/往返/版本协商用例 100% 通过。
3. 模糊 60s 无 panic；语料入库。
4. 序列化性能达标；覆盖率 ≥95%/85%；门禁全绿。

---

## 语言前端公共契约（LNG-GO/PY/JS/C/RS/CS/KT 共享）

### F1. 前端接口（七语言前端必须实现）

- Frontend 接口：
  - Parse(unit SourceUnit) (*air.Unit, []Diagnostic, error)：全量解析。
  - ParseIncremental(unit SourceUnit, old *air.Unit, edits []Edit) (*air.Unit, Diff, error)：增量解析（经 NodeID 复用未变子树）。
- SourceUnit：Files（Path/Content []byte）/Language/Dialect（方言选项：TS strict、C++20、Python 3.12 等）。
- Diagnostic：Line/Column/Severity（error/warning）/Code（LANG-<LANG>-E/W-xxx）/Message；错误上限 100（超出终止解析，防错误风暴）。
- 确定性承诺：同输入 100 次解析 AIR 哈希一致（每语言强制验收项）。

### F2. 五步并入法（统一流程，逐语言差异见各模块）

1. 源码 vendor：固定版本入 third_party/<lib>/（记录上游 commit 哈希、许可证审查结论、裁剪清单）。
2. 二次更改：确定性化（消除哈希迭代序等不稳定因素）、诊断结构化（错误归一为 Diagnostic）、位置对齐（统一 AIR Position 语义）。
3. 优化：剥离非解析路径（lint 规则/CLI/构建集成），只留解析器与必要类型信息。
4. 深度集成：AST → AIR 映射层（逐语言映射表入库）+ 语义注记附加。
5. 封装：只暴露 Frontend 接口；上游专有 AST 类型禁止出现在公开 API（审查清单项）。

### F3. 逐语言验收公共项（每语言前端必须全部满足）

1. 基准语料：≥1000 文件（开源项目快照 + 自构边界用例），解析成功率 ≥99.5%（C/C++ ≥99%），剩余文件必须有失败分析。
2. 确定性：100 次解析 AIR 哈希一致。
3. 保真：往返等价抽查 ≥100 文件（Go/Python/JS/C#/Kotlin：token 级等价；C/C++/Rust：位置区间全覆盖 + 语义等价）。
4. 吞吐：≥1MB/s（纯解析单线程；C/C++ 预算放宽至 ≥0.5MB/s）。
5. 映射层覆盖率 ≥85%/75%。

---

## 模块 LNG-GO：Go 前端

### 1. 模块定位与职责

Go 语言前端（M3 首语言，验证整个 AIR 架构的参考实现）与 Abate 自检的狗粮基础。源：go/ast + go/parser + go/types（标准库，零外部依赖——唯一无需 vendor 的语言，最安全）。

### 2. 包结构与文件布局

```
lang/gofront/
  frontend.go     // Frontend 接口实现
  mapdecl.go      // 声明类映射
  mapstmt.go      // 语句类映射
  mapexpr.go      // 表达式类映射
  maptype.go      // 类型类映射
  sem.go          // go/types 语义注记附加
  regen.go        // 源码再生（往返校验）
  gofront_test.go
```

### 3. 映射表要点（Go 专有构造 → AIR）

- go/defer/select → Defer/Select；channel 操作（send/receive/ok 型）→ ChanOp。
- 泛型（TypeParam/类型实参）→ GenericParam/GenericInst。
- struct tag → Attribute；//go: 指令 → Directive；iota 常量块 → ConstDecl（iota 语义经注记保留）。
- go/types 附加：TypeRef 解析至 NodeID；ScopeID 作用域链；Purity 启发式（无副作用调用 + 无赋值 → pure）。

### 4. 关键算法与流程

1. 增量解析：go/parser 无原生增量；经 AIR NodeID 子树复用实现（未变子树直接复用，变更路径重解析）。
2. 再生策略：token 级等价（保留 Comment 与 Position；v1 不承诺字节等价——ADR 记录决策）。

### 5. 依赖与约束

- 依赖：标准库（go/parser、go/types）、LNG-AIR。
- 禁止：第三方 Go 解析库；go/ast 类型泄漏进公开 API。

### 6. 测试要求

- 单元：映射表逐项测试（每个 Go 构造 → AIR 节点断言）；语义注记正确性（类型解析/作用域/纯度）。
- 狗粮：解析 Abate 仓库自身全部 .go 文件零错误。
- 确定性与基准：100 次哈希一致；1000 文件基准语料。

### 7. 验收标准（可量化）

1. 映射表逐项测试 100% 通过；Go 构造覆盖清单全勾选。
2. Abate 仓库自解析零错误；往返 token 等价抽查 100 文件通过。
3. 确定性 100 次一致；吞吐 ≥1MB/s。
4. 覆盖率 ≥85%/75%；门禁全绿。

---

## 模块 LNG-PY：Python 前端

### 1. 模块定位与职责

Python 语言前端。主源：ruff 的解析前端（rustpython-parser 血统，Rust 实现）——只 vendor 其 parser crate，绝不引入 ruff 的 lint 规则源码（规则自研铁律）。辅源：tree-sitter-python（容错增量解析，用于损坏文件降级路径）。

### 2. 五步并入法执行细节

1. vendor：ruff parser crate 固定版本入 third_party/ruff-parser/（MIT/Apache-2.0，审查结论可并入）；tree-sitter-python 入 third_party/tree-sitter-python/（MIT）。
2. 二次更改：确定性化（Rust HashMap 迭代序 → BTreeMap 规范化）；错误归一 Diagnostic；位置（0 起）对齐 AIR（1 起）。
3. 优化：剥离 ruff 的 linter/formatter/CLI，只留 parser + AST；tree-sitter 剥离语法生成产物，只留增量解析 API。
4. 深度集成：Rust → C ABI 静态库（cargo build --release + cbindgen 头文件）→ CGO 进程内链接（源码级，无进程调用）；Python AST → AIR 映射层。
5. 封装：Frontend 接口；ruff/tree-sitter 类型禁止泄漏。

### 3. 映射表要点（Python 专有构造 → AIR）

- 装饰器 → Decorator；推导式（list/dict/set）→ CompositeLit（注记 comprehension，for/if 子句保留为子节点）。
- match/case → Match/Pattern；async/await → Await（async 经注记）；海象 := → Assign（注记 walrus）。
- f-string → StrLit（注记 fstring，插值表达式保留为子节点）；缩进块 → Block（注记 indent-based）。

### 4. 关键算法与流程

1. 双解析器策略：主 ruff parser（全保真 AST，遇错即停）；失败降级 tree-sitter（容错，产出部分 AIR + warning 诊断）——降级在 AIR 显式标记（注记 partial=true），下游规则跳过 partial 单元。
2. 语义注记：Python 无静态类型系统；TypeRef 默认未解析（保留名），type hints 解析为 TypeRef。

### 5. 依赖与约束

- 依赖：third_party/ruff-parser、third_party/tree-sitter-python（CGO）、LNG-AIR。
- 禁止：ruff lint 规则源码（铁律）；运行时调用 python 解释器/pip（纯静态解析）。

### 6. 测试要求

- 单元：映射逐项；降级路径（损坏文件 → tree-sitter 部分 AIR + warning）。
- 基准：CPython 标准库 Lib/ 快照 + 自构边界用例凑足 1000 文件。
- 确定性：100 次哈希一致。

### 7. 验收标准（可量化）

1. 映射逐项与降级路径用例 100% 通过。
2. 基准成功率 ≥99.5%；往返 token 等价抽查 100 文件通过。
3. 确定性 100 次一致；吞吐 ≥1MB/s。
4. 覆盖率 ≥85%/75%；门禁全绿；ruff 规则源码零引入（扫描报告入库）。

---

## 模块 LNG-JS：JS/TS 前端

### 1. 模块定位与职责

JS/TS 语言前端。源：swc（Rust 实现，高性能）——只 vendor 其 parser（swc_ecma_parser + AST），不引入 transform/lint 路径。

### 2. 五步并入法执行细节

1. vendor：swc parser crate 固定版本入 third_party/swc/（Apache-2.0）。
2. 二次更改：确定性规范化；诊断归一；位置对齐。
3. 优化：剥离 transform/minify/lint/插件系统，只留 parser + AST。
4. 深度集成：Rust → C ABI 静态库 → CGO；JS/TS AST → AIR 映射。
5. 封装：Frontend 接口。

### 3. 映射表要点

- JSX → CompositeLit（注记 jsx）；TS 类型注解 → TypeRef；装饰器 → Decorator。
- 可选链 ?. 与空值合并 ?? → Field/Binary（注记 optional-chain/nullish）。
- TS 泛型 → GenericInst/GenericParam；enum → EnumDecl；interface → InterfaceDecl。
- async/await → Await；模板字面量 → StrLit（插值保留为子节点）。

### 4. 关键算法与流程

1. 方言协商：扩展名（.ts/.tsx/.js/.jsx/.mjs）+ Dialect 选项（strict/target ES 版本）→ swc 解析配置。
2. TS 类型保留策略：AIR 保留全量类型信息（TypeRef），不做类型擦除（AIR 面向分析而非编译——与 tsc emit 语义不同，ADR 记录）。

### 5. 依赖与约束

- 依赖：third_party/swc（CGO）、LNG-AIR。禁止：swc transform/lint 路径引入；运行时 node 调用。

### 6. 测试要求

- 单元：映射逐项（JSX/TS 专有构造各 ≥3 用例）；方言协商矩阵。
- 基准：TypeScript 编译器测试套件快照 + 自构边界用例凑足 1000 文件。
- 确定性：100 次哈希一致。

### 7. 验收标准（可量化）

1. 映射逐项与方言矩阵用例 100% 通过。
2. 基准成功率 ≥99.5%；往返 token 等价抽查 100 文件通过。
3. 确定性 100 次一致；吞吐 ≥1MB/s。
4. 覆盖率 ≥85%/75%；门禁全绿。

---

## 模块 LNG-C：C/C++ 前端

### 1. 模块定位与职责

C/C++ 语言前端。源：libclang（LLVM）——vendor LLVM/clang 源码，构建最小 libclang 子集（只含解析 + AST 访问 + 类型信息），CGO 进程内链接。最重的 vendor（LLVM 源码树庞大），优化步骤是关键战役。

### 2. 五步并入法执行细节

1. vendor：LLVM/clang 固定版本入 third_party/llvm/（Apache-2.0 + LLVM Exception，审查结论：友好可并入）。
2. 二次更改：确定性化（剔除编译产物中文件时间戳等非确定因素）；诊断结构化；位置对齐。
3. 优化（关键战役）：构建配置剥离 clang driver/Sema 代码生成/后端，只留 lexer/parser/AST + 基础类型查询；构建脚本产出 libclang-min.a（目标 ≤50MB；对照全量 libclang 约 150MB，对比报告入库）。
4. 深度集成：CGO 链接 libclang-min；CXCursor 遍历 → AIR 映射；类型信息（clang_getTypeDeclAsWritten 等）→ TypeRef。
5. 封装：Frontend 接口；CXString/CXCursor 禁止泄漏。

### 3. 映射表要点

- 预处理指令（#include/#define/#ifdef）→ Directive（注记 preprocessor；不展开宏——分析导向策略，展开信息按需附加）。
- C++ 模板 → GenericInst（注记 template，实例化信息按需附加）；命名空间 → ModuleDecl。
- 智能指针 → TypeRef（注记 smart-ptr）；#pragma → Directive。

### 4. 关键算法与流程

1. 头文件策略：头文件搜索路径由 SourceUnit 提供（虚拟 FS 映射——sim/fs 供头文件，禁止真实文件系统读取）。
2. 模板实例化深度上限 16（对齐 clang 默认），超出产生 warning 诊断。

### 5. 依赖与约束

- 依赖：third_party/llvm（CGO）、LNG-AIR、sim/fs（头文件虚拟化，经端口接口）。
- 禁止：真实文件系统读取；调用 clang 二进制；LLVM 后端代码路径引入。

### 6. 测试要求

- 单元：映射逐项（C/C++ 专有构造各 ≥3 用例）；头文件虚拟 FS 解析用例。
- 基准：LLVM 测试套件抽样 + 自构边界用例凑足 1000 文件。
- 确定性：100 次哈希一致。

### 7. 验收标准（可量化）

1. 映射逐项与虚拟头文件用例 100% 通过。
2. 基准成功率 ≥99%；位置区间全覆盖抽查 100 文件通过。
3. 确定性 100 次一致；吞吐 ≥0.5MB/s（C/C++ 预算放宽）。
4. libclang-min.a ≤50MB（构建报告入库）；覆盖率 ≥85%/75%；门禁全绿。

---

## 模块 LNG-RS：Rust 前端

### 1. 模块定位与职责

Rust 语言前端。源：syn（Rust 纯解析库，不含 clippy 检查逻辑——基于官方规范自研规则，铁律）。quote 面向代码生成，前端不需要，不引入。

### 2. 五步并入法执行细节

1. vendor：syn 固定版本入 third_party/syn/（MIT/Apache-2.0 双许可）。
2. 二次更改：确定性规范化；syn Error → Diagnostic。
3. 优化：只保留解析路径（syn 全量特性按需裁剪：保留 item/expr/stmt/pat/ty，剥离 proc-macro 支撑路径）。
4. 深度集成：Rust → C ABI 静态库 → CGO；syn AST → AIR 映射。
5. 封装：Frontend 接口。

### 3. 映射表要点

- 宏调用（含 macro_rules! 声明）→ Call（注记 macro）；属性 #[derive] → Attribute。
- 生命周期 → TypeRef（注记 lifetime）；trait → TraitDecl；impl 块 → MethodDecl（注记 impl）。
- 模式匹配 → Match/Pattern；? 运算符 → 注记 try-op；async fn → FuncDecl（注记 async）；unsafe 块 → Block（注记 unsafe）。

### 4. 依赖与约束

- 依赖：third_party/syn（CGO）、LNG-AIR。禁止：clippy 检查逻辑引入；运行时 cargo/rustc 调用。

### 5. 测试要求

- 单元：映射逐项（Rust 专有构造各 ≥3 用例）。
- 基准：Rust 标准库源码抽样 + 自构边界用例凑足 1000 文件。
- 确定性：100 次哈希一致。

### 6. 验收标准（可量化）

1. 映射逐项用例 100% 通过。
2. 基准成功率 ≥99.5%；位置区间全覆盖抽查 100 文件通过。
3. 确定性 100 次一致；吞吐 ≥1MB/s。
4. 覆盖率 ≥85%/75%；门禁全绿；clippy 逻辑零引入（扫描报告入库）。

---

## 模块 LNG-CS：C# 前端

### 1. 模块定位与职责

C# 语言前端。源：Roslyn（微软官方，MIT）——vendor Roslyn 源码，NativeAOT 编译为 C ABI 静态库（进程内，无 .NET 运行时进程调用），CGO 链接。只借用 Roslyn 解析与语义分析能力，不引入第三方分析器规则（铁律）。

### 2. 五步并入法执行细节

1. vendor：Roslyn 固定版本入 third_party/roslyn/（MIT）。
2. 二次更改：确定性规范化；Microsoft.CodeAnalysis.Diagnostic → AIR Diagnostic。
3. 优化：剥离 IDE 特性/分析器/代码修复，只留 SyntaxTree 解析 + 语义模型子集；NativeAOT 发布单文件静态库 libroslyn-abate.a（目标 ≤80MB）。
4. 深度集成：C ABI 包装层（extern C 导出 parse/语义附加入口）→ CGO；Roslyn SyntaxNode → AIR 映射。
5. 封装：Frontend 接口；Microsoft.CodeAnalysis 类型禁止泄漏。

### 3. 映射表要点

- LINQ 查询表达式 → Call（注记 linq）；特性 → Attribute；事件 → FieldDecl（注记 event）。
- 属性/索引器 → FuncDecl（注记 property/indexer）；可空引用类型 → Optional。
- using 指令/语句 → UsingDecl/Block（注记 using-stmt）；async/await → Await；record → ClassDecl（注记 record）。

### 4. 依赖与约束

- 依赖：third_party/roslyn（CGO）、LNG-AIR。禁止：第三方分析器规则引入；运行时 dotnet 进程调用。

### 5. 测试要求

- 单元：映射逐项（C# 专有构造各 ≥3 用例）。
- 基准：dotnet/runtime 源码抽样 + 自构边界用例凑足 1000 文件。
- 确定性：100 次哈希一致。

### 6. 验收标准（可量化）

1. 映射逐项用例 100% 通过。
2. 基准成功率 ≥99.5%；往返 token 等价抽查 100 文件通过。
3. 确定性 100 次一致；吞吐 ≥1MB/s。
4. libroslyn-abate.a ≤80MB（构建报告入库）；覆盖率 ≥85%/75%；门禁全绿。

---

## 模块 LNG-KT：Kotlin 前端

### 1. 模块定位与职责

Kotlin 语言前端。源：kotlin-compiler（JetBrains 官方，Apache-2.0）——vendor 其解析前端（lexer/parser/PSI）源码，Kotlin/Native 编译为原生静态库（C 互操作），CGO 链接。只拿 AST，规则自研（铁律）。

### 2. 五步并入法执行细节

1. vendor：kotlin-compiler 固定版本入 third_party/kotlin-compiler/（Apache-2.0）。
2. 二次更改：确定性规范化；编译器诊断 → Diagnostic。
3. 优化：剥离后端（JVM/JS/Native 代码生成）与 IDE 集成，只留前端（lexer/parser/PSI）；Kotlin/Native 产出 libkotlin-front.a（目标 ≤60MB）。
4. 深度集成：C ABI 包装层 → CGO；Kotlin PSI → AIR 映射。
5. 封装：Frontend 接口；JetBrains 专有类型禁止泄漏。

### 3. 映射表要点

- 协程（suspend fun）→ FuncDecl（注记 suspend）；密封类 → ClassDecl（注记 sealed）。
- 伴生对象 → ClassDecl（注记 companion）；智能转换 → Cast（注记 smart）。
- 扩展函数 → FuncDecl（注记 extension，接收者记为首参数）。
- when 表达式 → Match/MatchExpr；数据类 → ClassDecl（注记 data）。
- 空安全（?. 与 !!）→ Field/Binary（注记 null-safe/assert-not-null）。

### 4. 依赖与约束

- 依赖：third_party/kotlin-compiler（CGO）、LNG-AIR。禁止：JVM 运行时依赖；kotlinc 进程调用。

### 5. 测试要求

- 单元：映射逐项（Kotlin 专有构造各 ≥3 用例）。
- 基准：Kotlin 标准库源码抽样 + 自构边界用例凑足 1000 文件。
- 确定性：100 次哈希一致。

### 6. 验收标准（可量化）

1. 映射逐项用例 100% 通过。
2. 基准成功率 ≥99.5%；往返 token 等价抽查 100 文件通过。
3. 确定性 100 次一致；吞吐 ≥1MB/s。
4. libkotlin-front.a ≤60MB（构建报告入库）；覆盖率 ≥85%/75%；门禁全绿。

---

## 模块 LNG-RULE：规则引擎

### 1. 模块定位与职责

YAML 规则 DSL + AIR 结构模式匹配引擎：七语言官方规范规则集的加载、匹配、报告。每条规则同时服务检查（LNG-CHK）与变异派生（LNG-MUT）——规则双用途。

### 2. 包结构与文件布局

```
lang/rule/
  rule.go        // Rule 类型：规则模型全字段
  dsl.go         // YAML DSL 解析与校验
  pattern.go     // 结构模式编译器（模式 → 匹配器）
  matcher.go     // 匹配器：AIR 结构匹配
  predicate.go   // 谓词表达式（where 子句）
  registry.go    // 规则注册表（命名空间 ID）
  builtin/       // 官方规范规则集（内置）
    go/         // ≥100 条（Effective Go + CodeReviewComments + vet 语义）
    python/     // ≥60 条（PEP 8 + PEP 484）
    js/         // ≥60 条（ECMAScript + TS Handbook）
    c/          // ≥60 条（CERT + 官方编码标准）
    rust/       // ≥60 条（Rust API Guidelines）
    csharp/     // ≥60 条（.NET 设计规范）
    kotlin/     // ≥60 条（Kotlin 编码规范）
  rule_test.go
```

### 3. 核心类型与接口（必须实现）

- Rule 全字段：
  - id：命名空间（GO-EG-001 / GO-CRC-014 / PY-PEP8-003；格式 LANG-来源-序号）。
  - language / severity（error/warning/info）/ category（naming/style/error-prone/performance/concurrency/boundary）。
  - pattern：结构模式（节点 kind + 子模式 + 元变量 $x + 通配符 $_ + 深度匹配 $..）。
  - where：谓词列表（捕获类型约束、值匹配、作用域约束、语义注记检查）。
  - message：违规消息模板（引用捕获）。
  - fix：修复建议（描述 + 可选自动修复模式）。
  - reference：官方规范出处（强制字段——无出处规则拒绝加载）。
  - mutation：变异派生声明（可选：本规则可派生 M 族算子，派生类型与参数）。
- 模式编译器：模式 AST → 匹配器（一次编译，跨文件复用）。
- 匹配器：Match(node air.Node) []Capture；支持通配符（$_ 任意节点）、捕获（$x 绑定）、深度匹配（$.. 后代匹配）、负模式（!模式）、重复模式（$x* 零或多）。
- 谓词：类型约束（捕获为字面量/标识符）、值匹配（正则/数值区间）、语义注记检查（纯度/可变性/类型解析状态）。

### 4. 关键算法与流程

1. 匹配管线：文件 → AIR → 按语言过滤规则集 → 逐规则匹配器并行匹配（只读无锁）→ 捕获谓词过滤 → 违规列表。
2. 性能：匹配器一次编译；单规则单文件匹配 O(AIR 节点数)；全规则集（100 条）单文件 P99 ≤50ms（基准验收项）。
3. 规则加载四项校验：DSL schema 校验 + reference 非空 + id 唯一 + pattern 可编译——四项全过才注册（坏规则加载期拒绝，不带到运行期）。

### 5. 依赖与约束

- 依赖：LNG-AIR、third_party/yaml（源码级）。
- 被依赖：LNG-CHK（检查）、LNG-MUT（派生）。
- 禁止：规则内嵌语言专有逻辑（规则只面向 AIR）；正则滥用（结构模式优先——正则仅用于字面量值匹配）。

### 6. 测试要求

- 单元：DSL 每特性（通配/捕获/深度/负模式/重复）各 ≥3 用例；谓词各类型用例。
- 规则集测试：每条内置规则配正反代码夹具（各 ≥1 用例；七语言约 460 条 → ≥920 夹具）。
- 性能：全规则集单文件 P99 ≤50ms 基准。
- 模糊：随机模式不 panic。

### 7. 验收标准（可量化）

1. DSL 全特性用例通过；加载四项校验用例通过。
2. 内置规则集：Go ≥100、其余六语言各 ≥60（合计 ≥460），每条有正反夹具与官方规范出处（扫描检查：出处非空率 100%）。
3. 全规则集性能基准达标（报告入库）。
4. 覆盖率 ≥85%/75%；门禁全绿。

---

## 模块 LNG-CHK：检查器

### 1. 模块定位与职责

自研语言代码检查器：编排管线（文件 → 前端 → AIR → 规则引擎 → 违规 → 聚合 → 报告）。CLI 入口 abate check。狗粮基础：Abate 仓库用自研检查器自检。

### 2. 包结构与文件布局

```
lang/check/
  check.go       // Check 主入口：编排管线
  aggregate.go   // 违规聚合（去重/排序/严重级统计）
  sarif.go       // SARIF 2.1.02 输出
  console.go     // 控制台人读输出
  exit.go         // 退出码约定
  check_test.go
```

### 3. 核心类型与接口（必须实现）

- Check(paths []string, opts Options) (Result, error)：主入口。
- Options：Languages（语言过滤）/Rules（规则过滤）/Severity 阈值/Format（sarif/console/json）。
- Result：Violations（规则 id/文件/行/列/消息/捕获）+ Summary（分级统计/分规则统计）。
- 退出码约定：0 = 零违规；1 = 存在违规（门禁 FAIL 语义）；2 = 内部错误（与违规严格区分）。
- SARIF 输出：SARIF 2.1.02（runs[].results[].ruleId/location/message），过官方校验器（验收项）。

### 4. 关键算法与流程

1. 管线：路径展开 → 语言检测（扩展名 + 内容嗅探）→ 并行解析（文件级独立，结果按路径字典序确定性排序）→ 规则匹配 → 聚合（去重：同规则同位置；排序：文件字典序 → 行 → 列 → 规则 id）。
2. 增量检查：结合 AIR NodeID 只重查变更子树与受影响规则（v1 全量检查，增量列为 M5 增强项——ADR 记录）。

### 5. 依赖与约束

- 依赖：LNG-*（前端/规则）、RPT-SARIF（格式复用）。
- 禁止：检查器内嵌语言专有逻辑（纯编排层）。

### 6. 测试要求

- 单元：聚合去重/排序；退出码三情形。
- 端到端：夹具工程（含已知违规）→ check → 违规清单与期望完全一致。
- 狗粮（验收）：Abate 仓库自检 0 违规；与 golangci-lint（Go）对拍：差异报告逐项说明（规则语义差异/规则集差异）。

### 7. 验收标准（可量化）

1. 端到端夹具用例 100% 通过；退出码约定用例通过。
2. SARIF 过官方校验器。
3. 狗粮：Abate 仓库自检 0 违规；golangci-lint 对拍差异报告入库且逐项说明。
4. 覆盖率 ≥85%/75%；门禁全绿。

---

## 模块 LNG-MUT：算子派生

### 1. 模块定位与职责

规则到 M 族变异算子的桥梁：从规则 mutation 声明自动派生变异算子。规则即变异点：边界规则派生边界扰动算子、类型规则派生类型断言扰动算子、并发规则派生并发扰动算子。

### 2. 包结构与文件布局

```
lang/derive/
  derive.go      // Derive 主入口：规则集 → 算子集
  table.go       // 派生表：规则类别 → 算子类型映射
  validate.go    // 派生校验：算子在变异实验中的可用性
  derive_test.go
```

### 3. 核心类型与接口（必须实现）

- Derive(rules []rule.Rule) ([]MutationSpec, error)：主入口。
- MutationSpec：SourceRuleID（来源规则）/OperatorID（派生算子 id：m-derived-<规则 id>）/TargetKind（AIR 节点 kind）/Transform（扰动语义：算符替换/边界平移/常量扰动/条件取反/调用删除）。
- 派生表（首版五类）：
  - 边界规则（比较算符、区间检查）→ 边界扰动（< ↔ <=、> ↔ >=）。
  - 常量规则（魔法数字、字符串重复）→ 常量扰动（+1/-1/置零/置空）。
  - 条件规则（复合条件、短路求值）→ 条件取反/分支删除。
  - 类型规则（类型断言、转换）→ 类型扰动（断言目标扰动）。
  - 调用规则（错误忽略、返回值检查）→ 调用删除/返回值扰动。
- validate：每条派生算子必须在 KRN-MUT 通过变异实验可用性测试（生成变异体 → 运行 → 判定有效）。

### 4. 关键算法与流程

1. 派生管线：规则集过滤（含 mutation 声明）→ 派生表查表 → MutationSpec 生成 → 算子注册（入 KRN-MUT 注册表）。
2. 派生确定性：同规则集派生同算子集（排序输出，哈希可校验）。

### 5. 依赖与约束

- 依赖：LNG-RULE（规则模型）、KRN-MUT（算子注册表，经端口接口）。
- 禁止：派生逻辑硬编码具体语言（派生表只依赖规则类别与 AIR 节点 kind）。

### 6. 测试要求

- 单元：五类派生各有用例；确定性（同规则集双派生，算子集哈希一致）。
- 端到端：≥20 条规则派生算子，每条在 KRN-MUT 跑变异实验并产出有效判定。

### 7. 验收标准（可量化）

1. 五类派生表用例通过；派生确定性用例通过。
2. ≥20 条规则成功派生算子且在变异实验中可用（M3 验收项）。
3. 覆盖率 ≥85%/75%；门禁全绿。

---

## 附：L4 合入总验收（十一模块齐备后）

1. 端到端狗粮：Abate 仓库（Go）经自研检查器 0 违规；≥20 条规则派生变异算子，对 Abate 自身代码跑变异实验产出有效报告（完整狗粮闭环）。
2. 七语言基准：每语言 1000 文件解析成功率 ≥99.5%（C/C++ ≥99%）、确定性 100 次一致、吞吐达标（基准报告入库）。
3. AIR 序列化性能与版本兼容用例通过；七语言构造覆盖清单全勾选。
4. 十一模块覆盖率汇总报告入库；五段门禁 + 基准回归全绿。
