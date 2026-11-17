/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package gobuiltin

// egFixtures holds the Effective Go positive and negative fixtures.
func egFixtures() map[string]fixture {
	return map[string]fixture{
		"GO-EG-001": {"var foo_bar int", "var fooBar int"},
		"GO-EG-002": {"var FOO_BAR int", "var FooBar int"},
		"GO-EG-003": {"var userId int", "var userID int"},
		"GO-EG-004": {"var apiUrl string", "var apiURL string"},
		"GO-EG-005": {"var HttpClient int", "var HTTPClient int"},
		"GO-EG-006": {"var JsonData int", "var JSONData int"},
		"GO-EG-007": {"var XmlNode int", "var XMLNode int"},
		"GO-EG-008": {"var ApiKey int", "var APIKey int"},
		"GO-EG-009": {"var SqlDB int", "var SQLDB int"},
		"GO-EG-010": {"var UuidVal int", "var UUIDVal int"},
		"GO-EG-011": {"package foo_bar\n", "package foobar\n"},
		"GO-EG-012": {"func GetThing() int { return 1 }", "func Thing() int { return 1 }"},
		"GO-EG-013": {"func F() { x := 1;; _ = x }", "func F() { x := 1; _ = x }"},
		"GO-EG-014": {"func F() { panic(\"x\") }", "func F() { _ = \"x\" }"},
		"GO-EG-015": {"func F() { println(\"x\") }", "func F() { _ = \"x\" }"},
		"GO-EG-016": {"func F() { _ = new(int) }", "func F() { _ = 0 }"},
		"GO-EG-017": {
			"func F() { for i := 0; i < 2; i++ { defer G() } }",
			"func F() { defer G(); for i := 0; i < 2; i++ {} }",
		},
		"GO-EG-018": {"func F() (x int) { return }", "func F() int { return 1 }"},
		"GO-EG-019": {"func F() { select {} }", "func F() { select { default: } }"},
		"GO-EG-020": {"func F() { goto End; End: return }", "func F() { return }"},
		"GO-EG-021": {"func F() { var x interface{}; _ = x }", "func F() { var x any; _ = x }"},
		"GO-EG-022": {"func init() {}", "func Setup() {}"},
		"GO-EG-023": {
			"func F(s []int) { append(s, 1) }",
			"func F(s []int) { s = append(s, 1); _ = s }",
		},
		"GO-EG-024": {"var x = 12345", "var x = 8"},
		"GO-EG-025": {
			"func F() error { return errors.New(\"Bad input\") }",
			"func F() error { return errors.New(\"bad input\") }",
		},
		"GO-EG-026": {
			"func F(xs []int) { for range xs { defer G() } }",
			"func F(xs []int) { defer G(); for range xs {} }",
		},
		"GO-EG-027": {"func F() { _ = recover() }", "func F() { _ = 1 }"},
		"GO-EG-028": {
			"type T struct{}\nfunc (this *T) M() {}",
			"type T struct{}\nfunc (t *T) M() {}",
		},
		"GO-EG-029": {
			"func F() { s := \"\"; for i := 0; i < 2; i++ { s = s + \"x\" }; _ = s }",
			"func F() { s := \"x\"; _ = s }",
		},
		"GO-EG-030": {
			"package p\n// bad doc\nfunc G() {}\n",
			"package p\n// Good doc\nfunc G() {}\n",
		},
		"GO-EG-031": {
			"func F() { _ = make([]int, 0) }",
			"func F() { _ = make([]int, 4) }",
		},
		"GO-EG-032": {"type T struct{}", "type T struct{ X int }"},
		"GO-EG-033": {
			"func F() { go func() { panic(\"x\") }() }",
			"func F() { go func() { _ = 1 }() }",
		},
	}
}
