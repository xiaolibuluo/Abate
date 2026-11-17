/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package determinism

// Stride returns want indices equally spaced over [0, total) in ascending
// order. Both endpoints are included whenever want is at least two, and the
// result depends only on total and want, never on any random generator.
func Stride(total, want int) []int {
	if total <= 0 || want <= 0 {
		return []int{}
	}
	if want >= total {
		out := make([]int, total)
		for i := range out {
			out[i] = i
		}
		return out
	}
	if want == 1 {
		return []int{0}
	}
	out := make([]int, want)
	last := want - 1
	for i := 0; i < want; i++ {
		out[i] = i * (total - 1) / last
	}
	return out
}
