/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package signature

import (
	"crypto/ed25519"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"sync"
)

// demoAnchorID is the id of the built-in demo trust anchor.
const demoAnchorID = "abate-demo"

// demoBuiltinPublicKey is the PUBLIC key of the built-in demo trust anchor.
//
// It is a FIXED, NON-PRODUCTION Ed25519 key used only for tests and demos. The
// matching private key is intentionally absent from this repository; production
// deployments must register their own release keys via Add.
var demoBuiltinPublicKey = []byte{
	0x77, 0x76, 0xe8, 0x70, 0xb9, 0x33, 0x54, 0xf2,
	0xa0, 0xb2, 0x4c, 0x23, 0xf2, 0xa3, 0x6c, 0xc4,
	0xe8, 0x0e, 0x22, 0x32, 0x18, 0xc1, 0xb9, 0x79,
	0x26, 0xfd, 0xd0, 0x18, 0x39, 0x6a, 0x2b, 0x9b,
}

// TrustStore holds the set of trusted public keys used to verify plugin
// signatures. It is safe for concurrent use.
//
// Keys are partitioned into an immutable built-in set (shipped with the binary)
// and a mutable user-added set. Fingerprints are tracked to enforce dedup.
type TrustStore struct {
	mu      sync.RWMutex
	builtin map[string]ed25519.PublicKey
	extra   map[string]ed25519.PublicKey
	byFP    map[string]string
}

// NewTrustStore returns an initialized trust store seeded with the built-in demo
// anchor. Production callers should register their own release keys via Add.
func NewTrustStore() *TrustStore {
	ts := &TrustStore{
		builtin: make(map[string]ed25519.PublicKey),
		extra:   make(map[string]ed25519.PublicKey),
		byFP:    make(map[string]string),
	}
	pub := ed25519.PublicKey(append([]byte(nil), demoBuiltinPublicKey...))
	ts.builtin[demoAnchorID] = pub
	ts.byFP[ts.Fingerprint(pub)] = demoAnchorID
	return ts
}

// Add registers a user-provided public key under the given id.
//
// Fingerprints are deduplicated: registering a key whose fingerprint already
// exists under a different id is rejected. The built-in anchor cannot be
// shadowed because it lives in a separate map.
func (t *TrustStore) Add(id string, pub ed25519.PublicKey) error {
	if id == "" {
		return errors.New("plugin: trust anchor id must not be empty")
	}
	if len(pub) != ed25519.PublicKeySize {
		return fmt.Errorf("plugin: trust anchor %q has wrong key size %d", id, len(pub))
	}
	fp := t.Fingerprint(pub)
	t.mu.Lock()
	defer t.mu.Unlock()
	if existing, ok := t.byFP[fp]; ok && existing != id {
		return fmt.Errorf("plugin: trust anchor fingerprint %s already registered as %s", fp, existing)
	}
	t.extra[id] = pub
	t.byFP[fp] = id
	return nil
}

// Builtin returns a copy of the immutable built-in key set.
func (t *TrustStore) Builtin() map[string]ed25519.PublicKey {
	t.mu.RLock()
	defer t.mu.RUnlock()
	out := make(map[string]ed25519.PublicKey, len(t.builtin))
	for k, v := range t.builtin {
		out[k] = v
	}
	return out
}

// Remove deletes a user-added key by id.
//
// The built-in demo anchor cannot be removed; attempting to do so returns an
// error. Removing an unknown id also returns an error.
func (t *TrustStore) Remove(id string) error {
	t.mu.Lock()
	defer t.mu.Unlock()
	pub, ok := t.extra[id]
	if !ok {
		if _, bok := t.builtin[id]; bok {
			return fmt.Errorf("plugin: cannot remove built-in trust anchor %q", id)
		}
		return fmt.Errorf("plugin: trust anchor %q not found", id)
	}
	delete(t.extra, id)
	fp := t.Fingerprint(pub)
	if t.byFP[fp] == id {
		delete(t.byFP, fp)
	}
	return nil
}

// Fingerprint returns the first 16 hex characters of sha256(pub). It is
// deterministic and used for deduplication and human-readable key identity.
func (*TrustStore) Fingerprint(pub ed25519.PublicKey) string {
	sum := sha256.Sum256(pub)
	return hex.EncodeToString(sum[:])[:16]
}

// Resolve returns the public key for the given anchor id, preferring built-in
// keys, or nil if the id is unknown. It is the single lookup point used by
// Verify and never falls back to an unsigned state.
func (t *TrustStore) Resolve(id string) ed25519.PublicKey {
	if id == "" {
		return nil
	}
	t.mu.RLock()
	defer t.mu.RUnlock()
	if p, ok := t.builtin[id]; ok {
		return p
	}
	if p, ok := t.extra[id]; ok {
		return p
	}
	return nil
}
