/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package corpusgen

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestGenerateAndVerifySmallCorpus(t *testing.T) {
	t.Parallel()
	opts, ref, bands := tinyOptions(t)
	generated, err := Generate(opts)
	if err != nil {
		t.Fatalf("generate: %v", err)
	}
	if generated.Total != opts.Count || generated.Bytes < 1 || generated.Broken != 14 {
		t.Fatalf("generated = %+v", generated)
	}
	verified, err := Verify(opts)
	if err != nil {
		t.Fatalf("verify: %v", err)
	}
	if verified.Files != opts.Count || verified.DistinctSHA != opts.Count {
		t.Fatalf("verified = %+v", verified)
	}
	if len(verified.Dedication) != 64 {
		t.Fatalf("dedication digest = %q", verified.Dedication)
	}
	for _, language := range ref.Languages {
		if verified.PerLanguage[language] != generated.PerLanguage[language] {
			t.Fatalf("language %s mismatch", language)
		}
	}
	for stratum := 0; stratum < ref.StrataCount; stratum++ {
		key := "go:s" + itoa(stratum)
		if verified.PerStratum[key] != bands.Quota(generated.PerLanguage["go"], stratum) {
			t.Fatalf("%s = %d", key, verified.PerStratum[key])
		}
	}
	if _, err = os.Stat(filepath.Join(opts.Root, DedicationRelPath)); err != nil {
		t.Fatalf("dedication missing: %v", err)
	}
	for _, language := range ref.Languages {
		dir := filepath.Join(opts.BrokenRoot, language)
		entries, readErr := os.ReadDir(dir)
		if readErr != nil || len(entries) != 4 {
			t.Fatalf("broken %s entries = %d err = %v", language, len(entries), readErr)
		}
	}
}

func TestGenerationRejectsBadInputs(t *testing.T) {
	t.Parallel()
	opts, _, _ := tinyOptions(t)
	bad := opts
	bad.Strata = 3
	if _, err := Generate(bad); !errors.Is(err, errGenerate) {
		t.Fatalf("strata mismatch error = %v", err)
	}
	badCount := opts
	badCount.Count = 2
	if _, err := Generate(badCount); !errors.Is(err, errGenerate) {
		t.Fatalf("small count error = %v", err)
	}
	badPlan := opts
	badPlan.PlanPath = filepath.Join(t.TempDir(), "absent.json")
	if _, err := Generate(badPlan); err == nil {
		t.Fatal("missing plan accepted")
	}
	if err := assertEmptyRoot(opts.Root); err != nil {
		t.Fatalf("absent root rejected: %v", err)
	}
	if err := os.MkdirAll(opts.Root, 0o750); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	if err := assertEmptyRoot(opts.Root); err != nil {
		t.Fatalf("empty existing root rejected: %v", err)
	}
	if err := os.WriteFile(filepath.Join(opts.Root, "x.go"), []byte("x"), 0o600); err != nil {
		t.Fatalf("write: %v", err)
	}
	if err := assertEmptyRoot(opts.Root); !errors.Is(err, errGenerate) {
		t.Fatalf("non-empty root accepted: %v", err)
	}
	fileRoot := filepath.Join(t.TempDir(), "file")
	if err := os.WriteFile(fileRoot, []byte("x"), 0o600); err != nil {
		t.Fatalf("write file: %v", err)
	}
	if err := assertEmptyRoot(fileRoot); !errors.Is(err, errGenerate) {
		t.Fatalf("file root accepted: %v", err)
	}
}

func TestDamageSourceAndVerificationHelpers(t *testing.T) {
	t.Parallel()
	data := []byte("def anchor():\n    return 1\n\ndef step():\n    return 2\n")
	damaged, kind, err := damageSource("py", data)
	if err != nil || kind != "missing-colon" || string(damaged) == string(data) {
		t.Fatalf("python damage = %q %s %v", damaged, kind, err)
	}
	goData := []byte("package synth\n\nfunc a() {\n}\n")
	damaged, kind, err = damageSource("go", goData)
	if err != nil || kind != "unclosed-delimiter" || string(damaged) == string(goData) {
		t.Fatalf("go damage = %q %s %v", damaged, kind, err)
	}
	if _, _, err = damageSource("py", []byte("no def here")); err == nil {
		t.Fatal("python file without def accepted")
	}
	if _, _, err = damageSource("go", []byte("no brace")); err == nil {
		t.Fatal("go file without brace accepted")
	}
	ctx := verifyContext{ref: PlanRef{Languages: []string{"go", "py", "js", "c", "cs", "rs", "kt"},
		StrataCount: 2}, bands: SizeBands{MinBytes: 128, MaxBytes: 4096, Strata: 2}}
	if _, _, err = ctx.parseRelPath("go/s0/go_s0_00000.gotxt"); err != nil {
		t.Fatalf("parse relPath: %v", err)
	}
	if language, _, _ := ctx.parseRelPath("notes.md"); language != "" {
		t.Fatal("non-source file classified as source")
	}
	if _, _, err = ctx.parseRelPath("go/bad/x.gotxt"); err == nil {
		t.Fatal("bad stratum path accepted")
	}
	if _, _, err = ctx.parseRelPath("go/s9/x.gotxt"); err == nil {
		t.Fatal("out-of-range stratum accepted")
	}
}

func TestCheckCaptureAnchorRejectsInvalidFiles(t *testing.T) {
	t.Parallel()
	if err := checkCaptureAnchor([]byte("no anchor"), "go/s0/x.gotxt", ".gotxt"); err == nil {
		t.Fatal("missing anchor accepted")
	}
	late := []byte("1 + 2 == 3\nreturn 1 + 2 == 3\n")
	if err := checkCaptureAnchor(late, "go/s0/x.gotxt", ".gotxt"); err == nil {
		t.Fatal("early digit accepted")
	}
	marked := []byte("github.com/example\n1 + 2 == 3\n")
	if err := checkCaptureAnchor(marked, "go/s0/x.gotxt", ".gotxt"); err == nil {
		t.Fatal("third-party marker accepted")
	}
	cleanPath := "go/s0/go_s0_00000.gotxt"
	if err := checkCaptureAnchor([]byte("1 + 2 == 3\n"), cleanPath, ".gotxt"); err != nil {
		t.Fatalf("clean anchor rejected: %v", err)
	}
}

func TestVerifyRejectsMissingCorpusAndBadQuota(t *testing.T) {
	t.Parallel()
	opts, ref, _ := tinyOptions(t)
	missing := opts
	missing.Root = filepath.Join(t.TempDir(), "absent")
	if _, err := Verify(missing); err == nil {
		t.Fatal("missing corpus verified")
	}
	ctx := verifyContext{ref: ref, quotas: languageQuotas(opts.Count, ref.Languages),
		stats: VerifyStats{PerLanguage: map[string]int{}, PerStratum: map[string]int{}}}
	if err := ctx.checkQuotas(); err == nil {
		t.Fatal("empty corpus passed quota check")
	}
}
