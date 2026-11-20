<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->

> Language / 语言：[English](sandbox.en.md) | [中文](sandbox.md)

# Plugin Sandbox Model (plugin/sandbox)

English translation of `sandbox.md`, the section-by-section companion under
the plan-A side-by-side naming scheme; section order and tables correspond
one-to-one; no new conclusions are added; this document does not constitute
legal advice.

> **Honest statement: no runsc process was run this round; model wiring
> only.**
> `plugin/sandbox` provides only the **pure-function** configuration model
> and policy validation: it does not invoke the `runsc` binary, does not fork
> any process, does not do dynamic loading, and does not touch the network.
> It abstracts the **configuration semantics** of gVisor runsc into Abate's
> own pure-data documents that can be compared deterministically.

## 1. Model sources

| Abate name | gVisor runsc equivalent semantics | Upstream source (read-only reference) |
|---|---|---|
| `Platform` (`ptrace` / `systrap` / `kvm`) | `--platform` platform-layer selection | `runsc/config/config.go` (`Platform` field), `pkg/sentry/platform/{ptrace,systrap,kvm}` |
| `NetworkPolicy` (`none` / `sandbox` / `host`) | `--network` network-stack selection | `runsc/config/config.go` (`NetworkType`: `NetworkSandbox` / `NetworkHost` / `NetworkNone`) |
| `RootFS` (OCI-style rootfs path + digest) | rootfs bundle and gofer access | `runsc/config/config.go` (`FileAccess`, `Overlay`), `pkg/sentry/fsgofer` |
| `FilesystemPolicy` (read-only/writable mount whitelist) | bind mount and file-access semantics | `runsc/config/config.go` (`FileAccessType`: `FileAccessExclusive` / `FileAccessShared`) |
| `MemoryLimitMB` | Sandbox memory cap | `runsc/config/config.go` (semantics of `TotalMem` and related limit entries) |
| `Capability` (`fs` / `net` / `clock` / `env`) | runsc's capability-switch family | `runsc/config/config.go` (`EnableRaw`, `PauseExternalNetworking`, and related switch family) |

This package models the **configuration semantics** of gVisor `runsc` (an
Apache-2.0 project, historically used as a read-only reference). The repository
has since removed all third-party source; this package is a **pure Go
standard-library in-house rewrite** using Abate's own namespace, importing or
copying no upstream code (`go.mod` has zero dependencies).

## 2. Trim boundary (what was done, what was not)

**Done**:

- The **pure data model** of platform layer / network posture / filesystem
  whitelist / memory cap / capability set;
- **Deterministic derivation** from plugin declaration (`Manifest`) to
  sandbox configuration (`Config`) via `Build`;
- **Validation** of consistency between the capability list and the sandbox
  configuration (`Validate`, default deny-all);
- Deterministic JSON serialization (`Marshal` / `Unmarshal` /
  `Config.Canonical`).

**Explicitly not done (registered honestly)**:

| Not implemented | Reason |
|---|---|
| No `runsc` process invocation, no forking of any process | Project iron rules: no `os/exec`, no process invocation, no dynamic loading of external binaries |
| No runsc checkpoint/restore, cgroups, metrics server | Beyond this package's "parsing + AST + sandbox model" scope |
| No real network stack or real filesystem mounts | This package only produces the configuration model; real isolation is the host runtime's job |
| No implementation of gVisor's sentry platform layer (KVM/ptrace/systrap) | The platform layer is a host capability; this package only declares "which layer to select" |

## 3. Defaults and compatibility constants

| Constant | Value | Semantics |
|---|---|---|
| `DefaultPlatform` | `systrap` | Default platform layer (no KVM dependency, best compatibility) |
| `DefaultNetwork` | `none` | Loopback only by default, zero external network |
| `DefaultMemoryLimitMB` | `256` | Default memory cap (MiB) |
| `MaxMemoryLimitMB` | `4096` | Hard cap; anything above is rejected |
| `DefaultRootFSPath` | `/var/lib/abate/sandbox/rootfs` | Default rootfs bundle path |

The default posture is **deny-all by default**: `DefaultConfig` requests no
capabilities, the rootfs is read-only, and the clock is virtualized.

## 4. Validation caliber (rejection conditions)

`Validate(config, granted)` checks in a fixed order and rejects on the first
hit (errors are tested with `errors.Is`):

| # | Condition | Sentinel error |
|---|---|---|
| 1 | Empty plugin ID | `ErrInvalidManifest` |
| 2 | Platform layer not in the known set | `ErrUnknownPlatform` |
| 3 | Network posture not in the known set | `ErrUnknownNetwork` |
| 4 | Memory cap ≤ 0 | `ErrInvalidMemory` |
| 5 | Memory cap > `MaxMemoryLimitMB` | `ErrMemoryLimit` |
| 6 | A capability requested that the host has not granted | `ErrCapabilityDenied` |
| 7 | rootfs path not absolute or containing `..` | `ErrInvalidPath` |
| 8 | Mount point not absolute or containing `..` | `ErrInvalidPath` |
| 9 | Mount point escaping the rootfs root | `ErrPathEscape` |
| 10 | Non-`none` posture without the `net` capability | `ErrPolicyMismatch` |
| 11 | `none` posture with a network whitelist | `ErrPolicyMismatch` |
| 12 | Write access needed without the `fs` capability | `ErrPolicyMismatch` |
| 13 | Writable rootfs without the `fs` capability | `ErrPolicyMismatch` |
| 14 | Environment-variable whitelist without the `env` capability | `ErrPolicyMismatch` |
| 15 | `clock` capability requested but the virtual clock not enabled | `ErrPolicyMismatch` |

## 5. Determinism guarantees

- `Build` always produces the identical `Config` for the same `Manifest`;
  all whitelists are sorted during derivation.
- `Config.Canonical()` normalizes whitelist order, so **semantically
  identical configurations serialize to identical bytes**.
- `Marshal` uses the standard library `encoding/json`; field order is fixed
  by struct declaration order, with no maps involved.
- The package uses no `time.Now`, no global `math/rand`, and no reliance on
  map iteration order.

## 6. Tests and gates

```bash
GOFLAGS=-mod=vendor go test -race -count=1 ./plugin/sandbox/...
GOFLAGS=-mod=vendor golangci-lint run ./plugin/sandbox/...
```

- Ecosystem-package coverage threshold ≥85% (measured in the wrap-up
  output).
- Unit-test coverage: default configuration, illegal-configuration
  rejection, unauthorized-capability rejection, path-escape rejection,
  determinism (two serializations of the same input are byte-identical),
  JSON round-trip.
- `sandbox_chaos_test.go`: 100 re-runs of the same document + concurrent
  build/validate/serialize.
- `sandbox_mutation_test.go`: per-row assertion of sentinel errors over the
  policy matrix.

## 7. Relation to other components

- The `plugin/wasm` host (self-written WASM runtime) consumes this package's
  `Config` as a **declarative contract** (the host is responsible for real
  isolation); this package is unaware of the wasm runtime.
- Language frontends reuse the same "model only, never execute" boundary:
  this package only produces configuration models; real isolation is the host
  runtime's responsibility.
