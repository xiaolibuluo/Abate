/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package check orchestrates the language-neutral checking pipeline: paths to
// source buffers, a Frontend to AIR, a set of compiled Rules to findings, then
// deterministic aggregation and reporting.
//
// The package contains no language-specific logic. It depends only on the AIR
// Frontend contract and the AIR-facing Rule contract, and emits deterministic
// console, JSON and minimal SARIF 2.1.0 renderings.
package check
