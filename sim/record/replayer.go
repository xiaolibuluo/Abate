/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package record

import (
	"fmt"

	"abate/sim/clock"
)

// Replayer replays a sealed archive frame by frame and verifies its integrity.
type Replayer struct{}

// Replay verifies the archive and invokes the hooks for every frame in order, so a
// re-enactment engine can consume the recorded stream without re-reading the wire.
func (rp Replayer) Replay(arch *Archive, hooks ReplayHooks) error {
	if err := rp.Verify(arch); err != nil {
		return err
	}
	if hooks == nil {
		return nil
	}
	for i := range arch.entries {
		if err := hooks.OnFrame(arch.entries[i].frame); err != nil {
			return err
		}
	}
	return nil
}

// Verify checks the archive integrity, total-order sequence continuity and
// per-kind virtual-time monotonicity before any frame is handed to a consumer.
func (Replayer) Verify(arch *Archive) error {
	if err := arch.Verify(); err != nil {
		return err
	}
	if len(arch.entries) == 0 {
		return nil
	}
	prevByKind := make(map[FrameKind]clock.Instant)
	for i, e := range arch.entries {
		if e.frame.Seq != uint64(i)+1 { //nolint:gosec
			return fmt.Errorf("record: frame sequence discontinuity at index %d", i)
		}
		if e.frame.At < prevByKind[e.frame.Kind] {
			return fmt.Errorf("record: virtual time not monotonic for kind %s", e.frame.Kind)
		}
		prevByKind[e.frame.Kind] = e.frame.At
	}
	return nil
}
