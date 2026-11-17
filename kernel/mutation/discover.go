/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package mutation

import (
	"slices"

	"abate/sim/rand"
)

// Discover enumerates every applicable mutation point of one synthetic target.
//
// 参考 PITest 变异点发现（遍历字节码并匹配算子），改进点=节点先按字典序排序、
// 每个候选点的选择子由 SIM-RNG 种子派生，消除原生非确定的点选择顺序。
// The returned slice is deterministic for one seed, target and suite triple.
func Discover(target Target, suite Suite, reg *Registry, rng *rand.Stream) ([]MutationPoint, error) {
	if target == nil {
		return nil, ErrNilTarget
	}
	if suite == nil {
		return nil, ErrNilSuite
	}
	if reg == nil {
		return nil, ErrNilRegistry
	}
	nodes := append([]string(nil), target.Nodes()...)
	if len(nodes) == 0 {
		return nil, ErrNoNodes
	}
	if len(suite.Tests()) == 0 {
		return nil, ErrNoTests
	}
	slices.Sort(nodes)
	points := make([]MutationPoint, 0, len(nodes)*4)
	for _, node := range nodes {
		token := target.Token(node)
		for _, op := range reg.For(token) {
			points = append(points, NewMutationPoint(node, token, op.Params()["to"], op,
				suite.Coverage(node), rng.Uint64()))
		}
	}
	if len(points) == 0 {
		return nil, ErrNoMutants
	}
	return points, nil
}
