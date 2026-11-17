/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package pyfront

import (
	"flag"
	"path/filepath"
	"testing"

	"abate/lang/air"
	"abate/lang/frontkit/coverage"
	"abate/lang/frontkit/grammar"
)

// coverageUpdate 报告是否重写覆盖清单，用 go test -update 打开。
var coverageUpdate = flag.Bool("update", false,
	"rewrite the python frontend coverage checklist")

// pyCoverageCorpus 是真实 Python 语料：既有逐构造映射夹具，也有整文件样例。
func pyCoverageCorpus() []string {
	out := make([]string, 0, 96)
	for _, tc := range pyMappingCases {
		out = append(out, tc.src)
	}
	return append(out,
		"import os\nfrom sys import argv\nimport a.b.c as m\nfrom . import rel\n",
		"class C(Base, metaclass=Meta):\n    attr = 1\n\n    def m(self, x=1):\n        return x\n",
		"async def g():\n    async with open('f') as fh:\n        async for line in fh:\n            await h(line)\n",
		"def f(x: int = 1, *args: str, **kw: Any) -> Optional[int]:\n    return None\n",
		"try:\n    f()\nexcept (A, B) as e:\n    raise\nelse:\n    pass\nfinally:\n    g()\n",
		"with open('f') as fh, open('g') as gh:\n    pass\n",
		"match v:\n    case [1, 2, *rest]:\n        pass\n    case {'k': v, **rest}:\n        pass\n    case C(x=1) | D(y=2):\n        pass\n    case _:\n        pass\n",
		"x = [i for i in range(10) if i % 2 == 0]\n",
		"y = {k: v for k, v in items}\n",
		"z = {i for i in items}\n",
		"g = (i * 2 for i in items)\n",
		"s = f'{a!r:>10} literal {{escaped}}'\n",
		"t = b'bytes' rb'raw'\n",
		"n = 0x1f + 0o17 + 0b1010 + 1_000 + 1.5e3 + 3j\n",
		"b = True and False or None\n",
		"d = a if b else c\n",
		"w = (yield)\n",
		"lam = lambda a, b=1, *c, **d: a + b\n",
		"q = obj.attr[0].method(*args, **kwargs)\n",
		"p = a[1:5:2]\n",
		"del xs[0]\n",
		"global g\nnonlocal_holder = 1\n",
		"assert x, 'boom'\n",
		"print('x')\n",
		"exec('code')\n",
		"@staticmethod\ndef s():\n    pass\n",
		"type Alias = list[int]\n",
		"def g[T](x: T) -> T:\n    return x\n",
		"v: int | str = 1\n",
		"w: list[str] = []\n",
		"x: Optional[int] = None\n",
		"def f() -> bool:\n    return True\n",
		"type Alias = int\n",
		"def g[T](x: T) -> T:\n    return x\n",
		"v: int | str = 1\n",
		"w: int | None = None\n",
	)
}

// TestPyConstructCoverage 证明 Python 前端对上游语法的覆盖是可测量的。
//
// 覆盖数字来自对该语言真实语料的真实解析：上游 tree-sitter-python 的具体
// 构造全集与 AIR 实际产出的节点类型求差集，得到的清单入库为
// docs/dev/frontend-coverage/python.md。
func TestPyConstructCoverage(t *testing.T) {
	g, err := grammar.Load("py")
	if err != nil {
		t.Fatalf("grammar.Load: %v", err)
	}
	mappings := grammar.Map(g, grammar.DefaultRules("py"))
	observed := map[air.NodeKind]int{}
	for _, src := range pyCoverageCorpus() {
		unit, _, err := pyFrontend(t).Parse(air.SourceUnit{
			Language: "py", Dialect: "3.12",
			Files: []air.SourceFile{air.NewSourceFile("main.py", []byte(src))},
		})
		if err != nil {
			t.Fatalf("Parse(%q): %v", src, err)
		}
		observed = coverage.Merge(observed, coverage.Observe(unit))
	}
	cov := coverage.Compute(g, mappings, observed)
	t.Logf("python constructs=%d covered=%d (%.1f%%) exact=%d degraded=%d out-of-scope=%d unmapped=%d",
		cov.Total, cov.Covered, cov.Percent(), cov.Exact, cov.Degraded,
		cov.OutOfScope, cov.Unmapped)
	if cov.Unmapped != 0 {
		t.Fatalf("%d upstream constructs are not mapped", cov.Unmapped)
	}
	if cov.Covered == 0 {
		t.Fatal("the real parse produced no mapped AIR kind")
	}
	report := coverage.Markdown(cov)
	path := filepath.Join("..", "..", "docs", "dev", "frontend-coverage", "python.md")
	if err := coverage.Reconcile(path, report, *coverageUpdate); err != nil {
		t.Fatal(err)
	}
}
