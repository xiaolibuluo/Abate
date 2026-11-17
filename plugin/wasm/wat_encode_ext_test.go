/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package wasm

import (
	"bytes"
	"math"
	"testing"
)

// checkEncodedBytes compares an encoded buffer against its expected bytes.
func checkEncodedBytes(t *testing.T, label string, got, want []byte) {
	t.Helper()
	if !bytes.Equal(got, want) {
		t.Errorf("%s: got %v, want %v", label, got, want)
	}
}

// TestBlockTypeEncodesResultTypes verifies the void and value block type bytes.
func TestBlockTypeEncodesResultTypes(t *testing.T) {
	cases := []struct {
		name string
		in   valType
		want byte
	}{
		{name: "void", in: 0, want: 0x40},
		{name: "i32", in: vtI32, want: 0x7f},
		{name: "i64", in: vtI64, want: 0x7e},
	}
	for _, tc := range cases {
		if got := blockType(tc.in); got != tc.want {
			t.Errorf("%s: blockType(%#x) = %#x, want %#x", tc.name, tc.in, got, tc.want)
		}
	}
}

// TestAppendMemArgEncodesAlignThenOffset verifies the memarg immediate layout.
func TestAppendMemArgEncodesAlignThenOffset(t *testing.T) {
	cases := []struct {
		name string
		in   watInstr
		want []byte
	}{
		{name: "zero align and offset", in: watInstr{}, want: []byte{0x00, 0x00}},
		{name: "align two offset eight", in: watInstr{memAlgn: 2, memOff: 8}, want: []byte{0x02, 0x08}},
		{name: "large offset", in: watInstr{memAlgn: 3, memOff: 128}, want: []byte{0x03, 0x80, 0x01}},
	}
	for _, tc := range cases {
		instr := tc.in
		checkEncodedBytes(t, tc.name, appendMemArg(nil, &instr), tc.want)
	}
	prefixed := appendMemArg([]byte{0x28}, &watInstr{memAlgn: 2, memOff: 4})
	checkEncodedBytes(t, "prefixed buffer", prefixed, []byte{0x28, 0x02, 0x04})
}

// TestAppendLimitsEncodesMinAndMax verifies both limits variants.
func TestAppendLimitsEncodesMinAndMax(t *testing.T) {
	cases := []struct {
		name string
		lo   uint32
		hi   uint32
		want []byte
	}{
		{name: "min only", lo: 1, hi: 0, want: []byte{0x00, 0x01}},
		{name: "min and max", lo: 2, hi: 8, want: []byte{0x01, 0x02, 0x08}},
		{name: "multi byte", lo: 128, hi: 256, want: []byte{0x01, 0x80, 0x01, 0x80, 0x02}},
	}
	for _, tc := range cases {
		checkEncodedBytes(t, tc.name, appendLimits(nil, tc.lo, tc.hi), tc.want)
	}
}

// TestSameTypesComparesValueVectors verifies length and element comparisons.
func TestSameTypesComparesValueVectors(t *testing.T) {
	cases := []struct {
		name string
		a    []valType
		b    []valType
		want bool
	}{
		{name: "both nil", a: nil, b: nil, want: true},
		{name: "equal vectors", a: []valType{vtI32, vtI64}, b: []valType{vtI32, vtI64}, want: true},
		{name: "different lengths", a: []valType{vtI32}, b: []valType{vtI32, vtI64}, want: false},
		{name: "different element", a: []valType{vtI32}, b: []valType{vtI64}, want: false},
	}
	for _, tc := range cases {
		if got := sameTypes(tc.a, tc.b); got != tc.want {
			t.Errorf("%s: sameTypes(%v, %v) = %v, want %v", tc.name, tc.a, tc.b, got, tc.want)
		}
	}
}

// TestAppendULEDEncodesUnsignedIntegers verifies canonical LEB128 output.
func TestAppendULEDEncodesUnsignedIntegers(t *testing.T) {
	cases := []struct {
		name string
		in   uint64
		want []byte
	}{
		{name: "zero", in: 0, want: []byte{0x00}},
		{name: "one", in: 1, want: []byte{0x01}},
		{name: "max single byte", in: 127, want: []byte{0x7f}},
		{name: "two bytes", in: 128, want: []byte{0x80, 0x01}},
		{name: "two bytes high", in: 255, want: []byte{0xff, 0x01}},
		{name: "three bytes", in: 624485, want: []byte{0xe5, 0x8e, 0x26}},
		{name: "max uint64", in: math.MaxUint64, want: []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01}},
	}
	for _, tc := range cases {
		checkEncodedBytes(t, tc.name, appendULED(nil, tc.in), tc.want)
	}
	checkEncodedBytes(t, "prefixed buffer", appendULED([]byte{0xaa}, 130), []byte{0xaa, 0x82, 0x01})
}

// TestAppendILEDEncodesSignedIntegers verifies canonical signed LEB128 output.
func TestAppendILEDEncodesSignedIntegers(t *testing.T) {
	cases := []struct {
		name string
		in   int64
		want []byte
	}{
		{name: "zero", in: 0, want: []byte{0x00}},
		{name: "positive small", in: 42, want: []byte{0x2a}},
		{name: "negative one", in: -1, want: []byte{0x7f}},
		{name: "sign boundary positive", in: 64, want: []byte{0xc0, 0x00}},
		{name: "sign boundary negative", in: -64, want: []byte{0x40}},
		{name: "negative sixty five", in: -65, want: []byte{0xbf, 0x7f}},
		{name: "large positive", in: 128, want: []byte{0x80, 0x01}},
		{name: "negative one twenty eight", in: -128, want: []byte{0x80, 0x7f}},
		{name: "large negative", in: -123456, want: []byte{0xc0, 0xbb, 0x78}},
	}
	for _, tc := range cases {
		checkEncodedBytes(t, tc.name, appendILED(nil, tc.in), tc.want)
	}
}

// TestAppendByteHelpers verifies the shared byte-vector helpers.
func TestAppendByteHelpers(t *testing.T) {
	checkEncodedBytes(t, "appendULEN", appendULEN(nil, 130), []byte{0x82, 0x01})
	checkEncodedBytes(t, "appendValTypeVec", appendValTypeVec(nil, []valType{vtI32, vtI64}), []byte{0x02, 0x7f, 0x7e})
	checkEncodedBytes(t, "appendName", appendName(nil, "mem"), []byte{0x03, 0x6d, 0x65, 0x6d})
	checkEncodedBytes(t, "appendSection", appendSection([]byte{0x01}, 7, []byte{0x02, 0x03}), []byte{0x01, 0x07, 0x02, 0x02, 0x03})
}

// TestEncodeImportsCoversFuncAndMemory verifies both import kinds.
func TestEncodeImportsCoversFuncAndMemory(t *testing.T) {
	funcOnly := &watModule{imports: []watFuncImport{{module: "env", name: "f", params: []valType{vtI32}}}}
	checkEncodedBytes(t, "func import", encodeImports(funcOnly, []int{0}),
		[]byte{0x01, 0x03, 0x65, 0x6e, 0x76, 0x01, 0x66, 0x00, 0x00})
	memOnly := &watModule{memImport: &watMemImport{module: "m", name: "mem"}, memMin: 1, memMax: 4}
	checkEncodedBytes(t, "memory import", encodeImports(memOnly, nil),
		[]byte{0x01, 0x01, 0x6d, 0x03, 0x6d, 0x65, 0x6d, 0x02, 0x01, 0x01, 0x04})
	both := &watModule{
		imports:   []watFuncImport{{module: "env", name: "f"}},
		memImport: &watMemImport{module: "m", name: "mem"},
		memMin:    1,
	}
	checkEncodedBytes(t, "mixed imports", encodeImports(both, []int{0}),
		[]byte{0x02, 0x03, 0x65, 0x6e, 0x76, 0x01, 0x66, 0x00, 0x00,
			0x01, 0x6d, 0x03, 0x6d, 0x65, 0x6d, 0x02, 0x00, 0x01})
}

// TestEncodeMemorySkipsImportedMemory verifies the nil-memory branch.
func TestEncodeMemorySkipsImportedMemory(t *testing.T) {
	defined := &watModule{memMin: 1}
	checkEncodedBytes(t, "defined memory", encodeMemory(defined), []byte{0x01, 0x00, 0x01})
	imported := &watModule{memImport: &watMemImport{module: "m", name: "mem"}}
	if got := encodeMemory(imported); got != nil {
		t.Errorf("encodeMemory(imported) = %v, want nil", got)
	}
}

// TestEncodeExportsCoversFuncAndMemory verifies both export kinds and the empty case.
func TestEncodeExportsCoversFuncAndMemory(t *testing.T) {
	mod := &watModule{
		imports:   []watFuncImport{{module: "env", name: "f"}},
		funcs:     []watFunc{{exportName: "run"}, {exportName: ""}},
		memExport: "mem",
	}
	checkEncodedBytes(t, "func and memory exports", encodeExports(mod),
		[]byte{0x02, 0x03, 0x72, 0x75, 0x6e, 0x00, 0x01, 0x03, 0x6d, 0x65, 0x6d, 0x02, 0x00})
	checkEncodedBytes(t, "no exports", encodeExports(&watModule{}), []byte{0x00})
}

// TestEncodeCodePropagatesInstructionErrors verifies the success and failure paths.
func TestEncodeCodePropagatesInstructionErrors(t *testing.T) {
	good := &watModule{funcs: []watFunc{{body: []watInstr{{op: "nop"}}}}}
	checkEncodedBytes(t, "single function", mustEncodeCode(t, good), []byte{0x01, 0x03, 0x00, 0x01, 0x0b})
	bad := &watModule{funcs: []watFunc{{body: []watInstr{{op: "bogus"}}}}}
	if _, err := encodeCode(bad); err == nil {
		t.Fatal("encodeCode accepted an unsupported instruction")
	}
}

// mustEncodeCode calls encodeCode and fails the test on any error.
func mustEncodeCode(t *testing.T, mod *watModule) []byte {
	t.Helper()
	b, err := encodeCode(mod)
	if err != nil {
		t.Fatalf("encodeCode: %v", err)
	}
	return b
}

// TestEncodeModulePropagatesCodeErrors verifies the failure branch of encodeModule.
func TestEncodeModulePropagatesCodeErrors(t *testing.T) {
	bad := &watModule{funcs: []watFunc{{body: []watInstr{{op: "bogus"}}}}}
	if _, err := encodeModule(bad); err == nil {
		t.Fatal("encodeModule accepted an unsupported instruction")
	}
}

// TestEncodeModuleOmitsMemorySectionForImports verifies the exact section layout.
func TestEncodeModuleOmitsMemorySectionForImports(t *testing.T) {
	mod := &watModule{memImport: &watMemImport{module: "m", name: "mem"}, memMin: 1}
	b, err := encodeModule(mod)
	if err != nil {
		t.Fatalf("encodeModule: %v", err)
	}
	want := []byte{
		0x00, 0x61, 0x73, 0x6d, 0x01, 0x00, 0x00, 0x00,
		0x01, 0x01, 0x00,
		0x02, 0x0a, 0x01, 0x01, 0x6d, 0x03, 0x6d, 0x65, 0x6d, 0x02, 0x00, 0x01,
		0x03, 0x01, 0x00,
		0x07, 0x01, 0x00,
		0x0a, 0x01, 0x00,
	}
	checkEncodedBytes(t, "imported memory module", b, want)
}
