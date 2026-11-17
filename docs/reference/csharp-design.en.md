<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->

> Language / 语言：[English (abstract)](csharp-design.en.md) | [中文全文](csharp-design.md)

# C# Official / Community Standards Rules (.NET Framework Design Guidelines + StyleCop) — Abstract

**Abstract-only companion**: the full Chinese text is in
[`csharp-design.md`](csharp-design.md). External-specification links are
preserved as-is; no external-standard content is rewritten in this file.

## Scope

This archived reference records the official .NET Framework Design Guidelines
(naming, type design, member design, exceptions, usage guidelines) together
with the community StyleCop Analyzers rule families (SA1xxx documentation,
SA11xx readability, SA12xx ordering, SA13xx maintainability, SA14xx layout,
SA15xx spacing, SX1xxx naming) as the basis for the M5 C# frontend rule
engine. It also notes the tooling: Roslyn analyzers with per-rule severity
configuration and EditorConfig.

## Key links

- https://learn.microsoft.com/en-us/dotnet/standard/design-guidelines/ （官方）
- https://github.com/DotNetAnalyzers/StyleCopAnalyzers （StyleCop 规则 SAxxxx）

## Notes

- Archived 2026-11-08; Chinese full text remains the canonical version.
