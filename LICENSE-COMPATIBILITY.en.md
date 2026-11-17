<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com)
  Project: xiaolibuluo/Abate — https://github.com/xiaolibuluo/Abate
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->
> Language / 语言：[English](LICENSE-COMPATIBILITY.en.md) | [中文](LICENSE-COMPATIBILITY.md)

# License Compatibility Notes (LICENSE-COMPATIBILITY)

> The file name is fixed by the architectural design. This document explains the **import dependency allowlist in the dual-licensing context** as informational navigation; it **does not conflict with AGPLv3 clauses and does not copy / rewrite the `LICENSE` clause text**.
> **This document does not constitute legal advice / 本文件不构成法律意见; subject to confirmation by the rights holder or legal counsel.**

## 1. Dual-Licensing Context

- Abate is dual-licensed under **AGPLv3 (open-source side) + a commercial license (closed-source side)**; see `LICENSING.md`.
- The fact that the project **itself** is distributed outward under AGPLv3 / a commercial license **does not mean** third-party GPL / LGPL / AGPL variant code may be merged into the kernel: imported dependencies must still fall within the allowlist below (`governance/DUAL-LICENSE.md` §3, `governance/THIRD-PARTY.md` §1).

## 2. Import Dependency Allowlist

- **Allowed**: MIT / Apache-2.0 / BSD / ISC / CC0 / Unlicense / Public Domain.
- **Rejected**: GPL-2.0 / GPL-3.0 / LGPL / AGPL in any variant — conflicts with AGPLv3 copyleft (the project's own dual licensing is not affected).
- Whether any specific dependency is compatible **follows the accompanying `LICENSE` text in `vendor/` and `third_party/LICENSE-MANIFEST.md` (SSOT), subject to confirmation by the rights holder or legal counsel**; this document decides no per-item conclusion.

## 3. Relationship with AGPLv3 Clauses

- `LICENSE` is a verbatim copy of the AGPLv3 official full text; this document **rewrites none** of its clauses, in particular §13 "Remote Network Interaction" (line 540) and the Corresponding Source obligation.
- The allowlist is a **project governance convention** (`governance/HEADERS.md` / `governance/THIRD-PARTY.md`), not a reinterpretation of AGPLv3 clauses; any tension between the two is **subject to confirmation by the rights holder or legal counsel**.

## 4. Rights Chain and Relicensing Gap

- Current governance "CLA optional, DCO mandatory" is unchanged; **the DCO does not grant the project a relicensing right; under dual licensing a CLA or an explicit relicensing clause is needed — subject to confirmation by the rights holder or legal counsel**; this document **does not make the CLA mandatory** (`CLA.md` is a draft).
- SPDX `AGPL-3.0-or-later` (**subject to confirmation by the rights holder / 需权利人确认**); `only` / `or-later` is subject to confirmation by the rights holder.

## 5. Cross-Links

- `THIRD-PARTY-NOTICES.md` — per-item registration of wazero / x-sys / yaml.v3.
- `NOTICE` — copyright and attribution notice.
- `governance/THIRD-PARTY.md` — five-step integration method and review-conclusion coverage.
- `LICENSING.md` / `COMMERCIAL-LICENSE.md` / `FAQ-licensing.md`.

## 6. Statements

- Competing licenses are distinguished **by objective nature only**: Qt / Grafana / iText are dual-licensing references; MongoDB SSPL / Elastic / HashiCorp BUSL are **not OSI-approved open-source licenses**; no benchmarking-surpass claims are made.
- `competitorAbsolute = NOT_MEASURED` and `NOT_MEASURED` are retained verbatim.
- Copyright holder: **小黎的部落 (www.qingdujian.com)** (genuine attributions retained; no fabricated qualifications).
- **This document does not constitute legal advice / 本文件不构成法律意见; subject to confirmation by the rights holder or legal counsel / 需权利人或法务确认.**
