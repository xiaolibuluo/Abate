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

// loadReference reads the vendored PIT report of the test run.
func loadReference(t *testing.T) []Mutation {
	t.Helper()
	path, err := LocateReport(".")
	if err != nil {
		t.Fatalf("locate reference report: %v", err)
	}
	records, err := LoadMutations(path)
	if err != nil {
		t.Fatalf("load reference report: %v", err)
	}
	if len(records) == 0 {
		t.Fatal("reference report holds no mutation")
	}
	return records
}

// TestReferenceReportShape pins the size and status mix of the reference.
func TestReferenceReportShape(t *testing.T) {
	t.Parallel()
	records := loadReference(t)
	if len(records) != 28 {
		t.Fatalf("reference report holds %d mutations, want 28", len(records))
	}
	counts := map[string]int{}
	for _, record := range records {
		counts[record.Status]++
	}
	if counts[StatusKilled] != 23 || counts[StatusSurvived] != 3 ||
		counts[StatusNoCoverage] != 2 {
		t.Fatalf("status mix = %v, want 23 killed, 3 survived, 2 no coverage", counts)
	}
}

// TestMappingCoversEveryMutation checks that no mutation is silently dropped.
func TestMappingCoversEveryMutation(t *testing.T) {
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
		key := record.MutatedMethod + ":" + record.Mutator + ":" + itoa(record.LineNumber)
		if !registered[key] {
			t.Fatalf("mutation %s is not registered in the mapping table", key)
		}
	}
	if len(mappingTable) != len(records) {
		t.Fatalf("mapping table holds %d entries for %d mutations",
			len(mappingTable), len(records))
	}
}

// TestWitnessReproducesPITStatus checks the independent semantic oracle.
//
// The witness never looks at the PIT status, so agreeing with it on all 28
// mutations proves the ported baseline and the mapping are semantically right.
func TestWitnessReproducesPITStatus(t *testing.T) {
	t.Parallel()
	records := loadReference(t)
	for _, record := range records {
		caseEntry, ok := witnessFor(record.MutatedMethod, record.Mutator, record.LineNumber)
		if !ok {
			t.Fatalf("no witness for %s:%d %s",
				record.MutatedMethod, record.LineNumber, record.Mutator)
		}
		witness, err := classifyWitness(caseEntry)
		if err != nil {
			t.Fatalf("classify %s:%d %s: %v",
				record.MutatedMethod, record.LineNumber, record.Mutator, err)
		}
		if !witnessMatchesStatus(witness, record.Status) {
			t.Fatalf("%s:%d %s: witness %q does not match PIT status %q",
				record.MutatedMethod, record.LineNumber, record.Mutator,
				witness, record.Status)
		}
	}
}

// TestMappingNodesFollowTheWitness proves the mapping is not outcome tuned.
//
// A comparable mapping must pick the class node, except when the witness says
// the mutant is equivalent or unreachable, in which case the dedicated node
// of that branch is the only admissible choice.
func TestMappingNodesFollowTheWitness(t *testing.T) {
	t.Parallel()
	records := loadReference(t)
	for _, record := range records {
		mapping, err := requireMapping(record.MutatedMethod, record.Mutator, record.LineNumber)
		if err != nil {
			t.Fatalf("require mapping: %v", err)
		}
		if mapping.Excluded {
			continue
		}
		caseEntry, ok := witnessFor(record.MutatedMethod, record.Mutator, record.LineNumber)
		if !ok {
			t.Fatalf("no witness for %s:%d %s",
				record.MutatedMethod, record.LineNumber, record.Mutator)
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
				record.MutatedMethod, record.LineNumber, record.Mutator,
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
