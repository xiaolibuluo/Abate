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

```bash
# Requires Go 1.26.11 (offline build via -mod=vendor)
export GOFLAGS=-mod=vendor GOPROXY=off
go build -o abate .
./abate version
```

`abate version` prints build provenance (version / schema / license / watermark):

```
abate 0.1.0
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

- Writing more complex experiments: see `docs/user/ces-guide.md`
- All commands and exit codes: see `docs/user/cli-reference.md`
- The specification itself: see `docs/spec/ces-v1.md`
