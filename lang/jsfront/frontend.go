/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package jsfront

import (
	"fmt"
	"strings"

	"abate/lang/air"
	"abate/lang/frontkit"
)

// jsSpec 实现 frontkit.LanguageSpec，是 LNG-JS 前端的唯一语言钩子.
type jsSpec struct{}

// 编译期断言 jsSpec 满足冻结的语言契约.
var _ frontkit.LanguageSpec = (*jsSpec)(nil)

// NewJsFrontend 返回满足 air.Frontend 的 JS/TS 前端，内部复用 frontkit 骨架.
func NewJsFrontend() (air.Frontend, error) {
	base, err := frontkit.NewFrontendBase(&jsSpec{})
	if err != nil {
		return nil, err
	}
	return base, nil
}

// Tag 返回规范短标签 "js".
func (*jsSpec) Tag() string { return "js" }

// LexSpec 返回 JS/TS 词法表.
func (*jsSpec) LexSpec() frontkit.LexSpec { return jsLexSpec{} }

// ParseTokens 把 token 流真归约为有层级结构的 AIR 树，永不因畸形输入 panic.
func (*jsSpec) ParseTokens(tokens []frontkit.Token, path string) (air.Node,
	[]air.Diagnostic, error) {
	return newJsParser(tokens, path).parse()
}

// NormalizeDialect 把扩展名、strict 与 target 组合归一为确定性规范字符串.
func (*jsSpec) NormalizeDialect(d string) (string, error) {
	parsed, err := parseDialectString(d)
	if err != nil {
		return "", err
	}
	return parsed.canonical(), nil
}

// jsFlavorAlias 把方言别名归一到六种规范风味之一.
func jsFlavorAlias(text string) (string, bool) {
	switch text {
	case ".js", "js", "javascript", "ecmascript":
		return "js", true
	case ".jsx", "jsx":
		return "jsx", true
	case ".mjs", "mjs":
		return "mjs", true
	case ".cjs", "cjs", "commonjs":
		return "cjs", true
	case ".ts", "ts", "typescript":
		return "ts", true
	case ".tsx", "tsx":
		return "tsx", true
	default:
		return "", false
	}
}

// parseDialectString 解析方言描述串，纯 ASCII、纯函数、无全局状态.
func parseDialectString(raw string) (jsDialect, error) {
	d := jsDialect{flavor: "js", module: "auto", target: "es2022"}
	if len(raw) > 256 {
		return d, fmt.Errorf("jsfront: dialect description too long: %d", len(raw))
	}
	if err := jsValidateDialectBytes(raw); err != nil {
		return d, err
	}
	fields := strings.FieldsFunc(raw, jsIsDialectSeparator)
	explicitModule, explicitStrict := jsApplyDialectFields(&d, fields)
	if !explicitModule {
		jsInferDialectModule(&d)
	}
	if !explicitStrict {
		jsInferDialectStrict(&d)
	}
	return d, nil
}

// jsValidateDialectBytes 校验方言描述串仅含可打印 ASCII 与制表符.
func jsValidateDialectBytes(raw string) error {
	for i := 0; i < len(raw); i++ {
		c := raw[i]
		if (c < 0x20 && c != '\t') || c > 0x7e {
			return fmt.Errorf("jsfront: invalid dialect byte at %d", i)
		}
	}
	return nil
}

// jsIsDialectSeparator 判定方言描述串中的字段分隔字符.
func jsIsDialectSeparator(r rune) bool {
	return r == ' ' || r == ',' || r == ';' || r == ':' || r == '|' ||
		r == '/' || r == '\t'
}

// jsApplyDialectFields 逐字段应用风味、strict、module 与 target，返回显式设定标记.
func jsApplyDialectFields(d *jsDialect,
	fields []string) (explicitModule bool, explicitStrict bool) {
	for _, field := range fields {
		token := strings.ToLower(strings.TrimSpace(field))
		if token == "" {
			continue
		}
		if flavor, ok := jsFlavorAlias(token); ok {
			d.flavor = flavor
		}
		switch token {
		case "strict", "use-strict":
			d.strict, explicitStrict = true, true
		case "sloppy", "nonstrict", "no-strict":
			d.strict, explicitStrict = false, true
		case "module=esm", "esm":
			d.module, explicitModule = "esm", true
		case "module=cjs", "module=commonjs":
			d.module, explicitModule = "commonjs", true
		case "module=auto":
			d.module, explicitModule = "auto", true
		}
		target := strings.TrimPrefix(token, "target=")
		if jsInStringSlice(jsTargets, target) {
			d.target = target
		}
	}
	return explicitModule, explicitStrict
}

// jsInferDialectModule 按规范风味推导未显式设定的 module 制式.
func jsInferDialectModule(d *jsDialect) {
	switch d.flavor {
	case "mjs", "ts", "tsx":
		d.module = "esm"
	case "cjs":
		d.module = "commonjs"
	default:
		d.module = "auto"
	}
}

// jsInferDialectStrict 以 module 是否为 esm 推导未显式设定的 strict 标记.
func jsInferDialectStrict(d *jsDialect) {
	d.strict = d.module == "esm"
}
