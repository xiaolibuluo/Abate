<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->

> Language / 语言：[English](T05-cli-help.en.md) | [中文](T05-cli-help.md)

# T05 · Every CLI command, hands on

## 1 Audience

Users and operators who want to walk through all 8 subcommands: the goal is to
self-serve with `--help` and to read the command tree, the four-state exit
codes and the 6 common flags.

## 2 Prerequisites

- Build completed per T01;
- This tutorial only runs `--help`; no experiment is executed and no files
  are written;
- Command tree: `run / verify / replay / check / derive / plugin / web /
  version` (plus `help` / `--help` / `-h`).

## 3 Commands

```bash
/tmp/abate_verify --help
/tmp/abate_verify run --help
/tmp/abate_verify verify --help
/tmp/abate_verify replay --help
/tmp/abate_verify check --help
/tmp/abate_verify derive --help
/tmp/abate_verify plugin --help
/tmp/abate_verify web --help
/tmp/abate_verify version --help
```

## 4 Expected output

`--help` (top level, measured excerpt — the command tree and exit-code section
are quoted verbatim):

```
abate - deterministic experiment control plane

Usage:
  abate <command> [flags]

Commands:
  run <ces-file>     Execute a CES experiment and emit a deterministic report.
                     --record <path> also writes a SIM-REC archive.
  verify <a> <b>     Compare two run reports byte for byte.
  replay <archive> [other]  Verify an archive, optionally compare two.
  check <paths...>   Run the built-in Go rule set; emit console/JSON/SARIF.
  derive             Derive M-family operators from rules; run a mutation
                     experiment (static AIR-level killing approximation).
  plugin             Verify, load, call, list and hot-swap signed plugins.
  web                Serve the control API with a self-contained browser console.
  version            Print build version and provenance.

Exit codes (PLAN-07 §4):
  0  success (legal CES and steady state HELD).
  1  business failure (steady state VIOLATED during execution).
  2  internal error (scheduler or report output failure).
  3  usage error (invalid arguments, invalid CES document, or a CES-VALID
     error violation; --strict promotes warnings to this code too).
```

`run --help` (the 6 flags, verbatim):

```
  --seed uint        Deterministic seed override (0 uses the document seed).
  --format console|json  Report format (default console).
  --out path          Also write the report bytes to a file.
  --record path       Write the deterministic SIM-REC archive to a file.
  --dry-run           Parse and validate only; do not execute.
  --strict            Promote warning-severity violations to failures.
```

`derive --help` (the key sentence is kept **as-is**; do not paraphrase it into
"a mutation test that really runs go test"):

```
Killing is a static AIR-level approximation, not a real go test execution.
```

**Verification points**: all 9 `--help` calls exit 0; the 8 subcommands in the
command tree match the list above; the four-state exit codes 0/1/2/3 are
declared in the `run`/`replay`/`check` help texts.

## 5 Troubleshooting

| Symptom | Cause | Action |
|---|---|---|
| `--help` exits 3 | Misspelled subcommand | Run `abate --help` first to see the command tree |
| Help text differs from this tutorial | You ran a stale binary | Repeat the T01 build |
| Confusing `check` with `derive` | Both involve "rules" | `check` only reads sources and emits a report, **it never mutates sources**; `derive` derives M-family operators and runs a mutation experiment (static AIR-level approximation) |
| `web` serves but the console does not open | Bound a non-loopback address without a token | Default bind is `127.0.0.1:8430`; a non-loopback bind requires `--token`, otherwise the server refuses to start |
