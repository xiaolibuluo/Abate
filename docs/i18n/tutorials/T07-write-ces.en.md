<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->

> Language / 语言：[English](T07-write-ces.en.md) | [中文](T07-write-ces.md)

# T07 · Writing a CES file from scratch

## 1 Audience

Users writing their first CES experiment: the goal is to start from the
repository's ready-made 36-line sample, understand each field, then copy,
modify and pass a dry run.

## 2 Prerequisites

- Build completed per T01;
- Run from the repository root;
- `ok.yaml` is shorthand for `ctl/cli/testdata/ok.yaml` (the full text below
  is exactly those 36 lines of the real file; **field names follow the real
  file — nothing is invented**).

## 3 Commands

Read the real sample in full, then copy and adapt:

```bash
cat ctl/cli/testdata/ok.yaml
cp ctl/cli/testdata/ok.yaml /tmp/my-exp.yaml
/tmp/abate_verify run /tmp/my-exp.yaml --dry-run
/tmp/abate_verify run /tmp/my-exp.yaml
```

## 4 Expected output

`cat ctl/cli/testdata/ok.yaml` (the real file in full, 36 lines):

```yaml
apiVersion: ces.tencent.com/v1
kind: Experiment
metadata:
  name: cli-e2e-ok
  namespace: default
spec:
  target:
    kind: simulator
    selector:
      node: n1
  steadyState:
    - probeRef: health
      expectation:
        type: value
        value: "200"
      tolerance:
        type: absolute
        value: 1
      holdFor: 1s
  method:
    expression: "ε"
  probes:
    - id: health
      kind: http
      spec:
        url: http://sim/health
        method: GET
        expectStatus: 200
      interval: 1s
  rollback:
    strategy: auto
    verifySteadyState: true
  schedule:
    once:
      at: 5s
  seed: 12345
```

`--dry-run` report head (same measured anchors as T03):

```
status: VALIDATED
exitCode: 0
events: 0
observations: 0
```

Key lines of the full run (same measured anchors as T02):

```
status: COMPLETED
exitCode: 0
  health | HELD | held span 3000000000 satisfies holdFor 1000000000
```

**Verification points**: dry run `VALIDATED`; full run `COMPLETED` with the
verdict line HELD.

**Field-by-field quick read**: `apiVersion/kind` envelope rules (see
`docs/user/ces-guide.md`); `target.selector.node` picks a node inside the
simulator; `steadyState` references probe `health` and requires its value to
stay within tolerance 1 for 1s; `method.expression` uses the identity
perturbation `ε` (no real fault injected in this sample — the minimal legal
experiment); `probes[].id` is unique within the document; `rollback.strategy:
auto`; `schedule` is one-of-three (here `once`); `seed` is the document seed
(12345, overridable via `--seed`).

## 5 Troubleshooting

| Symptom | Rule | Fix |
|---|---|---|
| Expression syntax error | CES-VALID-METHOD-001 | Check operators/parentheses/symbols; the identity perturbation is `ε` |
| Non-invertible operator without compensation | CES-VALID-METHOD-002 | Add `compensation` or `invertible: true` |
| Nesting too deep | CES-VALID-METHOD-003 | Depth ≤ 64 |
| Multiple schedule modes present | CES-VALID-SCHEDULE-001 | Keep exactly one of `continuous/once/cron` |
| `external` target | CES-VALID-TARGET-003 | Use `simulator/source` instead |
| Renamed a field, got "unknown field" | The envelope is fail-fast; unknown fields are rejected outright | Compare field by field against the 36 real lines above |
