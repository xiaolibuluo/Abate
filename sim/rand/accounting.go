/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package rand

// Ledger chunk sizing. The first chunk is small so a stream that records only a
// handful of draws (the normal SIM-REC experiment) stays cheap, and every later
// chunk is large so a long experiment needs only a few allocations. Chunked
// storage removes slice grow-copies entirely: recording a draw never
// reallocates or copies records already captured.
const (
	drawChunkMin = 64
	drawChunkMax = 1024
)

// nextChunkSize returns the capacity of the chunk at ordinal n.
func nextChunkSize(n int) int {
	if n == 0 {
		return drawChunkMin
	}
	return drawChunkMax
}

// Internal kind codes identify the consuming method in the pointer-free ledger.
// They are stable: drawKindName restores the exact Kind string that encodeDraw
// and the determinism certification consume, so recorded RNG frames never shift.
const (
	drawKindUint64 uint8 = iota
	drawKindInt63
	drawKindFloat64
	drawKindIntN
	drawKindBool
	drawKindBytes
	drawKindFork
)

// drawKindName resolves an internal kind code to the exact Kind string constant
// carried by DrawRecord. Every string is a compile-time constant, so Snapshot
// performs no string allocation for the kind field.
func drawKindName(code uint8) string {
	switch code {
	case drawKindUint64:
		return "uint64"
	case drawKindInt63:
		return "int63"
	case drawKindFloat64:
		return "float64"
	case drawKindIntN:
		return "intn"
	case drawKindBool:
		return "bool"
	case drawKindBytes:
		return "bytes"
	case drawKindFork:
		return "fork"
	default:
		return ""
	}
}

// drawEntry is the pointer-free hot-path representation of one recorded draw.
//
// It carries no Go pointers, so the runtime GC never scans a ledger chunk; that
// scan was the dominant cost of the previous []DrawRecord design. The draw index
// is derived from the chunk offset instead of being stored, and the single
// integer argument is held inline so a draw never allocates an argument slice.
type drawEntry struct {
	result uint64
	arg    uint64
	kind   uint8
	hasArg uint8
}

// Accounting records every random draw of a stream for deterministic replay.
//
// The snapshot is consumed by SIM-REC to certify that an experiment's random
// consumption is reproducible from its seed. Records are held in fixed,
// pointer-free chunks so the hot record path never grows, copies or makes the
// GC scan the ledger, while a long experiment can still accumulate unbounded
// draws. Snapshot expands entries into the public DrawRecord view unchanged.
type Accounting struct {
	chunks [][]drawEntry
	active []drawEntry
	used   int
	length int
}

// newAccounting returns an empty accounting ledger.
func newAccounting() *Accounting {
	return &Accounting{}
}

// record appends an argument-free draw to the ledger with its consuming method.
//
// The common case writes into the active chunk without any slice growth; a new
// chunk is allocated only when the active one fills. Snapshot turns the zero
// hasArg field back into a nil Args slice exactly as the original variadic
// no-argument call did.
func (a *Accounting) record(code uint8, result uint64) {
	if a.used == len(a.active) {
		a.active = make([]drawEntry, nextChunkSize(len(a.chunks)))
		a.chunks = append(a.chunks, a.active)
		a.used = 0
	}
	e := &a.active[a.used]
	e.result = result
	e.kind = code
	a.used++
	a.length++
}

// recordArg appends a draw that carries a single integer argument to the ledger.
// The argument is stored inline and Snapshot expands it into a one-element Args
// slice with the same value as the original variadic argument call.
func (a *Accounting) recordArg(code uint8, arg, result uint64) {
	if a.used == len(a.active) {
		a.active = make([]drawEntry, nextChunkSize(len(a.chunks)))
		a.chunks = append(a.chunks, a.active)
		a.used = 0
	}
	e := &a.active[a.used]
	e.result = result
	e.kind = code
	e.arg = arg
	e.hasArg = 1
	a.used++
	a.length++
}

// Len returns the number of recorded draws.
func (a *Accounting) Len() int {
	return a.length
}

// Snapshot returns a defensive copy of the draw ledger.
//
// It flattens the pointer-free chunks into one contiguous []DrawRecord whose
// element values are identical to a ledger grown as []DrawRecord, so SIM-REC
// encodes the exact same RNG frame bytes and the determinism hash is unchanged.
func (a *Accounting) Snapshot() []DrawRecord {
	out := make([]DrawRecord, a.length)
	off := 0
	for _, chunk := range a.chunks {
		n := len(chunk)
		if off+n > a.length {
			n = a.length - off
		}
		for j := 0; j < n; j++ {
			out[off+j] = expandDrawRecord(off+j, chunk[j])
		}
		off += n
	}
	return out
}

// expandDrawRecord rebuilds the public view of one pointer-free ledger entry.
// An argument-free draw yields a nil Args slice; an argument-bearing draw yields
// a one-element slice with the same value as the original variadic slice.
func expandDrawRecord(index int, e drawEntry) DrawRecord {
	d := DrawRecord{
		Index:  index,
		Kind:   drawKindName(e.kind),
		Result: e.result,
	}
	if e.hasArg != 0 {
		d.Args = []uint64{e.arg}
	}
	return d
}
