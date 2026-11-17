/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package schema

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// oasisSchemaPath is the bundled OASIS SARIF 2.1.0 standard schema, resolved relative
// to the schema package directory (Go runs tests from the package directory). It is
// public standard fact data under internal/standards/sarif, not proprietary open-source
// implementation source; the validator here is 100% self-developed.
const oasisSchemaPath = "../../../internal/standards/sarif/sarif-2.1.0.schema.json"

// sarifDocument is a minimal but complete OASIS-valid log. It mirrors the shape
// emitted by plugin/report without importing that package (which would create a
// test import cycle, since report imports schema).
const sarifDocument = `{
  "$schema": "https://docs.oasis-open.org/sarif/sarif/v2.1.0/cos02/schemas/sarif-schema-2.1.0.json",
  "version": "2.1.0",
  "runs": [{
    "tool": {"driver": {
      "name": "Abate RPT",
      "version": "0.1.0",
      "informationUri": "https://github.com/xiaolibuluo/Abate",
      "rules": [{"id": "GO-EG-001", "name": "GO-EG-001",
                 "shortDescription": {"text": "Abate static rule GO-EG-001"}}]
    }},
    "results": [{
      "ruleId": "GO-EG-001", "level": "error", "kind": "fail",
      "message": {"text": "error not checked"},
      "locations": [{"physicalLocation": {
        "artifactLocation": {"uri": "a/b.go"},
        "region": {"startLine": 12, "startColumn": 4}
      }}]
    }]
  }]
}`

// decodeDocument decodes a JSON document into the generic tree the validator
// expects (map[string]any / []any / float64 / string / bool / nil).
func decodeDocument(t *testing.T, text string) any {
	t.Helper()
	var value any
	if err := json.Unmarshal([]byte(text), &value); err != nil {
		t.Fatalf("decode document: %v", err)
	}
	return value
}

func TestLoadOASISSchemaAndValidateDocument(t *testing.T) {
	root, err := LoadSchema(oasisSchemaPath)
	if err != nil {
		t.Fatalf("LoadSchema: %v", err)
	}
	if root.typ != "object" {
		t.Fatalf("root type = %q", root.typ)
	}
	if len(root.definitions) == 0 {
		t.Fatal("OASIS definitions were not compiled")
	}
	if err := root.Validate(decodeDocument(t, sarifDocument)); err != nil {
		t.Fatalf("OASIS SARIF document rejected: %v", err)
	}
	if err := root.Validate(decodeDocument(t, `"not-an-object"`)); err == nil {
		t.Fatal("validator accepted a non-object root")
	}
}

func TestLoadSchemaFailureModes(t *testing.T) {
	dir := t.TempDir()
	missing := filepath.Join(dir, "absent.json")
	if _, err := LoadSchema(missing); err == nil {
		t.Fatal("LoadSchema accepted a missing file")
	}
	malformed := filepath.Join(dir, "malformed.json")
	if err := os.WriteFile(malformed, []byte("{broken"), 0o600); err != nil {
		t.Fatalf("write malformed schema: %v", err)
	}
	if _, err := LoadSchema(malformed); err == nil {
		t.Fatal("LoadSchema accepted malformed JSON")
	}
	arrayRoot := filepath.Join(dir, "array.json")
	if err := os.WriteFile(arrayRoot, []byte("[]"), 0o600); err != nil {
		t.Fatalf("write array schema: %v", err)
	}
	if _, err := LoadSchema(arrayRoot); err == nil {
		t.Fatal("LoadSchema accepted a non-object root")
	}
}

func TestRefResolution(t *testing.T) {
	root := &Schema{
		definitions: map[string]*Schema{"thing": {typ: "string"}},
	}
	if root.lookupRef("#/definitions/thing") == nil {
		t.Fatal("valid definition reference did not resolve")
	}
	if root.lookupRef("#/definitions/") != nil {
		t.Fatal("empty definition name resolved")
	}
	if root.lookupRef("#/other/thing") != nil {
		t.Fatal("unknown reference prefix resolved")
	}
	if root.lookupRef("thing") != nil {
		t.Fatal("bare name resolved as a reference")
	}
	resolved := &Schema{ref: "#/definitions/thing", definitions: root.definitions}
	if err := resolved.Validate("text"); err != nil {
		t.Fatalf("resolved $ref rejected a string: %v", err)
	}
	if err := resolved.Validate(float64(1)); err == nil {
		t.Fatal("resolved $ref accepted a number")
	}
	unresolved := &Schema{ref: "#/definitions/absent", definitions: root.definitions}
	if err := unresolved.Validate("text"); err == nil {
		t.Fatal("unresolved $ref was accepted")
	}
}

func TestParseSchemaCompilesEveryKeyword(t *testing.T) {
	raw := `{
      "$ref": "#/definitions/thing",
      "type": "object",
      "pattern": "^a+$",
      "format": "uri",
      "enum": ["x", 1, true],
      "required": ["a", 5],
      "properties": {"a": {"type": "string"}},
      "items": {"type": "number"},
      "additionalProperties": {"type": "boolean"},
      "definitions": {"thing": {"type": "boolean"}},
      "oneOf": [{"type": "object"}],
      "anyOf": [{"type": "object"}],
      "allOf": [{"type": "object"}],
      "minimum": 1,
      "maximum": 9
    }`
	var node any
	if err := json.Unmarshal([]byte(raw), &node); err != nil {
		t.Fatalf("decode schema: %v", err)
	}
	s := &Schema{}
	parseSchema(node, s)
	if s.ref != "#/definitions/thing" || s.typ != "object" {
		t.Fatalf("ref/type = %q/%q", s.ref, s.typ)
	}
	if s.pattern != "^a+$" || s.format != "uri" || len(s.enum) != 3 {
		t.Fatalf("scalar keywords = %q/%q/%v", s.pattern, s.format, s.enum)
	}
	if len(s.required) != 1 || s.required[0] != "a" {
		t.Fatalf("required = %v", s.required)
	}
	if s.properties["a"] == nil || s.properties["a"].typ != "string" {
		t.Fatal("properties were not compiled")
	}
	if s.items == nil || s.items.typ != "number" {
		t.Fatal("items were not compiled")
	}
	if _, ok := s.additionalProperties.(*Schema); !ok {
		t.Fatal("additionalProperties subschema was not compiled")
	}
	if s.definitions["thing"] == nil || len(s.oneOf) != 1 || len(s.anyOf) != 1 || len(s.allOf) != 1 {
		t.Fatal("definitions or combiners were not compiled")
	}
	if s.minimum == nil || *s.minimum != 1 || s.maximum == nil || *s.maximum != 9 {
		t.Fatal("numeric bounds were not compiled")
	}
}

func TestBooleanAdditionalPropertiesAndNonObjectNode(t *testing.T) {
	trueSchema := &Schema{}
	parseSchema(map[string]any{"additionalProperties": true}, trueSchema)
	if b, ok := trueSchema.additionalProperties.(bool); !ok || !b {
		t.Fatal("boolean additionalProperties was not compiled")
	}
	if err := trueSchema.Validate(map[string]any{"extra": float64(1)}); err != nil {
		t.Fatalf("additionalProperties:true rejected an extra property: %v", err)
	}
	ignored := &Schema{}
	parseSchema("not-an-object", ignored)
	if ignored.typ != "" || ignored.properties != nil {
		t.Fatal("parseSchema mutated the target for a non-object node")
	}
	parseSchema(map[string]any{"type": float64(3)}, ignored)
	if ignored.typ != "" {
		t.Fatal("parseSchema accepted a non-string type")
	}
}

func TestHelperDeterminism(t *testing.T) {
	keys := sortedKeys(map[string]any{"c": nil, "a": nil, "b": nil})
	if strings.Join(keys, "") != "abc" {
		t.Fatalf("sortedKeys = %v", keys)
	}
	schemaKeys := sortedSchemaKeys(map[string]*Schema{"z": {}, "y": {}, "x": {}})
	if strings.Join(schemaKeys, "") != "xyz" {
		t.Fatalf("sortedSchemaKeys = %v", schemaKeys)
	}
	if got := summarize(strings.Repeat("a", 40)); len(got) != 35 {
		t.Fatalf("summarize(long) = %q", got)
	}
	if summarize(nil) != "null" || summarize(float64(3)) != "3" || summarize("s") != "s" {
		t.Fatal("summarize scalar branches changed")
	}
	if enumEqual(nil, nil) {
		t.Fatal("enumEqual must not equate two null values")
	}
	if !enumEqual("a", "a") || !enumEqual(float64(1), float64(1)) || !enumEqual(true, true) {
		t.Fatal("enumEqual scalar branches changed")
	}
	if enumEqual("a", float64(1)) || enumEqual(true, "a") {
		t.Fatal("enumEqual equated values of different types")
	}
	if err := checkType("tuple", float64(1)); err == nil {
		t.Fatal("checkType accepted an unknown type")
	}
	if err := checkType("null", nil); err != nil {
		t.Fatalf("checkType(null,nil): %v", err)
	}
}
