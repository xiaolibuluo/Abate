/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package frontkit

import (
	"errors"
	"strings"
	"testing"

	"abate/internal/determinism"
	"abate/lang/air"
)

func tokenList(kinds []TokenKind, texts []string) []Token {
	out := make([]Token, 0, len(texts)+1)
	offset := 0
	for i, text := range texts {
		out = append(out, Token{
			Kind: kinds[i], Text: text, Offset: offset,
			EndOffset: offset + len(text), Line: 1, Column: offset + 1,
		})
		offset += len(text) + 1
	}
	out = append(out, Token{
		Kind: TokenEOF, Offset: offset, EndOffset: offset,
		Line: 1, Column: offset + 1,
	})
	return out
}
func newExpressionParser(t *testing.T, tokens []Token) *PrattParser {
	t.Helper()
	stream, err := NewTokenStream(tokens)
	if err != nil {
		t.Fatal(err)
	}
	parser, err := NewPrattParser(stream, NewView(SourceFromTokens(tokens), "expr.fk"))
	if err != nil {
		t.Fatal(err)
	}
	literal := func(_ *PrattParser, tok Token) (air.Node, error) {
		start, end := tok.Span()
		return air.NewNodeAt(fakeNodeKind(tok.Kind), start, end, tok.Text), nil
	}
	binary := func(_ *PrattParser, left air.Node, op Token, right air.Node) (air.Node, error) {
		leftStart, _ := air.Span(left)
		_, rightEnd := air.Span(right)
		start := air.NewPosition(op.Line, op.Column, leftStart, rightEnd)
		end := air.NewPosition(op.Line, op.Column, rightEnd, rightEnd)
		return air.NewNodeAt(air.KindBinary, start, end, op.Text, left, right), nil
	}
	for _, kind := range []TokenKind{TokenIdent, TokenNum, TokenStr} {
		if err := parser.Register(kind, ParseRule{Prefix: literal}); err != nil {
			t.Fatal(err)
		}
	}
	for _, entry := range []struct {
		text  string
		power int
		right bool
	}{
		{"=", 5, true}, {"+", 10, false}, {"-", 10, false},
		{"*", 20, false}, {"/", 20, false}, {"^", 30, true},
	} {
		rule := ParseRule{InfixPower: entry.power, RightAssoc: entry.right, Infix: binary}
		if err := parser.RegisterText(TokenOp, entry.text, rule); err != nil {
			t.Fatal(err)
		}
	}
	unary := func(p *PrattParser, tok Token) (air.Node, error) {
		operand, err := p.ParsePrefixOperand(tok)
		if err != nil {
			return nil, err
		}
		start, _ := tok.Span()
		_, rightEnd := air.Span(operand)
		pos := air.NewPosition(tok.Line, tok.Column, start.Offset, rightEnd)
		end := air.NewPosition(tok.Line, tok.Column, rightEnd, rightEnd)
		return air.NewNodeAt(air.KindUnary, pos, end, tok.Text, operand), nil
	}
	if err := parser.RegisterText(TokenOp, "!",
		ParseRule{PrefixPower: 40, Prefix: unary}); err != nil {
		t.Fatal(err)
	}
	return parser
}
func expressionTokens(texts ...string) []Token {
	kinds := make([]TokenKind, 0, len(texts))
	for _, text := range texts {
		if strings.ContainsAny(text, "=+-*/^!") {
			kinds = append(kinds, TokenOp)
			continue
		}
		kinds = append(kinds, TokenIdent)
	}
	return tokenList(kinds, texts)
}
func TestPrattPrecedenceAndAssociativity(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name       string
		texts      []string
		root       string
		childIndex int
		child      string
	}{
		{"multiplication binds tighter", []string{"a", "+", "b", "*", "c"}, "+", 1, "*"},
		{"subtraction is left associative", []string{"a", "-", "b", "-", "c"}, "-", 0, "-"},
		{"assignment is right associative", []string{"a", "=", "b", "=", "c"}, "=", 1, "="},
		{"exponent binds tighter than plus", []string{"x", "^", "y", "+", "z"}, "+", 0, "^"},
	}
	for _, tc := range cases {
		parser := newExpressionParser(t, expressionTokens(tc.texts...))
		root, err := parser.Parse()
		if err != nil {
			t.Fatalf("%s: %v", tc.name, err)
		}
		if air.AttributeOf(root) != tc.root {
			t.Fatalf("%s: root = %q, want %q", tc.name, air.AttributeOf(root), tc.root)
		}
		children := root.Children()
		if len(children) != 2 || air.AttributeOf(children[tc.childIndex]) != tc.child {
			t.Fatalf("%s: unexpected children %v", tc.name, children)
		}
		if parser.Depth() != 0 {
			t.Fatalf("%s: depth leaked to %d", tc.name, parser.Depth())
		}
	}
}
func TestPrattDepthLimitAndRecovery(t *testing.T) {
	t.Parallel()
	count := maxParseDepth + 8
	kinds := make([]TokenKind, 0, count+1)
	texts := make([]string, 0, count+1)
	for i := 0; i < count; i++ {
		kinds = append(kinds, TokenOp)
		texts = append(texts, "!")
	}
	kinds = append(kinds, TokenIdent)
	texts = append(texts, "x")
	deep := newExpressionParser(t, tokenList(kinds, texts))
	if _, err := deep.Parse(); !errors.Is(err, ErrParseDepth) {
		t.Fatalf("deep parse error = %v, want ErrParseDepth", err)
	}
	broken := newExpressionParser(t, tokenList(
		[]TokenKind{TokenIdent, TokenKeyword, TokenIdent, TokenNewline, TokenIdent},
		[]string{"a", "let", "b", "\n", "c"}))
	if _, err := broken.Parse(); !errors.Is(err, ErrUnknownToken) {
		t.Fatalf("trailing token error = %v, want ErrUnknownToken", err)
	}
	if skipped := broken.Recover(); skipped != 2 {
		t.Fatalf("Recover skipped %d tokens, want 2", skipped)
	}
	if next, ok := broken.Stream().Peek(); !ok || next.Kind != TokenNewline {
		t.Fatalf("recovery did not stop at the newline: %+v %v", next, ok)
	}
	missing := newExpressionParser(t, tokenList([]TokenKind{TokenKeyword}, []string{"let"}))
	if _, err := missing.Parse(); !errors.Is(err, ErrUnknownToken) {
		t.Fatalf("missing prefix error = %v, want ErrUnknownToken", err)
	}
}
func TestStreamSaveRestoreExpectAndLimits(t *testing.T) {
	t.Parallel()
	tokens := tokenList(
		[]TokenKind{TokenIdent, TokenOp, TokenIdent}, []string{"a", "+", "b"})
	stream, err := NewTokenStream(tokens)
	if err != nil {
		t.Fatal(err)
	}
	first, ok := stream.Peek()
	if !ok || first.Text != "a" {
		t.Fatalf("Peek = %+v %v", first, ok)
	}
	mark, err := stream.Save()
	if err != nil {
		t.Fatal(err)
	}
	if _, advanced := stream.Advance(); !advanced {
		t.Fatal("Advance failed on a live stream")
	}
	if _, err := stream.Expect(TokenOp); err != nil {
		t.Fatal(err)
	}
	if err := stream.Restore(mark); err != nil {
		t.Fatal(err)
	}
	if stream.Depth() != 0 || stream.Remaining() != 4 || stream.Done() {
		t.Fatalf("stream state after Restore = depth %d remaining %d done %v",
			stream.Depth(), stream.Remaining(), stream.Done())
	}
	if err := stream.Restore(-1); !errors.Is(err, ErrUnknownToken) {
		t.Fatalf("bad mark error = %v", err)
	}
	if err := stream.Restore(0); !errors.Is(err, ErrUnknownToken) {
		t.Fatalf("restore without save error = %v", err)
	}
	if _, err := stream.Expect(TokenEOF); !errors.Is(err, ErrUnknownToken) {
		t.Fatalf("mismatched Expect error = %v", err)
	}
	tokens[0].Text = "mutated"
	if stream.Tokens()[0].Text != "a" {
		t.Fatal("NewTokenStream did not take a defensive copy")
	}
	deep, err := NewTokenStream(tokens)
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < maxStreamDepth; i++ {
		if _, err := deep.Save(); err != nil {
			t.Fatalf("save %d: %v", i, err)
		}
	}
	if _, err := deep.Save(); !errors.Is(err, ErrParseDepth) {
		t.Fatalf("save depth error = %v, want ErrParseDepth", err)
	}
	if _, err := NewTokenStream([]Token{{Kind: TokenIdent}}); !errors.Is(err, ErrUnknownToken) {
		t.Fatalf("invalid token error = %v", err)
	}
	reversed := tokenList([]TokenKind{TokenIdent, TokenIdent}, []string{"a", "b"})
	reversed[0].Offset, reversed[0].EndOffset = 10, 11
	if _, err := NewTokenStream(reversed); !errors.Is(err, ErrUnknownToken) {
		t.Fatalf("reversed token error = %v", err)
	}
	oversized := make([]Token, maxFileTokens+1)
	for i := range oversized {
		oversized[i] = Token{Kind: TokenIdent, Text: "x", Offset: i,
			EndOffset: i + 1, Line: 1, Column: i + 1}
	}
	if _, err := NewTokenStream(oversized); !errors.Is(err, ErrLexFailed) {
		t.Fatalf("oversized stream error = %v, want ErrLexFailed", err)
	}
}

func TestParserDeterminism100(t *testing.T) {
	t.Parallel()
	determinism.MustMatch100(t, 42, func() []byte {
		parser := newExpressionParser(t,
			expressionTokens("a", "+", "b", "*", "c", "=", "d"))
		root, err := parser.Parse()
		if err != nil {
			t.Fatal(err)
		}
		digest := determinism.HashState(root)
		return digest[:]
	})
}
