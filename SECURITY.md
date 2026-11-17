<!--
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
-->

# Security Policy

Abate is a deterministic chaos-and-mutation testing engine. This document is the
public surface of the XCU-SEC security体系：it publishes the signing trust
anchors, the vulnerability-disclosure process and the supported-version window.

## Supported Versions

Only the current kernel year-ring (LTS) receives security fixes. Per the
"十年内核" policy, a year-ring is maintained for 3 years.

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

## Trust Anchors (签名信任锚)

Release assets and tags are signed with the project's offline-custodied GPG
key. Verify any download against the fingerprint below before trusting it.

> 信任锚公钥指纹（指纹随版本更新并公示；轮换重叠期 2 个版本并存验签）：
>
> | 用途 | 指纹 (SHA-256, 占位待 GA 密钥生成后填入) |
> |---|---|
> | 发布签名主密钥 | `TODO_GA_PRIMARY_FINGERPRINT` |
> | 应急轮换密钥 | `TODO_GA_SECONDARY_FINGERPRINT` |
>
> Verification:
>
> ```bash
> gpg --recv-keys <FINGERPRINT>
> gpg --verify abate-vX.Y.Z-linux-amd64.sig abate-vX.Y.Z-linux-amd64
> sha256sum -c manifest.txt   # 校验和清单由 GPG 签名覆盖
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
