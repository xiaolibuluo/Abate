/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package frontkit

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"abate/lang/air"
)

// ErrTokenFidelity reports a disagreement between a source token stream and
// the AIR lexical projection of a unit.
var ErrTokenFidelity = errors.New("frontkit: token fidelity check failed")

// VerifyTokenFidelity proves a bidirectional match between a source token
// stream and the AIR lexical projection of a unit. Every projected AIR token
// must exist in the real stream at the same byte interval with the same text,
// and every identifier, literal or comment in the real stream must be projected
// into AIR. Keywords, operators, punctuation, indentation and inserted markers
// are outside this equivalence subset; byte-level source regeneration is not
// promised. The comparison key is kind-independent because AIR node kinds and
// frontkit token kinds are different frozen enumerations.
func VerifyTokenFidelity(tokens []Token, unit *air.Unit) error {
	if len(tokens) > maxFileTokens {
		return fmt.Errorf("%w: %d tokens exceed the per-file ceiling",
			ErrLexFailed, len(tokens))
	}
	all := make(map[string]struct{}, len(tokens))
	lexical := make([]Token, 0, len(tokens))
	for i, tok := range tokens {
		if !tok.Valid() {
			return fmt.Errorf("%w: invalid token at %d", ErrUnknownToken, i)
		}
		all[tokenKey(tok)] = struct{}{}
		if isLexicalToken(tok.Kind) && tok.Text != "" {
			lexical = append(lexical, tok)
		}
	}
	seen := make(map[string]struct{}, len(lexical))
	for _, node := range air.Collect(unit) {
		if !isLexicalNode(node.Kind()) {
			continue
		}
		text := air.AttributeOf(node)
		if text == "" {
			continue
		}
		pos := node.Pos()
		key := projectionKey(pos.Offset, pos.EndOffset, text)
		if _, ok := all[key]; !ok {
			return fmt.Errorf("%w: AIR invented %q at [%d,%d)",
				ErrTokenFidelity, text, pos.Offset, pos.EndOffset)
		}
		seen[key] = struct{}{}
	}
	for _, tok := range lexical {
		if _, ok := seen[tokenKey(tok)]; !ok {
			return fmt.Errorf("%w: source %s %q at [%d,%d) is absent from AIR",
				ErrTokenFidelity, tok.Kind.String(), tok.Text,
				tok.Offset, tok.EndOffset)
		}
	}
	return nil
}

// VerifyNodeFidelity applies the same bidirectional check to one root node by
// wrapping it in an otherwise empty unit.
func VerifyNodeFidelity(tokens []Token, root air.Node) error {
	unit := air.NewUnit(air.NewPackageRef("fidelity", nil))
	if root != nil {
		if err := unit.AddFile(air.NewFileRef("fidelity.src", nil), root); err != nil {
			return fmt.Errorf("%w: %w", ErrTokenFidelity, err)
		}
	}
	return VerifyTokenFidelity(tokens, unit)
}

// tokenKey renders one source token as its byte-exact comparison key.
func tokenKey(tok Token) string {
	return projectionKey(tok.Offset, tok.EndOffset, tok.Text)
}

// projectionKey renders one AIR lexical projection entry as the same key.
func projectionKey(offset, endOffset int, text string) string {
	return strconv.Itoa(offset) + "\x00" + strconv.Itoa(endOffset) + "\x00" + text
}

// isLexicalToken reports whether a token belongs to the preserved subset.
func isLexicalToken(kind TokenKind) bool {
	switch kind {
	case TokenIdent, TokenNum, TokenStr, TokenChar, TokenComment:
		return true
	default:
		return false
	}
}

// isLexicalNode reports whether an AIR kind may project a preserved token.
func isLexicalNode(kind air.NodeKind) bool {
	switch kind {
	case air.KindIdent, air.KindNumLit, air.KindStrLit, air.KindCharLit,
		air.KindBoolLit, air.KindNilLit, air.KindComment, air.KindDoc,
		air.KindDirective, air.KindAttribute:
		return true
	default:
		return false
	}
}

// BindTokenSpellings makes the AIR root carry the exact spelling of every
// source token that no vertex already projects at its own byte interval.
//
// The language reducers model identifiers, literals and comments as lexical
// leaves and fold operators into structural nodes, but keywords, punctuation
// and tokens swallowed by error recovery can still be dropped. Those dropped
// tokens are real mutation targets: rewriting "in" to "!=" or "2" to "1024"
// must move the AIR identifier. Writing the real token text through
// air.NewNodeAt puts it in the ComputeNodeID payload, exactly as
// BindOperator does for binary operators; no metric, threshold or hash rule
// is changed. Trivia markers are excluded because they carry no source token
// spelling of their own.
func BindTokenSpellings(tokens []Token, root air.Node) air.Node {
	if root == nil || len(tokens) == 0 {
		return root
	}
	projected := make(map[string]struct{}, len(tokens))
	air.Walk(root, func(current air.Node) bool {
		text := air.AttributeOf(current)
		if text != "" {
			pos := current.Pos()
			projected[projectionKey(pos.Offset, pos.EndOffset, text)] = struct{}{}
		}
		return true
	})
	var spelling strings.Builder
	for _, tok := range tokens {
		if !isBindableToken(tok) {
			continue
		}
		if _, found := projected[tokenKey(tok)]; found {
			continue
		}
		spelling.WriteString(tok.Text)
		spelling.WriteByte(0)
	}
	if spelling.Len() == 0 {
		return root
	}
	rebuilt := air.NewNodeAt(root.Kind(), root.Pos(), root.End(),
		spelling.String(), root.Children()...)
	semantics := air.SemanticsOf(root)
	for _, key := range semantics.Keys() {
		value, _ := semantics.Get(key)
		_ = air.Annotate(rebuilt, key, value)
	}
	return rebuilt
}

// isBindableToken reports whether a token carries a real source spelling.
func isBindableToken(tok Token) bool {
	if tok.Text == "" {
		return false
	}
	switch tok.Kind {
	case TokenNewline, TokenIndent, TokenDedent, TokenEOF:
		return false
	default:
		return true
	}
}
