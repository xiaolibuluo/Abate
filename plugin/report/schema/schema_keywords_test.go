/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package schema

import (
	"testing"
)

func TestTypeKeyword(t *testing.T) {
	cases := []struct {
		typ   string
		value any
	}{
		{"string", "s"},
		{"integer", float64(3)},
		{"number", float64(1.5)},
		{"boolean", true},
		{"object", map[string]any{}},
		{"array", []any{}},
		{"null", nil},
	}
	for _, c := range cases {
		s := &Schema{typ: c.typ}
		if err := s.Validate(c.value); err != nil {
			t.Fatalf("type %q rejected %v: %v", c.typ, c.value, err)
		}
		if err := s.Validate("wrong"); err == nil && c.typ != "string" {
			t.Fatalf("type %q accepted a string", c.typ)
		}
	}
	if err := (&Schema{typ: "integer"}).Validate(float64(1.5)); err == nil {
		t.Fatal("integer accepted a fractional value")
	}
	if err := (&Schema{typ: "string"}).Validate(float64(1)); err == nil {
		t.Fatal("string accepted a number")
	}
	if err := (&Schema{typ: "number"}).Validate("1"); err == nil {
		t.Fatal("number accepted a string")
	}
	if err := (&Schema{typ: "boolean"}).Validate(float64(1)); err == nil {
		t.Fatal("boolean accepted a number")
	}
	if err := (&Schema{typ: "object"}).Validate([]any{}); err == nil {
		t.Fatal("object accepted an array")
	}
	if err := (&Schema{typ: "array"}).Validate(map[string]any{}); err == nil {
		t.Fatal("array accepted an object")
	}
	if err := (&Schema{typ: "null"}).Validate("x"); err == nil {
		t.Fatal("null accepted a string")
	}
}

func TestEnumKeyword(t *testing.T) {
	s := &Schema{enum: []any{"a", float64(2), true}}
	for _, okValue := range []any{"a", float64(2), true} {
		if err := s.Validate(okValue); err != nil {
			t.Fatalf("enum rejected %v: %v", okValue, err)
		}
	}
	for _, badValue := range []any{"b", float64(3), false} {
		if err := s.Validate(badValue); err == nil {
			t.Fatalf("enum accepted %v", badValue)
		}
	}
}

func TestPatternKeyword(t *testing.T) {
	s := &Schema{typ: "string", pattern: "^a+$"}
	if err := s.Validate("aaa"); err != nil {
		t.Fatalf("pattern rejected a match: %v", err)
	}
	if err := s.Validate("bbb"); err == nil {
		t.Fatal("pattern accepted a non-match")
	}
	if err := s.Validate(float64(1)); err == nil {
		t.Fatal("typed schema accepted a number")
	}
	nonString := &Schema{pattern: "^a+$"}
	if err := nonString.Validate(float64(1)); err != nil {
		t.Fatalf("pattern must be skipped for non-strings: %v", err)
	}
	broken := &Schema{typ: "string", pattern: "("}
	if err := broken.Validate("a"); err == nil {
		t.Fatal("invalid regexp was accepted")
	}
	if _, err := matchPattern("^a+$", "aaa"); err != nil {
		t.Fatalf("matchPattern cache lookup: %v", err)
	}
}

func TestNumericBounds(t *testing.T) {
	low, high := float64(1), float64(9)
	s := &Schema{minimum: &low, maximum: &high}
	if err := s.Validate(float64(5)); err != nil {
		t.Fatalf("in-range value rejected: %v", err)
	}
	if err := s.Validate(float64(0)); err == nil {
		t.Fatal("value below minimum accepted")
	}
	if err := s.Validate(float64(10)); err == nil {
		t.Fatal("value above maximum accepted")
	}
	if err := s.Validate("not-a-number"); err != nil {
		t.Fatalf("numeric bounds must be skipped for non-numbers: %v", err)
	}
}

func TestFormatKeyword(t *testing.T) {
	cases := []struct {
		format string
		value  string
		valid  bool
	}{
		{"uri", "https://example.com/schema.json", true},
		{"uri", "://broken", false},
		{"uri-reference", "results/1", true},
		{"uri-reference", "", false},
		{"date-time", "2026-01-02T03:04:05Z", true},
		{"date-time", "2026-01-02T03:04:05.123456789Z", true},
		{"date-time", "not-a-time", false},
		{"email", "anything", true},
	}
	for _, c := range cases {
		s := &Schema{format: c.format}
		err := s.Validate(c.value)
		if c.valid && err != nil {
			t.Fatalf("format %q rejected %q: %v", c.format, c.value, err)
		}
		if !c.valid && err == nil {
			t.Fatalf("format %q accepted %q", c.format, c.value)
		}
	}
	if err := (&Schema{format: "uri"}).Validate(float64(1)); err != nil {
		t.Fatalf("format must be skipped for non-strings: %v", err)
	}
}

func TestCombinerKeywords(t *testing.T) {
	object := &Schema{typ: "object"}
	array := &Schema{typ: "array"}
	allPass := &Schema{allOf: []*Schema{object, {typ: "object"}}}
	if err := allPass.Validate(map[string]any{}); err != nil {
		t.Fatalf("allOf rejected a matching value: %v", err)
	}
	allFail := &Schema{allOf: []*Schema{object, array}}
	if err := allFail.Validate(map[string]any{}); err == nil {
		t.Fatal("allOf accepted a value failing one subschema")
	}
	anyPass := &Schema{anyOf: []*Schema{array, object}}
	if err := anyPass.Validate(map[string]any{}); err != nil {
		t.Fatalf("anyOf rejected a matching value: %v", err)
	}
	anyFail := &Schema{anyOf: []*Schema{array, {typ: "string"}}}
	if err := anyFail.Validate(map[string]any{}); err == nil {
		t.Fatal("anyOf accepted a value matching no subschema")
	}
	onePass := &Schema{oneOf: []*Schema{array, object}}
	if err := onePass.Validate(map[string]any{}); err != nil {
		t.Fatalf("oneOf rejected a unique match: %v", err)
	}
	oneNone := &Schema{oneOf: []*Schema{array, {typ: "string"}}}
	if err := oneNone.Validate(map[string]any{}); err == nil {
		t.Fatal("oneOf accepted zero matches")
	}
	oneMany := &Schema{oneOf: []*Schema{object, {typ: "object"}}}
	if err := oneMany.Validate(map[string]any{}); err == nil {
		t.Fatal("oneOf accepted two matches")
	}
}

func TestObjectAndItemsKeywords(t *testing.T) {
	s := &Schema{
		typ:                  "object",
		required:             []string{"name"},
		properties:           map[string]*Schema{"name": {typ: "string"}},
		additionalProperties: false,
	}
	if err := s.Validate(map[string]any{"name": "ok"}); err != nil {
		t.Fatalf("valid object rejected: %v", err)
	}
	if err := s.Validate(map[string]any{}); err == nil {
		t.Fatal("object missing a required property was accepted")
	}
	if err := s.Validate(map[string]any{"name": float64(1)}); err == nil {
		t.Fatal("property subschema violation was accepted")
	}
	if err := s.Validate(map[string]any{"name": "ok", "extra": true}); err == nil {
		t.Fatal("additionalProperties:false violation was accepted")
	}

	withExtraSchema := &Schema{
		typ:                  "object",
		properties:           map[string]*Schema{"name": {typ: "string"}},
		additionalProperties: &Schema{typ: "string"},
	}
	if err := withExtraSchema.Validate(map[string]any{"name": "ok", "extra": "text"}); err != nil {
		t.Fatalf("additionalProperties subschema rejected a string: %v", err)
	}
	if err := withExtraSchema.Validate(map[string]any{"name": "ok", "extra": float64(1)}); err == nil {
		t.Fatal("additionalProperties subschema accepted a number")
	}

	items := &Schema{typ: "array", items: &Schema{typ: "string"}}
	if err := items.Validate([]any{"a", "b"}); err != nil {
		t.Fatalf("valid array rejected: %v", err)
	}
	if err := items.Validate([]any{"a", float64(1)}); err == nil {
		t.Fatal("items subschema violation was accepted")
	}
}
