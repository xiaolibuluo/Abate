/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package rsfront

// This file holds the Pratt core of rsExpr: the ?, ::, closure, range and macro
// call reductions. Delimiter, literal and generic helpers live in rslexspec.go.

import (
	"abate/lang/air"
	"abate/lang/frontkit"
)

// rsExpr reduces one Rust expression interval with an injectable Pratt table.
// Each reduce call uses a fresh value cursor, so nested reductions never share
// mutable state and recursion stays bounded by an integer depth ceiling.
type rsExpr struct {
	tokens   []frontkit.Token
	mapper   *rsMapper
	infix    []rsBinding
	extra    []air.Node
	at       int
	hi       int
	depth    int
	maxDepth int
}

// newRsExpr binds a node mapper and the injected binding table.
func newRsExpr(mapper *rsMapper, cfg rsConfig) *rsExpr {
	limit := cfg.maxDepth
	if limit <= 0 {
		limit = rsMaxDepth
	}
	return &rsExpr{mapper: mapper, infix: cfg.infix, maxDepth: limit}
}

// reduce projects one token interval into ordered, non-overlapping AIR nodes.
func (e *rsExpr) reduce(tokens []frontkit.Token, lo, hi int) []air.Node {
	cur := rsExpr{
		tokens: tokens, mapper: e.mapper, infix: e.infix,
		at: lo, hi: hi, maxDepth: e.maxDepth, extra: make([]air.Node, 0),
	}
	return cur.reduceAll()
}

// reduceAll loops over the interval and reduces one expression at a time.
func (e *rsExpr) reduceAll() []air.Node {
	out := make([]air.Node, 0, 8)
	for e.at < e.hi {
		start := e.at
		tok := e.tokens[e.at]
		if rsIsTrivia(tok.Kind) || tok.Kind == frontkit.TokenEOF {
			e.at++
			continue
		}
		if tok.Kind == frontkit.TokenComment {
			out = append(out, e.mapper.comment(tok))
			e.at++
			continue
		}
		node := e.parseExpression(0)
		if node != nil {
			out = append(out, node)
		}
		if e.at <= start {
			e.at = start + 1
		}
	}
	out = append(out, e.extra...)
	return e.mapper.view.OrderChildren(out)
}

// parseExpression reduces one expression whose operators bind at minPower.
func (e *rsExpr) parseExpression(minPower int) air.Node {
	if e.at >= e.hi {
		return nil
	}
	if e.depth >= e.maxDepth {
		tok := e.tokens[e.at]
		e.at++
		return e.mapper.partialLeaf(tok)
	}
	e.depth++
	defer func() { e.depth-- }()
	left := e.parsePrefix()
	if left == nil {
		return nil
	}
	for e.at < e.hi {
		tok := e.tokens[e.at]
		if rsIsTrivia(tok.Kind) {
			e.at++
			continue
		}
		if tok.Kind == frontkit.TokenComment {
			e.extra = append(e.extra, e.mapper.comment(tok))
			e.at++
			continue
		}
		if tok.Kind == frontkit.TokenEOF {
			break
		}
		if tok.Text == "<" {
			if closeIdx, ok := e.findAngleClose(e.at); ok {
				if rsPostfixPower < minPower {
					break
				}
				// The parser recognized a type argument list while reducing
				// the generic instantiation of this expression.
				e.at++
				args := e.reduceSub(e.at, closeIdx)
				closeTok := e.tokens[closeIdx]
				e.at = closeIdx + 1
				left = e.mapper.genericInstMarked(left, closeTok,
					[]air.Node{e.mapper.dedicated(air.KindRsTypeArgs, tok,
						tok.Text, "rs-type-arguments")}, args...)
				continue
			}
		}
		power := 0
		suffix := false
		switch tok.Text {
		case "?", ".", "(", "[":
			power, suffix = rsPostfixPower, true
		case "::":
			power, suffix = rsSelectorPower, true
		case "!":
			if e.at+1 < e.hi && rsIsOpenBracket(e.tokens[e.at+1].Text) {
				power, suffix = rsPostfixPower, true
			}
		default:
			power = rsInfixPower(e.infix, tok.Text)
		}
		if power == 0 || power < minPower {
			break
		}
		e.at++
		if suffix {
			left = e.applySuffix(left, tok)
			continue
		}
		next := power
		if !rsInfixRight(e.infix, tok.Text) {
			next++
		}
		right := e.parseExpression(next)
		if right == nil {
			break
		}
		left = e.applyInfix(left, tok, right)
	}
	return left
}

// parsePrefix reduces one atom or one prefix operator.
func (e *rsExpr) parsePrefix() air.Node {
	if e.at >= e.hi {
		return nil
	}
	tok := e.tokens[e.at]
	e.at++
	switch tok.Kind {
	case frontkit.TokenIdent:
		return e.mapper.leaf(tok, air.KindIdent)
	case frontkit.TokenNum:
		return e.mapper.leaf(tok, air.KindNumLit)
	case frontkit.TokenStr:
		return e.mapper.stringNode(tok)
	case frontkit.TokenChar:
		return e.mapper.leaf(tok, air.KindCharLit)
	case frontkit.TokenKeyword:
		return e.parseKeywordPrefix(tok)
	case frontkit.TokenOp:
		return e.parseOpPrefix(tok)
	case frontkit.TokenPunct:
		return e.parsePunctPrefix(tok)
	case frontkit.TokenComment:
		e.extra = append(e.extra, e.mapper.comment(tok))
		return nil
	default:
		return e.mapper.partialLeaf(tok)
	}
}

// applySuffix reduces a postfix operator with its already-reduced left node.
func (e *rsExpr) applySuffix(left air.Node, tok frontkit.Token) air.Node {
	switch tok.Text {
	case "?":
		return e.mapper.tryNode(left, tok)
	case ".":
		if e.at >= e.hi {
			return e.mapper.partial(left)
		}
		name := e.tokens[e.at]
		e.at++
		if name.Text == "await" {
			return e.mapper.awaitNode(left, name)
		}
		return e.mapper.field(left, name)
	case "(":
		children, closeTok := e.parseDelimited(tok, ")")
		return e.mapper.call(left, tok, closeTok, children...)
	case "[":
		children, closeTok := e.parseDelimited(tok, "]")
		return e.mapper.index(left, tok, closeTok, children...)
	case "!":
		if e.at >= e.hi {
			return e.mapper.partial(left)
		}
		open := e.tokens[e.at]
		e.at++
		children, closeTok := e.parseDelimited(open, rsCloseOf(open.Text))
		return e.mapper.macroCall(left, tok, closeTok, children...)
	case "::":
		if e.at >= e.hi {
			return e.mapper.partial(left)
		}
		name := e.tokens[e.at]
		if name.Text == "<" {
			if closeIdx, ok := e.findAngleClose(e.at); ok {
				// The parser reduced a turbofish generic instantiation; the
				// angle list itself is also a genuine type-arguments node.
				args := e.reduceSub(e.at+1, closeIdx)
				closeTok := e.tokens[closeIdx]
				e.at = closeIdx + 1
				return e.mapper.genericInstMarked(left, closeTok,
					[]air.Node{
						e.mapper.dedicated(air.KindRsTurbofish, tok,
							tok.Text, "rs-turbofish"),
						e.mapper.dedicated(air.KindRsTypeArgs, name,
							name.Text, "rs-type-arguments"),
					}, args...)
			}
		}
		e.at++
		return e.mapper.selector(left, tok, name)
	default:
		return e.mapper.partial(left)
	}
}

// applyInfix reduces an infix operator with both operands already reduced.
func (e *rsExpr) applyInfix(left air.Node, op frontkit.Token,
	right air.Node) air.Node {
	switch {
	case rsContains(rsAssignOps, op.Text):
		return e.mapper.assign(left, op, right)
	case op.Text == ".." || op.Text == "..=":
		return e.mapper.rangeNode(left, op, right)
	case op.Text == "as":
		return e.mapper.cast(left, op, right)
	case op.Text == "=>" || op.Text == "->":
		return e.mapper.arrow(left, op, right)
	default:
		return e.mapper.binaryNodes(left, op, right)
	}
}
