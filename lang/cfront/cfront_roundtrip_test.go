/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package cfront

import (
	"testing"

	"abate/lang/frontkit"
)

// cRoundTripSource 是位置往返校验用的真实 C/C++ 样例。
const cRoundTripSource = "#include <stdio.h>\n// tail\nnamespace demo {\n" +
	"struct Point { int x; int y; };\n" +
	"int add(int a, int b) {\n    return a + b;\n}\n" +
	"}\n"

// TestCIntervalCoverage 证明 C/C++ 前端满足「源码 → AIR → 位置区间回放 →
// 区间覆盖检查」的往返契约：根覆盖整个文件，节点不越界不重叠，且每个保留的
// 词法 token 都被某个节点覆盖。
func TestCIntervalCoverage(t *testing.T) {
	corpus := append(cCoverageCorpus(), cRoundTripSource)
	for _, src := range corpus {
		unit, _ := cParseOne(t, src, "c++20")
		tokens := cLex(t, src)
		root := unit.Roots[0]
		if err := frontkit.VerifyIntervalCoverage(root, tokens, len(src)); err != nil {
			t.Fatalf("interval coverage for %q: %v", src, err)
		}
	}
}
