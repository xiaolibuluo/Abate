/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package report

import (
	"encoding/json"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"

	"abate/internal/e2e/metrics"
)

// recordOptions is the complete configuration of one drilldown record.
type recordOptions struct {
	lang     string
	operator string
	sha      string
	captured bool
	outcome  string
}

// record builds one deterministic drilldown record for the report tests.
func record(o recordOptions) metrics.DrilldownRecord {
	return metrics.DrilldownRecord{
		Language: o.lang, FileSHA256: o.sha, RelPath: o.lang + "/a.go",
		OperatorID: o.operator, Class: metrics.ColumnText,
		Offset: 4, EndOffset: 6, BeforeText: "a+b", AfterText: "a-b",
		AIRIDBefore: "before-" + o.sha, AIRIDAfter: "after-" + o.sha,
		Captured: o.captured, OutcomeKind: o.outcome,
		AllocDeltaBytes: 4096, SilentDrops: 0,
	}
}

// reportRecords is the fixed record set shared by the report tests.
func reportRecords() []metrics.DrilldownRecord {
	return []metrics.DrilldownRecord{
		record(recordOptions{lang: "go", operator: "cmp", sha: "aa",
			captured: true, outcome: metrics.OutcomeOK}),
		record(recordOptions{lang: "go", operator: "digit", sha: "aa",
			outcome: metrics.OutcomeTimeout}),
		record(recordOptions{lang: "py", operator: "cmp", sha: "bb",
			captured: true, outcome: metrics.OutcomeOK}),
		record(recordOptions{lang: "py", operator: "block", sha: "cc",
			captured: true, outcome: metrics.OutcomeMemExceeded}),
	}
}

// sampleReport builds a report whose aggregates come from real records.
func sampleReport() Report {
	recs := reportRecords()
	sum := metrics.Aggregate(recs)
	precision := []metrics.PrecisionRow{
		{Language: "go", DegradeNodePct: 1.5},
		{Language: "py", DegradeNodePct: 2.5},
		{Language: ""},
	}
	facts := ManifestFacts{RegisteredFiles: 3, SPDXAllowlisted: 2,
		SPDXBoundary: 1}
	return Report{
		SchemaVersion: 1, ToolVersion: "test-fixed",
		PlanSHA256:       strings.Repeat("a", 64),
		ManifestSHA256:   strings.Repeat("b", 64),
		ThresholdsSHA256: strings.Repeat("c", 64),
		DrilldownSHA256:  strings.Repeat("d", 64),
		Seed:             7, Languages: []string{"go", "py"},
		MaxFileBytes: 4096, ParseTimeoutMillis: 5000,
		MemLimitBytes: 4096, MaxRuns: 10, Merge: "merge=1",
		Columns: sum.Columns, Mutation: sum.Mutation,
		Robustness: sum.Robustness, Equivalence: sum.Equivalence,
		Precision: precision, Localizers: []string{"file=aa|operator=cmp"},
		CompetitorAbsolute: "NOT_MEASURED",
		CoverageStatus:     []string{"go=MEASURED"},
		Facts:              facts, Counters: []Counter{{Key: "timeoutCount", Value: 0}},
		BaselineValues: DeriveBaseline(recs, precision, facts),
		ProcessedFiles: 3, RegisteredFiles: 3,
		IncrementalConsistent: true,
	}
}

// TestDigestBytesAndFile proves the deterministic digest helpers.
func TestDigestBytesAndFile(t *testing.T) {
	t.Parallel()
	if DigestBytes([]byte("abc")) !=
		"ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad" {
		t.Fatal("sha256(abc) mismatch")
	}
	dir := t.TempDir()
	path := filepath.Join(dir, "cfg.json")
	if err := os.WriteFile(path, []byte("{}"), 0o600); err != nil {
		t.Fatalf("write config: %v", err)
	}
	sum, err := DigestFile(path)
	if err != nil || sum != DigestBytes([]byte("{}")) {
		t.Fatalf("DigestFile = %s, %v", sum, err)
	}
	if _, err = DigestFile(filepath.Join(dir, "absent.json")); err == nil {
		t.Fatal("missing digest source accepted")
	}
}

// TestWriteProducesThreeArtifactsAndExactChain proves the writer emits the
// report, the drilldown and the markdown, with the chain matching the bytes.
func TestWriteProducesThreeArtifactsAndExactChain(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()
	out := filepath.Join(dir, "e2e_report.json")
	rep := sampleReport()
	recs := reportRecords()
	if err := Write(out, rep, recs); err != nil {
		t.Fatalf("Write: %v", err)
	}
	raw, err := fs.ReadFile(os.DirFS(dir), "e2e_report.json")
	if err != nil {
		t.Fatalf("read report: %v", err)
	}
	var decoded Report
	if err = json.Unmarshal(raw, &decoded); err != nil {
		t.Fatalf("decode report: %v", err)
	}
	drill, err := fs.ReadFile(os.DirFS(dir), "e2e_drilldown.jsonl")
	if err != nil {
		t.Fatalf("read drilldown: %v", err)
	}
	if decoded.DrilldownSHA256 != DigestBytes(drill) {
		t.Fatal("drilldown chain does not match the written bytes")
	}
	for _, name := range []string{
		"e2e_report.json", "e2e_drilldown.jsonl", "e2e_report.md",
	} {
		if _, err = fs.ReadFile(os.DirFS(dir), name); err != nil {
			t.Fatalf("missing artifact %s: %v", name, err)
		}
	}
	loaded, err := LoadDrilldown(filepath.Join(dir, "e2e_drilldown.jsonl"))
	if err != nil || len(loaded) != len(recs) {
		t.Fatalf("LoadDrilldown = %d, %v", len(loaded), err)
	}
	for i := range loaded {
		if loaded[i].AllocDeltaBytes != 0 {
			t.Fatal("runtime observation leaked into the drilldown artifact")
		}
	}
}

// TestLoadDrilldownRejectsBrokenArtifacts covers read and decode failures.
func TestLoadDrilldownRejectsBrokenArtifacts(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()
	if _, err := LoadDrilldown(filepath.Join(dir, "absent.jsonl")); err == nil {
		t.Fatal("missing drilldown accepted")
	}
	path := filepath.Join(dir, "bad.jsonl")
	if err := os.WriteFile(path, []byte("{oops}\n"), 0o600); err != nil {
		t.Fatalf("write bad line: %v", err)
	}
	if _, err := LoadDrilldown(path); err == nil {
		t.Fatal("malformed drilldown line accepted")
	}
}

// TestRecomputeZeroDifference proves every aggregate and baseline value is
// reproduced with an exact zero difference from the drilldown records.
func TestRecomputeZeroDifference(t *testing.T) {
	t.Parallel()
	rep := sampleReport()
	got, err := Recompute(rep, reportRecords())
	if err != nil {
		t.Fatalf("Recompute error is not zero: %v", err)
	}
	if got.Columns != rep.Columns || got.Robustness != rep.Robustness ||
		got.Equivalence != rep.Equivalence ||
		!slices.Equal(got.Mutation, rep.Mutation) ||
		!slices.Equal(got.BaselineValues, rep.BaselineValues) {
		t.Fatal("recomputed report differs from the original")
	}
}

// TestRecomputeDetectsEachDifference proves every compared section rejects a
// forged value with a named error instead of a silent correction.
func TestRecomputeDetectsEachDifference(t *testing.T) {
	t.Parallel()
	base := sampleReport()
	recs := reportRecords()
	variants := []Report{}
	change := func(edit func(*Report)) {
		candidate := base
		edit(&candidate)
		variants = append(variants, candidate)
	}
	change(func(r *Report) { r.Columns.Text.Gross++ })
	change(func(r *Report) { r.Mutation = []metrics.MutationRow{} })
	change(func(r *Report) { r.Robustness.TimeoutCount++ })
	change(func(r *Report) { r.Equivalence.Gross++ })
	change(func(r *Report) { r.BaselineValues = []BaselineValue{} })
	for i := range variants {
		if _, err := Recompute(variants[i], recs); err == nil {
			t.Fatalf("variant %d accepted", i)
		}
	}
	if err := compareBaseline(
		[]BaselineValue{{Key: "a", Value: 1}, {Key: "b", Value: 2}},
		[]BaselineValue{{Key: "b", Value: 2}, {Key: "a", Value: 1}},
	); err == nil {
		t.Fatal("reordered baseline accepted")
	}
}

// TestMarkdownRendersEverySection proves the human-readable report carries
// the four chains, four columns, both robustness counters and the P2 label.
func TestMarkdownRendersEverySection(t *testing.T) {
	t.Parallel()
	text := Markdown(sampleReport())
	for _, want := range []string{
		"# Abate E2E deterministic report", "planSha256",
		"manifestSha256", "thresholdsSha256", "drilldownSha256",
		"## Four separate columns", "| text |", "| equivalent |",
		"## Robustness", "| timeoutCount |", "| oomCount |",
		"## Precision by language", "## Mutation rows",
		"## Zero tolerance", "## Baseline current values",
		"competitorAbsolute: `NOT_MEASURED`",
		"incrementalConsistent", "intervalFindings",
	} {
		if !strings.Contains(text, want) {
			t.Fatalf("markdown missing %q", want)
		}
	}
}

// TestPureHelpersCoverBoundaryValues exercises every small deterministic
// helper with empty and mixed inputs.
func TestPureHelpersCoverBoundaryValues(t *testing.T) {
	t.Parallel()
	if pct(1, 0) != 0 || pct(1, 4) != 25 {
		t.Fatal("pct boundary wrong")
	}
	if minFilesPerLanguage(nil) != 0 ||
		minBrokenPerLanguage(nil) != 0 ||
		minGrossPerOperator(nil) != 0 ||
		completeRecords(nil) != 0 ||
		minBrokenCount(map[string]int{}) != 0 ||
		minLanguageCount(map[string]map[string]bool{}, 4) != 0 {
		t.Fatal("empty helper inputs must yield zero")
	}
	if _, found := baselineValue(Report{}, "missing"); found {
		t.Fatal("missing baseline key reported as present")
	}
	if !strings.Contains(localizerFor(Report{}), "file=") ||
		localizerFor(sampleReport()) == "" {
		t.Fatal("localizer fallback or real localizer missing")
	}
	values := sortedBaseline(map[string]float64{"z": 1, "a": 2})
	if len(values) != 2 || values[0].Key != "a" {
		t.Fatalf("baseline not sorted: %+v", values)
	}
	if !slices.Equal(sortedIntKeys(map[string]int{"b": 1, "a": 2}),
		[]string{"a", "b"}) ||
		!slices.Equal(sortedFloatKeys(map[string]float64{"b": 1, "a": 2}),
			[]string{"a", "b"}) {
		t.Fatal("key sorting wrong")
	}
	line := toDrilldownLine(record(recordOptions{lang: "go",
		operator: "cmp", sha: "aa", captured: true,
		outcome: metrics.OutcomeOK}))
	back := line.record()
	if back.OperatorID != "cmp" || back.AllocDeltaBytes != 0 {
		t.Fatalf("drilldown round trip = %+v", back)
	}
	if data, err := encodeDrilldown(reportRecords()); err != nil ||
		!strings.Contains(string(data), `"operatorID":"cmp"`) {
		t.Fatalf("encodeDrilldown = %s, %v", data, err)
	}
}
