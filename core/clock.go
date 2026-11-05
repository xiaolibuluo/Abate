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

import "time"

// Clock 是时间抽象接口：引擎、评分、模拟器中一切时间读取必须经由此接口，
// 禁止直接调用 time.Now()。生产注入 RealClock，测试与确定性模拟注入
// VirtualClock——时间抖动在结构上被物理消除（零抖动底座三件套之一）。
//
// Clock is the time abstraction: every time read in the engine, scoring, and
// simulator must go through this interface; calling time.Now() directly is
// forbidden. Production injects RealClock; tests and the deterministic
// simulator inject VirtualClock — time jitter is structurally eliminated.
type Clock interface {
	// Now 返回当前时刻（Unix 毫秒）。
	Now() int64
}

// RealClock 是生产物理时钟。
// RealClock is the production physical clock.
type RealClock struct{}

// Now 返回物理当前时刻（Unix 毫秒）。
// Now returns the physical current time in Unix milliseconds.
func (RealClock) Now() int64 { return time.Now().UnixMilli() }

// VirtualClock 是确定性虚拟时钟：时间只在显式推进时流动。
// VirtualClock is a deterministic virtual clock: time flows only when
// explicitly advanced.
type VirtualClock struct {
	now int64
}

// NewVirtualClock 以初始时刻（Unix 毫秒）构造虚拟时钟。
// NewVirtualClock constructs a virtual clock from an initial instant.
func NewVirtualClock(startMs int64) *VirtualClock {
	return &VirtualClock{now: startMs}
}

// Now 实现 Clock：返回虚拟当前时刻。
// Now implements Clock: the virtual current instant.
func (c *VirtualClock) Now() int64 { return c.now }

// Advance 将虚拟时间推进 deltaMs 毫秒并返回新时刻。
// Advance moves virtual time forward by deltaMs and returns the new instant.
func (c *VirtualClock) Advance(deltaMs int64) int64 {
	c.now += deltaMs
	return c.now
}

// Set 直接设定虚拟时刻。
// Set directly sets the virtual instant.
func (c *VirtualClock) Set(ms int64) { c.now = ms }
