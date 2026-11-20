<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->

# Abate Architecture Overview

> This document describes the actual structure of the current code. Every component has source in
> the repo; claims are reproducible with `scripts/evidence-check.sh`.

## 1 Positioning and shape

Abate is a **single static binary** containing two cooperating subsystems:

1. **Deterministic experiment control plane**: describes experiments in CES (`ces.tencent.com/v1`),
   validates, schedules deterministically, injects faults, simulates, and emits byte-reproducible
   reports plus SIM-REC archives.
2. **Seven-language structured static analyzer**: parses source into a unified AIR syntax tree,
   matches structural patterns, emits console / JSON / SARIF reports.

Foundation: **pure Go standard library, in-house, zero external dependencies** (no `require` in
`go.mod`, no `vendor/`, no `third_party/`); a single binary cross-compiled for five platforms.

## 2 Layered structure

Code is layered by responsibility with one-way dependencies (upper calls lower, lower never calls back):

| Layer | Directories | Responsibility |
|---|---|---|
| Entry / control | `main.go`, `cmd/`, `ctl/cli/`, `ctl/rest/`, `ctl/web/` | CLI commands, REST API, browser console |
| CES frontend | `ces/schema/`, `ces/parse/`, `ces/validate/`, `ces/tck/` | CES parse, normalize, validate (CES-VALID / PARSE rules), conformance suite |
| Language frontends | `lang/air/`, `lang/frontkit/`, `lang/{go,py,js,c,cs,rs,kt}front/`, `lang/rule/` | 7-language lexer/parser → AIR AST → structural rule matching; `check` / `derive` in `lang/check/`, `lang/derive/` |
| Kernel | `kernel/algebra/`, `kernel/scheduler/`, `kernel/injection/`, `kernel/mutation/`, `kernel/observe/`, `kernel/verify/`, `kernel/abi/` | Fault algebra, deterministic scheduling, injection operators, mutation operators, observation, steady-state verdicts, ABI data |
| Simulator | `sim/clock/`, `sim/fs/`, `sim/net/`, `sim/rand/`, `sim/record/`, `sim/vsched/` | Virtual clock/fs/network (delay·partition·packet loss·bandwidth)/deterministic RNG/archive/scheduling |
| Security | `security/policy/`, `security/redact/`, `security/audit/`, `security/tests/` | Capability policy, redaction, security scanning |
| Plugins | `plugin/host/`, `plugin/signature/`, `plugin/capability/`, `plugin/sandbox/`, `plugin/{adapter,reporter,sample,wasm,hotswap}/` | Ed25519 signature verification, capability model, sandbox config, adapter/reporter/sample, self-written WASM host, hot-swap |
| Support | `internal/determinism/`, `internal/e2e/`, `internal/corpus/`, `internal/watermark/`, `internal/yamlmini/`, `ctl/rest/` | Determinism tools, E2E framework, corpus, watermark, in-house minimal YAML parser, REST |

## 3 Core pipelines

### 3.1 `check`: source → AIR → findings

```
<paths> → dispatch by extension → <lang>front (lexer → Pratt parser) → air.Node AST
        → rule matcher (structural patterns) → findings → console/json/sarif
```

- **AIR** (`lang/air/`): unified AST node types (kind / SemKey / attributes / children), schema versioned.
- **Rules** (`lang/rule/builtin/`): 72 shared AIR templates + per-language rules (Go 116, others 12–14),
  using compiled structural patterns (`$x` capture, `$..` descendant, `!P` negation, attribute
  constraints) matching AST nodes — not text regex.
- **Determinism**: identical input yields byte-identical output; `summary.fingerprint` supports `verify`.
- **Tolerance**: broken input yields partial AIR + warnings; not a syntax gate.

### 3.2 `run`: CES → experiment → report

```
CES YAML → parse → normalize → validate (CES-VALID/PARSE)
        → kernel scheduler (fixed seed) + injection + sim → observe → verify (steady state)
        → report + optional SIM-REC archive
```

- **Fault algebra** (`kernel/algebra/`): associativity, identity, inverse-compensation laws.
- **Determinism**: fixed seed + virtual clock + total-order events; reports and archives byte-reproducible.
- **SIM-REC archive** (`sim/record/`): hash chain + total-order sequence + virtual-time monotonicity, verifiable by `replay`.

### 3.3 `derive`: rules → mutation operators

Derives M-family mutation operators from built-in Go rules, registers them with `kernel/mutation/`,
and runs one deterministic mutation experiment. Killing is a **static AIR-level approximation**,
not a real go test execution.

## 4 Plugins and security

- **Signature**: `plugin/signature/` verifies the manifest public key with Ed25519; unsigned/tampered plugins are rejected.
- **Capability sandbox**: `plugin/capability/` + `plugin/sandbox/` model runsc config semantics
  (capability switches, read-only rootfs, virtual clock), default-deny and purely declarative; real
  isolation is the host's responsibility.
- **WASM host**: `plugin/wasm/` is a **self-written** WAT lexer/parser/encoder and runtime, zero third-party dependencies.
- **Hot-swap**: `plugin/hotswap/` atomically replaces instances without interrupting in-flight calls.

## 5 Determinism system

Determinism is a cross-cutting constraint, with reproducible gates:

| Mechanism | Implementation / script |
|---|---|
| Fixed seed / virtual clock | `kernel/scheduler/`, `sim/clock/`, `sim/rand/` (SplitMix64) |
| Byte-identical reports | `abate verify`, `summary.fingerprint`, run digest / stream digest |
| Archive integrity | `sim/record/` hash chain, `abate replay` |
| Determinism gate | `scripts/determinism-check.sh` (same-seed double-run compare of core packages) |
| Spec consistency | `scripts/drift-check.sh` |
| Formal verification | `scripts/tla/` (TLA+: PitParity / SandboxFuel / StreamDeterminism) |
| Integrity watermark | `internal/watermark/`, `scripts/hash-check.sh`, `scripts/watermark-verify.sh` |

## 6 Quality gates

`scripts/lint-gate.sh` chains: copyright header → style → license → gofmt → go vet → golangci-lint →
magic constants → `go test -race` → coverage → build. Zero-tolerance items include TODO/FIXME
(godox), unformatted code, ignored errors, hardcoded secrets, excessive complexity. Companion:
`coverage-gate.sh`, `vuln-scan.sh`, `m5-reference-scan.sh`, `audit-zero-third-party.sh`, `fullcorpus-degrade.sh`.

## 7 Boundaries and honest scope

- `check` is a structural linter, not compiler-grade semantic analysis (see [Capabilities and limits](user/capabilities-and-limits.en.md)).
- Full-corpus fidelity is 3%–5% for six languages, Go `NOT_MEASURED`; E2E "passing" applies only to
  the frozen single-file scope — not a "high fidelity" claim.
- Competitor absolute scores `NOT_MEASURED`; derive killing is a static approximation; community
  validation is still empty.

## 8 Map

- CLI and exit codes: [CLI reference](user/cli-reference.en.md)
- CES spec: [CES v1 spec](spec/ces-v1.en.md)
- Development: [Contributing](dev/contributing.en.md) · [Module dev](dev/module-dev-guide.en.md) · [Plugin dev](dev/plugin-dev-guide.en.md) · [Sandbox](dev/sandbox.en.md)
