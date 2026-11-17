/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package mutate

// RegisterLiteral adds the P1 integer-radix family to the registry. Every
// point is a genuine equivalent mutation: a decimal integer is rewritten into
// the hexadecimal spelling of the same integer value.
func RegisterLiteral(r *Registry) error {
	return r.Register(literalSpec())
}

// literalSpec builds the equivalent radix operator with 32 points per language.
func literalSpec() OperatorSpec {
	return newSpec(OperatorSpec{
		ID:     "literal",
		Class:  ClassEquivalent,
		Points: spreadPoints(radixPoints(32)),
	})
}
