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

// This file holds the Pratt core of pyExpr; the container, literal and
// callable-form reductions that consume the lexical tables live in pylexspec.go.

// pyExpr reduces Python expressions with an injectable Pratt binding table.
type pyExpr struct {
	stream  *frontkit.TokenStream
	mapper  *pyMap
	infix   []pyBinding
	trivia  *[]air.Node
	bracket int
	depth   int
}

// newPyExpr binds a token cursor, a node mapper and a binding table.
func newPyExpr(stream *frontkit.TokenStream, mapper *pyMap,
	infix []pyBinding, trivia *[]air.Node) *pyExpr {
	return &pyExpr{stream: stream, mapper: mapper, infix: infix, trivia: trivia}
}

// parse reduces one expression whose infix operators bind at minPower or tighter.
func (e *pyExpr) parse(minPower int) air.Node {
	if e.depth >= pyMaxDepth {
		tok, _ := e.stream.Peek()
		return e.mapper.partialLeaf(tok)
	}
	e.depth++
	defer func() { e.depth-- }()
	e.skipTrivia()
	left := e.parsePrefix()
	for {
		e.skipTrivia()
		tok, ok := e.stream.Peek()
		if !ok || tok.Kind == frontkit.TokenEOF {
			break
		}
		if pyContains(pySuffixTexts, tok.Text) && pySuffixPower >= minPower {
			if _, advanced := e.stream.Advance(); !advanced {
				break
			}
			left = e.parseSuffix(left, tok)
			continue
		}
		power := pyInfixPower(e.infix, tok.Text)
		if power == 0 || power < minPower {
			break
		}
		if _, advanced := e.stream.Advance(); !advanced {
			break
		}
		if tok.Text == "if" {
			left = e.parseTernary(left, tok, power)
			continue
		}
		next := power
		if !pyInfixRight(e.infix, tok.Text) {
			next++
		}
		left = e.mapper.binary(left, tok, e.parse(next))
	}
	return left
}

// skipTrivia consumes implicit line joins and projects comment tokens.
func (e *pyExpr) skipTrivia() {
	for {
		tok, ok := e.stream.Peek()
		if !ok {
			return
		}
		if tok.Kind == frontkit.TokenComment {
			e.stream.Advance()
			*e.trivia = append(*e.trivia, e.mapper.comment(tok))
			continue
		}
		if tok.Kind == frontkit.TokenNewline || tok.Kind == frontkit.TokenIndent ||
			tok.Kind == frontkit.TokenDedent {
			if e.bracket == 0 {
				return
			}
			e.stream.Advance()
			continue
		}
		if tok.Text == "\\" {
			e.stream.Advance()
			if following, ok := e.stream.Peek(); ok &&
				following.Kind == frontkit.TokenNewline {
				e.stream.Advance()
			}
			continue
		}
		return
	}
}

// parsePrefix reduces one atom or one prefix operator.
func (e *pyExpr) parsePrefix() air.Node {
	tok, ok := e.stream.Advance()
	if !ok || tok.Kind == frontkit.TokenEOF {
		return e.mapper.partialLeaf(tok)
	}
	switch tok.Kind {
	case frontkit.TokenIdent:
		return e.mapper.leaf(tok, air.KindIdent)
	case frontkit.TokenNum:
		return e.mapper.leaf(tok, air.KindNumLit)
	case frontkit.TokenStr:
		return e.parseString(tok)
	case frontkit.TokenChar:
		return e.mapper.leaf(tok, air.KindCharLit)
	case frontkit.TokenKeyword:
		return e.parseKeywordPrefix(tok)
	case frontkit.TokenOp:
		return e.parseOpPrefix(tok)
	case frontkit.TokenPunct:
		return e.parsePunctPrefix(tok)
	default:
		if pyIsLexicalKind(tok.Kind) {
			return e.mapper.lexical(tok)
		}
		return e.mapper.partialLeaf(tok)
	}
}

// parseKeywordPrefix lowers literals, await, lambda and keyword operands.
func (e *pyExpr) parseKeywordPrefix(tok frontkit.Token) air.Node {
	switch tok.Text {
	case "True", "False":
		return e.mapper.leaf(tok, air.KindBoolLit)
	case "None":
		return e.mapper.leaf(tok, air.KindNilLit)
	case "not":
		operand := e.parse(30)
		end := spanToken(operand, tok)
		return e.mapper.boundOperator(air.KindUnary, tok, end, tok,
			operand, e.mapper.dedicated(air.KindPyNotOp, tok, end,
				tok.Text, "py-not-op", operand))
	case "await", "async":
		return e.mapper.awaitNode(tok, e.parse(30))
	case "yield":
		return e.mapper.unary(tok, e.parse(30), air.KindUnary, "yield")
	case "lambda":
		return e.parseLambda(tok)
	case "match":
		return e.mapper.matchExpr(e.parse(30))
	default:
		// A keyword that reaches an expression-prefix position without a
		// reduction is a recovery position: honest partial node, never a
		// structural kind guessed from the keyword text.
		return e.mapper.partialLeaf(tok)
	}
}

// parseOpPrefix lowers prefix operators, containers and spread arguments.
func (e *pyExpr) parseOpPrefix(tok frontkit.Token) air.Node {
	switch tok.Text {
	case "(":
		return e.parseDelimited(tok, pyModeParen)
	case "[":
		return e.parseDelimited(tok, pyModeList)
	case "{":
		return e.parseDelimited(tok, pyModeDict)
	case "-", "+", "~":
		return e.mapper.unary(tok, e.parse(80), air.KindUnary, "")
	case "*", "**":
		return e.mapper.spread(tok, e.parse(80))
	case "...":
		// py.ellipsis 锁定降级：省略号在真实归约现场降级为区域节点。
		return e.mapper.annotate(
			e.mapper.build(air.KindRegion, tok, tok, tok.Text),
			air.SemNote, e.mapper.note("ellipsis"))
	default:
		return e.mapper.partialLeaf(tok)
	}
}

// parsePunctPrefix lowers bracketed prefixes and decorator marks.
func (e *pyExpr) parsePunctPrefix(tok frontkit.Token) air.Node {
	switch tok.Text {
	case "(":
		return e.parseDelimited(tok, pyModeParen)
	case "[":
		return e.parseDelimited(tok, pyModeList)
	case "{":
		return e.parseDelimited(tok, pyModeDict)
	case "@":
		return e.mapper.annotate(
			e.mapper.build(air.KindAttribute, tok, tok, ""),
			air.SemNote, e.mapper.note("decorator"))
	default:
		return e.mapper.partialLeaf(tok)
	}
}
