/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package air

import "testing"

// fuzzUnit builds a deterministic unit without requiring a testing.T helper.
func fuzzUnit() *Unit {
	u := NewUnit(NewPackageRef("fuzz", []string{"z", "a"}))
	leaf := NewNode(KindIdent, NewPosition(1, 1, 0, 3), "abc")
	Annotate(leaf, SemTypeRef, "string")
	root := NewNode(KindFuncDecl, NewPosition(1, 1, 0, 3), "main", leaf)
	if err := u.AddFile(NewFileRef("fuzz.go", []byte("abc")), root); err != nil {
		panic(err)
	}
	u.Annotations["seed"] = "deterministic"
	return u
}

// FuzzDecode feeds deterministic seed corpora and arbitrary bytes to the
// loader. It requires no panic, sentinel-classified errors, byte-stable double
// runs, and a lossless re-encode of every archive that decodes successfully.
func FuzzDecode(f *testing.F) {
	valid, err := NewEncoder().Encode(fuzzUnit())
	if err != nil {
		f.Fatalf("seed encode: %v", err)
	}
	seeds := [][]byte{
		{},
		[]byte("ABAT"),
		[]byte("ABAT\x00\x01\x00\x01"),
		make([]byte, headerSize+checksumSize),
		valid,
	}
	for _, seed := range seeds {
		f.Add(seed)
	}
	mutated := append([]byte(nil), valid...)
	mutated[0] ^= 0x01
	f.Add(mutated)
	truncated := append([]byte(nil), valid[:len(valid)/2]...)
	f.Add(truncated)
	body := archiveBody(valid)
	body = append(body, appendSection(nil, unknownTailID, []byte("future"))...)
	f.Add(reseal(body))

	f.Fuzz(func(t *testing.T, data []byte) {
		first, errA := NewDecoder().Decode(data)
		if errA != nil && !isAIRSentinel(errA) {
			t.Fatalf("unclassified error: %v", errA)
		}
		second, errB := NewDecoder().Decode(data)
		if errorText(errA) != errorText(errB) {
			t.Fatalf("diagnostic drift: %q vs %q", errorText(errA), errorText(errB))
		}
		if digestOfUnit(first) != digestOfUnit(second) {
			t.Fatal("decoded units drifted between identical inputs")
		}
		if errA != nil {
			return
		}
		reencoded, err := NewEncoder().Encode(first)
		if err != nil {
			t.Fatalf("re-encode: %v", err)
		}
		again, err := NewDecoder().Decode(reencoded)
		if err != nil {
			t.Fatalf("re-decode: %v", err)
		}
		if err := CompareUnits(first, again); err != nil {
			t.Fatalf("fuzz round trip: %v", err)
		}
	})
}
