/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package determinism

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

const (
	frozenGoldenPath = "../../docs/refactor/golden/plugin-semantics-before-migration.txt"
	goldenBodyMark   = "## FuelMeter\n"
	largeBudget      = 1 << 20
	epochNanos       = 1000
)

// TestSemanticsMatchFrozenGolden verifies that the migrated host-layer
// primitives reproduce the frozen T01 observation byte for byte. Only the
// semantic body is compared; the header is historical metadata.
func TestSemanticsMatchFrozenGolden(t *testing.T) {
	raw, err := os.ReadFile(frozenGoldenPath)
	if err != nil {
		t.Fatalf("read frozen golden: %v", err)
	}
	idx := strings.Index(string(raw), goldenBodyMark)
	if idx < 0 {
		t.Fatalf("frozen golden missing body marker %q", goldenBodyMark)
	}
	want := string(raw)[idx:]

	var b strings.Builder
	writeFuel(&b)
	writeClock(&b)
	writeConfig(&b)
	got := b.String()

	if got != want {
		t.Fatalf("migrated semantics differ from frozen golden:\n--- got ---\n%s\n--- want ---\n%s", got, want)
	}
}

// writeFuel records FuelMeter state transitions for the scenario matrix.
func writeFuel(b *strings.Builder) {
	b.WriteString("## FuelMeter\n")
	for _, budget := range []uint64{0, 1, 10, largeBudget} {
		m := NewFuelMeter(budget)
		fmt.Fprintf(b, "fuel budget=%d initial=%d remaining=%d used=%d halted=%v\n",
			budget, m.Initial(), m.Remaining(), m.Used(), m.Halted())
		for _, n := range []uint64{0, 1, 3, budget, budget + 1} {
			remaining, halted := m.Consume(n)
			fmt.Fprintf(b, "  consume=%d -> remaining=%d halted=%v | state remaining=%d used=%d halted=%v\n",
				n, remaining, halted, m.Remaining(), m.Used(), m.Halted())
		}
	}
	b.WriteString("\n")
}

// writeClock records VirtualClock advance/reset transitions.
func writeClock(b *strings.Builder) {
	b.WriteString("## VirtualClock\n")
	for _, epoch := range []uint64{0, epochNanos} {
		c := NewVirtualClock(epoch)
		fmt.Fprintf(b, "clock epoch=%d epochGet=%d now=%d\n", epoch, c.Epoch(), c.Now())
		for _, delta := range []uint64{0, 1, 7} {
			c.Advance(delta)
			fmt.Fprintf(b, "  advance=%d -> now=%d\n", delta, c.Now())
		}
		c.Reset()
		fmt.Fprintf(b, "  reset -> now=%d\n", c.Now())
	}
	b.WriteString("\n")
}

// writeConfig records the deterministic default config and validation outcomes.
func writeConfig(b *strings.Builder) {
	b.WriteString("## Config\n")
	def := DefaultConfig()
	fmt.Fprintf(b, "default fuel=%d min=%d max=%d clock=%v epoch=%d validate=%v\n",
		def.FuelLimit, def.MemoryMinPages, def.MemoryMaxPages,
		def.ClockEnabled, def.ClockEpoch, def.Validate())
	zeroMax := Config{FuelLimit: 10, MemoryMinPages: 1, MemoryMaxPages: 0}
	minAbove := Config{FuelLimit: 10, MemoryMinPages: 2, MemoryMaxPages: 1}
	fmt.Fprintf(b, "zero-max validate=%v\n", zeroMax.Validate())
	fmt.Fprintf(b, "min-above-max validate=%v\n", minAbove.Validate())
}
