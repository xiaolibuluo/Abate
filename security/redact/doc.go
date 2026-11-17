/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package redact provides deterministic, dependency-free secret redaction for
// log lines and URL paths. It is a single-pass, case-insensitive scanner built
// only on the Go standard library: it reads no clock, no random source, and no
// map iteration order, so identical inputs always yield identical outputs.
//
// The scanner is fail-safe. Unrecognized text is preserved byte for byte, while
// known secret shapes (Bearer tokens, Authorization headers, and key/value pairs
// such as token, secret, api_key, and password) have their values replaced with
// Placeholder. Text is intended for free-form log messages and panic values;
// Path additionally treats URL path separators as value terminators so the
// surrounding route structure survives redaction.
package redact
