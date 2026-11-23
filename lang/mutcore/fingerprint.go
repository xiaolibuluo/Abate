/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/
package mutcore

import "crypto/sha256"

// FingerprintTokens hashes the ordered token set into the Target digest.
// Every token contributes its node reference and original text, separated
// by zero bytes, so two sessions agree on the digest exactly when their
// scanned snapshots agree.
func FingerprintTokens(tokens []Token) [32]byte {
	h := sha256.New()
	for _, tok := range tokens {
		_, _ = h.Write([]byte(tok.Node()))
		_, _ = h.Write([]byte{0})
		_, _ = h.Write([]byte(tok.Text))
		_, _ = h.Write([]byte{0})
	}
	var out [32]byte
	copy(out[:], h.Sum(nil))
	return out
}

// FingerprintTests hashes the ordered test identifiers plus the covered
// file set into the Suite digest. Including the measured file set keeps
// the digest sensitive to coverage changes, which the report fingerprint
// then reflects.
func FingerprintTests(tests []TestRef, files []string) [32]byte {
	h := sha256.New()
	for _, test := range tests {
		_, _ = h.Write([]byte(test.ID))
		_, _ = h.Write([]byte{0})
	}
	for _, file := range files {
		_, _ = h.Write([]byte(file))
		_, _ = h.Write([]byte{0})
	}
	var out [32]byte
	copy(out[:], h.Sum(nil))
	return out
}
