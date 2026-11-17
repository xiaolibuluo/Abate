/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package capability

import "testing"

// BenchmarkCheck measures the cost of a single capability gate. Reported ns/op is
// real; the benchmark only records it and never fails on a threshold.
func BenchmarkCheck(b *testing.B) {
	var caps Capabilities
	caps.Grant(CapClock)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := Check(caps, CapClock); err != nil {
			b.Fatalf("check failed: %v", err)
		}
	}
}

// BenchmarkParseGrant measures declaration parsing cost. Reported ns/op is real.
func BenchmarkParseGrant(b *testing.B) {
	const decl = "clock,rand,observe,fs,net"
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := ParseGrant(decl); err != nil {
			b.Fatalf("parse failed: %v", err)
		}
	}
}

// BenchmarkAuditAppend measures audit-logging cost. Reported ns/op is real.
func BenchmarkAuditAppend(b *testing.B) {
	log := NewAuditLog()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.Append(AuditEvent{Kind: "use", PluginID: "p", Capability: CapClock, Result: true})
	}
}
