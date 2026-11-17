<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->

> Language / 语言：[English (abstract)](cpp-core.en.md) | [中文全文](cpp-core.md)

# C++ Community Standards Rules (C++ Core Guidelines, isocpp) — Abstract

**Abstract-only companion**: the full Chinese text is in [`cpp-core.md`](cpp-core.md).
External-specification links are preserved as-is; no external-standard content
is rewritten in this file.

## Scope

This archived reference records the isocpp C++ Core Guidelines as the
community baseline used by the M5 rule engine. It summarizes the P
(Philosophy), I (Interfaces), and F (Functions) sections — expressing intent
in code, RAII resource management, compile-time checking, precise strongly
typed interfaces with `Expects()`/`Ensures()`, ownership transfer via smart
pointers, and parameter-passing conventions — plus the remaining section
letters (C/EN/R/ES/Per/CP/E/Con/T/CPL/SF) that are catalogued but not yet
fully captured, and the key RAII/ownership points (`span<T>`, `zstring`,
`not_null<zstring>`).

## Key links

- https://isocpp.github.io/CppCoreGuidelines/CppCoreGuidelines （社区，M5 规则引擎落地依据）

## Notes

- Archived 2026-11-08; Chinese full text remains the canonical version.
