/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package frontkit

import (
	"bytes"
	"fmt"
	"slices"
	"unicode/utf8"
)

const lexTabWidth = 8
const lexPunct = "()[]{};,.:?@#$~\\`"

// Lexer turns source bytes into a deterministic token stream from normalized
// LexSpec tables, keeping integer cursors only.
type Lexer struct {
	src, lineComment, blockOpening, blockClosing []byte
	keywords, operators, stringQuotes            []string
	rawPrefixes, numberSeparators                []string
	indentSensitive, atLineStart                 bool
	off, line, column                            int
	indents                                      []int
}

// NewLexer builds a lexer over src for one lexical specification.
func NewLexer(spec LexSpec, src []byte) (*Lexer, error) {
	if spec == nil {
		return nil, fmt.Errorf("%w: lexical specification", ErrNilSpec)
	}
	keywords, operators := NormalizeTables(spec.Keywords(), spec.Operators())
	opening, closing := spec.BlockComment()
	return &Lexer{
		src: slices.Clone(src), keywords: keywords, operators: operators,
		lineComment: []byte(spec.LineComment()), blockOpening: []byte(opening),
		blockClosing: []byte(closing), stringQuotes: normalizeTable(spec.StringQuotes()),
		rawPrefixes:      normalizeTable(spec.RawStringPrefixes()),
		numberSeparators: normalizeTable(spec.NumberSeparators()),
		indentSensitive:  spec.IndentSensitive(),
	}, nil
}

// Tokenize returns source-ordered tokens ending in one zero-width EOF token.
func (l *Lexer) Tokenize() ([]Token, error) {
	if l == nil {
		return nil, fmt.Errorf("%w: nil lexer", ErrLexFailed)
	}
	l.off, l.line, l.column, l.atLineStart = 0, 1, 1, true
	l.indents = []int{0}
	if bytes.HasPrefix(l.src, []byte{0xEF, 0xBB, 0xBF}) {
		l.consume(3)
	}
	out := make([]Token, 0, 32)
	for l.off < len(l.src) {
		if err := l.scanNext(&out); err != nil {
			return nil, err
		}
	}
	if l.indentSensitive {
		if err := l.emitDedents(&out, 0); err != nil {
			return nil, err
		}
	}
	return appendChecked(out, Token{
		Kind: TokenEOF, Offset: l.off, EndOffset: l.off,
		Line: l.line, Column: l.column,
	})
}

// scanNext consumes exactly one lexical category from the current offset. It
// dispatches to the first handler that recognizes the leading bytes, returning
// immediately on the first match.
func (l *Lexer) scanNext(out *[]Token) error {
	if l.atLineStart && l.indentSensitive {
		l.atLineStart = false
		if err := l.scanIndent(out); err != nil {
			return err
		}
		if l.off >= len(l.src) {
			return nil
		}
	}
	start, sl, sc := l.off, l.line, l.column
	ch := l.src[l.off]
	_, runeWidth := utf8.DecodeRune(l.src[l.off:])
	if handled, err := l.scanNewline(out, start, sl, sc, ch); handled {
		return err
	}
	if handled, err := l.scanWhitespace(ch); handled {
		return err
	}
	if handled, err := l.scanLineComment(out, start, sl, sc); handled {
		return err
	}
	if handled, err := l.scanBlockComment(out, start, sl, sc); handled {
		return err
	}
	if handled, err := l.scanString(out, start, sl, sc); handled {
		return err
	}
	if handled, err := l.scanNumber(out, start, sl, sc); handled {
		return err
	}
	if handled, err := l.scanIdentifier(out, start, sl, sc, ch, runeWidth); handled {
		return err
	}
	if handled, err := l.scanOperator(out, start, sl, sc); handled {
		return err
	}
	if handled, err := l.scanPunct(out, start, sl, sc, ch); handled {
		return err
	}
	_, err := l.scanIllegal(out, start, sl, sc, runeWidth)
	return err
}

func (l *Lexer) consume(width int) {
	end := l.off + width
	if end > len(l.src) {
		end = len(l.src)
	}
	for l.off < end {
		current := l.src[l.off]
		previousWasCR := l.off > 0 && l.src[l.off-1] == '\r'
		l.off++
		switch current {
		case '\r':
			l.line++
			l.column = 1
		case '\n':
			if !previousWasCR {
				l.line++
			}
			l.column = 1
		default:
			l.column++
		}
	}
}

func (l *Lexer) emit(out *[]Token, kind TokenKind, start, sl, sc int) error {
	if start > l.off {
		start = l.off
	}
	next, err := appendChecked(*out, Token{
		Kind: kind, Text: string(l.src[start:l.off]),
		Offset: start, EndOffset: l.off, Line: sl, Column: sc,
	})
	if err != nil {
		return err
	}
	*out = next
	return nil
}

func (l *Lexer) scanIndent(out *[]Token) error {
	start, sl, sc := l.off, l.line, l.column
	depth, end := 0, l.off
	for end < len(l.src) {
		current := l.src[end]
		if current == ' ' {
			depth++
			end++
			continue
		}
		if current == '\t' {
			depth += lexTabWidth - depth%lexTabWidth
			end++
			continue
		}
		break
	}
	blank := end >= len(l.src)
	if !blank {
		rest := l.src[end:]
		blank = rest[0] == '\n' || rest[0] == '\r' ||
			(len(l.lineComment) > 0 && bytes.HasPrefix(rest, l.lineComment)) ||
			(len(l.blockOpening) > 0 && bytes.HasPrefix(rest, l.blockOpening))
	}
	l.consume(end - l.off)
	top := l.indents[len(l.indents)-1]
	if blank || depth == top {
		return nil
	}
	if depth > top {
		l.indents = append(l.indents, depth)
		return appendSynthetic(out, TokenIndent, start, sl, sc)
	}
	return l.emitDedents(out, depth)
}

func (l *Lexer) emitDedents(out *[]Token, target int) error {
	for len(l.indents) > 1 && l.indents[len(l.indents)-1] > target {
		l.indents = l.indents[:len(l.indents)-1]
		if err := appendSynthetic(out, TokenDedent, l.off, l.line, l.column); err != nil {
			return err
		}
	}
	return nil
}

func (l *Lexer) quoteMatch() (prefix, quote string, raw bool) {
	for _, candidate := range l.stringQuotes {
		if candidate != "" && bytes.HasPrefix(l.src[l.off:], []byte(candidate)) {
			return "", candidate, false
		}
	}
	for _, candidate := range l.rawPrefixes {
		at := l.off + len(candidate)
		if candidate == "" || at > len(l.src) ||
			!bytes.HasPrefix(l.src[l.off:], []byte(candidate)) {
			continue
		}
		for _, delimiter := range l.stringQuotes {
			if delimiter != "" && bytes.HasPrefix(l.src[at:], []byte(delimiter)) {
				return candidate, delimiter, true
			}
		}
	}
	return "", "", false
}
