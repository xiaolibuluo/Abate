<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->

> Language / 语言：[English](THREAT-MODEL.en.md) | [中文](THREAT-MODEL.md)

# Abate Threat Model (XCU-SEC)

English translation of `THREAT-MODEL.md`, the section-by-section companion
under the plan-A side-by-side naming scheme; section order and tables
correspond one-to-one; no new conclusions are added; this document does not
constitute legal advice.

> This document is the first full threat model of XCU-SEC, mapping to PLAN-08
> §3. Every threat must point to its owning module and executable test cases
> (alignment verified by scan check, 10/10, see "Alignment check" below). The
> test set lives in `security/tests/`, resident-green in CI.

## Threat coverage matrix

| ID | Threat | Attack surface | Mitigation | Owning module | Executable case |
|---|---|---|---|---|---|
| T1 | Malicious plugin loading | Plugin install channel | Mandatory signature verification (unsigned rejected) | `plugin/signature` | ✅ `TestThreatT1SignatureRejection` |
| T2 | Plugin privilege escalation (fs/net/privilege) | Host function calls | Capability sandbox default deny-all + runtime gating + audit | `plugin/capability` + `plugin/host` | ✅ `TestThreatT2CapabilityEscalation` |
| T3 | Plugin resource exhaustion | Plugin execution | 1 MiB result cap of the native Go host | `plugin/host` (WAS-HOST) | ✅ `TestThreatT3ResourceExhaustion` |
| T4 | Supply-chain poisoning | go.mod / third_party | Source-level inclusion + commit-hash pinning + license review | XCU-GOV | ✅ `security/audit/scan.sh` (version pinning + license review) |
| T5 | Signing-key compromise | Release channel | Offline key custody + rotation process + overlap window | XCU-REL | ✅ `security/policy/key-policy.md` (process rehearsal) |
| T6 | CES document injection | Document parsing | Depth limit 64 + input limits + fuzzing | `ces/parse` | ✅ `TestThreatT6ParseInjection` |
| T7 | Report XSS | Report output | html/template auto-escaping + template-variable whitelist | `report` / `ctl/web` | ✅ `TestThreatT7ReportXSS` |
| T8 | REST unauthorized access | API listener | Mandatory Bearer auth + remote listener refuses to start without a token | `ctl/rest` | ✅ `TestThreatT8UnauthorizedAccess` |
| T9 | Token/secret leakage into logs | Global logging | Redaction middleware + REST log-point integration + scan gate | `security/redact` + `ctl/rest` | ✅ `TestThreatT9SecretRedaction` |
| T10 | Archive tampering | Archive files | Frame hash chain + whole-archive SHA-256 | `sim/record` | ✅ `TestThreatT10ArchiveTamper` |

## Alignment check

- Delivered executable attack cases: T1 / T2 / T3 / T6 / T7 / T8 / T9 / T10,
  all passing with `-race`; T2/T3 go through the `plugin/host` native Go call
  chain; T9 is verified on both channels via the `security/redact` pure
  functions and the two log points in `ctl/rest`.
- Process/document controls: T4 (supply-chain scan script) and T5 (key
  policy) are delivered as executable/document controls.
- Threat-coverage alignment: 10/10 threats have a corresponding executable
  case or a delivered process/document control; the coverage matrix in
  `security/tests/doc.go` matches this section.
- Delivery boundary (kept as-is): the gVisor platform-level sandbox (ADR-001)
  and a true WASM runtime were not vendored in the air-gapped offline
  environment and are not part of this milestone's deliverables. The current
  T2/T3 mitigations are carried by the native Go implementation
  (`plugin/capability` default deny-all + `AuditLog` gating; `plugin/host`
  capability gating + the 1 MiB `MaxResultBytes` result cap;
  `plugin/host/hostfunc.go` explicitly implements host functions with
  SIM-CLK/SIM-RNG rather than a WASM runtime); T9 is implemented by the
  `security/redact` redaction middleware and wired into the panic and access
  `log.Printf` call sites of `ctl/rest/middleware.go`. The platform-level
  isolation and the true WASM runtime remain undelivered boundaries and must
  not be described as delivered.

## Design principles

1. Defense in depth: signature chain + capability sandbox + input limits +
   output escaping + authentication + tamper-proof evidence, layered
   blocking.
2. Default deny: the capability sandbox denies all by default; a remote
   listener refuses to start without a token by default.
3. Non-repudiable evidence: recorded-replay frame hash chain + whole-archive
   SHA-256; tampering is rejected.
4. Offline first: the kernel has zero network dependencies; release-signing
   keys are kept offline; the supply chain is included at source level.

## Audit cadence

At the end of each milestone (M1–M6), run the full security test set (≥1
attack case per threat T1–T10); before GA, complete a third-party external
security audit (report filed in the repository).
