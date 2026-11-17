/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package mutate

import (
	"errors"
	"fmt"
	"slices"
)

var (
	errStaleEdit       = errors.New("mutate: edit no longer matches source")
	errOverlappingEdit = errors.New("mutate: overlapping edits")
)

// Apply rewrites content in memory according to non-overlapping edits. It
// returns a fresh buffer and never writes to the read-only source tree.
func Apply(edits []Edit, content []byte) ([]byte, error) {
	ordered := slices.Clone(edits)
	slices.SortFunc(ordered, compareEdits)
	out := make([]byte, 0, len(content))
	previous := 0
	for _, edit := range ordered {
		if err := edit.Validate(len(content)); err != nil {
			return nil, err
		}
		if !edit.Matches(content) {
			return nil, fmt.Errorf("%w: %s", errStaleEdit, edit.OperatorID)
		}
		if edit.Offset < previous {
			return nil, fmt.Errorf("%w: at %d", errOverlappingEdit, edit.Offset)
		}
		out = append(out, content[previous:edit.Offset]...)
		out = append(out, edit.After...)
		previous = edit.EndOffset
	}
	out = append(out, content[previous:]...)
	return out, nil
}
