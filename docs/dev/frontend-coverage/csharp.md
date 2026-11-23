# C# 前端构造覆盖清单

- 上游语法事实：`https://github.com/tree-sitter/tree-sitter-c-sharp` 0.23.5
- 上游具体构造总数：**215**
- 映射完备率：**100.0%**（215/215，无未映射空洞）
- 实测覆盖：**113**（**52.6%**）
  - 精确映射覆盖：100
  - 降级映射覆盖：13
- 能力域外（显式登记并给出理由）：0
- 未映射：0
- 认可覆盖率（覆盖 + 能力域外）：**52.6%**
- 真实解析命中的 AIR 类型数：72

统计口径：覆盖率来自对该语言样例语料的**真实解析**，上游构造的 AIR 类型在解析产物中出现过才计为已覆盖。AIR 是中性中间表示，多个上游构造可能共用同一 AIR 类型，因此覆盖以 AIR 类型为粒度；每个上游构造在清单中单列一行，并标注其映射是精确映射还是降级映射。

未覆盖条目的含义：该上游构造在真实解析产物中从未产生过映射所指定的 AIR 类型。它既可能是前端尚未实现该构造，也可能是前端把它降级到了另一个 AIR 类型，二者都记为缺口。

## 未覆盖条目

| 上游构造 | 映射 AIR 类型 | 映射状态 | 原因 |
|---|---|---|---|
| `accessor_list` | Region | degraded | 访问器列表无独立节点，降级为区域节点 |
| `alias_qualified_name` | CsAliasName | exact | AIR 以独立节点建模该构造 |
| `and_pattern` | Region | degraded | 模式匹配在 AIR 无独立节点，降级为区域节点 |
| `anonymous_method_expression` | CsArrowClause | degraded | 匿名方法降级为独立箭头子句节点 |
| `anonymous_object_creation_expression` | CompositeLit | degraded | 匿名对象降级为复合字面量节点 |
| `argument` | Region | degraded | 实参列表在 AIR 无独立节点，降级为区域节点 |
| `argument_list` | Region | degraded | 实参列表在 AIR 无独立节点，降级为区域节点 |
| `array_rank_specifier` | Region | degraded | 数组秩说明符无独立节点，降级为区域节点 |
| `array_type` | Array | degraded | 数组类型降级为数组节点 |
| `arrow_expression_clause` | CsArrowClause | degraded | 表达式体成员降级为独立箭头子句节点 |
| `attribute_argument` | Region | degraded | 属性实参与目标无独立节点，降级为区域节点 |
| `attribute_argument_list` | Region | degraded | 属性实参与目标无独立节点，降级为区域节点 |
| `attribute_target_specifier` | Region | degraded | 属性实参与目标无独立节点，降级为区域节点 |
| `base_list` | Region | degraded | 基类列表无独立节点，降级为区域节点 |
| `bracketed_argument_list` | Region | degraded | 实参列表在 AIR 无独立节点，降级为区域节点 |
| `bracketed_parameter_list` | Region | degraded | 形参列表在 AIR 无独立节点，降级为区域节点 |
| `calling_convention` | Region | degraded | 调用约定无独立节点，降级为区域节点 |
| `catch_declaration` | Region | degraded | 捕获声明与过滤子句无独立节点，降级为区域节点 |
| `catch_filter_clause` | Region | degraded | 捕获声明与过滤子句无独立节点，降级为区域节点 |
| `checked_expression` | Region | degraded | checked 语句与表达式无独立节点，降级为区域节点 |
| `checked_statement` | Region | degraded | checked 语句与表达式无独立节点，降级为区域节点 |
| `collection_element` | Region | degraded | 集合元素无独立节点，降级为区域节点 |
| `collection_expression` | CompositeLit | degraded | 集合表达式降级为复合字面量节点 |
| `constant_pattern` | Region | degraded | 模式匹配在 AIR 无独立节点，降级为区域节点 |
| `constructor_constraint` | Region | degraded | 构造函数约束无独立节点，降级为区域节点 |
| `constructor_initializer` | Region | degraded | 构造函数初始化器无独立节点，降级为区域节点 |
| `declaration_expression` | CsDeclarator | degraded | 声明表达式降级为独立声明符节点 |
| `declaration_list` | Region | degraded | 声明列表无独立节点，降级为区域节点 |
| `declaration_pattern` | Region | degraded | 模式匹配在 AIR 无独立节点，降级为区域节点 |
| `default_expression` | Region | degraded | default 表达式无独立节点，降级为区域节点 |
| `delegate_declaration` | TypeDecl | degraded | 委托声明降级为类型声明节点 |
| `discard` | Region | degraded | 弃元无独立节点，降级为区域节点 |
| `element_binding_expression` | Region | degraded | 成员绑定表达式无独立节点，降级为区域节点 |
| `enum_member_declaration_list` | Region | degraded | 枚举成员列表无独立节点，降级为区域节点 |
| `explicit_interface_specifier` | Region | degraded | 显式接口限定无独立节点，降级为区域节点 |
| `expression_element` | Region | degraded | 集合元素无独立节点，降级为区域节点 |
| `fixed_statement` | Region | degraded | fixed 语句无独立节点，降级为区域节点 |
| `from_clause` | Region | degraded | LINQ 查询子句无独立节点，降级为区域节点 |
| `function_pointer_parameter` | Region | degraded | 形参列表在 AIR 无独立节点，降级为区域节点 |
| `function_pointer_type` | FuncType | degraded | 函数指针类型降级为函数类型节点 |
| `generic_name` | CsNamedType | degraded | 命名类型降级为独立命名类型节点 |
| `global_statement` | Region | degraded | global 语句无独立节点，降级为区域节点 |
| `group_clause` | Region | degraded | LINQ 查询子句无独立节点，降级为区域节点 |
| `implicit_array_creation_expression` | CompositeLit | degraded | 隐式对象创建降级为复合字面量节点 |
| `implicit_object_creation_expression` | CompositeLit | degraded | 隐式对象创建降级为复合字面量节点 |
| `implicit_parameter` | Region | degraded | 形参列表在 AIR 无独立节点，降级为区域节点 |
| `implicit_stackalloc_expression` | CompositeLit | degraded | 隐式 stackalloc 降级为复合字面量节点 |
| `implicit_type` | CsNamedType | degraded | 隐式类型降级为独立命名类型节点 |
| `interpolation` | Region | degraded | 插值片段无独立节点，降级为区域节点 |
| `interpolation_alignment_clause` | Region | degraded | 插值片段无独立节点，降级为区域节点 |
| `interpolation_brace` | Region | degraded | 插值片段无独立节点，降级为区域节点 |
| `interpolation_format_clause` | Region | degraded | 插值片段无独立节点，降级为区域节点 |
| `interpolation_quote` | Region | degraded | 插值片段无独立节点，降级为区域节点 |
| `interpolation_start` | Region | degraded | 插值片段无独立节点，降级为区域节点 |
| `join_clause` | Region | degraded | LINQ 查询子句无独立节点，降级为区域节点 |
| `join_into_clause` | Region | degraded | LINQ 查询子句无独立节点，降级为区域节点 |
| `let_clause` | Region | degraded | LINQ 查询子句无独立节点，降级为区域节点 |
| `list_pattern` | Region | degraded | 模式匹配在 AIR 无独立节点，降级为区域节点 |
| `makeref_expression` | Region | degraded | makeref 表达式无独立节点，降级为区域节点 |
| `member_binding_expression` | Region | degraded | 成员绑定表达式无独立节点，降级为区域节点 |
| `negated_pattern` | Region | degraded | 模式匹配在 AIR 无独立节点，降级为区域节点 |
| `or_pattern` | Region | degraded | 模式匹配在 AIR 无独立节点，降级为区域节点 |
| `order_by_clause` | Region | degraded | LINQ 查询子句无独立节点，降级为区域节点 |
| `parameter` | Region | degraded | 形参列表在 AIR 无独立节点，降级为区域节点 |
| `parameter_list` | Region | degraded | 形参列表在 AIR 无独立节点，降级为区域节点 |
| `parenthesized_expression` | Region | degraded | 括号表达式无独立节点，降级为区域节点 |
| `parenthesized_pattern` | Region | degraded | 模式匹配在 AIR 无独立节点，降级为区域节点 |
| `parenthesized_variable_designation` | Region | degraded | 模式匹配在 AIR 无独立节点，降级为区域节点 |
| `pointer_type` | Pointer | degraded | 指针类型降级为指针节点 |
| `positional_pattern_clause` | Region | degraded | 模式匹配在 AIR 无独立节点，降级为区域节点 |
| `predefined_type` | CsNamedType | degraded | 命名类型降级为独立命名类型节点 |
| `preproc_arg` | Region | degraded | 宏实参无独立节点，降级为区域节点 |
| `primary_constructor_base_type` | Region | degraded | 构造函数初始化器无独立节点，降级为区域节点 |
| `property_pattern_clause` | Region | degraded | 模式匹配在 AIR 无独立节点，降级为区域节点 |
| `qualified_name` | CsNamedType | degraded | 命名类型降级为独立命名类型节点 |
| `range_expression` | Region | degraded | 区间表达式无独立节点，降级为区域节点 |
| `raw_string_end` | Region | degraded | 字符串定界片段无独立节点，降级为区域节点 |
| `raw_string_start` | Region | degraded | 字符串定界片段无独立节点，降级为区域节点 |
| `recursive_pattern` | Region | degraded | 模式匹配在 AIR 无独立节点，降级为区域节点 |
| `relational_pattern` | Region | degraded | 模式匹配在 AIR 无独立节点，降级为区域节点 |
| `scoped_type` | CsNamedType | degraded | 命名类型降级为独立命名类型节点 |
| `select_clause` | Region | degraded | LINQ 查询子句无独立节点，降级为区域节点 |
| `stackalloc_expression` | CompositeLit | degraded | stackalloc 降级为复合字面量节点 |
| `string_literal_encoding` | Region | degraded | 字符串定界片段无独立节点，降级为区域节点 |
| `subpattern` | Region | degraded | 模式匹配在 AIR 无独立节点，降级为区域节点 |
| `switch_body` | Region | degraded | switch 体无独立节点，降级为区域节点 |
| `switch_expression` | Match | degraded | switch 表达式降级为 match 节点 |
| `tuple_element` | Region | degraded | 元组元素无独立节点，降级为区域节点 |
| `tuple_pattern` | Region | degraded | 模式匹配在 AIR 无独立节点，降级为区域节点 |
| `tuple_type` | CsTupleType | degraded | 元组类型降级为独立元组类型节点 |
| `type_argument_list` | GenericInst | degraded | 类型实参列表降级为泛型实例节点 |
| `type_parameter_constraint` | Region | degraded | 类型形参列表无独立节点，降级为区域节点 |
| `type_parameter_constraints_clause` | Region | degraded | 类型形参列表无独立节点，降级为区域节点 |
| `type_parameter_list` | Region | degraded | 类型形参列表无独立节点，降级为区域节点 |
| `type_pattern` | Region | degraded | 模式匹配在 AIR 无独立节点，降级为区域节点 |
| `unsafe_statement` | Region | degraded | unsafe 语句无独立节点，降级为区域节点 |
| `var_pattern` | Region | degraded | 模式匹配在 AIR 无独立节点，降级为区域节点 |
| `variable_declarator` | CsDeclarator | degraded | 变量声明符降级为独立声明符节点 |
| `when_clause` | Region | degraded | LINQ 查询子句无独立节点，降级为区域节点 |
| `where_clause` | Region | degraded | LINQ 查询子句无独立节点，降级为区域节点 |
| `with_initializer` | Region | degraded | with 初始化器无独立节点，降级为区域节点 |
| `yield_statement` | CsYield | exact | AIR 以独立节点建模该构造 |

## 能力域外条目

无。

## 全量构造映射清单

| 上游构造 | AIR 类型 | 映射状态 | 真实命中次数 | 映射规则 |
|---|---|---|---|---|
| `accessor_declaration` | CsAccessor | exact | 6 | `cs.accessor` |
| `accessor_list` | Region | degraded | 0 | `cs.accessor-list` |
| `alias_qualified_name` | CsAliasName | exact | 0 | `cs.alias-qualified` |
| `and_pattern` | Region | degraded | 0 | `cs.patterns` |
| `anonymous_method_expression` | CsArrowClause | degraded | 0 | `cs.anonymous-method` |
| `anonymous_object_creation_expression` | CompositeLit | degraded | 0 | `cs.anonymous-object` |
| `argument` | Region | degraded | 0 | `cs.arguments` |
| `argument_list` | Region | degraded | 0 | `cs.arguments` |
| `array_creation_expression` | CsObjectCreation | exact | 2 | `cs.array-creation` |
| `array_rank_specifier` | Region | degraded | 0 | `cs.array-rank` |
| `array_type` | Array | degraded | 0 | `cs.array-type` |
| `arrow_expression_clause` | CsArrowClause | degraded | 0 | `cs.arrow-clause` |
| `as_expression` | Cast | degraded | 1 | `cs.as` |
| `assignment_expression` | Assign | exact | 31 | `cs.assign` |
| `attribute` | Attribute | exact | 2 | `cs.attribute` |
| `attribute_argument` | Region | degraded | 0 | `cs.attribute-detail` |
| `attribute_argument_list` | Region | degraded | 0 | `cs.attribute-detail` |
| `attribute_list` | Attribute | exact | 2 | `cs.attribute` |
| `attribute_target_specifier` | Region | degraded | 0 | `cs.attribute-detail` |
| `await_expression` | Await | exact | 2 | `cs.await` |
| `base_list` | Region | degraded | 0 | `cs.base-list` |
| `binary_expression` | Binary | exact | 28 | `cs.binary` |
| `block` | Block | exact | 119 | `cs.block` |
| `boolean_literal` | BoolLit | exact | 3 | `cs.bool` |
| `bracketed_argument_list` | Region | degraded | 0 | `cs.arguments` |
| `bracketed_parameter_list` | Region | degraded | 0 | `cs.parameters` |
| `break_statement` | Branch | exact | 3 | `cs.branch` |
| `calling_convention` | Region | degraded | 0 | `cs.calling-convention` |
| `cast_expression` | Cast | exact | 1 | `cs.cast` |
| `catch_clause` | Catch | exact | 1 | `cs.catch` |
| `catch_declaration` | Region | degraded | 0 | `cs.catch-detail` |
| `catch_filter_clause` | Region | degraded | 0 | `cs.catch-detail` |
| `character_literal` | CharLit | exact | 1 | `cs.char` |
| `character_literal_content` | CsCharBody | exact | 1 | `cs.char-body` |
| `checked_expression` | Region | degraded | 0 | `cs.checked` |
| `checked_statement` | Region | degraded | 0 | `cs.checked` |
| `class_declaration` | ClassDecl | exact | 47 | `cs.class` |
| `collection_element` | Region | degraded | 0 | `cs.collection-element` |
| `collection_expression` | CompositeLit | degraded | 0 | `cs.collection` |
| `comment` | Comment | exact | 2 | `core.comment` |
| `compilation_unit` | Package | exact | 56 | `cs.unit` |
| `conditional_access_expression` | CsConditionalAccess | exact | 1 | `cs.conditional-access` |
| `conditional_expression` | Ternary | exact | 4 | `cs.ternary` |
| `constant_pattern` | Region | degraded | 0 | `cs.patterns` |
| `constructor_constraint` | Region | degraded | 0 | `cs.accessibility` |
| `constructor_declaration` | CsConstructor | exact | 1 | `cs.constructor` |
| `constructor_initializer` | Region | degraded | 0 | `cs.ctor-init` |
| `continue_statement` | Branch | exact | 3 | `cs.branch` |
| `conversion_operator_declaration` | CsOperatorDecl | exact | 1 | `cs.operator` |
| `declaration_expression` | CsDeclarator | degraded | 0 | `cs.declaration-expr` |
| `declaration_list` | Region | degraded | 0 | `cs.declaration-list` |
| `declaration_pattern` | Region | degraded | 0 | `cs.patterns` |
| `default_expression` | Region | degraded | 0 | `cs.default-expr` |
| `delegate_declaration` | TypeDecl | degraded | 0 | `cs.delegate` |
| `destructor_declaration` | CsDestructor | exact | 1 | `cs.destructor` |
| `discard` | Region | degraded | 0 | `cs.discard` |
| `do_statement` | DoWhile | exact | 1 | `cs.do` |
| `element_access_expression` | Index | exact | 7 | `cs.index` |
| `element_binding_expression` | Region | degraded | 0 | `cs.member-binding` |
| `empty_statement` | ExprStmt | degraded | 25 | `cs.empty-stmt` |
| `enum_declaration` | EnumDecl | exact | 1 | `cs.enum` |
| `enum_member_declaration` | EnumDecl | degraded | 1 | `cs.enum-member` |
| `enum_member_declaration_list` | Region | degraded | 0 | `cs.enum-member-list` |
| `escape_sequence` | CsStringFrag | exact | 9 | `cs.string-part` |
| `event_declaration` | CsEventDecl | exact | 1 | `cs.event` |
| `event_field_declaration` | CsEventDecl | exact | 1 | `cs.event` |
| `explicit_interface_specifier` | Region | degraded | 0 | `cs.explicit-interface` |
| `expression_element` | Region | degraded | 0 | `cs.collection-element` |
| `expression_statement` | ExprStmt | exact | 25 | `cs.expr-stmt` |
| `extern_alias_directive` | CsExternAlias | exact | 1 | `cs.extern-alias` |
| `field_declaration` | Field | exact | 5 | `cs.field` |
| `file_scoped_namespace_declaration` | ModuleDecl | exact | 2 | `cs.namespace` |
| `finally_clause` | Finally | exact | 1 | `cs.finally` |
| `fixed_statement` | Region | degraded | 0 | `cs.fixed` |
| `for_statement` | For | exact | 1 | `cs.for` |
| `foreach_statement` | RangeFor | exact | 2 | `cs.foreach` |
| `from_clause` | Region | degraded | 0 | `cs.linq-clause` |
| `function_pointer_parameter` | Region | degraded | 0 | `cs.parameters` |
| `function_pointer_type` | FuncType | degraded | 0 | `cs.fn-pointer-type` |
| `generic_name` | CsNamedType | degraded | 0 | `cs.named-type` |
| `global_attribute` | Attribute | exact | 2 | `cs.attribute` |
| `global_statement` | Region | degraded | 0 | `cs.global` |
| `goto_statement` | Branch | degraded | 3 | `cs.goto` |
| `group_clause` | Region | degraded | 0 | `cs.linq-clause` |
| `identifier` | Ident | exact | 404 | `core.identifier` |
| `if_statement` | If | exact | 3 | `cs.if` |
| `implicit_array_creation_expression` | CompositeLit | degraded | 0 | `cs.implicit-creation` |
| `implicit_object_creation_expression` | CompositeLit | degraded | 0 | `cs.implicit-creation` |
| `implicit_parameter` | Region | degraded | 0 | `cs.parameters` |
| `implicit_stackalloc_expression` | CompositeLit | degraded | 0 | `cs.implicit-stackalloc` |
| `implicit_type` | CsNamedType | degraded | 0 | `cs.implicit-type` |
| `indexer_declaration` | CsIndexer | exact | 1 | `cs.indexer` |
| `initializer_expression` | CsObjectCreation | exact | 2 | `cs.initializer` |
| `integer_literal` | NumLit | exact | 37 | `cs.number` |
| `interface_declaration` | InterfaceDecl | exact | 2 | `cs.interface` |
| `interpolated_string_expression` | CsInterpString | exact | 2 | `cs.interpolated-string` |
| `interpolation` | Region | degraded | 0 | `cs.interpolation` |
| `interpolation_alignment_clause` | Region | degraded | 0 | `cs.interpolation` |
| `interpolation_brace` | Region | degraded | 0 | `cs.interpolation` |
| `interpolation_format_clause` | Region | degraded | 0 | `cs.interpolation` |
| `interpolation_quote` | Region | degraded | 0 | `cs.interpolation` |
| `interpolation_start` | Region | degraded | 0 | `cs.interpolation` |
| `invocation_expression` | Call | exact | 92 | `cs.call` |
| `is_expression` | TypeAssert | exact | 3 | `cs.is` |
| `is_pattern_expression` | CsPattern | exact | 1 | `cs.is-pattern` |
| `join_clause` | Region | degraded | 0 | `cs.linq-clause` |
| `join_into_clause` | Region | degraded | 0 | `cs.linq-clause` |
| `labeled_statement` | Label | exact | 1 | `cs.label` |
| `lambda_expression` | Lambda | exact | 8 | `cs.lambda` |
| `let_clause` | Region | degraded | 0 | `cs.linq-clause` |
| `list_pattern` | Region | degraded | 0 | `cs.patterns` |
| `local_declaration_statement` | ConstDecl | exact | 1 | `cs.const` |
| `local_function_statement` | CsLocalFunc | exact | 1 | `cs.local-func` |
| `lock_statement` | With | degraded | 1 | `cs.lock` |
| `makeref_expression` | Region | degraded | 0 | `cs.makeref` |
| `member_access_expression` | Field | exact | 5 | `cs.member` |
| `member_binding_expression` | Region | degraded | 0 | `cs.member-binding` |
| `method_declaration` | MethodDecl | exact | 35 | `cs.method` |
| `modifier` | CsModifier | exact | 41 | `cs.modifier` |
| `namespace_declaration` | ModuleDecl | exact | 2 | `cs.namespace` |
| `negated_pattern` | Region | degraded | 0 | `cs.patterns` |
| `null_literal` | NilLit | exact | 4 | `cs.nil` |
| `nullable_type` | Optional | exact | 3 | `cs.nullable-type` |
| `object_creation_expression` | CsObjectCreation | exact | 2 | `cs.object-creation` |
| `operator_declaration` | CsOperatorDecl | exact | 1 | `cs.operator` |
| `or_pattern` | Region | degraded | 0 | `cs.patterns` |
| `order_by_clause` | Region | degraded | 0 | `cs.linq-clause` |
| `parameter` | Region | degraded | 0 | `cs.parameters` |
| `parameter_list` | Region | degraded | 0 | `cs.parameters` |
| `parenthesized_expression` | Region | degraded | 0 | `cs.paren-expr` |
| `parenthesized_pattern` | Region | degraded | 0 | `cs.patterns` |
| `parenthesized_variable_designation` | Region | degraded | 0 | `cs.patterns` |
| `pointer_type` | Pointer | degraded | 0 | `cs.pointer-type` |
| `positional_pattern_clause` | Region | degraded | 0 | `cs.patterns` |
| `postfix_unary_expression` | Unary | exact | 7 | `cs.unary` |
| `predefined_type` | CsNamedType | degraded | 0 | `cs.named-type` |
| `prefix_unary_expression` | Unary | exact | 7 | `cs.unary` |
| `preproc_arg` | Region | degraded | 0 | `cs.preproc-arg` |
| `preproc_define` | Directive | exact | 3 | `cs.preproc` |
| `preproc_elif` | Directive | exact | 3 | `cs.preproc` |
| `preproc_else` | Directive | exact | 3 | `cs.preproc` |
| `preproc_endregion` | Directive | exact | 3 | `cs.preproc` |
| `preproc_error` | Directive | exact | 3 | `cs.preproc` |
| `preproc_if` | Directive | exact | 3 | `cs.preproc` |
| `preproc_if_in_attribute_list` | Directive | exact | 3 | `cs.preproc` |
| `preproc_line` | Directive | exact | 3 | `cs.preproc` |
| `preproc_nullable` | Directive | exact | 3 | `cs.preproc` |
| `preproc_pragma` | Directive | exact | 3 | `cs.preproc` |
| `preproc_region` | Directive | exact | 3 | `cs.preproc` |
| `preproc_undef` | Directive | exact | 3 | `cs.preproc` |
| `preproc_warning` | Directive | exact | 3 | `cs.preproc` |
| `primary_constructor_base_type` | Region | degraded | 0 | `cs.ctor-init` |
| `property_declaration` | CsProperty | exact | 3 | `cs.property` |
| `property_pattern_clause` | Region | degraded | 0 | `cs.patterns` |
| `qualified_name` | CsNamedType | degraded | 0 | `cs.named-type` |
| `query_expression` | CsQuery | exact | 4 | `cs.query` |
| `range_expression` | Region | degraded | 0 | `cs.range-expr` |
| `raw_string_content` | CsStringFrag | exact | 9 | `cs.string-part` |
| `raw_string_end` | Region | degraded | 0 | `cs.string-bound` |
| `raw_string_literal` | StrLit | exact | 4 | `cs.string` |
| `raw_string_start` | Region | degraded | 0 | `cs.string-bound` |
| `real_literal` | NumLit | exact | 37 | `cs.number` |
| `record_declaration` | CsRecord | exact | 2 | `cs.record` |
| `recursive_pattern` | Region | degraded | 0 | `cs.patterns` |
| `ref_expression` | Ref | exact | 2 | `cs.ref` |
| `ref_type` | Ref | degraded | 2 | `cs.ref-family` |
| `reftype_expression` | Ref | degraded | 2 | `cs.ref-family` |
| `refvalue_expression` | Ref | degraded | 2 | `cs.ref-family` |
| `relational_pattern` | Region | degraded | 0 | `cs.patterns` |
| `return_statement` | Return | exact | 2 | `cs.return` |
| `scoped_type` | CsNamedType | degraded | 0 | `cs.named-type` |
| `select_clause` | Region | degraded | 0 | `cs.linq-clause` |
| `shebang_directive` | Directive | exact | 3 | `cs.shebang` |
| `sizeof_expression` | CsSizeof | degraded | 1 | `cs.sizeof` |
| `spread_element` | Spread | exact | 1 | `cs.spread` |
| `stackalloc_expression` | CompositeLit | degraded | 0 | `cs.stackalloc` |
| `string_content` | CsStringFrag | exact | 9 | `cs.string-part` |
| `string_literal` | StrLit | exact | 4 | `cs.string` |
| `string_literal_content` | CsStringFrag | exact | 9 | `cs.string-part` |
| `string_literal_encoding` | Region | degraded | 0 | `cs.string-bound` |
| `struct_declaration` | StructDecl | exact | 1 | `cs.struct` |
| `subpattern` | Region | degraded | 0 | `cs.patterns` |
| `switch_body` | Region | degraded | 0 | `cs.switch-body` |
| `switch_expression` | Match | degraded | 0 | `cs.switch-expr` |
| `switch_expression_arm` | Case | degraded | 2 | `cs.switch-arm` |
| `switch_section` | Case | degraded | 2 | `cs.switch-section` |
| `switch_statement` | Switch | exact | 1 | `cs.switch-stmt` |
| `throw_expression` | Throw | exact | 1 | `cs.throw` |
| `throw_statement` | Throw | exact | 1 | `cs.throw` |
| `try_statement` | Try | exact | 1 | `cs.try` |
| `tuple_element` | Region | degraded | 0 | `cs.tuple-element` |
| `tuple_expression` | Tuple | exact | 1 | `cs.tuple` |
| `tuple_pattern` | Region | degraded | 0 | `cs.patterns` |
| `tuple_type` | CsTupleType | degraded | 0 | `cs.tuple-type` |
| `type_argument_list` | GenericInst | degraded | 0 | `cs.type-args` |
| `type_parameter` | GenericParam | degraded | 2 | `cs.generic-param` |
| `type_parameter_constraint` | Region | degraded | 0 | `cs.type-params` |
| `type_parameter_constraints_clause` | Region | degraded | 0 | `cs.type-params` |
| `type_parameter_list` | Region | degraded | 0 | `cs.type-params` |
| `type_pattern` | Region | degraded | 0 | `cs.patterns` |
| `typeof_expression` | CsTypeof | exact | 1 | `cs.typeof` |
| `unary_expression` | Unary | exact | 7 | `cs.unary` |
| `unsafe_statement` | Region | degraded | 0 | `cs.unsafe` |
| `using_directive` | UsingDecl | exact | 1 | `cs.using-directive` |
| `using_statement` | With | exact | 1 | `cs.using-stmt` |
| `var_pattern` | Region | degraded | 0 | `cs.patterns` |
| `variable_declaration` | VarDecl | exact | 17 | `cs.var` |
| `variable_declarator` | CsDeclarator | degraded | 0 | `cs.declarator` |
| `verbatim_string_literal` | StrLit | exact | 4 | `cs.string` |
| `when_clause` | Region | degraded | 0 | `cs.linq-clause` |
| `where_clause` | Region | degraded | 0 | `cs.linq-clause` |
| `while_statement` | While | exact | 2 | `cs.while` |
| `with_expression` | With | degraded | 1 | `cs.with-expr` |
| `with_initializer` | Region | degraded | 0 | `cs.with-init` |
| `yield_statement` | CsYield | exact | 0 | `cs.yield` |

## 降级率口径说明（诚实披露）

`degrade*Pct` 按每语言**首个真实解析 unit（单文件）**计算，并对应冻结阈值；全语料使用同一 `Degradation` 公式，数值不同——两者公式相同、仅统计范围不同。

- 本语言：单文件 **3.26%**（≤5 达标）；全语料 **6.98%**（超阈）。
- 结论限定：达标仅限冻结的单文件口径，不代表全语料/语料级降级率达标；go 为 **NOT_MEASURED**。
- 数据来源：`docs/i18n/QA-ANTI-FRAUD-AUDIT.md` 第 7 维度；数值真实、口径不同，禁止把单文件达标外推为全语料达标。

