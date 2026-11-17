<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->

> Language / 语言：[English (abstract)](c-cert.en.md) | [中文全文](c-cert.md)

# C/C++ Official Standards Rules (SEI CERT C/C++ Secure Coding) — Abstract

**Abstract-only companion**: the full Chinese text is in [`c-cert.md`](c-cert.md).
External-specification links are preserved as-is; no external-standard content
is rewritten in this file.

## Scope

This archived reference records the SEI CERT C/C++ Secure Coding Standard as
the security-rule source for the M5 C/C++ frontend rule engine. It lists the
CERT rule groups (PRE/DCL/EXP/INT/FLP/ARR/STR/MEM/FIO/ENV/SIG/ERR/API/CON/
MSC/POS/WIN) and the core security themes the rule set is designed around:
memory safety, integer safety, bounded string handling, paired
allocation/deallocation, input validation, concurrency, and error handling.
CERT distinguishes normative Rules from advisory Recommendations, and ships
the Juliet test suite (61,387 cases) for static-analysis coverage verification.

## Key links

- https://wiki.sei.cmu.edu/confluence/display/c/SEI+CERT+C+Coding+Standard （官方安全编码基线）

## Notes

- Archived 2026-11-08; Chinese full text remains the canonical version.
