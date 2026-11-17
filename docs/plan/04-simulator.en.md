<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->
> Language / 语言：[中文](04-simulator.md) | [English](04-simulator.en.md)

# Abate L3 Deterministic Simulator Plan — Six Modules in Full Detail

- Document ID: PLAN-04　Status: Effective　Upstream: PLAN-00
- Covered modules: SIM-CLK / SIM-RNG / SIM-NET / SIM-FS / SIM-SCHED / SIM-REC

> **Design-target statement (mandatory for C-class translations)**: This document is a planning document. All `100%` / "pass" / "met"-style statements below are **Design targets / Acceptance targets**, **not measured results**. No claim is made that any target has been achieved.

Common constraints (observed jointly by the six modules):

1. The simulator is a purely software digital twin: any real system calls are forbidden (network/process/filesystem writes are all virtualized); the full test suite passes in an offline environment (acceptance item) (design target).
2. Determinism iron rule: with the same seed and configuration, any number of runs produce byte-identical event streams; all randomness goes through SIM-RNG, all time through SIM-CLK.
3. L3 depends only on L0 kernel types and the standard library; dependencies on L1/L2/L4–L7 are forbidden (L2's CES types do not enter L3 — the simulator only understands kernel operator types).
4. Coverage: statements ≥95%, branches ≥85% (kernel-layer standard); deterministic replay testing is a merge precondition.

---

## Module SIM-CLK: Virtual Clock

### 1. Module Position and Responsibilities

The engine's single time source: an integer-nanosecond virtual clock with four-state operations (fast-forward / slow-motion / freeze / offset). Real time does not exist inside the simulator; all scheduling, timeouts, intervals, and durations are measured in virtual time. It is the shared foundation of both determinism ("steady") and the chaos clock perturbation (clock offset injection).

### 2. Package Structure and File Layout

```
sim/clock/
  clock.go        // Clock interface and VirtualClock implementation
  instant.go      // Instant/Duration types (integer nanoseconds)
  state.go        // Four states: RUNNING/FAST/FREEZE/OFFSET
  clock_test.go
```

### 3. Core Types and Interfaces (must implement)

- `Instant`: int64 nanoseconds (floating point forbidden); `Add(Duration) Instant`, `Sub(Instant) Duration`, `Before/After(Instant) bool`, `UnixNano() int64`.
- `Duration`: int64 nanoseconds; constants `Nanosecond/Microsecond/Millisecond/Second/Minute/Hour`; `Abs()`.
- `Clock` interface:
  - `Now() Instant`: the current virtual instant.
  - `Advance(d Duration) Instant`: advances the clock (callable only by the scheduler; package-level encapsulation prevents misuse — an exported controlled handle `SchedulerHandle`).
  - `Freeze() / Unfreeze()`: freeze/unfreeze (while frozen, Advance is ineffective and returns an error).
  - `Offset(d Duration)`: clock offset injection (the implementation primitive for the chaos clock perturbation: Now() readings have the offset applied, underlying event times unchanged).
  - `Rate(multiplier int64, divisor int64) error`: fast-forward/slow-motion (virtual-time flow rate multiplier; only affects the semantic conversion of Advance, event ordering unchanged).
- `VirtualClock`: implements Clock; internal state (current instant / frozen flag / offset / multiplier numerator and denominator) is entirely integer arithmetic, lock-free (the single-threaded virtual-clock-driving iron rule — concurrency is bounded by the scheduler boundary).

### 4. Key Algorithms and Flows

1. Offset semantics: Offset affects only Now() readings, not the event queue's ordering time (observers see perturbed time while scheduling order is preserved — this is the correct semantics of chaos clock injection).
2. Rate semantics: Rate(2,1) means virtual time flows ×2 (fast-forward); Advance(d) actually advances d×multiplier/divisor; rate changes are recorded as clock events (replayable).
3. Overflow protection: the int64-nanosecond ceiling is roughly 292 years; Advance beyond the bound returns ErrClockOverflow (reject, not wrap).

### 5. Dependencies and Constraints

- Depends on: standard library only (math/bits for overflow checks).
- Depended on by: KRN-SCH (clock advancing), SIM-* (time source for all modules), KRN-INJ (clock offset injection), WAS-HOST (host function abate:clock).
- Forbidden: time.Now/time.Sleep (vet + gate banned-list scanning); floating-point time; real timers.

### 6. Testing Requirements

- Unit tests: full coverage of the four-state operations (running/freeze/offset/rate); semantic cases proving offset does not affect event ordering; overflow rejection cases.
- Determinism test: the same operation sequence executed 100 times, instant-series hash identical.
- Boundary tests: zero/negative durations rejected (Advance parameter validation); max-value boundaries.

### 7. Acceptance Criteria (quantifiable)

1. Four-state operations and combinations (freeze+offset, rate+offset) cases 100% passing (design target).
2. Offset-does-not-affect-ordering semantic test passing (chaos clock injection correctness).
3. Overflow protection cases passing (reject, not wrap).
4. Determinism: hash-identical across 100 runs; coverage ≥95%/85%; gates all green (design target).

---

## Module SIM-RNG: Deterministic PRNG

### 1. Module Position and Responsibilities

The engine's single random source: a seed-driven deterministic pseudorandom stream. The same seed necessarily produces the same sequence; all random consumption is fully accounted (which operator/probe drew how many random numbers at which virtual instant), guaranteeing replay can reproduce exactly. It is the other half of the determinism ("steady") foundation.

### 2. Package Structure and File Layout

```
sim/rand/
  rand.go        // Stream: deterministic random stream
  splitmix.go    // SplitMix64 core algorithm (public domain, merged at source level)
  accounting.go  // Consumption accounting: full record of random draws
  rand_test.go
```

### 3. Core Types and Interfaces (must implement)

- `Stream`:
  - `New(seed uint64) *Stream`: constructed from a seed (SplitMix64 initialization).
  - `Uint64() uint64`, `Int63() int64`, `Float64() float64` ([0,1)), `IntN(n int) int` ([0,n), unbiased rejection sampling), `Bool() bool`, `Bytes(n int) []byte`.
  - `Fork() *Stream`: derives a child stream (the child's seed is derived from the parent's next random number; used for independent random streams in parallel experiments — child streams do not interfere).
  - `State() (seed uint64, draws uint64)`: current state (seed + draws so far).
  - `Restore(seed, draws uint64) error`: state restore (the core of replay: skip draws draws, O(1) or O(log n) fast-forward).
- `Accounting`: a record per draw (virtual instant / caller ID / method / count); `Snapshot() []DrawRecord` for SIM-REC recording.

### 4. Key Algorithms and Flows

1. Core algorithm SplitMix64 (public domain, no license burden): 64-bit state, each output mixed by the golden ratio; period 2^64; passes TestU01 BigCrush (literature conclusion; at acceptance, re-verified with PractRand's 1TB fast tier).
2. Unbiased IntN: rejection sampling (re-draw in the rejection region; no modulo bias).
3. State restore: draws counter + replayed draws (O(draws)) or a precomputed skip table (O(log n); v1 uses the replay method, within the performance budget).
4. Fork semantics: child seed = one parent Uint64() draw; after Fork, parent and child evolve independently; Fork is accounted (guaranteeing an identical Fork sequence at replay).

### 5. Dependencies and Constraints

- Depends on: standard library only.
- Depended on by: KRN-SCH (scheduling randomness), KRN-MUT (mutation point selection), KRN-INJ (fault parameter randomness), SIM-NET (loss/latency jitter), WAS-HOST (host function abate:rand).
- Forbidden: the math/rand global source (global state is unreplayable); crypto/rand (non-deterministic; security scenarios are handled separately by XCU-SEC, not in the simulator); time as a seed.

### 6. Testing Requirements

- Unit tests: boundaries of each method (IntN(0)/IntN(1) rejected or degenerate; Float64 ∈ [0,1)); unbiasedness statistical testing (chi-square, 10^6 samples).
- Determinism test: the same seed's 100 draw sequences hash-identical; parent/child stream independence after Fork (same seed and Fork point ⇒ identical child sequences).
- Replay test: State/Restore round trip (after 10^6 draws, restore and continue — the sequence matches uninterrupted drawing).
- Statistical quality: PractRand fast tier (128MB) with no anomalies (tool output committed).

### 7. Acceptance Criteria (quantifiable)

1. Same-seed sequence determinism: hash-identical across 100 runs; Fork independence cases passing (design target).
2. State/Restore round-trip consistency cases passing (10^6-draw scale).
3. Chi-square unbiasedness test passing (significance level 0.01); PractRand 128MB with no alerts.
4. Draw performance: ≥10^8 draws/second (benchmark report committed); coverage ≥95%/85%; gates all green (design target).

---

## Module SIM-NET: Virtual Network

### 1. Module Position and Responsibilities

An all-software virtual network stack: an in-memory model of nodes/links/routing, supporting four fault injections (latency/partition/packet loss/bandwidth cap) with deterministic propagation emulation. Packet propagation, queuing, and dropping on virtual links are all event-driven (into scheduler queues); replay can reproduce packet by packet. It is the execution foundation of the C-family network operators.

### 2. Package Structure and File Layout

```
sim/net/
  net.go         // Network: topology management (nodes/links/routing)
  node.go        // Node: virtual endpoint (with protocol-stack stubs: TCP/UDP semantic subset)
  link.go        // Link: virtual link (bandwidth/latency/loss rate/partition state)
  packet.go      // Packet: virtual packet (payload/sequence/timestamps)
  transport.go   // Transport interface: the virtual channel for in-simulator http/shell probes
  fault.go       // Fault injection points: link-level implementations of the four fault classes
  net_test.go
```

### 3. Core Types and Interfaces (must implement)

- `Network`: `AddNode(id string) *Node`, `Link(a, b string, cfg LinkConfig) *Link`, `Route(from, to string) []string` (shortest path, deterministic tie-breaking: lexicographic by node ID), `Send(from, to string, payload []byte) (PacketID, error)`.
- `LinkConfig`: `BaseLatency Duration` (base latency), `Jitter Duration` (jitter cap, via SIM-RNG), `BandwidthBps uint64` (bandwidth, bytes/second), `LossRate float64` (loss rate [0,1], via SIM-RNG).
- `Link`: `Partition() error / Heal() error` (partition/heal — during partition, Send returns ErrPartitioned), `SetLatency(d Duration)`, `SetLossRate(r float64)`, `SetBandwidth(bps uint64)` — the link-level primitives for the four fault injections (called by KRN-INJ network.go).
- `Packet`: `ID PacketID` (global sequence), `From/To string`, `Payload []byte` (cap 1MB), `SentAt/DeliverAt Instant` (virtual time).
- `Transport` interface: `Dial(ctx, addr string) (Conn, error)` — the virtual connection for in-simulator HTTP probes and shell probes (Conn semantic subset: Read/Write/Close, all virtualized).

### 4. Key Algorithms and Flows

1. Propagation emulation: Send → route computation → per-hop enqueuing (each hop: loss decision → bandwidth queuing → latency propagation, all generating scheduler events) → delivery event triggers the receiver's callback.
2. Loss decision: an independent Bernoulli trial per packet (SIM-RNG Bool, weighted); a dropped packet generates a drop event (replayable).
3. Bandwidth queuing: token-bucket model (byte tokens, replenished in virtual time); queued packets are FIFO, same-instant ties broken by packet sequence number (deterministic tie-breaking).
4. Partition semantics: after Partition the link is unreachable in both directions (Send immediately returns ErrPartitioned, no propagation events); Heal restores. Partition state changes generate link events.

### 5. Dependencies and Constraints

- Depends on: SIM-CLK (virtual time), SIM-RNG (loss/jitter), KRN-SCH (event enqueuing).
- Depended on by: KRN-INJ (network fault operators), KRN-OBS (http probes via Transport).
- Forbidden: real sockets (net.Dial on the banned list); real DNS; system calls.

### 6. Testing Requirements

- Unit tests: inject → observe → heal full-chain cases for each of the four fault classes; routing determinism (the same topology routed 100 times with identical results).
- Determinism test: the same seed and traffic run twice 100 times, packet event stream (full ordering of send/deliver/drop) hash identical.
- Stress test: a 1000-node / 10000-link topology, 100k packets propagated, event throughput within budget (performance budget item).
- Semantic tests: Send errors during partition and works again after Heal; large-packet transfer time under bandwidth limits matches the token-bucket computation.

### 7. Acceptance Criteria (quantifiable)

1. Four fault injection/heal cases 100% passing; baseline recovery rate 100% after dual rollback (Heal/SetLatency restoration) (design target).
2. Packet event-stream determinism: hash-identical across 100 runs (design target).
3. Stress test: 100k-packet propagation P99 ≤1s (benchmark report committed).
4. Coverage ≥95%/85%; gates all green (design target).

---

## Module SIM-FS: Virtual File System

### 1. Module Position and Responsibilities

An in-memory virtual file system: an in-memory model of directory trees/files/permissions, supporting IO-latency and disk-exhaustion fault injection. Every file operation inside the simulator (probe-shell file access, target file read/write) lands on the virtual FS. It is the execution foundation of the C-family resource operators (IO/disk).

### 2. Package Structure and File Layout

```
sim/fs/
  fs.go          // FS: mount points and directory-tree management
  file.go        // File: virtual file (content/permissions/timestamps)
  dir.go         // Dir: virtual directory
  handle.go      // Handle: open-file handle (read/write cursor)
  fault.go       // Fault injection points: IO delay / disk exhaustion / bad blocks
  fs_test.go
```

### 3. Core Types and Interfaces (must implement)

- `FS`: `MkdirAll(path string) error`, `Create(path string) (*Handle, error)`, `Open(path string) (*Handle, error)`, `Remove(path string) error`, `Stat(path string) (FileInfo, error)`, `Walk(root string, fn func(path string) error) error` (lexicographic traversal, deterministic).
- `Handle`: `Read(p []byte) (int, error)`, `Write(p []byte) (int, error)`, `Seek(off int64) (int64, error)`, `Close() error` — semantics aligned to the os.File subset (error-code semantics aligned: ErrNotExist/ErrPermission/ErrNoSpace).
- `FileInfo`: `Size int64`, `Mode FileMode`, `ModTime Instant` (virtual time).
- Fault injection primitives (called by KRN-INJ resource.go): `SetIODelay(d Duration)` (a virtual delay added to every read/write), `SetQuota(bytes int64)` (disk quota; writes beyond quota return ErrNoSpace — the disk-exhaustion fault), `SetBadBlocks(rate float64)` (bad-block rate: reads return IO errors, via SIM-RNG).

### 4. Key Algorithms and Flows

1. Path normalization: `/` separator, `.` and `..` resolution, root traversal rejected (aligned with path.Clean semantics); path validation at the FS boundary (prevents probe-shell escape).
2. IO delay: every Read/Write generates a delay event (via the scheduler); the operation completes after virtual time advances — observers perceive a realistic IO rhythm.
3. Quota decision: checked before write (current usage + write amount > quota ⇒ ErrNoSpace); quota changes generate FS events.
4. Content storage: []byte entirely in memory (v1 caps: 100MB per file / 1GB total; over the cap is rejected — a memory-bomb guard; large-file scenarios belong to ecosystem-layer extensions).

### 5. Dependencies and Constraints

- Depends on: SIM-CLK (timestamps/delays), SIM-RNG (bad blocks), KRN-SCH (delay events).
- Depended on by: KRN-INJ (IO/disk operators), KRN-OBS (shell-probe file access), LNG frontends (source-target reading, via a port interface).
- Forbidden: real filesystem read/write (os.Create/os.Open on the banned list, test fixtures excepted); real permission system calls.

### 6. Testing Requirements

- Unit tests: all CRUD paths (including error-code semantics); path normalization (traversal/empty path/root escape rejected); inject→heal cases for the three fault classes (IO delay/quota/bad blocks).
- Determinism test: the same operation sequence run twice 100 times, FS event stream and final directory-tree snapshot hash identical.
- Boundary tests: quota filled exactly (boundary value); per-file/total caps rejected.

### 7. Acceptance Criteria (quantifiable)

1. CRUD and error-code semantics cases 100% passing; path-escape rejection cases passing (design target).
2. Three fault injection/heal cases passing; baseline recovery rate 100% after rollback.
3. FS event-stream and directory-tree snapshot determinism: hash-identical across 100 runs (design target).
4. Coverage ≥95%/85%; gates all green (design target).

---

## Module SIM-SCHED: Virtual Scheduling

### 1. Module Position and Responsibilities

A virtual process/goroutine model: creation/termination/suspension/resumption of virtual tasks (VTASK) and CPU-quota emulation. The target under test runs inside the simulator as virtual tasks; process termination/suspension faults (C-family process operators) act on virtual tasks. It is the execution foundation of "process/goroutine-level fault injection".

### 2. Package Structure and File Layout

```
sim/vsched/
  vsched.go      // VScheduler: virtual task table and CPU quota
  vtask.go       // VTask: virtual task (state machine)
  cgroup.go      // Virtual cgroup: CPU/memory quota and pressure
  fault.go       // Fault injection points: kill / suspend / CPU pressure / memory pressure
  vsched_test.go
```

### 3. Core Types and Interfaces (must implement)

- `VTask`: `ID VTaskID`, `Name string`, `State VTaskState` (READY/RUNNING/BLOCKED/SUSPENDED/TERMINATED), `Entry func(ctx VContext)` (task entry — executes within a virtual time slice, yielding on blocking primitives).
- `VScheduler`: `Spawn(name string, entry func(VContext)) VTaskID`, `Kill(id VTaskID) error` (termination — the process-kill fault primitive), `Suspend(id VTaskID) error / Resume(id VTaskID) error` (suspend/resume), `Tasks() []VTaskInfo` (lexicographic).
- `VContext`: the in-task context: `Sleep(d Duration)` (virtual sleep — yields the CPU and generates a wake event), `Yield()`, `Done() <-chan struct{}` (cancellation signal, virtualized channel).
- `VCgroup`: `SetCPUShare(permille int)` (CPU quota, per-mille), `SetMemoryLimit(bytes int64)` (memory quota; exceeding triggers an OOM-Kill event — the memory-pressure fault primitive), `PressureCPU(permille int)` (CPU pressure injection: quota compression emulation).

### 4. Key Algorithms and Flows

1. Scheduling model: cooperative scheduling (tasks yield on blocking primitives or explicit Yield); the ready queue is ordered by (priority, creation sequence) — deterministic tie-breaking; each virtual time slice runs ready tasks until they yield.
2. CPU quota: per virtual time slice, execution budget is distributed weighted by cgroup share (integer slicing, no floating point); over-budget tasks yield and continue next slice.
3. OOM semantics: memory usage over limit → OOM event generated → task state TERMINATED (reason OOM) — distinguished from Kill (fault types are observably distinguishable).
4. Kill/Suspend semantics: Kill terminates immediately (in-flight virtual operations are voided); Suspend freezes (not scheduled, state preserved); Resume returns to READY.

### 5. Dependencies and Constraints

- Depends on: SIM-CLK, SIM-RNG (pressure jitter), KRN-SCH (event enqueuing).
- Depended on by: KRN-INJ (process/resource operators), SIM-NET (protocol-stack stubs as tasks).
- Forbidden: real goroutines directly as the target under test (goroutines are only the host mechanism; virtual-task semantics are independent); real process operations (os/exec on the banned list); real signals (SIGKILL etc.).

### 6. Testing Requirements

- Unit tests: all state-machine transitions (READY→RUNNING→BLOCKED→…→TERMINATED legal paths + illegal transitions rejected); Kill/Suspend/Resume semantic cases; correctness of weighted CPU-quota distribution (integer-slice verification).
- Determinism test: the same seed and task set run twice 100 times, scheduling event stream hash identical.
- Stress test: 10000 virtual tasks scheduled concurrently, event throughput within budget (performance budget item).
- OOM test: memory over limit triggers an OOM event and terminates the task with the correct reason.

### 7. Acceptance Criteria (quantifiable)

1. All state-machine transition cases 100% passing (including illegal-transition rejection) (design target).
2. Scheduling event-stream determinism: hash-identical across 100 runs (design target).
3. Stress test: 10000-task P99 scheduling latency (virtual time) within budget; event throughput ≥100k/second (benchmark report committed).
4. Coverage ≥95%/85%; gates all green (design target).

---

## Module SIM-REC: Record and Replay

### 1. Module Position and Responsibilities

Full-course recording of observations and events (WAL-style append-only log) with frame-by-frame replay: records the entire event stream of an experiment (scheduling events / observation events / RNG draw accounting / clock operations / fault injections) in total order as a replayable archive; replay re-enacts frame by frame and verifies byte by byte. It is the judge and evidence chain for determinism acceptance ("same seed, two runs, hash identical").

### 2. Package Structure and File Layout

```
sim/record/
  record.go      // Recorder: the recorder (multi-source subscription, append-write)
  frame.go       // Frame: frame model (sequence / virtual time / type / payload)
  archive.go     // Archive: archive format (versioned binary, with index and checksum)
  replay.go      // Replayer: the replayer (frame-by-frame re-enactment + verification)
  verify.go      // Verify: dual-run comparison (hash-chain verification)
  record_test.go
```

### 3. Core Types and Interfaces (must implement)

- `Frame`: `Seq uint64` (total order, same source as the scheduler Seq), `At Instant`, `Kind FrameKind` (SCHED/OBS/RNG/CLK/FAULT/NET/FS/VTASK), `Payload []byte` (typed payload, versioned encoding).
- `Recorder`: `Subscribe(sources ...Source)` (subscribes to the scheduler / observation pipeline / RNG accounting / clock / injectors), `Start(ctx)` / `Stop() (*Archive, error)`.
- `Archive`: versioned binary format: header (magic + version + experiment fingerprint) + frame blocks (append-write) + index (by Seq) + trailing SHA-256 whole-archive checksum; `Save(w io.Writer) error`, `Load(r io.Reader) (*Archive, error)` (loading verifies integrity).
- `Replayer`: `Replay(arch *Archive, hooks ReplayHooks) error` — per-frame callbacks (consumed by the re-enactment engine); `Verify(arch *Archive) error` — archive self-verification (checksum + sequence continuity + time monotonicity).
- `Compare(a, b *Archive) (Diff, error)`: dual-run comparison: frame-by-frame hash comparison, outputting the first divergence point (Seq/Kind/expected/actual).

### 4. Key Algorithms and Flows

1. Recording chain: source events → frame encoding (typed + versioned) → append-write (in-memory buffer; flushed to Archive at Stop) → trailing checksum.
2. Replay chain: Load (verification) → per-frame decoding → hooks callbacks (the re-enactment engine replays per frame: restore RNG state / clock state / re-enact events) → per-frame verification (frames produced by replay hash-match the archive frames).
3. Hash chain: each frame carries the previous frame's hash (chained) — tampering with any frame causes all subsequent verification to fail (tamper-proof evidence chain).
4. Dual-run comparison: Compare outputs a structured Diff (first divergence + statistical summary), for CI determinism acceptance and fault diagnosis.

### 5. Dependencies and Constraints

- Depends on: KRN-SCH (Seq same source), KRN-OBS (observation subscription), SIM-CLK/SIM-RNG (state accounting), standard library (hash/sha256, encoding/binary).
- Depended on by: XCU-TST (determinism acceptance judge), CTL-CLI (replay command), RPT-* (replay reports).
- Forbidden: recording affecting the recorded event stream (zero intrusion: read-only subscription); real timestamps in the archive (only virtual time + sequence numbers inside).

### 6. Testing Requirements

- Unit tests: frame-encoding round trip (lossless payload per Kind); Archive Save/Load round trip; checksum tamper detection (flip any byte, Load rejects).
- Replay test: record a complete experiment (including M⊗C mixed perturbation) → replay → frame-by-frame hash identical; correctness of RNG/clock state restore (continuing after replay matches the original run).
- Hash-chain test: tamper a middle frame; all subsequent verification fails.
- Scale test: a 1M-frame archive (~1GB) Load+Verify P99 ≤10s (performance budget item).

### 7. Acceptance Criteria (quantifiable)

1. Frame-encoding and Save/Load round trips 100% lossless; tamper-detection cases 100% passing (design target).
2. Complete-experiment replay frame-by-frame hash identical (including mixed-perturbation scenarios).
3. Hash-chain tamper proof: single-frame tampering causes all subsequent failures — cases passing.
4. Scale performance: 1M frames Load+Verify P99 ≤10s (benchmark report committed).
5. Coverage ≥95%/85%; gates all green (design target).

---

## Appendix: L3 Overall Acceptance (after all six modules are complete)

1. End-to-end determinism: a mixed experiment with network faults (latency + packet loss) + resource pressure (CPU quota) + clock offset, run twice with the same seed 100 times, SIM-REC archives byte-identical (Compare zero differences).
2. Offline verification: in a network-disabled container, the full L3 test suite passes (the zero-real-network-dependency iron rule).
3. Banned-list scanning: time.Now / math/rand global source / net.Dial / os.Create / os.Exec appear zero times inside the sim/ tree (scan script committed, resident in CI).
4. Six-module coverage summary report committed; five-stage gates + benchmark regression all green.
