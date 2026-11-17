/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package gofront implements the Abate Go language frontend.
//
// It maps go/parser and go/ast syntax trees onto the language-neutral AIR and
// attaches go/types annotations from a purely in-memory type checker. Only the
// Go standard library go/parser, go/ast, go/token, go/types, and go/scanner
// packages are used; no third-party parser, CGO bridge, or subprocess is
// involved. No go/ast type appears in the exported API: callers only see the
// air.Frontend contract and AIR values.
//
// The frontend is Abate's dogfood foundation: it must parse every Go file in
// this repository without a syntax error and produce byte-stable AIR.
package gofront
