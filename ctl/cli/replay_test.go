/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package cli

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// recordArchive runs the CLI with --record and returns the archive path.
func recordArchive(t *testing.T, seed string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "run.arch")
	code, _, errs := invoke("run", "testdata/ok.yaml", "--seed", seed, "--record", path)
	if code != exitOK {
		t.Fatalf("run --record exit = %d (stderr=%s)", code, errs)
	}
	if _, err := os.Stat(path); err != nil {
		t.Fatalf("archive not written: %v", err)
	}
	return path
}

// TestReplayCommandCases locks the four-state exit-code contract of replay.
func TestReplayCommandCases(t *testing.T) {
	t.Parallel()
	good := recordArchive(t, "42")
	missing := filepath.Join(t.TempDir(), "absent.arch")
	corrupt := tempFile(t, []byte("not an archive"))
	badOut := filepath.Join(t.TempDir(), "absent-dir", "out.txt")
	for _, c := range []struct {
		name string
		args []string
		want int
	}{
		{"verify-one", []string{"replay", good}, exitOK},
		{"verify-two-equal", []string{"replay", good, good}, exitOK},
		{"json", []string{"replay", "--format", "json", good}, exitOK},
		{"format-equals", []string{"replay", "--format=json", good}, exitOK},
		{"flags-first", []string{"replay", "--format", "json", good, good}, exitOK},
		{"help", []string{"replay", "--help"}, exitOK},
		{"no-arg", []string{"replay"}, exitUsageError},
		{"too-many", []string{"replay", good, good, good}, exitUsageError},
		{"bad-flag", []string{"replay", "--bogus", good}, exitUsageError},
		{"bad-format", []string{"replay", "--format", "xml", good}, exitUsageError},
		{"missing-left", []string{"replay", missing}, exitUsageError},
		{"missing-right", []string{"replay", good, missing}, exitUsageError},
		{"corrupt-left", []string{"replay", corrupt}, exitInternalError},
		{"corrupt-right", []string{"replay", good, corrupt}, exitInternalError},
		{"internal-write-failure", []string{"replay", "--out", badOut, good}, exitInternalError},
	} {
		t.Run(c.name, func(t *testing.T) {
			code, _, errs := invoke(c.args...)
			if code != c.want {
				t.Fatalf("exit = %d, want %d (stderr=%s)", code, c.want, errs)
			}
		})
	}
}

// TestReplayDiffOutput proves a real difference prints the structured Diff.
func TestReplayDiffOutput(t *testing.T) {
	t.Parallel()
	a := recordArchive(t, "42")
	b := recordArchive(t, "43")
	code, out, errs := invoke("replay", a, b)
	if code != exitBusinessFailure {
		t.Fatalf("differing archives exit = %d (stderr=%s)", code, errs)
	}
	for _, want := range []string{
		"compared: true", "equal: false", "firstDiff: path=",
		"summary: leftFrames=", "exitCode: 1",
	} {
		if !strings.Contains(out, want) {
			t.Fatalf("replay output missing %q: %s", want, out)
		}
	}
	code, js, _ := invoke("replay", "--format", "json", a, b)
	if code != exitBusinessFailure ||
		!strings.Contains(js, `"equal": false`) ||
		!strings.Contains(js, `"firstDiff"`) ||
		!strings.Contains(js, `"exitCode": 1`) {
		t.Fatalf("JSON diff report = exit %d out %s", code, js)
	}
	code, same, _ := invoke("replay", a, a)
	if code != exitOK || !strings.Contains(same, "equal: true") {
		t.Fatalf("identical archives: exit=%d out=%s", code, same)
	}
}

// TestReplayReportRendering covers both report formats and the diff lines.
func TestReplayReportRendering(t *testing.T) {
	t.Parallel()
	rep := replayReport{
		Schema: reportSchema, Command: "replay", Left: "a",
		Frames: 3, Verified: true, Equal: true,
	}
	if !strings.Contains(string(renderReplayReport(rep, "xml")), "verified: true") {
		t.Fatal("unsupported format must fall back to console output")
	}
	if !strings.Contains(string(renderReplayReport(rep, formatJSON)), `"frames": 3`) {
		t.Fatal("JSON rendering lost the frame count")
	}
	rep.Equal = false
	rep.FirstDiff = "path=Frames[0].Kind"
	rep.Summary = "fieldsDifferent=1"
	console := string(renderReplayConsole(rep))
	if !strings.Contains(console, "firstDiff: path=Frames[0].Kind") ||
		!strings.Contains(console, "summary: fieldsDifferent=1") {
		t.Fatalf("console diff rendering = %s", console)
	}
}

// TestReplayHelpSemantics locks the documented replay and record help text.
func TestReplayHelpSemantics(t *testing.T) {
	t.Parallel()
	for _, want := range []string{
		"Exit codes (PLAN-07 §4):",
		"0  verified",
		"1  comparison difference",
		"2  internal error",
		"3  usage error",
		"abate replay /tmp/run.arch",
	} {
		if !strings.Contains(replayHelp, want) {
			t.Fatalf("replayHelp missing %q", want)
		}
	}
	if !strings.Contains(rootHelp, "replay <archive>") ||
		!strings.Contains(runHelp, "--record path") ||
		!strings.Contains(runHelp, "abate replay /tmp/run.arch") {
		t.Fatal("root or run help does not document replay and --record")
	}
}

// TestExecuteEntryPoint covers the public process entry point.
func TestExecuteEntryPoint(t *testing.T) {
	if code := Execute(nil); code != exitUsageError {
		t.Fatalf("Execute with no arguments = %d", code)
	}
}
