/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package clock

// State is the four-state clock mode derived from the virtual clock fields.
//
// The states are mutually exclusive snapshots of the clock configuration:
// Running (nominal), Fast (rate ≠ 1), Freeze (advances rejected), and Offset
// (Now readings shifted by a chaos-injected offset).
type State int

const (
	// StateRunning is the nominal clock with rate 1 and no offset.
	StateRunning State = iota

	// StateFast is a clock whose virtual-time rate differs from 1.
	StateFast

	// StateFreeze is a frozen clock that rejects advances.
	StateFreeze

	// StateOffset is a clock whose Now readings are shifted by an offset.
	StateOffset
)

// String renders the state as a stable uppercase identifier.
func (s State) String() string {
	switch s {
	case StateRunning:
		return "RUNNING"
	case StateFast:
		return "FAST"
	case StateFreeze:
		return "FREEZE"
	case StateOffset:
		return "OFFSET"
	default:
		return "UNKNOWN"
	}
}
