<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->

> Language / 语言：[English](02-wasm-runtime.en.md) | [中文](02-wasm-runtime.md)

# Abate L1 WASM Runtime Plan — Four Modules in Extreme Detail

**Design-target statement (mandatory for C-class translations)**: this file
is the English translation of an internal planning document. All `100%` /
"pass" / "met" style statements in the full text are **Design targets /
Acceptance targets**, **not measured results**; they must not be read as
established facts or capability claims. The translation corresponds
section by section, does not escalate the tone, and does not delete target
qualifiers.

- Document ID: PLAN-02　Status: in force　Upstream: PLAN-00
- Covered modules: WAS-HOST / WAS-CAP / WAS-HOT / WAS-SIG
- Common constraints (observed by all four modules):
  1. L1 is the only channel through which L4-L7 plugins enter the kernel
     (the ports-and-adapters iron rule); L1 depends only on L0 kernel
     interfaces and the standard library, and reverse dependence on L4-L7 is
     forbidden.
  2. The WASM host runtime uses source-level inclusion (the five-step
     process); candidate: wazero (Go implementation, zero dependencies,
     Apache-2.0, no CGO, supports WASI Preview2) — after vendoring: secondary
     changes (determinization: clock and random source rerouted to
     SIM-CLK/SIM-RNG), optimization (trim unneeded features), deep
     integration (the WIT binding generator aligns with Abate's plugin
     interface), encapsulation (expose only the Host type externally).
  3. Plugins are never trusted: zero capabilities by default, everything
     granted explicitly; unsigned plugins are refused.
  4. Coverage: statement ≥85%, branch ≥75% (ecosystem-layer standard);
     security test cases are a merge prerequisite.

---

## Module WAS-HOST: WASM host

### 1. Positioning and responsibilities

Provides loading, instantiation, invocation, and lifecycle management of WASM
modules; defines the unified WIT (WebAssembly Interface Types) interfaces for
the five plugin types (operator/injector/frontend/adapter/reporter); injects
host capabilities (logging, clock, randomness, observation reporting) into
plugins as host functions. It is the runtime base of the 「劲」 (agile,
hourly ecosystem extension) character.

### 2. Package structure and file layout

```
wasm/host/
  host.go          // Host: load/instantiate/call entry
  wit.go           // WIT interface definitions for the five plugin types (versioned)
  instance.go      // Instance: single-plugin-instance wrapper (memory table/call handle)
  hostfunc.go      // host function implementations: log/clock/rand/observation reporting
  registry.go      // Plugin registry: type → factory
  host_test.go
```

### 3. Core types and interfaces (must implement)

- `Host`:
  - `Load(ctx, path string, opts LoadOptions) (*Instance, error)`: loads a
    .wasm file (signature-verified via WAS-SIG first, then instantiated).
  - `Registry() *Registry`: access to the plugin registry.
- `LoadOptions`: `Capabilities Capabilities` (grants; default deny-all),
  `Fuel uint64` (execution fuel cap against infinite-loop plugins), `MemoryMB
  uint32` (memory cap, default 64MB).
- `Instance`:
  - `Call(ctx, entry string, args []byte) ([]byte, error)`: invokes a plugin
    export (entry-whitelist validated).
  - `Type() PluginType`: one of the five types.
  - `Version() SemVer`: the plugin's self-reported version (cross-checked
    against signature metadata).
  - `Close() error`: drains and releases (coordinated with WAS-HOT).
- WIT interfaces (one set per plugin type, versioned, e.g.
  `abate:plugin/operator@1.0.0`):
  - operator: `describe() -> OperatorSpec` (operator declaration:
    ID/family/parameter schema/invertibility).
  - injector: `inject(spec) -> Result` / `rollback(spec) -> Result`.
  - frontend: `parse(source) -> AirBytes` (produces AIR serialization bytes).
  - adapter: `notify(event) -> Result`.
  - reporter: `report(report) -> Result`.
- Host functions (host-injected, all capability-gated): `abate:log` (logs),
  `abate:clock` (virtual-clock reads), `abate:rand` (deterministic random
  stream), `abate:observe` (observation reporting).

### 4. Key algorithms and flows

1. Load chain: read file → WAS-SIG signature verification (reject on
   failure) → capability declaration parsing → WAS-CAP granting →
   instantiation (fuel + memory caps injected) → registry registration.
2. Call chain: entry-whitelist validation → fuel check → invoke → result
   validation (length cap 1MB against memory bombs) → return.
3. Fuel exhaustion: the call terminates immediately, returns
   ErrFuelExhausted, and the instance is marked degraded (not destroyed;
   reusable).

### 5. Dependencies and constraints

- Depends on: wazero included at source level (third_party/wazero,
  Apache-2.0), KRN-ALG (operator types), SIM-CLK/SIM-RNG (host-function
  backends).
- Depended on by: WAS-CAP/WAS-HOT/WAS-SIG, and the L4-L7 plugin loaders.
- Forbidden: CGO; real network listeners; plugins directly accessing the
  filesystem (must go through capability-granted host functions).

### 6. Test requirements

- Unit tests: one minimal .wasm test fixture per plugin type (Rust/TinyGo
  build artifacts committed), covering the full load → call → close chain.
- Boundary tests: fuel exhaustion, memory over-limit, over-long results,
  entry-whitelist rejection.
- Determinism tests: under the same seed, host functions (rand/clock) produce
  identical output sequences.

### 7. Acceptance criteria (quantifiable)

1. All five plugin types' WIT interfaces defined and versioned; all test
   fixtures pass the full chain.
2. Each of the fuel/memory/result-length caps has a triggering case with the
   correct error code.
3. Plugin-call P99 overhead ≤50μs (excluding first instantiation; benchmark
   report filed).
4. Coverage ≥85%/75%; all gates green.

---

## Module WAS-CAP: capability sandbox

### 1. Positioning and responsibilities

The capability-based security model: five capability classes
(fs/net/clock/env/observe), default deny-all, explicit granting, converged to
per-plugin declarations. It is the first wall of the 「强」 (strong, defense
in depth) character.

### 2. Package structure and file layout

```
wasm/capability/
  capability.go    // Capabilities bitset and the five capability definitions
  grant.go        // Grant parsing: declaration → capability set (least-privilege convergence)
  check.go        // Runtime check: gating before host-function calls
  audit.go        // Capability audit log (grant/deny/use all recorded)
  capability_test.go
```

### 3. Core types and interfaces (must implement)

- The `Capability` enum: `CapFS` (restricted file access, plugin data
  directory only), `CapNet` (restricted network, declared host whitelist
  only), `CapClock` (virtual-clock reads), `CapRand` (deterministic
  randomness), `CapObserve` (observation reporting).
- `Capabilities`: bitset type; `Has(Capability) bool`, `Grant(Capability)`,
  `Revoke(Capability)`.
- `ParseGrant(decl string) (Capabilities, error)`: parses a plugin capability
  declaration (manifest format: `capabilities: [clock, rand]`); unknown
  capability names refuse loading (a forward-compatible whitelist approach).
- `Check(cap Capabilities, c Capability) error`: gates host-function calls;
  unauthorized returns ErrCapabilityDenied and is audited.
- `AuditLog`: grant/deny/use events all recorded (plugin ID + capability +
  time + result), for security auditing and XCU-SEC consumption.

### 4. Key algorithms and flows

1. Least-privilege convergence: the plugin's declared capabilities intersect
   with host policy (the policy file is configurable); an empty intersection
   means zero capabilities.
2. Network whitelist: CapNet must come with a host whitelist (manifest
   `net-hosts`); DNS resolution happens on the host side (guarding against
   DNS rebinding).
3. File sandbox: CapFS maps only the plugin-specific data directory
   (`/var/lib/abate/plugins/<id>/`); path traversal (`..`) is validated after
   host-side normalization.

### 5. Dependencies and constraints

- Depends on: no external dependencies (pure standard library).
- Depended on by: WAS-HOST (load chain and host-function gating).
- Forbidden: any "default-on" capability; capability escalation requests —
  runtime capability requests from plugins are always rejected (capabilities
  are granted once, at load time).

### 6. Test requirements

- Unit tests: the five-capability grant/deny matrix (5×2 full combination);
  declaration parsing (legal/illegal/unknown capability names).
- Security tests (merge prerequisite): all unauthorized calls intercepted
  (one escalation case per capability class); path-traversal cases;
  DNS-whitelist bypass cases.
- Audit tests: grant/deny/use events all land in the audit log.

### 7. Acceptance criteria (quantifiable)

1. The 5×2 capability-matrix cases pass 100%; the escalation-interception
   cases pass 100% (design targets).
2. Capability-declaration parsing: legal/illegal/unknown-name classes each
   have cases; unknown names refuse loading 100%.
3. Audit-log event coverage of the three classes is 100% (cases assert log
   entries).
4. Coverage ≥85%/75%; all gates green.

---

## Module WAS-HOT: hot-loading

### 1. Positioning and responsibilities

Runtime plugin upgrades: atomic replacement (new version instantiated →
traffic switched → old instance drained → unloaded) without interrupting
running experiments. It is the direct landing point of the 「劲」 character
(seamless upgrades in the hourly ecosystem cadence).

### 2. Package structure and file layout

```
wasm/hotswap/
  swap.go        // Swap: the atomic-replacement main flow (state machine)
  drain.go       // Drain: waiting for old-instance in-flight calls to finish (with timeout)
  compat.go     // Compatibility matrix: semver range decisions
  hotswap_test.go
```

### 3. Core types and interfaces (must implement)

- `Swap(ctx, old, new *Instance) error`: the atomic-replacement state machine:
  `PREPARE` (new instance loaded and signature-verified) → `SWITCH` (atomic
  registry replacement, a single pointer write) → `DRAIN` (wait for old
  instance in-flight calls) → `UNLOAD` (unload the old instance).
- `DrainTimeout`: default 30s (real time); on timeout the old instance is
  force-Closed (in-flight calls return ErrDrainTimeout).
- `CompatMatrix`: `Allowed(pluginType, from, to SemVer) bool` — maintains
  semver compatibility ranges per plugin type (major changes are incompatible
  by default and require explicit host-policy approval).

### 4. Key algorithms and flows

1. Atomicity: the registry replacement is a single atomic write (sync/atomic
   or a single assignment inside an RWMutex write lock); the read side never
   blocks.
2. Draining: the old instance is unloaded only when its reference count
   reaches zero (in-flight calls hold references); reference-leak detection
   (a drain timeout reports the leaking plugin ID).
3. Failure rollback: a PREPARE failure (signature/instantiation/compatibility)
   leaves the old instance serving untouched.

### 5. Dependencies and constraints

- Depends on: WAS-HOST (instance management), WAS-SIG (new-version signature
  verification).
- Depended on by: CTL-CLI/CTL-REST (upgrade command entries).
- Forbidden: any intermediate registry state visible during replacement (the
  read side sees either all-old or all-new); unloading an instance with
  in-flight calls is forbidden.

### 6. Test requirements

- Concurrency tests: 100 concurrent readers + 1 replacement; zero errors and
  zero intermediate states on the read side (asserting each call hits a
  single instance version).
- Drain tests: slow in-flight calls (a sleep plugin) + replacement → the call
  completes → the old instance unloads; the timeout path (force close).
- Failure rollback tests: corrupted new version (bad wasm bytes) → the old
  instance continues serving unaffected.

### 7. Acceptance criteria (quantifiable)

1. Concurrent replacement: 100 concurrent readers, zero errors and zero
   intermediate states (passing under the `-race` gate).
2. Drain and timeout paths each have cases; reference-leak detection has a
   case.
3. Failure rollback: bad bytes / signature failure / compatibility failure —
   old-instance availability preserved 100%.
4. Coverage ≥85%/75%; all gates green.

---

## Module WAS-SIG: signature verification

### 1. Positioning and responsibilities

Plugin signature-chain verification: GPG/cosign dual-channel verification and
cross-checking of signature metadata (plugin ID/version/capability-declaration
hash). It is the second wall of the 「强」 character and the entrance of
supply-chain security.

### 2. Package structure and file layout

```
wasm/signature/
  verify.go       // Verify: the signature-verification main entry (GPG first, cosign alternate)
  gpg.go          // GPG channel: openpgp included at source level (golang.org/x/crypto/openpgp, BSD-3)
  cosign.go       // cosign channel: sigstore signature blob verification (source-level inclusion)
  manifest.go     // Plugin manifest parsing and hash cross-checking
  truststore.go    // Trust anchors: built-in official keys + user-added keys
  signature_test.go
```

### 3. Core types and interfaces (must implement)

- `Verify(path string) (*Manifest, error)`: the main entry — signature file
  (`<name>.wasm.sig`) + trust-anchor verification + manifest hash
  cross-check; any failure returns ErrSignatureInvalid and refuses loading.
- `Manifest`: `PluginID`, `Version SemVer`, `Type PluginType`,
  `CapabilitiesDecl string`, `SHA256 [32]byte` (hash of the wasm bytes),
  `SignedAt time.Time`.
- `TrustStore`: `Add(key)` (user-added), `Builtin()` (official release keys,
  updated with versions); key fingerprints deduplicated.
- Cross-check: the manifest's declared SHA256 matches the actual wasm byte
  hash; the signature covers all manifest fields (tamper protection).

### 4. Key algorithms and flows

1. Verification chain: read .wasm and .sig → find the public key in the trust
   anchor → GPG/cosign verification → manifest hash cross-check → only then
   return the Manifest to WAS-HOST.
2. Dual-channel policy: the signature file carries a channel marker; GPG
   signatures go through openpgp, cosign signatures through sigstore blob
   verification; the two channels never downgrade each other (there is no
   "fall back to unsigned on verification failure" path).
3. Key rotation: the trust anchor supports multiple coexisting keys (old/new
   overlap window); official keys update with Abate versions and their
   fingerprints are published.

### 5. Dependencies and constraints

- Depends on: openpgp included at source level (third_party/openpgp), the
  cosign verification library included at source level
  (third_party/sigstore, Apache-2.0).
- Depended on by: WAS-HOST (load chain), WAS-HOT (upgrade chain).
- Forbidden: any "skip verification" switch (except a test-only injection
  interface, disabled in production builds); private keys never enter the
  repository (only public-key trust anchors).

### 6. Test requirements

- Unit tests: one set each for GPG/cosign (valid signature passes, tampered
  bytes rejected, wrong key rejected, manifest-hash mismatch rejected).
- Trust-anchor tests: multi-key coexistence, user-added keys, fingerprint
  deduplication.
- Security tests (merge prerequisite): unsigned plugin loading refused 100%;
  tampered manifests refused 100%.

### 7. Acceptance criteria (quantifiable)

1. Dual-channel verification passes all positive/negative examples (≥4 cases
   per channel).
2. The four rejection classes (unsigned/tampered/wrong-key/hash-mismatch)
   pass 100%.
3. Verification overhead P99 ≤10ms (benchmark report filed).
4. Coverage ≥85%/75%; all gates green.

---

## Appendix: L1 overall acceptance (after all four modules)

1. End to end: a signed plugin walks the full chain — Load (verify → grant →
   instantiate) → Call → Swap (hot upgrade) → Close — without interrupting a
   running experiment.
2. Security audit: the three attack classes (escalation/unsigned/tampering)
   are all intercepted (security test report filed).
3. The plugin's full flow (commit → sign → publish → hot-load) measured at
   ≤1 hour (an M4 acceptance item).
