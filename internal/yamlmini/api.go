/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package yamlmini implements a deliberately small, self-contained YAML subset
// decoder/encoder on the Go standard library only. It is the project's own
// implementation and contains no third-party source.
//
// Supported grammar (whitelist):
//   - indentation based block mappings "key: value" and block sequences "- item"
//   - arbitrary nesting of the two, empty values ("null", "~" or nothing) and
//     blank lines
//   - line comments starting with '#' outside quotes
//   - single- and double-quoted scalars plus plain scalars; integers, floats,
//     YAML 1.2 booleans (true/false) and null
//   - Go bindings for structs (with "yaml" tags), pointers, map[string]string,
//     map[string]any, []string, []struct and encoding.TextUnmarshaler
//   - strict KnownFields mode that rejects keys absent from the target struct
//
// Controlled extensions (approved; outside the strict block-style whitelist,
// but required so existing corpus/TCK bytes decode byte-for-byte):
//   - literal block scalar headers "|", "|-" and "|+", used only for string
//     values (the TCK stores each CES document as "input: |")
//   - the empty flow sequence "[]" (and, symmetrically, "{}"), used for empty
//     collection values (the TCK emits "ruleIds: []")
//   - single-line non-empty flow sequences containing only scalars, e.g.
//     "[disk]", "[a, b]" or "[\"x\", 'y']" (one CES corpus document uses
//     "targets: [disk]")
//
// Newly authored CES documents remain block-style; these extensions exist
// solely to keep the frozen corpus/TCK bytes intact.
//
// Explicitly NOT supported:
//   - anchors &a, aliases *a and merge keys <<
//   - folded block scalars ">" (only the controlled "|" extension exists)
//   - custom tags "!!" and complex / non-string keys
//   - nested or multi-line flow collections (a flow sequence may contain only
//     scalars on a single line; "spec: [unterminated" remains an error)
//   - streaming multi-document Decode ("---" separators)
//   - YAML 1.1 implicit conversions (yes/no/on/off remain plain strings)
//
// Any future grammar extension must go through the "extend the subset plus new
// golden" process; this package must never widen the grammar silently.
package yamlmini

import (
	"fmt"
	"io"
	"strings"
)

// Decoder reads and decodes a single YAML document from an input stream.
type Decoder struct {
	r      io.Reader
	strict bool
}

// NewDecoder returns a Decoder that reads from r.
func NewDecoder(r io.Reader) *Decoder { return &Decoder{r: r} }

// KnownFields enables or disables strict field checking. When enabled, a
// mapping key that has no matching field in the destination struct produces an
// error whose text contains "unknown field".
func (d *Decoder) KnownFields(enable bool) *Decoder {
	d.strict = enable
	return d
}

// Decode reads the whole stream as one document and binds it into v.
//
// Multi-document streams are deliberately rejected; Decode must be called on a
// fresh Decoder for each document.
func (d *Decoder) Decode(v any) error {
	if d == nil || d.r == nil {
		return fmt.Errorf("yamlmini: nil decoder")
	}
	data, err := io.ReadAll(d.r)
	if err != nil {
		return fmt.Errorf("yamlmini: read document: %w", err)
	}
	// A stream with no content (blank/comment-only) is not a document. The
	// third-party decoder this package replaces returned io.EOF here, which the
	// CES caller classified as a malformed document rather than an empty struct.
	if blankDocument(data) {
		return fmt.Errorf("yamlmini: empty document")
	}
	return decodeBytes(data, v, d.strict)
}

// blankDocument reports whether data contains no content line at all (only
// whitespace, blank lines and full-line comments).
func blankDocument(data []byte) bool {
	for _, raw := range strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n") {
		t := strings.TrimSpace(raw)
		if t == "" || strings.HasPrefix(t, "#") {
			continue
		}
		return false
	}
	return true
}

// Unmarshal decodes one YAML document into v with strict field checking off.
func Unmarshal(data []byte, v any) error { return decodeBytes(data, v, false) }

func decodeBytes(data []byte, v any, strict bool) error {
	doc, err := parseDocument(data)
	if err != nil {
		return err
	}
	return bind(doc, v, strict)
}

// Marshal encodes v as a deterministic YAML document. Struct fields are
// emitted in declaration order, map keys are sorted, and scalar quoting is
// canonical so repeated round trips are byte-stable.
func Marshal(v any) ([]byte, error) {
	doc, err := emitValue("", v)
	if err != nil {
		return nil, err
	}
	var b strings.Builder
	emitNode(&b, doc, 0)
	return []byte(b.String()), nil
}
