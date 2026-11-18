<!--
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
-->

# Seven-Language Coverage, Quantified (measured)

> Data is **measured live from source** by `scripts/language-coverage.sh` (builds the
> binary, then runs it), not estimated. Reproduce: `bash scripts/language-coverage.sh`.
> The script output is authoritative for any number.

## Why this report

To answer "Go is the only solid one; the other languages' depth is unquantified".
Conclusion up front: **Go and the other six are two different kinds of depth**, not a
simple solid-vs-hollow split — Go takes the "specialized vet analyzer" route, the
other six take the "shared AIR structural templates + a handful of dedicated rules"
route. Both are real, but they cover different things; one ruler does not fit both.

## Quantified table (measured)

| Language | Dedicated rules (source `ID:` count) | Rule classes hit by sample | Findings in sample | Rule families |
|----------|--------------------------------------|----------------------------|--------------------|---------------|
| Go | 116 | 2 | 2 | EG / VET |
| Python | 12 | 16 | 34 | AIR / PEP8 |
| JavaScript | 12 | 15 | 32 | AIR / ES |
| C | 14 | 11 | 35 | AIR / CERT |
| C# | 12 | 15 | 26 | AIR / DGN |
| Rust | 14 | 13 | 34 | AIR |
| Kotlin | 12 | 15 | 30 | AIR / DGN |

Shared AIR structural templates: **72** (`summary.rules`, same for every language).

> Note: each sample is a small file with typical smells (naming violations, large
> literals, suspicious structures). "Rule classes hit" is the number of distinct
> rule IDs that sample triggered; different samples yield different numbers. This
> table shows "each language has triggerable rules", not full-spec coverage.

## How to read this (honestly)

**Go's 116 dedicated rules vs only 2 hits** — this is not "Go is weak"; the opposite:

- Go rules are the `rules_vet*` / `rules_eg*` / `rules_crc*` series, modeling go vet /
  Effective Go / CodeReviewComments — **precise analyzers for specific bug patterns**.
  The quantification sample only has naming violations (`foo_bar`, `BadName`), so it
  cannot trigger vet rules that need specific semantic context; only `GO-EG-001` and
  `GO-VET-043` fired.
- In short Go is "few but deep, needs context"; the other six are "shared templates,
  broadly triggered".

**11–16 classes for the other six** — mostly from the 72 shared AIR templates
(AIR-001/002/.../050), plus 12–14 dedicated rules each (PEP8 / ES / CERT / DGN).
These are **structural smells** (naming, literals, suspicious patterns); many hits
mean broad template coverage, not deep semantic analysis.

## Rule families

| Family | Language | Source focus |
|--------|----------|--------------|
| GO-VET / GO-EG / GO-CRC | Go | go vet analyzers / Effective Go / CodeReviewComments |
| PY-AIR / PY-PEP8 | Python | shared AIR structure / PEP 8 naming & style |
| JS-AIR / JS-ES | JavaScript | shared AIR structure / ES idioms |
| C-AIR / C-CERT | C | shared AIR structure / CERT C secure coding |
| CS-AIR / CS-DGN | C# | shared AIR structure / .NET design guidelines |
| RS-AIR | Rust | shared AIR structure (dedicated rules lean naming/style) |
| KT-AIR / KT-DGN | Kotlin | shared AIR structure / Kotlin design conventions |

## Honest scope statement

1. This is "how many rules can be triggered", **not "language spec coverage"**. AIR
   rules focus on naming/style/common structural smells, not type systems, concurrency
   safety or framework-specific API misuse.
2. **Go's depth hides in context**: few sample hits ≠ weak; dedicated samples are needed
   to trigger the vet series.
3. **Numbers vary by sample**: this table is one measured run with fixed samples; other
   samples give other hit counts — rerun the script for current values.
4. **No equal-capacity promise across languages**: differing rule counts/families reflect
   each ecosystem's conventions, not a target of "N rules per language".

## Reproduce

```bash
bash scripts/language-coverage.sh                       # print the table (live build + measure)
bash scripts/language-coverage.sh > /tmp/coverage.md    # save as Markdown
```

Rule counts can be checked by hand:

```bash
for d in go py js c cs rs kt common; do
  printf "%-8s " "$d"
  grep -rhoE 'ID:[[:space:]]*"' lang/rule/builtin/$d | wc -l
done
```
