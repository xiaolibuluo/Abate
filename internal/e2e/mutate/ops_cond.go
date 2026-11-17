/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package mutate

// condBefore lists the real condition spellings of each registered language.
var condBefore = map[string][]string{
	"go": {"true", "false", "&&", "||", "!", "if", "else"},
	"py": {"True", "False", "and", "or", "not", "if", "else"},
	"js": {"true", "false", "&&", "||", "!", "if", "else"},
	"c":  {"true", "false", "&&", "||", "!", "if", "else"},
	"cs": {"true", "false", "&&", "||", "!", "if", "else"},
	"rs": {"true", "false", "&&", "||", "!", "if", "else"},
	"kt": {"true", "false", "&&", "||", "!", "if", "else"},
}

// RegisterCondition adds the P1 condition family to the registry.
func RegisterCondition(r *Registry) error {
	return r.Register(conditionSpec())
}

// conditionSpec builds the condition operator from the per-language table.
func conditionSpec() OperatorSpec {
	return newSpec(OperatorSpec{
		ID:     "cond",
		Class:  ClassStructure,
		Points: crossByLanguage(condBefore, condBefore),
	})
}
