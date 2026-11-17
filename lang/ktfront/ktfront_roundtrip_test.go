/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package ktfront

import (
	"testing"

	"abate/lang/frontkit"
)

// ktRoundTripSource 是位置往返校验用的真实 Kotlin 样例。
const ktRoundTripSource = "package demo\n\nimport kotlin.math.abs\n\n" +
	"// tail\ndata class User(val name: String)\n\n" +
	"fun add(a: Int, b: Int): Int {\n    return a + b\n}\n"

// TestKtIntervalCoverage 证明 Kotlin 前端满足「源码 → AIR → 位置区间回放 →
// 区间覆盖检查」的往返契约：根覆盖整个文件，节点不越界不重叠，且每个保留的
// 词法 token 都被某个节点覆盖。
func TestKtIntervalCoverage(t *testing.T) {
	corpus := append(ktCoverageCorpus(), ktRoundTripSource)
	for _, src := range corpus {
		_, root, _ := ktParse(t, src)
		tokens := ktTokenize(t, src)
		if err := frontkit.VerifyIntervalCoverage(root, tokens, len(src)); err != nil {
			t.Fatalf("interval coverage for %q: %v", src, err)
		}
	}
}
