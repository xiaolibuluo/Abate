/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package mutation

import "abate/kernel/verify"

// Target is the synthetic code-space port that stands in for LNG-AIR in M2.
//
// 参考 PITest 变异点发现（MutationDetails 遍历字节码），改进点=暴露与语言无关的
// 节点端口（节点 ID + 记号），不引用任何 JVM 或 PIT 专有类型。
// Implementations must be pure: the same point and test always yield the same
// outcome, so Engine.Run stays byte-stable under replay and concurrency.
type Target interface {
	// Fingerprint returns the stable digest of the target program.
	Fingerprint() [32]byte

	// Nodes returns the mutable node references in deterministic order.
	Nodes() []string

	// Token returns the original lexical token of one node.
	Token(node string) string

	// Equivalent reports whether a mutant is semantically equivalent.
	Equivalent(point MutationPoint) bool

	// Evaluate reports the steady-state outcome of one test on one mutant.
	Evaluate(point MutationPoint, test string) verify.Outcome
}
