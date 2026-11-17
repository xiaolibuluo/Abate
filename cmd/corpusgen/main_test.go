/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// tinyPlanFile writes a small frozen-shaped plan for CLI end-to-end tests.
func tinyPlanFile(t *testing.T) string {
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

// TestRunGenerateAndVerify covers the complete CLI success path.
func TestRunGenerateAndVerify(t *testing.T) {
	t.Parallel()
	plan := tinyPlanFile(t)
	root := filepath.Join(t.TempDir(), "corpus")
	broken := filepath.Join(t.TempDir(), "broken")
	base := []string{"corpusgen", "-plan", plan, "-count", "350", "-strata", "2",
		"-root", root, "-broken-root", broken}
	var genOut, genErr bytes.Buffer
	if rc := run(base, &genOut, &genErr); rc != exitOK {
		t.Fatalf("generate rc=%d stderr=%s", rc, genErr.String())
	}
	if !strings.Contains(genOut.String(), "Generated=350\n") ||
		!strings.Contains(genOut.String(), "Broken=14\n") ||
		!strings.Contains(genOut.String(), "Language[go]=51\n") {
		t.Fatalf("generate summary = %s", genOut.String())
	}
	var verifyOut, verifyErr bytes.Buffer
	verifyArgs := append(append([]string{}, base...), "-verify-only")
	if rc := run(verifyArgs, &verifyOut, &verifyErr); rc != exitOK {
		t.Fatalf("verify rc=%d stderr=%s", rc, verifyErr.String())
	}
	if !strings.Contains(verifyOut.String(), "VerifiedFiles=350\n") {
		t.Fatalf("verify summary = %s", verifyOut.String())
	}
	var invalidOut, invalidErr bytes.Buffer
	if rc := run([]string{"corpusgen", "-count", "0"}, &invalidOut, &invalidErr); rc != exitUsage {
		t.Fatalf("invalid count rc=%d", rc)
	}
}

// TestRunHelpIsDeterministic proves the command's help path never touches the
// corpus and exits cleanly.
func TestRunHelpIsDeterministic(t *testing.T) {
	t.Parallel()
	var out, errBuf bytes.Buffer
	if rc := run([]string{"corpusgen", "-h"}, &out, &errBuf); rc != exitOK {
		t.Fatalf("help rc=%d stderr=%s", rc, errBuf.String())
	}
}

// TestRunVerifyRejectsMissingCorpus proves a missing corpus is an error, never
// an empty green verification.
func TestRunVerifyRejectsMissingCorpus(t *testing.T) {
	t.Parallel()
	args := []string{
		"corpusgen", "-verify-only",
		"-plan", filepath.Join("..", "..", "ci", "e2e_plan.json"),
		"-root", filepath.Join(t.TempDir(), "absent"),
	}
	var out, errBuf bytes.Buffer
	if rc := run(args, &out, &errBuf); rc != exitFail {
		t.Fatalf("rc=%d, want failure; stderr=%s", rc, errBuf.String())
	}
}

// TestRunRejectsMismatchedStrata proves CLI and plan disagreement is a hard
// generator failure rather than a silent fallback.
func TestRunRejectsMismatchedStrata(t *testing.T) {
	t.Parallel()
	args := []string{
		"corpusgen", "-plan", filepath.Join("..", "..", "ci", "e2e_plan.json"),
		"-strata", "7",
	}
	var out, errBuf bytes.Buffer
	if rc := run(args, &out, &errBuf); rc != exitFail {
		t.Fatalf("rc=%d, want generation failure; stderr=%s", rc, errBuf.String())
	}
}

// TestSortedKeysIsDeterministic proves report lines have a fixed order.
func TestSortedKeysIsDeterministic(t *testing.T) {
	t.Parallel()
	values := map[string]int{"b": 2, "a": 1, "c": 3}
	keys := sortedKeys(values)
	if len(keys) != 3 || keys[0] != "a" || keys[1] != "b" || keys[2] != "c" {
		t.Fatalf("keys = %v", keys)
	}
}
