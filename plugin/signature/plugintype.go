/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package signature

// PluginType is the kind of a plugin.
type PluginType string

// Known plugin types accepted by the runtime.
const (
	// PluginTypeOperator identifies an operator plugin.
	PluginTypeOperator PluginType = "operator"

	// PluginTypeReporter identifies a reporter plugin.
	PluginTypeReporter PluginType = "reporter"

	// PluginTypeAdapter identifies an adapter plugin.
	PluginTypeAdapter PluginType = "adapter"
)
