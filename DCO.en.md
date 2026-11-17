<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com)
  Project: xiaolibuluo/Abate — https://github.com/xiaolibuluo/Abate
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->
> Language / 语言：[English](DCO.en.md) | [中文](DCO.md)

# Developer Certificate of Origin (DCO 1.1)

> **This document does not constitute legal advice / 本文件不构成法律意见; subject to confirmation by the rights holder or legal counsel.**
> Cross-linked with `CONTRIBUTING.md` and `CLA.md`; the appendix is the standard official English DCO 1.1 full text (verbatim, identical in both language versions).

## 1. Status and Usage

- Current governance is "**CLA optional, DCO mandatory**"; this document **does not change that status and does not make the CLA mandatory**.
- **DCO is mandatory**: every commit must contain a `Signed-off-by` line certifying that you have the right to submit the contribution and agree to DCO 1.1.
- Documentation example (**example only; no git commit is executed in this round**):

```bash
git commit -s -m "docs: add bilingual quickstart"
```

- `git commit -s` appends a line of the form `Signed-off-by: <Name> <email>`.

## 2. Relicensing Gap under Dual Licensing (subject to confirmation by the rights holder or legal counsel)

- **The DCO does not grant the project a relicensing right; under dual licensing (AGPLv3 + commercial) a CLA or an explicit relicensing clause is needed — subject to confirmation by the rights holder or legal counsel**.
- This gap is not "already resolved" by this document; `CLA.md` is a **draft** and **whether it is required is subject to confirmation by the rights holder or legal counsel**; this document does not decide the matter.

## 3. Import Dependency Allowlist

- **Allowed**: MIT / Apache-2.0 / BSD / ISC / CC0 / Unlicense / Public Domain.
- **Rejected**: GPL / LGPL / AGPL in any variant.
- Review workflow: `governance/THIRD-PARTY.md`; inventory: `THIRD-PARTY-NOTICES.md`; compatibility: `LICENSE-COMPATIBILITY.md`.

## 4. Gates (wording per §6b)

- `header-check.sh` scans **`*.go` only and does not scan `.md` at all**; `license-check.sh` **does not scan the repository-root `LICENSE`**; a 6-line copyright header on new bilingual `.md` files is a **governance-consistency requirement, not a gate-enforced item**.

## 5. Statements

- Copyright holder: **小黎的部落 (www.qingdujian.com)**; source: `https://github.com/xiaolibuluo/Abate` (genuine attributions retained).
- `competitorAbsolute = NOT_MEASURED` and `NOT_MEASURED` are retained verbatim.
- **This document does not constitute legal advice / 本文件不构成法律意见; subject to confirmation by the rights holder or legal counsel / 需权利人或法务确认.**

## Appendix · Developer Certificate of Origin 1.1 (official English full text, verbatim)

> The following is the standard English DCO 1.1 full text, identical in both language versions and not translated; if it differs from the officially published text, the official text prevails — **subject to confirmation by the rights holder or legal counsel**.

```text
Developer Certificate of Origin
Version 1.1

Copyright (C) 2004, 2006 The Linux Foundation and its contributors.
1 Letterman Drive
Suite D4700
San Francisco, CA, 94129

Everyone is permitted to copy and distribute verbatim copies of this
license document, but changing it is not allowed.


Developer's Certificate of Origin 1.1

By making a contribution to this project, I certify that:

(a) The contribution was created in whole or in part by me and I
    have the right to submit it under the open source license
    indicated in the file; or

(b) The contribution is based upon previous work that, to the best
    of my knowledge, is covered under an appropriate open source
    license and I have the right under that license to submit that
    work with modifications, whether created in whole or in part
    by me, under the same open source license (unless I am
    permitted to submit under a different license), as indicated
    in the file; or

(c) The contribution was provided directly to me by some other
    person who certified (a), (b) or (c) and I have not modified
    it.

(d) I understand and agree that this project and the contribution
    are public and that a record of the contribution (including all
    personal information I submit with it, including my sign-off) is
    maintained indefinitely and may be redistributed consistent with
    this project or the open source license(s) involved.
```
