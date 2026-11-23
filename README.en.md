<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->

> Language / 语言：[English](README.en.md) | [中文](README.md)

# Abate

[![version](https://img.shields.io/badge/version-1.2.29-blue)](https://github.com/xiaolibuluo/Abate/releases/latest)
[![license](https://img.shields.io/badge/license-AGPLv3%20%7C%20commercial-green)](LICENSING.en.md)
[![checks](https://img.shields.io/badge/checks-7%20languages-orange)](docs/user/multilang-check-guide.en.md)
[![Go](https://img.shields.io/badge/Go-1.26.11-00ADD8?logo=go)](go.mod)

**Deterministic multi-language static analyzer + experiment control plane.** A single static binary provides two capabilities:

- **`abate check`**: dispatches **Go / Python / JavaScript / C / C# / Rust / Kotlin** by file extension,
  parses source with a self-written lexer + Pratt parser into a unified AIR syntax tree, matches
  structural patterns on the tree, and emits byte-reproducible console / JSON / SARIF reports.
- **`abate run`**: describes deterministic experiments in CES (`ces.tencent.com/v1`), drives a
  fixed-seed kernel and SIM-REC archive, and emits runs that compare byte for byte via
  `abate verify` / `abate replay`.

> **Capability boundary (stated up front):** `abate check` is a **structural linter**, not
> compiler-grade semantic analysis. It does no type checking, control-flow/data-flow analysis,
> and does not match the depth of go vet / eslint / clippy. It uses tolerant parsing, so syntax
> errors do not fail the command — it is **not a syntax gate**. Every claim is self-provable in
> one command — see [Evidence and boundaries](#7-evidence-and-boundaries).

## 1 Contents

| Section | Topic |
|---|---|
| [2 Requirements](#2-requirements) | Go version, zero dependencies, install |
| [3 Quick start](#3-quick-start) | Build + two real examples |
| [4 Commands](#4-commands) | 9 commands and exit codes |
| [5 Capability boundaries](#5-capability-boundaries-and-honest-scope) | Real parsing / tolerance / rule depth / unmeasured items |
| [6 Documentation map](#6-documentation-map) | User / dev / spec / legal index |
| [7 Evidence and boundaries](#7-evidence-and-boundaries) | One-command reproduction scripts and gates |
| [8 License and community](#8-license-and-community-status) | Dual license, contributing, community facts |

## 2 Requirements

- **Go 1.26.11** (the sole authoritative version, per the `go` directive in `go.mod`).
- **Zero external dependencies**: no `require` in `go.mod`, empty `go.sum`, no `vendor/`, no `third_party/`;
  offline builds work with `GOPROXY=off`.
- Optional: `golangci-lint` (version as installed locally).

Release binaries are static builds for five platforms: `linux/amd64` `linux/arm64` `darwin/amd64`
`darwin/arm64` `windows/amd64`, each with a GPG detached signature (key `5CF28F616D9888DB`).
See [Releases](https://github.com/xiaolibuluo/Abate/releases/latest).

## 3 Quick start

```bash
export GOPROXY=off
go build -o abate .
./abate version          # abate 1.2.29 / schema ces.tencent.com/v1 / ...
```

**Example 1: check seven-language source** (auto-dispatch by extension, no language flag needed):

```bash
./abate check ./path/to/your/code        # mixed dir: .go/.py/.js/.c/.cs/.rs/.kt in one run
./abate check app.py --format json       # machine-readable report
./abate check ./src --languages py,js --rules PY-PEP8
./abate check ./src --format sarif --out report.sarif
```

> `check` returns exit code **1** on any finding (including a warning); add `--severity error`
> to fail only on errors. See the [CLI reference](docs/user/cli-reference.en.md).

**Example 2: run a CES experiment** (a valid sample shipped in the repo):

```bash
./abate run ctl/cli/testdata/ok.yaml --seed 42 --format json --out report.json
./abate verify report.json report.json     # byte-for-byte compare, diff=0 exits 0
./abate run ctl/cli/testdata/ok.yaml --seed 42 --record /tmp/run.arch
./abate replay /tmp/run.arch               # verify archive hash chain / total order / monotonicity
```

## 4 Commands

| Command | Purpose |
|---|---|
| `run <ces-file>` | Parse/validate CES, run a deterministic experiment, emit a report; `--dry-run` validates only, `--record` writes a SIM-REC archive |
| `verify <a> <b>` | Byte-compare two reports; print SHA-256, sizes, first differing offset |
| `replay <archive> [other]` | Verify archive integrity, optionally compare two and print a structured Diff |
| `check <paths...>` | Seven-language structural check; `--languages` filter, `--rules` prefix filter, `--format console\|json\|sarif` |
| `derive` | Derive M-family mutation operators from built-in Go rules and run one deterministic mutation experiment (killing is a static approximation) |
| `plugin` | Signed plugins: `load` / `list` / `call` / `hotswap` (Ed25519 signature verification) |
| `web` | Local control API + embedded browser console (same-origin; `--token` required for non-loopback bind) |
| `version` | Print version and provenance |
| `--help` | Top-level help |

**Exit codes**: `0` success; `1` business failure (steady state VIOLATED / findings present);
`2` internal error; `3` usage error. `check` and `replay` have their own dedicated semantics —
see the [CLI reference](docs/user/cli-reference.en.md#exit-codes).

## 5 Capability boundaries and honest scope

Abate documents only facts reproducible from current code and scripts. Boundaries stated honestly:

| Dimension | Current fact |
|---|---|
| Parsing | Self-written lexer + Pratt parser → `air.Node` AST → structural pattern matching (**not regex over text**) |
| Syntax gate | Tolerant parsing: broken input yields partial AIR + warnings, **not a syntax-correctness gate** |
| Semantic depth | No type checking / CFG / data flow / interprocedural analysis; not equivalent to dedicated tools |
| Rule depth | 72 shared AIR templates + per-language rules (Go 116, others 12–14); Go hits need specific semantic context |
| Full-corpus construct coverage | Six languages measured in `docs/dev/frontend-coverage/` per-language reports (C 66.4% / C# 52.6% / JS 89.5% / Kotlin 47.0% / Python 91.8% / Rust 77.9%); Go has no embedded grammar mappings -> `NOT_MEASURED`; not a "high fidelity" claim |
| Competitor comparison | `competitorAbsolute = NOT_MEASURED` (never actually benchmarked; no absolute scores) |
| derive killing | Static AIR-level approximation, **not a real go test execution** |
| Community validation | 0 star / 0 fork / 0 watcher, individually maintained; no star buying, no fabricated benchmarks |

See [Capabilities and limits](docs/user/capabilities-and-limits.en.md).

## 6 Documentation map

**User docs** (bilingual)
- [Quick start](docs/user/quickstart.en.md) · [CLI reference](docs/user/cli-reference.en.md) · [CES guide](docs/user/ces-guide.en.md)
- [Seven-language check guide](docs/user/multilang-check-guide.en.md) · [Capabilities and limits](docs/user/capabilities-and-limits.en.md) · [Language coverage](docs/user/language-coverage.en.md)

**Architecture and spec**
- [Architecture](docs/ARCHITECTURE.en.md) · [CES v1 spec](docs/spec/ces-v1.en.md)

**Development**
- [Contributing guide](docs/dev/contributing.en.md) · [Plugin dev guide](docs/dev/plugin-dev-guide.en.md) · [Module dev guide](docs/dev/module-dev-guide.en.md) · [Sandbox](docs/dev/sandbox.en.md)

**Legal and governance** (not legal advice)
- [LICENSING.en.md](LICENSING.en.md) · [COMMERCIAL-LICENSE.en.md](COMMERCIAL-LICENSE.en.md) · [FAQ-licensing.en.md](FAQ-licensing.en.md) · [LICENSE-COMPATIBILITY.en.md](LICENSE-COMPATIBILITY.en.md)
- [CONTRIBUTING.en.md](CONTRIBUTING.en.md) · [DCO.en.md](DCO.en.md) · [CLA.en.md](CLA.en.md) · [CODE_OF_CONDUCT.en.md](CODE_OF_CONDUCT.en.md)
- [SECURITY.en.md](SECURITY.en.md) · [THIRD-PARTY-NOTICES.en.md](THIRD-PARTY-NOTICES.en.md) · [TRADEMARK.en.md](TRADEMARK.en.md)
- Governance details: `governance/`

## 7 Evidence and boundaries

Every capability claim is reproducible from source on-site — no need to trust the docs:

```bash
bash scripts/evidence-check.sh     # on-site build + 7 capability assertions, all PASS to exit 0
bash scripts/language-coverage.sh  # per-language rule counts and measured hit table
bash scripts/fullcorpus-degrade.sh # full-corpus degradation rate (honest scope)
bash scripts/lint-gate.sh          # gates: header/style/license/vet/golangci-lint/test-race/build
```

Also: `hash-check.sh` / `watermark-verify.sh` (integrity watermark), `determinism-check.sh`
(determinism), `drift-check.sh` (spec consistency), `audit-zero-third-party.sh` (zero third-party
source), `m5-reference-scan.sh`, `vuln-scan.sh`, `tla/` (TLA+ formal verification).

## 8 License and community status

Abate is **dual-licensed** under **AGPLv3 (open source) + commercial**: the open path is `LICENSE`
(full AGPLv3 text), SPDX `AGPL-3.0-or-later` (**subject to confirmation by the rights holder; not
legal advice**); the commercial path is `COMMERCIAL-LICENSE.md`.

Contributing: commits require `Signed-off-by` (`git commit -s`, DCO 1.1); CLA is currently optional
(note: **DCO does not grant relicensing rights; dual licensing needs a CLA or explicit relicensing
terms, pending rights-holder/legal confirmation**).

Community status: as of this update, the repo has **0 star / 0 fork / 0 watcher, individually
maintained** — an objective early-project fact, stated plainly. The project lowers the trial
barrier via one-command reproduction, public quantification, and machine-falsifiable determinism;
real failing test cases are welcome.
