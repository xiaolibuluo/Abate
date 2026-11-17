/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package rule

import "abate/lang/air"

// Rule is the AIR-facing public contract implemented by every compiled rule.
//
// Check accepts only an AIR unit, so no language-specific syntax type can ever
// enter the rule contract. Implementations are immutable and safe for
// concurrent read-only matching.
type Rule interface {
	// Spec returns the declaration the rule was compiled from.
	Spec() Spec
	// Check reports the findings produced for one AIR compilation unit.
	Check(unit *air.Unit) []Finding
}
