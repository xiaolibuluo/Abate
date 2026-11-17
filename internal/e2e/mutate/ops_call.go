/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package mutate

// callBefore and callAfter are real call-expression spellings. An empty
// after-token removes the call rather than inserting a formatting change.
var (
	callBefore = []string{"f()", "g()", "foo()", "bar()", "call()", "print()", "baz()"}
	callAfter  = []string{"", "f(1)", "g()", "foo()", "bar()", "call()", "print()"}
)

// RegisterCall adds the P1 call-expression family to the registry.
func RegisterCall(r *Registry) error {
	return r.Register(callSpec())
}

// callSpec builds the call operator from the two call tables.
func callSpec() OperatorSpec {
	return newSpec(OperatorSpec{
		ID:     "call",
		Class:  ClassStructure,
		Points: spreadPoints(crossPoints(callBefore, callAfter)),
	})
}
