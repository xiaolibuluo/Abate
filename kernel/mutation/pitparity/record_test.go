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
	"testing"
)

// TestParseMutationsRejectsBadInput covers the report validation branches.
func TestParseMutationsRejectsBadInput(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name string
		body string
	}{
		{"not-json", "{"},
		{"empty", "[]"},
		{"bad-status", `[{"status":"WEIRD","mutatedMethod":"a","mutator":"m","lineNumber":1}]`},
		{"no-method", `[{"status":"KILLED","mutator":"m","lineNumber":1}]`},
		{"no-mutator", `[{"status":"KILLED","mutatedMethod":"a","lineNumber":1}]`},
		{"bad-line", `[{"status":"KILLED","mutatedMethod":"a","mutator":"m","lineNumber":0}]`},
	}
	for _, tc := range cases {
		if _, err := ParseMutations([]byte(tc.body)); !errors.Is(err, ErrInvalidReport) {
			t.Fatalf("%s: error = %v, want ErrInvalidReport", tc.name, err)
		}
	}
}

// TestLocateReportWalksUp covers the report locator.
func TestLocateReportWalksUp(t *testing.T) {
	t.Parallel()
	path, err := LocateReport(".")
	if err != nil {
		t.Fatalf("locate: %v", err)
	}
	if filepath.Base(path) != "pit-mutations.json" {
		t.Fatalf("located %q, want the pit-mutations.json artifact", path)
	}
	if _, err := LocateReport(filepath.Join(t.TempDir(), "deep", "nested")); err == nil {
		t.Fatal("locate accepted a directory without the repository")
	}
}

// TestLoadMutationsReportsMissingFile covers the read failure branch.
func TestLoadMutationsReportsMissingFile(t *testing.T) {
	t.Parallel()
	if _, err := LoadMutations(filepath.Join(t.TempDir(), "absent.json")); err == nil {
		t.Fatal("load accepted a missing file")
	}
}

// TestConvertXMLReproducesTheReference proves the rerun mode is comparable.
func TestConvertXMLReproducesTheReference(t *testing.T) {
	t.Parallel()
	path, err := LocateReport(".")
	if err != nil {
		t.Fatalf("locate: %v", err)
	}
	xmlPath := filepath.Join(filepath.Dir(path), "pit-mutations.xml")
	//nolint:gosec // the artifact path is fixed by this test.
	raw, err := os.ReadFile(xmlPath)
	if err != nil {
		t.Fatalf("read XML: %v", err)
	}
	converted, err := ConvertXML(raw)
	if err != nil {
		t.Fatalf("convert: %v", err)
	}
	//nolint:gosec // the artifact path is fixed by this test.
	reference, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read reference: %v", err)
	}
	if string(converted) != string(reference) {
		t.Fatal("the XML conversion does not reproduce the vendored reference report")
	}
}

// TestConvertXMLRejectsBadInput covers the converter error branches.
func TestConvertXMLRejectsBadInput(t *testing.T) {
	t.Parallel()
	if _, err := ConvertXML([]byte("<mutations>")); !errors.Is(err, ErrInvalidReport) {
		t.Fatalf("error = %v, want ErrInvalidReport", err)
	}
	if _, err := ConvertXML([]byte("<mutations></mutations>")); !errors.Is(err, ErrInvalidReport) {
		t.Fatalf("error = %v, want ErrInvalidReport", err)
	}
}

// TestShortMutatorAndKillingTest covers the name shortening helpers.
func TestShortMutatorAndKillingTest(t *testing.T) {
	t.Parallel()
	if got := shortMutator("org.pitest.mutationtest.MathMutator"); got != "MathMutator" {
		t.Fatalf("short mutator = %q", got)
	}
	if got := shortMutator("MathMutator"); got != "MathMutator" {
		t.Fatalf("short mutator = %q", got)
	}
	if got := shortKillingTest("short"); got != "short" {
		t.Fatalf("short killing test = %q", got)
	}
	long := "com.abate.parity.ParityTargetTest.[method:t1ConstantAndReturn()]"
	if got := shortKillingTest(long); got != "t1ConstantAndReturn()]" {
		t.Fatalf("short killing test = %q", got)
	}
	if got := shortKillingTest("plain"); got != "plain" {
		t.Fatalf("short killing test = %q", got)
	}
}
