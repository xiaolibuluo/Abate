/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package gofront

import (
	"fmt"
	"go/scanner"
	"go/token"
	"slices"
	"strconv"
	"strings"

	"abate/lang/air"
)

// Token is one lexical token observed in a Go source buffer.
type Token struct {
	// Kind is the go/token name of the token.
	Kind string
	// Text is the exact source text of the token.
	Text string
	// Offset is the zero-based byte offset of the first byte.
	Offset int
	// EndOffset is the exclusive byte offset just past the token.
	EndOffset int
}

// ScanTokens scans a source buffer using only the standard go/scanner.
//
// An error callback is not installed: illegal characters are reported as
// ILLEGAL tokens rather than panicking or aborting the scan.
func ScanTokens(src []byte) ([]Token, error) {
	fset := token.NewFileSet()
	file := fset.AddFile("source.go", fset.Base(), len(src))
	var scannerState scanner.Scanner
	scannerState.Init(file, src, nil, scanner.ScanComments)
	out := make([]Token, 0, 32)
	for {
		pos, tok, literal := scannerState.Scan()
		if tok == token.EOF {
			break
		}
		text := literal
		if text == "" {
			text = tok.String()
		}
		offset := file.Offset(pos)
		out = append(out, Token{
			Kind: tok.String(), Text: text,
			Offset: offset, EndOffset: offset + len(text),
		})
	}
	return out, nil
}

// Regenerate projects AIR back into its lexical token view.
//
// Only nodes that stand for a real source token participate: identifiers,
// literals, comments, directives, and struct tags. Structural nodes such as
// calls and blocks are deliberately excluded because a neutral IR does not
// model separators or keywords as vertices.
func Regenerate(unit *air.Unit) []Token {
	out := make([]Token, 0, 32)
	for _, current := range air.Collect(unit) {
		if !isLexicalNode(current.Kind()) {
			continue
		}
		text := air.AttributeOf(current)
		if text == "" {
			continue
		}
		pos := current.Pos()
		out = append(out, Token{
			Kind: current.Kind().String(), Text: text,
			Offset: pos.Offset, EndOffset: pos.EndOffset,
		})
	}
	slices.SortStableFunc(out, func(a, b Token) int {
		switch {
		case a.Offset < b.Offset:
			return -1
		case a.Offset > b.Offset:
			return 1
		default:
			return strings.Compare(a.Text, b.Text)
		}
	})
	return out
}

// VerifyTokenFidelity proves a bidirectional match between the source token
// stream and the AIR lexical projection.
//
// Every AIR token must exist in the real stream, and every identifier,
// literal, or comment in the real stream must be represented in AIR. Keywords,
// separators, and inserted semicolons are outside this equivalence subset;
// byte-level source regeneration is explicitly not promised by version 1.
func VerifyTokenFidelity(src []byte, unit *air.Unit) error {
	scanned, err := ScanTokens(src)
	if err != nil {
		return err
	}
	all := make(map[string]struct{}, len(scanned))
	lexical := make([]Token, 0, len(scanned))
	for _, item := range scanned {
		all[tokenKey(item)] = struct{}{}
		if isLexicalToken(item.Kind) {
			lexical = append(lexical, item)
		}
	}
	regenerated := Regenerate(unit)
	seen := make(map[string]struct{}, len(regenerated))
	for _, item := range regenerated {
		if _, ok := all[tokenKey(item)]; !ok {
			return fmt.Errorf("%w: AIR invented %q at offset %d",
				ErrTokenFidelity, item.Text, item.Offset)
		}
		seen[tokenKey(item)] = struct{}{}
	}
	for _, item := range lexical {
		if _, ok := seen[tokenKey(item)]; !ok {
			return fmt.Errorf("%w: source %s %q at offset %d is absent from AIR",
				ErrTokenFidelity, item.Kind, item.Text, item.Offset)
		}
	}
	return nil
}

// isLexicalNode reports whether an AIR kind stands for a real source token.
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

// isLexicalToken reports whether a go/token name belongs to the subset that
// AIR promises to preserve verbatim.
func isLexicalToken(kind string) bool {
	switch kind {
	case "IDENT", "INT", "FLOAT", "IMAG", "CHAR", "STRING", "COMMENT":
		return true
	default:
		return false
	}
}

// tokenKey renders a token as a stable, order-independent comparison key.
func tokenKey(item Token) string {
	return strconv.Itoa(item.Offset) + "\x00" + item.Text
}
