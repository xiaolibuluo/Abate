/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package gofront

import (
	"testing"

	"abate/lang/air"
)

// TestGoOperatorFidelity asserts AIR-ID changes for cmp, op and digit mutations, with identical trailing comments on both variants.
func TestGoOperatorFidelity(t *testing.T) {
	cases := []struct{ name, base, mutated string }{
		{"cmp", "package f\n\nfunc eq(a int, b int) bool { return a == b }\n", "package f\n\nfunc eq(a int, b int) bool { return a != b }\n"},
		{"op", "package f\n\nfunc add(a int, b int) int { return a + b }\n", "package f\n\nfunc add(a int, b int) int { return a - b }\n"},
		{"digit", "package f\n\nfunc value() int { return 41 }\n", "package f\n\nfunc value() int { return 42 }\n"},
	}
	rootID := func(t *testing.T, src string) air.NodeID {
		t.Helper()
		unit, _ := parseSource(t, src)
		if len(unit.Roots) != 1 {
			t.Fatalf("got %d roots, want 1", len(unit.Roots))
		}
		return air.IDOf(unit.Roots[0])
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if rootID(t, tc.base) == rootID(t, tc.mutated) {
				t.Fatal("operator or digit mutation did not change AIR-ID")
			}
			baseComment, mutatedComment := tc.base+"// tail\n", tc.mutated+"// tail\n"
			if rootID(t, baseComment) == rootID(t, mutatedComment) {
				t.Fatal("identical comments changed the capture conclusion")
			}
		})
	}
}
