/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package wasm

import "testing"

// TestParseInstrAcceptsFlatAndParenthesizedForms verifies instruction parsing.
func TestParseInstrAcceptsFlatAndParenthesizedForms(t *testing.T) {
	cases := []struct {
		name    string
		src     string
		wantOp  string
		wantImm int64
		wantErr bool
	}{
		{name: "flat ident", src: "nop", wantOp: "nop"},
		{name: "flat with immediate", src: "i32.const 7", wantOp: "i32.const", wantImm: 7},
		{name: "flat block", src: "block (result i32) nop end", wantOp: "block"},
		{name: "parenthesized", src: "(nop)", wantOp: "nop"},
		{name: "parenthesized const", src: "(i32.const -3)", wantOp: "i32.const", wantImm: -3},
		{name: "integer token", src: "7", wantErr: true},
		{name: "right paren", src: ")", wantErr: true},
		{name: "eof", src: "", wantErr: true},
		{name: "missing immediate", src: "(i32.const)", wantErr: true},
		{name: "missing closing paren", src: "(nop", wantErr: true},
		{name: "bad inner ident", src: "(7)", wantErr: true},
		{name: "unexpected token in parens", src: "(nop 7)", wantErr: true},
	}
	for _, tc := range cases {
		p := newWatParser([]byte(tc.src))
		instr, err := p.parseInstr()
		if (err != nil) != tc.wantErr {
			t.Errorf("%s: error = %v, wantErr = %v", tc.name, err, tc.wantErr)
			continue
		}
		if tc.wantErr {
			continue
		}
		if instr.op != tc.wantOp || instr.imm != tc.wantImm {
			t.Errorf("%s: instr = %+v, want op %q imm %d", tc.name, instr, tc.wantOp, tc.wantImm)
		}
	}
}

// TestParseInstrBodyReadsImmediates verifies each immediate family and errors.
func TestParseInstrBodyReadsImmediates(t *testing.T) {
	cases := []struct {
		name      string
		op        string
		src       string
		wantImm   int64
		wantIdx   uint64
		wantOff   uint32
		wantAlign uint32
		wantErr   bool
	}{
		{name: "i32.const", op: "i32.const", src: "7", wantImm: 7},
		{name: "i64.const negative", op: "i64.const", src: "-3", wantImm: -3},
		{name: "br", op: "br", src: "2", wantIdx: 2},
		{name: "br_if", op: "br_if", src: "130", wantIdx: 130},
		{name: "call numeric", op: "call", src: "4", wantIdx: 4},
		{name: "memarg", op: "i32.load", src: "offset=8 align=2", wantOff: 8, wantAlign: 2},
		{name: "unknown numeric op", op: "i32.add", src: ""},
		{name: "end control token", op: "end", src: "", wantErr: true},
		{name: "else control token", op: "else", src: "", wantErr: true},
		{name: "missing const immediate", op: "i32.const", src: "nop", wantErr: true},
		{name: "negative br", op: "br", src: "-1", wantErr: true},
		{name: "missing call target", op: "call", src: ")", wantErr: true},
		{name: "bad memarg", op: "i32.store", src: "offset=xyz", wantErr: true},
	}
	for _, tc := range cases {
		p := newWatParser([]byte(tc.src))
		instr, err := p.parseInstrBody(tc.op)
		if (err != nil) != tc.wantErr {
			t.Errorf("%s: error = %v, wantErr = %v", tc.name, err, tc.wantErr)
			continue
		}
		if tc.wantErr {
			continue
		}
		if instr.imm != tc.wantImm || instr.idx != tc.wantIdx ||
			instr.memOff != tc.wantOff || instr.memAlgn != tc.wantAlign {
			t.Errorf("%s: instr = %+v, want imm %d idx %d off %d align %d",
				tc.name, instr, tc.wantImm, tc.wantIdx, tc.wantOff, tc.wantAlign)
		}
	}
}

// TestParseCallResolvesNumericAndNamedTargets verifies call target handling.
func TestParseCallResolvesNumericAndNamedTargets(t *testing.T) {
	numeric := newWatParser([]byte("7"))
	instr, err := numeric.parseCall()
	if err != nil {
		t.Fatalf("parseCall numeric: %v", err)
	}
	if instr.op != "call" || instr.idx != 7 {
		t.Errorf("parseCall numeric = %+v, want call 7", instr)
	}
	aliased := newWatParser([]byte("$f"))
	aliased.names["$f"] = 3
	instr, err = aliased.parseCall()
	if err != nil {
		t.Fatalf("parseCall alias: %v", err)
	}
	if instr.op != "call" || instr.idx != 3 {
		t.Errorf("parseCall alias = %+v, want call 3", instr)
	}
	bad := []struct {
		name string
		src  string
	}{
		{name: "unknown alias", src: "$missing"},
		{name: "plain identifier", src: "foo"},
		{name: "string target", src: `"f"`},
		{name: "negative index", src: "-1"},
		{name: "empty target", src: ""},
	}
	for _, tc := range bad {
		p := newWatParser([]byte(tc.src))
		if _, err := p.parseCall(); err == nil {
			t.Errorf("%s: expected an error", tc.name)
		}
	}
}

// TestParseBlockHandlesResultAndElse verifies structured control-flow parsing.
func TestParseBlockHandlesResultAndElse(t *testing.T) {
	result := newWatParser([]byte("(result i32) nop end"))
	instr, err := result.parseBlock("block")
	if err != nil {
		t.Fatalf("parseBlock result: %v", err)
	}
	if instr.op != "block" || instr.block == nil {
		t.Fatalf("parseBlock result = %+v, want a block instruction", instr)
	}
	if instr.block.result != vtI32 {
		t.Errorf("block result = %d, want i32", instr.block.result)
	}
	if len(instr.block.then) != 1 || instr.block.then[0].op != "nop" {
		t.Errorf("block then = %+v, want a single nop", instr.block.then)
	}
	if len(instr.block.els) != 0 {
		t.Errorf("block else = %+v, want empty", instr.block.els)
	}

	branch := newWatParser([]byte("nop else i32.const 5 end"))
	instr, err = branch.parseBlock("if")
	if err != nil {
		t.Fatalf("parseBlock else: %v", err)
	}
	if instr.block == nil || len(instr.block.els) != 1 ||
		instr.block.els[0].op != "i32.const" || instr.block.els[0].imm != 5 {
		t.Fatalf("parseBlock else = %+v, want one i32.const 5", instr)
	}

	unknownForm := newWatParser([]byte("(param i32) nop end"))
	instr, err = unknownForm.parseBlock("loop")
	if err != nil {
		t.Fatalf("parseBlock unknown sub-form: %v", err)
	}
	if instr.block == nil || instr.block.result != 0 || len(instr.block.then) != 1 {
		t.Fatalf("parseBlock unknown sub-form = %+v, want void result and one nop", instr)
	}
}

// TestParseBlockRejectsMalformedBodies verifies all failure paths.
func TestParseBlockRejectsMalformedBodies(t *testing.T) {
	cases := []struct {
		name string
		src  string
	}{
		{name: "no end", src: "nop"},
		{name: "bad result type", src: "(result f32) end"},
		{name: "unbalanced leading form", src: "(param i32 end"},
		{name: "bad instruction", src: "7 end"},
		{name: "bad else sequence", src: "else 7 end"},
		{name: "lexer error", src: `"open`},
	}
	for _, tc := range cases {
		p := newWatParser([]byte(tc.src))
		if _, err := p.parseBlock("if"); err == nil {
			t.Errorf("%s: expected an error", tc.name)
		}
	}
}

// TestParseSeqUntilEndStopsAtEnd verifies sequence parsing and stopping.
func TestParseSeqUntilEndStopsAtEnd(t *testing.T) {
	p := newWatParser([]byte("nop i32.const 3 end rest"))
	seq, err := p.parseSeqUntilEnd()
	if err != nil {
		t.Fatalf("parseSeqUntilEnd: %v", err)
	}
	if len(seq) != 2 || seq[0].op != "nop" || seq[1].op != "i32.const" || seq[1].imm != 3 {
		t.Fatalf("seq = %+v, want [nop i32.const 3]", seq)
	}
	tok, err := p.peekTok()
	if err != nil {
		t.Fatalf("peekTok after sequence: %v", err)
	}
	if tok.kind != tokIdent || tok.text != "rest" {
		t.Errorf("token after sequence = %+v, want ident rest", tok)
	}
	empty := newWatParser([]byte("end"))
	seq, err = empty.parseSeqUntilEnd()
	if err != nil {
		t.Fatalf("parseSeqUntilEnd empty: %v", err)
	}
	if len(seq) != 0 {
		t.Errorf("empty sequence = %+v, want no instructions", seq)
	}
	bad := newWatParser([]byte("7"))
	if _, err := bad.parseSeqUntilEnd(); err == nil {
		t.Fatal("parseSeqUntilEnd accepted a bad instruction")
	}
	truncated := newWatParser([]byte("nop"))
	if _, err := truncated.parseSeqUntilEnd(); err == nil {
		t.Fatal("parseSeqUntilEnd accepted a missing end")
	}
}

// TestParseMemArgReadsOffsetAndAlign verifies memory immediate parsing.
func TestParseMemArgReadsOffsetAndAlign(t *testing.T) {
	cases := []struct {
		name      string
		src       string
		wantOff   uint32
		wantAlign uint32
		wantErr   bool
	}{
		{name: "offset and align", src: "offset=8 align=2 nop", wantOff: 8, wantAlign: 2},
		{name: "hex offset", src: "offset=0x10", wantOff: 16},
		{name: "unknown key", src: "foo=3"},
		{name: "no arguments", src: "nop"},
		{name: "bad value", src: "offset=xyz", wantErr: true},
		{name: "negative offset", src: "offset=-1", wantErr: true},
		{name: "overflow align", src: "align=4294967296", wantErr: true},
		{name: "empty value", src: "align=", wantErr: true},
		{name: "lexer error", src: `"open`, wantErr: true},
	}
	for _, tc := range cases {
		p := newWatParser([]byte(tc.src))
		instr := &watInstr{op: "i32.load"}
		err := p.parseMemArg(instr)
		if (err != nil) != tc.wantErr {
			t.Errorf("%s: error = %v, wantErr = %v", tc.name, err, tc.wantErr)
			continue
		}
		if tc.wantErr {
			continue
		}
		if instr.memOff != tc.wantOff || instr.memAlgn != tc.wantAlign {
			t.Errorf("%s: memarg = off %d align %d, want %d %d",
				tc.name, instr.memOff, instr.memAlgn, tc.wantOff, tc.wantAlign)
		}
	}
	p := newWatParser([]byte("offset=8 nop"))
	instr := &watInstr{op: "i32.load"}
	if err := p.parseMemArg(instr); err != nil {
		t.Fatalf("parseMemArg: %v", err)
	}
	tok, err := p.peekTok()
	if err != nil {
		t.Fatalf("peekTok after memarg: %v", err)
	}
	if tok.kind != tokIdent || tok.text != "nop" {
		t.Errorf("token after memarg = %+v, want ident nop", tok)
	}
}
