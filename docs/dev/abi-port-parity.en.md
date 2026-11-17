<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->

> Language / 语言：[English](abi-port-parity.en.md) | [中文](abi-port-parity.md)

# kernel/abi Porting and Parity Notes

`kernel/abi/` is a minimal Linux syscall ABI and host-arch parameter subset
**trimmed and ported** from the in-repo gVisor sources
(`third_party/_sources/gvisor/`, snapshot
`v0.0.0-20261111045955-a42ad63bdbc8`, Apache-2.0), rewritten in the Abate
style; it does not import upstream.

## Porting scope

- amd64 / arm64 syscall number→name tables (upstream
  `pkg/sentry/syscalls/linux/linux64.go`).
- Linux errno constants (upstream `pkg/abi/linux/errno/errno.go`).
- Address and page-size primitives: `Addr` / `AddrRange` / page alignment /
  `PageSize` etc. (upstream `pkg/hostarch/*`, amd64 / arm64 build tags).

## Parity-check method

`kernel/abi/syscall_test.go` **directly reads** the upstream `linux64.go`
during `go test`, parses the AMD64 / ARM64 tables, and compares them item by
item (count, number, name) against this package's `AMD64Names` / `ARM64Names`.
If the upstream source file is updated, the test exposes the drift
immediately.

Parity-check conclusions (2026-11-13):

- amd64: **352 entries** (numbers 0–441), matching upstream item by item.
- arm64: **295 entries** (numbers 0–441, using a numbering system different
  from amd64), matching upstream item by item.
- errno: 131 base values + 2 aliases + 4 internal restart codes
  (`ERestartBlock` corresponds to upstream `ERESTART_RESTARTBLOCK`, value 516
  unchanged).

## Maintenance

When adding a syscall, take the number→name pair from the upstream table and
add it to the corresponding `*_low.go` / `*_high.go`;
`go test ./kernel/abi/...` automatically verifies consistency with upstream.
