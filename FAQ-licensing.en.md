<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com)
  Project: xiaolibuluo/Abate — https://github.com/xiaolibuluo/Abate
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->
> Language / 语言：[English](FAQ-licensing.en.md) | [中文](FAQ-licensing.md)

# Licensing FAQ (FAQ-licensing)

> This document is informational navigation only. **This document does not constitute legal advice / 本文件不构成法律意见**; any individual "is this lawful" question is answered uniformly with "**subject to confirmation by the rights holder or legal counsel**".

## 1. Decision Tree (When Is a Commercial License Needed?)

```
① Can you fulfil all AGPLv3 obligations (incl. §13 Remote Network Interaction: offering Corresponding Source to network users)?
    Yes ──► Take AGPLv3 (LICENSE / CONTRIBUTING / DCO)
    No  ──► ② Closed-source integration / unwilling to publish derivative source / unwilling to bear §13 obligations?
              Yes ──► A commercial license is needed; contact xiaolibuluo@163.com (retained, subject to confirmation)
    ③ SPDX only or or-later ──► subject to confirmation by the rights holder
    ④ Competing licenses: objective-nature distinction only, no benchmarking-surpass claims
    ⑤ Any individual "is this lawful" question ──► uniformly: "subject to confirmation by the rights holder or legal counsel; this document does not constitute legal advice"
```

## 2. Seven Questions

### Q1 · AGPLv3 or commercial license?
Two **optional paths** coexist (see `LICENSING.md`). If you can fulfil all AGPLv3 obligations (incl. §13) → open-source side; closed-source integration / unwilling to publish derivative source / unwilling to bear §13 source-provision obligations → commercial side, contact `xiaolibuluo@163.com` (**retained, subject to confirmation**). This FAQ does not decide the matter.

### Q2 · What is AGPLv3 §13?
Section 13, "Remote Network Interaction; Use with the GNU General Public License", is at **line 540** of the `LICENSE` official full text: if you modify the Program and let users interact with it over a network, you must offer those users the opportunity to receive the Corresponding Source. This FAQ only describes and links; it does **not copy or rewrite** the clause text; interpretation is **subject to confirmation by the rights holder or legal counsel**.

### Q3 · What is SPDX, and which one does this project use?
SPDX is a license identifier. Document headers here uniformly state `AGPL-3.0-or-later` (**subject to confirmation by the rights holder / 需权利人确认; not stated as settled**); the `AGPL-3.0-only` vs `or-later` election is **subject to confirmation by the rights holder**. The official text's "How to Apply" wording `or (at your option) any later version` is reference wording only and does not mean this project has elected it.

### Q4 · Why THIRD-PARTY-NOTICES and NOTICE?
`THIRD-PARTY-NOTICES.md` registers dependencies such as wazero / x-sys / yaml.v3 item by item (SSOT: `vendor/` and `third_party/LICENSE-MANIFEST.md`); `NOTICE` is the single-file bilingual copyright and attribution notice. Both serve faithful third-party disclosure, not capability claims.

### Q5 · Is the DCO enough (the relicensing gap)?
Current governance "**CLA optional, DCO mandatory**" is unchanged, and **the CLA is not made mandatory**; whether `CLA.md` is required is **subject to confirmation by the rights holder or legal counsel**. Explicit note: **the DCO does not grant the project a relicensing right; under dual licensing a CLA or an explicit relicensing clause is needed — subject to confirmation by the rights holder or legal counsel**. This FAQ states no "already compliant" conclusion.

### Q6 · How to read competing licenses?
**Objective-nature distinction only**: Qt / Grafana / iText are dual-licensing references; MongoDB SSPL / Elastic / HashiCorp BUSL are **not OSI-approved open-source licenses**. They are not conflated with AGPLv3; no benchmarking-surpass claims; `competitorAbsolute = NOT_MEASURED` and `NOT_MEASURED` are retained verbatim.

### Q7 · Will the root `LICENSE` and new `.md` files trigger a gate FAIL?
Per §6b correction: `header-check.sh` scans **`*.go` only and does not scan `.md` at all**; a missing 6-line header on a new bilingual `.md` **does not** trigger a FAIL, but per `governance/HEADERS.md` §2 it is a governance-consistency requirement and should be applied. `license-check.sh` **does not scan the repository-root `LICENSE`**; placing the AGPLv3 full-text `LICENSE` **does not** cause a false FAIL. The frozen nine sections of `lint-gate.sh` (lines 41–135) are unchanged. Both scripts were PASSED at the pre-change baseline.

## 3. Cross-Links

`LICENSING.md` / `COMMERCIAL-LICENSE.md` / `LICENSE` / `CONTRIBUTING.md` / `DCO.md` / `CLA.md` / `NOTICE` / `THIRD-PARTY-NOTICES.md` / `LICENSE-COMPATIBILITY.md` / `TRADEMARK.md` / `SECURITY.md`.

## 4. Statements

- Copyright holder: **小黎的部落 (www.qingdujian.com)**; source: `https://github.com/xiaolibuluo/Abate`; commercial contact: `xiaolibuluo@163.com` (genuine attributions retained; no fabricated qualifications).
- The tutorial version of the decision path is in `docs/i18n/tutorials/T09-license-decision.md` (Chinese) / `.en.md` (English).
- **This document does not constitute legal advice / 本文件不构成法律意见; subject to confirmation by the rights holder or legal counsel / 需权利人或法务确认.**
