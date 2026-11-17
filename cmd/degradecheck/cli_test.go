/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

// chdirToRepoRoot points the process CWD at the repository root so the
// manifest and corpus paths used by measure() resolve exactly as in production.
func chdirToRepoRoot(t *testing.T) {
	t.Helper()
	root, err := filepath.Abs(filepath.Join("..", ".."))
	if err != nil {
		t.Fatalf("abs root: %v", err)
	}
	t.Chdir(root)
}

// TestRunEndToEnd exercises run(), emit() and measure() against the real corpus
// and asserts the report is valid JSON with seven ordered language rows.
func TestRunEndToEnd(t *testing.T) {
	chdirToRepoRoot(t)
	var out, errOut bytes.Buffer
	if rc := run([]string{"degradecheck"}, &out, &errOut); rc != exitOK {
		t.Fatalf("run rc=%d stderr=%s", rc, errOut.String())
	}
	var rep measureReport
	if err := json.Unmarshal(out.Bytes(), &rep); err != nil {
		t.Fatalf("decode report: %v\n%s", err, out.String())
	}
	if rep.SchemaVersion != reportSchema || len(rep.Results) != len(reportLanguages) {
		t.Fatalf("bad report: schema=%d rows=%d", rep.SchemaVersion, len(rep.Results))
	}
	for i, lang := range reportLanguages {
		if rep.Results[i].Language != lang {
			t.Fatalf("row[%d]=%s want %s", i, rep.Results[i].Language, lang)
		}
	}
}

// TestRunRejectsExtraArgs proves run() returns usage exit for bad arity and
// still emits the error through the provided writer.
func TestRunRejectsExtraArgs(t *testing.T) {
	var out, errOut bytes.Buffer
	if rc := run([]string{"degradecheck", "extra"}, &out, &errOut); rc != exitUsage {
		t.Fatalf("run rc=%d want %d", rc, exitUsage)
	}
	if errOut.Len() == 0 {
		t.Fatal("expected usage error on stderr")
	}
}

// errWriter is an io.Writer that always fails, used to exercise encode errors.
type errWriter struct{}

// Write implements io.Writer by returning a fixed error.
func (errWriter) Write([]byte) (int, error) {
	return 0, os.ErrClosed
}

// TestRunMeasureFailure proves run() maps a measure error to the fail exit code
// and writes the reason to stderr.
func TestRunMeasureFailure(t *testing.T) {
	t.Chdir(t.TempDir())
	var out, errOut bytes.Buffer
	if rc := run([]string{"degradecheck"}, &out, &errOut); rc != exitFail {
		t.Fatalf("run rc=%d want %d", rc, exitFail)
	}
	if errOut.Len() == 0 {
		t.Fatal("expected measure error on stderr")
	}
}

// TestRunEncodeFailure proves run() maps an output encode failure to the fail
// exit code instead of silently claiming success.
func TestRunEncodeFailure(t *testing.T) {
	chdirToRepoRoot(t)
	var errOut bytes.Buffer
	if rc := run([]string{"degradecheck"}, errWriter{}, &errOut); rc != exitFail {
		t.Fatalf("run rc=%d want %d", rc, exitFail)
	}
}

// TestMeasureRealCorpus exercises measure() directly and checks that every
// language parsed with zero failures, keeping the full-corpus claim honest.
func TestMeasureRealCorpus(t *testing.T) {
	chdirToRepoRoot(t)
	rep, err := measure()
	if err != nil {
		t.Fatalf("measure: %v", err)
	}
	if rep.ManifestSHA256 == "" {
		t.Fatal("empty manifestSha256")
	}
	total := 0
	for _, row := range rep.Results {
		total += row.Files
		if row.ParseFail != 0 || row.EmptyUnit != 0 {
			t.Fatalf("%s parseFail=%d emptyUnit=%d, want 0/0",
				row.Language, row.ParseFail, row.EmptyUnit)
		}
	}
	if total != 20520 {
		t.Fatalf("total files=%d want 20520", total)
	}
}

// TestManifestUnreadable proves measure() surfaces a missing manifest error
// instead of returning a fabricated report.
func TestManifestUnreadable(t *testing.T) {
	dir := t.TempDir()
	t.Chdir(dir)
	if _, err := measure(); err == nil {
		t.Fatal("measure() with missing manifest must return an error")
	}
}

// TestEmitWrites proves emit() writes formatted text to the target writer.
func TestEmitWrites(t *testing.T) {
	var buf bytes.Buffer
	emit(&buf, "value=%d\n", 7)
	if buf.String() != "value=7\n" {
		t.Fatalf("emit wrote %q", buf.String())
	}
}

// TestChdirHelperSanity guards that the helper actually changed CWD to root by
// checking the manifest file is visible there.
func TestChdirHelperSanity(t *testing.T) {
	chdirToRepoRoot(t)
	if _, err := os.Stat(manifestFile); err != nil {
		t.Fatalf("manifest not visible from repo root: %v", err)
	}
}
