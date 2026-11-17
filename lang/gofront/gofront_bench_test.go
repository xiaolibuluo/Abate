/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package gofront

import (
	"testing"

	"abate/lang/air"
)

// BenchmarkParse records real syntax-to-AIR throughput.
func BenchmarkParse(b *testing.B) {
	frontend := NewGoFrontend().WithTypeCheck(false)
	unit := air.SourceUnit{
		Language: "go",
		Files:    []air.SourceFile{air.NewSourceFile("bench.go", []byte(sampleGo))},
	}
	b.SetBytes(int64(len(sampleGo)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, _, err := frontend.Parse(unit); err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkParseWithTypes records the cost including the in-memory checker.
func BenchmarkParseWithTypes(b *testing.B) {
	frontend := NewGoFrontend()
	unit := air.SourceUnit{
		Language: "go",
		Files:    []air.SourceFile{air.NewSourceFile("bench.go", []byte(sampleGo))},
	}
	b.SetBytes(int64(len(sampleGo)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, _, err := frontend.Parse(unit); err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkVerifyTokenFidelity records the real token equivalence check cost.
func BenchmarkVerifyTokenFidelity(b *testing.B) {
	frontend := NewGoFrontend().WithTypeCheck(false)
	parsed, _, err := frontend.Parse(air.SourceUnit{
		Language: "go",
		Files:    []air.SourceFile{air.NewSourceFile("bench.go", []byte(sampleGo))},
	})
	if err != nil {
		b.Fatal(err)
	}
	src := []byte(sampleGo)
	b.SetBytes(int64(len(src)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := VerifyTokenFidelity(src, parsed); err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkEncodeFrontendAIR records the cost of archiving frontend output.
func BenchmarkEncodeFrontendAIR(b *testing.B) {
	frontend := NewGoFrontend().WithTypeCheck(false)
	parsed, _, err := frontend.Parse(air.SourceUnit{
		Language: "go",
		Files:    []air.SourceFile{air.NewSourceFile("bench.go", []byte(sampleGo))},
	})
	if err != nil {
		b.Fatal(err)
	}
	encoder := air.NewEncoder()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := encoder.Encode(parsed); err != nil {
			b.Fatal(err)
		}
	}
}
