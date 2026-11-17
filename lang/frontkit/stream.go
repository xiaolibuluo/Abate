/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package frontkit

import (
	"fmt"
	"slices"
)

// The two constants below bound speculative parsing depth and the per-file
// token count, turning hostile inputs into classified errors.
const (
	maxStreamDepth = 128
	maxFileTokens  = 1 << 16
)

// TokenStream is a one-way cursor over an immutable token slice. Backtracking
// happens only through explicit integer marks, so nothing depends on global
// state or on map traversal order. It is not safe for concurrent mutation, and
// none of the frontkit parsers share a stream between goroutines.
type TokenStream struct {
	tokens []Token
	index  int
	depth  int
}

// NewTokenStream validates and copies a token slice, then returns a cursor at
// its first token. Reversed offsets, unknown kinds and oversized streams are
// rejected with a classified error instead of panicking.
func NewTokenStream(tokens []Token) (*TokenStream, error) {
	if len(tokens) > maxFileTokens {
		return nil, fmt.Errorf("%w: %d tokens exceed the per-file ceiling",
			ErrLexFailed, len(tokens))
	}
	cloned := slices.Clone(tokens)
	for i := range cloned {
		if !cloned[i].Valid() {
			return nil, fmt.Errorf("%w: invalid token at %d", ErrUnknownToken, i)
		}
		if i > 0 && cloned[i].Offset < cloned[i-1].Offset {
			return nil, fmt.Errorf("%w: token %d is out of source order",
				ErrUnknownToken, i)
		}
	}
	return &TokenStream{tokens: cloned}, nil
}

// Peek returns the current token without moving the cursor.
func (s *TokenStream) Peek() (Token, bool) {
	if s == nil || s.index < 0 || s.index >= len(s.tokens) {
		return Token{}, false
	}
	return s.tokens[s.index], true
}

// Advance returns the current token and moves the cursor one position right.
func (s *TokenStream) Advance() (Token, bool) {
	tok, ok := s.Peek()
	if !ok {
		return Token{}, false
	}
	s.index++
	return tok, true
}

// Expect consumes the current token when its kind matches, and reports a
// classified error otherwise.
func (s *TokenStream) Expect(kind TokenKind) (Token, error) {
	tok, ok := s.Peek()
	if !ok {
		return Token{}, fmt.Errorf("%w: expected %s at end of stream",
			ErrUnknownToken, kind.String())
	}
	if tok.Kind != kind {
		return Token{}, fmt.Errorf("%w: expected %s, found %s",
			ErrUnknownToken, kind.String(), tok.Kind.String())
	}
	s.index++
	return tok, nil
}

// Save returns a restore mark and pushes one level of backtracking depth. The
// depth ceiling turns unbounded speculative parsing into a classified error.
func (s *TokenStream) Save() (int, error) {
	if s == nil {
		return 0, fmt.Errorf("%w: nil stream", ErrUnknownToken)
	}
	if s.depth >= maxStreamDepth {
		return 0, fmt.Errorf("%w: save depth %d", ErrParseDepth, s.depth)
	}
	s.depth++
	return s.index, nil
}

// Restore moves the cursor back to mark and pops one backtracking level.
func (s *TokenStream) Restore(mark int) error {
	if s == nil {
		return fmt.Errorf("%w: nil stream", ErrUnknownToken)
	}
	if mark < 0 || mark > len(s.tokens) {
		return fmt.Errorf("%w: mark %d out of range", ErrUnknownToken, mark)
	}
	if s.depth <= 0 {
		return fmt.Errorf("%w: no outstanding save", ErrUnknownToken)
	}
	s.depth--
	s.index = mark
	return nil
}

// Depth returns the number of outstanding Save marks.
func (s *TokenStream) Depth() int {
	if s == nil {
		return 0
	}
	return s.depth
}

// Remaining returns the number of tokens from the cursor to the end.
func (s *TokenStream) Remaining() int {
	if s == nil || s.index >= len(s.tokens) {
		return 0
	}
	return len(s.tokens) - s.index
}

// Done reports whether the cursor has consumed every token.
func (s *TokenStream) Done() bool {
	return s == nil || s.index >= len(s.tokens)
}

// Tokens returns a defensive copy of the immutable token slice.
func (s *TokenStream) Tokens() []Token {
	if s == nil {
		return []Token{}
	}
	return slices.Clone(s.tokens)
}

// isDigit reports whether one ASCII byte is a decimal digit.
func isDigit(ch byte) bool { return ch >= '0' && ch <= '9' }

// isIdentStart reports whether one ASCII byte may open an identifier.
func isIdentStart(ch byte) bool {
	return ch == '_' || (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

// isIdentContinue reports whether a byte may continue an identifier, including
// validated UTF-8 continuation bytes.
func isIdentContinue(ch byte) bool {
	return isIdentStart(ch) || isDigit(ch) || ch >= 0x80
}

// appendChecked appends one token and enforces the per-file token ceiling.
func appendChecked(out []Token, tok Token) ([]Token, error) {
	if len(out) >= maxFileTokens {
		return out, fmt.Errorf("%w: %d tokens exceed the per-file ceiling",
			ErrLexFailed, len(out)+1)
	}
	return append(out, tok), nil
}

// appendSynthetic appends a zero-width token carrying no source text.
func appendSynthetic(out *[]Token, kind TokenKind, at, sl, sc int) error {
	next, err := appendChecked(*out, Token{
		Kind: kind, Offset: at, EndOffset: at, Line: sl, Column: sc,
	})
	if err != nil {
		return err
	}
	*out = next
	return nil
}
