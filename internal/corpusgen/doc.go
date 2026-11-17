/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package corpusgen produces the deterministic, fully self-researched E2E
// corpus. It depends only on the Go standard library and Abate's own
// internal/determinism package. The frozen plan, thresholds and gate scripts
// are read-only inputs and are never rewritten by this package.
package corpusgen
