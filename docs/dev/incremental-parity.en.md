<!--
  Copyright (c) 小黎的部落 (www.qingdujian.com) — Project: xiaolibuluo/Abate (https://github.com/xiaolibuluo/Abate)
  SPDX-License-Identifier: AGPL-3.0-or-later (subject to confirmation by the rights holder / 需权利人确认)
  Abate dual-license (AGPLv3 or commercial) bilingual documentation.
  本文件不构成法律意见 / This document does not constitute legal advice.
-->

> Language / 语言：[English](incremental-parity.en.md) | [中文](incremental-parity.md)

# Incremental Parsing Algorithm Cross-Check (tree-sitter ↔ Abate frontkit)

English translation of `incremental-parity.md`, the section-by-section
companion under the plan-A side-by-side naming scheme; section order and
tables correspond one-to-one; the comparison is a parity check only — no
competitor absolute assertions or absolute scores are added; no new
conclusions are added; this document does not constitute legal advice.

This document cross-checks upstream tree-sitter's incremental parsing
algorithm item by item against Abate's `lang/frontkit/incremental.go`
implementation, marking each item's status:

- **Implemented and equivalent**: Abate's behavior matches upstream
  semantics.
- **Implemented but simplified**: Abate has a corresponding mechanism, but
  its semantics are narrower or coarser than upstream.
- **Not implemented**: Abate has no corresponding mechanism.

## 0. Comparison material and fact sources

| Fact | Source | Read-only use |
|---|---|---|
| The incremental algorithm itself | `third_party/_sources/tree-sitter-runtime/lib/src/parser.c` | `ts_parser__reuse_node`, `ts_parser__breakdown_top_of_stack`, `ts_parser__breakdown_lookahead`, `ts_parser__can_reuse_first_leaf` |
| Tree edits and invalidation marking | `third_party/_sources/tree-sitter-runtime/lib/src/tree.c`, `subtree.c` | `ts_tree_edit`, `ts_subtree_edit`, `ts_subtree_has_changes` |
| Changed-range computation | `third_party/_sources/tree-sitter-runtime/lib/src/get_changed_ranges.c` | `ts_range_array_get_changed_ranges`, `ts_subtree_get_changed_ranges` |
| External scanner state | `third_party/_sources/tree-sitter-runtime/lib/src/parser.c` | `ts_parser__external_scanner_serialize` / `deserialize` / `scan` |
| Per-language grammar artifacts | `third_party/_sources/tree-sitter-*/src/{node-types.json,grammar.json,parser.c,scanner.c}` | Grammar facts and external-scanner contracts |
| Abate side | `lang/frontkit/incremental.go`, `lang/air/edit.go`, `lang/air/diff.go` | Incremental entry, edit application, diff computation |

The upstream C runtime (`tree-sitter-runtime/lib`) is committed to this
repository, so the cross-check in this document **depends on no external
network or inference**; every item points to a concrete function. The
generated per-language `src/parser.c` contains only parse tables and
`ts_symbol_metadata[]`, not the incremental algorithm itself.

## 1. Overview

| # | Upstream mechanism | Upstream entry | Abate counterpart | Status |
|---|---|---|---|---|
| 1 | Edit locating (byte range + line/column point) | `TSInputEdit` | `air.Edit{Path, Offset, EndOffset, Text}` | Implemented and equivalent |
| 2 | Edit application and overlap rejection | Buffer handling before `ts_subtree_edit` | `air.ApplyEdits` | Implemented and equivalent |
| 3 | Tree invalidation marking (propagating edits into subtrees) | `ts_tree_edit` → `ts_subtree_edit` | No node-level propagation | Not implemented |
| 4 | Node-level reuse | `ts_parser__reuse_node` | No node-level reuse | Not implemented |
| 5 | Stack breakdown before reuse | `ts_parser__breakdown_top_of_stack`, `ts_parser__breakdown_lookahead` | None | Not implemented |
| 6 | First-leaf reusability check | `ts_parser__can_reuse_first_leaf` | None | Not implemented |
| 7 | External scanner state consistency | `ts_parser__external_scanner_serialize` / `deserialize` | No external scanners; no state snapshot needed | Not implemented (not applicable) |
| 8 | File-level content-hash reuse | No counterpart in `ts_parser` | `reusableRoots` | Implemented but simplified |
| 9 | Invalidation on dialect/language config change | Implicitly handled by rebuilding `TSParser` | Dialect comparison in `reusableRoots` | Implemented and equivalent |
| 10 | Changed ranges | `ts_range_array_get_changed_ranges` | None | Not implemented |
| 11 | Diff computation | Beyond `ts_tree_get_changed_ranges`, plus structural comparison | `air.ComputeDiff` (NodeID multiset) | Implemented but simplified |
| 12 | Honest labeling of reuse mode | `LOG` debug logs | `lang.<tag>.incremental` annotation | Implemented and equivalent |

## 2. Item-by-item notes

### 2.1 Edit locating and application (#1, #2)

Upstream `TSInputEdit` carries both the byte range and `TSPoint`
(line/column) so ranges can be adjusted on subtrees. Abate's `air.Edit`
carries only the byte range; `air.ApplyEdits` sorts by ascending `Offset`,
applies edits in order, and explicitly rejects overlapping ranges:

- Sorting makes the result independent of caller slice order (deterministic).
- The overlap check returns `air.ErrInvalidEdit`; nothing is silently
  dropped.

**Conclusion: implemented and equivalent** (equivalent at the byte-range
level; line/column points are not modeled, but AIR's `air.Position` is
recomputed at node construction and does not depend on incremental
line/column adjustment).

### 2.2 Tree invalidation marking (#3, not implemented)

Upstream `ts_tree_edit` propagates edits down the tree: an edit before a
subtree only shifts its padding; an edit inside a subtree recomputes
size/padding and marks that subtree `has_changes`; only marked subtrees are
skipped during reuse.

Abate keeps no structured invalidation information for the old tree: the
incremental entry retains only **the old unit's root node and file
references**, with no per-node reusability metadata.

**Conclusion: not implemented.** The compensation is #8's file-level hash
reuse — if a file's content changed, the whole file is re-parsed, which is
the most conservative semantics of "every node in that file is invalid".

### 2.3 Node-level reuse and stack breakdown (#4, #5, #6, not implemented)

Upstream `ts_parser__reuse_node` does one pre-order scan of the old tree,
checking each candidate subtree in order:

1. Whether its byte position falls within the current parse position
   (`byte_offset > position` stops; `byte_offset < position` descends or
   skips);
2. Whether external scanner state is consistent
   (`ts_subtree_external_scanner_state_eq`);
3. Whether the subtree carries `has_changes` / `is_error` / `is_missing` /
   `is_fragile` marks;
4. Whether the first leaf's lexical mode under the current parse state is
   compatible (`ts_parser__can_reuse_first_leaf` compares `TSLexerMode`).

When conditions fail, upstream calls `ts_parser__breakdown_top_of_stack` to
break the reduced parent back into children and re-push them, degrading the
reuse granularity from "whole subtree" to "smaller subtrees".

Abate has no node-level reuse and no stack breakdown: the granularity is
"whole file" only.

**Conclusion: not implemented.** This is the largest gap versus upstream and
the main ceiling of Abate's incremental capability.

### 2.4 External scanner state consistency (#7, not applicable)

Upstream requires external scanners to implement
`serialize`/`deserialize`, saving scanner-private state along with tokens;
when reusing nodes it compares state consistency and skips reuse on
mismatch.

Abate's six language frontends have no external scanners: lexing is fully
deterministic and table-driven in `frontkit.Lexer`; there is no
scanner-private state to preserve across increments.

**Conclusion: not implemented (not applicable).** Semantically this equals
"state is always equal", so no reuse errors result.

### 2.5 File-level content-hash reuse (#8, implemented but simplified)

`reusableRoots` in `lang/frontkit/incremental.go` is Abate's only reuse
mechanism:

- Per file, compares `air.SourceFile.Digest()` against the content hash
  recorded for the old unit;
- On hash match, reuses the old root directly without re-parsing;
- On mismatch, the file goes into the re-parse set.

**Conclusion: implemented but simplified.** Semantically correct (unchanged
content must produce unchanged results), but the reuse granularity is the
file, not the node.

### 2.6 Invalidation on dialect change (#9, implemented and equivalent)

Upstream has no "dialect" concept: a dialect change is equivalent to
building a new parser over a different `TSLanguage`, invalidating the whole
old tree. Abate's six language frontends all carry dialects (e.g. Python
3.8–3.13, TypeScript versions); the same bytes may reduce differently under
different dialects, so a dialect change must invalidate old results.

`reusableRoots` now takes both `oldDialect` and `dialect`: if the dialects
differ, or the old unit recorded no dialect at all (missing
`lang.<tag>.dialect` annotation), nothing is reused.

> **The real defect fixed this round**: before the fix, `ParseIncremental`
> decided reuse by content hash only, ignoring dialects. When the same
> source switched between `3.11` and `3.12`, the old dialect's syntax tree
> was reused as-is and returned as the result — wrong, with no visible sign
> (the reuse-mode annotation still said `file-level-reuse`). Regression
> tests were added after the fix:
> `lang/frontkit/frontkit_incremental_test.go:TestIncrementalDialectChangeForcesFullRebuild`
> and
> `lang/pyfront/pyfront_incremental_test.go:TestPyIncrementalDialectChangeForcesRebuild`.

**Conclusion: implemented and equivalent.**

### 2.7 Changed ranges (#10, not implemented)

Upstream `ts_range_array_get_changed_ranges` walks the old and new trees
with two cursors and outputs the byte ranges where syntax structure changed,
for editors to do minimal redraws.

Abate produces no changed ranges: `air.Diff` only gives added/removed
`NodeID` multisets.

**Conclusion: not implemented.**

### 2.8 Diff computation (#11, implemented but simplified)

`air.ComputeDiff` compares the NodeID multisets of the old and new units to
get the added and removed node sets. Beyond range diffs, upstream also has
`ts_subtree_compare`'s structural comparison (to pick the better of two
candidate trees during error recovery).

**Conclusion: implemented but simplified.** The diff semantics suffice to
drive analyses like "which symbols disappeared", but are not equivalent to
upstream's range-level diffs.

### 2.9 Honest labeling of reuse mode (#12, implemented and equivalent)

Upstream prints reuse/rejection reasons with `LOG(...)`. Abate writes the
reuse mode into unit annotations:

| Annotation | Values | Meaning |
|---|---|---|
| `lang.<tag>.incremental` | `full-rebuild` | No files reused; the whole unit rebuilt |
| `lang.<tag>.incremental` | `file-level-reuse` | Some files reused by content hash |
| `lang.<tag>.reusedFiles` | decimal integer | Number of files actually reused |

**Conclusion: implemented and equivalent.** Every incremental result carries
a visible marker of "file-level reuse or full rebuild"; a full rebuild is
never disguised as node-level incrementality.

## 3. Known limitations and consequences

1. **Incrementality stops at file level.** Editing a file necessarily
   triggers a full re-parse of that file; other functions/classes in the
   same file are not reused. This is a performance ceiling, not a
   correctness issue.
2. **AIR units carry no diagnostics.** The `lang.<tag>.diagnostics`
   annotation only counts diagnostics of files re-parsed this time; the
   diagnostics of reused files are not regenerated. The diagnostic set of an
   incremental result is therefore not the full diagnostic set.
3. **No changed-range output.** Consumers cannot do minimal redraws from it;
   they can only replace wholesale.
4. **No cross-file dependency analysis for reuse.** Editing a file does not
   trigger re-parsing of other files that reference it.

## 4. Correctness verification

All assertions below are regression tests on **real editing scenarios**;
command:

```sh
GOFLAGS=-mod=vendor go test -race -count=1 ./lang/frontkit/ ./lang/pyfront/
```

| Test | Scenario | Assertion |
|---|---|---|
| `TestIncrementalEditsMatchFullParse` | Insert / delete / replace tokens | Incremental result matches full parse node for node |
| `TestIncrementalDialectChangeForcesFullRebuild` | Dialect switch | Forces a full rebuild; no old-dialect subtrees reused |
| `TestIncrementalReusesUnchangedFile` | Multiple files, one changed | Unchanged files reused by hash; reuse count correct |
| `TestIncrementalRejectsBadInput` | Empty frontend / invalid edits / nonexistent path | Returns a classified error; no panic |
| `TestPyIncrementalMatchesFullParse` | Python insert / delete / replace tokens | Incremental result matches full parse node for node |
| `TestPyIncrementalDialectChangeForcesRebuild` | Python dialect switch | Forces a rebuild |

The "node-for-node match" comparison caliber: node kinds, attributes,
position ranges, and child structure are all identical, excluding the
`lang.<tag>.incremental` / `lang.<tag>.reusedFiles` annotations (they
describe the reuse method, not the syntax result).

## 5. Follow-up optional work

Ordered by return on investment; none of these **are part of this round's
deliverables**:

1. Add "reusability" metadata to AIR nodes (edit ranges, lexical-state
   fingerprints), lowering reuse granularity from file to node.
2. Make `air.Diff` output changed ranges, aligning with
   `ts_range_array_get_changed_ranges`.
3. Make AIR units carry per-file diagnostics so an incremental result's
   diagnostic set matches the full one.
4. Add cross-file dependency analysis to the six language frontends so
   editing a referenced file triggers re-parsing of its referents.
