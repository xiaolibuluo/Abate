/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package csbuiltin

import (
	"crypto/sha256"
	"slices"
	"strings"

	"abate/lang/rule"
	"abate/lang/rule/builtin/common"
)

// All returns every declared C# specification in ascending identifier order.
//
// The declaration is pure data: calling All rebuilds the slice from the shared
// template set plus the twelve C#-specific rules, so no mutable package state
// can leak between invocations or between concurrent readers.
func All() []rule.Spec {
	out := common.Generate("CS", nil)
	out = append(out, csRules()...)
	slices.SortFunc(out, func(a, b rule.Spec) int {
		return strings.Compare(string(a.ID), string(b.ID))
	})
	return out
}

// Count returns the number of declared C# specifications.
func Count() int {
	return len(All())
}

// LoadBuiltins registers every C# specification into a rule registry.
func LoadBuiltins(registry *rule.Registry) error {
	if registry == nil {
		return common.ErrNilRegistry
	}
	return registry.Load(All())
}

// Fingerprint returns the SHA-256 digest of the fixed-order rule set.
func Fingerprint() [32]byte {
	var builder strings.Builder
	for _, spec := range All() {
		_, _ = builder.WriteString(spec.Canonical())
		_ = builder.WriteByte(0x1e)
	}
	return sha256.Sum256([]byte(builder.String()))
}
