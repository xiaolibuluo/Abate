/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package determinism

import (
	"sort"
	"testing"

	"abate/sim/rand"
)

// TestHashStateMapOrderIndependence asserts that maps hash identically regardless
// of the order in which their keys were inserted.
func TestHashStateMapOrderIndependence(t *testing.T) {
	t.Parallel()
	a := map[string]int{"x": 1, "y": 2, "z": 3}
	b := map[string]int{"z": 3, "y": 2, "x": 1}
	if HashState(a) != HashState(b) {
		t.Fatalf("map hash depends on insertion order")
	}
}

// TestHashStateSliceOrderSensitivity confirms order is significant for slices.
func TestHashStateSliceOrderSensitivity(t *testing.T) {
	t.Parallel()
	a := []int{1, 2, 3}
	b := []int{3, 2, 1}
	if HashState(a) == HashState(b) {
		t.Fatalf("slice hash ignored element order")
	}
}

// TestMustMatch100Deterministic runs a seeded producer one hundred times and
// asserts all outputs are byte-identical; then repeats the whole assertion to
// ensure the double-run is itself stable.
func TestMustMatch100Deterministic(t *testing.T) {
	t.Parallel()
	build := func() []byte {
		s := rand.New(4242)
		out := make([]byte, 64)
		for i := range out {
			out[i] = byte(s.Uint64() >> 56)
		}
		return out
	}
	MustMatch100(t, 4242, build)
	MustMatch100(t, 4242, build)
}

// TestMustMatch100SortedKeys exercises HashState on a value whose map ordering
// is shuffled between runs to reinforce order-independence.
func TestMustMatch100SortedKeys(t *testing.T) {
	t.Parallel()
	build := func() []byte {
		keys := []string{"c", "a", "b"}
		m := map[string]int{}
		for _, k := range keys {
			m[k] = int(k[0])
		}
		sorted := make([]string, 0, len(m))
		for k := range m {
			sorted = append(sorted, k)
		}
		sort.Strings(sorted)
		h := HashState(map[string]any{"k": sorted, "m": m})
		return h[:]
	}
	MustMatch100(t, 7, build)
}

// TestHashStateUnmarshalableFallsBack confirms HashState produces a stable
// fingerprint for values that json.Marshal cannot encode, instead of panicking.
func TestHashStateUnmarshalableFallsBack(t *testing.T) {
	t.Parallel()
	h := HashState(make(chan int))
	if h == [32]byte{} {
		t.Fatalf("HashState returned a zero hash for an unmarshalable value")
	}
	// The fallback must still be deterministic for the same input shape.
	if HashState(make(chan int)) != h {
		t.Fatalf("HashState fallback is not deterministic")
	}
}
