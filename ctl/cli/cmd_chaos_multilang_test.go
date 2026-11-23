/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/
package cli

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"abate/lang/targetgen"
)

// chaosWorkspaceLang materializes one deterministic target project of any
// supported language under t's temporary directory, mirroring chaosWorkspace
// for the non-Go frontends. Non-Go targets carry no module manifest, so the
// Spec.Module field stays empty. The seed and traversal order match the Go
// workspace helper, so the generated files are byte-identical across runs.
func chaosWorkspaceLang(t *testing.T, lang string, funcs int) string {
	t.Helper()
	target, err := targetgen.Generate(targetgen.Spec{
		Lang: lang, Funcs: funcs,
	}, 11)
	if err != nil {
		t.Fatalf("generate %s target: %v", lang, err)
	}
	dir := filepath.Join(t.TempDir(), "ws")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	for name, data := range target.Files {
		if err := os.WriteFile(filepath.Join(dir, name), data, 0o644); err != nil {
			t.Fatalf("write %s: %v", name, err)
		}
	}
	return dir
}

// TestChaosMutationMatchesPureMutationAcrossLangs extends the Go equivalence
// proof of TestChaosMutationMatchesPureMutation to every non-Go frontend.
// The L6 chaos phase injects and rolls back every registered fault operator
// inside a fully virtual world and consumes a random stream decorrelated by
// chaosSeedSalt, so it can never perturb the L7 mutation report regardless
// of the judging backend. Each subtest asserts the chaos run and the pure
// mutation run over the same materialized workspace produce the identical
// report fingerprint and the identical per-point verdict sequence.
//
// Funcs is kept at 3 (half of the Go workspace) because the slow toolchains
// (rustc -Cinstrument-coverage, Roslyn+dotnet-coverage, kotlinc) pay a real
// compile per mutant; the equivalence claim is an architectural property of
// the two-phase seed isolation and does not depend on the structure-family
// coverage breadth.
func TestChaosMutationMatchesPureMutationAcrossLangs(t *testing.T) {
	langs := []string{
		targetgen.LangPython,
		targetgen.LangJS,
		targetgen.LangC,
		targetgen.LangRust,
		targetgen.LangCSharp,
		targetgen.LangKotlin,
	}
	for _, lang := range langs {
		lang := lang
		t.Run(lang, func(t *testing.T) {
			dir := chaosWorkspaceLang(t, lang, 3)
			summary, chaosReport, _, err := runChaosMutation(context.Background(),
				42, dir, lang, 2, false, defaultChaosMutationPorts())
			if err != nil {
				t.Fatalf("chaos run: %v", err)
			}
			if summary.Operators != 12 {
				t.Errorf("Operators = %d, want 12", summary.Operators)
			}
			pureReport, _, err := runMutateExperiment(42, dir, lang, 2, false)
			if err != nil {
				t.Fatalf("pure mutation run: %v", err)
			}
			if chaosReport.Fingerprint != pureReport.Fingerprint {
				t.Errorf("chaos fingerprint %x != pure fingerprint %x",
					chaosReport.Fingerprint, pureReport.Fingerprint)
			}
			if len(chaosReport.Verdicts) != len(pureReport.Verdicts) {
				t.Fatalf("verdict count %d != %d",
					len(chaosReport.Verdicts), len(pureReport.Verdicts))
			}
			for i, v := range chaosReport.Verdicts {
				if v != pureReport.Verdicts[i] {
					t.Errorf("verdict[%d] = %s (chaos) vs %s (pure)",
						i, v, pureReport.Verdicts[i])
				}
			}
		})
	}
}
