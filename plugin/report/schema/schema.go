/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package schema is a self-researched, offline JSON-Schema (draft-07 subset)
// validator. It implements only the keywords the OASIS SARIF 2.1.0 schema
// actually uses: type, properties, required, enum, items, additionalProperties,
// $ref, oneOf/anyOf/allOf, pattern, minimum/maximum and format. It is a small
// dependency-free replacement for the external Microsoft SARIF validator, used to
// prove RPT-SARIF output conforms to the official OASIS schema. The validator is
// intentionally a subset: unknown keywords are ignored, which keeps it safe to
// run against the full SARIF schema without pulling in a third-party dependency.
package schema

import "fmt"

// Schema is a compiled JSON-Schema node. All fields are unexported: callers
// interact with a Schema only through LoadSchema and Validate, never by
// inspecting its internals.
type Schema struct {
	typ                  string
	properties           map[string]*Schema
	required             []string
	enum                 []any
	items                *Schema
	additionalProperties any
	ref                  string
	oneOf                []*Schema
	anyOf                []*Schema
	allOf                []*Schema
	pattern              string
	minimum              *float64
	maximum              *float64
	format               string
	definitions          map[string]*Schema
}

// LoadSchema reads a JSON-Schema document from path and compiles it into a
// Schema ready for validation. The document's definitions are retained so that
// $ref pointers of the form "#/definitions/Name" resolve.
func LoadSchema(path string) (*Schema, error) {
	data, err := readFile(path)
	if err != nil {
		return nil, err
	}
	var raw any
	if err := unmarshal(data, &raw); err != nil {
		return nil, fmt.Errorf("schema: cannot parse %q: %w", path, err)
	}
	m, ok := raw.(map[string]any)
	if !ok {
		return nil, fmt.Errorf("schema: root of %q is not a JSON object", path)
	}
	root := &Schema{}
	parseSchema(m, root)
	return root, nil
}

// lookupRef resolves a "#/definitions/Name" reference against the root schema.
func (r *Schema) lookupRef(ref string) *Schema {
	const prefix = "#/definitions/"
	if len(ref) <= len(prefix) || ref[:len(prefix)] != prefix {
		return nil
	}
	name := ref[len(prefix):]
	return r.definitions[name]
}

// parseSchema populates target from a decoded JSON node. Only the keywords this
// validator understands are compiled; others are silently ignored so the validator
// stays a safe subset of draft-07. The root node's definitions are attached to the
// root target, which is the object all $ref pointers resolve against.
func parseSchema(node any, target *Schema) {
	m, ok := node.(map[string]any)
	if !ok {
		return
	}
	if ref, ok := m["$ref"].(string); ok {
		target.ref = ref
	}
	if t, ok := m["type"].(string); ok {
		target.typ = t
	}
	if p, ok := m["pattern"].(string); ok {
		target.pattern = p
	}
	if f, ok := m["format"].(string); ok {
		target.format = f
	}
	if enum, ok := m["enum"].([]any); ok {
		target.enum = enum
	}
	if req, ok := m["required"].([]any); ok {
		for _, r := range req {
			if rs, ok := r.(string); ok {
				target.required = append(target.required, rs)
			}
		}
	}
	if props, ok := m["properties"].(map[string]any); ok {
		target.properties = make(map[string]*Schema, len(props))
		for k, v := range props {
			child := &Schema{}
			parseSchema(v, child)
			target.properties[k] = child
		}
	}
	if items, ok := m["items"]; ok {
		target.items = &Schema{}
		parseSchema(items, target.items)
	}
	if ap, ok := m["additionalProperties"]; ok {
		switch apv := ap.(type) {
		case bool:
			target.additionalProperties = apv
		case map[string]any:
			sub := &Schema{}
			parseSchema(apv, sub)
			target.additionalProperties = sub
		}
	}
	if defs, ok := m["definitions"].(map[string]any); ok {
		target.definitions = make(map[string]*Schema, len(defs))
		for k, v := range defs {
			child := &Schema{}
			parseSchema(v, child)
			target.definitions[k] = child
		}
	}
	collectCombos(m, "oneOf", &target.oneOf)
	collectCombos(m, "anyOf", &target.anyOf)
	collectCombos(m, "allOf", &target.allOf)
	if mn, ok := m["minimum"].(float64); ok {
		v := mn
		target.minimum = &v
	}
	if mx, ok := m["maximum"].(float64); ok {
		v := mx
		target.maximum = &v
	}
}

// collectCombos parses a combiner keyword (oneOf/anyOf/allOf) into subschemas.
func collectCombos(m map[string]any, key string, dst *[]*Schema) {
	if arr, ok := m[key].([]any); ok {
		for _, e := range arr {
			child := &Schema{}
			parseSchema(e, child)
			*dst = append(*dst, child)
		}
	}
}
