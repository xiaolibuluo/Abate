/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package air

// Frontend is the language-neutral parsing contract.
//
// Seven language frontends implement it: Go in M3 and Python, JavaScript, C,
// Rust, C#, and Kotlin in M5. Downstream rule, check, and derivation passes
// depend only on this interface and on AIR, never on a language-specific AST.
// This file defines the contract only; implementations live in their own
// packages.
type Frontend interface {
	// Parse performs a full parse and returns the AIR unit plus structured
	// diagnostics. A syntax error is reported through diagnostics rather than
	// by panicking; error is reserved for contract or resource failures.
	Parse(unit SourceUnit) (*Unit, []Diagnostic, error)

	// ParseIncremental reparses after applying edits, reusing unchanged
	// subtrees through NodeID. The returned Diff describes what changed.
	ParseIncremental(unit SourceUnit, old *Unit, edits []Edit) (*Unit, Diff, error)
}
