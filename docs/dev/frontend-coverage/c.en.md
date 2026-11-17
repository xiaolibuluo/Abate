<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->

> Language / 语言：[English](c.en.md) | [中文](c.md)

# C/C++ Frontend Construct Coverage List

English translation of `c.md`, the section-by-section companion under the
plan-A side-by-side naming scheme; section order and data tables correspond
one-to-one; no new conclusions are added; this document does not constitute
legal advice.

- Upstream grammar source: `https://github.com/tree-sitter/tree-sitter-c` 0.24.1
- Total upstream concrete constructs: **125**
- Mapping completeness: **100.0%** (125/125, no unmapped gaps)
- Measured coverage: **86** (**68.8%**)
  - Exact-mapping coverage: 51
  - Degraded-mapping coverage: 35
- Out-of-scope (explicitly registered with reasons): 0
- Unmapped: 0
- Accepted coverage rate (covered + out-of-scope): **68.8%**
- Number of AIR types actually hit by real parsing: 38

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
| `abstract_array_declarator` | Array | degraded | Abstract array declarator degraded to an array node |
| `abstract_function_declarator` | FuncType | degraded | Abstract function declarator degraded to a function-type node |
| `abstract_parenthesized_declarator` | Region | degraded | No dedicated AIR node for parenthesized declarators; degraded to a Region node |
| `abstract_pointer_declarator` | Pointer | degraded | Abstract pointer declarator degraded to a pointer node |
| `argument_list` | Region | degraded | No dedicated AIR node for argument lists; degraded to a Region node |
| `array_declarator` | Array | degraded | Zero output across the full corpus; reverted to degraded to avoid inflating the exact denominator |
| `attributed_declarator` | Region | degraded | No dedicated AIR node for attributed declarations; degraded to a Region node |
| `attributed_statement` | Region | degraded | No dedicated AIR node for attributed declarations; degraded to a Region node |
| `bitfield_clause` | Region | degraded | No dedicated AIR node for bitfield clauses; degraded to a Region node |
| `comma_expression` | Region | degraded | No dedicated AIR node for comma expressions; degraded to a Region node |
| `compound_literal_expression` | CompositeLit | degraded | Compound literal degraded to a composite-literal node |
| `declaration_list` | Region | degraded | No dedicated AIR node for declaration lists; degraded to a Region node |
| `enumerator_list` | Region | degraded | No dedicated AIR node for enumerator lists; degraded to a Region node |
| `field_declaration_list` | Region | degraded | No dedicated AIR node for field declaration lists; degraded to a Region node |
| `field_designator` | Region | degraded | No dedicated AIR node for initializer designators; degraded to a Region node |
| `gnu_asm_clobber_list` | Region | degraded | No dedicated AIR node for GNU inline asm; degraded to a Region node |
| `gnu_asm_expression` | Region | degraded | No dedicated AIR node for GNU inline asm; degraded to a Region node |
| `gnu_asm_goto_list` | Region | degraded | No dedicated AIR node for GNU inline asm; degraded to a Region node |
| `gnu_asm_input_operand` | Region | degraded | No dedicated AIR node for GNU inline asm; degraded to a Region node |
| `gnu_asm_input_operand_list` | Region | degraded | No dedicated AIR node for GNU inline asm; degraded to a Region node |
| `gnu_asm_output_operand` | Region | degraded | No dedicated AIR node for GNU inline asm; degraded to a Region node |
| `gnu_asm_output_operand_list` | Region | degraded | No dedicated AIR node for GNU inline asm; degraded to a Region node |
| `gnu_asm_qualifier` | Region | degraded | No dedicated AIR node for GNU inline asm; degraded to a Region node |
| `initializer_list` | CompositeLit | degraded | Initializer list degraded to a composite-literal node |
| `null` | NilLit | exact | AIR models this construct with a dedicated node |
| `parameter_declaration` | Region | degraded | No dedicated AIR node for parameter declarations; degraded to a Region node |
| `parameter_list` | Region | degraded | No dedicated AIR node for parameter declarations; degraded to a Region node |
| `parenthesized_declarator` | Region | degraded | No dedicated AIR node for parenthesized declarators; degraded to a Region node |
| `parenthesized_expression` | Region | degraded | No dedicated AIR node for parenthesized expressions; degraded to a Region node |
| `pointer_declarator` | Pointer | degraded | Zero output across the full corpus; reverted to degraded to avoid inflating the exact denominator |
| `preproc_arg` | Region | degraded | No dedicated AIR node for macro arguments; degraded to a Region node |
| `preproc_params` | Region | degraded | No dedicated AIR node for macro arguments; degraded to a Region node |
| `seh_except_clause` | Try | degraded | SEH exception structure degraded to a try node |
| `seh_finally_clause` | Try | degraded | SEH exception structure degraded to a try node |
| `seh_leave_statement` | Try | degraded | SEH exception structure degraded to a try node |
| `seh_try_statement` | Try | degraded | SEH exception structure degraded to a try node |
| `subscript_designator` | Region | degraded | No dedicated AIR node for initializer designators; degraded to a Region node |
| `subscript_range_designator` | Slice | degraded | Range subscript designator degraded to a slice node |
| `variadic_parameter` | Region | degraded | No dedicated AIR node for parameter declarations; degraded to a Region node |

## Out-of-scope entries

None.

## Full construct mapping list

| Upstream construct | AIR type | Mapping status | Real hit count | Mapping rule |
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
| `enumerator` | EnumDecl | degraded | 3 | `c.enumerator` |
| `enumerator_list` | Region | degraded | 0 | `c.enumerator-list` |
| `escape_sequence` | StrLit | degraded | 3 | `c.string-body` |
| `expression_statement` | ExprStmt | exact | 13 | `c.expr-stmt` |
| `extension_expression` | Cast | degraded | 3 | `c.extension-expr` |
| `false` | BoolLit | exact | 2 | `core.boolean` |
| `field_declaration` | Field | degraded | 9 | `c.field-decl` |
| `field_declaration_list` | Region | degraded | 0 | `c.field-decl-list` |
| `field_designator` | Region | degraded | 0 | `c.designator` |
| `field_expression` | Selector | exact | 1 | `c.field` |
| `field_identifier` | Ident | degraded | 342 | `c.field-identifier` |
| `for_statement` | For | exact | 2 | `c.for` |
| `function_declarator` | FuncDecl | degraded | 14 | `c.func-declarator` |
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
| `init_declarator` | VarDecl | degraded | 27 | `c.init-declarator` |
| `initializer_list` | CompositeLit | degraded | 0 | `c.initializer-list` |
| `initializer_pair` | Assign | degraded | 33 | `c.initializer-pair` |
| `labeled_statement` | Label | exact | 1 | `c.label` |
| `linkage_specification` | ModuleDecl | degraded | 2 | `c.linkage` |
| `macro_type_specifier` | TypeDecl | degraded | 2 | `c.macro-type` |
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
| `pointer_expression` | Star | degraded | 5 | `c.deref` |
| `preproc_arg` | Region | degraded | 0 | `c.preproc-params` |
| `preproc_call` | Directive | exact | 11 | `c.preproc` |
| `preproc_def` | Directive | exact | 11 | `c.preproc` |
| `preproc_defined` | Directive | degraded | 11 | `c.preproc-defined` |
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
| `sized_type_specifier` | Named | degraded | 120 | `c.sized-type` |
| `sizeof_expression` | Call | degraded | 3 | `c.sizeof` |
| `statement_identifier` | Ident | degraded | 342 | `c.statement-identifier` |
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
| `type_identifier` | Named | degraded | 120 | `c.type-identifier` |
| `type_qualifier` | Attribute | degraded | 1 | `c.qualifier` |
| `unary_expression` | Unary | exact | 3 | `c.unary` |
| `union_specifier` | StructDecl | degraded | 7 | `c.union` |
| `update_expression` | IncDec | exact | 1 | `c.incdec` |
| `variadic_parameter` | Region | degraded | 0 | `c.parameter-decl` |
| `while_statement` | While | exact | 3 | `c.while` |

## Degrade-rate caliber statement (honest disclosure)

`degrade*Pct` is computed per language over the **first real parsed unit
(single file)** and corresponds to the frozen threshold; the full corpus uses
the same `Degradation` formula and yields different numbers — the formula is
identical, only the statistical scope differs.

- This language: no frozen degrade-rate threshold; full corpus **28.00%**
  (above threshold).
- Conclusion qualifier: "meets threshold" holds only under the frozen
  single-file caliber and does not represent full-corpus/corpus-level degrade
  rates; go is **NOT_MEASURED**.
- Data source: dimension 7 of `docs/i18n/QA-ANTI-FRAUD-AUDIT.md`; the numbers
  are real but the calibers differ; extrapolating single-file compliance to
  full-corpus compliance is forbidden.
