/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package grammar

// Rule is one upstream production taken from grammar.json.
//
// A production is the upstream shape of one syntax rule: a symbol reference, a
// choice, a sequence, a repetition or a lexical pattern. The members are the
// production names the rule references, in source order.
type Rule struct {
	// Name is the upstream production name.
	Name string
	// Kind is the production shape, for example SYMBOL, CHOICE, SEQ, REPEAT,
	// STRING, PATTERN, TOKEN, PREC or ALIAS.
	Kind string
	// Members lists the referenced production names in source order without
	// duplicates.
	Members []string
}

// References reports whether the production references the named symbol.
func (r Rule) References(name string) bool {
	for _, member := range r.Members {
		if member == name {
			return true
		}
	}
	return false
}
