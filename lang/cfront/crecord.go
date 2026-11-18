/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package cfront

import (
	"abate/lang/air"
	"abate/lang/frontkit"
)

// This file holds the C/C++ constructs whose lowering depends on the
// surrounding declaration or record context: attribute prefixes, labels and
// record body members.

// cAttributeExtent returns the end index of the attribute prefix that begins
// the given token range. It returns lo when the range does not begin with an
// attribute prefix.
//
// Two spellings are recognized. The C++11 bracket form uses double square
// brackets. The GNU form uses the attribute or declspec keyword.
func cAttributeExtent(tokens []frontkit.Token, lo, hi int) int {
	if lo < 0 || lo+1 >= hi || lo+1 >= len(tokens) {
		return lo
	}
	if tokens[lo].Text == "[" && tokens[lo+1].Text == "[" {
		return cBalancedEnd(tokens, lo, hi, "[", "]")
	}
	switch tokens[lo].Text {
	case "__attribute__", "__declspec":
		if tokens[lo+1].Text != "(" {
			return lo + 1
		}
		return cBalancedEnd(tokens, lo+1, hi, "(", ")")
	default:
		return lo
	}
}

// cBalancedEnd returns the index just past the balanced group opened at lo.
func cBalancedEnd(tokens []frontkit.Token, lo, hi int, open, closeTok string) int {
	depth := 0
	for j := lo; j < hi && j < len(tokens); j++ {
		switch tokens[j].Text {
		case open:
			depth++
		case closeTok:
			depth--
			if depth == 0 {
				return j + 1
			}
		}
	}
	return lo
}

// cIsLabelUnit reports whether the range is a "label : statement" unit.
func cIsLabelUnit(tokens []frontkit.Token, lo, hi int) bool {
	if lo+1 >= hi || lo+1 >= len(tokens) {
		return false
	}
	if tokens[lo].Kind != frontkit.TokenIdent {
		return false
	}
	return tokens[lo+1].Text == ":"
}

// cRecordKind reports whether a mapping kind owns a record body whose
// members are fields rather than statements.
func cRecordKind(kind air.NodeKind) cBodyKind {
	switch kind {
	case air.KindStructDecl, air.KindClassDecl:
		return cBodyRecord
	case air.KindEnumDecl:
		return cBodyEnum
	default:
		return cBodyStatement
	}
}

// cMemberMapping maps a record body member to a field or a method.
//
// A member with a parameter list is a method declaration; every other member
// declaration is a field of the record.
func cMemberMapping(mapping cMapping, tokens []frontkit.Token,
	lo, hi int) cMapping {
	switch mapping.kind {
	case air.KindVarDecl, air.KindConstDecl:
		if cHasText(tokens, lo, hi, "(") {
			mapping.kind = air.KindMethodDecl
			return mapping
		}
		mapping.kind = air.KindField
		return mapping
	default:
		return mapping
	}
}
