/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package metrics

// Column keys of the four-way report.
const (
	// ColumnText is the key of the textual mutation column.
	ColumnText = "text"
	// ColumnStructure is the key of the structural mutation column.
	ColumnStructure = "structure"
	// ColumnBroken is the key of the broken mutation column.
	ColumnBroken = "broken"
	// ColumnEquivalent is the key of the equivalent mutation column.
	ColumnEquivalent = "equivalent"
)

type (
	// ColumnStats is one column of the four-way mutation report. Gross counts
	// every mutation of that class, Effective subtracts the equivalent ones,
	// Captured counts mutations whose AIR identifier changed, and CapturePct is
	// Captured divided by Gross in percent.
	ColumnStats struct {
		Gross      int
		Effective  int
		Captured   int
		CapturePct float64
	}

	// Columns is the four-way mutation report. Text, Structure, Broken and
	// Equivalent remain distinct fields by construction; no method of this
	// package folds them into a single score, so a report cannot silently merge
	// the four classes. Equivalent mutations are never deleted from Gross.
	Columns struct {
		Text       ColumnStats
		Structure  ColumnStats
		Broken     ColumnStats
		Equivalent ColumnStats
	}
)

// withPct returns a copy of s whose CapturePct is derived from Gross and
// Captured. A zero Gross yields a zero percentage instead of a division by zero.
func (s ColumnStats) withPct() ColumnStats {
	out := s
	if out.Gross <= 0 {
		out.CapturePct = 0
		return out
	}
	out.CapturePct = float64(out.Captured) * 100 / float64(out.Gross)
	return out
}

// column returns a pointer to the column selected by key. An unknown key maps
// to the broken column so no record can leave the report uncounted.
func (c *Columns) column(key string) *ColumnStats {
	switch key {
	case ColumnText:
		return &c.Text
	case ColumnStructure:
		return &c.Structure
	case ColumnEquivalent:
		return &c.Equivalent
	default:
		return &c.Broken
	}
}

// finalize fills the derived Effective and CapturePct fields of every column.
func (c Columns) finalize(equivalentCount int) Columns {
	c.Text.Effective = c.Text.Gross
	c.Structure.Effective = c.Structure.Gross
	c.Broken.Effective = c.Broken.Gross
	c.Equivalent.Effective = c.Equivalent.Gross - equivalentCount
	c.Text = c.Text.withPct()
	c.Structure = c.Structure.withPct()
	c.Broken = c.Broken.withPct()
	c.Equivalent = c.Equivalent.withPct()
	return c
}
