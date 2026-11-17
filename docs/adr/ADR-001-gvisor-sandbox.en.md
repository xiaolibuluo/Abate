<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->

> Language / 语言：[English](ADR-001-gvisor-sandbox.en.md) | [中文](ADR-001-gvisor-sandbox.md)

# ADR-001: Adopt gVisor as the capability sandbox (source-level inclusion)

English translation of `ADR-001-gvisor-sandbox.md`, the section-by-section
companion under the plan-A side-by-side naming scheme; section order and
tables correspond one-to-one; the decision context and conclusions are
preserved unchanged; no new conclusions are added; this document does not
constitute legal advice.

- Status: Accepted
- Date: 2026-11-08
- Proposed by: project owner (xiaolibuluo@163.com)
- Related: ARCHITECTURE.md §5.2 security model, PLAN-00 module map WAS-CAP,
  PLAN-08 crosscutting security system

---

## 1. Background and decision

Project requirements: **all admitted open source must be included at source
level (API access forbidden), must be written in Go, and GPL-licensed open
source is forbidden** (the supply-chain compliance iron rule under the
commercial dual-license premise).

The sandbox is central to Abate's defense in depth: the L1 `WAS-CAP`
capability sandbox grants fs/net/clock/env capabilities to WASM plugins with
default deny-all; the M4 WASM runtime needs process/namespace-level isolation
on the host side to prevent malicious or unsigned plugins from escaping.

After evaluation, **gVisor** (`gvisor.dev/gvisor`) is adopted as the sandbox
base.

---

## 2. Candidate comparison (facts)

| Candidate | Language | License | Go? | GPL? | Verdict |
|---|---|---|---|---|---|
| gVisor (runsc / pkg/sentry) | Go | Apache-2.0 | Yes | No | **Adopted** |
| firecracker (vmm) | Rust | Apache-2.0 / BSD-3 | No (Rust) | No | Excluded: not Go, violates the "must be Go" iron rule |
| containerd + runc (libcontainer) | Go | Apache-2.0 | Yes | No | Alternate: requires real kernel namespaces; nested sandboxing unavailable |
| bubblewrap (bwrap) | C | LGPL-2.1 | No (C) | **Yes (LGPL)** | **Excluded**: non-Go + GPL-family, double violation |

---

## 3. License review verdict

- gVisor's main license: **Apache License 2.0** (permissive; compatible with
  the project's own Apache-2.0-compatible terms; no copyleft contamination of
  the commercial license).
- Contains no GPL/LGPL/AGPL variant of any kind; no copyleft contagion risk.
  **Complies with the "no GPL inclusion" iron rule.**
- At inclusion time it is vendored under `third_party/gvisor/`, and
  `third_party/LICENSE-MANIFEST.md` records the source commit hash and the
  license-review verdict (DoD item 6).

---

## 4. Source-inclusion strategy (the five-step inclusion process)

Strictly following the five-step inclusion process of ARCHITECTURE.md §2.1;
any API/out-of-process/network access is forbidden:

1. **Vendor**: pull the relevant gVisor source (the `pkg/sentry` isolated
   kernel subset, not the full runsc CLI) into `third_party/gvisor/`, pinning
   the source commit hash.
2. **Secondary changes**: determinism adaptations (aligned to virtual
   clock/seed semantics); remove CLI/network paths unrelated to the kernel.
3. **Optimization**: trim non-core paths, converge to a library shape
   (sandbox primitives), remove external-dependency drift.
4. **Deep integration**: wrap the gVisor isolation primitives as the
   underlying implementation of `WAS-CAP` capability granting / default
   deny-all, without exposing gVisor-specific types directly.
5. **Encapsulation**: expose only the Abate capability-sandbox interfaces
   (fs/net/clock/env ports) externally; all gVisor concepts are translated
   into kernel vocabulary.

---

## 5. Scope and milestone assignment (important)

The gVisor sandbox belongs to **M4 (WASM runtime + minimal ecosystem loop)**,
corresponding to modules **WAS-CAP / WAS-HOST**, and is **not within the M1
minimal kernel loop**.

Reasons (facts, not expediency):

- Architectural iron rule: **L0/L3 kernels must be pure-software digital
  twins with zero external runtime dependencies** (PLAN-04 general constraint
  1, PLAN-00 dependency iron rule 3). Linking gVisor directly into the M1
  kernel would break that rule.
- M1's "sandbox" is already carried by the L3 deterministic simulator (fully
  virtualized clock/network/FS/scheduler) and the self-contained L0 kernel;
  no real platform sandbox is needed.
- This ADR completes only the four immediately actionable items: "decision +
  license review + source-inclusion strategy + gate constraints"; the actual
  source vendoring and encapsulation execute at M4, bound by this ADR and
  LICENSE-MANIFEST.

---

## 6. Gate constraints (executable immediately)

- `scripts/license-check.sh` lists gVisor as a reviewed whitelist entry
  (Apache-2.0); any new dependency judged GPL-family makes the gate **FAIL**
  directly.
- All in-house kernel packages import no external runtime other than
  `third_party/gvisor`; gVisor is accessed through ports only at L1, and L0
  is fully unaware of it (dependency iron rules 1 and 3).

---

## 7. Positive consequences

- Defense in depth gains a userspace-kernel isolation base validated at
  large production scale, without introducing GPL compliance risk.
- Full-chain source self-sufficiency, no supply-chain hijack surface
  (XCU-SEC goal).

## 8. Negative consequences / risks

- gVisor is large; after vendoring, an upstream security-patch backport
  process must be maintained (managed the same way as chaosblade/PIT).
- Nested sandboxing (running gVisor inside a container/sandbox) is
  unavailable in some CI environments; the runtime matrix must be clarified
  at M4 (ptrace platform first, KVM platform as needed).

## 9. Alternatives

- If the M4 runtime environment cannot carry gVisor, fall back to the pure-Go
  isolation-primitive subset of the containerd/runsc family, or an in-house
  lightweight syscall filtering layer; both are Apache-2.0 / Go and do not
  violate this ADR's license and language iron rules.
