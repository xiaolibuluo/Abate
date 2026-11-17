/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package mutate

// cmpBefore lists the real comparison spellings of each registered language.
// Multi-byte operators come first so <= is never shadowed by <.
var cmpBefore = map[string][]string{
	"go": {"==", "!=", "<=", ">=", "<", ">"},
	"py": {"==", "!=", "<=", ">=", "<", ">", "is", "is not", "in", "not in"},
	"js": {"===", "!==", "==", "!=", "<=", ">=", "<", ">"},
	"c":  {"==", "!=", "<=", ">=", "<", ">"},
	"cs": {"==", "!=", "<=", ">=", "<", ">", "is"},
	"rs": {"==", "!=", "<=", ">=", "<", ">"},
	"kt": {"===", "!==", "==", "!=", "<=", ">=", "<", ">"},
}

// RegisterComparison adds the P0 comparison family to the registry.
func RegisterComparison(r *Registry) error {
	return r.Register(comparisonSpec())
}

// comparisonSpec builds the comparison operator from the per-language table.
func comparisonSpec() OperatorSpec {
	return newSpec(OperatorSpec{
		ID:     "cmp",
		Class:  ClassText,
		Points: crossByLanguage(cmpBefore, cmpBefore),
	})
}
