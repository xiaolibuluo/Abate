/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package wasm

import (
	"bytes"
	"testing"
)

// instrEncoder is the shared signature of the instruction encoders under test.
type instrEncoder func(*watInstr, []byte) ([]byte, error)

// encodingCase is one expected instruction encoding.
type encodingCase struct {
	name  string
	instr watInstr
	want  []byte
}

// runInstrEncodings checks a table of instructions against their expected bytes.
func runInstrEncodings(t *testing.T, cases []encodingCase, enc instrEncoder) {
	t.Helper()
	for i := range cases {
		item := cases[i]
		instr := item.instr
		got, err := enc(&instr, nil)
		if err != nil {
			t.Errorf("%s: encode error = %v, want nil", item.name, err)
			continue
		}
		if !bytes.Equal(got, item.want) {
			t.Errorf("%s: encoded = %v, want %v", item.name, got, item.want)
		}
	}
}

// runInstrEncodingErrors checks that unsupported opcodes are rejected.
func runInstrEncodingErrors(t *testing.T, cases []watInstr, enc instrEncoder) {
	t.Helper()
	for i := range cases {
		instr := cases[i]
		if _, err := enc(&instr, nil); err == nil {
			t.Errorf("op %q: expected an encoding error", instr.op)
		}
	}
}

// TestEncodeInstrControlCoversOpcodes verifies every control-flow opcode encoding.
func TestEncodeInstrControlCoversOpcodes(t *testing.T) {
	cases := []encodingCase{
		{name: "loop void", instr: watInstr{op: "loop", block: &watBlock{kind: "loop"}}, want: []byte{0x03, 0x40, 0x0b}},
		{name: "block i32 result", instr: watInstr{op: "block", block: &watBlock{kind: "block", result: vtI32}}, want: []byte{0x02, 0x7f, 0x0b}},
		{name: "block with body", instr: watInstr{op: "block", block: &watBlock{kind: "block", then: []watInstr{{op: "nop"}}}}, want: []byte{0x02, 0x40, 0x01, 0x0b}},
		{name: "if without else", instr: watInstr{op: "if", block: &watBlock{kind: "if"}}, want: []byte{0x04, 0x40, 0x0b}},
		{name: "if with else", instr: watInstr{op: "if", block: &watBlock{kind: "if", then: []watInstr{{op: "nop"}}, els: []watInstr{{op: "i32.const", imm: 42}}}}, want: []byte{0x04, 0x40, 0x01, 0x05, 0x41, 0x2a, 0x0b}},
		{name: "else", instr: watInstr{op: "else"}, want: []byte{0x05}},
		{name: "end", instr: watInstr{op: "end"}, want: []byte{0x0b}},
		{name: "br", instr: watInstr{op: "br", idx: 3}, want: []byte{0x0c, 0x03}},
		{name: "br_if multi byte", instr: watInstr{op: "br_if", idx: 130}, want: []byte{0x0d, 0x82, 0x01}},
		{name: "return", instr: watInstr{op: "return"}, want: []byte{0x0f}},
		{name: "drop", instr: watInstr{op: "drop"}, want: []byte{0x1a}},
		{name: "call", instr: watInstr{op: "call", idx: 2}, want: []byte{0x10, 0x02}},
		{name: "unreachable", instr: watInstr{op: "unreachable"}, want: []byte{0x00}},
		{name: "nop", instr: watInstr{op: "nop"}, want: []byte{0x01}},
	}
	runInstrEncodings(t, cases, encodeInstrControl)
}

// TestEncodeInstrControlPropagatesErrors verifies nested and unknown opcode failures.
func TestEncodeInstrControlPropagatesErrors(t *testing.T) {
	bad := []watInstr{
		{op: "bogus"},
		{op: "block", block: &watBlock{kind: "block", then: []watInstr{{op: "bogus"}}}},
		{op: "if", block: &watBlock{kind: "if", then: []watInstr{{op: "bogus"}}}},
		{op: "if", block: &watBlock{kind: "if", els: []watInstr{{op: "bogus"}}}},
	}
	runInstrEncodingErrors(t, bad, encodeInstrControl)
}

// TestEncodeInstrNumericCoversOpcodes verifies every numeric and memory opcode.
func TestEncodeInstrNumericCoversOpcodes(t *testing.T) {
	cases := []encodingCase{
		{name: "i32.const positive", instr: watInstr{op: "i32.const", imm: 42}, want: []byte{0x41, 0x2a}},
		{name: "i32.const negative", instr: watInstr{op: "i32.const", imm: -1}, want: []byte{0x41, 0x7f}},
		{name: "i64.const multi byte", instr: watInstr{op: "i64.const", imm: 300}, want: []byte{0x42, 0xac, 0x02}},
		{name: "local.get", instr: watInstr{op: "local.get", idx: 1}, want: []byte{0x20, 0x01}},
		{name: "local.set", instr: watInstr{op: "local.set", idx: 2}, want: []byte{0x21, 0x02}},
		{name: "local.tee", instr: watInstr{op: "local.tee", idx: 3}, want: []byte{0x22, 0x03}},
		{name: "i32.eq", instr: watInstr{op: "i32.eq"}, want: []byte{0x46}},
		{name: "i32.ne", instr: watInstr{op: "i32.ne"}, want: []byte{0x47}},
		{name: "i32.lt_s", instr: watInstr{op: "i32.lt_s"}, want: []byte{0x48}},
		{name: "i32.gt_s", instr: watInstr{op: "i32.gt_s"}, want: []byte{0x4a}},
		{name: "i32.le_s", instr: watInstr{op: "i32.le_s"}, want: []byte{0x4c}},
		{name: "i32.ge_s", instr: watInstr{op: "i32.ge_s"}, want: []byte{0x4e}},
		{name: "i32.add", instr: watInstr{op: "i32.add"}, want: []byte{0x6a}},
		{name: "i32.sub", instr: watInstr{op: "i32.sub"}, want: []byte{0x6b}},
		{name: "i32.mul", instr: watInstr{op: "i32.mul"}, want: []byte{0x6c}},
		{name: "i32.and", instr: watInstr{op: "i32.and"}, want: []byte{0x71}},
		{name: "i32.or", instr: watInstr{op: "i32.or"}, want: []byte{0x72}},
		{name: "i32.xor", instr: watInstr{op: "i32.xor"}, want: []byte{0x73}},
		{name: "i32.shl", instr: watInstr{op: "i32.shl"}, want: []byte{0x74}},
		{name: "i32.shr_s", instr: watInstr{op: "i32.shr_s"}, want: []byte{0x75}},
		{name: "i32.clz", instr: watInstr{op: "i32.clz"}, want: []byte{0x67}},
		{name: "i32.ctz", instr: watInstr{op: "i32.ctz"}, want: []byte{0x68}},
		{name: "i32.popcnt", instr: watInstr{op: "i32.popcnt"}, want: []byte{0x69}},
		{name: "i32.load", instr: watInstr{op: "i32.load", memAlgn: 2}, want: []byte{0x28, 0x02, 0x00}},
		{name: "i64.load", instr: watInstr{op: "i64.load", memAlgn: 3, memOff: 8}, want: []byte{0x29, 0x03, 0x08}},
		{name: "i32.store", instr: watInstr{op: "i32.store", memAlgn: 2, memOff: 16}, want: []byte{0x36, 0x02, 0x10}},
		{name: "i64.store", instr: watInstr{op: "i64.store", memAlgn: 3, memOff: 128}, want: []byte{0x37, 0x03, 0x80, 0x01}},
	}
	runInstrEncodings(t, cases, encodeInstrNumeric)
	runInstrEncodingErrors(t, []watInstr{{op: "bogus"}}, encodeInstrNumeric)
}

// TestEncodeInstrMemoryCoversOpcodes verifies the four load/store opcodes directly.
func TestEncodeInstrMemoryCoversOpcodes(t *testing.T) {
	cases := []encodingCase{
		{name: "i32.load", instr: watInstr{op: "i32.load", memAlgn: 2, memOff: 4}, want: []byte{0x28, 0x02, 0x04}},
		{name: "i64.load", instr: watInstr{op: "i64.load", memAlgn: 3, memOff: 8}, want: []byte{0x29, 0x03, 0x08}},
		{name: "i32.store", instr: watInstr{op: "i32.store", memAlgn: 2, memOff: 16}, want: []byte{0x36, 0x02, 0x10}},
		{name: "i64.store", instr: watInstr{op: "i64.store", memAlgn: 3, memOff: 128}, want: []byte{0x37, 0x03, 0x80, 0x01}},
	}
	runInstrEncodings(t, cases, encodeInstrMemory)
	runInstrEncodingErrors(t, []watInstr{{op: "bogus"}, {op: "f32.load"}}, encodeInstrMemory)
}

// TestEncodeInstrDispatchesBothFamilies verifies the top-level opcode dispatch.
func TestEncodeInstrDispatchesBothFamilies(t *testing.T) {
	cases := []encodingCase{
		{name: "control route", instr: watInstr{op: "nop"}, want: []byte{0x01}},
		{name: "numeric route", instr: watInstr{op: "i32.add"}, want: []byte{0x6a}},
		{name: "memory route", instr: watInstr{op: "i32.load", memAlgn: 2, memOff: 4}, want: []byte{0x28, 0x02, 0x04}},
	}
	runInstrEncodings(t, cases, encodeInstr)
}

// TestEncodeInstrsSequentialAndErrors verifies sequence encoding and propagation.
func TestEncodeInstrsSequentialAndErrors(t *testing.T) {
	seq := []watInstr{{op: "nop"}, {op: "i32.const", imm: 7}, {op: "drop"}}
	got, err := encodeInstrs(seq, nil)
	if err != nil {
		t.Fatalf("encodeInstrs: %v", err)
	}
	checkEncodedBytes(t, "sequence", got, []byte{0x01, 0x41, 0x07, 0x1a})
	got, err = encodeInstrs(nil, []byte{0xaa})
	if err != nil {
		t.Fatalf("encodeInstrs(empty): %v", err)
	}
	checkEncodedBytes(t, "empty with prefix", got, []byte{0xaa})
	if _, err := encodeInstrs([]watInstr{{op: "bogus"}}, nil); err == nil {
		t.Fatal("encodeInstrs accepted an unsupported instruction")
	}
}

// TestAppendLocalsGroupsConsecutiveTypes verifies locals declaration grouping.
func TestAppendLocalsGroupsConsecutiveTypes(t *testing.T) {
	cases := []struct {
		name   string
		locals []valType
		want   []byte
	}{
		{name: "empty", locals: nil, want: []byte{0x00}},
		{name: "single", locals: []valType{vtI32}, want: []byte{0x01, 0x01, 0x7f}},
		{name: "grouped", locals: []valType{vtI32, vtI32, vtI64}, want: []byte{0x02, 0x02, 0x7f, 0x01, 0x7e}},
		{name: "interleaved", locals: []valType{vtI32, vtI64, vtI32}, want: []byte{0x03, 0x01, 0x7f, 0x01, 0x7e, 0x01, 0x7f}},
	}
	for _, tc := range cases {
		checkEncodedBytes(t, tc.name, appendLocals(nil, tc.locals), tc.want)
	}
}

// TestEncodeFuncBodyAndErrors verifies the function body layout and failure path.
func TestEncodeFuncBodyAndErrors(t *testing.T) {
	good := watFunc{locals: []valType{vtI32, vtI32}, body: []watInstr{{op: "i32.const", imm: 7}}}
	body, err := encodeFuncBody(&good)
	if err != nil {
		t.Fatalf("encodeFuncBody: %v", err)
	}
	checkEncodedBytes(t, "body", body, []byte{0x01, 0x02, 0x7f, 0x41, 0x07, 0x0b})
	bad := watFunc{body: []watInstr{{op: "bogus"}}}
	if _, err := encodeFuncBody(&bad); err == nil {
		t.Fatal("encodeFuncBody accepted an unsupported instruction")
	}
}
