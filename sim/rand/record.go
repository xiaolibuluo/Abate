/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package rand

// DrawRecord is a single random draw captured for replay accounting.
//
// It records the draw index, the consuming method kind, the integer arguments
// of that method and the raw 64-bit result, so SIM-REC can prove the exact
// random stream an experiment consumed from its seed.
type DrawRecord struct {
	// Index is the zero-based draw ordinal within the owning stream.
	Index int

	// Kind labels the consuming method (uint64, int63, float64, intn, bool,
	// bytes or fork).
	Kind string

	// Args carries the integer arguments of the consuming method, if any.
	Args []uint64

	// Result is the raw 64-bit output of the generator for this draw.
	Result uint64
}
