<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com)
  Project: xiaolibuluo/Abate — https://github.com/xiaolibuluo/Abate
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->
> Language / 语言：[English](THIRD-PARTY-NOTICES.en.md) | [中文](THIRD-PARTY-NOTICES.md)

# Third-Party Dependency License Notices (THIRD-PARTY-NOTICES)

> This document is the **per-item license inventory**; the governance-level integration register and review workflow are in `governance/THIRD-PARTY.md` (no conflict: the former covers workflow, the latter per-item attribution).
> **This document does not constitute legal advice / 本文件不构成法律意见; subject to confirmation by the rights holder or legal counsel.**

## 1. Scope of Registration

- Only third-party Go dependencies verified by the handoff and shipped in the repository's `vendor/` are registered; **no unverified components / versions / license texts are fabricated**.
- The single source of truth (SSOT) is each dependency's accompanying `LICENSE` text and `third_party/LICENSE-MANIFEST.md`; the license column below reflects existing `scripts/license-check.sh` determinations (ALLOW / MIT). The concrete license text **follows the SSOT** and is **subject to confirmation by the rights holder or legal counsel**.

## 2. Per-Item Registration

| Component | Version | Purpose | License (license-check determination) | Integration |
|---|---|---|---|---|
| wazero | v1.12.0 | Real WASM runtime | ALLOW (within allowlist; concrete license text per SSOT) | vendoring / module source merge |
| golang.org/x/sys | v0.44.0 | System-call primitives | ALLOW (within allowlist; concrete license text per SSOT) | Go module (source readable) |
| gopkg.in/yaml.v3 | v3.0.1 | YAML parsing | MIT | Go module (source readable) |

## 3. Import Dependency Allowlist

- **Allowed**: MIT / Apache-2.0 / BSD / ISC / CC0 / Unlicense / Public Domain.
- **Rejected**: GPL / LGPL / AGPL in any variant (the project's own outward dual licensing AGPLv3 + commercial is not affected).
- Review coverage: each entry must have an explicit "allowed / rejected / pending review" conclusion; GPL-family hits are rejected; unknown licenses are rejected by default, not silently allowed.

## 4. Cross-Links

- `governance/THIRD-PARTY.md` — five-step integration method, iron rules, governance index.
- `NOTICE` — copyright and attribution notice (single-file bilingual).
- `LICENSE-COMPATIBILITY.md` — dependency allowlist notes in the dual-licensing context.
- `LICENSING.md` — dual-licensing two-path overview.

## 5. Statements

- Copyright holder: **小黎的部落 (www.qingdujian.com)**; source: `https://github.com/xiaolibuluo/Abate` (genuine attributions retained).
- Competing licenses are distinguished **by objective nature only**; no benchmarking-surpass claims; `competitorAbsolute = NOT_MEASURED` and `NOT_MEASURED` are retained verbatim.
- SPDX `AGPL-3.0-or-later` (**subject to confirmation by the rights holder / 需权利人确认**).
- **This document does not constitute legal advice / 本文件不构成法律意见; subject to confirmation by the rights holder or legal counsel / 需权利人或法务确认.**
