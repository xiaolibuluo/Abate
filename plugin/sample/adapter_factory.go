/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package sample

import (
	"fmt"

	"abate/plugin/adapter"
	"abate/plugin/host"
)

// AdapterFactory instantiates the ADP-CI adapter from a verified manifest. It fails
// closed if the manifest is not an adapter.
func AdapterFactory(m host.Manifest) (host.Plugin, error) {
	if m.Type != host.PluginTypeAdapter {
		return nil, fmt.Errorf("sample: manifest %q is not an adapter (type %q)", m.ID, m.Type)
	}
	return adapter.NewCIAdapter(m), nil
}
