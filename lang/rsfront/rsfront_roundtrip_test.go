/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package rsfront

import (
	"testing"

	"abate/lang/frontkit"
)

// rsRoundTripSource 是位置往返校验用的真实 Rust 样例。
const rsRoundTripSource = "use std::fmt;\n\n/// doc\n#[derive(Clone)]\n" +
	"pub struct Point { pub x: i32 }\n\n" +
	"fn add(a: i32, b: i32) -> i32 {\n    // tail\n    a + b\n}\n"

// TestRsIntervalCoverage 证明 Rust 前端满足「源码 → AIR → 位置区间回放 →
// 区间覆盖检查」的往返契约：根覆盖整个文件，节点不越界不重叠，且每个保留的
// 词法 token 都被某个节点覆盖。
func TestRsIntervalCoverage(t *testing.T) {
	corpus := append(rsCoverageCorpus(), rsRoundTripSource)
	for _, src := range corpus {
		unit := rsParseOne(t, src)
		tokens := rsTokenize(t, src)
		root := unit.Roots[0]
		if err := frontkit.VerifyIntervalCoverage(root, tokens, len(src)); err != nil {
			t.Fatalf("interval coverage for %q: %v", src, err)
		}
	}
}
