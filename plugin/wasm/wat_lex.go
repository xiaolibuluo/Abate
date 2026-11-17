/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package wasm

import "fmt"

// watTokKind enumerates the lexical tokens of the supported WAT subset.
type watTokKind int

const (
	tokLParen watTokKind = iota
	tokRParen
	tokIdent
	tokString
	tokInt
	tokEOF
)

// watToken is a single lexical token produced by the WAT lexer.
type watToken struct {
	kind  watTokKind
	text  string
	value int64
	line  int
}

// watLexer turns WAT source bytes into a stream of tokens.
type watLexer struct {
	src  []byte
	pos  int
	line int
}

// newWatLexer builds a lexer over the given source.
func newWatLexer(src []byte) *watLexer {
	return &watLexer{src: src, line: 1}
}

// next returns the next token, or tokEOF at end of input.
func (l *watLexer) next() (*watToken, error) {
	for l.pos < len(l.src) {
		c := l.src[l.pos]
		switch c {
		case '\n':
			l.line++
			l.pos++
		case ' ', '\t', '\r':
			l.pos++
		case ';':
			for l.pos < len(l.src) && l.src[l.pos] != '\n' {
				l.pos++
			}
		case '(', ')':
			l.pos++
			kind := tokLParen
			if c == ')' {
				kind = tokRParen
			}
			return &watToken{kind: kind, line: l.line}, nil
		case '"':
			return l.lexString()
		default:
			return l.lexAtom()
		}
	}
	return &watToken{kind: tokEOF, line: l.line}, nil
}

// lexString reads a double-quoted string token.
func (l *watLexer) lexString() (*watToken, error) {
	l.pos++
	start := l.pos
	for l.pos < len(l.src) && l.src[l.pos] != '"' {
		if l.src[l.pos] == '\\' {
			l.pos++
		}
		l.pos++
	}
	if l.pos >= len(l.src) {
		return nil, fmt.Errorf("wasm: unterminated string at line %d", l.line)
	}
	text := string(l.src[start:l.pos])
	l.pos++
	return &watToken{kind: tokString, text: text, line: l.line}, nil
}

// lexAtom reads an identifier or integer atom.
func (l *watLexer) lexAtom() (*watToken, error) {
	start := l.pos
	for l.pos < len(l.src) {
		c := l.src[l.pos]
		if c == ' ' || c == '\t' || c == '\r' || c == '\n' || c == '(' || c == ')' || c == '"' || c == ';' {
			break
		}
		l.pos++
	}
	raw := string(l.src[start:l.pos])
	if raw == "" {
		return nil, fmt.Errorf("wasm: unexpected character at line %d", l.line)
	}
	if v, ok := parseWatInt(raw); ok {
		return &watToken{kind: tokInt, value: v, line: l.line}, nil
	}
	return &watToken{kind: tokIdent, text: raw, line: l.line}, nil
}

// parseWatInt parses a decimal or hexadecimal integer literal.
func parseWatInt(raw string) (int64, bool) {
	if raw == "" {
		return 0, false
	}
	neg := false
	s := raw
	if s[0] == '-' {
		neg = true
		s = s[1:]
	}
	if len(s) >= 2 && s[0] == '0' && (s[1] == 'x' || s[1] == 'X') {
		s = s[2:]
		var v int64
		for _, c := range s {
			d := hexDigit(c)
			if d < 0 {
				return 0, false
			}
			v = v*16 + int64(d)
		}
		if neg {
			v = -v
		}
		return v, true
	}
	var v int64
	for _, c := range s {
		if c < '0' || c > '9' {
			return 0, false
		}
		v = v*10 + int64(c-'0')
	}
	if neg {
		v = -v
	}
	return v, true
}

// hexDigit returns the value of a hex digit, or -1 if not hexadecimal.
func hexDigit(c rune) int {
	switch {
	case c >= '0' && c <= '9':
		return int(c - '0')
	case c >= 'a' && c <= 'f':
		return int(c-'a') + 10
	case c >= 'A' && c <= 'F':
		return int(c-'A') + 10
	default:
		return -1
	}
}
