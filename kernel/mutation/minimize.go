/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package mutation

// KillPredicate reports whether one test kills one mutant.
type KillPredicate func(point MutationPoint, test string) bool

// Minimize locates the first killing test by a deterministic binary search.
//
// 参考 PITest 最小化（线性回退寻找首个击杀测试），改进点=确定性二分下界，
// 在按击杀能力排序的合成套件上把探测次数从 O(n) 降为 O(log n)。
// It returns the index in tests, or -1 when no test kills the mutant.
func Minimize(point MutationPoint, tests []string, killed KillPredicate) (int, error) {
	if killed == nil {
		return -1, ErrNilKill
	}
	lo, hi := 0, len(tests)
	for lo < hi {
		mid := lo + (hi-lo)/2
		if killed(point, tests[mid]) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	if lo < len(tests) && killed(point, tests[lo]) {
		return lo, nil
	}
	return -1, nil
}
