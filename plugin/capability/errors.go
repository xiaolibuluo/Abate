/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package capability

import "errors"

// Sentinel errors returned by the capability sandbox. Callers must test them
// with errors.Is.
var (
	// ErrCapabilityDenied reports that a host-function call lacked the required
	// capability.
	ErrCapabilityDenied = errors.New("plugin: capability denied")

	// ErrUnknownCapability reports that a capability declaration named a
	// capability that is not in the known whitelist.
	ErrUnknownCapability = errors.New("plugin: unknown capability")
)
