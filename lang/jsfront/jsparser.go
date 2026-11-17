/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package jsfront

import (
	"path"
	"sort"
	"strconv"
	"strings"

	"abate/lang/air"
	"abate/lang/frontkit"
)

// jsDialect 是一次解析的方言协商结果，全部字段为确定性标量.
type jsDialect struct {
	flavor string
	jsx    bool
	ts     bool
	module string
	strict bool
	target string
}

// canonical 渲染方言的稳定规范串，写入单元注解与根节点标注.
func (d jsDialect) canonical() string {
	return "js/" + d.flavor + "/" + d.module + "/strict=" +
		strconv.FormatBool(d.strict) + "/target=" + d.target
}

// jsExtensionDialects 是六种扩展名的方言矩阵，声明顺序固定保证确定性.
var jsExtensionDialects = []struct {
	ext    string
	flavor string
	jsx    bool
	ts     bool
	module string
}{
	{".js", "js", false, false, "auto"},
	{".jsx", "jsx", true, false, "auto"},
	{".mjs", "mjs", false, false, "esm"},
	{".cjs", "cjs", false, false, "commonjs"},
	{".ts", "ts", false, true, "esm"},
	{".tsx", "tsx", true, true, "esm"},
}

// resolveDialect 按虚拟路径扩展名协商 JSX、TS 与模块制式，未知扩展名回退到 js.
func resolveDialect(p string) jsDialect {
	d := jsDialect{flavor: "js", module: "auto", target: "es2022"}
	ext := strings.ToLower(path.Ext(p))
	for i := range jsExtensionDialects {
		if jsExtensionDialects[i].ext == ext {
			row := jsExtensionDialects[i]
			d.flavor, d.jsx, d.ts, d.module = row.flavor, row.jsx, row.ts, row.module
			break
		}
	}
	d.strict = d.module == "esm"
	return d
}

// jsParser 是 token 流上的容错递归下降解析器，产出有层级的 AIR 树.
type jsParser struct {
	toks   []frontkit.Token
	i      int
	path   string
	d      jsDialect
	mapper *jsMapper
	sink   *frontkit.DiagnosticSink
	nodes  []air.Node
}

// newJsParser 绑定 token 流、方言、节点工厂与诊断槽.
func newJsParser(toks []frontkit.Token, srcPath string) *jsParser {
	d := resolveDialect(srcPath)
	return &jsParser{
		toks: toks, path: srcPath, d: d,
		mapper: newJsMapper(toks, srcPath, d),
		sink:   frontkit.NewDiagnosticSink(srcPath, "js"),
		nodes:  make([]air.Node, 0, 8),
	}
}

// parse 归约整个文件：顶层语句、注释、兜底投影，最后封装 KindPackage 根.
func (p *jsParser) parse() (air.Node, []air.Diagnostic, error) {
	if _, err := frontkit.NewTokenStream(p.toks); err != nil {
		_ = p.sink.Warning(1, 0, 1, 1, "token stream rejected: "+err.Error())
	}
	for p.i < len(p.toks) {
		tok := p.toks[p.i]
		if tok.Kind == frontkit.TokenEOF {
			break
		}
		if tok.Kind == frontkit.TokenNewline {
			p.i++
			continue
		}
		if tok.Kind == frontkit.TokenComment || tok.Kind == frontkit.TokenIllegal {
			p.nodes = append(p.nodes, p.mapper.leaf(p.i, tok))
			p.i++
			continue
		}
		at := p.i
		if node := p.parseStatement(); node != nil {
			p.nodes = append(p.nodes, node)
		}
		if p.i <= at {
			p.i = at + 1
		}
	}
	p.mapper.cover(0, len(p.toks), &p.nodes)
	p.nodes = append(p.nodes, p.dedicatedMarkers()...)
	sort.SliceStable(p.nodes, func(a, b int) bool {
		return p.nodes[a].Pos().Offset < p.nodes[b].Pos().Offset
	})
	root := p.mapper.node(0, len(p.toks)-1, air.KindPackage, "", p.nodes...)
	root = p.mapper.annotate(root, air.SemNote, noteModule)
	root = p.mapper.annotate(root, air.SemTag, p.d.canonical())
	root = p.mapper.annotate(root, air.SemDirective, p.d.module)
	return root, p.sink.Diagnostics(), nil
}

// significant 跳过注释、换行与 EOF，返回最近的有效 token 指针.
func (p *jsParser) significant(j, step int) *frontkit.Token {
	for j >= 0 && j < len(p.toks) {
		t := p.toks[j]
		if t.Kind == frontkit.TokenNewline || t.Kind == frontkit.TokenComment ||
			t.Kind == frontkit.TokenEOF {
			j += step
			continue
		}
		return &t
	}
	return nil
}

// statementEnd 用括号深度与 ASI 规则求当前语句的排他结束索引.
func (p *jsParser) statementEnd() int {
	depth := 0
	for j := p.i; j < len(p.toks); j++ {
		t := p.toks[j]
		if t.Kind == frontkit.TokenEOF {
			return j
		}
		switch t.Text {
		case "(", "[", "{":
			depth++
			continue
		case ")":
			if depth > 0 {
				depth--
				continue
			}
			return j + 1
		case "]":
			if depth > 0 {
				depth--
				continue
			}
			return j + 1
		case "}":
			if depth > 0 {
				depth--
				if depth == 0 {
					return j + 1
				}
				continue
			}
			return j
		case ";":
			if depth == 0 {
				return j + 1
			}
			continue
		}
		if t.Kind == frontkit.TokenNewline && depth == 0 {
			prev := p.significant(j-1, -1)
			next := p.significant(j+1, 1)
			if jsCanInsertSemicolonAfterNewline(prev, next) {
				return j
			}
		}
	}
	return len(p.toks)
}

// isTrivia 判定语句首 token 是否为不需要建节点的分隔符.
func (*jsParser) isTrivia(t frontkit.Token) bool {
	if t.Kind == frontkit.TokenNewline || t.Kind == frontkit.TokenEOF ||
		t.Kind == frontkit.TokenComment || t.Kind == frontkit.TokenIllegal {
		return true
	}
	return t.Text == ";" || t.Text == "}" || t.Text == ")" || t.Text == "]"
}

// parseStatement 归约一条顶层语句：声明类由映射表决定，表达式由 Pratt 解析器归约.
func (p *jsParser) parseStatement() air.Node {
	start := p.i
	end := p.statementEnd()
	if end <= start {
		end = start + 1
	}
	if end > len(p.toks) {
		end = len(p.toks)
	}
	first := p.toks[start]
	if p.isTrivia(first) {
		p.i = end
		return nil
	}
	kind, note, found := jsDeclMapping(first.Text)
	if !found {
		kind, note = air.KindExprStmt, "expression-statement"
	}
	from := start
	if found && first.Kind == frontkit.TokenKeyword {
		from = start + 1
	}
	if from > end {
		from = end
	}
	children := p.exprParser().parseRange(from, end)
	children = append(children, p.typeParser().decorate(start, end, note)...)
	p.mapper.cover(start, end, &children)
	p.i = end
	if first.Text == "for" && p.hasRangeFor(start, end) {
		kind, note = air.KindRangeFor, "for-of"
	}
	node := p.mapper.node(start, end-1, kind, "", children...)
	node = p.mapper.annotate(node, air.SemNote, note)
	return p.applySemantics(node, first.Text, start, end)
}

// hasRangeFor 判定 for 语句是否为 for-of / for-in 结构.
func (p *jsParser) hasRangeFor(start, end int) bool {
	for i := start + 1; i < end; i++ {
		if p.toks[i].Text == "of" || p.toks[i].Text == "in" {
			return true
		}
	}
	return false
}

// applySemantics 写入真实 SemKey：可变性、异步、指令与 TS 类型保留标记.
func (p *jsParser) applySemantics(node air.Node, head string,
	start, end int) air.Node {
	switch head {
	case "const":
		node = p.mapper.annotate(node, air.SemMutability, "immutable")
	case "let", "var":
		node = p.mapper.annotate(node, air.SemMutability, "mutable")
	case "async", "await":
		node = p.mapper.annotate(node, air.SemAsync, "true")
	case "import", "export", "declare":
		node = p.mapper.annotate(node, air.SemDirective, head)
	case "function":
		for i := start; i < end; i++ {
			if p.toks[i].Text == "async" {
				node = p.mapper.annotate(node, air.SemAsync, "true")
				break
			}
		}
	}
	if p.d.ts && (head == "interface" || head == "enum" || head == "type") {
		node = p.mapper.annotate(node, air.SemTypeRef, "preserved")
	}
	return node
}
