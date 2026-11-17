<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->

> Language / 语言：[English](T06-exit-codes.en.md) | [中文](T06-exit-codes.md)

# T06 · Exit codes, verified one by one

## 1 Audience

Users integrating Abate into CI/scripts: the goal is to measure the 0/1/3
states of the four-state exit-code contract (0/1/2/3) one by one so scripts
can gate reliably.

## 2 Prerequisites

- Build completed per T01;
- Run from the repository root;
- The four test documents are real repository files:
  `ctl/cli/testdata/{ok,violated,invalid,malformed}.yaml`;
- Exit code 2 (internal error: scheduler or report output failure) is a
  declared state that the four samples do not trigger; this tutorial does not
  construct internal faults and states that honestly.

## 3 Commands

```bash
/tmp/abate_verify run ctl/cli/testdata/ok.yaml        > /dev/null 2>&1; echo $?
/tmp/abate_verify run ctl/cli/testdata/violated.yaml  > /dev/null 2>&1; echo $?
/tmp/abate_verify run ctl/cli/testdata/invalid.yaml   > /dev/null 2>&1; echo $?
/tmp/abate_verify run ctl/cli/testdata/malformed.yaml > /dev/null 2>&1; echo $?
/tmp/abate_verify                                     > /dev/null 2>&1; echo $?
/tmp/abate_verify version                             > /dev/null 2>&1; echo $?
```

## 4 Expected output

The six `echo $?` lines print, in order (measured):

```
0
1
3
3
3
0
```

**Verification points** (mapping one-to-one to the declared PLAN-07 §4):

| Input | Exit code | Meaning |
|---|---|---|
| `ok.yaml` | 0 | legal CES and steady state HELD |
| `violated.yaml` | 1 | business failure: steady state VIOLATED during execution |
| `invalid.yaml` | 3 | usage error: invalid CES document (CES-VALID violation) |
| `malformed.yaml` | 3 | usage error: YAML cannot be parsed |
| no arguments | 3 | usage error: missing required argument |
| `version` | 0 | success |

## 5 Troubleshooting

| Symptom | Cause | Action |
|---|---|---|
| A line differs from this table | You ran a stale binary | Repeat the T01 build; if it still differs, keep the evidence and report to the maintainers |
| `violated.yaml` trips your CI | Exit 1 is the expected teaching state for **business failure** | In scripts, distinguish 1 (steady state broken) from 3 (illegal document/arguments) |
| Want to reproduce exit code 2 | Requires a scheduler/report internal error; samples do not trigger it | Not recommended to construct; read the declared help text as-is |
| `echo $?` lines come out out of order | Other commands ran between the command and `echo $?` | `$?` only reflects the most recent command; run line by line |
