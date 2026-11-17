/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package schema

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// VirtualDuration is an integer nanosecond span carried by CES documents.
//
// Every virtual-time field uses this type so that no floating point value can
// leak into scheduling, hashing or JSON/YAML round-trips.
type VirtualDuration int64

// Canonical duration units, all expressed in integer nanoseconds.
const (
	Nanosecond  VirtualDuration = 1
	Microsecond                 = 1000 * Nanosecond
	Millisecond                 = 1000 * Microsecond
	Second                      = 1000 * Millisecond
	Minute                      = 60 * Second
	Hour                        = 60 * Minute
)

// ParseDuration converts a CES duration literal into nanoseconds.
//
// A bare integer is read as nanoseconds; otherwise the literal is a sequence
// of segments with the suffixes ns, us, ms, s, m and h, such as "1h30m",
// "250ms" or "5s". An optional leading minus sign is accepted.
func ParseDuration(s string) (VirtualDuration, error) {
	if s == "" {
		return 0, errors.New("parse duration: empty literal")
	}
	neg := strings.HasPrefix(s, "-")
	if neg {
		s = s[1:]
	}
	if s == "" {
		return 0, errors.New("parse duration: missing magnitude")
	}
	total, err := parseSegments(s)
	if err != nil {
		return 0, err
	}
	if neg {
		return -VirtualDuration(total), nil
	}
	return VirtualDuration(total), nil
}

// parseSegments sums the numeric-unit segments of a duration literal.
func parseSegments(s string) (int64, error) {
	allNum := true
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			allNum = false
			break
		}
	}
	if allNum {
		n, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return 0, fmt.Errorf("parse duration %q: %w", s, err)
		}
		return n, nil
	}

	var total int64
	rest := s
	for rest != "" {
		i := 0
		for i < len(rest) && rest[i] >= '0' && rest[i] <= '9' {
			i++
		}
		if i == 0 {
			return 0, fmt.Errorf("parse duration %q: expected a number before %q", s, rest)
		}
		num, err := strconv.ParseInt(rest[:i], 10, 64)
		if err != nil {
			return 0, fmt.Errorf("parse duration %q: %w", s, err)
		}
		tail, mult, ok := cutUnit(rest[i:])
		if !ok {
			return 0, fmt.Errorf("parse duration %q: unknown unit in %q", s, rest[i:])
		}
		if num > (int64(1<<63-1)-total)/mult {
			return 0, fmt.Errorf("parse duration %q: overflow", s)
		}
		total += num * mult
		rest = tail
	}
	return total, nil
}

// cutUnit strips the leading unit suffix and returns its nanosecond factor.
func cutUnit(s string) (string, int64, bool) {
	units := []struct {
		name string
		mul  int64
	}{
		{"ns", int64(Nanosecond)},
		{"us", int64(Microsecond)},
		{"µs", int64(Microsecond)},
		{"ms", int64(Millisecond)},
		{"s", int64(Second)},
		{"m", int64(Minute)},
		{"h", int64(Hour)},
	}
	for _, u := range units {
		if strings.HasPrefix(s, u.name) {
			return strings.TrimPrefix(s, u.name), u.mul, true
		}
	}
	return s, 0, false
}

// String returns the canonical literal, using the largest exact unit.
func (d VirtualDuration) String() string {
	v := int64(d)
	sign := ""
	if v < 0 {
		sign = "-"
		v = -v
	}
	units := []struct {
		name string
		mul  int64
	}{
		{"h", int64(Hour)},
		{"m", int64(Minute)},
		{"s", int64(Second)},
		{"ms", int64(Millisecond)},
		{"us", int64(Microsecond)},
	}
	for _, u := range units {
		if v%u.mul == 0 {
			return sign + strconv.FormatInt(v/u.mul, 10) + u.name
		}
	}
	return sign + strconv.FormatInt(v, 10) + "ns"
}

// Nanoseconds returns the raw nanosecond count.
func (d VirtualDuration) Nanoseconds() int64 {
	return int64(d)
}

// MarshalText encodes the duration as a canonical literal.
func (d VirtualDuration) MarshalText() ([]byte, error) {
	return []byte(d.String()), nil
}

// UnmarshalText decodes a CES duration literal.
func (d *VirtualDuration) UnmarshalText(text []byte) error {
	v, err := ParseDuration(string(text))
	if err != nil {
		return err
	}
	*d = v
	return nil
}
