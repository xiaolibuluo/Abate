/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package coverage joins the six embedded upstream grammars with the explicit
// self-built Go construct inventory. Measured coverage is computed only from a
// real AIR parse through grammar.Load and coverage.Compute; the Go inventory is
// marked NOT_MEASURED rather than being presented as a seventh measured data
// set, and the StatusDegraded definition is never relaxed.
package coverage
