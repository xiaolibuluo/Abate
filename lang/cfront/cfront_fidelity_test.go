/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package cfront

import (
	"testing"

	"abate/lang/air"
	"abate/lang/frontkit"
)

// cSyntheticTokens 把一段 token 拼写序列转成合法的 frontkit.Token 切片。
func cSyntheticTokens(texts []string) []frontkit.Token {
	out := make([]frontkit.Token, 0, len(texts))
	offset := 0
	for _, text := range texts {
		kind := frontkit.TokenIdent
		if cIsKeyword(text) {
			kind = frontkit.TokenKeyword
		}
		out = append(out, frontkit.Token{
			Kind: kind, Text: text, Offset: offset, EndOffset: offset + len(text),
			Line: 1, Column: offset + 1,
		})
		offset += len(text) + 1
	}
	return out
}

// cLex 用本包词法表把源码切成 token 流。
func cLex(t *testing.T, source string) []frontkit.Token {
	t.Helper()
	lexer, err := frontkit.NewLexer(cLexSpec{}, []byte(source))
	if err != nil {
		t.Fatalf("NewLexer: %v", err)
	}
	tokens, err := lexer.Tokenize()
	if err != nil {
		t.Fatalf("Tokenize: %v", err)
	}
	return tokens
}

// cParseFiles 用 NewCFrontend 解析一组内存缓冲，返回 AIR unit 与诊断。
func cParseFiles(t *testing.T, files []air.SourceFile,
	dialect string) (*air.Unit, []air.Diagnostic) {
	t.Helper()
	frontend, err := NewCFrontend()
	if err != nil {
		t.Fatalf("NewCFrontend: %v", err)
	}
	unit, diagnostics, err := frontend.Parse(air.SourceUnit{
		Language: "c", Dialect: dialect, Files: files,
	})
	if err != nil {
		t.Fatalf("Parse: %v", err)
	}
	if unit == nil {
		t.Fatal("Parse returned a nil unit")
	}
	return unit, diagnostics
}

// cParseOne 解析一段单文件源码。
func cParseOne(t *testing.T, source, dialect string) (*air.Unit, []air.Diagnostic) {
	t.Helper()
	return cParseFiles(t, []air.SourceFile{
		air.NewSourceFile("sample.cpp", []byte(source)),
	}, dialect)
}

// cKindCounts 统计一个 unit 里每种 NodeKind 的数量。
func cKindCounts(unit *air.Unit) map[air.NodeKind]int {
	counts := make(map[air.NodeKind]int)
	for _, node := range air.Collect(unit) {
		counts[node.Kind()]++
	}
	return counts
}

// cCountKind 统计一棵节点树里指定种类的数量。
func cCountKind(root air.Node, kind air.NodeKind) int {
	count := 0
	air.Walk(root, func(node air.Node) bool {
		if node.Kind() == kind {
			count++
		}
		return true
	})
	return count
}

// cSemValue 读取节点上的一个字符串语义标注。
func cSemValue(node air.Node, key air.SemKey) (string, bool) {
	value, ok := node.Sem(key).(string)
	return value, ok
}

// cHasSemValue 判断 unit 里是否存在携带指定标注的节点。
func cHasSemValue(unit *air.Unit, key air.SemKey, want string) bool {
	for _, node := range air.Collect(unit) {
		if value, ok := cSemValue(node, key); ok && value == want {
			return true
		}
	}
	return false
}

// cAssertIntervalCoverage 实现设计文档 F3 对 C/C++ 的验收方式：
// 用「位置区间全覆盖」而非 token 等价。根必须覆盖 [0,size)，每个节点区间
// 必须合法、落在父节点内且兄弟不重叠，且每个非 trivia 词法 token 都必须
// 被某个非根 AIR 节点的区间包含。
func cAssertIntervalCoverage(t *testing.T, root air.Node,
	tokens []frontkit.Token, size int) {
	t.Helper()
	if root == nil {
		t.Fatal("nil root")
	}
	span := root.Pos()
	if span.Offset != 0 || span.EndOffset != size {
		t.Fatalf("root span [%d,%d) does not fully cover [0,%d)",
			span.Offset, span.EndOffset, size)
	}
	if err := air.CheckIntervals(root, size); err != nil {
		t.Fatalf("CheckIntervals: %v", err)
	}
	nodes := make([]air.Node, 0, 32)
	air.Walk(root, func(node air.Node) bool {
		if !node.Pos().Valid() || !node.End().Valid() {
			t.Fatalf("node %s has an invalid position", node.Kind())
		}
		nodes = append(nodes, node)
		return true
	})
	inner := nodes[1:] // 排除根，避免「根覆盖一切」的空断言
	for _, token := range tokens {
		switch token.Kind {
		case frontkit.TokenNewline, frontkit.TokenEOF,
			frontkit.TokenIndent, frontkit.TokenDedent:
			continue
		}
		covered := false
		for _, node := range inner {
			start, end := air.Span(node)
			if start <= token.Offset && token.EndOffset <= end {
				covered = true
				break
			}
		}
		if !covered {
			t.Fatalf("token %q %s at [%d,%d) has no covering AIR node",
				token.Text, token.Kind, token.Offset, token.EndOffset)
		}
	}
}

// cCorpus 是跨三方言的确定性位置区间夹具语料。
func cCorpus() []struct {
	name, dialect, source string
} {
	return []struct {
		name, dialect, source string
	}{
		{"cpp20", "c++20", cSampleSource()},
		{"cpp17", "c++17", "namespace n {\nclass C { public: int v; };\n}\n"},
		{"c11", "c11", "#include <stdio.h>\nint main(void) { int x = 1; return x; }\n"},
		{"preproc-only", "c++17", "#define A 1\n#ifdef A\n#undef A\n#endif\n#pragma once\n"},
		{"comments", "c11", "// line\n/* block */\nint g;\n"},
		{"literals", "c++20", "auto s = u8\"hi\";\nwchar_t c = L'x';\nint n = 1'000;\n"},
		{"empty", "c11", ""},
		{"only-newlines", "c++17", "\n\n\n"},
	}
}

// TestIntervalCoverageCorpus 对整份语料断言位置区间 100% 全覆盖。
func TestIntervalCoverageCorpus(t *testing.T) {
	for _, item := range cCorpus() {
		t.Run(item.name, func(t *testing.T) {
			tokens := cLex(t, item.source)
			unit, _ := cParseOne(t, item.source, item.dialect)
			if len(unit.Roots) != 1 {
				t.Fatalf("expected one root, got %d", len(unit.Roots))
			}
			cAssertIntervalCoverage(t, unit.Roots[0], tokens, len(item.source))
		})
	}
}

// TestStructureSpansMultipleTokens 证明归约是有层级的结构树，
// 而不是「一个 token 对应一个节点」的 token 等价投影。
func TestStructureSpansMultipleTokens(t *testing.T) {
	source := cSampleSource()
	tokens := cLex(t, source)
	unit, _ := cParseOne(t, source, "c++20")
	multi := false
	for _, node := range air.Collect(unit) {
		start, end := air.Span(node)
		contained := 0
		for _, token := range tokens {
			switch token.Kind {
			case frontkit.TokenNewline, frontkit.TokenEOF,
				frontkit.TokenIndent, frontkit.TokenDedent:
				continue
			}
			if start <= token.Offset && token.EndOffset <= end {
				contained++
			}
		}
		if contained >= 4 {
			multi = true
			break
		}
	}
	if !multi {
		t.Fatal("no structural node spans multiple tokens; reduction looks flat")
	}
}

// TestFingerprintStableAcrossParses 断言相同输入两次解析指纹与区间完全一致。
func TestFingerprintStableAcrossParses(t *testing.T) {
	source := cSampleSource()
	first, _ := cParseOne(t, source, "c++17")
	second, _ := cParseOne(t, source, "c++17")
	if first.Fingerprint() != second.Fingerprint() {
		t.Fatal("identical input produced different fingerprints")
	}
	if err := air.CompareUnits(first, second); err != nil {
		t.Fatalf("CompareUnits: %v", err)
	}
	if err := air.CheckIntervals(second.Roots[0], len(source)); err != nil {
		t.Fatalf("CheckIntervals: %v", err)
	}
}

// TestFileOrderIndependence 证明不同文件顺序产出同一指纹。
func TestFileOrderIndependence(t *testing.T) {
	frontend, err := NewCFrontend()
	if err != nil {
		t.Fatal(err)
	}
	files := []air.SourceFile{
		air.NewSourceFile("z.cpp", []byte("int z = 1;\n")),
		air.NewSourceFile("a.h", []byte("int a = 2;\n")),
	}
	first, _, err := frontend.Parse(air.SourceUnit{Language: "c", Files: files})
	if err != nil {
		t.Fatal(err)
	}
	reversed := []air.SourceFile{files[1], files[0]}
	second, _, err := frontend.Parse(air.SourceUnit{Language: "c", Files: reversed})
	if err != nil {
		t.Fatal(err)
	}
	if first.Fingerprint() != second.Fingerprint() {
		t.Fatal("file order leaked into the fingerprint")
	}
}

// TestCOperatorFidelity asserts AIR-ID changes for cmp, op and digit mutations, with identical trailing comments on both variants.
func TestCOperatorFidelity(t *testing.T) {
	cases := []struct{ name, base, mutated string }{
		{"cmp", "int f(int a, int b) { return a == b; }\n", "int f(int a, int b) { return a != b; }\n"},
		{"op", "int f(int a, int b) { return a + b; }\n", "int f(int a, int b) { return a - b; }\n"},
		{"digit", "int f() { return 41; }\n", "int f() { return 42; }\n"},
	}
	rootID := func(t *testing.T, src string) air.NodeID {
		t.Helper()
		unit, _ := cParseOne(t, src, "c++17")
		return air.IDOf(unit.Roots[0])
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if rootID(t, tc.base) == rootID(t, tc.mutated) {
				t.Fatal("operator or digit mutation did not change AIR-ID")
			}
			baseComment, mutatedComment := tc.base+"// tail\n", tc.mutated+"// tail\n"
			if rootID(t, baseComment) == rootID(t, mutatedComment) {
				t.Fatal("identical comments changed the capture conclusion")
			}
		})
	}
}
