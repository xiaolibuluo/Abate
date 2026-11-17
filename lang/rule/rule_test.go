/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package rule

import (
	"errors"
	"fmt"
	"slices"
	"strings"
	"testing"

	"abate/lang/air"
)

// TestPatternFeatures proves all five structural features compile and match.
func TestPatternFeatures(t *testing.T) {
	t.Parallel()
	cases := []struct {
		text    string
		feature string
		hit     bool
		onRoot  bool
	}{
		{"$_", "wildcard", true, false}, {"(Call $_)", "wildcard", true, false},
		{"(Call $_ $_)", "wildcard", true, false},
		{"$x", "capture", true, false}, {"(Call $x)", "capture", true, false},
		{"(Call $x $y)", "capture", true, false},
		{"(Call $..)", "deep", true, false},
		{"(Package $.. Call)", "deep", true, true},
		{"(Call $x $.. $y)", "deep", true, false},
		{"(Call !Ident=fmt)", "negative", true, false},
		{"(Call $x !StrLit)", "negative", false, false},
		{"(Return !ANY)", "negative", false, false},
		{"(Call $x*)", "repeat", true, false},
		{"(Call $x $y*)", "repeat", true, false},
		{"(Call $x* $..)", "repeat", true, false},
	}
	_, root, call := testTree()
	for _, c := range cases {
		pattern, err := CompilePattern(c.text)
		if err != nil {
			t.Errorf("CompilePattern(%q): %v", c.text, err)
			continue
		}
		if !slices.Contains(pattern.Features(), c.feature) {
			t.Errorf("%q lost feature %s", c.text, c.feature)
		}
		if !pattern.Valid() || pattern.Source() != c.text ||
			!strings.Contains(pattern.String(), c.text) {
			t.Errorf("%q has inconsistent accessors", c.text)
		}
		target := call
		if c.onRoot {
			target = root
		}
		if NewMatcher(pattern).Matches(target) != c.hit {
			t.Errorf("%q hit = %v, want %v", c.text,
				NewMatcher(pattern).Matches(target), c.hit)
		}
		if NewMatcher(pattern).Matches(root) && c.text != "$x" &&
			c.text != "$_" && c.text != "(Package $.. Call)" {
			t.Errorf("%q unexpectedly matched the package root", c.text)
		}
	}
}

// TestPatternRejections covers malformed pattern sources and metavariables.
func TestPatternRejections(t *testing.T) {
	t.Parallel()
	deepText := strings.Repeat("(Call ", 40) + "x" + strings.Repeat(")", 40)
	for _, text := range []string{
		"", "(", ")", "(Call", "(Call))", "( Call", "($..)", "($x*)",
		"Ident=", "Ident/regex", "Ident/^[/", "@Unknown", "#", "$",
		"$1bad", "(Call \"quoted\")", deepText,
	} {
		if _, err := CompilePattern(text); !errors.Is(err, ErrBadPattern) {
			t.Errorf("CompilePattern(%q) = %v, want ErrBadPattern", text, err)
		}
	}
	pattern, err := CompilePattern("(Call $alpha $beta* $.. !Ident=fmt)")
	if err != nil {
		t.Fatalf("rich pattern: %v", err)
	}
	if got := pattern.Metavariables(); !slices.Equal(got, []string{"alpha", "beta"}) {
		t.Fatalf("Metavariables = %v", got)
	}
	if _, err := CompilePattern("@Operator"); err != nil {
		t.Fatalf("semantic pattern: %v", err)
	}
	if (Pattern{}).Valid() || len(Pattern{}.Metavariables()) != 0 ||
		len(Pattern{}.Features()) != 0 {
		t.Fatal("zero pattern must report itself invalid")
	}
}

// TestPredicates covers every predicate family and its rejections.
func TestPredicates(t *testing.T) {
	t.Parallel()
	_, _, call := testTree()
	caps := []Capture{
		{Name: "x", Kind: "Ident", Text: "foo_bar"},
		{Name: "y", Kind: "StrLit", Text: "12"},
	}
	trueClauses := []string{
		"kind:Call", "kind-not:Return", "text-not:zzz", "text~:foo",
		"text^:foo", "text+:bar", "regex:^[a-z_]+$", "num-min:10",
		"num-max:99", "count-min:1", "count-max:4", "capture:x", "distinct",
	}
	for _, clause := range trueClauses {
		predicate, err := ParsePredicate(clause)
		if err != nil {
			t.Errorf("ParsePredicate(%q): %v", clause, err)
			continue
		}
		if !predicate.Valid() || predicate.String() != clause {
			t.Errorf("predicate %q lost its form", clause)
		}
		if !predicate.Evaluate(call, caps) {
			t.Errorf("predicate %q evaluated false", clause)
		}
	}
	falseClauses := []string{"kind:Return", "kind-not:Call", "text:zzz",
		"text~:zzz", "text^:zzz", "text+:zzz", "regex:^ZZZ+$",
		"num-min:99999", "num-max:0", "count-min:99", "count-max:0",
		"capture:missing", "sem:Mutability"}
	for _, clause := range falseClauses {
		predicate, err := ParsePredicate(clause)
		if err != nil {
			t.Errorf("ParsePredicate(%q): %v", clause, err)
			continue
		}
		if predicate.Evaluate(call, caps) {
			t.Errorf("predicate %q evaluated true", clause)
		}
		if predicate.Evaluate(nil, caps) {
			t.Errorf("predicate %q matched a nil node", clause)
		}
	}
	air.Annotate(call, air.SemPurity, "pure")
	annotation, err := ParsePredicate("sem:Purity")
	if err != nil || !annotation.Evaluate(call, caps) {
		t.Fatalf("semantic predicate failed: %v", err)
	}
	for _, bad := range []string{"", "unknown:x", "regex:", "regex:[",
		"num-min:abc", "count-max:abc"} {
		if _, err := ParsePredicate(bad); !errors.Is(err, ErrInvalidPredicate) {
			t.Errorf("ParsePredicate(%q) = %v, want ErrInvalidPredicate", bad, err)
		}
	}
	duplicated := []Capture{{Text: "a"}, {Text: "a"}}
	distinct, err := ParsePredicate("distinct")
	if err != nil || distinct.Evaluate(call, duplicated) {
		t.Fatal("distinct must reject repeated captures")
	}
	if (Predicate{Op: "unknown"}).Evaluate(call, caps) ||
		!capturesDistinct(nil) || hasCapture(nil, "x") {
		t.Fatal("predicate helper contracts are wrong")
	}
	if len(captureTexts(call, nil)) != 1 {
		t.Fatal("captureTexts must fall back to the node attribute")
	}
	badNumber := Predicate{Op: "num-min", Arg: "not-a-number"}
	if badNumber.Evaluate(call, caps) {
		t.Fatal("non-numeric limit must fail")
	}
	badCount := Predicate{Op: "count-min", Arg: "not-a-number"}
	if badCount.Evaluate(call, caps) {
		t.Fatal("non-numeric count limit must fail")
	}
}

// TestScopePredicates proves negative regex clauses exclude a named family
// while path clauses exclude test files without hiding library code.
func TestScopePredicates(t *testing.T) {
	t.Parallel()
	_, _, call := testTree()
	excluded := []Capture{{Name: "x", Kind: "Ident", Text: "abi_test"}}
	ordinary := []Capture{{Name: "x", Kind: "Ident", Text: "foo_bar"}}
	suffix, err := ParsePredicate(`text-not-regex:_test$`)
	if err != nil {
		t.Fatalf("parse text clause: %v", err)
	}
	if suffix.Evaluate(call, excluded) || !suffix.Evaluate(call, ordinary) {
		t.Fatal("text-not-regex must exclude only the _test suffix")
	}
	library, err := ParsePredicate(`path-not-regex:_test\.go$`)
	if err != nil {
		t.Fatalf("parse path clause: %v", err)
	}
	if !library.evaluateWithPath(call, ordinary, "pkg/library.go") ||
		library.evaluateWithPath(call, ordinary, "pkg/library_test.go") {
		t.Fatal("path-not-regex must exclude test files only")
	}
	for _, bad := range []string{"text-not-regex:[", "path-not-regex:["} {
		if _, parseErr := ParsePredicate(bad); !errors.Is(parseErr, ErrInvalidPredicate) {
			t.Errorf("ParsePredicate(%q) = %v, want ErrInvalidPredicate", bad, parseErr)
		}
	}
}

// TestRegistryFourValidations proves each load-time check rejects a bad rule.
func TestRegistryFourValidations(t *testing.T) {
	t.Parallel()
	registry := NewRegistry()
	if registry.Len() != 0 || registry.For("") == nil {
		t.Fatal("new registry must be empty with a non-nil rule slice")
	}
	good := validSpec("GO-EG-001")
	if err := registry.Register(good); err != nil {
		t.Fatalf("Register: %v", err)
	}
	if err := registry.Load([]Spec{validSpec("GO-EG-002")}); err != nil {
		t.Fatalf("Load: %v", err)
	}
	if registry.Len() != 2 {
		t.Fatalf("registry size = %d, want 2", registry.Len())
	}
	if _, ok := registry.Lookup("GO-EG-001"); !ok {
		t.Fatal("Lookup missed a registered rule")
	}
	if _, ok := registry.Lookup("GO-EG-777"); ok {
		t.Fatal("Lookup invented a rule")
	}
	ids := registry.IDs()
	if !slices.Equal(ids, []RuleID{"GO-EG-001", "GO-EG-002"}) {
		t.Fatalf("IDs = %v", ids)
	}
	if got := registry.For("go"); len(got) != 2 {
		t.Fatalf("For(go) returned %d rules", len(got))
	}
	if got := registry.For("python"); len(got) != 0 {
		t.Fatalf("For(python) returned %d rules", len(got))
	}
	rejections := []struct {
		spec Spec
		want error
	}{
		{validSpec("GO-EG-001"), ErrDuplicateID},
		{func() Spec { s := validSpec("GO-EG-004"); s.Reference = ""; return s }(), ErrEmptyReference},
		{func() Spec { s := validSpec("GO-EG-005"); s.Pattern = "("; return s }(), ErrBadPattern},
		{func() Spec { s := validSpec("GO-EG-006"); s.Severity = "bad"; return s }(), ErrInvalidSpec},
		{func() Spec { s := validSpec("GO-EG-007"); s.ID = "bad"; return s }(), ErrInvalidRuleID},
	}
	for _, item := range rejections {
		err := registry.Register(item.spec)
		if !errors.Is(err, item.want) {
			t.Errorf("Register(%s) = %v, want %v", item.spec.ID, err, item.want)
		}
	}
	if err := registry.Load(nil); err != nil {
		t.Fatalf("empty batch load: %v", err)
	}
	batch := []Spec{validSpec("GO-EG-008"), validSpec("GO-EG-008")}
	if err := registry.Load(batch); !errors.Is(err, ErrDuplicateID) {
		t.Fatalf("duplicate inside one batch = %v", err)
	}
	if registry.Len() != 2 {
		t.Fatal("failed batch load mutated the registry")
	}
}

// TestAggregateOrdering proves findings sort by file, line, column and id.
func TestAggregateOrdering(t *testing.T) {
	t.Parallel()
	first := Finding{RuleID: "GO-EG-002", Path: "b.go", Line: 1, Column: 1}
	second := Finding{RuleID: "GO-EG-001", Path: "b.go", Line: 1, Column: 1}
	third := Finding{RuleID: "GO-EG-001", Path: "a.go", Line: 9, Column: 1}
	fourth := Finding{RuleID: "GO-EG-001", Path: "a.go", Line: 2, Column: 8}
	fifth := Finding{RuleID: "GO-EG-001", Path: "a.go", Line: 2, Column: 3}
	ordered := []Finding{first, second, third, fourth, fifth}
	sortFindings(ordered)
	want := []string{"a.go:2:3", "a.go:2:8", "a.go:9:1", "b.go:1:1", "b.go:1:1"}
	for i, item := range ordered {
		got := fmt.Sprintf("%s:%d:%d", item.Path, item.Line, item.Column)
		if got != want[i] {
			t.Fatalf("ordering mismatch at %d: %s vs %s", i, got, want[i])
		}
	}
	if ordered[3].RuleID != "GO-EG-001" || ordered[4].RuleID != "GO-EG-002" {
		t.Fatal("rule identifier is not the final ordering key")
	}
}
