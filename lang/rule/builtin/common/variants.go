/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package common

// variant is one pattern/where pair inside a theme family.
type variant struct {
	pattern string
	where   []string
}

// templateCount is the frozen size of the common template family.
const templateCount = familyCount * 5

// variantGroups rebuilds the twelve ordered five-variant groups.
//
// The table is a deterministic slice of slices: no map iteration ever
// contributes to the generated rule order. Group i belongs to families()[i].
func variantGroups() [][]variant {
	return [][]variant{
		{
			{pattern: `Ident/^[a-z]+_[a-z]+$/`},
			{pattern: `Ident~_`},
			{pattern: `Ident^bad_`},
			{pattern: `$name`, where: []string{"regex:_[a-z]+$"}},
			{pattern: `Ident/^[a-z_]{4,}$/`},
		},
		{
			{pattern: `Ident/^$/`},
			{pattern: `$name`, where: []string{"regex:^$"}},
			{pattern: `Ident/^.{0,0}$/`},
			{pattern: `Ident!=value`},
			{pattern: `$name`, where: []string{"count-min:1", "regex:^$"}},
		},
		{
			{pattern: `NumLit/^[0-9]{4,}$/`},
			{pattern: `$n`, where: []string{"num-min:100"}},
			{pattern: `NumLit~00`},
			{pattern: `NumLit/^1[0-9]{3}$/`},
			{pattern: `NumLit+000`},
		},
		{
			{pattern: `Ident=_`},
			{pattern: `Ident^_`},
			{pattern: `Ident/^_{1}$/`},
			{pattern: `$name`, where: []string{"text:_"}},
			{pattern: `Ident/^[^a-zA-Z]+$/`},
		},
		{
			{pattern: `Binary===`},
			{pattern: `Binary/^==$/`},
			{pattern: `$op`, where: []string{"text:=="}},
			{pattern: `Binary/^={2}$/`},
			{pattern: `Binary/^[^=]*==[^=]*$/`},
		},
		{
			{pattern: `Star~unsafe`},
			{pattern: `Star^uns`},
			{pattern: `$ptr`, where: []string{"text~:unsafe"}},
			{pattern: `Star/^[a-z]*unsafe[a-z]*$/`},
			{pattern: `Star!=safe`},
		},
		{
			{pattern: `Binary+==true`},
			{pattern: `Binary~==true`},
			{pattern: `Binary/^.*==true$/`},
			{pattern: `$cmp`, where: []string{"regex:==true$"}},
			{pattern: `Binary^flag==`},
		},
		{
			{pattern: `Ident/^[A-Z0-9_]{3,}$/`},
			{pattern: `Ident~_`},
			{pattern: `Ident^MAX`},
			{pattern: `$name`, where: []string{"regex:^[A-Z_]+$"}},
			{pattern: `Ident/^[A-Z]+_[A-Z]+$/`},
		},
		{
			{pattern: `FuncDecl~arg`},
			{pattern: `FuncDecl^many_`},
			{pattern: `FuncDecl/^[a-z_]{10,}$/`},
			{pattern: `$fn`, where: []string{"regex:_[a-z]+$"}},
			{pattern: `FuncDecl+ments`},
		},
		{
			{pattern: `Ident^_`},
			{pattern: `Ident~unused`},
			{pattern: `$name`, where: []string{"text^:_"}},
			{pattern: `Ident/^_[a-z]+$/`},
			{pattern: `Ident!=used`},
		},
		{
			{pattern: `Comment=stub`},
			{pattern: `Comment^st`},
			{pattern: `$note`, where: []string{"text:stub"}},
			{pattern: `Comment/^stub$/`},
			{pattern: `Comment!=explain`},
		},
		{
			{pattern: `Binary+<=0`},
			{pattern: `Binary~<=0`},
			{pattern: `Binary/^.*<=[0-9]+$/`},
			{pattern: `$cmp`, where: []string{"regex:<=[0-9]+$"}},
			{pattern: `Binary^index<=`},
		},
	}
}

// builtinTemplates rebuilds the sixty common templates in declaration order.
func builtinTemplates() []Template {
	fams := families()
	groups := variantGroups()
	out := make([]Template, 0, templateCount)
	for i, f := range fams {
		for _, v := range groups[i] {
			out = append(out, newTemplate(f, v))
		}
	}
	return out
}
