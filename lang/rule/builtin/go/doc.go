/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package gobuiltin declares the built-in Go rule set for the Abate checker.
//
// The rule set is pure Go table data spanning Effective Go, the Go Code Review
// Comments and a semantic subset of cmd/vet. Every rule is AIR-facing: no
// go/ast or go/token type appears in the public contract, and the shared
// adapter parses in-memory source buffers into an air.Unit for matching.
package gobuiltin
