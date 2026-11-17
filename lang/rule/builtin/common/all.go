/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package common

import (
	"crypto/sha256"
	"errors"
	"strings"

	"abate/lang/rule"
)

// ErrNilRegistry reports a nil registry handed to LoadBuiltins.
var ErrNilRegistry = errors.New("common: nil rule registry")

// All returns the sixty common specifications in ascending identifier order.
//
// The set is pure data rebuilt on each call, so no mutable package state can
// leak between invocations or between concurrent readers.
func All() []rule.Spec {
	return Generate(DefaultTag, nil)
}

// Count returns the number of declared common specifications.
func Count() int {
	return len(All())
}

// LoadBuiltins registers the whole common template family into a registry.
//
// Loading goes through GenerateChecked so any schema, pattern or predicate
// failure is reported to the caller before the registry is touched.
func LoadBuiltins(registry *rule.Registry) error {
	if registry == nil {
		return ErrNilRegistry
	}
	specs, err := GenerateChecked(DefaultTag, nil)
	if err != nil {
		return err
	}
	return registry.Load(specs)
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
