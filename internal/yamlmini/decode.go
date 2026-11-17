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
		extra, err := appendEscape(&b, body, i, line)
		if err != nil {
			return "", err
		}
		i += extra
	}
	return b.String(), nil
}

// appendEscape writes the escape whose marker byte sits at body[i] into b. It
// returns the extra bytes consumed beyond the marker; a surrogate pair uses
// six more bytes for the low-surrogate escape.
func appendEscape(b *strings.Builder, body string, i int, line int) (int, error) {
	c := body[i]
	switch c {
	case '0', 'a', 'b', 't', '\t', 'n', 'v', 'f', 'r', 'e', ' ', '"', '/', '\\':
		b.WriteByte(escapedByte(c))
		return 0, nil
	case 'N', '_', 'L', 'P':
		_, _ = b.WriteRune(escapedRune(c))
		return 0, nil
	case 'x', 'u', 'U':
		return appendUnicodeEscape(b, body, i, line)
	default:
		return 0, fmt.Errorf("yamlmini: line %d: unknown escape \\%c", line, c)
	}
}

func escapedByte(c byte) byte {
	switch c {
	case '0':
		return 0
	case 'a':
		return '\a'
	case 'b':
		return '\b'
	case 't', '\t':
		return '\t'
	case 'n':
		return '\n'
	case 'v':
		return '\v'
	case 'f':
		return '\f'
	case 'r':
		return '\r'
	case 'e':
		return 0x1b
	default:
		return c // ' ', '"', '/', '\\' pass through as themselves
	}
}

func escapedRune(c byte) rune {
	switch c {
	case 'N':
		return '\u0085'
	case '_':
		return '\u00a0'
	case 'L':
		return '\u2028'
	default:
		return '\u2029' // 'P'
	}
}

func appendUnicodeEscape(b *strings.Builder, body string, i int, line int) (int, error) {
	marker := body[i]
	n := 2
	switch marker {
	case 'u':
		n = 4
	case 'U':
		n = 8
	}
	if i+n >= len(body) {
		return 0, fmt.Errorf("yamlmini: line %d: truncated unicode escape", line)
	}
	v, err := strconv.ParseUint(body[i+1:i+1+n], 16, 32)
	if err != nil {
		return 0, fmt.Errorf("yamlmini: line %d: invalid unicode escape: %w", line, err)
	}
	i += n
	if marker == 'U' && v > utf8.MaxRune {
		return 0, fmt.Errorf("yamlmini: line %d: invalid code point", line)
	}
	if v < 0xD800 || v > 0xDFFF {
		_, _ = b.WriteRune(rune(v))
		return n, nil
	}
	extra, err := appendSurrogate(b, body, i, line, marker, v)
	if err != nil {
		return 0, err
	}
	return n + extra, nil
}

func appendSurrogate(b *strings.Builder, body string, i int, line int, marker byte, high uint64) (int, error) {
	if marker == 'u' && i+6 < len(body) && body[i+1] == '\\' && body[i+2] == 'u' {
		lo, err := strconv.ParseUint(body[i+3:i+7], 16, 32)
		if err == nil && lo >= 0xDC00 && lo <= 0xDFFF {
			_, _ = b.WriteRune(utf16.DecodeRune(rune(high), rune(lo)))
			return 6, nil
		}
	}
	return 0, fmt.Errorf("yamlmini: line %d: lone surrogate", line)
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
