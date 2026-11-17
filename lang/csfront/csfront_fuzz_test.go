/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package csfront

import (
	"testing"

	"abate/lang/air"
)

// csFuzzSeeds 是原生 Go fuzz 的确定性种子语料。
func csFuzzSeeds() []string {
	return []string{
		"",
		"using System;\n",
		"namespace Demo { class C { } }\n",
		"class C { public int X { get; set; } }\n",
		"var q = from x in xs where x > 0 select x;\n",
		"string s = $\"hi {name}\";\n",
		"string p = @\"C:\\tmp\";\n",
		"class C { object M(object a) { return a?.X ?? b; } }\n",
		"#region R\n#pragma warning disable\n#endregion\n",
		"class C { void M( {\n",
		"\xff\xfe\n",
		"async Task M() { await F(); }\n",
		"record R(int X);\n",
	}
}

// FuzzCsFrontend 用 Go 原生 fuzz 归约任意字节缓冲。适配器必须不 panic，
// 未报告分类错误时必须返回合法 unit 与确定的指纹。
func FuzzCsFrontend(f *testing.F) {
	for _, seed := range csFuzzSeeds() {
		f.Add([]byte(seed))
	}
	f.Fuzz(func(t *testing.T, data []byte) {
		fe, err := NewCsFrontend()
		if err != nil {
			t.Fatalf("NewCsFrontend: %v", err)
		}
		unit := air.SourceUnit{
			Language: "cs", Dialect: "cs12",
			Files: []air.SourceFile{air.NewSourceFile("fuzz.cs", data)},
		}
		out, diags, parseErr := fe.Parse(unit)
		if parseErr != nil {
			return
		}
		if out == nil || len(out.Roots) != 1 || out.Roots[0] == nil {
			t.Fatalf("nil or malformed unit for %q", data[:csMinInt(len(data), 32)])
		}
		if err := out.Validate(); err != nil {
			t.Fatalf("invalid unit for %q: %v", data[:csMinInt(len(data), 32)], err)
		}
		if out.NodeCount() < 1 {
			t.Fatal("unit has no nodes")
		}
		if out.Roots[0].Kind() != air.KindPackage {
			t.Fatal("root is not KindPackage")
		}
		_ = out.Fingerprint()
		_ = diags
	})
}

// TestCsFuzzSeedsReproducible 把种子语料各解析两次，断言指纹逐字节一致。
func TestCsFuzzSeedsReproducible(t *testing.T) {
	fe := csFrontend(t)
	for _, src := range csFuzzSeeds() {
		request := air.SourceUnit{
			Language: "cs", Dialect: "cs12",
			Files: []air.SourceFile{air.NewSourceFile("seed.cs", []byte(src))},
		}
		a, _, errA := fe.Parse(request)
		b, _, errB := fe.Parse(request)
		if (errA == nil) != (errB == nil) || a == nil || b == nil {
			t.Fatalf("inconsistent seed parse for %q", src)
		}
		if a.Fingerprint() != b.Fingerprint() {
			t.Fatalf("seed %q is not deterministic", src)
		}
		if a.NodeCount() != b.NodeCount() {
			t.Fatalf("seed %q has non-deterministic node count", src)
		}
	}
}

// TestCsFuzzSeedsIntervals 断言每个种子都能通过位置区间校验。
func TestCsFuzzSeedsIntervals(t *testing.T) {
	for _, src := range csFuzzSeeds() {
		out, _, err := csFrontend(t).Parse(air.SourceUnit{
			Language: "cs", Dialect: "cs12",
			Files: []air.SourceFile{air.NewSourceFile("seed.cs", []byte(src))},
		})
		if err != nil {
			t.Fatalf("Parse(%q): %v", src, err)
		}
		if err := air.CheckIntervals(out.Roots[0], len(src)); err != nil {
			t.Fatalf("CheckIntervals for %q: %v", src, err)
		}
	}
}
