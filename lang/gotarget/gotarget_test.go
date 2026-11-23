/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/
package gotarget

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestGenerateDeterministic proves the core contract: one seed plus one
// spec always yields byte-identical file bytes and one fingerprint.
func TestGenerateDeterministic(t *testing.T) {
	t.Parallel()
	spec := Spec{Module: "example.com/target", Funcs: 14, Level: LevelMedium}
	first, err := Generate(spec, 42)
	if err != nil {
		t.Fatalf("Generate: %v", err)
	}
	second, err := Generate(spec, 42)
	if err != nil {
		t.Fatalf("Generate: %v", err)
	}
	if first.Fingerprint() != second.Fingerprint() {
		t.Fatal("same seed produced different fingerprints")
	}
	for name, data := range first.Files {
		if string(second.Files[name]) != string(data) {
			t.Fatalf("file %s differs across two runs of one seed", name)
		}
	}
	if first.Seed != 42 || len(first.Funcs) != 14 {
		t.Fatalf("bookkeeping wrong: seed=%d funcs=%d", first.Seed, len(first.Funcs))
	}
}

// TestGenerateSeedsDiverge proves different seeds produce different targets,
// so the seed is a real variability knob and not a constant.
func TestGenerateSeedsDiverge(t *testing.T) {
	t.Parallel()
	spec := Spec{Module: "example.com/target", Funcs: 14, Level: LevelMedium}
	a, err := Generate(spec, 1)
	if err != nil {
		t.Fatalf("Generate seed 1: %v", err)
	}
	b, err := Generate(spec, 2)
	if err != nil {
		t.Fatalf("Generate seed 2: %v", err)
	}
	if a.Fingerprint() == b.Fingerprint() {
		t.Fatal("different seeds produced identical targets")
	}
}

// TestGenerateStructureCoverage proves every structure family appears when
// the function count covers at least one full cycle of seven.
func TestGenerateStructureCoverage(t *testing.T) {
	t.Parallel()
	target, err := Generate(Spec{Module: "example.com/target", Funcs: 7, Level: LevelEasy}, 7)
	if err != nil {
		t.Fatalf("Generate: %v", err)
	}
	seen := make(map[string]bool)
	for _, fn := range target.Funcs {
		seen[fn.Structure] = true
		if !strings.HasPrefix(fn.TestName, "Test") {
			t.Fatalf("test name %q lacks Test prefix", fn.TestName)
		}
	}
	for _, structure := range structureOrder {
		if !seen[structure] {
			t.Fatalf("structure %s missing from 7-function target", structure)
		}
	}
	source := string(target.Files["target.go"])
	test := string(target.Files["target_test.go"])
	for _, fn := range target.Funcs {
		if !strings.Contains(source, "func "+fn.Name+"(") {
			t.Fatalf("source lacks function %s", fn.Name)
		}
		if !strings.Contains(test, "func "+fn.TestName+"(") {
			t.Fatalf("test file lacks %s", fn.TestName)
		}
	}
	if !strings.Contains(string(target.Files["go.mod"]), "module example.com/target") {
		t.Fatal("go.mod lacks the module path")
	}
}

// TestLevelShapesAssertionDensity proves the difficulty level steers the
// generated tests: easy pins more edge cases than hard on the same shape.
func TestLevelShapesAssertionDensity(t *testing.T) {
	t.Parallel()
	spec := Spec{Module: "example.com/target", Funcs: 28}
	spec.Level = LevelEasy
	easy, err := Generate(spec, 9)
	if err != nil {
		t.Fatalf("Generate easy: %v", err)
	}
	spec.Level = LevelHard
	hard, err := Generate(spec, 9)
	if err != nil {
		t.Fatalf("Generate hard: %v", err)
	}
	density := func(target Target) int {
		return strings.Count(string(target.Files["target_test.go"]), "t.Errorf")
	}
	if density(easy) <= density(hard) {
		t.Fatalf("easy assertions %d not above hard %d", density(easy), density(hard))
	}
}

// TestSpecValidation proves the spec errors fire before any generation.
func TestSpecValidation(t *testing.T) {
	t.Parallel()
	if _, err := Generate(Spec{Module: "", Funcs: 4}, 1); !strings.Contains(err.Error(), "module path") {
		t.Fatalf("empty module error = %v", err)
	}
	if _, err := Generate(Spec{Module: "m", Funcs: MaxFuncs + 1}, 1); !strings.Contains(err.Error(), "function count") {
		t.Fatalf("overflow error = %v", err)
	}
	if _, err := Generate(Spec{Module: "m", Funcs: 4, Level: "impossible"}, 1); !strings.Contains(err.Error(), "difficulty") {
		t.Fatalf("bad level error = %v", err)
	}
	if _, err := Generate(Spec{Module: "m", Package: "9bad", Funcs: 4}, 1); !strings.Contains(err.Error(), "package name") {
		t.Fatalf("bad package error = %v", err)
	}
}

// TestWriteSelfCheck proves Write materializes a module whose real
// `go test` run is green, and refuses to overwrite an existing directory.
func TestWriteSelfCheck(t *testing.T) {
	t.Parallel()
	target, err := Generate(Spec{Module: "example.com/target", Funcs: 7, Level: LevelEasy}, 3)
	if err != nil {
		t.Fatalf("Generate: %v", err)
	}
	dir := filepath.Join(t.TempDir(), "target")
	if err := Write(dir, target, true); err != nil {
		t.Fatalf("Write with self-check: %v", err)
	}
	for name := range target.Files {
		if _, err := os.Stat(filepath.Join(dir, name)); err != nil {
			t.Fatalf("written file %s missing: %v", name, err)
		}
	}
	if err := Write(dir, target, false); err == nil || !strings.Contains(err.Error(), "already exists") {
		t.Fatalf("second Write error = %v, want ErrDirExists", err)
	}
}
