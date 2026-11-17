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
)

// errBadThresholds reports a threshold document that cannot be evaluated.
var errBadThresholds = errors.New("report: invalid thresholds")

type (
	// Thresholds is the external, tighten-only threshold document. Maps are
	// used only for configuration lookup; no map is ever serialised into the
	// compared report bytes.
	Thresholds struct {
		SchemaVersion int                `json:"schemaVersion"`
		TightenOnly   bool               `json:"tightenOnly"`
		ZeroTolerance map[string]int     `json:"zeroTolerance"`
		Minima        map[string]float64 `json:"minima"`
		Maxima        map[string]float64 `json:"maxima"`
	}

	// Counter is one sorted zero-tolerance observation of the run.
	Counter struct {
		Key   string `json:"key"`
		Value int    `json:"value"`
	}

	// Failure is one zero-tolerance or tighten-only violation. Localizer
	// carries the file sha256, the operator and both AIR identifiers.
	Failure struct {
		Key       string  `json:"key"`
		Rule      string  `json:"rule"`
		Localizer string  `json:"localizer"`
		Actual    float64 `json:"actual"`
		Expected  float64 `json:"expected"`
	}

	// Verdict is the result of the tighten-only evaluation.
	Verdict struct {
		Passed   bool      `json:"passed"`
		Failures []Failure `json:"failures"`
	}
)

// evaluation bundles the immutable inputs of one verdict computation so every
// helper method stays within the three-parameter limit.
type evaluation struct {
	current  Thresholds
	baseline Thresholds
	report   Report
	loc      string
	observed map[string]int
	failures []Failure
}

// String renders one violation in the fixed Failure{Key,Rule,Localizer,
// Actual,Expected} form used by the gate's zero-tolerance localisation.
func (f Failure) String() string {
	return fmt.Sprintf(
		"Failure{Key:%s Rule:%s Localizer:%s Actual:%.6f Expected:%.6f}",
		f.Key, f.Rule, f.Localizer, f.Actual, f.Expected)
}

// String renders every violation on its own deterministic line.
func (v Verdict) String() string {
	lines := make([]string, 0, len(v.Failures))
	for _, failure := range v.Failures {
		lines = append(lines, failure.String())
	}
	return strings.Join(lines, "\n")
}

// LoadThresholds decodes an external threshold document from a relative path.
func LoadThresholds(path string) (Thresholds, error) {
	fsys := os.DirFS(filepath.Dir(path))
	data, err := fs.ReadFile(fsys, filepath.Base(path))
	if err != nil {
		return Thresholds{}, fmt.Errorf("%w: read %s: %w",
			errBadThresholds, path, err)
	}
	var out Thresholds
	if err = json.Unmarshal(data, &out); err != nil {
		return Thresholds{}, fmt.Errorf("%w: decode %s: %w",
			errBadThresholds, path, err)
	}
	return out, nil
}

// Evaluate implements the tighten-only invariant and the zero-tolerance gate.
// Zero-tolerance keys must match the baseline exactly and every observed
// counter must be zero; minima may only increase and maxima may only decrease;
// unregistered keys are refused. No failure is ever offset by another column.
func Evaluate(current, baseline Thresholds, r Report) Verdict {
	e := evaluation{
		current: current, baseline: baseline, report: r,
		loc: localizerFor(r), observed: make(map[string]int),
		failures: make([]Failure, 0),
	}
	for _, counter := range r.Counters {
		e.observed[counter.Key] = counter.Value
	}
	if !current.TightenOnly {
		e.add("tightenOnly", "tighten-only-required", 0, 1)
	}
	e.checkZero()
	e.checkMin()
	e.checkMax()
	return Verdict{Passed: len(e.failures) == 0, Failures: e.failures}
}

// add appends one violation with the fixed localizer.
func (e *evaluation) add(key, rule string, actual, expected float64) {
	e.failures = append(e.failures, Failure{
		Key: key, Rule: rule, Localizer: e.loc,
		Actual: actual, Expected: expected,
	})
}

// checkZero compares the zero-tolerance key sets and every observed counter.
func (e *evaluation) checkZero() {
	keys := sortedIntKeys(e.current.ZeroTolerance)
	baseKeys := sortedIntKeys(e.baseline.ZeroTolerance)
	if !slices.Equal(keys, baseKeys) {
		e.add("zeroTolerance", "zero-tolerance-key-set",
			float64(len(keys)), float64(len(baseKeys)))
	}
	for _, key := range keys {
		if e.current.ZeroTolerance[key] != 0 {
			e.add(key, "zero-tolerance-threshold-nonzero",
				float64(e.current.ZeroTolerance[key]), 0)
		}
		value, found := e.observed[key]
		if !found {
			e.add(key, "missing-zero-tolerance-counter", 0, 0)
			continue
		}
		if value != 0 {
			e.add(key, "zero-tolerance-nonzero", float64(value), 0)
		}
	}
}

// checkMin enforces monotone increase plus the current minimum per key.
func (e *evaluation) checkMin() {
	for _, key := range sortedFloatKeys(e.current.Minima) {
		bound := e.current.Minima[key]
		base, found := e.baseline.Minima[key]
		if !found {
			e.add(key, "unregistered-minima-key", bound, 0)
			continue
		}
		if bound < base {
			e.add(key, "relaxed-minima", bound, base)
		}
		e.checkBound(key, "minima-not-met", bound)
	}
}

// checkMax enforces monotone decrease plus the current maximum per key.
func (e *evaluation) checkMax() {
	for _, key := range sortedFloatKeys(e.current.Maxima) {
		bound := e.current.Maxima[key]
		base, found := e.baseline.Maxima[key]
		if !found {
			e.add(key, "unregistered-maxima-key", bound, 0)
			continue
		}
		if bound > base {
			e.add(key, "relaxed-maxima", bound, base)
		}
		e.checkBound(key, "maxima-exceeded", bound)
	}
}

// checkBound compares one recomputable baseline value against its bound.
func (e *evaluation) checkBound(key, rule string, bound float64) {
	value, found := baselineValue(e.report, key)
	if !found {
		e.add(key, "missing-baseline-value", 0, bound)
		return
	}
	if rule == "minima-not-met" && value < bound {
		e.add(key, rule, value, bound)
	}
	if rule == "maxima-exceeded" && value > bound {
		e.add(key, rule, value, bound)
	}
}

// baselineValue looks one recomputable key up in the sorted report slice.
func baselineValue(r Report, key string) (float64, bool) {
	for _, value := range r.BaselineValues {
		if value.Key == key {
			return value.Value, true
		}
	}
	return 0, false
}

// localizerFor returns a real drilldown localizer, or an explicit fallback.
func localizerFor(r Report) string {
	if len(r.Localizers) > 0 {
		return r.Localizers[0]
	}
	return fmt.Sprintf("file=%s|operator=none|airBefore=none|airAfter=none",
		r.ManifestSHA256)
}

// sortedIntKeys returns the keys of one integer map in ascending order.
func sortedIntKeys(in map[string]int) []string {
	keys := make([]string, 0, len(in))
	for key := range in {
		keys = append(keys, key)
	}
	slices.Sort(keys)
	return keys
}

// sortedFloatKeys returns the keys of one float map in ascending order.
func sortedFloatKeys(in map[string]float64) []string {
	keys := make([]string, 0, len(in))
	for key := range in {
		keys = append(keys, key)
	}
	slices.Sort(keys)
	return keys
}
