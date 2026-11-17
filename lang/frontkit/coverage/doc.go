/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package coverage measures how much of a language the Abate frontends
// really implement.
//
// The measurement joins two independent facts: the upstream node inventory of
// the language, which names every construct the grammar defines, and the AIR
// kinds a real parse of the language produced. A construct counts as covered
// only when the AIR kind it is lowered to actually shows up in a parsed tree,
// so the numbers come from an executed frontend instead of from a table.
//
// The granularity is the AIR node kind, because AIR is a neutral intermediate
// representation: several upstream constructs legitimately share one AIR
// kind. Every construct is still listed on its own row, and the status of its
// lowering states whether the mapping is exact or degraded, so a reader can
// tell a dedicated node from a shared one.
package coverage
