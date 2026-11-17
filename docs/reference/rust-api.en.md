<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->

> Language / 语言：[English (abstract)](rust-api.en.md) | [中文全文](rust-api.md)

# Rust Official Standards Rules (Rust API Guidelines, C-* Checklist) — Abstract

**Abstract-only companion**: the full Chinese text is in
[`rust-api.md`](rust-api.md). External-specification links are preserved
as-is; no external-standard content is rewritten in this file.

## Scope

This archived reference records the Rust API Guidelines C-* checklist as the
official baseline for the M5 Rust frontend rule engine. It summarizes the
checklist groups: Naming (C-CASE, C-CONV, C-GETTER, C-ITER, C-FEATURE, ...),
Interoperability (C-COMMON-TRAITS, C-CONV-TRAITS, C-SERDE, C-SEND-SYNC,
C-GOOD-ERR, ...), Macros, Documentation (C-CRATE-DOC, C-EXAMPLE,
C-QUESTION-MARK, C-FAILURE, ...), Predictability, Flexibility, Type safety
(C-NEWTYPE, C-CUSTOM-TYPE, C-BITFLAG, C-BUILDER), Dependability,
Debuggability, Future proofing (C-SEALED, ...), and Necessities
(C-STABLE, C-PERMISSIVE).

## Key links

- https://rust-lang.github.io/api-guidelines/checklist.html （官方，M5 规则引擎落地依据）

## Notes

- Archived 2026-11-08; Chinese full text remains the canonical version.
