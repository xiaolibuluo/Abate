/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package chaos

// DamageKind labels one deterministic corruption shape applied to a buffer.
type DamageKind string

// The four declared damage shapes.
const (
	// DamageTruncate marks a half truncation.
	DamageTruncate DamageKind = "truncate"
	// DamageSplice marks an inserted byte chunk.
	DamageSplice DamageKind = "splice"
	// DamageDrop marks the removal of a token.
	DamageDrop DamageKind = "drop"
	// DamageMerge marks a sha256-deduplicated corpus merge.
	DamageMerge DamageKind = "merge"
)

// Valid reports whether k is one of the four declared damage shapes.
func (k DamageKind) Valid() bool {
	switch k {
	case DamageTruncate, DamageSplice, DamageDrop, DamageMerge:
		return true
	default:
		return false
	}
}
