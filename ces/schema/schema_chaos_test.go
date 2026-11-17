/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package schema

import (
	"bytes"
	"encoding/json"
	"strconv"
	"testing"
)

// chaosRNG is a local deterministic LCG used only by chaos tests.
type chaosRNG struct{ state uint64 }

// next advances the generator with SplitMix-style integer arithmetic.
func (g *chaosRNG) next() uint64 {
	g.state = g.state*6364136223846793005 + 1442695040888963407
	return g.state
}

// chaosString draws a deterministic string from a fixed alphabet.
func chaosString(g *chaosRNG, n int, alphabet []rune) string {
	out := make([]rune, 0, n)
	for i := 0; i < n; i++ {
		out = append(out, alphabet[g.next()%uint64(len(alphabet))])
	}
	return string(out)
}

// TestChaosDurationRoundTrip hammers ParseDuration with seeded literals.
func TestChaosDurationRoundTrip(t *testing.T) {
	t.Parallel()
	g := chaosRNG{state: 0x9E3779B97F4A7C15}
	units := []string{"ns", "us", "µs", "ms", "s", "m", "h"}
	for i := 0; i < 5000; i++ {
		n := g.next() % 1000000
		u := units[g.next()%uint64(len(units))]
		lit := strconv.FormatUint(n, 10) + u
		d, err := ParseDuration(lit)
		if err != nil {
			t.Fatalf("ParseDuration(%q) err = %v", lit, err)
		}
		back, err := ParseDuration(d.String())
		if err != nil || back != d {
			t.Fatalf("round trip %q -> %d -> %q -> %d (err %v)",
				lit, d, d.String(), back, err)
		}
	}
}

// TestChaosParseNeverPanics feeds random byte strings to the duration parser.
func TestChaosParseNeverPanics(t *testing.T) {
	t.Parallel()
	g := chaosRNG{state: 0xD1B54A32D192ED03}
	alphabet := []rune("0123456789nsushm.-+é中")
	for i := 0; i < 5000; i++ {
		s := chaosString(&g, 1+int(g.next()%16), alphabet) //nolint:gosec // value is bounded by a modulo or bitmask; conversion is safe.
		_, _ = ParseDuration(s)
	}
}

// TestChaosDocumentRoundTrip fuzzes text fields with boundary characters.
func TestChaosDocumentRoundTrip(t *testing.T) {
	t.Parallel()
	g := chaosRNG{state: 0x2545F4914F6CDD1D}
	alphabet := []rune("中文特殊<>\"'\\\n\t\r`{}[]0123456789abcdefXYZ ")
	for i := 0; i < 200; i++ {
		exp := sampleExperiment()
		exp.Metadata.Name = chaosString(&g, 1+int(g.next()%24), alphabet)                //nolint:gosec // value is bounded by a modulo or bitmask; conversion is safe.
		exp.Metadata.Annotations["note"] = chaosString(&g, 1+int(g.next()%40), alphabet) //nolint:gosec // value is bounded by a modulo or bitmask; conversion is safe.
		first, err := json.Marshal(exp)
		if err != nil {
			t.Fatalf("json.Marshal err = %v", err)
		}
		var back Experiment
		// B3 (ADR-YAML-001): JSON interop uses encoding/json.
		if err := json.Unmarshal(first, &back); err != nil {
			t.Fatalf("json.Unmarshal err = %v on %s", err, first)
		}
		second, err := json.Marshal(back)
		if err != nil {
			t.Fatalf("json.Marshal err = %v", err)
		}
		if !bytes.Equal(first, second) {
			t.Fatalf("round-trip drift on iteration %d: %s", i, first)
		}
	}
}
