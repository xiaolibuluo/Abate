/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package adapter

import (
	"errors"
	"fmt"

	"abate/plugin/host"
)

// ErrM5WrongPluginType reports a manifest that is not an adapter manifest.
var ErrM5WrongPluginType = errors.New("adapter: manifest is not an adapter")

// Compile-time proofs that both factories match the host.Factory signature.
var (
	_ host.Factory = ITSFactory
	_ host.Factory = MSGFactory
)

// ITSFactory instantiates the ADP-ITS adapter from a verified manifest.
//
// It fails closed when the manifest type is not an adapter, mirroring the
// sample.AdapterFactory contract.
func ITSFactory(m host.Manifest) (host.Plugin, error) {
	if m.Type != host.PluginTypeAdapter {
		return nil, fmt.Errorf("%w: manifest %q has type %q",
			ErrM5WrongPluginType, m.ID, m.Type)
	}
	return newITSAdapter(m), nil
}

// MSGFactory instantiates the ADP-MSG adapter from a verified manifest.
//
// It fails closed when the manifest type is not an adapter, mirroring the
// sample.AdapterFactory contract.
func MSGFactory(m host.Manifest) (host.Plugin, error) {
	if m.Type != host.PluginTypeAdapter {
		return nil, fmt.Errorf("%w: manifest %q has type %q",
			ErrM5WrongPluginType, m.ID, m.Type)
	}
	return newMSGAdapter(m), nil
}
