<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->

> Language / 语言：[English](csharp.en.md) | [中文](csharp.md)

# C# Frontend Construct Coverage List

English translation of `csharp.md`, the section-by-section companion under
the plan-A side-by-side naming scheme; section order and data tables
correspond one-to-one; no new conclusions are added; this document does not
constitute legal advice.

- Upstream grammar source: `https://github.com/tree-sitter/tree-sitter-c-sharp` 0.23.5
- Total upstream concrete constructs: **215**
- Mapping completeness: **100.0%** (215/215, no unmapped gaps)
- Measured coverage: **109** (**50.7%**)
  - Exact-mapping coverage: 91
  - Degraded-mapping coverage: 18
- Out-of-scope (explicitly registered with reasons): 0
- Unmapped: 0
- Accepted coverage rate (covered + out-of-scope): **50.7%**
- Number of AIR types actually hit by real parsing: 66

Statistical caliber: coverage comes from **real parsing** of the language's
sample corpus; an upstream construct counts as covered only if the AIR type it
maps to appears in the parse output. AIR is a neutral intermediate
representation and multiple upstream constructs may share one AIR type, so
coverage is counted at AIR-type granularity; each upstream construct gets its
own row in the list, marked as an exact mapping or a degraded mapping.

Meaning of uncovered entries: the upstream construct never produced its
mapped AIR type in real parse output. Either the frontend has not implemented
the construct yet, or the frontend degraded it to another AIR type — both are
recorded as gaps.

## Uncovered entries

| Upstream construct | Mapped AIR type | Mapping status | Reason |
|---|---|---|---|
| `accessor_list` | Region | degraded | No dedicated AIR node for accessor lists; degraded to a Region node |
| `alias_qualified_name` | CsAliasName | exact | AIR models this construct with a dedicated node |
| `and_pattern` | Region | degraded | No dedicated AIR node for pattern matching; degraded to a Region node |
| `anonymous_object_creation_expression` | CompositeLit | degraded | Anonymous object degraded to a composite-literal node |
| `argument` | Region | degraded | No dedicated AIR node for argument lists; degraded to a Region node |
| `argument_list` | Region | degraded | No dedicated AIR node for argument lists; degraded to a Region node |
| `array_creation_expression` | CompositeLit | degraded | Zero output across the full corpus; reverted to degraded to avoid inflating the exact denominator |
| `array_rank_specifier` | Region | degraded | No dedicated AIR node for array rank specifiers; degraded to a Region node |
| `array_type` | Array | degraded | Array type degraded to an array node |
| `as_expression` | Cast | degraded | `as` expression degraded to a type-assertion node |
| `attribute_argument` | Region | degraded | No dedicated AIR node for attribute arguments/targets; degraded to a Region node |
| `attribute_argument_list` | Region | degraded | No dedicated AIR node for attribute arguments/targets; degraded to a Region node |
| `attribute_target_specifier` | Region | degraded | No dedicated AIR node for attribute arguments/targets; degraded to a Region node |
| `base_list` | Region | degraded | No dedicated AIR node for base lists; degraded to a Region node |
| `bracketed_argument_list` | Region | degraded | No dedicated AIR node for argument lists; degraded to a Region node |
| `bracketed_parameter_list` | Region | degraded | No dedicated AIR node for parameter lists; degraded to a Region node |
| `calling_convention` | Region | degraded | No dedicated AIR node for calling conventions; degraded to a Region node |
| `cast_expression` | Cast | degraded | Zero output across the full corpus; reverted to degraded to avoid inflating the exact denominator |
| `catch_declaration` | Region | degraded | No dedicated AIR node for catch declarations/filter clauses; degraded to a Region node |
| `catch_filter_clause` | Region | degraded | No dedicated AIR node for catch declarations/filter clauses; degraded to a Region node |
| `checked_expression` | Region | degraded | No dedicated AIR node for checked statements/expressions; degraded to a Region node |
| `checked_statement` | Region | degraded | No dedicated AIR node for checked statements/expressions; degraded to a Region node |
| `collection_element` | Region | degraded | No dedicated AIR node for collection elements; degraded to a Region node |
| `collection_expression` | CompositeLit | degraded | Collection expression degraded to a composite-literal node |
| `constant_pattern` | Region | degraded | No dedicated AIR node for pattern matching; degraded to a Region node |
| `constructor_constraint` | Region | degraded | No dedicated AIR node for constructor constraints; degraded to a Region node |
| `constructor_initializer` | Region | degraded | No dedicated AIR node for constructor initializers; degraded to a Region node |
| `declaration_list` | Region | degraded | No dedicated AIR node for declaration lists; degraded to a Region node |
| `declaration_pattern` | Region | degraded | No dedicated AIR node for pattern matching; degraded to a Region node |
| `default_expression` | Region | degraded | No dedicated AIR node for `default` expressions; degraded to a Region node |
| `delegate_declaration` | TypeDecl | degraded | Delegate declaration degraded to a type-declaration node |
| `discard` | Region | degraded | No dedicated AIR node for discards; degraded to a Region node |
| `element_binding_expression` | Region | degraded | No dedicated AIR node for member-binding expressions; degraded to a Region node |
| `empty_statement` | ExprStmt | degraded | Empty statement degraded to an expression-statement node |
| `enum_member_declaration_list` | Region | degraded | No dedicated AIR node for enum member lists; degraded to a Region node |
| `explicit_interface_specifier` | Region | degraded | No dedicated AIR node for explicit interface qualification; degraded to a Region node |
| `expression_element` | Region | degraded | No dedicated AIR node for collection elements; degraded to a Region node |
| `expression_statement` | ExprStmt | degraded | Zero output across the full corpus; reverted to degraded to avoid inflating the exact denominator |
| `extern_alias_directive` | CsExternAlias | degraded | Zero output across the full corpus; reverted to degraded to avoid inflating the exact denominator |
| `fixed_statement` | Region | degraded | No dedicated AIR node for `fixed` statements; degraded to a Region node |
| `from_clause` | Region | degraded | No dedicated AIR node for LINQ query clauses; degraded to a Region node |
| `function_pointer_parameter` | Region | degraded | No dedicated AIR node for parameter lists; degraded to a Region node |
| `function_pointer_type` | FuncType | degraded | Function pointer type degraded to a function-type node |
| `global_statement` | Region | degraded | No dedicated AIR node for `global` statements; degraded to a Region node |
| `group_clause` | Region | degraded | No dedicated AIR node for LINQ query clauses; degraded to a Region node |
| `implicit_array_creation_expression` | CompositeLit | degraded | Implicit creation degraded to a composite-literal node |
| `implicit_object_creation_expression` | CompositeLit | degraded | Implicit creation degraded to a composite-literal node |
| `implicit_parameter` | Region | degraded | No dedicated AIR node for parameter lists; degraded to a Region node |
| `implicit_stackalloc_expression` | CompositeLit | degraded | Implicit stackalloc degraded to a composite-literal node |
| `initializer_expression` | CompositeLit | degraded | Zero output across the full corpus; reverted to degraded to avoid inflating the exact denominator |
| `interpolation` | Region | degraded | No dedicated AIR node for interpolation fragments; degraded to a Region node |
| `interpolation_alignment_clause` | Region | degraded | No dedicated AIR node for interpolation fragments; degraded to a Region node |
| `interpolation_brace` | Region | degraded | No dedicated AIR node for interpolation fragments; degraded to a Region node |
| `interpolation_format_clause` | Region | degraded | No dedicated AIR node for interpolation fragments; degraded to a Region node |
| `interpolation_quote` | Region | degraded | No dedicated AIR node for interpolation fragments; degraded to a Region node |
| `interpolation_start` | Region | degraded | No dedicated AIR node for interpolation fragments; degraded to a Region node |
| `join_clause` | Region | degraded | No dedicated AIR node for LINQ query clauses; degraded to a Region node |
| `join_into_clause` | Region | degraded | No dedicated AIR node for LINQ query clauses; degraded to a Region node |
| `labeled_statement` | Label | degraded | Zero output across the full corpus; reverted to degraded to avoid inflating the exact denominator |
| `let_clause` | Region | degraded | No dedicated AIR node for LINQ query clauses; degraded to a Region node |
| `list_pattern` | Region | degraded | No dedicated AIR node for pattern matching; degraded to a Region node |
| `makeref_expression` | Region | degraded | No dedicated AIR node for makeref expressions; degraded to a Region node |
| `member_binding_expression` | Region | degraded | No dedicated AIR node for member-binding expressions; degraded to a Region node |
| `negated_pattern` | Region | degraded | No dedicated AIR node for pattern matching; degraded to a Region node |
| `or_pattern` | Region | degraded | No dedicated AIR node for pattern matching; degraded to a Region node |
| `order_by_clause` | Region | degraded | No dedicated AIR node for LINQ query clauses; degraded to a Region node |
| `parameter` | Region | degraded | No dedicated AIR node for parameter lists; degraded to a Region node |
| `parameter_list` | Region | degraded | No dedicated AIR node for parameter lists; degraded to a Region node |
| `parenthesized_expression` | Region | degraded | No dedicated AIR node for parenthesized expressions; degraded to a Region node |
| `parenthesized_pattern` | Region | degraded | No dedicated AIR node for pattern matching; degraded to a Region node |
| `parenthesized_variable_designation` | Region | degraded | No dedicated AIR node for pattern matching; degraded to a Region node |
| `pointer_type` | Pointer | degraded | Pointer type degraded to a pointer node |
| `positional_pattern_clause` | Region | degraded | No dedicated AIR node for pattern matching; degraded to a Region node |
| `preproc_arg` | Region | degraded | No dedicated AIR node for macro arguments; degraded to a Region node |
| `primary_constructor_base_type` | Region | degraded | No dedicated AIR node for constructor initializers; degraded to a Region node |
| `property_pattern_clause` | Region | degraded | No dedicated AIR node for pattern matching; degraded to a Region node |
| `range_expression` | Region | degraded | No dedicated AIR node for range expressions; degraded to a Region node |
| `raw_string_end` | Region | degraded | No dedicated AIR node for string delimiter fragments; degraded to a Region node |
| `raw_string_start` | Region | degraded | No dedicated AIR node for string delimiter fragments; degraded to a Region node |
| `recursive_pattern` | Region | degraded | No dedicated AIR node for pattern matching; degraded to a Region node |
| `ref_expression` | Ref | degraded | Zero output across the full corpus; reverted to degraded to avoid inflating the exact denominator |
| `ref_type` | Ref | degraded | `ref`/`ref readonly` expressions degraded to reference nodes |
| `reftype_expression` | Ref | degraded | `ref`/`ref readonly` expressions degraded to reference nodes |
| `refvalue_expression` | Ref | degraded | `ref`/`ref readonly` expressions degraded to reference nodes |
| `relational_pattern` | Region | degraded | No dedicated AIR node for pattern matching; degraded to a Region node |
| `select_clause` | Region | degraded | No dedicated AIR node for LINQ query clauses; degraded to a Region node |
| `spread_element` | Spread | degraded | Zero output across the full corpus; reverted to degraded to avoid inflating the exact denominator |
| `stackalloc_expression` | CompositeLit | degraded | stackalloc degraded to a composite-literal node |
| `string_literal_encoding` | Region | degraded | No dedicated AIR node for string delimiter fragments; degraded to a Region node |
| `subpattern` | Region | degraded | No dedicated AIR node for pattern matching; degraded to a Region node |
| `switch_body` | Region | degraded | No dedicated AIR node for switch bodies; degraded to a Region node |
| `switch_expression` | Match | degraded | switch expression degraded to a match node |
| `tuple_element` | Region | degraded | No dedicated AIR node for tuple elements; degraded to a Region node |
| `tuple_pattern` | Region | degraded | No dedicated AIR node for pattern matching; degraded to a Region node |
| `type_argument_list` | GenericInst | degraded | Type argument list degraded to a generic-instantiation node |
| `type_parameter` | GenericParam | degraded | Zero output across the full corpus; reverted to degraded to avoid inflating the exact denominator |
| `type_parameter_constraint` | Region | degraded | No dedicated AIR node for type parameter lists; degraded to a Region node |
| `type_parameter_constraints_clause` | Region | degraded | No dedicated AIR node for type parameter lists; degraded to a Region node |
| `type_parameter_list` | Region | degraded | No dedicated AIR node for type parameter lists; degraded to a Region node |
| `type_pattern` | Region | degraded | No dedicated AIR node for pattern matching; degraded to a Region node |
| `unsafe_statement` | Region | degraded | No dedicated AIR node for `unsafe` statements; degraded to a Region node |
| `var_pattern` | Region | degraded | No dedicated AIR node for pattern matching; degraded to a Region node |
| `when_clause` | Region | degraded | No dedicated AIR node for LINQ query clauses; degraded to a Region node |
| `where_clause` | Region | degraded | No dedicated AIR node for LINQ query clauses; degraded to a Region node |
| `with_initializer` | Region | degraded | No dedicated AIR node for `with` initializers; degraded to a Region node |
| `yield_statement` | CsYield | exact | AIR models this construct with a dedicated node |

## Out-of-scope entries

None.

## Full construct mapping list

| Upstream construct | AIR type | Mapping status | Real hit count | Mapping rule |
|---|---|---|---|---|
| `accessor_declaration` | CsAccessor | exact | 6 | `cs.accessor` |
| `accessor_list` | Region | degraded | 0 | `cs.accessor-list` |
| `alias_qualified_name` | CsAliasName | exact | 0 | `cs.alias-qualified` |
| `and_pattern` | Region | degraded | 0 | `cs.patterns` |
| `anonymous_method_expression` | Lambda | degraded | 8 | `cs.anonymous-method` |
| `anonymous_object_creation_expression` | CompositeLit | degraded | 0 | `cs.anonymous-object` |
| `argument` | Region | degraded | 0 | `cs.arguments` |
| `argument_list` | Region | degraded | 0 | `cs.arguments` |
| `array_creation_expression` | CompositeLit | degraded | 0 | `cs.array-creation` |
| `array_rank_specifier` | Region | degraded | 0 | `cs.array-rank` |
| `array_type` | Array | degraded | 0 | `cs.array-type` |
| `arrow_expression_clause` | Lambda | degraded | 8 | `cs.arrow-clause` |
| `as_expression` | Cast | degraded | 0 | `cs.as` |
| `assignment_expression` | Assign | exact | 30 | `cs.assign` |
| `attribute` | Attribute | exact | 2 | `cs.attribute` |
| `attribute_argument` | Region | degraded | 0 | `cs.attribute-detail` |
| `attribute_argument_list` | Region | degraded | 0 | `cs.attribute-detail` |
| `attribute_list` | Attribute | exact | 2 | `cs.attribute` |
| `attribute_target_specifier` | Region | degraded | 0 | `cs.attribute-detail` |
| `await_expression` | Await | exact | 2 | `cs.await` |
| `base_list` | Region | degraded | 0 | `cs.base-list` |
| `binary_expression` | Binary | exact | 28 | `cs.binary` |
| `block` | Block | exact | 113 | `cs.block` |
| `boolean_literal` | BoolLit | exact | 3 | `cs.bool` |
| `bracketed_argument_list` | Region | degraded | 0 | `cs.arguments` |
| `bracketed_parameter_list` | Region | degraded | 0 | `cs.parameters` |
| `break_statement` | Branch | exact | 3 | `cs.branch` |
| `calling_convention` | Region | degraded | 0 | `cs.calling-convention` |
| `cast_expression` | Cast | degraded | 0 | `cs.cast` |
| `catch_clause` | Catch | exact | 1 | `cs.catch` |
| `catch_declaration` | Region | degraded | 0 | `cs.catch-detail` |
| `catch_filter_clause` | Region | degraded | 0 | `cs.catch-detail` |
| `character_literal` | CharLit | exact | 1 | `cs.char` |
| `character_literal_content` | CsCharBody | exact | 1 | `cs.char-body` |
| `checked_expression` | Region | degraded | 0 | `cs.checked` |
| `checked_statement` | Region | degraded | 0 | `cs.checked` |
| `class_declaration` | ClassDecl | exact | 44 | `cs.class` |
| `collection_element` | Region | degraded | 0 | `cs.collection-element` |
| `collection_expression` | CompositeLit | degraded | 0 | `cs.collection` |
| `comment` | Comment | exact | 2 | `core.comment` |
| `compilation_unit` | Package | exact | 52 | `cs.unit` |
| `conditional_access_expression` | CsConditionalAccess | exact | 1 | `cs.conditional-access` |
| `conditional_expression` | Ternary | exact | 4 | `cs.ternary` |
| `constant_pattern` | Region | degraded | 0 | `cs.patterns` |
| `constructor_constraint` | Region | degraded | 0 | `cs.accessibility` |
| `constructor_declaration` | CsConstructor | exact | 1 | `cs.constructor` |
| `constructor_initializer` | Region | degraded | 0 | `cs.ctor-init` |
| `continue_statement` | Branch | exact | 3 | `cs.branch` |
| `conversion_operator_declaration` | CsOperatorDecl | exact | 1 | `cs.operator` |
| `declaration_expression` | VarDecl | degraded | 34 | `cs.declaration-expr` |
| `declaration_list` | Region | degraded | 0 | `cs.declaration-list` |
| `declaration_pattern` | Region | degraded | 0 | `cs.patterns` |
| `default_expression` | Region | degraded | 0 | `cs.default-expr` |
| `delegate_declaration` | TypeDecl | degraded | 0 | `cs.delegate` |
| `destructor_declaration` | CsDestructor | exact | 1 | `cs.destructor` |
| `discard` | Region | degraded | 0 | `cs.discard` |
| `do_statement` | DoWhile | exact | 1 | `cs.do` |
| `element_access_expression` | Index | exact | 6 | `cs.index` |
| `element_binding_expression` | Region | degraded | 0 | `cs.member-binding` |
| `empty_statement` | ExprStmt | degraded | 0 | `cs.empty-stmt` |
| `enum_declaration` | EnumDecl | exact | 1 | `cs.enum` |
| `enum_member_declaration` | EnumDecl | degraded | 1 | `cs.enum-member` |
| `enum_member_declaration_list` | Region | degraded | 0 | `cs.enum-member-list` |
| `escape_sequence` | CsStringFrag | exact | 9 | `cs.string-part` |
| `event_declaration` | CsEventDecl | exact | 1 | `cs.event` |
| `event_field_declaration` | CsEventDecl | exact | 1 | `cs.event` |
| `explicit_interface_specifier` | Region | degraded | 0 | `cs.explicit-interface` |
| `expression_element` | Region | degraded | 0 | `cs.collection-element` |
| `expression_statement` | ExprStmt | degraded | 0 | `cs.expr-stmt` |
| `extern_alias_directive` | CsExternAlias | degraded | 0 | `cs.extern-alias` |
| `field_declaration` | Field | exact | 5 | `cs.field` |
| `file_scoped_namespace_declaration` | ModuleDecl | exact | 2 | `cs.namespace` |
| `finally_clause` | Finally | exact | 1 | `cs.finally` |
| `fixed_statement` | Region | degraded | 0 | `cs.fixed` |
| `for_statement` | For | exact | 1 | `cs.for` |
| `foreach_statement` | RangeFor | exact | 2 | `cs.foreach` |
| `from_clause` | Region | degraded | 0 | `cs.linq-clause` |
| `function_pointer_parameter` | Region | degraded | 0 | `cs.parameters` |
| `function_pointer_type` | FuncType | degraded | 0 | `cs.fn-pointer-type` |
| `generic_name` | Named | degraded | 93 | `cs.named-type` |
| `global_attribute` | Attribute | exact | 2 | `cs.attribute` |
| `global_statement` | Region | degraded | 0 | `cs.global` |
| `goto_statement` | Branch | degraded | 3 | `cs.goto` |
| `group_clause` | Region | degraded | 0 | `cs.linq-clause` |
| `identifier` | Ident | exact | 386 | `core.identifier` |
| `if_statement` | If | exact | 3 | `cs.if` |
| `implicit_array_creation_expression` | CompositeLit | degraded | 0 | `cs.implicit-creation` |
| `implicit_object_creation_expression` | CompositeLit | degraded | 0 | `cs.implicit-creation` |
| `implicit_parameter` | Region | degraded | 0 | `cs.parameters` |
| `implicit_stackalloc_expression` | CompositeLit | degraded | 0 | `cs.implicit-stackalloc` |
| `implicit_type` | Named | degraded | 93 | `cs.implicit-type` |
| `indexer_declaration` | CsIndexer | exact | 1 | `cs.indexer` |
| `initializer_expression` | CompositeLit | degraded | 0 | `cs.initializer` |
| `integer_literal` | NumLit | exact | 34 | `cs.number` |
| `interface_declaration` | InterfaceDecl | exact | 1 | `cs.interface` |
| `interpolated_string_expression` | CsInterpString | exact | 2 | `cs.interpolated-string` |
| `interpolation` | Region | degraded | 0 | `cs.interpolation` |
| `interpolation_alignment_clause` | Region | degraded | 0 | `cs.interpolation` |
| `interpolation_brace` | Region | degraded | 0 | `cs.interpolation` |
| `interpolation_format_clause` | Region | degraded | 0 | `cs.interpolation` |
| `interpolation_quote` | Region | degraded | 0 | `cs.interpolation` |
| `interpolation_start` | Region | degraded | 0 | `cs.interpolation` |
| `invocation_expression` | Call | exact | 96 | `cs.call` |
| `is_expression` | TypeAssert | exact | 3 | `cs.is` |
| `is_pattern_expression` | CsPattern | exact | 1 | `cs.is-pattern` |
| `join_clause` | Region | degraded | 0 | `cs.linq-clause` |
| `join_into_clause` | Region | degraded | 0 | `cs.linq-clause` |
| `labeled_statement` | Label | degraded | 0 | `cs.label` |
| `lambda_expression` | Lambda | exact | 8 | `cs.lambda` |
| `let_clause` | Region | degraded | 0 | `cs.linq-clause` |
| `list_pattern` | Region | degraded | 0 | `cs.patterns` |
| `local_declaration_statement` | ConstDecl | exact | 1 | `cs.const` |
| `local_function_statement` | CsLocalFunc | exact | 1 | `cs.local-func` |
| `lock_statement` | With | degraded | 1 | `cs.lock` |
| `makeref_expression` | Region | degraded | 0 | `cs.makeref` |
| `member_access_expression` | Field | exact | 5 | `cs.member` |
| `member_binding_expression` | Region | degraded | 0 | `cs.member-binding` |
| `method_declaration` | MethodDecl | exact | 33 | `cs.method` |
| `modifier` | CsModifier | exact | 40 | `cs.modifier` |
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
| `postfix_unary_expression` | Unary | exact | 4 | `cs.unary` |
| `predefined_type` | Named | degraded | 93 | `cs.named-type` |
| `prefix_unary_expression` | Unary | exact | 4 | `cs.unary` |
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
| `qualified_name` | Named | degraded | 93 | `cs.named-type` |
| `query_expression` | CsQuery | exact | 4 | `cs.query` |
| `range_expression` | Region | degraded | 0 | `cs.range-expr` |
| `raw_string_content` | CsStringFrag | exact | 9 | `cs.string-part` |
| `raw_string_end` | Region | degraded | 0 | `cs.string-bound` |
| `raw_string_literal` | StrLit | exact | 4 | `cs.string` |
| `raw_string_start` | Region | degraded | 0 | `cs.string-bound` |
| `real_literal` | NumLit | exact | 34 | `cs.number` |
| `record_declaration` | CsRecord | degraded | 2 | `cs.record` |
| `recursive_pattern` | Region | degraded | 0 | `cs.patterns` |
| `ref_expression` | Ref | degraded | 0 | `cs.ref` |
| `ref_type` | Ref | degraded | 0 | `cs.ref-family` |
| `reftype_expression` | Ref | degraded | 0 | `cs.ref-family` |
| `refvalue_expression` | Ref | degraded | 0 | `cs.ref-family` |
| `relational_pattern` | Region | degraded | 0 | `cs.patterns` |
| `return_statement` | Return | exact | 2 | `cs.return` |
| `scoped_type` | Named | degraded | 93 | `cs.named-type` |
| `select_clause` | Region | degraded | 0 | `cs.linq-clause` |
| `shebang_directive` | Directive | exact | 3 | `cs.shebang` |
| `sizeof_expression` | CsSizeof | degraded | 1 | `cs.sizeof` |
| `spread_element` | Spread | degraded | 0 | `cs.spread` |
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
| `tuple_type` | Tuple | degraded | 1 | `cs.tuple-type` |
| `type_argument_list` | GenericInst | degraded | 0 | `cs.type-args` |
| `type_parameter` | GenericParam | degraded | 0 | `cs.generic-param` |
| `type_parameter_constraint` | Region | degraded | 0 | `cs.type-params` |
| `type_parameter_constraints_clause` | Region | degraded | 0 | `cs.type-params` |
| `type_parameter_list` | Region | degraded | 0 | `cs.type-params` |
| `type_pattern` | Region | degraded | 0 | `cs.patterns` |
| `typeof_expression` | CsTypeof | exact | 1 | `cs.typeof` |
| `unary_expression` | Unary | exact | 4 | `cs.unary` |
| `unsafe_statement` | Region | degraded | 0 | `cs.unsafe` |
| `using_directive` | UsingDecl | exact | 1 | `cs.using-directive` |
| `using_statement` | With | exact | 1 | `cs.using-stmt` |
| `var_pattern` | Region | degraded | 0 | `cs.patterns` |
| `variable_declaration` | VarDecl | exact | 34 | `cs.var` |
| `variable_declarator` | VarDecl | degraded | 34 | `cs.declarator` |
| `verbatim_string_literal` | StrLit | exact | 4 | `cs.string` |
| `when_clause` | Region | degraded | 0 | `cs.linq-clause` |
| `where_clause` | Region | degraded | 0 | `cs.linq-clause` |
| `while_statement` | While | exact | 2 | `cs.while` |
| `with_expression` | With | degraded | 1 | `cs.with-expr` |
| `with_initializer` | Region | degraded | 0 | `cs.with-init` |
| `yield_statement` | CsYield | exact | 0 | `cs.yield` |

## Degrade-rate caliber statement (honest disclosure)

`degrade*Pct` is computed per language over the **first real parsed unit
(single file)** and corresponds to the frozen threshold; the full corpus uses
the same `Degradation` formula and yields different numbers — the formula is
identical, only the statistical scope differs.

- This language: single file **3.26%** (≤5, meets threshold); full corpus
  **6.98%** (above threshold).
- Conclusion qualifier: "meets threshold" holds only under the frozen
  single-file caliber and does not represent full-corpus/corpus-level degrade
  rates; go is **NOT_MEASURED**.
- Data source: dimension 7 of `docs/i18n/QA-ANTI-FRAUD-AUDIT.md`; the numbers
  are real but the calibers differ; extrapolating single-file compliance to
  full-corpus compliance is forbidden.
