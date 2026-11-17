/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package verify

import (
	"testing"

	"abate/kernel/observe"
)

// TestNumericHelpers covers parsing fallbacks and unknown expectation types.
func TestNumericHelpers(t *testing.T) {
	t.Parallel()
	if n, ok := numeric(observe.NewText("12.5")); !ok || n != 12.5 {
		t.Fatalf("text numeric = (%g, %v)", n, ok)
	}
	if _, ok := numeric(observe.NewText("abc")); ok {
		t.Fatalf("non-numeric text must be rejected")
	}
	if _, ok := numeric(observe.NewBoolean(true)); ok {
		t.Fatalf("boolean must be rejected")
	}
	if _, ok := numeric(observe.NewHistogram([]uint64{1})); ok {
		t.Fatalf("histogram must be rejected")
	}
	if numericOrZero(observe.NewBoolean(true)) != 0 {
		t.Fatalf("numericOrZero fallback broken")
	}
	if _, _, ok := parsePredicate(">=abc"); ok {
		t.Fatalf("invalid predicate threshold must be rejected")
	}
	if _, _, ok := parsePredicate("!5"); ok {
		t.Fatalf("invalid predicate operator must be rejected")
	}
	if _, _, ok := parsePredicate(">=5"); !ok {
		t.Fatalf("valid predicate must parse")
	}
	if evalPredicate(observe.NewBoolean(true), ">5") {
		t.Fatalf("non-numeric predicate value must fail")
	}
	if evalPredicate(observe.NewNumber(6), "???") {
		t.Fatalf("malformed predicate must fail")
	}
	if matches(observe.NewBoolean(true), Expectation{Type: ExpectValue, Value: 1}, Tolerance{}) {
		t.Fatalf("non-numeric value expectation must fail")
	}
	if matches(observe.NewBoolean(true), Expectation{Type: ExpectRange, Min: 1, Max: 2}, Tolerance{}) {
		t.Fatalf("non-numeric range expectation must fail")
	}
	if matches(observe.NewNumber(1), Expectation{Type: "unknown"}, Tolerance{}) {
		t.Fatalf("unknown expectation type must fail")
	}
}
