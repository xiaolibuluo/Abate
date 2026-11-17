/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package pitparity

import (
	"errors"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

// goldenLine renders one outcome line of the self-developed golden.
func goldenLine(method string, line int, mutator, expect string) string {
	return strings.Join([]string{
		"outcome", "1", method, strconv.Itoa(line), mutator, expect,
		"node", "class", "KILLED", "detected", "comparable", "true",
	}, "\t")
}

// TestParseGoldenRejectsBadInput covers the self-developed golden validation.
func TestParseGoldenRejectsBadInput(t *testing.T) {
	t.Parallel()
	valid := goldenLine("accepts", 49, "Boundary", ExpectKilled)
	cases := []struct {
		name string
		body string
	}{
		{"no-case", "# header\ntotal=0\n"},
		{"few-fields", "outcome\t1\taccepts\n"},
		{"bad-line", strings.Replace(valid, "\t49\t", "\tx\t", 1)},
		{"zero-line", strings.Replace(valid, "\t49\t", "\t0\t", 1)},
		{"bad-expect", strings.Replace(valid, "\t"+ExpectKilled+"\t", "\tWEIRD\t", 1)},
		{"empty-method", strings.Replace(valid, "\taccepts\t", "\t\t", 1)},
		{"empty-mutator", strings.Replace(valid, "\tBoundary\t", "\t\t", 1)},
	}
	for _, tc := range cases {
		if _, err := ParseGolden([]byte(tc.body)); !errors.Is(err, ErrInvalidGolden) {
			t.Fatalf("%s: error = %v, want ErrInvalidGolden", tc.name, err)
		}
	}
}

// TestParseGoldenSkipsHeaders proves comments and header lines are ignored.
func TestParseGoldenSkipsHeaders(t *testing.T) {
	t.Parallel()
	body := "# comment\ntotal=1\n" + goldenLine("accepts", 49, "Boundary", ExpectKilled) + "\n"
	records, err := ParseGolden([]byte(body))
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if len(records) != 1 {
		t.Fatalf("records = %d, want 1", len(records))
	}
	if records[0].Method != "accepts" || records[0].Line != 49 {
		t.Fatalf("record = %+v", records[0])
	}
}

// TestLocateGoldenWalksUp covers the frozen golden locator.
func TestLocateGoldenWalksUp(t *testing.T) {
	t.Parallel()
	path, err := LocateGolden(".")
	if err != nil {
		t.Fatalf("locate: %v", err)
	}
	if filepath.Base(path) != "mutation-semantics.golden" {
		t.Fatalf("located %q, want the self-developed golden", path)
	}
	if _, err := LocateGolden(filepath.Join(t.TempDir(), "deep", "nested")); err == nil {
		t.Fatal("locate accepted a directory without the repository")
	}
}

// TestLoadGoldenReportsMissingFile covers the read failure branch.
func TestLoadGoldenReportsMissingFile(t *testing.T) {
	t.Parallel()
	if _, err := LoadGolden(filepath.Join(t.TempDir(), "absent.golden")); err == nil {
		t.Fatal("load accepted a missing file")
	}
}

// TestRegenerateGoldenReproducesTheFrozenFile proves the golden is a
// reproducible self-developed observation, not hand written data.
func TestRegenerateGoldenReproducesTheFrozenFile(t *testing.T) {
	t.Parallel()
	path, err := LocateGolden(".")
	if err != nil {
		t.Fatalf("locate: %v", err)
	}
	//nolint:gosec // the golden path is fixed by this test.
	frozen, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read golden: %v", err)
	}
	fresh, err := RegenerateGolden()
	if err != nil {
		t.Fatalf("regenerate: %v", err)
	}
	if string(fresh) != string(frozen) {
		t.Fatal("the regenerated golden differs from the frozen golden")
	}
}
