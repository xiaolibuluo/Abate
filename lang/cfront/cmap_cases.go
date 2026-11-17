/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package cfront

import (
	"abate/lang/air"
)

// cConstructCase 是「构造 → AIR」映射的可枚举契约，供构造断言与变异测试使用。
type cConstructCase struct {
	Name         string
	Tokens       []string
	Top          cNesting
	WantKind     air.NodeKind
	WantSemKey   air.SemKey
	WantSemValue string
}

// cConstructCases 返回 ≥20 个真实 C/C++ 构造的映射断言，全部只用冻结枚举。
func cConstructCases() []cConstructCase {
	return []cConstructCase{
		{"include-directive", []string{"#", "include", "\"a.h\""}, true, air.KindDirective, air.SemDirective, "include"},
		{"define-directive", []string{"#", "define", "X", "1"}, true, air.KindDirective, air.SemDirective, "define"},
		{"ifdef-directive", []string{"#", "ifdef", "FLAG"}, true, air.KindDirective, air.SemDirective, "ifdef"},
		{"pragma-directive", []string{"#", "pragma", "once"}, true, air.KindDirective, air.SemDirective, "pragma"},
		{"template", []string{"template", "<", "typename", "T", ">", "struct", "Box"}, true, air.KindGenericInst, air.SemTypeParam, "template"},
		{"namespace", []string{"namespace", "N", "{", "}"}, true, air.KindModuleDecl, air.SemInvalid, ""},
		{"class", []string{"class", "Foo", "{", "}"}, true, air.KindClassDecl, air.SemInvalid, ""},
		{"struct", []string{"struct", "S", "{", "}"}, true, air.KindStructDecl, air.SemInvalid, ""},
		{"union", []string{"union", "U", "{", "}"}, true, air.KindStructDecl, air.SemNote, "union"},
		{"enum", []string{"enum", "E", "{", "}"}, true, air.KindEnumDecl, air.SemInvalid, ""},
		{"enum-class", []string{"enum", "class", "E", "{", "}"}, true, air.KindEnumDecl, air.SemNote, "enum-class"},
		{"using", []string{"using", "Alias", "=", "int"}, true, air.KindUsingDecl, air.SemInvalid, ""},
		{"typedef", []string{"typedef", "int", "Word"}, true, air.KindTypeDecl, air.SemInvalid, ""},
		{"function-definition", []string{"int", "main", "(", ")", "{", "}"}, true, air.KindFuncDecl, air.SemInvalid, ""},
		{"function-declaration", []string{"void", "f", "(", "int", ")"}, true, air.KindFuncDecl, air.SemInvalid, ""},
		{"variable", []string{"int", "x", "=", "1"}, true, air.KindVarDecl, air.SemInvalid, ""},
		{"constant", []string{"const", "int", "N", "=", "4"}, true, air.KindConstDecl, air.SemInvalid, ""},
		{"static-variable", []string{"static", "int", "counter"}, true, air.KindVarDecl, air.SemNote, "static"},
		{"return", []string{"return", "0"}, false, air.KindReturn, air.SemInvalid, ""},
		{"if", []string{"if", "(", "x", ")", "{", "}"}, false, air.KindIf, air.SemInvalid, ""},
		{"for", []string{"for", "(", ";", ";", ")", "{", "}"}, false, air.KindFor, air.SemInvalid, ""},
		{"while", []string{"while", "(", "x", ")", "{", "}"}, false, air.KindWhile, air.SemInvalid, ""},
		{"do-while", []string{"do", "{", "}", "while", "(", "x", ")"}, false, air.KindDoWhile, air.SemInvalid, ""},
		{"switch", []string{"switch", "(", "x", ")", "{", "}"}, false, air.KindSwitch, air.SemInvalid, ""},
		{"case", []string{"case", "1", ":"}, false, air.KindCase, air.SemInvalid, ""},
		{"break", []string{"break"}, false, air.KindBranch, air.SemNote, "break"},
		{"continue", []string{"continue"}, false, air.KindBranch, air.SemNote, "continue"},
		{"throw", []string{"throw", "e"}, false, air.KindThrow, air.SemInvalid, ""},
		{"try", []string{"try", "{", "}"}, false, air.KindTry, air.SemInvalid, ""},
		{"catch", []string{"catch", "(", "e", ")", "{", "}"}, false, air.KindCatch, air.SemInvalid, ""},
		{"sizeof", []string{"sizeof", "(", "int", ")"}, false, air.KindCast, air.SemNote, "sizeof"},
		{"new", []string{"new", "Foo"}, false, air.KindUnary, air.SemNote, "new"},
		{"delete", []string{"delete", "p"}, false, air.KindUnary, air.SemNote, "delete"},
		{"call-expression", []string{"f", "(", "1", ")"}, false, air.KindCall, air.SemInvalid, ""},
		{"expression-statement", []string{"x", "+", "1"}, false, air.KindExprStmt, air.SemInvalid, ""},
	}
}
