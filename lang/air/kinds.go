/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package air

// NodeKind identifies the syntactic class of an AIR node.
//
// The enumeration is the union of the constructs found in the seven target
// languages. Values are part of the wire format and must never be renumbered.
type NodeKind uint16

// KindInvalid marks an unset, unknown, or unsupported node kind.
const KindInvalid NodeKind = 0

// Declaration kinds cover package, import, type, value, and function bindings.
const (
	KindPackage NodeKind = iota + 1
	KindImport
	KindTypeDecl
	KindConstDecl
	KindVarDecl
	KindFuncDecl
	KindMethodDecl
	KindStructDecl
	KindInterfaceDecl
	KindEnumDecl
	KindClassDecl
	KindTraitDecl
	KindModuleDecl
	KindAnnotationDecl
	KindUsingDecl
)

// Statement kinds cover control flow, blocks, and concurrency statements.
const (
	KindBlock NodeKind = iota + 16
	KindExprStmt
	KindDeclStmt
	KindIf
	KindFor
	KindWhile
	KindDoWhile
	KindRangeFor
	KindSwitch
	KindCase
	KindSelect
	KindDefer
	KindGoStmt
	KindReturn
	KindBranch
	KindThrow
	KindTry
	KindCatch
	KindFinally
	KindWith
	KindLabel
	KindMatch
)

// Expression kinds cover operands, operators, calls, and literal forms.
const (
	KindIdent NodeKind = iota + 38
	KindNumLit
	KindStrLit
	KindCharLit
	KindBoolLit
	KindNilLit
	KindBinary
	KindUnary
	KindCall
	KindIndex
	KindSlice
	KindField
	KindSelector
	KindCompositeLit
	KindFuncLit
	KindArrow
	KindLambda
	KindCast
	KindTypeAssert
	KindAssign
	KindIncDec
	KindStar
	KindAddr
	KindTernary
	KindSpread
	KindChanOp
	KindMatchExpr
	KindAwait
)

// Type kinds describe type expressions. SliceType and ChanType are suffixed to
// disambiguate them from the expression kinds Slice and ChanOp.
const (
	KindNamed NodeKind = iota + 66
	KindPointer
	KindSliceType
	KindArray
	KindMap
	KindChanType
	KindFuncType
	KindStructType
	KindInterfaceType
	KindGenericInst
	KindGenericParam
	KindUnion
	KindOptional
	KindTuple
	KindRef
)

// Other kinds cover trivia, metadata, and compiler directives.
const (
	KindComment NodeKind = iota + 81
	KindDoc
	KindAttribute
	KindDecorator
	KindDirective
	KindRegion
)

// Dedicated construct kinds appended after KindRegion.
//
// These values are the AIR capability that G4 adds for constructs the seven
// frontends used to lower onto KindRegion or another shared kind. Values are
// spelled out explicitly and only appended: the frozen wire enumeration never
// renumbers an existing value and SchemaVersion stays 2.
const (
	KindInterpolation   NodeKind = 87
	KindPattern         NodeKind = 88
	KindPropertyDecl    NodeKind = 89
	KindWhenBranch      NodeKind = 90
	KindAttributeArg    NodeKind = 91
	KindLinqClause      NodeKind = 92
	KindParameterList   NodeKind = 93
	KindStringFragment  NodeKind = 94
	KindMacroInvocation NodeKind = 95
	KindUseDetail       NodeKind = 96
	KindQualifiedName   NodeKind = 97
	KindCompoundAssign  NodeKind = 98
	KindBlockVariant    NodeKind = 99
	KindTypeArg         NodeKind = 100
	KindInvocation      NodeKind = 101
	KindSelectorDetail  NodeKind = 102
	KindDelimiter       NodeKind = 103
	KindOperatorToken   NodeKind = 104
	KindSyntaxDetail    NodeKind = 105
	KindPartialNode     NodeKind = 106
	KindSpecialIdent    NodeKind = 107
	KindComprehension   NodeKind = 108
)

// Genuine Python construct kinds appended after the G4 block.
//
// Unlike kinds 87-108, each value below is produced by a real syntax-node
// reduction in pyfront: the parser decides to emit it while lowering the
// corresponding Python construct, not by switching on a detached token's
// spelling. Values are appended explicitly so the frozen wire numbering of
// every earlier kind stays unchanged and SchemaVersion stays 2.
const (
	KindPyAssert         NodeKind = 109
	KindPyQualifiedName  NodeKind = 110
	KindPyStringFragment NodeKind = 111
	KindPyAugAssign      NodeKind = 112
	KindPyNamedExpr      NodeKind = 113
	KindPyKeywordArg     NodeKind = 114
	KindPyDictPair       NodeKind = 115
	KindPyCompClause     NodeKind = 116
	KindPyBoolOp         NodeKind = 117
	KindPyCompare        NodeKind = 118
	KindPyGenerator      NodeKind = 119
	KindPyImportPrefix   NodeKind = 120
	KindPyDelete         NodeKind = 121
	KindPyNotOp          NodeKind = 122
	KindPyYield          NodeKind = 123
	KindPyScopeDecl      NodeKind = 124
	KindPyPass           NodeKind = 125
	KindPyWithItem       NodeKind = 126
	KindPyTypeConversion NodeKind = 127
)

// kindLast is the largest allocated node kind.
const kindLast = KindJsEmptyStmt

// kindNames is a read-only lookup table indexed by NodeKind. It is never
// mutated after initialization, so it cannot leak nondeterministic state.
var kindNames = [...]string{
	"Invalid",
	"Package", "Import", "TypeDecl", "ConstDecl", "VarDecl", "FuncDecl",
	"MethodDecl", "StructDecl", "InterfaceDecl", "EnumDecl", "ClassDecl",
	"TraitDecl", "ModuleDecl", "AnnotationDecl", "UsingDecl",
	"Block", "ExprStmt", "DeclStmt", "If", "For", "While", "DoWhile",
	"RangeFor", "Switch", "Case", "Select", "Defer", "GoStmt", "Return",
	"Branch", "Throw", "Try", "Catch", "Finally", "With", "Label", "Match",
	"Ident", "NumLit", "StrLit", "CharLit", "BoolLit", "NilLit", "Binary",
	"Unary", "Call", "Index", "Slice", "Field", "Selector", "CompositeLit",
	"FuncLit", "Arrow", "Lambda", "Cast", "TypeAssert", "Assign", "IncDec",
	"Star", "Addr", "Ternary", "Spread", "ChanOp", "MatchExpr", "Await",
	"Named", "Pointer", "SliceType", "Array", "Map", "ChanType", "FuncType",
	"StructType", "InterfaceType", "GenericInst", "GenericParam", "Union",
	"Optional", "Tuple", "Ref",
	"Comment", "Doc", "Attribute", "Decorator", "Directive", "Region",
	"Interpolation", "Pattern", "PropertyDecl", "WhenBranch", "AttributeArg",
	"LinqClause", "ParameterList", "StringFragment", "MacroInvocation",
	"UseDetail", "QualifiedName", "CompoundAssign", "BlockVariant", "TypeArg",
	"Invocation", "SelectorDetail", "Delimiter", "OperatorToken",
	"SyntaxDetail", "PartialNode", "SpecialIdent", "Comprehension",
	"PyAssert", "PyQualifiedName", "PyStringFragment", "PyAugAssign",
	"PyNamedExpr", "PyKeywordArg", "PyDictPair", "PyCompClause",
	"PyBoolOp", "PyCompare", "PyGenerator", "PyImportPrefix",
	"PyDelete", "PyNotOp", "PyYield", "PyScopeDecl", "PyPass",
	"PyWithItem", "PyTypeConversion",
	"CsConstructor", "CsDestructor", "CsOperatorDecl", "CsLocalFunc",
	"CsAccessor", "CsProperty", "CsIndexer", "CsEventDecl", "CsRecord",
	"CsYield", "CsConditionalAccess", "CsObjectCreation", "CsSizeof",
	"CsTypeof", "CsQuery", "CsInterpString", "CsStringFrag", "CsCharBody",
	"CsModifier", "CsExternAlias", "CsPattern", "CsAliasName",
	"KtCompanion", "KtSecondaryCtor", "KtCtorDelegation", "KtAccessor",
	"KtImportAlias", "KtImportList", "KtInterp", "KtStringFrag",
	"KtProperty", "KtThis", "KtSuper", "KtWhenEntry", "KtCallableRef",
	"KtControlBody",
	"RsAsyncBlock", "RsConstBlock", "RsGenBlock", "RsTryBlock",
	"RsUnsafeBlock", "RsCompoundAssign", "RsFieldInit", "RsBaseField",
	"RsShorthandField", "RsShorthandIdent", "RsFieldIdent", "RsMetavar",
	"RsSelf", "RsCrate", "RsSuper", "RsEmptyStmt", "RsStringFrag",
	"RsExternCrate", "RsScopedUse", "RsUseList", "RsUseWildcard",
	"RsUseAs", "RsNegativeLit", "RsYield", "RsMacroCall", "RsMacroDef",
	"RsFuncSig", "RsGenericFunc", "RsTurbofish", "RsTypeArgs",
	"JsExport", "JsGeneratorFnDecl", "JsGeneratorFn", "JsClassExpr",
	"JsClassBody", "JsStaticBlock", "JsSwitchBody", "JsSwitchDefault",
	"JsLexicalDecl", "JsUsingDecl", "JsDeclarator", "JsAugAssign",
	"JsNew", "JsOptionalChain", "JsParenExpr", "JsSequence", "JsYield",
	"JsPair", "JsRestPattern", "JsTemplateSub", "JsStringPart", "JsRegex",
	"JsPropertyIdent", "JsThisSuper", "JsJsx",
	"JsClassHeritage", "JsIncDec", "JsArguments", "JsFormalParams",
	"KtForIn", "RsLoop", "RsImpl",
	"JsDebugger", "JsMetaProperty", "JsComputedName",
	"JsPairPattern", "JsObjectAssignPat",
	"CUnion", "CEnumMember", "CDefined", "CInitPair",
	"CLabel", "CDeref", "CFieldDecl",
	"PyStarList", "PySplatType", "PyExprList",
	"JsAssignPattern", "JsArrayPattern", "JsObjectPattern",
	"CMacroType", "CInitDeclarator", "CFuncDeclarator",
	"CFieldIdent", "CSizedType", "CTypeIdent",
	"RsUnion", "RsEnumVariant", "RsArray", "RsTupleType",
	"RsVisibility", "RsMutable", "RsExternMod", "RsFnModifiers",
	"CsDeclarator", "CsArrowClause", "CsNamedType", "CsTupleType",
	"KtMultiVar", "KtRange", "KtSimpleIdent", "KtInfixCtor", "KtAnnotatedLambda",
	"JsEmptyStmt",
}

// String returns the stable wire name of the kind.
func (k NodeKind) String() string {
	if !k.Valid() {
		return kindNames[0]
	}
	return kindNames[int(k)]
}

// Valid reports whether the kind is one of the allocated enumeration values.
func (k NodeKind) Valid() bool {
	return k > KindInvalid && k <= kindLast && int(k) < len(kindNames)
}

// CategoryName returns the family of the kind: declaration, statement,
// expression, type, other, or invalid.
func (k NodeKind) CategoryName() string {
	switch {
	case k >= KindPackage && k <= KindUsingDecl:
		return "declaration"
	case k >= KindBlock && k <= KindMatch:
		return "statement"
	case k >= KindIdent && k <= KindAwait:
		return "expression"
	case k >= KindNamed && k <= KindRef:
		return "type"
	case k >= KindComment && k <= kindLast:
		return "other"
	default:
		return "invalid"
	}
}
