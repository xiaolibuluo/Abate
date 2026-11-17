<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->

> Language / 语言：[English](HEADERS.en.md) | [中文](HEADERS.md)

# Copyright Header Convention and Watermark (XCU-GOV)

> Maps to PLAN-08 XCU-GOV §3. All source files must carry the unified
> copyright header; the gate is continuously enforced by
> `scripts/header-check.sh`.

> Scope precision note (read carefully): `scripts/header-check.sh` scans
> **only `*.go` files** (excluding `vendor/` and `third_party/`) and does not
> scan `.md` files. The governance convention below that covers `.md` is a
> **governance-consistency requirement, not a gate-enforced rule**: a missing
> header in a `.md` file will not cause `header-check.sh` to FAIL. Unified
> headers in markdown documents are still expected for consistency with the
> repository's existing documentation practice.

## 1. Go source file template (byte-for-byte reuse)

```go
/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/
```

The first 7 lines (including the blank line) of `main.go` are the authoritative
style; every new `.go` file must reuse this header byte for byte.

## 2. Coverage

- `.go` / `.sh` / `.yml` / `.md` (HTML comment at the document head or YAML
  front-matter) — as a governance convention for markdown, not a
  gate-enforced rule for `.md` (the script only scans `*.go`).
- `third_party/` is **excluded**: upstream copyright headers are preserved and
  registered in `third_party/LICENSE-MANIFEST.md`.

## 3. Gate rules

`scripts/header-check.sh` reads the first 8 lines of each scanned `.go` file
and FAILs if any of the following markers is missing:

- `Copyright`
- `xiaolibuluo/Abate` (repository origin)
- `小黎的部落` (digital watermark)

## 4. Digital code watermark (three layers)

1. **Copyright-header watermark**: the template above (static, distributed
   with the source code).
2. **Build-metadata watermark**: `-ldflags` injects version/commit/build-time
   and the watermark identifier 『小黎的部落』; `abate version` prints the full
   metadata (see `scripts/release.sh`).
3. **Release-manifest watermark**: `manifest.txt` contains the watermark line;
   covered by a GPG signature (tamper protection).
