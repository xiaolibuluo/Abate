/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package incremental probes the incremental-parse contract. Its Probe calls
// ParseIncremental and checks that files outside the edited byte ranges keep
// their AIR root identifiers, while reporting how many distinct ranges were
// edited. It is a P1 verification layer, not a replacement for a full parse.
package incremental
