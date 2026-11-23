/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/
package targetgen

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

// goSpec returns a minimal valid Go spec.
func goSpec() Spec {
	return Spec{Lang: LangGo, Module: "example.com/target"}
}

func TestValidateFillsDefaults(t *testing.T) {
	spec, err := goSpec().Validate()
	if err != nil {
		t.Fatalf("Validate: %v", err)
	}
	if spec.Package != "target" {
		t.Errorf("Package = %q, want target", spec.Package)
	}
	if spec.Funcs != 8 {
		t.Errorf("Funcs = %d, want 8", spec.Funcs)
	}
	if spec.Level != LevelMedium {
		t.Errorf("Level = %q, want %q", spec.Level, LevelMedium)
	}
	if spec.Mixed != 0 {
		t.Errorf("Mixed = %v, want 0", spec.Mixed)
	}
}

func TestValidateRejectsBadSpecs(t *testing.T) {
	cases := []struct {
		name string
		spec Spec
		want error
	}{
		{"unknown language", Spec{Lang: "brainfuck"}, ErrUnknownLang},
		{"go without module", Spec{Lang: LangGo}, ErrModulePath},
		{"bad package name", Spec{Lang: LangGo, Module: "m", Package: "9bad"}, ErrPackageName},
		{"funcs negative", Spec{Lang: LangPython, Funcs: -1}, ErrFuncCount},
		{"funcs too large", Spec{Lang: LangPython, Funcs: MaxFuncs + 1}, ErrFuncCount},
		{"mixed too large", Spec{Lang: LangPython, Mixed: 1.5}, ErrMixedRatio},
		{"mixed negative", Spec{Lang: LangPython, Mixed: -0.25}, ErrMixedRatio},
		{"unknown level", Spec{Lang: LangPython, Level: "nightmare"}, ErrUnknownLevel},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if _, err := tc.spec.Validate(); !errors.Is(err, tc.want) {
				t.Fatalf("Validate() error = %v, want %v", err, tc.want)
			}
		})
	}
}

func TestLanguagesCanonicalOrder(t *testing.T) {
	langs := Languages()
	want := []string{LangGo, LangPython, LangJS, LangC, LangCSharp, LangRust, LangKotlin}
	if len(langs) != len(want) {
		t.Fatalf("Languages() = %v, want %d entries", langs, len(want))
	}
	for i, lang := range want {
		if langs[i] != lang {
			t.Errorf("Languages()[%d] = %q, want %q", i, langs[i], lang)
		}
	}
}

// targetsEqual reports whether two generations agree on every observable.
func targetsEqual(a, b Target) bool {
	if a.Tokens != b.Tokens || a.PinnedStructures != b.PinnedStructures {
		return false
	}
	if len(a.Files) != len(b.Files) || len(a.Funcs) != len(b.Funcs) {
		return false
	}
	for name, data := range a.Files {
		other, ok := b.Files[name]
		if !ok || !bytes.Equal(data, other) {
			return false
		}
	}
	for i, fn := range a.Funcs {
		if fn != b.Funcs[i] {
			return false
		}
	}
	return true
}

func TestGenerateIsDeterministic(t *testing.T) {
	for _, lang := range Languages() {
		t.Run(lang, func(t *testing.T) {
			spec := Spec{Lang: lang, Module: "example.com/target", Funcs: 6, Mixed: 0.5, Complex: true}
			first, err := Generate(spec, 99)
			if err != nil {
				t.Fatalf("Generate: %v", err)
			}
			second, err := Generate(spec, 99)
			if err != nil {
				t.Fatalf("Generate: %v", err)
			}
			if !targetsEqual(first, second) {
				t.Fatal("same seed produced different output")
			}
		})
	}
}

func TestGenerateSeedChangesOutput(t *testing.T) {
	for _, lang := range Languages() {
		t.Run(lang, func(t *testing.T) {
			spec := Spec{Lang: lang, Module: "example.com/target"}
			a, err := Generate(spec, 1)
			if err != nil {
				t.Fatalf("Generate(seed=1): %v", err)
			}
			b, err := Generate(spec, 2)
			if err != nil {
				t.Fatalf("Generate(seed=2): %v", err)
			}
			if targetsEqual(a, b) {
				t.Fatal("different seeds produced identical output")
			}
		})
	}
}

func TestGenerateFunctionBookkeeping(t *testing.T) {
	target, err := Generate(Spec{Lang: LangJS, Funcs: 5, Mixed: 1}, 17)
	if err != nil {
		t.Fatalf("Generate: %v", err)
	}
	if len(target.Funcs) != 5 {
		t.Fatalf("len(Funcs) = %d, want 5", len(target.Funcs))
	}
	seen := map[string]bool{}
	for i, fn := range target.Funcs {
		if seen[fn.Name] {
			t.Errorf("duplicate function name %q", fn.Name)
		}
		seen[fn.Name] = true
		if fn.TestName != "test_"+fn.Name {
			t.Errorf("Funcs[%d].TestName = %q, want test_%s", i, fn.TestName, fn.Name)
		}
		if !bytes.Contains([]byte(fn.Structure), []byte("+")) {
			t.Errorf("Funcs[%d].Structure = %q, want mixed a+b form", i, fn.Structure)
		}
	}
}

func TestScriptLanguagesRenderSingleFile(t *testing.T) {
	files := map[string]string{
		LangPython: "target.py",
		LangJS:     "target.js",
		LangC:      "target.c",
		LangCSharp: "target.cs",
		LangRust:   "target.rs",
		LangKotlin: "target.kt",
	}
	markers := map[string]string{
		LangPython: "PYTARGET_OK",
		LangJS:     "JSTARGET_OK",
		LangC:      "CTARGET_OK",
		LangCSharp: "CSTARGET_OK",
		LangRust:   "RSTARGET_OK",
		LangKotlin: "KTTARGET_OK",
	}
	for lang, name := range files {
		t.Run(lang, func(t *testing.T) {
			target, err := Generate(Spec{Lang: lang}, 5)
			if err != nil {
				t.Fatalf("Generate: %v", err)
			}
			if len(target.Files) != 1 {
				t.Fatalf("len(Files) = %d, want 1", len(target.Files))
			}
			data, ok := target.Files[name]
			if !ok {
				keys := make([]string, 0, len(target.Files))
				for k := range target.Files {
					keys = append(keys, k)
				}
				t.Fatalf("missing file %q (have %v)", name, keys)
			}
			if !bytes.Contains(data, []byte(markers[lang])) {
				t.Errorf("%s does not contain marker %s", name, markers[lang])
			}
		})
	}
}

func TestGenerateGoProjectPassesGoTest(t *testing.T) {
	if _, err := exec.LookPath("go"); err != nil {
		t.Skip("go toolchain not in PATH")
	}
	target, err := Generate(Spec{
		Lang:    LangGo,
		Module:  "example.com/target",
		Funcs:   12,
		Mixed:   0.5,
		Complex: true,
	}, 7)
	if err != nil {
		t.Fatalf("Generate: %v", err)
	}
	for _, name := range []string{"target.go", "target_test.go", "go.mod"} {
		if _, ok := target.Files[name]; !ok {
			t.Fatalf("missing file %q", name)
		}
	}
	dir := t.TempDir()
	for name, data := range target.Files {
		if err := os.WriteFile(filepath.Join(dir, name), data, 0o644); err != nil {
			t.Fatalf("write %s: %v", name, err)
		}
	}
	cmd := exec.Command("go", "test", "./...")
	cmd.Dir = dir
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("generated project failed go test: %v\n%s", err, out)
	}
}

func TestPinnedStructuresScaleWithLevel(t *testing.T) {
	pin := func(level string) int {
		target, err := Generate(Spec{Lang: LangPython, Funcs: MaxFuncs, Level: level}, 3)
		if err != nil {
			t.Fatalf("Generate(%s): %v", level, err)
		}
		return target.PinnedStructures
	}
	easy, hard := pin(LevelEasy), pin(LevelHard)
	if easy <= hard {
		t.Fatalf("easy pinned %d, hard pinned %d; want easy > hard", easy, hard)
	}
}
