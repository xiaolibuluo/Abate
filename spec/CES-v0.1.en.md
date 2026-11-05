<!--
Abate — where chaos abates. Open Chaos Experiment Specification (CES) & reference implementation.
Copyright (C) 2026 小黎的部落 (www.qingdujian.com)

This file is published under the GNU Affero General Public License v3 (AGPLv3).
Commercial use requires prior written authorization from 小黎的部落.
See COMMERCIAL.md for details.
-->

# Abate Chaos Experiment Specification (CES) v0.1

> Brand: Abate — where chaos abates
> Publisher: 小黎的部落 (www.qingdujian.com)
> License: AGPLv3; commercial use requires prior written authorization from 小黎的部落
> Status: This specification is the 'weights and measures' of quality engineering. Implementations are interchangeable; the specification remains conserved.
> Scope: v0.1 covers dual kernels — dynamic perturbation (chaos experiments, Chapters 1–8) and static perturbation (mutation experiments, Chapter 9), unified under the HPOV perturbation algebra.

---

## Chapter 0 · Foreword and Governance

### 0.1 Specification Status
- CES defines the universal language, execution semantics, scoring algorithm, and certification methodology for perturbation experiments: dynamic perturbation = chaos experiments, static perturbation = mutation experiments, both unified under the HPOV paradigm (Hypothesis → Perturbation → Observation → Verdict).
- Any implementation conforming to CES is called a CES Compatible implementation. The Abate reference engine is the first implementation, but not the only one.
- Clause numbers, once published, are permanently valid; they may be deprecated but never recycled.

### 0.5 Perturbation Algebra and Unified Verdict
- [CES-0.5.1] All mutants and all chaos faults are perturbations: mutation is static perturbation, chaos is dynamic perturbation.
- [CES-0.5.2] Verdict semantics are unified to two values: Contained (controlled by the observation mesh) and Escaped (breached the observation mesh). Mutant killed = Contained; survived = Escaped. Steady state held = Contained; broken = Escaped.
- [CES-0.5.3] Mathematical definition of 100% full coverage: Escaped = 0.
- [CES-0.5.4] Unified Quality Score TQS = mutation score × chaos score × coverage score, geometrically aggregated (integer cube root); if any layer is zero, the total is zero; gating requires all three layers to independently meet the threshold — no long-board compensation.

### 0.2 Versioning Strategy
- Annual LTS cadence; semantic versioning MAJOR.MINOR.PATCH.
- From v1.0 the specification is frozen: additions only, backward-compatible.
- Compatibility commitments are written into GOVERNANCE; maintainers cannot unilaterally change them.

### 0.3 RFC Process
- Specification changes require an RFC: motivation, clause draft, impact analysis.
- Hard gate: two independent implementations corroborate feasibility + no less than 30-day public review.

### 0.4 License and Commercial Authorization
- The specification text and reference implementation are published under AGPLv3.
- Any commercial use (including providing Abate-based services as SaaS) requires prior written authorization from 小黎的部落.

## Chapter 1 · Terminology and Fault Taxonomy

### 1.1 Design Rationale
The fault domain is conserved: latency, crash, network partition, resource exhaustion, data corruption, clock skew — this taxonomy has not changed in thirty years and is the foundation for a decade of zero-maintenance specification. Platform API changes are isolated in the adapter layer (Chapter 7).

- [CES-1.1.1] The fault codes and terminology in this chapter constitute the specification's conserved layer: once published, semantics may not be modified; new entries may be added but must not overlap semantically with existing ones.
- [CES-1.1.2] Implementations, adapters, and toolchains must use the codes and terms defined in this chapter in all user-facing output (DSL, API, reports, logs, error codes, badges); coining synonyms is prohibited.

### 1.2 Fault Six-Classification (codes permanently fixed)
| Code | Class | Definition | Typical Parameters |
|---|---|---|---|
| F1 | Latency | Injection and jitter of call, I/O, and scheduling delays | delay, jitter, percent |
| F2 | Crash | Termination and restart of processes, instances, components | signal, gracePeriodMs |
| F3 | Partition | Network disconnection, packet loss, DNS pollution, route isolation | loss, direction, peer |
| F4 | Exhaustion | CPU, memory, disk, file-descriptor, connection-pool exhaustion | resource, utilization, fillRate |
| F5 | Corruption | Data bit-flip, reordering, duplication, loss | bitRate, pattern |
| F6 | ClockSkew | System clock offset and rollback | offsetMs |

- [CES-1.2.1] All dynamic perturbation faults must be classified under F1–F6; unclassifiable faults must not be injected and require an RFC per 0.3 to extend the taxonomy.
- [CES-1.2.2] Compound faults are expressed as fault sequences; mixed-class codes are prohibited; each element in the sequence independently declares its code and parameters and is independently counted in the coverage matrix.
- [CES-1.2.3] Fault codes are case-sensitive and must be verbatim identical across DSL, API, reports, error codes, and adapter capability declarations.

### 1.3 Sub-codes
- Sub-codes may be registered beneath the six classifications (e.g., F3.dns, F4.fd); the sub-code table is maintained in Appendix B (ChaosHub classification codes). New sub-codes follow the MINOR version process and do not touch the conserved layer.
- [CES-1.3.1] Sub-code format is '{F-code}.{lowercase letters or digits}', globally unique; adapters with incomplete mappings must explicitly declare degradation per Chapter 7 and must not silently map to the parent code.

### 1.4 Standard Terminology Table
One term, one meaning; terminology, error codes, test names, and documentation are cross-linked. The table below is normative:

| Term | English | Definition |
|---|---|---|
| Experiment | Experiment | One complete HPOV closed loop (hypothesis → perturbation → observation → verdict); the specification's minimum executable unit |
| Steady-State Hypothesis | Steady-State Hypothesis | The set of SLI assertions that should hold throughout the injection window; the experiment's falsifiable proposition |
| Blast Radius | Blast Radius | The maximum scope a perturbation may affect, declared as three-dimensional upper bounds: instance count, duration, traffic ratio (CES-2.5.1) |
| Abort Condition | Abort Condition | A stop condition that triggers unconditional transition to Recovering with no manual bypass (CES-3.1.1) |
| Weakness | Weakness | A defect detected and reproducible in an experiment, with lifecycle states per Chapter 6 |
| Regression Drill | Regression Drill | A periodic replay experiment targeting a fixed weakness to prevent regression |
| Evidence Level | Evidence Level | Observation method rigor grade A/B/C (Chapter 5), determining experiment weight discount (CES-4.3.1) |
| Criticality | Criticality | Target business importance, integer 1–5, a multiplier in experiment weight (CES-4.3.1) |
| Seed | seed | The global random source for deterministic replay: same experiment definition + same seed must produce the same fault sequence (CES-3.1.3) |
| Observation Mesh | Observation Mesh | The set of all SLI assertions, data invariants, and abort conditions declared by the experiment; the boundary for Escaped verdict |
| Verdict | Verdict | The experiment's terminal conclusion, unified two-value semantics Contained / Escaped (CES-0.5.2) |
| Adapter | Adapter | A plugin mapping third-party execution backend capabilities to F1–F6 codes (Chapter 7), the only part of the specification allowed to evolve with upstream |

- [CES-1.4.1] The terminology table prohibits polysemy and synonymy; synonymous expressions in implementation output are treated as documentation defects.
- [CES-1.4.2] Every term must be reversely traceable: term → specification clause → certification test case → error code; all four links must be present before a term enters the table.

## Chapter 2 · Experiment Model (DSL Specification)

### 2.1 Carrier and Grammar Subset
- Declarative YAML; files use .abate.yaml as the recommended suffix.
- The top-level five-tuple is a mandatory structure: target / fault / blastRadius / steadyState / abortOn.
- [CES-2.1.1] Implementations must not reject experiment files due to unknown fields (forward-compatible reservation).
- [CES-2.1.2] Validation failures must produce structured errors: field path + expected + actual + clause ID.
- [CES-2.1.3] The normative DSL is a strict line-oriented YAML subset: two-space indentation expresses hierarchy; only mappings, scalars, and list items are supported; nesting depth must not exceed three levels (top → section → list-item field). Flow syntax (braces and brackets), anchors, aliases, and multi-document markers are not in the subset; implementations must reject them with structured errors.
- [CES-2.1.4] Numeric fields are always decimal integers: durations in milliseconds (windowMs, maxDurationMs, recoveryBudgetMs), ratios in basis points (maxTraffic); numeric fields carrying unit suffixes (ms, m, %) are structured errors.

### 2.2 Top-Level Fields and Target Addressing
Top-level fields:

| Field | Type | Required | Semantics |
|---|---|---|---|
| id | string | yes | Experiment unique identifier, aggregation sort key (CES-4.1.3) |
| seed | int64 | yes | Deterministic replay seed (CES-3.1.3) |
| criticality | int | yes | Target criticality 1–5 (CES-4.3.1) |
| recoveryBudgetMs | int64 | no | Recovery budget, default 300000 (CES-4.4.2) |
| target | section | yes | Experiment target (this section) |
| fault | section | yes | Injected fault (2.3) |
| blastRadius | section | yes | Blast radius (2.5) |
| steadyState | list | yes | Steady-state assertion set (2.6) |
| abortOn | list | no | Abort condition set (2.7) |
| params | section | no | Implementation extension parameters, tolerated per CES-2.1.1 |

- [CES-2.2.1] target supports three addressing modes: service, instance, dependency; selector semantics must be deterministic.
- [CES-2.2.2] target.kind is limited to service / instance / dependency; target.name is required and non-empty; key-value pairs other than kind and name constitute the selector.
- [CES-2.2.3] Selectors are key-value equality matches, ANDed across keys; match results must be stably sorted in lexicographic order; the same experiment file must resolve to the same target set on any compliant implementation.

### 2.3 Fault Definition
- [CES-2.3.1] fault.type must use the Chapter 1 classification codes; the parameter set is defined per class.
- [CES-2.3.2] Key-value pairs other than fault.type constitute fault parameters; the parameter set must be non-empty; keys and values are strings for cross-plugin serializability; parameter validity is verified against the Appendix B classification parameter table and execution backend capability declarations.

### 2.4 Injection Semantics
- [CES-2.4.1] Injection actions are idempotent: repeated execution does not produce cumulative effects.

### 2.5 Blast Radius
- [CES-2.5.1] blastRadius must declare at least two of: instance upper bound, duration upper bound, traffic ratio upper bound.
- [CES-2.5.2] The three limit fields are maxInstances (instance count), maxDurationMs (milliseconds), maxTraffic (basis points); undeclared dimensions are treated as unbounded and must be flagged as risk items by Preflight.

### 2.6 Steady-State Assertions
- [CES-2.6.1] steadyState must contain at least one SLI assertion; each assertion includes metric source, window, threshold, and statistical method.
- [CES-2.6.2] Assertion fields: sli (metric source identifier) and assert (threshold expression) are required; method (statistical method identifier) and windowMs (statistical window in milliseconds) are optional; assert evaluation results must be normalized to basis-point integers for scoring (CES-4.1.2).

### 2.7 Abort Conditions
- [CES-2.7.1] abortOn triggers transition to Recovering with no manual bypass.
- [CES-2.7.2] abortOn assertion structure is identical to steadyState (CES-2.6.2); when abortOn is empty, blast radius and steady-state assertions remain mandatory; abort monitoring is covered by implementation safety policy and declared in the report.

### 2.8 Normative Example
The following example is parseable by the reference implementation dsl.Parse and passes all validations (all numeric fields are integers, CES-2.1.4):

```yaml
id: order-service-latency
seed: 20260101
criticality: 4
recoveryBudgetMs: 300000
target:
  kind: service
  name: order-service
  env: staging
fault:
  type: F1
  delay: 500
  jitter: 100
  percent: 3000
blastRadius:
  maxInstances: 1
  maxDurationMs: 300000
  maxTraffic: 1000
steadyState:
  - sli: http_success_rate
    assert: ">= 99.9%"
    method: rate
    windowMs: 60000
abortOn:
  - sli: error_rate
    assert: "> 5%"
    method: rate
    windowMs: 10000
```

## Chapter 3 · Execution Semantics and State Machine

### 3.1 Lifecycle
Pending → Preflight → Injecting → Observing → Recovering → Verdict. The state transition table is normative; illegal transitions are implementation defects.

State definitions:

| State | Definition |
|---|---|
| Pending | Experiment accepted, awaiting admission check |
| Preflight | Admission check in progress: target reachability, blast-radius compliance, observation mesh and credentials ready |
| Injecting | Fault injection in progress (idempotent, CES-2.4.1) |
| Observing | Steady-state observation: SLI assertions, data invariants, and abort conditions continuously monitored |
| Recovering | Rollback and residue cleanup in progress |
| Verdict | Terminal state: verdict complete, report produced |

Legal transition table (normative):

| Source | Legal targets |
|---|---|
| Pending | Preflight; Verdict (cancelled before admission, verdict CANCELLED) |
| Preflight | Injecting; Recovering (abort triggered); Verdict (admission failed, verdict REJECTED) |
| Injecting | Observing (injection window ended); Recovering (abort triggered) |
| Observing | Recovering (observation ended or abort triggered) |
| Recovering | Verdict |
| Verdict | none (terminal, no outgoing edges) |

- [CES-3.1.5] The initial state must be Pending; Verdict is the sole terminal state; no implementation may define outgoing edges from Verdict.
- [CES-3.1.6] Implementations must reject transitions outside the table above and record the source state, target state, and trigger event in a structured error; out-of-table transitions are implementation defects.
- [CES-3.1.7] Every state transition must record a timestamp and trigger reason; the experiment timeline must support complete reconstruction of the state sequence for Chapter 6 reports and Appendix A model checking.

### 3.2 Invariants
- [CES-3.1.1] From any state, a triggered abort condition must transition to Recovering; applicable states are Preflight, Injecting, and Observing. Abort events in Recovering and Verdict are recorded only — no further transition (recovery is non-reentrant).
- [CES-3.1.2] After Recovering completes, no injection side effect may remain; residue must be machine-detectable.
- [CES-3.1.3] The same experiment definition + the same seed must produce a replayable fault sequence (deterministic replay).
- [CES-3.1.4] If Preflight fails, Injecting must not be entered; the failure reason must be recorded in structured form.

### 3.3 Verdict Values
The Verdict state's conclusion takes four values:

| Verdict | Semantics |
|---|---|
| CONTAINED | Perturbation controlled by observation mesh: steady-state hypothesis held throughout injection window (maps to CES-0.5.2 Contained) |
| ESCAPED | Perturbation breached observation mesh: steady state broken or data invariant violated, weakness detected (maps to CES-0.5.2 Escaped) |
| REJECTED | Preflight admission failed; experiment did not execute injection |
| CANCELLED | Cancelled by operator before admission; experiment did not execute injection |

- [CES-3.3.1] Verdict values are limited to the four above; implementations must not extend verdict values; additional information is expressed via weakness lists and report fields.
- [CES-3.3.2] REJECTED and CANCELLED are not 'executed experiments' per Chapter 4 and are excluded from all dimension weight sums, but must be listed individually in the report.
- [CES-3.3.3] An ESCAPED verdict must be associated with at least one weakness record or abort event; evidence-free ESCAPED verdicts are prohibited.

### 3.4 Formalization
- The state machine and invariants are modeled in TLA+/Quint; the model file is published as Appendix A (mandatory).
- Implementations must derive test cases from the Appendix A model checking.
- [CES-3.4.1] The Appendix A model must cover all states, the transition table, and invariants in this chapter; if the model conflicts with the text, the stricter one prevails during the conflict and the discrepancy must be resolved in the next revision.
- [CES-3.4.2] The L1 certification suite (Chapter 8) must include all model-checking test cases derived from Appendix A, and the results must be deterministically reproducible.

## Chapter 4 · Reliability Scoring Algorithm (Mathematical Definition)

### 4.1 General
- [CES-4.1.1] The total reliability score Score ∈ [0.00, 100.00], derived from a weighted sum of five dimensions, displayed with exactly two decimal places.
- [CES-4.1.2] All scoring computations must use basis-point (‱) fixed-point integer arithmetic: floating-point numbers are forbidden in the scoring path, eliminating cross-language and cross-platform rounding drift.
- [CES-4.1.3] Before aggregation, experiment records are sorted ascending by experiment ID lexicographically; all multiplications precede divisions; integer division truncates toward zero. Identical inputs must produce identical outputs on any compliant implementation (reproducibility commitment).

### 4.2 Dimensions and Weights
| Dimension | Symbol | Weight (‱) |
|---|---|---|
| Availability Resilience | AR | 3000 |
| Recovery Speed | REC | 2000 |
| Degradation Capability | DC | 2000 |
| Data Integrity | DI | 2000 |
| Blast-Radius Control | BRC | 1000 |

- All dimension values are in basis points: full marks = 10000 (i.e., 100.00 points).
- [CES-4.2.1] Score_raw = (AR×3000 + REC×2000 + DC×2000 + DI×2000 + BRC×1000) ÷ 10000, the result is still a basis-point value.

### 4.3 Experiment Weight
- [CES-4.3.1] The weight of experiment e is w_e = c_e × e_e.
  - c_e: target criticality, integer 1–5, declared by the experiment with justification; the certification suite verifies consistency between the declaration and topology.
  - e_e: evidence-level coefficient (defined in Chapter 5): A = 10000, B = 8000, C = 5000.
- [CES-4.3.2] Lower-rigor evidence must be discounted in weight; it must not achieve the same scoring status as higher-rigor evidence.

### 4.4 Dimension Formulas
Convention: durations are integer milliseconds; ratios are basis-point integers.

#### 4.4.1 Availability Resilience AR
- hold_e = (cumulative milliseconds during the injection window when all steady-state SLIs hold simultaneously) × 10000 ÷ (total injection window milliseconds).
- [CES-4.4.1] AR = Σ(w_e × hold_e) ÷ Σ(w_e), covering all executed experiments.

#### 4.4.2 Recovery Speed REC
- T_e = milliseconds from abort trigger (or injection stop) to all SLIs returning to baseline and holding for 60s; T_budget is the experiment-declared recovery budget (default 300000ms).
- s_e = max(0, 10000 − T_e × 10000 ÷ T_budget).
- [CES-4.4.2] REC = Σ(w_e × s_e) ÷ Σ(w_e), counting only experiments with recovery events; no recovery event → unmeasured (see 4.6).

#### 4.4.3 Degradation Capability DC
- Applies only to dependency-failure experiments (F3 on dependency targets).
- f_e: degradation switch activated within 10s of injection = 10000, else 0 (machine-verifiable).
- q_e = degraded-mode compliant requests × 10000 ÷ degraded-mode total requests.
- DC_e = f_e × q_e ÷ 10000.
- [CES-4.4.3] DC = Σ(w_e × DC_e) ÷ Σ(w_e).

#### 4.4.4 Data Integrity DI
- d_e = invariants passed × 10000 ÷ total invariants (reconciliation, constraints, idempotency checks).
- [CES-4.4.4] Hard rule: any experiment detecting irreconcilable data loss → DI dimension capped at 2000 (20.00 points).
- DI = min(Σ(w_e × d_e) ÷ Σ(w_e), 2000 when hard rule triggered).

#### 4.4.5 Blast-Radius Control BRC
- r_e = min(10000, declared radius × 10000 ÷ observed radius); observed radius = 0 → 10000.
- [CES-4.4.5] Experiments with abort response latency exceeding the declared upper bound → r_e = 0.
- BRC = Σ(w_e × r_e) ÷ Σ(w_e).

### 4.5 Coverage Correction
- C = exercised combinations × 10000 ÷ total combinations in the 'service × fault class × dependency' coverage matrix; target ratio M = 9500 (i.e., 95%).
- Correction factor k = 8000 + 2000 × min(C, M) ÷ M.
- [CES-4.5.1] Score_final = Score_raw × k ÷ 10000. At 95% coverage k = 10000 (no penalty); at zero coverage the floor is 80%.
- [CES-4.5.2] ChaosHub TOP 20 high-frequency fault modes are mandatory coverage-matrix items; unexercised TOP 20 items count 0 in the numerator of C and must be listed individually in the report.

### 4.6 Unmeasured Dimensions
- [CES-4.6.1] Unmeasured dimensions are scored 0 and marked UNMEASURED in the report; this rule is linked to the 95% coverage target to drive drill-plan gap-filling.

### 4.7 Rounding and Display
- [CES-4.7.1] Display score = Score_final ÷ 100; the integer part is points, the remainder is two decimal places; basis-point integers are naturally exact to 0.01 with zero rounding error.
- [CES-4.7.2] The report must output: Score_final, five dimension raw values, C, k, and the unmeasured dimension list.

### 4.8 Freeze Clauses
- [CES-4.8.1] After v1.0, formulas in this chapter may only have dimensions added or typos corrected; any behavioral change requires a MAJOR version bump.
- [CES-4.8.2] The Appendix C golden dataset has the same force as this chapter: implementations must produce exactly the published values for golden inputs.

### 4.9 Benchmarks and Badges
- Scores include industry-anonymous benchmark percentiles (data from ChaosHub, opt-in).
- Badge embedding: README badge data must be produced by a certified implementation.

## Chapter 5 · Observation and Verdict

### 5.1 Evidence Levels
| Level | Method | Applicability | Weight coefficient e_e |
|---|---|---|---|
| A | Control-group comparison + significance testing (e.g., difference-in-differences) | Production and high-stakes verdicts | 10000 |
| B | Single-group time-series baseline comparison | Pre-production and routine drills | 8000 |
| C | Threshold verdict | Smoke tests and developer self-tests | 5000 |
- [CES-5.1.1] Reports must declare the evidence level; causal claims require level A.
- [CES-5.1.2] Evidence level maps to experiment weight coefficient e_e (CES-4.3.1); implementations must not weight above the actual observation method.
- [CES-5.1.3] When the declared level exceeds the actual observation method, the certification suite downgrades to the actual method and records the discrepancy in the compliance report.

### 5.2 Abort Response
- [CES-5.2.1] The upper bound on response latency from abort trigger to injection stop must be declared by the implementation in documentation; the certification suite verifies it by measurement.
- [CES-5.2.2] Response latency is defined as the millisecond difference between the moment the abort condition first holds and the moment the injection-stop instruction takes effect; the measurement must be machine-collectible and recorded in the report timeline (CES-3.1.7).
- [CES-5.2.3] Experiments whose response latency exceeds the declared upper bound have their BRC dimension processed per CES-4.4.5 (r_e = 0).

### 5.3 Observation Mesh and Sampling
- [CES-5.3.1] The observation mesh consists of the steady-state assertion set (steadyState), data invariant set (5.4), and abort condition set (abortOn); the observation mesh is the boundary for Escaped verdict (Chapter 1 terminology).
- [CES-5.3.2] Sampling order and sample points are driven by a seed-based deterministic RNG (CES-3.1.3); replays of the same experiment definition + same seed must produce point-wise identical observation sequences.
- [CES-5.3.3] SLI assertions are evaluated within windowMs using method statistics, compared against assert thresholds to yield hold/not-hold; evaluation and comparison use basis-point fixed-point integers throughout (CES-4.1.2).
- [CES-5.3.4] When observation data is missing (metric source unreachable or no samples within window), the assertion is treated as not-holding and marked DATA_MISSING in the report; silent skipping or imputation is prohibited.

### 5.4 Data Invariants
- [CES-5.4.1] Data invariants (reconciliation, constraints, idempotency checks) are fully verified after Recovering completes; results count toward the DI dimension (CES-4.4.4).
- [CES-5.4.2] Invariant verification must be replayable: verification inputs, data snapshot identifiers, and results are recorded in the report (CES-6.1.1).

## Chapter 6 · Report Format

### 6.1 Structured Schema
- JSON is the normative format: experiment metadata, timeline, SLI evidence, verdict, weakness list, evidence level, seed.
- Markdown is a recommended human-rendering profile; reports must be self-contained and verifiable offline.

Report top-level fields (mandatory):

| Field | Semantics |
|---|---|
| schemaVersion | Report schema version; this version is ces/v0.1 |
| experiment | Experiment metadata: id, seed, criticality, five-tuple full text, recoveryBudgetMs |
| timeline | State-transition timeline: timestamp and trigger reason for each transition (CES-3.1.7) |
| observations | SLI evidence: assertion text, window, evaluation result, DATA_MISSING flag (CES-5.3.4) |
| invariants | Data invariant verification results (CES-5.4.1) |
| verdict | Verdict: CONTAINED / ESCAPED / REJECTED / CANCELLED (CES-3.3.1) |
| weaknesses | Weakness list (6.2) |
| evidenceLevel | Evidence level A / B / C (CES-5.1.1) |
| score | Score breakdown: Score_final, five dimension raw values, C, k, unmeasured dimension list (CES-4.7.2) |

- [CES-6.1.1] Reports must contain seed and all parameters required for replay.
- [CES-6.1.3] Missing or invalid top-level fields are implementation defects; unknown fields in reports are tolerated per the same principle as CES-2.1.1.
- [CES-6.1.4] Reports for REJECTED and CANCELLED experiments must be listed individually; the score field is marked NOT_SCORED (CES-3.3.2).

### 6.2 Weakness Lifecycle
- [CES-6.1.2] Weakness entries must contain a lifecycle state: discovered → archived → fix-tracking → regression-verified → scenario-internalized.
- [CES-6.2.1] Lifecycle state transitions are unidirectional and irreversible; regression verification must be driven by regression drills (Chapter 1 terminology) and must pass before entering scenario-internalized.
- [CES-6.2.2] Weakness entries must contain: global ID, associated experiment ID, associated verdict (ESCAPED), reproduction steps or replay parameters, current state and timestamp.

### 6.3 Verification and Signing
- [CES-6.3.1] Reports must be self-contained: verdicts and scores can be recomputed offline from report data alone, without external services.
- [CES-6.3.2] Reports must contain a full-content hash (SHA-256); the signing chain is built on content hashes; any field change invalidates existing signatures.
- [CES-6.3.3] The Markdown profile is a rendering layer of the JSON source and must not add or remove semantic fields; both forms must point to the1 same JSON source and share its content hash.

## Chapter 7 · Plugin and Adapter Protocol

### 7.1 Plugin Protocol v1
- Process-isolated loading; protocolVersion exchanged at handshake; incompatibility is refused with an explicit error.
- Capability declaration: plugins declare supported fault codes (F1–F6) and target types at startup.
- Error code system: E1xxx stable directory, one documentation page per code (cause, impact, remediation).

- [CES-7.1.1] Plugins are loaded in process isolation; the protocol major version is v1; protocolVersion is exchanged at handshake; mismatch results in refused loading with an explicit error (error code E1001, Appendix D).
- [CES-7.1.2] Capability declaration fields: name, version, protocol, faultClasses, targetKinds; plugins with missing protocol or empty faultClasses must not accept injection tasks.
- [CES-7.1.3] Plugins are classified by four roles: Injector, Probe, Validator, Reporter; a plugin may declare multiple roles but must fully implement each role's interface.
- [CES-7.1.4] Injector.Inject must be idempotent for repeated calls (CES-2.4.1); Recover must be repeatable and ultimately residue-free (CES-3.1.2); the injection handle contains a globally unique ID and the owning experiment ID.
- [CES-7.1.5] Probe.Observe returns a sample triple: metric name, value, Unix millisecond timestamp; sampling timing is seed-driven (CES-5.3.2).
- [CES-7.1.6] Validator.Preflight returns issues that are each hard blockers (CES-3.1.4); Postflight performs recovery verification and residue scanning (CES-3.1.2).
- [CES-7.1.7] Error codes E1xxx are a stable directory: once published, never recycled; each code must have a documentation page (cause, impact, remediation); the full directory is in Appendix D.

### 7.2 Adapter Specification
- Third-party execution backends (Chaos Mesh / Litmus / cloud-vendor FIS, etc.) are integrated via adapters.
- Adapters must map their fault capabilities to the Chapter 1 classification codes; incomplete mappings must be explicitly declared as degraded.
- [CES-7.2.1] Adapters are the only part of the specification allowed to evolve with upstream platforms, released independently of the core version cadence.
- [CES-7.2.2] Adapter capability declarations must list individually the fault capabilities and target addressing modes that can be fully mapped to Chapter 1 codes; upstream capabilities that cannot be fully mapped must not be declared — only explicitly degraded.
- [CES-7.2.3] Adapters must implement the Injector role; Probe, Validator, and Reporter roles may be provided by core built-in fallbacks, with the fallback condition declared in capability declarations and reports.
- [CES-7.2.4] Adapter versions are independent of CES and core version cadence; adapter release, upgrade, and deprecation do not change the force of specification clauses; deprecated adapters must be announced in advance with a migration path.

### 7.3 MCP Server Protocol
- Abate exposes mutation testing and chaos engineering capabilities via an MCP (Model Context Protocol) Server using JSON-RPC 2.0 over stdio, for AI coding assistant and external tool integration.
- [CES-7.3.1] The MCP Server protocol version is `2024-11-05`, transport is JSON-RPC 2.0 over stdio; implementations must support `initialize`, `tools/list`, and `tools/call` methods, with error codes following JSON-RPC 2.0 standard (-32700/-32601/-32602/-32603).
- [CES-7.3.2] The tool set must include `generate_mutants` (mutant generation), `apply_mutant` (mutant application), and `list_fault_classes` (fault class listing); optionally include `generate_test` (LLM test generation, depends on CES-9.5.3 LLM client interface).
- [CES-7.3.3] The `generate_test` tool's provider parameter must support `openai` and `claude`; unsupported providers must return an error, no silent degradation.

## Chapter 8 · Conformance Certification

### 8.1 Certification Tiers
| Tier | Content |
|---|---|
| L1 | Deterministic simulator compliance: run the full specification test suite on the simulator |
| L2 | Real-backend compliance: at least one real adapter backend passes the test suite |

- [CES-8.1.1] L2 certification requires L1 as a prerequisite; implementations that have not passed L1 may not apply for L2.
- [CES-8.1.2] Certification results bind to specification version and plugin protocol major version; upgrading the core version requires re-certification; adapter upgrades do not trigger core re-certification (CES-7.2.1).

### 8.2 Certification Suite
- Publicly runnable: one command produces a compliance report; results are deterministically reproducible.
- Simulator as certification base: pure-software digital twin, never touches real platform APIs, never goes stale.

- [CES-8.2.1] The certification suite produces a compliance report via one command; reruns of the same implementation version + suite version must produce identical results (deterministic reproducibility, same discipline as CES-3.1.3).
- [CES-8.2.2] The certification base is a deterministic simulator with three modes: simulation / staging / production; L1 certification runs in simulation mode, never touching real platform APIs.
- [CES-8.2.3] The suite must cover: DSL conformance tests (Chapter 2), state-machine and invariant tests (Chapter 3, including Appendix A model-checking-derived cases, CES-3.4.2), golden dataset recomputation (Appendix C, CES-4.8.2), report format validation (Chapter 6).
- [CES-8.2.4] Any test case failure means certification failure; failed cases must be listed individually with clause IDs; partial pass is not permitted.

### 8.3 Badges
- Only certified implementations may display the CES Compatible badge; badge style and verification URL specification TBD.

- [CES-8.3.1] Badges are granted only to certified implementations and must annotate the certification tier (L1/L2) and specification version; badge data must be produced by the certified implementation (CES-4.9).
- [CES-8.3.2] Each badge must carry a verification URL: accessing it returns the implementation's certification report content hash (CES-6.3.2) and suite version; verification failure invalidates the badge.
- [CES-8.3.3] When certification expires or the implementation version exceeds the certification scope, the implementer must remove the badge within 30 days or complete re-certification.

## Chapter 9 · Mutation Experiment Specification (MES)

### 9.1 Positioning
Mutation experiments are static perturbations: modify code, run tests, determine kill/survive, proving test-set effectiveness. Mutation and chaos share the HPOV paradigm and unified verdict semantics (CES-0.5).

### 9.2 Mutation Operator Five-Classification (codes permanently fixed, conserved layer)
| Code | Class | Definition |
|---|---|---|
| M1 | Arithmetic | + - * / % interchange (AOR) |
| M2 | Relational | < <= > >= == != interchange and boundary offset (ROR) |
| M3 | Conditional | logical and/or interchange, condition negation (COR/LOR) |
| M4 | ReturnValue | return zero/nil/opposite (RVR) |
| M5 | Statement | delete or duplicate statement (SDL) |

- [CES-9.2.1] New operators must be classified under M1–M5; belonging to none implies a new operator class, a kernel-track change requiring an RFC.
- [CES-9.2.2] Operator implementations must be deterministic: same source location + same seed yields the same mutant.
- [CES-9.2.3] Mutant global identifier format: mut:{file}:{line}:{col}:{class}.

### 9.3 Mutation Score
- [CES-9.3.1] Mutation score = killed mutants × 10000 ÷ non-equivalent mutant total (basis-point fixed-point, CES-4.1.2 same discipline).
- [CES-9.3.1] Equivalent mutants are excluded from the denominator: must be declared via RFC and confirmed by the specification committee, entering the equivalent-mutant registry; no exemption switch at the gate.
- [CES-9.3.3] Full-coverage gate: mutation kill rate 100% (i.e., Escaped = 0, CES-0.5.3); below 100% must list every surviving mutant individually in the report.

### 9.4 Execution and Zero Jitter
- [CES-9.4.1] Mutants execute in isolated sandboxes; CPU instruction count replaces wall-clock for execution cost measurement, physically zeroing performance noise.
- [CES-9.4.2] Mutant selection and execution order are driven by a deterministic RNG (splitmix64) seeded by seed; same seed + same code version reruns must produce identical results.

### 9.5 AI Test-Supplement Flywheel
- [CES-9.5.1] Surviving mutants may serve as supplementary-test inputs to code-generation models; generated test cases must genuinely kill the target mutant and reproduce in CI before admission to the test suite.
- [CES-9.5.2] Ownership of AI-generated test cases belongs to the running project; the specification asserts no rights (license boundary in COMMERCIAL.md).
- [CES-9.5.3] LLM clients must implement the `LLMClient` interface (`GenerateTest(mutant, source) -> (testCode, err)`); the reference implementation provides OpenAI (Chat Completions API) and Anthropic Claude (Messages API) clients, both using the standard library net/http with zero external dependencies.
- [CES-9.5.4] LLM clients must support custom API endpoints (BaseURL) via configuration to be compatible with Azure OpenAI, proxy gateways, and private deployments; API keys must not be hardcoded or logged.
- [CES-9.5.3] Unreviewed AI batch output is prohibited from direct entry into the test suite (anti-poisoning, see SECURITY.md threat model).

## Appendix A · State Machine TLA+ Model (mandatory, CES-3.4)

The model covers all Chapter 3 states, the legal transition table, and invariants. Implementations must exhaustively model-check with TLC or equivalent and derive test cases (CES-3.4.2). Verdict is the sole terminal state: no deadlock except Verdict.

```tla
---------------- MODULE ces ----------------
(************************************************************)
(* Abate CES v0.1 state machine model (Appendix A, CES-3.4)*)
(* Covers Chapter 3 state definitions, transition table,  *)
(* and invariants.                                          *)
(************************************************************)
EXTENDS Naturals

CONSTANTS
    Pending, Preflight, Injecting, Observing, Recovering, Verdict

VARIABLES
    state,     (* current lifecycle state *)
    injected   (* whether injection side effects remain *)

vars == << state, injected >>

Init == /\ state = Pending
        /\ injected = FALSE

(* Pending -> Preflight: admitted to admission check *)
Admit == /\ state = Pending /\ state' = Preflight /\ injected' = injected

(* Pending -> Verdict: cancelled before admission, verdict CANCELLED *)
Cancel == /\ state = Pending /\ state' = Verdict /\ injected' = injected

(* Preflight -> Injecting: admission passed (CES-3.1.4) *)
Pass == /\ state = Preflight /\ state' = Injecting /\ injected' = injected

(* Preflight -> Verdict: admission failed, verdict REJECTED (CES-3.1.4) *)
Reject == /\ state = Preflight /\ state' = Verdict /\ injected' = injected

(* Abort trigger: applicable states force transition to Recovering (CES-3.1.1) *)
RedLine == /\ state \in {Preflight, Injecting, Observing}
           /\ state' = Recovering /\ injected' = injected

(* Injecting -> Observing: injection window ended *)
WindowEnd == /\ state = Injecting /\ state' = Observing /\ injected' = injected

(* Observing -> Recovering: observation ended *)
ObserveEnd == /\ state = Observing /\ state' = Recovering /\ injected' = injected

(* Recovering -> Verdict: recovery complete, no residue (CES-3.1.2) *)
Recover == /\ state = Recovering /\ state' = Verdict /\ injected' = FALSE

(* Verdict terminal self-loop: terminal state is not deadlock (CES-3.4) *)
Done == /\ state = Verdict /\ UNCHANGED vars

Next == Admit \/ Cancel \/ Pass \/ Reject \/ RedLine
        \/ WindowEnd \/ ObserveEnd \/ Recover \/ Done

Spec == Init /\ [][Next]_vars

(* Invariant 1: state value is legal *)
TypeOK == state \in {Pending, Preflight, Injecting, Observing, Recovering, Verdict}

(* Invariant 2: terminal state is residue-free (CES-3.1.2) *)
ResidueFree == state = Verdict => injected = FALSE
================
```

Model-checking obligations: exhaustively verify TypeOK and ResidueFree; no deadlock except Verdict; RedLine action is always enabled in applicable states (abort cannot be bypassed). An equivalent Quint model is acceptable provided its semantics are not weaker (CES-3.4.1).

## Appendix B · ChaosHub Classification Codes (F1–F6 Sub-code Table)

Sub-code format per CES-1.3.1; new sub-codes follow the MINOR version process. v0.1 seed table (20 entries):

| Sub-code | Fault Mode | Typical Parameters |
|---|---|---|
| F1.1 | Fixed delay injection | delay |
| F1.2 | Delay jitter | delay, jitter, percent |
| F1.3 | Timeout injection | timeoutMs |
| F2.1 | Process termination | signal |
| F2.2 | Process restart | signal, gracePeriodMs |
| F3.1 | Packet loss | loss, direction |
| F3.2 | DNS pollution | domain, mode |
| F3.3 | Network partition | peer, direction |
| F3.4 | Route isolation | cidr |
| F4.1 | CPU exhaustion | utilization |
| F4.2 | Memory exhaustion | utilization, fillRate |
| F4.3 | Disk space exhaustion | utilization |
| F4.4 | File-descriptor exhaustion | count |
| F4.5 | Connection-pool exhaustion | count |
| F4.6 | Disk IOPS exhaustion | iops |
| F5.1 | Bit flip | bitRate, pattern |
| F5.2 | Packet duplication | percent |
| F5.3 | Packet reordering | percent |
| F6.1 | Clock offset | offsetMs |
| F6.2 | Clock rollback | offsetMs |

- ChaosHub TOP 20 high-frequency fault modes (CES-4.5.2 mandatory test items) v0.1 seed list is all 20 entries above; the list is revised and announced with each MINOR version in sync with ChaosHub data.
- Adapters referencing ChaosHub fault modes must provide a mapping table between sub-codes and upstream mode identifiers; incomplete mappings are explicitly degraded per CES-7.2.2.
- [CES-B.1] The ChaosHub sub-code registry (Registry) is the sole authoritative source of sub-code data; adapters and scorers must obtain sub-code lists from the registry, not hardcode them.
- [CES-B.2] The remote syncer (Syncer) fetches sub-code table JSON via HTTP from the ChaosHub data source (format: `{version, codes: [{code, name, params}]}`); on sync failure, it falls back to the local seed table and must not block experiment execution.
- [CES-B.3] After synchronization, a version announcement (Announcement) must be generated listing three diff categories: added, updated, and removed; the announcement format is human-readable text for CI output and release notes.

## Appendix C · Scoring Golden Dataset (same force as Chapter 4, CES-4.8.2)

The expected values below are hand-computed per CES-4.1.3 arithmetic discipline (multiplication before division, integer division truncating toward zero) and verified by the reference implementation golden_test.go. Any compliant implementation must reproduce every value.

### GS-1 Single Experiment · Full Coverage

Input: experiment exp-a (criticality 4, evidence A, recoveryBudgetMs 300000); holdMs 90000 / windowMs 100000; recoveryMs 60000; degradation active, 90/100 compliant; invariants 9/10 passed, no data loss; declared radius 10 / observed radius 8; abort not late. Coverage: 95/100.

Computation: w_e = 4×10000 = 40000; AR = 90000×10000÷100000 = 9000; REC = 10000−60000×10000÷300000 = 8000; DC = 10000×(90×10000÷100)÷10000 = 9000; DI = 9×10000÷10 = 9000; BRC = min(10000, 10×10000÷8) = 10000.

| Expected | Value |
|---|---|
| AR / REC / DC / DI / BRC | 9000 / 8000 / 9000 / 9000 / 10000 |
| Score_raw | (9000×3000+8000×2000+9000×2000+9000×2000+10000×1000)÷10000 = 8900 |
| C / k | 9500 / 10000 |
| Score_final | 8900 (display 89.00) |

### GS-2 Two Experiments · Data-Loss Cap + Partial Coverage

Input: exp-a (criticality 5, evidence B, w_e = 5×8000 = 40000; hold 60000/100000; recoveryMs = −1 no recovery event; degradation inactive 0/0; invariants 10/10, data loss true; radius 5/0; abort not late); exp-b (criticality 3, evidence C, w_e = 3×5000 = 15000; hold 100000/100000; recoveryMs 90000; degradation active 100/100; invariants 8/10; radius 10/20; abort not late). Coverage: 47/100.

Computation (Σw = 55000):
- AR = (40000×6000+15000×10000)÷55000 = 390000000÷55000 = 7090 (truncated)
- REC: only exp-b counted = 10000−90000×10000÷300000 = 7000
- DC = (40000×0+15000×10000)÷55000 = 150000000÷55000 = 2727 (truncated)
- DI = (40000×10000+15000×8000)÷55000 = 520000000÷55000 = 9454 → data-loss cap 2000
- BRC = (40000×10000+15000×5000)÷55000 = 475000000÷55000 = 8636 (truncated)

| Expected | Value |
|---|---|
| AR / REC / DC / DI / BRC | 7090 / 7000 / 2727 / 2000 / 8636 |
| Score_raw | (7090×3000+7000×2000+2727×2000+2000×2000+8636×1000)÷10000 = 5336 |
| C / k | 4700 / 8000+2000×4700÷9500 = 8989 |
| Score_final | 5336×8989÷10000 = 4796 (display 47.96) |

### GS-3 Zero-Coverage Floor

Input: GS-1 experiment records; coverage 0/100.

| Expected | Value |
|---|---|
| C / k | 0 / 8000 |
| Score_final | 8900×8000÷10000 = 7120 (display 71.20) |

## Appendix D · Error Code Directory (E1xxx, stable directory, CES-7.1.7)

| Code | Cause | Impact | Remediation |
|---|---|---|---|
| E1001 | Plugin protocol major version mismatch with core | Plugin refused to load | Upgrade plugin or core to the same protocol major version |
| E1002 | Incomplete capability declaration (missing protocol or empty faultClasses) | Plugin must not accept injection tasks | Complete capability declaration and re-handshake |
| E1003 | Preflight admission failed | Experiment verdict REJECTED, no injection | Fix per structured failure reasons and resubmit |
| E1004 | Injection residue detected after Recovering completes | Verdict invalid, implementation defect | Immediate manual intervention to clean residue and investigate recovery implementation |
| E1005 | Abort response latency exceeds declared upper bound | BRC dimension r_e set to 0 for this experiment | Investigate abort-chain latency and adjust declaration or implementation |
| E1006 | DSL parse error (outside grammar subset) | Experiment file refused to load | Fix per structured error (path + reason + clause ID) |
| E1007 | Experiment validation failed (five-tuple or clause violation) | Experiment refused to execute | Fix per clause ID |
| E1008 | Deterministic replay result mismatch | Verdict invalid, implementation defect | Investigate seed chain and RNG implementation |
| E1009 | Report content hash verification failed | Report untrusted, signatures invalidated | Re-render from original JSON source and re-sign |
| E1010 | Dimension unmeasured (UNMEASURED) | Dimension scored 0 in total | Develop drill plans to fill coverage gaps |

Error codes, once published, are never recycled; new codes follow the MINOR version process and update this directory.

## Appendix E · Equivalent Mutant Registry (effective with Chapter 9)

Registry entry format (JSON):

```json
{
  "mutId": "mut:core/score.go:142:9:M2",
  "reason": "Equivalence argument summary",
  "evidence": "Evidence locator (test output or formal argument)",
  "status": "pending | confirmed | rejected",
  "rfcId": "RFC number"
}
```

- mutId must conform to the CES-9.2.3 global identifier format; the same mutId must not be declared twice.
- Submission template: mutId + equivalence argument + supporting materials + submitter; after RFC public review and specification committee confirmation, status is set to confirmed, entering the registry and excluded from the mutation-score denominator (CES-9.3.2).
- The registry publishes a snapshot with each specification version; the snapshot content hash (CES-6.3.2) is recorded in the release notes.

---

## Suggested Fill Order
1. Chapter 4 (scoring algorithm, the soul and hardest to change, finalized first) ✅ Finalized
2. Chapter 9 (MES, dual-kernel second pillar) ✅ Finalized
3. Chapter 1 (terminology and taxonomy, foundation) ✅ Finalized
4. Chapter 3 (state machine, TLA+ modeling in parallel) ✅ Finalized (Appendix A TLA+ model verified by TLC 2.16: 6 states exhaustively checked, TypeOK + ResidueFree invariants hold, no deadlock)
5. Chapter 2 (DSL experiment model) ✅ Finalized
6. Chapter 5 (observation and verdict) ✅ Finalized
7. Chapter 6 (report format) ✅ Finalized
8. Chapter 7 (plugin and adapter protocol) ✅ Finalized
9. Chapter 8 (conformance certification) ✅ Finalized
10. Appendices A/B/C/D/E ✅ Finalized (full specification finalized)

---

Copyright (C) 2026 小黎的部落 (www.qingdujian.com).
Published under the GNU Affero General Public License v3; commercial use requires prior written authorization from 小黎's tribe.
