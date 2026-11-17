/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package corpusgen

// goRecipe covers generics, interfaces, methods, concurrency, error handling,
// type switches, closures and embedded structs. Identifiers avoid every
// unexplained dogfood rule (MixedCaps, initialisms, receiver length, new(T),
// interface{}, goto, print, init, defer-in-loop, goroutine panic).
func goRecipe() Recipe {
	return Recipe{
		Language: "go",
		Ext:      "gotxt",
		Tags:     []string{"func", "generic", "concurrency", "controlflow", "error-handling", "interface", "closure", "method"},
		Header:   "package synth\n\n",
		AnchorClean: `func anchorProbe() bool {
	return 1 + 2 == 3
}

`,
		AnchorFull: `func anchorProbe() bool {
	return 1 + 2 == 3 && f() == g()
}

func f() int { return 1 }
func g() int { return 2 }

`,
		CleanBlock: "func step@@I@@x@@S@@() int {\n\treturn @@N@@ + @@S@@\n}\n",
		Blocks: []BlockTemplate{
			{Text: `type Store@@I@@x@@S@@[Key comparable, Value any] struct {
	items map[Key]Value
}

func (s *Store@@I@@x@@S@@[Key, Value]) Put(key Key, value Value) {
	if s.items == nil {
		s.items = make(map[Key]Value)
	}
	s.items[key] = value
}

func (s *Store@@I@@x@@S@@[Key, Value]) ValueAt(key Key) (Value, bool) {
	value, found := s.items[key]
	return value, found
}
`},
			{Text: `func relay@@I@@x@@S@@(in <-chan int, out chan<- int, done <-chan struct{}) {
	defer close(out)
	for {
		select {
		case value := <-in:
			if value >= 0 && value <= 64 {
				out <- value + @@N@@
			}
		case <-done:
			return
		}
	}
}
`},
			{Text: `func describe@@I@@x@@S@@(items []any) []string {
	result := make([]string, 0, len(items))
	for _, item := range items {
		switch value := item.(type) {
		case int:
			result = append(result, "int"+string(rune(value)))
		case string:
			result = append(result, value)
		default:
			result = append(result, "other")
		}
	}
	return result
}
`},
			{Text: `type probeErr@@I@@x@@S@@ struct {
	code int
}

func (e probeErr@@I@@x@@S@@) Error() string {
	return "probe failure"
}

func check@@I@@x@@S@@(value int) error {
	if value < 0 {
		return probeErr@@I@@x@@S@@{code: value}
	}
	return nil
}
`},
			{Text: `type Encoder@@I@@x@@S@@ interface {
	Encode(value int) int
}

type scaler@@I@@x@@S@@ struct {
	factor int
}

func (s scaler@@I@@x@@S@@) Encode(value int) int {
	return value * s.factor + @@N@@
}
`},
			{Text: `func fold@@I@@x@@S@@(values []int) int {
	total := 0
	visit := func(value int) {
		total = total + value + @@N@@
	}
	for _, value := range values {
		visit(value)
	}
	return total
}
`},
			{Text: `type base@@I@@x@@S@@ struct {
	seed int
}

type derived@@I@@x@@S@@ struct {
	base@@I@@x@@S@@
	step int
}

func (d derived@@I@@x@@S@@) Next() int {
	return d.seed + d.step + @@M@@
}
`},
			{Text: `func classify@@I@@x@@S@@(value int) int {
	defer func() {
		value = value + 1
	}()
	switch value {
	case @@N@@:
		return value * 2
	case @@M@@:
		return value + 2
	default:
		return value
	}
}
`},
		},
	}
}
