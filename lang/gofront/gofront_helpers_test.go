/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package gofront

import (
	"testing"

	"abate/lang/air"
)

// sampleGo exercises the full Go mapping table in one compilable file.
const sampleGo = `package sample

import (
	"fmt"
	"os"
)

//go:noinline
// Doc explains the constants.
const (
	First = iota
	Second
)

type Box[T any] struct {
	Value T ` + "`json:\"value\"`" + `
}

type Reader interface {
	Read(p []byte) (int, error)
}

var Global = map[string]int{"a": 1}

func (b *Box[T]) Get() T {
	return b.Value
}

func Run(ch chan int, done func()) {
	defer done()
	go func() { ch <- 1 }()
	for k, v := range Global {
		fmt.Println(k, v)
	}
	var value any = 7
	switch x := value.(type) {
	case int:
		_ = x
	default:
		break
	}
	select {
	case n := <-ch:
		_ = n
	case ch <- 2:
	default:
	}
	if len(os.Args) > 0 {
		return
	}
outer:
	for i := 0; i < 2; i++ {
		if i == 1 {
			break outer
		}
	}
	s := []int{1, 2, 3}
	_ = s[0:2]
	_ = &value
	_ = *(&value)
	_ = func(a int) int { return a + 1 }(1)
	_ = Box[int]{Value: 1}
	_ = []int{1, 2}[0]
	_ = [...]int{1, 2}
	_ = make(chan<- int)
	_ = (<-chan int)(nil)
	neg := -1
	not := !true
	_ = neg
	_ = not
}

// standalone trailing comment.
`

// parseSource parses one in-memory Go file with the default frontend.
func parseSource(t *testing.T, src string) (*air.Unit, []air.Diagnostic) {
	t.Helper()
	unit := air.SourceUnit{
		Language: "go",
		Files: []air.SourceFile{
			air.NewSourceFile("sample.go", []byte(src)),
		},
	}
	parsed, diagnostics, err := NewGoFrontend().Parse(unit)
	if err != nil {
		t.Fatalf("Parse: %v", err)
	}
	return parsed, diagnostics
}

// mustParse parses a source string that must be syntactically clean.
func mustParse(t *testing.T, src string) *air.Unit {
	t.Helper()
	unit, diagnostics := parseSource(t, src)
	for _, diagnostic := range diagnostics {
		if diagnostic.IsError() {
			t.Fatalf("unexpected syntax error: %s", diagnostic.String())
		}
	}
	return unit
}

// collectKinds returns every node of the requested kind in a unit.
func collectKinds(unit *air.Unit, kind air.NodeKind) []air.Node {
	out := make([]air.Node, 0)
	for _, current := range air.Collect(unit) {
		if current.Kind() == kind {
			out = append(out, current)
		}
	}
	return out
}

// hasKind reports whether a kind appears anywhere in the unit.
func hasKind(unit *air.Unit, kind air.NodeKind) bool {
	return len(collectKinds(unit, kind)) > 0
}

// annotationOf returns the first annotation value found on nodes of a kind.
func annotationOf(unit *air.Unit, kind air.NodeKind, key air.SemKey) string {
	for _, current := range collectKinds(unit, kind) {
		if value := current.Sem(key); value != nil {
			if text, ok := value.(string); ok {
				return text
			}
		}
	}
	return ""
}
