<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->

> Language / 语言：[English](python.en.md) | [中文](python.md)

# Python Frontend Construct Coverage List

English translation of `python.md`, the section-by-section companion under
the plan-A side-by-side naming scheme; section order and data tables
correspond one-to-one; no new conclusions are added; this document does not
constitute legal advice.

- Upstream grammar source: `https://github.com/tree-sitter/tree-sitter-python` 0.25.0
- Total upstream concrete constructs: **122**
- Mapping completeness: **100.0%** (122/122, no unmapped gaps)
- Measured coverage: **110** (**90.2%**)
  - Exact-mapping coverage: 75
  - Degraded-mapping coverage: 35
- Out-of-scope (explicitly registered with reasons): 3
- Unmapped: 0
- Accepted coverage rate (covered + out-of-scope): **92.6%**
- Number of AIR types actually hit by real parsing: 57

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
| `augmented_assignment` | PyAugAssign | exact | AIR models this construct with a dedicated node |
| `constrained_type` | GenericParam | degraded | PEP 695 constrained type parameter degraded to a generic-parameter node |
| `dotted_name` | PyQualifiedName | exact | AIR models this construct with a dedicated node |
| `member_type` | PyQualifiedName | exact | AIR models this construct with a dedicated node |
| `not_operator` | PyNotOp | exact | AIR models this construct with a dedicated node |
| `type_alias_statement` | TypeDecl | degraded | PEP 695 `type` alias: zero corpus output; locked at the degraded caliber |
| `type_conversion` | PyTypeConversion | exact | AIR models this construct with a dedicated node |
| `type_parameter` | GenericParam | degraded | PEP 695 `type` parameter: zero corpus output; locked at the degraded caliber |
| `union_type` | Union | degraded | PEP 604 union type: zero corpus output; locked at the degraded caliber |

## Out-of-scope entries

| Upstream construct | Reason |
|---|---|
| `chevron` | Python 2 legacy print chevron; outside the Python 3.8–3.12 capability scope |
| `exec_statement` | Python 2 legacy exec statement; outside the Python 3.8–3.12 capability scope |
| `print_statement` | Python 2 legacy print statement; outside the Python 3.8–3.12 capability scope |

## Full construct mapping list

| Upstream construct | AIR type | Mapping status | Real hit count | Mapping rule |
|---|---|---|---|---|
| `aliased_import` | Import | exact | 6 | `py.import` |
| `argument_list` | Region | degraded | 1 | `py.argument-list` |
| `as_pattern` | Region | degraded | 1 | `py.patterns` |
| `assert_statement` | PyAssert | exact | 2 | `py.assert` |
| `assignment` | Assign | exact | 44 | `py.assign` |
| `attribute` | Field | exact | 5 | `py.attribute` |
| `augmented_assignment` | PyAugAssign | exact | 0 | `py.assign-augmented` |
| `await` | Await | degraded | 3 | `py.await` |
| `binary_operator` | Binary | exact | 27 | `py.binary` |
| `block` | Block | exact | 54 | `py.block` |
| `boolean_operator` | PyBoolOp | exact | 2 | `py.boolean-op` |
| `break_statement` | Branch | exact | 4 | `py.branch` |
| `call` | Call | exact | 31 | `py.call` |
| `case_clause` | Case | degraded | 7 | `py.case` |
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
| `expression_list` | Tuple | degraded | 4 | `py.expr-list` |
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
| `list_splat` | Star | degraded | 6 | `py.list-splat` |
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
| `parenthesized_list_splat` | Star | degraded | 6 | `py.list-splat` |
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
| `splat_type` | Star | degraded | 6 | `py.splat-type` |
| `string` | StrLit | exact | 6 | `py.string` |
| `string_content` | PyStringFragment | exact | 10 | `py.string-part` |
| `string_end` | PyStringFragment | exact | 10 | `py.string-part` |
| `string_start` | PyStringFragment | exact | 10 | `py.string-part` |
| `subscript` | Index | exact | 4 | `py.subscript` |
| `true` | BoolLit | exact | 4 | `core.boolean` |
| `try_statement` | Try | exact | 4 | `py.try` |
| `tuple` | Tuple | exact | 4 | `py.tuple` |
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

## Degrade-rate caliber statement (honest disclosure)

`degrade*Pct` is computed per language over the **first real parsed unit
(single file)** and corresponds to the frozen threshold; the full corpus uses
the same `Degradation` formula and yields different numbers — the formula is
identical, only the statistical scope differs.

- This language: single file **0%** (≤3, meets threshold); full corpus
  **26.89%** (above threshold).
- Conclusion qualifier: "meets threshold" holds only under the frozen
  single-file caliber and does not represent full-corpus/corpus-level degrade
  rates; go is **NOT_MEASURED**.
- Data source: dimension 7 of `docs/i18n/QA-ANTI-FRAUD-AUDIT.md`; the numbers
  are real but the calibers differ; extrapolating single-file compliance to
  full-corpus compliance is forbidden.
