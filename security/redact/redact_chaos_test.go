/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package redact

import (
	"sort"
	"strings"
	"testing"
)

// The chaos vocabulary is fixed and ordered so every input is reproducible
// without a random source or map iteration.
var chaosFragments = []string{
	"Bearer abc",
	"token=abc",
	"SECRET=abc",
	"Authorization: Bearer abc",
	"authorization=Bearer abc",
	"api_key=abc",
	"token='abc'",
	"plain",
	"&",
	"/",
	"?",
	"next=1",
}

// ChaosInput assembles a deterministic input of index i without any random
// source, so the same index always yields the same byte string.
func chaosInput(i int) string {
	parts := make([]string, 0, 8)
	for j := 0; j < 8; j++ {
		parts = append(parts, chaosFragments[(i+j)%len(chaosFragments)])
	}
	return strings.Join(parts, " ")
}

// TestChaosRedactionIsIdempotent proves that over a wide deterministic input
// space no secret value survives and a second pass never changes the output.
func TestChaosRedactionIsIdempotent(t *testing.T) {
	for i := 0; i < 128; i++ {
		input := chaosInput(i)
		once := Text(input)
		if twice := Text(once); twice != once {
			t.Fatalf("iteration %d: Text not idempotent: %q -> %q", i, once, twice)
		}
		pathOnce := Path(input)
		if pathTwice := Path(pathOnce); pathTwice != pathOnce {
			t.Fatalf("iteration %d: Path not idempotent: %q -> %q", i, pathOnce, pathTwice)
		}
		if strings.Contains(once, "abc") || strings.Contains(pathOnce, "abc") {
			t.Fatalf("iteration %d: secret value survived: %q / %q", i, once, pathOnce)
		}
	}
}

// TestChaosDeterministicRepeatedRuns pins the determinism iron law: repeated
// calls on the same input are byte-identical for both entry points.
func TestChaosDeterministicRepeatedRuns(t *testing.T) {
	for i := 0; i < 64; i++ {
		input := chaosInput(i)
		want := Text(input)
		wantPath := Path(input)
		for k := 0; k < 4; k++ {
			if got := Text(input); got != want {
				t.Fatalf("iteration %d run %d: Text output drifted", i, k)
			}
			if got := Path(input); got != wantPath {
				t.Fatalf("iteration %d run %d: Path output drifted", i, k)
			}
		}
	}
}

// TestChaosSensitiveKeyTableIsSorted guards the precondition of the binary
// search and the ban on map-iteration-order dependence.
func TestChaosSensitiveKeyTableIsSorted(t *testing.T) {
	if !sort.StringsAreSorted(sensitiveKeys) {
		t.Fatal("sensitive key table is not sorted; binary search would be invalid")
	}
	for _, key := range sensitiveKeys {
		if !isSensitiveKey(key) {
			t.Fatalf("key %q not found by binary search", key)
		}
		if !isSensitiveKey(strings.ToUpper(key)) {
			t.Fatalf("upper-case key %q not recognized", key)
		}
	}
	if isSensitiveKey("") || isSensitiveKey("zzz_unknown") {
		t.Fatal("unknown key was accepted as sensitive")
	}
}
