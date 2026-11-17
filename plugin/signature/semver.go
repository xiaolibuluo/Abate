/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package signature

import (
	"fmt"
	"strconv"
	"strings"
)

// SemVer is a semantic version with major, minor, and patch components.
type SemVer struct {
	Major uint
	Minor uint
	Patch uint
}

// ParseSemVer parses a "major.minor.patch" version string into a SemVer.
//
// Exactly three non-negative integer components separated by dots are required;
// surrounding whitespace is trimmed and missing or non-numeric components are
// rejected with a descriptive error.
func ParseSemVer(s string) (SemVer, error) {
	parts := strings.Split(s, ".")
	if len(parts) != 3 {
		return SemVer{}, fmt.Errorf("plugin: invalid semver %q: expected major.minor.patch", s)
	}
	var v SemVer
	targets := []*uint{&v.Major, &v.Minor, &v.Patch}
	for i, raw := range parts {
		n, err := strconv.ParseUint(strings.TrimSpace(raw), 10, 64)
		if err != nil {
			return SemVer{}, fmt.Errorf("plugin: invalid semver component %q: %w", raw, err)
		}
		*targets[i] = uint(n)
	}
	return v, nil
}

// Compare returns -1 if a < b, 0 if a == b, and 1 if a > b, comparing major,
// then minor, then patch.
func Compare(a, b SemVer) int {
	if a.Major != b.Major {
		return cmpUint(a.Major, b.Major)
	}
	if a.Minor != b.Minor {
		return cmpUint(a.Minor, b.Minor)
	}
	return cmpUint(a.Patch, b.Patch)
}

// String renders the version as "major.minor.patch".
func (v SemVer) String() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}

// cmpUint returns the sign of (a - b) without risking overflow.
func cmpUint(a, b uint) int {
	switch {
	case a < b:
		return -1
	case a > b:
		return 1
	default:
		return 0
	}
}
