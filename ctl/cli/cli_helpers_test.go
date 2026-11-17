/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package cli

import (
	"bytes"
	"errors"
	"io"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"

	"abate/ces/validate"
)

// failingWriter always fails a write, used for internal-error branches.
type failingWriter struct{}

// Write returns a deterministic write failure.
func (failingWriter) Write(_ []byte) (int, error) {
	return 0, errors.New("write failed")
}

// TestBlockingSeverityMapping covers error, warning and strict promotion.
func TestBlockingSeverityMapping(t *testing.T) {
	t.Parallel()
	warning := []validate.Violation{{RuleID: "W", Severity: validate.SeverityWarning}}
	if blocking(warning) {
		t.Fatal("warning must not block a non-strict run")
	}
	if !blockingStrict(warning) {
		t.Fatal("warning must block a strict run")
	}
	errV := []validate.Violation{{RuleID: "E", Severity: validate.SeverityError}}
	if !blocking(errV) || !blockingStrict(errV) {
		t.Fatal("error-severity violations must always block")
	}
	if blockingStrict(nil) {
		t.Fatal("an empty violation list must not block")
	}
	other := []validate.Violation{{RuleID: "I", Severity: "info"}}
	if blockingStrict(other) {
		t.Fatal("unknown severity must not block")
	}
}

// TestFirstDifferenceBranches covers equal, differing and unequal-length data.
func TestFirstDifferenceBranches(t *testing.T) {
	t.Parallel()
	for _, c := range []struct {
		left  string
		right string
		want  int
		equal bool
	}{
		{"abc", "abc", -1, true},
		{"abc", "abd", 2, false},
		{"ab", "abc", 2, false},
		{"abc", "ab", 2, false},
		{"", "x", 0, false},
		{"", "", -1, true},
	} {
		rep := compareReports("l", "r", []byte(c.left), []byte(c.right))
		if rep.FirstDiff != c.want || rep.Equal != c.equal {
			t.Fatalf("compare(%q,%q) = firstDiff %d equal %v",
				c.left, c.right, rep.FirstDiff, rep.Equal)
		}
		if len(rep.LeftSHA) != 64 || rep.LeftBytes != len(c.left) {
			t.Fatalf("incomplete comparison record: %+v", rep)
		}
	}
}

// TestReorderArgsAndFormats covers flag permutation and format validation.
func TestReorderArgsAndFormats(t *testing.T) {
	t.Parallel()
	got := reorderArgs(
		[]string{"file.yaml", "--seed", "42", "--strict", "--format=json", "--out", "r.json"},
		[]string{"seed", "format", "out"})
	want := []string{"--seed", "42", "--strict", "--format=json", "--out", "r.json", "file.yaml"}
	if !slices.Equal(got, want) {
		t.Fatalf("reorderArgs = %v, want %v", got, want)
	}
	if !takesValue("--seed", []string{"seed"}) {
		t.Fatal("a value flag must consume the next argument")
	}
	if takesValue("--strict", []string{"seed"}) || takesValue("--seed=", []string{"seed"}) {
		t.Fatal("a boolean or inline-assignment flag must not consume a value")
	}
	if !validFormat(formatConsole) || !validFormat(formatJSON) || validFormat("xml") {
		t.Fatal("format validation disagrees with the supported formats")
	}
}

// TestWriteOutputBranches covers stdout failures, file writes and out failures.
func TestWriteOutputBranches(t *testing.T) {
	t.Parallel()
	if err := writeOutput([]byte("x"), "", &bytes.Buffer{}, io.Discard); err != nil {
		t.Fatalf("plain stdout write failed: %v", err)
	}
	if err := writeOutput([]byte("x"), "", failingWriter{}, io.Discard); err == nil {
		t.Fatal("a failing stdout must be reported")
	}
	p := tempFile(t, nil)
	if err := writeOutput([]byte("abc"), p, &bytes.Buffer{}, io.Discard); err != nil {
		t.Fatalf("file write failed: %v", err)
	}
	got, err := os.ReadFile(p) //nolint:gosec // the path was created by this test.
	if err != nil || string(got) != "abc" {
		t.Fatalf("file content = %q, err = %v", got, err)
	}
	bad := filepath.Join(t.TempDir(), "absent-dir", "out.json")
	var notWritten bytes.Buffer
	if err := writeOutput([]byte(`{"exitCode":0}`), bad, &notWritten, io.Discard); err == nil {
		t.Fatal("an unwritable output path must be reported")
	}
	if notWritten.Len() != 0 {
		t.Fatalf("a failed file write leaked misleading stdout bytes: %q", notWritten.String())
	}
}

// TestRenderVerifyFormats covers both report formats and the console fallback.
func TestRenderVerifyFormats(t *testing.T) {
	t.Parallel()
	rep := compareReports("a", "b", []byte("x"), []byte("y"))
	rep.ExitCode = exitBusinessFailure
	js := renderVerifyReport(rep, formatJSON)
	if !strings.Contains(string(js), `"equal": false`) ||
		!strings.Contains(string(js), `"firstDiff": 0`) {
		t.Fatalf("JSON comparison missing fields: %s", js)
	}
	console := renderVerifyReport(rep, formatConsole)
	if !strings.Contains(string(console), "equal: false") ||
		!strings.Contains(string(console), "exitCode: 1") {
		t.Fatalf("console comparison missing fields: %s", console)
	}
	if !bytes.Equal(renderVerifyReport(rep, "xml"), console) {
		t.Fatal("unsupported format must fall back to deterministic console output")
	}
}

// TestStrictWarningMapsToUsageError locks the PLAN-07 §4 rule that a promoted
// warning is a usage error (exit 3). The current 25 CES-VALID rules emit only
// error-severity violations, so this branch is verified on the blocking pure
// function; end-to-end strict mode shows no visible difference yet.
func TestStrictWarningMapsToUsageError(t *testing.T) {
	t.Parallel()
	for _, c := range []struct {
		name     string
		severity validate.Severity
		strict   bool
		want     bool
	}{
		{"warning-relaxed", validate.SeverityWarning, false, false},
		{"warning-strict", validate.SeverityWarning, true, true},
		{"error-relaxed", validate.SeverityError, false, true},
		{"error-strict", validate.SeverityError, true, true},
	} {
		t.Run(c.name, func(t *testing.T) {
			vs := []validate.Violation{{RuleID: "X", Severity: c.severity}}
			var got bool
			if c.strict {
				got = blockingStrict(vs)
			} else {
				got = blocking(vs)
			}
			if got != c.want {
				t.Fatalf("blocking = %v, want %v", got, c.want)
			}
			code := exitOK
			if got {
				code = exitUsageError
			}
			if c.want && code != exitUsageError {
				t.Fatalf("promoted violation exit = %d, want %d", code, exitUsageError)
			}
		})
	}
}

// TestRootHelpExitCodeSemantics locks the PLAN-07 §4 mapping in both help
// texts so the root help can never contradict the run help again.
func TestRootHelpExitCodeSemantics(t *testing.T) {
	t.Parallel()
	for _, want := range []string{
		"Exit codes (PLAN-07 §4):",
		"0  success (legal CES and steady state HELD).",
		"1  business failure (steady state VIOLATED during execution).",
		"2  internal error (scheduler or report output failure).",
		"3  usage error (invalid arguments, invalid CES document, or a CES-VALID",
		"error violation; --strict promotes warnings to this code too).",
	} {
		if !strings.Contains(rootHelp, want) {
			t.Fatalf("rootHelp missing %q", want)
		}
		if !strings.Contains(runHelp, want) {
			t.Fatalf("runHelp missing %q", want)
		}
	}
	if strings.Contains(rootHelp, "error-severity violation or steady state VIOLATED") {
		t.Fatal("rootHelp still maps validation violations to exit 1")
	}
	if !strings.Contains(rootHelp, "Examples:") || !strings.Contains(runHelp, "Examples:") {
		t.Fatal("help text lost its copyable examples")
	}
}
