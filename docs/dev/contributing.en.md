<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->

> Language / 语言：[English](contributing.en.md) | [中文](contributing.md)

# Contributing Guide

English translation of `contributing.md`, the section-by-section companion
under the plan-A side-by-side naming scheme; section order and code blocks
correspond one-to-one; no new conclusions are added; this document does not
constitute legal advice. This file is distinct from the root
`CONTRIBUTING.md`; both are kept.

## 1. Agreements

- **DCO mandatory**: every commit must carry a `Signed-off-by` line
  (`git commit -s`), declaring that you have the right to submit and agree to
  DCO 1.1.
- **CLA optional**: once signed, it explicitly grants the project the right to
  use and relicense the contribution.
- Contributing grants **dual-license** rights (AGPLv3 + commercial); see
  `governance/DUAL-LICENSE.md`.

## 2. Gates (all green required before merging)

```bash
# Five-stage gate chain
./scripts/lint-gate.sh
# Copyright-header gate
./scripts/header-check.sh
# License gate
./scripts/license-check.sh
# Determinism gate (same-seed double run on core packages)
./scripts/determinism-check.sh
# Drift-check gate
./scripts/drift-check.sh
```

Recommended local one-liner sequence: `go test -race -count=1 ./...` followed
by the gate scripts above.

## 3. Red lines (untouchable)

- **Deep in-house integration of open-source code**: third-party dependencies
  are always included at source level (vendoring / module source); API or
  network access is **forbidden** (except toolchain downloads).
- **Determinism iron rules**: no `time.Now` (real clock), no global
  `math/rand`, no reliance on map iteration order, no behavior driven by
  `os.Environ` output; kernel L0 has zero dependence on external networks.
- **Copyright header**: reuse byte for byte (see `governance/HEADERS.md`);
  every new file must carry it.
- **≤300 lines per file**; **no TODO/FIXME** (blocked by godox).
- **No leftover scratch/debug test files**.

## 4. Test requirements

- Table-driven, covering normal/boundary/error paths.
- New or modified code must include a determinism replay test (same-seed
  double run with identical hashes).
- Verify only your own package: `go test -race -count=1 ./<pkg>`
  (offline `-mod=vendor`).
- Coverage: kernel ≥95%/85% (statement/branch), ecosystem ≥85%/75%.

## 5. Submission flow

1. Branch off `main`; all gates green.
2. `git commit -s` (DCO).
3. PR description links the milestone and acceptance items; gate failures
   print the rule that was hit.
4. After merge, the CI full archive re-checks (nightly: fuzz + parity +
   scale).
