/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package jsfront

import (
	"strconv"
	"testing"

	"abate/lang/air"
	"abate/lang/frontkit"
)

// jsFidelityCases 是 token 保真抽查的源码矩阵。
var jsFidelityCases = []struct {
	name string
	path string
	src  string
}{
	{"basic", "f.js", "const x = 1; let s = \"hi\"; // tail\n/* block */\nvar q = x + 2;"},
	{"ts", "f.ts", "interface A { readonly id?: string } type B<T> = [T, T];"},
	{"jsx", "f.tsx", "const el = <div className=\"x\">hi</div>;"},
	{"template", "f.js", "const t = `hi ${name}!`;"},
	{"regex", "f.js", "const r = /ab+c/gi; const d = 10 / 2;"},
	{"chain", "f.ts", "const v = obj?.a?.[0] ?? fallback;"},
	{"comprehensive", "all.tsx", jsComprehensiveSource()},
}

// jsLexicalNodeKind 复制保真子集的 AIR 节点类判定。
func jsLexicalNodeKind(kind air.NodeKind) bool {
	switch kind {
	case air.KindIdent, air.KindNumLit, air.KindStrLit, air.KindCharLit,
		air.KindBoolLit, air.KindNilLit, air.KindComment, air.KindDoc,
		air.KindDirective, air.KindAttribute:
		return true
	default:
		return false
	}
}

// jsTokenKey 渲染 token 的字节精确比较键。
func jsTokenKey(offset, endOffset int, text string) string {
	return strconv.Itoa(offset) + "\x00" + strconv.Itoa(endOffset) + "\x00" + text
}

// jsSourceLexicalTokens 返回真实 token 流中属于保真子集的 token。
func jsSourceLexicalTokens(tokens []frontkit.Token) []frontkit.Token {
	out := make([]frontkit.Token, 0, len(tokens))
	for _, tok := range tokens {
		if !jsIsLexicalToken(tok.Kind) || tok.Text == "" {
			continue
		}
		out = append(out, tok)
	}
	return out
}

// jsProjectedKeys 返回 AIR 里全部 lexical 投影节点的比较键。
func jsProjectedKeys(u *air.Unit) map[string]bool {
	keys := make(map[string]bool)
	for _, n := range air.Collect(u) {
		if !jsLexicalNodeKind(n.Kind()) {
			continue
		}
		text := air.AttributeOf(n)
		if text == "" {
			continue
		}
		p := n.Pos()
		keys[jsTokenKey(p.Offset, p.EndOffset, text)] = true
	}
	return keys
}

// jsLexTokens 用真实词法表切分源码，失败即终止。
func jsLexTokens(t *testing.T, src string) []frontkit.Token {
	t.Helper()
	lexer, err := frontkit.NewLexer((&jsSpec{}).LexSpec(), []byte(src))
	if err != nil {
		t.Fatalf("NewLexer: %v", err)
	}
	tokens, err := lexer.Tokenize()
	if err != nil {
		t.Fatalf("Tokenize: %v", err)
	}
	return tokens
}

// TestTokenFidelity 验证双向 token 保真与位置区间覆盖。
func TestTokenFidelity(t *testing.T) {
	fe := jsFrontend(t)
	for _, tc := range jsFidelityCases {
		t.Run(tc.name, func(t *testing.T) {
			tokens := jsLexTokens(t, tc.src)
			out, _ := jsParse(t, fe, tc.path, "es2022", tc.src)
			if err := frontkit.VerifyTokenFidelity(tokens, out); err != nil {
				t.Fatalf("VerifyTokenFidelity: %v", err)
			}
			if err := frontkit.VerifyNodeFidelity(tokens, out.Roots[0]); err != nil {
				t.Fatalf("VerifyNodeFidelity: %v", err)
			}
			projected := jsProjectedKeys(out)
			for _, tok := range jsSourceLexicalTokens(tokens) {
				key := jsTokenKey(tok.Offset, tok.EndOffset, tok.Text)
				if !projected[key] {
					t.Errorf("source %s %q at [%d,%d) not projected",
						tok.Kind, tok.Text, tok.Offset, tok.EndOffset)
				}
			}
		})
	}
}

// TestFidelityAcrossDialects 对六种扩展名逐一验证 token 保真。
func TestFidelityAcrossDialects(t *testing.T) {
	fe := jsFrontend(t)
	src := "// lead\nconst x: number = 1;\nconst el = <section id=\"a\">x</section>;\n" +
		"const t = `v=${x}`;\n/* tail */\n"
	for _, ext := range jsDialectExtensions {
		path := "dialect" + ext
		tokens := jsLexTokens(t, src)
		out, _ := jsParse(t, fe, path, "", src)
		if err := frontkit.VerifyTokenFidelity(tokens, out); err != nil {
			t.Errorf("%s: %v", ext, err)
		}
		if out.NodeCount() < 10 {
			t.Errorf("%s: node count %d", ext, out.NodeCount())
		}
	}
}

// TestCommentAndTemplateProjection 验证注释与模板字符串精确投影到 AIR。
func TestCommentAndTemplateProjection(t *testing.T) {
	fe := jsFrontend(t)
	src := "// one\nconst a = 'single'; /* two */\nconst t = `x ${a} y`;\n"
	out, _ := jsParse(t, fe, "projection.js", "", src)
	facts := factsOf(out)
	if !facts["kind:Comment"] || !facts["kind:StrLit"] {
		t.Fatalf("missing comment/string projection: %v", sortedFactKeys(facts))
	}
	if !facts["note:template-string"] || !facts["note:template-interpolation"] {
		t.Fatal("template interpolation not preserved")
	}
	projected := jsProjectedKeys(out)
	for _, tok := range jsSourceLexicalTokens(jsLexTokens(t, src)) {
		if !projected[jsTokenKey(tok.Offset, tok.EndOffset, tok.Text)] {
			t.Errorf("token %q not projected", tok.Text)
		}
	}
}

// TestPositionIntervalsAreValid 验证所有 lexical 投影区间合法且落在源码范围内。
func TestPositionIntervalsAreValid(t *testing.T) {
	fe := jsFrontend(t)
	src := jsComprehensiveSource()
	out, _ := jsParse(t, fe, "positions.tsx", "", src)
	bound := len(src)
	for _, n := range air.Collect(out) {
		if !jsLexicalNodeKind(n.Kind()) || air.AttributeOf(n) == "" {
			continue
		}
		p := n.Pos()
		if !p.Valid() || p.EndOffset < p.Offset || p.EndOffset > bound {
			t.Fatalf("invalid lexical interval %s (bound %d)", p.String(), bound)
		}
	}
}

// TestFidelityEmptyUnit 验证空源码仍产出合法根且保真通过。
func TestFidelityEmptyUnit(t *testing.T) {
	fe := jsFrontend(t)
	tokens := jsLexTokens(t, "")
	out, _ := jsParse(t, fe, "empty.js", "", "")
	if out.NodeCount() < 1 {
		t.Fatal("empty source produced no root")
	}
	if err := frontkit.VerifyTokenFidelity(tokens, out); err != nil {
		t.Fatalf("empty fidelity: %v", err)
	}
}

// TestJsOperatorFidelity asserts AIR-ID changes for cmp, op and digit mutations, with identical trailing comments on both variants.
func TestJsOperatorFidelity(t *testing.T) {
	fe := jsFrontend(t)
	cases := []struct{ name, base, mutated string }{
		{"cmp", "const x = 1 == 2;\n", "const x = 1 != 2;\n"},
		{"op", "const x = 1 + 2;\n", "const x = 1 - 2;\n"},
		{"digit", "const x = 41;\n", "const x = 42;\n"},
	}
	rootID := func(t *testing.T, src string) air.NodeID {
		t.Helper()
		out, _ := jsParse(t, fe, "operator.js", "es2022", src)
		return air.IDOf(out.Roots[0])
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
