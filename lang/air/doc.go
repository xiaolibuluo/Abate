/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package air defines the versioned Abate Intermediate Representation.
//
// AIR is a language-neutral tree model shared by every language frontend and
// every downstream analysis pipeline. Nodes carry a syntactic kind, an exact
// source interval, ordered children, and optional semantic annotations. Node
// identifiers are content hashes, so structurally equal subtrees hash equally.
//
// Serialization is byte-stable and self-validating: an archive starts with a
// magic header and an explicit writer version range, stores length-prefixed
// sections in a fixed order, and ends with a SHA-256 checksum of every preceding
// byte. New fields are appended as trailing, length-prefixed sections so that a
// decoder can skip what it does not understand. The decoder declares the range
// of schema versions it accepts and rejects anything outside that range while
// reporting both the writer and reader versions.
//
// The package performs no I/O and imports no language-specific syntax types.
package air
