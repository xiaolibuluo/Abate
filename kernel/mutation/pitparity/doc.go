/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package pitparity implements the point-by-point parity harness between the
// KRN-MUT mutation kernel and PITest.
//
// 参考 PITest 1.19.6 官方实跑产物（kernel/mutation/testdata/pit/），改进点=
// 三方独立判定互校：PIT 官方状态、Abate 引擎真实判定、语义见证 oracle 逐点对齐，
// 任何不一致都以差异条目如实上报，绝不通过改写映射掩盖语义不符。
// The harness never starts a JVM: the PIT side replays the vendored official
// report and the Abate side is produced by the real mutation engine, so the
// whole comparison runs offline with the standard library only.
package pitparity
