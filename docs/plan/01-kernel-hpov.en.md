<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->

> Language / 语言：[English](01-kernel-hpov.en.md) | [中文](01-kernel-hpov.md)

# Abate L0 Kernel Plan — HPOV Kernel Six Modules in Extreme Detail

**Design-target statement (mandatory for C-class translations)**: this file
is the English translation of an internal planning document. All `100%` /
"pass" / "met" style statements in the full text are **Design targets /
Acceptance targets**, **not measured results**; they must not be read as
established facts or capability claims. The translation corresponds
section by section, does not escalate the tone, and does not delete target
qualifiers.

- Document ID: PLAN-01　Status: in force　Upstream: PLAN-00
- Covered modules: KRN-ALG / KRN-SCH / KRN-OBS / KRN-VER / KRN-MUT / KRN-INJ
- Common constraints (observed by all six modules):
  1. Depend only on the Go standard library and kernel-owned packages;
     importing any L1-L7 package is forbidden (enforced by the depguard
     gate).
  2. All public types and functions must carry Go doc comments (godot
     trailing-period gate).
  3. All randomness must be injected via SIM-RNG; direct use of the global
     math/rand source and time reads is forbidden (virtual-clock gate).
  4. Coverage: statement ≥95%, branch ≥85%; a determinism replay test
     (same-seed double run with identical hashes) is a merge prerequisite.

---

## Module KRN-ALG: perturbation algebra

### 1. Positioning and responsibilities

Defines the type system of perturbation operators and the four algebraic
operations (composition ∘, concurrent product ⊗, dual p⁻¹, identity ε), plus
an algebraic-law property verifier. It is the mathematical foundation of the
whole engine: mutation (M family) and chaos (C family) operators both
implement the same Operator interface; the scheduler and verifier program
against the algebraic structure only and are unaware of concrete operator
semantics.

### 2. Package structure and file layout

```
kernel/algebra/
  operator.go        // Operator interface, Family enum, operation error types
  composite.go       // Composition operation Compose
  concurrent.go      // Concurrent-product operation Concurrent
  dual.go            // Dual (rollback) operation Inverse
  identity.go        // Identity operator Identity
  laws.go            // Algebraic-law property verifier (associativity/identity/dual compensation)
  laws_test.go       // Property tests (1000 random combinations)
  doc.go             // Package doc: algebra definition and notation table
```

### 3. Core types and interfaces (must implement)

- The `Operator` interface:
  - `ID() OperatorID`: globally unique operator identifier (string; naming
    convention: family-prefix-name-version, e.g. m-arith-replace-v1,
    c-network-delay-v1).
  - `Family() Family`: returns the M or C family.
  - `Params() Params`: operator parameters (ordered key/values, serializable,
    participating in the deterministic hash).
  - `Inverse() (Operator, error)`: returns the dual operator;
    non-invertible operators return ErrNotInvertible (allowed only for
    C-family operators explicitly declared non-invertible, and a compensation
    strategy must be provided simultaneously).
  - `Fingerprint() [32]byte`: deterministic fingerprint (SHA-256 over
    ID+family+canonical parameter serialization), used for replay hashes and
    deduplication.
- `Compose(p1, p2 Operator) (Operator, error)`: composition; semantics are
  apply p2 first, then p1; verifies the two operators' target domains are
  compatible (incompatible returns ErrIncompatible).
- `Concurrent(p1, p2 Operator) (Operator, error)`: concurrent product;
  requires no target conflict between the two operators (conflict detection:
  rejection if target key sets intersect, returning ErrTargetConflict).
- `Identity(target TargetKey) Operator`: the identity operator.
- The algebraic-law verifier `Laws`:
  - `CheckAssociativity(p1, p2, p3) bool`:
    Compose(p1, Compose(p2, p3)) ≡ Compose(Compose(p1, p2), p3).
  - `CheckIdentity(p) bool`: Compose(p, ε) ≡ p ≡ Compose(ε, p).
  - `CheckDual(p) bool`: after applying p then p⁻¹, the target state returns
    to baseline (within tolerance).

### 4. Key algorithms and flows

1. Fingerprint computation: parameters sorted by key name, then canonical
   serialization (map iteration order leakage forbidden), guaranteeing same
   parameters → same fingerprint (determinism iron rule).
2. Composition expansion: the Compose result is a lazy wrapper node; Apply
   expands in right-associative order; depth limit 64 (operator-bomb guard),
   over-limit returns ErrTooDeep.
3. Concurrent-product expansion: on Apply, the two operators are applied
   concurrently within the same virtual time slice; event sequence numbers
   are totally ordered by the scheduler.

### 5. Dependencies and constraints

- Depends on: standard library only (crypto/sha256, sort, errors, fmt).
- Depended on by: KRN-SCH (schedules operators), KRN-VER (verifies dual
  compensation), KRN-MUT/KRN-INJ (operator implementers).
- Forbidden: any I/O, any time reads, any global state.

### 6. Test requirements

- Unit tests: table-driven, covering normal / boundary (empty params, max
  depth, non-invertible operators) / error (target conflict, incompatible)
  paths.
- Property tests: randomly generate 1000 operator combinations; all three
  algebraic laws must hold.
- Determinism tests: construct the same operator tree with the same seed 100
  times; fingerprints byte-identical.

### 7. Acceptance criteria (quantifiable)

1. All six Operator interface methods implemented with godoc; the four
   operations + three-law verifier all in place.
2. Property tests: 100% of 1000 random combinations pass (design target).
3. Coverage: statement ≥95%, branch ≥85%.
4. Fingerprint determinism: 100 double runs with identical hashes.
5. No TODO/FIXME (godox gate); five-stage gate chain all green.

---

## Module KRN-SCH: deterministic scheduler

### 1. Positioning and responsibilities

A discrete-event simulation (DES) scheduler with the virtual clock as the
only time source and a totally ordered event queue at its core. It is
responsible for the experiment's deterministic execution: same seed and same
configuration necessarily produce a byte-for-byte identical event stream. It
is the direct landing point of the 「稳」 (stable) character.

### 2. Package structure and file layout

```
kernel/scheduler/
  event.go          // Event type: virtual time + total-order sequence number + payload
  queue.go          // Priority queue: dequeues in lexicographic (virtual time, seq) order
  scheduler.go      // Scheduling loop: take event → execute → enqueue derived events
  budget.go         // Execution budget: event count / virtual duration / operator cost, triple caps
  scheduler_test.go // Determinism and throughput tests
```

### 3. Core types and interfaces (must implement)

- `Event`: `At VirtualTime` (integer nanoseconds; floats forbidden), `Seq
  uint64` (total-order sequence number, monotonically increasing), `Kind
  EventKind` (tick/probe/apply/rollback/timeout), `Payload`.
- `Queue`: `Push(Event)`, `Pop() (Event, bool)`; the sort key (At, Seq) is a
  strict total order; duplicate keys forbidden (validated at enqueue;
  conflict returns ErrSeqCollision).
- `Scheduler`:
  - `Run(ctx, Experiment) Result`: the main loop.
  - `Now() VirtualTime`: current virtual time.
  - `Emit(Event)`: enqueue a derived event (for operator and probe callbacks).
- `Budget`: `MaxEvents` (default 1,000,000), `MaxVirtualDuration` (default 24h
  virtual time), `MaxCost` (operator-cost sum cap); exceeding any cap
  terminates the experiment and marks BUDGET_EXCEEDED (runaway-experiment
  guard).

### 4. Key algorithms and flows

1. Main loop: Pop event → advance the virtual clock to At → dispatch by Kind
   → enqueue execution products (observation events / derived events) →
   loop until the queue is empty, the budget is exhausted, or the steady-state
   verdict completes.
2. Total-ordering iron rule: all concurrency (goroutines) is allowed only at
   simulator boundaries; before entering the queue, events must carry the
   scheduler-assigned Seq so the event stream can be replayed in total order.
3. Clock advancement: monotonic, never backwards; rollback events also take
   effect only at future virtual times (the eventually-consistent semantics
   of dual compensation).

### 5. Dependencies and constraints

- Depends on: SIM-CLK (clock interface), SIM-RNG (random stream), KRN-ALG
  (operator types).
- Depended on by: KRN-VER (consumes observation events), CTL-CLI (drives
  experiments).
- Forbidden: real time reads (time.Now disabled; vet + gate scan), float
  times, map iteration order participating in scheduling decisions.

### 6. Test requirements

- Unit tests: queue total ordering (100k events enqueued randomly, dequeued
  strictly in order); each of the three budget caps triggers.
- Determinism tests: same seed, same experiment, double run; the event
  stream's SHA-256 matches; 100 loops.
- Throughput benchmark: single process ≥100k events/second (first M2
  performance-budget item).

### 7. Acceptance criteria (quantifiable)

1. Event-stream determinism: 100% hash agreement over 100 double runs
   (design target).
2. Queue total-order test: 100k events, 0 conflicts (the negative case of
   Seq-conflict rejection covered separately).
3. Each of the three budget caps has a triggering case with the correct
   termination code.
4. Throughput benchmark ≥100k events/second (benchmark report filed).
5. Coverage ≥95%/85%; all gates green.

---

## Module KRN-OBS: observation pipeline

### 1. Positioning and responsibilities

Normalizes the three observation source classes — probe data, metrics,
traces — into an observation event stream with total-order sequence numbers,
for the verifier to consume and SIM-REC to record. It is the direct landing
point of the 「准」 (accurate, zero missed observations) character.

### 2. Package structure and file layout

```
kernel/observe/
  observation.go    // Observation type: virtual time + seq + probe ID + value + status
  probe.go          // Probe interface and registry
  sink.go           // Sink interface: unified consumption point for verifier/recorder/reporter
  pipeline.go       // Pipeline: normalize → assign seq → fan out
  pipeline_test.go
```

### 3. Core types and interfaces (must implement)

- `Observation`: `At VirtualTime`, `Seq uint64` (scheduler-assigned), `ProbeID
  string`, `Value Value` (four types: number/bool/string/histogram), `Status
  ProbeStatus` (ok/degraded/failed/timeout).
- The `Probe` interface: `ID() string`, `Kind() ProbeKind`
  (http/shell/metric/prom), `Collect(ctx, at VirtualTime) (Value, Status)` —
  executes inside the simulator; real network forbidden.
- The `Sink` interface: `Observe(Observation) error`; implementers: KRN-VER
  (assertions), SIM-REC (recording), RPT-* (reporting, injected via ports).
- `Pipeline`: `Register(Probe)`, `Attach(Sink)`, `Run(ctx)` — each tick
  collects all probes in registration order (stable sort), normalizes, and
  fans out to all Sinks.

### 4. Key algorithms and flows

1. Stable collection order: probes are collected in lexicographic ID order,
   eliminating registration-order nondeterminism.
2. Fan-out semantics: a Sink's consumption failure does not block other
   Sinks (the error is recorded into the observation stream itself as a meta
   observation).
3. Timeout: if probe collection exceeds the virtual-time budget, it is
   marked with the timeout status (without blocking the scheduling loop).

### 5. Dependencies and constraints

- Depends on: SIM-CLK, KRN-SCH (Seq assignment).
- Depended on by: KRN-VER, SIM-REC, RPT-* (via ports).
- Forbidden: real I/O; probe implementations always target the simulator
  (real-environment probes belong to ecosystem-layer injector plugins).

### 6. Test requirements

- Unit tests: four-type value normalization, timeout path, Sink failure
  isolation.
- Determinism tests: same-seed double run; the observation stream's hash
  matches.
- Integration tests: joint testing with KRN-SCH; the probe collection order
  matches the event order.

### 7. Acceptance criteria (quantifiable)

1. All four Value types have lossless encode/decode round-trip tests.
2. Observation-stream determinism: 100 double runs agree.
3. Sink failure isolation case: one failing Sink does not affect the other
   Sinks or the scheduling loop.
4. Coverage ≥95%/85%; all gates green.

---

## Module KRN-VER: verifier

### 1. Positioning and responsibilities

Consumes the observation stream and outputs experiment verdicts: steady-state
assertions (HELD/VIOLATED), mutation verdicts
(KILLED/SURVIVED/EQUIVALENT), and the survival score (mutation score). It is
the V of HPOV and the core value outlet of mutation testing.

### 2. Package structure and file layout

```
kernel/verify/
  steady.go         // Steady-state assertion: probe + tolerance + duration
  verdict.go        // Verdict type and verdict state machine
  score.go          // Survival score and equivalence analysis
  verify.go         // Verify main entry: observation stream → verdict
  verify_test.go
```

### 3. Core types and interfaces (must implement)

- `SteadyState`: `ProbeID`, `Expectation` (value/range/predicate),
  `Tolerance` (absolute/relative/percent), `HoldFor VirtualTime` (the minimum
  virtual duration the steady state must hold).
- `Verdict`: `Outcome` (HELD/VIOLATED/INCONCLUSIVE), `Reason` (the Seq of the
  first violating observation plus expected/actual values), `Timeline` (the
  violation timeline).
- `MutationVerdict`: `KILLED` (the tests killed the mutant) / `SURVIVED` (the
  tests did not kill it) / `EQUIVALENT` (equivalent mutant, judged by
  equivalence analysis) / `TIMEOUT` (over budget).
- `Score`: `Total`, `Killed`, `Survived`, `Equivalent`, `Timeout`;
  `Score = Killed / (Total - Equivalent)` (equivalents excluded from the
  denominator; the industry-standard caliber).
- `Verify(obs []Observation, ss []SteadyState) []Verdict`: pure function,
  stateless.

### 4. Key algorithms and flows

1. Steady-state verdict: HELD only when observations stay within tolerance
   and hold for the full HoldFor; any out-of-tolerance observation is
   VIOLATED, recording the first violating point.
2. Mutation verdict: run the test suite against the mutant (inside the
   simulator) → test failure means KILLED; all passing means SURVIVED.
3. Equivalence analysis (three-stage): static equivalence (AIR structural
   equivalence check) → dynamic equivalence (byte-for-byte comparison of the
   observation streams for identical inputs) → mark EQUIVALENT (suspected
   equivalents enter a manual-review queue, never auto-scored).

### 5. Dependencies and constraints

- Depends on: KRN-OBS (observation types), LNG-AIR (static equivalence
  checking, injected via ports; the kernel defines only the interface).
- Depended on by: RPT-* (report verdicts), CTL-* (display).
- Forbidden: any I/O or time reads inside Verify; verdicts must be pure
  functions (replayable and re-judgeable).

### 6. Test requirements

- Unit tests: three tolerance types (absolute/relative/percent) × boundaries
  (exactly at the tolerance boundary), full combination; HoldFor boundary
  (one tick short).
- Verdict tests: construction cases for all four MutationVerdict classes;
  score computation includes the equivalent-exclusion caliber.
- Replay re-judgment tests: replay a recorded observation stream and
  re-judge; the Verdicts are byte-for-byte identical.

### 7. Acceptance criteria (quantifiable)

1. Tolerance × boundary full-combination table-driven cases pass 100%
   (design target).
2. The four mutation verdict classes and the score caliber match the PIT
   reference implementation in the parity check (parity report filed).
3. Replay re-judgment byte-identical over 100 runs.
4. Coverage ≥95%/85%; all gates green.

---

## Module KRN-MUT: mutation core (PIT algorithm source-level port)

### 1. Positioning and responsibilities

The generation and scheduling engine for M-family operators: mutation-point
discovery, mutant generation, coverage-guided scheduling, and minimization
(locating the first killing test). The algorithm originates from PITest,
ported via the five-step inclusion process into kernel-owned Go code.

### 2. Five-step inclusion execution details (item by item)

1. **Source vendor**: the PIT source at a fixed version into
   `third_party/pit/` (read-only reference baseline; upstream commit hash and
   license-review verdict recorded: PIT is Apache-2.0 and may be included).
2. **Secondary changes**: the core algorithms rewritten in Go and
   determinized — mutation-point selection is driven by SIM-RNG seeds (all
   PIT-native nondeterministic points replaced); scheduling order totally
   ordered.
3. **Optimization**: port only the four core algorithm lines (mutation
   operator families / coverage guidance / minimization / priority ordering);
   trim reporting, CLI, Maven integration, and other non-core paths.
4. **Deep integration**: operator families implement KRN-ALG's Operator
   interface (M family); mutation-point discovery targets AIR (LNG-AIR) rather
   than JVM bytecode — the architectural difference from PIT: Abate mutations
   act on the source-AST layer, naturally multi-language.
5. **Encapsulation**: expose only the `Engine` type and
   `Run(ctx, Target, Suite) Report` externally; PIT concepts (MutationDetails
   etc.) must not appear in kernel APIs.

### 3. Package structure and file layout

```
kernel/mutation/
  engine.go        // Engine: orchestrates mutation experiments
  operators.go     // M-family operator registry (operator replacement / boundary value / constant perturbation / control flow / return value / negated condition / call removal)
  discover.go      // Mutation-point discovery: AIR traversal → candidate points
  schedule.go      // Coverage-guided scheduling: sort mutation points by test coverage (a port of PIT's thread-priority algorithm)
  minimize.go      // Minimization: binary-locate the first killing test (a port of PIT's minimization algorithm)
  engine_test.go
```

### 4. Core types and interfaces (must implement)

- `Engine`: `Run(ctx, Target, Suite) (Report, error)` — Target is an AIR unit
  (function/method); Suite is a test selector.
- `MutationPoint`: `AIRNodeRef` (node reference), `Operator OperatorID`,
  `Priority` (coverage-guided priority).
- `Report`: per-mutant `MutationVerdict` list + aggregate `Score` +
  deterministic fingerprint (seed + target fingerprint + suite fingerprint).
- The seven base M-family operator classes (the first version must have them
  all): operator replacement (+↔-, *↔/, &&↔||), boundary-value perturbation
  (<↔<=, >↔>=), constant perturbation (+1/-1/zero), control-flow perturbation
  (remove break/continue), return-value perturbation (true↔false, nil↔non-nil),
  negated conditions, call removal.

### 5. Dependencies and constraints

- Depends on: KRN-ALG, KRN-SCH, KRN-VER, LNG-AIR (via port interfaces).
- Forbidden: JVM runtime dependency (architecturally unnecessary); PIT
  proprietary type leakage (an API review item).

### 6. Test requirements

- Unit tests: each of the seven operator classes has positive/negative
  examples; scheduling priority order is deterministic.
- Parity tests (key acceptance): construct a Go baseline target set; the
  consistency rate between Abate's ported operators and the PIT reference
  implementation over the same mutation-point set is 100% (design target; the
  PIT side is produced via its official CLI; the parity script and report are
  filed).
- Determinism tests: same-seed double run; the Report hash matches.

### 7. Acceptance criteria (quantifiable)

1. All seven base operator classes land and the registry is enumerable.
2. PIT parity consistency rate 100% (design target; parity report filed with
  each difference explained).
3. Coverage-guided scheduling and minimization each have end-to-end cases.
4. Same-seed 100 double runs: Report hashes identical.
5. Coverage ≥95%/85%; all gates green; no PIT proprietary types in kernel
  APIs (review checklist checked).

---

## Module KRN-INJ: injection core (chaosblade source secondary development)

### 1. Positioning and responsibilities

The execution engine for C-family operators: injects chaos faults
(network/resource/clock/process/dependency) deterministically into simulator
targets and provides dual rollback. Based on chaosblade source secondary
development (also Go, natural affinity).

### 2. Five-step inclusion execution details (item by item)

1. **Source vendor**: the chaosblade source at a fixed version into
   `third_party/chaosblade/` (upstream commit hash recorded; license
   Apache-2.0, may be included).
2. **Secondary changes**: determinize the execution model — injection actions
   retargeted to simulator calls against SIM-NET/SIM-FS/SIM-SCHED (all
   original real-system call paths rerouted); random parameters driven by
   SIM-RNG seeds.
3. **Optimization**: trim the CLI shell, channel/cloud-vendor integrations,
   and the exec launch path; converge to a pure library shape.
4. **Deep integration**: injection actions translate directly into C-family
   Operators (implementing the KRN-ALG interface); rollback implements
   Inverse (the dual operation); scenario models map to CES vocabulary.
5. **Encapsulation**: expose only the `Injector` type and
   `Inject(ctx, op Operator) error` / `Rollback(ctx, op Operator) error`;
   chaosblade-specific concepts (Model/ExpModel etc.) must not leak.

### 3. Package structure and file layout

```
kernel/inject/
  injector.go      // Injector: inject/rollback entry
  network.go       // Network class: latency/partition/packet loss/limited bandwidth (lands on SIM-NET)
  resource.go      // Resource class: CPU/memory/IO/disk stress (lands on SIM-SCHED/SIM-FS)
  clock.go         // Clock class: offset/jitter (lands on SIM-CLK)
  process.go       // Process class: termination/suspension (lands on SIM-SCHED)
  dependency.go    // Dependency class: dependency faults (error injection/timeout)
  injector_test.go
```

### 4. Core types and interfaces (must implement)

- `Injector`: `Inject(ctx, op Operator) error`,
  `Rollback(ctx, op Operator) error` — all targeting the simulator; real
  system calls forbidden.
- The first-version C-family operator set across five classes (must be
  complete): network latency, network partition, packet loss, limited
  bandwidth, CPU stress, memory stress, IO latency, disk exhaustion, clock
  offset, process termination, dependency error injection, dependency
  timeout.
- Every operator must implement `Inverse` (dual rollback); non-invertible
  operators must provide a compensation strategy and declare it explicitly at
  CES validation.

### 5. Dependencies and constraints

- Depends on: KRN-ALG, SIM-NET/SIM-FS/SIM-SCHED/SIM-CLK, SIM-RNG.
- Forbidden: real network operations, real process operations, real
  filesystem writes (all rerouted to the simulator); chaosblade proprietary
  type leakage.

### 6. Test requirements

- Unit tests: per-class operator inject → observe → rollback full-chain
  cases; after rollback the target state returns to baseline (the dual
  compensation law).
- Determinism tests: same-seed double run; the injection event stream's hash
  matches.
- Upstream backport rehearsal: simulate an upstream security patch and walk
  the full review → backport → TCK regression flow (an M2 acceptance item).

### 7. Acceptance criteria (quantifiable)

1. All 12 first-version operators across the five C-family classes land, all
   implementing Inverse or explicit compensation.
2. Baseline recovery rate after inject → rollback is 100% (design target;
   dual-compensation test).
3. Same-seed 100 double runs: event-stream hashes identical.
4. chaosblade CLI shell removal confirmed (no exec launch paths remain in the
   repository; the scan report is filed).
5. Coverage ≥95%/85%; all gates green.

---
## Appendix: L0 kernel overall acceptance (after all six modules)

1. End to end: one CES YAML (steady state + M⊗C hybrid perturbation + probes
   + rollback) executed via CTL-CLI produces a deterministic report; 100
   same-seed double runs are byte-for-byte identical (design target).
2. Air-gapped test: the full kernel test suite passes with the network
   unplugged (in-container network-disabled simulation; the zero-external-
   dependency iron rule).
3. The all-module coverage summary report is filed; the five-stage gate chain
   + benchmark regression all green.
