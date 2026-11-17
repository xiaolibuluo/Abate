/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package report

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// thresholds builds a threshold document covering the report's real keys.
func thresholds() Thresholds {
	return Thresholds{
		SchemaVersion: 1, TightenOnly: true,
		ZeroTolerance: map[string]int{"timeoutCount": 0},
		Minima: map[string]float64{
			"drilldownCoveragePct": 0,
			"manifestCoveragePct":  0,
		},
		Maxima: map[string]float64{"timeoutRatePct": 100},
	}
}

// TestLoadThresholdsRejectsBrokenDocuments covers read and decode failures.
func TestLoadThresholdsRejectsBrokenDocuments(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()
	if _, err := LoadThresholds(filepath.Join(dir, "absent.json")); err == nil {
		t.Fatal("missing thresholds accepted")
	}
	path := filepath.Join(dir, "bad.json")
	if err := os.WriteFile(path, []byte("{oops"), 0o600); err != nil {
		t.Fatalf("write bad thresholds: %v", err)
	}
	if _, err := LoadThresholds(path); err == nil {
		t.Fatal("malformed thresholds accepted")
	}
	good := filepath.Join(dir, "good.json")
	if err := os.WriteFile(good, []byte(`{"tightenOnly":true}`), 0o600); err != nil {
		t.Fatalf("write good thresholds: %v", err)
	}
	decoded, err := LoadThresholds(good)
	if err != nil || !decoded.TightenOnly {
		t.Fatalf("good thresholds = %+v, %v", decoded, err)
	}
}

// TestEvaluatePassingVerdict proves a fully registered, tighten-only and
// zero-tolerance-clean report passes with no failures.
func TestEvaluatePassingVerdict(t *testing.T) {
	t.Parallel()
	rep := sampleReport()
	rep.Counters = []Counter{{Key: "timeoutCount", Value: 0}}
	cfg := thresholds()
	verdict := Evaluate(cfg, cfg, rep)
	if !verdict.Passed || len(verdict.Failures) != 0 ||
		verdict.String() != "" {
		t.Fatalf("passing verdict = %+v", verdict)
	}
}

// TestEvaluateRejectsEveryViolation enumerates each zero-tolerance and
// tighten-only rule and asserts the fixed Failure localisation format.
func TestEvaluateRejectsEveryViolation(t *testing.T) {
	t.Parallel()
	base := sampleReport()
	base.Counters = []Counter{{Key: "timeoutCount", Value: 0}}
	cfg := thresholds()
	cases := []struct {
		name    string
		current Thresholds
		base    Thresholds
		report  Report
	}{
		{"no_tighten", mutateThresholds(cfg, func(t *Thresholds) {
			t.TightenOnly = false
		}), cfg, base},
		{"zero_key_set", mutateThresholds(cfg, func(t *Thresholds) {
			t.ZeroTolerance = map[string]int{"other": 0}
		}), cfg, base},
		{"zero_threshold_nonzero", mutateThresholds(cfg,
			func(t *Thresholds) {
				t.ZeroTolerance = map[string]int{"timeoutCount": 1}
			}), cfg, base},
		{"missing_counter", cfg, cfg, countersReport(base, nil)},
		{"nonzero_counter", cfg, cfg,
			countersReport(base, []Counter{{Key: "timeoutCount", Value: 1}})},
		{"relaxed_minima", mutateThresholds(cfg, func(t *Thresholds) {
			t.Minima = map[string]float64{"drilldownCoveragePct": -1}
		}), cfg, base},
		{"unregistered_minima", mutateThresholds(cfg, func(t *Thresholds) {
			t.Minima = map[string]float64{"unknownKey": 0}
		}), cfg, base},
		{"minima_not_met", mutateThresholds(cfg, func(t *Thresholds) {
			t.Minima = map[string]float64{"drilldownCoveragePct": 101}
		}), cfg, base},
		{"relaxed_maxima", mutateThresholds(cfg, func(t *Thresholds) {
			t.Maxima = map[string]float64{"timeoutRatePct": 101}
		}), cfg, base},
		{"unregistered_maxima", mutateThresholds(cfg, func(t *Thresholds) {
			t.Maxima = map[string]float64{"unknownKey": 0}
		}), cfg, base},
		{"maxima_exceeded", mutateThresholds(cfg, func(t *Thresholds) {
			t.Maxima = map[string]float64{"timeoutRatePct": -1}
		}), cfg, base},
		{"missing_baseline", mutateThresholds(cfg, func(t *Thresholds) {
			t.Minima = map[string]float64{"unknownKey": 0}
			t.Maxima = map[string]float64{}
		}), thresholds(), base},
	}
	for _, tc := range cases {
		verdict := Evaluate(tc.current, tc.base, tc.report)
		if verdict.Passed || len(verdict.Failures) == 0 {
			t.Fatalf("%s unexpectedly passed: %+v", tc.name, verdict)
		}
		for _, failure := range verdict.Failures {
			if !strings.Contains(failure.String(), "Failure{Key:") ||
				!strings.Contains(failure.Localizer, "file=") {
				t.Fatalf("%s bad localisation: %s", tc.name,
					failure.String())
			}
		}
	}
}

// mutateThresholds returns an independently owned copy with one edit applied.
func mutateThresholds(in Thresholds, edit func(*Thresholds)) Thresholds {
	out := Thresholds{
		SchemaVersion: in.SchemaVersion, TightenOnly: in.TightenOnly,
		ZeroTolerance: map[string]int{},
		Minima:        map[string]float64{},
		Maxima:        map[string]float64{},
	}
	for key, value := range in.ZeroTolerance {
		out.ZeroTolerance[key] = value
	}
	for key, value := range in.Minima {
		out.Minima[key] = value
	}
	for key, value := range in.Maxima {
		out.Maxima[key] = value
	}
	edit(&out)
	return out
}

// countersReport returns a report carrying the supplied zero-tolerance rows.
func countersReport(in Report, counters []Counter) Report {
	out := in
	out.Counters = counters
	return out
}

// TestFailureAndVerdictStrings proves the fixed five-field failure rendering
// and the deterministic multi-line verdict rendering.
func TestFailureAndVerdictStrings(t *testing.T) {
	t.Parallel()
	failure := Failure{Key: "timeoutCount", Rule: "zero-tolerance-nonzero",
		Localizer: "file=aa|operator=cmp|airBefore=x|airAfter=y",
		Actual:    3, Expected: 0}
	if !strings.Contains(failure.String(),
		"Failure{Key:timeoutCount Rule:zero-tolerance-nonzero") ||
		!strings.Contains(failure.String(), "Actual:3.000000") ||
		!strings.Contains(failure.String(), "Expected:0.000000") {
		t.Fatalf("failure string = %s", failure.String())
	}
	verdict := Verdict{Passed: false, Failures: []Failure{failure, failure}}
	if len(strings.Split(verdict.String(), "\n")) != 2 {
		t.Fatalf("verdict string = %s", verdict.String())
	}
}
