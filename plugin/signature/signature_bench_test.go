/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package signature

import "testing"

// BenchmarkVerify measures the cost of a full verification (anchor resolve +
// Ed25519 verify + hash cross-check). The reported ns/op is real; this benchmark
// only records it and never fails on a threshold.
func BenchmarkVerify(b *testing.B) {
	store := NewTrustStore()
	priv := demoPrivateKey(b)
	m := mustSign(b, store, priv, baseManifest())
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := Verify(store, m); err != nil {
			b.Fatalf("verify failed: %v", err)
		}
	}
}

// BenchmarkSign measures signing cost. Reported ns/op is real.
func BenchmarkSign(b *testing.B) {
	store := NewTrustStore()
	priv := demoPrivateKey(b)
	m := baseManifest()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mm := m
		if err := Sign(store, priv, &mm); err != nil {
			b.Fatalf("sign failed: %v", err)
		}
	}
}

// BenchmarkCanonicalBytes measures canonical serialization cost. Reported ns/op
// is real.
func BenchmarkCanonicalBytes(b *testing.B) {
	m := baseManifest()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = m.CanonicalBytes()
	}
}
