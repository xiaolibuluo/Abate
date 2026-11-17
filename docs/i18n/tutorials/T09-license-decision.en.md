<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->

> Language / 语言：[English](T09-license-decision.en.md) | [中文](T09-license-decision.md)

# T09 · AGPLv3 or commercial license?

## 1 Audience

Technical decision-makers evaluating Abate for their own products: the goal
is to complete a license-path self-check in 10 minutes and know which file to
read and whom to ask.

## 2 Prerequisites

- No binary build strictly required, but running `abate version` once shows
  the license declaration directly;
- Governance files: `LICENSING.md` · `FAQ-licensing.md` (decision tree) ·
  `LICENSE` (verbatim official AGPLv3 text) · `COMMERCIAL-LICENSE.md`;
- **This tutorial does not constitute legal advice**; every individual
  legality judgment is subject to confirmation by the rights holder or legal
  counsel.

## 3 Commands

```bash
/tmp/abate_verify version
sed -n '536,544p' LICENSE
```

## 4 Expected output

The license line of the `version` output (one of the 6 provenance lines,
measured):

```
license: AGPLv3 or commercial
```

`sed -n '536,544p' LICENSE` lands on **§13 Remote Network Interaction; Use
with the GNU General Public License** of the official AGPLv3 text (the section
heading sits at line 540 of LICENSE):

```
  13. Notwithstanding any other provision of this License, if you modify the
Program, your modified version must prominently offer all users interacting
with it remotely through a computer network (if your version supports such
interaction) an opportunity to receive the Corresponding Source of your
modified version by providing access to the Corresponding Source through a
network server at no charge, through some standard or customary means of
facilitating copying of software.
```

(Excerpt from the official text; the complete obligation is governed by the
verbatim full text of the repository-root `LICENSE`.)

**Decision tree (informational navigation only; not legal advice)**:

1. **Can you fulfill all AGPLv3 obligations** (including §13: offering the
   Corresponding Source to users interacting remotely over a network)?
   - Yes → take the AGPLv3 path (`LICENSE` / `CONTRIBUTING.md` / `DCO.md`);
   - No → go to step 2.
2. **Closed-source integration / unwilling to publish derived-work source /
   unwilling to take on the §13 source-offering obligation**?
   - Yes → a commercial license is needed; contact `xiaolibuluo@163.com`
     (the existing contact, carried over as-is; applicability is subject to a
     separate written agreement between the parties).
3. **SPDX: only or or-later** → uniformly noted as `AGPL-3.0-or-later`
   (**subject to confirmation by the rights holder; not presented as settled**).
4. **How to view competitor licenses** → objective nature-differentiation
   only: Qt / Grafana / iText as dual-licensing references; MongoDB SSPL /
   Elastic / HashiCorp BUSL are **not OSI-approved open-source licenses**; no
   benchmarking or surpassing claims.

**Verification point**: you can say whether your scenario falls under step 1
("can fulfill") or step 2 ("needs a commercial license"), and which file to
read next.

## 5 Troubleshooting

| Symptom | Cause | Action |
|---|---|---|
| Confusion about "using AGPLv3 doesn't forbid commercial use" | Dual licensing = two optional paths, not a mutual gate | Read the two-path comparison table in `LICENSING.md` |
| §13 obligations feel uncertain | It is a legal question | Subject to confirmation by the rights holder or legal counsel; read only the official text, not second-hand summaries |
| No reply to a commercial inquiry | This document set contains no pricing/signed-deals/customer cases/SLA | Email `xiaolibuluo@163.com` directly |
| Any "is what I'm doing legal" question | This tutorial navigates but does not adjudicate | The uniform answer: subject to confirmation by the rights holder or legal counsel; this document does not constitute legal advice |
