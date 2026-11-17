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

// TestLoadPlanRefReadsFrozenPlan proves the generator binds the real plan's
// seed, byte window, strata and operator set without rewriting the file.
func TestLoadPlanRefReadsFrozenPlan(t *testing.T) {
	t.Parallel()
	ref, err := LoadPlanRef(filepath.Join("..", "..", "ci", "e2e_plan.json"))
	if err != nil {
		t.Fatalf("load frozen plan: %v", err)
	}
	if ref.Seed != 20261114 {
		t.Fatalf("seed = %d", ref.Seed)
	}
	if ref.FilesPerLanguage != 220 || ref.StrataCount != 8 {
		t.Fatalf("sampling contract = %+v", ref)
	}
	if ref.MinFileBytes != 128 || ref.MaxFileBytes != 184320 {
		t.Fatalf("byte window = [%d,%d]", ref.MinFileBytes, ref.MaxFileBytes)
	}
	if len(ref.Operators) != 8 || len(ref.Languages) != 7 {
		t.Fatalf("operator/language count = %d/%d", len(ref.Operators), len(ref.Languages))
	}
	if len(ref.PlanSHA256) != 64 {
		t.Fatalf("plan digest = %q", ref.PlanSHA256)
	}
}

// TestLoadPlanRefRejectsInvalidDocuments proves malformed and non-canonical
// plans are errors rather than silently skipped inputs.
func TestLoadPlanRefRejectsInvalidDocuments(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()
	bad := map[string]string{
		"broken.json":    `{`,
		"zero-seed.json": `{"seed":0}`,
		"one-lang.json":  `{"seed":1,"languages":["go"],"filesPerLanguage":1,"minFileBytes":1,"maxFileBytes":9,"strataCount":1,"brokenPerLanguage":1,"operators":["cmp"]}`,
	}
	for name, raw := range bad {
		path := filepath.Join(dir, name)
		if err := os.WriteFile(path, []byte(raw), 0o600); err != nil {
			t.Fatalf("write %s: %v", name, err)
		}
		if _, err := LoadPlanRef(path); err == nil {
			t.Fatalf("accepted invalid plan %s", name)
		}
	}
	if _, err := LoadPlanRef(filepath.Join(dir, "absent.json")); err == nil {
		t.Fatal("accepted missing plan")
	}
}

// TestLabeledRNGIsReproducible proves equal labels and seeds yield equal
// streams and distinct labels produce distinct scopes.
func TestLabeledRNGIsReproducible(t *testing.T) {
	t.Parallel()
	first := NewLabeledRNG(20261114, "go:s0")
	second := NewLabeledRNG(20261114, "go:s0")
	for i := 0; i < 16; i++ {
		if first.Intn(1000) != second.Intn(1000) {
			t.Fatalf("stream diverged at draw %d", i)
		}
	}
	other := NewLabeledRNG(20261114, "py:s0")
	if first.Label() == other.Label() {
		t.Fatal("distinct scopes share a label")
	}
}
