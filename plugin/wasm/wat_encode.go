/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package wasm

// wasmMagic is the four-byte WebAssembly binary header. It is retained here as
// pure data for the offline WAT encoder; loading and execution are not part of
// this package (see the package doc and plugin/host).
var wasmMagic = []byte{0x00, 0x61, 0x73, 0x6d}

// watToWasm translates the supported WAT subset into WebAssembly binary bytes.
//
// The translation is deterministic and side-effect free: identical WAT input
// always yields byte-identical modules. Unsupported constructs return an error
// rather than emitting invalid bytecode.
func watToWasm(src []byte) ([]byte, error) {
	p := newWatParser(src)
	mod, err := p.parseModule()
	if err != nil {
		return nil, err
	}
	return encodeModule(mod)
}

// encodeModule assembles the binary sections in canonical order.
func encodeModule(mod *watModule) ([]byte, error) {
	var types []funcType
	typeIdx := func(ft funcType) int {
		for i, t := range types {
			if sameTypes(t.params, ft.params) && sameTypes(t.results, ft.results) {
				return i
			}
		}
		types = append(types, ft)
		return len(types) - 1
	}
	importIdx := make([]int, len(mod.imports))
	for i := range mod.imports {
		importIdx[i] = typeIdx(funcType{params: mod.imports[i].params, results: mod.imports[i].results})
	}
	funcIdx := make([]int, len(mod.funcs))
	for i := range mod.funcs {
		funcIdx[i] = typeIdx(funcType{params: mod.funcs[i].params, results: mod.funcs[i].results})
	}

	var out []byte
	out = append(out, wasmMagic...)
	out = append(out, 0x01, 0x00, 0x00, 0x00)
	out = appendSection(out, 1, encodeTypes(types))
	out = appendSection(out, 2, encodeImports(mod, importIdx))
	out = appendSection(out, 3, encodeFunctions(funcIdx))
	if mem := encodeMemory(mod); mem != nil {
		out = appendSection(out, 5, mem)
	}
	out = appendSection(out, 7, encodeExports(mod))
	code, err := encodeCode(mod)
	if err != nil {
		return nil, err
	}
	out = appendSection(out, 10, code)
	return out, nil
}

// appendSection writes a section id, its LEB length and its content.
func appendSection(b []byte, id byte, content []byte) []byte {
	b = append(b, id)
	b = appendULEN(b, len(content))
	return append(b, content...)
}

// AppendULEN appends an unsigned LEB128 integer from a non-negative int length.
//
// Gosec G115 is suppressed because n is always a slice length or a validated
// Index, which is non-negative by construction.
func appendULEN(b []byte, n int) []byte {
	return appendULED(b, uint64(n)) //nolint:gosec // n is a length or validated index
}

// encodeTypes writes the type section body.
func encodeTypes(types []funcType) []byte {
	var b []byte
	b = appendULEN(b, len(types))
	for i := range types {
		b = append(b, 0x60)
		b = appendValTypeVec(b, types[i].params)
		b = appendValTypeVec(b, types[i].results)
	}
	return b
}

// encodeImports writes the import section body.
func encodeImports(mod *watModule, typeIdxs []int) []byte {
	count := len(mod.imports)
	if mod.memImport != nil {
		count++
	}
	var b []byte
	b = appendULEN(b, count)
	for i := range mod.imports {
		b = appendName(b, mod.imports[i].module)
		b = appendName(b, mod.imports[i].name)
		b = append(b, 0x00)
		b = appendULEN(b, typeIdxs[i])
	}
	if mod.memImport != nil {
		b = appendName(b, mod.memImport.module)
		b = appendName(b, mod.memImport.name)
		b = append(b, 0x02)
		b = appendLimits(b, mod.memMin, mod.memMax)
	}
	return b
}

// encodeFunctions writes the function section body.
func encodeFunctions(typeIdxs []int) []byte {
	var b []byte
	b = appendULEN(b, len(typeIdxs))
	for _, idx := range typeIdxs {
		b = appendULEN(b, idx)
	}
	return b
}

// encodeMemory writes the memory section body, or nil when memory is imported.
func encodeMemory(mod *watModule) []byte {
	if mod.memImport != nil {
		return nil
	}
	var b []byte
	b = appendULEN(b, 1)
	b = appendLimits(b, mod.memMin, mod.memMax)
	return b
}

// encodeExports writes the export section body.
func encodeExports(mod *watModule) []byte {
	count := 0
	for i := range mod.funcs {
		if mod.funcs[i].exportName != "" {
			count++
		}
	}
	if mod.memExport != "" {
		count++
	}
	var b []byte
	b = appendULEN(b, count)
	for i := range mod.funcs {
		if mod.funcs[i].exportName == "" {
			continue
		}
		b = appendName(b, mod.funcs[i].exportName)
		b = append(b, 0x00)
		b = appendULEN(b, len(mod.imports)+i)
	}
	if mod.memExport != "" {
		b = appendName(b, mod.memExport)
		b = append(b, 0x02)
		b = appendULED(b, 0)
	}
	return b
}

// encodeCode writes the code section body.
func encodeCode(mod *watModule) ([]byte, error) {
	var b []byte
	b = appendULEN(b, len(mod.funcs))
	for i := range mod.funcs {
		body, err := encodeFuncBody(&mod.funcs[i])
		if err != nil {
			return nil, err
		}
		b = appendULEN(b, len(body))
		b = append(b, body...)
	}
	return b, nil
}

// blockType encodes the block result as a WebAssembly block type byte.
func blockType(t valType) byte {
	if t == 0 {
		return 0x40
	}
	return byte(t)
}

// appendValTypeVec encodes a vector of value types.
func appendValTypeVec(b []byte, ts []valType) []byte {
	b = appendULEN(b, len(ts))
	for i := range ts {
		b = append(b, byte(ts[i]))
	}
	return b
}

// appendName encodes a UTF-8 name as a length-prefixed byte vector.
func appendName(b []byte, s string) []byte {
	b = appendULEN(b, len(s))
	return append(b, []byte(s)...)
}

// appendLimits encodes a memory/page limits structure.
func appendLimits(b []byte, lo, hi uint32) []byte {
	if hi == 0 {
		b = append(b, 0x00)
		return appendULED(b, uint64(lo))
	}
	b = append(b, 0x01)
	b = appendULED(b, uint64(lo))
	return appendULED(b, uint64(hi))
}

// appendMemArg encodes a memory immediate (align then offset).
func appendMemArg(b []byte, instr *watInstr) []byte {
	b = appendULED(b, uint64(instr.memAlgn))
	b = appendULED(b, uint64(instr.memOff))
	return b
}

// sameTypes reports whether two value-type vectors are element-wise equal.
func sameTypes(a, b []valType) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// appendULED appends an unsigned LEB128-encoded integer.
func appendULED(b []byte, v uint64) []byte {
	for {
		c := byte(v & 0x7f)
		v >>= 7
		if v != 0 {
			c |= 0x80
		}
		b = append(b, c)
		if v == 0 {
			return b
		}
	}
}

// appendILED appends a signed LEB128-encoded integer.
func appendILED(b []byte, value int64) []byte {
	more := true
	for more {
		c := byte(value & 0x7f)
		value >>= 7
		if (value == 0 && c&0x40 == 0) || (value == -1 && c&0x40 != 0) {
			more = false
		} else {
			c |= 0x80
		}
		b = append(b, c)
	}
	return b
}
