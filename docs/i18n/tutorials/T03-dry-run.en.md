<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->

> Language / 语言：[English](T03-dry-run.en.md) | [中文](T03-dry-run.md)

# T03 · Dry-run validation of a CES

## 1 Audience

Users authoring their own CES files: the goal is to parse + validate only,
without executing the experiment, to quickly iterate a CES document to a legal
state.

## 2 Prerequisites

- Build completed per T01;
- Run from the repository root;
- `ok.yaml` is shorthand for `ctl/cli/testdata/ok.yaml`.

## 3 Commands

```bash
/tmp/abate_verify run ctl/cli/testdata/ok.yaml --dry-run
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
dryRun: true
strict: false
status: VALIDATED
exitCode: 0
events: 0
observations: 0
```

**Verification points**: `dryRun: true`, `status: VALIDATED`, `exitCode: 0`,
`events: 0`, `observations: 0` — a dry run only validates the document and
does not execute the experiment, so events/observations are both 0.

## 5 Troubleshooting

| Symptom | Cause | Action |
|---|---|---|
| `exit 3` with CES-VALID-* violations | Illegal document (fail-fast; unknown fields are rejected outright) | Fix each item; quick-reference table at the end of `docs/user/ces-guide.md` |
| `status: VALIDATED` but you expected execution | Dry runs do not execute; that is expected | Drop `--dry-run` and run again (see T02) |
| YAML syntax error (`malformed.yaml` class) | Broken indentation/format | Use `ok.yaml` as a template and compare line by line |
| You want warnings treated as failures | By default warnings do not block | Add `--strict`; warning-severity violations are promoted to exit code 3 |
