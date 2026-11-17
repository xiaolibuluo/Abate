<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->

> Language / 语言：[English (abstract)](go-uber.en.md) | [中文全文](go-uber.md)

# Go Community Standards Rules (Uber Go Style Guide) — Abstract

**Abstract-only companion**: the full Chinese text is in
[`go-uber.md`](go-uber.md). External-specification links are preserved as-is;
no external-standard content is rewritten in this file.

## Scope

This archived reference records the Uber Go Style Guide as the community
baseline for the M3/M5 rule engine. It condenses package naming
(no `common/util/shared`), naming conventions (`Err` prefixes, no shadowing of
predeclared identifiers), declaration grouping and function ordering,
struct organization (no embedded mutexes, no embedded types in public
structs, field-name initialization), error handling (`%w` wrapping, handle an
error once, comma-ok assertions, no panics in production), concurrency
(zero-value mutexes, no unstopable goroutines), table-driven tests, and the
closing Do/Don't quick list.

## Key links

- https://github.com/uber-go/guide （社区，M3/M5 规则引擎落地依据）

## Notes

- Archived 2026-11-08; Chinese full text remains the canonical version.
