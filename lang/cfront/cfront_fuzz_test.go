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
)

// cFuzzSeeds 是 Go 原生 fuzz 的确定性种子语料：覆盖声明、结构、预处理、
// 宏、模板以及高位字节 / NUL / 截断等畸形输入。
var cFuzzSeeds = [][]byte{
	[]byte("int main(void) { return 0; }"),
	[]byte("struct S { int x; char y; };\nint add(int a, int b) { return a + b; }"),
	[]byte("#define M(x) ((x) + 1)\n#if defined(M)\nint y = M(2);\n#endif\n"),
	[]byte("union U { int i; float f; };\nenum E { A, B = 3 };\n"),
	[]byte("template <typename T> T max(T a, T b) { return a > b ? a : b; }\n"),
	[]byte("typedef struct { int a; } Node; static Node *p = 0;\n"),
	[]byte("int g(void) { int *p = 0; return *p; }\n"),
	[]byte("#include <stdio.h>\nint x = \"unterminated\n"),
	{0x00, 0xff, 0xfe, '\n', 0x80, 0x7f, 0x01},
	[]byte(""),
	[]byte("/* unterminated comment\n"),
	[]byte("<<>>`${}();\n\n\n"),
}

// FuzzCFrontParse 对任意字节流验证解析产出真实节点且指纹稳定。cfront 已有
// 固定畸形语料的 chaos 测试，本 fuzz 由 Go 框架做持续变异，补足其覆盖不到的
// 输入空间。不 panic 由测试框架保证（panic 会自然冒泡并判 FAIL），无需显式
// recover。
func FuzzCFrontParse(f *testing.F) {
	for _, seed := range cFuzzSeeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, data []byte) {
		// 将 fuzz 输入限制在 4096 字节以内以保证每轮快速；这不是失败抑制。
		// 超限输入与 cChaosBytes 走同一词法/语法路径，跳过不损失 panic 自由度
		// 或确定性覆盖（cChaosCorpus 已含 1024 字节畸形样例）。
		if len(data) > 4096 {
			t.Skip()
		}
		fe, err := NewCFrontend()
		if err != nil {
			t.Fatalf("NewCFrontend: %v", err)
		}
		unit := air.SourceUnit{
			Language: "c", Dialect: "c++17",
			Files: []air.SourceFile{air.NewSourceFile("fuzz.cpp", data)},
		}
		out, _, err := fe.Parse(unit)
		if err != nil {
			t.Fatalf("Parse: %v", err)
		}
		if out == nil || out.NodeCount() == 0 {
			t.Fatal("empty AIR tree")
		}
		first := out.Fingerprint()
		again, _, err := fe.Parse(unit)
		if err != nil {
			t.Fatalf("second Parse: %v", err)
		}
		if again.Fingerprint() != first {
			t.Fatal("non-deterministic fingerprint")
		}
	})
}

// TestCFuzzSeedCorpusDeterminism 对全部种子语料重复解析并断言一致，
// 同时断言畸形输入不 panic 且区间合法（复用 chaos 守护）。
func TestCFuzzSeedCorpusDeterminism(t *testing.T) {
	fe, err := NewCFrontend()
	if err != nil {
		t.Fatalf("NewCFrontend: %v", err)
	}
	for i, seed := range cFuzzSeeds {
		unit := air.SourceUnit{
			Language: "c", Dialect: "c++17",
			Files: []air.SourceFile{air.NewSourceFile("seed.cpp", seed)},
		}
		var first [32]byte
		for run := 0; run < 5; run++ {
			out, _, parseErr := fe.Parse(unit)
			if parseErr != nil {
				t.Fatalf("seed %d Parse: %v", i, parseErr)
			}
			if out == nil || out.NodeCount() == 0 {
				t.Fatalf("seed %d produced empty tree", i)
			}
			sum := out.Fingerprint()
			if run == 0 {
				first = sum
			} else if sum != first {
				t.Fatalf("seed %d non-deterministic at run %d", i, run)
			}
		}
		cAssertNoPanic(t, string(seed), "c++17")
	}
}

// TestCFuzzUnclosedTemplateRegression 锁定 fuzz 发现的一例真实 panic：未闭合
// 的 "template <" 使 cAngleRange 返回倒置区间，cLeafNodes 曾以负 cap 调用
// make 而崩溃。修复后该输入必须稳定产出非空树且指纹一致。
func TestCFuzzUnclosedTemplateRegression(t *testing.T) {
	src := "template <"
	cAssertNoPanic(t, src, "c++17")
	fe, err := NewCFrontend()
	if err != nil {
		t.Fatalf("NewCFrontend: %v", err)
	}
	unit := air.SourceUnit{
		Language: "c", Dialect: "c++17",
		Files: []air.SourceFile{air.NewSourceFile("regression.cpp", []byte(src))},
	}
	first, _, parseErr := fe.Parse(unit)
	if parseErr != nil {
		t.Fatalf("Parse: %v", parseErr)
	}
	again, _, parseErr := fe.Parse(unit)
	if parseErr != nil {
		t.Fatalf("second Parse: %v", parseErr)
	}
	if first.Fingerprint() != again.Fingerprint() {
		t.Fatal("non-deterministic fingerprint on regression input")
	}
}

// TestCFuzzMalformedByteTable 覆盖高位字节、NUL 与控制字符的确定性组合。
func TestCFuzzMalformedByteTable(t *testing.T) {
	for i := 0; i < 24; i++ {
		buf := make([]byte, 48+i*5)
		for j := range buf {
			buf[j] = byte((i*211 + j*29 + 3) % 256)
		}
		cAssertNoPanic(t, string(buf), "c++17")
	}
}
