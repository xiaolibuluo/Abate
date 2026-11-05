// Abate —— 止息混沌 · 混沌工程开放规范（CES）与参考实现
// Abate — where chaos abates. Open Chaos Experiment Specification (CES) & reference implementation.
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
//
// 本文件以 GNU Affero General Public License v3（AGPLv3）发布；商业用途须事先获得
// 小黎的部落书面授权，详见 COMMERCIAL.md。
// This file is published under the GNU Affero General Public License v3 (AGPLv3).
// Commercial use requires prior written authorization from 小黎的部落.
// See COMMERCIAL.md for details.

package hub

import (
	"fmt"
	"testing"
)

// BenchmarkSeedRegistryAllCodes 基准测试：种子注册表 AllCodes 排序输出。
func BenchmarkSeedRegistryAllCodes(b *testing.B) {
	r := NewSeedRegistry()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = r.AllCodes()
	}
}

// BenchmarkSeedRegistryAllCodeStrings 基准测试：种子注册表 AllCodeStrings。
func BenchmarkSeedRegistryAllCodeStrings(b *testing.B) {
	r := NewSeedRegistry()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = r.AllCodeStrings()
	}
}

// BenchmarkSeedRegistryLookup 基准测试：种子注册表 Lookup。
func BenchmarkSeedRegistryLookup(b *testing.B) {
	r := NewSeedRegistry()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.Lookup("F4.3")
	}
}

// BenchmarkRegistryRegister 基准测试：注册表 Register。
func BenchmarkRegistryRegister(b *testing.B) {
	r := NewRegistry("v0.2")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = r.Register(SubCode{Code: fmt.Sprintf("F9.%d", i%100), Name: "bench", Params: "x"})
	}
}

// BenchmarkComputeDiff 基准测试：差异计算。
func BenchmarkComputeDiff(b *testing.B) {
	old := NewSeedRegistry()
	new := NewSeedRegistry()
	_ = new.Register(SubCode{Code: "F7.1", Name: "新", Params: "x"})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ComputeDiff(old, new)
	}
}

// BenchmarkAnnouncementFormat 基准测试：公告格式化。
func BenchmarkAnnouncementFormat(b *testing.B) {
	old := NewSeedRegistry()
	new := NewSeedRegistry()
	_ = new.Register(SubCode{Code: "F7.1", Name: "新", Params: "x"})
	a := NewAnnouncement(old, new)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = a.Format()
	}
}
