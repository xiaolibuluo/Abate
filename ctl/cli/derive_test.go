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
	"slices"
	"strings"
	"testing"
)

// TestDeriveCommandCases covers legal, help, usage and internal-error states.
func TestDeriveCommandCases(t *testing.T) {
	t.Parallel()
	badOut := filepath.Join(t.TempDir(), "absent-dir", "out.json")
	for _, c := range []struct {
		name string
		args []string
		want int
	}{
		{"long-help", []string{"derive", "--help"}, exitOK},
		{"short-help", []string{"derive", "-h"}, exitOK},
		{"legal-console", []string{"derive"}, exitOK},
		{"legal-json", []string{"derive", "--format", "json"}, exitOK},
		{"seed-override", []string{"derive", "--seed", "7"}, exitOK},
		{"format-equals", []string{"derive", "--format=json"}, exitOK},
		{"bad-flag", []string{"derive", "--bogus"}, exitUsageError},
		{"bad-format", []string{"derive", "--format", "xml"}, exitUsageError},
		{"extra-arg", []string{"derive", "extra"}, exitUsageError},
		{"bad-out", []string{"derive", "--out", badOut}, exitInternalError},
	} {
		t.Run(c.name, func(t *testing.T) {
			code, out, errs := invoke(c.args...)
			if code != c.want {
				t.Fatalf("exit = %d, want %d (stderr=%s)", code, c.want, errs)
			}
			if out == "" && errs == "" {
				t.Fatal("derive produced no output on either stream")
			}
		})
	}
}

// TestDeriveJSONContent proves the real CLI reports at least 20 distinct
// derived operators with allocated verdicts and honest kill semantics.
func TestDeriveJSONContent(t *testing.T) {
	t.Parallel()
	code, out, errs := invoke("derive", "--seed", "42", "--format", "json")
	if code != exitOK {
		t.Fatalf("derive exit = %d (stderr=%s)", code, errs)
	}
	var view deriveView
	if err := json.Unmarshal([]byte(out), &view); err != nil {
		t.Fatalf("derive JSON is invalid: %v", err)
	}
	if view.Schema != deriveSchema || view.Command != "derive" ||
		view.Seed != 42 || view.ExitCode != exitOK {
		t.Fatalf("unexpected report header: %+v", view)
	}
	if view.Rules != 110 {
		t.Fatalf("derive used %d rules, want the 110 built-in rules", view.Rules)
	}
	if len(view.DerivedIDs) < 20 {
		t.Fatalf("report enumerates only %d derived operators", len(view.DerivedIDs))
	}
	distinct := make(map[string]bool, len(view.DerivedIDs))
	for _, id := range view.DerivedIDs {
		if !strings.HasPrefix(id, "m-derived-") {
			t.Fatalf("non-derived operator id leaked into the report: %s", id)
		}
		if distinct[id] {
			t.Fatalf("duplicate derived operator id %s", id)
		}
		distinct[id] = true
	}
	if !slices.IsSorted(view.DerivedIDs) {
		t.Fatal("derived identifier enumeration is not sorted")
	}
	if view.Mutants <= 0 || view.Killed == 0 || view.Survived == 0 {
		t.Fatalf("no allocated verdicts in the report: %+v", view)
	}
	total := view.Killed + view.Survived + view.Equivalent + view.Timeout
	if total != uint64(view.Mutants) { //nolint:gosec // the point count is a small non-negative integer.
		t.Fatalf("verdict total %d disagrees with %d mutants", total, view.Mutants)
	}
	if !strings.Contains(view.KillingModel, "static-air-approximation") ||
		!strings.Contains(view.KillingModel, "not a real go test kill") {
		t.Fatalf("kill model is not reported honestly: %q", view.KillingModel)
	}
	if len(view.Fingerprint) != 64 || len(view.ReportDigest) != 64 {
		t.Fatalf("report digests are incomplete: %q / %q",
			view.Fingerprint, view.ReportDigest)
	}
}

// TestDeriveConsoleEnumeratesIDs proves the console report really lists ids.
func TestDeriveConsoleEnumeratesIDs(t *testing.T) {
	t.Parallel()
	code, out, errs := invoke("derive")
	if code != exitOK {
		t.Fatalf("derive exit = %d (stderr=%s)", code, errs)
	}
	if strings.Count(out, "m-derived-") < 20 {
		t.Fatal("console report did not enumerate at least 20 derived operators")
	}
	for _, want := range []string{
		"abate derive report", "schema: abate.derive.v1", "derivedIds:",
		"killingModel: static-air-approximation", "fingerprint:", "reportDigest:",
	} {
		if !strings.Contains(out, want) {
			t.Fatalf("console report missing %q", want)
		}
	}
}

// TestDeriveDeterminism100 proves 100 identical CLI invocations yield one hash.
func TestDeriveDeterminism100(t *testing.T) {
	t.Parallel()
	seen := make(map[string]struct{}, 1)
	for i := 0; i < 100; i++ {
		code, out, errs := invoke("derive", "--seed", "42", "--format", "json")
		if code != exitOK {
			t.Fatalf("run %d exit = %d (stderr=%s)", i, code, errs)
		}
		seen[shaOf(out)] = struct{}{}
	}
	if len(seen) != 1 {
		t.Fatalf("derive produced %d distinct report hashes over 100 runs", len(seen))
	}
}

// TestDeriveOutFile proves --out writes exactly the stdout report bytes.
func TestDeriveOutFile(t *testing.T) {
	t.Parallel()
	target := filepath.Join(t.TempDir(), "derive.json")
	code, out, errs := invoke("derive", "--format", "json", "--out", target)
	if code != exitOK {
		t.Fatalf("exit = %d (stderr=%s)", code, errs)
	}
	data, err := os.ReadFile(target) //nolint:gosec // the target is a test-owned temporary file.
	if err != nil {
		t.Fatalf("read report: %v", err)
	}
	if string(data) != out {
		t.Fatal("--out bytes differ from stdout bytes")
	}
}

// TestDeriveHelpConsistency locks both help texts to the same command and the
// same shared exit-code table, so check and root help cannot contradict.
func TestDeriveHelpConsistency(t *testing.T) {
	t.Parallel()
	for _, want := range []string{
		"Exit codes (PLAN-07 §4):",
		"0  success", "1  business failure",
		"2  internal error", "3  usage error",
	} {
		if !strings.Contains(rootHelp, want) {
			t.Fatalf("rootHelp missing %q", want)
		}
		if !strings.Contains(deriveHelp, want) {
			t.Fatalf("deriveHelp missing %q", want)
		}
	}
	if !strings.Contains(rootHelp, "derive") ||
		!strings.Contains(rootHelp, "abate derive --seed 42 --format json") {
		t.Fatal("rootHelp does not advertise the derive command")
	}
	if !strings.Contains(checkHelp, "abate derive command") {
		t.Fatal("checkHelp does not point at the derive command")
	}
	if !strings.Contains(rootHelp, "static AIR-level") ||
		!strings.Contains(deriveHelp, "static AIR-level") {
		t.Fatal("help texts disagree about the static AIR-level kill approximation")
	}
	if !slices.Equal(deriveValueFlags(), []string{"seed", "format", "out", "source"}) {
		t.Fatal("derive value flags changed")
	}
}

// TestDeriveSourceFlag proves a caller-supplied Go source is parsed into AIR
// and that source read/parse failures map to the internal-error exit state.
func TestDeriveSourceFlag(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()
	good := filepath.Join(dir, "good.go")
	bad := filepath.Join(dir, "bad.go")
	missing := filepath.Join(dir, "missing.go")
	if err := os.WriteFile(good,
		[]byte("package sample\nfunc f(a int) bool { return a < 1 }\n"), 0o600); err != nil {
		t.Fatalf("write good source: %v", err)
	}
	if err := os.WriteFile(bad, []byte("this is not go source\n"), 0o600); err != nil {
		t.Fatalf("write bad source: %v", err)
	}
	for _, c := range []struct {
		name string
		path string
		want int
	}{
		{"valid-source", good, exitOK},
		{"invalid-source", bad, exitInternalError},
		{"missing-source", missing, exitInternalError},
	} {
		t.Run(c.name, func(t *testing.T) {
			code, out, errs := invoke("derive", "--format", "json", "--source", c.path)
			if code != c.want {
				t.Fatalf("exit = %d, want %d (stderr=%s)", code, c.want, errs)
			}
			if c.want == exitOK && !strings.Contains(out, "m-derived-") {
				t.Fatal("a valid source did not produce derived operators")
			}
			if c.want != exitOK && !strings.Contains(errs, "abate derive:") {
				t.Fatalf("source failure was not reported: %q", errs)
			}
		})
	}
}

// TestDeriveRenderFallback covers the console fallback and digest helper.
func TestDeriveRenderFallback(t *testing.T) {
	t.Parallel()
	view := deriveView{Schema: deriveSchema, Command: "derive"}
	console := renderDeriveReport(view, formatConsole)
	if !bytes.Equal(renderDeriveReport(view, "xml"), console) {
		t.Fatal("an unknown format must fall back to deterministic console output")
	}
	if len(digestHex([]byte("abate"))) != 64 {
		t.Fatal("report digest is not a full SHA-256 hex string")
	}
}
