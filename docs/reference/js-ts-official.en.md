<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->

> Language / 语言：[English (abstract)](js-ts-official.en.md) | [中文全文](js-ts-official.md)

# JS/TS Official Standards Rules (TypeScript Handbook / ECMAScript) — Abstract

**Abstract-only companion**: the full Chinese text is in
[`js-ts-official.md`](js-ts-official.md). External-specification links are
preserved as-is; no external-standard content is rewritten in this file.

## Scope

This archived reference records the TypeScript Handbook as the official
baseline for the M5 JS/TS frontend rule engine. It notes the positioning
(TS as a static type checker over JavaScript; the Handbook is not a full
language spec), the widely accepted official conventions (explicit
annotations, `interface` vs `type`, caution with `!`/`enum`/`any`,
`strict`/`noImplicitAny`/`strictNullChecks`, explicit return types on public
APIs, PascalCase types and camelCase values, generic constraints), and the
tooling baseline (`tsc --strict`, ESLint + typescript-eslint).

## Key links

- https://www.typescriptlang.org/docs/handbook/intro.html （官方，M5 规则引擎落地依据）

## Notes

- Archived 2026-11-08. The official TS "Do's and Don'ts" page (with the
  apostrophe URL) was unreachable at archiving time; the rule baseline is the
  Handbook plus the listed official conventions (as stated in the Chinese
  full text). Chinese full text remains the canonical version.
