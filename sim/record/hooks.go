/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package record

// ReplayHooks receives each frame during a replay so a re-enactment engine can
// rebuild simulator state from the recorded stream in total order.
type ReplayHooks interface {
	// OnFrame is called once per archived frame, in total order.
	OnFrame(f Frame) error
}
