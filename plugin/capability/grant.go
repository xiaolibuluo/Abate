/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package capability

import (
	"sort"
	"strings"
)

// whitelist maps the canonical capability name to its bit. It is the single
// source of truth for accepted names.
var whitelist = map[string]Capability{
	"clock":   CapClock,
	"rand":    CapRand,
	"observe": CapObserve,
	"fs":      CapFS,
	"net":     CapNet,
}

// sortedNames returns the whitelist names in deterministic order. Any code that
// needs to iterate the whitelist must use this to avoid map-order leakage.
func sortedNames() []string {
	names := make([]string, 0, len(whitelist))
	for n := range whitelist {
		names = append(names, n)
	}
	sort.Strings(names)
	return names
}

// KnownCapabilities returns the canonical capability names in deterministic
// (sorted) order. It is the authoritative list accepted by ParseGrant.
func KnownCapabilities() []string {
	return sortedNames()
}

// ParseGrant parses a comma-separated capability declaration such as
// "clock,rand" into a Capabilities bitmask.
//
// Rules:
//   - An empty or whitespace-only declaration yields a zero (all-deny) mask with
//     no error.
//   - Whitespace around names is trimmed and empty tokens are ignored, so
//     "clock, rand" and "clock," are accepted.
//   - Any unknown capability name rejects the entire declaration with
//     ErrUnknownCapability (fail-closed, forward-compatible whitelist).
func ParseGrant(decl string) (Capabilities, error) {
	var caps Capabilities
	if strings.TrimSpace(decl) == "" {
		return caps, nil
	}
	for _, raw := range strings.Split(decl, ",") {
		name := strings.TrimSpace(raw)
		if name == "" {
			continue
		}
		bit, ok := whitelist[name]
		if !ok {
			return 0, ErrUnknownCapability
		}
		caps.Grant(bit)
	}
	return caps, nil
}
