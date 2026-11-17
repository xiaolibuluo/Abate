/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package coverage

import (
	"abate/lang/air"
	"abate/lang/frontkit/grammar"
)

// Degradation returns the degraded-mapping share of a language in percent,
// counting only exact and degraded mappings. StatusDegraded keeps its original
// meaning: a degraded mapping is never reclassified as exact to lower the
// ratio, and constructs outside the AIR capability domain do not hide a gap.
func Degradation(mappings []grammar.Mapping,
	observed map[air.NodeKind]int) float64 {
	total := 0
	degraded := 0
	for _, mapping := range mappings {
		switch mapping.Status {
		case grammar.StatusExact:
			total++
		case grammar.StatusDegraded:
			total++
			if observed[mapping.Kind] > 0 {
				degraded++
			}
		}
	}
	if total == 0 {
		return 0
	}
	return float64(degraded) * 100 / float64(total)
}
