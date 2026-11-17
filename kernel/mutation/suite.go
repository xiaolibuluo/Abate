/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package mutation

// Suite is the synthetic test-selection port that stands in for LNG-AIR.
//
// 参考 PITest 覆盖率引导（测试覆盖矩阵），改进点=覆盖度以整数计数表达，
// 无浮点、无线程优先级、无 map 迭代序泄漏。
// Tests returns identifiers in a fixed order and Coverage is a pure integer
// count, so Schedule can resolve a total order without wall-clock or RNG races.
type Suite interface {
	// Fingerprint returns the stable digest of the test suite.
	Fingerprint() [32]byte

	// Tests returns the test identifiers in deterministic order.
	Tests() []string

	// Coverage returns the integer number of tests covering one node.
	Coverage(node string) uint64
}
