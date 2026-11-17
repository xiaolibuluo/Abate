/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package signature

import "errors"

// Sentinel errors returned by the verification chain. Callers must test them
// with errors.Is.
var (
	// ErrSignatureInvalid reports that the Ed25519 signature did not verify.
	ErrSignatureInvalid = errors.New("plugin: signature invalid")

	// ErrTrustAnchorUnknown reports that the manifest trust anchor did not
	// resolve to any known public key.
	ErrTrustAnchorUnknown = errors.New("plugin: trust anchor unknown")

	// ErrManifestHashMismatch reports that the manifest SHA256 field did not
	// equal the signable digest.
	ErrManifestHashMismatch = errors.New("plugin: manifest hash mismatch")
)
