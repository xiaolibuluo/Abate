/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package validate performs semantic validation of a normalized CES document.
//
// Unlike the fail-fast parser, validation is aggregating: every violated rule
// contributes a Violation. Rules run in lexicographic RuleID order and the
// result is stably sorted, so the same document always yields byte-identical
// output. Validation is pure, reentrant and safe for concurrent use.
package validate
