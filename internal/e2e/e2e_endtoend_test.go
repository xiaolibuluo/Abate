/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package e2e

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"abate/internal/corpus"
	"abate/internal/e2e/report"
)

// sourceFixture is one tiny, real, language-tagged source file.
type sourceFixture struct {
	language  string
	extension string
	body      string
}

// testSources is the seven-language synthetic corpus used by the end-to-end
// test. Every body is padded past the byte window minimum with real comments.
var testSources = []sourceFixture{
	{"go", "go", "package fixture\n\nfunc Add(a int, b int) int { return a + b + 16 }\n"},
	{"py", "py", "def add(a, b):\n    return a + b + 16\n"},
	{"js", "js", "function add(a, b) { return a + b + 16; }\n"},
	{"c", "c", "int add(int a, int b) { return a + b + 16; }\n"},
	{"cs", "cs", "class F { int Add(int a, int b) { return a + b + 16; } }\n"},
	{"rs", "rs", "fn add(a: i32, b: i32) -> i32 { a + b + 16 }\n"},
	{"kt", "kt", "fun add(a: Int, b: Int): Int { return a + b + 16 }\n"},
}

// TestDoubleRunIsByteIdentical proves the complete real pipeline (parse,
// mutate, reparse, damage, drilldown, report) is reproducible byte for byte.
func TestDoubleRunIsByteIdentical(t *testing.T) {
	t.Parallel()
	root := t.TempDir()
	manifest := buildTestManifest(t, root)
	plan := testPlan(root)
	first, records, err := runOnce(plan, manifest,
		filepath.Join(root, "run1"))
	if err != nil {
		t.Fatalf("first run failed: %v", err)
	}
	second, _, err := runOnce(plan, manifest,
		filepath.Join(root, "run2"))
	if err != nil {
		t.Fatalf("second run failed: %v", err)
	}
	if !bytes.Equal(first, second) {
		t.Fatalf("double run differs: %d vs %d bytes",
			len(first), len(second))
	}
	assertReport(t, first, filepath.Join(root, "run1"), records)
}

// runOnce constructs a fresh runner, executes it and returns the report bytes
// together with the number of real drilldown records produced.
func runOnce(plan Plan, manifest corpus.Manifest,
	dir string) ([]byte, int, error) {
	if err := os.MkdirAll(dir, DirPerm); err != nil {
		return nil, 0, err
	}
	trace := NewTrace(filepath.Join(dir, TraceFileName))
	runner, err := NewRunner(plan, manifest, trace)
	if err != nil {
		return nil, 0, err
	}
	rep, records, err := runner.Run()
	if err != nil {
		return nil, 0, err
	}
	if err = report.Write(filepath.Join(dir, "e2e_report.json"),
		rep, records); err != nil {
		return nil, 0, err
	}
	data, err := fs.ReadFile(os.DirFS(dir), "e2e_report.json")
	if err != nil {
		return nil, 0, err
	}
	return data, len(records), nil
}

// buildTestManifest writes the seven files and registers them with real
// SHA-256 digests and permissive SPDX declarations.
func buildTestManifest(t *testing.T, root string) corpus.Manifest {
	t.Helper()
	entries := make([]corpus.Entry, 0, len(testSources))
	for _, src := range testSources {
		pad := strings.Repeat("// padding\n", 40)
		if src.language == "py" {
			pad = strings.Repeat("# padding\n", 40)
		}
		content := []byte(src.body + pad)
		rel := filepath.Join(src.language, "sample."+src.extension)
		abs := filepath.Join(root, rel)
		if err := os.MkdirAll(filepath.Dir(abs), DirPerm); err != nil {
			t.Fatalf("mkdir: %v", err)
		}
		if err := os.WriteFile(abs, content, FilePerm); err != nil {
			t.Fatalf("write: %v", err)
		}
		sum := sha256.Sum256(content)
		entries = append(entries, corpus.Entry{
			Language: src.language, Project: "fixture",
			UpstreamURL: "https://example.invalid/fixture", Commit: "v1",
			SPDX: "MIT", SPDXEvidence: "LICENSE:fixture",
			RelPath: filepath.ToSlash(rel), Bytes: int64(len(content)),
			SHA256:        hex.EncodeToString(sum[:]),
			ConstructTags: []string{}, Source: corpus.SourceValidation,
		})
	}
	return corpus.Manifest{SchemaVersion: 1, Entries: entries}
}

// testPlan builds a small, fast yet fully real run configuration.
func testPlan(root string) Plan {
	ci := filepath.Join("..", "..", "ci")
	sum := sha256.Sum256([]byte("abate-e2e-fixture-plan"))
	return Plan{
		Seed:             20261114,
		Languages:        []string{"go", "py", "js", "c", "cs", "rs", "kt"},
		FilesPerLanguage: 1, MinFileBytes: 1, MaxFileBytes: 4096,
		StrataCount: 2, ParseTimeoutMillis: 5000,
		MemLimitBytes: 256 * corpus.Mib, MaxRuns: 4096,
		BrokenPerLanguage: 3,
		Operators: []string{"block", "call", "cmp", "cond", "digit",
			"literal", "op", "return"},
		PlanSHA256: hex.EncodeToString(sum[:]), CorpusRoot: root,
		RepoRoot:       filepath.Join("..", ".."),
		ThresholdsPath: filepath.Join(ci, ThresholdsName),
		BaselinePath:   filepath.Join(ci, BaselineName),
		ConstructsPath: filepath.Join(ci, ConstructsName),
		EquivalentPath: filepath.Join(ci, EquivalentName),
		SourcesPath:    filepath.Join(ci, SourcesName),
	}
}

// assertReport validates the four chains, the four columns, the P2 label and
// the zero-error recomputation from the written drilldown file.
func assertReport(t *testing.T, raw []byte,
	dir string, records int) {
	t.Helper()
	var decoded report.Report
	if err := json.Unmarshal(raw, &decoded); err != nil {
		t.Fatalf("report decode: %v", err)
	}
	for name, value := range map[string]string{
		"planSha256":       decoded.PlanSHA256,
		"manifestSha256":   decoded.ManifestSHA256,
		"thresholdsSha256": decoded.ThresholdsSHA256,
		"drilldownSha256":  decoded.DrilldownSHA256,
	} {
		if len(value) != 64 {
			t.Fatalf("%s is not a sha256 digest: %q", name, value)
		}
	}
	if decoded.CompetitorAbsolute != CompetitorNotMeasured {
		t.Fatalf("competitor label = %q", decoded.CompetitorAbsolute)
	}
	if len(decoded.Languages) != 7 || records == 0 {
		t.Fatalf("incomplete run: languages=%d records=%d",
			len(decoded.Languages), records)
	}
	if decoded.Columns.Text.Gross == 0 ||
		decoded.Columns.Structure.Gross == 0 ||
		decoded.Columns.Broken.Gross == 0 ||
		decoded.Columns.Equivalent.Gross == 0 {
		t.Fatalf("four columns incomplete: %+v", decoded.Columns)
	}
	if decoded.Robustness.ProcessedFiles !=
		decoded.Robustness.RegisteredFiles {
		t.Fatal("ProcessedFiles != RegisteredFiles")
	}
	drilldown, err := report.LoadDrilldown(
		filepath.Join(dir, "e2e_drilldown.jsonl"))
	if err != nil {
		t.Fatalf("load drilldown: %v", err)
	}
	if _, err = report.Recompute(decoded, drilldown); err != nil {
		t.Fatalf("recompute error is not zero: %v", err)
	}
	for _, file := range []string{
		"e2e_report.json", "e2e_trace.jsonl", "e2e_drilldown.jsonl",
	} {
		if _, err = fs.ReadFile(os.DirFS(dir), file); err != nil {
			t.Fatalf("missing artifact %s: %v", file, err)
		}
	}
}

// TestEvaluateTightenOnlyAndZeroTolerance proves the monotone verdict and
// the Failure{Key,Rule,Localizer,Actual,Expected} localisation.
func TestEvaluateTightenOnlyAndZeroTolerance(t *testing.T) {
	t.Parallel()
	root := t.TempDir()
	manifest := buildTestManifest(t, root)
	raw, _, err := runOnce(testPlan(root), manifest,
		filepath.Join(root, "run"))
	if err != nil {
		t.Fatalf("run failed: %v", err)
	}
	var decoded report.Report
	if err = json.Unmarshal(raw, &decoded); err != nil {
		t.Fatalf("decode: %v", err)
	}
	thresholds := permissiveThresholds(decoded)
	verdict := report.Evaluate(thresholds, thresholds, decoded)
	if !verdict.Passed {
		t.Fatalf("tighten-only verdict failed: %s", verdict.String())
	}
	relaxed := permissiveThresholds(decoded)
	for key := range relaxed.Minima {
		relaxed.Minima[key] = -1
	}
	if verdict = report.Evaluate(relaxed, thresholds, decoded); verdict.Passed {
		t.Fatal("relaxed minima must fail")
	}
	hot := decoded
	hot.Counters = append([]report.Counter{}, decoded.Counters...)
	for i := range hot.Counters {
		if hot.Counters[i].Key == "timeoutCount" {
			hot.Counters[i].Value = 1
		}
	}
	verdict = report.Evaluate(thresholds, thresholds, hot)
	if verdict.Passed || len(verdict.Failures) == 0 {
		t.Fatal("non-zero zero-tolerance counter must fail")
	}
	if !strings.Contains(verdict.Failures[0].Localizer, "file=") {
		t.Fatalf("failure localizer missing file sha256: %s",
			verdict.Failures[0].String())
	}
}

// permissiveThresholds derives a passing threshold set from real report data.
func permissiveThresholds(r report.Report) report.Thresholds {
	zeros := make(map[string]int, len(r.Counters))
	for _, counter := range r.Counters {
		zeros[counter.Key] = 0
	}
	minima := make(map[string]float64)
	maxima := make(map[string]float64)
	for _, value := range r.BaselineValues {
		if strings.HasPrefix(value.Key, "degrade") ||
			strings.HasSuffix(value.Key, "RatePct") {
			maxima[value.Key] = 100
		} else {
			minima[value.Key] = 0
		}
	}
	return report.Thresholds{
		SchemaVersion: 1, TightenOnly: true,
		ZeroTolerance: zeros, Minima: minima, Maxima: maxima,
	}
}
