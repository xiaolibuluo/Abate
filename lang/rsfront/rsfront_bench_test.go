/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package rsfront

import (
	"fmt"
	"testing"

	"abate/lang/air"
)

// rsBenchSource builds a deterministic Rust source of the requested number of
// structs plus functions by integer repetition only — no clock, no random
// seed — so the measured byte and node counts are reproducible.
func rsBenchSource(units int) string {
	var out []byte
	for i := 0; i < units; i++ {
		out = fmt.Appendf(out,
			"struct S%d { x: i32, y: i32 }\n"+
				"fn add%d(a: i32, b: i32) -> i32 { let s = a + b; s }\n", i, i)
	}
	return string(out)
}

// BenchmarkRsFrontParse records real syntax-to-AIR throughput for the rs
// frontend, closing the P8 "frontend parse throughput" budget for LNG-rs.
func BenchmarkRsFrontParse(b *testing.B) {
	fe, err := NewRsFrontend()
	if err != nil {
		b.Fatal(err)
	}
	src := rsBenchSource(24)
	unit := air.SourceUnit{
		Language: "rs", Dialect: "rust2021",
		Files: []air.SourceFile{air.NewSourceFile("bench.rs", []byte(src))},
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
