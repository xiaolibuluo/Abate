/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package rsbuiltin

import (
	"crypto/sha256"
	"errors"
	"slices"
	"strings"

	"abate/lang/rule"
	"abate/lang/rule/builtin/common"
)

// ErrNilRegistry reports a nil registry handed to LoadBuiltins.
var ErrNilRegistry = errors.New("rsbuiltin: nil rule registry")

// All returns every declared Rust specification in ascending identifier order.
//
// The declaration is pure data: it merges the sixty shared AIR templates
// (namespace "RS") with the twelve language-local rules and sorts the result, so
// no mutable package state can leak between calls or concurrent readers.
func All() []rule.Spec {
	out := common.Generate("RS", nil)
	out = append(out, rsRules()...)
	slices.SortFunc(out, func(a, b rule.Spec) int {
		return strings.Compare(string(a.ID), string(b.ID))
	})
	return out
}

// Count returns the number of declared Rust specifications.
func Count() int {
	return len(All())
}

// LoadBuiltins registers every Rust specification into a rule registry.
func LoadBuiltins(registry *rule.Registry) error {
	if registry == nil {
		return ErrNilRegistry
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
