/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package gobuiltin

import (
	"crypto/sha256"
	"fmt"
	"strings"

	"abate/lang/air"
	"abate/lang/gofront"
	"abate/lang/rule"
)

// LoadBuiltins registers every Go specification into a rule registry.
func LoadBuiltins(registry *rule.Registry) error {
	if registry == nil {
		return ErrNilRegistry
	}
	return registry.Load(All())
}

// ParseUnit parses one in-memory Go source buffer into an AIR unit.
//
// The frontend is constructed per call with semantic annotation disabled: the
// structural rules only need the syntax tree, and skipping the in-memory
// go/types pass keeps a whole-repository self-check fast and deterministic.
func ParseUnit(path string, source []byte) (*air.Unit, error) {
	unit := air.SourceUnit{
		Language: "go",
		Files:    []air.SourceFile{air.NewSourceFile(path, source)},
	}
	built, diagnostics, err := gofront.NewGoFrontend().
		WithTypeCheck(false).Parse(unit)
	if err != nil {
		return nil, fmt.Errorf("%w: %s: %w", ErrBadSource, path, err)
	}
	for _, diagnostic := range diagnostics {
		if diagnostic.IsError() {
			return nil, fmt.Errorf("%w: %s: %s", ErrBadSource,
				path, diagnostic.String())
		}
	}
	return built, nil
}

// CheckSpec compiles one specification and checks it against an AIR unit.
//
// This is the shared AIR-facing adapter used by the rule fixtures and by the
// self-check: it never exposes go/ast or go/token types to a caller.
func CheckSpec(spec rule.Spec, unit *air.Unit) ([]rule.Finding, error) {
	compiled, err := rule.Compile(spec)
	if err != nil {
		return nil, err
	}
	return compiled.Check(unit), nil
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
