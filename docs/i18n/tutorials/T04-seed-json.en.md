<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->

> Language / 语言：[English](T04-seed-json.en.md) | [中文](T04-seed-json.md)

# T04 · Fixed seed + JSON output

## 1 Audience

Users who want to plug experiment results into a pipeline or perform
byte-for-byte re-verification: the goal is to master `--seed` and
`--format json`, and to internalize the caliber that "determinism claims
require a fixed seed".

## 2 Prerequisites

- Build completed per T01;
- Run from the repository root;
- `ok.yaml` is shorthand for `ctl/cli/testdata/ok.yaml`.

## 3 Commands

```bash
/tmp/abate_verify run ctl/cli/testdata/ok.yaml --seed 42 --format json > /tmp/report-42.json
/tmp/abate_verify run ctl/cli/testdata/ok.yaml --seed 42 --format json > /tmp/report-42-again.json
/tmp/abate_verify verify /tmp/report-42.json /tmp/report-42-again.json
```

## 4 Expected output

Key fields of the JSON report (measured):

```json
{
  "seed": 42,
  "status": "COMPLETED",
  "runDigest": "32dbea3d9acc1c154a49a869dda11243db1a2e24cbeb64d5248619b3cb318339"
}
```

(Excerpt: the full JSON report contains more fields; only the three keys this
tutorial anchors on are listed.)

`verify` exits `0` (the two reports are byte-for-byte identical).
**Verification point**: with `--seed 42` fixed, `runDigest` is the same value
on every run (starting with `32dbea3d9`).

**Honest caliber**: determinism claims **must fix `--seed`**. A seed override
changes the digest — the same document without `--seed` uses the document seed
12345 and a digest starting with `6444a24ed2` (see T02); with `--seed 42` it
is `32dbea3d9…`. The two digests differ precisely because the seed takes
effect, not because of non-determinism.

## 5 Troubleshooting

| Symptom | Cause | Action |
|---|---|---|
| `verify` exits 1 | The reports differ (`verify` prints the first differing offset and SHA-256 digests) | Ensure both commands use the exact same `--seed` and file |
| Digest differs from this tutorial | A different seed or a different CES file was used | Copy the section 3 commands exactly |
| JSON parsing fails | Non-JSON output got mixed into the pipe | Confirm `--format json`; the report goes to stdout with no extra prefixes |
| You want an archived copy | In-memory reports are easy to lose | Add `--out /tmp/report.json` to also write the file (same bytes as stdout) |
