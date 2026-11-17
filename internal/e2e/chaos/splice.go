/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package chaos

import (
	"bytes"
	"slices"
)

// spliceChunk is the deterministic two-byte chunk inserted by the generator.
var spliceChunk = []byte{0x00, 0xff}

// Splice inserts chunk at a bounded byte offset of content and returns a fresh
// buffer. Offsets outside the source are clamped to its ends.
func Splice(content []byte, at int, chunk []byte) []byte {
	if at < 0 {
		at = 0
	}
	if at > len(content) {
		at = len(content)
	}
	out := make([]byte, 0, len(content)+len(chunk))
	out = append(out, content[:at]...)
	out = append(out, chunk...)
	out = append(out, content[at:]...)
	return out
}

// DropToken removes every occurrence of token from content. An empty token is
// a no-op that still returns an independent copy.
func DropToken(content []byte, token []byte) []byte {
	if len(token) == 0 {
		return slices.Clone(content)
	}
	return bytes.ReplaceAll(content, token, []byte{})
}
