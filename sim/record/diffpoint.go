/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package record

import "fmt"

// DiffPoint locates one divergence between two archives.
//
// Path is the complete structural path such as "Frames[2].Payload[7]", while
// Seq, Kind, Want and Got carry the sequence, category, expected and actual
// values required by the deterministic acceptance report. ByteOffset is the
// payload index for byte-level divergences and -1 for scalar fields.
type DiffPoint struct {
	// Path is the complete structural path of the divergent field.
	Path string

	// FrameIndex is the zero-based index of the divergent frame, or -1 for a header field.
	FrameIndex int

	// Seq is the total-order sequence of the expected frame.
	Seq uint64

	// Kind is the category of the expected frame.
	Kind FrameKind

	// Field names the compared field.
	Field string

	// Want is the expected value rendered deterministically.
	Want string

	// Got is the actual value rendered deterministically.
	Got string

	// ByteOffset is the payload index of the divergence, or -1 for scalar fields.
	ByteOffset int
}

// String renders the point as one deterministic line.
func (p DiffPoint) String() string {
	return fmt.Sprintf("path=%s frame=%d seq=%d kind=%s field=%s want=%s got=%s byteOffset=%d",
		p.Path, p.FrameIndex, p.Seq, p.Kind, p.Field, p.Want, p.Got, p.ByteOffset)
}
