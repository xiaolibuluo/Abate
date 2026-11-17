/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package wasm

import "fmt"

// encodeFuncBody writes one function's locals declaration and instruction stream.
func encodeFuncBody(f *watFunc) ([]byte, error) {
	var body []byte
	body = appendLocals(body, f.locals)
	var err error
	body, err = encodeInstrs(f.body, body)
	if err != nil {
		return nil, err
	}
	body = append(body, 0x0b)
	return body, nil
}

// appendLocals writes the locals declaration, grouping equal consecutive types.
func appendLocals(b []byte, locals []valType) []byte {
	type group struct {
		n uint64
		t valType
	}
	var groups []group
	for _, t := range locals {
		if len(groups) > 0 && groups[len(groups)-1].t == t {
			groups[len(groups)-1].n++
			continue
		}
		groups = append(groups, group{n: 1, t: t})
	}
	b = appendULEN(b, len(groups))
	for _, g := range groups {
		b = appendULED(b, g.n)
		b = append(b, byte(g.t))
	}
	return b
}

// encodeInstrs encodes a sequence of instructions.
func encodeInstrs(instrs []watInstr, b []byte) ([]byte, error) {
	var err error
	for i := range instrs {
		b, err = encodeInstr(&instrs[i], b)
		if err != nil {
			return nil, err
		}
	}
	return b, nil
}

// encodeInstr encodes a single instruction and its immediates.
func encodeInstr(instr *watInstr, b []byte) ([]byte, error) {
	switch instr.op {
	case "loop", "block", "if", "else", "end", "br", "br_if", "return", "drop", "call", "unreachable", "nop":
		return encodeInstrControl(instr, b)
	default:
		return encodeInstrNumeric(instr, b)
	}
}

// encodeInstrControl emits control-flow and stack instructions.
func encodeInstrControl(instr *watInstr, b []byte) ([]byte, error) {
	switch instr.op {
	case "loop", "block":
		op := byte(0x03)
		if instr.op == "block" {
			op = 0x02
		}
		b = append(b, op, blockType(instr.block.result))
		var err error
		b, err = encodeInstrs(instr.block.then, b)
		if err != nil {
			return nil, err
		}
		return append(b, 0x0b), nil
	case "if":
		b = append(b, 0x04, blockType(instr.block.result))
		var err error
		b, err = encodeInstrs(instr.block.then, b)
		if err != nil {
			return nil, err
		}
		if len(instr.block.els) > 0 {
			b = append(b, 0x05)
			b, err = encodeInstrs(instr.block.els, b)
			if err != nil {
				return nil, err
			}
		}
		return append(b, 0x0b), nil
	case "else":
		return append(b, 0x05), nil
	case "end":
		return append(b, 0x0b), nil
	case "br":
		b = append(b, 0x0c)
		b = appendULED(b, instr.idx)
		return b, nil
	case "br_if":
		b = append(b, 0x0d)
		b = appendULED(b, instr.idx)
		return b, nil
	case "return":
		return append(b, 0x0f), nil
	case "drop":
		return append(b, 0x1a), nil
	case "call":
		b = append(b, 0x10)
		b = appendULED(b, instr.idx)
		return b, nil
	case "unreachable":
		return append(b, 0x00), nil
	case "nop":
		return append(b, 0x01), nil
	default:
		return nil, fmt.Errorf("wasm: unsupported control opcode %q", instr.op)
	}
}

// encodeInstrNumeric emits constant, comparison, arithmetic and memory ops.
func encodeInstrNumeric(instr *watInstr, b []byte) ([]byte, error) {
	switch instr.op {
	case "i32.const":
		b = append(b, 0x41)
		b = appendILED(b, instr.imm)
		return b, nil
	case "i64.const":
		b = append(b, 0x42)
		b = appendILED(b, instr.imm)
		return b, nil
	case "local.get":
		b = append(b, 0x20)
		b = appendULED(b, instr.idx)
		return b, nil
	case "local.set":
		b = append(b, 0x21)
		b = appendULED(b, instr.idx)
		return b, nil
	case "local.tee":
		b = append(b, 0x22)
		b = appendULED(b, instr.idx)
		return b, nil
	case "i32.eq":
		return append(b, 0x46), nil
	case "i32.ne":
		return append(b, 0x47), nil
	case "i32.lt_s":
		return append(b, 0x48), nil
	case "i32.gt_s":
		return append(b, 0x4a), nil
	case "i32.le_s":
		return append(b, 0x4c), nil
	case "i32.ge_s":
		return append(b, 0x4e), nil
	case "i32.add":
		return append(b, 0x6a), nil
	case "i32.sub":
		return append(b, 0x6b), nil
	case "i32.mul":
		return append(b, 0x6c), nil
	case "i32.and":
		return append(b, 0x71), nil
	case "i32.or":
		return append(b, 0x72), nil
	case "i32.xor":
		return append(b, 0x73), nil
	case "i32.shl":
		return append(b, 0x74), nil
	case "i32.shr_s":
		return append(b, 0x75), nil
	case "i32.clz":
		return append(b, 0x67), nil
	case "i32.ctz":
		return append(b, 0x68), nil
	case "i32.popcnt":
		return append(b, 0x69), nil
	case "i32.load", "i64.load", "i32.store", "i64.store":
		return encodeInstrMemory(instr, b)
	default:
		return nil, fmt.Errorf("wasm: unsupported opcode %q", instr.op)
	}
}

// encodeInstrMemory emits the four linear-memory load and store opcodes.
func encodeInstrMemory(instr *watInstr, b []byte) ([]byte, error) {
	switch instr.op {
	case "i32.load":
		b = append(b, 0x28)
	case "i64.load":
		b = append(b, 0x29)
	case "i32.store":
		b = append(b, 0x36)
	case "i64.store":
		b = append(b, 0x37)
	default:
		return nil, fmt.Errorf("wasm: unsupported memory opcode %q", instr.op)
	}
	b = appendMemArg(b, instr)
	return b, nil
}
