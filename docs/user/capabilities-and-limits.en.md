<!--
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
-->

# What It Can and Cannot Do (Capabilities and Limits)

> This document explains the real boundaries of `abate check` across seven languages with
> **reproducible evidence**, so "multi-language linter" is not misread as "compiler-grade semantic
> analysis". Every conclusion can be reproduced locally with the commands at the end.

## In one sentence

Abate is a **structural static linter**: it **really parses** 7 languages into a unified AIR
abstract syntax tree and does **structural pattern matching** on the tree. It is not a compiler:
no type checking, no call graph, no interprocedural data-flow analysis.

## Real parsing — evidence

A common doubt: "zero dependencies, no tree-sitter / libclang — seven languages must be regex."
Half right: **indeed zero dependencies, no third-party parser, but not regex**. The parse chain is
complete and self-written:

| Stage | Implementation | Location |
|---|---|---|
| Lexing | Hand-written lexer + token spec (one per language) | `lang/<lang>front/*lexspec.go`, `*tokens.go` |
| Syntax | Pratt operator-precedence climbing parser, binding powers as integer data | `lang/frontkit/parser.go` |
| AST | Produces `air.Node` structured nodes (kind / SemKey / attributes / children) | `lang/air/node.go`, `lang/<lang>front/*map*.go` |
| Rules | Compiled structural patterns: `$_` any, `$x` capture, `$..` descendant, `!P` negation, `$x*` repeat, attribute constraints | `lang/rule/pattern.go`, `lang/rule/matcher.go` |

Rules match **AST node structure**, not text lines. Example Python rule:

```go
// lang/rule/builtin/py/rules.go
ID:      "PY-PEP8-002",
Pattern: `ClassDecl~_`,   // match kind=ClassDecl with an underscore in an attribute name
```

`Ident` / `ClassDecl` in patterns such as `Ident/[A-Z]/` are **AST node kinds**; the regex only
constrains a node's **attribute text** — different from running regex over a whole file.

**Measured evidence**: a Python file with `if → for → assignment → return` four-level nesting
produces `total=28` and hits five structural rule classes `PY-AIR-005/007/009/010/050` — proof that
the parser builds nested structure and rules match on the tree.

## What it is not — boundaries

Honest boundaries beat exaggeration:

| Boundary | Description |
|---|---|
| **Tolerant parsing, not strict syntax validation** | Broken input does not fail the command; it produces **partial AIR + warnings** and continues. **Not a "is syntax correct" gate.** |
| **No type checking** | No type inference or type matching; not mypy / pyright / rustc. |
| **No control-flow / data-flow analysis** | No CFG, no interprocedural analysis, no tracking variables across branches; findings are **local structural pattern** hits. |
| **Not go vet depth** | `go vet` is part of the build gate (`scripts/lint-gate.sh`), but `abate check` rules are **shared AIR templates + per-language rules**, not equivalent to go vet / eslint / clippy in depth. |
| **Limited rule count** | 72 shared AIR templates + 12–14 per-language rules (Go additionally has vet/EG/CRC series, 116 total). Covers naming/style/common smell, **not the full language spec**. |

## Per-language coverage

Dispatch by extension `.go` `.py` `.js` `.c` `.cs` `.rs` `.kt` (aliases `.jsx/.mjs/.cjs/.h/.kts`):
each language adds dedicated rules on top of 72 shared AIR templates. Measured details in
[Language coverage](language-coverage.en.md).

| Language | Dedicated rules | Focus |
|---|---|---|
| Go | 116 | vet / Effective Go / CodeReviewComments (hits need specific semantic context) |
| Python | 12 | PEP 8 naming / style / idioms |
| JavaScript | 12 | Naming / style / common structure |
| C | 14 | Naming / structure / CERT secure coding |
| C# | 12 | Naming / style / .NET design guidelines |
| Rust | 14 | Naming / style / idioms |
| Kotlin | 12 | Naming / style / design conventions |

## When to use / not use

**Good for**: unified structural style checks across languages (one report covers 7); deterministic
smell gates in CI; JSON/SARIF machine-readable reports.

**Not good for** (use dedicated tools): full syntax error diagnosis (use the language compiler/parser);
type safety / null / borrow checking (use mypy / rustc / Roslyn / go vet); deep semantic or
interprocedural analysis (use a dedicated static-analysis platform).

## Reproduction

```bash
bash scripts/evidence-check.sh     # on-site build + 7 assertions, all PASS to exit 0
bash scripts/language-coverage.sh  # per-language rule counts and measured hit table
```

Manual breakdown (inside the scripts): `abate check /tmp/nested.py --format json` to see nested hits;
`abate check /tmp/broken.py` to see tolerance (broken input still yields a report);
`grep -c require go.mod` to see zero dependencies.

## Honest scope (do not over-extrapolate)

- **Full-corpus fidelity**: six languages measured 3%–5% (see [Language coverage](language-coverage.en.md)),
  Go has no embedded grammar mappings → `NOT_MEASURED`; the E2E gate "passing" applies only to the
  frozen single-file (first-unit) scope; the full-corpus scope does not reach that threshold and is
  **not a "high fidelity" claim**.
- **Competitor comparison**: `competitorAbsolute = NOT_MEASURED`, never actually benchmarked.
- **derive killing**: static AIR-level approximation, **not a real go test execution**.
- **Community validation**: 0 star / 0 fork / 0 watcher, individually maintained — no star buying,
  no fork buying, no fabricated benchmarks. Lowering the trial barrier relies on one-command
  reproduction, public quantification and machine-falsifiable determinism; real failing cases welcome.

In short: an early project whose capability is self-provable but not yet community-validated.

## Version

The sole authoritative build version is `go 1.26.11` in `go.mod`, consistent with local `go version`.
Other go version strings that appeared in the repo are historical records, not the current build version.
