/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/
package mutcore

import (
	"fmt"
	"strconv"
	"strings"
)

// TokenClass enumerates the mutable lexical categories the operator
// registry rewrites. Classes are intentionally lexical, not syntactic, so
// every language backend can share the same classification rules.
type TokenClass int

const (
	// ClassUnknown marks tokens that never entered the mutable set.
	ClassUnknown TokenClass = iota
	// ClassOperator covers the binary comparison and arithmetic operators.
	ClassOperator
	// ClassNumber covers non-negative decimal integer literals.
	ClassNumber
	// ClassBool covers the boolean literals true and false.
	ClassBool
	// ClassBranch covers the loop-control keywords break and continue.
	ClassBranch
)

// Token is one mutable lexical unit discovered by a language backend. The
// byte interval [Offset, End) inside Path is the mutation site; Line and
// Column are kept for coverage joins and diagnostics.
type Token struct {
	Path   string
	Offset int
	End    int
	Line   int
	Column int
	Class  TokenClass
	Text   string
}

// Node returns the kernel-facing reference of this token in path:offset
// form. The format is identical across every language backend so reports
// stay comparable between toolchains.
func (t Token) Node() string {
	return fmt.Sprintf("%s:%d", t.Path, t.Offset)
}

// ParseNode splits a kernel-facing node reference back into its path and
// offset parts. The split happens on the last colon so workspace-relative
// paths never need escaping.
func ParseNode(node string) (string, int, error) {
	idx := strings.LastIndex(node, ":")
	if idx <= 0 || idx == len(node)-1 {
		return "", 0, fmt.Errorf("%w: %q", ErrInvalidNode, node)
	}
	offset, err := strconv.Atoi(node[idx+1:])
	if err != nil || offset < 0 {
		return "", 0, fmt.Errorf("%w: %q", ErrInvalidNode, node)
	}
	return node[:idx], offset, nil
}

// TestRef is one discovered test: its fully qualified identifier, the
// directory the runner must execute in and the runner-local name.
type TestRef struct {
	ID   string
	Dir  string
	Name string
}

// BoolText reports whether text spells a boolean literal in a language
// whose literals are exactly true and false.
func BoolText(text string) bool {
	return text == "true" || text == "false"
}

// IsMutableNumber reports whether text is a mutable integer literal: a
// non-negative decimal integer without any base prefix. Negative values
// are rejected because replacing the digits alone would leave a dangling
// sign; other bases are rejected because rewriting the digits would change
// the representation, not just the value.
func IsMutableNumber(text string) bool {
	if text == "" {
		return false
	}
	for _, r := range text {
		if r < '0' || r > '9' {
			return false
		}
	}
	if _, err := strconv.ParseInt(text, 10, 64); err != nil {
		return false
	}
	return true
}
