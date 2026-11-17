/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package grammar

import (
	"strings"

	"abate/lang/air"
)

// The three mapping statuses classify how faithful a lowering is.
//
// A construct is exact when AIR models it with a dedicated node kind, it is
// degraded when AIR has no dedicated kind and the construct shares one with
// its neighbours, and it is out of scope when the construct is outside the
// capability domain of Abate altogether.
const (
	// StatusExact marks a construct AIR models with a dedicated node kind.
	StatusExact = "exact"
	// StatusDegraded marks a construct lowered onto a shared AIR node kind.
	StatusDegraded = "degraded"
	// StatusOutOfScope marks a construct outside the AIR capability domain.
	StatusOutOfScope = "out-of-scope"
	// StatusUnmapped marks a construct that no rule selected.
	StatusUnmapped = "unmapped"
)

// MappingRule maps upstream node type names onto one AIR node kind.
//
// A rule selects a node name when the name equals one of the Exact entries,
// ends with one of the Suffix entries, starts with one of the Prefix entries
// or contains one of the Substr entries. The tables are evaluated in order and
// the first matching rule wins, so the rule list reads like a routing table:
// the specific entries come first and the generic families come last.
type MappingRule struct {
	// Name identifies the rule in reports and in tests.
	Name string
	// Kind is the AIR node kind the construct is lowered to.
	Kind air.NodeKind
	// Status classifies the fidelity of the lowering.
	Status string
	// Reason documents why the construct is lowered this way.
	Reason string
	// Exact lists upstream node names matched verbatim.
	Exact []string
	// Suffix lists upstream node name suffixes matched at the end.
	Suffix []string
	// Prefix lists upstream node name prefixes matched at the start.
	Prefix []string
	// Substr lists upstream node name fragments matched anywhere.
	Substr []string
}

// Matches reports whether the rule selects the given upstream node name.
func (r MappingRule) Matches(name string) bool {
	for _, exact := range r.Exact {
		if name == exact {
			return true
		}
	}
	for _, suffix := range r.Suffix {
		if strings.HasSuffix(name, suffix) {
			return true
		}
	}
	for _, prefix := range r.Prefix {
		if strings.HasPrefix(name, prefix) {
			return true
		}
	}
	for _, fragment := range r.Substr {
		if strings.Contains(name, fragment) {
			return true
		}
	}
	return false
}

// exact builds a rule for a construct AIR models with a dedicated kind.
func exact(name string, kind air.NodeKind, names ...string) MappingRule {
	return MappingRule{
		Name:   name,
		Kind:   kind,
		Status: StatusExact,
		Reason: "AIR 以独立节点建模该构造",
		Exact:  names,
	}
}

// degraded builds a rule for a construct lowered onto a shared AIR kind.
func degraded(name string, kind air.NodeKind, reason string, names ...string) MappingRule {
	return MappingRule{
		Name:   name,
		Kind:   kind,
		Status: StatusDegraded,
		Reason: reason,
		Exact:  names,
	}
}

// scoped builds a rule for a construct outside the AIR capability domain.
func scoped(name string, reason string, names ...string) MappingRule {
	return MappingRule{
		Name:   name,
		Kind:   air.KindInvalid,
		Status: StatusOutOfScope,
		Reason: reason,
		Exact:  names,
	}
}
