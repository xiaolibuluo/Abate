/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package corpus

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"testing/fstest"

	"abate/internal/determinism"
)

// syntheticFS builds a small read-only tree with both source tracks and a
// licence evidence file, covering all seven registered extensions.
func syntheticFS() fstest.MapFS {
	fsys := fstest.MapFS{
		"LICENSE-MANIFEST.md": &fstest.MapFile{Data: []byte("MIT Apache-2.0 BSD-3-Clause\n")},
	}
	langs := []string{"go", "py", "js", "c", "cs", "rs", "kt"}
	for langIndex, lang := range langs {
		for n := 0; n < 6; n++ {
			body := bytes.Repeat([]byte(lang), 40+n*20+langIndex)
			name := fmt.Sprintf("_sources/_validation/_%s/_demo/f%d.%s", lang, n, lang)
			fsys[name] = &fstest.MapFile{Data: body}
		}
		body := bytes.Repeat([]byte(lang), 55+langIndex)
		name := fmt.Sprintf("_sources/in_tree_%s/a.%s", lang, lang)
		fsys[name] = &fstest.MapFile{Data: body}
	}
	return fsys
}

// syntheticRefs returns one validation ref and one in-tree ref per language.
func syntheticRefs() []ProjectRef {
	langs := []string{"go", "py", "js", "c", "cs", "rs", "kt"}
	refs := make([]ProjectRef, 0, len(langs)*2)
	for _, lang := range langs {
		refs = append(refs, ProjectRef{
			Language: lang, RelDir: fmt.Sprintf("_sources/_validation/_%s/_demo", lang),
			UpstreamURL: "https://example.invalid/" + lang, Commit: "v1.0.0",
			SPDX: "MIT", LicenseRelPath: "LICENSE-MANIFEST.md",
			Project: "demo-" + lang, Source: SourceValidation,
		})
		refs = append(refs, ProjectRef{
			Language: lang, RelDir: fmt.Sprintf("_sources/in_tree_%s", lang),
			UpstreamURL: "https://example.invalid/" + lang, Commit: "v1.0.0",
			SPDX: "Apache-2.0", LicenseRelPath: "LICENSE-MANIFEST.md",
			Project: "tree-" + lang, Source: SourceInTree,
		})
	}
	return refs
}

// TestRegisterRecordsRealBytesAndSorts asserts every entry carries the real
// digest of its original bytes, both tracks are labelled and paths ascend.
func TestRegisterRecordsRealBytesAndSorts(t *testing.T) {
	t.Parallel()
	fsys := syntheticFS()
	m, err := Register(fsys, syntheticRefs())
	if err != nil {
		t.Fatalf("Register failed: %v", err)
	}
	if len(m.Entries) == 0 {
		t.Fatal("Register produced no entries")
	}
	for i := 1; i < len(m.Entries); i++ {
		if m.Entries[i-1].RelPath >= m.Entries[i].RelPath {
			t.Fatalf("entries not ascending at %d: %s then %s", i,
				m.Entries[i-1].RelPath, m.Entries[i].RelPath)
		}
	}
	tracks := map[Source]int{}
	for _, entry := range m.Entries {
		data, readErr := fsys.ReadFile(entry.RelPath)
		if readErr != nil {
			t.Fatalf("missing registered file %s: %v", entry.RelPath, readErr)
		}
		if entry.SHA256 != hashBytes(data) {
			t.Fatalf("digest mismatch for %s", entry.RelPath)
		}
		if entry.Bytes != int64(len(data)) {
			t.Fatalf("byte count mismatch for %s", entry.RelPath)
		}
		evidence := entry.SPDXEvidence
		track := string(entry.Source)
		if evidence == "" || track == "" {
			t.Fatalf("missing provenance for %s", entry.RelPath)
		}
		tracks[entry.Source]++
	}
	if tracks[SourceValidation] == 0 || tracks[SourceInTree] == 0 {
		t.Fatalf("both source tracks must be present: %v", tracks)
	}
}

// TestVerifyRecomputesEveryEntry asserts full recomputation leaves zero
// mismatches and zero SPDX violations with all seven counters populated.
func TestVerifyRecomputesEveryEntry(t *testing.T) {
	t.Parallel()
	fsys := syntheticFS()
	m, err := Register(fsys, syntheticRefs())
	if err != nil {
		t.Fatalf("Register failed: %v", err)
	}
	result, err := Verify(fsys, m)
	if err != nil {
		t.Fatalf("Verify failed: %v", err)
	}
	if result.Registered != len(m.Entries) || result.Recomputed != len(m.Entries) {
		t.Fatalf("incomplete recompute: %+v", result)
	}
	if result.SHA256Mismatches != 0 || result.SPDXViolations != 0 {
		t.Fatalf("unexpected violations: %+v", result)
	}
	for _, lang := range languageOrder {
		if result.ParseablePerLanguage[lang] == 0 {
			t.Fatalf("language %s has no registered files", lang)
		}
	}
}

// TestVerifyDetectsDigestMismatch asserts a changed byte is caught and the
// failing entry is localised by relative path.
func TestVerifyDetectsDigestMismatch(t *testing.T) {
	t.Parallel()
	fsys := syntheticFS()
	m, err := Register(fsys, syntheticRefs())
	if err != nil {
		t.Fatalf("Register failed: %v", err)
	}
	target := m.Entries[0].RelPath
	fsys[target] = &fstest.MapFile{Data: []byte("changed bytes")}
	result, err := Verify(fsys, m)
	if err != nil {
		t.Fatalf("Verify failed: %v", err)
	}
	if result.SHA256Mismatches != 1 {
		t.Fatalf("expected one mismatch, got %d", result.SHA256Mismatches)
	}
	if len(result.Mismatches) != 1 || !bytes.Contains([]byte(result.Mismatches[0]), []byte(target)) {
		t.Fatalf("mismatch lacks localiser: %v", result.Mismatches)
	}
}

// TestClassifySPDXCategories pins permissive, boundary, copyleft and unknown
// classifications, including a dual-licence expression.
func TestClassifySPDXCategories(t *testing.T) {
	t.Parallel()
	cases := []struct {
		expr string
		want string
	}{
		{"MIT", spdxAllowlisted},
		{"Apache-2.0", spdxAllowlisted},
		{"BSD-3-Clause", spdxAllowlisted},
		{"MIT OR Apache-2.0", spdxAllowlisted},
		{"MIT / Apache-2.0", spdxAllowlisted},
		{"Apache-2.0 WITH LLVM-exception", spdxBoundary},
		{"Apache-2.0 + LLVM exception", spdxBoundary},
		{"GPL-3.0-only", spdxGPL},
		{"GNU AFFERO GENERAL PUBLIC LICENSE", spdxGPL},
		{"", spdxUNKNOWN},
		{"Proprietary", spdxUNKNOWN},
	}
	for _, tc := range cases {
		if got := ClassifySPDX(tc.expr); got != tc.want {
			t.Fatalf("ClassifySPDX(%q) = %q, want %q", tc.expr, got, tc.want)
		}
	}
}

// TestVerifyGPLAndUNKNOWNHardFail asserts copyleft and unknown declarations
// both count as violations rather than being silently dropped.
func TestVerifyGPLAndUNKNOWNHardFail(t *testing.T) {
	t.Parallel()
	fsys := syntheticFS()
	m, err := Register(fsys, syntheticRefs())
	if err != nil {
		t.Fatalf("Register failed: %v", err)
	}
	m.Entries[0].SPDX = "GPL-3.0-only"
	m.Entries[1].SPDX = "UNKNOWN"
	result, err := Verify(fsys, m)
	if err != nil {
		t.Fatalf("Verify failed: %v", err)
	}
	if result.SPDXViolations != 2 || len(result.SPDXDetails) != 2 {
		t.Fatalf("expected two SPDX violations, got %+v", result)
	}
}

// TestStratifiedSampleSameSeedDoubleRun asserts two draws from the same
// manifest and plan are byte-identical, and that a same-seed RNG is stable.
func TestStratifiedSampleSameSeedDoubleRun(t *testing.T) {
	t.Parallel()
	fsys := syntheticFS()
	m, err := Register(fsys, syntheticRefs())
	if err != nil {
		t.Fatalf("Register failed: %v", err)
	}
	plan := SamplePlan{FilesPerLanguage: 3, StrataCount: 2, MinFileBytes: 1, MaxFileBytes: 4096}
	left := StratifiedSample(m, plan)
	right := StratifiedSample(m, plan)
	leftJSON, err := json.Marshal(left)
	if err != nil {
		t.Fatalf("marshal left: %v", err)
	}
	rightJSON, err := json.Marshal(right)
	if err != nil {
		t.Fatalf("marshal right: %v", err)
	}
	if !bytes.Equal(leftJSON, rightJSON) {
		t.Fatal("stratified sample is not byte-identical across runs")
	}
	if len(left) == 0 {
		t.Fatal("stratified sample is empty")
	}
	leftRNG := determinism.NewRNG(defaultSeed, m.Digest(), "mutation")
	rightRNG := determinism.NewRNG(defaultSeed, m.Digest(), "mutation")
	for i := 0; i < 32; i++ {
		if leftRNG.Next() != rightRNG.Next() {
			t.Fatalf("same-seed RNG diverged at draw %d", i)
		}
	}
}

// TestLoadRoundTripRejectsUnsorted asserts the loader recomputes the file
// digest and refuses a manifest whose entries are out of order.
func TestLoadRoundTripRejectsUnsorted(t *testing.T) {
	t.Parallel()
	fsys := syntheticFS()
	m, err := Register(fsys, syntheticRefs())
	if err != nil {
		t.Fatalf("Register failed: %v", err)
	}
	data := encodeManifest(m)
	dir := t.TempDir()
	path := filepath.Join(dir, "corpus_manifest.json")
	if err := os.WriteFile(path, data, 0o600); err != nil {
		t.Fatalf("write manifest: %v", err)
	}
	loaded, err := Load(path)
	if err != nil {
		t.Fatalf("Load failed: %v", err)
	}
	if loaded.Digest() != hashBytes(data) {
		t.Fatal("Load digest differs from the raw file bytes")
	}
	if len(loaded.Entries) != len(m.Entries) {
		t.Fatal("Load lost entries")
	}
	swapped := Manifest{SchemaVersion: manifestSchemaVersion, Entries: []Entry{m.Entries[1], m.Entries[0]}}
	bad, err := json.Marshal(swapped)
	if err != nil {
		t.Fatalf("marshal unsorted: %v", err)
	}
	if err := os.WriteFile(path, bad, 0o600); err != nil {
		t.Fatalf("write unsorted: %v", err)
	}
	if _, err := Load(path); !errors.Is(err, errUnsortedManifest) {
		t.Fatalf("expected errUnsortedManifest, got %v", err)
	}
}
