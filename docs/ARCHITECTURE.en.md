<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->

> Language / 语言：[English](ARCHITECTURE.en.md) | [中文](ARCHITECTURE.md)

# Abate Super Architecture — An Integrated Chaos-Testing and Mutation-Testing Engine
## Complete closed-loop refinement system · extreme detail

English translation of `ARCHITECTURE.md`, the section-by-section companion
under the plan-A side-by-side naming scheme; section order and tables
correspond one-to-one; the title 「超级架构」 is translated literally as
"Super Architecture" without upgrading it into a world-class or established
capability claim; competitor mentions are translated as the original
qualitative statements, with no absolute scores added; no new conclusions are
added; this document does not constitute legal advice.

- Project: Abate (https://github.com/xiaolibuluo/Abate)
- Website: www.qingdujian.com　Email: xiaolibuluo@163.com
- License: AGPLv3 + commercial dual licensing (LICENSE and DUAL-LICENSE.md)
- Code watermark: 小黎的部落 (triple watermark: copyright header + build
  metadata + release-signing manifest)

---

## 0. Overview

### 0.1 Four value pillars

| Pillar | Meaning | Differentiation |
|---|---|---|
| Unified dual kernel | Mutation and chaos unified on the HPOV perturbation algebra: one engine, two modes | Requires mastering both disciplines and unifying them deeply — a cognitive differentiator |
| WASM plugin ecosystem | Operators, injectors, language frontends, adapters, reporters all as WASM: any language, sandboxed, hot-loaded | Competitors ship core releases quarterly; community WASM plugins publish hourly |
| Open specification CES | The chaos-experiment specification is open and neutral; any implementation can conform | Competitors bind to proprietary platforms; Abate binds to an open specification |
| Deterministic simulator | Pure-software digital twin, never touching real platform APIs | Competitors depend on real clusters for attestation, costly and subject to obsolescence; Abate's models do not go stale |

### 0.2 The nine-character quality definition

| Character | Definition | Where it lands |
|---|---|---|
| 稳 (Stable) | Deterministic reproducibility | seed + virtual clock + record/replay |
| 快 (Fast) | Near-native execution | WASM AOT + parallel scheduling |
| 精 (Precise) | Precise perturbation | AST node-level / syscall-level |
| 准 (Accurate) | Zero missed observations | unified observation pipeline + survival score |
| 固 (Solid) | Fully strict gates | 110+ linters, five-stage gate chain |
| 全 (Complete) | All languages, all fault domains | 7 languages + M/C full operator families |
| 优 (Optimized) | Continuous benchmark regression | benchmark suite in the CI gate |
| 劲 (Agile) | Hourly ecosystem extension | WASM plugin hot-loading |
| 强 (Strong) | Defense in depth | capability sandbox + signature chain |

### 0.3 Three overarching design principles

1. Source-level inclusion iron rule: any third-party open source is included
   at source level (vendor → secondary changes → optimization → deep
   integration → encapsulation); API access is forbidden.
2. Ten-year kernel: the kernel must be strong enough — built right once,
   ten years without major updates, only simple yearly maintenance.
3. All structural evolution happens in the ecosystem layer: the kernel turns
   yearly; the ecosystem turns hourly.

## 1. Dual-kernel unification: the HPOV perturbation algebra

### 1.1 The HPOV loop

H (Hypothesis, steady-state assumption) → P (Perturbation) → O (Observation)
→ V (Verification) → back to H. Every experiment is an instance of the HPOV
loop.

### 1.2 Perturbation algebra definition

Let P be the set of perturbation operators, with:

| Operation | Notation | Semantics |
|---|---|---|
| Composition | p1 ∘ p2 | Apply p2 first, then p1 (temporal) |
| Concurrent product | p1 ⊗ p2 | Apply simultaneously (concurrency domain) |
| Dual (rollback) | p⁻¹ | Undo the perturbation, restore the steady state |
| Identity | ε | Empty perturbation (no-op) |

Algebraic properties: composition is associative; ε is the identity; rollback
satisfies p ∘ p⁻¹ ≈ ε (eventually-consistent compensation).

### 1.3 Operator families

- M family (mutation mode): acts on the code space. Operator replacement,
  boundary-value perturbation, constant perturbation, control-flow
  perturbation, return-value perturbation, negated conditions, call removal.
  Derived automatically from the language rule sets (see 3.4).
- C family (chaos mode): acts on the system space. Network
  latency/partition/packet loss, CPU/memory/IO/disk stress, clock skew,
  dependency failures, process/coroutine termination.
- M ⊗ C (hybrid mode): while mutating the code twin, inject faults into the
  system twin, verifying the test suite's survival verdict under dual
  perturbation (equivalent-mutant identification and race exposure).

### 1.4 The unified engine triad

One engine = the same scheduler (driven by a deterministic virtual clock) +
the same observation pipeline (probes/metrics/traces) + the same verifier
(steady-state assertions + survival score + equivalence analysis). Mutation
and chaos differ only in operator family; the execution path is fully
isomorphic.

## 2. Engine source inclusion: PIT + chaosblade, the five-step inclusion process

### 2.1 The five-step inclusion process

Source vendor → secondary changes → optimization → deep integration →
encapsulation as kernel primitives. API access is forbidden in any form: no
out-of-process calls, no gRPC, no network dependencies, no runtime dynamic
loading of external engines.

### 2.2 PITest (PIT) inclusion — the mutation core

- Vendor: the JVM mutation-testing engine's source into the kernel tree.
- Secondary changes: determinization (seed-driven, virtual-clock aligned).
- Optimization: trim non-core paths, remove external dependencies.
- Deep integration: mutation operator families connect directly to the HPOV
  algebra (M family); coverage-guided scheduling merges into the
  deterministic scheduler.
- Encapsulation: the kernel's M-family mutation primitives.

### 2.3 chaosblade inclusion — the injection core

- Vendor: the Go fault-injection engine's source into the kernel tree (same
  language as the kernel, natural affinity).
- Secondary changes: deterministic scheduling adaptation, unified observation
  events.
- Optimization: trim the CLI shell, converge to a library shape.
- Deep integration: injection actions translate directly into C-family
  operators; rollback merges into the algebra's dual operation.
- Encapsulation: the kernel's C-family injection primitives.

### 2.4 Inclusion iron rules

- Concept unification: external concepts are all translated into HPOV algebra
  and CES vocabulary; no PIT/chaosblade-specific types appear in the kernel.
- Upstream backport: upstream security patches and critical fixes are
  reviewed and selectively backported; the kernel version stays
  self-controlled.

## 3. Language frontend matrix and the in-house rule system

### 3.1 Seven-language low-level library matrix (all source-level inclusion)

| Language | Low-level library | License | Purpose | Inclusion strategy |
|---|---|---|---|---|
| Go | go/ast (standard library) | BSD | Syntax-tree parsing | Standard library native, zero external dependencies, safest; no third-party library needed at all |
| Python | ruff parsing frontend (rustpython-parser family), tree-sitter-python | MIT / Apache-2.0 | AST parsing | Take only the parser to generate syntax trees; do not include ruff's lint-rule sources; rules are in-house |
| JS/TS | swc | Apache-2.0 | AST parsing, source transforms | High-performance Rust library; reuse only the parsing part; rule logic is in-house |
| C/C++ | libclang (LLVM) | Apache-2.0 + LLVM Exception | AST and compile-time semantics | Use libclang only for syntax trees and type info; do not take clang-tidy's full source |
| Rust | syn + quote | MIT / Apache-2.0 dual | Rust source AST parsing | syn is only a parsing library without clippy check logic; rules are in-house per official standards |
| C# | Roslyn | MIT | C# syntax trees, semantic analysis | Borrow only Roslyn's compiler capabilities; no third-party analyzer rules |
| Kotlin | kotlin-compiler | Apache-2.0 | Kotlin syntax trees | JetBrains official; take the AST; rules are in-house |

### 3.2 Unified IR: AIR (Abate Intermediate Representation)

- All seven language frontends produce AIR: a node model (declarations/
  statements/expressions/types/positions/comments) + semantic annotations
  (type info, scopes, control-flow graphs).
- Both the rule engine and mutation operators act only on AIR — the
  language-neutral core is implemented once and all seven languages benefit.
- AIR serialization is stable (versioned schema), supporting caching and
  incremental parsing.

### 3.3 In-house rule system: official-standards rule sets

| Language | Official standard sources |
|---|---|
| Go | Effective Go + CodeReviewComments + full go vet |
| Python | PEP 8 + PEP 484 + official style guides |
| JS/TS | ECMAScript spec + TS Handbook |
| C/C++ | CERT C/C++ + official coding standards |
| Rust | Rust API Guidelines |
| C# | .NET design guidelines |
| Kotlin | Kotlin coding conventions |

### 3.4 Dual-use rules: checker × mutation operator

Every rule serves both engines:

- Checker use: AIR pattern matching → violation report (SARIF), forming the
  language code checker (in-house lint engine).
- Mutation-operator use: a rule is a mutation point — naming/boundary/type/
  concurrency rules automatically derive M-family operators (e.g. a
  boundary-value rule → a boundary-perturbation operator).

### 3.5 In-house language code checker

- Rule DSL: declarative YAML (pattern + predicate + severity + fix hint) +
  code-predicate extensions.
- Check pipeline: AIR → rule matching → violation aggregation → SARIF/HTML
  reports.
- Isomorphic with Abate's own gates: the Abate repository self-checks with its
  own Go rule set (eating its own dog food).

## 4. Layered architecture L0-L7

```
L7 control plane: CLI (abate) / REST API / Web console
L6 reporters: WASM reporter plugins (JSON/HTML/SARIF/JUnit/OTel)
L5 adapters: WASM adapter plugins (CI platforms, issue trackers, messaging channels)
L4 language frontends: WASM frontend plugins + built-in 7-language official-standards rule sets (source-level inclusion)
L3 deterministic simulator: virtual clock/network/FS/scheduler + record/replay
L2 CES specification layer: experiment description/steady state/probes/rollback (open and neutral, with TCK)
L1 WASM runtime: WASI Preview2, capability sandbox, hot-loading, signature verification
L0 HPOV kernel: perturbation algebra, deterministic scheduler, observation pipeline, verifier + source-included mutation core (PIT secondary development) and injection core (chaosblade secondary development)
```

Dependencies point downward only; L0–L3 are the kernel (yearly cadence),
L4–L7 are the ecosystem (hourly cadence).

## 5. WASM plugin ecosystem

### 5.1 Five plugin types (unified WIT interfaces)

| Type | Responsibility | Examples |
|---|---|---|
| operator | Declare perturbation operators (M or C family) | network-delay, arith-replace |
| injector | Land operators on targets (twin or real environment) | tc-injector, syscall-injector |
| frontend | Language frontend: parse source into AIR | rust-frontend, python-frontend |
| adapter | External-system integration | github-actions, gitlab-ci |
| reporter | Result output | sarif-reporter, html-reporter |

### 5.2 Security model

- Capability-based: fs/net/clock/env granted explicitly, default deny-all.
- Mandatory signing: plugins must be signed (GPG/cosign); unsigned plugins are
  refused.
- Hot-loading: atomic replacement (old instance drained before unload);
  upgrades are seamless.
- Version negotiation: semver + compatibility matrix, verified automatically
  in CI.

## 6. The CES open specification

Chaos Experiment Specification: YAML/JSON describing one HPOV experiment:

- apiVersion / kind / metadata
- steadyState: steady-state hypotheses (probe + tolerance + duration)
- method: perturbation sequence (operator composition/concurrency expressions)
- probes: observation probes (HTTP/shell/metric/prom)
- rollback: rollback strategy (auto/manual/timeout)
- schedule: once/cron/continuous

Open and neutral: the specification repository evolves independently
(semantic versioning); a TCK is provided, and any third-party implementation
that passes the TCK can declare CES conformance.

## 7. The deterministic simulator

Pure-software digital twin, never touching real platform APIs:

- Virtual clock: fast-forward/slow-motion/freeze/offset (the basis of chaos
  clock perturbation).
- Deterministic PRNG: seed-driven; same seed, same result (stability).
- Virtual network stack: latency/partition/packet loss/bandwidth fully
  simulated in software.
- Virtual filesystem and virtual scheduling: process/coroutine-level fault
  injection.
- Record/replay: the observation pipeline records everything and can be
  replayed frame by frame for verification.

## 8. The complete closed-loop system (extreme detail)

### 8.1 Data-flow loop: the nine-stage pipeline

Source/system → ① frontend parsing (AIR) → ② rule checking (lint gate) →
③ operator generation (M/C derivation) → ④ experiment orchestration (CES) →
⑤ deterministic scheduling (virtual clock) → ⑥ execution (simulator +
mutation core + injection core) → ⑦ observation (probes/metrics/traces) →
⑧ verification (steady-state assertions + survival score + equivalence
analysis) → ⑨ reporting (SARIF/HTML/OTel) → feedback to ③ (operator and rule
library evolution).

### 8.2 Experiment-execution loop: the HPOV state machine

IDLE → HYPOTHESIZED (steady-state hypothesis registered) → BASELINE (baseline
observation) → PERTURBED (perturbation applied) → OBSERVED (perturbation
observed) → VERIFIED (verdict: steady state held/broken + mutant
survived/killed) → ROLLED_BACK (dual rollback) → REPORTED (report archived) →
terminal state.

Every state transition is recorded deterministically (seed + clock + event
sequence); any experiment can be replayed frame by frame.

### 8.3 Ecosystem loop: the plugin lifecycle

Development (WIT interfaces + scaffold templates) → local testing (plugin
TCK) → signing (GPG/cosign) → publishing (plugin repository) → compatibility
matrix verification (automatic in CI) → distribution (hourly) → hot-loading
(atomic replacement) → telemetry (usage statistics) → feedback
(issues/ratings) → version iteration.

### 8.4 Quality loop: commit to release

Commit → copyright-header check (watermark mandatory) → five-stage gate chain
(gofmt → full go vet → full golangci-lint → test -race → build) → CI (GitHub
Actions) → benchmark regression (no performance regressions) → kernel TCK
(API compatibility) → GPG-signed release (tag + assets + checksums + SBOM).

### 8.5 Feedback loop: continuous evolution

Reports (survival score / equivalent mutants / escaped faults) → weak-spot
location (test blind spots / rule gaps / operator gaps) → rule-library
extension (new lint rules) → operator-library extension (new M/C operators) →
next experiment round (a self-reinforcing loop).

## 9. The ten-year kernel and clean architecture

Kernel design goal: not minimal, but strong enough — built right once, ten
years without major updates, only simple yearly maintenance.

Four kernel iron rules:

1. A strong kernel: L0 has all core capabilities built in — the four
   primitives of perturbation algebra, deterministic scheduler, observation
   pipeline, and verifier, plus the deeply source-included mutation core
   (PIT secondary development) and injection core (chaosblade secondary
   development). The kernel is self-sufficient, depending on no external
   runtime or external engine.
2. Frozen API: kernel APIs are semantically versioned with major-version gaps
   of at least 3 years (at most 2–3 majors in ten years); deprecation goes
   through a 2-version window; the kernel API has its own TCK.
3. Source self-sufficiency: third-party open source is always included at
   source level (five-step process); once included it is kernel-owned code —
   no external-dependency drift, no supply-chain hijack surface; upstream
   security patches are reviewed and selectively backported.
4. Yearly maintenance cadence: the kernel gets only simple maintenance each
   year (security patches, compiler/platform adaptation, performance
   fine-tuning) — no structural overhauls; all structural evolution happens
   in the ecosystem layer.

Clean architecture (the ceiling guarantee for ten years of large-scale
extension and iteration):

- Layers: entities (HPOV algebra/CES entities/AIR) ← use-cases (experiment
  orchestration) ← adapters (WASM plugins/CLI/REST).
- Dependency rule: source dependencies point inward only; outer layers are
  freely replaceable; inner layers are unaware of frameworks.
- Ports-and-adapters (hexagonal): every external capability is a port; the
  number of adapters is unbounded and the kernel is unaware of adapters.
- Framework-free use-case layer: business rules do not depend on Web/CLI/
  storage frameworks; framework replacement within ten years does not touch
  the use-case layer.
- Change isolation: the yearly kernel (L0-L3) and the hourly ecosystem
  (L4-L7) are physically separated, independently versioned, with
  compatibility-matrix gates — explosive ecosystem iteration never pollutes
  the kernel.

## 10. Governance and compliance

- License: AGPLv3 + commercial dual licensing (DUAL-LICENSE.md).
- Mandatory copyright header: every source file must carry the unified
  copyright header (`scripts/header-check.sh` + goheader double gate; missing
  → FAIL).
- Digital code watermark: 小黎的部落 — triple watermark of copyright header +
  build metadata (injected via -ldflags) + release-manifest SHA256.
- GPG signing: tags, release assets, checksum manifests, and SBOM are all GPG
  signed (`scripts/release.sh`).

## 11. Release cadence

- Kernel yearly: one major version per year (year ring), LTS 3 years; kernel
  APIs semantically versioned.
- Ecosystem hourly: plugin-repository CI publishes hourly, compatibility
  matrix verified automatically, hot-loading upgrades seamlessly.

## 12. Positioning

On the chaos side, referenced peers are Gremlin / Chaos Mesh / Litmus /
ChaosBlade; on the mutation side, Stryker / mull / PIT. Abate's stated
differentiation is dual-kernel unification + WASM ecosystem + open
specification + deterministic simulator: competitors ship core releases
quarterly while the Abate ecosystem publishes hourly; competitors bind to
proprietary platforms while Abate binds to the open CES; competitors depend
on real clusters while Abate's pure-software twin does not go stale. (These
are the original document's positioning statements, translated faithfully;
no absolute scores or surpassing claims are added.)
