/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package mutation

import (
	"bytes"
	"slices"
)

// Schedule orders mutation points into a coverage-guided total order.
//
// 参考 PITest 的线程优先级调度，改进点=整数覆盖计数 + 全序裁决 + 去重前先排
// 指纹键，彻底移除线程竞争、浮点权重与 map 迭代序泄漏。
// Duplicate fingerprints are collapsed, so the result is a strict total order.
func Schedule(points []MutationPoint) []MutationPoint {
	byFingerprint := make(map[[32]byte]MutationPoint, len(points))
	keys := make([][32]byte, 0, len(points))
	for _, point := range points {
		if _, exists := byFingerprint[point.Fingerprint]; !exists {
			keys = append(keys, point.Fingerprint)
		}
		byFingerprint[point.Fingerprint] = point
	}
	slices.SortFunc(keys, func(a, b [32]byte) int {
		return bytes.Compare(a[:], b[:])
	})
	out := make([]MutationPoint, 0, len(keys))
	for _, key := range keys {
		out = append(out, byFingerprint[key])
	}
	slices.SortFunc(out, func(a, b MutationPoint) int {
		if a.Less(b) {
			return -1
		}
		if b.Less(a) {
			return 1
		}
		return 0
	})
	return out
}
