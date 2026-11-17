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

// cSampleSource 是覆盖 ≥20 个 C/C++ 真实构造的确定性内联夹具。
func cSampleSource() string {
	return `#include "helper.h"
#include <vector>
#define MAX_SIZE 64
#ifdef VERBOSE
#pragma once
#endif
// a translation-unit comment
namespace demo {
template <typename T>
class Box {
public:
    T value;
    static int count;
};
struct Point { int x; int y; };
class Widget { public: int w; };
union Value { int i; double d; };
enum Color { Red, Green, Blue };
enum class Mode : int { Fast, Slow };
using Alias = Box<int>;
typedef unsigned long Word;
const int kLimit = MAX_SIZE;
extern int global_counter;
int add(int a, int b) {
    int sum = a + b;
    if (sum > kLimit) {
        return kLimit;
    }
    for (int i = 0; i < 3; ++i) {
        sum += i;
    }
    while (sum > 0) {
        --sum;
    }
    switch (sum) {
        case 0:
            break;
        default:
            continue;
    }
    int* ptr = new int(sum);
    int copy = *ptr;
    int& ref = copy;
    int addr = &copy;
    size_t n = sizeof(int);
    auto casted = static_cast<double>(copy);
    delete ptr;
    return sum;
}
}
`
}

// cMainWithIncludes 是引用虚拟头文件的翻译单元。
func cMainWithIncludes() string {
	return `#include "helper.h"
#include <vector>
#define SQUARE(x) ((x) * (x))
int use_macro(void) {
    return SQUARE(3);
}
`
}

// cHeaderSource 是内存提供的虚拟头文件，含 include guard 与声明。
func cHeaderSource() string {
	return `#pragma once
#ifndef HELPER_H
#define HELPER_H
namespace helper {
struct Header { int field; };
int header_fn(int v);
}
#endif
`
}

// cDefsHeaderSource 是第二个虚拟头文件，验证多文件与嵌套命名空间。
func cDefsHeaderSource() string {
	return `namespace defs {
enum class Kind : int { One, Two };
using Index = unsigned long;
}
`
}

// cHeaderFiles 构造一个主文件 + 两个虚拟头文件的内存 SourceUnit。
func cHeaderFiles() []air.SourceFile {
	return []air.SourceFile{
		air.NewSourceFile("main.cpp", []byte(cMainWithIncludes())),
		air.NewSourceFile("helper.h", []byte(cHeaderSource())),
		air.NewSourceFile("include/defs.hpp", []byte(cDefsHeaderSource())),
	}
}

// TestVirtualHeadersMultiFile 验证多文件解析、每个文件真实节点树与区间覆盖。
// 全流程只依赖 air.SourceUnit.Files 内存缓冲，零真实文件系统读取。
func TestVirtualHeadersMultiFile(t *testing.T) {
	files := cHeaderFiles()
	unit, diagnostics := cParseFiles(t, files, "c++20")
	if len(unit.Files) != 3 || len(unit.Roots) != 3 {
		t.Fatalf("expected 3 parsed files, got %d files / %d roots",
			len(unit.Files), len(unit.Roots))
	}
	if err := unit.Validate(); err != nil {
		t.Fatalf("unit.Validate: %v", err)
	}
	byPath := make(map[string][]byte, len(files))
	for _, file := range files {
		byPath[file.Path] = file.Content
	}
	for i, ref := range unit.Files {
		content, ok := byPath[ref.Path]
		if !ok {
			t.Fatalf("unexpected path %q", ref.Path)
		}
		root := unit.Roots[i]
		if root == nil || cCountKind(root, air.KindIdent) == 0 {
			t.Fatalf("file %q produced no real nodes", ref.Path)
		}
		tokens := cLex(t, string(content))
		cAssertIntervalCoverage(t, root, tokens, len(content))
	}
	counts := cKindCounts(unit)
	for _, kind := range []air.NodeKind{
		air.KindDirective, air.KindModuleDecl, air.KindStructDecl,
		air.KindEnumDecl, air.KindFuncDecl, air.KindUsingDecl,
	} {
		if counts[kind] == 0 {
			t.Errorf("virtual header corpus missing %s nodes", kind)
		}
	}
	_ = diagnostics
}

// TestIncludeDirectiveNoFS 断言 #include 只登记内存引用、宏绝不展开。
func TestIncludeDirectiveNoFS(t *testing.T) {
	files := cHeaderFiles()
	unit, _ := cParseFiles(t, files, "c++17")
	includeCount := 0
	hasHeaderNote := false
	for _, node := range air.Collect(unit) {
		if node.Kind() != air.KindDirective {
			continue
		}
		if value, ok := cSemValue(node, air.SemDirective); ok && value == "include" {
			includeCount++
			if note, exists := cSemValue(node, air.SemNote); exists &&
				note == "header:helper.h" {
				hasHeaderNote = true
			}
		}
	}
	if includeCount == 0 || !hasHeaderNote {
		t.Fatalf("include directives not reduced: count=%d headerNote=%v",
			includeCount, hasHeaderNote)
	}
	if !cHasSemValue(unit, air.SemDirective, "pragma") {
		t.Error("missing pragma directive")
	}
	if !cHasSemValue(unit, air.SemDirective, "ifndef") {
		t.Error("missing include-guard directive")
	}
	if !cHasSemValue(unit, air.SemNote, "macro:SQUARE") {
		t.Error("missing unexpanded macro registration")
	}
	// 宏调用处的宏名仍是普通标识符，宏体从未被代入。
	squareUses := 0
	for _, node := range air.Collect(unit) {
		if air.AttributeOf(node) == "SQUARE" {
			squareUses++
		}
	}
	if squareUses < 2 {
		t.Fatalf("macro name appeared %d times; macro was expanded", squareUses)
	}
	// 头文件只在内存缓冲里查找，从不触碰磁盘。
	if found, ok := ResolveInMemoryHeader("helper.h", files); !ok ||
		len(found.Content) == 0 {
		t.Fatal("virtual header not resolvable from SourceUnit.Files")
	}
	if found, ok := ResolveInMemoryHeader("defs.hpp", files); !ok ||
		string(found.Content) != cDefsHeaderSource() {
		t.Fatal("nested virtual header not resolvable by base name")
	}
	if _, ok := ResolveInMemoryHeader("missing.h", files); ok {
		t.Fatal("an absent header must not resolve")
	}
	if _, ok := ResolveInMemoryHeader("", files); ok {
		t.Fatal("an empty header name must not resolve")
	}
}

// TestHeaderDeterminism100 对多文件虚拟头文件单元做 100 次哈希一致断言。
func TestHeaderDeterminism100(t *testing.T) {
	frontend, err := NewCFrontend()
	if err != nil {
		t.Fatal(err)
	}
	request := air.SourceUnit{
		Language: "c", Dialect: "c++20", Files: cHeaderFiles(),
	}
	determinism.MustMatch100(t, 0xC05F0002, func() []byte {
		unit, _, parseErr := frontend.Parse(request)
		if parseErr != nil {
			t.Fatal(parseErr)
		}
		fingerprint := unit.Fingerprint()
		return fingerprint[:]
	})
}
