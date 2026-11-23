# C/C++ 前端构造覆盖清单

- 上游语法事实：`https://github.com/tree-sitter/tree-sitter-c` 0.24.1
- 上游具体构造总数：**125**
- 映射完备率：**100.0%**（125/125，无未映射空洞）
- 实测覆盖：**83**（**66.4%**）
  - 精确映射覆盖：61
  - 降级映射覆盖：22
- 能力域外（显式登记并给出理由）：0
- 未映射：0
- 认可覆盖率（覆盖 + 能力域外）：**66.4%**
- 真实解析命中的 AIR 类型数：46

统计口径：覆盖率来自对该语言样例语料的**真实解析**，上游构造的 AIR 类型在解析产物中出现过才计为已覆盖。AIR 是中性中间表示，多个上游构造可能共用同一 AIR 类型，因此覆盖以 AIR 类型为粒度；每个上游构造在清单中单列一行，并标注其映射是精确映射还是降级映射。

未覆盖条目的含义：该上游构造在真实解析产物中从未产生过映射所指定的 AIR 类型。它既可能是前端尚未实现该构造，也可能是前端把它降级到了另一个 AIR 类型，二者都记为缺口。

## 未覆盖条目

| 上游构造 | 映射 AIR 类型 | 映射状态 | 原因 |
|---|---|---|---|
| `abstract_array_declarator` | Array | degraded | 抽象数组声明符降级为数组节点 |
| `abstract_function_declarator` | FuncType | degraded | 抽象函数声明符降级为函数类型节点 |
| `abstract_parenthesized_declarator` | Region | degraded | 带括号声明符无独立节点，降级为区域节点 |
| `abstract_pointer_declarator` | Pointer | degraded | 抽象指针声明符降级为指针节点 |
| `argument_list` | Region | degraded | 实参列表在 AIR 无独立节点，降级为区域节点 |
| `array_declarator` | Array | degraded | 全语料零产出，回退 degraded 避免 exact 分母虚增 |
| `attributed_declarator` | Region | degraded | 带属性声明无独立节点，降级为区域节点 |
| `attributed_statement` | Region | degraded | 带属性声明无独立节点，降级为区域节点 |
| `bitfield_clause` | Region | degraded | 位域子句无独立节点，降级为区域节点 |
| `comma_expression` | Region | degraded | 逗号表达式无独立节点，降级为区域节点 |
| `compound_literal_expression` | CompositeLit | degraded | 复合字面量降级为复合字面量节点 |
| `declaration_list` | Region | degraded | 声明列表无独立节点，降级为区域节点 |
| `enumerator_list` | Region | degraded | 枚举成员列表无独立节点，降级为区域节点 |
| `field_declaration_list` | Region | degraded | 字段声明列表无独立节点，降级为区域节点 |
| `field_designator` | Region | degraded | 初始化设计符无独立节点，降级为区域节点 |
| `gnu_asm_clobber_list` | Region | degraded | GNU 内联汇编在 AIR 无独立节点，降级为区域节点 |
| `gnu_asm_expression` | Region | degraded | GNU 内联汇编在 AIR 无独立节点，降级为区域节点 |
| `gnu_asm_goto_list` | Region | degraded | GNU 内联汇编在 AIR 无独立节点，降级为区域节点 |
| `gnu_asm_input_operand` | Region | degraded | GNU 内联汇编在 AIR 无独立节点，降级为区域节点 |
| `gnu_asm_input_operand_list` | Region | degraded | GNU 内联汇编在 AIR 无独立节点，降级为区域节点 |
| `gnu_asm_output_operand` | Region | degraded | GNU 内联汇编在 AIR 无独立节点，降级为区域节点 |
| `gnu_asm_output_operand_list` | Region | degraded | GNU 内联汇编在 AIR 无独立节点，降级为区域节点 |
| `gnu_asm_qualifier` | Region | degraded | GNU 内联汇编在 AIR 无独立节点，降级为区域节点 |
| `initializer_list` | CompositeLit | degraded | 初始化列表降级为复合字面量节点 |
| `macro_type_specifier` | CMacroType | exact | AIR 以独立节点建模该构造 |
| `null` | NilLit | exact | AIR 以独立节点建模该构造 |
| `parameter_declaration` | Region | degraded | 形参声明在 AIR 无独立节点，降级为区域节点 |
| `parameter_list` | Region | degraded | 形参声明在 AIR 无独立节点，降级为区域节点 |
| `parenthesized_declarator` | Region | degraded | 带括号声明符无独立节点，降级为区域节点 |
| `parenthesized_expression` | Region | degraded | 括号表达式无独立节点，降级为区域节点 |
| `pointer_declarator` | Pointer | degraded | 全语料零产出，回退 degraded 避免 exact 分母虚增 |
| `preproc_arg` | Region | degraded | 宏参数无独立节点，降级为区域节点 |
| `preproc_defined` | CDefined | exact | AIR 以独立节点建模该构造 |
| `preproc_params` | Region | degraded | 宏参数无独立节点，降级为区域节点 |
| `seh_except_clause` | Try | degraded | SEH 异常结构降级为 try 节点 |
| `seh_finally_clause` | Try | degraded | SEH 异常结构降级为 try 节点 |
| `seh_leave_statement` | Try | degraded | SEH 异常结构降级为 try 节点 |
| `seh_try_statement` | Try | degraded | SEH 异常结构降级为 try 节点 |
| `subscript_designator` | Region | degraded | 初始化设计符无独立节点，降级为区域节点 |
| `subscript_range_designator` | Slice | degraded | 范围下标设计符降级为切片节点 |
| `type_identifier` | CTypeIdent | exact | AIR 以独立节点建模该构造 |
| `variadic_parameter` | Region | degraded | 形参声明在 AIR 无独立节点，降级为区域节点 |

## 能力域外条目

无。

## 全量构造映射清单

| 上游构造 | AIR 类型 | 映射状态 | 真实命中次数 | 映射规则 |
|---|---|---|---|---|
| `abstract_array_declarator` | Array | degraded | 0 | `c.abstract-array-declarator` |
| `abstract_function_declarator` | FuncType | degraded | 0 | `c.abstract-func-declarator` |
| `abstract_parenthesized_declarator` | Region | degraded | 0 | `c.paren-declarator` |
| `abstract_pointer_declarator` | Pointer | degraded | 0 | `c.abstract-pointer-declarator` |
| `alignas_qualifier` | Attribute | degraded | 1 | `c.qualifier` |
| `alignof_expression` | Call | degraded | 3 | `c.alignof` |
| `argument_list` | Region | degraded | 0 | `c.argument-list` |
| `array_declarator` | Array | degraded | 0 | `c.array-declarator` |
| `assignment_expression` | Assign | exact | 33 | `c.assign` |
| `attribute` | Attribute | exact | 1 | `c.attribute` |
| `attribute_declaration` | Attribute | exact | 1 | `c.attribute` |
| `attribute_specifier` | Attribute | exact | 1 | `c.attribute` |
| `attributed_declarator` | Region | degraded | 0 | `c.attributed` |
| `attributed_statement` | Region | degraded | 0 | `c.attributed` |
| `binary_expression` | Binary | exact | 36 | `c.binary` |
| `bitfield_clause` | Region | degraded | 0 | `c.bitfield` |
| `break_statement` | Branch | exact | 3 | `c.branch` |
| `call_expression` | Call | exact | 3 | `c.call` |
| `case_statement` | Case | exact | 5 | `c.case` |
| `cast_expression` | Cast | exact | 3 | `c.cast` |
| `char_literal` | CharLit | exact | 1 | `c.char` |
| `character` | CharLit | degraded | 1 | `c.char-body` |
| `comma_expression` | Region | degraded | 0 | `c.comma-expr` |
| `comment` | Comment | exact | 1 | `core.comment` |
| `compound_literal_expression` | CompositeLit | degraded | 0 | `c.compound-literal` |
| `compound_statement` | Block | exact | 39 | `c.compound-stmt` |
| `concatenated_string` | StrLit | exact | 3 | `c.string` |
| `conditional_expression` | Ternary | exact | 2 | `c.ternary` |
| `continue_statement` | Branch | exact | 3 | `c.branch` |
| `declaration` | VarDecl | exact | 27 | `c.decl` |
| `declaration_list` | Region | degraded | 0 | `c.decl-list` |
| `do_statement` | DoWhile | exact | 1 | `c.do` |
| `else_clause` | If | exact | 1 | `c.else` |
| `enum_specifier` | EnumDecl | exact | 3 | `c.enum` |
| `enumerator` | CEnumMember | exact | 3 | `c.enumerator` |
| `enumerator_list` | Region | degraded | 0 | `c.enumerator-list` |
| `escape_sequence` | StrLit | degraded | 3 | `c.string-body` |
| `expression_statement` | ExprStmt | exact | 13 | `c.expr-stmt` |
| `extension_expression` | Cast | degraded | 3 | `c.extension-expr` |
| `false` | BoolLit | exact | 2 | `core.boolean` |
| `field_declaration` | CFieldDecl | exact | 12 | `c.field-decl` |
| `field_declaration_list` | Region | degraded | 0 | `c.field-decl-list` |
| `field_designator` | Region | degraded | 0 | `c.designator` |
| `field_expression` | Selector | exact | 1 | `c.field` |
| `field_identifier` | CFieldIdent | exact | 1 | `c.field-identifier` |
| `for_statement` | For | exact | 2 | `c.for` |
| `function_declarator` | CFuncDeclarator | exact | 14 | `c.func-declarator` |
| `function_definition` | FuncDecl | exact | 14 | `c.func` |
| `generic_expression` | Call | degraded | 3 | `c.generic-expr` |
| `gnu_asm_clobber_list` | Region | degraded | 0 | `c.gnu-asm` |
| `gnu_asm_expression` | Region | degraded | 0 | `c.gnu-asm` |
| `gnu_asm_goto_list` | Region | degraded | 0 | `c.gnu-asm` |
| `gnu_asm_input_operand` | Region | degraded | 0 | `c.gnu-asm` |
| `gnu_asm_input_operand_list` | Region | degraded | 0 | `c.gnu-asm` |
| `gnu_asm_output_operand` | Region | degraded | 0 | `c.gnu-asm` |
| `gnu_asm_output_operand_list` | Region | degraded | 0 | `c.gnu-asm` |
| `gnu_asm_qualifier` | Region | degraded | 0 | `c.gnu-asm` |
| `goto_statement` | Branch | degraded | 3 | `c.goto` |
| `identifier` | Ident | exact | 342 | `core.identifier` |
| `if_statement` | If | exact | 1 | `c.if` |
| `init_declarator` | CInitDeclarator | exact | 21 | `c.init-declarator` |
| `initializer_list` | CompositeLit | degraded | 0 | `c.initializer-list` |
| `initializer_pair` | CInitPair | exact | 2 | `c.initializer-pair` |
| `labeled_statement` | Label | exact | 1 | `c.label` |
| `linkage_specification` | ModuleDecl | degraded | 2 | `c.linkage` |
| `macro_type_specifier` | CMacroType | exact | 0 | `c.macro-type` |
| `ms_based_modifier` | Attribute | degraded | 1 | `c.qualifier` |
| `ms_call_modifier` | Attribute | degraded | 1 | `c.qualifier` |
| `ms_declspec_modifier` | Attribute | degraded | 1 | `c.qualifier` |
| `ms_pointer_modifier` | Attribute | degraded | 1 | `c.qualifier` |
| `ms_restrict_modifier` | Attribute | degraded | 1 | `c.qualifier` |
| `ms_signed_ptr_modifier` | Attribute | degraded | 1 | `c.qualifier` |
| `ms_unaligned_ptr_modifier` | Attribute | degraded | 1 | `c.qualifier` |
| `ms_unsigned_ptr_modifier` | Attribute | degraded | 1 | `c.qualifier` |
| `null` | NilLit | exact | 0 | `core.nil` |
| `number_literal` | NumLit | exact | 45 | `c.number` |
| `offsetof_expression` | Call | degraded | 3 | `c.offsetof` |
| `parameter_declaration` | Region | degraded | 0 | `c.parameter-decl` |
| `parameter_list` | Region | degraded | 0 | `c.parameter-decl` |
| `parenthesized_declarator` | Region | degraded | 0 | `c.paren-declarator` |
| `parenthesized_expression` | Region | degraded | 0 | `c.paren-expr` |
| `pointer_declarator` | Pointer | degraded | 0 | `c.pointer-declarator` |
| `pointer_expression` | CDeref | exact | 7 | `c.deref` |
| `preproc_arg` | Region | degraded | 0 | `c.preproc-params` |
| `preproc_call` | Directive | exact | 11 | `c.preproc` |
| `preproc_def` | Directive | exact | 11 | `c.preproc` |
| `preproc_defined` | CDefined | exact | 0 | `c.preproc-defined` |
| `preproc_directive` | Directive | exact | 11 | `c.preproc` |
| `preproc_elif` | Directive | exact | 11 | `c.preproc` |
| `preproc_elifdef` | Directive | exact | 11 | `c.preproc` |
| `preproc_else` | Directive | exact | 11 | `c.preproc` |
| `preproc_function_def` | Directive | exact | 11 | `c.preproc` |
| `preproc_if` | Directive | exact | 11 | `c.preproc` |
| `preproc_ifdef` | Directive | exact | 11 | `c.preproc` |
| `preproc_include` | Directive | exact | 11 | `c.preproc` |
| `preproc_params` | Region | degraded | 0 | `c.preproc-params` |
| `primitive_type` | Named | exact | 120 | `c.primitive-type` |
| `return_statement` | Return | exact | 5 | `c.return` |
| `seh_except_clause` | Try | degraded | 0 | `c.seh` |
| `seh_finally_clause` | Try | degraded | 0 | `c.seh` |
| `seh_leave_statement` | Try | degraded | 0 | `c.seh` |
| `seh_try_statement` | Try | degraded | 0 | `c.seh` |
| `sized_type_specifier` | CSizedType | exact | 1 | `c.sized-type` |
| `sizeof_expression` | Call | degraded | 3 | `c.sizeof` |
| `statement_identifier` | CLabel | exact | 1 | `c.statement-identifier` |
| `storage_class_specifier` | Attribute | degraded | 1 | `c.qualifier` |
| `string_content` | StrLit | degraded | 3 | `c.string-body` |
| `string_literal` | StrLit | exact | 3 | `c.string` |
| `struct_specifier` | StructDecl | exact | 7 | `c.struct` |
| `subscript_designator` | Region | degraded | 0 | `c.designator` |
| `subscript_expression` | Index | exact | 1 | `c.index` |
| `subscript_range_designator` | Slice | degraded | 0 | `c.subscript-range` |
| `switch_statement` | Switch | exact | 3 | `c.switch` |
| `system_lib_string` | StrLit | degraded | 3 | `c.system-lib-string` |
| `translation_unit` | Package | exact | 27 | `c.unit` |
| `true` | BoolLit | exact | 2 | `core.boolean` |
| `type_definition` | TypeDecl | exact | 2 | `c.typedef` |
| `type_descriptor` | Named | exact | 120 | `c.type-descriptor` |
| `type_identifier` | CTypeIdent | exact | 0 | `c.type-identifier` |
| `type_qualifier` | Attribute | degraded | 1 | `c.qualifier` |
| `unary_expression` | Unary | exact | 3 | `c.unary` |
| `union_specifier` | CUnion | exact | 2 | `c.union` |
| `update_expression` | IncDec | exact | 1 | `c.incdec` |
| `variadic_parameter` | Region | degraded | 0 | `c.parameter-decl` |
| `while_statement` | While | exact | 3 | `c.while` |

## 降级率口径说明（诚实披露）

`degrade*Pct` 按每语言**首个真实解析 unit（单文件）**计算，并对应冻结阈值；全语料使用同一 `Degradation` 公式，数值不同——两者公式相同、仅统计范围不同。

- 本语言：无冻结降级率阈值；全语料 **28.00%**（超阈）。
- 结论限定：达标仅限冻结的单文件口径，不代表全语料/语料级降级率达标；go 为 **NOT_MEASURED**。
- 数据来源：`docs/i18n/QA-ANTI-FRAUD-AUDIT.md` 第 7 维度；数值真实、口径不同，禁止把单文件达标外推为全语料达标。

