/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package record

import "abate/sim/clock"

// Frame is one totally ordered record of the experiment stream: a scheduler
// sequence, the virtual instant it occurred, its category and a versioned payload.
type Frame struct {
	Seq     uint64
	At      clock.Instant
	Kind    FrameKind
	Payload []byte
}
