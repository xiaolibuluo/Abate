/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package pyfront

import (
	"abate/lang/air"
	"abate/lang/frontkit"
)

// This file holds the with-statement reductions: the comma-separated context
// manager list, the optional as targets and the suite.

// parseWithHeader reduces a with statement header and its suite.
func (p *pyParser) parseWithHeader(tok frontkit.Token) air.Node {
	start, modifier := tok, p.modifier
	if modifier.Text == "async" {
		start = modifier
	}
	p.stream.Advance()
	children := p.parseWithItems()
	block, end := p.parseSuite()
	if block != nil {
		children = append(children, block)
	}
	node := p.finish(air.KindWith, start, end, tok.Text, children...)
	if modifier.Text == "async" {
		p.mapper.annotate(node, air.SemAsync, "true")
	}
	return node
}

// parseWithItems reduces the comma-separated context managers of a with
// statement. Every item is one expression with an optional as target, and
// every item keeps its own byte-exact span inside the header.
func (p *pyParser) parseWithItems() []air.Node {
	items := make([]air.Node, 0, 4)
	for {
		p.expr.skipTrivia()
		cur, ok := p.stream.Peek()
		if !ok || !pyWithItemFollows(cur) {
			return items
		}
		before := p.stream.Remaining()
		item := p.expr.parse(pyAsPower)
		if item == nil || p.stream.Remaining() >= before {
			return items
		}
		items = append(items, p.mapper.withItem(item, cur))
		p.expr.skipTrivia()
		next, ok := p.stream.Peek()
		if !ok {
			return items
		}
		if next.Text == "as" {
			p.stream.Advance()
			if target := p.parseWithTarget(); target != nil {
				items = append(items, target)
			}
			continue
		}
		if next.Text == "," {
			p.stream.Advance()
			continue
		}
		return items
	}
}

// parseWithTarget reduces the as target of one context manager.
func (p *pyParser) parseWithTarget() air.Node {
	p.expr.skipTrivia()
	cur, ok := p.stream.Peek()
	if !ok || !pyWithItemFollows(cur) {
		return nil
	}
	before := p.stream.Remaining()
	target := p.expr.parse(pyAsPower)
	if target == nil || p.stream.Remaining() >= before {
		return nil
	}
	return target
}

// pyWithItemFollows reports whether one token can open a context manager or
// its as target. The colon, the line break and the indentation end the header.
func pyWithItemFollows(tok frontkit.Token) bool {
	switch tok.Kind {
	case frontkit.TokenNewline, frontkit.TokenIndent, frontkit.TokenDedent,
		frontkit.TokenEOF:
		return false
	}
	return tok.Text != ":"
}
