<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->

> Language / 语言：[English](T02-first-run.en.md) | [中文](T02-first-run.md)

# T02 · Your first chaos experiment

## 1 Audience

Users who just finished building (see T01): the goal is to run your first
experiment with the repository's sample CES and learn to read the report head.

## 2 Prerequisites

- Build completed per T01 (this tutorial refers to the binary as
  `/tmp/abate_verify`);
- Run from the repository root;
- In this tutorial, `ok.yaml` is shorthand for `ctl/cli/testdata/ok.yaml`
  (a real 36-line file in the repository).

## 3 Commands

```bash
/tmp/abate_verify run ctl/cli/testdata/ok.yaml
```

## 4 Expected output

Report head (measured, verbatim):

```
abate run report
schema: abate.report.v1
apiVersion: ces.tencent.com/v1
name: cli-e2e-ok
namespace: default
seed: 12345
dryRun: false
strict: false
status: COMPLETED
exitCode: 0
events: 6
observations: 4
end: 10000000000
runDigest: 6444a24ed20f880d4b3b9c0815ecfc429afcc7b3793619ba5e5cc733058043ef
streamDigest: 20220908d7fcaeddb5eb3141230b7908e428b370bf4479362c6d0add98570e1d
violations: 0
verdicts: 1
  health | HELD | held span 3000000000 satisfies holdFor 1000000000
```

**Verification points**: `status: COMPLETED`, `exitCode: 0`,
`violations: 0`; the verdict line `health | HELD | ...` means the steady-state
assertion held within holdFor (`HELD` is a status enum value and is kept as-is).
`runDigest` starts with `6444a24ed2` — for the same command and file, the
digest is byte-for-byte reproducible.

## 5 Troubleshooting

| Symptom | Cause | Action |
|---|---|---|
| `exit 3` / usage error | Wrong path (not run from the repository root) | `cd` to the repository root, or use the absolute path `ctl/cli/testdata/ok.yaml` |
| `status: VIOLATED` | You ran `violated.yaml` instead of `ok.yaml` (exit 1) | Check the file name; `violated.yaml` is a teaching fixture whose steady state deliberately fails |
| Report fields differ from this tutorial | You ran a stale binary | Repeat the T01 build steps |
| You want machine-readable output | console is the default format | Add `--format json` (see T04) |
