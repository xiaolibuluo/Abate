/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package e2e

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"abate/internal/corpus"
	"abate/internal/e2e/metrics"
)

// TestRunnerRejectsEmptySample proves a byte window that selects no corpus
// file is returned as a hard error, never as a silently reduced report.
func TestRunnerRejectsEmptySample(t *testing.T) {
	t.Parallel()
	root := t.TempDir()
	manifest := buildTestManifest(t, root)
	plan := testPlan(root)
	plan.MinFileBytes = 1 << 40
	plan.MaxFileBytes = 1 << 41
	trace := NewTrace(filepath.Join(root, TraceFileName))
	runner, err := NewRunner(plan, manifest, trace)
	if err != nil {
		t.Fatalf("runner: %v", err)
	}
	if _, _, err = runner.Run(); !errors.Is(err, errEmptySample) {
		t.Fatalf("empty sample error = %v", err)
	}
}

// TestRunnerRejectsUnregisteredOperator proves a plan operator absent from
// the registry fails the run instead of being skipped.
func TestRunnerRejectsUnregisteredOperator(t *testing.T) {
	t.Parallel()
	root := t.TempDir()
	manifest := buildTestManifest(t, root)
	plan := testPlan(root)
	plan.Operators = []string{"unknown"}
	trace := NewTrace(filepath.Join(root, TraceFileName))
	runner, err := NewRunner(plan, manifest, trace)
	if err != nil {
		t.Fatalf("runner: %v", err)
	}
	if _, _, err = runner.Run(); !errors.Is(err, errMissingOperator) {
		t.Fatalf("operator error = %v", err)
	}
}

// TestNewRunnerRejectsBrokenConfigPaths proves both external definition files
// must load before any measurement starts.
func TestNewRunnerRejectsBrokenConfigPaths(t *testing.T) {
	t.Parallel()
	root := t.TempDir()
	manifest := buildTestManifest(t, root)
	plan := testPlan(root)
	plan.EquivalentPath = filepath.Join(root, "absent-equivalent.json")
	if _, err := NewRunner(plan, manifest, NewTrace(
		filepath.Join(root, TraceFileName))); err == nil {
		t.Fatal("missing equivalent definitions accepted")
	}
	plan = testPlan(root)
	plan.ConstructsPath = filepath.Join(root, "absent-constructs.json")
	if _, err := NewRunner(plan, manifest, NewTrace(
		filepath.Join(root, TraceFileName))); err == nil {
		t.Fatal("missing construct inventory accepted")
	}
}

// TestSourceListAndTraceErrors covers the read and flush error paths.
func TestSourceListAndTraceErrors(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()
	if _, err := readSourceList(filepath.Join(dir, "absent.txt")); err == nil {
		t.Fatal("missing source list accepted")
	}
	empty := filepath.Join(dir, "empty.txt")
	if err := os.WriteFile(empty, []byte("  \n\n"), FilePerm); err != nil {
		t.Fatalf("write empty list: %v", err)
	}
	if _, err := readSourceList(empty); err == nil {
		t.Fatal("empty source list accepted")
	}
	trace := NewTrace(filepath.Join(dir, "absent", "trace.jsonl"))
	for i := 0; i < traceFlushInterval; i++ {
		trace.Mark(corpus.Entry{Language: "go"}, "parse.before", 0)
	}
	if trace.Err() == nil {
		t.Fatal("trace write failure was swallowed")
	}
}

// TestSmallHelpersCoverDeterministicProjections exercises the pure helpers
// that assemble report facts from parsed states.
func TestSmallHelpersCoverDeterministicProjections(t *testing.T) {
	t.Parallel()
	states := []fileState{
		{spdxClass: "allowlisted", shaOK: true, intervalError: true,
			changedRanges: 2, consistent: true},
		{spdxClass: "boundary", shaOK: false, changedRanges: 3},
		{spdxClass: "violation", shaOK: true, consistent: true},
	}
	facts := manifestFacts(make([]corpus.Entry, 3), states)
	if facts.SPDXAllowlisted != 1 || facts.SPDXBoundary != 1 ||
		facts.SPDXViolations != 1 || facts.SHA256Mismatches != 1 ||
		facts.RegisteredFiles != 3 {
		t.Fatalf("facts = %+v", facts)
	}
	if intervalCount(states) != 1 || changedRanges(states) != 5 ||
		consistent(states) {
		t.Fatalf("state helpers disagree: %+v", states)
	}
	if rootIDOf(nil) != "" || len(localizersOf(nil)) != 0 {
		t.Fatal("empty projections must stay empty")
	}
	if recs := sortRecords([]metrics.DrilldownRecord{}); len(recs) != 0 {
		t.Fatal("empty record sort must stay empty")
	}
	if entries := sortedEntries([]corpus.Entry{}); len(entries) != 0 {
		t.Fatal("empty entry sort must stay empty")
	}
	lines, err := fs.ReadFile(os.DirFS(filepath.Join("..", "..", "ci")),
		SourcesName)
	if err != nil || !strings.Contains(string(lines), "cmd/e2e/main.go") {
		t.Fatalf("source anchor unreadable or incomplete: %v", err)
	}
}
