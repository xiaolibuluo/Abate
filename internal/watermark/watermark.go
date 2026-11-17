/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package watermark defines the Abate copyright watermark and its verifiable
// SHA-256 fingerprint. It is the source layer of the three-layer watermark.
package watermark

import (
	"crypto/sha256"
	"encoding/hex"
)

const (
	// OwnerName is the copyright owner carried by every Abate source header.
	OwnerName = "小黎的部落"
	// SourceURL is the canonical upstream repository of this project.
	SourceURL = "https://github.com/xiaolibuluo/Abate"
	// Canonical is the exact byte string whose digest is the digital fingerprint.
	Canonical = "xiaolibuluo/Abate © 小黎的部落"
	// FingerprintSHA256 is the lowercase hex SHA-256 digest of Canonical.
	FingerprintSHA256 = "e92fa6a0a80b33498123d5545e07e5b2c8b1ce30f220e3a933644bd8118b81bc"
	// VersionLabel is the provenance version exposed by the watermark package.
	VersionLabel = "0.1.0"
)

// Independent anchors used to detect drift of the exported constants above.
const (
	ownerAnchor     = "小黎的部落"
	sourceAnchor    = "https://github.com/xiaolibuluo/Abate"
	canonicalAnchor = "xiaolibuluo/Abate © 小黎的部落"
)

// Owner returns the copyright owner string of the watermark.
func Owner() string { return OwnerName }

// Source returns the canonical upstream repository URL.
func Source() string { return SourceURL }

// Fingerprint returns the lowercase hex SHA-256 digital fingerprint.
func Fingerprint() string { return FingerprintSHA256 }

// Version returns the build provenance version label.
func Version() string { return VersionLabel }

// Verify proves the exported constants agree with the independent anchors and
// that the fingerprint is the real SHA-256 digest of the canonical string.
func Verify() bool {
	return verifyCandidate(Canonical, FingerprintSHA256, OwnerName, SourceURL)
}

// verifyCandidate checks one candidate set against the immutable anchors.
func verifyCandidate(canonical, fingerprint, owner, source string) bool {
	if owner != ownerAnchor || source != sourceAnchor || canonical != canonicalAnchor {
		return false
	}
	digest := sha256.Sum256([]byte(canonical))
	return hex.EncodeToString(digest[:]) == fingerprint
}
