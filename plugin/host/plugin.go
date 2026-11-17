/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package host

// Plugin is the minimal contract every loaded plugin satisfies. A Plugin is
// opaque to the host except through its concrete sub-interface (OperatorPlugin,
// ReporterPlugin, or AdapterPlugin), which the host dispatches to at call time.
type Plugin interface {
	// Manifest returns the signed metadata of the plugin.
	Manifest() Manifest

	// Type returns the plugin kind (operator, reporter, or adapter).
	Type() PluginType

	// Close releases plugin resources. It is called by Host.Close and during
	// hot-swap drain (see plugin/hotswap).
	Close() error
}
