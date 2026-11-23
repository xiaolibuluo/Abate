<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->

> Language / 语言：[English](CHANGELOG.en.md) | [中文](CHANGELOG.md)

# Changelog

All notable changes to Abate are documented here. The format follows a
milestone-based record; released engine versions use `v<year-ring>.<minor>.<patch>`
per the XCU-REL version policy (kernel year-ring + hourly ecosystem cadence).

Version-number consistency is enforced by CI: the `ldflags` version, this
CHANGELOG entry and the OpenAPI `info.version` must agree (drift → FAIL).

## [v1.2.29]

Patch release — fix nine zero-yield C# constructs, coverage 46.0% → 52.6% (2026-11-23).

- **C# frontend zero-yield fixes (all nine flipped to exact)**: `extern_alias_directive`→`CsExternAlias`,
  `labeled_statement`→`Label`, `ref_expression`→`Ref`, `expression_statement`→`ExprStmt`,
  `array_creation_expression`/`initializer_expression`→`CsObjectCreation` (rule retargeting),
  `cast_expression`→`Cast` (parenPrefix cast detection + `isCastOperandStart` helper),
  `spread_element`→`Spread` (direct `..` handling in `reduce()`, `..` added to the operator
  table with PrefixPower=4), `type_parameter`→`GenericParam` (`csparser.genericParamMarkers()`
  scanning the declaration header `<...>` span). Changes land in `lang/csfront/`
  (csparser/csexpr/csexpr_prefix/csmap/cstokens/csexpr_dedicated) and
  `lang/frontkit/grammar/rules_csharp.go` (nine degraded→exact rules).
- **Measured coverage gain (`docs/dev/frontend-coverage/csharp.md`)**: covered 99→113 of 215
  constructs, exact mappings 92→100, distinct AIR kinds hit 65→72. The degraded count moving
  12→13 is the expected side effect of newly added generic corpus samples activating existing
  degraded rules; the six-language `degradedaudit` stays at zero, so no new degradation was
  introduced.
- **Regression checks**: all `lang/` frontend packages (air/cfront/csfront/frontkit/jsfront/
  ktfront/pyfront/rsfront) and `frontkit/grammar` tests pass; `ci/gate_e2e.sh` passes all seven
  stages (7 languages, 1540 corpus entries, zero failures, byte-identical double-run with the
  same seed); `go build ./...` and `go vet ./lang/...` both exit 0.
- **Version bump**: `abate version` emits `abate 1.2.29`, `MetaVersion` synced to `1.2.29`;
  README zh/en badges and user docs (quickstart / cli-reference / multilang-check-guide, both
  languages) version examples updated accordingly.

## [v1.2.28]

Patch release — ship the post-v1.2.27 usability fixes in the binaries (2026-11-19).

- **check help text now shipped**: after v1.2.27, source replaced the shared
  "steady-state violated" exit-code help with a check-specific one (0 clean /
  1 findings incl. warning / 2 internal error / 3 usage error, filterable via
  `--severity error`). The previous binary lacked this text; after rebuilding
  all 5 targets, the downloaded binary **shows the new help**. Report output,
  rule sets, AIR and deterministic fingerprints are unaffected (help text is
  not a report artifact).
- **usability doc fixes shipped**: README quick-start copy-paste seven-language
  check examples, `GOFLAGS=-mod=vendor`→`GOPROXY=off` correction and T01 version
  sample `0.1.0`→`1.2.27` — documentation-only fixes already in source, now
  paired with a matching binary release.
- **Version bump**: `abate version` emits `abate 1.2.28`, `MetaVersion` synced
  to `1.2.28`; user docs (quickstart / cli-reference / multilang-check-guide /
  T01, both languages) version examples and README zh/en badges synced. Watermark
  `VersionLabel` / three SHA anchors unchanged; `SHA256SUMS` / `WATERMARK.sha256`
  regenerated per established practice due to the README badge change.
- **Republished**: 5 binaries + `manifest.txt` + SPDX SBOM + 7 detached GPG
  signatures, tagged and signed as `v1.2.28` and set as Latest; v1.2.27 /
  v1.2.26 / v1.0.1 / v1.0.0 tags and releases kept unmodified. Details in
  release readiness doc §十六.
- **Credibility review & honesty fixes (post-release docs, binary untouched)**:
  two external claims were checked empirically. "Seven languages are just regex"
  was disproven — `lang/frontkit/parser.go` is a real Pratt parser and rules match
  structurally on the AST (nested-structure probe hit multiple rule classes) —
  while the **error-tolerant, non-compiler-grade** boundary was explicitly
  acknowledged. "Go version contradiction" was confirmed — `go.mod`/local are
  both `1.26.11`, README/T01's self-contradicting "old values" phrasing was
  replaced with a neutral single source of truth. Added
  `docs/user/capabilities-and-limits.md` / `.en.md` with reproducible evidence,
  a README boundary note, and changed the GitHub description to "structural
  linter ... not a compiler-grade analyzer". Details in §十七.
- **Second-round credibility fixes (evidence/quantification/repro; docs & tests,
  binary untouched)**: added `scripts/evidence-check.sh` (live build + 7 asserted
  reproductions, exit 0 only if all pass) and `scripts/language-coverage.sh`
  (measured per-language rule/hit table + `docs/user/language-coverage.md`/.en.md).
  Capabilities doc gained a "community validation status" section stating 0 stars /
  0 forks honestly and refusing to fake it. Full `go test ./...` fixed a real
  regression introduced earlier (`TestCheckHelpContent` stale help assertion) and
  added two regression tests for error-tolerant parsing and multi-class AST hits;
  final run 0 FAIL. Details in §十八.

## [v1.2.27]

Patch release — make the published binaries actually carry seven-language check (2026-11-19).

- **Fix the code/release timing gap**: the v1.2.26 seven-language wiring source
  and docs were committed at 02:18, but the release binaries were built at 01:13
  and uploaded at 01:14, so the 5 downloadable binaries were still the old
  Go-only build (`check` help read "Run the built-in Go rule set"). This version
  rebuilds all 5 targets from current source; `abate check` is verified to
  dispatch 7 languages by extension (`.go/.py/.js/.c/.cs/.rs/.kt`), with
  `--languages` filtering, JSON/SARIF output and mixed-directory checks all
  passing.
- **Version bump**: `abate version` emits `abate 1.2.27`, `MetaVersion` synced
  to `1.2.27`; user docs (quickstart / cli-reference / multilang-check-guide,
  both zh and en) version examples updated. Watermark `VersionLabel` / three
  SHA anchors left unchanged; `SHA256SUMS` / `WATERMARK.sha256` were regenerated
  and synced per established practice after the README hero enhancement, with
  all hash gates passing.
- **Republished**: 5 binaries + `manifest.txt` + SPDX SBOM + 7 detached GPG
  signatures, tagged and signed as `v1.2.27` and set as Latest; v1.2.26 /
  v1.0.1 / v1.0.0 tags and releases kept unmodified.
- **Repository discoverability**: description leads with the "seven-language
  static analysis / CLI" search intent; topics expanded from 10 to 20 (all 7
  languages — python/javascript/golang/rust/kotlin/c-sharp/c — plus
  static-analysis/code-analysis/linter/code-quality/ast/sarif/cli/developer-tools
  and chaos-engineering/mutation-testing long-tail terms, selected by measured
  repository counts via GitHub API); README zh/en heroes gained version /
  license / seven-language / Go badges and a value proposition. See release
  readiness doc §十四.
- **Post-release usability fixes (source/docs, not in the v1.2.27 binary)**:
  README quick start now includes a copy-paste seven-language check example
  with an upfront exit-code note; README/T01 build command changed from
  `GOFLAGS=-mod=vendor` to `GOPROXY=off` to match the zero-dependency reality,
  and T01 version sample `0.1.0`→`1.2.27`; `check --help` exit codes switched
  from the shared steady-state wording to check-specific text (warnings also
  return 1; `--severity error` filters them). These are presentation/doc-only
  changes — report output and determinism are unchanged; `SHA256SUMS` /
  `WATERMARK.sha256` regenerated. Note: the released v1.2.27 binary still has
  the old generic exit-code help, carried into the next version. See §十五.

## [v1.2.26]

Release version bump with rebuilt binaries and public repository publication.

- **Version bump**: user-facing runtime version constants unified to `1.2.26` —
  `abate version` now emits `abate 1.2.26`, and the health/OpenAPI metadata
  `MetaVersion` is synced to `1.2.26`. Frozen artifacts (watermark
  `VersionLabel` / three SHA anchors) are left unchanged per the established
  freeze policy, so hash gates stay intact.
- **Rebuilt binaries and release**: 5-target cross-compiled binaries
  (linux/darwin/windows × amd64/arm64) + `manifest.txt` + SPDX SBOM rebuilt,
  7 detached GPG signatures verified 7/7, `sha256sum -c` manifest check
  passed, uploaded via GitHub Release.
- **Public repository**: visibility changed from private to public, with
  description / topics / homepage settings aligned; merge-protection ruleset
  remains active.
- Marked by a separate signed tag `v1.2.26`; prior v1.0.0 / v1.0.1 tags kept
  unmodified.
- **Seven-language check wiring**: `abate check` now dispatches by file
  extension to 7 language frontends (`.go/.py/.js/.c/.cs/.rs/.kt`, plus
  `.jsx/.mjs/.cjs/.h/.kts`), runs each language's in-house rule set and emits a
  unified console/JSON/SARIF report. Added `--languages` filtering and a generic
  `check.MultiChecker` (language-neutral dispatch + aggregation reusing the
  existing deterministic pipeline), with guard tests for 7-language dispatch,
  mixed directories, determinism, filtering and extension routing. Step-by-step
  guide: `docs/user/multilang-check-guide.en.md`. Frozen artifacts untouched.

## [v1.0.1]

Patch release — post-GA real quality consolidation and formal signed release
wrap-up (2026-11-19).

- **Full-corpus degradation closed out**: 20520 self-authored corpus entries
  parsed one by one; all six languages below 5% (py 3.36% / js 3.51% / c 4.80%
  / cs 4.65% / rs 4.91% / kt 4.48%, go `NOT_MEASURED`). Every exact promotion
  is backed by real frontend dedicated-kind production (minimal sample + real
  corpus double guard tests); remaining degraded entries are kept honestly.
  The frozen single-unit caliber is unchanged.
- **Quality consolidation**: parse benchmarks for the js/c/cs/rs/kt frontends
  (deterministic source generation + real `Parse` + non-zero `NodeCount` guard);
  cfront Go-native fuzz added and a real fuzz-discovered panic fixed (negative
  cap on unterminated `template <`), with a regression lock. Full
  `go test ./...` (73 packages) passed, `golangci-lint` 0 issues.
- **libclang C API functional-level verification closed**: `scripts/libclang-api-verify.sh`
  really calls the self-built `libclang-18.so.18` (ldd asserts self-built
  linkage), diagnostics=0, CXCursor walk recovers 30 cursors / struct Point 1 /
  functions add+main 2, rc=0. Iron-rule caliber 100% closed.
- **SBOM / Release pipeline real runs**: fixed `sbom.sh` crash on missing
  manifest, SPDX 2.3 emitted honestly; `release.sh --dry-run v1.0.0` 5-target
  cross-compilation + manifest 5/5 OK + reproducible dual-build SHA256.
- **Formal signed release**: after rights-holder authorization, 7 GPG-signed
  commits pushed to `origin/main`, signed tag `v1.0.0` synced, 7 release
  artifacts detached-signed with 7/7 verification passed; this `v1.0.1` marks
  the wrap-up with a separate signed tag (the contemporary final-readiness report is archived).

## [v1.0.0]

GA release. Closes the M1–M6 program.

- **Kernel (L0)**: the four perturbation-algebra operations and algebraic-law
  checks, deterministic scheduler, observation pipeline, verifier, mutation
  core (PIT algorithm port), injection core (chaosblade source merged in).
- **CES (L2)**: canonical schema / parsing / validation / TCK in place; 164 TCK
  cases (parse 68 / validate 46 / roundtrip 50), including a sample
  third-party conforming implementation.
- **Simulator (L3)**: virtual clock / deterministic PRNG / virtual network /
  virtual filesystem / virtual scheduling / record-and-replay (frame hash
  chain + full-tier SHA-256).
- **Language frontends (L4)**: AIR + seven-language frontends for
  Go/Python/JS/C/C#/Rust/Kotlin + rule engine + in-house checker (SARIF) +
  operator derivation.
- **Control plane (L7)**: CLI / REST / Web console; REST enforces Bearer
  authentication and refuses to start on a remote listener without a token.
- **Security (XCU-SEC)**: threat model T1–T10 established; executable attack
  case sets for T1/T6/T7/T8/T10 (including -race) in CI; signature chain /
  report escaping / archive tamper resistance verified.
- **Security increments (T2/T3/T9 landed)**: added the `security/redact`
  deterministic redaction middleware and wired it into the two log points in
  `ctl/rest/middleware.go` (panic log and access log); added three executable
  attack cases under `security/tests`:
  `TestThreatT2CapabilityEscalation`, `TestThreatT3ResourceExhaustion`,
  `TestThreatT9SecretRedaction`; updated the threat-model alignment check to
  10/10. The gVisor platform-level sandbox (`kernel/abi/`) and the real WASM
  runtime (`plugin/wasm/` + wazero) were delivered incrementally after GA;
  see "Post-GA incremental wrap-up" below. The version number was not changed;
  the single source of truth stays consistent.
- **Governance (XCU-GOV)**: copyright-header gate + watermark + dual
  licensing + third-party source incorporation registry.
- **Release (XCU-REL)**: `release.sh` reproducible build matrix + checksum
  manifest (watermark) + SBOM (SPDX 2.3) + GPG signing rehearsal.

### Post-GA incremental wrap-up (2026-11-13 ~ 2026-11-14)

The following are incremental deliveries and governance wrap-up after v1.0.0;
**the version number is unchanged** and the single source of truth stays
consistent.

- **Real WASM runtime**: `plugin/wasm/` contains 32 `.go` files (WAT lexing /
  parsing / encoding, module loading, fuel/gas limits, deterministic clock,
  host function binding), with wazero v1.12.0 fully vendored
  (`go.mod` / `vendor/modules.txt` / `third_party/_sources/wazero/`).
  Corrects the earlier boundary statement that "the real WASM runtime has not
  been delivered".
- **gVisor ABI-trimmed port**: `kernel/abi/` syscall tables (amd64 352 / arm64
  295), errno, hostarch primitives; `syscall_test.go` reads upstream values at
  runtime for an item-by-item cross-check (the parity report is archived).
  Corrects the earlier statement that "the gVisor platform-level sandbox has not been delivered".
- **Lint fully cleared**: 7 high-complexity production functions and 6
  over-limit test functions were split into single-responsibility helpers,
  removing 8 knock-on findings; 9 companion source files were split out at
  ≤300 lines each; `golangci-lint` reports **0 issues** (revive with
  enable-all-rules on).
- **Security upgrade**: Go 1.26.10 → **1.26.11**, fixing GO-2026-6859
  (`net/http` MaxBytesReader header race); `govulncheck` actually run and
  reported `No vulnerabilities found`.
- **Rule precision**: section ③ of `m5-reference-scan.sh` changed to a
  `go.mod`-authoritative bidirectional vendor/require consistency check;
  `license-check.sh` gained standard BSD legal-text signature phrases.
  Neither check was loosened; no whitelist or skip was introduced.
- **All gates green**: all nine CI gate sections `PASSED`; 3 TLA+ specs;
  determinism / drift / pit-parity(100%) / m5-reference-scan /
  `go mod verify` all passed.
- **Repository governance**: 1013 CMake/Ninja build artifacts under
  `libclang-min/build/` were removed from version control and added to
  `.gitignore` (workspace files retained); the completion-status document was
  rewritten to reflect the as-is state (historical process docs are archived).
- **All-language degradation wrap-up (2026-11-18)**: all 20520 self-researched
  corpus entries re-parsed individually; py 3.36% / js 3.51% / c 4.80% /
  cs 4.65% / rs 4.91% / kt 4.48%, with go remaining `NOT_MEASURED` (no embedded
  grammar mappings). All six languages are below 5%. Every exact mapping is
  backed by a genuine dedicated-kind frontend emission (dual minimal-sample and
  real-corpus guards); residual degraded entries are shared-kind carry-overs or
  zero-production frozen-corpus constructs and stay degraded honestly — no
  labels were rewritten.
- **Quality-dimension consolidation (2026-11-18)**:
  - The P8 `frontend parse throughput` budget is owned by `LNG-*`, but only the
    go/py frontends had parse benchmarks. Added parse benchmarks for the js/c/
    cs/rs/kt frontends (`*_bench_test.go`) with deterministic source generation,
    real `Parse`, a `NodeCount` non-zero guard, and reported MB/s and nodes/op —
    no performance hard assertions. Measured: js 1.23 / c 1.19 / cs 0.73 /
    rs 2.85 / kt 0.80 MB/s.
  - Added Go-native fuzzing for cfront (`cfront_fuzz_test.go`), covering the
    input space beyond the fixed chaos corpus. **Fuzzing found and fixed a real
    panic**: an unclosed `template <` made `cAngleRange` return an inverted
    interval and `cLeafNodes` call `make` with a negative cap; added inverted-
    range guards to cfront's `cLeafNodes`/`cExprReducer.reduce` and csfront/
    ktfront's `leafNodes`, with a `TestCFuzzUnclosedTemplateRegression` lock.
    20s / 200k fuzz executions, zero crashes.
  - Full `go test ./...` (73 packages) passed, `golangci-lint` 0 issues,
    lint-gate ten sections all green.

## [M6] GA wrap-up (before v1.0.0)

- CES-TCK released and the sample third-party conformance process walked
  through.
- Security audit gates (signature chain / sandbox escape / dependency
  vulnerability scan) all green.
- Full documentation set (user / developer / specification / ADR) in place.
- Release engineering rehearsal (tag + assets + checksums + SBOM).

## [M5] Seven languages + control plane

- LNG-PY/JS/C/RS/CS/KT all landed; CTL-REST online; HTML/JSON/JUnit/OTel
  reports.

## [M4] WASM runtime + minimal ecosystem loop

- WAS-HOST/CAP/HOT/SIG four modules; plugin signature chain + capability
  sandbox; SARIF report validation.

## [M3] Go language frontend loop (dogfooding)

- LNG-AIR/GO/RULE/CHK/MUT complete; Go rule set ≥100 rules; repository
  self-check 0 violations.

## [M2] Simulator complete + dual cores merged

- SIM-NET/FS/SCHED/REC + KRN-MUT/INJ; PIT/chaosblade merged at source level.

## [M1] Minimal kernel loop

- KRN-ALG/SCH/OBS/VER + CES-SCHEMA/PARSE/VALID + SIM-CLK/RNG + CTL-CLI.
- 100 runs with the same seed produce byte-identical reports; the
  perturbation-algebra laws are fully covered.
