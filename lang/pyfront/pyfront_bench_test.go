/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package pyfront

import (
	"strings"
	"testing"

	"abate/lang/air"
)

// pyBenchSource builds a deterministic Python source of the requested number
// of statements. It uses integer repetition only, never a clock or random seed.
func pyBenchSource(statements int) string {
	const letters = "abcdefghijklmnopqrstuvwxyz"
	var out strings.Builder
	for i := 0; i < statements; i++ {
		out.WriteString("def f")
		out.WriteByte(letters[i%len(letters)])
		out.WriteString("(x: int, y=0) -> int:\n")
		out.WriteString("    xs = [v * 2 for v in range(8) if v > 1]\n")
		out.WriteString("    return sum(xs) + x + y\n\n")
	}
	return out.String()
}

// BenchmarkPyFrontendBytes measures throughput in source bytes per operation
// and reports the integer-derived node count; no wall-clock float is asserted.
func BenchmarkPyFrontendBytes(b *testing.B) {
	fe, err := NewPyFrontend()
	if err != nil {
		b.Fatal(err)
	}
	src := pyBenchSource(24)
	unit := air.SourceUnit{
		Language: "py", Dialect: "3.12",
		Files: []air.SourceFile{air.NewSourceFile("bench.py", []byte(src))},
	}
	b.SetBytes(int64(len(src)))
	b.ReportAllocs()
	nodes := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		out, _, parseErr := fe.Parse(unit)
		if parseErr != nil {
			b.Fatal(parseErr)
		}
		nodes += out.NodeCount()
	}
	b.StopTimer()
	if nodes == 0 {
		b.Fatal("benchmark produced no AIR nodes")
	}
	b.ReportMetric(float64(nodes)/float64(b.N), "nodes/op")
}

// BenchmarkPyFrontendLargeBuffer repeats the byte/node accounting on a larger
// buffer to exercise the deterministic token and node ceilings.
func BenchmarkPyFrontendLargeBuffer(b *testing.B) {
	fe, err := NewPyFrontend()
	if err != nil {
		b.Fatal(err)
	}
	src := pyBenchSource(200)
	unit := air.SourceUnit{
		Language: "py", Dialect: "3.12",
		Files: []air.SourceFile{air.NewSourceFile("large.py", []byte(src))},
	}
	b.SetBytes(int64(len(src)))
	b.ReportAllocs()
	total := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		out, _, parseErr := fe.Parse(unit)
		if parseErr != nil {
			b.Fatal(parseErr)
		}
		total += out.NodeCount()
	}
	b.StopTimer()
	if total < b.N {
		b.Fatal("large benchmark produced fewer than one node per run")
	}
	b.ReportMetric(float64(total)/float64(b.N), "nodes/op")
}
