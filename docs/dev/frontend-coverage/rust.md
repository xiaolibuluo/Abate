# Rust 前端构造覆盖清单

- 上游语法事实：`https://github.com/tree-sitter/tree-sitter-rust` 0.24.0
- 上游具体构造总数：**163**
- 映射完备率：**100.0%**（163/163，无未映射空洞）
- 实测覆盖：**127**（**77.9%**）
  - 精确映射覆盖：70
  - 降级映射覆盖：57
- 能力域外（显式登记并给出理由）：0
- 未映射：0
- 认可覆盖率（覆盖 + 能力域外）：**77.9%**
- 真实解析命中的 AIR 类型数：59

统计口径：覆盖率来自对该语言样例语料的**真实解析**，上游构造的 AIR 类型在解析产物中出现过才计为已覆盖。AIR 是中性中间表示，多个上游构造可能共用同一 AIR 类型，因此覆盖以 AIR 类型为粒度；每个上游构造在清单中单列一行，并标注其映射是精确映射还是降级映射。

未覆盖条目的含义：该上游构造在真实解析产物中从未产生过映射所指定的 AIR 类型。它既可能是前端尚未实现该构造，也可能是前端把它降级到了另一个 AIR 类型，二者都记为缺口。

## 未覆盖条目

| 上游构造 | 映射 AIR 类型 | 映射状态 | 原因 |
|---|---|---|---|
| `array_expression` | RsArray | degraded | 全语料零产出，回退 degraded 避免 exact 分母虚增 |
| `array_type` | Array | degraded | 全语料零产出，回退 degraded 避免 exact 分母虚增 |
| `async_block` | RsAsyncBlock | degraded | 全语料零产出，回退 degraded 避免 exact 分母虚增 |
| `base_field_initializer` | RsBaseField | exact | AIR 以独立节点建模该构造 |
| `block_comment` | Comment | exact | AIR 以独立节点建模该构造 |
| `const_block` | RsConstBlock | degraded | 全语料零产出，回退 degraded 避免 exact 分母虚增 |
| `const_item` | ConstDecl | exact | AIR 以独立节点建模该构造 |
| `const_parameter` | ConstDecl | degraded | const 形参降级为常量声明节点 |
| `empty_statement` | RsEmptyStmt | degraded | 全语料零产出，回退 degraded 避免 exact 分母虚增 |
| `extern_crate_declaration` | RsExternCrate | exact | AIR 以独立节点建模该构造 |
| `extern_modifier` | RsExternMod | degraded | 全语料零产出，回退 degraded 避免 exact 分母虚增 |
| `foreign_mod_item` | ModuleDecl | degraded | 外部模块项降级为模块声明节点 |
| `function_modifiers` | RsFnModifiers | degraded | 全语料零产出，回退 degraded 避免 exact 分母虚增 |
| `function_type` | FuncType | degraded | 全语料零产出，回退 degraded 避免 exact 分母虚增 |
| `gen_block` | RsGenBlock | degraded | 全语料零产出，回退 degraded 避免 exact 分母虚增 |
| `generic_type_with_turbofish` | RsTurbofish | exact | AIR 以独立节点建模该构造 |
| `label` | Label | degraded | 全语料零产出，回退 degraded 避免 exact 分母虚增 |
| `line_comment` | Comment | exact | AIR 以独立节点建模该构造 |
| `macro_definition` | RsMacroDef | exact | AIR 以独立节点建模该构造 |
| `metavariable` | RsMetavar | exact | AIR 以独立节点建模该构造 |
| `mod_item` | ModuleDecl | exact | AIR 以独立节点建模该构造 |
| `negative_literal` | RsNegativeLit | exact | AIR 以独立节点建模该构造 |
| `pointer_type` | Pointer | degraded | 全语料零产出，回退 degraded 避免 exact 分母虚增 |
| `reference_expression` | Addr | exact | AIR 以独立节点建模该构造 |
| `return_expression` | Return | exact | AIR 以独立节点建模该构造 |
| `shebang` | Directive | degraded | 全语料零产出，回退 degraded 避免 exact 分母虚增 |
| `shorthand_field_identifier` | RsShorthandIdent | exact | AIR 以独立节点建模该构造 |
| `shorthand_field_initializer` | RsShorthandField | exact | AIR 以独立节点建模该构造 |
| `static_item` | ConstDecl | exact | AIR 以独立节点建模该构造 |
| `super` | RsSuper | exact | AIR 以独立节点建模该构造 |
| `try_block` | RsTryBlock | degraded | 全语料零产出，回退 degraded 避免 exact 分母虚增 |
| `try_expression` | Try | exact | AIR 以独立节点建模该构造 |
| `tuple_type` | RsTupleType | degraded | 全语料零产出，回退 degraded 避免 exact 分母虚增 |
| `union_item` | RsUnion | degraded | 全语料零产出，回退 degraded 避免 exact 分母虚增 |
| `use_as_clause` | RsUseAs | exact | AIR 以独立节点建模该构造 |
| `yield_expression` | RsYield | exact | AIR 以独立节点建模该构造 |

## 能力域外条目

无。

## 全量构造映射清单

| 上游构造 | AIR 类型 | 映射状态 | 真实命中次数 | 映射规则 |
|---|---|---|---|---|
| `abstract_type` | Region | degraded | 1 | `rs.abstract-type` |
| `arguments` | Region | degraded | 1 | `rs.arguments` |
| `array_expression` | RsArray | degraded | 0 | `rs.array` |
| `array_type` | Array | degraded | 0 | `rs.array-type` |
| `assignment_expression` | Assign | exact | 34 | `rs.assign` |
| `associated_type` | TypeDecl | degraded | 2 | `rs.associated-type` |
| `async_block` | RsAsyncBlock | degraded | 0 | `rs.async-block` |
| `attribute` | Attribute | exact | 3 | `rs.attribute` |
| `attribute_item` | Attribute | exact | 3 | `rs.attribute` |
| `await_expression` | Await | exact | 1 | `rs.await` |
| `base_field_initializer` | RsBaseField | exact | 0 | `rs.field-init-base` |
| `binary_expression` | Binary | exact | 6 | `rs.binary` |
| `block` | Block | exact | 55 | `rs.block` |
| `block_comment` | Comment | exact | 0 | `rs.comment` |
| `boolean_literal` | BoolLit | exact | 1 | `rs.bool` |
| `bounded_type` | Region | degraded | 1 | `rs.bounded-type` |
| `bracketed_type` | Named | degraded | 1 | `rs.bracketed-type` |
| `break_expression` | Branch | exact | 2 | `rs.branch` |
| `call_expression` | Call | exact | 54 | `rs.call` |
| `captured_pattern` | Region | degraded | 1 | `rs.patterns` |
| `char_literal` | CharLit | exact | 1 | `rs.char` |
| `closure_expression` | Lambda | exact | 3 | `rs.closure` |
| `closure_parameters` | Region | degraded | 1 | `rs.parameters` |
| `compound_assignment_expr` | RsCompoundAssign | exact | 1 | `rs.assign-compound` |
| `const_block` | RsConstBlock | degraded | 0 | `rs.const-block` |
| `const_item` | ConstDecl | exact | 0 | `rs.const` |
| `const_parameter` | ConstDecl | degraded | 0 | `rs.const-param` |
| `continue_expression` | Branch | exact | 2 | `rs.branch` |
| `crate` | RsCrate | exact | 1 | `rs.literals-crate` |
| `declaration_list` | Region | degraded | 1 | `rs.misc` |
| `doc_comment` | Doc | exact | 3 | `rs.doc-marker` |
| `dynamic_type` | Named | degraded | 1 | `rs.dynamic-type` |
| `else_clause` | If | exact | 1 | `rs.else` |
| `empty_statement` | RsEmptyStmt | degraded | 0 | `rs.empty-stmt` |
| `enum_item` | EnumDecl | exact | 2 | `rs.enum` |
| `enum_variant` | RsEnumVariant | exact | 5 | `rs.enum-variant` |
| `enum_variant_list` | Region | degraded | 1 | `rs.enum-variant-list` |
| `escape_sequence` | RsStringFrag | exact | 12 | `rs.string-part` |
| `expression_statement` | ExprStmt | exact | 35 | `rs.expr-stmt` |
| `extern_crate_declaration` | RsExternCrate | exact | 0 | `rs.extern-crate` |
| `extern_modifier` | RsExternMod | degraded | 0 | `rs.extern-modifier` |
| `field_declaration` | Field | exact | 13 | `rs.field` |
| `field_declaration_list` | Region | degraded | 1 | `rs.field-list` |
| `field_expression` | Field | exact | 13 | `rs.field-expr` |
| `field_identifier` | RsFieldIdent | exact | 5 | `rs.literals-field` |
| `field_initializer` | RsFieldInit | exact | 5 | `rs.field-init` |
| `field_initializer_list` | Region | degraded | 1 | `rs.field-init-list` |
| `field_pattern` | Region | degraded | 1 | `rs.patterns` |
| `float_literal` | NumLit | exact | 31 | `rs.number` |
| `for_expression` | RangeFor | exact | 2 | `rs.for` |
| `for_lifetimes` | Region | degraded | 1 | `rs.lifetime` |
| `foreign_mod_item` | ModuleDecl | degraded | 0 | `rs.foreign-mod` |
| `fragment_specifier` | Region | degraded | 1 | `rs.misc` |
| `function_item` | FuncDecl | exact | 23 | `rs.func` |
| `function_modifiers` | RsFnModifiers | degraded | 0 | `rs.function-modifiers` |
| `function_signature_item` | RsFuncSig | exact | 2 | `rs.func-signature` |
| `function_type` | FuncType | degraded | 0 | `rs.fn-type` |
| `gen_block` | RsGenBlock | degraded | 0 | `rs.gen-block` |
| `generic_function` | RsGenericFunc | exact | 3 | `rs.generic-func` |
| `generic_pattern` | Region | degraded | 1 | `rs.patterns` |
| `generic_type` | GenericInst | exact | 9 | `rs.generic-inst` |
| `generic_type_with_turbofish` | RsTurbofish | exact | 0 | `rs.turbofish` |
| `higher_ranked_trait_bound` | Region | degraded | 1 | `rs.trait-bounds` |
| `identifier` | Ident | exact | 341 | `core.identifier` |
| `if_expression` | If | exact | 1 | `rs.if` |
| `impl_item` | RsImpl | exact | 2 | `rs.impl` |
| `index_expression` | Index | exact | 3 | `rs.index` |
| `inner_attribute_item` | Attribute | exact | 3 | `rs.attribute` |
| `inner_doc_comment_marker` | Doc | exact | 3 | `rs.doc-marker` |
| `integer_literal` | NumLit | exact | 31 | `rs.number` |
| `label` | Label | degraded | 0 | `rs.label` |
| `let_chain` | Region | degraded | 1 | `rs.misc` |
| `let_condition` | Region | degraded | 1 | `rs.misc` |
| `let_declaration` | VarDecl | exact | 28 | `rs.let` |
| `lifetime` | Region | degraded | 1 | `rs.lifetime` |
| `lifetime_parameter` | Region | degraded | 1 | `rs.lifetime` |
| `line_comment` | Comment | exact | 0 | `rs.comment` |
| `loop_expression` | RsLoop | exact | 1 | `rs.loop` |
| `macro_definition` | RsMacroDef | exact | 0 | `rs.macro-def` |
| `macro_invocation` | RsMacroCall | exact | 6 | `rs.macro-call` |
| `macro_rule` | Region | degraded | 1 | `rs.token-tree` |
| `match_arm` | Case | exact | 6 | `rs.case` |
| `match_block` | Match | exact | 2 | `rs.match` |
| `match_expression` | Match | exact | 2 | `rs.match` |
| `match_pattern` | Region | degraded | 1 | `rs.patterns` |
| `metavariable` | RsMetavar | exact | 0 | `rs.literals-metavar` |
| `mod_item` | ModuleDecl | exact | 0 | `rs.mod` |
| `mut_pattern` | Region | degraded | 1 | `rs.patterns` |
| `mutable_specifier` | RsMutable | exact | 3 | `rs.mutable-specifier` |
| `negative_literal` | RsNegativeLit | exact | 0 | `rs.negative-literal` |
| `never_type` | Named | degraded | 1 | `rs.never-type` |
| `or_pattern` | Region | degraded | 1 | `rs.patterns` |
| `ordered_field_declaration_list` | Region | degraded | 1 | `rs.field-list` |
| `outer_doc_comment_marker` | Doc | exact | 3 | `rs.doc-marker` |
| `parameter` | Region | degraded | 1 | `rs.parameters` |
| `parameters` | Region | degraded | 1 | `rs.parameters` |
| `parenthesized_expression` | Region | degraded | 1 | `rs.paren-expr` |
| `pointer_type` | Pointer | degraded | 0 | `rs.pointer-type` |
| `primitive_type` | Named | exact | 1 | `rs.named-type` |
| `qualified_type` | Named | degraded | 1 | `rs.qualified-type` |
| `range_expression` | Region | degraded | 1 | `rs.range-expr` |
| `range_pattern` | Region | degraded | 1 | `rs.patterns` |
| `raw_string_literal` | StrLit | exact | 6 | `rs.string` |
| `ref_pattern` | Region | degraded | 1 | `rs.patterns` |
| `reference_expression` | Addr | exact | 0 | `rs.addr` |
| `reference_pattern` | Region | degraded | 1 | `rs.patterns` |
| `reference_type` | Ref | exact | 10 | `rs.ref-type` |
| `remaining_field_pattern` | Region | degraded | 1 | `rs.patterns` |
| `removed_trait_bound` | Region | degraded | 1 | `rs.trait-bounds` |
| `return_expression` | Return | exact | 0 | `rs.return` |
| `scoped_identifier` | Selector | exact | 9 | `rs.scoped-identifier` |
| `scoped_type_identifier` | Named | degraded | 1 | `rs.qualified-type` |
| `scoped_use_list` | RsScopedUse | exact | 4 | `rs.scoped-use` |
| `self` | RsSelf | exact | 7 | `rs.literals-self` |
| `self_parameter` | Region | degraded | 1 | `rs.parameters` |
| `shebang` | Directive | degraded | 0 | `rs.shebang` |
| `shorthand_field_identifier` | RsShorthandIdent | exact | 0 | `rs.literals-shorthand` |
| `shorthand_field_initializer` | RsShorthandField | exact | 0 | `rs.field-init-shorthand` |
| `slice_pattern` | Region | degraded | 1 | `rs.patterns` |
| `source_file` | Package | exact | 27 | `rs.file` |
| `static_item` | ConstDecl | exact | 0 | `rs.const` |
| `string_content` | RsStringFrag | exact | 12 | `rs.string-part` |
| `string_literal` | StrLit | exact | 6 | `rs.string` |
| `struct_expression` | CompositeLit | exact | 3 | `rs.struct-expr` |
| `struct_item` | StructDecl | exact | 5 | `rs.struct` |
| `struct_pattern` | Region | degraded | 1 | `rs.patterns` |
| `super` | RsSuper | exact | 0 | `rs.literals-super` |
| `token_binding_pattern` | Region | degraded | 1 | `rs.patterns` |
| `token_repetition` | Region | degraded | 1 | `rs.token-tree` |
| `token_repetition_pattern` | Region | degraded | 1 | `rs.patterns` |
| `token_tree` | Region | degraded | 1 | `rs.token-tree` |
| `token_tree_pattern` | Region | degraded | 1 | `rs.patterns` |
| `trait_bounds` | Region | degraded | 1 | `rs.trait-bounds` |
| `trait_item` | TraitDecl | exact | 2 | `rs.trait` |
| `try_block` | RsTryBlock | degraded | 0 | `rs.try-block` |
| `try_expression` | Try | exact | 0 | `rs.try` |
| `tuple_expression` | Tuple | exact | 1 | `rs.tuple` |
| `tuple_pattern` | Region | degraded | 1 | `rs.patterns` |
| `tuple_struct_pattern` | Region | degraded | 1 | `rs.patterns` |
| `tuple_type` | RsTupleType | degraded | 0 | `rs.tuple-type` |
| `type_arguments` | RsTypeArgs | exact | 9 | `rs.type-args` |
| `type_binding` | TypeDecl | degraded | 2 | `rs.type-binding` |
| `type_cast_expression` | Cast | exact | 2 | `rs.cast` |
| `type_identifier` | Named | exact | 1 | `rs.type-identifier` |
| `type_item` | TypeDecl | exact | 2 | `rs.type` |
| `type_parameter` | GenericParam | exact | 2 | `rs.generic-param` |
| `type_parameters` | Region | degraded | 1 | `rs.type-params` |
| `unary_expression` | Unary | exact | 2 | `rs.unary` |
| `union_item` | RsUnion | degraded | 0 | `rs.union` |
| `unit_expression` | Region | degraded | 1 | `rs.unit-expr` |
| `unit_type` | Named | degraded | 1 | `rs.unit-type` |
| `unsafe_block` | RsUnsafeBlock | exact | 2 | `rs.unsafe-block` |
| `use_as_clause` | RsUseAs | exact | 0 | `rs.use-as` |
| `use_bounds` | Import | degraded | 4 | `rs.use-bounds` |
| `use_declaration` | Import | exact | 4 | `rs.use` |
| `use_list` | RsUseList | exact | 1 | `rs.use-list` |
| `use_wildcard` | RsUseWildcard | exact | 1 | `rs.use-wildcard` |
| `variadic_parameter` | Region | degraded | 1 | `rs.parameters` |
| `visibility_modifier` | RsVisibility | exact | 6 | `rs.visibility-modifier` |
| `where_clause` | Region | degraded | 1 | `rs.where-clause` |
| `where_predicate` | Region | degraded | 1 | `rs.where-clause` |
| `while_expression` | While | exact | 2 | `rs.while` |
| `yield_expression` | RsYield | exact | 0 | `rs.yield` |

## 降级率口径说明（诚实披露）

`degrade*Pct` 按每语言**首个真实解析 unit（单文件）**计算，并对应冻结阈值；全语料使用同一 `Degradation` 公式，数值不同——两者公式相同、仅统计范围不同。

- 本语言：单文件 **0.61%**（≤3 达标）；全语料 **42.33%**（超阈）。
- 结论限定：达标仅限冻结的单文件口径，不代表全语料/语料级降级率达标；go 为 **NOT_MEASURED**。
- 数据来源：`docs/i18n/QA-ANTI-FRAUD-AUDIT.md` 第 7 维度；数值真实、口径不同，禁止把单文件达标外推为全语料达标。

