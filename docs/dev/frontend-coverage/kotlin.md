# Kotlin 前端构造覆盖清单

- 上游语法事实：`https://github.com/tree-sitter/tree-sitter-kotlin` 0.3.8
- 上游具体构造总数：**134**
- 映射完备率：**100.0%**（134/134，无未映射空洞）
- 实测覆盖：**63**（**47.0%**）
  - 精确映射覆盖：59
  - 降级映射覆盖：4
- 能力域外（显式登记并给出理由）：0
- 未映射：0
- 认可覆盖率（覆盖 + 能力域外）：**47.0%**
- 真实解析命中的 AIR 类型数：44

统计口径：覆盖率来自对该语言样例语料的**真实解析**，上游构造的 AIR 类型在解析产物中出现过才计为已覆盖。AIR 是中性中间表示，多个上游构造可能共用同一 AIR 类型，因此覆盖以 AIR 类型为粒度；每个上游构造在清单中单列一行，并标注其映射是精确映射还是降级映射。

未覆盖条目的含义：该上游构造在真实解析产物中从未产生过映射所指定的 AIR 类型。它既可能是前端尚未实现该构造，也可能是前端把它降级到了另一个 AIR 类型，二者都记为缺口。

## 未覆盖条目

| 上游构造 | 映射 AIR 类型 | 映射状态 | 原因 |
|---|---|---|---|
| `annotated_lambda` | KtAnnotatedLambda | degraded | 带注解 lambda 降级为独立注解 lambda 节点 |
| `anonymous_initializer` | CompositeLit | degraded | 匿名初始化块降级为复合字面量节点 |
| `as_expression` | Cast | degraded | as 表达式降级为类型断言节点 |
| `binding_pattern_kind` | Region | degraded | 绑定模式无独立节点，降级为区域节点 |
| `call_suffix` | Region | degraded | 调用与导航后缀无独立节点，降级为区域节点 |
| `check_expression` | Region | degraded | check 表达式无独立节点，降级为区域节点 |
| `class_body` | Region | degraded | 类体无独立节点，降级为区域节点 |
| `class_modifier` | Attribute | degraded | 声明修饰符降级为属性节点 |
| `class_parameter` | Region | degraded | 主构造函数形参无独立节点，降级为区域节点 |
| `collection_literal` | CompositeLit | degraded | 全语料零产出，回退 degraded 避免 exact 分母虚增 |
| `constructor_delegation_call` | KtCtorDelegation | exact | AIR 以独立节点建模该构造 |
| `constructor_invocation` | KtInfixCtor | degraded | 构造调用降级为独立中缀/构造调用节点 |
| `delegation_specifier` | Region | degraded | 委派规格无独立节点，降级为区域节点 |
| `directly_assignable_expression` | Region | degraded | 可赋值表达式标记无独立节点，降级为区域节点 |
| `elvis_expression` | Ternary | degraded | 全语料零产出，回退 degraded 避免 exact 分母虚增 |
| `enum_class_body` | Region | degraded | 枚举体无独立节点，降级为区域节点 |
| `explicit_delegation` | Region | degraded | 委派规格无独立节点，降级为区域节点 |
| `function_modifier` | Attribute | degraded | 声明修饰符降级为属性节点 |
| `function_type` | FuncType | degraded | 全语料零产出，回退 degraded 避免 exact 分母虚增 |
| `function_type_parameters` | Region | degraded | 函数类型形参列表无独立节点，降级为区域节点 |
| `function_value_parameters` | Region | degraded | 函数类型形参列表无独立节点，降级为区域节点 |
| `indexing_suffix` | Region | degraded | 调用与导航后缀无独立节点，降级为区域节点 |
| `infix_expression` | KtInfixCtor | degraded | 中缀调用降级为独立中缀/构造调用节点 |
| `inheritance_modifier` | Attribute | degraded | 声明修饰符降级为属性节点 |
| `interpolated_identifier` | Region | degraded | 模板标识符片段无独立节点，降级为区域节点 |
| `label` | Label | degraded | 全语料零产出，回退 degraded 避免 exact 分母虚增 |
| `lambda_parameters` | Region | degraded | lambda 形参无独立节点，降级为区域节点 |
| `member_modifier` | Attribute | degraded | 声明修饰符降级为属性节点 |
| `modifiers` | Attribute | degraded | 声明修饰符降级为属性节点 |
| `multi_variable_declaration` | KtMultiVar | degraded | 多变量声明降级为独立多变量声明节点 |
| `navigation_suffix` | Region | degraded | 调用与导航后缀无独立节点，降级为区域节点 |
| `not_nullable_type` | Named | degraded | 非空类型降级为命名类型节点 |
| `object_literal` | CompositeLit | degraded | 全语料零产出，回退 degraded 避免 exact 分母虚增 |
| `parameter` | Region | degraded | 主构造函数形参无独立节点，降级为区域节点 |
| `parameter_modifier` | Attribute | degraded | 声明修饰符降级为属性节点 |
| `parameter_modifiers` | Attribute | degraded | 声明修饰符降级为属性节点 |
| `parameter_with_optional_type` | Region | degraded | 主构造函数形参无独立节点，降级为区域节点 |
| `parenthesized_expression` | Region | degraded | 括号表达式无独立节点，降级为区域节点 |
| `parenthesized_type` | Named | degraded | 带括号类型降级为命名类型节点 |
| `parenthesized_user_type` | Named | degraded | 带括号类型降级为命名类型节点 |
| `platform_modifier` | Attribute | degraded | 声明修饰符降级为属性节点 |
| `primary_constructor` | Region | degraded | 主构造函数无独立节点，降级为区域节点 |
| `property_delegate` | Region | degraded | 属性委托无独立节点，降级为区域节点 |
| `property_modifier` | Attribute | degraded | 声明修饰符降级为属性节点 |
| `range_expression` | KtRange | degraded | 区间表达式降级为独立区间节点 |
| `range_test` | Region | degraded | 区间判定无独立节点，降级为区域节点 |
| `reification_modifier` | Attribute | degraded | 声明修饰符降级为属性节点 |
| `shebang_line` | Directive | degraded | 全语料零产出，回退 degraded 避免 exact 分母虚增 |
| `simple_identifier` | KtSimpleIdent | degraded | 简单标识符降级为独立简单标识符节点 |
| `spread_expression` | Spread | degraded | 全语料零产出，回退 degraded 避免 exact 分母虚增 |
| `statements` | Region | degraded | 语句列表无独立节点，降级为区域节点 |
| `super_expression` | KtSuper | exact | AIR 以独立节点建模该构造 |
| `type_arguments` | GenericInst | degraded | 类型实参列表降级为泛型实例节点 |
| `type_constraint` | Region | degraded | 类型形参列表无独立节点，降级为区域节点 |
| `type_constraints` | Region | degraded | 类型形参列表无独立节点，降级为区域节点 |
| `type_identifier` | Named | degraded | 全语料零产出，回退 degraded 避免 exact 分母虚增 |
| `type_modifiers` | Attribute | degraded | 类型修饰符降级为属性节点 |
| `type_parameter` | GenericParam | degraded | 全语料零产出，回退 degraded 避免 exact 分母虚增 |
| `type_parameter_modifiers` | Region | degraded | 类型形参列表无独立节点，降级为区域节点 |
| `type_parameters` | Region | degraded | 类型形参列表无独立节点，降级为区域节点 |
| `type_projection` | GenericInst | degraded | 类型实参列表降级为泛型实例节点 |
| `type_projection_modifiers` | Region | degraded | 类型形参列表无独立节点，降级为区域节点 |
| `use_site_target` | Region | degraded | use-site 目标无独立节点，降级为区域节点 |
| `user_type` | Named | degraded | 全语料零产出，回退 degraded 避免 exact 分母虚增 |
| `value_argument` | Region | degraded | 实参列表在 AIR 无独立节点，降级为区域节点 |
| `value_arguments` | Region | degraded | 实参列表在 AIR 无独立节点，降级为区域节点 |
| `variance_modifier` | Attribute | degraded | 声明修饰符降级为属性节点 |
| `visibility_modifier` | Attribute | degraded | 声明修饰符降级为属性节点 |
| `when_condition` | Region | degraded | when 条件无独立节点，降级为区域节点 |
| `when_entry` | KtWhenEntry | exact | AIR 以独立节点建模该构造 |
| `when_expression` | Match | exact | AIR 以独立节点建模该构造 |

## 能力域外条目

无。

## 全量构造映射清单

| 上游构造 | AIR 类型 | 映射状态 | 真实命中次数 | 映射规则 |
|---|---|---|---|---|
| `additive_expression` | Binary | exact | 25 | `kt.additive` |
| `annotated_lambda` | KtAnnotatedLambda | degraded | 0 | `kt.annotated-lambda` |
| `annotation` | AnnotationDecl | exact | 1 | `kt.annotation` |
| `anonymous_function` | Lambda | exact | 10 | `kt.anonymous-func` |
| `anonymous_initializer` | CompositeLit | degraded | 0 | `kt.anonymous-init` |
| `as_expression` | Cast | degraded | 0 | `kt.as` |
| `assignment` | Assign | exact | 51 | `kt.assign` |
| `bin_literal` | NumLit | exact | 34 | `kt.number` |
| `binding_pattern_kind` | Region | degraded | 0 | `kt.binding-pattern` |
| `boolean_literal` | BoolLit | exact | 4 | `kt.bool` |
| `call_expression` | Call | exact | 60 | `kt.call` |
| `call_suffix` | Region | degraded | 0 | `kt.suffix` |
| `callable_reference` | KtCallableRef | exact | 1 | `kt.callable-ref` |
| `catch_block` | Catch | exact | 1 | `kt.catch` |
| `character_escape_seq` | CharLit | degraded | 2 | `kt.char-escape` |
| `character_literal` | CharLit | exact | 2 | `kt.char` |
| `check_expression` | Region | degraded | 0 | `kt.check` |
| `class_body` | Region | degraded | 0 | `kt.class-body` |
| `class_declaration` | ClassDecl | exact | 14 | `kt.class` |
| `class_modifier` | Attribute | degraded | 0 | `kt.modifiers` |
| `class_parameter` | Region | degraded | 0 | `kt.class-parameter` |
| `collection_literal` | CompositeLit | degraded | 0 | `kt.collection` |
| `companion_object` | KtCompanion | exact | 2 | `kt.companion` |
| `comparison_expression` | Binary | exact | 25 | `kt.comparison` |
| `conjunction_expression` | Binary | exact | 25 | `kt.conjunction` |
| `constructor_delegation_call` | KtCtorDelegation | exact | 0 | `kt.ctor-delegation` |
| `constructor_invocation` | KtInfixCtor | degraded | 0 | `kt.ctor-call` |
| `control_structure_body` | KtControlBody | exact | 4 | `kt.control-body` |
| `delegation_specifier` | Region | degraded | 0 | `kt.delegation` |
| `directly_assignable_expression` | Region | degraded | 0 | `kt.assignable` |
| `disjunction_expression` | Binary | exact | 25 | `kt.disjunction` |
| `do_while_statement` | DoWhile | exact | 1 | `kt.do` |
| `elvis_expression` | Ternary | degraded | 0 | `kt.ternary` |
| `enum_class_body` | Region | degraded | 0 | `kt.enum-body` |
| `enum_entry` | EnumDecl | degraded | 2 | `kt.enum-entry` |
| `equality_expression` | Binary | exact | 25 | `kt.equality` |
| `explicit_delegation` | Region | degraded | 0 | `kt.delegation` |
| `file_annotation` | AnnotationDecl | exact | 1 | `kt.annotation` |
| `finally_block` | Finally | exact | 1 | `kt.finally` |
| `for_statement` | KtForIn | exact | 2 | `kt.for` |
| `function_body` | Block | exact | 36 | `kt.block` |
| `function_declaration` | FuncDecl | exact | 15 | `kt.func` |
| `function_modifier` | Attribute | degraded | 0 | `kt.modifiers` |
| `function_type` | FuncType | degraded | 0 | `kt.fn-type` |
| `function_type_parameters` | Region | degraded | 0 | `kt.fn-type-params` |
| `function_value_parameters` | Region | degraded | 0 | `kt.fn-type-params` |
| `getter` | KtAccessor | exact | 1 | `kt.accessor` |
| `hex_literal` | NumLit | exact | 34 | `kt.number` |
| `identifier` | Ident | exact | 361 | `core.identifier` |
| `if_expression` | If | exact | 1 | `kt.if` |
| `import_alias` | KtImportAlias | exact | 1 | `kt.import-alias` |
| `import_header` | Import | exact | 3 | `kt.import` |
| `import_list` | KtImportList | exact | 2 | `kt.import-list` |
| `indexing_expression` | Index | exact | 1 | `kt.index` |
| `indexing_suffix` | Region | degraded | 0 | `kt.suffix` |
| `infix_expression` | KtInfixCtor | degraded | 0 | `kt.infix-call` |
| `inheritance_modifier` | Attribute | degraded | 0 | `kt.modifiers` |
| `integer_literal` | NumLit | exact | 34 | `kt.number` |
| `interpolated_expression` | KtInterp | exact | 6 | `kt.interpolated` |
| `interpolated_identifier` | Region | degraded | 0 | `kt.interpolated-ident` |
| `jump_expression` | Branch | degraded | 2 | `kt.jump` |
| `label` | Label | degraded | 0 | `kt.label` |
| `lambda_literal` | Lambda | exact | 10 | `kt.lambda` |
| `lambda_parameters` | Region | degraded | 0 | `kt.lambda-params` |
| `line_comment` | Comment | exact | 4 | `kt.comment` |
| `long_literal` | NumLit | exact | 34 | `kt.number` |
| `member_modifier` | Attribute | degraded | 0 | `kt.modifiers` |
| `modifiers` | Attribute | degraded | 0 | `kt.modifiers` |
| `multi_variable_declaration` | KtMultiVar | degraded | 0 | `kt.multi-var` |
| `multiline_comment` | Comment | exact | 4 | `kt.comment` |
| `multiplicative_expression` | Binary | exact | 25 | `kt.multiplicative` |
| `navigation_expression` | Selector | exact | 1 | `kt.navigation` |
| `navigation_suffix` | Region | degraded | 0 | `kt.suffix` |
| `not_nullable_type` | Named | degraded | 0 | `kt.not-null-type` |
| `nullable_type` | Optional | exact | 3 | `kt.nullable-type` |
| `object_declaration` | ClassDecl | exact | 14 | `kt.object` |
| `object_literal` | CompositeLit | degraded | 0 | `kt.object-literal` |
| `package_header` | Package | exact | 55 | `kt.package` |
| `parameter` | Region | degraded | 0 | `kt.class-parameter` |
| `parameter_modifier` | Attribute | degraded | 0 | `kt.modifiers` |
| `parameter_modifiers` | Attribute | degraded | 0 | `kt.modifiers` |
| `parameter_with_optional_type` | Region | degraded | 0 | `kt.class-parameter` |
| `parenthesized_expression` | Region | degraded | 0 | `kt.paren-expr` |
| `parenthesized_type` | Named | degraded | 0 | `kt.paren-type` |
| `parenthesized_user_type` | Named | degraded | 0 | `kt.paren-type` |
| `platform_modifier` | Attribute | degraded | 0 | `kt.modifiers` |
| `postfix_expression` | Unary | degraded | 3 | `kt.postfix` |
| `prefix_expression` | Unary | exact | 3 | `kt.prefix` |
| `primary_constructor` | Region | degraded | 0 | `kt.primary-ctor` |
| `property_declaration` | KtProperty | exact | 41 | `kt.property` |
| `property_delegate` | Region | degraded | 0 | `kt.property-delegate` |
| `property_modifier` | Attribute | degraded | 0 | `kt.modifiers` |
| `range_expression` | KtRange | degraded | 0 | `kt.range` |
| `range_test` | Region | degraded | 0 | `kt.range-test` |
| `real_literal` | NumLit | exact | 34 | `kt.number` |
| `reification_modifier` | Attribute | degraded | 0 | `kt.modifiers` |
| `secondary_constructor` | KtSecondaryCtor | exact | 1 | `kt.secondary-ctor` |
| `setter` | KtAccessor | exact | 1 | `kt.accessor` |
| `shebang_line` | Directive | degraded | 0 | `kt.shebang` |
| `simple_identifier` | KtSimpleIdent | degraded | 0 | `kt.simple-identifier` |
| `source_file` | Package | exact | 55 | `kt.file` |
| `spread_expression` | Spread | degraded | 0 | `kt.spread` |
| `statements` | Region | degraded | 0 | `kt.statements` |
| `string_content` | KtStringFrag | exact | 15 | `kt.string-part` |
| `string_literal` | StrLit | exact | 7 | `kt.string` |
| `super_expression` | KtSuper | exact | 0 | `kt.super-expr` |
| `this_expression` | KtThis | exact | 1 | `kt.this-expr` |
| `try_expression` | Try | exact | 1 | `kt.try` |
| `type_alias` | TypeDecl | exact | 1 | `kt.type-alias` |
| `type_arguments` | GenericInst | degraded | 0 | `kt.type-args` |
| `type_constraint` | Region | degraded | 0 | `kt.type-params` |
| `type_constraints` | Region | degraded | 0 | `kt.type-params` |
| `type_identifier` | Named | degraded | 0 | `kt.named-type` |
| `type_modifiers` | Attribute | degraded | 0 | `kt.type-modifiers` |
| `type_parameter` | GenericParam | degraded | 0 | `kt.generic-param` |
| `type_parameter_modifiers` | Region | degraded | 0 | `kt.type-params` |
| `type_parameters` | Region | degraded | 0 | `kt.type-params` |
| `type_projection` | GenericInst | degraded | 0 | `kt.type-args` |
| `type_projection_modifiers` | Region | degraded | 0 | `kt.type-params` |
| `type_test` | TypeAssert | exact | 4 | `kt.type-test` |
| `unsigned_literal` | NumLit | exact | 34 | `kt.number` |
| `use_site_target` | Region | degraded | 0 | `kt.use-site` |
| `user_type` | Named | degraded | 0 | `kt.named-type` |
| `value_argument` | Region | degraded | 0 | `kt.value-args` |
| `value_arguments` | Region | degraded | 0 | `kt.value-args` |
| `variable_declaration` | VarDecl | exact | 45 | `kt.var` |
| `variance_modifier` | Attribute | degraded | 0 | `kt.modifiers` |
| `visibility_modifier` | Attribute | degraded | 0 | `kt.modifiers` |
| `when_condition` | Region | degraded | 0 | `kt.when-condition` |
| `when_entry` | KtWhenEntry | exact | 0 | `kt.when-entry` |
| `when_expression` | Match | exact | 0 | `kt.when` |
| `when_subject` | MatchExpr | exact | 2 | `kt.when-subject` |
| `while_statement` | While | exact | 2 | `kt.while` |
| `wildcard_import` | Import | exact | 3 | `kt.import` |

## 降级率口径说明（诚实披露）

`degrade*Pct` 按每语言**首个真实解析 unit（单文件）**计算，并对应冻结阈值；全语料使用同一 `Degradation` 公式，数值不同——两者公式相同、仅统计范围不同。

- 本语言：单文件 **4.48%**（≤5 达标）；全语料 **8.21%**（超阈）。
- 结论限定：达标仅限冻结的单文件口径，不代表全语料/语料级降级率达标；go 为 **NOT_MEASURED**。
- 数据来源：`docs/i18n/QA-ANTI-FRAUD-AUDIT.md` 第 7 维度；数值真实、口径不同，禁止把单文件达标外推为全语料达标。

