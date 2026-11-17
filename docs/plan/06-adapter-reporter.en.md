<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->

> Language / 语言：[English](06-adapter-reporter.en.md) | [中文](06-adapter-reporter.md)

# Abate L5/L6 Adapter and Reporter Plan — Eight Modules in Extreme Detail

**Design-target statement (mandatory for C-class translations)**: this file is the English translation of an internal planning document. All `100%` / "pass" / "met"-style statements in the full text are **Design targets / Acceptance targets**, **not measured results**; they must not be read as established facts or capability claims. The translation corresponds section by section, does not escalate the tone, and does not delete target qualifiers.

- Document ID: PLAN-06　Status: in force　Upstream: PLAN-00
- Covered modules: ADP-CI / ADP-ITS / ADP-MSG / RPT-SARIF / RPT-HTML / RPT-JSON / RPT-JUNIT / RPT-OTEL
- Common constraints (observed by all eight modules):
  1. All L5/L6 components are delivered as WASM plugins (WIT interfaces `adapter`/`reporter`), with official built-in implementations also provided (native Go, distributed as default plugins alongside the engine) — built-in implementations implement the WIT interface semantics as well, so plugins can replace built-ins (ecosystem openness iron rule) (design target).
  2. Adapters and reporters are pure consumers: they read experiment results and violation data only, and are forbidden from influencing kernel execution in the reverse direction.
  3. Output determinism: the same input (experiment report / violation list) produces byte-identical output (except HTML generation timestamps — the timestamp is placed in the footer and can be disabled, and is forcibly disabled in CI mode) (design target).
  4. Coverage: statement ≥85%, branch ≥75% (ecosystem-layer standard) (design target).
  5. Network boundary: external network calls by adapters (CI platform APIs, issue systems, messaging channels) are allowed only inside adapters; reporters have zero network use except RPT-OTEL.

---

## Module ADP-CI: CI adapter

### 1. Module positioning and responsibilities

Connects Abate gates to CI platforms: GitHub Actions and GitLab CI in the first version. Responsibilities: drive `abate check` and `abate run` (experiments) in CI pipelines, and return exit codes and result artifacts (SARIF/JSON reports) to the platform (PR comments, check runs, artifact uploads).

### 2. Package structure and file layout

```
adapter/ci/
  ci.go           // CIAdapter interface and registry
  github/        // GitHub Actions adaptation
    actions.go   // workflow action entry (composite action)
    api.go       // GitHub API client (check runs / PR comments / artifact uploads)
    sarif_upload.go // code scanning SARIF upload
  gitlab/
    gitlab.go    // GitLab CI job entry
    api.go       // GitLab API client (MR discussions / artifacts)
  config.go      // adapter configuration (platform detection / token source / artifact naming)
  ci_test.go
```

### 3. Core types and interfaces (must implement)

- `CIAdapter` interface:
  - `Detect(env Env) (Platform, bool)`: platform detection (sniffing `GITHUB_ACTIONS`/`GITLAB_CI` environment variables).
  - `ReportCheck(ctx, result CheckResult) error`: return check results (GitHub: check run + conclusion; GitLab: MR report + conclusion).
  - `UploadArtifact(ctx, files []string) error`: upload artifacts (SARIF/JSON report archives).
  - `Comment(ctx, summary Summary) error`: PR/MR summary comment (violation statistics / survival score / gate verdict).
- `CheckResult`: ExitCode (0/1/2 three-state semantics aligned with LNG-CHK) / ViolationSummary / MutationScore / GateVerdict.
- `Summary`: Markdown text (templated: verdict badge + severity statistics table + top violating rules + report artifact links).
- Configuration: tokens injected from environment variables (`GITHUB_TOKEN`/`GITLAB_TOKEN`); storing tokens in config files is forbidden.

### 4. Key algorithms and flows

1. Platform detection: environment-variable sniffing (deterministic: fixed detection order GitHub → GitLab; when both platforms are present, explicit configuration takes precedence).
2. Summary generation: template rendering (Go `text/template`; data is structured `Summary`, and string-concatenated HTML is forbidden — injection prevention).
3. Failure semantics: on gate FAIL (exit 1) the CI job FAILs and the comment lists the violated rules; exit 2 (internal error) is strictly distinguished from FAIL (the comment marks `internal error`, not a misleading code violation).
4. SARIF upload: GitHub code scanning API (SARIF 2020-21 format, produced by RPT-SARIF).

### 5. Dependencies and constraints

- Depends on: RPT-SARIF/RPT-JSON (artifacts), LNG-CHK (exit-code semantics).
- Network boundary: GitHub/GitLab API calls only in this module; tokens via environment variables only.
- Forbidden: executing check logic inside the adapter (pure return layer); tokens written to disk or into logs (masking gate).

### 6. Test requirements

- Unit: platform detection matrix (with/without environment variables, both platforms present); summary template rendering (including special-character escaping cases).
- Integration: real workflow runs on GitHub Actions (test repository, PR trigger → check run + comment + artifact assertions); likewise for GitLab CI.
- Failure semantics: comment-content assertions for both exit 1 and exit 2 paths.

### 7. Acceptance criteria (quantifiable)

1. Platform detection matrix cases 100% passing; both-platform explicit-configuration case passing (acceptance target).
2. GitHub Actions and GitLab CI real-run integration tests passing (three assertions: check run / comment / artifact) (acceptance target).
3. Summary template special-character escaping cases passing (XSS/injection protection) (acceptance target).
4. Coverage ≥85%/75% (design target); all gates green (acceptance target).

---

## Module ADP-ITS: Issue-tracking adapter

### 1. Module positioning and responsibilities

Connects check violations and mutation results to issue trackers (GitHub Issues in the first version): automatically opens issues for violations (deduplicated), opens improvement issues for SURVIVED mutants, and opens review issues for equivalent mutants (design target).

### 2. Package structure and file layout

```
adapter/its/
  its.go          // ITSAdapter interface
  github/
    issues.go     // GitHub Issues client (create / query / close)
    dedupe.go     // deduplication: fingerprint matching against existing issues
  fingerprint.go  // violation fingerprint (rule + file + function-level normalization)
  its_test.go
```

### 3. Core types and interfaces (must implement)

- `ITSAdapter` interface: `CreateIssue(ctx, issue Issue) (IssueRef, error)`, `FindOpen(ctx, fp Fingerprint) (IssueRef, bool)`, `CloseIssue(ctx, ref IssueRef, reason string) error`.
- `Issue`: Title/Body (Markdown template)/Labels (rule category/severity/language)/Fingerprint.
- `Fingerprint`: SHA-256(rule id + file path + function name + normalized violation summary) — line numbers do not participate (line drift does not create duplicate issues).
- Synchronization policy: full reconciliation — on every run: create (new fingerprint) / close (fixed: issue open last run with no corresponding violation this run) / keep (still present).

### 4. Key algorithms and flows

1. Fingerprint normalization: violation location normalized to function level (AIR NodeID locates the function); multiple violations of the same rule in the same function merge into one issue (all locations listed in the body).
2. Reconciliation flow: fetch all open issues (label filter) → compare with this run's fingerprint set → three-way operations (create / close / keep).
3. Rate limiting: API token bucket (GitHub 5000/h); reconciliation report outputs operation statistics.

### 5. Dependencies and constraints

- Depends on: LNG-CHK (violation data), KRN-MUT (mutation verdicts).
- Network boundary: GitHub API only in this module.
- Forbidden: issue bodies containing full source code (only the violating snippet ±3 lines of context).

### 6. Test requirements

- Unit: fingerprint normalization (same fingerprint on line drift / different fingerprint on function rename); three-way reconciliation logic.
- Integration: real-run on a test repository (create 3 issues → fix 1 → rerun → assert 1 closed and 2 kept).

### 7. Acceptance criteria (quantifiable)

1. Fingerprint normalization cases (line drift / function rename / same-function merge) 100% passing (acceptance target).
2. Three-way reconciliation integration test passing (operation statistics assertions) (acceptance target).
3. Coverage ≥85%/75% (design target); all gates green (acceptance target).

---

## Module ADP-MSG: Messaging adapter

### 1. Module positioning and responsibilities

Messaging-channel notifications for experiment and gate results (first version: generic Webhook + Feishu + WeCom): three event types — gate FAIL, experiment completion, and score below threshold.

### 2. Package structure and file layout

```
adapter/msg/
  msg.go         // MsgAdapter interface and event model
  webhook/       // generic Webhook (POST JSON)
  feishu/        // Feishu bot
  wecom/         // WeCom bot
  template.go    // message templates (Markdown / cards)
  msg_test.go
```

### 3. Core types and interfaces (must implement)

- `MsgAdapter` interface: `Send(ctx, event Event) error`.
- `Event`: Kind (GATE_FAIL/EXPERIMENT_DONE/SCORE_BELOW) / Payload (structured summary) / Level (info/warn/critical).
- Configuration: channel list (type + endpoint + secret); secrets only via environment variables / controlled configuration, with log masking.
- Threshold: SCORE_BELOW carries a threshold (e.g. trigger when survival score < 0.7).

### 4. Key algorithms and flows

1. Event routing: event → subscribed channel list (configuration-driven) → parallel sending (one channel's failure does not block others; errors aggregated and reported).
2. Retry: retry 3 times on 5xx/timeout (exponential backoff 1 s/2 s/4 s, real time); no retry on 4xx (configuration error).
3. Templates: a template per channel (Webhook structured JSON; Feishu/WeCom Markdown cards); template-variable allowlist (arbitrary HTML injection forbidden).

### 5. Dependencies and constraints

- Depends on: KRN-VER (score), LNG-CHK (gate result).
- Network boundary: the three channel endpoints only in this module.
- Forbidden: messages containing source-code snippets (statistics and links only).

### 6. Test requirements

- Unit: event routing matrix; retry policy (5xx retried / 4xx not retried); template-variable allowlist.
- Integration: local mock server (Webhook) real-run payload assertions; Feishu/WeCom actual sends via official test bots (manual acceptance record).

### 7. Acceptance criteria (quantifiable)

1. Routing matrix and retry cases 100% passing (acceptance target).
2. Webhook mock integration payload-structure assertions passing (acceptance target).
3. Coverage ≥85%/75% (design target); all gates green (acceptance target).

---

## Module RPT-SARIF: SARIF reporter

### 1. Module positioning and responsibilities

SARIF 2.1.02 standard output of check violations and mutation results: a standard format consumed by GitHub code scanning, IDEs, and CI platforms.

### 2. Package structure and file layout

```
report/sarif/
  sarif.go       // SARIF document model (2.1.02 subset)
  build.go       // violations / mutation results → SARIF build
  validate.go    // self-validation (schema + semantics)
  sarif_test.go
```

### 3. Core types and interfaces (must implement)

- `Build(violations []Violation, mutations []MutationResult) (*SarifDoc, error)`: main entry point.
- `SarifDoc` model (mandatory 2.1.02 subset): version/$schema/runs[] (`tool.driver{name,version,rules[]}`, `results[]{ruleId,level,message,locations[]}`, `artifacts[]`, `columnKind: unicodeCodePoints`).
- Rule metadata: every rule goes into `rules[]` (id/shortDescription/helpUri — pointing at the official-specification source).
- Mutation-result mapping: SURVIVED → level `note` + properties area (score details); KILLED is not included in `results` (summary goes into the properties area).

### 4. Key algorithms and flows

1. Build pipeline: violation grouping (by rule) → `rules[]` generation → `results[]` generation (location: file URI + line/column, 1-based, `unicodeCodePoints`) → `artifacts` registration.
2. Self-validation: built-in schema validation after build (JSON Schema draft 2020-12) + semantic validation (`ruleId` must exist in `rules[]`; locations must not be out of bounds) — validation failure returns an error rather than producing a bad document.
3. Determinism: the same input produces byte-identical output (fixed key order: struct order, not map order).

### 5. Dependencies and constraints

- Depends on: LNG-CHK/KRN-MUT (data types), standard library `encoding/json`.
- Forbidden: network; non-standard extension fields (to preserve maximum compatibility).

### 6. Test requirements

- Unit: build pipeline (three scenarios: violations / mutations / mixed); determinism (two runs byte-identical).
- Compliance: output passes the official SARIF JSON Schema + the Microsoft sarif-sdk validator (double validation) (acceptance target).
- Boundaries: empty violations (valid empty document), overlong messages, Unicode filenames/paths.

### 7. Acceptance criteria (quantifiable)

1. Both official validators pass (report to be archived) (acceptance target).
2. Two-run byte-identical output; empty/boundary cases passing (acceptance target).
3. Coverage ≥85%/75% (design target); all gates green (acceptance target).

---

## Module RPT-HTML: HTML reporter

### 1. Module positioning and responsibilities

A human-readable interactive report: violation list (filterable/sortable), mutation matrix (per-mutant verdict), experiment timeline (HPOV state transitions), and survival-score gauge. Single self-contained file (inline CSS/JS, zero external dependencies, openable offline).

### 2. Package structure and file layout

```
report/html/
  html.go        // HTMLDoc build main entry
  layout.go      // page skeleton (template)
  section_violation.go // violation section
  section_mutation.go  // mutation matrix section
  section_timeline.go  // experiment timeline section
  section_score.go     // score gauge section
  assets/       // inline assets (CSS/JS sources)
  html_test.go
```

### 3. Core types and interfaces (must implement)

- `Build(report FullReport, opts HTMLOptions) ([]byte, error)`: main entry; FullReport = violations + mutations + experiment archive metadata + scores.
- `HTMLOptions`: EmbedSource (violation snippet embedding switch, default on ±3 lines) / Timestamp (footer timestamp, forcibly off in CI mode).
- Interactive capabilities (inline JS, no framework dependency): violation table filtering (language/severity/rule) and sorting; mutation matrix expansion (per mutant: location/operator/verdict/killing test); timeline collapse/expand.
- Security: all dynamic content HTML-escaped (`html/template` automatic escaping; `text/template` forbidden for rendering user data); source-snippet syntax highlighting is pure CSS-class annotation (no eval).

### 4. Key algorithms and flows

1. Rendering pipeline: data model → template tree (`html/template`, compiled once) → byte output; all templates auto-escape (XSS protection).
2. Mutation matrix: rows = mutants, columns = tests (verdict matrix with KILLED marking the killing test); pagination for large data (>1000 mutants rendered in pages of 100) (design target).
3. Determinism: byte-identical except for the footer timestamp; in CI mode (Timestamp off) two runs are byte-identical (acceptance target).

### 5. Dependencies and constraints

- Depends on: LNG-CHK/KRN-MUT/KRN-VER/SIM-REC (archive metadata).
- Forbidden: external CDN/fonts/scripts (fully self-contained); network.

### 6. Test requirements

- Unit: rendering pipeline (each of the four sections' data model → output assertions); escaping cases (script-injection payload escaping assertions).
- Determinism: two-run byte-identical in CI mode.
- Scale: 100,000 violations / 10,000 mutants rendering P99 ≤5 s (benchmark archived) (design target).

### 7. Acceptance criteria (quantifiable)

1. Four-section rendering and escaping cases 100% passing (acceptance target).
2. Two-run consistency in CI mode (acceptance target).
3. Scale performance met (design target); coverage ≥85%/75% (design target); all gates green (acceptance target).

---

## Module RPT-JSON: JSON reporter

### 1. Module positioning and responsibilities

A machine-readable canonical report: a stable, versioned JSON Schema for the full result set (violations / mutations / experiment metadata / scores / archive references), for downstream tools and CI consumption.

### 2. Package structure and file layout

```
report/json/
  json.go        // Build main entry
  schema.go      // JSON Schema definition and version
  json_test.go
```

### 3. Core types and interfaces (must implement)

- `Build(report FullReport) ([]byte, error)`.
- Document structure: schemaVersion (semver) / generated (deterministic metadata: seed / fingerprint / tool version) / violations[] / mutations[] / experiments[] / scores / summary.
- Schema: publish `report-schema-v1.json` with each version; field deprecation goes through two minor windows (aligned with the kernel API policy).
- Determinism: same input produces byte-identical output (struct order, no map order; timestamps only in `generated.stable` metadata, never real time).

### 4. Key algorithms and flows

1. Build pipeline: FullReport → structured model → `json.Marshal` (fixed struct field order).
2. Self-validation: built-in schema validation after build (prevent model/schema drift; CI drift check).

### 5. Dependencies and constraints

- Depends on: same data types as RPT-HTML. Forbidden: network; schema drift (enforced in CI).

### 6. Test requirements

- Unit: build pipeline; schema self-validation; two-run determinism.
- Compatibility: v1.x old documents readable by new code (backward-compatibility cases).

### 7. Acceptance criteria (quantifiable)

1. Schema self-validation and CI drift check with zero differences (acceptance target).
2. Two-run byte-identical output; backward-compatibility cases passing (acceptance target).
3. Coverage ≥85%/75% (design target); all gates green (acceptance target).

---

## Module RPT-JUNIT: JUnit reporter

### 1. Module positioning and responsibilities

JUnit XML output (the standard consumed by CI platform test reports): maps mutant verdicts to test cases (SURVIVED = failed case, KILLED = passed case) and gate verdicts to a testsuite.

### 2. Package structure and file layout

```
report/junit/
  junit.go       // Build main entry
  junit_test.go
```

### 3. Core types and interfaces (must implement)

- `Build(report MutationReport) ([]byte, error)`.
- Mapping: `testsuite name=abate-mutation`, `tests` = total mutants, `failures` = SURVIVED count, `skipped` = EQUIVALENT+TIMEOUT count; one `testcase` per mutant (`name=operator:location`; SURVIVED generates a `failure` node with a message and a killing-attempt summary).
- Follows common JUnit XML dialects (the subset consumed by GitHub/GitLab: testsuite/testcase/failure/skipped attribute sets).

### 4. Key algorithms and flows

1. Mapping pipeline: mutation results → testcase list (deterministic sort: location → operator) → XML serialization (`encoding/xml`, correct escaping).
2. Determinism: same input produces byte-identical output.

### 5. Dependencies and constraints

- Depends on: KRN-MUT data types. Forbidden: network; dialect-specific extension attributes (to preserve general compatibility).

### 6. Test requirements

- Unit: mapping (cases for all four verdict classes); XML escaping (messages containing `<>&"`).
- Consumption verification: real-run parsing by GitHub Actions test-report annotations and GitLab unit-test reports passing.

### 7. Acceptance criteria (quantifiable)

1. Four-verdict mapping and escaping cases 100% passing (acceptance target).
2. Both platforms' real-run parsing passing (record archived) (acceptance target).
3. Two-run consistency (acceptance target); coverage ≥85%/75% (design target); all gates green (acceptance target).

---

## Module RPT-OTEL: OTel reporter

### 1. Module positioning and responsibilities

OpenTelemetry export: execution traces (HPOV state spans), observation metrics (probe-value metrics), and verdict events (logs), for consumption by observability platforms (Jaeger/Prometheus/Grafana).

### 2. Package structure and file layout

```
report/otel/
  otel.go        // Exporter construction (OTLP exporter wrapper)
  trace.go       // HPOV states → span tree
  metrics.go     // observation stream → metric series
  logs.go        // verdicts / violations → log records
  otel_test.go
```

### 3. Core types and interfaces (must implement)

- `NewExporter(ctx, cfg OTLPConfig) (*Exporter, error)`: OTLP gRPC/HTTP exporter (endpoint/protocol/headers/batch parameters).
- `Exporter.Export(ctx, report FullReport, archive *sim/record.Archive) error`: three-stream export.
- Trace mapping: experiment root span (experiment fingerprint as a trace attribute) → HPOV state spans (state name / virtual-time interval) → perturbation spans (operator id / parameter hash) → probe spans.
- Metric mapping: probe observations → gauge (virtual time as a timeline attribute); survival score → gauge.
- Log mapping: verdict events (VERDICT) and violations (VIOLATION) → log records (severity mapping: VIOLATED→WARN, SURVIVED→ERROR).

### 4. Key algorithms and flows

1. Batched export: three streams (span/metric/log) batched independently (default 512 records/batch or a 5 s real-time flush); after 3 failed retries, export is dropped and counted (does not block the experiment).
2. Virtual-time semantics: OTel timestamps are real export times; virtual time is carried as span/log attributes (`virtualTime`) — preserving fidelity without conflating the two.

### 5. Dependencies and constraints

- Depends on: OTel Go SDK brought in at source level (`third_party/otel`, Apache-2.0, five-step inclusion process: strip the collector, keep only the SDK export path).
- Network boundary: the OTLP endpoint only in this module; export failures do not affect experiment results (bystander semantics).
- Forbidden: export blocking the main experiment flow (async queue capped at 100,000 records; drop and count when full).

### 6. Test requirements

- Unit: three-stream mapping (≥3 cases each); batching and retry policy; queue-full drop counting.
- Integration: local OTLP collector (out-of-container process, test fixture) receiving assertions (span tree structure / metric values / log records).
- Bystander semantics: experiment completes normally when the collector is unreachable (assert results unaffected).

### 7. Acceptance criteria (quantifiable)

1. Three-stream mapping and batching/retry cases 100% passing (acceptance target).
2. Local collector integration assertions passing (structural validation) (acceptance target).
3. Bystander-semantics case passing (unreachability does not affect the experiment) (acceptance target).
4. Coverage ≥85%/75% (design target); all gates green (acceptance target).

---

## Appendix: Overall L5/L6 merge acceptance (after all eight modules are in place)

1. End-to-end: one complete experiment (check + mutation + chaos) produces five report formats (SARIF/HTML/JSON/JUnit/OTel) plus CI return-path (real GitHub Actions run: check run + comment + artifact + code scanning) in a full-chain demonstration (acceptance target).
2. Standard compliance: SARIF passes both official validators; JUnit passes both platforms' parsers; JSON schema has zero drift (acceptance target).
3. Determinism: SARIF/JSON/JUnit two runs byte-identical; HTML two runs identical in CI mode (acceptance target).
4. An aggregate coverage report for all eight modules is archived; the five-section gate chain + baseline regression are all green (acceptance target).
