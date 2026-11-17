/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package csfront

import (
	"testing"

	"abate/lang/frontkit"
)

// csRoundTripSource 是位置往返校验用的真实 C# 样例。
const csRoundTripSource = "using System;\n\nnamespace Demo {\n  // tail\n" +
	"  public class C {\n    public int M(int x) { return x + 1; }\n  }\n}\n"

// TestCsIntervalCoverage 证明 C# 前端满足「源码 → AIR → 位置区间回放 →
// 区间覆盖检查」的往返契约：根覆盖整个文件，节点不越界不重叠，且每个保留的
// 词法 token 都被某个节点覆盖。
func TestCsIntervalCoverage(t *testing.T) {
	corpus := append(csCoverageCorpus(), csRoundTripSource)
	for _, src := range corpus {
		_, root, _ := csParse(t, src)
		tokens := csTokenize(t, src)
		if err := frontkit.VerifyIntervalCoverage(root, tokens, len(src)); err != nil {
			t.Fatalf("interval coverage for %q: %v", src, err)
		}
	}
}
