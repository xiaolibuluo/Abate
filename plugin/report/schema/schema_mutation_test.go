/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package schema

import (
	"strings"
	"testing"
)

// TestKeywordMutationsAreRejected mutates one keyword at a time and proves the
// validator detects each mutation. A vacuous validator would pass all cases.
func TestKeywordMutationsAreRejected(t *testing.T) {
	low, high := float64(1), float64(9)
	object := map[string]any{"name": "ok"}
	stringValue := "aaa"
	cases := []struct {
		name   string
		schema *Schema
		value  any
	}{
		{"type", &Schema{typ: "array"}, object},
		{"required", &Schema{required: []string{"absent"}}, object},
		{"enum", &Schema{enum: []any{"other"}}, "aaa"},
		{"pattern", &Schema{pattern: "^z+$"}, stringValue},
		{"minimum", &Schema{minimum: &high}, float64(2)},
		{"maximum", &Schema{maximum: &low}, float64(8)},
		{"additionalProperties", &Schema{properties: map[string]*Schema{}, additionalProperties: false}, object},
		{"items", &Schema{items: &Schema{typ: "boolean"}}, []any{"a"}},
		{"oneOf", &Schema{oneOf: []*Schema{{typ: "array"}, {typ: "string"}}}, object},
		{"anyOf", &Schema{anyOf: []*Schema{{typ: "array"}, {typ: "string"}}}, object},
		{"allOf", &Schema{allOf: []*Schema{{typ: "object"}, {typ: "array"}}}, object},
		{"ref", &Schema{ref: "#/definitions/absent"}, object},
		{"format", &Schema{format: "date-time"}, "not-a-time"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if err := c.schema.Validate(c.value); err == nil {
				t.Fatalf("mutation %q was not detected", c.name)
			}
		})
	}
}

// TestOASISDocumentMutationsAreRejected mutates the bundled SARIF document and
// proves the compiled official schema rejects each change.
func TestOASISDocumentMutationsAreRejected(t *testing.T) {
	root, err := LoadSchema(oasisSchemaPath)
	if err != nil {
		t.Fatalf("LoadSchema: %v", err)
	}
	if err := root.Validate(decodeDocument(t, sarifDocument)); err != nil {
		t.Fatalf("baseline document rejected: %v", err)
	}
	mutations := []string{
		strings.Replace(sarifDocument, `"version": "2.1.0"`, `"version": "9.9.9"`, 1),
		strings.Replace(sarifDocument, `"version": "2.1.0",`, ``, 1),
		strings.Replace(sarifDocument, `"version": "2.1.0",`, `"version": "2.1.0", "unexpectedRoot": true,`, 1),
		strings.Replace(sarifDocument, `"level": "error"`, `"level": "not-a-level"`, 1),
		strings.Replace(sarifDocument, `"startLine": 12`, `"startLine": -1`, 1),
	}
	for i, mutated := range mutations {
		t.Run(string(rune('a'+i)), func(t *testing.T) {
			if err := root.Validate(decodeDocument(t, mutated)); err == nil {
				t.Fatalf("document mutation %d was not detected", i)
			}
		})
	}
}
