/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package gofront

import (
	"go/ast"
	"reflect"
)

// IsNilASTNode reports whether an AST node interface holds a nil concrete value.
//
// A truncated Go source leaves nodes such as a nil block statement stored in an
// AST node interface. The go/ast walker dereferences such a value, so callers
// must reject it before any inspection. The nilable kinds are checked explicitly
// to avoid calling IsNil on a non-nilable reflection value, which would panic.
func isNilASTNode(node ast.Node) bool {
	if node == nil {
		return true
	}
	value := reflect.ValueOf(node)
	switch value.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface,
		reflect.Map, reflect.Pointer, reflect.Slice:
		return value.IsNil()
	default:
		return false
	}
}
