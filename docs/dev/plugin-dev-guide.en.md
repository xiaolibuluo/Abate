<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->

> Language / 语言：[English](plugin-dev-guide.en.md) | [中文](plugin-dev-guide.md)

# Plugin Dev Guide

English translation of `plugin-dev-guide.md`, the section-by-section
companion under the plan-A side-by-side naming scheme; section order and
tables correspond one-to-one; no new conclusions are added; this document
does not constitute legal advice.

> Maps to PLAN-02 (WASM runtime) / PLAN-07 (control plane). Plugins are the
> vehicle for ecosystem extensions at hourly cadence: the whole flow from
> commit to publish takes ≤ 1 hour (see the `scripts/plugin-publish.sh`
> rehearsal; this is a process description, not a capability guarantee).

## 1. Plugin types

| Type | Purpose |
|---|---|
| `operator` | Perturbation operators (injection/mutation primitives) |
| `reporter` | Reporters (report formats/exports) |
| `adapter` | Adapters (CI / issue / messaging channels) |

## 2. Signature chain (threat T1)

- All plugins **must be signed**; unsigned loads are rejected
  (`plugin/signature.Verify`).
- The trust-anchor public-key fingerprint is published in `SECURITY.md`;
  signing keys are kept offline (see `security/policy/key-policy.md`).
- Release: the GPG section of `scripts/release.sh` signs assets + manifest
  (`--demo-sign` available for rehearsal).

## 3. Capability sandbox (threats T2/T3)

- Default is **deny-all**: host function calls are rejected unless explicitly
  granted (`plugin/capability`).
- Runtime gating + auditing; triple resource caps: fuel + memory + result
  length (`plugin/host`).
- Unauthorized calls are intercepted (covered by security test cases;
  executable T2/T3 cases are to be completed once the runtime host ships —
  stated as-is, not claimed as done).

## 4. Development flow

1. Define the interface with WIT (the host layer isolates WASI details; WIT
   is versioned with a compatibility matrix).
2. Implement the operator/reporter/adapter and compile to a WASI Preview 2
   component.
3. Sign locally: `scripts/plugin-publish.sh` (build → race → CLI tests → sign
   manifest).
4. `abate plugin load --manifest <m> --type <operator|reporter|adapter>` to
   verify loading.
5. `abate plugin call --id <id> --entry <entry> --args <base64>` to invoke.
6. Hot-swap: `abate plugin` hot-loads the new version and drains the old
   instance before unloading (running experiments are not interrupted).

## 5. Security self-checks

- Unsigned plugin loading must be rejected (T1).
- Capability escalation must be intercepted by the sandbox (T2).
- Resource exhaustion must be intercepted by the fuel/memory caps (T3).
- Signing-key compromise is mitigated by offline custody + a rotation
  process (T5).

## 6. Release cadence

- Plugin repositories use independent semver + a compatibility matrix
  (host version ranges).
- CI uses hourly release windows (ecosystem hourly cadence); decoupled from
  the kernel's yearly versioning.
