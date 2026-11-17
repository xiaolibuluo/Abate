/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package determinism

import "testing"

// TestFuelMeterInitialState verifies a fresh meter reports the full budget.
func TestFuelMeterInitialState(t *testing.T) {
	m := NewFuelMeter(100)
	if m.Initial() != 100 {
		t.Fatalf("initial = %d, want 100", m.Initial())
	}
	if m.Remaining() != 100 || m.Used() != 0 || m.Halted() {
		t.Fatalf("remaining=%d used=%d halted=%v", m.Remaining(), m.Used(), m.Halted())
	}
}

// TestFuelMeterConsumeDecrements verifies normal consumption preserves invariants.
func TestFuelMeterConsumeDecrements(t *testing.T) {
	m := NewFuelMeter(10)
	rem, halted := m.Consume(3)
	if rem != 7 || halted {
		t.Fatalf("rem=%d halted=%v", rem, halted)
	}
	if m.Used() != 3 || m.Remaining() != 7 {
		t.Fatalf("used=%d remaining=%d", m.Used(), m.Remaining())
	}
	if m.Used() > m.Initial() || m.Remaining() > m.Initial() {
		t.Fatalf("invariant violated: used=%d remaining=%d initial=%d", m.Used(), m.Remaining(), m.Initial())
	}
}

// TestFuelMeterClampsOnExhaustion verifies the SandboxFuel invariants at the edge.
func TestFuelMeterClampsOnExhaustion(t *testing.T) {
	m := NewFuelMeter(5)
	rem, halted := m.Consume(3)
	if rem != 2 || halted {
		t.Fatalf("rem=%d halted=%v", rem, halted)
	}
	rem, halted = m.Consume(10)
	if rem != 0 || !halted {
		t.Fatalf("rem=%d halted=%v, want 0 true", rem, halted)
	}
	if m.Used() != 5 || m.Remaining() != 0 {
		t.Fatalf("used=%d remaining=%d, want 5 0", m.Used(), m.Remaining())
	}
	rem, halted = m.Consume(1)
	if rem != 0 || !halted {
		t.Fatalf("post-halt rem=%d halted=%v", rem, halted)
	}
}

// TestFuelMeterDisabledNeverHalts verifies a zero budget disables metering.
func TestFuelMeterDisabledNeverHalts(t *testing.T) {
	m := NewFuelMeter(0)
	if m.Halted() {
		t.Fatalf("disabled meter must not start halted")
	}
	for i := 0; i < 3; i++ {
		rem, halted := m.Consume(1)
		if halted || rem == 0 {
			t.Fatalf("disabled meter halted unexpectedly: rem=%d", rem)
		}
	}
}

// TestFuelMeterZeroConsumeIsNoOp verifies zero consumption leaves the budget untouched.
func TestFuelMeterZeroConsumeIsNoOp(t *testing.T) {
	m := NewFuelMeter(10)
	rem, halted := m.Consume(0)
	if rem != 10 || halted {
		t.Fatalf("Consume(0) = %d/%v, want 10/false", rem, halted)
	}
	if m.Used() != 0 || m.Remaining() != 10 {
		t.Fatalf("used = %d remaining = %d, want 0/10", m.Used(), m.Remaining())
	}
}
