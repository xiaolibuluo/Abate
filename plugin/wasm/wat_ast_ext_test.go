/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package wasm

import (
	"math"
	"testing"
)

// TestToU32AcceptsInRangeValues verifies the non-negative uint32 window.
func TestToU32AcceptsInRangeValues(t *testing.T) {
	cases := []struct {
		name string
		in   int64
		want uint32
	}{
		{name: "zero", in: 0, want: 0},
		{name: "one", in: 1, want: 1},
		{name: "page boundary", in: 65_536, want: 65_536},
		{name: "max uint32", in: math.MaxUint32, want: math.MaxUint32},
	}
	for _, tc := range cases {
		got, err := toU32(tc.in)
		if err != nil {
			t.Errorf("%s: toU32(%d) error = %v, want nil", tc.name, tc.in, err)
			continue
		}
		if got != tc.want {
			t.Errorf("%s: toU32(%d) = %d, want %d", tc.name, tc.in, got, tc.want)
		}
	}
}

// TestToU32RejectsOutOfRangeValues verifies the negative and overflow branches.
func TestToU32RejectsOutOfRangeValues(t *testing.T) {
	cases := []struct {
		name string
		in   int64
	}{
		{name: "negative one", in: -1},
		{name: "large negative", in: math.MinInt64},
		{name: "one above max", in: math.MaxUint32 + 1},
		{name: "max int64", in: math.MaxInt64},
	}
	for _, tc := range cases {
		got, err := toU32(tc.in)
		if err == nil {
			t.Errorf("%s: toU32(%d) = %d, want an error", tc.name, tc.in, got)
			continue
		}
		if got != 0 {
			t.Errorf("%s: toU32(%d) = %d on error, want 0", tc.name, tc.in, got)
		}
	}
}

// TestToU64AcceptsNonNegativeValues verifies widening of valid indices.
func TestToU64AcceptsNonNegativeValues(t *testing.T) {
	cases := []struct {
		name string
		in   int64
		want uint64
	}{
		{name: "zero", in: 0, want: 0},
		{name: "one", in: 1, want: 1},
		{name: "large index", in: 1 << 40, want: 1 << 40},
		{name: "max int64", in: math.MaxInt64, want: math.MaxInt64},
	}
	for _, tc := range cases {
		got, err := toU64(tc.in)
		if err != nil {
			t.Errorf("%s: toU64(%d) error = %v, want nil", tc.name, tc.in, err)
			continue
		}
		if got != tc.want {
			t.Errorf("%s: toU64(%d) = %d, want %d", tc.name, tc.in, got, tc.want)
		}
	}
}

// TestToU64RejectsNegativeValues verifies the negative-index rejection branch.
func TestToU64RejectsNegativeValues(t *testing.T) {
	cases := []struct {
		name string
		in   int64
	}{
		{name: "negative one", in: -1},
		{name: "large negative", in: -1 << 40},
		{name: "min int64", in: math.MinInt64},
	}
	for _, tc := range cases {
		got, err := toU64(tc.in)
		if err == nil {
			t.Errorf("%s: toU64(%d) = %d, want an error", tc.name, tc.in, got)
			continue
		}
		if got != 0 {
			t.Errorf("%s: toU64(%d) = %d on error, want 0", tc.name, tc.in, got)
		}
	}
}
