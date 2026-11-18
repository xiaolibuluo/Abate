/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package csfront

import (
	"fmt"
	"testing"

	"abate/lang/air"
)

// csBenchSource builds a deterministic C# source of the requested number of
// classes by integer repetition only — no clock, no random seed — so the
// measured byte and node counts are reproducible.
func csBenchSource(classes int) string {
	var out []byte
	for i := 0; i < classes; i++ {
		out = fmt.Appendf(out,
			"public class C%d { public int Add(int a, int b) { var s = a + b; return s; } }\n", i)
	}
	return string(out)
}

// BenchmarkCsFrontParse records real syntax-to-AIR throughput for the cs
// frontend, closing the P8 "frontend parse throughput" budget for LNG-cs.
func BenchmarkCsFrontParse(b *testing.B) {
	fe, err := NewCsFrontend()
	if err != nil {
		b.Fatal(err)
	}
	src := csBenchSource(24)
	unit := air.SourceUnit{
		Language: "cs", Dialect: "cs12",
		Files: []air.SourceFile{air.NewSourceFile("bench.cs", []byte(src))},
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
