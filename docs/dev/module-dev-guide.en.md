<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->

> Language / 语言：[English](module-dev-guide.en.md) | [中文](module-dev-guide.md)

# Module Dev Guide

English translation of `module-dev-guide.md`, the section-by-section
companion under the plan-A side-by-side naming scheme; section order and
tables correspond one-to-one; no new conclusions are added; this document
does not constitute legal advice.

> The module map and dependency layers are in `docs/ARCHITECTURE.en.md`.
> This guide supplements the development DoD and layering iron rules.

## 1. Layers (L0–L7)

| Layer | Example modules | Constraint |
|---|---|---|
| L0 kernel | KRN-ALG/SCH/OBS/VER/MUT/INJ | Depends only on the Go standard library + in-house source-level code; **ten-year no-major-update** |
| L1 WASM | WAS-HOST/CAP/HOT/SIG | The only channel for plugins to enter the kernel (ports-and-adapters) |
| L2 CES | CES-SCHEMA/PARSE/VALID/TCK | Specification layer, same source as JSON Schema / fields.md |
| L3 simulator | SIM-CLK/RNG/NET/FS/SCHED/REC | Deterministic virtual environment |
| L4 language frontend | LNG-AIR/GO/PY/JS/.../RULE/CHK/MUT | Parsing reuse; rules share the AIR layer |
| L5/L6 adapters/reporting | ADP-*/RPT-* | Loaded via WAS-HOST |
| L7 control plane | CTL-CLI/REST/WEB | Entry points |

## 2. Dependency iron rules

- Only top-down, outside-in dependencies are allowed; **L0 does not import
  L1–L7**.
- No circular dependencies (enforced by CI depguard).
- L0 depends only on the standard library and in-house source (the ten-year
  kernel iron rule).

## 3. Module DoD (Definition of Done)

1. Merged into `main` with all five gate stages green; no TODO/FIXME;
   copyright header + watermark present.
2. Tests: unit (table-driven, three paths) + integration + determinism
   replay (same-seed double run with identical hashes). Coverage: kernel
   ≥95%/85%, ecosystem ≥85%/75%.
3. Interfaces: public APIs carry Go doc comments (godot enforces trailing
   periods); breaking changes go through an ADR + a 2-version deprecation
   window.
4. Documentation: module README (responsibilities/interfaces/acceptance
   checklist) + a CHANGELOG entry.
5. Performance: hot paths carry benchmarks within the performance budget
   (see `bench/budget.md`, 15 metrics).
6. Security: new dependencies are included at source level with the source
   commit and license-review verdict recorded.

## 4. Determinism development checklist

- Do not use `time.Now` / `time.Date` to read the real clock; use the
  `sim/clock` virtual clock.
- Do not use global `math/rand`; use the seed-driven PRNG in `sim/rand`.
- Do not depend on map iteration order; sort before comparing/outputting.
- Do not read `os.Environ` to decide behavior; pass configuration explicitly.
- The kernel is single-threaded and driven by the virtual clock; concurrency
  exists only at simulator boundaries, with event sequence numbers
  totally ordered.

## 5. Files and naming

- Each file ≤300 lines; split packages when over-long.
- No TODO/FIXME/godox hits.
- New packages reuse the copyright header byte for byte at the top
  (4-line comment + blank line + package comment).
- Verify only your own package: `go test -race -count=1 ./<pkg>`
  (offline `-mod=vendor`).
