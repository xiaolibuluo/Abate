/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package cfront

import (
	"testing"

	"abate/internal/determinism"
	"abate/lang/air"
)

// TestNewCFrontendSatisfiesAirFrontend 证明 NewCFrontend 满足 air.Frontend
// 且 Parse 产出有真实节点、结构合法的 *air.Unit。
func TestNewCFrontendSatisfiesAirFrontend(t *testing.T) {
	frontend, err := NewCFrontend()
	if err != nil || frontend == nil {
		t.Fatalf("NewCFrontend returned %v, %v", frontend, err)
	}
	asFrontend := frontend
	unit, diagnostics, err := asFrontend.Parse(air.SourceUnit{
		Language: "c", Dialect: "c++17",
		Files: []air.SourceFile{
			air.NewSourceFile("main.cpp", []byte(cSampleSource())),
		},
	})
	if err != nil {
		t.Fatalf("Parse: %v", err)
	}
	if unit == nil || unit.NodeCount() < 40 {
		t.Fatalf("expected a real node tree, got NodeCount=%d", unit.NodeCount())
	}
	if err := unit.Validate(); err != nil {
		t.Fatalf("unit.Validate: %v", err)
	}
	if unit.Roots[0].Kind() != air.KindPackage {
		t.Fatalf("root kind = %s, want Package", unit.Roots[0].Kind())
	}
	if _, ok := cSemValue(unit.Roots[0], air.SemScopeID); !ok {
		t.Fatal("root is missing the translation-unit scope annotation")
	}
	if len(diagnostics) != 0 {
		t.Logf("sample produced %d diagnostics", len(diagnostics))
	}
}

// TestConstructMappings 遍历 ≥20 个构造映射契约，逐一验证归约结果。
func TestConstructMappings(t *testing.T) {
	cases := cConstructCases()
	if len(cases) < 20 {
		t.Fatalf("only %d construct cases, want at least 20", len(cases))
	}
	mapper := newCMapper()
	for _, item := range cases {
		tokens := cSyntheticTokens(item.Tokens)
		got := mapper.match(tokens, 0, len(tokens), item.Top)
		if got.kind != item.WantKind {
			t.Errorf("%s: kind = %s, want %s",
				item.Name, got.kind, item.WantKind)
		}
		if !got.kind.Valid() {
			t.Errorf("%s: produced an invalid NodeKind", item.Name)
		}
		node := air.NewNode(got.kind, air.NewPosition(1, 1, 0, 0), got.name)
		mapper.annotate(node, got)
		if item.WantSemKey.Valid() {
			value, ok := cSemValue(node, item.WantSemKey)
			if !ok || value != item.WantSemValue {
				t.Errorf("%s: sem[%s] = %q, want %q",
					item.Name, item.WantSemKey, value, item.WantSemValue)
			}
		}
	}
	// nil 节点与空映射的降级分支也必须安全。
	mapper.annotate(nil, cMapping{kind: air.KindVarDecl})
	mapper.annotate(air.NewNode(air.KindVarDecl, air.NewPosition(1, 1, 0, 0), ""),
		cMapping{kind: air.KindVarDecl})
}

// TestParseProducesRealConstructTree 断言大夹具解析出多样化真实节点。
func TestParseProducesRealConstructTree(t *testing.T) {
	unit, _ := cParseOne(t, cSampleSource(), "c++20")
	counts := cKindCounts(unit)
	required := []air.NodeKind{
		air.KindPackage, air.KindDirective, air.KindGenericInst,
		air.KindModuleDecl, air.KindClassDecl, air.KindStructDecl,
		air.KindEnumDecl, air.KindFuncDecl, air.KindVarDecl, air.KindBlock,
		air.KindReturn, air.KindIf, air.KindFor, air.KindWhile,
		air.KindSwitch, air.KindNumLit, air.KindStrLit, air.KindComment,
	}
	for _, kind := range required {
		if counts[kind] == 0 {
			t.Errorf("missing %s nodes in the parsed tree", kind)
		}
	}
	if !cHasSemValue(unit, air.SemDirective, "include") {
		t.Error("missing SemDirective=include")
	}
	if !cHasSemValue(unit, air.SemTypeParam, "template") {
		t.Error("missing SemTypeParam=template")
	}
	if !cHasSemValue(unit, air.SemNote, "header:helper.h") {
		t.Error("missing in-memory header reference")
	}
	if !cHasSemValue(unit, air.SemNote, "macro:MAX_SIZE") {
		t.Error("missing unexpanded macro registration")
	}
	if !cHasSemValue(unit, air.SemOperator, "+") {
		t.Error("missing binary operator annotation")
	}
	if err := air.CheckIntervals(unit.Roots[0], len(cSampleSource())); err != nil {
		t.Fatalf("CheckIntervals: %v", err)
	}
}

// TestDeterminism100 要求 100 次重新解析的指纹与节点数逐字节一致。
func TestDeterminism100(t *testing.T) {
	frontend, err := NewCFrontend()
	if err != nil {
		t.Fatal(err)
	}
	source := air.SourceUnit{
		Language: "c", Dialect: "c++17",
		Files: []air.SourceFile{
			air.NewSourceFile("b.cpp", []byte(cSampleSource())),
			air.NewSourceFile("a.h", []byte("#pragma once\nint shared;\n")),
		},
	}
	determinism.MustMatch100(t, 0xC05F0001, func() []byte {
		unit, _, parseErr := frontend.Parse(source)
		if parseErr != nil {
			t.Fatal(parseErr)
		}
		fingerprint := unit.Fingerprint()
		return fingerprint[:]
	})
	digests := make(map[[32]byte]struct{})
	counts := make(map[int]struct{})
	for i := 0; i < 100; i++ {
		unit, _, parseErr := frontend.Parse(source)
		if parseErr != nil {
			t.Fatal(parseErr)
		}
		digests[determinism.HashState(unit)] = struct{}{}
		counts[unit.NodeCount()] = struct{}{}
	}
	if len(digests) != 1 || len(counts) != 1 {
		t.Fatalf("nondeterministic: %d digests, %d node counts",
			len(digests), len(counts))
	}
}
