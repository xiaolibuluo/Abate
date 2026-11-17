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
	if text, ok, err := marshaledText(rv); ok {
		if err != nil {
			return nil, fmt.Errorf("yamlmini: %s: MarshalText: %w", path, err)
		}
		return scalarNode(string(text), sString, 0), nil
	}
	return emitByKind(path, rv)
}

// marshaledText reports whether rv carries a text codec and returns its text.
// The two-value assertions keep the interface conversion panic-free.
func marshaledText(rv reflect.Value) ([]byte, bool, error) {
	// TextMarshaler first: VirtualDuration emits "1s" rather than an integer.
	if rv.CanInterface() && rv.Type().Implements(textMarshalerType) {
		if m, ok := rv.Interface().(encoding.TextMarshaler); ok {
			text, err := m.MarshalText()
			return text, true, err
		}
	}
	if rv.CanAddr() && rv.Addr().Type().Implements(textMarshalerType) {
		if m, ok := rv.Addr().Interface().(encoding.TextMarshaler); ok {
			text, err := m.MarshalText()
			return text, true, err
		}
	}
	return nil, false, nil
}

func emitByKind(path string, rv reflect.Value) (*Node, error) {
	switch rv.Kind() {
	case reflect.Interface, reflect.Pointer:
		if rv.IsNil() {
			return nullNode(0), nil
		}
		return emitReflect(path, rv.Elem())
	case reflect.Struct:
		return emitStruct(path, rv)
	case reflect.Map:
		return emitMap(path, rv)
	case reflect.Slice, reflect.Array:
		return emitSlice(path, rv)
	default:
		return emitPrimitive(path, rv)
	}
}

func emitStruct(path string, rv reflect.Value) (*Node, error) {
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
}

func emitMap(path string, rv reflect.Value) (*Node, error) {
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
}

func emitSlice(path string, rv reflect.Value) (*Node, error) {
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
}

func emitPrimitive(path string, rv reflect.Value) (*Node, error) {
	switch rv.Kind() {
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
