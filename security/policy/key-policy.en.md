<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->

> Language / 语言：[English](key-policy.en.md) | [中文](key-policy.md)

# Key Management Policy (XCU-SEC / T5)

English translation of `key-policy.md`, the section-by-section companion
under the plan-A side-by-side naming scheme; section order and tables
correspond one-to-one; no new conclusions are added; this document does not
constitute legal advice.

> Maps to PLAN-08 XCU-SEC §4.3 and threat T5 (signing-key compromise). This
> policy is process control, executed by the GPG-signing section of
> `scripts/release.sh`; one rotation rehearsal is completed before GA.

## 1. Key categories

| Key | Purpose | Custody |
|---|---|---|
| Release-signing master key | Signs release assets + manifest + tag | Offline hardware media (HSM / offline key card) |
| Emergency rotation key | Temporary signing when the master key fails | Separate offline media, physically isolated from the master key |
| Trust-anchor public key | User-side signature verification (fingerprint published) | Repository `SECURITY.md` + release announcements |

## 2. Offline custody

- Private keys **never** enter CI runners, images, or any networked host; only
  release jobs use them in an isolated environment.
- Build and signing are separated: CI produces assets; signing happens in an
  offline/controlled environment.
- Private-key passphrases are ≥ 30 characters, split across different
  custodians (Shamir optional).

## 3. Rotation process

1. The new key is generated offline; the trust-anchor public-key fingerprint
   is published in `SECURITY.md` and the release announcement.
2. **Overlap of 2 versions**: old and new keys coexist for verification; old
   assets remain verifiable by both keys.
3. After the overlap ends, the old key is removed from the trust-anchor list
   and its private-key medium destroyed.
4. Rotation events are recorded in `CHANGELOG.md` and the audit record.

## 4. Compromise response

- Suspected master-key compromise → immediately re-sign in-service assets
  with the emergency rotation key and announce.
- Revoke trust in the compromised key (remove it from the `SECURITY.md`
  trust anchor).
- Trigger an external security-audit re-review before GA.

## 5. Rehearsal record

- Before GA, complete one full key-generation → signing → verification →
  rotation-overlap → destruction rehearsal, recorded in the repository.
- `scripts/release.sh --demo-sign` provides a one-off-key signing-chain
  rehearsal (local evidence only, not published).
