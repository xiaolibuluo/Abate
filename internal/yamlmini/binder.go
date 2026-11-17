/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package yamlmini

import (
	"encoding"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// A bindPolicy carries the strictness decision for one decode operation.
// Using a policy value keeps helper signatures free of boolean control flags.
type bindPolicy struct {
	strict bool
}

func bind(n *Node, v any, strict bool) error {
	if v == nil {
		return errors.New("yamlmini: decode target must be a non-nil pointer, got nil")
	}
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Pointer || rv.IsNil() {
		return fmt.Errorf("yamlmini: decode target must be a non-nil pointer, got %T", v)
	}
	return bindPolicy{strict: strict}.bindValue(n, rv.Elem(), "$")
}

var textUnmarshalerType = reflect.TypeOf((*encoding.TextUnmarshaler)(nil)).Elem()

func (p bindPolicy) bindValue(n *Node, out reflect.Value, path string) error {
	if u, ok := textUnmarshalerOf(out); ok {
		return bindTextValue(n, u, path)
	}
	return p.bindByKind(n, out, path)
}

// textUnmarshalerOf returns the text codec of out when it implements one.
// The two-value type assertion makes the conversion explicit and panic-free.
func textUnmarshalerOf(out reflect.Value) (encoding.TextUnmarshaler, bool) {
	if !out.CanAddr() || !out.Addr().Type().Implements(textUnmarshalerType) {
		return nil, false
	}
	u, ok := out.Addr().Interface().(encoding.TextUnmarshaler)
	return u, ok
}

func bindTextValue(n *Node, u encoding.TextUnmarshaler, path string) error {
	text, err := nodeText(n, path)
	if err != nil {
		return err
	}
	if err := u.UnmarshalText([]byte(text)); err != nil {
		return fmt.Errorf("yamlmini: %s: %w", path, err)
	}
	return nil
}

func (p bindPolicy) bindByKind(n *Node, out reflect.Value, path string) error {
	switch out.Kind() {
	case reflect.Pointer:
		return p.bindPointer(n, out, path)
	case reflect.Struct:
		return p.bindStruct(n, out, path)
	case reflect.Map:
		return p.bindMap(n, out, path)
	case reflect.Slice:
		return p.bindSlice(n, out, path)
	case reflect.Array:
		return p.bindArray(n, out, path)
	case reflect.Interface:
		return bindInterface(n, out, path)
	case reflect.String, reflect.Bool,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Float32, reflect.Float64:
		return bindScalar(n, out, path)
	}
	return fmt.Errorf("yamlmini: %s: unsupported destination type %s", path, out.Type())
}

func bindScalar(n *Node, out reflect.Value, path string) error {
	switch out.Kind() {
	case reflect.String:
		text, err := nodeText(n, path)
		if err != nil {
			return err
		}
		out.SetString(text)
		return nil
	case reflect.Bool:
		if n.Sk != sBool {
			return scalarTypeError(path, out.Type(), n)
		}
		out.SetBool(n.Value == "true")
		return nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return bindSigned(n, out, path)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return bindUnsigned(n, out, path)
	case reflect.Float32, reflect.Float64:
		return bindFloat(n, out, path)
	}
	return fmt.Errorf("yamlmini: %s: unsupported destination type %s", path, out.Type())
}

func bindSigned(n *Node, out reflect.Value, path string) error {
	if n.Sk != sInt {
		return scalarTypeError(path, out.Type(), n)
	}
	v, err := strconv.ParseInt(n.Value, 10, out.Type().Bits())
	if err != nil {
		return fmt.Errorf("yamlmini: %s: %w", path, err)
	}
	out.SetInt(v)
	return nil
}

func bindUnsigned(n *Node, out reflect.Value, path string) error {
	if n.Sk != sInt || strings.HasPrefix(n.Value, "-") {
		return scalarTypeError(path, out.Type(), n)
	}
	v, err := strconv.ParseUint(n.Value, 10, out.Type().Bits())
	if err != nil {
		return fmt.Errorf("yamlmini: %s: %w", path, err)
	}
	out.SetUint(v)
	return nil
}

func bindFloat(n *Node, out reflect.Value, path string) error {
	if n.Sk != sInt && n.Sk != sFloat {
		return scalarTypeError(path, out.Type(), n)
	}
	v, err := strconv.ParseFloat(n.Value, out.Type().Bits())
	if err != nil {
		return fmt.Errorf("yamlmini: %s: %w", path, err)
	}
	out.SetFloat(v)
	return nil
}
