/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package example is a reference third-party integration for the CES
// Technology Compatibility Kit (CES-TCK).
//
// It demonstrates the compliance flow a downstream project follows: it builds
// CES documents with the public schema API (the producer side) and adapts the
// public parse/normalize/validate pipeline behind the tck.Driver interface (the
// consumer side). Running the bundled kit against this driver proves that an
// external implementation can be certified conformant without importing any
// Abate internal type. The kit itself is implementation-agnostic: it only
// exchanges neutral byte streams and Rule-ID sets.
package example
