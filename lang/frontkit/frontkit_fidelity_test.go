/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package frontkit

import (
	"errors"
	"slices"
	"testing"

	"abate/internal/determinism"
	"abate/lang/air"
)

func fakeSourceUnit(files ...air.SourceFile) air.SourceUnit {
	return air.SourceUnit{Language: "go", Dialect: "FAKE1", Files: files}
}

func mustFrontend(t *testing.T, spec LanguageSpec) *FrontendBase {
	t.Helper()
	frontend, err := NewFrontendBase(spec)
	if err != nil {
		t.Fatal(err)
	}
	return frontend
}

func fixtureFiles() []air.SourceFile {
	return []air.SourceFile{
		air.NewSourceFile("b.fk", []byte("let b = 2\n")),
		air.NewSourceFile("a.fk", []byte(lexerFixture)),
	}
}

func unitWithRoot(t *testing.T, src string, root air.Node) *air.Unit {
	t.Helper()
	unit := air.NewUnit(air.NewPackageRef("fake", nil))
	if err := unit.AddFile(air.NewFileRef("a.fk", []byte(src)), root); err != nil {
		t.Fatal(err)
	}
	return unit
}

func TestFidelityBidirectional(t *testing.T) {
	t.Parallel()
	tokens := mustLex(t, newFakeSpec(), lexerFixture)
	root := fakeReduce(tokens, "a.fk")
	unit := unitWithRoot(t, lexerFixture, root)
	if err := VerifyTokenFidelity(tokens, unit); err != nil {
		t.Fatalf("source to AIR fidelity: %v", err)
	}
	if err := VerifyNodeFidelity(tokens, root); err != nil {
		t.Fatalf("node-level fidelity: %v", err)
	}
	if err := air.CheckIntervals(root, len(lexerFixture)); err != nil {
		t.Fatalf("AIR intervals: %v", err)
	}
}

func TestFidelityRejectsMissingAndInventedTokens(t *testing.T) {
	t.Parallel()
	tokens := mustLex(t, newFakeSpec(), lexerFixture)
	root := fakeReduce(tokens, "a.fk")
	children := root.Children()
	if len(children) < 2 {
		t.Fatal("fixture produced too few lexical nodes")
	}
	missing := air.NewNodeAt(air.KindPackage, root.Pos(), root.End(), "", children[1:]...)
	if err := VerifyTokenFidelity(tokens,
		unitWithRoot(t, lexerFixture, missing)); !errors.Is(err, ErrTokenFidelity) {
		t.Fatalf("a dropped token was accepted: %v", err)
	}
	invented := leafNode("ghost", 3, 7)
	rootWithGhost := air.NewNodeAt(air.KindPackage, root.Pos(), root.End(), "",
		append(slices.Clone(children), invented)...)
	if err := VerifyTokenFidelity(tokens,
		unitWithRoot(t, lexerFixture, rootWithGhost)); !errors.Is(err, ErrTokenFidelity) {
		t.Fatalf("an invented token was accepted: %v", err)
	}
}

func TestFidelityRejectsEmptyUnitAndBadTokens(t *testing.T) {
	t.Parallel()
	tokens := mustLex(t, newFakeSpec(), lexerFixture)
	empty := air.NewUnit(air.NewPackageRef("x", nil))
	if err := VerifyTokenFidelity(tokens, empty); !errors.Is(err, ErrTokenFidelity) {
		t.Fatalf("empty unit was accepted: %v", err)
	}
	if err := VerifyNodeFidelity(tokens, nil); !errors.Is(err, ErrTokenFidelity) {
		t.Fatalf("nil root was accepted: %v", err)
	}
	invalid := []Token{{Kind: TokenIdent, Text: "x", Offset: 0, EndOffset: 1}}
	if err := VerifyTokenFidelity(invalid, empty); !errors.Is(err, ErrUnknownToken) {
		t.Fatalf("invalid token error = %v, want ErrUnknownToken", err)
	}
	oversized := make([]Token, maxFileTokens+1)
	if err := VerifyTokenFidelity(oversized, empty); !errors.Is(err, ErrLexFailed) {
		t.Fatalf("oversized stream error = %v, want ErrLexFailed", err)
	}
	for _, kind := range []TokenKind{
		TokenIdent, TokenNum, TokenStr, TokenChar, TokenComment,
	} {
		if !isLexicalToken(kind) {
			t.Fatalf("%s must be lexical", kind.String())
		}
	}
	for _, kind := range []TokenKind{TokenKeyword, TokenOp, TokenPunct, TokenEOF} {
		if isLexicalToken(kind) {
			t.Fatalf("%s must not be lexical", kind.String())
		}
	}
	if !isLexicalNode(air.KindComment) || isLexicalNode(air.KindBinary) {
		t.Fatal("AIR lexical classification drifted")
	}
}

func TestFrontendBaseParseAndInterface(t *testing.T) {
	t.Parallel()
	frontend := mustFrontend(t, newFakeSpec())
	var port air.Frontend = frontend
	unit, diagnostics, err := port.Parse(fakeSourceUnit(fixtureFiles()...))
	if err != nil {
		t.Fatal(err)
	}
	if len(diagnostics) != 0 {
		t.Fatalf("unexpected diagnostics: %v", diagnostics)
	}
	if err := unit.Validate(); err != nil {
		t.Fatalf("unit validation: %v", err)
	}
	if len(unit.Files) != 2 || unit.Files[0].Path != "a.fk" ||
		unit.Files[1].Path != "b.fk" {
		t.Fatalf("files not path-sorted: %v", unit.Files)
	}
	if unit.Annotations["lang.go.dialect"] != "fake1" ||
		unit.Annotations["lang.go.files"] != "2" ||
		unit.Annotations["lang.go.diagnostics"] != "0" {
		t.Fatalf("annotations = %v", unit.Annotations)
	}
	if err := air.CheckIntervals(unit.Roots[0], len(lexerFixture)); err != nil {
		t.Fatalf("AIR intervals: %v", err)
	}
	single := air.NewUnit(air.NewPackageRef("fake", nil))
	if err := single.AddFile(unit.Files[0], unit.Roots[0]); err != nil {
		t.Fatal(err)
	}
	tokens := mustLex(t, newFakeSpec(), lexerFixture)
	if err := VerifyTokenFidelity(tokens, single); err != nil {
		t.Fatalf("end-to-end fidelity: %v", err)
	}
	reversed := slices.Clone(fixtureFiles())
	slices.Reverse(reversed)
	other, _, err := frontend.Parse(fakeSourceUnit(reversed...))
	if err != nil {
		t.Fatal(err)
	}
	if err := air.CompareUnits(unit, other); err != nil {
		t.Fatalf("file order leaked into AIR: %v", err)
	}
}

func TestFrontendBaseContractErrors(t *testing.T) {
	t.Parallel()
	if _, err := NewFrontendBase(nil); !errors.Is(err, ErrNilSpec) {
		t.Fatalf("nil spec error = %v", err)
	}
	badDialect := air.SourceUnit{
		Language: "go", Dialect: "broken!",
		Files: []air.SourceFile{air.NewSourceFile("a.fk", []byte("x"))},
	}
	if _, _, err := mustFrontend(t, newFakeSpec()).Parse(badDialect); !errors.Is(err, ErrBadDialect) {
		t.Fatalf("dialect error = %v, want ErrBadDialect", err)
	}
	unknown := newFakeSpec()
	unknown.tag = "ruby"
	if _, _, err := mustFrontend(t, unknown).Parse(fakeSourceUnit(fixtureFiles()...)); !errors.Is(err, ErrUnsupportedLanguage) {
		t.Fatalf("tag error = %v, want ErrUnsupportedLanguage", err)
	}
	failing := newFakeSpec()
	failing.failParse = true
	unit, diagnostics, err := mustFrontend(t, failing).Parse(
		fakeSourceUnit(air.NewSourceFile("a.fk", []byte("x"))))
	if err != nil {
		t.Fatal(err)
	}
	if len(diagnostics) == 0 || !diagnostics[0].IsError() {
		t.Fatalf("parse failure produced no error diagnostic: %v", diagnostics)
	}
	if err := unit.Validate(); err != nil {
		t.Fatal(err)
	}
	if unit.Roots[0].Sem(air.SemPartial) != "true" {
		t.Fatal("failed reduction is not honestly marked partial")
	}
	limited := mustFrontend(t, newFakeSpec()).WithMaxDiagnostics(1)
	if limited == nil || limited.maxDiagnostics != 1 {
		t.Fatal("WithMaxDiagnostics did not apply the cap")
	}
	duplicate := air.SourceUnit{
		Language: "go", Files: []air.SourceFile{
			air.NewSourceFile("a.fk", []byte("x")),
			air.NewSourceFile("a.fk", []byte("y")),
		},
	}
	if _, _, err := mustFrontend(t, newFakeSpec()).Parse(duplicate); !errors.Is(err, air.ErrDuplicatePath) {
		t.Fatalf("duplicate path error = %v", err)
	}
}

func TestIncrementalReuseAndHonestFullRebuild(t *testing.T) {
	t.Parallel()
	frontend := mustFrontend(t, newFakeSpec())
	files := fixtureFiles()
	old, _, err := frontend.Parse(fakeSourceUnit(files...))
	if err != nil {
		t.Fatal(err)
	}
	warm, warmDiff, err := frontend.ParseIncremental(fakeSourceUnit(files...), old, nil)
	if err != nil {
		t.Fatal(err)
	}
	if warmDiff.FullRebuild || !warmDiff.Empty() ||
		warm.NodeCount() != old.NodeCount() ||
		warm.Roots[0] != old.Roots[0] || warm.Roots[1] != old.Roots[1] {
		t.Fatalf("unchanged unit was not reused verbatim: %s", warmDiff.String())
	}
	if warm.Annotations["lang.go.incremental"] != "file-level-reuse" ||
		warm.Annotations["lang.go.reusedFiles"] != "2" {
		t.Fatalf("reuse annotations = %v", warm.Annotations)
	}
	cold, coldDiff, err := frontend.ParseIncremental(fakeSourceUnit(files...), nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	if !coldDiff.FullRebuild || cold.Annotations["lang.go.incremental"] != "full-rebuild" {
		t.Fatal("cold parse must be honestly marked as a full rebuild")
	}
	if len(cold.Files) != 2 {
		t.Fatalf("cold unit has %d files", len(cold.Files))
	}
	edits := []air.Edit{{Path: "a.fk", Offset: 0, EndOffset: 0,
		Text: []byte("let c = 3\n")}}
	edited, editedDiff, err := frontend.ParseIncremental(
		fakeSourceUnit(files...), old, edits)
	if err != nil {
		t.Fatal(err)
	}
	if editedDiff.Empty() || editedDiff.FullRebuild ||
		edited.Annotations["lang.go.reusedFiles"] != "1" {
		t.Fatalf("edited diff = %s, annotations = %v",
			editedDiff.String(), edited.Annotations)
	}
	if _, _, err := frontend.ParseIncremental(fakeSourceUnit(files...), old,
		[]air.Edit{{Path: "missing.fk", Offset: 0, EndOffset: 0}}); !errors.Is(err, errUnknownEditPath) {
		t.Fatalf("unknown edit path error = %v", err)
	}
}

func TestFrontendDeterminism100(t *testing.T) {
	t.Parallel()
	frontend := mustFrontend(t, newFakeSpec())
	files := fixtureFiles()
	determinism.MustMatch100(t, 42, func() []byte {
		unit, _, err := frontend.Parse(fakeSourceUnit(files...))
		if err != nil {
			t.Fatal(err)
		}
		digest := unit.Fingerprint()
		return digest[:]
	})
}

func TestIncrementalDeterminism100(t *testing.T) {
	t.Parallel()
	frontend := mustFrontend(t, newFakeSpec())
	files := fixtureFiles()
	old, _, err := frontend.Parse(fakeSourceUnit(files...))
	if err != nil {
		t.Fatal(err)
	}
	edits := []air.Edit{{Path: "a.fk", Offset: 0, EndOffset: 0,
		Text: []byte("let c = 3\n")}}
	determinism.MustMatch100(t, 42, func() []byte {
		out, _, err := frontend.ParseIncremental(
			fakeSourceUnit(files...), old, edits)
		if err != nil {
			t.Fatal(err)
		}
		digest := out.Fingerprint()
		return digest[:]
	})
}
