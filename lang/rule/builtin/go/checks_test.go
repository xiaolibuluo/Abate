/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package gobuiltin

import (
	"errors"
	"testing"

	"abate/lang/rule"
)

// TestAdapterContracts covers the failure paths of the shared AIR adapter.
func TestAdapterContracts(t *testing.T) {
	t.Parallel()
	if err := LoadBuiltins(nil); !errors.Is(err, ErrNilRegistry) {
		t.Fatalf("LoadBuiltins(nil) = %v, want ErrNilRegistry", err)
	}
	if _, err := ParseUnit("", []byte("package p\n")); !errors.Is(err, ErrBadSource) {
		t.Fatalf("ParseUnit with an empty path = %v, want ErrBadSource", err)
	}
	if _, err := ParseUnit("broken.go", []byte("package p\nfunc F( {\n")); !errors.Is(err, ErrBadSource) {
		t.Fatalf("ParseUnit with broken syntax = %v, want ErrBadSource", err)
	}
	unit, err := ParseUnit("clean.go", []byte("package p\nvar foo_bar int\n"))
	if err != nil {
		t.Fatalf("ParseUnit clean: %v", err)
	}
	spec := rule.Spec{}
	for _, candidate := range All() {
		if candidate.ID == "GO-EG-001" {
			spec = candidate
			break
		}
	}
	findings, err := CheckSpec(spec, unit)
	if err != nil {
		t.Fatalf("CheckSpec: %v", err)
	}
	if len(findings) == 0 {
		t.Fatal("CheckSpec found no violation in its own positive fragment")
	}
	bad := rule.Spec{
		ID: "GO-EG-999", Reference: "Effective Go",
		Severity: rule.SeverityWarning, Category: rule.CategoryStyle,
		Pattern: "(", Message: "x",
	}
	if _, err := CheckSpec(bad, unit); !errors.Is(err, rule.ErrBadPattern) {
		t.Fatalf("CheckSpec with a bad pattern = %v, want ErrBadPattern", err)
	}
	if _, err := CheckSpec(spec, nil); err != nil {
		t.Fatalf("CheckSpec with a nil unit = %v, want no error", err)
	}
}
