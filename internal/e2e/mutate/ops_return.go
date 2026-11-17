/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package mutate

// returnBefore lists the real return-value spellings of each language.
var returnBefore = map[string][]string{
	"go": {"return nil", "return 0", "return false", "return true", "return 1", "return -1", "return err"},
	"py": {"return None", "return 0", "return False", "return True", "return 1", "return -1", "return value"},
	"js": {"return null", "return 0", "return false", "return true", "return 1", "return -1", "return value"},
	"c":  {"return 0", "return 1", "return -1", "return NULL", "return true", "return false", "return ptr"},
	"cs": {"return null", "return 0", "return false", "return true", "return 1", "return -1", "return value"},
	"rs": {"return 0", "return 1", "return false", "return true", "return None", "return Some", "return value"},
	"kt": {"return null", "return 0", "return false", "return true", "return 1", "return -1", "return value"},
}

// RegisterReturn adds the P1 return-value family to the registry.
func RegisterReturn(r *Registry) error {
	return r.Register(returnSpec())
}

// returnSpec builds the return-value operator from the per-language table.
func returnSpec() OperatorSpec {
	return newSpec(OperatorSpec{
		ID:     "return",
		Class:  ClassStructure,
		Points: crossByLanguage(returnBefore, returnBefore),
	})
}
