/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package yamlmini

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf16"
	"unicode/utf8"
)

// Kind identifies the shape of a Node.
type Kind uint8

const (
	// Scalar is a quoted or plain scalar carried as text in Value.
	Scalar Kind = iota
	// Mapping is an ordered key/value map; Keys parallels Vals.
	Mapping
	// Sequence is an ordered list carried in Items.
	Sequence
	// Null is an empty value ("null", "~" or an absent value).
	Null
)

func (k Kind) String() string {
	switch k {
	case Scalar:
		return "scalar"
	case Mapping:
		return "mapping"
	case Sequence:
		return "sequence"
	case Null:
		return "null"
	default:
		return fmt.Sprintf("kind(%d)", uint8(k))
	}
}

// scalarKind records how a plain scalar was lexed so that the binder keeps the
// distinction between, for example, the string "42" and the integer 42.
type scalarKind uint8

const (
	sString scalarKind = iota
	sInt
	sFloat
	sBool
	sNull
)

// Node is one node of the ordered parse tree. Mappings keep source order in
// Keys/Vals; sequences use Items; scalars and nulls use Value and Sk.
type Node struct {
	Kind  Kind
	Value string
	Sk    scalarKind
	Keys  []string
	Vals  []*Node
	Items []*Node
	Line  int
}

func (n *Node) keyIndex(key string) int {
	if n == nil || n.Kind != Mapping {
		return -1
	}
	for i, k := range n.Keys {
		if k == key {
			return i
		}
	}
	return -1
}

func (n *Node) lookup(key string) (*Node, bool) {
	i := n.keyIndex(key)
	if i < 0 {
		return nil, false
	}
	return n.Vals[i], true
}

func nullNode(line int) *Node { return &Node{Kind: Null, Sk: sNull, Line: line} }

func scalarNode(value string, sk scalarKind, line int) *Node {
	return &Node{Kind: Scalar, Value: value, Sk: sk, Line: line}
}

func emptySequence(line int) *Node { return &Node{Kind: Sequence, Line: line} }
func emptyMapping(line int) *Node  { return &Node{Kind: Mapping, Line: line} }

// parseScalar decodes one scalar token. Quoted tokens always yield sString;
// plain tokens are classified as null/bool/int/float/string in that order.
func parseScalar(raw string, line int) (*Node, error) {
	if raw == "" {
		return nullNode(line), nil
	}
	switch raw[0] {
	case '"':
		s, err := decodeDoubleQuoted(raw, line)
		if err != nil {
			return nil, err
		}
		return scalarNode(s, sString, line), nil
	case '\'':
		s, err := decodeSingleQuoted(raw, line)
		if err != nil {
			return nil, err
		}
		return scalarNode(s, sString, line), nil
	}
	switch raw {
	case "null", "Null", "NULL", "~":
		return nullNode(line), nil
	case "true", "True", "TRUE":
		return scalarNode("true", sBool, line), nil
	case "false", "False", "FALSE":
		return scalarNode("false", sBool, line), nil
	case "[]":
		return emptySequence(line), nil
	case "{}":
		return emptyMapping(line), nil
	}
	if raw[0] == '[' {
		return parseFlowSequence(raw, line)
	}
	if integerLiteral(raw) {
		if _, err := strconv.ParseInt(raw, 10, 64); err == nil {
			return scalarNode(raw, sInt, line), nil
		}
		if _, err := strconv.ParseUint(raw, 10, 64); err == nil {
			return scalarNode(raw, sInt, line), nil
		}
	}
	if floatLiteral(raw) {
		if _, err := strconv.ParseFloat(raw, 64); err == nil {
			return scalarNode(raw, sFloat, line), nil
		}
	}
	return scalarNode(raw, sString, line), nil
}

// parseFlowSequence decodes a controlled single-line flow sequence whose items
// are all scalars (string/number/bool/null). Nested flow collections and
// multi-line sequences are out of scope and rejected by the caller.
func parseFlowSequence(raw string, line int) (*Node, error) {
	if len(raw) < 2 || raw[0] != '[' || raw[len(raw)-1] != ']' {
		return nil, fmt.Errorf("yamlmini: line %d: unterminated flow sequence", line)
	}
	body := raw[1 : len(raw)-1]
	if strings.TrimSpace(body) == "" {
		return emptySequence(line), nil
	}
	parts, err := splitFlowItems(body, line)
	if err != nil {
		return nil, err
	}
	s := &Node{Kind: Sequence, Line: line}
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			return nil, fmt.Errorf("yamlmini: line %d: empty element in flow sequence", line)
		}
		if part[0] == '[' || part[0] == '{' {
			return nil, fmt.Errorf("yamlmini: line %d: nested flow collections are not supported", line)
		}
		item, err := parseScalar(part, line)
		if err != nil {
			return nil, err
		}
		s.Items = append(s.Items, item)
	}
	return s, nil
}

// splitFlowItems splits a flow collection body on top-level commas, respecting
// single and double quotes. An unterminated quote or a nested collection is an
// error; the strictness here is what prevents an extension from becoming a
// silent widening of the grammar.
func splitFlowItems(body string, line int) ([]string, error) {
	var out []string
	var quote byte
	start := 0
	for i := 0; i < len(body); i++ {
		c := body[i]
		if quote != 0 {
			if quote == '"' && c == '\\' && i+1 < len(body) {
				i++
				continue
			}
			if c == quote {
				if quote == '\'' && i+1 < len(body) && body[i+1] == '\'' {
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
		if c == '[' || c == '{' {
			return nil, fmt.Errorf("yamlmini: line %d: nested flow collections are not supported", line)
		}
		if c == ',' {
			out = append(out, body[start:i])
			start = i + 1
		}
	}
	if quote != 0 {
		return nil, fmt.Errorf("yamlmini: line %d: unterminated quote in flow sequence", line)
	}
	out = append(out, body[start:])
	return out, nil
}

// integerLiteral reports whether s has the shape of a base-10 integer. YAML
// 1.1 forms such as 0x1f or 1_000 are deliberately not recognised.
func integerLiteral(s string) bool {
	if s == "" {
		return false
	}
	i := 0
	if s[0] == '-' || s[0] == '+' {
		i++
	}
	if i == len(s) {
		return false
	}
	for ; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return true
}

// floatLiteral is a conservative pre-filter: the actual decision is always
// made by strconv.ParseFloat, this only avoids misclassifying ordinary text.
func floatLiteral(s string) bool {
	if s == "" {
		return false
	}
	seenDigit, seenDot, seenExp := false, false, false
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch {
		case c >= '0' && c <= '9':
			seenDigit = true
		case c == '+' || c == '-':
			if i != 0 && s[i-1] != 'e' && s[i-1] != 'E' {
				return false
			}
		case c == '.':
			if seenDot || seenExp {
				return false
			}
			seenDot = true
		case c == 'e' || c == 'E':
			if seenExp || !seenDigit {
				return false
			}
			seenExp = true
		default:
			return false
		}
	}
	return seenDigit && (seenDot || seenExp)
}

// decodeDoubleQuoted unquotes a YAML double-quoted scalar. A trailing comment
// has already been stripped by the lexer, so raw must end with a closing '"'.
func decodeDoubleQuoted(raw string, line int) (string, error) {
	if len(raw) < 2 || raw[len(raw)-1] != '"' {
		return "", fmt.Errorf("yamlmini: line %d: unterminated double-quoted scalar", line)
	}
	body := raw[1 : len(raw)-1]
	var b strings.Builder
	for i := 0; i < len(body); i++ {
		c := body[i]
		if c != '\\' {
			b.WriteByte(c)
			continue
		}
		i++
		if i >= len(body) {
			return "", fmt.Errorf("yamlmini: line %d: dangling escape", line)
		}
		switch body[i] {
		case '0':
			b.WriteByte(0)
		case 'a':
			b.WriteByte('\a')
		case 'b':
			b.WriteByte('\b')
		case 't', '\t':
			b.WriteByte('\t')
		case 'n':
			b.WriteByte('\n')
		case 'v':
			b.WriteByte('\v')
		case 'f':
			b.WriteByte('\f')
		case 'r':
			b.WriteByte('\r')
		case 'e':
			b.WriteByte(0x1b)
		case ' ', '"', '/', '\\':
			b.WriteByte(body[i])
		case 'N':
			b.WriteRune('\u0085')
		case '_':
			b.WriteRune('\u00a0')
		case 'L':
			b.WriteRune('\u2028')
		case 'P':
			b.WriteRune('\u2029')
		case 'x', 'u', 'U':
			n := 2
			if body[i] == 'u' {
				n = 4
			} else if body[i] == 'U' {
				n = 8
			}
			if i+n >= len(body) {
				return "", fmt.Errorf("yamlmini: line %d: truncated unicode escape", line)
			}
			v, err := strconv.ParseUint(body[i+1:i+1+n], 16, 32)
			if err != nil {
				return "", fmt.Errorf("yamlmini: line %d: invalid unicode escape: %w", line, err)
			}
			i += n
			if body[i-n] == 'U' && v > utf8.MaxRune {
				return "", fmt.Errorf("yamlmini: line %d: invalid code point", line)
			}
			if v >= 0xD800 && v <= 0xDFFF {
				// Surrogate pairs are only meaningful when followed by \u low.
				if body[i-n] == 'u' && i+6 < len(body) && body[i+1] == '\\' && body[i+2] == 'u' {
					lo, err := strconv.ParseUint(body[i+3:i+7], 16, 32)
					if err == nil && lo >= 0xDC00 && lo <= 0xDFFF {
						r := utf16.DecodeRune(rune(v), rune(lo))
						b.WriteRune(r)
						i += 6
						continue
					}
				}
				return "", fmt.Errorf("yamlmini: line %d: lone surrogate", line)
			}
			b.WriteRune(rune(v))
		default:
			return "", fmt.Errorf("yamlmini: line %d: unknown escape \\%c", line, body[i])
		}
	}
	return b.String(), nil
}

// decodeSingleQuoted unquotes a YAML single-quoted scalar: the only escape is
// a doubled quote.
func decodeSingleQuoted(raw string, line int) (string, error) {
	if len(raw) < 2 || raw[len(raw)-1] != '\'' {
		return "", fmt.Errorf("yamlmini: line %d: unterminated single-quoted scalar", line)
	}
	body := raw[1 : len(raw)-1]
	var b strings.Builder
	for i := 0; i < len(body); i++ {
		if body[i] == '\'' {
			if i+1 >= len(body) || body[i+1] != '\'' {
				return "", fmt.Errorf("yamlmini: line %d: unexpected quote in single-quoted scalar", line)
			}
			b.WriteByte('\'')
			i++
			continue
		}
		b.WriteByte(body[i])
	}
	return b.String(), nil
}

// nodeToAny converts a node into the generic shape used by map[string]any
// destinations: map[string]any, []any, string, int, uint64, float64, bool or
// nil. This matches the type set the TCK reference driver type-asserts on.
func nodeToAny(n *Node) any {
	if n == nil {
		return nil
	}
	switch n.Kind {
	case Null:
		return nil
	case Mapping:
		m := make(map[string]any, len(n.Keys))
		for i, k := range n.Keys {
			m[k] = nodeToAny(n.Vals[i])
		}
		return m
	case Sequence:
		s := make([]any, 0, len(n.Items))
		for _, item := range n.Items {
			s = append(s, nodeToAny(item))
		}
		return s
	default:
		switch n.Sk {
		case sInt:
			if v, err := strconv.ParseInt(n.Value, 10, 64); err == nil {
				return int(v)
			}
			if v, err := strconv.ParseUint(n.Value, 10, 64); err == nil {
				return v
			}
		case sFloat:
			if v, err := strconv.ParseFloat(n.Value, 64); err == nil {
				return v
			}
		case sBool:
			return n.Value == "true"
		}
		return n.Value
	}
}
