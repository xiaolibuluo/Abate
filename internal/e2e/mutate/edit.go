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
)

// errMalformedEdit reports an edit whose interval is outside the source buffer.
var errMalformedEdit = errors.New("mutate: malformed edit")

// Edit is one in-memory source replacement produced by an operator. Offsets
// always refer to the original buffer, never to an earlier edit's result.
type Edit struct {
	OperatorID string
	Before     string
	After      string
	Offset     int
	EndOffset  int
}

// Validate checks the half-open interval against a source of srcLen bytes and
// requires the operator and original text to be recorded.
func (e Edit) Validate(srcLen int) error {
	if e.OperatorID == "" || e.Before == "" || e.Before == e.After ||
		e.Offset < 0 || e.EndOffset <= e.Offset || e.EndOffset > srcLen {
		return fmt.Errorf("%w: [%d,%d) over %d bytes", errMalformedEdit,
			e.Offset, e.EndOffset, srcLen)
	}
	return nil
}

// Matches reports whether the original buffer still carries Before inside the
// edit interval, so a stale edit can never be applied.
func (e Edit) Matches(content []byte) bool {
	if e.Offset < 0 || e.EndOffset > len(content) || e.EndOffset <= e.Offset {
		return false
	}
	return string(content[e.Offset:e.EndOffset]) == e.Before
}
