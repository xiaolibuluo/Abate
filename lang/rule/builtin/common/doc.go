/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package common declares the cross-language AIR rule template family.
//
// 参考 PEP 8 / TS Handbook / CERT / Rust API Guidelines / .NET Design /
// Kotlin Conventions 的规则分类法，改进点=以 AIR 结构模式统一表达，不含语言专有逻辑。
// The declarations are pure deterministic data rebuilt on every call, and the
// generator binds each template to a traceable archived specification section.
package common
