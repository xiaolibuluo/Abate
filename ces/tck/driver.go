/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package tck

// Driver is the adaptation point for a third-party CES implementation under test.
//
// The kit never inspects the implementation's internal types; it only exchanges
// neutral byte streams and Rule-ID sets. This keeps the kit portable to any
// language or engine that speaks the CES specification.
type Driver interface {
	// Parse decodes a CES document, normalizes it and reports acceptance.
	//
	// On rejection (accept=false) RuleID carries the deterministic CES-PARSE
	// rule identifier and Error carries a human-readable, byte-stable message.
	Parse(data []byte, format string) (ParseResult, error)

	// Validate checks an already-decoded CES document and returns every
	// violation as a neutral Violation. An empty slice means compliance.
	Validate(data []byte) ([]Violation, error)
}
