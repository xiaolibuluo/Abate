/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package hotswap

import "time"

// DrainTimeout is the one and only wall-clock (non-deterministic) boundary in
// the plugin runtime. It bounds how long the DRAIN phase waits for real,
// concurrently executing goroutines that hold in-flight calls on the old
// plugin instance.
//
// A virtual clock cannot be used here: virtual time only advances when the
// simulator explicitly advances it, whereas an actual blocked goroutine can
// only be waited on with a real timer. Without a wall-clock bound, a plugin
// that never returns from a call would stall the pipeline forever, which
// violates the bounded commit-to-publish SLA. The timeout never influences
// plugin output: SWITCH has already completed and the new instance serves all
// subsequent calls, so experiment results stay deterministic. Tests may lower
// DrainTimeout to exercise the timeout path without waiting 30 seconds.
var DrainTimeout = 30 * time.Second
