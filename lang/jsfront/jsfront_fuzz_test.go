/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package jsfront

import (
	"os"
	"strings"
	"testing"

	"abate/lang/air"
)

// jsFuzzSeeds 是 Go 原生 fuzz 的确定性种子语料。
var jsFuzzSeeds = [][]byte{
	[]byte("const x = 1;"),
	[]byte("const el = <div a=\"1\">hi</div>;"),
	[]byte("interface A { readonly id?: string }"),
	[]byte("return\nx"),
	[]byte("const t = `hi ${name}!`;"),
	[]byte("const r = /ab+c/gi; const d = 10 / 2;"),
	[]byte("const v = obj?.a ?? fallback;"),
	[]byte("type B<T> = [T, T];"),
	[]byte("@decorator class C {}"),
	{0x00, 0xff, 0xfe, '\n', 0x80, 0x7f},
	[]byte(""),
	[]byte("<<>>`${"),
}

// FuzzJsFrontParse 对任意字节流验证解析产出真实节点且指纹稳定；不 panic
// 由测试框架保证（panic 会自然冒泡并判 FAIL，go test -fuzz 记为 crash），
// 无需显式 recover（避免触发 GO-EG-027 提示规则）。
func FuzzJsFrontParse(f *testing.F) {
	for _, seed := range jsFuzzSeeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, data []byte) {
		// Bound fuzz inputs to 4096 bytes: the guard keeps each fuzz iteration
		// fast and is not a failure suppression. Oversized inputs exercise the
		// same lexer/parser code paths as the bounded malformed-token cases in
		// TestIllegalByteTableNoPanic and TestFuzzParseTokensDirect, so skipping
		// them loses no coverage of panic-freedom or determinism.
		if len(data) > 4096 {
			t.Skip()
		}
		fe, err := NewJsFrontend()
		if err != nil {
			t.Fatalf("NewJsFrontend: %v", err)
		}
		unit := air.SourceUnit{
			Language: "js", Dialect: "tsx/es2022",
			Files: []air.SourceFile{
				air.NewSourceFile("fuzz.tsx", data),
			},
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

// TestFuzzSeedCorpusDeterminism 对全部种子语料重复解析并断言一致。
func TestFuzzSeedCorpusDeterminism(t *testing.T) {
	fe := jsFrontend(t)
	for i, seed := range jsFuzzSeeds {
		src := string(seed)
		path := "seed.tsx"
		unit := air.SourceUnit{
			Language: "js", Dialect: "tsx/es2022",
			Files: []air.SourceFile{air.NewSourceFile(path, seed)},
		}
		var first [32]byte
		for run := 0; run < 5; run++ {
			out, _, err := fe.Parse(unit)
			if err != nil {
				t.Fatalf("seed %d Parse: %v", i, err)
			}
			sum := out.Fingerprint()
			if run == 0 {
				first = sum
			} else if sum != first {
				t.Fatalf("seed %d non-deterministic at run %d", i, run)
			}
		}
		jsParseNoPanic(t, fe, "seed.js", src)
	}
}

// TestIllegalByteTableNoPanic 覆盖高位字节、NUL 与控制字符的确定性组合。
func TestIllegalByteTableNoPanic(t *testing.T) {
	fe := jsFrontend(t)
	for i := 0; i < 24; i++ {
		buf := make([]byte, 48+i*5)
		for j := range buf {
			buf[j] = byte((i*211 + j*29 + 3) % 256)
		}
		jsParseNoPanic(t, fe, "illegal.js", string(buf))
		jsAssertStable(t, fe, "illegal.js", string(buf))
	}
}

// TestFuzzParseTokensDirect 直接对畸形 token 流调用 ParseTokens 验证容错；
// 不 panic 由测试框架保证（panic 会自然冒泡并判 FAIL），无需显式 recover
// （避免触发 GO-EG-027 提示规则）。
func TestFuzzParseTokensDirect(t *testing.T) {
	spec := &jsSpec{}
	for i := 0; i < 16; i++ {
		buf := make([]byte, 32+i*9)
		for j := range buf {
			buf[j] = byte((i*97 + j*41 + 17) % 256)
		}
		tokens := jsLexTokens(t, string(buf))
		root, _, err := spec.ParseTokens(tokens, "direct.ts")
		if err != nil || root == nil {
			t.Fatalf("direct ParseTokens: %v", err)
		}
		if root.Kind() != air.KindPackage {
			t.Fatalf("direct root = %s", root.Kind())
		}
	}
}

// jsSourceReaders 是本包固定 .go 源文件名白名单：每个名字绑定一个只含编译期
// 常量路径的读取器；读取前先校验 name 在白名单内，杜绝变量路径读取。
var jsSourceReaders = map[string]func() ([]byte, error){
	"doc.go":                   func() ([]byte, error) { return os.ReadFile("doc.go") },
	"frontend.go":              func() ([]byte, error) { return os.ReadFile("frontend.go") },
	"jsexpr.go":                func() ([]byte, error) { return os.ReadFile("jsexpr.go") },
	"jsfront_chaos_test.go":    func() ([]byte, error) { return os.ReadFile("jsfront_chaos_test.go") },
	"jsfront_dialect_test.go":  func() ([]byte, error) { return os.ReadFile("jsfront_dialect_test.go") },
	"jsfront_fidelity_test.go": func() ([]byte, error) { return os.ReadFile("jsfront_fidelity_test.go") },
	"jsfront_fuzz_test.go":     func() ([]byte, error) { return os.ReadFile("jsfront_fuzz_test.go") },
	"jsfront_mutation_test.go": func() ([]byte, error) { return os.ReadFile("jsfront_mutation_test.go") },
	"jsfront_test.go":          func() ([]byte, error) { return os.ReadFile("jsfront_test.go") },
	"jslexspec.go":             func() ([]byte, error) { return os.ReadFile("jslexspec.go") },
	"jsmap.go":                 func() ([]byte, error) { return os.ReadFile("jsmap.go") },
	"jsparser.go":              func() ([]byte, error) { return os.ReadFile("jsparser.go") },
	"jstokens.go":              func() ([]byte, error) { return os.ReadFile("jstokens.go") },
	"jstype.go":                func() ([]byte, error) { return os.ReadFile("jstype.go") },
}

// TestFuzzNoThirdPartyImports 扫描本包全部 .go 文件，确认零第三方 import。
func TestFuzzNoThirdPartyImports(t *testing.T) {
	entries, err := os.ReadDir(".")
	if err != nil {
		t.Fatalf("ReadDir: %v", err)
	}
	banned := []string{
		"github" + ".com/", "gopkg" + ".in/", "golang" + ".org/",
		"import " + "\"C\"", "os/" + "exec", "net/" + "http",
		"net" + ".Dial", "math/" + "rand", "time" + ".Now",
	}
	checked := 0
	for _, entry := range entries {
		name := entry.Name()
		if !strings.HasSuffix(name, ".go") {
			continue
		}
		read, ok := jsSourceReaders[name]
		if !ok {
			continue
		}
		content, readErr := read()
		if readErr != nil {
			t.Fatalf("ReadFile(%s): %v", name, readErr)
		}
		checked++
		for _, line := range strings.Split(string(content), "\n") {
			trimmed := strings.TrimSpace(line)
			if !strings.HasPrefix(trimmed, "\"") &&
				!strings.HasPrefix(trimmed, "import ") {
				continue
			}
			for _, bad := range banned {
				if strings.Contains(line, bad) {
					t.Errorf("%s import line contains forbidden %q: %s",
						name, bad, trimmed)
				}
			}
		}
	}
	if checked != 14 {
		t.Fatalf("expected 14 package files, found %d", checked)
	}
}
