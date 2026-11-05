// Abate —— 止息混沌 · 混沌工程开放规范（CES）与参考实现
// Abate — where chaos abates. Open Chaos Experiment Specification (CES) & reference implementation.
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
//
// 本文件以 GNU Affero General Public License v3（AGPLv3）发布；商业用途须事先获得
// 小黎的部落书面授权，详见 COMMERCIAL.md。
// This file is published under the GNU Affero General Public License v3 (AGPLv3).
// Commercial use requires prior written authorization from 小黎的部落.
// See COMMERCIAL.md for details.

package core

import (
	"reflect"
	"sort"
	"testing"
)

// TestRNGDeterministic 验证同一 seed 产生完全相同序列（零抖动底座，CES-3.1.3）。
func TestRNGDeterministic(t *testing.T) {
	a, b := NewRNG(42), NewRNG(42)
	for i := 0; i < 1000; i++ {
		if a.Uint64() != b.Uint64() {
			t.Fatalf("sequence diverged at step %d", i)
		}
	}
}

// TestRNGSeedSensitivity 验证不同 seed 产生不同序列。
func TestRNGSeedSensitivity(t *testing.T) {
	a, b := NewRNG(1), NewRNG(2)
	same := 0
	for i := 0; i < 100; i++ {
		if a.Uint64() == b.Uint64() {
			same++
		}
	}
	if same > 0 {
		t.Errorf("different seeds produced %d identical values", same)
	}
}

// TestRNGZeroSeed 验证零种子归一为非零状态，不陷入退化序列。
func TestRNGZeroSeed(t *testing.T) {
	r := NewRNG(0)
	if r.Uint64() == 0 {
		t.Error("zero seed must not yield zero first output")
	}
	r2 := NewRNG(0)
	r3 := NewRNG(0)
	if r3.Uint64() != r2.Uint64() {
		t.Error("zero seed normalization must be deterministic")
	}
}

// TestRNGIntn 验证区间约束与非法入参。
func TestRNGIntn(t *testing.T) {
	r := NewRNG(7)
	for i := 0; i < 1000; i++ {
		v := r.Intn(10)
		if v < 0 || v >= 10 {
			t.Fatalf("Intn(10)=%d out of range", v)
		}
	}
	if r.Intn(0) != 0 || r.Intn(-5) != 0 {
		t.Error("non-positive n must yield 0")
	}
}

// TestRNGPermDeterministic 验证确定性排列：同 seed 同排列，且为合法置换。
func TestRNGPermDeterministic(t *testing.T) {
	r99a, r99b := NewRNG(99), NewRNG(99)
	p1 := r99a.Perm(50)
	p2 := r99b.Perm(50)
	if !reflect.DeepEqual(p1, p2) {
		t.Error("same seed must yield identical permutation")
	}
	sorted := make([]int, 50)
	copy(sorted, p1)
	sort.Ints(sorted)
	for i, v := range sorted {
		if v != i {
			t.Fatalf("permutation is not a valid 0..49 permutation: %v", sorted)
		}
	}
	r100 := NewRNG(100)
	p3 := r100.Perm(50)
	if reflect.DeepEqual(p1, p3) {
		t.Error("different seeds unexpectedly yielded identical permutations")
	}
}

// TestRNGShuffleEmpty 验证空与单元素输入不崩溃。
func TestRNGShuffleEmpty(t *testing.T) {
	r := NewRNG(1)
	r.Shuffle(0, func(i, j int) { t.Error("swap called on empty") })
	r.Shuffle(1, func(i, j int) { t.Error("swap called on singleton") })
	if got := r.Perm(0); len(got) != 0 {
		t.Errorf("Perm(0)=%v, want empty", got)
	}
}
