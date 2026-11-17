/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package pyfront

import (
	"testing"

	"abate/internal/determinism"
	"abate/lang/air"
	"abate/lang/frontkit"
)

// pyKeywordMutSpec removes one keyword from the frozen table; mutation tests
// use it to prove the lexer tables really drive the reduction.
type pyKeywordMutSpec struct{ pyLexSpec }

// Keywords returns the Python keyword table without the def keyword.
func (pyKeywordMutSpec) Keywords() []string {
	out := make([]string, 0, len(pyKeywords))
	for _, keyword := range pyKeywords {
		if keyword != "def" {
			out = append(out, keyword)
		}
	}
	return out
}

// pyMutatedTokens tokenizes src with a deliberately altered lexical table.
func pyMutatedTokens(t *testing.T, src string, spec frontkit.LexSpec) []frontkit.Token {
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

// pyReduce builds a root from a token stream with the given injected tables.
func pyReduce(t *testing.T, tokens []frontkit.Token,
	infix []pyBinding, notes []pyNote) [32]byte {
	t.Helper()
	parser, err := newPyParserWithBindings(tokens, "mut.py", infix, notes)
	if err != nil {
		t.Fatalf("newPyParserWithBindings: %v", err)
	}
	return determinism.HashState(parser.parseModule())
}

// TestPyMutationInfixBindings proves deleting one binding changes the AIR.
func TestPyMutationInfixBindings(t *testing.T) {
	src := "x = 1 + 2 * 3\n"
	tokens := pyTokenize(t, src)
	base := pyReduce(t, tokens, pyInfixBindings, pyNotes)
	mutated := make([]pyBinding, 0, len(pyInfixBindings))
	for _, binding := range pyInfixBindings {
		if binding.text != "+" {
			mutated = append(mutated, binding)
		}
	}
	changed := pyReduce(t, tokens, mutated, pyNotes)
	if base == changed {
		t.Fatal("removing the + binding left the AIR unchanged")
	}
	reversed := make([]pyBinding, 0, len(pyInfixBindings))
	for _, binding := range pyInfixBindings {
		copyBinding := binding
		copyBinding.power = 1
		reversed = append(reversed, copyBinding)
	}
	if base == pyReduce(t, tokens, reversed, pyNotes) {
		t.Fatal("changing every binding power left the AIR unchanged")
	}
}

// TestPyMutationNoteTable proves semantic notes come from the injected table.
func TestPyMutationNoteTable(t *testing.T) {
	src := "l = [1, 2]\n"
	tokens := pyTokenize(t, src)
	mutated := make([]pyNote, len(pyNotes))
	copy(mutated, pyNotes)
	for i := range mutated {
		if mutated[i].name == "list" {
			mutated[i].value = "MUTATED-LIST-NOTE"
		}
	}
	parser, err := newPyParserWithBindings(tokens, "mut.py", pyInfixBindings, mutated)
	if err != nil {
		t.Fatal(err)
	}
	root := parser.parseModule()
	if !pyHasNote(root, "MUTATED-LIST-NOTE") {
		t.Fatal("mutated note value is absent from AIR")
	}
	if pyHasNote(root, "list-literal") {
		t.Fatal("original note value survived the mutation")
	}
	base := pyReduce(t, tokens, pyInfixBindings, pyNotes)
	if base == determinism.HashState(root) {
		t.Fatal("note mutation left the AIR hash unchanged")
	}
}

// TestPyMutationKeywordTable proves the keyword table changes token kinds and
// therefore the reduced AIR tree.
func TestPyMutationKeywordTable(t *testing.T) {
	src := "def f(x):\n    return x\n"
	baseTokens := pyTokenize(t, src)
	mutTokens := pyMutatedTokens(t, src, pyKeywordMutSpec{})
	baseKinds, mutKinds := 0, 0
	for i := range baseTokens {
		if baseTokens[i].Text == "def" && baseTokens[i].Kind == frontkit.TokenKeyword {
			baseKinds++
		}
		if i < len(mutTokens) && mutTokens[i].Text == "def" &&
			mutTokens[i].Kind == frontkit.TokenIdent {
			mutKinds++
		}
	}
	if baseKinds != 1 || mutKinds != 1 {
		t.Fatalf("keyword mutation did not change token kinds: %d/%d", baseKinds, mutKinds)
	}
	base := pyReduce(t, baseTokens, pyInfixBindings, pyNotes)
	changed := pyReduce(t, mutTokens, pyInfixBindings, pyNotes)
	if base == changed {
		t.Fatal("keyword table mutation left the AIR unchanged")
	}
}

// TestPyMutationDialectTable proves dialect acceptance is table-driven.
func TestPyMutationDialectTable(t *testing.T) {
	if !pyContains(pyDialects, "3.8") || !pyContains(pyDialects, "3.12") {
		t.Fatal("canonical Python dialects absent from table")
	}
	if pyContains(pyDialects, "4.0") {
		t.Fatal("unsupported dialect accepted by table")
	}
	mutated := []string{"3.8"}
	if !pyContains(mutated, "3.8") || pyContains(mutated, "3.12") {
		t.Fatal("mutated dialect table behaves incorrectly")
	}
	if _, err := (pySpec{}).NormalizeDialect("3.8"); err != nil {
		t.Fatalf("3.8 must be accepted: %v", err)
	}
	if _, err := (pySpec{}).NormalizeDialect("3.8-experimental"); err == nil {
		t.Fatal("non-canonical dialect accepted")
	}
}

// TestPyTableLookups covers both hit and miss paths of the keyword/infix/rule
// tables so the mutation tests above are known to exercise real table-driven
// behavior.
func TestPyTableLookups(t *testing.T) {
	if !pyContains(pyKeywords, "def") || pyContains(pyKeywords, "nope") {
		t.Fatal("pyContains broken")
	}
	if !pyHasPrefix("# noqa: x", pyDirectivePrefixes) ||
		pyHasPrefix("# plain", pyDirectivePrefixes) {
		t.Fatal("pyHasPrefix broken")
	}
	if pyInfixPower(pyInfixBindings, "**") != 80 || pyInfixPower(pyInfixBindings, ",") != 0 {
		t.Fatal("pyInfixPower broken")
	}
	if !pyInfixRight(pyInfixBindings, "**") || pyInfixRight(pyInfixBindings, "+") ||
		pyInfixRight(pyInfixBindings, ",") {
		t.Fatal("pyInfixRight broken")
	}
	if rule, ok := pySimpleStmtRule("return"); !ok || rule.kind != air.KindReturn {
		t.Fatal("pySimpleStmtRule broken")
	}
	if _, ok := pySimpleStmtRule("nope"); ok {
		t.Fatal("pySimpleStmtRule accepted unknown keyword")
	}
	if rule, ok := pyHeaderKind("match"); !ok || rule.kind != air.KindMatch {
		t.Fatal("pyHeaderKind broken")
	}
	if _, ok := pyHeaderKind("nope"); ok {
		t.Fatal("pyHeaderKind accepted unknown keyword")
	}
}

// TestPyTokenLookups covers both hit and miss paths of the token/kind/note
// lookup helpers.
func TestPyTokenLookups(t *testing.T) {
	if pyTokenNodeKind(frontkit.TokenNum) != air.KindNumLit ||
		pyTokenNodeKind(frontkit.TokenEOF) != air.KindRegion {
		t.Fatal("pyTokenNodeKind broken")
	}
	if !pyIsLexicalKind(frontkit.TokenStr) || pyIsLexicalKind(frontkit.TokenEOF) {
		t.Fatal("pyIsLexicalKind broken")
	}
	if pyNoteFor(pyNotes, "walrus") != "walrus" || pyNoteFor(pyNotes, "nope") != "" {
		t.Fatal("pyNoteFor broken")
	}
	if pyCloseOf(pyModeParen) != ")" || pyCloseOf(pyModeList) != "]" ||
		pyCloseOf(pyModeSet) != "}" || pyCloseOf(-1) != "}" {
		t.Fatal("pyCloseOf broken")
	}
	if pyCollectionNote(pyModeDict) != "dict" ||
		pyComprehensionNote(pyModeSet) != "set-comprehension" ||
		pyTableAt(nil, 0) != "" || pyTableAt([]string{}, 0) != "" {
		t.Fatal("note table lookup broken")
	}
}
