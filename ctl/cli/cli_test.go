/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package cli

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// invoke runs the CLI with captured output streams.
func invoke(args ...string) (code int, out string, errOut string) {
	var stdout, stderr bytes.Buffer
	code = dispatch(args, &stdout, &stderr)
	out = stdout.String()
	errOut = stderr.String()
	return code, out, errOut
}

// shaOf returns the hex SHA-256 digest of a string.
func shaOf(s string) string {
	sum := sha256.Sum256([]byte(s))
	return hex.EncodeToString(sum[:])
}

// tempFile writes data to a temporary file and returns its path.
func tempFile(t *testing.T, data []byte) string {
	t.Helper()
	p := filepath.Join(t.TempDir(), "report.json")
	if err := os.WriteFile(p, data, 0o600); err != nil {
		t.Fatalf("write temp file: %v", err)
	}
	return p
}

// TestVersionCommandCases covers the legal, extra-argument and invalid-flag
// cases of the version command, plus both help spellings.
func TestVersionCommandCases(t *testing.T) {
	t.Parallel()
	for _, c := range []struct {
		name string
		args []string
		want int
	}{
		{"legal", []string{"version"}, exitOK},
		{"extra-positional", []string{"version", "extra"}, exitUsageError},
		{"invalid-flag", []string{"version", "--bogus"}, exitUsageError},
		{"long-help", []string{"version", "--help"}, exitOK},
		{"short-help", []string{"version", "-h"}, exitOK},
	} {
		t.Run(c.name, func(t *testing.T) {
			code, out, errs := invoke(c.args...)
			if code != c.want {
				t.Fatalf("exit = %d, want %d (stderr=%s)", code, c.want, errs)
			}
			if out == "" && errs == "" {
				t.Fatal("version produced no output on either stream")
			}
		})
	}
}

// TestVerifyCommandCases covers legal, missing-argument, invalid-argument,
// missing-file and internal-error cases of the double-run comparison.
func TestVerifyCommandCases(t *testing.T) {
	t.Parallel()
	a := tempFile(t, []byte("report-a\n"))
	b := tempFile(t, []byte("report-b\n"))
	missing := filepath.Join(t.TempDir(), "absent.json")
	badOut := filepath.Join(t.TempDir(), "absent-dir", "out.json")
	for _, c := range []struct {
		name string
		args []string
		want int
	}{
		{"equal", []string{"verify", a, a}, exitOK},
		{"different", []string{"verify", a, b}, exitBusinessFailure},
		{"missing-arg", []string{"verify", a}, exitUsageError},
		{"extra-arg", []string{"verify", a, b, a}, exitUsageError},
		{"invalid-flag", []string{"verify", "--bogus", a, b}, exitUsageError},
		{"invalid-format", []string{"verify", "--format", "xml", a, b}, exitUsageError},
		{"format-equals", []string{"verify", "--format=json", a, a}, exitOK},
		{"missing-right-file", []string{"verify", a, missing}, exitUsageError},
		{"missing-left-file", []string{"verify", missing, a}, exitUsageError},
		{"internal-write-failure", []string{"verify", "--out", badOut, a, a}, exitInternalError},
		{"help", []string{"verify", "--help"}, exitOK},
	} {
		t.Run(c.name, func(t *testing.T) {
			code, _, errs := invoke(c.args...)
			if code != c.want {
				t.Fatalf("exit = %d, want %d (stderr=%s)", code, c.want, errs)
			}
		})
	}
}

// TestRootHelpAndUnregisteredCommands proves the command tree exposes the
// registered commands and rejects every unregistered one. The plugin command
// is registered but requires a subcommand, so a bare invocation is a usage
// error (verified separately by cmd_plugin_test.go).
func TestRootHelpAndUnregisteredCommands(t *testing.T) {
	t.Parallel()
	for _, c := range []struct {
		name string
		args []string
		want int
	}{
		{"root-help-long", []string{"--help"}, exitOK},
		{"root-help-short", []string{"-h"}, exitOK},
		{"root-help-word", []string{"help"}, exitOK},
		{"no-command", nil, exitUsageError},
		{"unknown", []string{"bogus"}, exitUsageError},
		{"unregistered-check", []string{"check", "foo"}, exitUsageError},
		{"unregistered-mutate", []string{"mutate"}, exitUsageError},
		{"unregistered-chaos", []string{"chaos"}, exitUsageError},
		{"unregistered-replay", []string{"replay"}, exitUsageError},
		{"plugin-without-subcommand", []string{"plugin"}, exitUsageError},
		{"unregistered-report", []string{"report"}, exitUsageError},
		{"unregistered-gate", []string{"gate"}, exitUsageError},
	} {
		t.Run(c.name, func(t *testing.T) {
			code, _, _ := invoke(c.args...)
			if code != c.want {
				t.Fatalf("exit = %d, want %d", code, c.want)
			}
		})
	}
}

// TestExitCodeFourStates asserts all four CI-consumed exit codes are reachable.
func TestExitCodeFourStates(t *testing.T) {
	t.Parallel()
	a := tempFile(t, []byte("a\n"))
	b := tempFile(t, []byte("b\n"))
	badOut := filepath.Join(t.TempDir(), "absent-dir", "out.json")
	for _, c := range []struct {
		name string
		args []string
		want int
	}{
		{"success", []string{"version"}, exitOK},
		{"business-failure", []string{"verify", a, b}, exitBusinessFailure},
		{"internal-error", []string{"verify", "--out", badOut, a, a}, exitInternalError},
		{"usage-error", []string{"bogus"}, exitUsageError},
	} {
		t.Run(c.name, func(t *testing.T) {
			code, _, _ := invoke(c.args...)
			if code != c.want {
				t.Fatalf("exit = %d, want %d", code, c.want)
			}
		})
	}
}

// TestHelpExamplesAreRunnable proves a documented example really executes.
func TestHelpExamplesAreRunnable(t *testing.T) {
	t.Parallel()
	if !strings.Contains(rootHelp, "abate run ctl/cli/testdata/ok.yaml") {
		t.Fatal("root help lost its copyable run example")
	}
	code, out, errs := invoke("run", "testdata/ok.yaml", "--format", "json")
	if code != exitOK {
		t.Fatalf("documented example exit = %d (stderr=%s)", code, errs)
	}
	var doc map[string]any
	if err := json.Unmarshal([]byte(out), &doc); err != nil {
		t.Fatalf("example output is not JSON: %v", err)
	}
	if doc["status"] != "COMPLETED" || doc["schema"] != reportSchema {
		t.Fatalf("unexpected example report: %v", doc)
	}
}

// TestVersionDeterminism100 proves version bytes are stable across 100 runs.
func TestVersionDeterminism100(t *testing.T) {
	t.Parallel()
	first := ""
	for i := 0; i < 100; i++ {
		code, out, _ := invoke("version")
		if code != exitOK {
			t.Fatalf("version exit = %d on iteration %d", code, i)
		}
		if i == 0 {
			first = shaOf(out)
			continue
		}
		if shaOf(out) != first {
			t.Fatalf("version digest drifted on iteration %d", i)
		}
	}
}
