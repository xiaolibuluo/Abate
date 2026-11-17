/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Command corpusctl registers and verifies the read-only upstream corpora.
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"

	"abate/internal/corpus"
)

// Paths and exit codes are fixed so the tenth gate segment can recompute them.
const (
	manifestPath = "ci/corpus_manifest.json"
	digestPath   = "ci/corpus_manifest.sha256"
	// The corpusRoot constant points at the self-researched E2E parser-test fixtures, not a
	// product dependency. The path is the single source of truth exported by
	// internal/corpus so the runner and corpusctl cannot drift apart.
	corpusRoot   = corpus.DefaultCorpusRoot
	manifestPerm = 0o600
	exitOK       = 0
	exitFail     = 1
	exitUsage    = 2
)

func main() {
	os.Exit(run(os.Args, os.Stdout, os.Stderr))
}

func run(args []string, stdout, stderr io.Writer) int {
	if len(args) != 2 {
		emit(stderr, "usage: corpusctl register|verify (got %d arguments)\n", len(args))
		return exitUsage
	}
	switch args[1] {
	case "register":
		return runRegister(stdout, stderr)
	case "verify":
		return runVerify(stdout, stderr)
	default:
		emit(stderr, "%v: %s\n", corpus.ErrUnknownCommand, args[1])
		return exitUsage
	}
}

func runRegister(stdout, stderr io.Writer) int {
	fsys := os.DirFS(corpusRoot)
	m, err := corpus.Register(fsys, projectRefs())
	if err != nil {
		emit(stderr, "register failed: %v\n", err)
		return exitFail
	}
	data, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		emit(stderr, "register encode failed: %v\n", err)
		return exitFail
	}
	if err := os.WriteFile(manifestPath, data, manifestPerm); err != nil {
		emit(stderr, "register write failed: %v\n", err)
		return exitFail
	}
	if err := os.WriteFile(digestPath, []byte(m.Digest()+"\n"), manifestPerm); err != nil {
		emit(stderr, "register digest write failed: %v\n", err)
		return exitFail
	}
	emit(stdout, "Registered=%d ManifestSHA256=%s\n", len(m.Entries), m.Digest())
	return exitOK
}

func runVerify(stdout, stderr io.Writer) int {
	m, err := corpus.Load(manifestPath)
	if err != nil {
		emit(stderr, "verify load failed: %v\n", err)
		return exitFail
	}
	result, err := corpus.Verify(os.DirFS(corpusRoot), m)
	if err != nil {
		emit(stderr, "verify recompute failed: %v\n", err)
		return exitFail
	}
	printResult(stdout, result, m.Digest())
	if result.SHA256Mismatches != 0 || result.SPDXViolations != 0 {
		emit(stderr, "%v: sha256Mismatches=%d spdxViolations=%d\n",
			corpus.ErrVerifyFailed, result.SHA256Mismatches, result.SPDXViolations)
		return exitFail
	}
	return exitOK
}

// printResult emits a deterministic, fully sorted verification report.
func printResult(w io.Writer, result corpus.VerifyResult, digest string) {
	emit(w, "ManifestSHA256=%s\n", digest)
	emit(w, "Registered=%d\n", result.Registered)
	emit(w, "Recomputed=%d\n", result.Recomputed)
	emit(w, "SHA256Mismatches=%d\n", result.SHA256Mismatches)
	emit(w, "SPDXViolations=%d\n", result.SPDXViolations)
	emit(w, "BoundarySPDXCount=%d\n", result.BoundarySPDXCount)
	langs := make([]string, 0, len(result.ParseablePerLanguage))
	for lang := range result.ParseablePerLanguage {
		langs = append(langs, lang)
	}
	sort.Strings(langs)
	for _, lang := range langs {
		emit(w, "ParseablePerLanguage[%s]=%d\n", lang, result.ParseablePerLanguage[lang])
	}
	for _, line := range result.Mismatches {
		emit(w, "Mismatch: %s\n", line)
	}
	for _, line := range result.SPDXDetails {
		emit(w, "SPDXViolation: %s\n", line)
	}
}

// emit writes one line. The write error is explicitly discarded with a blank
// assignment, as prescribed by the cross-file error convention (§8.9).
func emit(w io.Writer, format string, args ...any) {
	_, _ = io.WriteString(w, fmt.Sprintf(format, args...))
}

// projectRefs lists the seven first-party, self-researched source trees. The
// SPDX evidence is the real CC0 dedication file beside the corpus; no upstream
// URL, third-party project name or fabricated commit hash is registered.
func projectRefs() []corpus.ProjectRef {
	refs := []corpus.ProjectRef{
		{Language: "go", RelDir: "go", ConstructTags: []string{"func", "generic", "concurrency", "interface"}},
		{Language: "py", RelDir: "py", ConstructTags: []string{"func", "class", "async", "match", "typing"}},
		{Language: "js", RelDir: "js", ConstructTags: []string{"class", "async", "generator", "module"}},
		{Language: "c", RelDir: "c", ConstructTags: []string{"struct", "macro", "pointer", "bitops"}},
		{Language: "cs", RelDir: "cs", ConstructTags: []string{"class", "generic", "linq", "async", "property"}},
		{Language: "rs", RelDir: "rs", ConstructTags: []string{"trait", "generic", "lifetime", "macro"}},
		{Language: "kt", RelDir: "kt", ConstructTags: []string{"class", "extension", "when", "null-safety"}},
	}
	for i := range refs {
		refs[i].Project = "Abate Self-Research Corpus"
		refs[i].UpstreamURL = ""
		refs[i].Commit = "self-researched-v1"
		refs[i].SPDX = "CC0-1.0"
		refs[i].LicenseRelPath = "SELF_RESEARCHED_CORPUS_DEDICATION.md"
		refs[i].Source = corpus.SourceSelfResearched
	}
	return refs
}
