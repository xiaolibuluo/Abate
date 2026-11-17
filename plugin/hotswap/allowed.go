/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package hotswap

import (
	"abate/plugin/host"
	"abate/plugin/signature"
)

// Allowed reports whether a hot-swap from version from to version to is
// permitted for plugin type t.
//
// Policy (fail-closed, documented for every reviewer):
//   - Only the three frozen plugin kinds are swappable. An unknown type is
//     rejected, so a future plugin kind cannot silently bypass the version gate.
//   - The major version must be identical: ABI-breaking upgrades require a new
//     plugin id and a fresh Load, never an in-place swap.
//   - The target version must not be lower than the current version, i.e.
//     signature.Compare(from, to) <= 0. Downgrades and same-major rollbacks are
//     rejected because an experiment already observed the newer behaviour.
//
// The function is pure: it reads only its arguments and never touches global
// state, clocks or randomness.
func Allowed(t host.PluginType, from, to signature.SemVer) bool {
	switch t {
	case host.PluginTypeOperator, host.PluginTypeReporter, host.PluginTypeAdapter:
	default:
		return false
	}
	if from.Major != to.Major {
		return false
	}
	return signature.Compare(from, to) <= 0
}
