/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package watermark

// BuildTime is the UTC build timestamp injected through -ldflags at build time.
var BuildTime = "unknown"

// GitCommit is the git revision injected through -ldflags at build time.
var GitCommit = "unknown"

// BuildInfo is a snapshot of the build-layer provenance of this binary.
type BuildInfo struct {
	// Version is the provenance version label.
	Version string
	// BuildTime is the injected build timestamp.
	BuildTime string
	// GitCommit is the injected source revision.
	GitCommit string
	// Owner is the copyright watermark owner.
	Owner string
	// Fingerprint is the digital fingerprint of the canonical watermark string.
	Fingerprint string
}

// Build returns a BuildInfo snapshot built from the current constants and
// the values injected through -ldflags.
func Build() BuildInfo {
	return BuildInfo{
		Version:     VersionLabel,
		BuildTime:   BuildTime,
		GitCommit:   GitCommit,
		Owner:       OwnerName,
		Fingerprint: FingerprintSHA256,
	}
}
