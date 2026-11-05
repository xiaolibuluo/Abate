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

import "testing"

// TestVirtualClockDeterministic 验证虚拟时钟只在显式推进时流动（零抖动底座）。
func TestVirtualClockDeterministic(t *testing.T) {
	c := NewVirtualClock(1000)
	if c.Now() != 1000 {
		t.Errorf("initial Now=%d, want 1000", c.Now())
	}
	if c.Now() != 1000 {
		t.Error("virtual time must not flow without Advance")
	}
	if got := c.Advance(500); got != 1500 {
		t.Errorf("Advance(500)=%d, want 1500", got)
	}
	if c.Now() != 1500 {
		t.Errorf("Now=%d, want 1500", c.Now())
	}
	if got := c.Advance(0); got != 1500 {
		t.Errorf("Advance(0)=%d, want 1500", got)
	}
}

// TestVirtualClockSet 验证直接设定虚拟时刻。
func TestVirtualClockSet(t *testing.T) {
	c := NewVirtualClock(0)
	c.Set(42)
	if c.Now() != 42 {
		t.Errorf("Now=%d after Set(42)", c.Now())
	}
	c.Set(1)
	if c.Now() != 1 {
		t.Errorf("Now=%d after Set(1)", c.Now())
	}
}

// TestRealClock 验证物理时钟满足 Clock 接口且返回合理毫秒值。
func TestRealClock(t *testing.T) {
	var clk Clock = RealClock{}
	now := clk.Now()
	if now <= 0 {
		t.Errorf("RealClock.Now=%d, want positive Unix ms", now)
	}
	again := clk.Now()
	if again < now {
		t.Error("physical time must be monotonic non-decreasing")
	}
}

// TestClockInterfaceCompliance 编译期与运行期双重确认两种时钟都实现 Clock 接口。
func TestClockInterfaceCompliance(t *testing.T) {
	var clocks = []Clock{RealClock{}, NewVirtualClock(0)}
	for _, c := range clocks {
		_ = c.Now()
	}
}
