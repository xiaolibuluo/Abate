/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package jsfront

import (
	"fmt"
	"testing"

	"abate/lang/air"
)

// jsBenchSource builds a deterministic JavaScript/TSX source of the requested
// number of function declarations by integer repetition only — no clock, no
// random seed — so the measured byte count and node count are reproducible.
func jsBenchSource(functions int) string {
	var out []byte
	for i := 0; i < functions; i++ {
		out = fmt.Appendf(out,
			"function f%d(a, b) { const s = a + b * 2; return s > 0 ? a : b; }\n", i)
	}
	return string(out)
}

// BenchmarkJsFrontParse records real syntax-to-AIR throughput for the js
// frontend, closing the P8 "frontend parse throughput" budget for LNG-js.
func BenchmarkJsFrontParse(b *testing.B) {
	fe, err := NewJsFrontend()
	if err != nil {
		b.Fatal(err)
	}
	src := jsBenchSource(24)
	unit := air.SourceUnit{
		Language: "js", Dialect: "tsx/es2022",
		Files: []air.SourceFile{air.NewSourceFile("bench.ts", []byte(src))},
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
