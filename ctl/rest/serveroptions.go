/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package rest

// ServerOptions customizes a server beyond its configuration.
type ServerOptions struct {
	// Runner overrides the default in-memory check runner.
	Runner TaskRunner
	// Plugins overrides the default static plugin manager.
	Plugins PluginManager
}
