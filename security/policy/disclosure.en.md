<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->

> Language / 语言：[English](disclosure.en.md) | [中文](disclosure.md)

# Vulnerability Disclosure Process (XCU-SEC / T4)

English translation of `disclosure.md`, the section-by-section companion
under the plan-A side-by-side naming scheme; section order and tables
correspond one-to-one; no new conclusions are added; this document does not
constitute legal advice.

> Maps to PLAN-08 XCU-SEC §4.1. This policy is published with the repository;
 reports are received through an independent security channel.

## 1. Reporting channel

- Security email: **security@qingdujian.com** (independent channel;
  xiaolibuluo@163.com also works).
- **Do not** report security vulnerabilities through public issues, to avoid
  0-day exposure.

## 2. Handling timeline

| Stage | Time limit | Action |
|---|---|---|
| Acknowledgement | **48 hours** | Confirm receipt and classify (critical/high/medium/low) |
| Fix Window | **90 days** | Complete the fix for critical severity and prepare the advisory |
| CVE request | After fix confirmation | For critical severity, request a CVE from MITRE |
| Coordinated Advisory | After fix + verification | Publish the security advisory (scope and mitigations) |

## 3. Severity definitions

- **Critical**: remote unauthenticated execution, signature-chain bypass,
  forged archive evidence.
- **High**: sandbox escape, key compromise, REST auth bypass.
- **Medium**: information leakage, known CVEs in dependencies.
- **Low**: documentation/configuration defects.

## 4. Advisories and trust anchor

- Advisories are published via GitHub Security Advisory plus an update to the
  repository's `SECURITY.md`.
- If signing-key rotation is involved, update the published trust-anchor
  fingerprint in sync.

## 5. Safe Harbor

Good-faith researchers who follow this process, do not exploit the
vulnerability, and do not disclose unpublished information are treated as
white-hat behavior and will not be pursued. We commit to working with
researchers rather than against them: fixing first, not assigning blame.
