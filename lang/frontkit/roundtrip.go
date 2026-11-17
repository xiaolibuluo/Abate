/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package frontkit

import (
	"fmt"

	"abate/lang/air"
)

// VerifyIntervalCoverage proves the round-trip position contract of one file.
//
// The contract has three parts. The root must cover the whole source, so the
// tree can be replayed against the file. Every node must stay inside its
// parent and inside the file, and no sibling may overlap another one, so no
// interval runs out of the source or over a neighbour. Finally every
// preserved lexical token must be contained by some node of the tree, so no
// identifier, literal or comment is lost between two nodes.
//
// The preserved subset is the lexical one, because a neutral intermediate
// representation does not model keywords, operators and punctuation as
// vertices. Together the three parts are the "no gap, no out-of-range
// overlap" contract of the language frontend milestone.
func VerifyIntervalCoverage(root air.Node, tokens []Token, size int) error {
	if root == nil {
		return fmt.Errorf("%w: nil root", ErrTokenFidelity)
	}
	pos := root.Pos()
	if pos.Offset != 0 || pos.EndOffset != size {
		return fmt.Errorf("%w: root span [%d,%d) does not cover [0,%d)",
			ErrTokenFidelity, pos.Offset, pos.EndOffset, size)
	}
	if err := air.CheckIntervals(root, size); err != nil {
		return fmt.Errorf("%w: %w", ErrTokenFidelity, err)
	}
	inner := make([]air.Node, 0, 32)
	air.Walk(root, func(node air.Node) bool {
		if node != root {
			inner = append(inner, node)
		}
		return true
	})
	for _, tok := range tokens {
		if !isLexicalToken(tok.Kind) || tok.Text == "" {
			continue
		}
		if !tokenCovered(inner, tok) {
			return fmt.Errorf("%w: token %q at [%d,%d) has no covering node",
				ErrTokenFidelity, tok.Text, tok.Offset, tok.EndOffset)
		}
	}
	return nil
}

// tokenCovered reports whether some node contains the whole token interval.
func tokenCovered(nodes []air.Node, tok Token) bool {
	for _, node := range nodes {
		start, end := air.Span(node)
		if start <= tok.Offset && tok.EndOffset <= end {
			return true
		}
	}
	return false
}
