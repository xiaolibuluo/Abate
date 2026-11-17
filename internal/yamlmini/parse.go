/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package yamlmini

import (
	"fmt"
	"strings"
)

// parser builds an ordered node tree from logical lines. It is deliberately a
// recursive-descent parser over indentation, not a general YAML parser.
type parser struct {
	lines []line
	pos   int
}

func parseDocument(data []byte) (*Node, error) {
	lines, err := lexLines(data)
	if err != nil {
		return nil, err
	}
	p := &parser{lines: lines}
	p.skipBlank()
	if p.pos >= len(lines) {
		return nullNode(0), nil
	}
	doc, err := p.parseNode(lines[p.pos].indent)
	if err != nil {
		return nil, err
	}
	p.skipBlank()
	if p.pos < len(lines) {
		return nil, p.errf(lines[p.pos], "unexpected content after document")
	}
	return doc, nil
}

func (p *parser) skipBlank() {
	for p.pos < len(p.lines) && p.lines[p.pos].blank {
		p.pos++
	}
}

func (*parser) errf(l line, format string, args ...any) error {
	return fmt.Errorf("yamlmini: line %d: %s", l.num, fmt.Sprintf(format, args...))
}

func (p *parser) parseNode(indent int) (*Node, error) {
	p.skipBlank()
	if p.pos >= len(p.lines) {
		return nullNode(0), nil
	}
	l := p.lines[p.pos]
	if l.indent != indent {
		return nil, p.errf(l, "inconsistent indentation (want %d, got %d)", indent, l.indent)
	}
	if isSeqItem(l.text) {
		return p.parseSequence(indent)
	}
	return p.parseMapping(indent)
}

func isSeqItem(text string) bool {
	return text == "-" || strings.HasPrefix(text, "- ")
}

func (p *parser) parseMapping(indent int) (*Node, error) {
	m := &Node{Kind: Mapping, Line: p.lines[p.pos].num}
	for p.pos < len(p.lines) {
		p.skipBlank()
		if p.pos >= len(p.lines) {
			break
		}
		l := p.lines[p.pos]
		if l.indent < indent {
			break
		}
		if l.indent > indent {
			return nil, p.errf(l, "unexpected indentation in mapping")
		}
		if isSeqItem(l.text) {
			return nil, p.errf(l, "sequence item where a mapping key was expected")
		}
		p.pos++
		if err := p.addPair(m, l.text, indent, l); err != nil {
			return nil, err
		}
	}
	return m, nil
}

func (p *parser) parseSequence(indent int) (*Node, error) {
	s := &Node{Kind: Sequence, Line: p.lines[p.pos].num}
	for p.pos < len(p.lines) {
		p.skipBlank()
		if p.pos >= len(p.lines) {
			break
		}
		l := p.lines[p.pos]
		if l.indent < indent {
			break
		}
		if l.indent > indent {
			return nil, p.errf(l, "unexpected indentation in sequence")
		}
		if !isSeqItem(l.text) {
			return nil, p.errf(l, "expected a sequence item")
		}
		rest := strings.TrimSpace(strings.TrimPrefix(l.text, "-"))
		p.pos++
		var item *Node
		var err error
		switch {
		case isBlockHeader(rest):
			item, err = p.parseBlockScalar(indent, l, rest)
		case rest == "":
			p.skipBlank()
			if p.pos < len(p.lines) && p.lines[p.pos].indent > indent {
				item, err = p.parseNode(p.lines[p.pos].indent)
			} else {
				item = nullNode(l.num)
			}
		case splitableKey(rest):
			// "- key: value" starts an inline mapping whose following keys are
			// indented two columns past the dash (the "- " prefix).
			item, err = p.parseInlineMapping(rest, l)
		default:
			if err := rejectUnsupportedScalar(rest, l); err != nil {
				return nil, err
			}
			item, err = parseScalar(rest, l.num)
		}
		if err != nil {
			return nil, err
		}
		s.Items = append(s.Items, item)
	}
	return s, nil
}

// parseInlineMapping parses a mapping whose first pair appeared after a dash.
func (p *parser) parseInlineMapping(first string, seqLine line) (*Node, error) {
	m := &Node{Kind: Mapping, Line: seqLine.num}
	childIndent := seqLine.indent + 2
	if err := p.addPair(m, first, childIndent, seqLine); err != nil {
		return nil, err
	}
	for p.pos < len(p.lines) {
		p.skipBlank()
		if p.pos >= len(p.lines) {
			break
		}
		l := p.lines[p.pos]
		if l.indent < childIndent {
			break
		}
		if l.indent > childIndent {
			return nil, p.errf(l, "unexpected indentation in sequence mapping")
		}
		if isSeqItem(l.text) {
			return nil, p.errf(l, "unexpected sequence item in mapping")
		}
		p.pos++
		if err := p.addPair(m, l.text, childIndent, l); err != nil {
			return nil, err
		}
	}
	return m, nil
}

// addPair splits a "key: value" line and appends the parsed value to m. When
// the value is absent it parses a nested node from the following deeper lines,
// defaulting to null when there is none.
func (p *parser) addPair(m *Node, text string, indent int, ln line) error {
	key, rest, ok := splitKey(text)
	if !ok {
		return p.errf(ln, "expected 'key: value', got %q", text)
	}
	if key == "<<" {
		return p.errf(ln, "merge keys are not supported")
	}
	if m.keyIndex(key) >= 0 {
		return p.errf(ln, "duplicate key %q", key)
	}
	var val *Node
	var err error
	switch {
	case isBlockHeader(rest):
		val, err = p.parseBlockScalar(indent, ln, rest)
	case rest == "":
		p.skipBlank()
		if p.pos < len(p.lines) && p.lines[p.pos].indent > indent {
			val, err = p.parseNode(p.lines[p.pos].indent)
		} else {
			val = nullNode(ln.num)
		}
	default:
		if err := rejectUnsupportedScalar(rest, ln); err != nil {
			return err
		}
		val, err = parseScalar(rest, ln.num)
	}
	if err != nil {
		return err
	}
	m.Keys = append(m.Keys, key)
	m.Vals = append(m.Vals, val)
	return nil
}
