/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package host

import (
	"context"
	"errors"
	"unicode/utf8"

	"abate/kernel/algebra"
	"abate/kernel/mutation"
	"abate/lang/rule"
	"abate/plugin/report"
)

// errTestTypeMismatch is returned by the test factories when a manifest kind
// disagrees with the factory. It keeps the fail-closed contract observable.
var errTestTypeMismatch = errors.New("host test: manifest type mismatch")

// errTestPayload is returned by the test adapter for an invalid payload.
var errTestPayload = errors.New("host test: invalid adapter payload")

// testOperator is a minimal, deterministic OperatorPlugin used by the host
// package tests. It lives in the internal test package so the tests never
// import abate/plugin/sample, which would create a test import cycle.
type testOperator struct {
	manifest Manifest
}

// Manifest returns the signed metadata of the plugin.
func (p *testOperator) Manifest() Manifest { return p.manifest }

// Type returns the plugin kind.
func (*testOperator) Type() PluginType { return PluginTypeOperator }

// Close releases no resources.
func (*testOperator) Close() error { return nil }

// Describe returns a stable algebra operator descriptor.
func (*testOperator) Describe() algebra.Operator {
	return algebra.NewBasic(
		algebra.FormatOperatorID(algebra.FamilyMutation, "host-test", "1"),
		algebra.FamilyMutation,
		algebra.Params{"algo": "copy"},
		[]algebra.TargetKey{"plugin/host"},
		false,
	)
}

// Apply copies the input and appends one byte derived from the injected virtual
// clock. It never draws from the RNG, so repeated calls on one instance are
// byte-identical and the deterministic assertions remain sound.
func (*testOperator) Apply(ctx context.Context, input []byte) ([]byte, error) {
	out := make([]byte, len(input)+1)
	copy(out, input)
	if be := Backend(ctx); be != nil && be.Clock != nil {
		out[len(input)] = byte(int64(be.Clock.Now()) & 0xff)
	}
	return out, nil
}

// testReporter is a minimal ReporterPlugin that delegates to the deterministic
// SARIF encoder so host-side EntryBuild tests can validate the artifact.
type testReporter struct {
	manifest Manifest
}

// Manifest returns the signed metadata of the plugin.
func (p *testReporter) Manifest() Manifest { return p.manifest }

// Type returns the plugin kind.
func (*testReporter) Type() PluginType { return PluginTypeReporter }

// Close releases no resources.
func (*testReporter) Close() error { return nil }

// Build renders findings into a SARIF 2.1.0 document.
func (*testReporter) Build(findings []rule.Finding, _ mutation.Report) ([]byte, error) {
	return report.BuildSARIF(findings)
}

// testAdapter is a minimal, air-gapped AdapterPlugin for host dispatch tests.
type testAdapter struct {
	manifest Manifest
}

// Manifest returns the signed metadata of the plugin.
func (p *testAdapter) Manifest() Manifest { return p.manifest }

// Type returns the plugin kind.
func (*testAdapter) Type() PluginType { return PluginTypeAdapter }

// Close releases no resources.
func (*testAdapter) Close() error { return nil }

// Detect matches the canonical CI environment marker.
func (*testAdapter) Detect(env map[string]string) (string, bool) {
	if env["CI"] == "true" {
		return "adp-ci", true
	}
	return "", false
}

// Emit validates the payload locally and never performs I/O.
func (*testAdapter) Emit(payload []byte) error {
	if len(payload) == 0 || !utf8.Valid(payload) {
		return errTestPayload
	}
	return nil
}

// testOperatorFactory instantiates the in-package test operator.
func testOperatorFactory(m Manifest) (Plugin, error) {
	if m.Type != PluginTypeOperator {
		return nil, errTestTypeMismatch
	}
	return &testOperator{manifest: m}, nil
}

// testReporterFactory instantiates the in-package test reporter.
func testReporterFactory(m Manifest) (Plugin, error) {
	if m.Type != PluginTypeReporter {
		return nil, errTestTypeMismatch
	}
	return &testReporter{manifest: m}, nil
}

// testAdapterFactory instantiates the in-package test adapter.
func testAdapterFactory(m Manifest) (Plugin, error) {
	if m.Type != PluginTypeAdapter {
		return nil, errTestTypeMismatch
	}
	return &testAdapter{manifest: m}, nil
}
