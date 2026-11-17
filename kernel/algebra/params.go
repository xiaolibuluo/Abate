/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package algebra

import (
	"slices"
	"strings"
)

// Params carries the operator parameters.
//
// Params participates in the deterministic fingerprint, so it must be
// canonicalized before hashing; see Params.Canonical.
type Params map[string]string

// Canonical returns the canonical serialization of the params.
//
// Keys are sorted so that map iteration order never leaks into the output.
func (params Params) Canonical() string {
	keys := make([]string, 0, len(params))

	for key := range params {
		keys = append(keys, key)
	}

	slices.Sort(keys)

	pairs := make([]string, 0, len(keys))

	for _, key := range keys {
		pairs = append(pairs, key+"="+params[key])
	}

	return strings.Join(pairs, separator)
}
