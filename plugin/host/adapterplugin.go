/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package host

// AdapterPlugin bridges host events to an external sink. Host.Call dispatches
// EntryDetect and EntryEmit to this interface, gated on CapObserve.
type AdapterPlugin interface {
	Plugin

	// Detect inspects an environment and returns the matched adapter id plus
	// whether detection succeeded.
	Detect(env map[string]string) (string, bool)

	// Emit pushes a payload to the external sink.
	Emit(payload []byte) error
}
