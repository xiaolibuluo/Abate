/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package pitparity

// The functions below observe the mutated baseline methods. Each one models
// exactly the construct the PIT mutator replaces, so the witness can decide
// whether the Java suite distinguishes the mutant from the original.

// addSubtracted observes add with the addition replaced by a subtraction.
func addSubtracted(a int) observation { return obsInt(a - 5) }

// addZeroReturn observes add with the integer return replaced by zero.
func addZeroReturn(_ int) observation { return "0" }

// clampBoundary observes clamp with x < 10 shifted to x <= 10.
func clampBoundary(x int) observation {
	if x <= 10 {
		return obsInt(x)
	}
	return obsInt(10)
}

// clampOrderElse observes clamp with the comparison removed, so the method
// always returns the constant branch.
func clampOrderElse(_ int) observation { return obsInt(10) }

// clampZeroReturn observes clamp with the integer return replaced by zero.
func clampZeroReturn(_ int) observation { return "0" }

// capBoundary observes cap with x > limit shifted to x >= limit.
func capBoundary(x int) observation {
	const limit = 7
	if x >= limit {
		return obsInt(limit)
	}
	return obsInt(x)
}

// capOrderElse observes cap with the comparison removed, so the method always
// returns its argument.
func capOrderElse(x int) observation { return obsInt(x) }

// capZeroReturn observes cap with the integer return replaced by zero.
func capZeroReturn(_ int) observation { return "0" }

// sumUntilBoundary observes sumUntil with i < n shifted to i <= n.
func sumUntilBoundary(n int) observation {
	sum := 0
	for i := 0; i <= n; i++ {
		if i == 3 {
			break
		}
		sum += i
	}
	return obsInt(sum)
}

// sumUntilOrderElse observes sumUntil with the loop comparison removed, so the
// loop body never runs.
func sumUntilOrderElse(_ int) observation { return obsInt(0) }

// sumUntilNoGuard observes sumUntil with the break guard removed.
func sumUntilNoGuard(n int) observation {
	sum := 0
	for i := 0; i < n; i++ {
		sum += i
	}
	return obsInt(sum)
}

// sumUntilSubtracted observes sumUntil with the accumulation replaced by a
// subtraction.
func sumUntilSubtracted(n int) observation {
	sum := 0
	for i := 0; i < n; i++ {
		if i == 3 {
			break
		}
		sum -= i
	}
	return obsInt(sum)
}

// sumUntilZeroReturn observes sumUntil with the integer return replaced by zero.
func sumUntilZeroReturn(_ int) observation { return "0" }

// acceptsBoundary observes accepts with score >= 5 shifted to score > 5.
func acceptsBoundary(score int) observation {
	if score > 5 {
		return "true"
	}
	return "false"
}

// acceptsOrderElse observes accepts with the comparison removed.
func acceptsOrderElse(_ int) observation { return "false" }

// acceptsTrueReturn observes accepts with the boolean return replaced by true.
func acceptsTrueReturn(_ int) observation { return "true" }

// labelNoGuard observes label with the equality guard removed.
func labelNoGuard(_ int) observation { return "other" }

// labelEmptyReturn observes label with the string return replaced by an empty
// string.
func labelEmptyReturn(_ int) observation { return "" }

// auditedBoundary observes audited with i < n shifted to i <= n.
func auditedBoundary(n int) observation { return obsInt(n + 1) }

// auditedOrderElse observes audited with the loop comparison removed.
func auditedOrderElse(_ int) observation { return obsInt(0) }

// auditedNoCall observes audited with the void call removed.
func auditedNoCall(_ int) observation { return obsInt(0) }

// logSubtracted observes the log counter with the increment replaced by a
// decrement.
func logSubtracted(n int) observation { return obsInt(-n) }

// logCallsZeroReturn observes logCalls with the integer return replaced by zero.
func logCallsZeroReturn(_ int) observation { return "0" }

// uncoveredDivided observes uncovered with x * 3 replaced by x / 3.
func uncoveredDivided(x int) observation { return obsInt(x / 3) }

// uncoveredZeroReturn observes uncovered with the integer return replaced by
// zero.
func uncoveredZeroReturn(_ int) observation { return "0" }
