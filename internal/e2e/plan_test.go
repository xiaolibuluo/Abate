/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package e2e

import (
	"errors"
	"os"
	"path/filepath"
	"slices"
	"testing"

	"abate/internal/corpus"
	"abate/internal/e2e/metrics"
	"abate/lang/air"
)

// validPlan is the smallest configuration accepted by Validate.
func validPlan() Plan {
	return Plan{
		Seed: 7, Languages: slices.Clone(languageOrder),
		FilesPerLanguage: 1, StrataCount: 1,
		MinFileBytes: 1, MaxFileBytes: 4096,
		ParseTimeoutMillis: 1, MemLimitBytes: 4096,
		MaxRuns: 1, BrokenPerLanguage: 1,
		Operators: []string{"cmp"},
	}
}

// writePlanFile stores raw plan bytes in a fresh directory.
func writePlanFile(t *testing.T, raw string) string {
	t.Helper()
	dir := t.TempDir()
	path := filepath.Join(dir, "plan.json")
	if err := os.WriteFile(path, []byte(raw), FilePerm); err != nil {
		t.Fatalf("write plan: %v", err)
	}
	return path
}

// TestLoadPlanDecodesRealAndMinimalPlans proves the digest, derived paths and
// default filling of both the checked-in plan and a zeroed plan document.
func TestLoadPlanDecodesRealAndMinimalPlans(t *testing.T) {
	t.Parallel()
	loadedPlan, err := LoadPlan(filepath.Join("..", "..", "ci", "e2e_plan.json"))
	if err != nil {
		t.Fatalf("real plan rejected: %v", err)
	}
	if len(loadedPlan.PlanSHA256) != 64 || loadedPlan.Seed == 0 {
		t.Fatalf("real plan digest or seed missing: %+v", loadedPlan)
	}
	if loadedPlan.ThresholdsPath != filepath.Join("..", "..", "ci", ThresholdsName) {
		t.Fatalf("derived thresholds path = %q", loadedPlan.ThresholdsPath)
	}
	minimal, err := LoadPlan(writePlanFile(t, `{"seed":9}`))
	if err != nil {
		t.Fatalf("minimal plan rejected: %v", err)
	}
	if minimal.FilesPerLanguage != defaultFilesPerLanguage ||
		minimal.StrataCount != defaultStrataCount ||
		minimal.MinFileBytes != defaultMinFileBytes ||
		minimal.MaxFileBytes != defaultMaxFileBytes ||
		minimal.ParseTimeoutMillis != defaultParseTimeoutMillis ||
		minimal.MaxRuns != defaultMaxRuns ||
		minimal.BrokenPerLanguage != defaultBrokenPerLanguage {
		t.Fatalf("defaults not filled: %+v", minimal)
	}
	if !slices.Equal(minimal.Languages, languageOrder) ||
		!slices.Equal(minimal.Operators, defaultOperators) {
		t.Fatalf("canonical lists not filled: %+v", minimal)
	}
	if minimal.PlanSHA256 == loadedPlan.PlanSHA256 {
		t.Fatal("different plan bytes produced the same digest")
	}
	if err = minimal.Validate(); err != nil {
		t.Fatalf("normalized plan invalid: %v", err)
	}
}

// TestLoadPlanRejectsBrokenDocuments covers decode and contract failures.
func TestLoadPlanRejectsBrokenDocuments(t *testing.T) {
	t.Parallel()
	missing := filepath.Join(t.TempDir(), "absent.json")
	if _, err := LoadPlan(missing); err == nil {
		t.Fatal("missing plan file accepted")
	}
	bad := []string{
		`{`,
		`{"seed":0}`,
		`{"seed":1,"languages":["go"]}`,
		`{"seed":1,"operators":["bogus"]}`,
	}
	for _, raw := range bad {
		if _, err := LoadPlan(writePlanFile(t, raw)); err == nil {
			t.Fatalf("broken plan accepted: %s", raw)
		}
	}
}

// TestValidateRejectsEveryContractViolation covers all budget, language and
// operator branches of the monotone plan validator.
func TestValidateRejectsEveryContractViolation(t *testing.T) {
	t.Parallel()
	variants := []Plan{}
	mutate := func(change func(*Plan)) {
		candidate := validPlan()
		change(&candidate)
		variants = append(variants, candidate)
	}
	mutate(func(p *Plan) { p.Seed = 0 })
	mutate(func(p *Plan) { p.FilesPerLanguage = 0 })
	mutate(func(p *Plan) { p.StrataCount = 0 })
	mutate(func(p *Plan) { p.MaxFileBytes = p.MinFileBytes - 1 })
	mutate(func(p *Plan) { p.ParseTimeoutMillis = 0 })
	mutate(func(p *Plan) { p.MemLimitBytes = 0 })
	mutate(func(p *Plan) { p.MaxRuns = 0 })
	mutate(func(p *Plan) { p.BrokenPerLanguage = 0 })
	mutate(func(p *Plan) { p.Languages = []string{"go"} })
	mutate(func(p *Plan) { p.Operators = []string{} })
	mutate(func(p *Plan) { p.Operators = []string{"op", "cmp"} })
	mutate(func(p *Plan) { p.Operators = []string{"unknown"} })
	for i := range variants {
		if err := variants[i].Validate(); err == nil {
			t.Fatalf("variant %d accepted: %+v", i, variants[i])
		}
	}
	if err := validPlan().Validate(); err != nil {
		t.Fatalf("valid plan rejected: %v", err)
	}
}

// TestPlanProjections proves the sampler and resource-budget projections.
func TestPlanProjections(t *testing.T) {
	t.Parallel()
	p := validPlan()
	sp := p.SamplePlan()
	if sp.FilesPerLanguage != 1 || sp.StrataCount != 1 ||
		sp.MinFileBytes != 1 || sp.MaxFileBytes != 4096 {
		t.Fatalf("sample plan = %+v", sp)
	}
	b := p.Budget()
	if b.TimeoutMillis != 1 || b.MemLimitBytes != 4096 {
		t.Fatalf("budget = %+v", b)
	}
}

// TestFrontendTableCoverage proves all seven frontends register and unknown
// language tags are refused instead of being skipped.
func TestFrontendTableCoverage(t *testing.T) {
	t.Parallel()
	table, err := NewFrontendTable()
	if err != nil {
		t.Fatalf("frontend table: %v", err)
	}
	if !slices.Equal(table.Languages(), languageOrder) {
		t.Fatalf("languages = %v", table.Languages())
	}
	for _, lang := range languageOrder {
		frontend, _, err := table.For(lang)
		if err != nil || frontend == nil {
			t.Fatalf("frontend %s: %v", lang, err)
		}
	}
	if _, _, err = table.For("unknown"); !errors.Is(err, errNoFrontend) {
		t.Fatalf("unknown language error = %v", err)
	}
}

// TestDefaultCorpusRootMatchesRegistry proves the E2E runner, corpusctl and
// the loaded plan all resolve the same single-source corpus root, and that the
// deleted third_party tree is no longer referenced.
func TestDefaultCorpusRootMatchesRegistry(t *testing.T) {
	t.Parallel()
	if DefaultCorpusRoot != corpus.DefaultCorpusRoot {
		t.Fatalf("e2e root %q != corpus root %q",
			DefaultCorpusRoot, corpus.DefaultCorpusRoot)
	}
	if DefaultCorpusRoot != "testdata/corpus" {
		t.Fatalf("default corpus root = %q", DefaultCorpusRoot)
	}
	loaded, err := LoadPlan(filepath.Join("..", "..", "ci", "e2e_plan.json"))
	if err != nil {
		t.Fatalf("load plan: %v", err)
	}
	if loaded.CorpusRoot != corpus.DefaultCorpusRoot {
		t.Fatalf("plan corpus root = %q", loaded.CorpusRoot)
	}
}

// TestResourceProbeClassifiesRealOutcomes covers success, contract failure,
// empty-unit failure and the deterministic allocation-budget verdict.
func TestResourceProbeClassifiesRealOutcomes(t *testing.T) {
	t.Parallel()
	ok := NewResourceProbe(ProbeBudget{
		TimeoutMillis: 5000, MemLimitBytes: 1 << 40,
	}).Measure(func() (*air.Unit, []air.Diagnostic, error) {
		return &air.Unit{}, nil, nil
	})
	if ok.Kind != metrics.OutcomeOK || ok.Diagnostics != 0 {
		t.Fatalf("success outcome = %+v", ok)
	}
	failed := NewResourceProbe(ProbeBudget{
		TimeoutMillis: 5000, MemLimitBytes: 1 << 40,
	}).Measure(func() (*air.Unit, []air.Diagnostic, error) {
		return nil, nil, errors.New("parse rejected")
	})
	if failed.Kind != metrics.OutcomeContractError || failed.Errors != 1 {
		t.Fatalf("contract outcome = %+v", failed)
	}
	empty := NewResourceProbe(ProbeBudget{
		TimeoutMillis: 5000, MemLimitBytes: 1 << 40,
	}).Measure(func() (*air.Unit, []air.Diagnostic, error) {
		return nil, []air.Diagnostic{{}}, nil
	})
	if empty.Kind != metrics.OutcomeContractError || empty.Diagnostics != 1 {
		t.Fatalf("empty-unit outcome = %+v", empty)
	}
	over := NewResourceProbe(ProbeBudget{
		TimeoutMillis: 5000, MemLimitBytes: 1,
	}).Measure(func() (*air.Unit, []air.Diagnostic, error) {
		data := make([]byte, 4096)
		data[0] = 1
		return &air.Unit{}, nil, nil
	})
	if over.Kind != metrics.OutcomeMemExceeded ||
		over.AllocDeltaBytes <= 1 {
		t.Fatalf("memory outcome = %+v", over)
	}
}
