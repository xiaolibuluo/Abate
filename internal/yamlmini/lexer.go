/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package yamlmini

import (
	"bytes"
	"fmt"
	"strings"
)

// line is one logical input line after indentation measurement and comment
// stripping. text excludes leading indentation; indent is the leading space
// count. blank lines are retained because they may belong to a block scalar.
type line struct {
	text   string
	indent int
	num    int
	blank  bool
}

// lexLines splits data into logical lines, strips comments that appear outside
// quotes and rejects tabs used as indentation. CRLF is normalised to LF.
func lexLines(data []byte) ([]line, error) {
	if bytes.Contains(data, []byte("\x00")) {
		return nil, fmt.Errorf("yamlmini: NUL byte in document")
	}
	text := strings.ReplaceAll(string(data), "\r\n", "\n")
	raw := strings.Split(text, "\n")
	out := make([]line, 0, len(raw))
	for i, s := range raw {
		num := i + 1
		indent := 0
		for indent < len(s) && s[indent] == ' ' {
			indent++
		}
		if indent < len(s) && s[indent] == '\t' {
			return nil, fmt.Errorf("yamlmini: line %d: tab in indentation", num)
		}
		content := stripComment(s[indent:])
		trimmed := strings.TrimRight(content, " \t")
		// Multi-document markers are explicitly out of scope; fail rather than
		// silently mis-lexing them as a scalar.
		if trimmed == "---" || strings.HasPrefix(trimmed, "--- ") ||
			trimmed == "..." || strings.HasPrefix(trimmed, "... ") {
			return nil, fmt.Errorf("yamlmini: line %d: multi-document markers are not supported", num)
		}
		// A complex key ("? key") is outside the string-key whitelist.
		if trimmed == "?" || strings.HasPrefix(trimmed, "? ") {
			return nil, fmt.Errorf("yamlmini: line %d: complex keys are not supported", num)
		}
		out = append(out, line{
			text:   trimmed,
			indent: indent,
			num:    num,
			blank:  strings.TrimSpace(trimmed) == "",
		})
	}
	return out, nil
}

// stripComment removes a trailing "# ..." comment. A '#' starts a comment only
// at the beginning of the content or after whitespace, and never inside quotes.
func stripComment(s string) string {
	var quote byte
	for i := 0; i < len(s); i++ {
		c := s[i]
		if quote != 0 {
			if quote == '"' && c == '\\' && i+1 < len(s) {
				i++ // skip the escaped byte
				continue
			}
			if c == quote {
				if quote == '\'' && i+1 < len(s) && s[i+1] == '\'' {
					i++ // doubled quote inside a single-quoted scalar
					continue
				}
				quote = 0
			}
			continue
		}
		if c == '"' || c == '\'' {
			quote = c
			continue
		}
		if c == '#' && (i == 0 || s[i-1] == ' ' || s[i-1] == '\t') {
			return s[:i]
		}
	}
	return s
}

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

func (p *parser) errf(l line, format string, args ...any) error {
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

// rejectUnsupportedScalar fails explicitly on plain scalars that begin with a
// YAML indicator belonging to an unsupported feature: anchors (&), aliases
// (*), custom tags (!) and folded block scalars (>). Treating them as ordinary
// strings would be a silent downgrade, which this package forbids.
func rejectUnsupportedScalar(rest string, ln line) error {
	if rest == "" {
		return nil
	}
	switch rest[0] {
	case '&':
		return fmt.Errorf("yamlmini: line %d: anchors are not supported", ln.num)
	case '*':
		return fmt.Errorf("yamlmini: line %d: aliases are not supported", ln.num)
	case '!':
		return fmt.Errorf("yamlmini: line %d: custom tags are not supported", ln.num)
	case '>':
		return fmt.Errorf("yamlmini: line %d: folded block scalars are not supported", ln.num)
	case '[':
		if rest == "[]" {
			return nil // controlled empty-flow-sequence extension
		}
		// Controlled single-line scalar flow sequence. Pre-validate here so an
		// unterminated "[unterminated" still fails (never silently widened).
		if !strings.HasSuffix(rest, "]") {
			return fmt.Errorf("yamlmini: line %d: unterminated flow sequence", ln.num)
		}
		if _, err := splitFlowItems(rest[1:len(rest)-1], ln.num); err != nil {
			return err
		}
		return nil
	case '{':
		if rest == "{}" {
			return nil // controlled empty-flow-mapping extension
		}
		return fmt.Errorf("yamlmini: line %d: flow mappings are not supported", ln.num)
	}
	return nil
}

func isBlockHeader(rest string) bool {
	if rest == "" || rest[0] != '|' {
		return false
	}
	for _, c := range rest[1:] {
		if c != '-' && c != '+' {
			return false
		}
	}
	return true
}

// parseBlockScalar collects all following lines indented deeper than parent.
// The common indent of the first non-blank line is stripped while preserving
// deeper relative indentation inside the block.
func (p *parser) parseBlockScalar(parent int, header line, rest string) (*Node, error) {
	var collected []line
	base := -1
	for p.pos < len(p.lines) {
		l := p.lines[p.pos]
		if !l.blank && l.indent <= parent {
			break
		}
		if !l.blank && base < 0 {
			base = l.indent
		}
		collected = append(collected, l)
		p.pos++
	}
	if base < 0 {
		base = parent + 1
	}
	var b strings.Builder
	for _, l := range collected {
		if l.blank {
			b.WriteByte('\n')
			continue
		}
		if l.indent < base {
			return nil, p.errf(l, "block scalar line is less indented than its content")
		}
		b.WriteString(strings.Repeat(" ", l.indent-base))
		b.WriteString(l.text)
		b.WriteByte('\n')
	}
	text := b.String()
	if strings.Contains(rest, "-") {
		text = strings.TrimRight(text, "\n")
	} else if !strings.Contains(rest, "+") {
		// "clip": trim trailing blank lines but keep one final newline.
		text = strings.TrimRight(text, "\n")
		if text != "" {
			text += "\n"
		}
	}
	return scalarNode(text, sString, header.num), nil
}

// splitableKey reports whether text looks like a mapping pair so that a "- "
// sequence item can decide between a scalar item and an inline mapping.
func splitableKey(text string) bool {
	_, _, ok := splitKey(text)
	return ok
}

// splitKey splits "key: value" at the first colon followed by a space or end
// of line, ignoring colons inside quotes. It returns the unquoted key and the
// trimmed raw value.
func splitKey(text string) (string, string, bool) {
	var quote byte
	for i := 0; i < len(text); i++ {
		c := text[i]
		if quote != 0 {
			if quote == '"' && c == '\\' && i+1 < len(text) {
				i++
				continue
			}
			if c == quote {
				if quote == '\'' && i+1 < len(text) && text[i+1] == '\'' {
					i++
					continue
				}
				quote = 0
			}
			continue
		}
		if c == '"' || c == '\'' {
			quote = c
			continue
		}
		if c == ':' && (i+1 == len(text) || text[i+1] == ' ' || text[i+1] == '\t') {
			key := strings.TrimSpace(text[:i])
			if key == "" {
				return "", "", false
			}
			// Unquote a quoted key using the scalar decoder when possible.
			if key[0] == '"' && strings.HasSuffix(key, "\"") {
				if decoded, err := decodeDoubleQuoted(key, 0); err == nil {
					key = decoded
				}
			} else if key[0] == '\'' && strings.HasSuffix(key, "'") {
				if decoded, err := decodeSingleQuoted(key, 0); err == nil {
					key = decoded
				}
			}
			return key, strings.TrimSpace(text[i+1:]), true
		}
	}
	return "", "", false
}
