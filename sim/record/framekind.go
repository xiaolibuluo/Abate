/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package record

// FrameKind classifies a recorded frame within the experiment event stream.
type FrameKind string

// Frame kind constants classify the recorded event categories.
const (
	KindSched FrameKind = "SCHED"
	KindObs   FrameKind = "OBS"
	KindRNG   FrameKind = "RNG"
	KindClk   FrameKind = "CLK"
	KindFault FrameKind = "FAULT"
	KindNet   FrameKind = "NET"
	KindFS    FrameKind = "FS"
	KindVTask FrameKind = "VTASK"
)

// frameKindOrder fixes the byte tag assigned to each kind so the on-disk layout
// is stable and independent of declaration or map traversal order.
var frameKindOrder = []FrameKind{
	KindSched, KindObs, KindRNG, KindClk,
	KindFault, KindNet, KindFS, KindVTask,
}

// String returns the canonical name of the frame kind.
func (k FrameKind) String() string { return string(k) }

// Marshal returns the single-byte deterministic tag for the frame kind.
func (k FrameKind) Marshal() []byte {
	for i, candidate := range frameKindOrder {
		if candidate == k {
			return []byte{byte(i)} //nolint:gosec // the index is bounded by the fixed table length.
		}
	}
	return []byte{0xFF}
}

// FrameKindFromTag decodes the byte tag produced by Marshal, reporting success.
func FrameKindFromTag(tag byte) (FrameKind, bool) {
	if int(tag) >= len(frameKindOrder) {
		return "", false
	}
	return frameKindOrder[tag], true
}
