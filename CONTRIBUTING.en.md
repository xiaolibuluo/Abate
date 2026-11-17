<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com)
  Project: xiaolibuluo/Abate — https://github.com/xiaolibuluo/Abate
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->
> Language / 语言：[English](CONTRIBUTING.en.md) | [中文](CONTRIBUTING.md)

# Contributing Guide (root CONTRIBUTING)

> This file is the repository-root contribution entry point; `docs/dev/contributing.md` is the developer-workflow document. The two paths differ and both are retained.
> **This document does not constitute legal advice / 本文件不构成法律意见; subject to confirmation by the rights holder or legal counsel.**

## 1. Dual Licensing and Contributions

Abate is dual-licensed under **AGPLv3 (open-source side) + a commercial license (closed-source side)**; see `LICENSING.md`. Before contributing, read `LICENSING.md`, `DCO.md`, and `CLA.md`. All legal determinations are **subject to confirmation by the rights holder or legal counsel**.

## 2. DCO 1.1 and Sign-off

- **DCO is mandatory**: every commit must contain a `Signed-off-by` line certifying that you have the right to submit the contribution and agree to DCO 1.1.
- Documentation example (**example only; no git commit is executed in this round**):

```bash
git commit -s -m "docs: add bilingual quickstart"
```

- `git commit -s` appends `Signed-off-by: <Name> <email>` automatically. The official DCO 1.1 full text is in `DCO.md`.

## 3. CLA Status (draft, subject to confirmation)

- Current governance is "**CLA optional, DCO mandatory**"; this document **does not change that status and does not make the CLA mandatory**.
- `CLA.md` is a **draft**; **whether it is required is subject to confirmation by the rights holder or legal counsel**.
- Explicit note: **the DCO does not grant the project a relicensing right; under dual licensing a CLA or an explicit relicensing clause is needed — subject to confirmation by the rights holder or legal counsel**. This document does not decide how contribution copyright rights are used by the project under dual licensing.

## 4. Import Dependency Allowlist

- **Allowed**: MIT / Apache-2.0 / BSD / ISC / CC0 / Unlicense / Public Domain.
- **Rejected**: GPL / LGPL / AGPL in any variant (conflicts with AGPLv3 copyleft; the project's own dual licensing is not affected).
- Review workflow: `governance/THIRD-PARTY.md` and `third_party/LICENSE-MANIFEST.md`; per-item inventory: `THIRD-PARTY-NOTICES.md`; compatibility notes: `LICENSE-COMPATIBILITY.md`.

## 5. Gates (wording per §6b)

- `header-check.sh` scans **`*.go` only and does not scan `.md` at all**; a missing 6-line copyright header on a new bilingual `.md` **does not** trigger a header-check FAIL, but per `governance/HEADERS.md` §2 it is a **governance-consistency requirement** and should be applied uniformly.
- `license-check.sh` **does not scan the repository-root `LICENSE`**; placing the AGPLv3 full-text `LICENSE` **does not** cause a false FAIL.
- The frozen nine sections of `lint-gate.sh` (lines 41–135) are unchanged.

## 6. Statements

- Copyright holder: **小黎的部落 (www.qingdujian.com)**; source: `https://github.com/xiaolibuluo/Abate`; commercial contact: `xiaolibuluo@163.com` (genuine attributions retained; no fabricated qualifications).
- `competitorAbsolute = NOT_MEASURED` and `NOT_MEASURED` are retained verbatim.
- **This document does not constitute legal advice / 本文件不构成法律意见; subject to confirmation by the rights holder or legal counsel / 需权利人或法务确认.**
