/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package gobuiltin

// vetFixtures holds the vet semantic subset positive and negative fixtures.
func vetFixtures() map[string]fixture {
	return map[string]fixture{
		"GO-VET-001": {
			"func F() { fmt.Printf(\"v %d\") }",
			"func F() { fmt.Printf(\"v %d\", 1) }",
		},
		"GO-VET-002": {
			"func F() string { return fmt.Sprintf(\"v %d\") }",
			"func F() string { return fmt.Sprintf(\"v %d\", 1) }",
		},
		"GO-VET-003": {
			"func F() error { return fmt.Errorf(\"v %d\") }",
			"func F() error { return fmt.Errorf(\"v %d\", 1) }",
		},
		"GO-VET-004": {
			"func F(t *T) { t.Fatalf(\"v %d\") }",
			"func F(t *T) { t.Fatalf(\"v %d\", 1) }",
		},
		"GO-VET-005": {
			"func F(l *L) { l.Logf(\"v %d\") }",
			"func F(l *L) { l.Logf(\"v %d\", 1) }",
		},
		"GO-VET-006": {
			"func F() { fmt.Println(\"v %d\") }",
			"func F() { fmt.Println(\"v\") }",
		},
		"GO-VET-007": {
			"func F() { fmt.Print(\"v %d\") }",
			"func F() { fmt.Print(\"v\") }",
		},
		"GO-VET-008": {
			"func F() string { return fmt.Sprint(\"v %d\") }",
			"func F() string { return fmt.Sprint(\"v\") }",
		},
		"GO-VET-009": {
			"func F(err error) error { return fmt.Errorf(\"wrap %v\", err) }",
			"func F(err error) error { return fmt.Errorf(\"wrap %w\", err) }",
		},
		"GO-VET-010": {
			"func F(err error) { var t error; _ = errors.As(err, &t) }",
			"func F(err error) { _ = errors.Is(err, nil) }",
		},
		"GO-VET-011": {
			"func F(err error) { _ = errors.Unwrap(err) }",
			"func F(err error) { _ = err }",
		},
		"GO-VET-012": {
			"func F(s []int) { _ = append(s) }",
			"func F(s []int) { _ = append(s, 1) }",
		},
		"GO-VET-013": {
			"func F(s []int) { _ = copy(s) }",
			"func F(s, d []int) { _ = copy(d, s) }",
		},
		"GO-VET-014": {"func F() string { return string(65) }", "func F() string { return \"A\" }"},
		"GO-VET-015": {"func F() int { return int(\"x\") }", "func F() int { return 1 }"},
		"GO-VET-016": {
			"func F() { _ = make([]int) }",
			"func F() { _ = make([]int, 2) }",
		},
		"GO-VET-017": {
			"func F() string { return strings.Replace(\"a\", \"a\", \"b\", 0) }",
			"func F() string { return strings.Replace(\"a\", \"a\", \"b\", 1) }",
		},
		"GO-VET-018": {
			"func F() []string { return strings.Split(\"a\", \",\", 0) }",
			"func F() []string { return strings.Split(\"a\", \",\") }",
		},
		"GO-VET-019": {"func F(w *W) { w.Wait() }", "func F(w *W) { w.Add(1) }"},
		"GO-VET-020": {"func F(w *W) { w.Add(0) }", "func F(w *W) { w.Add(1) }"},
		"GO-VET-021": {"func F() { time.Sleep(0) }", "func F() { time.Sleep(1) }"},
		"GO-VET-022": {"func F(c C) { _ = c.Done() }", "func F(c C) { _ = c.Err() }"},
		"GO-VET-023": {
			"type T struct { X int \x60json:\"x\"\x60 }",
			"type T struct { X int }",
		},
		"GO-VET-024": {
			"func F(b []byte, v any) { _ = json.Unmarshal(b, v) }",
			"func F(b []byte) { _ = b }",
		},
		"GO-VET-025": {
			"func F(v any) { _, _ = json.Marshal(v) }",
			"func F(v any) { _ = v }",
		},
		"GO-VET-026": {
			"func F() { _, _ = http.Get(\"http://x\") }",
			"func F() { _, _ = http.Head(\"http://x\") }",
		},
		"GO-VET-027": {
			"func F(t T) string { return t.Format(\"2006\") }",
			"func F(t T) string { return t.Format(\"Jan\") }",
		},
		"GO-VET-028": {
			"func F(s string) { _, _ = time.Parse(\"2006\", s) }",
			"func F(s string) { _, _ = time.Parse(\"Jan\", s) }",
		},
		"GO-VET-029": {
			"func F(l L) L { return l.With(\"k\", 1) }",
			"func F(l L) L { return l.Without(\"k\") }",
		},
		"GO-VET-030": {"func F(l L) { l.Log(\"v %d\") }", "func F(l L) { l.Log(\"v\") }"},
		"GO-VET-031": {"func F(l L) { l.LogAttrs(1) }", "func F(l L) { l.Log() }"},
		"GO-VET-032": {"func F(l L) { l.Info(\"v %d\") }", "func F(l L) { l.Info(\"v\") }"},
		"GO-VET-033": {"func F(l L) { l.Warn(\"v %d\") }", "func F(l L) { l.Warn(\"v\") }"},
		"GO-VET-034": {
			"func F() int { return cap([]int{1}) }",
			"func F(s []int) int { return cap(s) }",
		},
		"GO-VET-035": {
			"func F() int { return len([]int{1}) }",
			"func F(s []int) int { return len(s) }",
		},
		"GO-VET-036": {
			"func F(p *int) bool { return p == nil }",
			"func F(p *int) bool { return p != nil }",
		},
		"GO-VET-037": {
			"func F(p *int) bool { return p != nil }",
			"func F(p *int) bool { return p == nil }",
		},
		"GO-VET-038": {
			"func F(b bool) bool { return b == true }",
			"func F(b bool) bool { return b }",
		},
		"GO-VET-039": {"func F(r R) { r.Copy() }", "func F(r R) { r.Clone() }"},
		"GO-VET-040": {"func F(m M) { m.Lock() }", "func F(m M) { m.TryLock() }"},
		"GO-VET-041": {"func F(m M) { m.Unlock() }", "func F(m M) { m.TryLock() }"},
		"GO-VET-042": {"func F(c C) { _ = c.Close() }", "func F(c C) { _ = c.Open() }"},
		"GO-VET-043": {"func F() {}", "func F() { _ = 1 }"},
		"GO-VET-044": {"func F(r C) { defer r.Close() }", "func F(r C) { _ = r }"},
	}
}
