/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package corpusgen

// GenerateStats records the raw output numbers used in the delivery report.
type GenerateStats struct {
	Total       int
	PerLanguage map[string]int
	PerStratum  map[string]int
	Bytes       int64
	Broken      int
}
