/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package schema

import (
	"encoding/json"
	"errors"
	"strings"
	"testing"

	"abate/internal/yamlmini"
)

// TestMutationVersionStrings rejects every single-character version mutation.
func TestMutationVersionStrings(t *testing.T) {
	t.Parallel()
	v := NewVersion()
	base := APIVersion
	for i := 0; i < len(base); i++ {
		mut := base[:i] + "X" + base[i+1:]
		if mut == base {
			continue
		}
		if _, err := v.Negotiate(mut); !errors.Is(err, ErrUnsupportedVersion) {
			t.Fatalf("Negotiate(%q) must reject a mutated version", mut)
		}
	}
	for _, mut := range []string{
		base + "x", base[:len(base)-1], " " + base, strings.ToUpper(base),
	} {
		if _, err := v.Negotiate(mut); !errors.Is(err, ErrUnsupportedVersion) {
			t.Fatalf("Negotiate(%q) must reject a mutated version", mut)
		}
	}
}

// TestMutationDurationLiterals checks that character mutations never yield an
// unstable value that cannot be reproduced from its canonical literal.
func TestMutationDurationLiterals(t *testing.T) {
	t.Parallel()
	base := "1h30m250ms"
	for i := 0; i < len(base); i++ {
		mutations := []string{
			base[:i] + base[i+1:],
			base[:i] + "9" + base[i:],
			base[:i] + "z" + base[i+1:],
		}
		for _, m := range mutations {
			d, err := ParseDuration(m)
			if err != nil {
				continue
			}
			back, err2 := ParseDuration(d.String())
			if err2 != nil || back != d {
				t.Fatalf("unstable mutation %q -> %d -> %q", m, d, d.String())
			}
		}
	}
}

// TestMutationUnknownFieldRejection proves the CES v1 strict-decoding contract.
func TestMutationUnknownFieldRejection(t *testing.T) {
	t.Parallel()
	doc, err := json.Marshal(sampleExperiment())
	if err != nil {
		t.Fatalf("json.Marshal err = %v", err)
	}
	tampered := strings.Replace(string(doc),
		`"kind":"Experiment"`, `"kind":"Experiment","unexpected":1`, 1)
	if tampered == string(doc) {
		t.Fatalf("mutation did not apply to the document")
	}
	dec := json.NewDecoder(strings.NewReader(tampered))
	dec.DisallowUnknownFields()
	var viaJSON Experiment
	if err := dec.Decode(&viaJSON); err == nil {
		t.Fatalf("strict JSON decoder must reject an unknown field")
	}
	// The tampered document is JSON (flow style), which yamlmini deliberately
	// does not parse. Route it through encoding/json into a generic tree and
	// re-emit as block YAML, so the strict YAML decoder still sees the injected
	// unknown key and must reject it.
	var generic any
	if err := json.Unmarshal([]byte(tampered), &generic); err != nil {
		t.Fatalf("json.Unmarshal(tampered) err = %v", err)
	}
	block, err := yamlmini.Marshal(generic)
	if err != nil {
		t.Fatalf("yamlmini.Marshal err = %v", err)
	}
	ydec := yamlmini.NewDecoder(strings.NewReader(string(block))).KnownFields(true)
	var viaYAML Experiment
	yerr := ydec.Decode(&viaYAML)
	if yerr == nil {
		t.Fatalf("strict YAML decoder must reject an unknown field")
	}
	if !strings.Contains(yerr.Error(), "unknown field") {
		t.Fatalf("strict YAML error %q does not classify as unknown field", yerr)
	}
}
