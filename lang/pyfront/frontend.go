/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package pyfront

// This file holds the public air.Frontend wiring plus the compound-statement,
// indentation-suite and synchronization reductions of pyParser.

import (
	"fmt"
	"strings"

	"abate/lang/air"
	"abate/lang/frontkit"
)

// PyFrontend is the Python frontend handle, an alias of the shared frontkit
// base bound to pySpec by NewPyFrontend.
type PyFrontend = frontkit.FrontendBase

// Compile-time proofs of the two contracts this package implements.
var (
	_ frontkit.LanguageSpec = pySpec{}
	_ air.Frontend          = (*PyFrontend)(nil)
)

// NewPyFrontend returns an air.Frontend that reduces Python sources to AIR.
func NewPyFrontend() (air.Frontend, error) {
	return frontkit.NewFrontendBase(pySpec{})
}

// pySpec implements frontkit.LanguageSpec for Python.
type pySpec struct{}

// Tag returns the canonical short Python language tag.
func (pySpec) Tag() string { return "py" }

// LexSpec returns the Python lexical tables.
func (pySpec) LexSpec() frontkit.LexSpec { return pyLexSpec{} }

// NormalizeDialect maps one dialect spelling onto a canonical 3.x string.
func (pySpec) NormalizeDialect(d string) (string, error) {
	normalized := strings.ToLower(strings.TrimSpace(d))
	if normalized == "" {
		return pyDefaultDialect, nil
	}
	for i := range pyDialectPrefixes {
		normalized = strings.TrimPrefix(normalized, pyDialectPrefixes[i])
	}
	if pyContains(pyDialects, normalized) {
		return normalized, nil
	}
	return "", fmt.Errorf("pyfront: unsupported dialect %q", d)
}

// ParseTokens reduces one file's token stream into a layered AIR tree.
func (pySpec) ParseTokens(tokens []frontkit.Token, path string) (air.Node,
	[]air.Diagnostic, error) {
	parser, err := newPyParser(tokens, path)
	if err != nil {
		return nil, nil, err
	}
	root := parser.parseModule()
	return root, parser.sink.Diagnostics(), nil
}

// parseHeader reduces a compound statement header and its suite.
func (p *pyParser) parseHeader(tok frontkit.Token, rule pyHeaderRule) air.Node {
	start, modifier := tok, p.modifier
	if modifier.Text == "async" {
		start = modifier
	}
	p.stream.Advance()
	var condition air.Node
	if cur, ok := p.stream.Peek(); ok && cur.Text != ":" &&
		cur.Kind != frontkit.TokenNewline && cur.Kind != frontkit.TokenEOF {
		condition = p.expr.parse(0)
	}
	if condition != nil && tok.Text == "match" {
		condition = p.mapper.matchExpr(condition)
	}
	block, end := p.parseSuite()
	children := make([]air.Node, 0, 2)
	if condition != nil {
		children = append(children, condition)
	}
	if block != nil {
		children = append(children, block)
	}
	node := p.finish(rule.kind, start, end, tok.Text, children...)
	if rule.note != "" {
		p.mapper.annotate(node, air.SemNote, p.mapper.note(rule.note))
	}
	if modifier.Text == "async" {
		p.mapper.annotate(node, air.SemAsync, "true")
	}
	return node
}

// parseFuncDef reduces a def or async def declaration.
func (p *pyParser) parseFuncDef(tok frontkit.Token) air.Node {
	start, modifier := tok, p.modifier
	if modifier.Text == "async" {
		start = modifier
	}
	p.stream.Advance()
	nameTok, ok := p.stream.Peek()
	if !ok || nameTok.Kind == frontkit.TokenEOF {
		return p.partialStatement(start, start)
	}
	p.stream.Advance()
	name := p.mapper.leaf(nameTok, air.KindIdent)
	children := []air.Node{name}
	if cur, ok := p.stream.Peek(); ok && cur.Text == "(" {
		p.stream.Advance()
		p.expr.bracket++
		params, _, _ := p.expr.parseList(")", pyListSlice)
		p.expr.bracket--
		children = append(children, params...)
	}
	if cur, ok := p.stream.Peek(); ok && cur.Text == "->" {
		p.stream.Advance()
		if annotation := p.expr.parse(0); annotation != nil {
			children = append(children, p.mapper.typeNode(annotation))
		}
	}
	block, end := p.parseSuite()
	if block != nil {
		children = append(children, block)
	}
	node := p.finish(air.KindFuncDecl, start, end, air.AttributeOf(name), children...)
	if modifier.Text == "async" {
		p.mapper.annotate(node, air.SemAsync, "true")
	}
	return node
}

// parseClassDef reduces a class declaration and its base list.
func (p *pyParser) parseClassDef(tok frontkit.Token) air.Node {
	p.stream.Advance()
	nameTok, ok := p.stream.Peek()
	if !ok || nameTok.Kind == frontkit.TokenEOF {
		return p.partialStatement(tok, tok)
	}
	p.stream.Advance()
	name := p.mapper.leaf(nameTok, air.KindIdent)
	children := []air.Node{name}
	if cur, ok := p.stream.Peek(); ok && cur.Text == "(" {
		p.stream.Advance()
		p.expr.bracket++
		bases, _, _ := p.expr.parseList(")", pyListPlain)
		p.expr.bracket--
		children = append(children, bases...)
	}
	block, end := p.parseSuite()
	if block != nil {
		children = append(children, block)
	}
	return p.finish(air.KindClassDecl, tok, end, air.AttributeOf(name), children...)
}

// parseDecorator reduces one decorator line.
func (p *pyParser) parseDecorator(tok frontkit.Token) air.Node {
	p.stream.Advance()
	target := p.expr.parse(0)
	p.synchronize()
	end := p.triviaEnd(spanToken(target, tok))
	children := append([]air.Node{target}, p.trivia...)
	return p.mapper.decorator(tok, end, children...)
}

// parseSuite reduces either an indented suite or a same-line simple suite.
func (p *pyParser) parseSuite() (air.Node, frontkit.Token) {
	if cur, ok := p.stream.Peek(); ok && cur.Text == ":" {
		p.stream.Advance()
	}
	p.consumeHeaderTail()
	cur, ok := p.stream.Peek()
	if !ok || cur.Kind == frontkit.TokenEOF {
		return p.mapper.partialLeaf(cur), cur
	}
	if cur.Kind == frontkit.TokenIndent {
		p.stream.Advance()
		body, end := p.parseBody()
		if len(body) == 0 {
			return p.mapper.partialLeaf(cur), cur
		}
		return p.mapper.indentBlock(cur, end, body...), end
	}
	stmt := p.parseStatement()
	if stmt == nil {
		return p.mapper.partialLeaf(cur), cur
	}
	end := spanToken(stmt, cur)
	return p.mapper.build(air.KindBlock, cur, end, "", stmt), end
}

// consumeHeaderTail projects header-line comments and eats physical line breaks.
func (p *pyParser) consumeHeaderTail() {
	for {
		tok, ok := p.stream.Peek()
		if !ok {
			return
		}
		if tok.Kind == frontkit.TokenComment {
			p.stream.Advance()
			p.trivia = append(p.trivia, p.mapper.comment(tok))
			continue
		}
		if tok.Kind == frontkit.TokenNewline {
			p.stream.Advance()
			continue
		}
		return
	}
}

// parseBody reduces the statements of one indented suite up to DEDENT.
func (p *pyParser) parseBody() ([]air.Node, frontkit.Token) {
	body := make([]air.Node, 0, 4)
	end := frontkit.Token{}
	for {
		p.skipBlank()
		cur, ok := p.stream.Peek()
		if !ok || cur.Kind == frontkit.TokenEOF {
			return body, end
		}
		if cur.Kind == frontkit.TokenDedent {
			p.stream.Advance()
			return body, end
		}
		before := p.stream.Remaining()
		stmt := p.parseStatement()
		if stmt != nil {
			body = append(body, stmt)
			end = spanToken(stmt, cur)
		}
		if p.stream.Remaining() >= before {
			if _, advanced := p.stream.Advance(); !advanced {
				return body, end
			}
		}
	}
}

// synchronize consumes one logical line, projecting skipped lexical tokens.
func (p *pyParser) synchronize() {
	for {
		tok, ok := p.stream.Peek()
		if !ok {
			return
		}
		if tok.Kind == frontkit.TokenEOF || tok.Kind == frontkit.TokenDedent ||
			tok.Kind == frontkit.TokenIndent {
			return
		}
		p.stream.Advance()
		if tok.Kind == frontkit.TokenNewline {
			return
		}
		p.collect(tok)
	}
}

// collect projects a skipped token and marks the recovery as partial.
func (p *pyParser) collect(tok frontkit.Token) {
	if pyIsLexicalKind(tok.Kind) {
		p.trivia = append(p.trivia, p.mapper.partial(p.mapper.lexical(tok)))
	}
	if tok.Kind == frontkit.TokenIllegal {
		_ = p.sink.Warning(2, tok.Offset, tok.Line, tok.Column,
			"illegal byte sequence skipped during recovery")
	}
}
