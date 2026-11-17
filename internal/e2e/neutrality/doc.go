/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package neutrality performs a zero-exemption go/ast reverse static check of
// every .go file registered in ci/e2e_sources.txt. Eight syntax-level
// invariants reject provenance discrimination, corpus-loop skips, type-check
// bypasses, recover swallowing, deterministic red lines, deep exits outside
// command mains, boolean flag parameters and early corpus-loop returns.
package neutrality
