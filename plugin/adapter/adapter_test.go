/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package adapter

import (
	"errors"
	"regexp"
	"testing"

	"abate/plugin/host"
)

// newTestAdapter builds an adapter with a fixed manifest.
func newTestAdapter() *CIAdapter {
	return NewCIAdapter(host.Manifest{
		ID:   host.PluginID("sample/adp-ci"),
		Type: host.PluginTypeAdapter,
	})
}

func TestAdapterMetadata(t *testing.T) {
	a := newTestAdapter()
	if a.Manifest().ID != host.PluginID("sample/adp-ci") {
		t.Fatalf("manifest id = %q", a.Manifest().ID)
	}
	if a.Type() != host.PluginTypeAdapter {
		t.Fatalf("type = %q", a.Type())
	}
	if err := a.Close(); err != nil {
		t.Fatalf("Close: %v", err)
	}
	if ADPAdapterID != "adp-ci" {
		t.Fatalf("adapter id = %q", ADPAdapterID)
	}
}

func TestDetectKnownProviders(t *testing.T) {
	cases := []struct {
		name string
		env  map[string]string
		want bool
	}{
		{"ci-true", map[string]string{"CI": "true"}, true},
		{"github", map[string]string{"GITHUB_ACTIONS": "true"}, true},
		{"gitlab", map[string]string{"GITLAB_CI": "true"}, true},
		{"jenkins", map[string]string{"JENKINS_URL": "true"}, true},
		{"multiple", map[string]string{"CI": "true", "GITLAB_CI": "true"}, true},
		{"ci-false", map[string]string{"CI": "false"}, false},
		{"ci-one", map[string]string{"CI": "1"}, false},
		{"unrelated", map[string]string{"HOME": "/root"}, false},
		{"empty", map[string]string{}, false},
		{"nil", nil, false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			id, ok := newTestAdapter().Detect(c.env)
			if ok != c.want {
				t.Fatalf("Detect = (%q,%v), want ok=%v", id, ok, c.want)
			}
			if c.want && id != ADPAdapterID {
				t.Fatalf("detected id = %q, want %q", id, ADPAdapterID)
			}
			if !c.want && id != "" {
				t.Fatalf("unmatched detection returned id %q", id)
			}
		})
	}
}

func TestEmitBaselineAssertions(t *testing.T) {
	a := newTestAdapter()
	cases := []struct {
		name    string
		payload []byte
		wantOK  bool
	}{
		{"ascii", []byte("build-42"), true},
		{"utf8", []byte("构建完成"), true},
		{"embedded-nul", []byte{'a', 0x00, 'b'}, true},
		{"empty", []byte{}, false},
		{"nil", nil, false},
		{"invalid-utf8", []byte{0xff, 0xfe}, false},
		{"truncated-utf8", []byte{0xe4, 0xb8}, false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := a.Emit(c.payload)
			if c.wantOK && err != nil {
				t.Fatalf("Emit rejected a valid payload: %v", err)
			}
			if !c.wantOK && !errors.Is(err, ErrAdapterPayloadRejected) {
				t.Fatalf("Emit = %v, want ErrAdapterPayloadRejected", err)
			}
		})
	}
}

func TestEmitOptionalPatternAssertion(t *testing.T) {
	a := newTestAdapter()
	pattern, err := regexp.Compile(`^OK-[0-9]+$`)
	if err != nil {
		t.Fatalf("compile payload pattern: %v", err)
	}
	a.payloadPattern = pattern
	if err := a.Emit([]byte("OK-1234")); err != nil {
		t.Fatalf("pattern-matching payload rejected: %v", err)
	}
	if err := a.Emit([]byte("BAD-1234")); !errors.Is(err, ErrAdapterPayloadRejected) {
		t.Fatalf("non-matching payload = %v", err)
	}
	if err := a.Emit(nil); !errors.Is(err, ErrAdapterPayloadRejected) {
		t.Fatalf("empty payload = %v", err)
	}
	if err := a.Emit([]byte{0xff}); !errors.Is(err, ErrAdapterPayloadRejected) {
		t.Fatalf("invalid UTF-8 payload = %v", err)
	}
}
