/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package mutate

// arithmeticTokens lists the binary operators shared by the seven frontends.
var arithmeticTokens = []string{
	"<<", ">>", "&&", "||", "+", "-", "*", "/", "%", "&", "|", "^",
}

// RegisterArithmetic adds the P0 arithmetic family to the registry.
func RegisterArithmetic(r *Registry) error {
	return r.Register(arithmeticSpec())
}

// arithmeticSpec builds the arithmetic operator from the shared token table.
func arithmeticSpec() OperatorSpec {
	return newSpec(OperatorSpec{
		ID:     "op",
		Class:  ClassText,
		Points: spreadPoints(crossPoints(arithmeticTokens, arithmeticTokens)),
	})
}
