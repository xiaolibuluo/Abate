/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package pyfront

// This file holds the module loop and the statement dispatch of pyParser;
// the simple-statement productions live in pysimple.go and the compound
// statements, suites and recovery live in frontend.go.

import (
	"abate/lang/air"
	"abate/lang/frontkit"
)

// pyParser reduces a token stream into a layered Python AIR tree.
type pyParser struct {
	stream   *frontkit.TokenStream
	view     *frontkit.View
	mapper   *pyMap
	expr     *pyExpr
	sink     *frontkit.DiagnosticSink
	first    frontkit.Token
	last     frontkit.Token
	modifier frontkit.Token
	trivia   []air.Node
	depth    int
}

// newPyParser builds the default Python reducer for one file.
func newPyParser(tokens []frontkit.Token, path string) (*pyParser, error) {
	return newPyParserWithBindings(tokens, path, pyInfixBindings, pyNotes)
}

// newPyParserWithBindings builds a reducer with injectable binding and note
// tables, which mutation tests use to prove the tables affect the output.
func newPyParserWithBindings(tokens []frontkit.Token, path string,
	infix []pyBinding, notes []pyNote) (*pyParser, error) {
	stream, err := frontkit.NewTokenStream(tokens)
	if err != nil {
		return nil, err
	}
	view := frontkit.NewView(frontkit.SourceFromTokens(tokens), path)
	mapper := newPyMap(view)
	mapper.notes = notes
	first, last := frontkit.Token{}, frontkit.Token{}
	if len(tokens) > 0 {
		first, last = tokens[0], tokens[len(tokens)-1]
	}
	p := &pyParser{
		stream: stream, view: view, mapper: mapper,
		sink:  frontkit.NewDiagnosticSink(path, "py"),
		first: first, last: last,
	}
	p.expr = newPyExpr(stream, mapper, infix, &p.trivia)
	return p, nil
}

// parseModule reduces the whole file into one KindPackage root.
func (p *pyParser) parseModule() air.Node {
	stmts := make([]air.Node, 0, 8)
	for {
		p.skipBlank()
		tok, ok := p.stream.Peek()
		if !ok || tok.Kind == frontkit.TokenEOF {
			break
		}
		if tok.Kind == frontkit.TokenDedent {
			p.stream.Advance()
			continue
		}
		before := p.stream.Remaining()
		stmt := p.parseStatement()
		if stmt != nil {
			stmts = append(stmts, stmt)
		}
		if p.stream.Remaining() >= before {
			if _, advanced := p.stream.Advance(); !advanced {
				break
			}
		}
	}
	return p.mapper.build(air.KindPackage, p.first, p.last, "", stmts...)
}

// skipBlank consumes line breaks and stray indentation markers between statements.
func (p *pyParser) skipBlank() {
	for {
		tok, ok := p.stream.Peek()
		if !ok || (tok.Kind != frontkit.TokenNewline &&
			tok.Kind != frontkit.TokenIndent) {
			return
		}
		p.stream.Advance()
	}
}

// finish builds a statement node and attaches collected comment projections.
// The statement span always reaches the last collected trivia node, so a token
// swallowed by recovery never escapes the statement that absorbed it.
func (p *pyParser) finish(kind air.NodeKind, start, end frontkit.Token,
	text string, children ...air.Node) air.Node {
	all := append(children, p.trivia...)
	return p.mapper.build(kind, start, p.triviaEnd(end), text, all...)
}

// triviaEnd returns the token that reaches the last collected trivia node.
func (p *pyParser) triviaEnd(fallback frontkit.Token) frontkit.Token {
	end := fallback
	for i := range p.trivia {
		token := spanToken(p.trivia[i], end)
		if token.EndOffset > end.EndOffset {
			end = token
		}
	}
	return end
}

// parseStatement reduces one statement, applying single-pass recovery on damage.
func (p *pyParser) parseStatement() air.Node {
	p.trivia = p.trivia[:0]
	p.modifier = frontkit.Token{}
	tok, ok := p.stream.Peek()
	if !ok || tok.Kind == frontkit.TokenEOF {
		return nil
	}
	if tok.Kind == frontkit.TokenComment {
		p.stream.Advance()
		p.synchronize()
		return p.mapper.comment(tok)
	}
	if tok.Kind == frontkit.TokenKeyword && tok.Text == "async" {
		p.modifier = tok
		p.stream.Advance()
		tok, ok = p.stream.Peek()
		if !ok || tok.Kind == frontkit.TokenEOF {
			return p.partialStatement(p.modifier, p.modifier)
		}
	}
	if p.depth >= pyMaxDepth {
		return p.partialStatement(tok, tok)
	}
	p.depth++
	defer func() { p.depth-- }()
	if tok.Kind == frontkit.TokenKeyword {
		if tok.Text == "with" {
			return p.parseWithHeader(tok)
		}
		if rule, found := pyHeaderKind(tok.Text); found {
			return p.parseHeader(tok, rule)
		}
		if tok.Text == "def" {
			return p.parseFuncDef(tok)
		}
		if tok.Text == "class" {
			return p.parseClassDef(tok)
		}
		if rule, found := pySimpleStmtRule(tok.Text); found {
			return p.parseSimpleKeyword(tok, rule)
		}
	}
	if tok.Text == "@" {
		return p.parseDecorator(tok)
	}
	return p.parseSimple(tok)
}

// partialStatement records one depth-limited or malformed reduction.
func (p *pyParser) partialStatement(start, end frontkit.Token) air.Node {
	p.synchronize()
	return p.mapper.partial(p.finish(air.KindPartialNode, start, end, ""))
}
