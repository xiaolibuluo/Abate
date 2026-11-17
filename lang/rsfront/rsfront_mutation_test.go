/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package rsfront

import (
	"testing"

	"abate/internal/determinism"
	"abate/lang/air"
	"abate/lang/frontkit"
)

// rsKeywordMutSpec removes one keyword from the frozen table; mutation tests
// use it to prove the lexer tables really drive the reduction.
type rsKeywordMutSpec struct{ rsLexSpec }

// Keywords returns the Rust keyword table without the fn keyword.
func (rsKeywordMutSpec) Keywords() []string {
	out := make([]string, 0, len(rsKeywords))
	for _, keyword := range rsKeywords {
		if keyword != "fn" {
			out = append(out, keyword)
		}
	}
	return out
}

// rsMutatedTokens tokenizes src with a deliberately altered lexical table.
func rsMutatedTokens(t *testing.T, src string,
	spec frontkit.LexSpec) []frontkit.Token {
	t.Helper()
	lexer, err := frontkit.NewLexer(spec, []byte(src))
	if err != nil {
		t.Fatalf("NewLexer: %v", err)
	}
	tokens, err := lexer.Tokenize()
	if err != nil {
		t.Fatalf("Tokenize: %v", err)
	}
	return tokens
}

// rsReduceRoot builds one AIR root from tokens with injected tables.
func rsReduceRoot(tokens []frontkit.Token, cfg rsConfig) air.Node {
	return newRsParser(tokens, "mut.rs", cfg).parseFile()
}

// rsReduce hashes one reduced root so mutations can be compared exactly.
func rsReduce(tokens []frontkit.Token, cfg rsConfig) [32]byte {
	return determinism.HashState(rsReduceRoot(tokens, cfg))
}

// TestRsMutationMacroTable proves the injected macro table controls SemTag.
func TestRsMutationMacroTable(t *testing.T) {
	src := "println!(\"hi\");\n"
	tokens := rsTokenize(t, src)
	base := defaultRsConfig()
	mutated := defaultRsConfig()
	mutated.macroNames = make([]string, 0, len(rsMacroNames))
	for _, name := range rsMacroNames {
		if name != "println" {
			mutated.macroNames = append(mutated.macroNames, name)
		}
	}
	if !rsHasSem(rsReduceRoot(tokens, base), air.SemTag, "known-macro") {
		t.Fatal("base macro table did not tag the built-in macro")
	}
	if rsHasSem(rsReduceRoot(tokens, mutated), air.SemTag, "known-macro") {
		t.Fatal("mutated macro table still tags the removed macro")
	}
	if rsReduce(tokens, base) == rsReduce(tokens, mutated) {
		t.Fatal("macro table mutation left the AIR unchanged")
	}
}

// TestRsMutationTryOpNote proves the try (? ) mapping is table-driven.
func TestRsMutationTryOpNote(t *testing.T) {
	src := "let v = foo()?;\n"
	tokens := rsTokenize(t, src)
	base := defaultRsConfig()
	if !rsHasNote(rsReduceRoot(tokens, base), "try-op") {
		t.Fatal("base try-op note is missing")
	}
	mutated := defaultRsConfig()
	mutated.notes = append([]rsNote(nil), rsNotes...)
	for i := range mutated.notes {
		if mutated.notes[i].name == "try-op" {
			mutated.notes[i].value = "MUTATED-TRY-OP"
		}
	}
	root := rsReduceRoot(tokens, mutated)
	if !rsHasNote(root, "MUTATED-TRY-OP") {
		t.Fatal("mutated try-op note is absent from AIR")
	}
	if rsHasNote(root, "try-op") {
		t.Fatal("original try-op note survived the mutation")
	}
	if rsReduce(tokens, base) == determinism.HashState(root) {
		t.Fatal("try-op note mutation left the AIR hash unchanged")
	}
}

// TestRsMutationInfixBindings proves deleting or rewriting a binding changes
// the expression tree instead of leaving a constant output.
func TestRsMutationInfixBindings(t *testing.T) {
	src := "let x = 1 + 2 * 3;\n"
	tokens := rsTokenize(t, src)
	base := defaultRsConfig()
	without := defaultRsConfig()
	without.infix = make([]rsBinding, 0, len(rsInfixBindings))
	for _, binding := range rsInfixBindings {
		if binding.text != "+" {
			without.infix = append(without.infix, binding)
		}
	}
	if rsReduce(tokens, base) == rsReduce(tokens, without) {
		t.Fatal("removing the + binding left the AIR unchanged")
	}
	flattened := defaultRsConfig()
	flattened.infix = make([]rsBinding, 0, len(rsInfixBindings))
	for _, binding := range rsInfixBindings {
		copyBinding := binding
		copyBinding.power = 1
		flattened.infix = append(flattened.infix, copyBinding)
	}
	if rsReduce(tokens, base) == rsReduce(tokens, flattened) {
		t.Fatal("changing every binding power left the AIR unchanged")
	}
}

// TestRsMutationNoteTable proves semantic notes come from the injected table.
func TestRsMutationNoteTable(t *testing.T) {
	src := "let a = [1, 2];\n"
	tokens := rsTokenize(t, src)
	base := defaultRsConfig()
	mutated := defaultRsConfig()
	mutated.notes = append([]rsNote(nil), rsNotes...)
	for i := range mutated.notes {
		if mutated.notes[i].name == "array" {
			mutated.notes[i].value = "MUTATED-ARRAY-NOTE"
		}
	}
	root := rsReduceRoot(tokens, mutated)
	if !rsHasNote(root, "MUTATED-ARRAY-NOTE") {
		t.Fatal("mutated array note is absent from AIR")
	}
	if rsHasNote(root, "array") {
		t.Fatal("original array note survived the mutation")
	}
	if rsReduce(tokens, base) == determinism.HashState(root) {
		t.Fatal("note table mutation left the AIR hash unchanged")
	}
}

// TestRsMutationItemRuleTable proves the item rule table drives node kinds.
func TestRsMutationItemRuleTable(t *testing.T) {
	src := "fn f() {}\n"
	tokens := rsTokenize(t, src)
	base := defaultRsConfig()
	if !rsHasKind(rsReduceRoot(tokens, base), air.KindFuncDecl) {
		t.Fatal("base item rule table lost KindFuncDecl")
	}
	mutated := defaultRsConfig()
	mutated.itemRules = make([]rsItemRule, 0, len(rsItemRules))
	for _, rule := range rsItemRules {
		if rule.keyword != "fn" {
			mutated.itemRules = append(mutated.itemRules, rule)
		}
	}
	if rsHasKind(rsReduceRoot(tokens, mutated), air.KindFuncDecl) {
		t.Fatal("removed item rule still produced KindFuncDecl")
	}
	if rsReduce(tokens, base) == rsReduce(tokens, mutated) {
		t.Fatal("item rule mutation left the AIR unchanged")
	}
}

// TestRsMutationKeywordTable proves the keyword table changes token kinds and
// therefore the reduced AIR tree.
func TestRsMutationKeywordTable(t *testing.T) {
	src := "fn f(x: i32) -> i32 { x }\n"
	baseTokens := rsTokenize(t, src)
	mutTokens := rsMutatedTokens(t, src, rsKeywordMutSpec{})
	baseKinds, mutKinds := 0, 0
	for i := range baseTokens {
		if baseTokens[i].Text == "fn" && baseTokens[i].Kind == frontkit.TokenKeyword {
			baseKinds++
		}
		if i < len(mutTokens) && mutTokens[i].Text == "fn" &&
			mutTokens[i].Kind == frontkit.TokenIdent {
			mutKinds++
		}
	}
	if baseKinds != 1 || mutKinds != 1 {
		t.Fatalf("keyword mutation did not change token kinds: %d/%d",
			baseKinds, mutKinds)
	}
	cfg := defaultRsConfig()
	if rsReduce(baseTokens, cfg) == rsReduce(mutTokens, cfg) {
		t.Fatal("keyword table mutation left the AIR unchanged")
	}
}

// TestRsMutationDialectTable proves dialect acceptance is table-driven.
func TestRsMutationDialectTable(t *testing.T) {
	if !rsContains(rsDialects, "2015") || !rsContains(rsDialects, "2024") {
		t.Fatal("canonical Rust editions absent from table")
	}
	if rsContains(rsDialects, "2022") {
		t.Fatal("unsupported edition accepted by table")
	}
	mutated := []string{"2015"}
	if !rsContains(mutated, "2015") || rsContains(mutated, "2021") {
		t.Fatal("mutated edition table behaves incorrectly")
	}
	if _, err := defaultRsSpec().NormalizeDialect("2015"); err != nil {
		t.Fatalf("2015 must be accepted: %v", err)
	}
	if _, err := defaultRsSpec().NormalizeDialect("2015-experimental"); err == nil {
		t.Fatal("non-canonical edition accepted")
	}
}

// TestRsMutationDepthCeiling proves the injected depth ceiling caps recursion.
func TestRsMutationDepthCeiling(t *testing.T) {
	src := "((((((((((1))))))))))\n"
	tokens := rsTokenize(t, src)
	base := defaultRsConfig()
	shallow := defaultRsConfig()
	shallow.maxDepth = 2
	if rsReduce(tokens, base) == rsReduce(tokens, shallow) {
		t.Fatal("depth ceiling mutation left the AIR unchanged")
	}
	if !rsHasSem(rsReduceRoot(tokens, shallow), air.SemPartial, "true") {
		t.Fatal("shallow depth ceiling did not mark partial reductions")
	}
}
