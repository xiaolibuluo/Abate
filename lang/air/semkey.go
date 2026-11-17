/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package air

// SemKey identifies one slot of the language-neutral semantic annotation map.
//
// Values are part of the wire format. New slots are appended only; a decoder
// silently ignores keys it does not know, which is the forward-compatibility
// mechanism for additive semantic changes.
type SemKey uint16

// SemInvalid marks an unset annotation key.
const SemInvalid SemKey = 0

// Semantic annotation keys. The first five are the core annotation slots;
// the rest carry frontend-specific facts in a language-neutral form.
const (
	SemTypeRef SemKey = iota + 1
	SemScopeID
	SemCFGEdges
	SemMutability
	SemPurity
	SemIota
	SemAsync
	SemDirective
	SemTag
	SemNote
	SemPartial
	SemChanDir
	SemOperator
	SemLabel
	SemTypeParam
)

// SemKeyLast is the largest allocated annotation key.
const SemKeyLast = SemTypeParam

// semKeyNames is a read-only lookup table indexed by SemKey.
var semKeyNames = [...]string{
	"Invalid",
	"TypeRef", "ScopeID", "CFGEdges", "Mutability", "Purity", "Iota",
	"Async", "Directive", "Tag", "Note", "Partial", "ChanDir",
	"Operator", "Label", "TypeParam",
}

// String returns the stable wire name of the annotation key.
func (k SemKey) String() string {
	if !k.Valid() {
		return semKeyNames[0]
	}
	return semKeyNames[int(k)]
}

// Valid reports whether the key is one of the allocated annotation slots.
func (k SemKey) Valid() bool {
	return k > SemInvalid && k <= SemKeyLast && int(k) < len(semKeyNames)
}
