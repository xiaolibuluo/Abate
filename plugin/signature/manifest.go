/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package signature

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
)

// Manifest is the signed metadata of a plugin.
//
// The signature covers the canonical bytes of every field except Signature and
// SHA256. SHA256 is the content hash of those canonical bytes and is
// cross-checked during verification (see Verify). Because SHA256 is excluded
// from the signed payload, a manifest with a valid signature but a corrupted
// SHA256 field is rejected as ErrManifestHashMismatch rather than
// ErrSignatureInvalid, keeping the two failure modes distinct.
type Manifest struct {
	// ID is the plugin identifier, formatted as "scope/name".
	ID PluginID

	// Version is the plugin semantic version.
	Version SemVer

	// Type is the plugin kind (operator, reporter, or adapter).
	Type PluginType

	// CapabilitiesDecl is the raw capability declaration string (for example
	// "clock,rand") consumed by the capability sandbox.
	CapabilitiesDecl string

	// SHA256 is the content hash of the canonical (signature-excluded) bytes.
	SHA256 [32]byte

	// Signature is the detached Ed25519 signature over SignableDigest.
	Signature []byte

	// TrustAnchor is the id of the public key that must verify the signature.
	TrustAnchor string

	// BuildRef is an opaque build reference (commit, tag, or digest).
	BuildRef string
}

// CanonicalBytes returns the deterministic serialization of the manifest for
// signing and hashing.
//
// The field order is fixed (ID, Version, Type, CapabilitiesDecl, TrustAnchor,
// BuildRef) and every string is length-prefixed with a big-endian uint32 while
// integers use big-endian encoding. The Signature and SHA256 fields are
// deliberately excluded. Two equal manifests always yield byte-identical
// canonical bytes.
func (m Manifest) CanonicalBytes() []byte {
	var buf bytes.Buffer
	putString(&buf, string(m.ID))
	putUint64(&buf, uint64(m.Version.Major))
	putUint64(&buf, uint64(m.Version.Minor))
	putUint64(&buf, uint64(m.Version.Patch))
	putString(&buf, string(m.Type))
	putString(&buf, m.CapabilitiesDecl)
	putString(&buf, m.TrustAnchor)
	putString(&buf, m.BuildRef)
	return buf.Bytes()
}

// SignableDigest returns the SHA-256 digest over CanonicalBytes. This is the
// exact byte slice the private key signs and the public key verifies.
func (m Manifest) SignableDigest() [32]byte {
	return sha256.Sum256(m.CanonicalBytes())
}

// putUint64 appends v to b using big-endian encoding. The width is fixed at
// eight bytes so the serialization is byte-stable across 32/64-bit platforms.
func putUint64(b *bytes.Buffer, v uint64) {
	var tmp [8]byte
	binary.BigEndian.PutUint64(tmp[:], v)
	_, _ = b.Write(tmp[:])
}

// putString appends a length-prefixed string to b (big-endian uint64 length
// followed by the UTF-8 bytes). This avoids delimiter ambiguity.
func putString(b *bytes.Buffer, s string) {
	putUint64(b, uint64(len(s)))
	_, _ = b.WriteString(s)
}
