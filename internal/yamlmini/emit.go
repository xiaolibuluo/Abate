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

// emitNode writes a node tree with a fixed two-space indentation. Scalars and
// empty collections are written inline; non-empty collections are expanded.
func emitNode(b *strings.Builder, n *Node, indent int) {
	pad := strings.Repeat(" ", indent)
	switch n.Kind {
	case Mapping:
		for i, key := range n.Keys {
			v := n.Vals[i]
			b.WriteString(pad)
			b.WriteString(emitKey(key))
			b.WriteByte(':')
			switch {
			case isInline(v):
				b.WriteByte(' ')
				b.WriteString(emitScalar(v))
				b.WriteByte('\n')
			default:
				b.WriteByte('\n')
				emitNode(b, v, indent+2)
			}
		}
	case Sequence:
		for _, item := range n.Items {
			b.WriteString(pad)
			b.WriteByte('-')
			switch {
			case isInline(item):
				b.WriteByte(' ')
				b.WriteString(emitScalar(item))
				b.WriteByte('\n')
			default:
				b.WriteByte('\n')
				emitNode(b, item, indent+2)
			}
		}
	case Scalar, Null:
		b.WriteString(pad)
		b.WriteString(emitScalar(n))
		b.WriteByte('\n')
	}
}

func isInline(n *Node) bool {
	if n == nil {
		return true
	}
	if n.Kind == Scalar || n.Kind == Null {
		return true
	}
	if n.Kind == Sequence && len(n.Items) == 0 {
		return true
	}
	return n.Kind == Mapping && len(n.Keys) == 0
}

func emitScalar(n *Node) string {
	if n == nil || n.Kind == Null {
		return "null"
	}
	if n.Kind != Scalar {
		if n.Kind == Sequence {
			return "[]"
		}
		return "{}"
	}
	if n.Sk == sString && (!plainSafe(n.Value) || looksNonString(n.Value)) {
		return quoteString(n.Value)
	}
	return n.Value
}

func emitKey(key string) string {
	if plainSafe(key) && !looksNonString(key) {
		return key
	}
	return quoteString(key)
}

func quoteString(s string) string {
	var b strings.Builder
	b.WriteByte('"')
	for _, r := range s {
		switch r {
		case '"':
			b.WriteString(`\"`)
		case '\\':
			b.WriteString(`\\`)
		case '\n':
			b.WriteString(`\n`)
		case '\r':
			b.WriteString(`\r`)
		case '\t':
			b.WriteString(`\t`)
		case '\b':
			b.WriteString(`\b`)
		case '\f':
			b.WriteString(`\f`)
		default:
			if r < 0x20 || r == 0x7f {
				_, _ = fmt.Fprintf(&b, `\u%04X`, r)
			} else {
				_, _ = b.WriteRune(r)
			}
		}
	}
	b.WriteByte('"')
	return b.String()
}

// plainSafe reports whether s can be emitted as an unquoted plain scalar that
// will round-trip back to the same string. Anything ambiguous is quoted.
func plainSafe(s string) bool {
	if s == "" || strings.TrimSpace(s) != s {
		return false
	}
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c < 0x20 || c == 0x7f {
			return false
		}
	}
	if strings.ContainsAny(s[:1], "-?:,[]{}#&*!|>'\"%@`~") {
		return false
	}
	if strings.HasPrefix(s, "- ") || strings.HasPrefix(s, "? ") || strings.HasPrefix(s, ": ") {
		return false
	}
	if strings.Contains(s, ": ") || strings.Contains(s, " #") || strings.HasSuffix(s, ":") {
		return false
	}
	return true
}

// looksNonString reports whether a plain scalar would be re-lexed as a
// non-string value, in which case it must be quoted to preserve its type.
func looksNonString(s string) bool {
	switch s {
	case "null", "Null", "NULL", "~", "true", "True", "TRUE", "false", "False", "FALSE":
		return true
	}
	if integerLiteral(s) {
		return true
	}
	if floatLiteral(s) {
		if _, err := strconv.ParseFloat(s, 64); err == nil {
			return true
		}
	}
	return false
}
