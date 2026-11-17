<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->

> Language / 语言：[English](07-control-plane.en.md) | [中文](07-control-plane.md)

# Abate L7 Control Plane Plan — Three Modules in Extreme Detail

**Design-target statement (mandatory for C-class translations)**: this file is the English translation of an internal planning document. All `100%` / "pass" / "met"-style statements in the full text are **Design targets / Acceptance targets**, **not measured results**; they must not be read as established facts or capability claims. The translation corresponds section by section, does not escalate the tone, and does not delete target qualifiers.

- Document ID: PLAN-07　Status: in force　Upstream: PLAN-00
- Covered modules: CTL-CLI / CTL-REST / CTL-WEB
- Common constraints (observed by all three modules):
  1. The control plane is a pure orchestration layer: it calls only the port interfaces exposed by L0–L6, and no business logic may be embedded (check rules, operator semantics, and verdict logic must never appear in L7).
  2. The control plane itself is stateless: all experiment state lives in the kernel (scheduler/archive); each CLI run is independent, and REST/Web operate through kernel session interfaces.
  3. Exit-code convention (unified across all three entry points): 0 = success; 1 = gate/experiment verdict failure (business failure); 2 = internal error; 3 = usage error (invalid arguments/configuration). Business failures are strictly distinguished from internal errors (CI consumption semantics).
  4. Coverage: statement ≥85%, branch ≥75% (ecosystem-layer standard) (design target).

---

## Module CTL-CLI: Command line

### 1. Module positioning and responsibilities

The `abate` command line: the main entry point for users and CI. Responsibilities: command parsing and dispatch, configuration loading, experiment driving (`run`/`check`/`mutate`/`chaos`), archive management (`replay`/`verify`), plugin management (`install`/`list`/`upgrade`), and report output. M1 delivers a minimal set (`run`/`check`), with milestones completing the rest progressively.

### 2. Package structure and file layout

```
ctl/cli/
  main.go          // entry: build the root command and Execute
  root.go          // root command: global flags (--config/--verbose/--no-color)
  cmd_run.go       // run: execute a CES experiment
  cmd_check.go     // check: language code checking
  cmd_mutate.go    // mutate: mutation experiment
  cmd_chaos.go     // chaos: chaos experiment (semantic shortcut for run)
  cmd_replay.go    // replay: archive playback
  cmd_verify.go    // verify: two-run comparison (determinism acceptance)
  cmd_plugin.go    // plugin install/list/upgrade/remove
  cmd_report.go    // report: generate reports from archives (format conversion)
  cmd_version.go   // version: version and build metadata
  config.go        // configuration loading (default path search + merge precedence)
  exit.go          // exit code mapping
  cli_test.go
```

### 3. Full command-tree enumeration (must be complete in the first version)

```
abate
├── run <ces-file>              # execute a CES experiment
│   --seed uint                  # determinism seed (generated and written back into the report if absent)
│   --format sarif|json|junit|console|html
│   --out path                   # report output path
│   --record path                # archive recording path
│   --budget-events/-duration/-cost  # budget overrides
│   --dry-run                    # parse + validate only, do not execute
│   --strict                     # warning violations also FAIL (error only by default)
├── check <paths...>             # language code checking
│   --lang go|python|js|c|rust|csharp|kotlin|all
│   --rules id-filter            # rule filter (comma-separated / prefix wildcard)
│   --severity error|warning|info  # threshold
│   --format sarif|json|console
│   --exit-zero                  # violations do not set FAIL (report mode)
├── mutate <paths...>            # mutation experiment
│   --suite test-selector        # test suite selector
│   --operators id-filter        # operator filter
│   --max-mutants int             # mutant cap
│   --format sarif|json|junit|html
├── chaos <ces-file>              # chaos experiment (= C-family shortcut for run)
├── replay <archive>              # archive playback
│   --verify                     # frame-by-frame verification during playback
├── verify <a> <b>               # two-run comparison (determinism acceptance)
│   --diff-out path              # difference report output
├── plugin
│   ├── install <path|ref>       # install plugin (signature check → capability review → registration)
│   ├── list                     # installed plugin list (type/version/capabilities/signature status)
│   ├── upgrade [id]             # upgrade (hot reload)
│   └── remove <id>              # remove
├── report <archive|json-report>  # generate other formats from an archive/JSON report
│   --format html|sarif|junit
├── version                      # version/commit/build metadata/watermark
└── gate                         # full gate chain (built-in equivalent of scripts/lint-gate.sh)
```

### 4. Core types and interfaces (must implement)

- Root-command global flags: `--config` (config file path; default search order: `./abate.yml` → `~/.config/abate/abate.yml` → built-in default), `--verbose` (log level), `--no-color` (TTY auto-detection in CI turns color off).
- Configuration model (`abate.yml`): plugins (install directory / additional trust-anchor paths), report (default format / output directory), budget (three default budget caps), net (offline-mode switch — offline by default, iron rule).
- Exit-code mapping: `run`/`check`/`mutate`/`chaos`/`gate` → four states 0/1/2/3; `plugin`/`report`/`replay`/`verify` → 0/2/3 (no business-failure state; `verify` differences return 1).
- `run` execution chain: load config → CES-PARSE (failure exit 3 + structured error) → CES-VALID (error violation exit 3; warning also exit 3 with `--strict`) → scheduled execution (KRN-SCH) → report output (RPT-*) → verdict-to-exit-code mapping (steady-state VIOLATED or mutation SURVIVED above threshold → exit 1).
- `verify` execution chain: load two archives → SIM-REC Compare → zero differences exit 0; if differences exist, output a structured Diff (first difference point + statistics) exit 1.

### 5. Key algorithms and flows

1. Command dispatch: cobra-style subcommand registration — no, a self-developed lightweight command tree on the standard-library `flag` package, avoiding vendoring (ADR records the decision: L7 permits third-party source inclusion, but the CLI framework is self-developed to control the dependency surface).
2. Help system: each command's `--help` outputs usage/examples/exit-code semantics; all examples must be directly copyable and executable (verified by documentation tests).
3. Progress output: progress bar for long tasks (large-scale `mutate`) (per-record output with `--verbose`; single-line refresh by default; plain-text percentage with `--no-color`).
4. Signal handling: SIGINT/SIGTERM trigger graceful termination (scheduler stops → archive flushed to disk → partial report marked `incomplete`) — direct exit without saving the archive is forbidden.

### 6. Dependencies and constraints

- Depends on: all L0–L6 port interfaces (via the use-case orchestration layer), standard library.
- Forbidden: business verdicts embedded in the CLI; network (except explicit `plugin install` network operations, refused with `--offline`); vendoring a third-party CLI framework (self-developed command tree).

### 7. Test requirements

- Unit: command-tree parsing (≥3 cases per command: valid / missing argument / invalid argument); four-state exit-code mapping; config lookup and merge precedence.
- End-to-end (CLI-level integration): `run` (valid/invalid CES cases, asserting exit codes and output); `check` (fixture project, asserting the violation list); `verify` (same-seed two-run zero difference + different-seed difference detected); `plugin install` (unsigned refused).
- Documentation tests: all example commands in `--help` output actually run successfully.
- Signal tests: after SIGINT interruption the archive is fully flushed (with the `incomplete` marker).

### 8. Acceptance criteria (quantifiable)

1. Full command-tree enumeration implemented; per-command help and example documentation tests 100% passing (acceptance target).
2. Four-state exit-code mapping cases 100% passing (correct CI consumption semantics) (acceptance target).
3. End-to-end cases (`run`/`check`/`verify`/`plugin`, ≥12 cases) 100% passing (acceptance target).
4. Graceful signal-termination case passing (archive not lost) (acceptance target).
5. Coverage ≥85%/75% (design target); all gates green (acceptance target).

---

## Module CTL-REST: REST API

### 1. Module positioning and responsibilities

Local/remote control API: experiment submission and querying, report retrieval, plugin management, and health probes. For long-lived service scenarios (Web console backend, remote orchestration system integration). Listens on localhost by default (remote listening requires explicit configuration + mandatory authentication).

### 2. Package structure and file layout

```
ctl/rest/
  server.go       // Server: route registration and lifecycle
  auth.go         // authentication (token; mandatory for remote listening)
  middleware.go   // middleware chain: logging / rate limiting / recovery / timeout
  api_experiment.go // experiment resource endpoints
  api_report.go    // report resource endpoints
  api_plugin.go    // plugin resource endpoints
  api_health.go    // health probes
  openapi.go      // OpenAPI 3.1 document generation
  rest_test.go
```

### 3. Full REST resource and endpoint definitions (must be complete in the first version)

- Experiments:
  - `POST /api/v1/experiments`: submit a CES document (body: YAML/JSON) → 202 + `Location: /experiments/{id}` (asynchronous execution).
  - `GET /api/v1/experiments/{id}`: status (pending/running/completed/failed/incomplete) + summary (score / verdict).
  - `GET /api/v1/experiments/{id}/events`: event stream (SSE, real-time progress).
  - `POST /api/v1/experiments/{id}/cancel`: cancel (graceful termination).
  - `GET /api/v1/experiments`: list (pagination / filtering: status / time window).
- Reports:
  - `GET /api/v1/experiments/{id}/report?format=sarif|json|junit|html`: report retrieval.
  - `GET /api/v1/experiments/{id}/archive`: archive download (SIM-REC binary).
- Checks:
  - `POST /api/v1/checks`: submit a check task (paths + options) → 202 asynchronous + same semantics as CLI `check`.
  - `GET /api/v1/checks/{id}`: result (violation list + summary).
- Plugins:
  - `GET /api/v1/plugins`: list (type/version/capabilities/signature status).
  - `POST /api/v1/plugins`: install (multipart upload `.wasm`+`.sig` → signature check → registration).
  - `DELETE /api/v1/plugins/{id}`: remove.
  - `POST /api/v1/plugins/{id}/upgrade`: upgrade (hot reload).
- System:
  - `GET /api/v1/health`: liveness (process responsive).
  - `GET /api/v1/ready`: readiness (kernel initialized + trust anchors loaded).
  - `GET /api/v1/version`: version / build metadata.
  - `GET /openapi.json`: OpenAPI 3.1 document (generated artifact).

### 4. Core types and interfaces (must implement)

- `Server`: `New(cfg ServerConfig) (*Server, error)`, `ListenAndServe() error`, `Shutdown(ctx) error` (graceful: drain in-flight requests → terminate running experiments → flush archives).
- `ServerConfig`: Listen (default `127.0.0.1:8420`), Auth (Token; mandatory when listening remotely on `0.0.0.0`, otherwise startup is refused), MaxConcurrent (concurrent experiment cap, default CPU count), RequestTimeout (default 30 s; long operations are always made asynchronous tasks).
- Authentication: `Authorization: Bearer <token>`; mandatory for all endpoints except `/health` `/ready` `/version` (also mandatory on local listening — protects against hostile local processes; documented explicitly).
- Task model: asynchronous execution (202 + polling/SSE); task queue cap of 100 (429 + `Retry-After` when full).
- Error model (RFC 9457 `problem+json`): type/title/status/detail/instance; error codes aligned with CLI exit-code semantics (validation→400, not-found→404, conflict→409, internal→500).

### 5. Key algorithms and flows

1. Middleware chain (fixed order): recovery (panic→500 + log) → logging (structured: method/path/status/duration/request ID) → timeout → rate limiting (token bucket: 100 req/s/IP by default) → authentication → routing.
2. SSE event stream: experiment events (state transitions/progress) forwarded by subscribing to the SIM-REC frame stream; reconnect (`Last-Event-ID` resume, with the archive Seq as cursor).
3. OpenAPI generation: declare metadata at route registration (parameters/response schemas) → aggregate at runtime to generate `openapi.json`; CI verifies zero drift between the generated artifact and routes.
4. Concurrency safety: experiment tasks isolated from each other (separate scheduler instances); plugin operations serialized (single write lock — atomic hot reload).

### 6. Dependencies and constraints

- Depends on: L0–L6 port interfaces, standard library `net/http` (Go 1.22+ routing enhancements, no third-party routing framework — recorded in an ADR).
- Forbidden: embedded business logic; remote listening without authentication (startup refused); plaintext tokens in logs (masking).

### 7. Test requirements

- Unit: full endpoint table (≥2 cases per endpoint: valid/invalid); authentication (no token/wrong token → 401); rate limiting (over limit → 429); error model (`problem+json` structure assertions).
- Integration: `httptest` full chain (submit experiment → poll to completion → fetch report → fetch archive); SSE (reconnect and resume); cancel (cancel→incomplete state + complete archive).
- Security: remote-listening-without-token startup refusal case; token log-masking assertions.
- OpenAPI: generated-artifact schema validation + zero-drift CI check against routes.

### 8. Acceptance criteria (quantifiable)

1. All endpoints (5 experiments + 2 reports + 2 checks + 4 plugins + 4 system = 17 endpoints) with cases 100% passing (acceptance target).
2. OpenAPI 3.1 generated artifact passes the official validator; zero-drift CI check against routes passing (acceptance target).
3. SSE reconnect-and-resume case passing; graceful-cancel case passing (complete archive) (acceptance target).
4. Security cases (mandatory authentication / masking / remote refusal) 100% passing (acceptance target).
5. Coverage ≥85%/75% (design target); all gates green (acceptance target).

---

## Module CTL-WEB: Web console

### 1. Module positioning and responsibilities

A browser console: experiment submission and monitoring, report browsing, interactive mutation matrix, and plugin management. A pure static frontend (SPA) + CTL-REST backend; frontend assets are statically served by the `abate web` command (same origin, no CORS surface).

### 2. Package structure and file layout

```
ctl/web/
  serve.go        // abate web command: start REST + static serving
  embed.go        // go:embed frontend build artifacts (dist/)
  dist/           // frontend build artifacts (archived to guarantee offline availability)
  web_test.go
frontend/         // frontend sources (independent subproject; build output goes to ctl/web/dist/)
  src/
    pages/        # page components
    api/          # REST client (OpenAPI-generated)
    components/   # shared components
```

### 3. Page inventory (must be complete in the first version)

1. Experiment list page: status filter/pagination; inline summary (score / verdict / duration).
2. Experiment detail page: HPOV state timeline (state-transition visualization) + real-time event stream (SSE) + cancel button.
3. Experiment submission page: CES document editor (syntax highlighting / inline validation-error markers — reusing CES-PARSE error positions) + seed/budget form.
4. Report page: violation list (filter/sort / source snippet ±3 lines) + mutation matrix (rows = mutants, columns = tests, kill markers) + score gauge.
5. Plugin page: list (type/version/capabilities/signature status) + install (upload `.wasm`+`.sig`) + upgrade/remove.
6. Settings page: service address / token configuration (stored locally in the browser).

### 4. Core types and interfaces (must implement)

- `abate web` command: `--listen` (default `127.0.0.1:8430`) + `--api` (REST address, defaults to an in-process embedded start); static serving (`go:embed dist`; SPA route fallback to `index.html`).
- REST client: OpenAPI-generated (type-safe; generated artifact archived; CI verifies zero drift against `openapi.json`).
- Frontend technical constraints: no runtime framework CDN dependency (bundled at build time, self-contained artifact); build toolchain version locked (`pnpm-lock` archived).

### 5. Key algorithms and flows

1. Same-origin serving: the `web` command embeds REST by default or reverse-proxies a specified API (`--api`); same origin eliminates the CORS configuration surface.
2. SSE consumption: the experiment detail page subscribes to the event stream; automatic reconnect (`Last-Event-ID`) + status indication.
3. Report rendering: SARIF/JSON data-driven rendering; virtual scrolling for large mutation matrices (>1000 rows without lag) (design target).

### 6. Dependencies and constraints

- Depends on: CTL-REST (API), frontend build toolchain (build time only, not in runtime).
- Forbidden: external runtime CDN/fonts/scripts (fully self-contained); business logic embedded in the frontend (pure display and orchestration — all verdicts come from the API).

### 7. Test requirements

- Frontend unit tests: page components (≥3 cases per page: data rendering / empty state / error state); API-client generated-artifact type checks.
- E2E: Playwright (source-level inclusion in the toolchain, dev dependency) across all six pages' full flows (submit → monitor → report → plugin management).
- Build: `dist` artifact archived and builds reproducible (two builds byte-identical — frontend determinism) (acceptance target).

### 8. Acceptance criteria (quantifiable)

1. Six-page component unit tests 100% passing; E2E full flow passing (acceptance target).
2. API-client generated artifact and `openapi.json` with zero drift (CI check) (acceptance target).
3. `dist` build reproducible (two builds byte-identical); artifact self-contained (no external requests — network-panel assertion) (acceptance target).
4. Large-matrix virtual-scrolling performance: 10,000 mutants rendered interactively P99 <100 ms (E2E performance assertion) (design target).
5. Coverage ≥85%/75% (Go side) (design target); all gates green (acceptance target).

---

## Appendix: Overall L7 merge acceptance (after all three modules are in place)

1. End-to-end: CLI submits an experiment → REST queries status → Web monitors → reports agree across all three ends (asserting CLI output / REST report / Web rendering are from the same source data) (acceptance target).
2. Semantic consistency across exit codes / API error codes / OpenAPI (mapping table archived, CI-verified) (acceptance target).
3. Offline verification: with the network disconnected, all CLI commands and all Web pages work (zero-external-dependency iron rule) (acceptance target).
4. An aggregate coverage report for all three modules is archived; the five-section gate chain + baseline regression are all green (acceptance target).
