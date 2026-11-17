/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package gobuiltin

// crcFixtures holds the Go Code Review Comments positive and negative fixtures.
func crcFixtures() map[string]fixture {
	return map[string]fixture{
		"GO-CRC-001": {"var TcpConn int", "var TCPConn int"},
		"GO-CRC-002": {"var IpAddr int", "var IPAddr int"},
		"GO-CRC-003": {"var CpuTime int", "var CPUTime int"},
		"GO-CRC-004": {"var Utf8Text int", "var UTF8Text int"},
		"GO-CRC-005": {"var AsciiText int", "var ASCIIText int"},
		"GO-CRC-006": {"var CliArgs int", "var CLIArgs int"},
		"GO-CRC-007": {"var GuiView int", "var GUIView int"},
		"GO-CRC-008": {"var item01 int", "var item int"},
		"GO-CRC-009": {"var s = \"Hello\"", "var s = \"hello\""},
		"GO-CRC-010": {"package Foo\n", "package foo\n"},
		"GO-CRC-011": {"var CsvRow int", "var CSVRow int"},
		"GO-CRC-012": {
			"func F() error { return fmt.Errorf(\"Bad\") }",
			"func F() error { return fmt.Errorf(\"bad\") }",
		},
		"GO-CRC-013": {"import _ \"os\"", "import \"os\""},
		"GO-CRC-014": {"func F() { print(\"x\") }", "func F() { _ = \"x\" }"},
		"GO-CRC-015": {
			"func F() { _ = G() }",
			"func F() { v := G(); if v != nil { return } }",
		},
		"GO-CRC-016": {"import . \"fmt\"", "import \"fmt\""},
		"GO-CRC-017": {"func F() (x int) { return }", "func F() int { return 1 }"},
		"GO-CRC-018": {
			"type T struct{}\nfunc (self *T) M() {}",
			"type T struct{}\nfunc (t *T) M() {}",
		},
		"GO-CRC-019": {
			"// bad line\npackage p\n",
			"// Good line\npackage p\n",
		},
		"GO-CRC-020": {
			"func F() { _ = reflect.DeepEqual(1, 2) }",
			"func F() { _ = 1 }",
		},
		"GO-CRC-021": {"func F() { time.Sleep(1) }", "func F() { _ = 1 }"},
		"GO-CRC-022": {
			"type T struct{}\nfunc (t *T) M() {}",
			"type T struct{}\nfunc (t T) M() {}",
		},
		"GO-CRC-023": {"import \"Example.com/x\"", "import \"example.com/x\""},
		"GO-CRC-024": {"var s = \"hello \"", "var s = \"hello\""},
		"GO-CRC-025": {
			"func F() error { return errors.New(\"bad.\") }",
			"func F() error { return errors.New(\"bad\") }",
		},
		"GO-CRC-026": {
			"func F() { _ = regexp.MustCompile(\"x\") }",
			"func F() { _, _ = regexp.Compile(\"x\") }",
		},
		"GO-CRC-027": {"func F() { _ = time.Now() }", "func F() { _ = 1 }"},
		"GO-CRC-028": {"func F() { _ = os.Getenv(\"X\") }", "func F() { _ = \"X\" }"},
		"GO-CRC-029": {"func F() { os.Exit(1) }", "func F() { return }"},
		"GO-CRC-030": {
			"func F(t *T) { t.Fatal(\"no\") }",
			"func F(t *T) { t.Log(\"ok\") }",
		},
		"GO-CRC-031": {"func F() { _ = os.Setenv(\"X\", \"1\") }", "func F() { _ = 1 }"},
		"GO-CRC-032": {
			"func F() { _, _ = os.ReadFile(\"x\") }",
			"func F() { _ = \"x\" }",
		},
		"GO-CRC-033": {
			"func F() { _, _ = http.Post(\"http://x\") }",
			"func F() { _, _ = http.Head(\"http://x\") }",
		},
	}
}
