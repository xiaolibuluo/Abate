/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package schema

import (
	"fmt"
	"math"
	"regexp"
	"sort"
	"sync"
)

// patternCache memoizes compiled regular expressions so repeated validation of a
// large document (for example a full SARIF log) does not recompile the same
// pattern on every node. A sync.Map keeps concurrent validation race-free
// without a process-wide mutable map guarded by hand.
var patternCache sync.Map

// Validate checks value against the schema, returning a descriptive error on the
// first violation. The value is a decoded JSON document (map[string]any for
// objects, []any for arrays, and string/float64/bool/nil for scalars).
func (s *Schema) Validate(value any) error {
	return s.validateNode(value, "$", s)
}

// validateNode walks value against the compiled schema. The root argument
// carries the definitions used to resolve $ref pointers.
func (s *Schema) validateNode(value any, path string, root *Schema) error {
	if s.ref != "" {
		target := root.lookupRef(s.ref)
		if target == nil {
			return fmt.Errorf("%s: unresolved $ref %q", path, s.ref)
		}
		return target.validateNode(value, path, root)
	}
	if err := s.validateScalar(value, path); err != nil {
		return err
	}
	if err := s.validateCombinators(value, path, root); err != nil {
		return err
	}
	return s.validateContainer(value, path, root)
}

// validateScalar enforces the type, enum, pattern, range and format constraints
// that apply to a single scalar value.
func (s *Schema) validateScalar(value any, path string) error {
	if s.typ != "" {
		if err := checkType(s.typ, value); err != nil {
			return fmt.Errorf("%s: %w", path, err)
		}
	}
	if len(s.enum) > 0 && !enumContains(s.enum, value) {
		return fmt.Errorf("%s: value %v is not one of the allowed enum values", path, summarize(value))
	}
	if s.pattern != "" {
		if str, ok := value.(string); ok {
			matched, err := matchPattern(s.pattern, str)
			if err != nil {
				return fmt.Errorf("%s: invalid pattern %q: %w", path, s.pattern, err)
			}
			if !matched {
				return fmt.Errorf("%s: value %q does not match pattern %q", path, str, s.pattern)
			}
		}
	}
	if s.minimum != nil {
		if f, ok := value.(float64); ok && f < *s.minimum {
			return fmt.Errorf("%s: value %v is below minimum %v", path, f, *s.minimum)
		}
	}
	if s.maximum != nil {
		if f, ok := value.(float64); ok && f > *s.maximum {
			return fmt.Errorf("%s: value %v is above maximum %v", path, f, *s.maximum)
		}
	}
	if s.format != "" {
		if str, ok := value.(string); ok {
			if err := checkFormat(s.format, str); err != nil {
				return fmt.Errorf("%s: %w", path, err)
			}
		}
	}
	return nil
}

// validateCombinators enforces allOf, anyOf and oneOf subschemas.
func (s *Schema) validateCombinators(value any, path string, root *Schema) error {
	for _, sub := range s.allOf {
		if err := sub.validateNode(value, path, root); err != nil {
			return err
		}
	}
	if len(s.anyOf) > 0 {
		matched := false
		for _, sub := range s.anyOf {
			if sub.validateNode(value, path, root) == nil {
				matched = true
				break
			}
		}
		if !matched {
			return fmt.Errorf("%s: value does not match any anyOf subschema", path)
		}
	}
	if len(s.oneOf) > 0 {
		matches := 0
		for _, sub := range s.oneOf {
			if sub.validateNode(value, path, root) == nil {
				matches++
			}
		}
		if matches != 1 {
			return fmt.Errorf("%s: oneOf requires exactly one match, got %d", path, matches)
		}
	}
	return nil
}

// validateContainer enforces object properties and array items.
func (s *Schema) validateContainer(value any, path string, root *Schema) error {
	if obj, ok := value.(map[string]any); ok {
		if err := s.validateObject(obj, path, root); err != nil {
			return err
		}
	}
	if arr, ok := value.([]any); ok {
		if s.items != nil {
			for i, item := range arr {
				if err := s.items.validateNode(item, fmt.Sprintf("%s[%d]", path, i), root); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// sortedKeys returns the keys of m in lexical order. Every map walk in the
// validator uses this helper so error messages and traversal order never leak
// Go's randomized map iteration order.
func sortedKeys(m map[string]any) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// sortedSchemaKeys returns the keys of m in lexical order.
func sortedSchemaKeys(m map[string]*Schema) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// validateObject enforces required properties, additionalProperties and property
// subschemas for an object value.
func (s *Schema) validateObject(obj map[string]any, path string, root *Schema) error {
	for _, r := range s.required {
		if _, present := obj[r]; !present {
			return fmt.Errorf("%s: missing required property %q", path, r)
		}
	}
	objKeys := sortedKeys(obj)
	if s.additionalProperties != nil {
		if b, ok := s.additionalProperties.(bool); ok && !b {
			for _, k := range objKeys {
				if _, known := s.properties[k]; !known {
					return fmt.Errorf("%s: additional property %q is not allowed", path, k)
				}
			}
		} else if sub, ok := s.additionalProperties.(*Schema); ok {
			for _, k := range objKeys {
				if _, known := s.properties[k]; !known {
					if err := sub.validateNode(obj[k], path+"."+k, root); err != nil {
						return err
					}
				}
			}
		}
	}
	for _, name := range sortedSchemaKeys(s.properties) {
		if val, present := obj[name]; present {
			if err := s.properties[name].validateNode(val, path+"."+name, root); err != nil {
				return err
			}
		}
	}
	return nil
}

// checkType verifies that value matches the JSON-Schema type name. Integers are
// validated as integral float64 values, the representation produced by the
// standard library JSON decoder.
func checkType(t string, value any) error {
	switch t {
	case "string":
		if _, ok := value.(string); !ok {
			return fmt.Errorf("expected string, got %T", value)
		}
	case "integer":
		f, ok := value.(float64)
		if !ok {
			return fmt.Errorf("expected integer, got %T", value)
		}
		if f != math.Trunc(f) {
			return fmt.Errorf("expected integer, got %v", f)
		}
	case "number":
		if _, ok := value.(float64); !ok {
			return fmt.Errorf("expected number, got %T", value)
		}
	case "boolean":
		if _, ok := value.(bool); !ok {
			return fmt.Errorf("expected boolean, got %T", value)
		}
	case "object":
		if _, ok := value.(map[string]any); !ok {
			return fmt.Errorf("expected object, got %T", value)
		}
	case "array":
		if _, ok := value.([]any); !ok {
			return fmt.Errorf("expected array, got %T", value)
		}
	case "null":
		if value != nil {
			return fmt.Errorf("expected null, got %T", value)
		}
	default:
		return fmt.Errorf("unknown schema type %q", t)
	}
	return nil
}

// enumContains reports whether value equals one of the enum entries.
func enumContains(enum []any, value any) bool {
	for _, e := range enum {
		if enumEqual(e, value) {
			return true
		}
	}
	return false
}

// enumEqual compares two decoded JSON scalars for equality.
func enumEqual(a, b any) bool {
	switch av := a.(type) {
	case string:
		bv, ok := b.(string)
		return ok && av == bv
	case float64:
		bv, ok := b.(float64)
		return ok && av == bv
	case bool:
		bv, ok := b.(bool)
		return ok && av == bv
	}
	return false
}

// matchPattern compiles and matches a regular expression against str. Compilation
// results are cached by pattern text.
func matchPattern(pattern, str string) (bool, error) {
	if cached, ok := patternCache.Load(pattern); ok {
		return cached.(*regexp.Regexp).MatchString(str), nil
	}
	compiled, err := regexp.Compile(pattern)
	if err != nil {
		return false, err
	}
	actual, _ := patternCache.LoadOrStore(pattern, compiled)
	return actual.(*regexp.Regexp).MatchString(str), nil
}
