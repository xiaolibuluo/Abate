/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package cfront

import (
	"fmt"
	"testing"

	"abate/lang/air"
)

// cBenchSource builds a deterministic C/C++ source of the requested number of
// struct declarations plus functions by integer repetition only — no clock,
// no random seed — so the measured byte and node counts are reproducible.
func cBenchSource(units int) string {
	var out []byte
	for i := 0; i < units; i++ {
		out = fmt.Appendf(out,
			"struct S%d { int x; char y; };\n"+
				"int add%d(int a, int b) { int s = a + b; return s; }\n", i, i)
	}
	return string(out)
}

// BenchmarkCFrontParse records real syntax-to-AIR throughput for the c
// frontend, closing the P8 "frontend parse throughput" budget for LNG-c.
func BenchmarkCFrontParse(b *testing.B) {
	fe, err := NewCFrontend()
	if err != nil {
		b.Fatal(err)
	}
	src := cBenchSource(24)
	unit := air.SourceUnit{
		Language: "c", Dialect: "c++17",
		Files: []air.SourceFile{air.NewSourceFile("bench.cpp", []byte(src))},
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
