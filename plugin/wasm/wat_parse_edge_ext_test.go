/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package wasm

import "testing"

// TestParserLoopFunctionsPropagateErrors verifies malformed inputs fail closed.
func TestParserLoopFunctionsPropagateErrors(t *testing.T) {
	cases := []struct {
		name string
		src  string
		call func(*watParser) error
	}{
		{name: "skipSExpr lexer error", src: `"open`, call: func(p *watParser) error {
			return p.skipSExpr()
		}},
		{name: "parseTypeList lexer error", src: `"open`, call: func(p *watParser) error {
			var dst []valType
			return p.parseTypeList(&dst)
		}},
		{name: "parseFuncSig lexer error", src: `"open`, call: func(p *watParser) error {
			var params []valType
			var results []valType
			return p.parseFuncSig(&params, &results)
		}},
		{name: "parseFuncSig bad result", src: "(result f32))", call: func(p *watParser) error {
			var params []valType
			var results []valType
			return p.parseFuncSig(&params, &results)
		}},
		{name: "parseFuncSig unknown form", src: "(local i32", call: func(p *watParser) error {
			var params []valType
			var results []valType
			return p.parseFuncSig(&params, &results)
		}},
		{name: "parseFuncSig non-ident", src: "(7)", call: func(p *watParser) error {
			var params []valType
			var results []valType
			return p.parseFuncSig(&params, &results)
		}},
		{name: "parseSeqUntilEnd lexer error", src: `nop "open`, call: func(p *watParser) error {
			_, err := p.parseSeqUntilEnd()
			return err
		}},
		{name: "parseBlock lexer error in loop", src: `nop "open`, call: func(p *watParser) error {
			_, err := p.parseBlock("if")
			return err
		}},
		{name: "parseBlock result without close", src: "(result i32", call: func(p *watParser) error {
			_, err := p.parseBlock("block")
			return err
		}},
		{name: "parseCall lexer error", src: `"open`, call: func(p *watParser) error {
			_, err := p.parseCall()
			return err
		}},
		{name: "parseInstr lexer error", src: `"open`, call: func(p *watParser) error {
			_, err := p.parseInstr()
			return err
		}},
		{name: "parseInstrBody br without immediate", src: "nop", call: func(p *watParser) error {
			_, err := p.parseInstrBody("br")
			return err
		}},
		{name: "parseFunc lexer error", src: `"open`, call: func(p *watParser) error {
			return p.parseFunc(&watModule{})
		}},
		{name: "parseFunc non-ident sub-form", src: "(7)", call: func(p *watParser) error {
			return p.parseFunc(&watModule{})
		}},
		{name: "parseMemory lexer error", src: `"open`, call: func(p *watParser) error {
			return p.parseMemory(&watModule{})
		}},
		{name: "parseMemory lexer error after min", src: `1 "open`, call: func(p *watParser) error {
			return p.parseMemory(&watModule{})
		}},
		{name: "parseModule lexer error in loop", src: `(module "open`, call: func(p *watParser) error {
			_, err := p.parseModule()
			return err
		}},
		{name: "parseImport alias lexer error", src: `"env" "f" (func "open`, call: func(p *watParser) error {
			return p.parseImport(&watModule{})
		}},
		{name: "parseImport memory lexer error", src: `"env" "mem" (memory 1 "open`, call: func(p *watParser) error {
			return p.parseImport(&watModule{})
		}},
		{name: "parseImport missing final paren", src: `"env" "f" (func)`, call: func(p *watParser) error {
			return p.parseImport(&watModule{})
		}},
		{name: "parseImport non-ident kind", src: `"env" "f" (7)`, call: func(p *watParser) error {
			return p.parseImport(&watModule{})
		}},
		{name: "parseImport memory non-integer min", src: `"env" "mem" (memory word)`, call: func(p *watParser) error {
			return p.parseImport(&watModule{})
		}},
		{name: "parseImport memory missing outer close", src: `"env" "mem" (memory 1)`, call: func(p *watParser) error {
			return p.parseImport(&watModule{})
		}},
		{name: "parseBlock non-ident leading form", src: "(7)", call: func(p *watParser) error {
			_, err := p.parseBlock("block")
			return err
		}},
		{name: "parseModule non-ident keyword", src: "(7)", call: func(p *watParser) error {
			_, err := p.parseModule()
			return err
		}},
		{name: "parseMemory non-ident sub-form", src: "(7)", call: func(p *watParser) error {
			return p.parseMemory(&watModule{})
		}},
		{name: "parseMemory skipSExpr error", src: "(data 0", call: func(p *watParser) error {
			return p.parseMemory(&watModule{})
		}},
		{name: "parseFunc lexer error in loop", src: `nop "open`, call: func(p *watParser) error {
			return p.parseFunc(&watModule{})
		}},
		{name: "parseFunc bad sub-form immediate", src: "(i32.const)", call: func(p *watParser) error {
			return p.parseFunc(&watModule{})
		}},
	}
	for _, tc := range cases {
		p := newWatParser([]byte(tc.src))
		if err := tc.call(p); err == nil {
			t.Errorf("%s: expected an error", tc.name)
		}
	}
}
