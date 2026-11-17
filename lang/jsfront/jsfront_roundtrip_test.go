/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package jsfront

import (
	"testing"

	"abate/lang/frontkit"
)

// jsRoundTripSource 是位置往返校验用的真实 JS/TS 样例。
const jsRoundTripSource = "import { a } from \"m\";\n\n" +
	"// tail\nclass C extends B {\n  m(x: number): string { return `v${x}`; }\n}\n" +
	"const t = [1, 2].map((v) => v * 2);\n"

// TestJsIntervalCoverage 证明 JS/TS 前端满足「源码 → AIR → 位置区间回放 →
// 区间覆盖检查」的往返契约：根覆盖整个文件，节点不越界不重叠，且每个保留的
// 词法 token 都被某个节点覆盖。
func TestJsIntervalCoverage(t *testing.T) {
	corpus := append(jsCoverageCorpus(), jsRoundTripSource)
	for _, src := range corpus {
		unit, _ := jsParse(t, jsFrontend(t), "f.ts", "", src)
		lexer, err := frontkit.NewLexer((&jsSpec{}).LexSpec(), []byte(src))
		if err != nil {
			t.Fatalf("NewLexer: %v", err)
		}
		tokens, err := lexer.Tokenize()
		if err != nil {
			t.Fatalf("Tokenize: %v", err)
		}
		root := unit.Roots[0]
		if err := frontkit.VerifyIntervalCoverage(root, tokens, len(src)); err != nil {
			t.Fatalf("interval coverage for %q: %v", src, err)
		}
	}
}
