/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package hotswap

// SwapState is a phase of the hot-swap state machine. The zero value is PREPARE,
// so a freshly constructed transaction begins in the validating phase.
type SwapState int

// The four frozen M4 hot-swap phases. Names are kept verbatim from the milestone
// contract so control-plane logs and evidence files stay comparable.
//
//nolint:revive // ALL_CAPS phase names are frozen by the M4 contract.
const (
	// PREPARE validates identities, registries and the allowed version step.
	PREPARE SwapState = iota
	// SWITCH marks the old instance superseded and atomically replaces it.
	SWITCH
	// DRAIN waits for in-flight calls on the old instance to finish.
	DRAIN
	// UNLOAD closes the old plugin after its calls have drained.
	UNLOAD
)

// String renders the phase name. Unknown values render as "UNKNOWN" so a
// corrupted state can never masquerade as a valid phase.
func (s SwapState) String() string {
	switch s {
	case PREPARE:
		return "PREPARE"
	case SWITCH:
		return "SWITCH"
	case DRAIN:
		return "DRAIN"
	case UNLOAD:
		return "UNLOAD"
	default:
		return "UNKNOWN"
	}
}
