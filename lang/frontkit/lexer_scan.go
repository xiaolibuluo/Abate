/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package frontkit

import (
	"bytes"
	"slices"
	"strings"
)

// scanNewline handles CR/LF line endings.
func (l *Lexer) scanNewline(out *[]Token, start, sl, sc int, ch byte) (bool, error) {
	if ch != '\n' && ch != '\r' {
		return false, nil
	}
	width := 1
	if ch == '\r' && l.off+1 < len(l.src) && l.src[l.off+1] == '\n' {
		width = 2
	}
	l.consume(width)
	l.atLineStart = true
	return true, l.emit(out, TokenNewline, start, sl, sc)
}

// scanWhitespace skips horizontal/vertical whitespace.
func (l *Lexer) scanWhitespace(ch byte) (bool, error) {
	if ch != ' ' && ch != '\t' && ch != '\f' && ch != '\v' {
		return false, nil
	}
	l.consume(1)
	return true, nil
}

// scanLineComment consumes a single-line comment.
func (l *Lexer) scanLineComment(out *[]Token, start, sl, sc int) (bool, error) {
	if len(l.lineComment) == 0 || !bytes.HasPrefix(l.src[l.off:], l.lineComment) {
		return false, nil
	}
	for l.off < len(l.src) && l.src[l.off] != '\n' && l.src[l.off] != '\r' {
		l.consume(1)
	}
	return true, l.emit(out, TokenComment, start, sl, sc)
}

// scanBlockComment consumes a block comment, honoring nested-agnostic pairing.
func (l *Lexer) scanBlockComment(out *[]Token, start, sl, sc int) (bool, error) {
	if len(l.blockOpening) == 0 || !bytes.HasPrefix(l.src[l.off:], l.blockOpening) {
		return false, nil
	}
	l.consume(len(l.blockOpening))
	for l.off < len(l.src) {
		if len(l.blockClosing) > 0 && bytes.HasPrefix(l.src[l.off:], l.blockClosing) {
			l.consume(len(l.blockClosing))
			break
		}
		l.consume(1)
	}
	return true, l.emit(out, TokenComment, start, sl, sc)
}

// scanString consumes a (possibly raw) string or character literal.
func (l *Lexer) scanString(out *[]Token, start, sl, sc int) (bool, error) {
	prefix, quote, raw := l.quoteMatch()
	if quote == "" {
		return false, nil
	}
	kind := TokenStr
	if quote == "'" {
		kind = TokenChar
	}
	delimiter := []byte(quote)
	l.consume(len(prefix) + len(delimiter))
	for l.off < len(l.src) {
		current := l.src[l.off]
		if !raw && current == '\\' && l.off+1 < len(l.src) {
			l.consume(2)
			continue
		}
		if bytes.HasPrefix(l.src[l.off:], delimiter) {
			l.consume(len(delimiter))
			break
		}
		if len(delimiter) == 1 && (current == '\n' || current == '\r') {
			break
		}
		l.consume(1)
	}
	return true, l.emit(out, kind, start, sl, sc)
}

// scanNumber consumes a numeric literal with optional separators.
func (l *Lexer) scanNumber(out *[]Token, start, sl, sc int) (bool, error) {
	if !isDigit(l.src[l.off]) {
		return false, nil
	}
	for l.off < len(l.src) {
		separator := 0
		for _, candidate := range l.numberSeparators {
			if candidate != "" && bytes.HasPrefix(l.src[l.off:], []byte(candidate)) {
				separator = len(candidate)
				break
			}
		}
		if separator > 0 {
			l.consume(separator)
			continue
		}
		current := l.src[l.off]
		if isIdentContinue(current) || current == '.' {
			l.consume(1)
			continue
		}
		break
	}
	return true, l.emit(out, TokenNum, start, sl, sc)
}

// scanIdentifier consumes an identifier or keyword token.
func (l *Lexer) scanIdentifier(out *[]Token, start, sl, sc int, ch byte, runeWidth int) (bool, error) {
	if !isIdentStart(ch) && (ch < 0x80 || runeWidth <= 1) {
		return false, nil
	}
	for l.off < len(l.src) && isIdentContinue(l.src[l.off]) {
		l.consume(1)
	}
	kind := TokenIdent
	if slices.Contains(l.keywords, string(l.src[start:l.off])) {
		kind = TokenKeyword
	}
	return true, l.emit(out, kind, start, sl, sc)
}

// scanOperator matches the longest registered operator prefix.
func (l *Lexer) scanOperator(out *[]Token, start, sl, sc int) (bool, error) {
	for _, operator := range l.operators {
		if operator == "" || !bytes.HasPrefix(l.src[l.off:], []byte(operator)) {
			continue
		}
		l.consume(len(operator))
		return true, l.emit(out, TokenOp, start, sl, sc)
	}
	return false, nil
}

// scanPunct emits a single punctuation character.
func (l *Lexer) scanPunct(out *[]Token, start, sl, sc int, ch byte) (bool, error) {
	if strings.IndexByte(lexPunct, ch) < 0 {
		return false, nil
	}
	l.consume(1)
	return true, l.emit(out, TokenPunct, start, sl, sc)
}

// scanIllegal consumes one ill-formed rune as an illegal token.
func (l *Lexer) scanIllegal(out *[]Token, start, sl, sc, runeWidth int) (bool, error) {
	width := runeWidth
	if width < 1 {
		width = 1
	}
	l.consume(width)
	return true, l.emit(out, TokenIllegal, start, sl, sc)
}
