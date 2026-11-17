/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package frontkit

import (
	"errors"
	"testing"

	"abate/lang/air"
)

// roundTripSource is the deterministic fixture of the position contract.
const roundTripSource = "let a = 1 // tail\nlet b = \"two\"\n"

// TestVerifyIntervalCoverageAcceptsAFaithfulTree proves a faithful tree
// satisfies the round-trip position contract.
func TestVerifyIntervalCoverageAcceptsAFaithfulTree(t *testing.T) {
	spec := newFakeSpec()
	tokens := mustLex(t, spec, roundTripSource)
	root, _, err := spec.ParseTokens(tokens, "a.go")
	if err != nil {
		t.Fatalf("ParseTokens: %v", err)
	}
	if err := VerifyIntervalCoverage(root, tokens, len(roundTripSource)); err != nil {
		t.Fatalf("interval coverage: %v", err)
	}
}

// TestVerifyIntervalCoverageRejectsBrokenTrees proves the contract actually
// rejects a root that does not cover the file and a tree that loses a token.
func TestVerifyIntervalCoverageRejectsBrokenTrees(t *testing.T) {
	spec := newFakeSpec()
	tokens := mustLex(t, spec, roundTripSource)
	root, _, err := spec.ParseTokens(tokens, "a.go")
	if err != nil {
		t.Fatalf("ParseTokens: %v", err)
	}
	if err := VerifyIntervalCoverage(nil, tokens, len(roundTripSource)); !errors.Is(err, ErrTokenFidelity) {
		t.Fatalf("a nil root must fail: %v", err)
	}
	if err := VerifyIntervalCoverage(root, tokens, len(roundTripSource)+1); !errors.Is(err, ErrTokenFidelity) {
		t.Fatalf("a short root must fail: %v", err)
	}
	hollow := air.NewNodeAt(air.KindPackage,
		air.NewPosition(1, 1, 0, len(roundTripSource)),
		air.NewPosition(1, 1, len(roundTripSource), len(roundTripSource)),
		"")
	if err := VerifyIntervalCoverage(hollow, tokens, len(roundTripSource)); !errors.Is(err, ErrTokenFidelity) {
		t.Fatalf("an empty tree must fail: %v", err)
	}
}
