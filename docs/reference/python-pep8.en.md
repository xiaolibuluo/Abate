<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->

> Language / 语言：[English (abstract)](python-pep8.en.md) | [中文全文](python-pep8.md)

# Python Official Standards Rules (PEP 8) — Abstract

**Abstract-only companion**: the full Chinese text is in
[`python-pep8.md`](python-pep8.md). External-specification links are
preserved as-is; no external-standard content is rewritten in this file.

## Scope

This archived reference records PEP 8 as the official baseline for the M5
Python frontend rule engine. It condenses the guide's sections: naming (no
single-char `l`/`O`/`I`, CapWords classes, `Error` suffix, `self`/`cls`,
trailing-underscore conflicts), indentation (4 spaces, no tab/space mixing,
hanging indents), line length (79/72, team relaxation to 99), imports (one
per line, top of file, three-group ordering, avoid `import *`), whitespace
rules, comments and docstrings (PEP 257), and function-level advice
(`is`/`is not`, `isinstance()`, `if not seq:`, `def` over bound lambdas).

## Key links

- https://peps.python.org/pep-0008/ （官方，M5 规则引擎落地依据）

## Notes

- Archived 2026-11-08; Chinese full text remains the canonical version.
