<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->

> Language / 语言：[English](T08-contribute.en.md) | [中文](T08-contribute.md)

# T08 · How to contribute (DCO/CLA)

## 1 Audience

Contributors who want to submit a patch to Abate: the goal is to get the DCO
sign-off commit format right and to understand the dependency whitelist and
the relicensing gap as they currently stand.

## 2 Prerequisites

- `git` installed;
- Repository forked/cloned;
- Governance files: `CONTRIBUTING.md` · `DCO.md` · `CLA.md` (draft) ·
  `governance/DUAL-LICENSE.md`;
- **The `git commit` in this tutorial is a format example only; no git write
  operation is executed in this round.**

## 3 Commands

```bash
git clone https://github.com/xiaolibuluo/Abate.git
cd Abate
# (after editing files)
git add <your-files>
git commit -s -m "docs: add bilingual quickstart"
git log -1 --format=%B
```

## 4 Expected output

`git commit -s` appends a `Signed-off-by` line to the commit message; the
message shown by `git log -1 --format=%B` looks like:

```
docs: add bilingual quickstart

Signed-off-by: Your Name <your.email@example.com>
```

**Verification point**: the last line of the commit message is
`Signed-off-by: your real name <your real email>` (the DCO 1.1 declaration:
you have the right to submit the patch and agree to DCO 1.1). **No such line
= non-conforming commit**.

**Governance status (stated as-is; no legal conclusions drawn here)**:

- **DCO mandatory**: every commit must use `git commit -s`;
- **CLA optional**: currently optional; once signed it explicitly grants the
  project the right to use and relicense the contribution;
- **Relicensing gap**: **the DCO does not grant the project a relicensing
  right; under a dual license a CLA or explicit relicensing terms are needed —
  subject to confirmation by the rights holder or legal counsel**;
- **Dependency whitelist**: only MIT / Apache-2.0 / BSD / ISC / CC0 /
  Unlicense / Public Domain are allowed; GPL-family licenses (GPL-2.0 /
  GPL-3.0 / LGPL / AGPL, any variant) are rejected. Before importing a
  dependency, check `third_party/LICENSE-MANIFEST.md` and
  `LICENSE-COMPATIBILITY.md`.

## 5 Troubleshooting

| Symptom | Cause | Action |
|---|---|---|
| Commit lacks `Signed-off-by` | Forgot `-s` | Re-sign: `git commit --amend -s` (before pushing) |
| `-s` reports "unable to auto-detect email address" | git identity not configured | Set `git config user.name` / `git config user.email`, then retry |
| You want to import an LGPL library | The whitelist rejects GPL-family licenses | Do not import; if truly necessary, go through an ADR assessment, default is rejection |
| Unsure whether your contribution can be used commercially | Relicensing rights are a legal question | Subject to confirmation by the rights holder or legal counsel; this tutorial does not constitute legal advice |
