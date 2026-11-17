/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package rule

import (
	"strings"

	"abate/lang/air"
)

// Matcher matches one compiled structural pattern against AIR nodes.
//
// A Matcher is immutable and safe for concurrent read-only use: every match
// builds its own capture slice and never mutates the compiled pattern.
type Matcher struct {
	pattern Pattern
}

// NewMatcher returns a matcher for a compiled pattern.
func NewMatcher(pattern Pattern) Matcher {
	return Matcher{pattern: pattern}
}

// Pattern returns the compiled pattern backing the matcher.
func (m Matcher) Pattern() Pattern {
	return m.pattern
}

// Match returns the captures of a successful match at node, or nil when the
// pattern does not match. A pattern without explicit captures still yields one
// implicit "_match" capture, so a non-empty result always means a hit.
func (m Matcher) Match(node air.Node) []Capture {
	if node == nil || !m.pattern.Valid() {
		return nil
	}
	caps, ok := m.matchAt(node, m.pattern.root, nil)
	if !ok {
		return nil
	}
	if len(caps) == 0 {
		return []Capture{makeCapture("_match", node)}
	}
	return caps
}

// Matches reports whether the pattern matches node.
func (m Matcher) Matches(node air.Node) bool {
	return len(m.Match(node)) > 0
}

// matchAt matches one vertex and its ordered child patterns.
func (m Matcher) matchAt(node air.Node, p *patternNode,
	caps []Capture) ([]Capture, bool) {
	if node == nil || p == nil {
		return nil, false
	}
	if p.neg {
		_, hit := m.matchAt(node, p.inner, nil)
		if hit {
			return nil, false
		}
		return caps, true
	}
	if p.ellipsis || p.repeat || !vertexMatches(node, p) {
		return nil, false
	}
	out := caps
	if p.capture != "" {
		out = append(append([]Capture(nil), caps...), makeCapture(p.capture, node))
	}
	if len(p.children) == 0 {
		return out, true
	}
	candidates := childCandidates(node, p.children)
	return matchSequence(m, candidates, 0, p.children, 0, out)
}

// vertexMatches checks the kind, annotation and attribute constraints.
func vertexMatches(node air.Node, p *patternNode) bool {
	if p.kind != "" && p.kind != "ANY" && node.Kind().String() != p.kind {
		return false
	}
	if p.semKey != air.SemInvalid && node.Sem(p.semKey) == nil {
		return false
	}
	if p.attrOp == "" {
		return true
	}
	text := air.AttributeOf(node)
	switch p.attrOp {
	case "=":
		return text == p.attrVal
	case "!=":
		return text != p.attrVal
	case "~":
		return strings.Contains(text, p.attrVal)
	case "^":
		return strings.HasPrefix(text, p.attrVal)
	case "+":
		return strings.HasSuffix(text, p.attrVal)
	case "regex":
		return p.regex != nil && p.regex.MatchString(text)
	default:
		return false
	}
}

// childCandidates returns direct children, or all descendants when a deep
// ellipsis is present so depth matching can see grandchildren.
func childCandidates(node air.Node, pats []*patternNode) []air.Node {
	if !hasEllipsis(pats) {
		return node.Children()
	}
	out := make([]air.Node, 0, 16)
	air.Walk(node, func(current air.Node) bool {
		out = append(out, current)
		return true
	})
	if len(out) > 0 {
		out = out[1:]
	}
	return out
}

// hasEllipsis reports whether a child pattern list contains a deep marker.
func hasEllipsis(pats []*patternNode) bool {
	for _, p := range pats {
		if p.ellipsis {
			return true
		}
	}
	return false
}

// matchSequence matches an ordered child pattern list against candidates.
func matchSequence(m Matcher, cands []air.Node, ci int,
	pats []*patternNode, pi int, caps []Capture) ([]Capture, bool) {
	if pi >= len(pats) {
		return caps, true
	}
	p := pats[pi]
	switch {
	case p.ellipsis:
		for skip := 0; ci+skip <= len(cands); skip++ {
			if out, ok := matchSequence(m, cands, ci+skip, pats, pi+1, caps); ok {
				return out, true
			}
		}
		return nil, false
	case p.repeat:
		return matchRepeat(m, cands, ci, pats, pi, caps)
	case p.neg:
		for k := ci; k < len(cands); k++ {
			if _, hit := m.matchAt(cands[k], p.inner, nil); hit {
				return nil, false
			}
		}
		return matchSequence(m, cands, ci, pats, pi+1, caps)
	default:
		if ci >= len(cands) {
			return nil, false
		}
		out, ok := m.matchAt(cands[ci], p, caps)
		if !ok {
			return nil, false
		}
		return matchSequence(m, cands, ci+1, pats, pi+1, out)
	}
}

// matchRepeat consumes zero or more candidates into one repeated capture.
func matchRepeat(m Matcher, cands []air.Node, ci int,
	pats []*patternNode, pi int, caps []Capture) ([]Capture, bool) {
	atom := *pats[pi]
	atom.repeat = false
	if out, ok := matchSequence(m, cands, ci, pats, pi+1, caps); ok {
		return out, true
	}
	acc := append([]Capture(nil), caps...)
	for k := ci; k < len(cands); k++ {
		next, ok := m.matchAt(cands[k], &atom, acc)
		if !ok {
			break
		}
		acc = next
		if out, ok := matchSequence(m, cands, k+1, pats, pi+1, acc); ok {
			return out, true
		}
	}
	return nil, false
}

// makeCapture binds one AIR node to a metavariable.
func makeCapture(name string, node air.Node) Capture {
	pos := node.Pos()
	return Capture{
		Name: name, Kind: node.Kind().String(),
		Text: air.AttributeOf(node), Line: pos.Line, Column: pos.Column,
	}
}
