/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package yamlmini

import (
	"fmt"
	"reflect"
	"strings"
)

func (p bindPolicy) bindPointer(n *Node, out reflect.Value, path string) error {
	if n.Kind == Null {
		out.Set(reflect.Zero(out.Type()))
		return nil
	}
	if out.IsNil() {
		out.Set(reflect.New(out.Type().Elem()))
	}
	return p.bindValue(n, out.Elem(), path)
}

func (p bindPolicy) bindStruct(n *Node, out reflect.Value, path string) error {
	if n.Kind == Null {
		return nil
	}
	if n.Kind != Mapping {
		return shapeError(path, out.Type(), n.Kind)
	}
	fields := fieldsOf(out.Type())
	for i, key := range n.Keys {
		f, ok := fieldByName(fields, key)
		if !ok {
			if p.strict {
				return fmt.Errorf("yamlmini: %s: unknown field %q in type %s", path, key, out.Type())
			}
			continue
		}
		if err := p.bindValue(n.Vals[i], out.Field(f.index), path+"."+key); err != nil {
			return err
		}
	}
	return nil
}

func (p bindPolicy) bindMap(n *Node, out reflect.Value, path string) error {
	if n.Kind == Null {
		return nil
	}
	if n.Kind != Mapping {
		return shapeError(path, out.Type(), n.Kind)
	}
	t := out.Type()
	if t.Key().Kind() != reflect.String {
		return fmt.Errorf("yamlmini: %s: map key must have string kind, got %s", path, t.Key())
	}
	m := reflect.MakeMapWithSize(t, len(n.Keys))
	for i, key := range n.Keys {
		kv := reflect.New(t.Key()).Elem()
		kv.SetString(key)
		vv := reflect.New(t.Elem()).Elem()
		if err := p.bindValue(n.Vals[i], vv, path+"."+key); err != nil {
			return err
		}
		m.SetMapIndex(kv, vv)
	}
	out.Set(m)
	return nil
}

func (p bindPolicy) bindSlice(n *Node, out reflect.Value, path string) error {
	if n.Kind == Null {
		return nil
	}
	if n.Kind != Sequence {
		return shapeError(path, out.Type(), n.Kind)
	}
	s := reflect.MakeSlice(out.Type(), 0, len(n.Items))
	for i, item := range n.Items {
		elem := reflect.New(out.Type().Elem()).Elem()
		if err := p.bindValue(item, elem, fmt.Sprintf("%s[%d]", path, i)); err != nil {
			return err
		}
		s = reflect.Append(s, elem)
	}
	out.Set(s)
	return nil
}

func (p bindPolicy) bindArray(n *Node, out reflect.Value, path string) error {
	if n.Kind != Sequence {
		return shapeError(path, out.Type(), n.Kind)
	}
	if len(n.Items) != out.Len() {
		return fmt.Errorf("yamlmini: %s: sequence has %d items, want %d", path, len(n.Items), out.Len())
	}
	for i, item := range n.Items {
		if err := p.bindValue(item, out.Index(i), fmt.Sprintf("%s[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}

func bindInterface(n *Node, out reflect.Value, path string) error {
	if n.Kind == Null || out.NumMethod() == 0 {
		generic := nodeToAny(n)
		if generic == nil {
			out.Set(reflect.Zero(out.Type()))
		} else {
			out.Set(reflect.ValueOf(generic))
		}
		return nil
	}
	return fmt.Errorf("yamlmini: %s: cannot decode into non-empty interface %s", path, out.Type())
}

func shapeError(path string, t reflect.Type, got Kind) error {
	return fmt.Errorf("yamlmini: %s: cannot decode %s into %s", path, got, t)
}

func scalarTypeError(path string, t reflect.Type, n *Node) error {
	return fmt.Errorf("yamlmini: %s: cannot decode %q into %s", path, n.Value, t)
}

func nodeText(n *Node, path string) (string, error) {
	switch n.Kind {
	case Scalar:
		return n.Value, nil
	case Null:
		return "", nil
	default:
		return "", shapeError(path, reflect.TypeOf(""), n.Kind)
	}
}

type fieldInfo struct {
	name  string
	index int
}

// fieldsOf caches the yaml-tag view of a struct type. Field order is the
// declaration order, which is also the order used by the encoder.
func fieldsOf(t reflect.Type) []fieldInfo {
	var out []fieldInfo
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if !f.IsExported() {
			continue
		}
		name := f.Name
		if tag := f.Tag.Get("yaml"); tag != "" {
			parts := strings.Split(tag, ",")
			if parts[0] == "-" {
				continue
			}
			if parts[0] != "" {
				name = parts[0]
			}
		}
		out = append(out, fieldInfo{name: name, index: i})
	}
	return out
}

func fieldByName(fields []fieldInfo, name string) (fieldInfo, bool) {
	for _, f := range fields {
		if f.name == name {
			return f, true
		}
	}
	return fieldInfo{}, false
}
