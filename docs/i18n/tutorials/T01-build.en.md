<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->

> Language / 语言：[English](T01-build.en.md) | [中文](T01-build.md)

# T01 · Offline install and build from scratch

## 1 Audience

First-time users of Abate: the goal is to build the binary from source and
confirm the build provenance.

## 2 Prerequisites

- **Go 1.26.11** installed (the real build version of this repository; the
  older `go1.27` and `Go 1.24+` values in legacy documents are outdated);
- The repository source, including the `vendor/` directory (all dependencies
  are vendored; no module downloads needed);
- All commands in this tutorial are run from the repository root.

## 3 Commands

```bash
go version
cd /path/to/abate
export GOFLAGS=-mod=vendor
go build -o /tmp/abate_verify .
/tmp/abate_verify version
```

## 4 Expected output

`go version`:

```
go version go1.26.11 linux/amd64
```

(`linux/amd64` varies by platform; the key is `go1.26.11`.)

`/tmp/abate_verify version` (6 lines, byte-for-byte stable):

```
abate 0.1.0
schema: ces.tencent.com/v1
build: deterministic
license: AGPLv3 or commercial
source: https://github.com/xiaolibuluo/Abate
watermark: 小黎的部落
```

**Verification points**: exit code `0`; exactly 6 lines containing
`license: AGPLv3 or commercial` and `watermark: 小黎的部落`.

## 5 Troubleshooting

| Symptom | Cause | Action |
|---|---|---|
| `go: downloading ...` hangs or network error | `GOFLAGS=-mod=vendor` not set, so Go tries to fetch modules | Run `export GOFLAGS=-mod=vendor` and retry (dependencies are in `vendor/`) |
| `go: go.mod requires go >= 1.26.11` style error | Local Go is older than 1.26.11 | Upgrade to Go 1.26.11 |
| Built binary reports `permission denied` | Target directory not writable or missing the execute bit | Output to a writable directory, or `chmod +x` |
| `abate version` line count differs | You ran a stale binary | Delete the old artifact and rebuild per section 3 |
