# Python 前端构造覆盖清单

- 上游语法事实：`https://github.com/tree-sitter/tree-sitter-python` 0.25.0
- 上游具体构造总数：**122**
- 映射完备率：**100.0%**（122/122，无未映射空洞）
- 实测覆盖：**109**（**89.3%**）
  - 精确映射覆盖：80
  - 降级映射覆盖：29
- 能力域外（显式登记并给出理由）：3
- 未映射：0
- 认可覆盖率（覆盖 + 能力域外）：**91.8%**
- 真实解析命中的 AIR 类型数：58

统计口径：覆盖率来自对该语言样例语料的**真实解析**，上游构造的 AIR 类型在解析产物中出现过才计为已覆盖。AIR 是中性中间表示，多个上游构造可能共用同一 AIR 类型，因此覆盖以 AIR 类型为粒度；每个上游构造在清单中单列一行，并标注其映射是精确映射还是降级映射。

未覆盖条目的含义：该上游构造在真实解析产物中从未产生过映射所指定的 AIR 类型。它既可能是前端尚未实现该构造，也可能是前端把它降级到了另一个 AIR 类型，二者都记为缺口。

## 未覆盖条目

| 上游构造 | 映射 AIR 类型 | 映射状态 | 原因 |
|---|---|---|---|
| `augmented_assignment` | PyAugAssign | exact | AIR 以独立节点建模该构造 |
| `constrained_type` | GenericParam | degraded | PEP695 受约束类型形参降级为泛型形参节点 |
| `dotted_name` | PyQualifiedName | exact | AIR 以独立节点建模该构造 |
| `member_type` | PyQualifiedName | exact | AIR 以独立节点建模该构造 |
| `not_operator` | PyNotOp | exact | AIR 以独立节点建模该构造 |
| `splat_type` | PySplatType | degraded | 类型形参解包降级，语料零产出 |
| `type_alias_statement` | TypeDecl | degraded | PEP695 type 别名语料零产出，锁定降级口径 |
| `type_conversion` | PyTypeConversion | exact | AIR 以独立节点建模该构造 |
| `type_parameter` | GenericParam | degraded | PEP695 type 形参语料零产出，锁定降级口径 |
| `union_type` | Union | degraded | PEP604 联合类型语料零产出，锁定降级口径 |

## 能力域外条目

| 上游构造 | 原因 |
|---|---|
| `chevron` | Python 2 遗留 print chevron，不在 Python 3.8-3.12 能力域 |
| `exec_statement` | Python 2 遗留 exec 语句，不在 Python 3.8-3.12 能力域 |
| `print_statement` | Python 2 遗留 print 语句，不在 Python 3.8-3.12 能力域 |

## 全量构造映射清单

| 上游构造 | AIR 类型 | 映射状态 | 真实命中次数 | 映射规则 |
|---|---|---|---|---|
| `aliased_import` | Import | exact | 6 | `py.import` |
| `argument_list` | Region | degraded | 1 | `py.argument-list` |
| `as_pattern` | Region | degraded | 1 | `py.patterns` |
| `assert_statement` | PyAssert | exact | 2 | `py.assert` |
| `assignment` | Assign | exact | 44 | `py.assign` |
| `attribute` | Field | exact | 5 | `py.attribute` |
| `augmented_assignment` | PyAugAssign | exact | 0 | `py.assign-augmented` |
| `await` | Await | exact | 3 | `py.await` |
| `binary_operator` | Binary | exact | 27 | `py.binary` |
| `block` | Block | exact | 54 | `py.block` |
| `boolean_operator` | PyBoolOp | exact | 2 | `py.boolean-op` |
| `break_statement` | Branch | exact | 4 | `py.branch` |
| `call` | Call | exact | 31 | `py.call` |
| `case_clause` | Case | exact | 7 | `py.case` |
| `case_pattern` | Region | degraded | 1 | `py.patterns` |
| `chevron` | Invalid | out-of-scope | 0 | `py.chevron` |
| `class_definition` | ClassDecl | exact | 2 | `py.class` |
| `class_pattern` | Region | degraded | 1 | `py.patterns` |
| `comment` | Comment | exact | 1 | `core.comment` |
| `comparison_operator` | PyCompare | exact | 6 | `py.compare` |
| `complex_pattern` | Region | degraded | 1 | `py.patterns` |
| `concatenated_string` | StrLit | exact | 6 | `py.string` |
| `conditional_expression` | Ternary | exact | 2 | `py.ternary` |
| `constrained_type` | GenericParam | degraded | 0 | `py.type-constrained` |
| `continue_statement` | Branch | exact | 4 | `py.branch` |
| `decorated_definition` | Decorator | exact | 2 | `py.decorator` |
| `decorator` | Decorator | exact | 2 | `py.decorator` |
| `default_parameter` | Region | degraded | 1 | `py.parameters` |
| `delete_statement` | PyDelete | exact | 2 | `py.delete` |
| `dict_pattern` | Region | degraded | 1 | `py.patterns` |
| `dictionary` | CompositeLit | exact | 19 | `py.dict` |
| `dictionary_comprehension` | CompositeLit | exact | 19 | `py.dict-comp` |
| `dictionary_splat` | Spread | exact | 6 | `py.dict-splat` |
| `dictionary_splat_pattern` | Region | degraded | 1 | `py.patterns` |
| `dotted_name` | PyQualifiedName | exact | 0 | `py.dotted-name` |
| `elif_clause` | If | exact | 9 | `py.if` |
| `ellipsis` | Region | degraded | 1 | `py.ellipsis` |
| `else_clause` | If | exact | 9 | `py.if` |
| `escape_interpolation` | Region | degraded | 1 | `py.interpolation` |
| `escape_sequence` | PyStringFragment | exact | 10 | `py.string-part` |
| `except_clause` | Catch | exact | 4 | `py.catch` |
| `exec_statement` | Invalid | out-of-scope | 0 | `py.legacy-exec` |
| `expression_list` | PyExprList | exact | 3 | `py.expr-list` |
| `expression_statement` | ExprStmt | exact | 26 | `py.expr-stmt` |
| `false` | BoolLit | exact | 4 | `core.boolean` |
| `finally_clause` | Finally | exact | 4 | `py.finally` |
| `float` | NumLit | exact | 53 | `py.float` |
| `for_in_clause` | PyCompClause | exact | 10 | `py.for-in-clause` |
| `for_statement` | RangeFor | exact | 13 | `py.for` |
| `format_expression` | Region | degraded | 1 | `py.interpolation` |
| `format_specifier` | Region | degraded | 1 | `py.interpolation` |
| `function_definition` | FuncDecl | exact | 13 | `py.func` |
| `future_import_statement` | Import | exact | 6 | `py.import` |
| `generator_expression` | PyGenerator | exact | 2 | `py.generator` |
| `generic_type` | GenericInst | exact | 2 | `py.generic-inst` |
| `global_statement` | PyScopeDecl | exact | 3 | `py.scope-decl` |
| `identifier` | Ident | exact | 258 | `core.identifier` |
| `if_clause` | PyCompClause | exact | 10 | `py.if-clause` |
| `if_statement` | If | exact | 9 | `py.if` |
| `import_from_statement` | Import | exact | 6 | `py.import` |
| `import_prefix` | PyImportPrefix | exact | 1 | `py.import-prefix` |
| `import_statement` | Import | exact | 6 | `py.import` |
| `integer` | NumLit | exact | 53 | `py.integer` |
| `interpolation` | Region | degraded | 1 | `py.interpolation` |
| `keyword_argument` | PyKeywordArg | exact | 5 | `py.keyword-arg` |
| `keyword_pattern` | Region | degraded | 1 | `py.patterns` |
| `keyword_separator` | Region | degraded | 1 | `py.separators` |
| `lambda` | Lambda | exact | 2 | `py.lambda` |
| `lambda_parameters` | Region | degraded | 1 | `py.parameters` |
| `line_continuation` | Region | degraded | 1 | `py.line-continuation` |
| `list` | CompositeLit | exact | 19 | `py.list` |
| `list_comprehension` | CompositeLit | exact | 19 | `py.list-comp` |
| `list_pattern` | Region | degraded | 1 | `py.patterns` |
| `list_splat` | PyStarList | exact | 6 | `py.list-splat` |
| `list_splat_pattern` | Region | degraded | 1 | `py.patterns` |
| `match_statement` | Match | exact | 4 | `py.match` |
| `member_type` | PyQualifiedName | exact | 0 | `py.type-member` |
| `module` | Package | exact | 105 | `py.module` |
| `named_expression` | PyNamedExpr | degraded | 1 | `py.named-expr` |
| `none` | NilLit | exact | 8 | `core.nil` |
| `nonlocal_statement` | PyScopeDecl | exact | 3 | `py.scope-decl` |
| `not_operator` | PyNotOp | exact | 0 | `py.not` |
| `pair` | PyDictPair | exact | 5 | `py.pair` |
| `parameters` | Region | degraded | 1 | `py.parameters` |
| `parenthesized_expression` | Region | degraded | 1 | `py.paren-expr` |
| `parenthesized_list_splat` | PyStarList | exact | 6 | `py.list-splat` |
| `pass_statement` | PyPass | exact | 21 | `py.pass` |
| `pattern_list` | Region | degraded | 1 | `py.patterns` |
| `positional_separator` | Region | degraded | 1 | `py.separators` |
| `print_statement` | Invalid | out-of-scope | 0 | `py.legacy-print` |
| `raise_statement` | Throw | exact | 5 | `py.throw` |
| `relative_import` | Import | exact | 6 | `py.import` |
| `return_statement` | Return | exact | 7 | `py.return` |
| `set` | CompositeLit | exact | 19 | `py.set` |
| `set_comprehension` | CompositeLit | exact | 19 | `py.set-comp` |
| `slice` | Slice | exact | 2 | `py.slice` |
| `splat_pattern` | Region | degraded | 1 | `py.patterns` |
| `splat_type` | PySplatType | degraded | 0 | `py.splat-type` |
| `string` | StrLit | exact | 6 | `py.string` |
| `string_content` | PyStringFragment | exact | 10 | `py.string-part` |
| `string_end` | PyStringFragment | exact | 10 | `py.string-part` |
| `string_start` | PyStringFragment | exact | 10 | `py.string-part` |
| `subscript` | Index | exact | 4 | `py.subscript` |
| `true` | BoolLit | exact | 4 | `core.boolean` |
| `try_statement` | Try | exact | 4 | `py.try` |
| `tuple` | Tuple | exact | 2 | `py.tuple` |
| `tuple_pattern` | Region | degraded | 1 | `py.patterns` |
| `type` | Named | exact | 12 | `py.type` |
| `type_alias_statement` | TypeDecl | degraded | 0 | `py.type-alias` |
| `type_conversion` | PyTypeConversion | exact | 0 | `py.type-conversion` |
| `type_parameter` | GenericParam | degraded | 0 | `py.generic-param` |
| `typed_default_parameter` | Region | degraded | 1 | `py.parameters` |
| `typed_parameter` | Region | degraded | 1 | `py.parameters` |
| `unary_operator` | Unary | exact | 5 | `py.unary` |
| `union_pattern` | Region | degraded | 1 | `py.patterns` |
| `union_type` | Union | degraded | 0 | `py.type-union` |
| `while_statement` | While | exact | 1 | `py.while` |
| `wildcard_import` | Import | exact | 6 | `py.import` |
| `with_clause` | With | exact | 3 | `py.with` |
| `with_item` | PyWithItem | exact | 4 | `py.with-item` |
| `with_statement` | With | exact | 3 | `py.with` |
| `yield` | PyYield | exact | 1 | `py.yield` |

## 降级率口径说明（诚实披露）

`degrade*Pct` 按每语言**首个真实解析 unit（单文件）**计算，并对应冻结阈值；全语料使用同一 `Degradation` 公式，数值不同——两者公式相同、仅统计范围不同。

- 本语言：单文件 **0%**（≤3 达标）；全语料 **26.89%**（超阈）。
- 结论限定：达标仅限冻结的单文件口径，不代表全语料/语料级降级率达标；go 为 **NOT_MEASURED**。
- 数据来源：`docs/i18n/QA-ANTI-FRAUD-AUDIT.md` 第 7 维度；数值真实、口径不同，禁止把单文件达标外推为全语料达标。

