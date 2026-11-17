/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package adapter

import "abate/plugin/host"

// itsAdapterID is the stable identifier returned by ITS detection.
const itsAdapterID = "adp-its"

// itsProviderKeys is the ordered list of injected environment markers that
// enable the issue-tracking adapter. Detection never reads process state.
var itsProviderKeys = []string{"ABATE_ADP_ITS", "ITS_ENABLED", "ISSUE_TRACKER"}

// itsRequiredFields is the issue payload assertion contract.
var itsRequiredFields = []string{"type", "fingerprint", "action", "labels"}

// itsAdapter is the air-gapped ADP-ITS adapter.
//
// Detect consults only the caller-injected environment map and Emit performs
// only a local payload assertion. No external service is ever contacted.
type itsAdapter struct {
	manifest  host.Manifest
	assertion PayloadAssertion
}

// Compile-time proof that itsAdapter satisfies the host contract.
var _ host.AdapterPlugin = (*itsAdapter)(nil)

// newITSAdapter constructs an ITS adapter from its verified manifest.
func newITSAdapter(m host.Manifest) *itsAdapter {
	return &itsAdapter{
		manifest:  m,
		assertion: NewPayloadAssertion(itsRequiredFields),
	}
}

// Manifest returns the signed metadata of the plugin.
func (a *itsAdapter) Manifest() host.Manifest {
	return a.manifest
}

// Type returns the plugin kind.
func (*itsAdapter) Type() host.PluginType {
	return host.PluginTypeAdapter
}

// Close releases no resources for this in-memory adapter.
func (*itsAdapter) Close() error {
	return nil
}

// Detect inspects the injected environment map and returns a stable adapter id.
func (*itsAdapter) Detect(env map[string]string) (string, bool) {
	for _, key := range itsProviderKeys {
		value, present := env[key]
		if !present {
			continue
		}
		if value == "true" || value == "its" {
			return itsAdapterID, true
		}
	}
	return "", false
}

// Emit validates an issue payload locally and performs no I/O.
func (a *itsAdapter) Emit(payload []byte) error {
	return a.assertion.Assert(payload)
}
