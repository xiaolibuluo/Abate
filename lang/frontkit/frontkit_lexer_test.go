/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package frontkit

import (
	"errors"
	"strconv"
	"strings"
	"testing"

	"abate/internal/determinism"
	"abate/lang/air"
)

// fakeSpec is the single minimal LanguageSpec and LexSpec used by every B02
// test. It is deliberately not a real language frontend.
type fakeSpec struct {
	tag             string
	keywords        []string
	operators       []string
	indentSensitive bool
	failParse       bool
}

func newFakeSpec() *fakeSpec {
	return &fakeSpec{
		tag:       "go",
		keywords:  []string{"let", "fn", "return"},
		operators: []string{"===", "==", "=", "+", "-", "*", "/", "!", "->"},
	}
}

func (s *fakeSpec) Keywords() []string                    { return s.keywords }
func (s *fakeSpec) Operators() []string                   { return s.operators }
func (*fakeSpec) LineComment() string                     { return "//" }
func (*fakeSpec) BlockComment() (open, closeDelim string) { return "/*", "*/" }
func (*fakeSpec) StringQuotes() []string                  { return []string{`"""`, `"`, "'"} }
func (*fakeSpec) RawStringPrefixes() []string             { return []string{"r", "f"} }
func (*fakeSpec) NumberSeparators() []string              { return []string{"_"} }
func (s *fakeSpec) IndentSensitive() bool                 { return s.indentSensitive }
func (s *fakeSpec) Tag() string                           { return s.tag }
func (s *fakeSpec) LexSpec() LexSpec                      { return s }
func (*fakeSpec) NormalizeDialect(d string) (string, error) {
	if strings.Contains(d, "!") {
		return "", ErrBadDialect
	}
	if d == "" {
		return "default", nil
	}
	return strings.ToLower(strings.TrimSpace(d)), nil
}
func (s *fakeSpec) ParseTokens(tokens []Token, path string) (air.Node, []air.Diagnostic, error) {
	if s.failParse {
		return nil, nil, ErrLexFailed
	}
	return fakeReduce(tokens, path), nil, nil
}

// fakeReduce projects every lexical token straight into an AIR node.
func fakeReduce(tokens []Token, path string) air.Node {
	view := NewView(SourceFromTokens(tokens), path)
	view.SetPackageName("fake")
	children := make([]air.Node, 0, len(tokens))
	for i, tok := range tokens {
		if !isLexicalToken(tok.Kind) {
			continue
		}
		node := view.NewNodeAt(fakeNodeKind(tok.Kind), tok, tok, tok.Text)
		view.Bind(i, node)
		children = append(children, node)
	}
	if len(tokens) == 0 {
		return view.Synthetic(air.KindPackage, "")
	}
	return view.NewNodeAt(air.KindPackage,
		tokens[0], tokens[len(tokens)-1], "", children...)
}

func fakeNodeKind(kind TokenKind) air.NodeKind {
	switch kind {
	case TokenIdent, TokenKeyword:
		return air.KindIdent
	case TokenNum:
		return air.KindNumLit
	case TokenStr:
		return air.KindStrLit
	case TokenChar:
		return air.KindCharLit
	case TokenComment:
		return air.KindComment
	default:
		return air.KindRegion
	}
}

func mustLex(t *testing.T, spec LexSpec, src string) []Token {
	t.Helper()
	lexer, err := NewLexer(spec, []byte(src))
	if err != nil {
		t.Fatalf("NewLexer: %v", err)
	}
	tokens, err := lexer.Tokenize()
	if err != nil {
		t.Fatalf("Tokenize: %v", err)
	}
	return tokens
}

func hasKind(tokens []Token, kind TokenKind) bool {
	for _, tok := range tokens {
		if tok.Kind == kind {
			return true
		}
	}
	return false
}

func tokenSnapshot(tokens []Token) []byte {
	parts := make([]string, 0, len(tokens))
	for _, tok := range tokens {
		parts = append(parts, strings.Join([]string{
			tok.Kind.String(), tok.Text,
			strconv.Itoa(tok.Offset), strconv.Itoa(tok.EndOffset),
			strconv.Itoa(tok.Line), strconv.Itoa(tok.Column),
		}, ":"))
	}
	return []byte(strings.Join(parts, "|"))
}

func assertLexicalInvariants(t *testing.T, src string, tokens []Token) {
	t.Helper()
	if len(tokens) == 0 || tokens[len(tokens)-1].Kind != TokenEOF {
		t.Fatal("token stream must end with exactly one EOF token")
	}
	for i, tok := range tokens {
		if !tok.Valid() {
			t.Fatalf("token %d is invalid: %+v", i, tok)
		}
		if tok.Text != src[tok.Offset:tok.EndOffset] {
			t.Fatalf("token %d text %q differs from its source interval", i, tok.Text)
		}
		if i > 0 && tok.Offset < tokens[i-1].Offset {
			t.Fatalf("token %d is out of source order", i)
		}
	}
}

const lexerFixture = "let total = 12_5 == \"text\" 'c' // note\n/* block */\nx(1);\n"

func TestLexerSixTokenClasses(t *testing.T) {
	t.Parallel()
	tokens := mustLex(t, newFakeSpec(), lexerFixture)
	assertLexicalInvariants(t, lexerFixture, tokens)
	for _, kind := range []TokenKind{
		TokenKeyword, TokenIdent, TokenNum, TokenStr, TokenChar,
		TokenOp, TokenPunct, TokenComment, TokenNewline, TokenEOF,
	} {
		if !hasKind(tokens, kind) {
			t.Fatalf("missing token kind %s in %v", kind.String(), tokenSnapshot(tokens))
		}
	}
}

func TestLexerLongestMatchAndKeywordBoundary(t *testing.T) {
	t.Parallel()
	const src = "a === b == c = d letter letting -> e"
	tokens := mustLex(t, newFakeSpec(), src)
	assertLexicalInvariants(t, src, tokens)
	wantOperators := []string{"===", "==", "=", "->"}
	seen := make([]string, 0, len(wantOperators))
	for _, tok := range tokens {
		if tok.Kind == TokenOp {
			seen = append(seen, tok.Text)
		}
	}
	if strings.Join(seen, ",") != strings.Join(wantOperators, ",") {
		t.Fatalf("operators = %v, want %v", seen, wantOperators)
	}
	ident, keyword := 0, 0
	for _, tok := range tokens {
		if tok.Kind == TokenIdent && (tok.Text == "letter" || tok.Text == "letting") {
			ident++
		}
		if tok.Kind == TokenKeyword && tok.Text == "let" {
			keyword++
		}
	}
	if ident != 2 || keyword != 0 {
		t.Fatalf("identifier boundary failed: idents=%d keywords=%d", ident, keyword)
	}
}

func TestLexerIndentDedentSynthesis(t *testing.T) {
	t.Parallel()
	spec := newFakeSpec()
	spec.indentSensitive = true
	const src = "fn f\n  x\n  y\n// ignored\nz\n"
	tokens := mustLex(t, spec, src)
	assertLexicalInvariants(t, src, tokens)
	indents, dedents := 0, 0
	for _, tok := range tokens {
		if tok.Kind == TokenIndent {
			indents++
		}
		if tok.Kind == TokenDedent {
			dedents++
		}
	}
	if indents != 1 || dedents != 1 {
		t.Fatalf("INDENT=%d DEDENT=%d, want 1 and 1: %s",
			indents, dedents, tokenSnapshot(tokens))
	}
}

func TestLexerBoundariesBOMCRLFIllegalAndEOF(t *testing.T) {
	t.Parallel()
	src := "\xEF\xBB\xBF" + "let a\r\nlet b\r\n" + string([]byte{0xff, 0x00})
	tokens := mustLex(t, newFakeSpec(), src)
	assertLexicalInvariants(t, src, tokens)
	if !hasKind(tokens, TokenIllegal) {
		t.Fatal("invalid UTF-8 bytes must become Illegal tokens")
	}
	if tokens[0].Offset != 3 {
		t.Fatalf("BOM was not skipped: first offset %d", tokens[0].Offset)
	}
	crlf := 0
	for _, tok := range tokens {
		if tok.Kind == TokenNewline && tok.Text == "\r\n" {
			crlf++
		}
	}
	if crlf != 2 {
		t.Fatalf("CRLF newline count = %d, want 2", crlf)
	}
	eof := tokens[len(tokens)-1]
	if eof.Offset != len(src) || eof.EndOffset != len(src) || eof.Line != 3 {
		t.Fatalf("EOF position = %+v, want offset %d on line 3", eof, len(src))
	}
}

func TestLexerEmptyAndUnterminatedInput(t *testing.T) {
	t.Parallel()
	for _, src := range []string{"", "\xEF\xBB\xBF", `"unterminated`, "/* unterminated", "\\"} {
		tokens := mustLex(t, newFakeSpec(), src)
		assertLexicalInvariants(t, src, tokens)
	}
	empty := mustLex(t, newFakeSpec(), "")
	if len(empty) != 1 || empty[0].Kind != TokenEOF {
		t.Fatalf("empty source produced %v", tokenSnapshot(empty))
	}
}

func TestLexerNilSpecResetAndCeiling(t *testing.T) {
	t.Parallel()
	if _, err := NewLexer(nil, nil); !errors.Is(err, ErrNilSpec) {
		t.Fatalf("nil spec error = %v, want ErrNilSpec", err)
	}
	lexer, err := NewLexer(newFakeSpec(), []byte(lexerFixture))
	if err != nil {
		t.Fatal(err)
	}
	first, err := lexer.Tokenize()
	if err != nil {
		t.Fatal(err)
	}
	second, err := lexer.Tokenize()
	if err != nil {
		t.Fatal(err)
	}
	if string(tokenSnapshot(first)) != string(tokenSnapshot(second)) {
		t.Fatal("Tokenize is not idempotent on one lexer")
	}
	huge, err := NewLexer(newFakeSpec(), []byte(strings.Repeat(";", maxFileTokens+1)))
	if err != nil {
		t.Fatal(err)
	}
	if _, err := huge.Tokenize(); !errors.Is(err, ErrLexFailed) {
		t.Fatalf("token ceiling error = %v, want ErrLexFailed", err)
	}
}

func TestLexerDeterminism100(t *testing.T) {
	t.Parallel()
	determinism.MustMatch100(t, 42, func() []byte {
		return tokenSnapshot(mustLex(t, newFakeSpec(), lexerFixture))
	})
}
