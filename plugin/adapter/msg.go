/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package adapter

import "abate/plugin/host"

// msgAdapterID is the stable identifier returned by MSG detection.
const msgAdapterID = "adp-msg"

// msgProviderKeys is the ordered list of injected environment markers that
// enable the message adapter. Detection never reads process state.
var msgProviderKeys = []string{"ABATE_ADP_MSG", "MSG_ENABLED", "MESSAGE_SINK"}

// msgRequiredFields is the event payload assertion contract.
var msgRequiredFields = []string{"kind", "level", "summary"}

// msgAdapter is the air-gapped ADP-MSG adapter.
//
// Detect consults only the caller-injected environment map and Emit performs
// only a local payload assertion. No external service is ever contacted.
type msgAdapter struct {
	manifest  host.Manifest
	assertion PayloadAssertion
}

// Compile-time proof that msgAdapter satisfies the host contract.
var _ host.AdapterPlugin = (*msgAdapter)(nil)

// newMSGAdapter constructs a MSG adapter from its verified manifest.
func newMSGAdapter(m host.Manifest) *msgAdapter {
	return &msgAdapter{
		manifest:  m,
		assertion: NewPayloadAssertion(msgRequiredFields),
	}
}

// Manifest returns the signed metadata of the plugin.
func (a *msgAdapter) Manifest() host.Manifest {
	return a.manifest
}

// Type returns the plugin kind.
func (*msgAdapter) Type() host.PluginType {
	return host.PluginTypeAdapter
}

// Close releases no resources for this in-memory adapter.
func (*msgAdapter) Close() error {
	return nil
}

// Detect inspects the injected environment map and returns a stable adapter id.
func (*msgAdapter) Detect(env map[string]string) (string, bool) {
	for _, key := range msgProviderKeys {
		value, present := env[key]
		if !present {
			continue
		}
		if value == "true" || value == "msg" {
			return msgAdapterID, true
		}
	}
	return "", false
}

// Emit validates an event payload locally and performs no I/O.
func (a *msgAdapter) Emit(payload []byte) error {
	return a.assertion.Assert(payload)
}
