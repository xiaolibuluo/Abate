/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package cbuiltin declares the built-in C and C++ rule set for the Abate
// checker.
//
// The rule set is pure AIR structural data spanning CERT C and the C++ Core
// Guidelines. Every rule is AIR-facing: no clang or libclang type appears in the
// public contract, and the shared adapter parses in-memory C/C++ source buffers
// into an air.Unit for matching through the same matcher used by every language.
// 参考 CERT C / C++ Core Guidelines 或 Rust API Guidelines / rustfmt 归档规范，改进点=以共享 AIR 结构模式生成并叠加 C / C++ 专有规则。
package cbuiltin
