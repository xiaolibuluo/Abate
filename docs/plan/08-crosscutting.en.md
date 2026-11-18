<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->

> Language / 语言：[English](08-crosscutting.en.md) | [中文](08-crosscutting.md)

# Abate Crosscutting Systems Plan — Seven Systems in Extreme Detail

**Design-target statement (mandatory for C-class translations)**: this file is the English translation of an internal planning document. All `100%` / "pass" / "met"-style statements in the full text are **Design targets / Acceptance targets**, **not measured results**; they must not be read as established facts or capability claims. The translation corresponds section by section, does not escalate the tone, and does not delete target qualifiers.

- Document ID: PLAN-08　Status: in force　Upstream: PLAN-00
- Covered modules: XCU-SEC / XCU-GOV / XCU-CI / XCU-REL / XCU-TST / XCU-PRF / XCU-DOC
- Common constraints (observed by all seven systems):
  1. Crosscutting systems act on all L0–L7 layers but introduce no new layer dependencies (they are tools and processes, not runtime components; the only exception is XCU-SEC's runtime portion, which belongs to the respective layer modules).
  2. Crosscutting systems are themselves subject to the five-section gate chain (tool scripts are code too: gofmt/vet/lint/test/build all green) (acceptance target).
  3. Each system has a single owner and a quarterly review cadence (registered in the system README).

---

## System XCU-SEC: Security

### 1. System positioning and responsibilities

Defense in depth for the whole engine: threat modeling, signature chain, vulnerability management process, dependency supply-chain security, and key management. Runtime security mechanisms (capability sandbox / plugin signature verification) belong to the respective L1 modules; this system covers policies, processes, and audits.

### 2. File layout

```
security/
  THREAT-MODEL.md     // threat model document (quarterly review)
  policy/             // security policy files
    key-policy.md     // key management policy
    disclosure.md     // vulnerability disclosure process
  audit/              // audit tools
    scan.sh           // dependency vulnerability scan entry
  tests/              // security test case set
```

### 3. Threat model (must be fully established in the first version)

| ID | Threat | Attack surface | Mitigation | Owning module |
|---|---|---|---|---|
| T1 | Malicious plugin loading | plugin install channel | mandatory signature verification (unsigned refused) | WAS-SIG |
| T2 | Plugin privilege escalation (fs/net/elevation) | host function calls | capability sandbox deny-by-default + runtime gating + audit | WAS-CAP |
| T3 | Plugin resource exhaustion (memory/fuel) | plugin execution | three-way caps on fuel + memory + result length | WAS-HOST |
| T4 | Supply-chain poisoning (dependencies) | go.mod/third_party | source-level inclusion + commit hash pinning + license review | XCU-GOV |
| T5 | Signing-key compromise | release channel | offline key custody + rotation process + overlap period | XCU-REL |
| T6 | CES document injection (malicious expressions / excessive nesting) | document parsing | depth cap 64 + input cap 10 MB + fuzz testing | CES-PARSE |
| T7 | Report XSS (HTML / message injection) | report output | `html/template` automatic escaping + template-variable allowlist | RPT-HTML/ADP-* |
| T8 | Unauthorized REST access | API listener | mandatory Bearer authentication + startup refusal for remote listening without a token | CTL-REST |
| T9 | Token/key leakage into logs | global logging | masking middleware + scan gate (regex scan of log calls) | global |
| T10 | Archive tampering (forged replay evidence) | archive files | frame hash chain + full-archive SHA-256 | SIM-REC |

### 4. Key processes (must be implemented)

1. Vulnerability disclosure process: `security@` mailbox (`xiaolibuluo@163.com`, an independent security channel) → confirmation within 48 h → 90-day fix window → CVE request (for severity) → announcement publication. The disclosure policy text (`disclosure.md`) is public with the repository.
2. Dependency vulnerability scanning: periodic CI job (govulncheck brought in at source level + manual review list for third_party pinned versions); a new CVE hitting a third_party pinned version → review and backport process (aligned with KRN-INJ upstream backporting).
3. Key management: release signing keys held offline (hardware media); trust-anchor public keys updated with each release with published fingerprints (release announcement + repository `SECURITY.md`); rotation overlap period of 2 releases (old and new keys both accepted for verification).
4. Security audit cadence: at the end of each milestone (M1–M6), run the full security test case set (≥1 attack case for each of T1–T10); third-party security audit before GA (external).

### 5. Dependencies and constraints

- Depends on: modules at all layers (policy consumers).
- Forbidden: security policies diverging from implementation (every threat in the threat model must point at a concrete owning module and test case — verified by scan).

### 6. Test requirements

- ≥1 attack case per threat T1–T10 (negative: assertion that the attack is blocked); the case set goes into `security/tests/` and into CI.
- Log-masking scan: regex scan of all repository log calls (token/key patterns) with zero hits (acceptance target).

### 7. Acceptance criteria (quantifiable)

1. All 10 threats in the threat model have an owning module + mitigation + attack case (scan-verified 100% alignment) (acceptance target).
2. Security test case set all green in CI; log-masking scan with zero hits (acceptance target).
3. Vulnerability disclosure process documented publicly; key-rotation process rehearsed once (record archived) (acceptance target).
4. External security audit completed before GA (report archived) (acceptance target).

---

## System XCU-GOV: Governance and compliance

### 1. System positioning and responsibilities

Copyright and licensing governance: mandatory copyright headers, digital code watermark (`小黎的部落`), AGPLv3 + commercial dual licensing, and compliance of third-party source inclusion (license review + source registration).

### 2. File layout

```
governance/
  HEADERS.md          // copyright header specification and template
  THIRD-PARTY.md     // third-party source inclusion registry
  DUAL-LICENSE.md    // dual-licensing explanation (repository root)
scripts/
  header-check.sh     // copyright header gate
  watermark.sh        // build watermark injection
```

### 3. Copyright header specification (mandatory; headers must be enforced wherever users are involved)

- Template (Go source files):

```go
// Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
// Abate is dual-licensed under AGPLv3 and a commercial license.
// For commercial licensing, contact xiaolibuluo@163.com.
// Source: https://github.com/xiaolibuluo/Abate
```

- Coverage: all `.go`/`.sh`/`.yml`/`.md` files (document headers use YAML front-matter or a first-line comment); `third_party/` is excluded (upstream copyright headers preserved, managed via the registry).
- Gate: `header-check.sh` matches the header of every file (missing/malformed → FAIL); permanently resident in CI.

### 4. Digital code watermark (threefold)

1. Copyright-header watermark: the template above (static, distributed with the source).
2. Build-metadata watermark: injected via `-ldflags` (version/commit/build time/watermark marker `小黎的部落`); `abate version` outputs the full metadata.
3. Release-manifest watermark: release-asset SHA256 manifest (`manifest.txt`) includes a watermark line; covered by GPG signature (tamper protection).

### 5. Third-party source inclusion registry (`THIRD-PARTY.md`, mandatory)

Each entry records: library name / version / upstream commit hash / license / review conclusion / trimming list / backport policy. First-version registry list (aligned with the respective volumes):

| Library | Purpose | License | Inclusion location |
|---|---|---|---|
| PIT | mutation algorithm reference | Apache-2.0 | third_party/pit (read-only reference) |
| chaosblade | injection core source | Apache-2.0 | third_party/chaosblade → kernel/inject |
| wazero | WASM runtime | Apache-2.0 | third_party/wazero → wasm/host |
| yaml.v3 | YAML parsing | Apache-2.0 | third_party/yaml → ces/parse |
| openpgp (x/crypto) | GPG signature verification | BSD-3-Clause | third_party/openpgp → wasm/signature |
| sigstore/cosign verification library | cosign verification | Apache-2.0 | third_party/sigstore → wasm/signature |
| ruff parser crate | Python parsing | MIT/Apache-2.0 | third_party/ruff-parser → lang/pyfront |
| tree-sitter-python | Python fault-tolerant parsing | MIT | third_party/tree-sitter-python → lang/pyfront |
| swc parser crate | JS/TS parsing | Apache-2.0 | third_party/swc → lang/jsfront |
| LLVM/clang (libclang) | C/C++ parsing | Apache-2.0 + LLVM Exception | third_party/llvm → lang/cfront |
| syn | Rust parsing | MIT/Apache-2.0 | third_party/syn → lang/rsfront |
| Roslyn | C# parsing | MIT | third_party/roslyn → lang/csfront |
| kotlin-compiler frontend | Kotlin parsing | Apache-2.0 | third_party/kotlin-compiler → lang/ktfront |
| otel-go SDK | OTel export | Apache-2.0 | third_party/otel → report/otel |
| SplitMix64 | PRNG algorithm | public domain | sim/rand (algorithm port) |

- Review process: every new inclusion must go through a registry-entry review (license compatibility: AGPLv3 copyleft review — Apache-2.0/MIT/BSD compatible; GPL-2-only incompatible → refused).

### 6. Dual licensing (`DUAL-LICENSE.md`)

- AGPLv3 (open-source side) + commercial license (closed-source side; contact `xiaolibuluo@163.com`).
- Contributor agreements: CLA optional + DCO mandatory (`Signed-off-by`); contribution grants dual-licensing rights (contributor-agreement text ships with the repository). **Subject to confirmation by the rights holder or legal counsel** — DCO does not itself grant a relicense right.

### 7. Test requirements and acceptance criteria

1. `header-check.sh` scans the whole repository with zero missing headers (permanently resident in CI; the deliberately-deleted-header case is blocked).
2. Threefold watermark verification: `version` metadata contains the watermark; the manifest includes a watermark line covered by the signature; build artifact `-ldflags` injection asserted.
3. `THIRD-PARTY.md` and `third_party/` with zero drift (scan check: every library in the directory has a registry entry).
4. License-compatibility review conclusions cover 100% of the registry list (acceptance target).

---

## System XCU-CI: CI/CD gates

### 1. System positioning and responsibilities

The full quality gate between merge and release: five-section gate chain (already built) + copyright-header gate + baseline-regression gate + TCK gate + determinism gate + drift-check gate. Any rule hit → FAIL with the violated rule printed.

### 2. File layout

```
scripts/
  lint-gate.sh       // five-section gate (already built: gofmt→vet→golangci-lint→test -race→build)
  header-check.sh    // copyright header gate
  drift-check.sh     // drift check (generated artifacts in sync with sources)
  determinism-check.sh // determinism gate (two-run hash comparison)
.github/workflows/
  ci.yml             // main gate (already built)
  benchmarks.yml     // baseline regression (from M2)
  nightly.yml        // nightly full run (fuzzing + parity check + large scale)
```

### 3. Full gate-chain definition (complete checkpoints before merging into main)

1. Five-section gate (`lint-gate.sh`, already built): gofmt → go vet (36 analyzers) → golangci-lint (110+ linters, `issues-exit-code=1`) → `go test -race -count=1` → `go build`.
2. Copyright-header gate (`header-check.sh`): first-line template match on every file, missing → FAIL.
3. Determinism gate (`determinism-check.sh`): two runs of core packages (same seed) → SIM-REC Compare zero difference; hit list: `kernel/*`, `sim/*`, `ces/*` (progressively included from M1).
4. Drift-check gate (`drift-check.sh`): four categories of generated artifacts in zero drift with sources — CES JSON Schema/`fields.md` (CES-SCHEMA generator), OpenAPI (CTL-REST routes), frontend API client (generated from `openapi.json`), frontend `dist` (frontend build).
5. Baseline-regression gate (`benchmarks.yml`, from M2): benchmark suite run compared with baselines (performance budget thresholds); regression beyond threshold → FAIL.
6. TCK gate: CES-TCK self-bootstrap (Abate Driver 100% pass) + kernel API TCK all green (acceptance target).
7. Nightly full run (`nightly.yml`): long fuzzing profiles (30-minute profile per parser) + PIT parity check + scale tests (1-million-frame archive / 100,000-packet network / 10,000 tasks).

### 4. Key processes

1. Merge flow: this project does not accept PRs; the rights holder runs all gates green locally, then directly commits with a GPG signature and pushes to `main` (branch protection enforces only required signatures and no-force-push); gate failures print the violated rule (aligned with the requirement to print the rule explicitly on trigger).
2. Gate tiers: quick profile (pre-commit locally, ≤10 minutes: five sections + copyright header + quick determinism) and full profile (nightly); baseline regression runs on a dedicated runner (performance jitter isolation).
3. Gate self-testing: gate scripts have unit tests (injected violation samples asserting FAIL — the verified godox/unused interception is to be extended to all gate sections).

### 5. Test requirements and acceptance criteria

1. All seven gate categories built and permanently in CI; each gate category has a negative case of "injected violation blocked" (aligned with the verified stub-code interception test).
2. PR quick profile P95 ≤10 minutes (CI duration statistics archived) (design target).
3. Four generated-artifact categories with zero drift (CI record).
4. Gate scripts themselves pass all five gate sections (tools are code too).

---

## System XCU-REL: Release engineering

### 1. System positioning and responsibilities

Version release process: tag signing, asset building, checksum manifests, SBOM, and announcements. Implements the version strategy of annual kernel releases + hourly ecosystem releases.

### 2. File layout

```
scripts/
  release.sh        // release process orchestration (build→verify→sign→manifest)
  sbom.sh           // SBOM generation (SPDX format)
SECURITY.md         // security policy and trust-anchor fingerprint publication
CHANGELOG.md        // changelog (one entry per version)
```

### 3. Version strategy

1. Kernel annual cadence: one major per year (annual number); LTS 3-year maintenance window; kernel API semver (major ≥3-year interval, aligned with the ten-year kernel iron rule).
2. Ecosystem hourly cadence: independent plugin-repository versions (semver) + compatibility matrix (host version interval); plugin CI on hourly release windows.
3. Engine version: `v<annual>.<minor>.<patch>` (e.g. `v1.2.3`); pre-releases (alpha/beta/rc) do not enter LTS.

### 4. Release process (`release.sh`, must be fully automated)

1. Pre-checks: all gates green (full XCU-CI chain) + complete CHANGELOG entry + version consistency (ldflags/CHANGELOG/OpenAPI agree in three places).
2. Build: target matrix (linux/amd64, linux/arm64, darwin/amd64, darwin/arm64, windows/amd64); reproducible builds (two builds byte-identical — `GOFLAGS=-trimpath` + fixed toolchain) (design target).
3. Assets: binaries + archives + SHA256 manifest (`manifest.txt`, with watermark line) + SBOM (SPDX 2.3, including all third_party registrations).
4. Signing: all assets + manifest GPG-signed (`.sig` files); tag signed (annotated + signed).
5. Announcement: GitHub Release (assets + signatures + SBOM) + trust-anchor fingerprint publication (`SECURITY.md` updated if rotated).

### 5. Test requirements and acceptance criteria

1. Release-process rehearsal (M6 acceptance): full flow walked through (verify→build→sign→manifest→Release draft); reproducible builds (two builds byte-identical) (acceptance target).
2. SBOM passes the SPDX validator; signature chain verified (public key verifies all assets) (acceptance target).
3. Three-place version consistency check in CI (drift → FAIL) (acceptance target).
4. Rollback process: one rollback rehearsal for a release defect (record archived) (acceptance target).

---

## System XCU-TST: Testing

### 1. System positioning and responsibilities

Engine-wide test strategy and shared infrastructure: the test pyramid, parity checks, fuzzing, and chaos self-testing (the ultimate form of eating our own dog food — Abate runs chaos experiments against itself).

### 2. Test pyramid (proportions and ownership)

1. Unit tests 70%: table-driven; full coverage of normal/boundary/error paths; defined in each module volume (estimated ≥2000 cases in total) (design target).
2. Integration tests 20%: cross-module chains (end-to-end cases defined in the volumes' "overall merge acceptance").
3. End-to-end tests 10%: full flows across the three entry points CLI/REST/Web.

### 3. Specialized testing (four types)

1. Parity checks: PIT parity check (mutation-verdict agreement 100%, KRN-MUT acceptance item); golangci-lint parity check (Go rule-set difference report, LNG-CHK acceptance item); upstream-engine parity check (chaosblade behavior locking, KRN-INJ TCK).
2. Fuzzing: CES-PARSE (document input), all seven LNG-* frontends (source input), AIR serialization (random trees), WAS-HOST (wasm bytes) — four targets each with a 30-minute nightly profile + a 60 s PR quick profile; the corpus is archived and evolves continuously (design target).
3. Chaos self-testing (ultimate dogfooding): Abate runs mutation experiments (M family) against its own codebase and chaos experiments (C family, inside the simulator) against its own test environment; the survival score goes on the quality dashboard (target: core-package survival score ≥0.85; SURVIVED mutants automatically open improvement issues — linked with ADP-ITS) (design target).
4. Determinism testing: all core packages two-run hash-identical (XCU-CI determinism gate); SIM-REC playback frame-by-frame verification.

### 4. Shared infrastructure

- `testutil` package: determinism assertion helpers (two-run hash / serialization comparison), fixture loaders (`testdata`), baseline corpus management.
- Corpus: unified `testdata/` layout (defined in each module volume); corpus provenance and license registration (license review of open-source snapshots — aligned with XCU-GOV).

### 5. Acceptance criteria

1. Pyramid proportions met (case-count statistics report: unit/integration/end-to-end); total case scale ≥2000 (design target).
2. All four specialized testing types built and in CI; fuzz corpus archived and continuously growing (quarterly statistics).
3. Chaos self-testing: core-package survival score ≥0.85 (dashboard archived); 100% of SURVIVED mutants get an issue (acceptance target).
4. Three parity reports (PIT/golangci-lint/chaosblade) archived and updated at milestones.

---

## System XCU-PRF: Performance baselines

### 1. System positioning and responsibilities

Setting performance budgets, maintaining the benchmark suite, and the regression gate. The quantifiable landing point of the words "fast" and "excellent".

### 2. File layout

```
bench/
  suite_test.go     // benchmark suite entry (go test -bench)
  baseline/         // baseline data (archived per platform)
  budget.md         // performance budget table (this table)
```

### 3. Performance budget table (full, aggregated acceptance items from the volumes)

| ID | Metric | Budget | Owner |
|---|---|---|---|
| P1 | Scheduler event throughput | ≥100,000 events/s | KRN-SCH |
| P2 | PRNG extraction throughput | ≥10^8 draws/s | SIM-RNG |
| P3 | Virtual network propagation (100,000 packets) | P99 ≤1 s | SIM-NET |
| P4 | Virtual scheduling (10,000 tasks) | throughput ≥100,000 events/s | SIM-SCHED |
| P5 | Archive Load+Verify (1 million frames) | P99 ≤10 s | SIM-REC |
| P6 | AIR serialization (1 million nodes) | P99 ≤500 ms | LNG-AIR |
| P7 | AIR deserialization (1 million nodes) | P99 ≤700 ms | LNG-AIR |
| P8 | Frontend parsing throughput (six languages) | ≥1 MB/s (C/C++ ≥0.5 MB/s) | LNG-* |
| P9 | Rule matching (100 rules/file) | P99 ≤50 ms | LNG-RULE |
| P10 | HTML report (100,000 violations) | P99 ≤5 s | RPT-HTML |
| P11 | Plugin call overhead | P99 ≤50 μs | WAS-HOST |
| P12 | Plugin signature verification | P99 ≤10 ms | WAS-SIG |
| P13 | CES parsing (100 KB document) | P99 ≤10 ms | CES-PARSE |
| P14 | CES validation (100 KB document) | P99 ≤5 ms | CES-VALID |
| P15 | CI quick-profile total duration | P95 ≤10 minutes | XCU-CI |

### 4. Key processes

1. Benchmark runs: `benchmarks.yml` (dedicated runner, fixed CPU model); full benchmarks on every merge to main.
2. Regression verdict: comparison against baselines (statistically robust: median of 5 runs, benchstat style); regression >10% → FAIL (budget-table metrics) (acceptance target).
3. Baseline updates: when a performance improvement is merged, the baseline is updated with the version (recorded in CHANGELOG).

### 5. Acceptance criteria

1. All 15 budget metrics have benchmark cases and baseline data (progressively included from M2, complete at M6) (design target).
2. Regression gate running (regression >10% FAIL case verified) (acceptance target).
3. Benchmark reports archived per version (trend traceable).

---

## System XCU-DOC: Documentation system

### 1. System positioning and responsibilities

A complete system of four documentation categories: user documentation, developer documentation, specification texts, and architecture decision records (ADRs).

### 2. File layout

```
docs/
  ARCHITECTURE.md    // architecture overview (already built)
  plan/              // planning document system (00-08, this system)
  spec/              // specification texts
    ces-v1.md        // CES specification (in sync with CES-SCHEMA)
  adr/               // architecture decision records
    0001-record.md   // template and index
    0002-fail-fast-parse.md
    0003-token-level-roundtrip.md
    ... (all ADR items referenced in the volumes are archived)
  user/              // user documentation
    quickstart.md
    cli-reference.md // full command-tree reference (generated from --help)
    ces-guide.md     // CES authoring guide
    rules/           // rule-set directory (seven languages, with source links)
  dev/               // developer documentation
    contributing.md
    module-dev-guide.md // module development guide (including acceptance DoD)
    plugin-dev-guide.md // plugin development guide (WIT interface + scaffolding)
```

### 3. Documentation synchronization iron rules

1. `cli-reference.md` is generated from `--help` output (CI drift check); `ces-v1.md` is synchronized with CES-SCHEMA generated artifacts (drift check).
2. The rule-set directory is synchronized with `builtin/` rules (per rule: id/description/source/positive+negative examples — generated from rule YAML).
3. The ADR index agrees with the bodies (continuous numbering; status proposed/accepted/superseded).

### 4. Acceptance criteria

1. All four documentation categories fully in place (M6 acceptance item); two generated synchronizations with zero drift (CI) (acceptance target).
2. All `quickstart` commands copyable and executable (documentation tests); the rule directory covers all built-in rules (≥460).
3. All ADRs archived (ADR-0001~000N referenced in the volumes numbered continuously with no gaps).

---

## Appendix: Overall crosscutting-systems merge acceptance (after all seven systems are in place)

1. Full-chain rehearsal: one commit → all seven gate categories green → release process walked through (signed assets + SBOM + announcement) → chaos self-test report produced → documentation synchronization with zero drift — a complete closed-loop verification (acceptance target).
2. Security audit: all T1–T10 attack cases blocked; external security audit completed (GA prerequisite) (acceptance target).
3. All 15 performance-budget items meeting their targets and the regression gate running (design target).
4. Governance compliance: zero missing copyright headers, threefold watermark verification, `THIRD-PARTY` zero drift, dual-license texts in place (acceptance target).
5. An aggregate coverage and case-statistics report for all seven systems is archived; the five-section gate chain + full gate chain are all green (acceptance target).
