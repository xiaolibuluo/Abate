/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package chaos

import "slices"

// g1Minimal holds the two truncated Go bodies that reproduce the G1 panic in
// the regression suite. They are the first halves of complete synthetic files.
var g1Minimal = [][]byte{
	[]byte("package replica\n\nfunc (r *Replica) syncFrom(peer string) error"),
	[]byte("package stats\n\nfunc TestReplicaStats(t *testing.T)"),
}

// g1Tail is repeated to build a complete source twice as long as a minimal
// reproduction, so a half truncation cuts exactly at that reproduction.
const g1Tail = "\nvalue = value + 1\n"

// TruncateHalf returns the first half of content at a byte boundary. The result
// is an independent buffer, so the caller's source bytes are never modified.
func TruncateHalf(content []byte) []byte {
	cut := len(content) / 2
	return slices.Clone(content[:cut])
}

// G1Reproductions returns the two truncated bodies used by the G1 regression
// suite, each produced by TruncateHalf on a complete synthetic source.
func G1Reproductions() [][]byte {
	out := make([][]byte, 0, len(g1Minimal))
	for _, minimal := range g1Minimal {
		out = append(out, TruncateHalf(g1FullSource(minimal)))
	}
	return out
}

// g1FullSource builds a source whose first half is exactly minimal.
func g1FullSource(minimal []byte) []byte {
	want := len(minimal) * 2
	tail := []byte(g1Tail)
	out := make([]byte, 0, want)
	out = append(out, minimal...)
	for len(out) < want {
		out = append(out, tail...)
	}
	return out[:want]
}
