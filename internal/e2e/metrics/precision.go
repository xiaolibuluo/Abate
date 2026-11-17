/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package metrics

import "errors"

// errInvalidPrecisionRow reports a row whose counters or percentages are
// outside their declared ranges.
var errInvalidPrecisionRow = errors.New("metrics: invalid precision row")

// PrecisionRow is one language's precision observation. ZeroDiagnosticZeroError
// counts clean parses, DegradeNodePct is the degraded-node share, and
// OutsideDamageRetentionPct is the share of structure retained outside a
// damaged byte range. The three values are independent of each other.
type PrecisionRow struct {
	Language                  string
	ZeroDiagnosticZeroError   int
	Diagnostics               int
	ParseErrors               int
	DegradeNodePct            float64
	OutsideDamageRetentionPct float64
	ConstructCoveragePct      float64
}

// Validate checks that counters are non-negative and every percentage lies in
// the closed [0, 100] interval.
func (r PrecisionRow) Validate() error {
	if r.Language == "" || r.ZeroDiagnosticZeroError < 0 ||
		r.Diagnostics < 0 || r.ParseErrors < 0 {
		return errInvalidPrecisionRow
	}
	for _, pct := range []float64{r.DegradeNodePct,
		r.OutsideDamageRetentionPct, r.ConstructCoveragePct} {
		if pct < 0 || pct > 100 {
			return errInvalidPrecisionRow
		}
	}
	return nil
}
