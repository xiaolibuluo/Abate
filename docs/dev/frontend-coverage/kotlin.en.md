<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->

> Language / 语言：[English](kotlin.en.md) | [中文](kotlin.md)

# Kotlin Frontend Construct Coverage List

English translation of `kotlin.md`, the section-by-section companion under
the plan-A side-by-side naming scheme; section order and data tables
correspond one-to-one; no new conclusions are added; this document does not
constitute legal advice.

- Upstream grammar source: `https://github.com/tree-sitter/tree-sitter-kotlin` 0.3.8
- Total upstream concrete constructs: **134**
- Mapping completeness: **100.0%** (134/134, no unmapped gaps)
- Measured coverage: **69** (**51.5%**)
  - Exact-mapping coverage: 58
  - Degraded-mapping coverage: 11
- Out-of-scope (explicitly registered with reasons): 0
- Unmapped: 0
- Accepted coverage rate (covered + out-of-scope): **51.5%**
- Number of AIR types actually hit by real parsing: 44

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
| `anonymous_initializer` | CompositeLit | degraded | Anonymous initializer block degraded to a composite-literal node |
| `as_expression` | Cast | degraded | `as` expression degraded to a type-assertion node |
| `binding_pattern_kind` | Region | degraded | No dedicated AIR node for binding patterns; degraded to a Region node |
| `call_suffix` | Region | degraded | No dedicated AIR node for call/navigation suffixes; degraded to a Region node |
| `check_expression` | Region | degraded | No dedicated AIR node for check expressions; degraded to a Region node |
| `class_body` | Region | degraded | No dedicated AIR node for class bodies; degraded to a Region node |
| `class_modifier` | Attribute | degraded | Declaration modifier degraded to an attribute node |
| `class_parameter` | Region | degraded | No dedicated AIR node for primary-constructor parameters; degraded to a Region node |
| `collection_literal` | CompositeLit | degraded | Zero output across the full corpus; reverted to degraded to avoid inflating the exact denominator |
| `constructor_delegation_call` | KtCtorDelegation | exact | AIR models this construct with a dedicated node |
| `delegation_specifier` | Region | degraded | No dedicated AIR node for delegation specifiers; degraded to a Region node |
| `directly_assignable_expression` | Region | degraded | No dedicated AIR node for assignable-expression markers; degraded to a Region node |
| `elvis_expression` | Ternary | degraded | Zero output across the full corpus; reverted to degraded to avoid inflating the exact denominator |
| `enum_class_body` | Region | degraded | No dedicated AIR node for enum bodies; degraded to a Region node |
| `explicit_delegation` | Region | degraded | No dedicated AIR node for delegation specifiers; degraded to a Region node |
| `function_modifier` | Attribute | degraded | Declaration modifier degraded to an attribute node |
| `function_type` | FuncType | degraded | Zero output across the full corpus; reverted to degraded to avoid inflating the exact denominator |
| `function_type_parameters` | Region | degraded | No dedicated AIR node for function-type parameter lists; degraded to a Region node |
| `function_value_parameters` | Region | degraded | No dedicated AIR node for function-type parameter lists; degraded to a Region node |
| `indexing_suffix` | Region | degraded | No dedicated AIR node for call/navigation suffixes; degraded to a Region node |
| `inheritance_modifier` | Attribute | degraded | Declaration modifier degraded to an attribute node |
| `interpolated_identifier` | Region | degraded | No dedicated AIR node for template identifier fragments; degraded to a Region node |
| `label` | Label | degraded | Zero output across the full corpus; reverted to degraded to avoid inflating the exact denominator |
| `lambda_parameters` | Region | degraded | No dedicated AIR node for lambda parameters; degraded to a Region node |
| `member_modifier` | Attribute | degraded | Declaration modifier degraded to an attribute node |
| `modifiers` | Attribute | degraded | Declaration modifier degraded to an attribute node |
| `navigation_suffix` | Region | degraded | No dedicated AIR node for call/navigation suffixes; degraded to a Region node |
| `not_nullable_type` | Named | degraded | Non-nullable type degraded to a named-type node |
| `object_literal` | CompositeLit | degraded | Zero output across the full corpus; reverted to degraded to avoid inflating the exact denominator |
| `parameter` | Region | degraded | No dedicated AIR node for primary-constructor parameters; degraded to a Region node |
| `parameter_modifier` | Attribute | degraded | Declaration modifier degraded to an attribute node |
| `parameter_modifiers` | Attribute | degraded | Declaration modifier degraded to an attribute node |
| `parameter_with_optional_type` | Region | degraded | No dedicated AIR node for primary-constructor parameters; degraded to a Region node |
| `parenthesized_expression` | Region | degraded | No dedicated AIR node for parenthesized expressions; degraded to a Region node |
| `parenthesized_type` | Named | degraded | Parenthesized type degraded to a named-type node |
| `parenthesized_user_type` | Named | degraded | Parenthesized type degraded to a named-type node |
| `platform_modifier` | Attribute | degraded | Declaration modifier degraded to an attribute node |
| `primary_constructor` | Region | degraded | No dedicated AIR node for primary constructors; degraded to a Region node |
| `property_delegate` | Region | degraded | No dedicated AIR node for property delegation; degraded to a Region node |
| `property_modifier` | Attribute | degraded | Declaration modifier degraded to an attribute node |
| `range_test` | Region | degraded | No dedicated AIR node for range tests; degraded to a Region node |
| `reification_modifier` | Attribute | degraded | Declaration modifier degraded to an attribute node |
| `shebang_line` | Directive | degraded | Zero output across the full corpus; reverted to degraded to avoid inflating the exact denominator |
| `spread_expression` | Spread | degraded | Zero output across the full corpus; reverted to degraded to avoid inflating the exact denominator |
| `statements` | Region | degraded | No dedicated AIR node for statement lists; degraded to a Region node |
| `super_expression` | KtSuper | exact | AIR models this construct with a dedicated node |
| `type_arguments` | GenericInst | degraded | Type argument list degraded to a generic-instantiation node |
| `type_constraint` | Region | degraded | No dedicated AIR node for type parameter lists; degraded to a Region node |
| `type_constraints` | Region | degraded | No dedicated AIR node for type parameter lists; degraded to a Region node |
| `type_identifier` | Named | degraded | Zero output across the full corpus; reverted to degraded to avoid inflating the exact denominator |
| `type_modifiers` | Attribute | degraded | Type modifier degraded to an attribute node |
| `type_parameter` | GenericParam | degraded | Zero output across the full corpus; reverted to degraded to avoid inflating the exact denominator |
| `type_parameter_modifiers` | Region | degraded | No dedicated AIR node for type parameter lists; degraded to a Region node |
| `type_parameters` | Region | degraded | No dedicated AIR node for type parameter lists; degraded to a Region node |
| `type_projection` | GenericInst | degraded | Type argument list degraded to a generic-instantiation node |
| `type_projection_modifiers` | Region | degraded | No dedicated AIR node for type parameter lists; degraded to a Region node |
| `use_site_target` | Region | degraded | No dedicated AIR node for use-site targets; degraded to a Region node |
| `user_type` | Named | degraded | Zero output across the full corpus; reverted to degraded to avoid inflating the exact denominator |
| `value_argument` | Region | degraded | No dedicated AIR node for argument lists; degraded to a Region node |
| `value_arguments` | Region | degraded | No dedicated AIR node for argument lists; degraded to a Region node |
| `variance_modifier` | Attribute | degraded | Declaration modifier degraded to an attribute node |
| `visibility_modifier` | Attribute | degraded | Declaration modifier degraded to an attribute node |
| `when_condition` | Region | degraded | No dedicated AIR node for when conditions; degraded to a Region node |
| `when_entry` | KtWhenEntry | exact | AIR models this construct with a dedicated node |
| `when_expression` | Match | exact | AIR models this construct with a dedicated node |

## Out-of-scope entries

None.

## Full construct mapping list

| Upstream construct | AIR type | Mapping status | Real hit count | Mapping rule |
|---|---|---|---|---|
| `additive_expression` | Binary | exact | 25 | `kt.additive` |
| `annotated_lambda` | Lambda | degraded | 10 | `kt.annotated-lambda` |
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
| `constructor_invocation` | Call | degraded | 60 | `kt.ctor-call` |
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
| `for_statement` | RangeFor | degraded | 2 | `kt.for` |
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
| `infix_expression` | Call | degraded | 60 | `kt.infix-call` |
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
| `multi_variable_declaration` | VarDecl | degraded | 45 | `kt.multi-var` |
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
| `range_expression` | RangeFor | degraded | 2 | `kt.range` |
| `range_test` | Region | degraded | 0 | `kt.range-test` |
| `real_literal` | NumLit | exact | 34 | `kt.number` |
| `reification_modifier` | Attribute | degraded | 0 | `kt.modifiers` |
| `secondary_constructor` | KtSecondaryCtor | exact | 1 | `kt.secondary-ctor` |
| `setter` | KtAccessor | exact | 1 | `kt.accessor` |
| `shebang_line` | Directive | degraded | 0 | `kt.shebang` |
| `simple_identifier` | Ident | degraded | 361 | `kt.simple-identifier` |
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

## Degrade-rate caliber statement (honest disclosure)

`degrade*Pct` is computed per language over the **first real parsed unit
(single file)** and corresponds to the frozen threshold; the full corpus uses
the same `Degradation` formula and yields different numbers — the formula is
identical, only the statistical scope differs.

- This language: single file **4.48%** (≤5, meets threshold); full corpus
  **8.21%** (above threshold).
- Conclusion qualifier: "meets threshold" holds only under the frozen
  single-file caliber and does not represent full-corpus/corpus-level degrade
  rates; go is **NOT_MEASURED**.
- Data source: dimension 7 of `docs/i18n/QA-ANTI-FRAUD-AUDIT.md`; the numbers
  are real but the calibers differ; extrapolating single-file compliance to
  full-corpus compliance is forbidden.
