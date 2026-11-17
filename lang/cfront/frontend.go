/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package cfront

import (
	"fmt"
	"strings"

	"abate/lang/air"
	"abate/lang/frontkit"
)

// cDefaultMaxTemplateDepth 是模板嵌套深度上限，超限时产出 warning 并降级归约。
const cDefaultMaxTemplateDepth = 16

// cSpec 实现 frontkit.LanguageSpec，是 C/C++ 前端注入共享解析骨架的适配器。
//
// 字段可注入以便变异测试改写预处理表与模板深度上限；NewCFrontend 使用默认值。
type cSpec struct {
	preprocKeywords  []string
	maxTemplateDepth int
}

// 编译期证明适配器满足冻结契约，且 NewCFrontend 的返回类型是 air.Frontend。
var _ frontkit.LanguageSpec = (*cSpec)(nil)
var _ air.Frontend = (*frontkit.FrontendBase)(nil)

// NewCFrontend 构造 C/C++ 前端，返回满足 air.Frontend 的共享解析骨架。
//
// 骨架负责文件排序、方言归一、词法化、诊断聚合与确定性注释；本包只负责
// 词法表与 ParseTokens 的真归约。零真实文件系统读取、零子进程、零 CGO。
func NewCFrontend() (air.Frontend, error) {
	spec := &cSpec{
		preprocKeywords:  append([]string(nil), cPreprocKeywords...),
		maxTemplateDepth: cDefaultMaxTemplateDepth,
	}
	base, err := frontkit.NewFrontendBase(spec)
	if err != nil {
		return nil, err
	}
	return base, nil
}

// Tag 返回 C/C++ 的规范短标签。
func (*cSpec) Tag() string { return "c" }

// LexSpec 返回 C/C++ 词法表。
func (*cSpec) LexSpec() frontkit.LexSpec { return cLexSpec{} }

// NormalizeDialect 把别名归一到 C11 / C++17 / C++20 三方言之一，确定性。
func (*cSpec) NormalizeDialect(dialect string) (string, error) {
	key := strings.ToLower(strings.TrimSpace(dialect))
	for _, entry := range cDialectAliases {
		if entry.alias == key {
			return entry.canonical, nil
		}
	}
	return "", fmt.Errorf("unsupported C/C++ dialect %q (want c11, c++17 or c++20)", dialect)
}

// ParseTokens 把一文件的 token 流真归约成有层级的 AIR 节点树。
func (s *cSpec) ParseTokens(tokens []frontkit.Token, path string) (air.Node, []air.Diagnostic, error) {
	parser := newCParser(tokens, path, s.preprocKeywords, s.maxTemplateDepth)
	return parser.parse()
}
