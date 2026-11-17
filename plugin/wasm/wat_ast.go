/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package wasm

import (
	"fmt"
	"math"
)

// valType is the WebAssembly value-type encoding used by the translator.
type valType byte

const (
	vtI32 valType = 0x7f
	vtI64 valType = 0x7e
)

// watMemImport records an imported linear memory.
type watMemImport struct {
	module string
	name   string
}

// watFuncImport records an imported host function signature.
type watFuncImport struct {
	module  string
	name    string
	params  []valType
	results []valType
}

// watFunc is a defined function with its body instructions.
type watFunc struct {
	exportName string
	params     []valType
	results    []valType
	locals     []valType
	body       []watInstr
}

// watBlock is a structured control-flow body (loop/block/if).
type watBlock struct {
	kind   string
	result valType
	then   []watInstr
	els    []watInstr
}

// watInstr is a single WebAssembly instruction in the parsed tree.
type watInstr struct {
	op      string
	imm     int64
	idx     uint64
	memOff  uint32
	memAlgn uint32
	block   *watBlock
}

// watModule is the parsed representation of a WAT module.
type watModule struct {
	memMin    uint32
	memMax    uint32
	memImport *watMemImport
	memExport string
	imports   []watFuncImport
	funcs     []watFunc
}

// funcType is a function signature used by the type section.
type funcType struct {
	params  []valType
	results []valType
}

// ToU32 narrows a parser value to uint32 after a non-negative, in-range check.
//
// Gosec G115 is suppressed on the widening because the value is validated
// immediately above: WASM memory limits and offsets are unsigned by grammar and
// The parser rejects any out-of-range literal rather than emitting it.
func toU32(v int64) (uint32, error) {
	if v < 0 || v > math.MaxUint32 {
		return 0, fmt.Errorf("wasm: value %d out of uint32 range", v)
	}
	return uint32(v), nil //nolint:gosec // range checked above
}

// ToU64 widens a parser value to uint64 after a non-negative check.
//
// Gosec G115 is suppressed because the value is validated above: call/branch
// Targets are unsigned by grammar and the parser rejects negative literals.
func toU64(v int64) (uint64, error) {
	if v < 0 {
		return 0, fmt.Errorf("wasm: negative index %d", v)
	}
	return uint64(v), nil //nolint:gosec // v >= 0 checked above
}
