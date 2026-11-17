/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package yamlmini

import (
	"encoding"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

var textMarshalerType = reflect.TypeOf((*encoding.TextMarshaler)(nil)).Elem()

// emitValue converts an arbitrary Go value into the ordered node tree. Struct
// fields follow declaration order; map keys are sorted; TextMarshaler values
// are emitted as their text scalar.
func emitValue(path string, v any) (*Node, error) {
	if v == nil {
		return nullNode(0), nil
	}
	return emitReflect(path, reflect.ValueOf(v))
}

func emitReflect(path string, rv reflect.Value) (*Node, error) {
	if !rv.IsValid() {
		return nullNode(0), nil
	}
	// TextMarshaler first: VirtualDuration emits "1s" rather than an integer.
	if rv.CanInterface() && rv.Type().Implements(textMarshalerType) {
		text, err := rv.Interface().(encoding.TextMarshaler).MarshalText()
		if err != nil {
			return nil, fmt.Errorf("yamlmini: %s: MarshalText: %w", path, err)
		}
		return scalarNode(string(text), sString, 0), nil
	}
	if rv.CanAddr() && rv.Addr().Type().Implements(textMarshalerType) {
		text, err := rv.Addr().Interface().(encoding.TextMarshaler).MarshalText()
		if err != nil {
			return nil, fmt.Errorf("yamlmini: %s: MarshalText: %w", path, err)
		}
		return scalarNode(string(text), sString, 0), nil
	}

	switch rv.Kind() {
	case reflect.Interface:
		if rv.IsNil() {
			return nullNode(0), nil
		}
		return emitReflect(path, rv.Elem())

	case reflect.Pointer:
		if rv.IsNil() {
			return nullNode(0), nil
		}
		return emitReflect(path, rv.Elem())

	case reflect.Struct:
		m := &Node{Kind: Mapping}
		t := rv.Type()
		for _, f := range fieldsOf(t) {
			fv := rv.Field(f.index)
			if hasOmitEmpty(t.Field(f.index)) && isZero(fv) {
				continue
			}
			child, err := emitReflect(path+"."+f.name, fv)
			if err != nil {
				return nil, err
			}
			m.Keys = append(m.Keys, f.name)
			m.Vals = append(m.Vals, child)
		}
		return m, nil

	case reflect.Map:
		if rv.IsNil() {
			return nullNode(0), nil
		}
		if rv.Type().Key().Kind() != reflect.String {
			return nil, fmt.Errorf("yamlmini: %s: map key must have string kind, got %s", path, rv.Type().Key())
		}
		keys := make([]string, 0, rv.Len())
		byKey := make(map[string]reflect.Value, rv.Len())
		iter := rv.MapRange()
		for iter.Next() {
			k := iter.Key().String()
			if _, dup := byKey[k]; dup {
				return nil, fmt.Errorf("yamlmini: %s: duplicate map key %q", path, k)
			}
			keys = append(keys, k)
			byKey[k] = iter.Value()
		}
		sort.Strings(keys)
		m := &Node{Kind: Mapping}
		for _, k := range keys {
			child, err := emitReflect(path+"."+k, byKey[k])
			if err != nil {
				return nil, err
			}
			m.Keys = append(m.Keys, k)
			m.Vals = append(m.Vals, child)
		}
		return m, nil

	case reflect.Slice, reflect.Array:
		if rv.Kind() == reflect.Slice && rv.IsNil() {
			return nullNode(0), nil
		}
		s := &Node{Kind: Sequence}
		for i := 0; i < rv.Len(); i++ {
			child, err := emitReflect(fmt.Sprintf("%s[%d]", path, i), rv.Index(i))
			if err != nil {
				return nil, err
			}
			s.Items = append(s.Items, child)
		}
		return s, nil

	case reflect.String:
		return scalarNode(rv.String(), sString, 0), nil

	case reflect.Bool:
		return scalarNode(strconv.FormatBool(rv.Bool()), sBool, 0), nil

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return scalarNode(strconv.FormatInt(rv.Int(), 10), sInt, 0), nil

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return scalarNode(strconv.FormatUint(rv.Uint(), 10), sInt, 0), nil

	case reflect.Float32, reflect.Float64:
		bits := 64
		if rv.Kind() == reflect.Float32 {
			bits = 32
		}
		return scalarNode(strconv.FormatFloat(rv.Float(), 'g', -1, bits), sFloat, 0), nil
	}
	return nil, fmt.Errorf("yamlmini: %s: cannot encode type %s", path, rv.Type())
}

func hasOmitEmpty(f reflect.StructField) bool {
	for _, part := range strings.Split(f.Tag.Get("yaml"), ",")[1:] {
		if part == "omitempty" {
			return true
		}
	}
	return false
}

func isZero(rv reflect.Value) bool {
	if !rv.IsValid() {
		return true
	}
	switch rv.Kind() {
	case reflect.Pointer, reflect.Interface, reflect.Map, reflect.Slice:
		return rv.IsNil()
	case reflect.Array:
		for i := 0; i < rv.Len(); i++ {
			if !isZero(rv.Index(i)) {
				return false
			}
		}
		return true
	case reflect.Struct:
		// Treat a struct as empty only when every exported field is empty; this
		// keeps "rollback: {}" out of documents when the whole struct is zero.
		for i := 0; i < rv.NumField(); i++ {
			if rv.Type().Field(i).IsExported() && !isZero(rv.Field(i)) {
				return false
			}
		}
		return true
	default:
		return rv.IsZero()
	}
}

// emitNode writes a node tree with a fixed two-space indentation. Scalars and
// empty collections are written inline; non-empty collections are expanded.
func emitNode(b *strings.Builder, n *Node, indent int) {
	pad := strings.Repeat(" ", indent)
	switch n.Kind {
	case Mapping:
		for i, key := range n.Keys {
			v := n.Vals[i]
			b.WriteString(pad)
			b.WriteString(emitKey(key))
			b.WriteByte(':')
			switch {
			case isInline(v):
				b.WriteByte(' ')
				b.WriteString(emitScalar(v))
				b.WriteByte('\n')
			default:
				b.WriteByte('\n')
				emitNode(b, v, indent+2)
			}
		}
	case Sequence:
		for _, item := range n.Items {
			b.WriteString(pad)
			b.WriteByte('-')
			switch {
			case isInline(item):
				b.WriteByte(' ')
				b.WriteString(emitScalar(item))
				b.WriteByte('\n')
			default:
				b.WriteByte('\n')
				emitNode(b, item, indent+2)
			}
		}
	case Scalar, Null:
		b.WriteString(pad)
		b.WriteString(emitScalar(n))
		b.WriteByte('\n')
	}
}

func isInline(n *Node) bool {
	if n == nil {
		return true
	}
	if n.Kind == Scalar || n.Kind == Null {
		return true
	}
	if n.Kind == Sequence && len(n.Items) == 0 {
		return true
	}
	return n.Kind == Mapping && len(n.Keys) == 0
}

func emitScalar(n *Node) string {
	if n == nil || n.Kind == Null {
		return "null"
	}
	if n.Kind != Scalar {
		if n.Kind == Sequence {
			return "[]"
		}
		return "{}"
	}
	if n.Sk == sString && (!plainSafe(n.Value) || looksNonString(n.Value)) {
		return quoteString(n.Value)
	}
	return n.Value
}

func emitKey(key string) string {
	if plainSafe(key) && !looksNonString(key) {
		return key
	}
	return quoteString(key)
}

func quoteString(s string) string {
	var b strings.Builder
	b.WriteByte('"')
	for _, r := range s {
		switch r {
		case '"':
			b.WriteString(`\"`)
		case '\\':
			b.WriteString(`\\`)
		case '\n':
			b.WriteString(`\n`)
		case '\r':
			b.WriteString(`\r`)
		case '\t':
			b.WriteString(`\t`)
		case '\b':
			b.WriteString(`\b`)
		case '\f':
			b.WriteString(`\f`)
		default:
			if r < 0x20 || r == 0x7f {
				fmt.Fprintf(&b, `\u%04X`, r)
			} else {
				b.WriteRune(r)
			}
		}
	}
	b.WriteByte('"')
	return b.String()
}

// plainSafe reports whether s can be emitted as an unquoted plain scalar that
// will round-trip back to the same string. Anything ambiguous is quoted.
func plainSafe(s string) bool {
	if s == "" || strings.TrimSpace(s) != s {
		return false
	}
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c < 0x20 || c == 0x7f {
			return false
		}
	}
	if strings.ContainsAny(s[:1], "-?:,[]{}#&*!|>'\"%@`~") {
		return false
	}
	if strings.HasPrefix(s, "- ") || strings.HasPrefix(s, "? ") || strings.HasPrefix(s, ": ") {
		return false
	}
	if strings.Contains(s, ": ") || strings.Contains(s, " #") || strings.HasSuffix(s, ":") {
		return false
	}
	return true
}

// looksNonString reports whether a plain scalar would be re-lexed as a
// non-string value, in which case it must be quoted to preserve its type.
func looksNonString(s string) bool {
	switch s {
	case "null", "Null", "NULL", "~", "true", "True", "TRUE", "false", "False", "FALSE":
		return true
	}
	if integerLiteral(s) {
		return true
	}
	if floatLiteral(s) {
		if _, err := strconv.ParseFloat(s, 64); err == nil {
			return true
		}
	}
	return false
}
