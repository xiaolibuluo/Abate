<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->

> Language / 语言：[English](00-MASTER-PLAN.en.md) | [中文](00-MASTER-PLAN.md)

# Abate Architecture Master Plan

**Design-target statement (mandatory for C-class translations)**: this file
is the English translation of an internal planning document. All `100%` /
"pass" / "met" style statements in the full text are **Design targets /
Acceptance targets**, **not measured results**; they must not be read as
established facts or capability claims. The translation corresponds
section by section, does not escalate the tone, and does not delete
target qualifiers.

- Document ID: PLAN-00　Status: in force　Owner: xiaolibuluo
  (xiaolibuluo@163.com)
- Review cadence: mandatory review at the end of every milestone (M1–M6);
  major architecture changes require an ADR
- This document is the master plan for all module plans; the extreme detail
  of each module is in the docs/plan/01–08 volumes

---

## 1. Purpose and scope

This document answers four questions:

1. Which modules Abate must build in total (the full module map; no
   omissions).
2. Who depends on whom (the dependency graph; circular dependencies
   forbidden).
3. In what order to build (M0–M6 milestones; each phase has explicit
   acceptance criteria and exit conditions).
4. What "done" means (the acceptance general rules: DoD, coverage,
   performance budget, determinism, security, documentation).

Scope: all of Abate's L0–L7 layers and crosscutting concerns. Excluded:
website operations, community-operations processes (separate documents).

## 2. Overall goal and success criteria

Overall goal: build a deterministic engine integrating chaos testing and
mutation testing, with a kernel that goes ten years without major updates and
only simple yearly maintenance, and an ecosystem that extends hourly.

Success criteria (quantifiable):

1. Same seed, same configuration, byte-for-byte identical results on any
   number of runs (deterministic-replay verification pass rate 100% — design
   target).
2. The kernel (L0-L3) has zero dependence on external networks: the full test
   suite passes in an air-gapped environment.
3. All seven-language official-standards rule sets land, and the Abate
   repository itself passes the in-house checker at 100% (design target;
   eating its own dog food).
4. The whole plugin flow from signature to hot-loaded deployment takes under
   1 hour (hourly ecosystem cadence).
5. Kernel API breaking changes per year: 0 (TCK all green).

## 3. Full module map

Module numbering rule: layer prefix-module name. Once assigned, a number is
never reused (ten-year traceability).

| ID | Module | Layer | Responsibility | Detail volume | Milestone |
|---|---|---|---|---|---|
| KRN-ALG | Perturbation algebra | L0 | Operator algebra structure: composition/concurrent product/dual/identity, four operations, and algebraic-law verification | 01 | M1 |
| KRN-SCH | Deterministic scheduler | L0 | Virtual-clock-driven deterministic experiment scheduling | 01 | M1 |
| KRN-OBS | Observation pipeline | L0 | Unified observation event stream: probe data/metrics/traces normalized | 01 | M1 |
| KRN-VER | Verifier | L0 | Steady-state assertions, survival score, equivalence analysis | 01 | M1 |
| KRN-MUT | Mutation core | L0 | PIT algorithm source-level port: mutation operator families + coverage-guided scheduling + minimization | 01 | M2 |
| KRN-INJ | Injection core | L0 | chaosblade source secondary development: C-family injection primitives + dual rollback | 01 | M2 |
| WAS-HOST | WASM host | L1 | WASI Preview2 runtime host and WIT bindings | 02 | M4 |
| WAS-CAP | Capability sandbox | L1 | fs/net/clock/env capability granting with default deny-all | 02 | M4 |
| WAS-HOT | Hot-loading | L1 | Atomic replacement, old-instance draining, version negotiation | 02 | M4 |
| WAS-SIG | Signature verification | L1 | GPG/cosign plugin signature-chain verification | 02 | M4 |
| CES-SCHEMA | Specification schema | L2 | CES YAML/JSON Schema definition and versioning | 03 | M1 |
| CES-PARSE | Specification parsing | L2 | CES document parsing into kernel experiment entities | 03 | M1 |
| CES-VALID | Specification validation | L2 | Semantic validation: steady-state/method/probe/rollback/schedule legality | 03 | M1 |
| CES-TCK | Conformance suite | L2 | Third-party implementation conformance test suite | 03 | M6 |
| SIM-CLK | Virtual clock | L3 | Deterministic clock with fast-forward/slow-motion/freeze/offset | 04 | M1 |
| SIM-RNG | Deterministic PRNG | L3 | Seed-driven pseudo-random stream (replayable) | 04 | M1 |
| SIM-NET | Virtual network | L3 | Latency/partition/packet-loss/bandwidth fully simulated in software | 04 | M2 |
| SIM-FS | Virtual filesystem | L3 | In-memory FS and fault-injection points | 04 | M2 |
| SIM-SCHED | Virtual scheduling | L3 | Virtual processes/coroutines and termination faults | 04 | M2 |
| SIM-REC | Record/replay | L3 | Full observation recording and frame-by-frame replay | 04 | M2 |
| LNG-AIR | Unified IR | L4 | Abate intermediate representation: node model + semantic annotations + versioned serialization | 05 | M3 |
| LNG-GO | Go frontend | L4 | go/ast to AIR frontend | 05 | M3 |
| LNG-PY | Python frontend | L4 | ruff parsing frontend / tree-sitter source inclusion | 05 | M5 |
| LNG-JS | JS/TS frontend | L4 | swc source inclusion | 05 | M5 |
| LNG-C | C/C++ frontend | L4 | libclang source inclusion | 05 | M5 |
| LNG-RS | Rust frontend | L4 | syn+quote source inclusion | 05 | M5 |
| LNG-CS | C# frontend | L4 | Roslyn source inclusion | 05 | M5 |
| LNG-KT | Kotlin frontend | L4 | kotlin-compiler source inclusion | 05 | M5 |
| LNG-RULE | Rule engine | L4 | YAML rule DSL + AIR pattern matching | 05 | M3 |
| LNG-CHK | Checker | L4 | In-house language code checker (SARIF output) | 05 | M3 |
| LNG-MUT | Operator derivation | L4 | Automatic derivation of rules into M-family mutation operators | 05 | M3 |
| ADP-CI | CI adapter | L5 | GitHub Actions/GitLab CI adaptation | 06 | M4 |
| ADP-ITS | Tracker adapter | L5 | Issue-tracking system integration | 06 | M5 |
| ADP-MSG | Messaging adapter | L5 | Messaging-channel notifications | 06 | M5 |
| RPT-SARIF | SARIF report | L6 | SARIF 2.1.02 standard report | 06 | M4 |
| RPT-HTML | HTML report | L6 | Human-readable interactive report | 06 | M5 |
| RPT-JSON | JSON report | L6 | Machine-readable specification report | 06 | M4 |
| RPT-JUNIT | JUnit report | L6 | CI-consumable format | 06 | M5 |
| RPT-OTEL | OTel report | L6 | OpenTelemetry export | 06 | M5 |
| CTL-CLI | Command line | L7 | abate command-line entry | 07 | M1 |
| CTL-REST | REST API | L7 | Local/remote control API | 07 | M5 |
| CTL-WEB | Web console | L7 | Browser console | 07 | M6 |
| XCU-SEC | Security system | Crosscutting | Threat model, signature chain, vulnerability process | 08 | ongoing |
| XCU-GOV | Governance & compliance | Crosscutting | Copyright header, watermark, dual licensing, CLA/DCO | 08 | M1 |
| XCU-CI | CI/CD gates | Crosscutting | Five-stage gate chain + benchmark regression + TCK gates | 08 | M0 |
| XCU-REL | Release engineering | Crosscutting | GPG-signed releases, SBOM, checksums | 08 | M6 |
| XCU-TST | Testing system | Crosscutting | Test pyramid, parity checks, fuzzing, chaos self-testing | 08 | ongoing |
| XCU-PRF | Performance benchmark | Crosscutting | Benchmark suite and performance-budget gates | 08 | M2 |
| XCU-DOC | Documentation system | Crosscutting | User/dev/spec/ADR documentation | 08 | ongoing |

## 4. Dependency graph

```
CTL-CLI / CTL-REST / CTL-WEB（L7）
        │ depends on
        ▼
ADP-* / RPT-*（L5/L6, loaded via WAS-HOST）
        │ depends on
        ▼
LNG-*（L4: AIR → rules/checks/operator derivation）
        │ depends on
        ▼
CES-*（L2）──► KRN-ALG/SCH/OBS/VER（L0）◄── SIM-*（L3）
                    │ contains
                    ▼
              KRN-MUT / KRN-INJ（L0 kernel-owned）
```

Dependency iron rules:

1. Dependencies run top-down and outside-in only; L0 imports no L1-L7
   packages.
2. No circular dependencies: enforced by the CI depguard (enabled in
   .golangci.yml).
3. L0 depends only on the Go standard library and in-house source-level code
   (the ten-year kernel iron rule).
4. WAS-HOST is the only channel through which L4-L7 plugins enter the kernel
   (ports-and-adapters).

## 5. Phases and milestone plan

### M0 (completed): foundation

- Deliverables: repository skeleton, five-stage gate chain
  (gofmt→vet→golangci-lint→test -race→build), CI workflow, architecture
  master plan.
- Acceptance: all gates green; stub-code injection measured and intercepted
  (a godox/unused hit means FAIL).

### M1: minimal kernel loop (a deterministic experiment runs end to end)

- Goal: run one deterministic HPOV experiment in a single process (no
  language frontends, no WASM).
- Deliverables: KRN-ALG, KRN-SCH, KRN-OBS, KRN-VER, CES-SCHEMA/PARSE/VALID,
  SIM-CLK, SIM-RNG, and the minimal CTL-CLI set.
- Acceptance criteria:
  1. One CES YAML declaring steady state + perturbation + probes; the CLI
     executes it and outputs a deterministic report.
  2. 100 runs with the same seed produce byte-for-byte identical reports
     (hash-verified).
  3. Perturbation-algebra unit tests cover all algebraic laws of the four
     operations (associativity/identity/dual compensation).
  4. Kernel packages (L0/L2/L3) test coverage: statement ≥95%, branch ≥85%.
- Exit condition: all acceptance items pass and all gates green.

### M2: full simulator + dual-core inclusion

- Goal: the deterministic simulator in full + PIT/chaosblade source-level
  inclusion complete.
- Deliverables: SIM-NET/FS/SCHED/REC, KRN-MUT, KRN-INJ, and the XCU-PRF
  benchmark suite.
- Acceptance criteria:
  1. The virtual network's four fault classes (latency/partition/packet
     loss/limited bandwidth) can be injected inside the simulator and replay
     deterministically and consistently.
  2. PIT parity: over the same mutation-point set, the consistency rate of
     Abate's ported operators vs the PIT reference implementation is 100%
     (design target; parity report filed in the repository).
  3. After chaosblade inclusion: the original CLI shell removed; the library
     form passes the kernel TCK; one upstream security-patch backport
     rehearsal.
  4. First performance-budget measurement: scheduler per-experiment overhead
     and simulator event throughput meet the budget (figures in the 08
     volume's performance budget table).

### M3: Go language frontend loop (eating its own dog food)

- Goal: LNG-AIR+LNG-GO+LNG-RULE+LNG-CHK+LNG-MUT in full; the Abate repository
  self-checked.
- Acceptance criteria:
  1. AIR serialization schema versioned; cross-version compatibility tests
     pass.
  2. Go official-standards rule set (Effective Go+CodeReviewComments+vet
     semantic subset): ≥100 rules land.
  3. The Abate repository passes the in-house checker with 0 violations
     (parity-checked against golangci-lint results, with each difference
     explained).
  4. Rule→operator derivation: at least 20 rules automatically derive
     M-family operators usable in mutation experiments.

### M4: WASM runtime + minimal ecosystem loop

- Goal: the plugin system runs end to end (at least 1 official plugin each
  for operator/reporter/adapter).
- Acceptance criteria:
  1. Unsigned plugin loading is rejected; capability-escalation calls are
     intercepted by the sandbox (security test cases).
  2. Hot-loading: replacing a plugin version does not interrupt a running
     experiment; the old instance is drained before unload.
  3. The whole plugin flow from commit to publish takes ≤1 hour (design
     target; CI-measured record).
  4. RPT-SARIF output passes the official SARIF validator.

### M5: all seven languages + control plane

- Goal: LNG-PY/JS/C/RS/CS/KT all land; CTL-REST goes live.
- Acceptance criteria:
  1. Each language frontend's parsing throughput and AIR fidelity meet the
     bar (100% of each language's benchmark case set passing — design
     target).
  2. Each language's official-standards rule set ≥60 rules; the checker
     outputs SARIF.
  3. REST API OpenAPI documentation generated, with implementation-consistency
     tests passing.

### M6: GA

- Goal: quality, security, documentation, and release all meet the bar.
- Acceptance criteria:
  1. CES-TCK released; a third-party implementation completes the conformance
     process (at least 1 example implementation).
  2. Security audit: signature chain, sandbox-escape tests, dependency
     vulnerability scanning all green.
  3. XCU-REL: GPG-signed release process rehearsal (tag + assets + checksums +
     SBOM).
  4. Full documentation: user guide/dev guide/spec text/ADR all in place.

## 6. Acceptance general rules (the DoD common to all modules)

A module counts as done only if it simultaneously satisfies:

1. Code: merged into main with all five gate stages green; no TODO/FIXME
   (godox interception); copyright header + watermark present.
2. Tests: unit tests (table-driven, covering normal/boundary/error paths) +
   integration tests + determinism replay tests (same-seed double run with
   identical hashes). Coverage: kernel ≥95%/85% (statement/branch), ecosystem
   ≥85%/75%.
3. Interfaces: public APIs carry Go doc comments (godot enforces trailing
   periods); breaking changes go through an ADR + a 2-version deprecation
   window.
4. Documentation: a module README (responsibilities/interfaces/acceptance
   checklist) + a CHANGELOG entry.
5. Performance: modules touching hot paths carry benchmarks within the
   performance budget (volume 08).
6. Security: new dependencies must be included at source level (the five-step
   process) with the source commit hash and license-review verdict recorded.

Test pyramid: unit 70% / integration 20% / end-to-end 10%; plus parity tests
(PIT/upstream engines), fuzzing (parsers and CES inputs), and chaos
self-testing (Abate runs chaos experiments on itself).

## 7. Risk register

| ID | Risk | Level | Mitigation |
|---|---|---|---|
| R1 | PIT is implemented in Java — a language gap with the Go kernel | High | PIT is positioned as an algorithm reference + parity baseline: vendor the source into third_party/pit (read-only reference); port the core algorithms (operator families/coverage scheduling/minimization) into KRN-MUT at Go source level, verifying operator-by-operator with a 100% consistency-rate design target; Abate mutations act on the AIR source layer rather than JVM bytecode, so no runtime JVM dependency is architecturally needed |
| R2 | chaosblade upstream evolution fork | Medium | Freeze the baseline at inclusion; review and backport upstream security patches; lock behavior with the kernel TCK |
| R3 | WASM ecosystem interface instability (WASI Preview2 evolution) | Medium | Version the WIT interfaces + compatibility matrix; the host layer isolates WASI details |
| R4 | Seven-language frontend effort exceeds expectations | High | M3 first validates the architecture with the Go loop; the other languages proceed in parallel batches at M5; frontends only reuse parsing, all rules are in-house sharing the AIR layer |
| R5 | Determinism broken by concurrency | High | The whole kernel is single-threaded and virtual-clock driven; real concurrency exists only at simulator boundaries with totally ordered event sequence numbers; `-race` gate + replay hash verification as double insurance |
| R6 | AGPLv3 ecosystem acceptance | Medium | Dual licensing (commercial license) + the open CES specification lower the adoption barrier |

## 8. Document index

- 01-kernel-hpov.md: L0 kernel six modules in extreme detail
  (KRN-ALG/SCH/OBS/VER/MUT/INJ)
- 02-wasm-runtime.md: L1 WASM runtime four modules
- 03-ces-spec.md: L2 CES specification layer four modules
- 04-simulator.md: L3 deterministic simulator six modules
- 05-language-frontend.md: L4 language frontend ten modules
  (AIR/7 languages/rules/checks/derivation)
- 06-adapter-reporter.md: L5/L6 adapters and reporters, eight modules
- 07-control-plane.md: L7 control plane three modules
- 08-crosscutting.md: seven crosscutting systems
  (security/governance/CI/release/testing/performance/documentation)
