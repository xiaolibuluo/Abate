<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->

> Language / 语言：[English](quickstart.en.md) | [中文](quickstart.md)

# User Quickstart

This guide gets you from a clean machine to building Abate and running your
first deterministic chaos experiment in 5 minutes. Every command can be copied
and executed as-is.

## 1. Build

The project has **zero external dependencies** (`go.mod` declares only the module
and Go version; `go.sum` is empty and there is no `vendor/`). Go 1.26.11+ builds
it offline:

```bash
export GOPROXY=off
go build -o abate .
./abate version
```

`abate version` prints build provenance (version / schema / license / watermark):

```
abate 1.2.27
schema: ces.tencent.com/v1
build: deterministic
license: AGPLv3 or commercial
source: https://github.com/xiaolibuluo/Abate
watermark: 小黎的部落
```

## 2. Write a CES experiment

Create `hello.yaml`:

```yaml
apiVersion: ces.tencent.com/v1
kind: Experiment
metadata:
  name: hello-ces
spec:
  target:
    kind: simulator
    selector: { node: n1 }
  steadyState:
    - probeRef: p1
      expectation: { type: value, value: ok }
      tolerance: { type: absolute, value: 1 }
      holdFor: 2s
  method: { expression: ε }
  probes:
    - id: p1
      kind: metric
      spec: { query: up, expect: "1" }
      interval: 1s
  rollback: { strategy: auto, verifySteadyState: false }
  schedule: { continuous: { duration: 1s } }
```

## 3. Run

```bash
./abate run hello.yaml
```

Add `--format json --out report.json` for a machine-readable report; add
`--record /tmp/run.arch` to write a deterministic SIM-REC archive.

## 4. Verify determinism (core property)

With the same seed and the same configuration, any number of runs produce
**byte-for-byte identical** results:

```bash
./abate run hello.yaml --seed 42 --format json --out a.json
./abate run hello.yaml --seed 42 --format json --out b.json
./abate verify a.json b.json      # exit 0 if byte-for-byte identical
```

## 5. Replay evidence

```bash
./abate run hello.yaml --seed 42 --record /tmp/run.arch
./abate replay /tmp/run.arch       # verifies the frame hash chain + whole-archive SHA-256
```

If the last byte of the archive is tampered with, `replay` rejects it (see
threat T10 in the security threat model).

## 6. Next steps

- Checking source in all 7 languages (step-by-step guide): see `docs/user/multilang-check-guide.en.md`
- Writing more complex experiments: see `docs/user/ces-guide.md`
- All commands and exit codes: see `docs/user/cli-reference.md`
- The specification itself: see `docs/spec/ces-v1.md`

## 7. Open-Box Readiness Statement

> Measurement baseline: 2026-11-19. The source backup was extracted into a clean
> empty directory (without `.git` / `dist` / `.artifacts`) and each item below was
> executed as if by a developer receiving the project for the first time. The
> result is stated in layers, without overstatement or extrapolation.

### 7.1 Core path: genuinely ready out of the box

| Item | Measured result |
|---|---|
| Dependencies | Zero external dependencies; offline `go build` with `GOPROXY=off` succeeds on the first try |
| Build artifact | Single ~16 MB binary; no code generation or toolchain download needed |
| Version | `./abate version` → `abate 1.2.27` |
| Gates | header / style / hash / watermark / drift — all five PASS |
| Tests | Full `go test ./...` passes (73 packages) |
| Functional commands | `derive` works; `check` dispatches 7 languages by extension (`.go/.py/.js/.c/.cs/.rs/.kt`) and reports real rule violations; `run` executes a CES and emits a report + digest; invalid input yields precise parse errors carrying rule IDs |

### 7.2 Three boundaries that are NOT ready out of the box (must be aware)

1. **Prebuilt binaries in `dist/` are not in the repository.** That directory is
   excluded by `.gitignore` and can be rebuilt with `release.sh --dry-run`;
   ready-made 5-platform binaries (linux / darwin / windows × amd64 / arm64) are
   available from the [GitHub Release v1.2.27](https://github.com/xiaolibuluo/Abate/releases/tag/v1.2.27).
   Building from source with `go build` is unaffected.

2. **Advanced verification needs extra tooling — it is not clone-and-run**:
   - TLA+ formal verification requires Java + `tla2tools.jar`;
   - libclang / C++ function-level verification requires the corresponding toolchain
     (the ~776 MB build artifacts under `.artifacts/` are platform-bound, excluded by
     `.gitignore`, and reproducible via scripts);
   - full-corpus E2E regression requires the ~264 MB `testdata/` corpus.
   The core CLI and the five gates require none of this extra tooling.

3. **`check` reports info / error-level style findings on the project's own source**
   (measured: `check main.go` reports `GO-CRC-029`). This is the rule engine working
   correctly, not a fault; therefore "check passes cleanly on its own source" is not an
   acceptance criterion. Acceptance uses the frozen rule set in `scripts/lint-gate.sh`.

### 7.3 Conclusion

As a Go CLI and deterministic chaos-engineering tool, Abate **is ready out of the
box**: after a clone / extract, `go build && ./abate` just works. As a complete quality
system (TLC / clang / full-corpus E2E), the extra tooling in §7.2 must be provided.

> Full evidence for release and verification is in `docs/i18n/RELEASE-READINESS-FINAL.md` §§11–12.
