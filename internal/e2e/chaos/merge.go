/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package chaos

import (
	"crypto/sha256"
	"encoding/hex"
	"slices"
	"sort"
)

// Merge deduplicates inputs by the SHA-256 digest of their bytes and returns
// the survivors in ascending digest order. This matches -merge=1 semantics:
// identical corpus inputs collapse to one entry and the output has no map
// iteration order in it.
func Merge(inputs [][]byte) [][]byte {
	byDigest := make(map[string][]byte, len(inputs))
	digests := make([]string, 0, len(inputs))
	for _, input := range inputs {
		sum := sha256.Sum256(input)
		digest := hex.EncodeToString(sum[:])
		if _, exists := byDigest[digest]; !exists {
			byDigest[digest] = slices.Clone(input)
			digests = append(digests, digest)
		}
	}
	sort.Strings(digests)
	out := make([][]byte, 0, len(digests))
	for _, digest := range digests {
		out = append(out, byDigest[digest])
	}
	return out
}
