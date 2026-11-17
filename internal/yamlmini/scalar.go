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
)

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
