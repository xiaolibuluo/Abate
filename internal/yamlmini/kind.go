/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package yamlmini

import "fmt"

// Kind identifies the shape of a Node.
type Kind uint8

const (
	// Scalar is a quoted or plain scalar carried as text in Value.
	Scalar Kind = iota
	// Mapping is an ordered key/value map; Keys parallels Vals.
	Mapping
	// Sequence is an ordered list carried in Items.
	Sequence
	// Null is an empty value ("null", "~" or an absent value).
	Null
)

func (k Kind) String() string {
	switch k {
	case Scalar:
		return "scalar"
	case Mapping:
		return "mapping"
	case Sequence:
		return "sequence"
	case Null:
		return "null"
	default:
		return fmt.Sprintf("kind(%d)", uint8(k))
	}
}

// scalarKind records how a plain scalar was lexed so that the binder keeps the
// distinction between, for example, the string "42" and the integer 42.
type scalarKind uint8

const (
	sString scalarKind = iota
	sInt
	sFloat
	sBool
	sNull
)
