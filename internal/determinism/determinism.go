/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package determinism provides canonical hashing and a same-seed double-run
// assertion used to certify that simulator components are reproducible.
package determinism

import (
	"crypto/sha256"
	"encoding/json"
	"testing"
)

// HashState returns the SHA-256 digest of a canonical JSON encoding of v.
//
// The encoding/json package orders map keys lexicographically, so two values that
// are equal up to map traversal order produce identical digests. The function
// targets value types (structs, slices, maps of scalars); values that cannot be
// encoded fall back to a stable textual fingerprint derived from the error.
func HashState(v any) [32]byte {
	data, err := json.Marshal(v)
	if err != nil {
		data = []byte(err.Error())
	}
	return sha256.Sum256(data)
}

// MustMatch100 runs fn one hundred times under the same seed and fails the test
// if any run's output hash differs from the first, certifying that fn is
// reproducible for the given seed.
func MustMatch100(t *testing.T, seed uint64, fn func() []byte) {
	t.Helper()
	var prev [32]byte
	seen := false
	for i := 0; i < 100; i++ {
		out := fn()
		sum := sha256.Sum256(out)
		if seen {
			if sum != prev {
				t.Fatalf("determinism broken at run %d (seed %d): outputs differ", i, seed)
			}
		} else {
			prev = sum
			seen = true
		}
	}
}
