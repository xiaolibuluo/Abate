<!--
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
-->

# What Abate Can and Cannot Do (Capabilities and Limits)

> Purpose: state the real capability boundary of `abate check`'s seven-language
> support with **reproducible evidence**, so that "multi-language linter" is not
> misread as "compiler-grade semantic analysis / go vet-level understanding of
> seven languages". Every claim below can be reproduced with the commands at the
> end.

## One-line positioning

Abate is a **structural linter**: it **really parses** seven languages into a
unified AIR abstract syntax tree, then finds issues by **structural pattern
matching** on the AST. It is not a compiler — no type checking, no call graph,
no interprocedural data-flow analysis.

## It really parses — evidence

A common suspicion: "zero dependencies, no tree-sitter / libclang — seven
languages must be regex matching." Half of that is right: **it is zero-dependency
and uses no third-party parser, but it is not regex matching.**

The pipeline is complete and in-house:

| Stage | Implementation | Location |
|-------|----------------|----------|
| Lexing | hand-written lexer + token spec (one per language) | `lang/<lang>front/*lexspec.go`, `*tokens.go` |
| Parsing | Pratt precedence-climbing parser, integer binding powers, no float/clock/random state | `lang/frontkit/parser.go` |
| AST | produces structured `air.Node` (kind / SemKey / attributes / children) | `lang/air/node.go`, `lang/<lang>front/*map*.go` |
| Rules | compiled structural patterns: `$_` any node, `$x` capture, `$..` any-depth, `!P` negation, `$x*` repeat, attribute constraints `= != ~ ^ + /regex/` | `lang/rule/pattern.go`, `lang/rule/matcher.go` |

Rules match **AST node structure**, not text lines. Example Python rule:

```go
// lang/rule/builtin/py/rules.go
ID:      "PY-PEP8-002",
Pattern: `ClassDecl~_`,   // match kind=ClassDecl with underscore in attribute name
```

In patterns like `Ident/[A-Z]/`, `ClassDecl~_`, `NumLit/^[0-9]{4,}$/`, the kind
(`Ident` / `ClassDecl` / `NumLit`) is an **AST node type**; the regex only
constrains the node's **attribute text** — different from running regex over the
whole file.

**Empirical confirmation**: a Python file with four nesting levels
(`if → for → assign → return`) yields `files=1, total=28`, hitting five
structural rule classes `PY-AIR-005/007/009/010/050` — the parser built the
nested structure and rules matched on the tree, not a line-by-line text scan.

## What it is NOT — limits (equally important)

Honest limits earn more trust than inflated claims:

| Limit | Detail |
|-------|--------|
| **Error-tolerant parsing, not strict syntax validation** | Syntactically broken input does not fail the command; it produces a **partial AIR + diagnostics** and continues (`lang/frontkit/frontend.go` documents this by design). Tested: broken `def f(:` / `x =` still emits 12 warnings. **It is not a syntax-correctness gate.** |
| **No type checking** | No type inference or type-match validation. It is not mypy / pyright / rustc's type system. |
| **No control-flow / data-flow semantics** | No CFG, no interprocedural analysis, no tracking variable propagation across branches. Findings are **local structural pattern** hits. |
| **Not go vet-grade depth** | The Go-side `go vet` analyzers are part of the build gate (`scripts/lint-gate.sh`), but `abate check`'s seven-language rules are **shared AIR templates + ~a dozen per-language rules** — not equivalent in depth to purpose-built tools (go vet / eslint / clippy / Pylint). |
| **Limited rule count** | 72 shared AIR structural templates + 12–14 per-language rules (Go also has vet/effector/CRC series). Coverage targets naming/style/common structural smells, **not the full language spec**. |

## Per-language actual coverage

`abate check` dispatches by extension: `.go` `.py` `.js` `.c` `.cs` `.rs` `.kt`
(aliases `.jsx/.mjs/.cjs/.h/.kts`). Each language adds dedicated rules on top of
72 shared AIR templates:

| Language | Dedicated rules | Focus |
|----------|-----------------|-------|
| Go | 116 (vet / effector / CRC series) | closest to go vet style, multiple vet groups |
| Python | 12 | PEP 8 naming / style / idioms |
| JavaScript | 12 | naming / style / common structures |
| C | 14 | naming / structure / idioms |
| C# | 12 | naming / style / idioms |
| Rust | 14 | naming / structure / idioms |
| Kotlin | 12 | naming / style / idioms |

> Rule counts from `ID:` declarations under `lang/rule/builtin/`. Dedicated rules
> plus shared templates make up the `summary.rules` field in `check --format json`
> reports (a sample file shows 72, the shared-template count).

## When to use / when not

**Good for**:
- unified structural style sweeps across a polyglot repo — one report, 7 languages;
- deterministic smell gates in CI (fixed input → byte-identical output, comparable via `verify`);
- machine-readable reports (JSON / SARIF) for downstream processing.

**Not suitable** (use purpose-built tools instead):
- full syntax error diagnostics → the language's own compiler / parser;
- type safety, null checks, borrow checking → mypy / rustc / Roslyn / go vet;
- deep semantics, interprocedural analysis → a dedicated static-analysis platform.

## Reproduce

**首选：one command reproduces all 7 claims** (builds from source + asserts):

```bash
bash scripts/evidence-check.sh     # exits 0 only if all PASS; parsing/tolerance/determinism/7-lang/help
bash scripts/language-coverage.sh  # prints the quantified per-language rule/hit table
```

Manual step-by-step breakdown (what the scripts do internally):

```bash
# 1. real parsing: nested structure yields multiple structural rule hits
printf 'def outer(x):\n    if x:\n        for i in range(3):\n            y = i\n            return y\n    return x\n' > /tmp/nested.py
abate check /tmp/nested.py --format json --out /tmp/nested.json
# observe summary.total > 0 and byRule with several PY-AIR-* classes

# 2. error-tolerance limit: broken syntax does not fail, emits partial diagnostics
printf 'def f(:\n  x =\n' > /tmp/broken.py
abate check /tmp/broken.py --format json --out /tmp/broken.json
# observe: a report is still produced (not a syntax gate)

# 3. zero-dependency evidence: no vendor/, no third-party require in go.mod
grep -c 'require' go.mod   # prints 0
ls vendor 2>/dev/null      # does not exist

# 4. structural matching evidence: rule DSL matches AST node kinds
sed -n '20,40p' lang/rule/builtin/py/rules.go
```

## Version

The single authoritative build version is **`go 1.26.11`** from `go.mod`, matching
the local `go version` output. Other go version numbers that appeared in the repo
(e.g. old values in design drafts) are historical records, not the current build
version.

## Community validation status (stated honestly)

As of this update the repository has **0 stars / 0 forks / 0 watchers and a
single maintainer**. That is an objective fact — not hidden and not sugar-coated.
Three points:

1. **What it means**: no external endorsement means the "community validation"
   dimension is currently empty — nobody has run it at scale on real projects,
   no third-party reproduction, no issue-driven defect discovery. Star count is
   not itself proof of capability, but it does reflect "how many people have
   checked it". **Current answer: few. Do not assume battle-tested.**
2. **What we will not do**: no star padding, no bought forks, no fabricated users
   or benchmarks. Any claimed "community recognition" that is not real would
   damage this project's credibility — the very thing this document exists to
   repair.
3. **How we lower the trial cost (without faking)**:
   - every capability claim is reproducible via `scripts/evidence-check.sh` —
     anyone can build their own evidence instead of trusting the docs;
   - `scripts/language-coverage.sh` publishes real per-language rule counts, no
     hiding;
   - determinism is byte-verifiable with `abate verify` — claims are machine-falsifiable;
   - run it on a real repo and open an issue — **a concrete failing case is worth
     more than a star**.

In short: an early-stage project whose **capabilities are self-verifiable but
community validation is pending**. Please evaluate it with early-stage standards.
