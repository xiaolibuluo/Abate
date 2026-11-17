/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package observe

import "errors"

// Sentinel errors of the observation pipeline.
var (
	// ErrNilProbe reports a nil probe registration.
	ErrNilProbe = errors.New("observe: probe must not be nil")

	// ErrNilSink reports a nil sink attachment.
	ErrNilSink = errors.New("observe: sink must not be nil")

	// ErrInvalidProbeID reports an empty probe identifier.
	ErrInvalidProbeID = errors.New("observe: probe ID must not be empty")

	// ErrDuplicateProbe reports a probe ID registered twice.
	ErrDuplicateProbe = errors.New("observe: duplicate probe ID")
)
