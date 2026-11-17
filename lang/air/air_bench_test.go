/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package air

import "testing"

// benchmarkUnit builds a moderately sized unit for codec benchmarks.
func benchmarkUnit() *Unit {
	u := NewUnit(NewPackageRef("bench", []string{"fmt", "os"}))
	children := make([]Node, 0, 256)
	for i := 0; i < 256; i++ {
		children = append(children,
			NewNode(KindIdent, NewPosition(1, i+1, i, i+1), "ident"))
	}
	root := NewNode(KindBlock, NewPosition(1, 1, 0, 256), "", children...)
	if err := u.AddFile(NewFileRef("bench.go", make([]byte, 256)), root); err != nil {
		panic(err)
	}
	return u
}

// BenchmarkEncode records real encode throughput and allocations.
func BenchmarkEncode(b *testing.B) {
	u := benchmarkUnit()
	encoder := NewEncoder()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := encoder.Encode(u); err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkDecode records real decode throughput and allocations.
func BenchmarkDecode(b *testing.B) {
	archive, err := NewEncoder().Encode(benchmarkUnit())
	if err != nil {
		b.Fatal(err)
	}
	decoder := NewDecoder()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := decoder.Decode(archive); err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkFingerprint records the cost of canonical unit hashing.
func BenchmarkFingerprint(b *testing.B) {
	u := benchmarkUnit()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = u.Fingerprint()
	}
}

// BenchmarkNodeID records the cost of deterministic identifier derivation.
func BenchmarkNodeID(b *testing.B) {
	ids := []NodeID{1, 2, 3, 4}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ComputeNodeID(KindCall, "operator", ids)
	}
}
