/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package metrics

import "errors"

// errInvalidMutationRow reports a row whose counters disagree with its Gross
// denominator or whose derived percentages are out of range.
var errInvalidMutationRow = errors.New("metrics: inconsistent mutation row")

// MutationRow is one (language, operator, class) row of the mutation report.
// Gross is the strict denominator: equivalent mutations stay inside it and are
// never removed. ValidOnlyPct is an auxiliary competitor-style observation and
// never replaces CapturePct. SilentDrops is always zero.
type MutationRow struct {
	Language     string
	OperatorID   string
	Class        string
	Gross        int
	Effective    int
	Equivalent   int
	Captured     int
	Missed       int
	Timeout      int
	Unviable     int
	SilentDrops  int
	CapturePct   float64
	ValidOnlyPct float64
}

// Validate checks that every mutation is accounted for exactly once, that no
// mutation was silently dropped, and that both percentages are in [0, 100].
func (r MutationRow) Validate() error {
	if r.Language == "" || r.OperatorID == "" || r.Class == "" ||
		r.Gross < 0 || r.Effective < 0 || r.Equivalent < 0 ||
		r.Captured < 0 || r.Missed < 0 || r.Timeout < 0 ||
		r.Unviable < 0 || r.SilentDrops != 0 {
		return errInvalidMutationRow
	}
	if r.Gross != r.Captured+r.Missed+r.Timeout+r.Unviable {
		return errInvalidMutationRow
	}
	if r.Effective != r.Gross-r.Equivalent {
		return errInvalidMutationRow
	}
	if r.CapturePct < 0 || r.CapturePct > 100 ||
		r.ValidOnlyPct < 0 || r.ValidOnlyPct > 100 {
		return errInvalidMutationRow
	}
	return nil
}
