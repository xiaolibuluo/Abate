/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package coverage

import "abate/lang/air"

// Entry is the coverage verdict for one upstream construct.
//
// The row keeps the lowering decision and the measurement apart, so a reader
// can tell what the frontend intends to do from what it actually did.
type Entry struct {
	// Node is the upstream node type name.
	Node string
	// Kind is the AIR node kind the construct is lowered to.
	Kind air.NodeKind
	// Status classifies the fidelity of the lowering.
	Status string
	// Reason documents the lowering decision.
	Reason string
	// Rule is the name of the mapping rule that selected the construct.
	Rule string
	// Observed reports whether a real parse produced the mapped kind.
	Observed bool
	// Count is how often the mapped kind appeared in the real parse.
	Count int
}
