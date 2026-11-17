<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->

> Language / 语言：[English](README.en.md) | [中文](README.md)

# Abate

Abate is a deterministic experiment control plane: chaos experiments are
described in CES (`ces.tencent.com/v1`), then driven through the AIR frontend,
a deterministic kernel and SIM-REC archives, producing run reports and evidence
chains that are byte-for-byte reproducible under a fixed seed.

## 1 License

Abate is dual-licensed under **AGPLv3 (open source) and a commercial license**:

- Open-source path: `LICENSE` (the verbatim official AGPLv3 text), SPDX
  identifier `AGPL-3.0-or-later` (**subject to confirmation by the rights
  holder; this document does not constitute legal advice**);
- Commercial path: `COMMERCIAL-LICENSE.md` (contact `xiaolibuluo@163.com`);
- Dual-license overview and the two-path comparison table: `LICENSING.md`;
- Decision aid: `FAQ-licensing.md` (informational navigation only).

## 2 Requirements

- Go **1.26.11** (the real build version; the older `go1.27` / `Go 1.24+`
  values in historical documents are outdated)
- Offline build: `GOFLAGS=-mod=vendor` (dependencies are vendored; no git remote)
- Optional: `golangci-lint` (use whatever version is installed locally; this
  document does not assert any unverified version number)

## 3 Quick start

```bash
export GOFLAGS=-mod=vendor
go build -o abate .
./abate version
```

Step-by-step tutorials:

- T01 Offline install and build from scratch: `docs/i18n/tutorials/T01-build.en.md`
- T02 Your first chaos experiment: `docs/i18n/tutorials/T02-first-run.en.md`

## 4 Tutorials

| Tutorial | English | 中文 |
|---|---|---|
| T01 Offline install and build from scratch | [T01-build.en.md](docs/i18n/tutorials/T01-build.en.md) | [T01-build.md](docs/i18n/tutorials/T01-build.md) |
| T02 Your first chaos experiment | [T02-first-run.en.md](docs/i18n/tutorials/T02-first-run.en.md) | [T02-first-run.md](docs/i18n/tutorials/T02-first-run.md) |
| T03 Dry-run validation of a CES | [T03-dry-run.en.md](docs/i18n/tutorials/T03-dry-run.en.md) | [T03-dry-run.md](docs/i18n/tutorials/T03-dry-run.md) |
| T04 Fixed seed + JSON output | [T04-seed-json.en.md](docs/i18n/tutorials/T04-seed-json.en.md) | [T04-seed-json.md](docs/i18n/tutorials/T04-seed-json.md) |
| T05 Every CLI command, hands on | [T05-cli-help.en.md](docs/i18n/tutorials/T05-cli-help.en.md) | [T05-cli-help.md](docs/i18n/tutorials/T05-cli-help.md) |
| T06 Exit codes, verified one by one | [T06-exit-codes.en.md](docs/i18n/tutorials/T06-exit-codes.en.md) | [T06-exit-codes.md](docs/i18n/tutorials/T06-exit-codes.md) |
| T07 Writing a CES file from scratch | [T07-write-ces.en.md](docs/i18n/tutorials/T07-write-ces.en.md) | [T07-write-ces.md](docs/i18n/tutorials/T07-write-ces.md) |
| T08 How to contribute (DCO/CLA) | [T08-contribute.en.md](docs/i18n/tutorials/T08-contribute.en.md) | [T08-contribute.md](docs/i18n/tutorials/T08-contribute.md) |
| T09 AGPLv3 or commercial license? | [T09-license-decision.en.md](docs/i18n/tutorials/T09-license-decision.en.md) | [T09-license-decision.md](docs/i18n/tutorials/T09-license-decision.md) |
| T10 FAQ map + how to run the gates | [T10-faq-gates.en.md](docs/i18n/tutorials/T10-faq-gates.en.md) | [T10-faq-gates.md](docs/i18n/tutorials/T10-faq-gates.md) |

## 5 Documentation

- Licensing and compliance: `LICENSING.md` · `COMMERCIAL-LICENSE.md` · `FAQ-licensing.md` · `LICENSE-COMPATIBILITY.md`
- Contributing and sign-off: `CONTRIBUTING.md` · `DCO.md` · `CLA.md` (draft) · `CODE_OF_CONDUCT.md`
- Security and notices: `SECURITY.md` · `NOTICE` · `THIRD-PARTY-NOTICES.md` · `TRADEMARK.md`
- Governance details: `governance/DUAL-LICENSE.md` · `governance/HEADERS.md` · `governance/THIRD-PARTY.md`
- User documents: `docs/user/quickstart.md` · `docs/user/cli-reference.md` · `docs/user/ces-guide.md`

## 6 Contributing

- Every commit must carry `Signed-off-by` (`git commit -s`, DCO 1.1);
- CLA is currently **optional**; note that the **DCO does not grant the project
  a relicense right; under a dual license a CLA or explicit relicensing terms
  are needed — subject to confirmation by the rights holder or legal counsel**;
- Dependency whitelist: MIT / Apache-2.0 / BSD / ISC / CC0 / Unlicense /
  Public Domain; GPL-family licenses are rejected (see
  `LICENSE-COMPATIBILITY.md`).

## 7 Gates

| Stage | Tool | Rule source | Failure condition |
|------|------|----------|----------|
| 1 | gofmt | CodeReviewComments · Gofmt | Any unformatted file |
| 2 | go vet | All 36 official analyzers (printf/copylocks/loopclosure/lostcancel etc.) | Any analyzer error |
| 3 | golangci-lint | Full linter set from `.golangci.yml` (110+ rules) | Any linter hit (--issues-exit-code=1) |
| 4 | go test -race -count=1 | All tests + race detector | Test failure or data race |
| 5 | go build | Compilation check | Build failure |

### Zero-tolerance items

- **Stubs/placeholders**: `TODO` / `FIXME` / `BUG` / `HACK` comments → blocked by `godox`
- **Fake examples**: Example without `// Output:` → blocked by `testableexamples`
- **Ignored errors**: unchecked return values → `errcheck` (incl. blank/type assertions), `wrapcheck`, `err113`, `errorlint`
- **Security risks**: hardcoded secrets, weak randomness, injection surface → `gosec`, `bidichk`
- **Resource leaks**: unclosed HTTP body / sql.Rows → `bodyclose`, `rowserrcheck`, `sqlclosecheck`
- **Complexity**: cyclomatic >15 (`gocyclo`/`cyclop`), cognitive >20 (`gocognit`), nesting >4 (`nestif`), functions >60 lines (`funlen`), duplication >100 tokens (`dupl`)
- **Style gaps**: naming/comment/export rules → `revive` (all rules), `stylecheck`, `godot` (comments end with a period)
- **Concurrency/test discipline**: missing `t.Parallel()` → `paralleltest`; helpers missing `t.Helper()` → `thelper`; races → `-race`
- **Half-finished work**: non-compiling or unused code → `go build`, `unused`, dead-code rules

### Usage

```bash
# Run the full gate chain locally
make gate
# Or invoke the script directly
bash scripts/lint-gate.sh
```

Any failing stage exits non-zero immediately and prints the rule and location
that was hit; CI fails in sync.

### Rule baselines

- [CodeReviewComments](https://go.dev/wiki/CodeReviewComments)
- [Effective Go](https://go.dev/doc/effective_go)
- [go vet analyzer list](https://pkg.go.dev/cmd/vet) (all 36 enabled)
- [Go language spec](https://go.dev/ref/spec) (Go 1.26.11)
- golangci-lint: starts from `default: none`, enables all non-deprecated
  linters; deprecated entries (exhaustruct/gomodguard/wsl) have been migrated
  to v5/v2 replacements or removed per official guidance

### Exact boundary of the header and license scripts (do not misread)

- `scripts/header-check.sh` **scans only `*.go`, never `.md`**: a newly added
  bilingual `.md` without a copyright header will **not** trigger a FAIL from
  this gate; uniformly carrying the 6-line HTML copyright header is a
  **governance-consistency requirement** (see `governance/HEADERS.md`), not a
  gate-enforced rule.
- The GPL scan of `scripts/license-check.sh` targets the license texts of
  `vendor/modules.txt` dependencies and `third_party/LICENSE-MANIFEST.md`; it
  **does not scan the repository-root `LICENSE`**: placing the verbatim AGPLv3
  text at the root will not cause a false FAIL.
- Lines 41–135 of `scripts/lint-gate.sh` are the frozen nine sections and are
  not modified.

## 8 Language & Status

- This file is the English section-by-section counterpart of the Chinese main
  facade `README.md`;
- This document does not constitute legal advice; `AGPL-3.0-or-later` is
  subject to confirmation by the rights holder;
- Honest-caliber statements (limit faithfully; do not extrapolate):
  - Competitor absolute score: `competitorAbsolute = NOT_MEASURED` (no
    absolute competitor scores are written before a real cross-check);
  - The E2E degrade "meets threshold" statement holds **only under the frozen
    single-file (first-unit) caliber**. On the early third-party corpus, the same
    formula put all six languages above threshold (js 49.12% / c 28.00% /
    py 26.89% / rs 42.33% / kt 8.21% / cs 6.98%), with go `NOT_MEASURED` — these
    figures are a **historical measurement taken before the corpus was fully
    self-researched**, preserved only in historical audit evidence such as
    `docs/i18n/QA-ANTI-FRAUD-AUDIT.md`, and **do not represent the current
    corpus**. The test corpus is now 100% replaced with self-researched sources
    (see `docs/refactor/`); the **full-corpus** degrade rate on the self-researched
    corpus has been re-measured by parsing all 20520 entries with their real
    frontends via `scripts/fullcorpus-degrade.sh`: py 5.04% / js 46.49% /
    c 10.40% / cs 5.12% / rs 6.13% / kt 5.22%, with go `NOT_MEASURED` (no embedded
    grammar mappings). Raw data and SHA-256 are under
    `docs/refactor/fullcorpus-degrade/`. These full-corpus values **do not meet**
    the frozen single-file thresholds and are disclosed honestly; they are not a
    "high-fidelity / meets-threshold" claim. The single-file and full-corpus
    calibers must not be mixed or extrapolated, nor may the historical
    third-party figures above be presented as current self-research measurements;
  - `derive` killing is a **static AIR-level approximation, not a real go test
    execution**.
