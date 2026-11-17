/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package record

// Source is a read-only, non-intrusive provider of frames for one category of the
// experiment stream. The recorder pulls from every subscribed source once at Stop,
// so attaching a source never alters the event flow being recorded.
type Source interface {
	// Collect returns the frames contributed by this source in deterministic order.
	Collect() []Frame
}
