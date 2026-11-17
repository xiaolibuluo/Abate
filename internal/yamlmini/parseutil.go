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

// RejectUnsupportedScalar rejects plain scalars that begin with an indicator
// belonging to an unsupported YAML feature. Anchors, aliases, custom tags and
// folded block scalars are all rejected. Treating them as ordinary strings
// would be a silent downgrade, which this package forbids.
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
func splitKey(text string) (key string, rest string, ok bool) {
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
