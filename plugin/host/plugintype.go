/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package host

import "abate/plugin/signature"

// PluginType re-exports signature.PluginType so the host API is self-contained.
type PluginType = signature.PluginType

// The plugin type constants re-export the canonical signature values so callers can
// express a plugin's kind through the host API without reaching into signature.
const (
	// PluginTypeOperator identifies an operator plugin.
	PluginTypeOperator = signature.PluginTypeOperator
	// PluginTypeReporter identifies a reporter plugin.
	PluginTypeReporter = signature.PluginTypeReporter
	// PluginTypeAdapter identifies an adapter plugin.
	PluginTypeAdapter = signature.PluginTypeAdapter
)
