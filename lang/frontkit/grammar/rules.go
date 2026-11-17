/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package grammar

import "abate/lang/air"

// DefaultRules returns the upstream-to-AIR mapping table of one language.
//
// The language table comes first so a language can override a shared entry,
// and the shared table closes the list with the constructs every grammar
// spells the same way.
func DefaultRules(tag string) []MappingRule {
	var language []MappingRule
	switch tag {
	case "c":
		language = cRules()
	case "cs":
		language = csharpRules()
	case "js":
		language = javascriptRules()
	case "kt":
		language = kotlinRules()
	case "py":
		language = pythonRules()
	case "rs":
		language = rustRules()
	default:
		language = nil
	}
	return append(language, sharedRules()...)
}

// sharedRules maps the constructs every upstream grammar spells identically.
func sharedRules() []MappingRule {
	return []MappingRule{
		exact("core.identifier", air.KindIdent, "identifier"),
		exact("core.comment", air.KindComment, "comment"),
		exact("core.boolean", air.KindBoolLit, "true", "false"),
		exact("core.nil", air.KindNilLit, "null", "none", "undefined"),
	}
}
