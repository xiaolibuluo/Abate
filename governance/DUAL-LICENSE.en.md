<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->

> Language / 语言：[English](DUAL-LICENSE.en.md) | [中文](DUAL-LICENSE.md)

# Dual-Licensing Statement (XCU-GOV)

> Maps to PLAN-08 XCU-GOV §6. Abate is dual-licensed under **AGPLv3
> (open-source side) + a commercial license (closed-source side)**.
> This document does not constitute legal advice.

## 1. Why dual licensing

- **AGPLv3 (open-source side)**: strong copyleft for network/SaaS scenarios,
  protecting community and user freedoms and preventing a SaaS loophole.
  Suited to self-hosting, research, and contribution.
- **Commercial license (closed-source side)**: permits closed-source
  integrations that do not publish derived code, lowering the enterprise
  compliance barrier. Contact **xiaolibuluo@163.com** for a commercial license.

The two licenses coexist: you may either accept the AGPLv3 obligations or
obtain a commercial license that exempts you from its copyleft obligations.

## 2. Contributor agreements

- **CLA (Contributor License Agreement)**: optional. Once signed, it explicitly
  grants the project the right to use and relicense the contribution.
- **DCO (Developer Certificate of Origin)**: mandatory. Every commit must carry
  a `Signed-off-by` line (`git commit -s`) declaring that the contributor has
  the right to submit and agrees to DCO 1.1.
- Contributing grants dual-license rights: the project may use your
  contribution under both AGPLv3 and the commercial license.
- Gap statement (stated as-is, not resolved here): **the DCO does not grant
  the project a relicensing (relicense) right; under a dual license a CLA or
  explicit relicensing terms are needed — subject to confirmation by the
  rights holder or legal counsel.** This document does not change CLA's
  current optional status, nor does it constitute legal advice.

## 3. License compatibility (imported dependencies)

- Allowed only: MIT / Apache-2.0 / BSD / ISC / CC0 / Unlicense / Public Domain.
- Rejected: GPL-2.0 / GPL-3.0 / LGPL / AGPL (any variant) — conflicts with
  AGPLv3's copyleft.
- The review process is described in `governance/THIRD-PARTY.md` and
  `third_party/LICENSE-MANIFEST.md`; see also the root
  `THIRD-PARTY-NOTICES.md` and `LICENSE-COMPATIBILITY.md`.

## 4. Trademark and watermark

- 『小黎的部落』 is the project's digital watermark and brand identifier; it
  does not constitute a standalone trademark grant.
- The copyright-header watermark is distributed with the source code; see
  `governance/HEADERS.md`. Whether a formal trademark policy exists is
  subject to confirmation by the rights holder or legal counsel (see the root
  `TRADEMARK.md`).
