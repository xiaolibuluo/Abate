/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package common

import (
	"testing"

	"abate/lang/rule"
)

// FuzzGenerate proves arbitrary tags and pattern fragments never panic during
// generation, validation or compilation. Only the seed corpus runs under the
// normal test gate; extended fuzzing is available through go test -fuzz.
func FuzzGenerate(f *testing.F) {
	f.Add("PY", "Ident~x")
	f.Add("", "(")
	f.Add("js", "Ident/^[a-z]+$/")
	f.Add("bad-tag", "ANY")
	f.Add("RS", "$name")
	f.Add("COM", `Ident/^[A-Z0-9_]{3,}$/`)
	f.Fuzz(func(t *testing.T, tag string, pattern string) {
		local := []Template{{
			Family:   "naming",
			Pattern:  pattern,
			Message:  "fuzz local template",
			Severity: rule.SeverityInfo,
			Category: rule.CategoryStyle,
			Positive: "Ident:bad_name",
			Negative: "Ident:goodName",
		}}
		specs := Generate(tag, local)
		checked, err := GenerateChecked(tag, local)
		if err != nil {
			if specs != nil || checked != nil {
				t.Fatalf("a rejected batch returned specs: tag=%q pattern=%q", tag, pattern)
			}
			return
		}
		if len(specs) != len(checked) {
			t.Fatalf("Generate and GenerateChecked disagree for tag %q", tag)
		}
		for i := range specs {
			if specs[i].ID != checked[i].ID {
				t.Fatalf("spec %d differs between the two generators", i)
			}
			if err := specs[i].Validate(); err != nil {
				t.Fatalf("accepted spec %s is invalid: %v", specs[i].ID, err)
			}
		}
	})
}
