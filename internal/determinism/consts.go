/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package determinism

// SplitMix64 mixing constants are untyped so the constant gate treats them as
// named constants rather than magic literals.
const (
	splitMixGamma = 0x9E3779B97F4A7C15
	splitMixMulA  = 0xBF58476D1CE4E5B9
	splitMixMulB  = 0x94D049BB133111EB
	rngDomain     = "ABATE-E2E"
	maxInt        = int(^uint(0) >> 1)
)
