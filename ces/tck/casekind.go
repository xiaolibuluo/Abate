/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package tck

// CaseKind enumerates the three TCK case families.
type CaseKind string

// Supported case families.
const (
	// KindParse exercises Driver.Parse (accept/reject + parse RuleID).
	KindParse CaseKind = "parse"
	// KindValidate exercises Driver.Validate (accept/reject + validate RuleIDs).
	KindValidate CaseKind = "validate"
	// KindRoundtrip exercises lossless YAML/JSON round-tripping.
	KindRoundtrip CaseKind = "roundtrip"
)
