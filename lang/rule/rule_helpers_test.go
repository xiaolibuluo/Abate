/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package rule

import (
	"errors"
	"slices"
	"strings"
	"testing"

	"abate/lang/air"
)

// testTree builds a small AIR unit with one package root and one call child.
func testTree() (unit *air.Unit, root air.Node, call air.Node) {
	ident := air.NewNode(air.KindIdent,
		air.NewPosition(2, 3, 10, 16), "foo_bar")
	lit := air.NewNode(air.KindStrLit,
		air.NewPosition(2, 8, 17, 22), `"value"`)
	call = air.NewNode(air.KindCall,
		air.NewPosition(2, 1, 8, 23), "", ident, lit)
	root = air.NewNode(air.KindPackage,
		air.NewPosition(1, 1, 0, 23), "p", call)
	unit = air.NewUnit(air.NewPackageRef("p", nil))
	if err := unit.AddFile(air.NewFileRef("p.go", []byte("package p")), root); err != nil {
		panic(err)
	}
	return unit, root, call
}

func validSpec(id string) Spec {
	return Spec{
		ID: RuleID(id), Reference: "Effective Go, Names",
		Severity: SeverityWarning, Category: CategoryNaming,
		Pattern: "Ident~foo", Message: "bad identifier $x",
	}
}

// TestSeverityAndCategory exercises every allocated value and the defaults.
func TestSeverityAndCategory(t *testing.T) {
	t.Parallel()
	for _, sev := range []Severity{SeverityError, SeverityWarning, SeverityInfo} {
		if !sev.Valid() || sev.Rank() < 1 || sev.String() != string(sev) {
			t.Fatalf("severity %q is inconsistent", sev)
		}
	}
	if Severity("bogus").Valid() || Severity("bogus").Rank() != 0 {
		t.Fatal("unknown severity must be invalid with rank zero")
	}
	for _, cat := range []Category{
		CategoryNaming, CategoryStyle, CategoryErrorProne,
		CategoryPerformance, CategoryConcurrency, CategoryBoundary,
	} {
		if !cat.Valid() || cat.String() != string(cat) {
			t.Fatalf("category %q is inconsistent", cat)
		}
	}
	if Category("bogus").Valid() {
		t.Fatal("unknown category must be invalid")
	}
}

// TestRuleID covers parsing, accessors and every malformed shape.
func TestRuleID(t *testing.T) {
	t.Parallel()
	id, err := ParseRuleID("GO-CRC-014")
	if err != nil || id != "GO-CRC-014" {
		t.Fatalf("ParseRuleID = %q, %v", id, err)
	}
	if !id.Valid() || id.Language() != "go" || id.Source() != "CRC" ||
		id.Sequence() != 14 || id.String() != "GO-CRC-014" {
		t.Fatalf("accessors disagree for %s", id)
	}
	for _, bad := range []string{"", "GO", "GO-CRC", "go-CRC-001",
		"GO-CRC-01", "GO-CRC-0001", "GO-CR_c-001", "GO-CRC-abc"} {
		if _, err := ParseRuleID(bad); !errors.Is(err, ErrInvalidRuleID) {
			t.Errorf("ParseRuleID(%q) = %v, want ErrInvalidRuleID", bad, err)
		}
	}
	broken := RuleID("nope")
	if broken.Valid() || broken.Language() != "" || broken.Source() != "" ||
		broken.Sequence() != 0 {
		t.Fatal("malformed RuleID must return zero accessors")
	}
	if !isUpperLetter('A') || isUpperLetter('1') || !isDigit('9') || isDigit('x') {
		t.Fatal("byte classification is wrong")
	}
}

// TestSpecValidation covers every schema rejection and the canonical form.
func TestSpecValidation(t *testing.T) {
	t.Parallel()
	good := validSpec("GO-EG-001")
	if err := good.Validate(); err != nil {
		t.Fatalf("valid spec rejected: %v", err)
	}
	if good.Language() != "go" || good.Canonical() == "" {
		t.Fatal("spec accessors are incomplete")
	}
	variants := map[string]func(*Spec){
		"bad-id":        func(s *Spec) { s.ID = "bad" },
		"no-reference":  func(s *Spec) { s.Reference = " " },
		"bad-severity":  func(s *Spec) { s.Severity = "bogus" },
		"bad-category":  func(s *Spec) { s.Category = "bogus" },
		"empty-pattern": func(s *Spec) { s.Pattern = " " },
		"empty-message": func(s *Spec) { s.Message = "" },
		"bad-mutation":  func(s *Spec) { s.Mutation = MutationDecl{Family: "M"} },
	}
	names := make([]string, 0, len(variants))
	for name := range variants {
		names = append(names, name)
	}
	slices.Sort(names)
	for _, name := range names {
		candidate := good
		variants[name](&candidate)
		if err := candidate.Validate(); err == nil {
			t.Errorf("variant %s was accepted", name)
		}
	}
}

// TestMutationDecl covers enabled, partial and canonical forms.
func TestMutationDecl(t *testing.T) {
	t.Parallel()
	empty := MutationDecl{}
	if empty.Enabled() || empty.EnabledPartially() || !empty.Valid() {
		t.Fatal("zero mutation declaration must be valid and disabled")
	}
	partial := MutationDecl{Family: "M-NAM"}
	if !partial.EnabledPartially() || partial.Valid() {
		t.Fatal("partial mutation declaration must be rejected")
	}
	full := MutationDecl{Family: "M-NAM", Kind: "rename", Params: []string{"x"}}
	if !full.Enabled() || !full.Valid() || full.Canonical() == "" {
		t.Fatal("full mutation declaration must be valid")
	}
	withBlank := full
	withBlank.Params = []string{" "}
	if withBlank.Valid() {
		t.Fatal("blank mutation parameter must be rejected")
	}
}

// TestFindingAndCapture covers keys, strings and canonical forms.
func TestFindingAndCapture(t *testing.T) {
	t.Parallel()
	capture := Capture{Name: "x", Kind: "Ident", Text: "foo", Line: 2, Column: 3}
	if !strings.Contains(capture.String(), "foo") || capture.Canonical() == "" {
		t.Fatal("capture rendering is incomplete")
	}
	finding := Finding{
		RuleID: "GO-EG-001", Path: "p.go", Line: 2, Column: 3,
		Severity: SeverityWarning, Category: CategoryNaming,
		Message: "bad", Reference: "Effective Go", Captures: []Capture{capture},
	}
	if finding.Key() == "" || !strings.Contains(finding.String(), "p.go:2:3") {
		t.Fatal("finding key or console form is wrong")
	}
	if !strings.Contains(finding.Canonical(), "foo") {
		t.Fatal("finding canonical form omitted the capture")
	}
}

// TestMatcherSemantics covers attribute operators and annotation checks.
func TestMatcherSemantics(t *testing.T) {
	t.Parallel()
	_, _, call := testTree()
	ident := call.Children()[0]
	air.Annotate(call, air.SemPurity, "pure")
	for _, text := range []string{
		"Call", "ANY", "Call!=zzz", "Ident=foo_bar", "Ident~foo",
		"Ident^foo", "Ident+bar", "Ident/^foo_/", "#foo_bar", "@Purity",
		"(Call Ident=foo_bar $..)", "(Call $.. Ident~foo)",
	} {
		pattern, err := CompilePattern(text)
		if err != nil {
			t.Fatalf("compile %q: %v", text, err)
		}
		matcher := NewMatcher(pattern)
		target := call
		if strings.HasPrefix(text, "Ident") || strings.HasPrefix(text, "#") {
			target = ident
		}
		if !matcher.Matches(target) {
			t.Errorf("%q did not match its target", text)
		}
		if len(matcher.Match(nil)) != 0 {
			t.Errorf("%q matched a nil node", text)
		}
	}
	for _, text := range []string{"Ident=other", "Ident~zzz", "Ident^zz", "Ident+zz",
		"StrLit=other", "@Mutability", "(Call Ident=other)", "(Return !ANY)"} {
		pattern, err := CompilePattern(text)
		if err != nil {
			t.Fatalf("compile %q: %v", text, err)
		}
		if NewMatcher(pattern).Matches(call) {
			t.Errorf("%q unexpectedly matched the call", text)
		}
	}
	if vertexMatches(call, &patternNode{attrOp: "unknown"}) {
		t.Fatal("unknown attribute operator must never match")
	}
	if patternCaptureNames(nil) == nil {
		t.Fatal("nil pattern tree must still return a slice")
	}
	collectCaptureNames(nil, map[string]struct{}{})
	collectFeatures(nil, map[string]struct{}{})
	if hasEllipsis(nil) || childCandidates(call, nil) == nil {
		t.Fatal("child candidate helpers misbehave")
	}
}

// TestCompiledRuleCheck covers check ordering, path mapping and templates.
func TestCompiledRuleCheck(t *testing.T) {
	t.Parallel()
	unit, _, _ := testTree()
	spec := validSpec("GO-EG-001")
	spec.Pattern = `(Call $name $..)`
	spec.Message = "bad identifier $name"
	spec.Where = []string{"capture:name"}
	compiled, err := Compile(spec)
	if err != nil {
		t.Fatalf("Compile: %v", err)
	}
	if compiled.Spec().ID != spec.ID {
		t.Fatal("Spec() lost the declaration")
	}
	if compiled.Check(nil) != nil {
		t.Fatal("Check(nil) must return nil")
	}
	findings := compiled.Check(unit)
	if len(findings) == 0 {
		t.Fatal("compiled rule produced no findings")
	}
	if findings[0].Path != "p.go" {
		t.Fatalf("finding path = %q, want p.go", findings[0].Path)
	}
	if !strings.Contains(findings[0].Message, "foo_bar") ||
		strings.Contains(findings[0].Message, "$name") {
		t.Fatalf("message template was not expanded: %q", findings[0].Message)
	}
	filtered := spec
	filtered.Where = []string{"capture:missing"}
	filteredRule, err := Compile(filtered)
	if err != nil {
		t.Fatalf("Compile filtered: %v", err)
	}
	if len(filteredRule.Check(unit)) != 0 {
		t.Fatal("a false predicate must suppress every finding")
	}
	badWhere := spec
	badWhere.Where = []string{"unknown:x"}
	if _, err := Compile(badWhere); !errors.Is(err, ErrInvalidPredicate) {
		t.Fatalf("Compile with a bad where clause = %v", err)
	}
}
