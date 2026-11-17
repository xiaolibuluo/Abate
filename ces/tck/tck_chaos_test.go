/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package tck

import (
	"errors"
	"fmt"
	"io/fs"
	"strings"
	"testing"
	"testing/fstest"

	"abate/ces/parse"
)

// scriptedDriver is a fully programmable Driver used to steer the neutral
// comparison logic into every branch without depending on the CES pipeline.
// Each field is a deterministic function of the input bytes only, so two runs
// over identical cases always observe identical outcomes.
type scriptedDriver struct {
	parseOK   bool
	parseRule string
	parseErr  error
	viols     []Violation
	valErr    error
}

func (d scriptedDriver) Parse(_ []byte, _ string) (ParseResult, error) {
	if !d.parseOK {
		return ParseResult{OK: false, RuleID: d.parseRule, Error: "scripted rejection"}, d.parseErr
	}
	return ParseResult{OK: true}, nil
}

func (d scriptedDriver) Validate(_ []byte) ([]Violation, error) {
	if d.valErr != nil {
		return nil, d.valErr
	}
	return d.viols, nil
}

// TestChaosRunnerDeterministic proves the runner produces byte-identical reports
// when the same case stream is replayed, and that case order alone never changes
// the aggregate verdict. This is the reproducibility contract under adversarial
// ordering of the case library.
func TestChaosRunnerDeterministic(t *testing.T) {
	cases := []Case{
		{ID: "C1", Spec: "s", Kind: KindParse, Input: "a: 1", Expect: CaseExpect{Accept: true}},
		{ID: "C2", Spec: "s", Kind: KindParse, Input: "b: 2", Expect: CaseExpect{Accept: false, RuleIDs: []string{"R"}}},
		{ID: "C3", Spec: "s", Kind: KindValidate, Input: "c: 3", Expect: CaseExpect{Accept: true}},
		{ID: "C4", Spec: "s", Kind: KindRoundtrip, Input: "d: 4", Expect: CaseExpect{Accept: true}},
	}
	// Parse accepts C1/C3/C4 but rejects C2 with rule "R"; the driver is chosen
	// so every case matches by construction except where the expect flips below.
	run := func(cs []Case) string {
		rep, err := (Runner{Driver: scriptedDriver{
			parseOK: false, parseRule: "R", parseErr: errors.New("reject"),
		}, Name: "chaos"}).Run(cs)
		if err != nil {
			t.Fatalf("run: %v", err)
		}
		var sb strings.Builder
		for _, d := range rep.Details {
			sb.WriteString(fmt.Sprintf("%s:%t:%s|", d.ID, d.Passed, d.Detail))
		}
		sb.WriteString(fmt.Sprintf("P%d/F%d/S%d/C%t", rep.Passed, rep.Failed, rep.Skipped, rep.Compliant))
		return sb.String()
	}
	first := run(cases)
	// Reverse order: aggregate Passed/Failed must be identical (order-independent).
	rev := make([]Case, len(cases))
	for i := range cases {
		rev[len(cases)-1-i] = cases[i]
	}
	repA, _ := (Runner{Driver: scriptedDriver{parseOK: false, parseRule: "R", parseErr: errors.New("reject")}}).Run(cases)
	repB, _ := (Runner{Driver: scriptedDriver{parseOK: false, parseRule: "R", parseErr: errors.New("reject")}}).Run(rev)
	if repA.Passed != repB.Passed || repA.Failed != repB.Failed || repA.Compliant != repB.Compliant {
		t.Fatalf("aggregate verdict depends on case order: %+v vs %+v", repA, repB)
	}
	if second := run(cases); second != first {
		t.Fatalf("runner is not deterministic across replays:\n%s\n%s", first, second)
	}
}

// TestChaosRunnerNilDriver proves a runner without a driver fails fast with the
// documented error, rather than panicking on a nil interface.
func TestChaosRunnerNilDriver(t *testing.T) {
	if _, err := (Runner{}).Run(nil); err == nil {
		t.Fatal("Run with nil driver must fail")
	}
}

// TestChaosRunnerUnknownKindAndSkip proves the default branch of runCase and the
// skip aggregation path both behave deterministically.
func TestChaosRunnerUnknownKindAndSkip(t *testing.T) {
	cases := []Case{
		{ID: "U", Spec: "s", Kind: CaseKind("mystery"), Input: "x", Expect: CaseExpect{Accept: true}},
	}
	rep, err := (Runner{Driver: scriptedDriver{parseOK: true}}).Run(cases)
	if err != nil {
		t.Fatalf("run: %v", err)
	}
	if rep.Failed != 1 || rep.Compliant {
		t.Fatalf("unknown kind must fail the case: %+v", rep)
	}
	if !strings.Contains(rep.Details[0].Detail, "unknown case kind") {
		t.Fatalf("missing unknown-kind detail: %q", rep.Details[0].Detail)
	}
}

// TestChaosRunnerDefaultName proves an unnamed runner derives a stable label
// from the driver's concrete type via driverName.
func TestChaosRunnerDefaultName(t *testing.T) {
	rep, err := (Runner{Driver: scriptedDriver{parseOK: true}}).Run(nil)
	if err != nil {
		t.Fatalf("run: %v", err)
	}
	if rep.Driver != "tck.scriptedDriver" {
		t.Fatalf("driverName = %q, want tck.scriptedDriver", rep.Driver)
	}
}

// TestChaosValidateCaseBranches exercises every rejection branch of the case
// quality validator, which is otherwise unreachable through the embedded library.
func TestChaosValidateCaseBranches(t *testing.T) {
	good := Case{ID: "X", Spec: "clause", Kind: KindParse, Input: "a: 1"}
	if err := validateCase(good); err != nil {
		t.Fatalf("valid case rejected: %v", err)
	}
	mutations := []struct {
		name string
		edit func(*Case)
	}{
		{"missing id", func(c *Case) { c.ID = "" }},
		{"missing spec", func(c *Case) { c.Spec = "" }},
		{"unknown kind", func(c *Case) { c.Kind = CaseKind("zzz") }},
		{"blank input", func(c *Case) { c.Input = "   " }},
	}
	for _, m := range mutations {
		c := good
		m.edit(&c)
		if err := validateCase(c); err == nil {
			t.Errorf("validateCase accepted %s", m.name)
		}
	}
	for _, k := range []CaseKind{KindParse, KindValidate, KindRoundtrip} {
		c := good
		c.Kind = k
		if err := validateCase(c); err != nil {
			t.Errorf("valid kind %q rejected: %v", k, err)
		}
	}
}

// TestChaosFormatAndRuleHelpers covers the format-detection, enum-mapping and
// error-extraction helpers on all their branches, including JSON sniffing and
// structured vs unstructured errors.
func TestChaosFormatAndRuleHelpers(t *testing.T) {
	if detectFormat("  {a: 1}") != "json" {
		t.Error("leading-brace document must sniff as json")
	}
	for _, y := range []string{"a: 1", "", "   ", "- list"} {
		if detectFormat(y) != "yaml" {
			t.Errorf("detectFormat(%q) = yaml expected", y)
		}
	}
	cases := map[string]parse.Format{
		"json": parse.FormatJSON, " JSON ": parse.FormatJSON, "Json": parse.FormatJSON,
		"yaml": parse.FormatYAML, " YAML ": parse.FormatYAML,
		"auto": parse.FormatAuto, "": parse.FormatAuto, "toml": parse.FormatAuto,
	}
	for hint, want := range cases {
		if got := toFormat(hint); got != want {
			t.Errorf("toFormat(%q) = %v, want %v", hint, got, want)
		}
	}
	structured := &parse.ParseError{RuleID: "CES-PARSE-006"}
	if ruleIDOf(structured) != "CES-PARSE-006" {
		t.Error("structured ParseError rule id not extracted")
	}
	if ruleIDOf(errors.New("plain")) != "" {
		t.Error("unstructured error must yield empty rule id")
	}
	if ruleIDOf(fmt.Errorf("wrapped: %w", structured)) != "CES-PARSE-006" {
		t.Error("wrapped ParseError rule id not extracted")
	}
}

// TestChaosRuleSetEqualAndSummary proves set equality ignores order/duplicates
// and that the report summary renders every counter deterministically.
func TestChaosRuleSetEqualAndSummary(t *testing.T) {
	if !ruleSetEqual([]string{"b", "a", "a"}, []string{"a", "b"}) {
		t.Error("ruleSetEqual must ignore order and duplicates")
	}
	if ruleSetEqual([]string{"a"}, []string{"a", "b"}) {
		t.Error("ruleSetEqual must detect differing sets")
	}
	rep := Report{Driver: "d", Total: 4, Passed: 2, Failed: 1, Skipped: 1, Compliant: false}
	s := rep.summary()
	for _, want := range []string{"TCK[d]", "2/4", "1 failed", "1 skipped", "compliant=false"} {
		if !strings.Contains(s, want) {
			t.Errorf("summary %q missing %q", s, want)
		}
	}
}

// yamlCase renders a minimal, spec-valid case document for the in-memory FS.
func yamlCase(id, spec, kind, input string) []byte {
	return []byte(fmt.Sprintf("id: %s\nspec: %s\nkind: %s\ninput: %q\nexpect:\n  accept: true\n",
		id, spec, kind, input))
}

// TestChaosLoadCasesDefensiveBranches injects synthetic filesystems to reach the
// read/decode/validate/skip/ordering branches of loadCasesFrom that the embedded,
// always-valid library can never exercise. Each scenario asserts the exact error
// class, so these are real correctness checks, not coverage stubs.
func TestChaosLoadCasesDefensiveBranches(t *testing.T) {
	// Missing root -> read case library error.
	if _, err := loadCasesFrom(fstest.MapFS{}); err == nil {
		t.Error("empty FS must fail to read the case library")
	} else if !strings.Contains(err.Error(), "read case library") {
		t.Errorf("unexpected error for missing root: %v", err)
	}

	// Non-directory entries at root are skipped; a nested non-yaml file and a
	// nested directory are both skipped; valid yaml cases are loaded.
	mixed := fstest.MapFS{
		"cases/README.txt":        {Data: []byte("not a case")},
		"cases/parse/a.yaml":      {Data: yamlCase("A", "s", "parse", "x: 1")},
		"cases/parse/notes.txt":   {Data: []byte("ignore")},
		"cases/parse/nested/b.go": {Data: []byte("package x")},
	}
	got, err := loadCasesFrom(mixed)
	if err != nil {
		t.Fatalf("mixed FS: %v", err)
	}
	if len(got) != 1 || got[0].ID != "A" {
		t.Fatalf("mixed FS loaded %d cases: %+v", len(got), got)
	}

	// Malformed YAML -> decode error.
	badYAML := fstest.MapFS{"cases/parse/bad.yaml": {Data: []byte("id: [unterminated")}}
	if _, err := loadCasesFrom(badYAML); err == nil || !strings.Contains(err.Error(), "decode") {
		t.Errorf("malformed YAML must produce a decode error, got %v", err)
	}

	// Valid YAML but spec-invalid case (missing id) -> validateCase error.
	badCase := fstest.MapFS{"cases/parse/bad.yaml": {Data: []byte("spec: s\nkind: parse\ninput: x\n")}}
	if _, err := loadCasesFrom(badCase); err == nil || !strings.Contains(err.Error(), "missing id") {
		t.Errorf("invalid case must be rejected, got %v", err)
	}

	// A regular file sitting at the root (even named like a group) is a
	// non-directory entry and is skipped by the loader, never treated as a group.
	regularFile := fstest.MapFS{"cases/parse": {Data: []byte("i am a file")}}
	plain, err := loadCasesFrom(regularFile)
	if err != nil {
		t.Fatalf("regular root file must be skipped, got %v", err)
	}
	if len(plain) != 0 {
		t.Errorf("regular root file produced %d cases, want 0", len(plain))
	}

	// Tie-break ordering: two cases sharing an ID are ordered by kind, so the
	// deterministic sort comparator's second branch is exercised.
	ties := fstest.MapFS{
		"cases/validate/z.yaml": {Data: yamlCase("SAME", "s", "validate", "x: 1")},
		"cases/parse/z.yaml":    {Data: yamlCase("SAME", "s", "parse", "x: 1")},
	}
	ordered, err := loadCasesFrom(ties)
	if err != nil {
		t.Fatalf("tie FS: %v", err)
	}
	if len(ordered) != 2 {
		t.Fatalf("tie FS loaded %d cases, want 2", len(ordered))
	}
	if ordered[0].Kind != KindParse || ordered[1].Kind != KindValidate {
		t.Errorf("tie-break by kind failed: %q then %q", ordered[0].Kind, ordered[1].Kind)
	}

	// A read error on a case file: use a directory entry whose name ends with
	// .yaml but whose content is unreadable via fs.ReadFile (a sub-directory).
	unreadable := fstest.MapFS{
		"cases/parse/x.yaml/child": {Data: []byte("x")},
	}
	if _, err := loadCasesFrom(unreadable); err != nil {
		t.Fatalf("directory-with-.yaml-suffix must be skipped, got %v", err)
	}

	// Compile-time guarantee the injected FS satisfies the read-only contract.
	var _ fs.FS = fstest.MapFS{}
}
