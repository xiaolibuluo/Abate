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

// ParseFunc parses a func form, recording its name and body.
func (p *watParser) parseFunc(mod *watModule) error {
	f := watFunc{}
	if err := p.recordFuncName(mod); err != nil {
		return err
	}
	for {
		tt, err := p.peekTok()
		if err != nil {
			return err
		}
		if tt.kind == tokRParen {
			_, _ = p.next()
			mod.funcs = append(mod.funcs, f)
			return nil
		}
		if tt.kind == tokLParen {
			_, _ = p.next()
			t, err := p.nextIdent()
			if err != nil {
				return err
			}
			handled, err := p.parseFuncDirective(&f, t)
			if err != nil {
				return err
			}
			if handled {
				continue
			}
			instr, err := p.parseInstrBody(t)
			if err != nil {
				return err
			}
			if err := p.expect(tokRParen); err != nil {
				return err
			}
			f.body = append(f.body, instr)
			continue
		}
		instr, err := p.parseInstr()
		if err != nil {
			return err
		}
		f.body = append(f.body, instr)
	}
}

// recordFuncName binds a $alias to the defined function's eventual index.
func (p *watParser) recordFuncName(mod *watModule) error {
	tt, err := p.peekTok()
	if err != nil {
		return err
	}
	if tt.kind != tokIdent || !strings.HasPrefix(tt.text, "$") {
		return nil
	}
	_, _ = p.next()
	p.names[tt.text] = p.impN + len(mod.funcs)
	return nil
}

// parseInstr parses a single instruction in either flat or parenthesized form.
func (p *watParser) parseInstr() (watInstr, error) {
	t, err := p.next()
	if err != nil {
		return watInstr{}, err
	}
	if t.kind == tokLParen {
		op, err := p.nextIdent()
		if err != nil {
			return watInstr{}, err
		}
		instr, err := p.parseInstrBody(op)
		if err != nil {
			return watInstr{}, err
		}
		if err := p.expect(tokRParen); err != nil {
			return watInstr{}, err
		}
		return instr, nil
	}
	if t.kind == tokIdent {
		return p.parseInstrBody(t.text)
	}
	return watInstr{}, fmt.Errorf("wasm: line %d: expected instruction, got %q", t.line, t.text)
}

// parseInstrBody parses the immediate operands of one instruction.
func (p *watParser) parseInstrBody(op string) (watInstr, error) {
	switch op {
	case "if", "loop", "block":
		return p.parseBlock(op)
	case "end", "else":
		return watInstr{}, fmt.Errorf("wasm: unexpected control token %q", op)
	}
	instr := watInstr{op: op}
	switch op {
	case "i32.const", "i64.const":
		v, err := p.nextInt()
		if err != nil {
			return instr, err
		}
		instr.imm = v
	case "br", "br_if":
		v, err := p.nextInt()
		if err != nil {
			return instr, err
		}
		idx, err := toU64(v)
		if err != nil {
			return instr, err
		}
		instr.idx = idx
	case "call":
		return p.parseCall()
	case "i32.load", "i64.load", "i32.store", "i64.store":
		if err := p.parseMemArg(&instr); err != nil {
			return instr, err
		}
	}
	return instr, nil
}

// parseCall parses a call target as either a numeric index or a $name alias.
func (p *watParser) parseCall() (watInstr, error) {
	v, err := p.next()
	if err != nil {
		return watInstr{}, err
	}
	switch {
	case v.kind == tokInt:
		idx, err := toU64(v.value)
		if err != nil {
			return watInstr{}, err
		}
		return watInstr{op: "call", idx: idx}, nil
	case v.kind == tokIdent && strings.HasPrefix(v.text, "$"):
		idx, ok := p.names[v.text]
		if !ok {
			return watInstr{}, fmt.Errorf("wasm: unknown function %q", v.text)
		}
		return watInstr{op: "call", idx: uint64(idx)}, nil //nolint:gosec // index from a validated alias map
	default:
		return watInstr{}, fmt.Errorf("wasm: bad call target %q", v.text)
	}
}

// parseBlock parses a structured control-flow body (loop/block/if).
func (p *watParser) parseBlock(op string) (watInstr, error) {
	blk := &watBlock{kind: op}
	tt, err := p.peekTok()
	if err != nil {
		return watInstr{}, err
	}
	if tt.kind == tokLParen {
		_, _ = p.next()
		t, err := p.nextIdent()
		if err != nil {
			return watInstr{}, err
		}
		if t == "result" {
			vt, err := p.nextType()
			if err != nil {
				return watInstr{}, err
			}
			blk.result = vt
			if err := p.expect(tokRParen); err != nil {
				return watInstr{}, err
			}
		} else {
			if err := p.skipSExpr(); err != nil {
				return watInstr{}, err
			}
		}
	}
	for {
		tt, err := p.peekTok()
		if err != nil {
			return watInstr{}, err
		}
		if tt.kind == tokIdent && tt.text == "else" {
			_, _ = p.next()
			els, err := p.parseSeqUntilEnd()
			if err != nil {
				return watInstr{}, err
			}
			blk.els = els
			return watInstr{op: op, block: blk}, nil
		}
		if tt.kind == tokIdent && tt.text == "end" {
			_, _ = p.next()
			return watInstr{op: op, block: blk}, nil
		}
		instr, err := p.parseInstr()
		if err != nil {
			return watInstr{}, err
		}
		blk.then = append(blk.then, instr)
	}
}

// parseSeqUntilEnd parses instructions until the matching 'end' token.
func (p *watParser) parseSeqUntilEnd() ([]watInstr, error) {
	var seq []watInstr
	for {
		tt, err := p.peekTok()
		if err != nil {
			return nil, err
		}
		if tt.kind == tokIdent && tt.text == "end" {
			_, _ = p.next()
			return seq, nil
		}
		instr, err := p.parseInstr()
		if err != nil {
			return nil, err
		}
		seq = append(seq, instr)
	}
}

// parseMemArg reads optional offset=/align= keyword arguments on a memory op.
func (p *watParser) parseMemArg(instr *watInstr) error {
	for {
		tt, err := p.peekTok()
		if err != nil {
			return err
		}
		if tt.kind != tokIdent || !strings.Contains(tt.text, "=") {
			return nil
		}
		_, _ = p.next()
		kv := strings.SplitN(tt.text, "=", 2)
		v, ok := parseWatInt(kv[1])
		if !ok {
			return fmt.Errorf("wasm: bad memory argument %q", tt.text)
		}
		switch kv[0] {
		case "offset":
			off, err := toU32(v)
			if err != nil {
				return err
			}
			instr.memOff = off
		case "align":
			algn, err := toU32(v)
			if err != nil {
				return err
			}
			instr.memAlgn = algn
		}
	}
}
