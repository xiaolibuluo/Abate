/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Command gen emits the CES v1 JSON Schema and the field reference from the
// Go data model. The generated files are the normative publishing artifacts.
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"abate/ces/schema"
)

const schemaID = "https://github.com/xiaolibuluo/Abate/ces/v1.schema.json"

// fieldDoc is one row of the generated field reference.
type fieldDoc struct {
	Owner    string
	Name     string
	GoType   string
	Required bool
	Default  string
	Rule     string
	Example  string
}

// builder accumulates schema definitions and field rows during the walk.
type builder struct {
	defs map[string]map[string]any
	docs []fieldDoc
	seen map[reflect.Type]bool
}

func main() {
	dir := "ces/schema/dist"
	if len(os.Args) > 1 {
		dir = os.Args[1]
	}
	if err := os.MkdirAll(dir, 0o750); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	b := &builder{
		defs: map[string]map[string]any{},
		seen: map[reflect.Type]bool{},
	}
	root := reflect.TypeOf(schema.Experiment{})
	b.seen[root] = true
	rootSchema := b.buildObject(root)
	rootSchema["$schema"] = "https://json-schema.org/draft/2020-12/schema"
	rootSchema["$id"] = schemaID
	rootSchema["title"] = "CES Experiment v1"
	rootSchema["additionalProperties"] = false

	defs := make(map[string]any, len(b.defs))
	for name, def := range b.defs {
		defs[name] = def
	}
	rootSchema["$defs"] = defs

	if err := writeJSON(filepath.Join(dir, "ces-v1.schema.json"), rootSchema); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fieldsPath := filepath.Join(dir, "fields.md")
	if err := os.WriteFile(fieldsPath, []byte(b.fieldsMarkdown()), 0o600); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// buildObject produces an object schema for a struct and records its fields.
func (b *builder) buildObject(t reflect.Type) map[string]any {
	props := map[string]any{}
	var required []string
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		name, omit := jsonField(f)
		if name == "-" {
			continue
		}
		props[name] = b.typeSchema(f.Type)
		if !omit {
			required = append(required, name)
		}
		b.docs = append(b.docs, makeDoc(t.Name(), f, name, omit))
	}
	out := map[string]any{
		"type":                 "object",
		"properties":           props,
		"additionalProperties": false,
	}
	if len(required) > 0 {
		out["required"] = required
	}
	return out
}

// typeSchema maps a Go type to its JSON Schema fragment.
func (b *builder) typeSchema(t reflect.Type) map[string]any {
	if t.Name() == "VirtualDuration" {
		return map[string]any{
			"type":        "string",
			"pattern":     "^-?[0-9]+(ns|us|µs|ms|s|m|h)*$",
			"description": "Integer nanosecond duration with optional unit suffixes.",
		}
	}
	switch t.Kind() {
	case reflect.String:
		out := map[string]any{"type": "string"}
		if values := enumOf(t); len(values) > 0 {
			out["enum"] = values
		}
		return out
	case reflect.Int, reflect.Int64, reflect.Uint, reflect.Uint64:
		return map[string]any{"type": "integer"}
	case reflect.Float64:
		return map[string]any{"type": "number"}
	case reflect.Bool:
		return map[string]any{"type": "boolean"}
	case reflect.Pointer:
		return b.typeSchema(t.Elem())
	case reflect.Slice:
		return map[string]any{"type": "array", "items": b.typeSchema(t.Elem())}
	case reflect.Map:
		return map[string]any{
			"type":                 "object",
			"additionalProperties": b.typeSchema(t.Elem()),
		}
	case reflect.Struct:
		return b.structRef(t)
	default:
		return map[string]any{}
	}
}

// structRef registers a named struct once and returns a reference to it.
func (b *builder) structRef(t reflect.Type) map[string]any {
	ref := "#/$defs/" + t.Name()
	if b.seen[t] {
		return map[string]any{"$ref": ref}
	}
	b.seen[t] = true
	b.defs[t.Name()] = b.buildObject(t)
	return map[string]any{"$ref": ref}
}

// enumOf returns the closed value set of the CES enumerated types.
func enumOf(t reflect.Type) []string {
	switch t.Name() {
	case "ProbeKind":
		return []string{"http", "shell", "metric", "prom"}
	case "Strategy":
		return []string{"auto", "manual", "timeout"}
	default:
		return nil
	}
}

// jsonField extracts the JSON name and the omitempty flag from a struct tag.
func jsonField(f reflect.StructField) (string, bool) {
	tag := f.Tag.Get("json")
	if tag == "" {
		return f.Name, false
	}
	parts := strings.Split(tag, ",")
	name := parts[0]
	if name == "" {
		name = f.Name
	}
	omit := false
	for _, p := range parts[1:] {
		if p == "omitempty" {
			omit = true
		}
	}
	return name, omit
}

// makeDoc derives one field-reference row from a struct field.
func makeDoc(owner string, f reflect.StructField, name string, omit bool) fieldDoc {
	d := fieldDoc{
		Owner:    owner,
		Name:     name,
		GoType:   f.Type.String(),
		Required: !omit,
		Default:  "—",
		Rule:     "见 CES-VALID 规则清单",
		Example:  "—",
	}
	switch name {
	case "apiVersion":
		d.Default, d.Rule, d.Example = "必填", "固定 ces.tencent.com/v1", "ces.tencent.com/v1"
	case "kind":
		d.Default, d.Rule, d.Example = "必填", "固定 Experiment", "Experiment"
	case "namespace":
		d.Default, d.Example = "default", "default"
	case "strategy":
		d.Default, d.Rule, d.Example = "auto", "auto/manual/timeout", "auto"
	case "verifySteadyState":
		d.Default, d.Example = "true", "true"
	case "seed":
		d.Default, d.Rule, d.Example = "CLI 生成", "uint64 确定性种子", "12345"
	case "type":
		d.Rule, d.Example = "枚举见对应类型说明", "value"
	case "interval", "holdFor", "timeout", "at", "duration":
		d.Rule, d.Example = "整数纳秒时长，须 > 0", "5s"
	}
	if !d.Required && d.Default == "—" {
		d.Default = "省略"
	}
	return d
}

// fieldsMarkdown renders the deterministic field-reference document.
func (b *builder) fieldsMarkdown() string {
	var out strings.Builder
	_, _ = out.WriteString("# CES v1 字段参考\n\n")
	_, _ = out.WriteString("本文件由 `ces/schema/gen` 从 Go 数据模型自动生成，请勿手工编辑。\n\n")
	_, _ = out.WriteString("| 类型 | 字段 | Go 类型 | 必填 | 默认 | 约束 | 示例 |\n")
	_, _ = out.WriteString("| --- | --- | --- | --- | --- | --- | --- |\n")
	for _, d := range b.docs {
		req := "否"
		if d.Required {
			req = "是"
		}
		_, _ = fmt.Fprintf(&out, "| %s | `%s` | `%s` | %s | %s | %s | %s |\n",
			d.Owner, d.Name, d.GoType, req, d.Default, d.Rule, d.Example)
	}
	return out.String()
}

// writeJSON writes a stable, indented JSON document.
func writeJSON(path string, v any) error {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	data = append(data, '\n')
	return os.WriteFile(path, data, 0o600)
}
