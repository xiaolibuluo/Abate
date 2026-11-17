/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package observe

import (
	"errors"
	"testing"
)

// TestMutationObservationChangesDigest proves stream tampering is detected.
func TestMutationObservationChangesDigest(t *testing.T) {
	t.Parallel()
	base := []Observation{
		{At: 1, Seq: 1, ProbeID: "p-a", Value: NewNumber(42), Status: StatusOK},
		{At: 2, Seq: 2, ProbeID: "p-b", Value: NewText("ok"), Status: StatusOK},
	}
	baseline := StreamDigest(base)
	copyOf := func() []Observation {
		return append([]Observation(nil), base...)
	}
	valueMut := copyOf()
	valueMut[0].Value = NewNumber(43)
	seqMut := copyOf()
	seqMut[1].Seq = 99
	atMut := copyOf()
	atMut[0].At = 99
	idMut := copyOf()
	idMut[1].ProbeID = "p-x"
	statusMut := copyOf()
	statusMut[0].Status = StatusFailed
	for i, m := range [][]Observation{valueMut, seqMut, atMut, idMut, statusMut} {
		if StreamDigest(m) == baseline {
			t.Fatalf("mutation %d was not detected", i)
		}
	}
}

// TestMutationSinkFailureInjectsMeta verifies isolation under a mutated sink.
func TestMutationSinkFailureInjectsMeta(t *testing.T) {
	t.Parallel()
	p := NewPipeline()
	if err := p.Register(&fakeProbe{
		id: "p-a", kind: ProbeHTTP, value: NewBoolean(true), status: StatusOK,
	}); err != nil {
		t.Fatalf("Register err = %v", err)
	}
	healthy := &fakeSink{}
	if err := p.Attach(healthy); err != nil {
		t.Fatalf("Attach err = %v", err)
	}
	clean := StreamDigest(p.Collect(1, 1))
	if err := p.Attach(&fakeSink{fail: errors.New("mutated sink")}); err != nil {
		t.Fatalf("Attach err = %v", err)
	}
	broken := p.Collect(1, 1)
	if len(broken) != 2 {
		t.Fatalf("len = %d, want 2 (sample plus meta)", len(broken))
	}
	if StreamDigest(broken) == clean {
		t.Fatalf("sink mutation did not change the stream")
	}
	if len(healthy.seen) != 2 {
		t.Fatalf("healthy sink was blocked: %d observations", len(healthy.seen))
	}
}
