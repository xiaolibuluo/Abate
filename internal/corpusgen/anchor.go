/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package corpusgen

import (
	"bytes"
	"errors"
	"fmt"
)

// errAnchor reports a generated file whose first mutation token escaped the
// capture anchor. Such a file is a generator failure, never a threshold tweak.
var errAnchor = errors.New("corpusgen: capture anchor preflight failed")

// Token tables mirror internal/e2e/mutate/ops_*.go so preflight uses the same
// byte-matching semantics as the real scanner (word boundaries included).
var (
	anchorDigitTokens = []string{"0", "1", "2", "3", "10", "16", "32", "64"}
	anchorOpTokens    = []string{"<<", ">>", "&&", "||", "+", "-", "*", "/", "%", "&", "|", "^"}
	anchorCmpTokens   = []string{"===", "!==", "==", "!=", "<=", ">=", "<", ">", "is", "is not", "in", "not in"}
	anchorCallTokens  = []string{"f()", "g()", "foo()", "bar()", "call()", "print()", "baz()"}
)

// CaptureAnchor is one language-specific anchor and its expected token class.
type CaptureAnchor struct {
	Language string
	Clean    bool
	Bounds   AnchorBounds
}

// isWordByte matches the scanner's identifier-extension predicate.
func isWordByte(b byte) bool {
	return b == '_' || (b >= 'a' && b <= 'z') ||
		(b >= 'A' && b <= 'Z') || (b >= '0' && b <= '9')
}

// firstTokenOffset returns the earliest byte offset of any token in the list,
// or -1 when no token occurs. Multi-byte tokens are preferred at equal offsets.
func firstTokenOffset(content []byte, tokens []string) int {
	first := -1
	for _, token := range tokens {
		needle := []byte(token)
		if len(needle) == 0 {
			continue
		}
		lead, tail := needle[0], needle[len(needle)-1]
		for at := 0; at <= len(content); {
			index := bytes.Index(content[at:], needle)
			if index < 0 {
				break
			}
			start := at + index
			end := start + len(needle)
			leftOK := !isWordByte(lead) || start == 0 || !isWordByte(content[start-1])
			rightOK := !isWordByte(tail) || end == len(content) || !isWordByte(content[end])
			if leftOK && rightOK && (first < 0 || start < first) {
				first = start
			}
			at = end
		}
	}
	return first
}

// Preflight asserts that the first digit, arithmetic and comparison tokens of
// the whole file lie inside the anchor, and that non-clean anchors carry the
// first call token as well. A clean first-unit anchor carries no call token.
func (a CaptureAnchor) Preflight(content []byte) error {
	classes := []struct {
		name   string
		tokens []string
	}{
		{"digit", anchorDigitTokens},
		{"op", anchorOpTokens},
		{"cmp", anchorCmpTokens},
	}
	if !a.Clean {
		classes = append(classes, struct {
			name   string
			tokens []string
		}{"call", anchorCallTokens})
	}
	for _, class := range classes {
		at := firstTokenOffset(content, class.tokens)
		if at < 0 {
			return fmt.Errorf("%w: %s/%s: token absent", errAnchor, a.Language, class.name)
		}
		if at < a.Bounds.Start || at >= a.Bounds.End {
			return fmt.Errorf("%w: %s/%s: first offset %d outside [%d,%d)",
				errAnchor, a.Language, class.name, at, a.Bounds.Start, a.Bounds.End)
		}
	}
	return nil
}

// HeaderHasToken reports whether the pre-anchor header already contains a
// mutation token, which would invalidate every generated file of that recipe.
func HeaderHasToken(header string) bool {
	data := []byte(header)
	for _, tokens := range [][]string{
		anchorDigitTokens, anchorOpTokens, anchorCmpTokens, anchorCallTokens,
	} {
		if firstTokenOffset(data, tokens) >= 0 {
			return true
		}
	}
	return false
}
