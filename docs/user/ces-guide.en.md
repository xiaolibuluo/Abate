<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->

> Language / 语言：[English](ces-guide.en.md) | [中文](ces-guide.md)

# CES Authoring Guide (User CES Guide)

> Companion specification: `docs/spec/ces-v1.md`. All examples can be
> validated with `abate run --dry-run`.

## Envelope rules

- `apiVersion` must be `ces.tencent.com/v1` and `kind` must be `Experiment`.
- Unknown fields are **rejected outright** (fail-fast); do not rely on
  defaults to paper over typos.
- `metadata.name` is a DNS subdomain; `namespace` defaults to `default`.

## Target

| kind | Meaning | Status |
|---|---|---|
| `simulator` | Target inside the simulator (default, deterministic) | ✅ supported |
| `source` | Source-level target (language frontend) | ✅ supported |
| `external` | External system | ❌ reserved, rejected by validation |

`selector` is a set of stable key/value pairs: `node` / `process` / `package` /
`function`.

## Steady state

At least one hypothesis is required; each entry has the structure:

```yaml
- probeRef: p1                       # references probes[].id
  expectation: { type: value, value: ok }   # value | range | predicate
  tolerance:   { type: absolute, value: 1 } # absolute | relative | percent
  holdFor: 2s                        # virtual duration, must be positive
```

`expectation.type` values: `value` (uses `value`), `range` (uses `min`/`max`),
`predicate` (uses `predicate`). `tolerance.type` values: `absolute`/
`relative`/`percent`, with a non-negative `value`.

## Method

Perturbation algebra (see specification §3 for details): `∘` composition,
`⊗` concurrent product, `a⁻¹` dual, `ε` identity, parentheses for grouping.
Nesting depth limit is **64**. Non-invertible operators must declare a
`compensation`.

```yaml
method:
  expression: "kill ∘ restart⁻¹"
  operators:
    - id: kill
      family: chaos
      targets: [n1]
      invertible: false
      compensation: restart
```

## Probes

`id` is unique within the document; `kind` ∈ {`http`,`shell`,`metric`,`prom`},
each with its own fields:

```yaml
- id: p1
  kind: metric
  spec: { query: up, expect: "1" }
  interval: 1s
- id: p2
  kind: http
  spec: { url: http://svc/health, method: GET, expectStatus: 200 }
  interval: 5s
```

The sampling `interval` must be positive; `shell` commands are validated
against a security whitelist (read-only/safe commands only).

## Rollback

`strategy` ∈ {`auto`,`manual`,`timeout`}; the `timeout` strategy must also
fill in `timeout`. `verifySteadyState: true` re-runs the steady-state
assertions after the dual operator is applied.

## Schedule

One of three (mutually exclusive):

```yaml
schedule: { continuous: { duration: 30s } }   # continuous perturbation
schedule: { once: { at: 2026-01-01T00:00:00Z } }  # one-shot
schedule: { cron: { expr: "0 * * * *" } }     # periodic
```

## Quick reference: common validation errors

| Symptom | Rule | Fix |
|---|---|---|
| Expression syntax error | CES-VALID-METHOD-001 | Check operators/parentheses/symbols |
| Non-invertible operator without compensation | CES-VALID-METHOD-002 | Add `compensation` or `invertible: true` |
| Expression nesting too deep | CES-VALID-METHOD-003 | Depth ≤ 64 |
| Multiple schedule modes present | CES-VALID-SCHEDULE-001 | Keep exactly one |
| `external` target | CES-VALID-TARGET-003 | Use `simulator`/`source` instead |
