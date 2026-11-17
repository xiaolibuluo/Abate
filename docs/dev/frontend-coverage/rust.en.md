<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->

> Language / 语言：[English](rust.en.md) | [中文](rust.md)

# Rust Frontend Construct Coverage List

English translation of `rust.md`, the section-by-section companion under the
plan-A side-by-side naming scheme; section order and data tables correspond
one-to-one; no new conclusions are added; this document does not constitute
legal advice.

- Upstream grammar source: `https://github.com/tree-sitter/tree-sitter-rust` 0.24.0
- Total upstream concrete constructs: **163**
- Mapping completeness: **100.0%** (163/163, no unmapped gaps)
- Measured coverage: **132** (**81.0%**)
  - Exact-mapping coverage: 65
  - Degraded-mapping coverage: 67
- Out-of-scope (explicitly registered with reasons): 0
- Unmapped: 0
- Accepted coverage rate (covered + out-of-scope): **81.0%**
- Number of AIR types actually hit by real parsing: 55

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
| `array_type` | Array | degraded | Zero output across the full corpus; reverted to degraded to avoid inflating the exact denominator |
| `async_block` | RsAsyncBlock | degraded | Zero output across the full corpus; reverted to degraded to avoid inflating the exact denominator |
| `base_field_initializer` | RsBaseField | exact | AIR models this construct with a dedicated node |
| `block_comment` | Comment | exact | AIR models this construct with a dedicated node |
| `const_block` | RsConstBlock | degraded | Zero output across the full corpus; reverted to degraded to avoid inflating the exact denominator |
| `const_item` | ConstDecl | exact | AIR models this construct with a dedicated node |
| `const_parameter` | ConstDecl | degraded | const parameter degraded to a constant-declaration node |
| `empty_statement` | RsEmptyStmt | degraded | Zero output across the full corpus; reverted to degraded to avoid inflating the exact denominator |
| `extern_crate_declaration` | RsExternCrate | exact | AIR models this construct with a dedicated node |
| `foreign_mod_item` | ModuleDecl | degraded | Foreign-module item degraded to a module-declaration node |
| `function_type` | FuncType | degraded | Zero output across the full corpus; reverted to degraded to avoid inflating the exact denominator |
| `gen_block` | RsGenBlock | degraded | Zero output across the full corpus; reverted to degraded to avoid inflating the exact denominator |
| `generic_type_with_turbofish` | RsTurbofish | exact | AIR models this construct with a dedicated node |
| `label` | Label | degraded | Zero output across the full corpus; reverted to degraded to avoid inflating the exact denominator |
| `line_comment` | Comment | exact | AIR models this construct with a dedicated node |
| `macro_definition` | RsMacroDef | exact | AIR models this construct with a dedicated node |
| `metavariable` | RsMetavar | exact | AIR models this construct with a dedicated node |
| `mod_item` | ModuleDecl | exact | AIR models this construct with a dedicated node |
| `negative_literal` | RsNegativeLit | exact | AIR models this construct with a dedicated node |
| `pointer_type` | Pointer | degraded | Zero output across the full corpus; reverted to degraded to avoid inflating the exact denominator |
| `reference_expression` | Addr | exact | AIR models this construct with a dedicated node |
| `return_expression` | Return | exact | AIR models this construct with a dedicated node |
| `shebang` | Directive | degraded | Zero output across the full corpus; reverted to degraded to avoid inflating the exact denominator |
| `shorthand_field_identifier` | RsShorthandIdent | exact | AIR models this construct with a dedicated node |
| `shorthand_field_initializer` | RsShorthandField | exact | AIR models this construct with a dedicated node |
| `static_item` | ConstDecl | exact | AIR models this construct with a dedicated node |
| `super` | RsSuper | exact | AIR models this construct with a dedicated node |
| `try_block` | RsTryBlock | degraded | Zero output across the full corpus; reverted to degraded to avoid inflating the exact denominator |
| `try_expression` | Try | exact | AIR models this construct with a dedicated node |
| `use_as_clause` | RsUseAs | exact | AIR models this construct with a dedicated node |
| `yield_expression` | RsYield | exact | AIR models this construct with a dedicated node |

## Out-of-scope entries

None.

## Full construct mapping list

| Upstream construct | AIR type | Mapping status | Real hit count | Mapping rule |
|---|---|---|---|---|
| `abstract_type` | Region | degraded | 1 | `rs.abstract-type` |
| `arguments` | Region | degraded | 1 | `rs.arguments` |
| `array_expression` | CompositeLit | degraded | 3 | `rs.array` |
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
| `enum_variant` | EnumDecl | degraded | 2 | `rs.enum-variant` |
| `enum_variant_list` | Region | degraded | 1 | `rs.enum-variant-list` |
| `escape_sequence` | RsStringFrag | exact | 12 | `rs.string-part` |
| `expression_statement` | ExprStmt | exact | 35 | `rs.expr-stmt` |
| `extern_crate_declaration` | RsExternCrate | exact | 0 | `rs.extern-crate` |
| `extern_modifier` | Attribute | degraded | 3 | `rs.modifier` |
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
| `function_modifiers` | Attribute | degraded | 3 | `rs.modifier` |
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
| `impl_item` | MethodDecl | degraded | 4 | `rs.impl` |
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
| `loop_expression` | While | degraded | 2 | `rs.loop` |
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
| `mutable_specifier` | Attribute | degraded | 3 | `rs.modifier` |
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
| `tuple_type` | Tuple | degraded | 1 | `rs.tuple-type` |
| `type_arguments` | RsTypeArgs | exact | 9 | `rs.type-args` |
| `type_binding` | TypeDecl | degraded | 2 | `rs.type-binding` |
| `type_cast_expression` | Cast | exact | 2 | `rs.cast` |
| `type_identifier` | Named | exact | 1 | `rs.type-identifier` |
| `type_item` | TypeDecl | exact | 2 | `rs.type` |
| `type_parameter` | GenericParam | exact | 2 | `rs.generic-param` |
| `type_parameters` | Region | degraded | 1 | `rs.type-params` |
| `unary_expression` | Unary | exact | 2 | `rs.unary` |
| `union_item` | StructDecl | degraded | 5 | `rs.union` |
| `unit_expression` | Region | degraded | 1 | `rs.unit-expr` |
| `unit_type` | Named | degraded | 1 | `rs.unit-type` |
| `unsafe_block` | RsUnsafeBlock | exact | 2 | `rs.unsafe-block` |
| `use_as_clause` | RsUseAs | exact | 0 | `rs.use-as` |
| `use_bounds` | Import | degraded | 4 | `rs.use-bounds` |
| `use_declaration` | Import | exact | 4 | `rs.use` |
| `use_list` | RsUseList | exact | 1 | `rs.use-list` |
| `use_wildcard` | RsUseWildcard | exact | 1 | `rs.use-wildcard` |
| `variadic_parameter` | Region | degraded | 1 | `rs.parameters` |
| `visibility_modifier` | Attribute | degraded | 3 | `rs.modifier` |
| `where_clause` | Region | degraded | 1 | `rs.where-clause` |
| `where_predicate` | Region | degraded | 1 | `rs.where-clause` |
| `while_expression` | While | exact | 2 | `rs.while` |
| `yield_expression` | RsYield | exact | 0 | `rs.yield` |

## Degrade-rate caliber statement (honest disclosure)

`degrade*Pct` is computed per language over the **first real parsed unit
(single file)** and corresponds to the frozen threshold; the full corpus uses
the same `Degradation` formula and yields different numbers — the formula is
identical, only the statistical scope differs.

- This language: single file **0.61%** (≤3, meets threshold); full corpus
  **42.33%** (above threshold).
- Conclusion qualifier: "meets threshold" holds only under the frozen
  single-file caliber and does not represent full-corpus/corpus-level degrade
  rates; go is **NOT_MEASURED**.
- Data source: dimension 7 of `docs/i18n/QA-ANTI-FRAUD-AUDIT.md`; the numbers
  are real but the calibers differ; extrapolating single-file compliance to
  full-corpus compliance is forbidden.
