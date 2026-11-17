/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package air

import "fmt"

// Diff is the deterministic node-level difference between two AIR units.
//
// Added and Removed hold identifiers in ascending order. Reused counts the
// identifiers present in both multisets. FullRebuild is true when there was no
// prior unit to reuse, so callers can tell a genuine incremental parse from a
// cold parse without guessing.
type Diff struct {
	// Added lists identifiers that appear more often in the new unit.
	Added []NodeID
	// Removed lists identifiers that appear more often in the old unit.
	Removed []NodeID
	// Reused is the size of the identifier multiset intersection.
	Reused int
	// FullRebuild reports that every subtree had to be parsed from scratch.
	FullRebuild bool
}

// Empty reports whether both units described the same node multiset.
func (d Diff) Empty() bool {
	return len(d.Added) == 0 && len(d.Removed) == 0
}

// String renders the diff summary in a stable form.
func (d Diff) String() string {
	return fmt.Sprintf("added=%d removed=%d reused=%d full=%t",
		len(d.Added), len(d.Removed), d.Reused, d.FullRebuild)
}
