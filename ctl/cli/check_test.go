/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package cli

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// checkFixture writes one small Go file with a known naming violation.
func checkFixture(t *testing.T) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "sample.go")
	if err := os.WriteFile(path,
		[]byte("package p\n\nvar foo_bar int\n"), 0o600); err != nil {
		t.Fatalf("write fixture: %v", err)
	}
	return path
}

// TestCheckCommandCases covers legal, help, missing-argument, invalid-flag and
// failure cases of the check command, including the four CLI exit states.
func TestCheckCommandCases(t *testing.T) {
	t.Parallel()
	for _, c := range []struct {
		name string
		args []string
		want int
	}{
		{"legal-clean", []string{"check", ".", "--rules", "GO-EG-001"}, exitOK},
		{"legal-violated", []string{"check", ".", "--rules", "GO-VET-037"}, exitBusinessFailure},
		{"long-help", []string{"check", "--help"}, exitOK},
		{"short-help", []string{"check", "-h"}, exitOK},
		{"no-paths", []string{"check"}, exitUsageError},
		{"bad-format", []string{"check", ".", "--format", "bogus"}, exitUsageError},
		{"bad-severity", []string{"check", ".", "--severity", "bogus"}, exitUsageError},
		{"bad-flag", []string{"check", ".", "--bogus"}, exitUsageError},
		{"missing-path", []string{"check", "./no-such-dir-m3"}, exitUsageError},
		{"no-active-rules", []string{"check", ".", "--rules", "ZZZ"}, exitUsageError},
	} {
		t.Run(c.name, func(t *testing.T) {
			code, out, errs := invoke(c.args...)
			if code != c.want {
				t.Fatalf("exit = %d, want %d (stderr=%s)", code, c.want, errs)
			}
			if out == "" && errs == "" {
				t.Fatal("check produced no output on either stream")
			}
		})
	}
}

// TestCheckHelpContent proves the check help documents check-specific exit codes. The check command describes zero for clean, one for findings, two for internal errors and three for usage errors. It must not reuse the run command steady-state wording.
func TestCheckHelpContent(t *testing.T) {
	t.Parallel()
	code, out, errs := invoke("check", "--help")
	if code != exitOK || errs != "" {
		t.Fatalf("check help exit = %d (stderr=%s)", code, errs)
	}
	for _, want := range []string{
		"abate check", "--format", "--severity", "--rules", "--out", "--languages",
		"Exit codes:", "including warnings", "--severity error", "abate check ./ctl/cli",
	} {
		if !strings.Contains(out, want) {
			t.Fatalf("check help missing %q", want)
		}
	}
	if !strings.Contains(rootHelp, "check <paths...>") {
		t.Fatal("root help does not advertise the check command")
	}
}

// TestCheckSARIFStructure proves SARIF field order, version and fingerprints.
func TestCheckSARIFStructure(t *testing.T) {
	t.Parallel()
	path := checkFixture(t)
	code, out, errs := invoke("check", path, "--format", "sarif")
	if code != exitBusinessFailure {
		t.Fatalf("fixture check exit = %d, want %d (stderr=%s)",
			code, exitBusinessFailure, errs)
	}
	var document map[string]any
	if err := json.Unmarshal([]byte(out), &document); err != nil {
		t.Fatalf("SARIF is not valid JSON: %v", err)
	}
	if document["version"] != "2.1.0" {
		t.Fatalf("SARIF version = %v", document["version"])
	}
	if !strings.Contains(out, `"$schema"`) ||
		!strings.Contains(out, `"partialFingerprints"`) ||
		!strings.Contains(out, `"abateFinding/v1"`) {
		t.Fatal("SARIF lost a required deterministic field")
	}
	schemaAt := strings.Index(out, `"$schema"`)
	versionAt := strings.Index(out, `"version"`)
	runsAt := strings.Index(out, `"runs"`)
	if schemaAt < 0 || versionAt < 0 || runsAt < 0 ||
		schemaAt > versionAt || versionAt > runsAt {
		t.Fatal("SARIF top-level field order is not fixed")
	}
	if strings.Contains(out, "timestamp") {
		t.Fatal("SARIF must not embed a wall-clock timestamp")
	}
}

// TestCheckOutFile proves --out writes exactly the stdout report bytes.
func TestCheckOutFile(t *testing.T) {
	t.Parallel()
	path := checkFixture(t)
	target := filepath.Join(t.TempDir(), "report.json")
	code, out, errs := invoke("check", path, "--format", "json", "--out", target)
	if code != exitBusinessFailure {
		t.Fatalf("exit = %d (stderr=%s)", code, errs)
	}
	data, err := os.ReadFile(target) //nolint:gosec // target is a test-owned temporary file.
	if err != nil {
		t.Fatalf("read report: %v", err)
	}
	if string(data) != out {
		t.Fatal("--out bytes differ from stdout bytes")
	}
	var document map[string]any
	if err := json.Unmarshal(data, &document); err != nil {
		t.Fatalf("JSON report is invalid: %v", err)
	}
	if document["schema"] != "abate.check.v1" {
		t.Fatalf("JSON schema = %v", document["schema"])
	}
}

// TestCheckDeterminism100 proves 100 identical invocations yield one hash.
func TestCheckDeterminism100(t *testing.T) {
	t.Parallel()
	path := checkFixture(t)
	seen := make(map[string]struct{}, 1)
	for i := 0; i < 100; i++ {
		code, out, errs := invoke("check", path, "--format", "sarif",
			"--rules", "GO-EG-001")
		if code != exitBusinessFailure {
			t.Fatalf("run %d exit = %d (stderr=%s)", i, code, errs)
		}
		seen[shaOf(out)] = struct{}{}
	}
	if len(seen) != 1 {
		t.Fatalf("check output produced %d distinct hashes", len(seen))
	}
}
