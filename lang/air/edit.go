/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package air

import (
	"fmt"
	"slices"
)

// Edit replaces the half-open byte interval [Offset, EndOffset) of one source
// file with Text. Offsets always refer to the original buffer, not to the
// result of an earlier edit.
type Edit struct {
	// Path selects the source buffer to edit.
	Path string
	// Offset is the inclusive start byte offset.
	Offset int
	// EndOffset is the exclusive end byte offset.
	EndOffset int
	// Text is the replacement byte stream.
	Text []byte
}

// Validate checks the interval against a source buffer of srcLen bytes.
func (e Edit) Validate(srcLen int) error {
	if e.Offset < 0 || e.EndOffset < e.Offset || e.EndOffset > srcLen {
		return fmt.Errorf("%w: [%d,%d) exceeds %d bytes",
			ErrInvalidEdit, e.Offset, e.EndOffset, srcLen)
	}
	return nil
}

// ApplyEdits applies a set of non-overlapping edits to a source buffer.
//
// Edits are sorted by start offset before application, so the result is
// independent of the caller's slice order. Overlapping edits are rejected.
func ApplyEdits(src []byte, edits []Edit) ([]byte, error) {
	ordered := slices.Clone(edits)
	slices.SortStableFunc(ordered, func(a, b Edit) int {
		switch {
		case a.Offset < b.Offset:
			return -1
		case a.Offset > b.Offset:
			return 1
		default:
			return 0
		}
	})

	out := make([]byte, 0, len(src))
	previous := 0
	for _, edit := range ordered {
		if err := edit.Validate(len(src)); err != nil {
			return nil, err
		}
		if edit.Offset < previous {
			return nil, fmt.Errorf("%w: edits overlap at %d",
				ErrInvalidEdit, edit.Offset)
		}
		out = append(out, src[previous:edit.Offset]...)
		out = append(out, edit.Text...)
		previous = edit.EndOffset
	}
	out = append(out, src[previous:]...)
	return out, nil
}
