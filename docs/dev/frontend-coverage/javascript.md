# JavaScript/TypeScript 前端构造覆盖清单

- 上游语法事实：`https://github.com/tree-sitter/tree-sitter-javascript` 0.25.0
- 上游具体构造总数：**114**
- 映射完备率：**100.0%**（114/114，无未映射空洞）
- 实测覆盖：**109**（**95.6%**）
  - 精确映射覆盖：54
  - 降级映射覆盖：55
- 能力域外（显式登记并给出理由）：0
- 未映射：0
- 认可覆盖率（覆盖 + 能力域外）：**95.6%**
- 真实解析命中的 AIR 类型数：43

统计口径：覆盖率来自对该语言样例语料的**真实解析**，上游构造的 AIR 类型在解析产物中出现过才计为已覆盖。AIR 是中性中间表示，多个上游构造可能共用同一 AIR 类型，因此覆盖以 AIR 类型为粒度；每个上游构造在清单中单列一行，并标注其映射是精确映射还是降级映射。

未覆盖条目的含义：该上游构造在真实解析产物中从未产生过映射所指定的 AIR 类型。它既可能是前端尚未实现该构造，也可能是前端把它降级到了另一个 AIR 类型，二者都记为缺口。

## 未覆盖条目

| 上游构造 | 映射 AIR 类型 | 映射状态 | 原因 |
|---|---|---|---|
| `class_static_block` | Block | degraded | 静态初始化块降级为块节点 |
| `labeled_statement` | Label | degraded | 全语料零产出，回退 degraded 避免 exact 分母虚增 |
| `sequence_expression` | Tuple | degraded | 逗号表达式降级为元组节点 |
| `statement_block` | Block | exact | AIR 以独立节点建模该构造 |
| `update_expression` | IncDec | degraded | 全语料零产出，回退 degraded 避免 exact 分母虚增 |

## 能力域外条目

无。

## 全量构造映射清单

| 上游构造 | AIR 类型 | 映射状态 | 真实命中次数 | 映射规则 |
|---|---|---|---|---|
| `arguments` | Region | degraded | 401 | `js.arguments` |
| `array` | CompositeLit | exact | 79 | `js.array` |
| `array_pattern` | Region | degraded | 401 | `js.array-pattern` |
| `arrow_function` | Arrow | exact | 3 | `js.arrow` |
| `assignment_expression` | Assign | exact | 97 | `js.assign` |
| `assignment_pattern` | Assign | degraded | 97 | `js.assign-pattern` |
| `augmented_assignment_expression` | Assign | degraded | 97 | `js.assign-augmented` |
| `await_expression` | Await | exact | 8 | `js.await` |
| `binary_expression` | Binary | exact | 59 | `js.binary` |
| `break_statement` | Branch | exact | 10 | `js.branch` |
| `call_expression` | Call | exact | 45 | `js.call` |
| `catch_clause` | Catch | exact | 6 | `js.catch` |
| `class` | ClassDecl | degraded | 12 | `js.class-expr` |
| `class_body` | Region | degraded | 401 | `js.class-body` |
| `class_declaration` | ClassDecl | exact | 12 | `js.class` |
| `class_heritage` | Region | degraded | 401 | `js.class-heritage` |
| `class_static_block` | Block | degraded | 0 | `js.static-block` |
| `comment` | Comment | exact | 2 | `core.comment` |
| `computed_property_name` | Region | degraded | 401 | `js.computed-name` |
| `continue_statement` | Branch | exact | 10 | `js.branch` |
| `debugger_statement` | Region | degraded | 401 | `js.debugger` |
| `decorator` | Decorator | exact | 6 | `js.decorator` |
| `do_statement` | DoWhile | exact | 6 | `js.do` |
| `else_clause` | If | exact | 8 | `js.else` |
| `empty_statement` | Region | degraded | 401 | `js.empty-stmt` |
| `escape_sequence` | StrLit | degraded | 22 | `js.string-part` |
| `export_clause` | ModuleDecl | degraded | 22 | `js.export` |
| `export_specifier` | ModuleDecl | degraded | 22 | `js.export` |
| `export_statement` | ModuleDecl | degraded | 22 | `js.export` |
| `expression_statement` | ExprStmt | exact | 17 | `js.expr-stmt` |
| `false` | BoolLit | exact | 2 | `core.boolean` |
| `field_definition` | Field | degraded | 9 | `js.field-def` |
| `finally_clause` | Finally | exact | 6 | `js.finally` |
| `for_in_statement` | RangeFor | exact | 4 | `js.for-in` |
| `for_statement` | For | exact | 7 | `js.for` |
| `formal_parameters` | Region | degraded | 401 | `js.formal-parameters` |
| `function_declaration` | FuncDecl | exact | 20 | `js.func` |
| `function_expression` | FuncLit | exact | 4 | `js.func-lit` |
| `generator_function` | FuncLit | degraded | 4 | `js.generator-func` |
| `generator_function_declaration` | FuncDecl | degraded | 20 | `js.generator-func-decl` |
| `hash_bang_line` | Directive | exact | 4 | `js.hash-bang` |
| `html_character_reference` | StrLit | degraded | 22 | `js.string-part` |
| `html_comment` | Comment | exact | 2 | `js.html-comment` |
| `identifier` | Ident | exact | 340 | `core.identifier` |
| `if_statement` | If | exact | 8 | `js.if` |
| `import` | Import | exact | 10 | `js.import` |
| `import_attribute` | Import | exact | 10 | `js.import` |
| `import_clause` | Import | exact | 10 | `js.import` |
| `import_specifier` | Import | exact | 10 | `js.import` |
| `import_statement` | Import | exact | 10 | `js.import` |
| `jsx_attribute` | CompositeLit | degraded | 79 | `js.jsx` |
| `jsx_closing_element` | CompositeLit | degraded | 79 | `js.jsx` |
| `jsx_element` | CompositeLit | degraded | 79 | `js.jsx` |
| `jsx_expression` | CompositeLit | degraded | 79 | `js.jsx` |
| `jsx_namespace_name` | CompositeLit | degraded | 79 | `js.jsx` |
| `jsx_opening_element` | CompositeLit | degraded | 79 | `js.jsx` |
| `jsx_self_closing_element` | CompositeLit | degraded | 79 | `js.jsx` |
| `jsx_text` | CompositeLit | degraded | 79 | `js.jsx` |
| `labeled_statement` | Label | degraded | 0 | `js.label` |
| `lexical_declaration` | VarDecl | degraded | 12 | `js.lexical-decl` |
| `member_expression` | Field | exact | 9 | `js.member` |
| `meta_property` | Region | degraded | 401 | `js.meta-property` |
| `method_definition` | MethodDecl | exact | 7 | `js.method` |
| `named_imports` | Import | exact | 10 | `js.import` |
| `namespace_export` | ModuleDecl | degraded | 22 | `js.export` |
| `namespace_import` | Import | exact | 10 | `js.import` |
| `new_expression` | Call | degraded | 45 | `js.new` |
| `null` | NilLit | exact | 5 | `core.nil` |
| `number` | NumLit | exact | 80 | `js.number` |
| `object` | CompositeLit | exact | 79 | `js.object` |
| `object_assignment_pattern` | Assign | degraded | 97 | `js.object-assign-pattern` |
| `object_pattern` | Region | degraded | 401 | `js.object-pattern` |
| `optional_chain` | Field | degraded | 9 | `js.optional-chain` |
| `pair` | Assign | degraded | 97 | `js.pair` |
| `pair_pattern` | Assign | degraded | 97 | `js.pair-pattern` |
| `parenthesized_expression` | Region | degraded | 401 | `js.paren-expr` |
| `private_property_identifier` | Ident | degraded | 340 | `js.ident-family` |
| `program` | Package | exact | 37 | `js.program` |
| `property_identifier` | Ident | degraded | 340 | `js.ident-family` |
| `regex` | StrLit | degraded | 22 | `js.regex` |
| `regex_flags` | StrLit | degraded | 22 | `js.regex` |
| `regex_pattern` | StrLit | degraded | 22 | `js.regex` |
| `rest_pattern` | Spread | degraded | 10 | `js.rest-pattern` |
| `return_statement` | Return | exact | 8 | `js.return` |
| `sequence_expression` | Tuple | degraded | 0 | `js.sequence` |
| `shorthand_property_identifier` | Ident | degraded | 340 | `js.ident-family` |
| `shorthand_property_identifier_pattern` | Ident | degraded | 340 | `js.ident-family` |
| `spread_element` | Spread | exact | 10 | `js.spread` |
| `statement_block` | Block | exact | 0 | `js.block` |
| `statement_identifier` | Ident | degraded | 340 | `js.ident-family` |
| `string` | StrLit | exact | 22 | `js.string` |
| `string_fragment` | StrLit | degraded | 22 | `js.string-part` |
| `subscript_expression` | Index | exact | 9 | `js.index` |
| `super` | Ident | degraded | 340 | `js.this-super` |
| `switch_body` | Region | degraded | 401 | `js.switch-body` |
| `switch_case` | Case | exact | 7 | `js.case` |
| `switch_default` | Case | degraded | 7 | `js.switch-default` |
| `switch_statement` | Switch | exact | 6 | `js.switch` |
| `template_string` | StrLit | exact | 22 | `js.template` |
| `template_substitution` | Region | degraded | 401 | `js.template-substitution` |
| `ternary_expression` | Ternary | exact | 6 | `js.ternary` |
| `this` | Ident | degraded | 340 | `js.this-super` |
| `throw_statement` | Throw | exact | 2 | `js.throw` |
| `true` | BoolLit | exact | 2 | `core.boolean` |
| `try_statement` | Try | exact | 6 | `js.try` |
| `unary_expression` | Unary | exact | 14 | `js.unary` |
| `undefined` | NilLit | exact | 5 | `core.nil` |
| `update_expression` | IncDec | degraded | 0 | `js.incdec` |
| `using_declaration` | VarDecl | degraded | 12 | `js.using-decl` |
| `variable_declaration` | VarDecl | exact | 12 | `js.var` |
| `variable_declarator` | Region | degraded | 401 | `js.declarator` |
| `while_statement` | While | exact | 10 | `js.while` |
| `with_statement` | With | exact | 2 | `js.with` |
| `yield_expression` | Unary | degraded | 14 | `js.yield` |

## 降级率口径说明（诚实披露）

`degrade*Pct` 按每语言**首个真实解析 unit（单文件）**计算，并对应冻结阈值；全语料使用同一 `Degradation` 公式，数值不同——两者公式相同、仅统计范围不同。

- 本语言：无冻结降级率阈值；全语料 **49.12%**（超阈）。
- 结论限定：达标仅限冻结的单文件口径，不代表全语料/语料级降级率达标；go 为 **NOT_MEASURED**。
- 数据来源：`docs/i18n/QA-ANTI-FRAUD-AUDIT.md` 第 7 维度；数值真实、口径不同，禁止把单文件达标外推为全语料达标。

