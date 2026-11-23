<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->

# CLI Reference

> This document describes all commands and exit codes of the current binary (`abate 1.2.29`).
> Each command's help is available live via `abate <cmd> --help`.

## Command overview

| Command | Purpose |
|---|---|
| `run <ces-file>` | Execute a CES experiment and emit a deterministic report |
| `verify <a> <b>` | Byte-compare two reports |
| `replay <archive> [other]` | Verify a SIM-REC archive, optionally compare two |
| `check <paths...>` | Seven-language structural check |
| `derive` | Derive mutation operators from Go rules and run one mutation experiment |
| `mutate <path>` | Run a deterministic mutation experiment on a real Go workspace (judged by real go test) |
| `targetgen` | Generate a deterministic mutation-testing target project in one of 7 languages |
| `plugin <subcmd>` | Signed plugin management |
| `web` | Local control API + browser console |
| `version` | Version and provenance |
| `--help` | Help |

---

## run

Parse, normalize and validate a CES document, drive a deterministic kernel experiment and emit a report.
`--dry-run` validates only.

```
abate run <ces-file> [flags]
```

| Flag | Description |
|---|---|
| `--seed uint` | Deterministic seed override (0 uses the document seed) |
| `--format console\|json` | Report format (default console) |
| `--out path` | Also write report bytes to a file |
| `--record path` | Write the SIM-REC deterministic archive |
| `--dry-run` | Parse and validate only; do not execute |
| `--strict` | Promote warning-severity violations to failures |
| `--help, -h` | Help |

Examples:
```bash
abate run ctl/cli/testdata/ok.yaml
abate run ctl/cli/testdata/ok.yaml --seed 42 --format json --out report.json
abate run ctl/cli/testdata/ok.yaml --seed 42 --record /tmp/run.arch
abate run ctl/cli/testdata/ok.yaml --dry-run
```

## verify

Compare two run reports byte for byte, printing both SHA-256 digests, sizes, the first differing
offset and the mapped exit code. Exits 0 if identical, 1 on any difference, 3 if a file cannot be read.

```
abate verify <a> <b> [--format console|json] [--out path]
```

## replay

Load a SIM-REC archive, verify its hash chain, total-order sequence and virtual-time monotonicity;
optionally compare a second archive and print a structured Diff. Archives are read-only.

```
abate replay <archive> [compare-archive] [flags]
```
Flags: `--format console|json`, `--out path`, `--help`.

## check

Dispatch seven-language AIR frontends by extension, run the corresponding rule set, and emit a
deterministic console / JSON / SARIF 2.1.0 report. A path may be a file or directory; `vendor`,
`third_party` and `.git` are skipped.

```
abate check <paths...> [flags]
```

| Flag | Description |
|---|---|
| `--format console\|json\|sarif` | Report format (default console) |
| `--severity error\|warning\|info` | Minimum severity (default info) |
| `--rules prefix[,prefix...]` | Restrict to rule ID prefixes |
| `--languages lang[,lang...]` | Restrict to languages (go,py,js,c,cs,rs,kt) |
| `--out path` | Also write report bytes to a file |
| `--help, -h` | Help |

`check` never mutates sources; rule-to-operator derivation is the separate `derive` command.

Examples:
```bash
abate check ./ctl/cli
abate check app.py report.js --format sarif --out /tmp/check.sarif
abate check ./src --languages py,js --rules PY-PEP8
abate check ./lang/air --rules GO-EG --format json
```

## derive

Derive M-family mutation operators from the built-in Go rule set, parse an in-memory Go sample into
AIR, register the operators and run one deterministic mutation experiment.
**Killing is a static AIR-level approximation, not a real go test execution.**

```
abate derive [flags]
```
Flags: `--seed uint` (default 42), `--format console|json`, `--out path`,
`--source path` (optional Go source instead of the sample), `--help`.

Example: `abate derive --seed 42 --format json`


## mutate

Run a deterministic mutation experiment against a real Go workspace: mutable tokens are rewritten in a private staging copy, and each mutant is judged KILLED / SURVIVED by real `go build` + `go test -run` execution, unlike the static approximation of `derive`.

```
abate mutate <path> [flags]
```

| Flag | Description |
|---|---|
| `--seed uint` | Deterministic seed (default 42) |
| `--format console\|json` | Report format (default console) |
| `--out path` | Also write report bytes to a file |
| `--parallel n` | Parallelism; clamped to CPU count when >1, report byte-identical to sequential |
| `--chaos` | Inject and roll back the 12 C-family chaos faults of the L6 injection kernel; the report gains a deterministic chaos summary block and stays byte-identical to a no-chaos run |
| `--native` | Copy `vendor`/`node_modules`/`target` once and symlink into each mutant workspace |
| `--help, -h` | Help |

Examples:
```bash
abate targetgen --out /tmp/tgt --seed 42
abate mutate /tmp/tgt --seed 42
abate mutate /tmp/tgt --seed 42 --parallel 4
abate mutate ./lang/gomutate --seed 42 --format json --out /tmp/mut.json
```

Common errors:
- `path` is not a Go workspace (no go.mod / no tests) -> exit code 3; generate one with `targetgen` first to walk the flow.
- `--parallel` larger than the CPU count is clamped automatically; no error is reported and the report content is unaffected.

## targetgen

Generate a deterministic mutation-testing target project: 7 languages (go/py/js/c/cs/rs/kt) available, byte-identical output for the same seed; every generated structure is guarded by real tests and self-checked by the real toolchain before delivery.

```
abate targetgen --out <dir> [flags]
```

| Flag | Description |
|---|---|
| `--out dir` | Output directory, **must not exist** (overwrite protection) |
| `--lang` | Target language: go/py/js/c/cs/rs/kt (default go) |
| `--seed uint` | Deterministic seed (default 42) |
| `--module path` | Go module path (default example.com/target) |
| `--package name` | Package name (default target) |
| `--funcs n` | Number of generated functions, 1..64 (default 8) |
| `--level` | Difficulty: easy/medium/hard (default medium) |
| `--mixed f` | Mix ratio in [0,1] (default 0) |
| `--complex` | Enable the slice/map/error/nested/method/iface structure families |
| `--skip-check` | Skip the pre-delivery toolchain self-check (not recommended) |
| `--help, -h` | Help |

Examples:
```bash
abate targetgen --out /tmp/tgt --seed 42 --funcs 14 --level easy
abate targetgen --out /tmp/tgt-py --lang py --seed 42
abate targetgen --out /tmp/tgt2 --seed 42
```

Common errors:
- --out directory already exists -> exit code 3; delete it or pick another path and retry.
- --lang cs/rs without the matching toolchain installed -> the self-check stage fails; add --skip-check to generate sources only.

## plugin

Signed plugin management in the native plugin host. Every load verifies the Ed25519 signature against
the public key shipped in the manifest; an unsigned or tampered manifest is rejected and never registered.

```
abate plugin <subcommand> [flags]
```

| Subcommand | Description |
|---|---|
| `load --manifest <file> [--type operator\|reporter\|adapter]` | Verify and load a signed manifest |
| `list` | List registered plugin ids in deterministic order |
| `call --id <id> --entry <entry> [--args <base64>]` | Call an entry (apply / build / detect / emit) |
| `hotswap --id <id> --new-manifest <file>` | Atomically replace an instance without interrupting in-flight calls |

## web

Serve the local control API together with an embedded browser console from a single process,
same-origin, with no CORS configuration; the API is never exposed separately.

```
abate web [flags]
```

| Flag | Description |
|---|---|
| `--listen host:port` | Bind address (default 127.0.0.1:8430) |
| `--token string` | Bearer token; **required** for non-loopback bind; refuses exposed + unauthenticated |
| `--rest string` | External REST base URL (/api reverse-proxied, console still served locally) |
| `--static dir` | Serve static assets from a directory (development only) |
| `--help, -h` | Help |

## version

Prints version and provenance: `abate 1.2.29` / `schema: ces.tencent.com/v1` / `build: deterministic` /
`license` / `source` / `watermark`.

---

<a id="exit-codes"></a>

## Exit codes

Common semantics for top-level / run / derive (PLAN-07 §4):

| Code | Meaning |
|---|---|
| 0 | Success (legal CES and steady state HELD) |
| 1 | Business failure (steady state VIOLATED during execution) |
| 2 | Internal error (scheduler or report output failure) |
| 3 | Usage error (invalid arguments/CES, CES-VALID error; `--strict` promotes warnings) |

**check**: 0 no findings; 1 one or more findings (including warnings; `--severity error` fails only on errors);
2 internal error (frontend/report output); 3 usage error (invalid arguments, missing paths, no rules matched).

**replay**: 0 verified (and equal when two archives supplied); 1 comparison difference (structured Diff printed);
2 internal error (archive parse/integrity/report output); 3 usage error (invalid arguments or unreadable archive).

**plugin**: 0 success; 1 business failure (signature/capability/hot-swap rejection);
2 internal error (output/encoding); 3 usage error (bad flags, unreadable or malformed manifest).

**verify**: 0 identical; 1 difference; 3 file unreadable.
