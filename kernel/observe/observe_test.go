/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package observe

import (
	"errors"
	"math"
	"strings"
	"testing"
)

// fakeProbe is a deterministic in-test probe.
type fakeProbe struct {
	id     string
	kind   string
	value  Value
	status string
}

// ID returns the probe identifier.
func (f *fakeProbe) ID() string { return f.id }

// Kind returns the probe kind.
func (f *fakeProbe) Kind() string { return f.kind }

// Collect returns the configured value and status.
func (f *fakeProbe) Collect(int64) (Value, string) { return f.value, f.status }

// fakeSink records observations or fails on demand.
type fakeSink struct {
	seen []Observation
	fail error
}

// Observe records an observation unless the sink is configured to fail.
func (f *fakeSink) Observe(o Observation) error {
	if f.fail != nil {
		return f.fail
	}
	f.seen = append(f.seen, o)
	return nil
}

// TestValueConstructors covers all four tagged value kinds.
func TestValueConstructors(t *testing.T) {
	t.Parallel()
	values := []Value{
		NewNumber(3.5),
		NewBoolean(true),
		NewText("中文\"'\n"),
		NewHistogram([]uint64{1, 2, 3}),
	}
	want := []string{ValueNumber, ValueBoolean, ValueText, ValueHistogram}
	for i, v := range values {
		if v.Kind != want[i] {
			t.Fatalf("value %d kind = %s, want %s", i, v.Kind, want[i])
		}
		if !v.Equal(values[i]) || v.Equal(Value{}) {
			t.Fatalf("Equal broken for value %d", i)
		}
	}
}

// TestValueRoundTrip checks lossless encode/decode for all four kinds.
func TestValueRoundTrip(t *testing.T) {
	t.Parallel()
	values := []Value{
		NewNumber(1e18),
		NewNumber(-0.000001),
		NewBoolean(false),
		NewText("特殊字符 <>&\"'\n\t中文"),
		NewHistogram([]uint64{0, 999, 1 << 40}),
	}
	for i, v := range values {
		data, err := v.Encode()
		if err != nil {
			t.Fatalf("Encode %d err = %v", i, err)
		}
		back, err := DecodeValue(data)
		if err != nil {
			t.Fatalf("Decode %d err = %v", i, err)
		}
		if !v.Equal(back) {
			t.Fatalf("round trip %d mismatch: %+v vs %+v", i, v, back)
		}
	}
}

// TestValueDecodeError checks malformed input is rejected.
func TestValueDecodeError(t *testing.T) {
	t.Parallel()
	if _, err := DecodeValue([]byte("{not json")); err == nil {
		t.Fatalf("DecodeValue must reject malformed JSON")
	}
}

// TestValueHistogramCopies proves the constructor does not alias the input.
func TestValueHistogramCopies(t *testing.T) {
	t.Parallel()
	buckets := []uint64{1, 2, 3}
	v := NewHistogram(buckets)
	buckets[0] = 99
	if v.Buckets[0] != 1 {
		t.Fatalf("NewHistogram aliased the caller slice")
	}
}

// TestObservationEncodeStable checks byte-stable encoding.
func TestObservationEncodeStable(t *testing.T) {
	t.Parallel()
	o := Observation{At: 7, Seq: 3, ProbeID: "p-a", Value: NewText("ok"), Status: StatusOK}
	first, err := o.Encode()
	if err != nil {
		t.Fatalf("Encode err = %v", err)
	}
	second, err := o.Encode()
	if err != nil {
		t.Fatalf("Encode err = %v", err)
	}
	if string(first) != string(second) {
		t.Fatalf("encoding is not stable")
	}
}

// TestPipelineRegister covers the registration error branches.
func TestPipelineRegister(t *testing.T) {
	t.Parallel()
	var p Pipeline
	if err := p.Register(nil); !errors.Is(err, ErrNilProbe) {
		t.Fatalf("nil probe err = %v", err)
	}
	if err := p.Register(&fakeProbe{id: ""}); !errors.Is(err, ErrInvalidProbeID) {
		t.Fatalf("empty ID err = %v", err)
	}
	probe := &fakeProbe{id: "p-a", kind: ProbeMetric, value: NewNumber(1), status: StatusOK}
	if err := p.Register(probe); err != nil {
		t.Fatalf("first Register err = %v", err)
	}
	if err := p.Register(probe); !errors.Is(err, ErrDuplicateProbe) {
		t.Fatalf("duplicate err = %v", err)
	}
}

// TestPipelineAttach covers the sink attachment error branch.
func TestPipelineAttach(t *testing.T) {
	t.Parallel()
	p := NewPipeline()
	if err := p.Attach(nil); !errors.Is(err, ErrNilSink) {
		t.Fatalf("nil sink err = %v", err)
	}
	if err := p.Attach(&fakeSink{}); err != nil {
		t.Fatalf("Attach err = %v", err)
	}
}

// TestPipelineCollectOrder verifies stable ID order and sequence assignment.
func TestPipelineCollectOrder(t *testing.T) {
	t.Parallel()
	p := NewPipeline()
	for _, id := range []string{"p-c", "p-a", "p-b"} {
		err := p.Register(&fakeProbe{
			id: id, kind: ProbeMetric, value: NewText(id), status: StatusOK,
		})
		if err != nil {
			t.Fatalf("Register(%s) err = %v", id, err)
		}
	}
	obs := p.Collect(42, 10)
	if len(obs) != 3 {
		t.Fatalf("len = %d, want 3", len(obs))
	}
	wantSeq := uint64(10)
	for i, want := range []string{"p-a", "p-b", "p-c"} {
		if obs[i].ProbeID != want || obs[i].Seq != wantSeq || obs[i].At != 42 {
			t.Fatalf("obs[%d] = %+v", i, obs[i])
		}
		wantSeq++
	}
	if obs[0].Value.Text != "p-a" || obs[0].Status != StatusOK {
		t.Fatalf("value/status not carried: %+v", obs[0])
	}
}

// TestPipelineSinkFailureIsolation proves one bad sink cannot block others.
func TestPipelineSinkFailureIsolation(t *testing.T) {
	t.Parallel()
	boom := errors.New("sink down")
	p := NewPipeline()
	if err := p.Register(&fakeProbe{
		id: "p-a", kind: ProbeHTTP, value: NewNumber(1), status: StatusOK,
	}); err != nil {
		t.Fatalf("Register err = %v", err)
	}
	first := &fakeSink{}
	third := &fakeSink{}
	if err := p.Attach(first); err != nil {
		t.Fatalf("Attach err = %v", err)
	}
	if err := p.Attach(&fakeSink{fail: boom}); err != nil {
		t.Fatalf("Attach err = %v", err)
	}
	if err := p.Attach(third); err != nil {
		t.Fatalf("Attach err = %v", err)
	}
	obs := p.Collect(10, 5)
	if len(obs) != 2 {
		t.Fatalf("len = %d, want 2 (sample + meta)", len(obs))
	}
	if obs[1].ProbeID != "meta/sink-error" || obs[1].Status != StatusFailed ||
		!strings.Contains(obs[1].Value.Text, "sink down") {
		t.Fatalf("meta observation = %+v", obs[1])
	}
	if len(first.seen) != 1 || len(third.seen) != 1 {
		t.Fatalf("sinks were blocked: first=%d third=%d", len(first.seen), len(third.seen))
	}
}

// TestPipelineTimeoutPassthrough verifies timeout status is preserved.
func TestPipelineTimeoutPassthrough(t *testing.T) {
	t.Parallel()
	p := NewPipeline()
	if err := p.Register(&fakeProbe{
		id: "p-t", kind: ProbeShell, value: NewText("late"), status: StatusTimeout,
	}); err != nil {
		t.Fatalf("Register err = %v", err)
	}
	obs := p.Collect(1, 1)
	if len(obs) != 1 || obs[0].Status != StatusTimeout {
		t.Fatalf("timeout status lost: %+v", obs)
	}
}

// TestStreamDigestDeterministic proves 100 runs hash identically.
func TestStreamDigestDeterministic(t *testing.T) {
	t.Parallel()
	batch := []Observation{
		{At: 1, Seq: 1, ProbeID: "p-a", Value: NewNumber(1.5), Status: StatusOK},
		{At: 2, Seq: 2, ProbeID: "p-b", Value: NewText("ok"), Status: StatusDegraded},
	}
	baseline := StreamDigest(batch)
	for i := 0; i < 100; i++ {
		if got := StreamDigest(batch); got != baseline {
			t.Fatalf("digest diverged at run %d", i)
		}
	}
	other := []Observation{
		{At: 1, Seq: 1, ProbeID: "p-a", Value: NewNumber(2.5), Status: StatusOK},
	}
	if StreamDigest(other) == baseline {
		t.Fatalf("different batch produced the same digest")
	}
}

// TestStreamDigestSkipsNaN covers the unencodable-value guard.
func TestStreamDigestSkipsNaN(t *testing.T) {
	t.Parallel()
	bad := Observation{At: 1, Seq: 1, ProbeID: "p-nan",
		Value: Value{Kind: ValueNumber, Number: math.NaN()}, Status: StatusFailed}
	good := Observation{At: 2, Seq: 2, ProbeID: "p-ok",
		Value: NewNumber(1), Status: StatusOK}
	if StreamDigest([]Observation{bad, good}) != StreamDigest([]Observation{good}) {
		t.Fatalf("unencodable record must be skipped, not hashed")
	}
}
