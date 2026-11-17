<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->

> Language / 语言：[English](05-language-frontend.en.md) | [中文](05-language-frontend.md)

# Abate L4 Language Frontend Plan — Eleven Modules in Extreme Detail

**Design-target statement (mandatory for C-class translations)**: this file is the English translation of an internal planning document. All `100%` / "pass" / "met"-style statements in the full text are **Design targets / Acceptance targets**, **not measured results**; they must not be read as established facts or capability claims. The translation corresponds section by section, does not escalate the tone, and does not delete target qualifiers.

- Document ID: PLAN-05　Status: in force　Upstream: PLAN-00
- Covered modules: LNG-AIR / LNG-GO / LNG-PY / LNG-JS / LNG-C / LNG-RS / LNG-CS / LNG-KT / LNG-RULE / LNG-CHK / LNG-MUT
- Common constraints (observed by all eleven modules):
  1. Every frontend produces AIR: language differences stop at the frontend; rules, checks, and mutation derivation target AIR only — implemented once, benefiting seven languages (design target).
  2. Source-level inclusion iron rule: third-party parsing libraries are always brought in via the five-step inclusion process; API-based access is forbidden (no subprocess calls, no gRPC, no network, no runtime dynamic loading).
  3. Frontend determinism: parsing the same source byte stream 100 times yields identical AIR serialization hashes; position information is byte-exact (acceptance target).
  4. Layered quality standards: LNG-AIR is the shared foundation of the whole engine and follows the kernel-layer standard (coverage ≥95%/85%); the other L4 modules follow the ecosystem-layer standard (≥85%/75%) (design target).
  5. Incrementality: AIR NodeIDs are stable across parses (isomorphic nodes carry the same ID), supporting incremental parsing and incremental checks (design target).

---

## Module LNG-AIR: Unified IR

### 1. Module positioning and responsibilities

Defines the Abate intermediate representation: a node model covering the superset of common constructs across seven languages, a semantic annotation system, stable node IDs, and versioned binary serialization. AIR is the convergence point of the three pipelines for rules, checks, and mutation derivation: language frontends produce only AIR, and downstream consumers consume only AIR. Its design quality directly determines the engine's ceiling (a ten-year architecture).

### 2. Package structure and file layout

```
lang/air/
  node.go        // Node interface, node registry
  kinds.go       // full NodeKind enumeration (five major categories)
  pos.go         // Position: line / column / byte offset
  sem.go         // semantic annotations: type / scope / CFG / purity
  id.go          // NodeID: content hash (isomorphism-aware)
  unit.go        // Unit/File/Package: compilation unit model
  encode.go      // versioned binary serialization
  decode.go      // deserialization + version negotiation
  roundtrip.go   // round-trip verification helpers
  air_test.go
```

### 3. Core types and interfaces (must implement)

- Full NodeKind enumeration (must be complete in the first version, five major categories):
  - Declarations (15): Package/Import/TypeDecl/ConstDecl/VarDecl/FuncDecl/MethodDecl/StructDecl/InterfaceDecl/EnumDecl/ClassDecl/TraitDecl/ModuleDecl/AnnotationDecl/UsingDecl
  - Statements (22): Block/ExprStmt/DeclStmt/If/For/While/DoWhile/RangeFor/Switch/Case/Select/Defer/GoStmt/Return/Branch/Throw/Try/Catch/Finally/With/Label/Match
  - Expressions (26): Ident/NumLit/StrLit/CharLit/BoolLit/NilLit/Binary/Unary/Call/Index/Slice/Field/Selector/CompositeLit/FuncLit/Arrow/Lambda/Cast/TypeAssert/Assign/IncDec/Star/Addr/Ternary/Spread/ChanOp/MatchExpr/Await
  - Types (15): Named/Pointer/Slice/Array/Map/Chan/FuncType/StructType/InterfaceType/GenericInst/GenericParam/Union/Optional/Tuple/Ref
  - Others (6): Comment/Doc/Attribute/Decorator/Directive/Region
- `Node` interface: `Kind() NodeKind` / `Pos() Position` / `End() Position` / `Children() []Node` / `Sem(key SemKey) any`
- `Position`: Line (1-based) / Column (1-based) / Offset (0-based byte offset) / EndOffset; file ownership via Unit
- Semantic annotation `SemInfo`: TypeRef (type reference: resolved points to a NodeID, unresolved keeps the name) / ScopeID (scope chain) / CFGEdges (control-flow edges: Reachable/Fallthrough/LoopBack/Unreachable) / Mutability (immutable/mutable/unknown) / Purity (pure/side-effect/unknown)
- `NodeID`: uint64 content hash = FNV-1a(kind + normalized own attributes + ordered child IDs); isomorphic subtrees must share the same ID (design target)
- `Unit`: Files (Path/ContentHash/LineEnding/Encoding) + Package (Name/Imports); the Unit fingerprint participates in the experiment determinism hash
- Serialization: binary format = magic (4 B) + version (2 B) + unit header + node table + string table + semantic table + SHA-256 checksum; version negotiation: the decoder declares a supported interval [min, current], and out-of-interval input is refused with both versions reported

### 4. Key algorithms and flows

1. Stable NodeID algorithm: bottom-up recursive hashing (leaf nodes hashed first, internal nodes combine ordered child IDs); incremental parsing only rehashes the affected path, at O(edit distance) rather than O(file size) cost (design target).
2. String table: deduplicate identifiers/literals/type names and reference them by index; baseline acceptance: typical projects see ≥60% reduction in serialized size (acceptance target).
3. Lazy semantic attachment: parsing produces purely syntactic AIR (zero semantic overhead); type/scope/CFG annotations are attached on demand in a second pass by the language frontend.
4. Round-trip verification: source → AIR → position-interval playback → interval coverage check (no gaps, no out-of-bounds overlaps); the per-language regeneration strategy is declared by each frontend (see common contract F3).

### 5. Dependencies and constraints

- Depends on: standard library only (hash/fnv, encoding/binary, sort).
- Depended on by: all LNG-* modules, KRN-MUT (mutation points target AIR nodes), KRN-VER (static equivalence checks).
- Forbidden: any language-specific type leaking into AIR (superset-neutral model); any I/O (only in-memory bytes are accepted); map iteration order participating in any hash (all normalized by sorting).

### 6. Test requirements

- Unit: construction/traversal/annotation read-write for every NodeKind; ID stability (isomorphic nodes share an ID, a one-byte difference must yield a different ID); serialization round-trip (lossless payloads for each kind).
- Version compatibility: older encoding decoded by newer code (interval boundaries); newer encoding decoded by older code (refusal with version reported).
- Fuzzing: random AIR trees serialize/deserialize losslessly and without panic (60-second quick profile, corpus archived).
- Performance: 1-million-node Unit serialization P99 ≤500 ms, deserialization P99 ≤700 ms (benchmark report to be archived) (design target).

### 7. Acceptance criteria (quantifiable)

1. Full NodeKind set covers common constructs of all seven languages (a seven-language construct coverage list is to be archived; uncovered constructs must state their reasons and degrade mappings).
2. ID stability / round-trip / version-negotiation cases 100% passing (acceptance target).
3. No panic in 60 seconds of fuzzing; corpus archived (acceptance target).
4. Serialization performance meets the budget (design target); coverage ≥95%/85% (design target); all gates green (acceptance target).

---

## Language frontend common contract (shared by LNG-GO/PY/JS/C/RS/CS/KT)

### F1. Frontend interface (must be implemented by all seven frontends)

- `Frontend` interface:
  - `Parse(unit SourceUnit) (*air.Unit, []Diagnostic, error)`: full parse.
  - `ParseIncremental(unit SourceUnit, old *air.Unit, edits []Edit) (*air.Unit, Diff, error)`: incremental parse (reusing unchanged subtrees by NodeID).
- `SourceUnit`: Files (Path/Content []byte) / Language / Dialect (dialect options: TS strict, C++20, Python 3.12, etc.).
- `Diagnostic`: Line/Column/Severity (error/warning)/Code (LANG-<LANG>-E/W-xxx)/Message; error cap of 100 (parsing stops when exceeded, preventing error storms).
- Determinism commitment: 100 parses of the same input yield identical AIR hashes (mandatory acceptance item for every language) (acceptance target).

### F2. Five-step inclusion process (unified flow; per-language differences are detailed in each module)

1. Vendor the source: fixed version placed under `third_party/<lib>/` (record the upstream commit hash, license review conclusion, and trimming list).
2. Secondary changes: determinization (eliminate unstable factors such as hash iteration order), structured diagnostics (errors normalized into `Diagnostic`), position alignment (unified AIR Position semantics).
3. Optimization: strip non-parsing paths (lint rules / CLI / build integration), keeping only the parser and necessary type information.
4. Deep integration: AST → AIR mapping layer (per-language mapping tables archived) + semantic annotation attachment.
5. Encapsulation: expose only the `Frontend` interface; upstream-proprietary AST types must not appear in the public API (review checklist item).

### F3. Common per-language acceptance items (every language frontend must satisfy all)

1. Baseline corpus: ≥1000 files (open-source project snapshots + self-constructed boundary cases), parse success rate ≥99.5% (C/C++ ≥99%), with failure analysis required for every remaining file (design target).
2. Determinism: 100 parses yield identical AIR hashes (acceptance target).
3. Fidelity: ≥100 files spot-checked for round-trip equivalence (Go/Python/JS/C#/Kotlin: token-level equivalence; C/C++/Rust: full position-interval coverage + semantic equivalence) (design target).
4. Throughput: ≥1 MB/s (pure parsing, single-threaded; C/C++ budget relaxed to ≥0.5 MB/s) (design target).
5. Mapping-layer coverage ≥85%/75% (design target).

---

## Module LNG-GO: Go frontend

### 1. Module positioning and responsibilities

The Go language frontend (first language of M3, the reference implementation validating the entire AIR architecture) and the foundation of Abate's self-check dogfooding. Source: `go/ast` + `go/parser` + `go/types` (standard library, zero external dependencies — the only language requiring no vendoring, and the safest).

### 2. Package structure and file layout

```
lang/gofront/
  frontend.go     // Frontend interface implementation
  mapdecl.go      // declaration mapping
  mapstmt.go      // statement mapping
  mapexpr.go      // expression mapping
  maptype.go      // type mapping
  sem.go          // go/types semantic annotation attachment
  regen.go        // source regeneration (round-trip verification)
  gofront_test.go
```

### 3. Mapping highlights (Go-specific constructs → AIR)

- `go`/`defer`/`select` → Defer/Select; channel operations (send/receive/ok forms) → ChanOp.
- Generics (TypeParam/type arguments) → GenericParam/GenericInst.
- struct tags → Attribute; `//go:` directives → Directive; iota constant blocks → ConstDecl (iota semantics preserved via annotations).
- `go/types` attachment: TypeRef resolved to NodeID; ScopeID scope chain; Purity heuristic (no side-effecting calls + no assignments → pure).

### 4. Key algorithms and flows

1. Incremental parsing: `go/parser` has no native incremental support; implemented via AIR NodeID subtree reuse (unchanged subtrees reused directly, changed paths re-parsed) (design target).
2. Regeneration strategy: token-level equivalence (Comment and Position preserved; byte equivalence not promised in v1 — decision recorded in an ADR).

### 5. Dependencies and constraints

- Depends on: standard library (`go/parser`, `go/types`), LNG-AIR.
- Forbidden: third-party Go parsing libraries; `go/ast` types leaking into the public API.

### 6. Test requirements

- Unit: item-by-item mapping tests (each Go construct → AIR node assertion); semantic annotation correctness (type resolution / scope / purity).
- Dogfooding: parse every `.go` file in the Abate repository itself with zero errors (acceptance target).
- Determinism and baseline: 100 identical hashes; 1000-file baseline corpus.

### 7. Acceptance criteria (quantifiable)

1. Item-by-item mapping tests 100% passing; Go construct coverage checklist fully checked (acceptance target).
2. Self-parsing of the Abate repository with zero errors; 100-file token-equivalence round-trip spot checks passing (acceptance target).
3. 100-run determinism; throughput ≥1 MB/s (acceptance target).
4. Coverage ≥85%/75% (design target); all gates green (acceptance target).

---

## Module LNG-PY: Python frontend

### 1. Module positioning and responsibilities

The Python language frontend. Primary source: ruff's parsing frontend (rustpython-parser lineage, implemented in Rust) — vendor only its parser crate; never bring in ruff's lint rule sources (rule self-development iron rule). Secondary source: tree-sitter-python (fault-tolerant incremental parsing, used for the degraded path on damaged files).

### 2. Five-step inclusion process — execution details

1. Vendor: fixed version of the ruff parser crate into `third_party/ruff-parser/` (MIT/Apache-2.0, reviewed as includable); tree-sitter-python into `third_party/tree-sitter-python/` (MIT).
2. Secondary changes: determinization (Rust HashMap iteration order → BTreeMap normalization); errors normalized into `Diagnostic`; positions (0-based) aligned to AIR (1-based).
3. Optimization: strip ruff's linter/formatter/CLI, keeping only parser + AST; strip tree-sitter's grammar generation artifacts, keeping only the incremental parsing API.
4. Deep integration: Rust → C ABI static library (`cargo build --release` + cbindgen header) → CGO in-process linking (source-level, no subprocess calls); Python AST → AIR mapping layer.
5. Encapsulation: `Frontend` interface; ruff/tree-sitter types must not leak.

### 3. Mapping highlights (Python-specific constructs → AIR)

- Decorators → Decorator; comprehensions (list/dict/set) → CompositeLit (annotated `comprehension`, with `for`/`if` clauses preserved as child nodes).
- `match`/`case` → Match/Pattern; `async`/`await` → Await (async via annotation); walrus `:=` → Assign (annotated `walrus`).
- f-strings → StrLit (annotated `fstring`, interpolation expressions preserved as child nodes); indented blocks → Block (annotated `indent-based`).

### 4. Key algorithms and flows

1. Dual-parser strategy: primary ruff parser (full-fidelity AST, stops on error); on failure, degrade to tree-sitter (fault-tolerant, producing partial AIR + warning diagnostics) — the degrade is explicitly marked in AIR (annotation `partial=true`), and downstream rules skip partial units (design target).
2. Semantic annotation: Python has no static type system; TypeRef is unresolved by default (name preserved), and type hints resolve to TypeRef.

### 5. Dependencies and constraints

- Depends on: `third_party/ruff-parser`, `third_party/tree-sitter-python` (CGO), LNG-AIR.
- Forbidden: ruff lint rule sources (iron rule); calling the python interpreter/pip at runtime (pure static parsing).

### 6. Test requirements

- Unit: item-by-item mapping; degrade path (damaged file → tree-sitter partial AIR + warning).
- Baseline: CPython standard library `Lib/` snapshot + self-constructed boundary cases to reach 1000 files.
- Determinism: 100 identical hashes.

### 7. Acceptance criteria (quantifiable)

1. Item-by-item mapping and degrade-path cases 100% passing (acceptance target).
2. Baseline success rate ≥99.5%; 100-file token-equivalence round-trip spot checks passing (acceptance target).
3. 100-run determinism; throughput ≥1 MB/s (acceptance target).
4. Coverage ≥85%/75% (design target); all gates green (acceptance target); zero introduction of ruff rule sources (scan report to be archived).

---

## Module LNG-JS: JS/TS frontend

### 1. Module positioning and responsibilities

The JS/TS language frontend. Source: swc (implemented in Rust, high performance) — vendor only its parser (`swc_ecma_parser` + AST); do not bring in transform/lint paths.

### 2. Five-step inclusion process — execution details

1. Vendor: fixed version of the swc parser crate into `third_party/swc/` (Apache-2.0).
2. Secondary changes: deterministic normalization; diagnostic normalization; position alignment.
3. Optimization: strip transform/minify/lint/plugin systems, keeping only parser + AST.
4. Deep integration: Rust → C ABI static library → CGO; JS/TS AST → AIR mapping.
5. Encapsulation: `Frontend` interface.

### 3. Mapping highlights

- JSX → CompositeLit (annotated `jsx`); TS type annotations → TypeRef; decorators → Decorator.
- Optional chaining `?.` and nullish coalescing `??` → Field/Binary (annotated `optional-chain`/`nullish`).
- TS generics → GenericInst/GenericParam; `enum` → EnumDecl; `interface` → InterfaceDecl.
- `async`/`await` → Await; template literals → StrLit (interpolations preserved as child nodes).

### 4. Key algorithms and flows

1. Dialect negotiation: extension (`.ts`/`.tsx`/`.js`/`.jsx`/`.mjs`) + Dialect options (strict / target ES version) → swc parsing configuration.
2. TS type preservation policy: AIR retains full type information (TypeRef), with no type erasure (AIR is for analysis, not compilation — semantics differ from `tsc emit`; recorded in an ADR).

### 5. Dependencies and constraints

- Depends on: `third_party/swc` (CGO), LNG-AIR. Forbidden: bringing in swc transform/lint paths; calling node at runtime.

### 6. Test requirements

- Unit: item-by-item mapping (≥3 cases each for JSX/TS-specific constructs); dialect negotiation matrix.
- Baseline: TypeScript compiler test suite snapshot + self-constructed boundary cases to reach 1000 files.
- Determinism: 100 identical hashes.

### 7. Acceptance criteria (quantifiable)

1. Item-by-item mapping and dialect-matrix cases 100% passing (acceptance target).
2. Baseline success rate ≥99.5%; 100-file token-equivalence round-trip spot checks passing (acceptance target).
3. 100-run determinism; throughput ≥1 MB/s (acceptance target).
4. Coverage ≥85%/75% (design target); all gates green (acceptance target).

---

## Module LNG-C: C/C++ frontend

### 1. Module positioning and responsibilities

The C/C++ language frontend. Source: libclang (LLVM) — vendor LLVM/clang sources and build a minimal libclang subset (parsing + AST access + type information only), linked in-process via CGO. The heaviest vendoring effort (the LLVM source tree is large), so the optimization step is the critical campaign.

### 2. Five-step inclusion process — execution details

1. Vendor: fixed version of LLVM/clang into `third_party/llvm/` (Apache-2.0 + LLVM Exception, reviewed as includable).
2. Secondary changes: determinization (remove non-deterministic factors such as build-artifact timestamps); structured diagnostics; position alignment.
3. Optimization (critical campaign): build configuration strips the clang driver / Sema code generation / backends, keeping only lexer/parser/AST + basic type queries; the build script produces `libclang-min.a` (target ≤50 MB; compared with full libclang at roughly 150 MB, with the comparison report to be archived) (design target).
4. Deep integration: CGO links `libclang-min`; CXCursor traversal → AIR mapping; type information (e.g. `clang_getTypeDeclAsWritten`) → TypeRef.
5. Encapsulation: `Frontend` interface; CXString/CXCursor must not leak.

### 3. Mapping highlights

- Preprocessor directives (`#include`/`#define`/`#ifdef`) → Directive (annotated `preprocessor`; macros are not expanded — analysis-oriented policy, with expansion information attached on demand).
- C++ templates → GenericInst (annotated `template`, with instantiation information attached on demand); namespaces → ModuleDecl.
- Smart pointers → TypeRef (annotated `smart-ptr`); `#pragma` → Directive.

### 4. Key algorithms and flows

1. Header policy: header search paths are supplied by SourceUnit (virtual FS mapping — `sim/fs` provides headers; real filesystem reads are forbidden).
2. Template instantiation depth cap of 16 (aligned with the clang default); exceeding it produces a warning diagnostic.

### 5. Dependencies and constraints

- Depends on: `third_party/llvm` (CGO), LNG-AIR, `sim/fs` (header virtualization via a port interface).
- Forbidden: real filesystem reads; calling the clang binary; bringing in LLVM backend code paths.

### 6. Test requirements

- Unit: item-by-item mapping (≥3 cases each for C/C++-specific constructs); virtual-FS header resolution cases.
- Baseline: LLVM test suite sampling + self-constructed boundary cases to reach 1000 files.
- Determinism: 100 identical hashes.

### 7. Acceptance criteria (quantifiable)

1. Item-by-item mapping and virtual-header cases 100% passing (acceptance target).
2. Baseline success rate ≥99%; 100-file full position-interval coverage spot checks passing (acceptance target).
3. 100-run determinism; throughput ≥0.5 MB/s (relaxed C/C++ budget) (acceptance target).
4. `libclang-min.a` ≤50 MB (build report to be archived) (design target); coverage ≥85%/75% (design target); all gates green (acceptance target).

---

## Module LNG-RS: Rust frontend

### 1. Module positioning and responsibilities

The Rust language frontend. Source: syn (a pure Rust parsing library, containing no clippy checking logic — rules are self-developed from the official specification, iron rule). quote is for code generation and is not needed by the frontend, so it is not included.

### 2. Five-step inclusion process — execution details

1. Vendor: fixed version of syn into `third_party/syn/` (MIT/Apache-2.0 dual license).
2. Secondary changes: deterministic normalization; syn `Error` → `Diagnostic`.
3. Optimization: keep only the parsing path (trim syn's full feature set as needed: retain item/expr/stmt/pat/ty, strip proc-macro support paths).
4. Deep integration: Rust → C ABI static library → CGO; syn AST → AIR mapping.
5. Encapsulation: `Frontend` interface.

### 3. Mapping highlights

- Macro invocations (including `macro_rules!` declarations) → Call (annotated `macro`); `#[derive]` attributes → Attribute.
- Lifetimes → TypeRef (annotated `lifetime`); `trait` → TraitDecl; `impl` blocks → MethodDecl (annotated `impl`).
- Pattern matching → Match/Pattern; the `?` operator → annotation `try-op`; `async fn` → FuncDecl (annotated `async`); `unsafe` blocks → Block (annotated `unsafe`).

### 4. Dependencies and constraints

- Depends on: `third_party/syn` (CGO), LNG-AIR. Forbidden: bringing in clippy checking logic; calling cargo/rustc at runtime.

### 5. Test requirements

- Unit: item-by-item mapping (≥3 cases each for Rust-specific constructs).
- Baseline: Rust standard library source sampling + self-constructed boundary cases to reach 1000 files.
- Determinism: 100 identical hashes.

### 6. Acceptance criteria (quantifiable)

1. Item-by-item mapping cases 100% passing (acceptance target).
2. Baseline success rate ≥99.5%; 100-file full position-interval coverage spot checks passing (acceptance target).
3. 100-run determinism; throughput ≥1 MB/s (acceptance target).
4. Coverage ≥85%/75% (design target); all gates green (acceptance target); zero introduction of clippy logic (scan report to be archived).

---

## Module LNG-CS: C# frontend

### 1. Module positioning and responsibilities

The C# language frontend. Source: Roslyn (Microsoft official, MIT) — vendor Roslyn sources, compile to a C ABI static library with NativeAOT (in-process, no .NET runtime subprocess calls), linked via CGO. Borrow only Roslyn's parsing and semantic analysis capabilities; do not bring in third-party analyzer rules (iron rule).

### 2. Five-step inclusion process — execution details

1. Vendor: fixed version of Roslyn into `third_party/roslyn/` (MIT).
2. Secondary changes: deterministic normalization; `Microsoft.CodeAnalysis.Diagnostic` → AIR `Diagnostic`.
3. Optimization: strip IDE features/analyzers/code fixes, keeping only SyntaxTree parsing + a semantic model subset; NativeAOT publishes a single-file static library `libroslyn-abate.a` (target ≤80 MB) (design target).
4. Deep integration: C ABI wrapper layer (`extern C` exports for parse / semantic attachment entry points) → CGO; Roslyn SyntaxNode → AIR mapping.
5. Encapsulation: `Frontend` interface; `Microsoft.CodeAnalysis` types must not leak.

### 3. Mapping highlights

- LINQ query expressions → Call (annotated `linq`); attributes → Attribute; events → FieldDecl (annotated `event`).
- Properties/indexers → FuncDecl (annotated `property`/`indexer`); nullable reference types → Optional.
- `using` directives/statements → UsingDecl/Block (annotated `using-stmt`); `async`/`await` → Await; `record` → ClassDecl (annotated `record`).

### 4. Dependencies and constraints

- Depends on: `third_party/roslyn` (CGO), LNG-AIR. Forbidden: bringing in third-party analyzer rules; calling dotnet as a subprocess at runtime.

### 5. Test requirements

- Unit: item-by-item mapping (≥3 cases each for C#-specific constructs).
- Baseline: dotnet/runtime source sampling + self-constructed boundary cases to reach 1000 files.
- Determinism: 100 identical hashes.

### 6. Acceptance criteria (quantifiable)

1. Item-by-item mapping cases 100% passing (acceptance target).
2. Baseline success rate ≥99.5%; 100-file token-equivalence round-trip spot checks passing (acceptance target).
3. 100-run determinism; throughput ≥1 MB/s (acceptance target).
4. `libroslyn-abate.a` ≤80 MB (build report to be archived) (design target); coverage ≥85%/75% (design target); all gates green (acceptance target).

---

## Module LNG-KT: Kotlin frontend

### 1. Module positioning and responsibilities

The Kotlin language frontend. Source: kotlin-compiler (JetBrains official, Apache-2.0) — vendor its parsing frontend (lexer/parser/PSI) sources, compile to a native static library with Kotlin/Native (C interop), linked via CGO. Take only the AST; rules are self-developed (iron rule).

### 2. Five-step inclusion process — execution details

1. Vendor: fixed version of kotlin-compiler into `third_party/kotlin-compiler/` (Apache-2.0).
2. Secondary changes: deterministic normalization; compiler diagnostics → `Diagnostic`.
3. Optimization: strip backends (JVM/JS/Native code generation) and IDE integration, keeping only the frontend (lexer/parser/PSI); Kotlin/Native produces `libkotlin-front.a` (target ≤60 MB) (design target).
4. Deep integration: C ABI wrapper layer → CGO; Kotlin PSI → AIR mapping.
5. Encapsulation: `Frontend` interface; JetBrains-proprietary types must not leak.

### 3. Mapping highlights

- Coroutines (`suspend fun`) → FuncDecl (annotated `suspend`); sealed classes → ClassDecl (annotated `sealed`).
- Companion objects → ClassDecl (annotated `companion`); smart casts → Cast (annotated `smart`).
- Extension functions → FuncDecl (annotated `extension`, with the receiver recorded as the first parameter).
- `when` expressions → Match/MatchExpr; data classes → ClassDecl (annotated `data`).
- Null safety (`?.` and `!!`) → Field/Binary (annotated `null-safe`/`assert-not-null`).

### 4. Dependencies and constraints

- Depends on: `third_party/kotlin-compiler` (CGO), LNG-AIR. Forbidden: JVM runtime dependencies; calling kotlinc as a subprocess.

### 5. Test requirements

- Unit: item-by-item mapping (≥3 cases each for Kotlin-specific constructs).
- Baseline: Kotlin standard library source sampling + self-constructed boundary cases to reach 1000 files.
- Determinism: 100 identical hashes.

### 6. Acceptance criteria (quantifiable)

1. Item-by-item mapping cases 100% passing (acceptance target).
2. Baseline success rate ≥99.5%; 100-file token-equivalence round-trip spot checks passing (acceptance target).
3. 100-run determinism; throughput ≥1 MB/s (acceptance target).
4. `libkotlin-front.a` ≤60 MB (build report to be archived) (design target); coverage ≥85%/75% (design target); all gates green (acceptance target).

---

## Module LNG-RULE: Rule engine

### 1. Module positioning and responsibilities

A YAML rule DSL + AIR structural pattern matching engine: loading, matching, and reporting of official-specification rule sets for seven languages. Every rule serves both checks (LNG-CHK) and mutation derivation (LNG-MUT) — dual-purpose rules.

### 2. Package structure and file layout

```
lang/rule/
  rule.go        // Rule type: full rule model fields
  dsl.go         // YAML DSL parsing and validation
  pattern.go     // structural pattern compiler (pattern → matcher)
  matcher.go     // matcher: AIR structural matching
  predicate.go   // predicate expressions (where clauses)
  registry.go    // rule registry (namespaced IDs)
  builtin/       // official-specification rule sets (built in)
    go/         // ≥100 rules (Effective Go + CodeReviewComments + vet semantics)
    python/     // ≥60 rules (PEP 8 + PEP 484)
    js/         // ≥60 rules (ECMAScript + TS Handbook)
    c/          // ≥60 rules (CERT + official coding standards)
    rust/       // ≥60 rules (Rust API Guidelines)
    csharp/     // ≥60 rules (.NET design guidelines)
    kotlin/     // ≥60 rules (Kotlin coding conventions)
  rule_test.go
```

### 3. Core types and interfaces (must implement)

- Full `Rule` fields:
  - `id`: namespace (GO-EG-001 / GO-CRC-014 / PY-PEP8-003; format LANG-source-serial).
  - `language` / `severity` (error/warning/info) / `category` (naming/style/error-prone/performance/concurrency/boundary).
  - `pattern`: structural pattern (node kind + subpatterns + metavariable `$x` + wildcard `$_` + deep match `$..`).
  - `where`: predicate list (capture type constraints, value matching, scope constraints, semantic annotation checks).
  - `message`: violation message template (referencing captures).
  - `fix`: fix suggestion (description + optional automated fix pattern).
  - `reference`: official-specification source (mandatory field — rules without a source are refused at load time).
  - `mutation`: mutation derivation declaration (optional: M-family operators derivable from this rule, with derivation types and parameters).
- Pattern compiler: pattern AST → matcher (compiled once, reused across files).
- Matcher: `Match(node air.Node) []Capture`; supports wildcards (`$_` any node), captures (`$x` binding), deep matching (`$..` descendants), negative patterns (`!pattern`), and repetition (`$x*` zero or more).
- Predicates: type constraints (capture is a literal/identifier), value matching (regex/numeric ranges), semantic annotation checks (purity/mutability/type-resolution state).

### 4. Key algorithms and flows

1. Matching pipeline: file → AIR → filter rule sets by language → parallel per-rule matcher execution (read-only, lock-free) → capture predicate filtering → violation list.
2. Performance: matchers compiled once; per-rule per-file matching O(AIR node count); full rule set (100 rules) per file P99 ≤50 ms (baseline acceptance item) (design target).
3. Four rule-load checks: DSL schema validation + non-empty `reference` + unique `id` + compilable `pattern` — all four must pass before registration (bad rules refused at load time, never carried into runtime).

### 5. Dependencies and constraints

- Depends on: LNG-AIR, `third_party/yaml` (source-level).
- Depended on by: LNG-CHK (checks), LNG-MUT (derivation).
- Forbidden: language-specific logic embedded in rules (rules target AIR only); regex abuse (structural patterns preferred — regex only for literal value matching).

### 6. Test requirements

- Unit: ≥3 cases for each DSL feature (wildcard/capture/deep/negative/repetition); cases for each predicate type.
- Rule-set tests: positive and negative code fixtures for every built-in rule (≥1 case each; ~460 rules across seven languages → ≥920 fixtures) (design target).
- Performance: full rule set per file P99 ≤50 ms baseline (design target).
- Fuzzing: random patterns do not panic.

### 7. Acceptance criteria (quantifiable)

1. All DSL feature cases pass; all four load-check cases pass (acceptance target).
2. Built-in rule sets: Go ≥100, each of the other six languages ≥60 (≥460 total), each with positive/negative fixtures and an official-specification source (scan check: source non-empty rate 100%) (acceptance target).
3. Full rule-set performance baseline met (report to be archived) (design target).
4. Coverage ≥85%/75% (design target); all gates green (acceptance target).

---

## Module LNG-CHK: Checker

### 1. Module positioning and responsibilities

A self-developed language code checker: orchestration pipeline (file → frontend → AIR → rule engine → violations → aggregation → report). CLI entry point `abate check`. Dogfooding foundation: Abate checks its own repository with its own checker.

### 2. Package structure and file layout

```
lang/check/
  check.go       // Check main entry: pipeline orchestration
  aggregate.go   // violation aggregation (dedupe / sort / severity statistics)
  sarif.go       // SARIF 2.1.02 output
  console.go     // human-readable console output
  exit.go         // exit code convention
  check_test.go
```

### 3. Core types and interfaces (must implement)

- `Check(paths []string, opts Options) (Result, error)`: main entry point.
- `Options`: Languages (language filter) / Rules (rule filter) / Severity threshold / Format (sarif/console/json).
- `Result`: Violations (rule id / file / line / column / message / captures) + Summary (per-severity / per-rule statistics).
- Exit-code convention: 0 = zero violations; 1 = violations present (gate FAIL semantics); 2 = internal error (strictly distinguished from violations).
- SARIF output: SARIF 2.1.02 (`runs[].results[].ruleId/location/message`), passing the official validator (acceptance item).

### 4. Key algorithms and flows

1. Pipeline: path expansion → language detection (extension + content sniffing) → parallel parsing (file-level independence, results deterministically sorted by path lexicographic order) → rule matching → aggregation (dedupe: same rule at same position; sort: file lexicographic order → line → column → rule id).
2. Incremental checking: combine AIR NodeIDs to re-check only changed subtrees and affected rules (v1 performs full checks; incrementality is listed as an M5 enhancement — recorded in an ADR).

### 5. Dependencies and constraints

- Depends on: LNG-* (frontends / rules), RPT-SARIF (format reuse).
- Forbidden: language-specific logic embedded in the checker (pure orchestration layer).

### 6. Test requirements

- Unit: aggregation dedupe / sort; three exit-code cases.
- End-to-end: fixture project (with known violations) → check → violation list exactly matching expectations.
- Dogfooding (acceptance): Abate repository self-check with zero violations; parity check against golangci-lint (Go) with a difference report explaining each item (rule semantic differences / rule-set differences).

### 7. Acceptance criteria (quantifiable)

1. End-to-end fixture cases 100% passing; exit-code convention cases passing (acceptance target).
2. SARIF passes the official validator (acceptance target).
3. Dogfooding: Abate repository self-check with zero violations; golangci-lint parity difference report archived with item-by-item explanations (acceptance target).
4. Coverage ≥85%/75% (design target); all gates green (acceptance target).

---

## Module LNG-MUT: Operator derivation

### 1. Module positioning and responsibilities

The bridge from rules to M-family mutation operators: automatically deriving mutation operators from rule `mutation` declarations. Rules are mutation points: boundary rules derive boundary-perturbation operators, type rules derive type-assertion perturbation operators, and concurrency rules derive concurrency-perturbation operators.

### 2. Package structure and file layout

```
lang/derive/
  derive.go      // Derive main entry: rule set → operator set
  table.go       // derivation table: rule category → operator type mapping
  validate.go    // derivation validation: operator usability in mutation experiments
  derive_test.go
```

### 3. Core types and interfaces (must implement)

- `Derive(rules []rule.Rule) ([]MutationSpec, error)`: main entry point.
- `MutationSpec`: SourceRuleID / OperatorID (derived operator id: `m-derived-<rule id>`) / TargetKind (AIR node kind) / Transform (perturbation semantics: operator replacement / boundary shift / constant perturbation / condition negation / call deletion).
- Derivation table (five categories in the first version):
  - Boundary rules (comparison operators, range checks) → boundary perturbation (`<` ↔ `<=`, `>` ↔ `>=`).
  - Constant rules (magic numbers, repeated strings) → constant perturbation (+1/-1/zero/empty).
  - Condition rules (compound conditions, short-circuit evaluation) → condition negation / branch deletion.
  - Type rules (type assertions, casts) → type perturbation (assertion target perturbation).
  - Call rules (ignored errors, return-value checks) → call deletion / return-value perturbation.
- `validate`: every derived operator must pass a usability test in KRN-MUT mutation experiments (mutant generation → execution → validity verdict) (acceptance target).

### 4. Key algorithms and flows

1. Derivation pipeline: filter rule set (those with `mutation` declarations) → derivation table lookup → MutationSpec generation → operator registration (into the KRN-MUT registry).
2. Derivation determinism: the same rule set derives the same operator set (sorted output, hash-verifiable) (design target).

### 5. Dependencies and constraints

- Depends on: LNG-RULE (rule model), KRN-MUT (operator registry, via a port interface).
- Forbidden: derivation logic hardcoding specific languages (the derivation table depends only on rule categories and AIR node kinds).

### 6. Test requirements

- Unit: cases for each of the five derivation categories; determinism (double derivation of the same rule set, identical operator-set hashes).
- End-to-end: ≥20 rules derive operators, each run through a mutation experiment in KRN-MUT producing a validity verdict (acceptance target).

### 7. Acceptance criteria (quantifiable)

1. Five-category derivation table cases pass; derivation determinism cases pass (acceptance target).
2. ≥20 rules successfully derive operators usable in mutation experiments (M3 acceptance item).
3. Coverage ≥85%/75% (design target); all gates green (acceptance target).

---

## Appendix: Overall L4 merge acceptance (after all eleven modules are in place)

1. End-to-end dogfooding: the Abate repository (Go) produces zero violations under the self-developed checker; ≥20 rules derive mutation operators and run mutation experiments against Abate's own code producing a valid report (complete dogfooding loop) (acceptance target).
2. Seven-language baseline: per-language 1000-file parse success rate ≥99.5% (C/C++ ≥99%), 100-run determinism, throughput meeting budgets (baseline reports to be archived) (acceptance target).
3. AIR serialization performance and version-compatibility cases pass; the seven-language construct coverage checklist is fully checked (acceptance target).
4. An aggregate coverage report for all eleven modules is archived; the five-section gate chain + baseline regression are all green (acceptance target).
