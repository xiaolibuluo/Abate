/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package pitparity

import (
	"testing"

	"abate/kernel/verify"
)

// TestVerdictDifferenceExplainsEveryStatus covers the status mapping rules.
//
// Every branch of the mapping must name the mismatch it reports, so a review
// of the parity report never sees a silent difference.
func TestVerdictDifferenceExplainsEveryStatus(t *testing.T) {
	t.Parallel()
	target, suite := newFixture()
	if _, err := newVerdictIndex(mutationReport(t, target, suite)); err != nil {
		t.Fatalf("index: %v", err)
	}
	mapping := Mapping{Node: "a-arith", Class: "arith-replace"}
	if reason := verdictDifference(verify.MutantKilled, StatusKilled, WitnessDetected,
		suite, mapping); reason != "" {
		t.Fatalf("killed point reported %q", reason)
	}
	if reason := verdictDifference(verify.MutantSurvived, StatusKilled, WitnessDetected,
		suite, mapping); reason == "" {
		t.Fatal("a surviving mutant was accepted for a killed status")
	}
	equivalent := Mapping{Node: "h-equivalent", Class: "arith-replace"}
	if reason := verdictDifference(verify.MutantEquivalent, StatusSurvived,
		WitnessEquivalent, suite, equivalent); reason != "" {
		t.Fatalf("equivalent point reported %q", reason)
	}
	if reason := verdictDifference(verify.MutantSurvived, StatusSurvived,
		WitnessEquivalent, suite, equivalent); reason == "" {
		t.Fatal("a plain survivor was accepted for an equivalent status")
	}
	if reason := verdictDifference(verify.MutantSurvived, StatusSurvived,
		WitnessDetected, suite, equivalent); reason == "" {
		t.Fatal("a detected mutant was accepted for a survived status")
	}
	uncovered := Mapping{Node: "j-survived", Class: "arith-replace"}
	if reason := verdictDifference(verify.MutantSurvived, StatusNoCoverage,
		WitnessUncovered, suite, uncovered); reason != "" {
		t.Fatalf("no coverage point reported %q", reason)
	}
	if reason := verdictDifference(verify.MutantKilled, StatusNoCoverage,
		WitnessUncovered, suite, uncovered); reason == "" {
		t.Fatal("a killed mutant was accepted for a no coverage status")
	}
	if reason := verdictDifference(verify.MutantSurvived, StatusNoCoverage,
		WitnessUncovered, suite, mapping); reason == "" {
		t.Fatal("a covered node accepted a no coverage status")
	}
	if reason := verdictDifference(verify.MutantTimeout, "WEIRD", WitnessDetected,
		suite, mapping); reason == "" {
		t.Fatal("an unknown status was accepted")
	}
}

// TestVerdictIndexLookupModes covers the point level and node level lookups.
func TestVerdictIndexLookupModes(t *testing.T) {
	t.Parallel()
	target, suite := newFixture()
	index, err := newVerdictIndex(mutationReport(t, target, suite))
	if err != nil {
		t.Fatalf("index: %v", err)
	}
	verdict, mode, ok := index.lookup(Mapping{Node: "a-arith", Class: "arith-replace"})
	if !ok || verdict != verify.MutantKilled || mode != VerdictModePoint {
		t.Fatalf("lookup = %v/%v/%v", verdict, mode, ok)
	}
	verdict, mode, ok = index.lookup(Mapping{Node: "j-survived", Class: "constant-perturb"})
	if !ok || verdict != verify.MutantSurvived || mode != VerdictModeNode {
		t.Fatalf("node lookup = %v/%v/%v", verdict, mode, ok)
	}
	if _, _, ok = index.lookup(Mapping{Node: "missing", Class: "arith-replace"}); ok {
		t.Fatal("lookup found a verdict for an unknown node")
	}
}
