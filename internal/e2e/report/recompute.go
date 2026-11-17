/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package report

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"abate/internal/e2e/metrics"
)

// errRecompute reports a non-zero difference between the report and the
// drilldown-derived recomputation.
var errRecompute = errors.New("report: recompute mismatch")

// BaselineValue is one recomputable current value of the threshold baseline.
type BaselineValue struct {
	Key   string  `json:"key"`
	Value float64 `json:"value"`
}

// DeriveBaseline recomputes every current baseline value from the drilldown
// records, the per-language precision rows and the manifest facts. It is the
// single source used by both the runner and Recompute, so the two agree by
// construction and Recompute can prove the difference is exactly zero.
func DeriveBaseline(recs []metrics.DrilldownRecord,
	precision []metrics.PrecisionRow, facts ManifestFacts) []BaselineValue {
	sum := metrics.Aggregate(recs)
	values := make(map[string]float64, 16)
	for _, pair := range []struct{ key, operator string }{
		{"airIDCmpCapturePct", "cmp"},
		{"airIDOpCapturePct", "op"},
		{"airIDDigitCapturePct", "digit"},
	} {
		gross, captured := 0, 0
		for _, row := range sum.Mutation {
			if row.OperatorID == pair.operator {
				gross += row.Gross
				captured += row.Captured
			}
		}
		values[pair.key] = pct(captured, gross)
	}
	values["parseableFilesPerLanguage"] = minFilesPerLanguage(recs)
	values["validMutationPointsPerOperator"] = minGrossPerOperator(sum.Mutation)
	values["brokenCorpusPerLanguage"] = minBrokenPerLanguage(recs)
	values["spdxCoveragePct"] = pct(
		facts.SPDXAllowlisted+facts.SPDXBoundary, facts.RegisteredFiles)
	values["drilldownCoveragePct"] = pct(completeRecords(recs), len(recs))
	values["manifestCoveragePct"] = pct(
		facts.RegisteredFiles-facts.SHA256Mismatches, facts.RegisteredFiles)
	values["timeoutRatePct"] = pct(sum.Robustness.TimeoutCount, len(recs))
	values["oomRatePct"] = pct(sum.Robustness.OOMCount, len(recs))
	for _, row := range precision {
		if len(row.Language) == 0 {
			continue
		}
		values["degrade"+strings.ToUpper(row.Language[:1])+
			row.Language[1:]+"Pct"] = row.DegradeNodePct
	}
	return sortedBaseline(values)
}

// Recompute rebuilds every aggregate and baseline value from the drilldown
// records and returns an error on any non-zero difference. Zero difference is
// asserted, never assumed.
func Recompute(r Report, recs []metrics.DrilldownRecord) (Report, error) {
	sum := metrics.Aggregate(recs)
	recomputed := r
	recomputed.Columns = sum.Columns
	recomputed.Mutation = sum.Mutation
	recomputed.Robustness = sum.Robustness
	recomputed.Equivalence = sum.Equivalence
	recomputed.BaselineValues = DeriveBaseline(recs, r.Precision, r.Facts)
	if err := compareReport(r, recomputed); err != nil {
		return Report{}, err
	}
	return recomputed, nil
}

// compareReport checks every recomputed section with an exact comparison.
func compareReport(want, got Report) error {
	if want.Columns != got.Columns {
		return fmt.Errorf("%w: columns want=%+v got=%+v",
			errRecompute, want.Columns, got.Columns)
	}
	if !slices.Equal(want.Mutation, got.Mutation) {
		return fmt.Errorf("%w: mutation rows differ want=%d got=%d",
			errRecompute, len(want.Mutation), len(got.Mutation))
	}
	if want.Robustness != got.Robustness {
		return fmt.Errorf("%w: robustness want=%+v got=%+v",
			errRecompute, want.Robustness, got.Robustness)
	}
	if want.Equivalence != got.Equivalence {
		return fmt.Errorf("%w: equivalence want=%+v got=%+v",
			errRecompute, want.Equivalence, got.Equivalence)
	}
	return compareBaseline(want.BaselineValues, got.BaselineValues)
}

// compareBaseline names the first key whose recomputed value differs.
func compareBaseline(want, got []BaselineValue) error {
	if !slices.Equal(want, got) {
		for _, value := range want {
			actual, found := baselineValue(
				Report{BaselineValues: got}, value.Key)
			if !found || actual != value.Value {
				return fmt.Errorf("%w: baseline %s want=%.6f got=%.6f",
					errRecompute, value.Key, value.Value, actual)
			}
		}
		return fmt.Errorf("%w: baseline ordering differs", errRecompute)
	}
	return nil
}

// LoadDrilldown decodes the canonical e2e_drilldown.jsonl back into records.
func LoadDrilldown(path string) ([]metrics.DrilldownRecord, error) {
	fsys := os.DirFS(filepath.Dir(path))
	data, err := fs.ReadFile(fsys, filepath.Base(path))
	if err != nil {
		return nil, fmt.Errorf("%w: read %s: %w", errRecompute, path, err)
	}
	out := make([]metrics.DrilldownRecord, 0)
	for _, line := range strings.Split(string(data), "\n") {
		text := strings.TrimSpace(line)
		if text == "" {
			continue
		}
		var decoded drilldownLine
		if err = json.Unmarshal([]byte(text), &decoded); err != nil {
			return nil, fmt.Errorf("%w: decode line: %w", errRecompute, err)
		}
		out = append(out, decoded.record())
	}
	return out, nil
}

// pct returns part divided by total in percent, or zero for an empty total.
func pct(part, total int) float64 {
	if total <= 0 {
		return 0
	}
	return float64(part) * 100 / float64(total)
}

// minFilesPerLanguage returns the smallest distinct-file count per language.
func minFilesPerLanguage(recs []metrics.DrilldownRecord) float64 {
	files := make(map[string]map[string]bool)
	for _, rec := range recs {
		if files[rec.Language] == nil {
			files[rec.Language] = make(map[string]bool)
		}
		files[rec.Language][rec.RelPath] = true
	}
	return minLanguageCount(files, len(recs))
}

// minBrokenPerLanguage returns the smallest broken-column count per language.
func minBrokenPerLanguage(recs []metrics.DrilldownRecord) float64 {
	counts := make(map[string]int)
	for _, rec := range recs {
		if rec.ColumnKey() == metrics.ColumnBroken {
			counts[rec.Language]++
		}
	}
	return minBrokenCount(counts)
}

// minBrokenCount returns the minimum broken count over the language keys.
func minBrokenCount(counts map[string]int) float64 {
	if len(counts) == 0 {
		return 0
	}
	minimum := -1
	for _, value := range counts {
		if minimum < 0 || value < minimum {
			minimum = value
		}
	}
	return float64(minimum)
}

// minLanguageCount returns the smallest per-language set size.
func minLanguageCount(files map[string]map[string]bool,
	recordCount int) float64 {
	if len(files) == 0 {
		return 0
	}
	minimum := recordCount + 1
	for _, set := range files {
		if len(set) < minimum {
			minimum = len(set)
		}
	}
	return float64(minimum)
}

// minGrossPerOperator returns the smallest gross denominator per operator.
func minGrossPerOperator(rows []metrics.MutationRow) float64 {
	gross := make(map[string]int)
	for _, row := range rows {
		gross[row.OperatorID] += row.Gross
	}
	if len(gross) == 0 {
		return 0
	}
	minimum := -1
	for _, value := range gross {
		if minimum < 0 || value < minimum {
			minimum = value
		}
	}
	return float64(minimum)
}

// completeRecords counts records carrying every recomputation key.
func completeRecords(recs []metrics.DrilldownRecord) int {
	count := 0
	for _, rec := range recs {
		if rec.FileSHA256 != "" && rec.OperatorID != "" &&
			rec.AIRIDBefore != "" && rec.AIRIDAfter != "" {
			count++
		}
	}
	return count
}

// sortedBaseline renders the values in ascending key order.
func sortedBaseline(values map[string]float64) []BaselineValue {
	keys := make([]string, 0, len(values))
	for key := range values {
		keys = append(keys, key)
	}
	slices.Sort(keys)
	out := make([]BaselineValue, 0, len(keys))
	for _, key := range keys {
		out = append(out, BaselineValue{Key: key, Value: values[key]})
	}
	return out
}
