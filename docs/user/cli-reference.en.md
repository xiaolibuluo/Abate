<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->

> Language / 语言：[English](cli-reference.en.md) | [中文](cli-reference.md)

# CLI Reference

> Compiled from the real output of `abate --help` and `abate <cmd> --help`
> (rebuilt from source with `GOFLAGS=-mod=vendor`; commands and flags below
> match the actual binary). The command tree is stateless: identical
> arguments produce identical bytes on every run (deterministic).
>
> Note: this document's flag/command descriptions follow the real `--help`
> output of a fresh source build; where an older Chinese original and the real
> help text differ, the real `abate <cmd> --help` output is authoritative.

## Command tree

| Command | Purpose |
|---|---|
| `run <ces-file>` | Parse/normalize/validate a CES document, execute a deterministic experiment, emit a report. `--record` also writes a SIM-REC archive. |
| `verify <a> <b>` | Compare two run reports byte for byte. |
| `replay <archive> [other]` | Verify an archive (frame hash chain + whole-archive SHA-256), optionally compare two. |
| `check <paths...>` | Dispatch by file extension to 7 language frontends (`.go/.py/.js/.c/.cs/.rs/.kt`), run each language's rule set; emit console/JSON/SARIF. |
| `derive` | Derive M-family operators from rules and run a mutation experiment (static AIR-level killing approximation). |
| `plugin` | Verify, load, call, list and hot-swap signed plugins. |
| `web` | Serve the control API with a self-contained browser console. |
| `version` | Print build version and provenance. |
| `help` / `--help` / `-h` | Print help. |

## abate check

Seven-language source checking, dispatched by file extension to the matching
frontend and rule set:

```
Usage: abate check <paths...> [flags]
  --format console|json|sarif  Report format (default console).
  --severity error|warning|info  Minimum severity (default info).
  --rules prefix[,prefix...]   Restrict to rule identifier prefixes.
  --languages lang[,lang...]   Restrict to language tags (go,py,js,c,cs,rs,kt).
  --out path                   Also write the report bytes to a file.
```

| Extension | Language | Rule prefix |
|---|---|---|
| `.go` | go | `GO-` |
| `.py` | py | `PY-` |
| `.js` `.jsx` `.mjs` `.cjs` | js | `JS-` |
| `.c` `.h` | c | `C-` |
| `.cs` | cs | `CS-` |
| `.rs` | rs | `RS-` |
| `.kt` `.kts` | kt | `KT-` |

Four exit codes: `0` no violations; `1` violations found (business failure);
`2` internal error; `3` usage error.

For a step-by-step walkthrough see `docs/user/multilang-check-guide.en.md`.

## abate run

```
Usage: abate run <ces-file> [flags]
Flags:
  --seed uint        Deterministic seed override (0 uses the document seed).
  --format console|json  Report format (default console).
  --out path          Also write the report bytes to a file.
  --record path       Write the deterministic SIM-REC archive to a file.
  --dry-run           Parse and validate only; do not execute.
  --strict            Promote warning-severity violations to failures.
  --help, -h           Print this help text.
```

Exit codes (PLAN-07 §4):

| Code | Meaning |
|---|---|
| 0 | success (legal CES and steady state HELD) |
| 1 | business failure (steady state VIOLATED during execution) |
| 2 | internal error (scheduler or report output failure) |
| 3 | usage error (invalid arguments, invalid CES document, or a CES-VALID error violation; `--strict` promotes warnings to this code too) |

## abate version

```
abate 1.2.26
schema: ces.tencent.com/v1
build: deterministic
license: AGPLv3 or commercial
source: https://github.com/xiaolibuluo/Abate
watermark: 小黎的部落
```

> The engine's internal version is `0.1.0` (separate dimension from the release
> version `vX.Y.Z`: one is build provenance, the other the release cadence);
> release assets get version/commit metadata injected via `ldflags` by
> `scripts/release.sh`.

## Copy-and-run examples

```bash
abate run ctl/cli/testdata/ok.yaml
abate run ctl/cli/testdata/ok.yaml --seed 42 --format json --out report.json
abate run ctl/cli/testdata/ok.yaml --seed 42 --record /tmp/run.arch
abate replay /tmp/run.arch
abate verify report-a.json report-b.json
abate check ./ctl/cli
abate check ./lang/rule --format sarif --out /tmp/check.sarif
abate derive --seed 42 --format json
abate plugin load --manifest plugin-manifest.json --type operator
abate plugin call --id abate/operator --entry apply --args ZWJhdGU=
abate web --listen 127.0.0.1:8430
abate version
```
