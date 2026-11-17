/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package frontkit

import "strconv"

// TokenKind classifies the lexical category of a token. The wire names in
// tokenKindNames are frozen: an existing value is never reused or renumbered.
type TokenKind uint16

const (
	// TokenIdent marks an identifier.
	TokenIdent TokenKind = iota
	// TokenKeyword marks a reserved word.
	TokenKeyword
	// TokenNum marks a numeric literal.
	TokenNum
	// TokenStr marks a string literal.
	TokenStr
	// TokenChar marks a character or rune literal.
	TokenChar
	// TokenOp marks an operator.
	TokenOp
	// TokenPunct marks punctuation such as braces, commas and semicolons.
	TokenPunct
	// TokenComment marks a comment trivia token.
	TokenComment
	// TokenNewline marks a logical line break.
	TokenNewline
	// TokenIndent marks a synthesized indentation increase.
	TokenIndent
	// TokenDedent marks a synthesized indentation decrease.
	TokenDedent
	// TokenEOF marks the end of the token stream.
	TokenEOF
	// TokenIllegal marks a byte sequence that could not be classified.
	TokenIllegal
	// The tokenKindCount value is the exclusive upper bound of the enum range.
	tokenKindCount
)

// tokenKindNames maps each frozen TokenKind to its stable wire name.
var tokenKindNames = [...]string{
	"Ident",
	"Keyword",
	"Num",
	"Str",
	"Char",
	"Op",
	"Punct",
	"Comment",
	"Newline",
	"Indent",
	"Dedent",
	"EOF",
	"Illegal",
}

// String returns the frozen wire name of the token kind. Out-of-range values
// render as TokenKind(n) instead of panicking, so damaged input stays
// deterministic and inspectable.
func (k TokenKind) String() string {
	if int(k) < len(tokenKindNames) {
		return tokenKindNames[k]
	}
	return "TokenKind(" + strconv.FormatUint(uint64(k), 10) + ")"
}

// Valid reports whether the token kind is one of the frozen enum values.
func (k TokenKind) Valid() bool {
	return k < tokenKindCount
}
