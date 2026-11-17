/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package corpusgen

import (
	"os"
	"path/filepath"
	"testing"
)

// tinyPlan writes a structurally valid plan with a small byte window so that
// generation tests finish in milliseconds while exercising the real code.
func tinyPlan(t *testing.T) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "plan.json")
	raw := `{"seed":20261114,"languages":["go","py","js","c","cs","rs","kt"],` +
		`"filesPerLanguage":2,"minFileBytes":128,"maxFileBytes":4096,` +
		`"strataCount":2,"parseTimeoutMillis":5000,"memLimitBytes":268435456,` +
		`"maxRuns":200000,"brokenPerLanguage":2,` +
		`"operators":["cmp","op","digit","call"],"reportDir":".artifacts/e2e"}`
	if err := os.WriteFile(path, []byte(raw), 0o600); err != nil {
		t.Fatalf("write plan: %v", err)
	}
	return path
}

// tinyOptions returns the generator options used by all generation tests.
func tinyOptions(t *testing.T) (GenOptions, PlanRef, SizeBands) {
	t.Helper()
	opts := GenOptions{
		PlanPath: tinyPlan(t), Root: filepath.Join(t.TempDir(), "corpus"),
		BrokenRoot: filepath.Join(t.TempDir(), "broken"), Count: 350, Strata: 2,
	}
	ref, err := LoadPlanRef(opts.PlanPath)
	if err != nil {
		t.Fatalf("load plan: %v", err)
	}
	bands, err := NewSizeBands(ref.MinFileBytes, ref.MaxFileBytes, ref.StrataCount)
	if err != nil {
		t.Fatalf("bands: %v", err)
	}
	return opts, ref, bands
}

func TestSizeBandsBoundariesAndQuotas(t *testing.T) {
	t.Parallel()
	bands, err := NewSizeBands(128, 184320, 8)
	if err != nil {
		t.Fatalf("bands: %v", err)
	}
	if bands.Step() != 23024 || bands.Low(0) != 128 || bands.High(0) != 23151 {
		t.Fatalf("band 0 = %+v", bands)
	}
	if bands.High(7) != 184320 || bands.StratumOf(184320) != 7 {
		t.Fatal("final band clamp failed")
	}
	if bands.StratumOf(0) != 0 || bands.StratumOf(23152) != 1 {
		t.Fatal("stratum classification failed")
	}
	if bands.Quota(2932, 0) != 2652 || bands.Quota(2932, 7) != 40 {
		t.Fatal("quota mismatch")
	}
	if bands.Target(0, 0) < 128 || bands.Target(7, 1) > 184320 {
		t.Fatal("target outside window")
	}
	for _, bad := range [][3]int{{0, 8, 1}, {8, 7, 8}, {7, 8, 0}} {
		if _, err = NewSizeBands(bad[0], bad[1], bad[2]); err == nil {
			t.Fatalf("accepted invalid bands %v", bad)
		}
	}
}

func TestNamesAreUniqueAndCleanFirstIsHashSmallest(t *testing.T) {
	t.Parallel()
	bands := SizeBands{MinBytes: 128, MaxBytes: 4096, Strata: 2}
	names := bands.Names("go", "gotxt", 0, 60)
	if len(names) != 60 || names[0] != "go/s0/go_s0_00000.gotxt" {
		t.Fatalf("names = %v", names[:2])
	}
	seen := map[string]bool{}
	for _, name := range names {
		if seen[name] {
			t.Fatalf("duplicate name %s", name)
		}
		seen[name] = true
	}
	if other := bands.Names("go", "gotxt", 1, 3); len(other) != 3 {
		t.Fatalf("other stratum names = %v", other)
	}
	// Design §2.2: the first six languages receive base+1 and the last
	// language absorbs the remainder, so the frozen 20520 splits as
	// 6x2932 + kt 2928. With count=287 (base=41) that is 6x42 + kt 35.
	quotas := languageQuotas(287, []string{"go", "py", "js", "c", "cs", "rs", "kt"})
	sum := 0
	for _, q := range quotas {
		sum += q
	}
	if quotas["go"] != 42 || quotas["kt"] != 35 || sum != 287 {
		t.Fatalf("quotas = %v sum=%d", quotas, sum)
	}
}

func TestLoadPlanRefReadsAndRejects(t *testing.T) {
	t.Parallel()
	ref, err := LoadPlanRef(filepath.Join("..", "..", "ci", "e2e_plan.json"))
	if err != nil || ref.Seed != 20261114 || len(ref.PlanSHA256) != 64 {
		t.Fatalf("ref = %+v err = %v", ref, err)
	}
	broken := filepath.Join(t.TempDir(), "bad.json")
	if err = os.WriteFile(broken, []byte(`{`), 0o600); err != nil {
		t.Fatalf("write: %v", err)
	}
	if _, err = LoadPlanRef(broken); err == nil {
		t.Fatal("accepted broken plan")
	}
	if err = ref.validate(); err != nil {
		t.Fatalf("validated plan rejected: %v", err)
	}
	invalid := ref
	invalid.Seed = 0
	if err = invalid.validate(); err == nil {
		t.Fatal("accepted zero seed")
	}
	invalid = ref
	invalid.Languages = []string{"go"}
	if err = invalid.validate(); err == nil {
		t.Fatal("accepted wrong language order")
	}
	invalid = ref
	invalid.Operators = nil
	if err = invalid.validate(); err == nil {
		t.Fatal("accepted empty operator set")
	}
}
