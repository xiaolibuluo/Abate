/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package yamlmini

import (
	"bytes"
	"strings"
	"testing"
)

type testNested struct {
	ProbeRef string `yaml:"probeRef"`
	Value    string `yaml:"value"`
}

type testDoc struct {
	APIVersion string            `yaml:"apiVersion"`
	Replicas   int               `yaml:"replicas"`
	Ratio      float64           `yaml:"ratio"`
	Enabled    bool              `yaml:"enabled"`
	Name       string            `yaml:"name"`
	Labels     map[string]string `yaml:"labels"`
	Targets    []string          `yaml:"targets"`
	Nested     []testNested      `yaml:"nested"`
	Optional   string            `yaml:"optional,omitempty"`
}

func TestScalarWhitelist(t *testing.T) {
	t.Parallel()
	data := []byte(`apiVersion: ces/v1
replicas: 3
ratio: 1.5
enabled: true
name: "a: b # not a comment"
labels:
  team: 内核组
targets:
  - n1
  - "n2"
nested:
  - probeRef: p1
    value: "42"
  - probeRef: p2
    value: ok
`)
	var got testDoc
	if err := Unmarshal(data, &got); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if got.APIVersion != "ces/v1" || got.Replicas != 3 || got.Ratio != 1.5 || !got.Enabled {
		t.Fatalf("scalars: %+v", got)
	}
	if got.Name != "a: b # not a comment" {
		t.Fatalf("quoted scalar = %q", got.Name)
	}
	if got.Labels["team"] != "内核组" || len(got.Targets) != 2 || got.Targets[1] != "n2" {
		t.Fatalf("map/slice: %+v", got)
	}
	if len(got.Nested) != 2 || got.Nested[0].Value != "42" || got.Nested[1].Value != "ok" {
		t.Fatalf("nested: %+v", got)
	}
}

func TestCommentsAndNull(t *testing.T) {
	t.Parallel()
	var got testDoc
	err := Unmarshal([]byte("apiVersion: v1 # trailing comment\n# whole line\nname: ~\noptional:\n"), &got)
	if err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if got.APIVersion != "v1" || got.Name != "" || got.Optional != "" {
		t.Fatalf("got %+v", got)
	}
}

func TestKnownFieldsStrict(t *testing.T) {
	t.Parallel()
	data := []byte("apiVersion: v1\nunexpected: 1\n")
	var got testDoc
	err := NewDecoder(bytes.NewReader(data)).KnownFields(true).Decode(&got)
	if err == nil || !strings.Contains(err.Error(), "unknown field") {
		t.Fatalf("strict decode error = %v, want 'unknown field'", err)
	}
	// Non-strict mode tolerates the same key.
	if err := Unmarshal(data, &got); err != nil {
		t.Fatalf("non-strict decode: %v", err)
	}
}

func TestTextUnmarshaler(t *testing.T) {
	t.Parallel()
	type td struct {
		Hold string `yaml:"hold"`
	}
	// A duration-like plain scalar binds as a string; the schema duration test
	// exercises the real TextUnmarshaler through UnmarshalText in schema tests.
	var got td
	if err := Unmarshal([]byte("hold: 250ms\n"), &got); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if got.Hold != "250ms" {
		t.Fatalf("hold = %q", got.Hold)
	}
}

func TestBlockScalarControlledExtension(t *testing.T) {
	t.Parallel()
	var got struct {
		ID    string   `yaml:"id"`
		Input string   `yaml:"input"`
		Empty []string `yaml:"ids"`
	}
	raw := "id: TCK-001\ninput: |\n  apiVersion: v1\n  metadata:\n      name: x\nids: []\n"
	if err := Unmarshal([]byte(raw), &got); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if got.ID != "TCK-001" {
		t.Fatalf("id = %q", got.ID)
	}
	if got.Input != "apiVersion: v1\nmetadata:\n    name: x\n" {
		t.Fatalf("input = %q", got.Input)
	}
	if got.Empty == nil || len(got.Empty) != 0 {
		t.Fatalf("empty flow seq = %#v", got.Empty)
	}
}

func TestGenericAnyShape(t *testing.T) {
	t.Parallel()
	var got map[string]any
	raw := "apiVersion: v1\ncount: 2\nratio: 1.5\nok: true\nnothing: null\nitems:\n  - a\n  - b\nmeta:\n  k: v\n"
	if err := Unmarshal([]byte(raw), &got); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if got["apiVersion"] != "v1" || got["count"] != int(2) || got["ratio"] != 1.5 || got["ok"] != true {
		t.Fatalf("scalars: %#v", got)
	}
	if _, ok := got["nothing"]; !ok {
		t.Fatalf("null key missing")
	}
	items, ok := got["items"].([]any)
	if !ok || len(items) != 2 {
		t.Fatalf("items: %#v", got["items"])
	}
	meta, ok := got["meta"].(map[string]any)
	if !ok || meta["k"] != "v" {
		t.Fatalf("meta: %#v", got["meta"])
	}
}

func TestDeterministicMarshalRoundTrip(t *testing.T) {
	t.Parallel()
	in := testDoc{
		APIVersion: "ces/v1",
		Replicas:   2,
		Ratio:      0.5,
		Enabled:    true,
		Name:       "name with <>&\"'\n\t chars",
		Labels:     map[string]string{"z": "last", "a": "first"},
		Targets:    []string{"n1", "n2"},
		Nested:     []testNested{{ProbeRef: "p1", Value: "42"}},
	}
	first, err := Marshal(in)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	second, err := Marshal(in)
	if err != nil {
		t.Fatalf("Marshal again: %v", err)
	}
	if !bytes.Equal(first, second) {
		t.Fatalf("Marshal not deterministic:\n%s\n%s", first, second)
	}
	var back testDoc
	if err := Unmarshal(first, &back); err != nil {
		t.Fatalf("Unmarshal(%s): %v", first, err)
	}
	if back.Name != in.Name || !bytes.Equal([]byte(back.Name), []byte(in.Name)) {
		t.Fatalf("string round trip drift: %q -> %q", in.Name, back.Name)
	}
	if back.Replicas != in.Replicas || back.Ratio != in.Ratio || back.Enabled != in.Enabled {
		t.Fatalf("scalar round trip drift: %+v", back)
	}
	if back.Labels["a"] != "first" || back.Labels["z"] != "last" || len(back.Targets) != 2 {
		t.Fatalf("map/slice round trip drift: %+v", back)
	}
	third, err := Marshal(back)
	if err != nil || !bytes.Equal(first, third) {
		t.Fatalf("byte round trip drift (err %v):\n%s\n%s", err, first, third)
	}
}

func TestFlowSequenceScalarOnly(t *testing.T) {
	t.Parallel()
	var got struct {
		Targets []string `yaml:"targets"`
	}
	if err := Unmarshal([]byte("targets: [disk, \"net\", 'cpu', 7, true, null]\n"), &got); err != nil {
		t.Fatalf("Unmarshal flow sequence: %v", err)
	}
	if len(got.Targets) != 6 || got.Targets[0] != "disk" || got.Targets[1] != "net" ||
		got.Targets[2] != "cpu" || got.Targets[3] != "7" || got.Targets[4] != "true" || got.Targets[5] != "" {
		t.Fatalf("flow targets = %#v", got.Targets)
	}
	// Unterminated flow must still be an error, not accepted as a scalar.
	var bad struct {
		Spec string `yaml:"spec"`
	}
	if err := Unmarshal([]byte("spec: [unterminated\n"), &bad); err == nil {
		t.Fatal("unterminated flow sequence must be rejected")
	}
	// Nested flow collections are out of scope.
	var nested struct {
		X []string `yaml:"x"`
	}
	if err := Unmarshal([]byte("x: [[a], b]\n"), &nested); err == nil {
		t.Fatal("nested flow sequence must be rejected")
	}
}

func TestUnsupportedSyntaxRejected(t *testing.T) {
	t.Parallel()
	cases := []string{
		"a: &anchor v\nb: *anchor\n", // anchors/aliases
		"a: !!str v\n",               // custom tags
		"a: |\n  x\nb: >\n  y\n",     // folded block scalar
		"---\na: v\n",                // multi-document marker
		"? [a, b]: v\n",              // complex key
	}
	for _, raw := range cases {
		var out any
		if err := Unmarshal([]byte(raw), &out); err == nil {
			t.Fatalf("expected rejection for %q", raw)
		}
	}
}
