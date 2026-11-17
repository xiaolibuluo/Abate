/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package ktfront

// 本文件承载对外的 air.Frontend 接线。声明/语句主产生式、Pratt 表达式与
// 构造到冻结枚举的映射分别留在 ktparser.go、ktexpr.go 与 ktmap.go，与
// csfront 的分工一致。

import (
	"fmt"
	"strings"

	"abate/lang/air"
	"abate/lang/frontkit"
)

// KtFrontend 是 Kotlin 前端句柄，是共享 frontkit 骨架绑定到 ktSpec 的别名。
type KtFrontend = frontkit.FrontendBase

// 编译期证明适配器满足冻结契约，且返回值满足 air.Frontend。
var (
	_ frontkit.LanguageSpec = (*ktSpec)(nil)
	_ air.Frontend          = (*KtFrontend)(nil)
)

// NewKtFrontend 构造 Kotlin 前端，返回满足 air.Frontend 的共享解析骨架。
//
// 骨架负责文件排序、方言归一、词法化、诊断聚合与确定性注释；本包只负责
// 词法表与 ParseTokens 的真归约。零真实文件系统读取、零子进程、零 CGO。
func NewKtFrontend() (air.Frontend, error) {
	spec := &ktSpec{
		keywords:  append([]string(nil), ktKeywords...),
		infix:     append([]ktBinding(nil), ktInfixBindings...),
		notes:     append([]ktNote(nil), ktNotes...),
		nullSafe:  append([]ktNullSafeRule(nil), ktNullSafeRules...),
		whenWords: append([]string(nil), ktWhenBranchWords...),
	}
	return frontkit.NewFrontendBase(spec)
}

// ktSpec 实现 frontkit.LanguageSpec，是注入共享解析骨架的 Kotlin 适配器。
//
// 字段可注入以便变异测试改写词法表、绑定表与标注表；零值字段回退到冻结默认值。
type ktSpec struct {
	keywords  []string
	infix     []ktBinding
	notes     []ktNote
	nullSafe  []ktNullSafeRule
	whenWords []string
}

// Tag 返回 Kotlin 的规范短标签。
func (*ktSpec) Tag() string { return "kt" }

// LexSpec 返回 Kotlin 词法表。
func (s *ktSpec) LexSpec() frontkit.LexSpec {
	return ktLexSpec{keywords: s.specKeywords()}
}

// specKeywords 返回关键字表的独立副本，空表回退默认值。
func (s *ktSpec) specKeywords() []string {
	if len(s.keywords) == 0 {
		return append([]string(nil), ktKeywords...)
	}
	return append([]string(nil), s.keywords...)
}

// NormalizeDialect 把别名归一到 kt1.6..kt2.1 六种规范方言之一，确定性。
func (*ktSpec) NormalizeDialect(dialect string) (string, error) {
	key := strings.ToLower(strings.TrimSpace(dialect))
	if canonical, ok := ktDialectFor(key); ok {
		return canonical, nil
	}
	return "", fmt.Errorf("ktfront: unsupported dialect %q (want kt1.6..kt2.1)", dialect)
}

// ParseTokens 把一文件的 token 流真归约成有层级的 AIR 节点树。
func (s *ktSpec) ParseTokens(tokens []frontkit.Token,
	path string) (air.Node, []air.Diagnostic, error) {
	return newKtParser(tokens, path, s).parse()
}
