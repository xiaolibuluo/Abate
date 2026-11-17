/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package mutate

// blockBefore and blockAfter are real block or statement spellings. An empty
// after-token deletes the construct instead of adding a comment or whitespace.
var (
	blockBefore = []string{"{}", "{ }", "pass", "else", "break", "return", "end"}
	blockAfter  = []string{"", "{ }", "{}", "pass", "else", "break", "return"}
)

// RegisterBlock adds the P1 statement-block family to the registry.
func RegisterBlock(r *Registry) error {
	return r.Register(blockSpec())
}

// blockSpec builds the block operator from the two statement tables.
func blockSpec() OperatorSpec {
	return newSpec(OperatorSpec{
		ID:     "block",
		Class:  ClassStructure,
		Points: spreadPoints(crossPoints(blockBefore, blockAfter)),
	})
}
