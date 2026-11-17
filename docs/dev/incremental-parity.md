# 增量解析算法对照（tree-sitter ↔ Abate frontkit）

本文逐条对照上游 tree-sitter 的增量解析算法与 Abate `lang/frontkit/incremental.go`
的实现，标注每一项的状态：

- **已实现且等价**：Abate 侧的行为与上游语义一致。
- **已实现但简化**：Abate 侧有对应机制，但语义比上游窄或更粗。
- **未实现**：Abate 侧没有对应机制。

## 0. 对照材料与事实来源

| 事实 | 来源 | 只读用途 |
|---|---|---|
| 增量算法本体 | `third_party/_sources/tree-sitter-runtime/lib/src/parser.c` | `ts_parser__reuse_node`、`ts_parser__breakdown_top_of_stack`、`ts_parser__breakdown_lookahead`、`ts_parser__can_reuse_first_leaf` |
| 树编辑与失效标记 | `third_party/_sources/tree-sitter-runtime/lib/src/tree.c`、`subtree.c` | `ts_tree_edit`、`ts_subtree_edit`、`ts_subtree_has_changes` |
| 变更区间计算 | `third_party/_sources/tree-sitter-runtime/lib/src/get_changed_ranges.c` | `ts_range_array_get_changed_ranges`、`ts_subtree_get_changed_ranges` |
| 外部扫描器状态 | `third_party/_sources/tree-sitter-runtime/lib/src/parser.c` | `ts_parser__external_scanner_serialize` / `deserialize` / `scan` |
| 各语言语法产物 | `third_party/_sources/tree-sitter-*/src/{node-types.json,grammar.json,parser.c,scanner.c}` | 语法事实与外部扫描器契约 |
| Abate 侧 | `lang/frontkit/incremental.go`、`lang/air/edit.go`、`lang/air/diff.go` | 增量入口、编辑应用、差异计算 |

上游 C 运行时（`tree-sitter-runtime/lib`）已随仓库入库，因此本文的对照**不依赖外部
网络或推断**，每一条都能指到具体函数。生成的语言 `src/parser.c` 只是解析表与
`ts_symbol_metadata[]`，不含增量算法本体。

## 1. 总览

| # | 上游机制 | 上游入口 | Abate 对应 | 状态 |
|---|---|---|---|---|
| 1 | 编辑定位（字节区间 + 行列点） | `TSInputEdit` | `air.Edit{Path, Offset, EndOffset, Text}` | 已实现且等价 |
| 2 | 编辑应用与重叠拒绝 | `ts_subtree_edit` 之前的缓冲处理 | `air.ApplyEdits` | 已实现且等价 |
| 3 | 树失效标记（把编辑传播到子树） | `ts_tree_edit` → `ts_subtree_edit` | 无节点级传播 | 未实现 |
| 4 | 节点级复用 | `ts_parser__reuse_node` | 无节点级复用 | 未实现 |
| 5 | 复用前拆栈（breakdown） | `ts_parser__breakdown_top_of_stack`、`ts_parser__breakdown_lookahead` | 无 | 未实现 |
| 6 | 首叶可复用性判定 | `ts_parser__can_reuse_first_leaf` | 无 | 未实现 |
| 7 | 外部扫描器状态一致性 | `ts_parser__external_scanner_serialize` / `deserialize` | 无外部扫描器，无需状态快照 | 未实现（不适用） |
| 8 | 文件级内容哈希复用 | `ts_parser` 无对应 | `reusableRoots` | 已实现但简化 |
| 9 | 方言/语言配置变更导致失效 | 由 `TSParser` 重建隐式承担 | `reusableRoots` 的方言比对 | 已实现且等价 |
| 10 | 变更区间（changed ranges） | `ts_range_array_get_changed_ranges` | 无 | 未实现 |
| 11 | 差异计算 | `ts_tree_get_changed_ranges` 之外另有结构比较 | `air.ComputeDiff`（NodeID 多重集） | 已实现但简化 |
| 12 | 复用模式诚实标注 | `LOG` 调试日志 | `lang.<tag>.incremental` 注记 | 已实现且等价 |

## 2. 逐条说明

### 2.1 编辑定位与应用（#1、#2）

上游 `TSInputEdit` 同时携带字节区间与 `TSPoint`（行列），以便在子树上做区间
调整。Abate 的 `air.Edit` 只带字节区间，`air.ApplyEdits` 按 `Offset` 升序排序后
依次应用，并显式拒绝重叠区间：

- 排序保证结果与调用方切片顺序无关（确定性）。
- 重叠检查返回 `air.ErrInvalidEdit`，不静默丢弃。

**结论：已实现且等价**（在字节区间这一层等价；行列点未建模，但 AIR 的
`air.Position` 在节点构造时重新计算，不依赖增量阶段的行列调整）。

### 2.2 树失效标记（#3，未实现）

上游 `ts_tree_edit` 把编辑沿树向下传播：落在子树之前的编辑只平移 padding，
落在子树内部的编辑会重算 size/padding 并把该子树标记为 `has_changes`；只有
标记过的子树才会在复用阶段被跳过。

Abate 侧没有保留旧树的结构化失效信息：增量入口只保留**旧单元的根节点与文件
引用**，不保留每个节点的可复用性元数据。

**结论：未实现。** 补偿手段是 #8 的文件级哈希复用——只要文件内容变了，整棵
文件重解析，等价于「该文件的所有节点都失效」这一最保守语义。

### 2.3 节点级复用与拆栈（#4、#5、#6，未实现）

上游 `ts_parser__reuse_node` 在旧树上做一次前序扫描，对每个候选子树依次检查：

1. 字节位置是否落在当前解析位置（`byte_offset > position` 停止，
   `byte_offset < position` 则下潜或跳过）；
2. 外部扫描器状态是否一致（`ts_subtree_external_scanner_state_eq`）；
3. 子树是否带 `has_changes` / `is_error` / `is_missing` / `is_fragile` 标记；
4. 首叶在当前解析状态下的词法模式是否兼容
   （`ts_parser__can_reuse_first_leaf` 比较 `TSLexerMode`）。

不满足条件时，上游调用 `ts_parser__breakdown_top_of_stack` 把已归约的父节点
拆回子节点重新入栈，从而把复用粒度从「整棵子树」退化为「更小的子树」。

Abate 侧没有节点级复用，也没有拆栈：粒度只有「整个文件」。

**结论：未实现。** 这是与上游差距最大的一项，也是 Abate 增量能力的主要上限。

### 2.4 外部扫描器状态一致性（#7，不适用）

上游要求外部扫描器实现 `serialize`/`deserialize`，把扫描器私有状态随 token
一起保存，复用节点时比对状态是否一致，不一致则不复用。

Abate 的六语言前端没有外部扫描器：词法完全由 `frontkit.Lexer` 的确定性表驱动，
不存在需要在增量之间保持的扫描器私有状态。

**结论：未实现（不适用）。** 这一点在语义上等价于「状态总是相等」，因此不会
造成复用错误。

### 2.5 文件级内容哈希复用（#8，已实现但简化）

`lang/frontkit/incremental.go` 的 `reusableRoots` 是 Abate 侧唯一的复用机制：

- 逐文件比较 `air.SourceFile.Digest()` 与旧单元记录的内容哈希；
- 哈希一致则直接复用旧根节点，不重新解析该文件；
- 哈希不一致则把该文件放进待重解析集合。

**结论：已实现但简化。** 语义正确（内容未变则结果必然未变），但复用粒度是文件，
不是节点。

### 2.6 方言变更导致失效（#9，已实现且等价）

上游没有「方言」概念：方言变化等价于换一个 `TSLanguage` 重新建解析器，旧树整体
作废。Abate 侧的六语言前端都带方言（如 Python 3.8–3.13、TypeScript 版本），同一
段字节在不同方言下的归约结果可以不同，因此方言变化必须让旧结果失效。

`reusableRoots` 现在接收 `oldDialect` 与 `dialect` 两个参数：方言不同、或旧单元
根本没有记录方言（`lang.<tag>.dialect` 注记缺失）时，一律不复用。

> **本轮修复的真实缺陷**：在修复前，`ParseIncremental` 只按内容哈希决定复用，
> 不看方言。同一份源码在 `3.11` 与 `3.12` 之间切换时，旧方言的语法树会被原样
> 复用并作为结果返回——结果是错的，而且没有任何可见迹象（复用模式注记仍然是
> `file-level-reuse`）。修复后补了回归测试
> `lang/frontkit/frontkit_incremental_test.go:TestIncrementalDialectChangeForcesFullRebuild`
> 与 `lang/pyfront/pyfront_incremental_test.go:TestPyIncrementalDialectChangeForcesRebuild`。

**结论：已实现且等价。**

### 2.7 变更区间（#10，未实现）

上游 `ts_range_array_get_changed_ranges` 用双游标遍历新旧两棵树，输出「语法结构
发生变化的字节区间」列表，供编辑器做最小重绘。

Abate 侧不产出变更区间：`air.Diff` 只给出新增/删除的 `NodeID` 多重集。

**结论：未实现。**

### 2.8 差异计算（#11，已实现但简化）

`air.ComputeDiff` 比较新旧单元的节点 ID 多重集，得到新增与删除的节点集合。
上游除区间差异外还有 `ts_subtree_compare` 的结构比较（用于在错误恢复时择优选择
两棵候选树）。

**结论：已实现但简化。** 差异语义足以驱动「哪些符号消失了」这一类分析，但不
等价于上游的区间级差异。

### 2.9 复用模式诚实标注（#12，已实现且等价）

上游用 `LOG(...)` 打印复用/拒绝原因。Abate 侧把复用模式写进单元注记：

| 注记 | 取值 | 含义 |
|---|---|---|
| `lang.<tag>.incremental` | `full-rebuild` | 没有复用任何文件，整单元重建 |
| `lang.<tag>.incremental` | `file-level-reuse` | 按文件内容哈希复用了若干文件 |
| `lang.<tag>.reusedFiles` | 十进制整数 | 实际复用的文件数 |

**结论：已实现且等价。** 任何一次增量结果都自带「这是文件级复用还是全量重建」
的可见标记，不会把全量重建伪装成节点级增量。

## 3. 已知限制与后果

1. **增量只到文件级。** 编辑一个文件必然触发该文件全量重解析；同一文件内的
   其它函数/类不会复用。这是性能上限，不是正确性问题。
2. **AIR 单元不承载诊断。** `lang.<tag>.diagnostics` 注记只统计本次被重解析文件
   的诊断，被复用文件的诊断不会重新产生。增量结果的诊断集合因此不是全量诊断。
3. **无变更区间输出。** 消费方无法据此做最小重绘，只能整体替换。
4. **跨文件复用不做依赖分析。** 一个文件改动不会触发引用它的其它文件重解析。

## 4. 正确性验证

以下断言全部是**真实编辑场景**下的回归测试，命令：

```sh
GOFLAGS=-mod=vendor go test -race -count=1 ./lang/frontkit/ ./lang/pyfront/
```

| 测试 | 场景 | 断言 |
|---|---|---|
| `TestIncrementalEditsMatchFullParse` | 插入 / 删除 / 替换 token | 增量结果与全量解析逐节点一致 |
| `TestIncrementalDialectChangeForcesFullRebuild` | 方言切换 | 强制全量重建，不复用旧方言子树 |
| `TestIncrementalReusesUnchangedFile` | 多文件只改一个 | 未改文件按哈希复用，复用数正确 |
| `TestIncrementalRejectsBadInput` | 空前端 / 非法编辑 / 不存在的路径 | 返回分类错误，不 panic |
| `TestPyIncrementalMatchesFullParse` | Python 插入 / 删除 / 替换 token | 增量结果与全量解析逐节点一致 |
| `TestPyIncrementalDialectChangeForcesRebuild` | Python 方言切换 | 强制重建 |

「逐节点一致」的比较口径：节点种类、属性、位置区间、子节点结构全部相同，且剔除
`lang.<tag>.incremental` / `lang.<tag>.reusedFiles` 两个注记（它们描述的是复用
方式，不是语法结果）。

## 5. 后续可选工作

按投入产出排序，均**不属于本轮交付范围**：

1. 在 AIR 节点上增加「可复用性」元数据（编辑区间、词法状态指纹），把复用粒度
   从文件降到节点。
2. 让 `air.Diff` 输出变更区间，对齐 `ts_range_array_get_changed_ranges`。
3. 让 AIR 单元承载逐文件诊断，使增量结果的诊断集合与全量一致。
4. 对六语言前端做跨文件依赖分析，编辑被引用文件时触发引用方重解析。
