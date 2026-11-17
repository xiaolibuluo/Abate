/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package parse

import (
	"errors"
	"testing"
)

// FuzzParse feeds deterministic seed documents to the strict CES parser and
// requires no panic, ParseError-typed diagnostics and byte-stable double runs.
func FuzzParse(f *testing.F) {
	for _, seed := range [][]byte{
		[]byte(baseYAML),
		[]byte(baseJSON),
		[]byte(""),
		[]byte("   "),
		[]byte("{"),
		[]byte("apiVersion: ces.tencent.com/v1\nkind: Experiment\n"),
		[]byte(`{"apiVersion":"ces.tencent.com/v1","kind":"Experiment","bogus":1}`),
		[]byte("ε ⊗ a⁻¹ operator://missing"),
	} {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, data []byte) {
		first, errA := Parse(data, FormatAuto)
		if errA != nil {
			var pe *ParseError
			if !errors.As(errA, &pe) {
				t.Fatalf("parser returned a non-ParseError failure: %v", errA)
			}
		}
		second, errB := Parse(data, FormatAuto)
		if fuzzErrorText(errA) != fuzzErrorText(errB) {
			t.Fatalf("diagnostics drifted: %q vs %q",
				fuzzErrorText(errA), fuzzErrorText(errB))
		}
		if digestOf(t, first) != digestOf(t, second) {
			t.Fatal("parsed documents drifted between identical inputs")
		}
	})
}

// fuzzErrorText renders a comparable diagnostic string.
func fuzzErrorText(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}
