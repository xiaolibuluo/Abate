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
	"testing"
)

// TestMulDiv 验证乘法先于除法、向零取整（CES-4.1.3）。
func TestMulDiv(t *testing.T) {
	cases := []struct {
		a, b, d, want BP
	}{
		{10000, 10000, 10000, 10000},
		{9999, 9999, 10000, 9998}, // 99980001/10000 截断
		{5000, 1, 3, 1666},        // 5000/3 向零取整
		{8000, 2000, 10000, 1600},
		{1234, 0, 10000, 0},
		{1234, 10000, 0, 0}, // 除零按缺测语义返回 0
	}
	for _, c := range cases {
		if got := c.a.MulDiv(c.b, c.d); got != c.want {
			t.Errorf("%d.MulDiv(%d,%d)=%d, want %d", c.a, c.b, c.d, got, c.want)
		}
	}
}

// TestClampMinMax 验证区间收敛与取极值。
func TestClampMinMax(t *testing.T) {
	if got := BP(12000).Clamp(0, BPFull); got != BPFull {
		t.Errorf("Clamp high = %d", got)
	}
	if got := BP(-5).Clamp(0, BPFull); got != 0 {
		t.Errorf("Clamp low = %d", got)
	}
	if BP(3).Min(4) != 3 || BP(3).Max(4) != 4 {
		t.Error("Min/Max wrong")
	}
	if BP(5).Min(4) != 4 || BP(5).Max(4) != 5 {
		t.Error("Min/Max reverse-order branches wrong")
	}
}

// TestRatio 验证万分位比率与除零。
func TestRatio(t *testing.T) {
	if got := Ratio(1, 2); got != 5000 {
		t.Errorf("Ratio(1,2)=%d, want 5000", got)
	}
	if got := Ratio(99, 100); got != 9900 {
		t.Errorf("Ratio(99,100)=%d, want 9900", got)
	}
	if got := Ratio(1, 0); got != 0 {
		t.Errorf("Ratio(1,0)=%d, want 0", got)
	}
}

// TestWeighted 验证加权聚合与缺测信号。
func TestWeighted(t *testing.T) {
	v, ok := Weighted([]BP{10000, 5000}, []BP{8000, 4000})
	if !ok || v != 6666 {
		t.Errorf("Weighted=(%d,%v), want (6666,true)", v, ok)
	}
	if _, ok := Weighted(nil, nil); ok {
		t.Error("empty input must be unmeasured")
	}
	if _, ok := Weighted([]BP{1}, []BP{1, 2}); ok {
		t.Error("mismatched lengths must be unmeasured")
	}
}

// TestSortedIDs 验证字典序升序且不修改入参（CES-4.1.3）。
func TestSortedIDs(t *testing.T) {
	in := []string{"b", "a10", "a2", "a"}
	got := SortedIDs(in)
	want := []string{"a", "a10", "a2", "b"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("SortedIDs=%v, want %v", got, want)
	}
	if in[0] != "b" {
		t.Error("input slice mutated")
	}
}
