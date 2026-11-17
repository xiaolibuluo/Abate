/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package tck

// ParseResult is the neutral outcome of Driver.Parse.
type ParseResult struct {
	// OK is true when the document parsed and normalized successfully.
	OK bool
	// RuleID is the CES-PARSE rule identifier when OK is false.
	RuleID string
	// Error is a byte-stable diagnostic message when OK is false.
	Error string
}
