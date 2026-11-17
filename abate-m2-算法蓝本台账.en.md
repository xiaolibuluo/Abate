<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->

> Language / 语言：[English](abate-m2-算法蓝本台账.en.md) | [中文](abate-m2-算法蓝本台账.md)

# Deep-integration algorithm blueprints — read-only reference baseline (Abate supply-chain compliance ledger)

> This directory is a **read-only algorithm reference** for "deep improvement
> and optimization + thorough in-house refinement": after fully understanding
> the upstream algorithm models, reimplement them from scratch in the Abate
> kernel using only the Go standard library, with deep improvements such as
> determinization, algebraization and dual compensation.
> **No import, no vendoring into `go.mod`, no compilation, no subprocess
> invocation.** Files are stored as `.go.txt` inside `_`-prefixed directories;
> the Go toolchain and all gates skip them; upstream copyright headers are
> preserved verbatim.

## Locked facts (fetched 2026-11-09; GitHub anonymous REST hit a 403 rate limit — items not obtained are marked truthfully)

| Upstream project | License | Commit snapshot | Branch |
|---|---|---|---|
| Shopify/toxiproxy | MIT | 49299d8838525126e09c85ebe0ba89940a3261e5 | main |
| go-gremlins/gremlins | Apache-2.0 | b48a4aad17eff33dc7262e0d1487b1a4d6eec322 | main |
| google/go-cmp | BSD-3-Clause (manifest marks Apache-2.0, disputed; LICENSE prevails) | 5deca1c7dab7d33c102d8d7a26794da148485563 | master |
| panjf2005/ants | MIT | Not obtained (API rate limit) | - |
| pingcap/failpoint | Apache-2.0 | 55ac33a48e3b10f21fdb30e9d29368e98bcd70fc | master |
| spf13/cobra | Apache-2.0 | adbc8813901bba65827259daa8e22ff94ec1f30e | main |

## Fetched algorithm files (file-level SHA-256 first 16 digits locked, re-verifiable)

| Upstream project | File | Algorithm purpose | Bytes | SHA-256 (first 16) |
|---|---|---|---|---|
| Shopify/toxiproxy | `toxics__latency.go.txt` | Latency/jitter toxic model | 1096 | `a85cd31f54de1d59` |
| Shopify/toxiproxy | `toxics__bandwidth.go.txt` | Token-bucket bandwidth rate limiting | 2701 | `c34d6fe8402d22c5` |
| Shopify/toxiproxy | `toxics__slicer.go.txt` | Packet slicing / bandwidth-limited scheduling | 2559 | `ff993f521b6bd720` |
| google/go-cmp | `cmp__compare.go.txt` | Deep comparison / diff-report algorithm | 22812 | `547d779d6ecb7e9a` |
| go-gremlins/gremlins | `internal__mutator__mutator.go.txt` | Go mutation-operator registration model | 4746 | `4cf48e79a490e12f` |
| spf13/cobra | `command.go.txt` | Command tree / argument parsing design | 61142 | `f79613721ad04c9b` |

## Deep-improvement directions (optimization points for the in-house rewrite, not verbatim copying)

- **Determinization**: all upstream randomness / time / concurrency ordering is
  converted to SIM-RNG seed-driven + virtual clock + total-order arbitration.
- **Algebraic unification**: faults and mutations uniformly implement
  `kernel/algebra.Operator`, with dual compensation and deterministic
  fingerprints.
- **Zero intrusion**: recording / observation subscribe read-only and do not
  modify the observed event stream.
- **Trimming**: only the core algorithm line is taken; upstream CLI / network /
  report / JVM / framework glue is removed.
- **Boundary hardening**: explicit branches and rejection are added for
  overflow / empty-input / out-of-authority paths not covered upstream.
