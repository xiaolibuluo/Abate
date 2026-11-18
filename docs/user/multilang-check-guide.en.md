<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->

# Seven-Language Source Check — Step-by-Step Guide

> This guide is for first-time users of `abate check`. From building the binary
> to checking all 7 languages individually, scanning a mixed directory, emitting
> machine-readable reports and filtering by language — every step gives a
> copy-pasteable command and its expected output. All commands were verified
> locally.

## Contents

- [0. What it does](#0-what-it-does)
- [1. Prepare: build the binary](#1-prepare-build-the-binary)
- [2. First example: check a Go file](#2-first-example-check-a-go-file)
- [3. All seven languages, hands-on](#3-all-seven-languages-hands-on)
- [4. Check a whole project at once (mixed directory)](#4-check-a-whole-project-at-once-mixed-directory)
- [5. Filter by language or rule](#5-filter-by-language-or-rule)
- [6. Emit JSON / SARIF reports](#6-emit-json--sarif-reports)
- [7. Understand exit codes](#7-understand-exit-codes)
- [8. Troubleshooting](#8-troubleshooting)
- [9. Extension-to-language table](#9-extension-to-language-table)

---

## 0. What it does

`abate check` reads source → parses it with the matching language frontend into
AIR (a unified intermediate representation) → runs that language's in-house rule
set → emits a deterministic report.

- Supports **7 languages**: Go, Python, JavaScript, C, C#, Rust, Kotlin;
- Dispatches automatically **by file extension** — no manual language flag needed;
- Report formats: console (human), json (machine), sarif 2.1.0 (CI / code platforms);
- Identical input produces **byte-for-byte identical** output — repeatable,
  comparable, git-friendly.

It **never modifies your source** (check only — no formatting, no auto-fix).

---

## 1. Prepare: build the binary

Zero external dependencies; build offline:

```bash
cd /path/to/Abate
export GOPROXY=off
go build -o abate .
./abate version
```

Expected first line:

```
abate 1.2.27
```

Confirm `check` help lists the seven languages:

```bash
./abate check --help
```

You should see `Seven languages are dispatched by file extension: .go .py .js .c .cs .rs .kt`.

> Below, `./abate` refers to the binary just built. You may put it on PATH and use `abate` directly.

---

## 2. First example: check a Go file

Create a deliberately imperfect Go file (underscore naming triggers a rule):

```bash
mkdir -p /tmp/check-demo
cat > /tmp/check-demo/demo.go <<'EOF'
package demo

var foo_bar int
EOF
```

Run the check:

```bash
./abate check /tmp/check-demo/demo.go
```

Expected output (line/column may differ):

```
/tmp/check-demo/demo.go:3:5: GO-EG-001: identifier foo_bar should not contain underscores [info]
```

Reading it:
- `/tmp/check-demo/demo.go:3:5` → file, line, column;
- `GO-EG-001` → rule ID; the `GO-` prefix means a Go-language rule;
- `[info]` → severity.

The exit code is `1` (violation found) — this is a **normal business failure**, not a crash:

```bash
echo $?   # prints 1
```

---

## 3. All seven languages, hands-on

Create a minimal sample per language and check it. Each command is independent.

### 3.1 Go

```bash
cat > /tmp/check-demo/sample.go <<'EOF'
package p
var foo_bar int
EOF
./abate check /tmp/check-demo/sample.go
```

### 3.2 Python

```bash
cat > /tmp/check-demo/sample.py <<'EOF'
def greet(name):
    result = 1
    return result
EOF
./abate check /tmp/check-demo/sample.py
```

### 3.3 JavaScript

```bash
cat > /tmp/check-demo/sample.js <<'EOF'
function greet(name) {
  var result = 1;
  return result;
}
EOF
./abate check /tmp/check-demo/sample.js
```

### 3.4 C

```bash
cat > /tmp/check-demo/sample.c <<'EOF'
int add(int a, int b) {
    int foo_bar = a + b;
    return foo_bar;
}
EOF
./abate check /tmp/check-demo/sample.c
```

### 3.5 C#

```bash
cat > /tmp/check-demo/sample.cs <<'EOF'
public class Greeter {
    public int Add(int a, int b) {
        var foo_bar = a + b;
        return foo_bar;
    }
}
EOF
./abate check /tmp/check-demo/sample.cs
```

### 3.6 Rust

```bash
cat > /tmp/check-demo/sample.rs <<'EOF'
fn add(a: i32) -> i32 {
    let foo_bar = a;
    foo_bar
}
EOF
./abate check /tmp/check-demo/sample.rs
```

### 3.7 Kotlin

```bash
cat > /tmp/check-demo/sample.kt <<'EOF'
fun add(a: Int): Int {
    val foo_bar = a
    return foo_bar
}
EOF
./abate check /tmp/check-demo/sample.kt
```

Each command emits rule IDs with the matching language prefix (`PY-`, `JS-`,
`C-`, `CS-`, `RS-`, `KT-`), proving the file was routed to the correct frontend
and rule set.

---

## 4. Check a whole project at once (mixed directory)

With all seven samples in `/tmp/check-demo/`, check the entire directory:

```bash
./abate check /tmp/check-demo --format json
```

`abate` will:
1. Recursively scan the directory (skipping `vendor/`, `third_party/`, `.git/`);
2. Dispatch each file by extension to its language frontend;
3. Aggregate all findings in a fixed order (path → line → column → rule ID);
4. Emit one unified report.

From the JSON you will see findings from every language (the `ruleId` prefix is
the language): `GO-` / `PY-` / `JS-` / `C-` / `CS-` / `RS-` / `KT-`.

Multiple files/directories in one run work too:

```bash
./abate check ./src ./tests extra.py --format json
```

---

## 5. Filter by language or rule

**By language** (`--languages`, comma-separated tags):

```bash
# check only Python and JavaScript; other languages in the directory are ignored
./abate check /tmp/check-demo --languages py,js
```

**By rule** (`--rules`, comma-separated rule ID prefixes):

```bash
# only rules starting with PY-PEP8
./abate check /tmp/check-demo --rules PY-PEP8
```

**By severity** (`--severity` is a lower bound):

```bash
# report error and above only
./abate check /tmp/check-demo --severity error
```

Language tags and rule prefixes are in [section 9](#9-extension-to-language-table).

---

## 6. Emit JSON / SARIF reports

**JSON** (for scripts):

```bash
./abate check /tmp/check-demo --format json --out report.json
head -20 report.json
```

The report contains `schema: abate.check.v1`, `summary` (files/rules/total),
`findings` and `fingerprint`. `--out` writes the exact stdout bytes to a file.

**SARIF 2.1.0** (for GitHub/GitLab code scanning):

```bash
./abate check /tmp/check-demo --format sarif --out report.sarif
```

The SARIF report embeds no timestamp and has fixed field order — safe to commit
and diff.

**Verify determinism** (two runs byte-identical):

```bash
./abate check /tmp/check-demo --format json --out a.json
./abate check /tmp/check-demo --format json --out b.json
cmp a.json b.json && echo "identical"   # expect "identical"
```

---

## 7. Understand exit codes

`check` follows the four-state PLAN-07 §4 contract:

| Code | Meaning | What to do |
|---|---|---|
| `0` | No violations | Passed |
| `1` | Violations found (business failure) | Read the report, fix code; the common "findings" state |
| `2` | Internal error (frontend/output failure) | File a bug with the command and output |
| `3` | Usage error (bad args, missing path, no matching rules) | Check spelling, path, `--rules`/`--languages` values |

In CI, treat `0` as pass, `1` as findings needing review, `2`/`3` as config errors.

---

## 8. Troubleshooting

**Q: `check` returns 1 — is it broken?**
No. 1 means "violations found", i.e. the rule engine worked. Read the rule ID and
location. To get 0, narrow with `--rules` to a clean subset or fix the code.

**Q: Why was a file not checked?**
Only the extensions in [section 9](#9-extension-to-language-table) are supported.
`.ts`, `.cpp`, `.java`, etc. are not; files under `vendor/`/`third_party/`/`.git/`
are skipped.

**Q: `--languages xx` reports a usage error?**
Tags must be one of `go,py,js,c,cs,rs,kt` (lowercase). Others (e.g. `python`,
`golang`) match no rules and return 3.

**Q: No findings for a language?**
Verify the extension and that rules loaded. Check alone with
`./abate check <file> --languages <tag> --format json` and confirm
`summary.rules` > 0.

**Q: Spaces in the path?**
Quote it: `./abate check "my dir/sample.py"`.

**Q: Does `check` modify my source?**
No. It is strictly read-only (`--out` writes only the report file).

---

## 9. Extension-to-language table

| Extension | Language tag | Rule prefix |
|---|---|---|
| `.go` | `go` | `GO-` |
| `.py` | `py` | `PY-` |
| `.js` `.jsx` `.mjs` `.cjs` | `js` | `JS-` |
| `.c` `.h` | `c` | `C-` |
| `.cs` | `cs` | `CS-` |
| `.rs` | `rs` | `RS-` |
| `.kt` `.kts` | `kt` | `KT-` |

> More commands: `docs/user/cli-reference.en.md`; the specification:
> `docs/spec/ces-v1.md`; language frontends and fidelity:
> `docs/ARCHITECTURE.md` §3.1 and `docs/i18n/RELEASE-READINESS-FINAL.md`.
