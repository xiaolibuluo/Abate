<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->

> Language / 语言：[English](procurement.en.md) | [中文](procurement.md)

# Abate Procurement Manifest

English translation of `procurement.md`, the section-by-section companion
under the plan-A side-by-side naming scheme; section order and tables
correspond one-to-one; no new conclusions are added; this document does not
constitute legal advice.

> Status: procurement phase (coding forbidden; M1 implementation starts once
> everything is in place). Red lines: tool versions align with the project
> (go 1.26.11 / golangci-lint v2); **downgrading is forbidden**; any open
> source admitted must be **Go + source-level inclusion + non-GPL**; all
> downloads go through mirrors.

## 1. Primary toolchain (project-aligned versions)

| Component | Declared project version | Source (mirror) | Status |
|---|---|---|---|
| Go toolchain | go 1.26.11 | Official mirror `golang.google.cn/dl` download + SHA256 check, `GOTOOLCHAIN=local` pinned | ✅ Installed (go1.26.11 linux/amd64, checksum verified) |
| golangci-lint | v2 (CI uses `/v2/cmd/golangci-lint`) | `go install .../golangci-lint/v2/cmd/golangci-lint@v2.1.6` (mirror) | ✅ Installed v2.1.6 (existing binary built with go1.26.10, compatible with the 1.26.11 toolchain) |

## 2. Auxiliary toolchain (all Go, via mirrors, non-GPL)

| Component | Purpose | Install command | Status |
|---|---|---|---|
| gofumpt | Stricter formatting than gofmt (optional gate enhancement) | `go install mvdan.cc/gofumpt@latest` | ✅ Installed (v0.12.0, mirror) |
| goheader | Copyright-header gate (XCU-GOV) | `go install github.com/denis-tingajkin/goheader@latest` | ⚠️ Mirror 404, unavailable → **replaced by the equivalent `scripts/header-check.sh` script** (wired into the gate) |
| go-licenses | Dependency license classification (LICENSE gate enhancement, GPL detection) | `go install github.com/google/go-licenses@latest` | ✅ Installed (v1.6.0, mirror) |

## 3. Open source requiring source-level inclusion (Go / non-GPL / mirror)

| Component | License | Language | Milestone | Inclusion method | Status |
|---|---|---|---|---|---|
| yaml.v3 | MIT | Go | M1 (CES YAML parsing) | Go module with readable source (module cache) | ✅ Source in module cache (v3.0.1) |
| gVisor (isolated kernel subset) | Apache-2.0 | Go | M4 (WAS-CAP sandbox, user-specified) | `go mod download gvisor.dev/gvisor` for source; five-step inclusion into `third_party/gvisor` at M4 | ✅ Source in module cache (latest; note: gVisor@latest requires Go ≥1.27.3 to build; M4 will pin a compatible version or upgrade Go) |
| chaosblade (injection core) | Apache-2.0 | Go | M2 | Land in `third_party/chaosblade` at M2 | 📅 M2 |
| SplitMix64 | Public Domain | Go (self-implemented) | M1 (SIM-RNG) | Algorithm constants; self-implemented in sim/rand, zero external dependencies | 📅 During M1 implementation |

## Procurement completion record (2026-11-08)

- All gate scripts ready and `chmod +x`: `header-check.sh` / `style-check.sh`
  / `license-check.sh` / `coverage-gate.sh` / `lint-gate.sh`.
- `.golangci.yml` rewritten to **v2 syntax** and verified by a real run
  (0 issues); `Makefile` / `.github/workflows/ci.yml` aligned to Go 1.26 +
  golangci-lint v2.1.6.
- The seven-language official + community standards rules archived under
  `docs/reference/` (go-effective/go-uber/python-pep8/python-google/js-airbnb/
  js-ts-official/c-cert/cpp-core/rust-api/csharp-design/kotlin-conventions/
  kotlin-ktlint).
- **All gates verified by real runs** (header/style/license/gofmt/vet/
  golangci-lint/test-race/coverage/build all green).
- ⚠️ Pending user decision: see §5, the conflict between non-Go low-level
  libraries and the "must be Go" red line (affects M3/M5 only; does not block
  M1).

## 4. Seven-language official + community standards rules (reference rule base, archived in this phase)

> These rules are the authoritative sources for the in-house rule set
> (PLAN-05 §3.3). This phase fetches the official/community standard texts
> into `docs/reference/` for item-by-item mapping to AIR patterns when the
> M3/M5 rule engines land.

| Language | Official standards | Community standards | Archive path |
|---|---|---|---|
| Go | Effective Go + CodeReviewComments + go vet | Uber Go Style Guide | docs/reference/go-* |
| Python | PEP 8 / PEP 257 / PEP 484 | Google Python Style Guide | docs/reference/python-* |
| JS/TS | ECMAScript spec + TS Handbook/Do's & Don'ts | Airbnb JavaScript Style / StandardJS | docs/reference/js-* |
| C/C++ | CERT C/C++ Secure Coding | C++ Core Guidelines (isocpp) | docs/reference/c-* |
| Rust | Rust API Guidelines + Rust Reference | clippy lints | docs/reference/rust-* |
| C# | .NET Framework Design Guidelines + C# Coding Conventions | IDesign / StyleCop | docs/reference/csharp-* |
| Kotlin | Kotlin Coding Conventions | ktlint / Android Kotlin Guide | docs/reference/kotlin-* |

## 5. ⚠️ Decision required: the seven-language frontend "low-level libraries" vs the "must be Go" red line

Of the seven-language parsing low-level libraries planned in PLAN-05 §3.1,
**5 are not implemented in Go**:

| Language | Low-level library in the plan | Language | vs the "must be Go" red line |
|---|---|---|---|
| Python | ruff (rustpython-parser family) / tree-sitter-python | **Rust / C** | ❌ conflict |
| JS/TS | swc | **Rust** | ❌ conflict |
| C/C++ | libclang (LLVM) | **C++** | ❌ conflict |
| Rust | syn + quote | Rust (the language itself; no substitute) | ⚠️ necessarily non-Go |
| C# | Roslyn | **C#** | ❌ conflict |
| Kotlin | kotlin-compiler | **Kotlin/JVM** | ❌ conflict |
| Go | go/ast (standard library) | Go | ✅ no conflict |

In addition, the mutation core **PIT** planned by PLAN-00 R1 is implemented
in Java, which also conflicts with the "must be Go" red line.

**Scope of impact**: only M3 (Go frontend) / M5 (the other six language
frontends) are affected; **M1 is unaffected** (M1 is the Go kernel + CES +
CLI only, all Go standard library + yaml.v3).

**Pending user decision (does not block M1)**:
- Option A: use **Go implementations** for all seven-language frontend
  parsers (e.g. Go tree-sitter bindings, in-house Go parsers, or `go/ast`
  for Go only), fully satisfying "must be Go"; effort moves into M5.
- Option B: grant an exception for "self-implemented-language" libraries
  (Rust's syn, the Kotlin compiler), allowing them only as parse-output
  sources but **not linked into the engine**; parse results are injected into
  AIR through ports.
- Option C: at M3/M5, formally revise PLAN-05 and replace non-Go low-level
  libraries with Go-ecosystem equivalents.

This procurement phase **downloads no non-Go source** (per the red line);
the conflicts above are decided by the user before M3/M5 starts.

## 6. Auxiliary verification tools (optional, not engine dependencies)

| Component | Purpose | Note |
|---|---|---|
| PractRand | RNG statistical quality (PLAN-04 requires the 128MB tier) | C++ tool, **non-Go**, not an engine dependency; unrelated to the "must be Go" red line (test aid only). M1 first self-verifies with Go's built-in chi-square/frequency tests; PractRand is a later optional enhancement. |
