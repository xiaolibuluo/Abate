/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package gomutate implements the real go test execution backend for the
// mutation kernel (KRN-MUT).
//
// 参考 go-gremlins / PITest 的真实变异执行（改写源码后跑测试判定生死），
// 改进点=接入 abate 的确定性内核契约：变异点按 (相对路径, 偏移, 记号) 全序
// 枚举，测试列表经 go test -list 收集后排序，覆盖度由逐测试 coverprofile
// 真实计算，评估结果按 (point, test) 缓存以满足 Target 接口的纯性要求。
//
// A mutant is killed by the real toolchain: its rewritten source fails to
// compile, or at least one selected test exits non-zero. A mutant survives
// only when every selected test passes against the mutated tree. No outcome
// is ever approximated, sampled or inferred statically.
package gomutate
