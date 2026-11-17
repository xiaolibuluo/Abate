/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package cli

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"

	"abate/ces/parse"
	"abate/ces/schema"
	"abate/kernel/scheduler"
	"abate/kernel/verify"
)

// runFixture executes a CES fixture with extra flags through the CLI.
func runFixture(t *testing.T, name string, extra ...string) (code int, stdoutText string, stderrText string) {
	t.Helper()
	args := append([]string{"run", filepath.Join("testdata", name)}, extra...)
	return invoke(args...)
}

// loadFixture parses and normalizes a CES fixture for direct unit tests.
func loadFixture(t *testing.T, name string) *schema.Experiment {
	t.Helper()
	data, err := os.ReadFile(filepath.Join("testdata", name)) //nolint:gosec // the fixture path is fixed by this test.
	if err != nil {
		t.Fatalf("read fixture: %v", err)
	}
	parsed, err := parse.Parse(data, parse.FormatAuto)
	if err != nil {
		t.Fatalf("parse fixture: %v", err)
	}
	normalized, err := parse.Normalize(parsed)
	if err != nil {
		t.Fatalf("normalize fixture: %v", err)
	}
	return normalized
}

// TestRunEndToEndCases covers legal, invalid, dry-run, strict, seed, format,
// out, help, missing-argument and internal-error paths of the run command.
func TestRunEndToEndCases(t *testing.T) {
	t.Parallel()
	badOut := filepath.Join(t.TempDir(), "absent-dir", "out.json")
	for _, c := range []struct {
		name    string
		fixture string
		flags   []string
		want    int
	}{
		{"legal", "ok.yaml", nil, exitOK},
		{"legal-json", "ok.yaml", []string{"--format", "json"}, exitOK},
		{"legal-seed", "ok.yaml", []string{"--seed", "42"}, exitOK},
		{"legal-seed-equals", "ok.yaml", []string{"--seed=7"}, exitOK},
		{"dry-run", "ok.yaml", []string{"--dry-run"}, exitOK},
		{"dry-run-strict", "ok.yaml", []string{"--dry-run", "--strict"}, exitOK},
		{"strict-legal", "ok.yaml", []string{"--strict"}, exitOK},
		{"violated", "violated.yaml", nil, exitBusinessFailure},
		{"invalid-rule", "invalid.yaml", nil, exitUsageError},
		{"malformed", "malformed.yaml", nil, exitUsageError},
		{"missing-file", "absent.yaml", nil, exitUsageError},
		{"bad-seed", "ok.yaml", []string{"--seed", "abc"}, exitUsageError},
		{"bad-format", "ok.yaml", []string{"--format", "xml"}, exitUsageError},
		{"bad-flag", "ok.yaml", []string{"--bogus"}, exitUsageError},
		{"extra-positional", "ok.yaml", []string{"extra"}, exitUsageError},
		{"internal-write-failure", "ok.yaml", []string{"--out", badOut}, exitInternalError},
		{"help", "ok.yaml", []string{"--help"}, exitOK},
	} {
		t.Run(c.name, func(t *testing.T) {
			code, out, errs := runFixture(t, c.fixture, c.flags...)
			if code != c.want {
				t.Fatalf("exit = %d, want %d (stderr=%s)", code, c.want, errs)
			}
			if (c.want == exitOK || c.want == exitBusinessFailure) && out == "" {
				t.Fatal("run produced no report bytes")
			}
		})
	}
	code, _, _ := invoke("run")
	if code != exitUsageError {
		t.Fatalf("run without a CES file exit = %d, want %d", code, exitUsageError)
	}
}

// TestRunDeterminism100 asserts the SHA-256 of a JSON report is byte-stable
// across 100 runs of the same document and seed.
func TestRunDeterminism100(t *testing.T) {
	t.Parallel()
	first := ""
	for i := 0; i < 100; i++ {
		code, out, errs := invoke("run", "testdata/ok.yaml", "--seed", "12345", "--format", "json")
		if code != exitOK {
			t.Fatalf("run exit = %d on iteration %d (stderr=%s)", code, i, errs)
		}
		if i == 0 {
			first = shaOf(out)
			continue
		}
		if shaOf(out) != first {
			t.Fatalf("report SHA-256 drifted on iteration %d", i)
		}
	}
}

// TestRunConcurrentDeterminism100 proves the control plane holds no shared
// mutable state: 100 concurrent runs all agree byte for byte.
func TestRunConcurrentDeterminism100(t *testing.T) {
	t.Parallel()
	const n = 100
	digests := make([]string, n)
	codes := make([]int, n)
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(k int) {
			defer wg.Done()
			code, out, _ := invoke("run", "testdata/ok.yaml", "--seed", "99")
			codes[k] = code
			digests[k] = shaOf(out)
		}(i)
	}
	wg.Wait()
	for i := 0; i < n; i++ {
		if codes[i] != exitOK || digests[i] != digests[0] {
			t.Fatalf("concurrent run %d: code=%d digest=%s", i, codes[i], digests[i])
		}
	}
}

// TestRunVerifyDoubleRun performs two real runs, writes their reports, and
// verifies the double-run comparison with and without tampering.
func TestRunVerifyDoubleRun(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()
	a := filepath.Join(dir, "report-a.json")
	b := filepath.Join(dir, "report-b.json")
	code1, _, _ := invoke("run", "testdata/ok.yaml", "--format", "json", "--out", a)
	code2, _, _ := invoke("run", "testdata/ok.yaml", "--format", "json", "--out", b)
	if code1 != exitOK || code2 != exitOK {
		t.Fatalf("double run exits = %d, %d", code1, code2)
	}
	code, out, _ := invoke("verify", a, b)
	if code != exitOK || !strings.Contains(out, "equal: true") {
		t.Fatalf("identical reports: exit=%d out=%s", code, out)
	}
	if err := os.WriteFile(b, []byte("tampered"), 0o600); err != nil {
		t.Fatalf("tamper report: %v", err)
	}
	code, out, _ = invoke("verify", a, b)
	if code != exitBusinessFailure || !strings.Contains(out, "equal: false") {
		t.Fatalf("tampered reports: exit=%d out=%s", code, out)
	}
}

// TestRunReportContent asserts the JSON report carries real kernel evidence.
func TestRunReportContent(t *testing.T) {
	t.Parallel()
	code, out, _ := invoke("run", "testdata/ok.yaml", "--format", "json")
	if code != exitOK {
		t.Fatalf("run exit = %d", code)
	}
	var rep report
	if err := json.Unmarshal([]byte(out), &rep); err != nil {
		t.Fatalf("report is not JSON: %v", err)
	}
	if rep.Status != scheduler.StatusCompleted || rep.Events != 6 || rep.Observations != 4 {
		t.Fatalf("report lacks scheduler evidence: %+v", rep)
	}
	if rep.RunDigest == "" || rep.StreamDigest == "" || rep.End == 0 {
		t.Fatalf("report lacks deterministic digests: %+v", rep)
	}
	if len(rep.Verdicts) != 1 || rep.Verdicts[0].Outcome != string(verify.OutcomeHeld) {
		t.Fatalf("report lacks the verifier verdict: %+v", rep.Verdicts)
	}
	if len(rep.Violations) != 0 {
		t.Fatalf("legal document reported violations: %+v", rep.Violations)
	}
}

// TestDryRunAndOutFile assert dry-run stops before execution and --out bytes
// match stdout exactly.
func TestDryRunAndOutFile(t *testing.T) {
	t.Parallel()
	code, out, _ := invoke("run", "testdata/ok.yaml", "--dry-run", "--format", "json")
	if code != exitOK {
		t.Fatalf("dry-run exit = %d", code)
	}
	var rep report
	if err := json.Unmarshal([]byte(out), &rep); err != nil {
		t.Fatalf("dry-run report is not JSON: %v", err)
	}
	if rep.Status != statusValidated || rep.Events != 0 || rep.RunDigest != "" {
		t.Fatalf("dry-run executed an experiment: %+v", rep)
	}
	p := filepath.Join(t.TempDir(), "report.json")
	code, stdout, _ := invoke("run", "testdata/ok.yaml", "--format", "json", "--out", p)
	if code != exitOK {
		t.Fatalf("run with --out exit = %d", code)
	}
	data, err := os.ReadFile(p) //nolint:gosec // the path was created by this test.
	if err != nil || !bytes.Equal(data, []byte(stdout)) {
		t.Fatalf("--out bytes disagree with stdout: err=%v", err)
	}
}

// TestRunExitCodeFourStates locks the PLAN-07 §4 contract on the real run
// chain. Code 2 uses the existing unwritable --out path, not an injected stub.
func TestRunExitCodeFourStates(t *testing.T) {
	t.Parallel()
	badOut := filepath.Join(t.TempDir(), "absent-dir", "out.json")
	for _, c := range []struct {
		name string
		args []string
		want int
	}{
		{"success-held", []string{"run", "testdata/ok.yaml"}, exitOK},
		{"business-failure-violated", []string{"run", "testdata/violated.yaml"}, exitBusinessFailure},
		{"internal-error-out", []string{"run", "testdata/ok.yaml", "--out", badOut}, exitInternalError},
		{"usage-parse-failure", []string{"run", "testdata/malformed.yaml"}, exitUsageError},
		{"usage-validation-error", []string{"run", "testdata/invalid.yaml"}, exitUsageError},
		{"usage-strict-validation", []string{"run", "testdata/invalid.yaml", "--strict"}, exitUsageError},
		{"usage-invalid-arguments", []string{"run"}, exitUsageError},
	} {
		t.Run(c.name, func(t *testing.T) {
			code, _, errs := invoke(c.args...)
			if code != c.want {
				t.Fatalf("exit = %d, want %d (stderr=%s)", code, c.want, errs)
			}
		})
	}
	code, out, _ := invoke("run", "testdata/invalid.yaml", "--format", "json")
	if code != exitUsageError || !strings.Contains(out, `"exitCode": 3`) ||
		!strings.Contains(out, "CES-VALID-ROLLBACK-001") {
		t.Fatalf("validation report = exit %d out %s", code, out)
	}
	code, out, _ = invoke("run", "testdata/violated.yaml", "--format", "json")
	if code != exitBusinessFailure || !strings.Contains(out, `"exitCode": 1`) ||
		!strings.Contains(out, `"outcome": "VIOLATED"`) {
		t.Fatalf("experiment report = exit %d out %s", code, out)
	}
}

// TestRunInternalErrorHasNoMisleadingReport proves a failed --out write exits
// 2 without emitting a report whose exitCode still claims success.
func TestRunInternalErrorHasNoMisleadingReport(t *testing.T) {
	t.Parallel()
	badOut := filepath.Join(t.TempDir(), "absent-dir", "out.json")
	code, out, errs := invoke("run", "testdata/ok.yaml", "--format", "json", "--out", badOut)
	if code != exitInternalError {
		t.Fatalf("exit = %d, want %d", code, exitInternalError)
	}
	if out != "" {
		t.Fatalf("internal error emitted misleading stdout bytes: %s", out)
	}
	if !strings.Contains(errs, "cannot write report") {
		t.Fatalf("internal error produced no diagnostic: %s", errs)
	}
}
