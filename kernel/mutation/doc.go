/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package mutation implements the KRN-MUT mutation kernel as a deterministic,
// JVM-free Go rewrite of four PITest algorithm lines: operator families,
// coverage-guided scheduling, killing-test minimization and priority ordering.
//
// 参考 PITest 四条算法线（算子族 / 覆盖引导 / 最小化 / 优先级），
// 改进点=确定性化（SIM-RNG 种子驱动 + 全序裁决）+ 去 JVM + 合成 Target 端口。
// Every decision is derived from one seed, sorted keys and integer counters, so
// the same seed always yields a byte-identical Report and no PIT-specific type
// ever enters the public API.
package mutation
