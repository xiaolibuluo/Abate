/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package grammar reads the upstream tree-sitter grammar facts of one
// language and maps them onto the frozen Abate AIR node model.
//
// The package is the mechanical answer to one question: does a self-built
// frontend really cover the constructs of its language? The upstream
// node-types.json is the official node inventory of a grammar and the
// upstream grammar.json is the official production inventory, so both are
// treated as read-only facts rather than as documentation.
//
// Three inventories are derived from those facts:
//
//   - Grammar.Nodes holds every upstream node type, and Grammar.Constructs
//     narrows the list to the concrete named constructs.
//   - Grammar.Rules holds every upstream production with the symbols it
//     references.
//   - Map resolves every concrete construct onto one AIR node kind through an
//     explicit rule table, and reports the constructs that no rule selected
//     instead of silently falling back to a generic node.
//
// The upstream data is embedded from the data directory so the package works
// offline and never depends on the read-only third_party reference tree.
//
// # Data provenance
//
// The embedded files are verbatim copies of the upstream grammar packages:
//
//	| Tag        | Upstream package          | Version | License |
//	|------------|---------------------------|---------|---------|
//	| python     | tree-sitter-python        | 0.25.0  | MIT     |
//	| javascript | tree-sitter-javascript    | 0.25.0  | MIT     |
//	| c          | tree-sitter-c             | 0.24.1  | MIT     |
//	| rust       | tree-sitter-rust          | 0.24.0  | MIT     |
//	| csharp     | tree-sitter-c-sharp       | 0.23.5  | MIT     |
//	| kotlin     | tree-sitter-kotlin        | 0.3.8   | MIT     |
//
// The copies live under third_party/_sources/ as the read-only reference
// tree; data/ holds the identical bytes so the Go toolchain can embed them.
package grammar
