/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package rsbuiltin declares the built-in Rust rule set for the Abate checker.
//
// The rule set is pure AIR structural data spanning the Rust API Guidelines and
// the rustfmt archive conventions. Every rule is AIR-facing: no rustc or syn type
// appears in the public contract, and the shared adapter parses in-memory Rust
// source buffers into an air.Unit for matching through the same matcher used by
// every language.
// 参考 CERT C / C++ Core Guidelines 或 Rust API Guidelines / rustfmt 归档规范，改进点=以共享 AIR 结构模式生成并叠加 Rust 专有规则。
package rsbuiltin
