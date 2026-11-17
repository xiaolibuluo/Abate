/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package schema defines the CES (Chaos Experiment Specification) data model.
//
// It is the single source of truth for the parser, the validator and the
// schema generators. Every duration is an integer nanosecond value so that
// JSON and YAML encodings round-trip byte for byte.
package schema
