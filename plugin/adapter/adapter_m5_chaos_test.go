/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package adapter

import (
	"errors"
	"strings"
	"sync"
	"testing"
)

// TestM5ChaosMalformedPayloadsRejected proves hostile inputs are rejected
// without a panic on both adapter assertion contracts.
func TestM5ChaosMalformedPayloadsRejected(t *testing.T) {
	t.Parallel()
	its := NewPayloadAssertion(itsRequiredFields)
	msg := NewPayloadAssertion(msgRequiredFields)
	truncated := append([]byte(`{"type":"`), 0xe4, 0xb8)
	oversized := []byte(`{"type":"` + strings.Repeat("x", 70000) +
		`","fingerprint":"a","action":"open","labels":[]}`)
	cases := []struct {
		name      string
		assertion PayloadAssertion
		payload   []byte
	}{
		{"its-nil", its, nil},
		{"its-empty", its, []byte{}},
		{"its-invalid-utf8", its, []byte{0xff, 0xfe}},
		{"its-truncated-utf8", its, truncated},
		{"its-oversized", its, oversized},
		{"its-missing-all", its, []byte(`{}`)},
		{"its-missing-one", its, []byte(`{"type":"x","fingerprint":"a","action":"open"}`)},
		{"its-nested-endpoint", its, []byte(`{"type":"x","fingerprint":"a","action":"open","labels":[],"meta":{"endpoint":"x"}}`)},
		{"its-array", its, []byte(`[{"endpoint":"x"}]`)},
		{"its-number", its, []byte(`42`)},
		{"its-string", its, []byte(`"text"`)},
		{"msg-nil", msg, nil},
		{"msg-oversized", msg, oversized},
		{"msg-missing-summary", msg, []byte(`{"kind":"report","level":"info"}`)},
		{"msg-token", msg, []byte(`{"kind":"report","level":"info","summary":"ok","token":"x"}`)},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if err := c.assertion.Assert(c.payload); !errors.Is(err, ErrAdapterPayloadRejected) {
				t.Fatalf("hostile payload was not rejected: %v", err)
			}
		})
	}
}

// TestM5ChaosConcurrentEmit proves concurrent assertions are race-free and
// keep their accept/reject decisions.
func TestM5ChaosConcurrentEmit(t *testing.T) {
	t.Parallel()
	its := newITSAdapter(m5AdapterManifest("abate/adp-its"))
	msg := newMSGAdapter(m5AdapterManifest("abate/adp-msg"))
	type emitResult struct {
		index int
		err   error
	}
	var group sync.WaitGroup
	results := make(chan emitResult, 32)
	for i := 0; i < 32; i++ {
		group.Add(1)
		go func(index int) {
			defer group.Done()
			var err error
			switch index % 4 {
			case 0:
				err = its.Emit(m5ITSPayload())
			case 1:
				err = msg.Emit(m5MSGPayload())
			case 2:
				err = its.Emit([]byte("broken"))
			default:
				err = msg.Emit([]byte(`{"endpoint":"x"}`))
			}
			results <- emitResult{index: index, err: err}
		}(i)
	}
	group.Wait()
	close(results)
	for result := range results {
		wantOK := result.index%4 == 0 || result.index%4 == 1
		if wantOK && result.err != nil {
			t.Fatalf("concurrent valid emit %d failed: %v", result.index, result.err)
		}
		if !wantOK && !errors.Is(result.err, ErrAdapterPayloadRejected) {
			t.Fatalf("concurrent invalid emit %d was accepted", result.index)
		}
	}
}

// TestM5ChaosDetectRepeatable proves detection is deterministic across one
// hundred identical injected environments.
func TestM5ChaosDetectRepeatable(t *testing.T) {
	t.Parallel()
	its := newITSAdapter(m5AdapterManifest("abate/adp-its"))
	env := map[string]string{"ABATE_ADP_ITS": "true", "HOME": "/root"}
	for i := 0; i < 100; i++ {
		id, ok := its.Detect(env)
		if !ok || id != itsAdapterID {
			t.Fatalf("detect drifted at iteration %d: (%q,%v)", i, id, ok)
		}
	}
}
