<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->

> Language / 语言：[English](T10-faq-gates.en.md) | [中文](T10-faq-gates.md)

# T10 · FAQ map + how to run the gates

## 1 Audience

Newly onboarded maintainers and local environment builders: the goal is to
know where each FAQ answer lives (navigation) and to run the header and
license script gates locally.

## 2 Prerequisites

- Build completed per T01 (`make test` etc. need the full toolchain);
- Run from the repository root;
- FAQ entry: `FAQ-licensing.md` (license decision tree + 7 questions); gate
  details: `README.md` §7 and `governance/HEADERS.md`.

## 3 Commands

```bash
bash scripts/header-check.sh
bash scripts/license-check.sh
make gate
make test
make build
```

## 4 Expected output

`bash scripts/header-check.sh` (measured, verbatim):

```
[PASS] 全部源文件版权头齐全（含小黎的部落水印）

== HEADER GATE: PASSED ==
```

`bash scripts/license-check.sh` (measured, verbatim; the WARN is expected and
does not FAIL):

```
[PASS] 全部导入依赖许可证合规（无 GPL）
[WARN] LICENSE-MANIFEST 含 GPL 字样，请确认仅指代项目自身双授权而非导入依赖。
[PASS] third_party/LICENSE-MANIFEST.md 存在且已登记

== LICENSE GATE: PASSED - 无 GPL 系依赖 ==
```

**Verification points**: both scripts exit 0 and end with `PASSED` (the
pre-change baseline was already PASSED; re-verified by measurement while
writing this tutorial).

**Exact wording of gate boundaries (do not misread)**:

- `header-check.sh` **scans only `*.go`, never `.md`**: a newly added
  bilingual `.md` without a header will **not** trigger a FAIL from this gate;
  uniform headers are a governance-consistency requirement
  (`governance/HEADERS.md` §2), **not a gate-enforced rule**;
- The GPL scan of `license-check.sh` targets the license texts of
  `vendor/modules.txt` dependencies and `third_party/LICENSE-MANIFEST.md`
  only; it **does not scan the repository-root `LICENSE`**: placing the
  verbatim AGPLv3 text at the root will not cause a false FAIL (the GPL
  mentions in LICENSE-MANIFEST are the expected WARN for the project's own
  dual license);
- `make gate` (`scripts/lint-gate.sh`) is the full five-stage gate chain
  (gofmt → go vet → golangci-lint → go test -race → go build); lines 41–135
  are the frozen nine sections and are not modified; `golangci-lint` must be
  installed locally (use whatever version is installed);
- `make test` / `make build` map to `go test -race -count=1 ./...` and
  `go build ./...`; output varies by environment — read "exit 0 = pass".

**FAQ navigation (7-question map)**: AGPLv3 or commercial / what is §13 /
what is SPDX and which one this project uses / why THIRD-PARTY-NOTICES and
NOTICE / is the DCO enough (relicensing gap pending confirmation) / how to
view competitor licenses / will the root LICENSE and new .md files trigger a
gate FAIL — all in `FAQ-licensing.md`.

## 5 Troubleshooting

| Symptom | Cause | Action |
|---|---|---|
| HEADER GATE: FAILED | A new/modified `.go` file is missing its copyright header, or it is not within the first 8 lines | Restore it byte for byte per the template in `governance/HEADERS.md` §1 |
| LICENSE GATE: FAILED | A `vendor/modules.txt` dependency license text hits the GPL family | Do not import that dependency; process in `governance/THIRD-PARTY.md` |
| `.md` missing a header but the gate PASSED | Expected: header-check does not scan `.md` | For governance consistency, add the 6-line HTML template header — optional, not enforced |
| `make gate` reports golangci-lint missing | Not installed locally | Install before `make lint` (path under `go env GOPATH`/bin); use whatever version is installed |
| WARN mistaken for FAIL | WARN does not change the exit code | Judge by the trailing `PASSED/FAILED` line and the exit code |
