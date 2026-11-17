/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package mutate

// MutationClass names one of the four separate report columns.
type MutationClass string

// The four mutation classes of the four-way report.
const (
	// ClassText marks a textual mutation.
	ClassText MutationClass = "text"
	// ClassStructure marks a structural mutation.
	ClassStructure MutationClass = "structure"
	// ClassBroken marks a deliberately broken mutation.
	ClassBroken MutationClass = "broken"
	// ClassEquivalent marks a semantically equivalent mutation.
	ClassEquivalent MutationClass = "equivalent"
)

// Valid reports whether c is one of the four declared classes.
func (c MutationClass) Valid() bool {
	switch c {
	case ClassText, ClassStructure, ClassBroken, ClassEquivalent:
		return true
	default:
		return false
	}
}

// Column returns the lowercase column key of c.
func (c MutationClass) Column() string {
	return string(c)
}

// ParseClass resolves a lowercase column key into a mutation class.
func ParseClass(key string) (MutationClass, bool) {
	class := MutationClass(key)
	if !class.Valid() {
		return "", false
	}
	return class, true
}
