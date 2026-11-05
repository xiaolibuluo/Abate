// Abate —— 止息混沌 · 混沌工程开放规范（CES）与参考实现
// Abate — where chaos abates. Open Chaos Experiment Specification (CES) & reference implementation.
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
//
// 本文件以 GNU Affero General Public License v3（AGPLv3）发布；商业用途须事先获得
// 小黎的部落书面授权，详见 COMMERCIAL.md。
// This file is published under the GNU Affero General Public License v3 (AGPLv3).
// Commercial use requires prior written authorization from 小黎的部落.
// See COMMERCIAL.md for details.

// Package mut 实现 Go 源码变异引擎：基于标准库 go/ast 遍历生成变异体。
// 零依赖红线：仅使用标准库；确定性生成（CES-9.2.2）：同一文件内容 +
// 同一 seed 产出同一变异体集合，顺序由确定性 RNG 洗牌决定。
//
// Package mut implements the Go source mutation engine: mutant generation by
// AST traversal using only the standard library. Deterministic generation
// (CES-9.2.2): the same file content plus the same seed yields the same
// mutant set, ordered by a deterministic RNG shuffle.
package mut

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"sort"

	"github.com/qingdujian/abate/core"
)

// arithSwap 是 M1 算术变异映射表：+/-/*/÷ 互换（取确定性轮换目标）。
// arithSwap is the M1 arithmetic mutation map (deterministic rotation).
var arithSwap = map[token.Token]token.Token{
	token.ADD: token.SUB,
	token.SUB: token.MUL,
	token.MUL: token.QUO,
	token.QUO: token.ADD,
	token.REM: token.MUL,
}

// relSwap 是 M2 关系变异映射表：比较运算符互换与边界偏移。
// relSwap is the M2 relational mutation map (interchange and boundary shift).
var relSwap = map[token.Token]token.Token{
	token.LSS: token.LEQ,
	token.LEQ: token.LSS,
	token.GTR: token.GEQ,
	token.GEQ: token.GTR,
	token.EQL: token.NEQ,
	token.NEQ: token.EQL,
}

// Generate 解析 Go 源码并生成全部可判定变异体（M1+M2），
// 返回顺序经 seed 驱动的确定性洗牌（CES-9.2.2、CES-9.4.2）。
//
// Generate parses Go source and produces all decidable mutants (M1+M2),
// deterministically shuffled by the seed (CES-9.2.2, CES-9.4.2).
func Generate(filename, src string, seed int64) ([]core.Mutation, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filename, src, parser.SkipObjectResolution)
	if err != nil {
		return nil, fmt.Errorf("mut: parse %s: %w", filename, err)
	}
	var out []core.Mutation
	ast.Inspect(f, func(n ast.Node) bool {
		expr, ok := n.(*ast.BinaryExpr)
		if !ok {
			return true
		}
		pos := fset.Position(expr.OpPos)
		if to, hit := arithSwap[expr.Op]; hit {
			out = append(out, core.Mutation{
				File: filename, Line: pos.Line, Col: pos.Column,
				Class: core.OpArithmetic,
				From:  expr.Op.String(), To: to.String(),
			})
		}
		if to, hit := relSwap[expr.Op]; hit {
			out = append(out, core.Mutation{
				File: filename, Line: pos.Line, Col: pos.Column,
				Class: core.OpRelational,
				From:  expr.Op.String(), To: to.String(),
			})
		}
		return true
	})
	// 生成顺序先按位置排序（遍历序的稳定化），再按 seed 洗牌（CES-9.2.2）。
	sort.Slice(out, func(i, j int) bool { return out[i].ID() < out[j].ID() })
	rng := core.NewRNG(seed)
	rng.Shuffle(len(out), func(i, j int) { out[i], out[j] = out[j], out[i] })
	return out, nil
}

// Apply 将变异体应用到源码上，返回变异后的源码文本。
// 采用 AST 重解析+文本替换的最小实现：按变异点行/列定位运算符并替换。
//
// Apply applies a mutant to the source, returning the mutated source text.
// Minimal implementation: locate the operator by the mutation's line/column
// and replace it textually.
func Apply(src string, m core.Mutation) (string, error) {
	if err := m.Validate(); err != nil {
		return "", err
	}
	lines := splitLines(src)
	if m.Line > len(lines) {
		return "", fmt.Errorf("mut: line %d beyond file (%d lines)", m.Line, len(lines))
	}
	line := lines[m.Line-1]
	col := m.Col - 1
	if col < 0 || col+len(m.From) > len(line) {
		return "", fmt.Errorf("mut: col %d out of range in line %d", m.Col, m.Line)
	}
	if line[col:col+len(m.From)] != m.From {
		return "", fmt.Errorf("mut: source drift at %d:%d, want %q", m.Line, m.Col, m.From)
	}
	lines[m.Line-1] = line[:col] + m.To + line[col+len(m.From):]
	return joinLines(lines), nil
}

// splitLines 按行切分，保留行尾语义（不附加额外换行）。
// splitLines splits source into lines without appending extra newlines.
func splitLines(s string) []string {
	var out []string
	start := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			out = append(out, s[start:i])
			start = i + 1
		}
	}
	if start < len(s) {
		out = append(out, s[start:])
	}
	return out
}

// joinLines 以换行符合并行切片。
// joinLines joins lines with newline characters.
func joinLines(lines []string) string {
	out := ""
	for i, l := range lines {
		if i > 0 {
			out += "\n"
		}
		out += l
	}
	return out
}
