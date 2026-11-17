<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->

> Language / 语言：[English (abstract)](go-effective.en.md) | [中文全文](go-effective.md)

# Go Official Standards Rules (Effective Go) — Abstract

**Abstract-only companion**: the full Chinese text is in
[`go-effective.md`](go-effective.md). External-specification links are
preserved as-is; no external-standard content is rewritten in this file.

## Scope

This archived reference records Effective Go as the official baseline of the
in-house Go rule set (M3/M5). It condenses the guide's sections — naming
(lowercase single-word packages, `MixedCaps`, `-er` single-method interfaces),
formatting (mandatory gofmt, opening braces on the same line), doc comments,
control structures (single `for`, `range`, no automatic fallthrough), functions
(multiple returns, `defer`), allocation, interfaces and receiver choice,
concurrency (share memory by communicating), and the closing Do/Don't list
including "never ignore errors".

## Key links

- https://go.dev/doc/effective_go （官方，M3/M5 规则引擎落地依据）

## Notes

- Archived 2026-11-08; Chinese full text remains the canonical version.
