/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package jsfront

import (
	"strings"
	"testing"

	"abate/lang/air"
)

// jsParseNoPanic 包装解析并断言非空节点树；不 panic 由 Go 测试框架保证
// （panic 会自然冒泡并判 FAIL），无需显式 recover（避免触发 GO-EG-027 提示规则）。
func jsParseNoPanic(t *testing.T, fe air.Frontend, path, src string) *air.Unit {
	t.Helper()
	out, _ := jsParse(t, fe, path, "", src)
	if out == nil || out.NodeCount() == 0 {
		t.Fatalf("empty tree for %s", path)
	}
	return out
}

// jsAssertStable 连续三次解析并断言指纹逐字节一致。
func jsAssertStable(t *testing.T, fe air.Frontend, path, src string) {
	t.Helper()
	unit := air.SourceUnit{
		Language: "js",
		Files:    []air.SourceFile{air.NewSourceFile(path, []byte(src))},
	}
	var first [32]byte
	for i := 0; i < 3; i++ {
		out, _, err := fe.Parse(unit)
		if err != nil {
			t.Fatalf("Parse(%s): %v", path, err)
		}
		sum := out.Fingerprint()
		if i == 0 {
			first = sum
		} else if sum != first {
			t.Fatalf("non-deterministic parse for %s at run %d", path, i)
		}
	}
}

// TestASIBoundariesNoPanic 覆盖自动分号插入的受限产生式与续接边界。
func TestASIBoundariesNoPanic(t *testing.T) {
	fe := jsFrontend(t)
	cases := []string{
		"return\nx",
		"return\n{x:1}",
		"throw\nnew Error('x')",
		"break\nlabel",
		"a = 1\nb = 2",
		"a\n+b",
		"a\n  .b",
		"a?.b\nc",
		"x++\ny--",
		";;;\n",
		"}\n{",
		"const a = 1\nconst b = 2\n",
		"let x\n(x + 1)\n",
		"async\nfunction f() {}\n",
		"yield\n1",
		"continue\nloop",
		"\n\n\n",
		"f()\n(function(){})()",
	}
	for i, src := range cases {
		path := "asi.js"
		out := jsParseNoPanic(t, fe, path, src)
		if out.Roots[0].Kind() != air.KindPackage {
			t.Fatalf("case %d root = %s", i, out.Roots[0].Kind())
		}
		jsAssertStable(t, fe, path, src)
	}
}

// TestBrokenJSXNoPanic 覆盖损坏 JSX、孤立闭合标签与 fragment。
func TestBrokenJSXNoPanic(t *testing.T) {
	fe := jsFrontend(t)
	cases := []string{
		"<div><span>",
		"<>",
		"</div>",
		"<div attr=",
		"<<>>",
		"<div",
		"const el = <div><span></div>;",
		"<a b={1} c",
		"const el = </>;",
		"const el = <div className=;",
		"<Foo.Bar",
		"const el = <>text</>;",
	}
	partialSeen := false
	for i, src := range cases {
		path := "broken.jsx"
		unit := air.SourceUnit{
			Language: "js",
			Files:    []air.SourceFile{air.NewSourceFile(path, []byte(src))},
		}
		out := jsParseNoPanic(t, fe, path, src)
		_, diags, err := fe.Parse(unit)
		if err != nil {
			t.Fatalf("case %d Parse: %v", i, err)
		}
		if len(diags) > 0 || factsOf(out)["partial:true"] {
			partialSeen = true
		}
		jsAssertStable(t, fe, path, src)
	}
	if !partialSeen {
		t.Fatal("broken JSX produced neither diagnostics nor SemPartial")
	}
}

// TestDeepNestingNoPanic 覆盖超深括号、数组与成员链。
func TestDeepNestingNoPanic(t *testing.T) {
	fe := jsFrontend(t)
	cases := []string{
		strings.Repeat("(", 200) + "1" + strings.Repeat(")", 200) + ";",
		"[" + strings.Repeat("[", 200) + "1" + strings.Repeat("]", 200) + "];",
		"f" + strings.Repeat(".a", 300) + ";",
		strings.Repeat("if(x){", 120) + "y();" + strings.Repeat("}", 120),
		"const t = " + strings.Repeat("`a${", 60) + "1" +
			strings.Repeat("}`", 60) + ";",
	}
	for _, src := range cases {
		path := "deep.ts"
		jsParseNoPanic(t, fe, path, src)
		jsAssertStable(t, fe, path, src)
	}
}

// TestDeterministicByteChaos 用确定性字节序列覆盖非法字节与畸形 token 流。
func TestDeterministicByteChaos(t *testing.T) {
	fe := jsFrontend(t)
	for i := 0; i < 32; i++ {
		buf := make([]byte, 64+i*7)
		for j := range buf {
			buf[j] = byte((i*131 + j*37 + 11) % 256)
		}
		path := "bytes.js"
		jsParseNoPanic(t, fe, path, string(buf))
		jsAssertStable(t, fe, path, string(buf))
	}
}

// TestEmptyAndTriviaNoPanic 覆盖空源、纯空白、纯注释、BOM 与未闭合块注释。
func TestEmptyAndTriviaNoPanic(t *testing.T) {
	fe := jsFrontend(t)
	cases := []string{
		"",
		"   \n\t\n",
		"// only comment\n",
		"/* block */",
		"/* unterminated",
		"\xef\xbb\xbfconst x=1;",
		"//\n//\n//\n",
	}
	for _, src := range cases {
		path := "trivia.js"
		jsParseNoPanic(t, fe, path, src)
		jsAssertStable(t, fe, path, src)
	}
}

// TestUnterminatedStringsNoPanic 覆盖未闭合字符串与模板字符串。
func TestUnterminatedStringsNoPanic(t *testing.T) {
	fe := jsFrontend(t)
	cases := []string{
		"const a = 'unterminated",
		"const b = \"unterminated",
		"const c = `unterminated",
		"const d = `a ${b",
		"const e = '\\",
	}
	for _, src := range cases {
		path := "strings.js"
		jsParseNoPanic(t, fe, path, src)
		jsAssertStable(t, fe, path, src)
	}
}
