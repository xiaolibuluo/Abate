<!--
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
-->

# Seven-Language Coverage Quantified (Measured)

> Data is **generated live** by `scripts/language-coverage.sh` (builds the binary from source, then
> measures), not estimated. Re-run: `bash scripts/language-coverage.sh`. Script output is authoritative.

## Measured table

| Language | Dedicated rules (source `ID:`) | Rule classes hit (sample) | Hits | Rule families |
|---|---|---|---|---|
| Go | 116 | 2 | 2 | EG / VET |
| Python | 12 | 16 | 34 | AIR / PEP8 |
| JavaScript | 12 | 15 | 32 | AIR / ES |
| C | 14 | 11 | 35 | AIR / CERT |
| C# | 12 | 15 | 26 | AIR / DGN |
| Rust | 14 | 13 | 34 | AIR |
| Kotlin | 12 | 15 | 30 | AIR / DGN |

Shared AIR structural templates: **72** (`summary.rules`, identical across languages).

> Note: each sample is a small file with typical smells (naming violations, large literals, suspicious
> structures). "Rule classes hit" is the number of distinct rule IDs the sample triggered; different
> samples yield different numbers. The table shows "each language has triggerable rules", not full
> rule-set coverage.

## How to read this data

**Go's 116 dedicated rules vs only 2 classes hit in the sample** — this is not "Go is weak"; it is two
different depths. Go rules are the `rules_vet*` / `rules_eg*` / `rules_crc*` families, matching go vet /
Effective Go / CodeReviewComments — **precise analyzers for specific bug patterns**. The quantification
sample only has naming violations (`foo_bar`, `BadName`), which cannot trigger vet rules needing
specific semantic context; it hits only `GO-EG-001` and `GO-VET-043`.
Go follows a "**few and precise, context-dependent**" path; the other six follow "**shared structural
templates, broadly triggered**".

**The other six hit 11–16 classes** — mainly from the 72 shared AIR templates (AIR-001/002/.../050)
plus 12–14 dedicated rules each. These are **structural smells** (naming, literals, suspicious patterns);
many hits mean broad template coverage, not deep semantic analysis.

## Rule family meanings

| Family | Language | Source focus |
|---|---|---|
| GO-VET / GO-EG / GO-CRC | Go | go vet analyzers / Effective Go / CodeReviewComments |
| PY-AIR / PY-PEP8 | Python | shared AIR structure / PEP 8 naming and style |
| JS-AIR / JS-ES | JavaScript | shared AIR structure / ES idioms |
| C-AIR / C-CERT | C | shared AIR structure / CERT C secure coding |
| CS-AIR / CS-DGN | C# | shared AIR structure / .NET design guidelines |
| RS-AIR | Rust | shared AIR structure (dedicated rules lean naming/style) |
| KT-AIR / KT-DGN | Kotlin | shared AIR structure / Kotlin design conventions |

## Honest coverage statement

1. This is "how many rules can be triggered", **not "language spec coverage"**: AIR rules focus on
   naming/style/common smells, not type systems, concurrency safety or framework API misuse.
2. Go's depth hides in context: few sample hits does not mean weak; dedicated samples are needed to trigger vet rules.
3. Numbers vary with samples: this table is one fixed-sample measurement; re-running reproduces current values.
4. No cross-language parity promise: differing rule counts/families fit each ecosystem's conventions.

## Full-corpus fidelity (a different scope — do not mix with the table above)

Besides rule-trigger quantification, `scripts/fullcorpus-degrade.sh` parses all 20520 self-developed
corpus entries live and computes the full-corpus degradation rate with the E2E formula:
py 3.36% / js 3.51% / c 4.80% / cs 4.65% / rs 4.91% / kt 4.48%, Go `NOT_MEASURED` (no embedded
grammar mappings). All six are under 5%; this scope **does not reach** the single-file frozen threshold
and is disclosed as fact only — not a "high fidelity/passing" claim.

## Reproduction

```bash
bash scripts/language-coverage.sh                      # print the table (live build + measurement)
bash scripts/fullcorpus-degrade.sh                     # full-corpus degradation rate
for d in go py js c cs rs kt common; do printf "%-8s " "$d"; \
  grep -rhoE 'ID:[[:space:]]*"' lang/rule/builtin/$d | wc -l; done   # count rules manually
```
