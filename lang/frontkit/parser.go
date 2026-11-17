/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package frontkit

import (
	"fmt"

	"abate/lang/air"
)

// maxParseDepth bounds expression nesting so deeply nested input returns a
// classified error instead of exhausting the goroutine stack.
const maxParseDepth = 64

// ParseRule and PrattParser form the shared precedence-climbing skeleton. All
// binding powers are integer data; no floating point value, clock or randomized
// state takes part in a reduction.
type (
	// ParseRule holds the deterministic binding powers and reduction hooks for
	// one token kind or one exact token spelling.
	ParseRule struct {
		// PrefixPower is the minimum power a prefix hook passes to
		// ParsePrefixOperand while reducing its operand.
		PrefixPower int
		// InfixPower is the power used to decide whether an infix token binds
		// at the caller's current minimum power.
		InfixPower int
		// RightAssoc makes an infix hook parse its right operand at the same
		// power rather than one level tighter.
		RightAssoc bool
		// Prefix reduces a token that opens an expression.
		Prefix func(*PrattParser, Token) (air.Node, error)
		// Infix reduces an operator with its already-reduced left and right
		// operands, so recursion depth stays under the parser's control.
		Infix func(*PrattParser, air.Node, Token, air.Node) (air.Node, error)
		// Postfix reduces a token that follows an already-reduced left
		// operand, such as a subscript. It is optional: a rule without a
		// Postfix hook never consumes a postfix token, so adding the hook
		// leaves every existing rule table unchanged.
		Postfix func(*PrattParser, air.Node, Token) (air.Node, error)
	}

	// PrattParser reduces an expression with a table-driven precedence climb.
	// The rule table is an array indexed by the frozen TokenKind enum, so no
	// map traversal can influence a reduction or an emitted node.
	PrattParser struct {
		stream    *TokenStream
		view      *View
		rules     [tokenKindCount]ParseRule
		textRules []textRule
		depth     int
	}
)

// textRule binds one exact token spelling to a more specific rule.
type textRule struct {
	kind TokenKind
	text string
	rule ParseRule
}

// NewPrattParser binds a token cursor and a node-building view into a parser.
func NewPrattParser(stream *TokenStream, view *View) (*PrattParser, error) {
	if stream == nil || view == nil {
		return nil, fmt.Errorf("%w: parser needs a stream and a view", ErrUnknownToken)
	}
	return &PrattParser{
		stream: stream, view: view, textRules: []textRule{},
	}, nil
}

// Register binds a rule to a whole token kind.
func (p *PrattParser) Register(kind TokenKind, rule ParseRule) error {
	if p == nil || !kind.Valid() {
		return fmt.Errorf("%w: invalid token kind for rule", ErrUnknownToken)
	}
	p.rules[int(kind)] = rule
	return nil
}

// RegisterText binds a rule to one exact spelling, overriding the kind rule.
// Registrations are kept in an ordered slice and resolved by a linear scan in
// registration order, so priority never depends on map iteration.
func (p *PrattParser) RegisterText(kind TokenKind, text string, rule ParseRule) error {
	if p == nil || !kind.Valid() || text == "" {
		return fmt.Errorf("%w: invalid text rule", ErrUnknownToken)
	}
	p.textRules = append(p.textRules, textRule{kind: kind, text: text, rule: rule})
	return nil
}

// ParseExpression reduces one expression whose operators bind at minPower or
// tighter. Prefix and infix hooks build AIR nodes directly; no language-specific
// expression type is ever materialized.
func (p *PrattParser) ParseExpression(minPower int) (air.Node, error) {
	if p == nil || p.stream == nil {
		return nil, fmt.Errorf("%w: nil parser", ErrUnknownToken)
	}
	if p.depth >= maxParseDepth {
		return nil, fmt.Errorf("%w: depth %d", ErrParseDepth, p.depth)
	}
	p.depth++
	defer func() { p.depth-- }()

	head, ok := p.stream.Advance()
	if !ok {
		return nil, fmt.Errorf("%w: unexpected end of token stream", ErrUnknownToken)
	}
	rule := p.rulesFor(head)
	if rule.Prefix == nil {
		return nil, fmt.Errorf("%w: no prefix rule for %s %q",
			ErrUnknownToken, head.Kind.String(), head.Text)
	}
	left, err := rule.Prefix(p, head)
	if err != nil {
		return nil, err
	}
	for {
		next, available := p.stream.Peek()
		if !available || next.Kind == TokenEOF {
			break
		}
		rule := p.rulesFor(next)
		if rule.Postfix != nil {
			if _, advanced := p.stream.Advance(); !advanced {
				return nil, fmt.Errorf("%w: operator vanished", ErrUnknownToken)
			}
			left, err = rule.Postfix(p, left, next)
			if err != nil {
				return nil, err
			}
			continue
		}
		if rule.Infix == nil || rule.InfixPower < minPower {
			break
		}
		if _, advanced := p.stream.Advance(); !advanced {
			return nil, fmt.Errorf("%w: operator vanished", ErrUnknownToken)
		}
		rightPower := rule.InfixPower
		if !rule.RightAssoc {
			rightPower++
		}
		right, err := p.ParseExpression(rightPower)
		if err != nil {
			return nil, err
		}
		left, err = rule.Infix(p, left, next, right)
		if err != nil {
			return nil, err
		}
	}
	return left, nil
}

// Parse reduces a whole expression and rejects any trailing non-EOF token.
func (p *PrattParser) Parse() (air.Node, error) {
	root, err := p.ParseExpression(0)
	if err != nil {
		return nil, err
	}
	if next, ok := p.stream.Peek(); ok && next.Kind != TokenEOF {
		return nil, fmt.Errorf("%w: trailing %s %q",
			ErrUnknownToken, next.Kind.String(), next.Text)
	}
	return root, nil
}

// ParsePrefixOperand reduces the operand of a prefix token at its declared
// prefix binding power.
func (p *PrattParser) ParsePrefixOperand(tok Token) (air.Node, error) {
	return p.ParseExpression(p.rulesFor(tok).PrefixPower)
}

// Recover skips tokens until a synchronization point and returns the skipped
// count. It never panics on an exhausted stream and never rewinds past EOF.
func (p *PrattParser) Recover() int {
	if p == nil || p.stream == nil {
		return 0
	}
	skipped := 0
	for {
		tok, ok := p.stream.Peek()
		if !ok || isRecoveryPoint(tok) {
			return skipped
		}
		if _, advanced := p.stream.Advance(); !advanced {
			return skipped
		}
		skipped++
	}
}

// Depth returns the current expression nesting level.
func (p *PrattParser) Depth() int {
	if p == nil {
		return 0
	}
	return p.depth
}

// Stream returns the token cursor bound to the parser.
func (p *PrattParser) Stream() *TokenStream {
	if p == nil {
		return nil
	}
	return p.stream
}

// View returns the node-building view bound to the parser.
func (p *PrattParser) View() *View {
	if p == nil {
		return nil
	}
	return p.view
}

// rulesFor resolves the most specific rule for one token.
func (p *PrattParser) rulesFor(tok Token) ParseRule {
	for i := range p.textRules {
		if p.textRules[i].kind == tok.Kind && p.textRules[i].text == tok.Text {
			return p.textRules[i].rule
		}
	}
	if !tok.Kind.Valid() {
		return ParseRule{}
	}
	return p.rules[int(tok.Kind)]
}

// isRecoveryPoint reports whether a token is a safe place to resume parsing.
func isRecoveryPoint(tok Token) bool {
	return tok.Kind == TokenEOF || tok.Kind == TokenNewline ||
		tok.Kind == TokenDedent ||
		(tok.Kind == TokenPunct && (tok.Text == ";" || tok.Text == "}"))
}
