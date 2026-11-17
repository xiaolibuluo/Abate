/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package pitparity

import (
	"errors"
	"strings"
	"testing"
)

// primaryNode maps an Abate operator class to its canonical fixture node.
//
// The mapping table must derive its node from this rule, except for the
// equivalent and uncovered branches, which are derived from the witness.
var primaryNode = map[string]string{
	"arith-replace":    "a-arith",
	"boundary-value":   "b-bound",
	"constant-perturb": "c-const",
	"control-flow":     "d-control",
	"return-value":     "e-return",
	"negate-condition": "f-negate",
	"remove-call":      "g-call",
}

// loadReference reads the frozen self-developed golden of the test run.
func loadReference(t *testing.T) []Case {
	t.Helper()
	path, err := LocateGolden(".")
	if err != nil {
		t.Fatalf("locate golden: %v", err)
	}
	records, err := LoadGolden(path)
	if err != nil {
		t.Fatalf("load golden: %v", err)
	}
	if len(records) == 0 {
		t.Fatal("golden holds no case")
	}
	return records
}

// TestGoldenShape pins the size and expectation mix of the frozen golden.
func TestGoldenShape(t *testing.T) {
	t.Parallel()
	records := loadReference(t)
	if len(records) != 28 {
		t.Fatalf("golden holds %d cases, want 28", len(records))
	}
	counts := map[string]int{}
	for _, record := range records {
		counts[record.Expect]++
	}
	if counts[ExpectKilled] != 23 || counts[ExpectSurvived] != 3 ||
		counts[ExpectNoCoverage] != 2 {
		t.Fatalf("expectation mix = %v, want 23 killed, 3 survived, 2 no coverage", counts)
	}
}

// TestMappingCoversEveryCase checks that no golden case is silently dropped.
func TestMappingCoversEveryCase(t *testing.T) {
	t.Parallel()
	records := loadReference(t)
	registered := make(map[string]bool, len(mappingTable))
	for _, entry := range mappingTable {
		key := entry.Method + ":" + entry.Mutator + ":" + itoa(entry.Line)
		if registered[key] {
			t.Fatalf("duplicate mapping for %s", key)
		}
		registered[key] = true
		if strings.TrimSpace(entry.Reason) == "" {
			t.Fatalf("mapping %s has no reason", key)
		}
		if entry.Excluded {
			continue
		}
		if entry.Node == "" || entry.Class == "" {
			t.Fatalf("comparable mapping %s has no node or class", key)
		}
		if _, ok := primaryNode[entry.Class]; !ok {
			t.Fatalf("mapping %s uses unknown class %q", key, entry.Class)
		}
	}
	for _, record := range records {
		key := record.Method + ":" + record.Mutator + ":" + itoa(record.Line)
		if !registered[key] {
			t.Fatalf("case %s is not registered in the mapping table", key)
		}
	}
	if len(mappingTable) != len(records) {
		t.Fatalf("mapping table holds %d entries for %d cases",
			len(mappingTable), len(records))
	}
}

// TestWitnessReproducesFrozenExpectation checks the independent semantic oracle.
//
// The witness never looks at the frozen expectation, so agreeing with it on all
// 28 cases proves the baseline port and the mapping are semantically right.
func TestWitnessReproducesFrozenExpectation(t *testing.T) {
	t.Parallel()
	records := loadReference(t)
	for _, record := range records {
		caseEntry, ok := witnessFor(record.Method, record.Mutator, record.Line)
		if !ok {
			t.Fatalf("no witness for %s:%d %s",
				record.Method, record.Line, record.Mutator)
		}
		witness, err := classifyWitness(caseEntry)
		if err != nil {
			t.Fatalf("classify %s:%d %s: %v",
				record.Method, record.Line, record.Mutator, err)
		}
		if !witnessMatchesExpect(witness, record.Expect) {
			t.Fatalf("%s:%d %s: witness %q does not match frozen expectation %q",
				record.Method, record.Line, record.Mutator,
				witness, record.Expect)
		}
	}
}

// TestMappingNodesFollowTheWitness proves the mapping is not outcome tuned.
//
// A comparable mapping must pick the class node, except when the witness says
// the case is equivalent or unreachable, in which case the dedicated node of
// that branch is the only admissible choice.
func TestMappingNodesFollowTheWitness(t *testing.T) {
	t.Parallel()
	records := loadReference(t)
	for _, record := range records {
		mapping, err := requireMapping(record.Method, record.Mutator, record.Line)
		if err != nil {
			t.Fatalf("require mapping: %v", err)
		}
		if mapping.Excluded {
			continue
		}
		caseEntry, ok := witnessFor(record.Method, record.Mutator, record.Line)
		if !ok {
			t.Fatalf("no witness for %s:%d %s",
				record.Method, record.Line, record.Mutator)
		}
		witness, err := classifyWitness(caseEntry)
		if err != nil {
			t.Fatalf("classify: %v", err)
		}
		want := primaryNode[mapping.Class]
		switch witness {
		case WitnessEquivalent:
			want = "h-equivalent"
		case WitnessUncovered:
			want = "j-survived"
		case WitnessDetected:
		default:
			t.Fatalf("unexpected witness %q", witness)
		}
		if mapping.Node != want {
			t.Fatalf("%s:%d %s maps to node %q, the witness derives %q",
				record.Method, record.Line, record.Mutator,
				mapping.Node, want)
		}
	}
}

// TestRequireMappingRejectsUnknownSites covers the missing mapping branch.
func TestRequireMappingRejectsUnknownSites(t *testing.T) {
	t.Parallel()
	_, err := requireMapping("nowhere", "NoMutator", 42)
	if !errors.Is(err, ErrUnmappedMutation) {
		t.Fatalf("error = %v, want ErrUnmappedMutation", err)
	}
	if _, ok := lookupMapping("nowhere", "NoMutator", 42); ok {
		t.Fatal("lookup found a mapping that was never registered")
	}
}
