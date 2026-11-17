/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package pitparity

import "strconv"

// observation is one observable value produced by the baseline project.
type observation string

// The helpers below are a faithful Go port of the Java parity baseline that
// ships in kernel/mutation/testdata/pit/baseline/.
//
// The port is anchored by the assertions of ParityTargetTest.java: every
// original observation must reproduce the value the Java suite asserts, which
// is what makes the semantic witness trustworthy instead of hand written.
func obsInt(value int) observation { return observation(strconv.Itoa(value)) }

// baselineAdd observes add(a, b); the suite fixes the second operand at 5.
func baselineAdd(a int) observation { return obsInt(a + 5) }

// baselineClamp observes clamp(x) with its strict boundary at 10.
func baselineClamp(x int) observation {
	if x < 10 {
		return obsInt(x)
	}
	return obsInt(10)
}

// baselineCap observes cap(x) with the constant limit at 7.
func baselineCap(x int) observation {
	const limit = 7
	if x > limit {
		return obsInt(limit)
	}
	return obsInt(x)
}

// baselineSumUntil observes sumUntil(n) with the break guard at i == 3.
func baselineSumUntil(n int) observation {
	sum := 0
	for i := 0; i < n; i++ {
		if i == 3 {
			break
		}
		sum += i
	}
	return obsInt(sum)
}

// baselineAccepts observes accepts(score) with the threshold at 5.
func baselineAccepts(score int) observation {
	if score >= 5 {
		return "true"
	}
	return "false"
}

// baselineLabel observes label(v) with the equality guard at v == 1.
func baselineLabel(v int) observation {
	if v == 1 {
		return "one"
	}
	return "other"
}

// baselineAudited observes the log counter reached by audited(n) calls.
func baselineAudited(n int) observation { return obsInt(n) }

// baselineLogCalls observes the counter exposed by logCalls() after the suite
// has called audited(2).
func baselineLogCalls(_ int) observation { return "2" }
