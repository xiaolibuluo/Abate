/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package main

import (
	"encoding/json"
	"strings"
	"testing"
	"testing/fstest"

	"abate/lang/air"
	"abate/lang/frontkit/grammar"
)

// TestNewTallyHasAllLanguages proves every canonical language has an
// observation map so a language with zero files cannot hide a nil deref.
func TestNewTallyHasAllLanguages(t *testing.T) {
	s := newTally()
	for _, lang := range reportLanguages {
		if s.observed[lang] == nil {
			t.Fatalf("newTally: observed[%s] is nil", lang)
		}
	}
}

// TestBuildRowGoNotMeasured proves Go keeps its honest NOT_MEASURED status even
// when it has real files and observed kinds, never a fabricated zero-degrade claim.
func TestBuildRowGoNotMeasured(t *testing.T) {
	s := newTally()
	s.files[goLang] = 100
	s.parseOK[goLang] = 100
	s.observed[goLang][air.KindFuncDecl] = 5
	row := buildRow(goLang, s)
	if row.Status != notMeasured {
		t.Fatalf("buildRow(go).Status = %s, want %s", row.Status, notMeasured)
	}
	if row.DegradePct != 0 {
		t.Fatalf("buildRow(go).DegradePct = %v, want 0", row.DegradePct)
	}
}

// TestBuildRowDegradedNeverReclassified proves a degraded mapping with an
// observed kind is counted as degraded, so the ratio cannot be lowered by
// silently promoting it to exact.
func TestBuildRowDegradedNeverReclassified(t *testing.T) {
	facts, err := grammar.Load("js")
	if err != nil {
		t.Fatalf("grammar.Load(js): %v", err)
	}
	var degradedKind air.NodeKind
	found := false
	for _, m := range grammar.Map(facts, grammar.DefaultRules("js")) {
		if m.Status == grammar.StatusDegraded {
			degradedKind = m.Kind
			found = true
			break
		}
	}
	if !found {
		t.Fatal("js grammar has no degraded mapping to test")
	}
	s := newTally()
	s.observed["js"][degradedKind] = 3
	row := buildRow("js", s)
	if row.Status != measured {
		t.Fatalf("buildRow(js).Status = %s, want %s", row.Status, measured)
	}
	if row.DegradedObserved < 1 {
		t.Fatal("buildRow(js): observed degraded mapping was reclassified away")
	}
	if row.DegradePct <= 0 || row.DegradePct > 100 {
		t.Fatalf("buildRow(js).DegradePct = %v outside (0,100]", row.DegradePct)
	}
}

// TestObserveEntryRecordsParseFailure proves a missing corpus file is counted
// as a parse failure rather than skipped, keeping the report honest.
func TestObserveEntryRecordsParseFailure(t *testing.T) {
	ft, err := e2eFrontendsForTest()
	if err != nil {
		t.Fatalf("frontend table: %v", err)
	}
	s := newTally()
	fsys := fstest.MapFS{}
	observeEntry(fsys, ft, s, corpusEntryForTest(goLang, "missing.go"))
	if s.files[goLang] != 1 || s.parseFail[goLang] != 1 || s.parseOK[goLang] != 0 {
		t.Fatalf("missing file: files=%d parseFail=%d parseOK=%d, want 1/1/0",
			s.files[goLang], s.parseFail[goLang], s.parseOK[goLang])
	}
}

// TestObserveEntryUnknownLanguage proves an unsupported language is recorded
// as a parse failure through the frontend lookup branch, never skipped.
func TestObserveEntryUnknownLanguage(t *testing.T) {
	ft, err := e2eFrontendsForTest()
	if err != nil {
		t.Fatalf("frontend table: %v", err)
	}
	s := newTally()
	fsys := fstest.MapFS{"x.zzz": {Data: []byte("anything")}}
	observeEntry(fsys, ft, s, corpusEntryForTest("zzz", "x.zzz"))
	if s.files["zzz"] != 1 || s.parseFail["zzz"] != 1 {
		t.Fatalf("unknown lang: files=%d parseFail=%d, want 1/1",
			s.files["zzz"], s.parseFail["zzz"])
	}
}

// TestObserveEntrySuccess proves a real Go source file is parsed into observed
// kinds through the success path of observeEntry.
func TestObserveEntrySuccess(t *testing.T) {
	ft, err := e2eFrontendsForTest()
	if err != nil {
		t.Fatalf("frontend table: %v", err)
	}
	src := "package p\n\nfunc F() int { return 1 }\n"
	s := newTally()
	fsys := fstest.MapFS{"p.go": {Data: []byte(src)}}
	observeEntry(fsys, ft, s, corpusEntryForTest(goLang, "p.go"))
	if s.files[goLang] != 1 || s.parseOK[goLang] != 1 || s.parseFail[goLang] != 0 {
		t.Fatalf("success: files=%d parseOK=%d parseFail=%d, want 1/1/0",
			s.files[goLang], s.parseOK[goLang], s.parseFail[goLang])
	}
	if len(s.observed[goLang]) == 0 {
		t.Fatal("success path produced no observed kinds")
	}
}

// TestDeterministicJSON proves the report serializes to stable field order so
// two runs are byte-identical.
func TestDeterministicJSON(t *testing.T) {
	r := measureReport{
		SchemaVersion:  reportSchema,
		ManifestSHA256: "abc",
		Formula:        degradeFormula,
		Results:        []langRow{{Language: goLang, Status: notMeasured}},
	}
	data, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	first := string(data)
	again, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		t.Fatalf("marshal2: %v", err)
	}
	if first != string(again) {
		t.Fatal("report JSON is not deterministic across calls")
	}
	if !strings.Contains(first, `"schemaVersion"`) {
		t.Fatalf("report JSON missing schemaVersion: %s", first)
	}
}
