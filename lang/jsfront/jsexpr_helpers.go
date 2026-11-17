/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package jsfront

import (
	"abate/lang/air"
	"abate/lang/frontkit"
)

// parseLoop 承载 Pratt 优先级攀升的后缀、中缀与终止判定循环.
func (e *jsExprParser) parseLoop(left air.Node, minPower int) air.Node {
	for e.i < e.end {
		tok, ok := e.peek()
		if !ok || tok.Kind == frontkit.TokenEOF {
			break
		}
		if tok.Kind == frontkit.TokenNewline {
			e.i++
			continue
		}
		if jsIsExprTerminator(tok.Text) {
			return left
		}
		if next, handled := e.parseChainToken(left, tok); handled {
			left = next
			continue
		}
		rule, found := jsLookupInfix(tok.Text)
		if !found || rule.power < minPower {
			return left
		}
		e.i++
		nextPower := rule.power
		if !rule.right {
			nextPower++
		}
		right := e.parse(nextPower)
		left = e.p.mapper.binary(rule, left, right, tok)
	}
	return left
}

// jsIsExprTerminator 判定 token 文本是否终止当前表达式归约.
func jsIsExprTerminator(text string) bool {
	switch text {
	case ";", "}", ",", ":":
		return true
	}
	return false
}

// parseChainToken 归约后缀链 token，返回更新后的 left 与是否已消费该 token.
func (e *jsExprParser) parseChainToken(left air.Node,
	tok frontkit.Token) (air.Node, bool) {
	switch tok.Text {
	case "?.":
		return e.optionalChainPostfix(left, tok), true
	case ".":
		e.i++
		return e.p.mapper.chain(left, e.parseMember(),
			air.KindSelector, "member-access", tok), true
	case "(":
		return e.parseCallLike(left, e.i, air.KindCall, "call"), true
	case "[":
		return e.parseCallLike(left, e.i, air.KindIndex, "index"), true
	case "=>":
		e.i++
		return e.p.mapper.arrow(left, e.parse(2), tok), true
	case "?":
		return e.parseTernary(left, e.i), true
	case "as", "satisfies":
		e.i++
		return e.p.mapper.typeAssert(left, e.parse(3), tok), true
	case "<":
		return e.anglePostfix(left, tok)
	}
	return left, false
}

// optionalChainPostfix 归约 "?." 可选链或可选调用.
func (e *jsExprParser) optionalChainPostfix(left air.Node,
	tok frontkit.Token) air.Node {
	e.i++
	next, found := e.peek()
	if found && next.Text == "(" {
		return e.parseCallLike(left, e.i, air.KindCall, "optional-call")
	}
	return e.p.mapper.chain(left, e.parseMember(),
		air.KindField, "optional-chain", tok)
}

// anglePostfix 在 JSX 方言下把 "<" 归约为起始标签，否则交回中缀小于运算符.
func (e *jsExprParser) anglePostfix(left air.Node,
	tok frontkit.Token) (air.Node, bool) {
	next, found := e.peekAt(e.i + 1)
	if e.p.d.jsx && found && jsIsJSXOpen(next, e.p.d) {
		at := e.i
		e.i++
		return e.jsxPrefix(at, tok, left), true
	}
	return left, false
}

// skipNewlines 跳过前缀位置之前的换行 token.
func (e *jsExprParser) skipNewlines() {
	for e.i < e.end {
		tok, ok := e.peek()
		if !ok || tok.Kind != frontkit.TokenNewline {
			break
		}
		e.i++
	}
}

// prefixDispatch 按 token 类别把前缀归约分派到字面量、关键字或符号构造.
func (e *jsExprParser) prefixDispatch(idx int, tok frontkit.Token) air.Node {
	if jsIsLexicalToken(tok.Kind) {
		return e.lexicalPrefix(idx, tok)
	}
	if tok.Kind == frontkit.TokenKeyword {
		return e.keywordPrefix(idx, tok)
	}
	return e.symbolPrefix(idx, tok)
}

// lexicalPrefix 归约保真子集 token，模板字符串额外升级为模板节点.
func (e *jsExprParser) lexicalPrefix(idx int, tok frontkit.Token) air.Node {
	leaf := e.p.mapper.leaf(idx, tok)
	if jsIsTemplate(tok) {
		return e.p.mapper.template(leaf, tok)
	}
	return leaf
}

// symbolPrefix 归约分组、字面量、装饰器、一元运算与专有符号构造.
func (e *jsExprParser) symbolPrefix(idx int, tok frontkit.Token) air.Node {
	switch {
	case tok.Text == "(":
		return e.groupPrefix(idx, tok)
	case tok.Text == "[":
		return e.arrayPrefix(idx, tok)
	case tok.Text == "{":
		return e.objectPrefix(idx, tok)
	case tok.Text == "<" && e.p.d.jsx:
		return e.jsxAnglePrefix(idx, tok)
	case tok.Text == "@":
		return e.p.mapper.decorator(tok, e.parse(14))
	case tok.Text == "...":
		return e.p.mapper.spread(tok, e.parse(14))
	case jsIsPrefixUnary(tok.Text):
		return e.p.mapper.unary(tok, e.parse(14))
	case tok.Text == "/":
		return e.regexOrDivision(idx, tok)
	default:
		return e.p.mapper.opaque(idx, tok)
	}
}

// jsxAnglePrefix 在 JSX 方言下归约 "<" 起始标签，否则退回容错节点.
func (e *jsExprParser) jsxAnglePrefix(idx int, tok frontkit.Token) air.Node {
	next, found := e.peekAt(e.i)
	if found && jsIsJSXOpen(next, e.p.d) {
		return e.jsxPrefix(idx, tok, nil)
	}
	return e.p.mapper.opaque(idx, tok)
}

// jsIsPrefixUnary 判定 token 文本是否为前缀一元运算符.
func jsIsPrefixUnary(text string) bool {
	switch text {
	case "!", "~", "++", "--", "+", "-":
		return true
	}
	return false
}
