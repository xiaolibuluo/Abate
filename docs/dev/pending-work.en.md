<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->

> Language / 语言：[English](pending-work.en.md) | [中文](pending-work.md)

# Abate Project · Completion Status Archive (formerly the "full pending-work list")

English translation of `pending-work.md`, the section-by-section companion
under the plan-A side-by-side naming scheme; section order and tables
correspond one-to-one; the pending/planning tone is preserved and nothing is
translated into "already done"; no new conclusions are added; this document
does not constitute legal advice.

> Verified: 2026-11-14 (GMT+8)
> Source: the 2026-11-11 18:20 todo snapshot (then marked "engineering
> implementation about 87%").
> This document was rewritten after item-by-item re-verification against the
> actual source code and artifacts; each item gives a conclusion, evidence
> path, and remaining boundary.
> **No fabrication**: unfinished items are kept as-is, never masked with
> "completed" wording.

---

## I. Overview

| # | Item | Priority | Conclusion | Key evidence |
|---|------|--------|------|---------|
| 1 | libclang-min real build | P0 | ⚠️ artifact-level achieved; functional-level loop not closed | `.artifacts/libclang-min/build/lib/libclang-18.so.18` |
| 2 | WASM host (`plugin/wasm`) | P0 | ✅ done | 32 .go files + wazero v1.12.0 vendoring |
| 3 | gVisor `pkg/abi` trimmed port | P1 | ✅ done | `kernel/abi/` + `docs/dev/abi-port-parity.md` |
| 4 | `jsfront/jsmap.go` split | P1 | ✅ done | all 3 files ≤300 lines |
| 5 | Final nine-section gate + `-race` + govulncheck | P0 | ✅ done | `CI GATE: PASSED`, auxiliary gates all green |
| 6 | Documentation consolidation | P2 | ✅ done this round | this document + `CHANGELOG.md` |
| 7 | SBOM + Release pipeline real run | P2 | ⬜ not executed (not a gate item) | scripts ready; no artifacts from this round |
| 8 | Final completion-review report filed | P0 | ⬜ not filed (verdict in §V) | standalone report file missing |

**Conclusion**: engineering implementation and the frozen-gate caliber
(including the degrade-rate single-file caliber) are fully met (items 1–6,
where item 1's functional-level verification is a known boundary);
the iron-rule caliber (functional-level verification) and the full-corpus
caliber are stated separately — see §V of this document and dimension 7 of
`docs/i18n/QA-ANTI-FRAUD-AUDIT.md`: the degrade rate "meets threshold" only
under the frozen single-file caliber, and under the same formula all six
languages exceed the threshold on the full corpus; items 7 and 8 are
procedural wrap-ups that enter no CI gate and are honestly listed as
remaining items.

---

## II. Item details

### 1. libclang-min real build (P0) — ⚠️ artifact-level achieved; functional-level loop not closed

**Achieved**
- `libclang-18.so.18` really compiled from the LLVM 18.1.3 sources
  (79,650,944 bytes, 2026-11-13 18:55); all `ldd` dependencies resolve to
  system `libstdc++ / libm / libgcc_s / libc`.
- Static libraries `libclang*.a` / `libLLVM*.a` produced in the same
  directory (incl. `libclangAST.a` 22MB, `libclangSema.a` 30MB, etc.).
- Build configuration traceable: `.artifacts/libclang-min/build/CMakeCache.txt`
  contains `CMAKE_HOME_DIRECTORY=/tmp/llvm-src/llvm-toolchain-18-18.1.3/llvm`.

**Not closed (stated honestly)**
- `build-report.txt` and `build.log` still hold the 2026-11-11 failure
  snapshot (round 9 `[563/2380]` compiler OOM); not updated with the
  successful 11-13 build.
- The original acceptance criteria's `libclang-min.a` static-library artifact
  and the C API functional-level verification (a minimal C program using
  `clang_getCXCursor`-style calls) **were not produced**; the source has no
  CGO bridge, and `lang/rule/builtin/c/doc.go` states explicitly that AIR
  does not expose clang/libclang types.
- Build artifacts are platform-bound; they have been removed from version
  control and added to `.gitignore` (see §IV).

**Verdict**: the artifact really exists and is parseable, but the iron-rule
caliber of "functional-level verification" is not satisfied and is not
counted toward iron-rule completion.

### 2. WASM host implementation (P0) — ✅ done

- `plugin/wasm/` has **32 .go files** in total (WAT lexing / parsing /
  encoding, module loading, fuel/gas limits, deterministic clock, host
  function binding, runtime); `go test ./...` all green.
- wazero v1.12.0 fully vendored: `go.mod`, `vendor/modules.txt`,
  `third_party/_sources/wazero/`, verified via `go mod verify` and the
  non-GPL verdict of `license-check`.
- The old list's "0 .go files / wazero not yet vendored" was the stale
  11-11 state and is obsolete.
- The dependency-surface gate `scripts/m5-reference-scan.sh` stage ③ was
  changed from the outdated "yaml.v3-only" snapshot to a bidirectional
  consistency check with `go.mod` as the authority (commit `cbedd1fc`),
  eliminating false positives on wazero / x/sys.

### 3. gVisor `pkg/abi` trimmed port (P1) — ✅ done

- `kernel/abi/`: syscall number→name tables (amd64 / arm64), errno,
  `Addr` / `AddrRange` / `PageSize`, architecture constants, and
  `syscall_test.go`. Rewritten in the Abate style; does not import upstream.
- `docs/dev/abi-port-parity.md` records the item-by-item parity check:
  amd64 **352 entries**, arm64 **295 entries**, errno 131 + 2 + 4, matching
  the upstream in `third_party/_sources/gvisor/` item by item; the test
  reads the upstream `linux64.go` directly during `go test`, so upstream
  changes expose drift immediately.

### 4. `jsfront/jsmap.go` split (P1) — ✅ done

- The original 384-line single file was split into `jsmap.go` (106) /
  `jsmap_children.go` (118) / `jsmap_node.go` (193), all ≤300 lines;
  `style-check.sh` and `header-check.sh` pass.

### 5. Final gate full real run (P0) — ✅ done

- `scripts/lint-gate.sh`: `CI GATE: PASSED` (header / style / license / gofmt
  / vet / golangci-lint **0 issues** / constant-lint / `go test -race` /
  coverage / build).
- Auxiliary gates: `determinism-check` PASSED, `drift-check` PASSED,
  `pit-parity` **100%** (i.e. 21/21 comparable points consistent, see
  `docs/dev/pit-parity-report.md`), `m5-reference-scan` PASSED, `vuln-scan`
  **PASS**, `go mod verify` all modules verified.
- TLA+ `scripts/tla/check.sh`: PitParity 44 / SandboxFuel 61 /
  StreamDeterminism 4800 states, all 3 specs pass.
- **GO-2026-6859** (`net/http` MaxBytesReader header race, affecting
  go1.26.10, fixed in 1.26.11): the toolchain was uniformly upgraded to
  **Go 1.26.11**; the real `govulncheck` result is
  `No vulnerabilities found` (commit `cbedd1fc`). The old list's
  "govulncheck pending final run" is closed.
- Coverage maintained at the existing level (mutation 98.3% / pitparity
  95.5% / sandbox 96.2%).

### 6. Documentation consolidation (P2) — ✅ done this round

- This document was rewritten against the actual state, removing stale
  statements such as "WASM not started" and "wazero not vendored".
- The `CHANGELOG.md` `[v1.0.0]` section corrected the statement "the gVisor
  platform-level sandbox and a true WASM runtime remain undelivered
  boundaries", which contradicted the code's actual state, and recorded this
  round's incremental wrap-up (version number unchanged; the drift gate is
  unaffected).
- `docs/dev/frontend-coverage/` six-language coverage lists and
  `docs/dev/pit-parity-report.md` are both in a produced state; cross
  references unchanged.

### 7. SBOM + Release pipeline real run (P2) — ⬜ not executed

- `scripts/sbom.sh`, `scripts/release.sh`, `scripts/plugin-publish.sh` are
  all ready; there are **no** SBOM / release artifacts from this round under
  `.artifacts/`; directories such as `plugin-publish/` hold acceptance
  evidence from earlier rehearsals.
- This item does not enter the nine-section gate and is an optional
  release-process wrap-up; "pending execution" is kept honestly, not marked
  done.

### 8. Final completion-review report filed (P0) — ⬜ not filed

- The standalone final-review report has not been generated; §V of this
  document gives the immediate dual-caliber verdict as the substitute
  conclusion.
- The original list's procedural condition "the final review must pass
  before 100% completion may be claimed" has not been closed in the form of
  a standalone document.

---

## III. Closed loop (formerly "dependencies and order")

```
[2 WASM host] ──┐
[3 gVisor port] ├─→ [5 final gates] ─→ [6 doc consolidation] ─→ engineering/gate caliber 100%
[4 jsmap split] ─┤        ↑
[1 libclang artifact]┘   [TLA+ / 5 auxiliary gates]
        │
        └── ⚠️ functional-level verification not closed (iron-rule caliber remaining boundary)

Remaining (non-gate items): [7 SBOM / Release real run] · [8 standalone final-review report filed]
```

---

## IV. Completed baseline (for the final review to compare against)

| Item | Status |
|------|------|
| Six categories of upstream source vendoring (gVisor / PIT / LLVM / Roslyn / Kotlin / tree-sitter family + syn / swc_ecma_parser) | ✅ `third_party/_sources/` |
| tree-sitter grammar data + runtime functional-level verification (gcc-compiled Python parsing) | ✅ sha256 matches upstream |
| PIT 1.19.6 official parity check | ✅ 28 mutations → 21 comparable → 21 consistent → **100%** (`.artifacts/pit-parity.json`) |
| `kernel/mutation/pitparity` (15 .go) + `plugin/sandbox` (14 .go) | ✅ delivered and re-verified |
| Six-language frontend coverage lists (`docs/dev/frontend-coverage/`) | ✅ produced |
| `lang/pyfront` pyAsPower prefix symbol | ✅ fixed (`pytokens.go:26`) |
| Codebase TODO/FIXME residue | ✅ 0 occurrences |
| GitHub full upload (xiaolibuluo/Abate) | ✅ 11-11 snapshot uploaded; later local commits not pushed (no remote) |
| **wazero v1.12.0 vendoring + `plugin/wasm/` true WASM runtime (32 .go)** | ✅ delivered |
| **`kernel/abi/` gVisor trimmed port + upstream item-by-item parity** | ✅ amd64 352 / arm64 295 / errno 137 |
| **`lang/jsfront/jsmap.go` split (≤300 lines)** | ✅ all 3 files meet the limit |
| **Cleared all lint hits (7 high-complexity + 8 consequential + 6 over-limit file splits)** | ✅ `golangci-lint` 0 issues (`b69f2c1c`) |
| **Nine-section gate + TLA+ 3 specs + 5 auxiliary gates all green** | ✅ `CI GATE: PASSED` |
| **Go 1.26.11 upgrade fixing GO-2026-6859** | ✅ `vuln-scan` PASS (`cbedd1fc`) |
| **Build-artifact cleanup**: 1013 `libclang-min/build/` files removed from version control and added to `.gitignore` | ✅ this round (workspace files kept; local builds unaffected) |

---

## V. Completion verdict (dual caliber, stated honestly)

| Caliber | Conclusion |
|------|------|
| **Engineering implementation + gate verification** | **100%**: items 1–6 closed; the nine-section CI gate, TLA+ 3 specs, determinism / drift / pit-parity(100%) / m5 / vuln-scan / `go mod verify` all green, `golangci-lint` 0 issues, 59 packages `go test -race` 0 FAIL |
| **Iron-rule caliber (real vendoring + functional-level verification)** | **100% not claimed**: the `libclang` C API functional-level verification and the `build-report.txt` update are not closed; the functional-level verification of the six upstream vendorings + wazero + gVisor abi are closed |

**Remaining boundaries (fabrication forbidden)**
1. `libclang-18.so.18` was really compiled, but CXCursor traversal / CGO
   bridging / report updates are not filed.
2. The `SBOM` + `Release` pipelines were not really run this round (scripts
   ready, no artifacts).
3. The standalone final-review report is not filed (§V of this document is
   the immediate verdict).
4. `.artifacts/libclang-min/build/` was removed from the index, but history
   commits still contain these artifacts; fully slimming history requires a
   separate `git filter-repo` operation (rewriting all history hashes), not
   executed this time.

**To complete the iron-rule caliber 100%**, the suggested order is:
① libclang C API minimal functional-level verification + rewrite
`build-report.txt`;
② real runs of `sbom.sh` / `release.sh` with artifacts retained;
③ file the standalone final-review report.
