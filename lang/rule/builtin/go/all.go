/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package gobuiltin

import (
	"slices"
	"strings"

	"abate/lang/rule"
)

// All returns every declared Go specification in ascending identifier order.
//
// The declaration is pure data: calling All rebuilds the slice from the ten
// table-driven source files, so no mutable package state can leak between
// invocations or between concurrent readers.
func All() []rule.Spec {
	groups := [][]rule.Spec{
		goEG1Rules(), goEG2Rules(), goEG3Rules(),
		goCRC1Rules(), goCRC2Rules(), goCRC3Rules(),
		goVET1Rules(), goVET2Rules(), goVET3Rules(), goVET4Rules(),
	}
	out := make([]rule.Spec, 0, 128)
	for _, group := range groups {
		out = append(out, group...)
	}
	slices.SortFunc(out, func(a, b rule.Spec) int {
		return strings.Compare(string(a.ID), string(b.ID))
	})
	return out
}

// Count returns the number of declared Go specifications.
func Count() int {
	return len(All())
}
