/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/
package gomutate

import (
	"strconv"
	"strings"
)

// Lexical classes of a mutable token. The classes mirror the operator
// dispatch categories of the mutation registry: every discovered token must
// be rewritable into a compilable Go source, so classes whose mutation
// requires type-aware rewriting (nil literals, unary operators) are
// deliberately not discovered.
const (
	// ClassOperator marks a binary operator token of a BinaryExpr.
	ClassOperator = "operator"
	// ClassNumber marks an integer basic literal.
	ClassNumber = "number"
	// ClassBool marks the built-in true and false identifiers.
	ClassBool = "bool"
	// ClassBranch marks a break or continue branch statement.
	ClassBranch = "branch"
	// ClassCall marks a call expression used as an expression statement.
	ClassCall = "call"
)

// Token is one mutable lexical token discovered in a Go source file. The
// byte interval [Offset, End) is the exact replacement range used when a
// mutant is materialized on disk, and Line/Column are kept for coverage
// profile matching, which reports blocks in line/column coordinates.
type Token struct {
	// Path is the slash-separated path relative to the workspace root.
	Path string
	// Offset is the start byte offset of the token inside its file.
	Offset int
	// End is the end byte offset of the token inside its file.
	End int
	// Line is the 1-based line of the token start.
	Line int
	// Column is the 1-based column of the token start.
	Column int
	// Class is one of the Class* constants above.
	Class string
	// Text is the original source text of the token.
	Text string
}

// Node returns the kernel-facing node reference in "path:offset" form.
// The form is stable for a fixed workspace snapshot, which the mutation
// kernel requires for deterministic point enumeration.
func (t Token) Node() string {
	return t.Path + ":" + strconv.Itoa(t.Offset)
}

// parseNode decodes one "path:offset" node reference, splitting on the last
// colon so paths never need escaping on supported platforms.
func parseNode(node string) (string, int, bool) {
	i := strings.LastIndexByte(node, ':')
	if i <= 0 || i+1 >= len(node) {
		return "", 0, false
	}
	offset, err := strconv.Atoi(node[i+1:])
	if err != nil || offset < 0 {
		return "", 0, false
	}
	return node[:i], offset, true
}
