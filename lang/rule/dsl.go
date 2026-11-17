/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package rule

import (
	"fmt"
	"regexp"

	"abate/lang/air"
)

// patternNode is one compiled vertex of a structural pattern.
type patternNode struct {
	kind     string
	attrOp   string
	attrVal  string
	regex    *regexp.Regexp
	semKey   air.SemKey
	capture  string
	wild     bool
	ellipsis bool
	repeat   bool
	neg      bool
	inner    *patternNode
	children []*patternNode
}

// patternParser is a single-use recursive descent parser over pattern tokens.
type patternParser struct {
	tokens []string
	pos    int
	depth  int
}

// tokenizePattern splits a pattern source into words and parentheses.
func tokenizePattern(src string) ([]string, error) {
	out := make([]string, 0, 8)
	for i := 0; i < len(src); {
		switch {
		case isPatternSpace(src[i]):
			i++
		case src[i] == '(' || src[i] == ')':
			out = append(out, string(src[i]))
			i++
		case src[i] == '"':
			return nil, fmt.Errorf("%w: quotes are not supported", ErrBadPattern)
		default:
			j := i
			for j < len(src) && !isPatternSpace(src[j]) &&
				src[j] != '(' && src[j] != ')' && src[j] != '"' {
				j++
			}
			out = append(out, src[i:j])
			i = j
		}
	}
	return out, nil
}

// isPatternSpace reports whether a byte separates pattern tokens.
func isPatternSpace(c byte) bool {
	return c == ' ' || c == '\t' || c == '\n' || c == '\r'
}

// parsePatternSource compiles a whole pattern into its root vertex.
func parsePatternSource(src string) (*patternNode, error) {
	tokens, err := tokenizePattern(src)
	if err != nil {
		return nil, err
	}
	if len(tokens) == 0 {
		return nil, fmt.Errorf("%w: empty pattern", ErrBadPattern)
	}
	parser := &patternParser{tokens: tokens}
	root, err := parser.parseExpr()
	if err != nil {
		return nil, err
	}
	if parser.pos != len(tokens) {
		return nil, fmt.Errorf("%w: trailing tokens", ErrBadPattern)
	}
	if root.ellipsis || root.repeat || root.neg {
		return nil, fmt.Errorf("%w: invalid root vertex", ErrBadPattern)
	}
	if err := validatePatternTree(root, 0); err != nil {
		return nil, err
	}
	return root, nil
}

// parseExpr parses either a bare word or a parenthesized node with children.
func (p *patternParser) parseExpr() (*patternNode, error) {
	if p.depth >= maxPatternDepth {
		return nil, fmt.Errorf("%w: pattern too deep", ErrBadPattern)
	}
	token, ok := p.advance()
	if !ok || token == ")" {
		return nil, fmt.Errorf("%w: expected a pattern vertex", ErrBadPattern)
	}
	if token != "(" {
		return parseWord(token)
	}
	head, ok := p.advance()
	if !ok || head == "(" || head == ")" {
		return nil, fmt.Errorf("%w: missing node word", ErrBadPattern)
	}
	node, err := parseWord(head)
	if err != nil {
		return nil, err
	}
	p.depth++
	groupErr := p.parseChildren(node)
	p.depth--
	if groupErr != nil {
		return nil, groupErr
	}
	return node, nil
}

// parseChildren parses the child vertices of a parenthesized node.
func (p *patternParser) parseChildren(node *patternNode) error {
	for {
		next, present := p.peek()
		if !present {
			return fmt.Errorf("%w: unclosed group", ErrBadPattern)
		}
		if next == ")" {
			_, _ = p.advance()
			return nil
		}
		child, err := p.parseExpr()
		if err != nil {
			return err
		}
		node.children = append(node.children, child)
	}
}

// advance consumes and returns the next token.
func (p *patternParser) advance() (string, bool) {
	if p.pos >= len(p.tokens) {
		return "", false
	}
	token := p.tokens[p.pos]
	p.pos++
	return token, true
}

// peek returns the next token without consuming it.
func (p *patternParser) peek() (string, bool) {
	if p.pos >= len(p.tokens) {
		return "", false
	}
	return p.tokens[p.pos], true
}
