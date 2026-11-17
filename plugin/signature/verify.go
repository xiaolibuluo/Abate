/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package signature

import (
	"crypto/ed25519"
	"crypto/rand"
	"errors"
)

// Verify validates a manifest against the supplied trust store.
//
// It applies three independent checks, each failing with its own sentinel:
//
//  1. ErrTrustAnchorUnknown  — the manifest TrustAnchor does not resolve to a
//     known public key.
//  2. ErrSignatureInvalid     — the Ed25519 signature does not verify over the
//     signable digest.
//  3. ErrManifestHashMismatch — the manifest SHA256 field does not equal the
//     signable digest.
//
// There is no fall-through to an "unsigned" state: a manifest that fails any
// check is rejected. Verification never discloses the private key.
func Verify(store *TrustStore, m Manifest) error {
	pub := store.Resolve(m.TrustAnchor)
	if pub == nil {
		return ErrTrustAnchorUnknown
	}
	digest := m.SignableDigest()
	if !ed25519.Verify(pub, digest[:], m.Signature) {
		return ErrSignatureInvalid
	}
	if m.SHA256 != digest {
		return ErrManifestHashMismatch
	}
	return nil
}

// Sign attaches an Ed25519 signature to m in place, using priv.
//
// This helper exists for tests and signing tooling only; production code must
// never handle private keys. Sign requires that m.TrustAnchor already resolves
// to a known key in store (nil store skips that precondition) and sets m.SHA256
// to the signable digest so the hash cross-check passes during Verify.
func Sign(store *TrustStore, priv ed25519.PrivateKey, m *Manifest) error {
	if store != nil {
		if store.Resolve(m.TrustAnchor) == nil {
			return ErrTrustAnchorUnknown
		}
	}
	digest := m.SignableDigest()
	m.SHA256 = digest
	m.Signature = ed25519.Sign(priv, digest[:])
	return nil
}

// GenerateKey produces a fresh Ed25519 key pair from the crypto/rand source.
//
// The crypto/rand source is used (not math/rand) because this is a
// security-critical key generation path. This function is intended for signing
// tooling and tests, not for per-request runtime use.
func GenerateKey() (ed25519.PublicKey, ed25519.PrivateKey, error) {
	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, nil, errors.New("plugin: failed to generate ed25519 key: " + err.Error())
	}
	return pub, priv, nil
}
