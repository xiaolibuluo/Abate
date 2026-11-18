/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package ktfront

import (
	"fmt"
	"testing"

	"abate/lang/air"
)

// ktBenchSource builds a deterministic Kotlin source of the requested number
// of data classes plus functions by integer repetition only — no clock, no
// random seed — so the measured byte and node counts are reproducible.
func ktBenchSource(units int) string {
	var out []byte
	for i := 0; i < units; i++ {
		out = fmt.Appendf(out,
			"data class C%d(val x: Int, val y: Int)\n"+
				"fun add%d(a: Int, b: Int): Int { val s = a + b; return s }\n", i, i)
	}
	return string(out)
}

// BenchmarkKtFrontParse records real syntax-to-AIR throughput for the kt
// frontend, closing the P8 "frontend parse throughput" budget for LNG-kt.
func BenchmarkKtFrontParse(b *testing.B) {
	fe, err := NewKtFrontend()
	if err != nil {
		b.Fatal(err)
	}
	src := ktBenchSource(24)
	unit := air.SourceUnit{
		Language: "kt", Dialect: "kt2.1",
		Files: []air.SourceFile{air.NewSourceFile("bench.kt", []byte(src))},
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
