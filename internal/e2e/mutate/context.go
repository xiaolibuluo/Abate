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

// errInvalidContext reports an operator context that cannot drive a
// deterministic scan.
var errInvalidContext = errors.New("mutate: invalid operator context")

// OpContext is the complete configuration handed to one operator. The seed and
// manifest digest only choose mutation points; they never choose which file is
// processed. All source bytes are supplied in memory.
type OpContext struct {
	Content        []byte
	Language       string
	ManifestSHA256 string
	Seed           uint64
	Offset         int
}

// Validate checks that the context carries a non-empty buffer, a registered
// language tag and a bounded offset.
func (c OpContext) Validate() error {
	if len(c.Content) == 0 || c.Language == "" || c.Offset < 0 ||
		c.Offset > len(c.Content) {
		return fmt.Errorf("%w: language=%q bytes=%d offset=%d",
			errInvalidContext, c.Language, len(c.Content), c.Offset)
	}
	return nil
}
