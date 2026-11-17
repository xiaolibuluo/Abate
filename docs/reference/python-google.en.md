<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->

> Language / 语言：[English (abstract)](python-google.en.md) | [中文全文](python-google.md)

# Python Community Standards Rules (Google Python Style Guide) — Abstract

**Abstract-only companion**: the full Chinese text is in
[`python-google.md`](python-google.md). External-specification links are
preserved as-is; no external-standard content is rewritten in this file.

## Scope

This archived reference records the Google Python Style Guide as the
community baseline for the M5 Python frontend rule engine. It condenses the
guide's sections: line length (≤80, implicit continuation in parentheses),
import rules (packages/modules only, absolute imports, four-group ordering),
formatting (4-space indent, blank-line conventions), naming (PEP 8 style),
docstrings (PEP 257 with `Args/Returns/Yields/Raises`), type annotations
(PEP 484, `X | None`, no implicit `None` defaults), error handling (no bare
`except`, minimal `try` blocks, `with` for resources), and the tooling
requirements (Pylint mandatory; Black/Pyink and pyrefly recommended).

## Key links

- https://google.github.io/styleguide/pyguide.html （社区，M5 规则引擎落地依据）

## Notes

- Archived 2026-11-08; Chinese full text remains the canonical version.
