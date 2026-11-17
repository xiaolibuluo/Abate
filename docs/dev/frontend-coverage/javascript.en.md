<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->

> Language / 语言：[English](javascript.en.md) | [中文](javascript.md)

# JavaScript/TypeScript Frontend Construct Coverage List

English translation of `javascript.md`, the section-by-section companion
under the plan-A side-by-side naming scheme; section order and data tables
correspond one-to-one; no new conclusions are added; this document does not
constitute legal advice.

- Upstream grammar source: `https://github.com/tree-sitter/tree-sitter-javascript` 0.25.0
- Total upstream concrete constructs: **114**
- Mapping completeness: **100.0%** (114/114, no unmapped gaps)
- Measured coverage: **109** (**95.6%**)
  - Exact-mapping coverage: 54
  - Degraded-mapping coverage: 55
- Out-of-scope (explicitly registered with reasons): 0
- Unmapped: 0
- Accepted coverage rate (covered + out-of-scope): **95.6%**
- Number of AIR types actually hit by real parsing: 43

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
| `class_static_block` | Block | degraded | Static initialization block degraded to a block node |
| `labeled_statement` | Label | degraded | Zero output across the full corpus; reverted to degraded to avoid inflating the exact denominator |
| `sequence_expression` | Tuple | degraded | Comma expression degraded to a tuple node |
| `statement_block` | Block | exact | AIR models this construct with a dedicated node |
| `update_expression` | IncDec | degraded | Zero output across the full corpus; reverted to degraded to avoid inflating the exact denominator |

## Out-of-scope entries

None.

## Full construct mapping list

| Upstream construct | AIR type | Mapping status | Real hit count | Mapping rule |
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

## Degrade-rate caliber statement (honest disclosure)

`degrade*Pct` is computed per language over the **first real parsed unit
(single file)** and corresponds to the frozen threshold; the full corpus uses
the same `Degradation` formula and yields different numbers — the formula is
identical, only the statistical scope differs.

- This language: no frozen degrade-rate threshold; full corpus **49.12%**
  (above threshold).
- Conclusion qualifier: "meets threshold" holds only under the frozen
  single-file caliber and does not represent full-corpus/corpus-level degrade
  rates; go is **NOT_MEASURED**.
- Data source: dimension 7 of `docs/i18n/QA-ANTI-FRAUD-AUDIT.md`; the numbers
  are real but the calibers differ; extrapolating single-file compliance to
  full-corpus compliance is forbidden.
