/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package schema

import (
	"bytes"
	"encoding/json"
	"errors"
	"testing"

	"abate/internal/yamlmini"
)

// sampleExperiment builds a fully populated CES document for round-trip tests.
func sampleExperiment() Experiment {
	return Experiment{
		APIVersion: APIVersion,
		Kind:       KindExperiment,
		Metadata: Metadata{
			Name:        "网络延迟实验",
			Namespace:   "default",
			Labels:      map[string]string{"team": "内核组", "env": "prod"},
			Annotations: map[string]string{"note": "特殊字符 <>&\"'\n\t"},
		},
		Spec: ExperimentSpec{
			Target: Target{
				Kind:     TargetSimulator,
				Selector: map[string]string{"node": "n-1", "process": "api"},
			},
			SteadyState: []SteadyState{{
				ProbeRef:    "latency",
				Expectation: Expectation{Type: ExpectValue, Value: "42"},
				Tolerance:   Tolerance{Type: TolerancePercent, Value: 5},
				HoldFor:     30 * Second,
			}},
			Method: Method{
				Expression: "operator://c-delay-v1 ∘ ε",
				Operators: []OperatorDef{{
					ID:         "c-delay-v1",
					Family:     FamilyChaos,
					Params:     map[string]string{"ms": "100"},
					Targets:    []string{"n-1"},
					Invertible: true,
					TargetKind: TargetSimulator,
				}},
			},
			Probes: []Probe{{
				ID:       "latency",
				Kind:     ProbeHTTP,
				Spec:     ProbeSpec{URL: "http://sim/health", Method: "GET", ExpectStatus: 200},
				Interval: Second,
			}},
			Rollback: Rollback{Strategy: StrategyAuto, VerifySteadyState: true},
			Schedule: Schedule{Once: &OnceSpec{At: 5 * Second}},
			Timeout:  Hour,
			Seed:     12345,
		},
	}
}

// TestParseDurationUnits covers every unit suffix and composed literals.
func TestParseDurationUnits(t *testing.T) {
	t.Parallel()
	cases := []struct {
		lit  string
		want VirtualDuration
	}{
		{"0", 0},
		{"1000000000", Second},
		{"1ns", Nanosecond},
		{"1us", Microsecond},
		{"1µs", Microsecond},
		{"1ms", Millisecond},
		{"1s", Second},
		{"1m", Minute},
		{"1h", Hour},
		{"1h30m", Hour + 30*Minute},
		{"250ms", 250 * Millisecond},
		{"1s500ms", Second + 500*Millisecond},
		{"-5s", -5 * Second},
	}
	for _, c := range cases {
		got, err := ParseDuration(c.lit)
		if err != nil {
			t.Fatalf("ParseDuration(%q) err = %v", c.lit, err)
		}
		if got != c.want {
			t.Fatalf("ParseDuration(%q) = %d, want %d", c.lit, got, c.want)
		}
	}
}

// TestParseDurationErrors covers malformed and overflowing literals.
func TestParseDurationErrors(t *testing.T) {
	t.Parallel()
	bad := []string{
		"", "-", "abc", "5x", "1h30", "999999999999999999999",
		"999999999999999999999h", "1ss", "1.5s", "9223372036854775807ns1h",
	}
	for _, s := range bad {
		if _, err := ParseDuration(s); err == nil {
			t.Fatalf("ParseDuration(%q) must fail", s)
		}
	}
}

// TestVirtualDurationString checks canonical literal rendering.
func TestVirtualDurationString(t *testing.T) {
	t.Parallel()
	cases := []struct {
		d    VirtualDuration
		want string
	}{
		{0, "0h"},
		{1, "1ns"},
		{1000, "1us"},
		{1500, "1500ns"},
		{Second, "1s"},
		{90 * Second, "90s"},
		{1500 * Millisecond, "1500ms"},
		{-5 * Second, "-5s"},
		{Hour, "1h"},
	}
	for _, c := range cases {
		if got := c.d.String(); got != c.want {
			t.Fatalf("(%d).String() = %q, want %q", c.d, got, c.want)
		}
	}
}

// TestVirtualDurationTextRoundTrip checks the TextMarshaler contract.
func TestVirtualDurationTextRoundTrip(t *testing.T) {
	t.Parallel()
	values := []VirtualDuration{
		0, 1, 999, Microsecond, Second, -Second, Hour + 30*Minute, 123456789,
	}
	for _, d := range values {
		text, err := d.MarshalText()
		if err != nil {
			t.Fatalf("MarshalText(%d) err = %v", d, err)
		}
		var back VirtualDuration
		if err := back.UnmarshalText(text); err != nil {
			t.Fatalf("UnmarshalText(%q) err = %v", text, err)
		}
		if back != d {
			t.Fatalf("round trip %d -> %q -> %d", d, text, back)
		}
		if d.Nanoseconds() != int64(d) {
			t.Fatalf("Nanoseconds() mismatch for %d", d)
		}
	}
	var bad VirtualDuration
	if err := bad.UnmarshalText([]byte("nope")); err == nil {
		t.Fatalf("UnmarshalText(nope) must fail")
	}
}

// TestJSONYAMLRoundTrip proves byte-stable JSON/YAML interop.
func TestJSONYAMLRoundTrip(t *testing.T) {
	t.Parallel()
	exp := sampleExperiment()
	first, err := json.Marshal(exp)
	if err != nil {
		t.Fatalf("json.Marshal err = %v", err)
	}
	// B3 (ADR-YAML-001): JSON interop goes through encoding/json, not the
	// YAML subset decoder (which intentionally does not parse flow-style JSON).
	var viaYAML Experiment
	if err := json.Unmarshal(first, &viaYAML); err != nil {
		t.Fatalf("json.Unmarshal(json) err = %v", err)
	}
	second, err := json.Marshal(viaYAML)
	if err != nil {
		t.Fatalf("json.Marshal(viaYAML) err = %v", err)
	}
	if !bytes.Equal(first, second) {
		t.Fatalf("JSON->YAML->JSON drift:\n%s\n%s", first, second)
	}
	y, err := yamlmini.Marshal(exp)
	if err != nil {
		t.Fatalf("yamlmini.Marshal err = %v", err)
	}
	var back Experiment
	if err := yamlmini.Unmarshal(y, &back); err != nil {
		t.Fatalf("yamlmini.Unmarshal(yaml) err = %v", err)
	}
	third, err := json.Marshal(back)
	if err != nil {
		t.Fatalf("json.Marshal(back) err = %v", err)
	}
	if !bytes.Equal(first, third) {
		t.Fatalf("YAML round trip drift:\n%s\n%s", first, third)
	}
}

// TestVersionNegotiate covers supported, empty and unknown version requests.
func TestVersionNegotiate(t *testing.T) {
	t.Parallel()
	v := NewVersion()
	if v.Current() != APIVersion {
		t.Fatalf("Current = %s", v.Current())
	}
	got, err := v.Negotiate("")
	if err != nil || got != APIVersion {
		t.Fatalf("Negotiate(empty) = (%s, %v)", got, err)
	}
	got, err = v.Negotiate(APIVersion)
	if err != nil || got != APIVersion {
		t.Fatalf("Negotiate(v1) = (%s, %v)", got, err)
	}
	if _, err := v.Negotiate("ces.tencent.com/v2"); !errors.Is(err, ErrUnsupportedVersion) {
		t.Fatalf("Negotiate(v2) err = %v, want ErrUnsupportedVersion", err)
	}
	if len(v.SupportedVersions()) != 1 || v.SupportedVersions()[0] != APIVersion {
		t.Fatalf("SupportedVersions = %v", v.SupportedVersions())
	}
}
