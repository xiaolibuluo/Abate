/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package air

import (
	"errors"
	"fmt"
)

// SchemaVersion is the AIR archive version produced by this build.
//
// Version 2 adds the trailing extension section for unit-level annotations.
// Version 1 archives remain readable because the extension is purely additive.
const SchemaVersion uint16 = 2

// MinSchemaVersion is the oldest archive version this build still accepts.
const MinSchemaVersion uint16 = 1

// ErrUnsupportedAIRVersion is the sentinel reported by version negotiation.
//
// Every rejection wraps this sentinel, so callers can use errors.Is while the
// message still names both the writer and the reader version numbers.
var ErrUnsupportedAIRVersion = errors.New("air: unsupported schema version")

// Version describes the half-open range of AIR schema versions a producer or
// consumer understands.
//
// The negotiation algorithm is self-contained: a reader accepts a writer only
// when the writer's version lies inside [MinSupported, Current] and the
// writer's own minimum is not newer than the reader's current version. This
// mirrors the CES version negotiation idea without importing the CES layer,
// which would invert the L4 to L2 dependency direction.
type Version struct {
	min     uint16
	current uint16
}

// NewVersion returns the version range of this build.
func NewVersion() Version {
	return Version{min: MinSchemaVersion, current: SchemaVersion}
}

// NewVersionRange builds an explicit version range, used to emulate older or
// future readers in compatibility tests.
func NewVersionRange(minimum, current uint16) (Version, error) {
	if minimum == 0 || current == 0 || minimum > current {
		return Version{}, fmt.Errorf("%w: minimum %d, current %d",
			ErrInvalidVersionRange, minimum, current)
	}
	return Version{min: minimum, current: current}, nil
}

// Current returns the newest version this range understands.
func (v Version) Current() uint16 {
	return v.current
}

// MinSupported returns the oldest version this range understands.
func (v Version) MinSupported() uint16 {
	return v.min
}

// Negotiate resolves a requested writer version against this range.
//
// A zero request selects the current version. Any other request must lie in
// [MinSupported, Current], otherwise the sentinel is returned.
func (v Version) Negotiate(requested uint16) (uint16, error) {
	if requested == 0 {
		return v.current, nil
	}
	if requested < v.min || requested > v.current {
		return 0, fmt.Errorf("%w: writer version v%d, reader supports v%d..v%d",
			ErrUnsupportedAIRVersion, requested, v.min, v.current)
	}
	return requested, nil
}

// Accepts validates the two version numbers found in an archive header.
//
// Both sides are named in every rejection so an operator can tell whether the
// archive is too old, too new, or requires a minimum the reader cannot meet.
func (v Version) Accepts(writerMin, writerVersion uint16) error {
	reject := func(reason string) error {
		return fmt.Errorf("%w: %s (writer version v%d, writer minimum v%d; reader supports v%d..v%d)",
			ErrUnsupportedAIRVersion, reason, writerVersion, writerMin, v.min, v.current)
	}
	if writerVersion == 0 {
		return reject("zero writer version")
	}
	if writerMin == 0 {
		return reject("zero writer minimum")
	}
	if writerMin > writerVersion {
		return reject("writer minimum exceeds writer version")
	}
	if writerMin > v.current {
		return reject("writer minimum is newer than the reader")
	}
	if writerVersion > v.current {
		return reject("writer version is newer than the reader")
	}
	if writerVersion < v.min {
		return reject("writer version is older than the reader minimum")
	}
	return nil
}

// String renders the version range in a stable form.
func (v Version) String() string {
	return fmt.Sprintf("v%d..v%d", v.min, v.current)
}
