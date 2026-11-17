/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package ktfront

import (
	"abate/lang/frontkit"
)

// ktUnitBoundary reports a unit/statement boundary at zero bracket depth,
// returning the offset to stop at and whether one was found.
func ktUnitBoundary(tok frontkit.Token, j, paren, bracket, brace int) (int, bool) {
	if tok.Kind == frontkit.TokenEOF {
		return j, true
	}
	if tok.Kind == frontkit.TokenNewline && paren == 0 &&
		bracket == 0 && brace == 0 {
		return j + 1, true
	}
	return 0, false
}

// ktBalanceOrTerminate updates bracket depths and body bounds. It returns true
// when tok terminates the construct (a closing brace or a semicolon at zero
// depth) together with the offset to stop at. Kotlin statements end at a
// newline (handled by ktUnitBoundary), never at a top-level comma.
func (p *ktParser) ktBalanceOrTerminate(tok frontkit.Token, j, hi int,
	paren, bracket, brace, bodyOpen, bodyClose *int) (bool, int) {
	switch tok.Text {
	case "(":
		*paren++
	case ")":
		if *paren > 0 {
			*paren--
		}
	case "[":
		*bracket++
	case "]":
		if *bracket > 0 {
			*bracket--
		}
	case "{":
		if *brace == 0 && *bodyOpen < 0 {
			*bodyOpen = j
		}
		*brace++
	case "}":
		if *brace > 0 {
			*brace--
		}
		if *brace == 0 && *bodyOpen >= 0 {
			next := j + 1
			if next < hi && p.tokens[next].Text == ";" {
				next++
			}
			*bodyClose = j
			return true, next
		}
	case ";":
		if *paren == 0 && *bracket == 0 && *brace == 0 {
			return true, j + 1
		}
	}
	return false, 0
}
