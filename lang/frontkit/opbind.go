/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package frontkit

import "abate/lang/air"

// BindOperator builds one AIR node whose own attribute text is the exact
// source spelling of an operator token, so ComputeNodeID can distinguish two
// otherwise identical trees that differ only in that operator.
//
// The token text is passed as NewNodeAt's attribute, which is part of the
// NodeID payload. SemOperator is only an auxiliary annotation and never enters
// the identifier: annotations alone cannot make an operator distinguishable.
// Every binary, unary, assignment and increment node is routed through this
// entry so the seven frontends share one binding contract.
func BindOperator(kind air.NodeKind, pos, end air.Position, tok Token,
	children ...air.Node) air.Node {
	n := air.NewNodeAt(kind, pos, end, tok.Text, children...)
	_ = air.Annotate(n, air.SemOperator, tok.Text)
	return n
}
