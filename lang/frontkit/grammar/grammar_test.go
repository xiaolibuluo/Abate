/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package grammar

import (
	"errors"
	"slices"
	"strings"
	"testing"

	"abate/lang/air"
)

// TestLoadEveryEmbeddedLanguage proves the embedded upstream data decodes for
// every supported language and carries a complete node and rule inventory.
func TestLoadEveryEmbeddedLanguage(t *testing.T) {
	t.Parallel()
	tags := SupportedLanguages()
	if len(tags) != 6 {
		t.Fatalf("expected six embedded grammars, got %d", len(tags))
	}
	if !slices.IsSorted(tags) {
		t.Fatalf("supported languages are not sorted: %v", tags)
	}
	for _, tag := range tags {
		g, err := Load(tag)
		if err != nil {
			t.Fatalf("Load(%q): %v", tag, err)
		}
		if g.Language != tag || g.Version == "" || g.Source == "" {
			t.Fatalf("Load(%q): incomplete provenance %+v", tag, g)
		}
		if len(g.Nodes) == 0 || len(g.Rules) == 0 {
			t.Fatalf("Load(%q): empty inventory", tag)
		}
		if len(g.Constructs()) < 50 {
			t.Fatalf("Load(%q): only %d concrete constructs", tag,
				len(g.Constructs()))
		}
		if !slices.IsSorted(g.FieldNames()) {
			t.Fatalf("Load(%q): field inventory is not sorted", tag)
		}
	}
}

// TestConstructsExcludeSupertypes proves the concrete inventory never mixes
// an abstract supertype into the construct list.
func TestConstructsExcludeSupertypes(t *testing.T) {
	t.Parallel()
	for _, tag := range SupportedLanguages() {
		g, err := Load(tag)
		if err != nil {
			t.Fatalf("Load(%q): %v", tag, err)
		}
		supertypes := make(map[string]struct{}, len(g.Supertypes))
		for _, name := range g.Supertypes {
			supertypes[name] = struct{}{}
		}
		for _, node := range g.Constructs() {
			if _, ok := supertypes[node.Name]; ok {
				t.Fatalf("%s: supertype %q leaked into the constructs", tag, node.Name)
			}
			if !node.Named {
				t.Fatalf("%s: anonymous node %q leaked into the constructs",
					tag, node.Name)
			}
		}
	}
}

// TestDefaultRulesCoverEveryConstruct proves the mapping table is total.
//
// A construct no rule selects would be a silent hole in the coverage report,
// so the table must resolve every concrete construct of every language.
func TestDefaultRulesCoverEveryConstruct(t *testing.T) {
	t.Parallel()
	for _, tag := range SupportedLanguages() {
		g, err := Load(tag)
		if err != nil {
			t.Fatalf("Load(%q): %v", tag, err)
		}
		mappings := Map(g, DefaultRules(tag))
		missing := Unmapped(mappings)
		if len(missing) != 0 {
			t.Fatalf("%s: %d unmapped constructs: %s",
				tag, len(missing), strings.Join(missing, ", "))
		}
		if len(mappings) != len(g.Constructs()) {
			t.Fatalf("%s: %d mappings for %d constructs",
				tag, len(mappings), len(g.Constructs()))
		}
	}
}

// TestMappingIsDeterministic proves the same grammar and table always yield
// the same mapping, in the same order, with the same verdicts.
func TestMappingIsDeterministic(t *testing.T) {
	t.Parallel()
	for _, tag := range SupportedLanguages() {
		g, err := Load(tag)
		if err != nil {
			t.Fatalf("Load(%q): %v", tag, err)
		}
		rules := DefaultRules(tag)
		first := Map(g, rules)
		second := Map(g, rules)
		if !slices.Equal(first, second) {
			t.Fatalf("%s: mapping is not reproducible", tag)
		}
		if Summary(first) != Summary(second) {
			t.Fatalf("%s: summary is not reproducible", tag)
		}
		for _, mapping := range first {
			if mapping.Status == StatusUnmapped {
				t.Fatalf("%s: unmapped construct %q", tag, mapping.Node)
			}
			if mapping.Status == StatusOutOfScope {
				continue
			}
			if !mapping.Kind.Valid() {
				t.Fatalf("%s: %q maps onto an invalid kind", tag, mapping.Node)
			}
			if mapping.Rule == "" || mapping.Reason == "" {
				t.Fatalf("%s: %q lacks a rule or a reason", tag, mapping.Node)
			}
		}
	}
}

// TestParseRejectsMalformedDocuments proves the decoder classifies bad input
// instead of returning a half-built fact set.
func TestParseRejectsMalformedDocuments(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name      string
		nodeTypes string
		grammar   string
	}{
		{"empty", "", ""},
		{"broken-node-types", "{", "{}"},
		{"broken-grammar", "[]", "{"},
		{"grammar-without-rules", "[]", `{"name":"x"}`},
		{"node-type-without-name", `[{"named":true}]`, `{"name":"x","rules":{"a":{"type":"SYMBOL","name":"b"}}}`},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := Parse("x", "1", "s", []byte(tc.nodeTypes), []byte(tc.grammar))
			if !errors.Is(err, ErrMalformedGrammar) {
				t.Fatalf("expected a malformed grammar error, got %v", err)
			}
		})
	}
}

// TestLoadRejectsUnknownLanguage proves an unsupported tag is refused.
func TestLoadRejectsUnknownLanguage(t *testing.T) {
	t.Parallel()
	if _, err := Load("cobol"); !errors.Is(err, ErrUnknownLanguage) {
		t.Fatalf("expected an unknown language error, got %v", err)
	}
	if _, err := LoadFromFS("cobol", nil); !errors.Is(err, ErrUnknownLanguage) {
		t.Fatalf("expected an unknown language error, got %v", err)
	}
}

// TestParseBuildsRuleMembers proves the production members are collected in
// source order and that the referenced symbols are the upstream rules.
func TestParseBuildsRuleMembers(t *testing.T) {
	t.Parallel()
	nodeTypes := `[{"type":"call","named":true,"fields":{"function":{"multiple":false,"required":true,"types":[{"type":"identifier","named":true}]}}}]`
	grammar := `{"name":"demo","rules":{"call":{"type":"SYMBOL","name":"identifier"},` +
		`"identifier":{"type":"CHOICE","members":[{"type":"SYMBOL","name":"a"},` +
		`{"type":"SEQ","members":[{"type":"SYMBOL","name":"b"},{"type":"STRING","value":"("}]}]}}}`
	g, err := Parse("demo", "1", "test", []byte(nodeTypes), []byte(grammar))
	if err != nil {
		t.Fatalf("Parse: %v", err)
	}
	if len(g.Rules) != 2 {
		t.Fatalf("expected two rules, got %d", len(g.Rules))
	}
	rule, ok := g.Production("identifier")
	if !ok {
		t.Fatal("production identifier is missing")
	}
	if rule.Kind != "CHOICE" {
		t.Fatalf("unexpected production kind %q", rule.Kind)
	}
	if !slices.Equal(rule.Members, []string{"a", "b"}) {
		t.Fatalf("unexpected members %v", rule.Members)
	}
	if !rule.References("a") || rule.References("missing") {
		t.Fatalf("reference lookup is wrong: %v", rule.Members)
	}
	node, ok := g.Node("call")
	if !ok {
		t.Fatal("node call is missing")
	}
	if !node.HasField("function") || node.HasField("missing") {
		t.Fatalf("unexpected fields %v", node.Fields)
	}
	if !node.Concrete() {
		t.Fatal("a named node without subtypes must stay concrete")
	}
}

// TestKindsAndStatusFilters proves the mapping helpers filter and sort.
func TestKindsAndStatusFilters(t *testing.T) {
	t.Parallel()
	mappings := []Mapping{
		{Node: "a", Kind: 3, Status: StatusExact},
		{Node: "b", Kind: 1, Status: StatusDegraded},
		{Node: "c", Kind: 3, Status: StatusOutOfScope},
		{Node: "d", Kind: 0, Status: StatusUnmapped},
	}
	kinds := Kinds(mappings)
	if !slices.Equal(kinds, []air.NodeKind{0, 1, 3}) {
		t.Fatalf("unexpected kinds %v", kinds)
	}
	if got := len(ByStatus(mappings, StatusExact)); got != 1 {
		t.Fatalf("unexpected exact count %d", got)
	}
	if got := len(Unmapped(mappings)); got != 1 || got == 0 {
		t.Fatalf("unexpected unmapped count %d", got)
	}
}
