/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package air

// Genuine C#, Kotlin and Rust construct kinds appended after the pyfront block.
//
// Like the KindPy* block above, every value below is produced by a real
// syntax-node reduction in its frontend: the parser emits it while lowering
// the corresponding construct, anchored on a real token position of that
// construct, never by switching on a detached token's spelling. Values are
// appended explicitly so the frozen wire numbering of every earlier kind
// stays unchanged and SchemaVersion stays 2.
const (
	KindCsConstructor       NodeKind = 128
	KindCsDestructor        NodeKind = 129
	KindCsOperatorDecl      NodeKind = 130
	KindCsLocalFunc         NodeKind = 131
	KindCsAccessor          NodeKind = 132
	KindCsProperty          NodeKind = 133
	KindCsIndexer           NodeKind = 134
	KindCsEventDecl         NodeKind = 135
	KindCsRecord            NodeKind = 136
	KindCsYield             NodeKind = 137
	KindCsConditionalAccess NodeKind = 138
	KindCsObjectCreation    NodeKind = 139
	KindCsSizeof            NodeKind = 140
	KindCsTypeof            NodeKind = 141
	KindCsQuery             NodeKind = 142
	KindCsInterpString      NodeKind = 143
	KindCsStringFrag        NodeKind = 144
	KindCsCharBody          NodeKind = 145
	KindCsModifier          NodeKind = 146
	KindCsExternAlias       NodeKind = 147
	KindCsPattern           NodeKind = 148
	KindCsAliasName         NodeKind = 149
	KindKtCompanion         NodeKind = 150
	KindKtSecondaryCtor     NodeKind = 151
	KindKtCtorDelegation    NodeKind = 152
	KindKtAccessor          NodeKind = 153
	KindKtImportAlias       NodeKind = 154
	KindKtImportList        NodeKind = 155
	KindKtInterp            NodeKind = 156
	KindKtStringFrag        NodeKind = 157
	KindKtProperty          NodeKind = 158
	KindKtThis              NodeKind = 159
	KindKtSuper             NodeKind = 160
	KindKtWhenEntry         NodeKind = 161
	KindKtCallableRef       NodeKind = 162
	KindKtControlBody       NodeKind = 163
	KindRsAsyncBlock        NodeKind = 164
	KindRsConstBlock        NodeKind = 165
	KindRsGenBlock          NodeKind = 166
	KindRsTryBlock          NodeKind = 167
	KindRsUnsafeBlock       NodeKind = 168
	KindRsCompoundAssign    NodeKind = 169
	KindRsFieldInit         NodeKind = 170
	KindRsBaseField         NodeKind = 171
	KindRsShorthandField    NodeKind = 172
	KindRsShorthandIdent    NodeKind = 173
	KindRsFieldIdent        NodeKind = 174
	KindRsMetavar           NodeKind = 175
	KindRsSelf              NodeKind = 176
	KindRsCrate             NodeKind = 177
	KindRsSuper             NodeKind = 178
	KindRsEmptyStmt         NodeKind = 179
	KindRsStringFrag        NodeKind = 180
	KindRsExternCrate       NodeKind = 181
	KindRsScopedUse         NodeKind = 182
	KindRsUseList           NodeKind = 183
	KindRsUseWildcard       NodeKind = 184
	KindRsUseAs             NodeKind = 185
	KindRsNegativeLit       NodeKind = 186
	KindRsYield             NodeKind = 187
	KindRsMacroCall         NodeKind = 188
	KindRsMacroDef          NodeKind = 189
	KindRsFuncSig           NodeKind = 190
	KindRsGenericFunc       NodeKind = 191
	KindRsTurbofish         NodeKind = 192
	KindRsTypeArgs          NodeKind = 193
)
