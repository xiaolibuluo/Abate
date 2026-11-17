<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->

> Language / 语言：[English](THIRD-PARTY.en.md) | [中文](THIRD-PARTY.md)

# Third-Party Source Inclusion Register (XCU-GOV)

> The governance-level master register for third-party inclusion. The
> **single source of truth (SSOT)** for per-entry registration is
> `third_party/LICENSE-MANIFEST.md` (the operational register, checked in CI by
> `scripts/license-check.sh` and `security/audit/scan.sh`). This document is
> its governance index and process description. For a user-facing dependency
> notice, see also the root `THIRD-PARTY-NOTICES.md`.

## 1. Inclusion rules (project red lines)

1. **Source-code inclusion, no API-only integration**: always vendoring /
   module source inclusion; out-of-process calls, gRPC, network dependencies,
   or runtime dynamic fetching of external engines are forbidden.
2. **Must be Go**: non-Go implementations (Rust/C-family) are not admitted;
   gVisor is a Go implementation, and exceptions are limited to Go subsets
   explicitly approved by an ADR.
3. **No GPL-family licenses**: GPL-2.0 / GPL-3.0 / LGPL / AGPL in any variant
   are rejected (the project's own AGPLv3+commercial dual license is not
   affected by this rule). Allowed: MIT / Apache-2.0 / BSD / ISC / CC0 /
   Unlicense / Public Domain.
4. **Register the source commit hash**: vendoring must pin the upstream commit
   to enable security-patch backporting and traceability.

## 2. The five-step inclusion process

1. Assess: license compatibility + language + functional necessity (decision
   recorded in an ADR).
2. Register: add an entry to `third_party/LICENSE-MANIFEST.md`
   (component/source/version or commit/license/language/integration
   method/milestone/status/review verdict).
3. Include: vendoring or module source inclusion, pinning the commit;
   preserve upstream copyright headers.
4. Trim: remove unrelated files and record the trim list; no non-Go
   dependencies introduced.
5. Backport: upstream security patches are reviewed and backported, recorded
   in the audit trail (aligned with the KRN-INJ upstream backport mechanism).

## 3. Current register index (points to the SSOT)

| Component | Purpose | License | Integration point | Status |
|---|---|---|---|---|
| gVisor (isolated kernel subset) | Sandbox isolation (M4) | Apache-2.0 | vendor → five-step inclusion | Decision adopted, not yet vendored |
| yaml.v3 | YAML parsing | MIT | Go module (source readable) | Planned |
| SplitMix64 | PRNG algorithm | Public Domain | Algorithm source included in sim/rand | Planned |

> The actual registrations and statuses are governed by
> `third_party/LICENSE-MANIFEST.md`; this index is a governance overview only.

## 4. Review verdicts

- Every registration must carry an explicit license review verdict
  (allowed / rejected / pending review).
- A GPL-family hit means rejection, zero exceptions (other than the project's
  own outward dual license).
- Unknown licenses default to rejection; never granted by default.
