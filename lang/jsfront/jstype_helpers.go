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

// decorateToken 按连续 case 段竖切归约单个 token 的降级节点.
func (t *jsTypeParser) decorateToken(out []air.Node, i, start, end int,
	tok frontkit.Token) []air.Node {
	if next, ok := t.decorateLeading(out, i, tok); ok {
		return next
	}
	if next, ok := t.decorateAnnotation(out, i, start, end, tok); ok {
		return next
	}
	if next, ok := t.decorateTSOperator(out, i, end, tok); ok {
		return next
	}
	if next, ok := t.decorateTSKeyword(out, i, tok); ok {
		return next
	}
	if next, ok := t.decorateJSXSymbol(out, i, start, tok); ok {
		return next
	}
	return t.decorateKeyword(out, i, tok)
}

// decorateLeading 归约装饰器与指令序言 token.
func (t *jsTypeParser) decorateLeading(out []air.Node, i int,
	tok frontkit.Token) ([]air.Node, bool) {
	switch {
	case tok.Text == "@":
		return append(out, t.simpleNode(i, air.KindDecorator, noteDecorator)), true
	case isUseStrictToken(tok):
		n := t.simpleNode(i, air.KindDirective, noteUseStrict)
		_ = air.Annotate(n, air.SemDirective, "use strict")
		return append(out, n), true
	}
	return out, false
}

// decorateAnnotation 归约 ":" 类型注解、"<" 泛型参数列表与 JSX 起始标签.
func (t *jsTypeParser) decorateAnnotation(out []air.Node, i, start, end int,
	tok frontkit.Token) ([]air.Node, bool) {
	p := t.p
	switch {
	case tok.Text == ":" && p.d.ts && typeFollows(p.toks, i, end):
		return append(out, t.typeNode(i, end)), true
	case tok.Text == "<" && p.d.ts && i > start &&
		genericPrefix(p.toks, i) && looksGeneric(p.toks, i, end):
		return append(out, t.genericNodes(i, end)...), true
	case tok.Text == "<" && p.d.jsx && jsxFollows(p.toks, i, end, p.d):
		return append(out, t.simpleNode(i, air.KindCompositeLit, noteJSX)), true
	}
	return out, false
}

// decorateTSOperator 归约 TS 联合类型、可选标记与元组类型.
func (t *jsTypeParser) decorateTSOperator(out []air.Node, i, end int,
	tok frontkit.Token) ([]air.Node, bool) {
	p := t.p
	switch {
	case tok.Text == "|" && p.d.ts:
		return append(out, t.simpleNode(i, air.KindUnion, noteUnion)), true
	case tok.Text == "?" && p.d.ts && optionalFollows(p.toks, i, end):
		return append(out, t.simpleNode(i, air.KindOptional, noteOptional)), true
	case tok.Text == "[" && p.d.ts && tupleFollows(p.toks, i):
		return append(out, t.simpleNode(i, air.KindTuple, noteTuple)), true
	}
	return out, false
}

// decorateTSKeyword 归约 readonly 修饰符与 as/satisfies 类型断言.
func (t *jsTypeParser) decorateTSKeyword(out []air.Node, i int,
	tok frontkit.Token) ([]air.Node, bool) {
	p := t.p
	switch {
	case tok.Text == "readonly" && p.d.ts:
		n := t.simpleNode(i, air.KindNamed, noteReadonly)
		_ = air.Annotate(n, air.SemMutability, "readonly")
		return append(out, n), true
	case (tok.Text == "as" || tok.Text == "satisfies") && p.d.ts:
		return append(out, t.simpleNode(i, air.KindTypeAssert,
			"type-assert-"+tok.Text)), true
	}
	return out, false
}

// decorateJSXSymbol 归约 JSX 属性等号、成员访问点与展开运算符.
func (t *jsTypeParser) decorateJSXSymbol(out []air.Node, i, start int,
	tok frontkit.Token) ([]air.Node, bool) {
	p := t.p
	switch {
	case tok.Text == "=" && p.d.jsx && jsxAttrFollows(p.toks, i, start):
		return append(out, t.simpleNode(i, air.KindField, noteJSXAttribute)), true
	case tok.Text == ".":
		return append(out, t.simpleNode(i, air.KindSelector, "member-access")), true
	case tok.Text == "...":
		return append(out, t.simpleNode(i, air.KindSpread, noteSpread)), true
	}
	return out, false
}

// decorateKeyword 归约 default 分支中的声明类关键字映射节点.
func (t *jsTypeParser) decorateKeyword(out []air.Node, i int,
	tok frontkit.Token) []air.Node {
	if tok.Kind == frontkit.TokenKeyword {
		if kind, note, ok := jsDeclMapping(tok.Text); ok {
			return append(out, t.simpleNode(i, kind, note))
		}
	}
	return out
}
