/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package observe

import "testing"

// chaosSrc is a local deterministic integer source used only by chaos tests.
type chaosSrc struct{ state uint64 }

// next advances the generator with integer arithmetic.
func (g *chaosSrc) next() uint64 {
	g.state = g.state*6364136223846793005 + 1442695040888963407
	return g.state
}

// pick returns a deterministic value in [0, n).
func (g *chaosSrc) pick(n uint64) uint64 {
	return g.next() % n
}

// index returns a deterministic slice index in [0, n).
func (g *chaosSrc) index(n int) int {
	return int(g.next() % uint64(n)) //nolint:gosec // the modulo result is below n and fits int.
}

// chaosText draws a deterministic string from a fixed alphabet.
func chaosText(g *chaosSrc, n int, alphabet []rune) string {
	out := make([]rune, 0, n)
	for i := 0; i < n; i++ {
		out = append(out, alphabet[g.index(len(alphabet))])
	}
	return string(out)
}

// TestChaosValueRoundTrip fuzzes all four value kinds.
func TestChaosValueRoundTrip(t *testing.T) {
	t.Parallel()
	g := chaosSrc{state: 0x9E3779B97F4A7C15}
	alphabet := []rune("中文特殊<>\"'\\\n\t0123456789abcdefXYZ ")
	numbers := []float64{0, 1, -1, 3.14, 1e9, 1e-9, 123456.789}
	for i := 0; i < 2000; i++ {
		var v Value
		switch g.pick(4) {
		case 0:
			v = NewNumber(numbers[g.index(len(numbers))])
		case 1:
			v = NewBoolean(g.pick(2) == 1)
		case 2:
			v = NewText(chaosText(&g, 1+g.index(32), alphabet))
		default:
			buckets := make([]uint64, 0, 8)
			for j := 0; j < 8; j++ {
				buckets = append(buckets, g.pick(1000))
			}
			v = NewHistogram(buckets)
		}
		data, err := v.Encode()
		if err != nil {
			t.Fatalf("iteration %d Encode err = %v", i, err)
		}
		back, err := DecodeValue(data)
		if err != nil || !v.Equal(back) {
			t.Fatalf("iteration %d round trip failed", i)
		}
	}
}

// TestChaosCollectOrderIsStable shuffles registration order 100 times.
func TestChaosCollectOrderIsStable(t *testing.T) {
	t.Parallel()
	g := chaosSrc{state: 0xD1B54A32D192ED03}
	ids := []string{"p-a", "p-b", "p-c", "p-d", "p-e"}
	var baseline [32]byte
	for run := 0; run < 100; run++ {
		shuffled := append([]string(nil), ids...)
		for j := len(shuffled) - 1; j > 0; j-- {
			k := g.index(j + 1)
			shuffled[j], shuffled[k] = shuffled[k], shuffled[j]
		}
		p := NewPipeline()
		for _, id := range shuffled {
			if err := p.Register(&fakeProbe{
				id: id, kind: ProbeMetric, value: NewText(id), status: StatusOK,
			}); err != nil {
				t.Fatalf("Register err = %v", err)
			}
		}
		obs := p.Collect(100, 1)
		if len(obs) != len(ids) {
			t.Fatalf("run %d len = %d, want %d", run, len(obs), len(ids))
		}
		for i, want := range ids {
			if obs[i].ProbeID != want {
				t.Fatalf("run %d obs[%d] = %s, want %s", run, i, obs[i].ProbeID, want)
			}
		}
		digest := StreamDigest(obs)
		if run == 0 {
			baseline = digest
		} else if digest != baseline {
			t.Fatalf("digest diverged at run %d", run)
		}
	}
}
