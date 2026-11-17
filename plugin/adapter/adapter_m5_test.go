/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package adapter

import (
	"bytes"
	"encoding/json"
	"errors"
	"go/parser"
	"go/token"
	"os"
	"strings"
	"testing"

	"abate/plugin/capability"
	"abate/plugin/host"
)

// m5AdapterManifest builds an adapter manifest with a fixed short id.
func m5AdapterManifest(id string) host.Manifest {
	return host.Manifest{ID: host.PluginID(id), Type: host.PluginTypeAdapter}
}

// m5ITSPayload is a valid issue payload satisfying the ITS assertion.
func m5ITSPayload() []byte {
	return []byte(`{"type":"defect","fingerprint":"a1b2c3","action":"open","labels":["m5","local"]}`)
}

// m5MSGPayload is a valid event payload satisfying the MSG assertion.
func m5MSGPayload() []byte {
	return []byte(`{"kind":"report","level":"info","summary":"deterministic build"}`)
}

// TestM5PayloadAcceptanceMatrix proves the shared assertion accepts well-formed
// issue/event payloads and rejects every malformed shape.
func TestM5PayloadAcceptanceMatrix(t *testing.T) {
	t.Parallel()
	its := NewPayloadAssertion(itsRequiredFields)
	msg := NewPayloadAssertion(msgRequiredFields)
	cases := []struct {
		name      string
		assertion PayloadAssertion
		payload   []byte
		wantOK    bool
	}{
		{"its-valid", its, m5ITSPayload(), true},
		{"msg-valid", msg, m5MSGPayload(), true},
		{"its-nested-safe", its, []byte(`{"type":"defect","fingerprint":"a1","action":"open","labels":[],"details":{"owner":"qa"}}`), true},
		{"its-empty", its, nil, false},
		{"its-invalid-utf8", its, []byte{0xff, 0xfe}, false},
		{"its-not-json", its, []byte("plain text"), false},
		{"its-json-array", its, []byte(`[]`), false},
		{"its-json-null", its, []byte(`null`), false},
		{"its-missing-field", its, []byte(`{"type":"defect"}`), false},
		{"its-nil-field", its, []byte(`{"type":null,"fingerprint":"a1","action":"open","labels":[]}`), false},
		{"msg-missing-field", msg, []byte(`{"kind":"report","level":"info"}`), false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := c.assertion.Assert(c.payload)
			if c.wantOK && err != nil {
				t.Fatalf("valid payload was rejected: %v", err)
			}
			if !c.wantOK && !errors.Is(err, ErrAdapterPayloadRejected) {
				t.Fatalf("invalid payload error = %v, want ErrAdapterPayloadRejected", err)
			}
		})
	}
}

// TestM5OutboundFieldRejectionMatrix proves every outbound field name is
// rejected both at the top level and at a nested level.
func TestM5OutboundFieldRejectionMatrix(t *testing.T) {
	t.Parallel()
	its := NewPayloadAssertion(itsRequiredFields)
	markers := []string{
		"endpoint", "url", "token", "host", "address", "webhook",
		"uri", "secret", "credential", "apikey", "api_key",
		"authorization", "apiEndpoint", "webhook_url", "HOST", "X-Api-Token",
	}
	for _, marker := range markers {
		payload := []byte(`{"type":"defect","fingerprint":"a1","action":"open","labels":[],"` +
			marker + `":"x"}`)
		if err := its.Assert(payload); !errors.Is(err, ErrAdapterPayloadRejected) {
			t.Errorf("outbound field %q was not rejected", marker)
		}
	}
	nested := []byte(`{"type":"defect","fingerprint":"a1","action":"open","labels":[],"metadata":{"endpoint":"x"}}`)
	if err := its.Assert(nested); !errors.Is(err, ErrAdapterPayloadRejected) {
		t.Error("a nested outbound field was not rejected")
	}
	inArray := []byte(`{"type":"defect","fingerprint":"a1","action":"open","labels":[{"url":"x"}]}`)
	if err := its.Assert(inArray); !errors.Is(err, ErrAdapterPayloadRejected) {
		t.Error("an outbound field inside an array was not rejected")
	}
	msg := NewPayloadAssertion(msgRequiredFields)
	msgPayload := []byte(`{"kind":"report","level":"info","summary":"ok","token":"x"}`)
	if err := msg.Assert(msgPayload); !errors.Is(err, ErrAdapterPayloadRejected) {
		t.Error("MSG accepted an outbound token field")
	}
}

// TestM5DetectUsesInjectedEnv proves detection depends only on the injected map.
func TestM5DetectUsesInjectedEnv(t *testing.T) {
	t.Parallel()
	its := newITSAdapter(m5AdapterManifest("abate/adp-its"))
	msg := newMSGAdapter(m5AdapterManifest("abate/adp-msg"))
	itsCases := []struct {
		env    map[string]string
		wantOK bool
		wantID string
	}{
		{map[string]string{"ABATE_ADP_ITS": "true"}, true, itsAdapterID},
		{map[string]string{"ITS_ENABLED": "its"}, true, itsAdapterID},
		{map[string]string{"ISSUE_TRACKER": "true"}, true, itsAdapterID},
		{map[string]string{"ABATE_ADP_ITS": "false"}, false, ""},
		{map[string]string{"HOME": "/root"}, false, ""},
		{map[string]string{}, false, ""},
		{nil, false, ""},
	}
	for i, c := range itsCases {
		id, ok := its.Detect(c.env)
		if ok != c.wantOK || id != c.wantID {
			t.Fatalf("ITS detect case %d = (%q,%v), want (%q,%v)",
				i, id, ok, c.wantID, c.wantOK)
		}
	}
	msgCases := []struct {
		env    map[string]string
		wantOK bool
	}{
		{map[string]string{"ABATE_ADP_MSG": "true"}, true},
		{map[string]string{"MSG_ENABLED": "msg"}, true},
		{map[string]string{"MESSAGE_SINK": "true"}, true},
		{map[string]string{"ABATE_ADP_MSG": "1"}, false},
		{nil, false},
	}
	for i, c := range msgCases {
		id, ok := msg.Detect(c.env)
		if ok != c.wantOK {
			t.Fatalf("MSG detect case %d ok = %v, want %v", i, ok, c.wantOK)
		}
		if c.wantOK && id != msgAdapterID {
			t.Fatalf("MSG detect case %d returned id %q", i, id)
		}
	}
}

// TestM5HostCallEntries proves the real host dispatch reaches Detect and Emit.
func TestM5HostCallEntries(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name    string
		factory host.Factory
		envKey  string
		wantID  string
		payload []byte
	}{
		{"its", ITSFactory, "ABATE_ADP_ITS", itsAdapterID, m5ITSPayload()},
		{"msg", MSGFactory, "ABATE_ADP_MSG", msgAdapterID, m5MSGPayload()},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			plugin, err := c.factory(m5AdapterManifest("abate/" + c.name))
			if err != nil {
				t.Fatalf("factory: %v", err)
			}
			var caps host.Capabilities
			caps.Grant(capability.CapObserve)
			inst := &host.Instance{
				Plugin: plugin, Caps: caps,
				Entries: []string{host.EntryDetect, host.EntryEmit},
			}
			env, err := json.Marshal(map[string]string{c.envKey: "true"})
			if err != nil {
				t.Fatalf("encode env: %v", err)
			}
			h := &host.Host{}
			out, err := h.Call(inst, host.EntryDetect, env, nil)
			if err != nil {
				t.Fatalf("Host.Call(detect): %v", err)
			}
			var result host.AdapterDetectResult
			if err := json.Unmarshal(out, &result); err != nil {
				t.Fatalf("decode detect result: %v", err)
			}
			if !result.OK || result.ID != c.wantID {
				t.Fatalf("detect result = %+v", result)
			}
			if _, err := h.Call(inst, host.EntryEmit, c.payload, nil); err != nil {
				t.Fatalf("Host.Call(emit): %v", err)
			}
			if _, err := h.Call(inst, host.EntryEmit, []byte(`{"endpoint":"x"}`), nil); err == nil {
				t.Fatal("Host.Call(emit) accepted an outbound payload")
			}
		})
	}
}

// TestM5FactoriesFailClosed proves both factories reject a wrong plugin type.
func TestM5FactoriesFailClosed(t *testing.T) {
	t.Parallel()
	wrong := m5AdapterManifest("abate/wrong")
	wrong.Type = host.PluginTypeReporter
	for _, factory := range []host.Factory{ITSFactory, MSGFactory} {
		if _, err := factory(wrong); !errors.Is(err, ErrM5WrongPluginType) {
			t.Fatalf("wrong type error = %v, want ErrM5WrongPluginType", err)
		}
		plugin, err := factory(m5AdapterManifest("abate/valid"))
		if err != nil {
			t.Fatalf("valid factory: %v", err)
		}
		if plugin.Type() != host.PluginTypeAdapter {
			t.Fatalf("factory produced type %q", plugin.Type())
		}
		if _, ok := plugin.(host.AdapterPlugin); !ok {
			t.Fatal("factory did not produce an AdapterPlugin")
		}
	}
}

// TestM5AdapterMetadata proves both adapters satisfy the base host contract.
func TestM5AdapterMetadata(t *testing.T) {
	t.Parallel()
	m := m5AdapterManifest("abate/adp-m5")
	for _, adapter := range []host.AdapterPlugin{
		newITSAdapter(m), newMSGAdapter(m),
	} {
		if adapter.Manifest().ID != m.ID || adapter.Type() != host.PluginTypeAdapter {
			t.Fatalf("metadata mismatch: %+v", adapter.Manifest())
		}
		if err := adapter.Close(); err != nil {
			t.Fatalf("Close: %v", err)
		}
	}
}

// TestM5AirGapStaticAssertion proves the implementations contain no network
// client, no dial, no subprocess, no environment read and no wall-clock use.
func TestM5AirGapStaticAssertion(t *testing.T) {
	t.Parallel()
	files := []string{"payload.go", "its.go", "msg.go", "factories.go", "ci.go"}
	forbidden := []string{
		"net/http", "net.Dial", "os/exec", "os.Environ", "http.Get",
		"http.Post", "http.DefaultClient", "exec.Command", "time.Now",
		"math/rand", "os.Getenv",
	}
	bannedImports := []string{"net", "net/http", "os/exec", "os", "math/rand", "time"}
	for _, name := range files {
		data, err := os.ReadFile(name) //nolint:gosec // fixed package-local source names.
		if err != nil {
			t.Fatalf("read %s: %v", name, err)
		}
		for _, marker := range forbidden {
			if bytes.Contains(data, []byte(marker)) {
				t.Errorf("%s leaks forbidden marker %q", name, marker)
			}
		}
		parsed, err := parser.ParseFile(token.NewFileSet(), name, nil, parser.ImportsOnly)
		if err != nil {
			t.Fatalf("parse %s: %v", name, err)
		}
		for _, spec := range parsed.Imports {
			path := strings.Trim(spec.Path.Value, `"`)
			for _, banned := range bannedImports {
				if path == banned {
					t.Errorf("%s imports forbidden package %q", name, path)
				}
			}
		}
	}
}
