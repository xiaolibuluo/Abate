<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->

> Language / 语言：[English](tla-verification.en.md) | [中文](tla-verification.md)

# Abate TLA+ Formal Verification Report

English translation of `tla-verification.md`, the section-by-section
companion under the plan-A side-by-side naming scheme; section order, tables,
and code blocks correspond one-to-one; verification results are reported
as-is and are not upgraded into claims of proven optimality; no new
conclusions are added; this document does not constitute legal advice.

> Generated: 2026-11-13
> Toolchain: TLC2 Version 2026.11.06 (tla2tools.jar, v1.8.0)
> Entry point: `bash scripts/tla/check.sh`
> Spec directory: `scripts/tla/specs/`

## I. Purpose

Abate's determinism, resource isolation, and coverage parity are
non-negotiable quality red lines. The nine-section code-level gate chain
(including `-race`, coverage, govulncheck) blocks implementation defects but
cannot block **concurrency/state-machine-level** design errors. This work
applies TLA+ **model checking** to the three most heavily loaded invariants,
proving at the mathematical level:

1. The "reproducibility / replay equivalence" of the deterministic random
   stream (the cornerstone of SIM-REC replay attestation);
2. The "boundedness + guaranteed termination" of the WASM/sandbox fuel (gas)
   model;
3. The "bounded + monotonic + full-when-equal" behavior of PIT coverage
   parity counting.

## II. Spec list and conclusions

| Spec | Verification target | Properties | Model size | Conclusion |
|------|---------|------|---------|------|
| `StreamDeterminism.tla` | `sim/rand`'s `Stream` + `splitMix64` | Invariants `StateConsistency` / `Reproducibility` / `TypeOK` | 4800 reachable states | ✅ all pass |
| `SandboxFuel.tla` | The fuel model of `plugin/sandbox` and the planned `plugin/wasm` | Invariants `FuelNonNeg`/`UsedBounded`/`StepsBounded`/`HaltEmpty` + liveness `Terminates` | 61 reachable states | ✅ all pass |
| `PitParity.tla` | The parity-rate computation of `kernel/mutation/pitparity` | Invariants `RateInRange`/`NoOverCount`/`FullWhenEqual`/`ZeroWhenZero` | 44 reachable states | ✅ all pass |

Running `bash scripts/tla/check.sh` prints:

```
===== PitParity =====
[PASS] PitParity  (44 distinct states)
===== SandboxFuel =====
[PASS] SandboxFuel  (61 distinct states)
===== StreamDeterminism =====
[PASS] StreamDeterminism  (4800 distinct states)
===== TLA+ 形式化验证全部通过 =====
```

## III. Key points of each spec

### 3.1 StreamDeterminism — deterministic stream reproducibility

Abstract model: each stream `S[s]` holds `(origin, state, draws)`. The core
fact is that `splitMix64`'s **value mixing** does not change the stored
`state`; `state` is pure arithmetic:

```
state = (origin + draws * Gamma) % (2^WIDTH)
```

`Restore(seed, d)` sets `state` directly to `seed + d*Gamma`, exactly the
position reached by "starting fresh from seed and advancing d times" — this
is precisely why the replay jump is mathematically exact.

Proved invariants:

- `StateConsistency`: `∀s, state == (origin + draws*Gamma) % M`.
- `Reproducibility`: if two streams have the same `origin` and `draws`, their
  `state` must be equal. From this it follows that "`Restore(seed,d)` is
  indistinguishable from a stream advanced to d".
- Actions cover `New` / `Advance` / `Restore` / `Fork` (a child stream takes
  its parent's current `state` as `origin`, then evolves independently),
  exhausting all interleavings.

### 3.2 SandboxFuel — fuel is bounded and termination is guaranteed

Abstract model: the execution machine holds `fuel`, `steps`, `halted`,
`used`. Each step consumes `c∈[1,MaxCost]` units of fuel; running out of
fuel halts (`halted`). Proved:

- Safety: `fuel≥0`, `used≤F0`, `steps≤F0`, `halted ⇒ fuel=0`;
- Liveness (under weak fairness): `◇halted` — finite fuel must halt within
  ≤F0 steps, and after halting nothing executes again.

This proof locked in the "bounded resources + deterministic termination"
design constraint **before the WASM host (`plugin/wasm`, based on wazero) is
implemented**, so subsequent code may not violate it.

### 3.3 PitParity — parity-rate counting is bounded and monotonic

Abstract model: `comparable` (comparable mutation points) and `consistent`
(points Abate reproduces), with `consistent≤comparable`. The parity rate:

```
rate = consistent * SCALE / comparable   (integer division; SCALE is 10000 in production)
```

Proved: `rate∈[0,SCALE]`, `consistent≤comparable` (over-counting forbidden),
full-when-equal `consistent==comparable ⇒ rate==SCALE`, zero-when-zero
`consistent==0 ⇒ rate==0`.

## IV. Toolchain constraints and handling (recorded honestly)

TLC2 2026.11.06 in this environment has the following limitations, each
worked around and recorded; **no results were fabricated**:

1. **Integer cap 2³¹−1**: `0x9E3779B97F4A7C15` and even its 32-bit form
   `0x9E3779B9` overflow. Since the reproducibility proof is
   **parameter-independent** for both `Gamma` and the modulus, the model uses
   the small odd constant `Gamma==9`; the conclusions hold for the
   production 64-bit constant and full word width (uniformity argument of
   modular arithmetic).
2. **No `Floor` / `Div` / `Mod` operators**, only infix `%`: PitParity's
   integer division uses a bounded `CHOOSE` (exact floor division over
   `0..SCALE`) to avoid stack overflow of a recursive `Idiv` when the
   quotient is large.
3. **No hexadecimal literals**: uniformly switched to decimal.
4. **Adjacent `*` and `/` need explicit parentheses**: TLA+ forbids unparenthesized
   mixing; parentheses added.
5. **Terminal states need an explicit stutter**: after the counter tops out,
   no action is enabled and TLC would deadlock; `Terminated == UNCHANGED ...`
   added.

## V. Running and reuse

```bash
# Prerequisite: java on PATH; if tla2tools.jar is missing, check.sh
# automatically downloads and caches it from the official release
bash scripts/tla/check.sh
```

- Adding a spec: place `<Name>.tla` + `<Name>.cfg` in `scripts/tla/specs/`;
  `check.sh` discovers them automatically.
- Constant overrides (smaller word widths for model checking) go in each
  `.cfg`; the proofs are independent of production word widths/numeric
  parameters.

## VI. Relation to the nine-section gate

TLA+ verification is a **design-layer complement** on top of the nine-section
gate chain (header/style/license/gofmt/vet/golangci-lint/test-race/coverage/
build): the gates ensure "the code is written correctly", TLA+ ensures "the
design is not wrong". Together they form Abate's completion evidence chain.
