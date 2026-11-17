/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package pitparity implements the self-developed golden semantic regression
// of the KRN-MUT mutation kernel.
//
// 本包以 kernel/mutation/testdata/pit/golden/mutation-semantics.golden 为权威冻结数据源，
// 由本仓变异核对同构夹具的真实判定与自研语义见证逐点互校，复跑可逐字节复现。
// 这是纯自研 golden 语义回归，不是 PIT 官方对拍；竞品对拍分数维持 NOT_MEASURED。
// 全流程只用标准库、不启动 JVM、不读取任何上游产物。
package pitparity
