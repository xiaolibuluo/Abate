<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->

> Language / 语言：[English](SECURITY.en.md) | [中文](SECURITY.md)

# Security Policy

Abate is a deterministic chaos-and-mutation testing engine. This document is
the public surface of the XCU-SEC security system: it publishes the signing
trust anchors, the vulnerability-disclosure process and the supported-version
window.

## Supported Versions

Only the current kernel year-ring (LTS) receives security fixes. Per the
"ten-year kernel" policy, a year-ring is maintained for 3 years.

| Version | Supported |
|---|---|
| v1.x (year-ring 1, LTS) | ✅ |
| < v1.0.0 (pre-GA) | ❌ |

## Reporting a Vulnerability

Please report security issues privately to the independent channel:

- **Email**: security@qingdujian.com (also: xiaolibuluo@163.com)
- Do **not** open a public issue for security reports.

Disclosure workflow (see `security/policy/disclosure.md` for the full text):

1. Acknowledgement within **48 hours** of a well-formed report.
2. **90-day** fix window for confirmed vulnerabilities.
3. CVE request for severity-class issues, then coordinated public advisory.
4. External third-party security audit before each GA.

## Trust Anchors (signing trust anchors)

Release assets and tags are signed with the project's offline-custodied GPG
key. Verify any download against the fingerprint below before trusting it.

> Trust-anchor public-key fingerprint (the fingerprint is updated with each
> release and published; during a rotation overlap two versions coexist for
> signature verification):
>
> | Purpose | Fingerprint (SHA-256, placeholder pending GA key generation) |
> |---|---|
> | Primary release-signing key | `TODO_GA_PRIMARY_FINGERPRINT` |
> | Emergency rotation key | `TODO_GA_SECONDARY_FINGERPRINT` |
>
> Verification:
>
> ```bash
> gpg --recv-keys <FINGERPRINT>
> gpg --verify abate-vX.Y.Z-linux-amd64.sig abate-vX.Y.Z-linux-amd64
> sha256sum -c manifest.txt   # the checksum manifest is covered by the GPG signature
> ```

## Threat Model Summary

The full threat model lives in `security/THREAT-MODEL.md`. Each threat (T1–T10)
has a named owner module and at least one negative test in `security/tests/`
that proves the mitigation intercepts the attack (signature rejection, CES
injection depth cap, report XSS auto-escaping, REST auth guard, archive
tamper rejection).

## Supply Chain

All third-party source is merged at source level (no runtime API/network
dependency) and pinned to an upstream commit, with a license review recorded in
`third_party/LICENSE-MANIFEST.md`. GPL-family licenses are rejected. The
dependency-vulnerability scan entry point is `security/audit/scan.sh`.
