/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package parse

import (
	"fmt"
	"unicode"
)

// MaxExprDepth bounds the nesting of a CES operator expression.
const MaxExprDepth = 64

// tokenKind enumerates the lexical tokens of the operator calculus.
type tokenKind int

// Token kinds of the expression grammar.
const (
	tokEOF tokenKind = iota
	tokLParen
	tokRParen
	tokCompose
	tokProduct
	tokInverse
	tokIdentity
	tokRef
	tokBare
)

// token is one positioned lexical unit.
type token struct {
	kind tokenKind
	text string
	pos  int
}

// nodeKind enumerates the expression tree node types.
type nodeKind int

// Node kinds of the expression tree.
const (
	nodeIdentity nodeKind = iota
	nodeOperand
	nodeInverse
	nodeCompose
	nodeProduct
)

// exprNode is one node of the parsed operator expression tree.
type exprNode struct {
	kind  nodeKind
	left  *exprNode
	right *exprNode
	text  string
	ref   bool
	pos   int
}

// lex splits an expression into positioned tokens.
func lex(expr string) ([]token, error) {
	runes := []rune(expr)
	tokens := make([]token, 0, len(runes)/2+1)
	i := 0
	for i < len(runes) {
		r := runes[i]
		switch {
		case unicode.IsSpace(r):
			i++
		case r == '(':
			tokens = append(tokens, token{kind: tokLParen, pos: i})
			i++
		case r == ')':
			tokens = append(tokens, token{kind: tokRParen, pos: i})
			i++
		case r == '∘':
			tokens = append(tokens, token{kind: tokCompose, pos: i})
			i++
		case r == '⊗':
			tokens = append(tokens, token{kind: tokProduct, pos: i})
			i++
		case r == 'ε':
			tokens = append(tokens, token{kind: tokIdentity, pos: i})
			i++
		case r == '⁻' && i+1 < len(runes) && runes[i+1] == '¹':
			tokens = append(tokens, token{kind: tokInverse, pos: i})
			i += 2
		case hasPrefixAt(runes, i, "operator://"):
			j := i + len("operator://")
			start := j
			for j < len(runes) && isIDRune(runes[j]) {
				j++
			}
			if j == start {
				return nil, fmt.Errorf("empty operator reference at %d", i)
			}
			tokens = append(tokens, token{kind: tokRef, text: string(runes[start:j]), pos: i})
			i = j
		case isIDRune(r):
			j := i
			for j < len(runes) && isIDRune(runes[j]) {
				j++
			}
			tokens = append(tokens, token{kind: tokBare, text: string(runes[i:j]), pos: i})
			i = j
		default:
			return nil, fmt.Errorf("unexpected rune %q at %d", r, i)
		}
	}
	tokens = append(tokens, token{kind: tokEOF, pos: len(runes)})
	return tokens, nil
}

// hasPrefixAt reports whether a rune slice starts with a prefix at offset i.
func hasPrefixAt(runes []rune, i int, prefix string) bool {
	p := []rune(prefix)
	if i+len(p) > len(runes) {
		return false
	}
	for j, r := range p {
		if runes[i+j] != r {
			return false
		}
	}
	return true
}

// isIDRune reports whether a rune may appear in an operator identifier.
func isIDRune(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r) ||
		r == '-' || r == '_' || r == '.'
}

// exprParser is a hand-written recursive-descent parser for the calculus.
type exprParser struct {
	tokens []token
	pos    int
	depth  int
}

// peek returns the current token without consuming it.
func (p *exprParser) peek() token {
	return p.tokens[p.pos]
}

// next consumes and returns the current token.
func (p *exprParser) next() token {
	t := p.tokens[p.pos]
	if p.pos < len(p.tokens)-1 {
		p.pos++
	}
	return t
}

// parse parses a complete expression and rejects trailing tokens.
func (p *exprParser) parse() (*exprNode, error) {
	if p.peek().kind == tokEOF {
		return nil, newExprError(RuleExprSyntax, "empty expression", 0)
	}
	n, err := p.parseExpr()
	if err != nil {
		return nil, err
	}
	if p.peek().kind != tokEOF {
		return nil, newExprError(RuleExprSyntax, "unexpected trailing token", p.peek().pos)
	}
	return n, nil
}

// parseExpr handles composition, the lowest-precedence operator.
func (p *exprParser) parseExpr() (*exprNode, error) {
	left, err := p.parseTerm()
	if err != nil {
		return nil, err
	}
	for p.peek().kind == tokCompose {
		p.next()
		right, err := p.parseTerm()
		if err != nil {
			return nil, err
		}
		left = &exprNode{kind: nodeCompose, left: left, right: right}
	}
	return left, nil
}

// parseTerm handles the concurrent product.
func (p *exprParser) parseTerm() (*exprNode, error) {
	left, err := p.parseFactor()
	if err != nil {
		return nil, err
	}
	for p.peek().kind == tokProduct {
		p.next()
		right, err := p.parseFactor()
		if err != nil {
			return nil, err
		}
		left = &exprNode{kind: nodeProduct, left: left, right: right}
	}
	return left, nil
}

// parseFactor handles grouping, operands and the postfix dual suffix.
func (p *exprParser) parseFactor() (*exprNode, error) {
	var n *exprNode
	if p.peek().kind == tokLParen {
		p.depth++
		if p.depth > MaxExprDepth {
			return nil, newExprError(RuleExprDepth,
				"expression nesting exceeds 64", p.peek().pos)
		}
		p.next()
		inner, err := p.parseExpr()
		if err != nil {
			return nil, err
		}
		if p.peek().kind != tokRParen {
			return nil, newExprError(RuleExprSyntax,
				"missing closing parenthesis", p.peek().pos)
		}
		p.next()
		p.depth--
		n = inner
	} else {
		operand, err := p.parseOperand()
		if err != nil {
			return nil, err
		}
		n = operand
	}
	for p.peek().kind == tokInverse {
		p.next()
		n = &exprNode{kind: nodeInverse, left: n}
	}
	return n, nil
}

// parseOperand consumes an identity, an explicit reference or an inline id.
func (p *exprParser) parseOperand() (*exprNode, error) {
	t := p.peek()
	switch t.kind {
	case tokIdentity:
		p.next()
		return &exprNode{kind: nodeIdentity, pos: t.pos}, nil
	case tokRef, tokBare:
		p.next()
		return &exprNode{
			kind: nodeOperand, text: t.text, ref: t.kind == tokRef, pos: t.pos,
		}, nil
	default:
		return nil, newExprError(RuleExprSyntax, "expected an operand", t.pos)
	}
}
