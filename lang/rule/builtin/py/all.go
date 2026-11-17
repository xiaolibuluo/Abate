/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package pybuiltin

import (
	"crypto/sha256"
	"slices"
	"strings"

	"abate/lang/rule"
	"abate/lang/rule/builtin/common"
)

// All returns every declared Python specification in ascending identifier
// order. It composes the sixty shared AIR templates (namespace PY, ids
// PY-AIR-NNN) with the twelve Python-specific rules declared by pyRules.
func All() []rule.Spec {
	out := common.Generate("PY", nil)
	out = append(out, pyRules()...)
	slices.SortFunc(out, func(a, b rule.Spec) int {
		return strings.Compare(string(a.ID), string(b.ID))
	})
	return out
}

// Count returns the number of declared Python specifications. It is the sum of
// the sixty common templates and the twelve local rules, so it is at least 60.
func Count() int {
	return len(All())
}

// LoadBuiltins registers the whole Python rule set into a registry.
//
// A nil registry is rejected with common.ErrNilRegistry before any rule is
// touched, mirroring the contract of the shared common package.
func LoadBuiltins(registry *rule.Registry) error {
	if registry == nil {
		return common.ErrNilRegistry
	}
	return registry.Load(All())
}

// Fingerprint returns the SHA-256 digest of the fixed-order rule set so two
// builds of the same rules produce a byte-identical gate key.
func Fingerprint() [32]byte {
	var builder strings.Builder
	for _, spec := range All() {
		_, _ = builder.WriteString(spec.Canonical())
		_ = builder.WriteByte(0x1e)
	}
	return sha256.Sum256([]byte(builder.String()))
}
