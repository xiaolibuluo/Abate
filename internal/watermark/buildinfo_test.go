/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package watermark

import "testing"

// TestBuildSnapshot proves the build snapshot carries every watermark field.
func TestBuildSnapshot(t *testing.T) {
	t.Parallel()
	info := Build()
	if info.Version != VersionLabel || info.Owner != OwnerName {
		t.Fatalf("unexpected provenance: version=%q owner=%q", info.Version, info.Owner)
	}
	if info.Fingerprint != FingerprintSHA256 {
		t.Fatalf("fingerprint = %q, want %q", info.Fingerprint, FingerprintSHA256)
	}
	if info.BuildTime != BuildTime || info.GitCommit != GitCommit {
		t.Fatalf("build metadata drifted: %q / %q", info.BuildTime, info.GitCommit)
	}
	if info.BuildTime == "" || info.GitCommit == "" {
		t.Fatal("build metadata must never be empty")
	}
}
