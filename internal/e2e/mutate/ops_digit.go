/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package mutate

// digitBefore and digitAfter are concrete integer literal spellings. Every
// rewrite changes the literal text, which is exactly what the AIR-ID capture
// criterion observes.
var (
	digitBefore = []string{"0", "1", "2", "3", "10", "16", "32", "64"}
	digitAfter  = []string{"9", "8", "7", "41", "128", "256", "512", "1024"}
)

// RegisterDigit adds the P0 numeric-literal family to the registry.
func RegisterDigit(r *Registry) error {
	return r.Register(digitSpec())
}

// digitSpec builds the numeric-literal operator from the two literal tables.
func digitSpec() OperatorSpec {
	return newSpec(OperatorSpec{
		ID:     "digit",
		Class:  ClassText,
		Points: spreadPoints(crossPoints(digitBefore, digitAfter)),
	})
}
