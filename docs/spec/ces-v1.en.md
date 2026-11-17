<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->

> Language / 语言：[English](ces-v1.en.md) | [中文](ces-v1.md)

# CES v1 Specification Text (CES-SCHEMA)

English translation of `ces-v1.md`, the section-by-section companion under
the plan-A side-by-side naming scheme; section order and tables correspond
one-to-one; no new conclusions are added; this document does not constitute
legal advice.

> Maps to PLAN-03 / PLAN-08 XCU-DOC. CES (Chaos Experiment Specification) is
> Abate's declarative specification for chaos experiments, version
> `ces.tencent.com/v1`. This text, `ces/schema/dist/ces-v1.schema.json`
> (JSON Schema), and `fields.md` share a single source of truth; CI enforces
> zero drift via `scripts/drift-check.sh`.

## 1. Document envelope

Every CES document is a top-level object with a fixed envelope plus the `spec`
executable body. Unknown fields are **always rejected** (fail-fast, to avoid
silently ignoring typos).

| Field | Type | Required | Description |
|---|---|---|---|
| `apiVersion` | string | ✅ | Fixed `ces.tencent.com/v1` |
| `kind` | string | ✅ | Fixed `Experiment` |
| `metadata.name` | string | ✅ | DNS subdomain; namespace defaults to `default` |
| `metadata.namespace` | string | ⬜ | Defaults to `default` |
| `spec` | object | ✅ | Executable body (see below) |

## 2. The spec executable body

| Field | Required | Description |
|---|---|---|
| `target` | ✅ | Perturbation target: `kind` ∈ {`simulator`,`source`,`external`}; `external` is a reserved field, unsupported at runtime, **rejected by validation**. `selector` is a stable key/value set (package/function/node/process). |
| `steadyState` | ✅ | At least one steady-state hypothesis. Each entry references one probe and contains `expectation`/`tolerance`/`holdFor`. |
| `method` | ✅ | Perturbation algebra expression (see §3). |
| `probes` | ⬜ | Observation source list; `id` is unique within the document; registered in ID order (order does not affect the observation stream). |
| `rollback` | ✅ | Compensation strategy: `strategy` ∈ {`auto`,`manual`,`timeout`} (normalized default `auto`); the `timeout` strategy must fill `timeout`. |
| `schedule` | ✅ | Mutually exclusive one-of-three: `once` / `cron` / `continuous` (pointer-level exclusivity, see CES-VALID-SCHEDULE-001). |
| `timeout` | ⬜ | Optional; normalized with defaults filled. |
| `seed` | ⬜ | Optional; normalization derives a deterministic seed. |

## 3. Perturbation algebra (Method Expression)

`method.expression` uses the CES operator-calculus notation:

| Symbol | Meaning |
|---|---|
| `∘` | composition |
| `⊗` | concurrent product |
| `a⁻¹` | dual (inverse compensation) |
| `ε` | identity |
| `( )` | grouping |

- Operands are `operator://` references (resolving inline definitions in
  `method.operators` or references).
- Depth limit **64** (the CES-PARSE/METHOD-003 attack surface, guarding
  against over-deep nested injection).
- Non-invertible operators must declare `compensation`, otherwise
  CES-VALID-METHOD-002 rejects.

## 4. Validation rule catalog (CES-VALID-*, 31 rules total)

Grouped by family; each rule can report `error` or `warning` severity
(`--strict` promotes warnings to failures).

| Family | Rules | Concerns |
|---|---|---|
| STEADY | 001–005 | Steady-state hypothesis completeness, probe reference existence, expectation/tolerance type matrix, holdFor legality |
| ROLLBACK | 001–003 | Strategy legality, timeout strategy must fill timeout, compensation consistency |
| SCHEDULE | 001–004 | Three-mode exclusivity, cron expression legality, once/continuous field legality |
| METHOD | 001–004 | Expression syntax, dual compensation (invertibility), depth limit (003), operator resolution |
| PROBE | 001–005 | Probe id uniqueness, kind-field matching, positive sampling interval, command/URL safety |
| TARGET | 001–004 | Target kind legality, selector key legality, `external` rejected, domain matching |

## 5. Parsing rule catalog (CES-PARSE-*, 13 rules total)

Structural and envelope validation during parsing/normalization (e.g.
envelope required fields, version negotiation, YAML/JSON decoding,
unknown-field rejection, field-type legality). Normalization fills defaults
(namespace/strategy/seed/timeout) and derives the deterministic seed.

## 6. Conformance (TCK)

`ces/tck` provides a data-driven conformance test suite (164 cases: parse 68 /
validate 46 / roundtrip 50). `ces/tck/example` is the reference third-party
implementation, demonstrating that it consumes the neutral case format and is
fully conformant. A third-party implementation only needs to implement the
`tck.Driver` interface (Parse/Validate) to be certified against the same case
set.
