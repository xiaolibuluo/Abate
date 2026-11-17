/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package frontkit

import (
	"slices"
	"strings"
	"testing"

	"abate/internal/determinism"
	"abate/lang/air"
)

// kindCases enumerates the whole frozen TokenKind enum with its wire name.
var kindCases = []struct {
	kind TokenKind
	want string
}{
	{TokenIdent, "Ident"},
	{TokenKeyword, "Keyword"},
	{TokenNum, "Num"},
	{TokenStr, "Str"},
	{TokenChar, "Char"},
	{TokenOp, "Op"},
	{TokenPunct, "Punct"},
	{TokenComment, "Comment"},
	{TokenNewline, "Newline"},
	{TokenIndent, "Indent"},
	{TokenDedent, "Dedent"},
	{TokenEOF, "EOF"},
	{TokenIllegal, "Illegal"},
}

func TestTokenKindStringAllEnum(t *testing.T) {
	t.Parallel()
	if len(kindCases) != len(tokenKindNames) {
		t.Fatalf("enum count %d != name table count %d", len(kindCases), len(tokenKindNames))
	}
	for _, tc := range kindCases {
		if got := tc.kind.String(); got != tc.want {
			t.Fatalf("TokenKind(%d).String() = %q, want %q", uint16(tc.kind), got, tc.want)
		}
		if got := tokenKindNames[tc.kind]; got != tc.want {
			t.Fatalf("tokenKindNames[%d] = %q, want %q", uint16(tc.kind), got, tc.want)
		}
	}
}

func TestTokenKindStringOutOfRange(t *testing.T) {
	t.Parallel()
	cases := []struct {
		kind TokenKind
		want string
	}{
		{TokenKind(13), "TokenKind(13)"},
		{TokenKind(99), "TokenKind(99)"},
		{TokenKind(65535), "TokenKind(65535)"},
	}
	for _, tc := range cases {
		if got := tc.kind.String(); got != tc.want {
			t.Fatalf("TokenKind(%d).String() = %q, want %q",
				uint16(tc.kind), got, tc.want)
		}
		if tc.kind.Valid() {
			t.Fatalf("TokenKind(%d).Valid() = true, want false", uint16(tc.kind))
		}
	}
}

func TestTokenKindValid(t *testing.T) {
	t.Parallel()
	for _, tc := range kindCases {
		if !tc.kind.Valid() {
			t.Fatalf("TokenKind(%d).Valid() = false, want true", uint16(tc.kind))
		}
	}
	for _, bad := range []TokenKind{13, 14, 1000} {
		if bad.Valid() {
			t.Fatalf("TokenKind(%d).Valid() = true, want false", uint16(bad))
		}
	}
}

func TestTokenSpanPosition(t *testing.T) {
	t.Parallel()
	tok := Token{
		Kind:      TokenIdent,
		Text:      "abc",
		Offset:    4,
		EndOffset: 7,
		Line:      3,
		Column:    9,
	}
	start, end := tok.Span()
	wantStart := air.NewPosition(3, 9, 4, 7)
	wantEnd := air.NewPosition(3, 9, 7, 7)
	if start != wantStart || end != wantEnd {
		t.Fatalf("Span() = (%s, %s), want (%s, %s)",
			start.Canonical(), end.Canonical(),
			wantStart.Canonical(), wantEnd.Canonical())
	}
	if start.Canonical() != "3/9/4/7" || end.Canonical() != "3/9/7/7" {
		t.Fatalf("unexpected canonical positions: %s / %s",
			start.Canonical(), end.Canonical())
	}
	if !start.Valid() || !end.Valid() {
		t.Fatal("Span() produced an invalid air.Position")
	}
}

func TestTokenSpanHalfOpen(t *testing.T) {
	t.Parallel()
	tok := Token{Kind: TokenKeyword, Text: "let", Offset: 10, EndOffset: 13, Line: 1, Column: 2}
	start, end := tok.Span()
	if start.Offset != 10 || start.EndOffset != 13 {
		t.Fatalf("start interval is not [10,13): %s", start.Canonical())
	}
	if end.Offset != 13 || end.EndOffset != 13 {
		t.Fatalf("end interval is not the zero-width [13,13): %s", end.Canonical())
	}
	if !start.Contains(end) {
		t.Fatal("start span must contain the zero-width end position")
	}
	if start.Overlaps(end) {
		t.Fatal("half-open spans must not overlap at the shared boundary")
	}
	if !tok.Valid() {
		t.Fatal("a well-formed token must be valid")
	}
}

func TestTokenValidLegalAndIllegalSpans(t *testing.T) {
	t.Parallel()
	legal := Token{Kind: TokenNum, Text: "7", Offset: 0, EndOffset: 1, Line: 1, Column: 1}
	if !legal.Valid() {
		t.Fatal("legal token reported invalid")
	}
	zeroWidth := Token{Kind: TokenEOF, Text: "", Offset: 5, EndOffset: 5, Line: 2, Column: 6}
	if !zeroWidth.Valid() {
		t.Fatal("a legal zero-width EOF token reported invalid")
	}
	illegal := []struct {
		name string
		tok  Token
	}{
		{"unknown kind", Token{Kind: TokenKind(400), Offset: 0, EndOffset: 1, Line: 1, Column: 1}},
		{"negative offset", Token{Kind: TokenIdent, Offset: -1, EndOffset: 2, Line: 1, Column: 1}},
		{"reversed interval", Token{Kind: TokenIdent, Offset: 5, EndOffset: 4, Line: 1, Column: 1}},
		{"zero line", Token{Kind: TokenIdent, Offset: 0, EndOffset: 1, Line: 0, Column: 1}},
		{"zero column", Token{Kind: TokenIdent, Offset: 0, EndOffset: 1, Line: 1, Column: 0}},
		{"negative line", Token{Kind: TokenIdent, Offset: 0, EndOffset: 1, Line: -3, Column: 1}},
	}
	for _, tc := range illegal {
		if tc.tok.Valid() {
			t.Fatalf("%s: Valid() = true, want false", tc.name)
		}
	}
}

// tokenOrderingSnapshot folds a deterministically sorted token slice into one
// canonical byte string, covering Span, String and stable ordering together.
func tokenOrderingSnapshot() []byte {
	tokens := []Token{
		{Kind: TokenKeyword, Text: "let", Offset: 10, EndOffset: 13, Line: 1, Column: 11},
		{Kind: TokenIdent, Text: "x", Offset: 0, EndOffset: 1, Line: 1, Column: 1},
		{Kind: TokenOp, Text: "==", Offset: 4, EndOffset: 6, Line: 1, Column: 5},
		{Kind: TokenNum, Text: "7", Offset: 7, EndOffset: 8, Line: 1, Column: 8},
		{Kind: TokenEOF, Text: "", Offset: 13, EndOffset: 13, Line: 1, Column: 14},
	}
	slices.SortStableFunc(tokens, func(a, b Token) int {
		if a.Offset != b.Offset {
			if a.Offset < b.Offset {
				return -1
			}
			return 1
		}
		if a.Kind != b.Kind {
			if a.Kind < b.Kind {
				return -1
			}
			return 1
		}
		return strings.Compare(a.Text, b.Text)
	})
	parts := make([]string, 0, len(tokens))
	for _, tok := range tokens {
		start, end := tok.Span()
		parts = append(parts, tok.Kind.String()+":"+tok.Text+":"+
			start.Canonical()+":"+end.Canonical())
	}
	return []byte(strings.Join(parts, "|"))
}

func TestTokenDeterministicOrdering100(t *testing.T) {
	t.Parallel()
	determinism.MustMatch100(t, 42, tokenOrderingSnapshot)
}

// TestBindOperatorPutsTokenTextIntoNodeID proves the shared binding entry enters the NodeID payload and SemOperator annotation.
func TestBindOperatorPutsTokenTextIntoNodeID(t *testing.T) {
	t.Parallel()
	start := air.NewPosition(1, 1, 0, 2)
	end := air.NewPosition(1, 1, 2, 2)
	equal := BindOperator(air.KindBinary, start, end, Token{Kind: TokenOp, Text: "=="})
	notEqual := BindOperator(air.KindBinary, start, end, Token{Kind: TokenOp, Text: "!="})
	if air.AttributeOf(equal) != "==" {
		t.Fatalf("attribute = %q, want ==", air.AttributeOf(equal))
	}
	if air.IDOf(equal) == air.IDOf(notEqual) {
		t.Fatal("operator spelling did not enter the AIR NodeID")
	}
	value, ok := equal.Sem(air.SemOperator).(string)
	if !ok || value != "==" {
		t.Fatalf("SemOperator = %q, want ==", value)
	}
}
