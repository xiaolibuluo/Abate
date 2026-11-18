/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package check

import (
	"abate/lang/air"
	"abate/lang/rule"
)

// LanguageSpec binds one AIR frontend, its compiled rule set and the source
// file extensions that route to it. A MultiChecker is composed entirely from
// these values, keeping the orchestration layer free of any language package
// import (the CLI owns the concrete frontend/rule construction).
type LanguageSpec struct {
	// Language is the lowercase tag written into each source unit and expected
	// by rule.Spec.Language() filters (go, py, js, c, cs, rs, kt).
	Language string
	// Frontend parses this language into AIR.
	Frontend air.Frontend
	// Rules is the compiled rule set for this language.
	Rules []rule.Rule
	// Extensions are lowercased file suffixes including the dot (.go, .py...).
	Extensions []string
}
