/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package wasm

import (
	"errors"
	"fmt"
)

// ParseFuncSig reads param and result sub-forms until the closing parenthesis.
func (p *watParser) parseFuncSig(params, results *[]valType) error {
	for {
		tt, err := p.peekTok()
		if err != nil {
			return err
		}
		if tt.kind == tokRParen {
			_, _ = p.next()
			return nil
		}
		if err := p.expect(tokLParen); err != nil {
			return err
		}
		t, err := p.nextIdent()
		if err != nil {
			return err
		}
		switch t {
		case "param":
			if err := p.parseTypeList(params); err != nil {
				return err
			}
		case "result":
			if err := p.parseTypeList(results); err != nil {
				return err
			}
		default:
			if err := p.skipSExpr(); err != nil {
				return err
			}
		}
	}
}

// ParseTypeList reads a list of value types until the closing ')'.
func (p *watParser) parseTypeList(dst *[]valType) error {
	for {
		tt, err := p.peekTok()
		if err != nil {
			return err
		}
		if tt.kind == tokRParen {
			_, _ = p.next()
			return nil
		}
		vt, err := p.nextType()
		if err != nil {
			return err
		}
		*dst = append(*dst, vt)
	}
}

// NextType consumes a value-type identifier.
func (p *watParser) nextType() (valType, error) {
	t, err := p.nextIdent()
	if err != nil {
		return 0, err
	}
	switch t {
	case "i32":
		return vtI32, nil
	case "i64":
		return vtI64, nil
	default:
		return 0, fmt.Errorf("wasm: unsupported value type %q", t)
	}
}

// SkipSExpr consumes a balanced s-expression starting just after '('.
func (p *watParser) skipSExpr() error {
	depth := 1
	for {
		t, err := p.next()
		if err != nil {
			return err
		}
		switch t.kind {
		case tokEOF:
			return errors.New("wasm: unexpected EOF in form")
		case tokLParen:
			depth++
		case tokRParen:
			depth--
			if depth == 0 {
				return nil
			}
		}
	}
}

// Parses one directive sub-form inside a func (export, param, result, local). It
// reports whether the sub-form was a directive; any other parenthesized
// sub-form is an instruction handled by the caller.
func (p *watParser) parseFuncDirective(f *watFunc, t string) (bool, error) {
	switch t {
	case "export":
		name, err := p.nextString()
		if err != nil {
			return false, err
		}
		f.exportName = name
		if err := p.expect(tokRParen); err != nil {
			return false, err
		}
		return true, nil
	case "param":
		if err := p.parseTypeList(&f.params); err != nil {
			return false, err
		}
		return true, nil
	case "result":
		if err := p.parseTypeList(&f.results); err != nil {
			return false, err
		}
		return true, nil
	case "local":
		if err := p.parseTypeList(&f.locals); err != nil {
			return false, err
		}
		return true, nil
	default:
		return false, nil
	}
}
