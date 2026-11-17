/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package gofront

import (
	"errors"
	"strings"
	"testing"

	"abate/lang/air"
)

// TestMappingTableDeclarations checks every declaration mapping.
func TestMappingTableDeclarations(t *testing.T) {
	t.Parallel()
	unit := mustParse(t, sampleGo)
	for _, kind := range []air.NodeKind{
		air.KindPackage, air.KindImport, air.KindTypeDecl, air.KindConstDecl,
		air.KindVarDecl, air.KindFuncDecl, air.KindMethodDecl,
	} {
		if !hasKind(unit, kind) {
			t.Fatalf("declaration kind %s is missing", kind)
		}
	}
	if unit.Package.Name != "sample" || len(unit.Package.Imports) != 2 {
		t.Fatalf("package metadata: %+v", unit.Package)
	}
	if len(unit.Files) != 1 || unit.Files[0].Path != "sample.go" {
		t.Fatalf("file metadata: %+v", unit.Files)
	}
	if unit.NodeCount() < 100 {
		t.Fatalf("mapping produced only %d nodes", unit.NodeCount())
	}
}

// TestMappingTableStatements checks every statement mapping family.
func TestMappingTableStatements(t *testing.T) {
	t.Parallel()
	unit := mustParse(t, sampleGo)
	for _, kind := range []air.NodeKind{
		air.KindBlock, air.KindExprStmt, air.KindDeclStmt, air.KindIf,
		air.KindFor, air.KindRangeFor, air.KindSwitch, air.KindCase,
		air.KindSelect, air.KindDefer, air.KindGoStmt, air.KindReturn,
		air.KindBranch, air.KindLabel, air.KindAssign, air.KindIncDec,
		air.KindChanOp,
	} {
		if !hasKind(unit, kind) {
			t.Fatalf("statement kind %s is missing", kind)
		}
	}
}

// TestMappingTableExpressions checks every expression mapping family.
func TestMappingTableExpressions(t *testing.T) {
	t.Parallel()
	unit := mustParse(t, sampleGo)
	for _, kind := range []air.NodeKind{
		air.KindIdent, air.KindNumLit, air.KindStrLit, air.KindNilLit,
		air.KindBinary, air.KindUnary, air.KindAddr, air.KindCall,
		air.KindIndex, air.KindSlice, air.KindField, air.KindSelector,
		air.KindCompositeLit, air.KindFuncLit, air.KindTypeAssert,
		air.KindStar, air.KindSpread, air.KindGenericInst, air.KindRegion,
	} {
		if !hasKind(unit, kind) {
			t.Fatalf("expression kind %s is missing", kind)
		}
	}
}

// TestMappingTableTypes checks every type mapping family.
func TestMappingTableTypes(t *testing.T) {
	t.Parallel()
	unit := mustParse(t, sampleGo)
	for _, kind := range []air.NodeKind{
		air.KindNamed, air.KindPointer, air.KindSliceType, air.KindArray,
		air.KindMap, air.KindChanType, air.KindFuncType,
		air.KindStructType, air.KindInterfaceType, air.KindGenericParam,
		air.KindAttribute, air.KindComment, air.KindDoc, air.KindDirective,
	} {
		if !hasKind(unit, kind) {
			t.Fatalf("type or trivia kind %s is missing", kind)
		}
	}
}

// TestDirectivesIotaTagsAndNotes checks the Go-specific annotations.
func TestDirectivesIotaTagsAndNotes(t *testing.T) {
	t.Parallel()
	unit := mustParse(t, sampleGo)
	if got := annotationOf(unit, air.KindDirective, air.SemDirective); !strings.Contains(got, "//go:noinline") {
		t.Fatalf("directive annotation = %q", got)
	}
	if got := annotationOf(unit, air.KindConstDecl, air.SemIota); got == "" {
		t.Fatal("iota position not annotated")
	}
	if got := annotationOf(unit, air.KindAttribute, air.SemTag); !strings.Contains(got, "json") {
		t.Fatalf("struct tag annotation = %q", got)
	}
	if got := annotationOf(unit, air.KindChanOp, air.SemChanDir); got == "" {
		t.Fatal("channel direction not annotated")
	}
	if got := annotationOf(unit, air.KindSwitch, air.SemNote); got != "type-switch" {
		t.Fatalf("type switch annotation = %q", got)
	}
	if got := annotationOf(unit, air.KindAssign, air.SemNote); got != "key-value" {
		t.Fatalf("key-value annotation = %q", got)
	}
	if got := annotationOf(unit, air.KindGenericParam, air.SemTypeParam); got != "" {
		t.Fatalf("generic parameter wrapper carries an unexpected annotation: %q", got)
	}
}

// typedGo is a self-contained file for go/types annotation assertions.
const typedGo = `package typed

func Add(a int, b int) int {
	return a + b
}

func Effect(ch chan int) {
	ch <- 1
}
`

// TestGoTypesAnnotations proves the in-memory checker enriches AIR nodes.
func TestGoTypesAnnotations(t *testing.T) {
	t.Parallel()
	unit := mustParse(t, typedGo)
	pure := ""
	effect := ""
	for _, fn := range collectKinds(unit, air.KindFuncDecl) {
		switch air.AttributeOf(fn) {
		case "Add":
			if value, ok := fn.Sem(air.SemPurity).(string); ok {
				pure = value
			}
		case "Effect":
			if value, ok := fn.Sem(air.SemPurity).(string); ok {
				effect = value
			}
		}
	}
	if pure != "pure" || effect != "side-effect" {
		t.Fatalf("purity heuristic = %q / %q", pure, effect)
	}
	typed := false
	scoped := false
	for _, current := range air.Collect(unit) {
		if value, ok := current.Sem(air.SemTypeRef).(string); ok && value == "int" {
			typed = true
		}
		if value, ok := current.Sem(air.SemScopeID).(string); ok && value != "" {
			scoped = true
		}
	}
	if !typed || !scoped {
		t.Fatalf("type or scope annotation missing: typed=%t scoped=%t", typed, scoped)
	}
}

// TestSyntaxErrorsBecomeDiagnostics checks graceful handling of bad sources.
func TestSyntaxErrorsBecomeDiagnostics(t *testing.T) {
	t.Parallel()
	sources := []string{
		"",
		"package",
		"func (",
		"package p\nfunc f( {\n",
		"\xef\xbb\xbfpackage p\n",
		"// only a comment\n",
	}
	for i, src := range sources {
		unit := air.SourceUnit{Language: "go", Files: []air.SourceFile{
			air.NewSourceFile("bad.go", []byte(src)),
		}}
		parsed, diagnostics, err := NewGoFrontend().WithTypeCheck(false).Parse(unit)
		if err != nil {
			t.Fatalf("source %d returned a contract failure: %v", i, err)
		}
		if parsed == nil || len(parsed.Files) != 1 || parsed.Roots[0] == nil {
			t.Fatalf("source %d produced an incomplete unit", i)
		}
		for _, diagnostic := range diagnostics {
			if !diagnostic.IsError() || diagnostic.Code == "" ||
				diagnostic.Line < 0 || diagnostic.Column < 0 {
				t.Fatalf("source %d malformed diagnostic: %+v", i, diagnostic)
			}
		}
	}
	broken := air.SourceUnit{Language: "go", Files: []air.SourceFile{
		air.NewSourceFile("bad.go", []byte("func (")),
	}}
	_, diagnostics, err := NewGoFrontend().WithMaxDiagnostics(1).Parse(broken)
	if err != nil || len(diagnostics) > 1 {
		t.Fatalf("diagnostic cap ignored: %d %v", len(diagnostics), err)
	}
}

// TestFrontendOptionsAndLanguage checks contract-level validation.
func TestFrontendOptionsAndLanguage(t *testing.T) {
	t.Parallel()
	frontend := NewGoFrontend()
	bad := air.SourceUnit{Language: "python", Files: []air.SourceFile{
		air.NewSourceFile("x.py", []byte("print(1)")),
	}}
	if _, _, err := frontend.Parse(bad); !errors.Is(err, ErrUnsupportedLanguage) {
		t.Fatalf("unsupported language: %v", err)
	}
	duplicate := air.SourceUnit{Language: "go", Files: []air.SourceFile{
		air.NewSourceFile("a.go", []byte("package p\n")),
		air.NewSourceFile("a.go", []byte("package p\n")),
	}}
	if _, _, err := frontend.Parse(duplicate); !errors.Is(err, air.ErrDuplicatePath) {
		t.Fatalf("duplicate path: %v", err)
	}
	empty := air.SourceUnit{Language: "go"}
	unit, diagnostics, err := frontend.WithMaxDiagnostics(0).Parse(empty)
	if err != nil || unit == nil || len(diagnostics) != 0 || unit.NodeCount() != 0 {
		t.Fatalf("empty unit: %+v %v %v", unit, diagnostics, err)
	}
}
