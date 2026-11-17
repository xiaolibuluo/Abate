/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package wasm

import (
	"fmt"
	"strings"
)

// watParser performs a recursive-descent parse of the supported WAT subset.
type watParser struct {
	lex   *watLexer
	peek  *watToken
	names map[string]int
	impN  int
}

// newWatParser builds a parser over WAT source.
func newWatParser(src []byte) *watParser {
	return &watParser{lex: newWatLexer(src), names: map[string]int{}}
}

// next returns the next token, honoring a single-token lookahead.
func (p *watParser) next() (*watToken, error) {
	if p.peek != nil {
		t := p.peek
		p.peek = nil
		return t, nil
	}
	return p.lex.next()
}

// peekTok returns the next token without consuming it.
func (p *watParser) peekTok() (*watToken, error) {
	if p.peek == nil {
		t, err := p.lex.next()
		if err != nil {
			return nil, err
		}
		p.peek = t
	}
	return p.peek, nil
}

// expect consumes a token of the expected kind or returns an error.
func (p *watParser) expect(kind watTokKind) error {
	t, err := p.next()
	if err != nil {
		return err
	}
	if t.kind != kind {
		return fmt.Errorf("wasm: line %d: expected token %d, got %q", t.line, kind, t.text)
	}
	return nil
}

// nextIdent consumes an identifier token and returns its text.
func (p *watParser) nextIdent() (string, error) {
	t, err := p.next()
	if err != nil {
		return "", err
	}
	if t.kind != tokIdent {
		return "", fmt.Errorf("wasm: line %d: expected identifier, got %q", t.line, t.text)
	}
	return t.text, nil
}

// nextString consumes a string token and returns its text.
func (p *watParser) nextString() (string, error) {
	t, err := p.next()
	if err != nil {
		return "", err
	}
	if t.kind != tokString {
		return "", fmt.Errorf("wasm: line %d: expected string, got %q", t.line, t.text)
	}
	return t.text, nil
}

// nextInt consumes an integer token and returns its value.
func (p *watParser) nextInt() (int64, error) {
	t, err := p.next()
	if err != nil {
		return 0, err
	}
	if t.kind != tokInt {
		return 0, fmt.Errorf("wasm: line %d: expected integer, got %q", t.line, t.text)
	}
	return t.value, nil
}

// ParseModule parses the top-level module form.
func (p *watParser) parseModule() (*watModule, error) {
	mod := &watModule{}
	if err := p.expect(tokLParen); err != nil {
		return nil, err
	}
	kw, err := p.nextIdent()
	if err != nil {
		return nil, err
	}
	if kw != "module" {
		return nil, fmt.Errorf("wasm: expected 'module', got %q", kw)
	}
	for {
		tt, err := p.peekTok()
		if err != nil {
			return nil, err
		}
		if tt.kind == tokRParen {
			_, _ = p.next()
			return mod, nil
		}
		if err := p.parseTopLevel(mod); err != nil {
			return nil, err
		}
	}
}

// parseTopLevel dispatches a top-level module form.
func (p *watParser) parseTopLevel(mod *watModule) error {
	if err := p.expect(tokLParen); err != nil {
		return err
	}
	kw, err := p.nextIdent()
	if err != nil {
		return err
	}
	switch kw {
	case "memory":
		return p.parseMemory(mod)
	case "import":
		return p.parseImport(mod)
	case "func":
		return p.parseFunc(mod)
	default:
		return p.skipSExpr()
	}
}

// ParseMemory parses a memory form for defined memory.
func (p *watParser) parseMemory(mod *watModule) error {
	for {
		tt, err := p.peekTok()
		if err != nil {
			return err
		}
		if tt.kind == tokRParen {
			_, _ = p.next()
			return nil
		}
		if tt.kind == tokLParen {
			_, _ = p.next()
			t, err := p.nextIdent()
			if err != nil {
				return err
			}
			if t == "export" {
				name, err := p.nextString()
				if err != nil {
					return err
				}
				mod.memExport = name
				if err := p.expect(tokRParen); err != nil {
					return err
				}
				continue
			}
			if err := p.skipSExpr(); err != nil {
				return err
			}
			continue
		}
		v, err := p.nextInt()
		if err != nil {
			return err
		}
		mn, err := toU32(v)
		if err != nil {
			return err
		}
		mod.memMin = mn
		tt2, err := p.peekTok()
		if err != nil {
			return err
		}
		if tt2.kind == tokInt {
			mx, err := p.nextInt()
			if err != nil {
				return err
			}
			mxv, err := toU32(mx)
			if err != nil {
				return err
			}
			mod.memMax = mxv
		}
	}
}

// ParseImport parses an import form with a module name and signature.
func (p *watParser) parseImport(mod *watModule) error {
	modName, err := p.nextString()
	if err != nil {
		return err
	}
	impName, err := p.nextString()
	if err != nil {
		return err
	}
	if err := p.expect(tokLParen); err != nil {
		return err
	}
	t, err := p.nextIdent()
	if err != nil {
		return err
	}
	switch t {
	case "func":
		fi := watFuncImport{module: modName, name: impName}
		if err := p.bindImportAlias(mod); err != nil {
			return err
		}
		if err := p.parseFuncSig(&fi.params, &fi.results); err != nil {
			return err
		}
		mod.imports = append(mod.imports, fi)
		p.impN = len(mod.imports)
		return p.expect(tokRParen)
	case "memory":
		mn, err := p.nextInt()
		if err != nil {
			return err
		}
		mnv, err := toU32(mn)
		if err != nil {
			return err
		}
		mod.memImport = &watMemImport{module: modName, name: impName}
		mod.memMin = mnv
		tt, err := p.peekTok()
		if err != nil {
			return err
		}
		if tt.kind == tokInt {
			mx, err := p.nextInt()
			if err != nil {
				return err
			}
			mxv, err := toU32(mx)
			if err != nil {
				return err
			}
			mod.memMax = mxv
		}
		if err := p.expect(tokRParen); err != nil {
			return err
		}
		return p.expect(tokRParen)
	default:
		return p.skipSExpr()
	}
}

// bindImportAlias records a $alias on the current (not-yet-appended) import.
func (p *watParser) bindImportAlias(mod *watModule) error {
	tt, err := p.peekTok()
	if err != nil {
		return err
	}
	if tt.kind == tokIdent && strings.HasPrefix(tt.text, "$") {
		_, _ = p.next()
		p.names[tt.text] = len(mod.imports)
	}
	return nil
}
