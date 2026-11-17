/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package pyfront

import (
	"testing"

	"abate/lang/air"
)

// FuzzPyFrontend reduces arbitrary byte buffers with the native Go fuzzer.
// The adapter must never panic, must return a valid unit when it reports no
// classified error, and must keep the fingerprint well defined.
func FuzzPyFrontend(f *testing.F) {
	seeds := []string{
		"",
		"x = 1\n",
		"def f(a):\n    return a\n",
		"class C:\n    pass\n",
		"xs = [x for x in ys if x]\n",
		"if x:\n\tpass\n",
		"f\"{name}\"\n",
		"# type: ignore\n",
		"x = ((((\n",
		"\xff\xfe\n",
		"async def g():\n    await h()\n",
		"match v:\n    case 1:\n        pass\n",
	}
	for _, seed := range seeds {
		f.Add([]byte(seed))
	}
	f.Fuzz(func(t *testing.T, data []byte) {
		fe, err := NewPyFrontend()
		if err != nil {
			t.Fatalf("NewPyFrontend: %v", err)
		}
		unit := air.SourceUnit{
			Language: "py", Dialect: "3.12",
			Files: []air.SourceFile{air.NewSourceFile("fuzz.py", data)},
		}
		out, diags, parseErr := fe.Parse(unit)
		if parseErr != nil {
			return
		}
		if out == nil || len(out.Roots) != 1 || out.Roots[0] == nil {
			t.Fatalf("nil or malformed unit for %q", data[:minInt(len(data), 32)])
		}
		if err := out.Validate(); err != nil {
			t.Fatalf("invalid unit for %q: %v", data[:minInt(len(data), 32)], err)
		}
		if out.NodeCount() < 1 {
			t.Fatal("unit has no nodes")
		}
		if out.Roots[0].Kind() != air.KindPackage {
			t.Fatal("root is not KindPackage")
		}
		_ = out.Fingerprint()
		_ = diags
	})
}

// TestPyFuzzSeedsReproducible runs the seed corpus twice through the adapter
// and checks byte-identical fingerprints for each seed.
func TestPyFuzzSeedsReproducible(t *testing.T) {
	seeds := []string{
		"x = 1\n", "def f():\n    pass\n", "xs = [1, 2, 3]\n",
		"f\"{a}\"\n", "\xff\n", "if x:\n\tpass\n",
	}
	fe := pyFrontend(t)
	for _, src := range seeds {
		unit := air.SourceUnit{
			Language: "py", Dialect: "3.12",
			Files: []air.SourceFile{air.NewSourceFile("seed.py", []byte(src))},
		}
		a, _, errA := fe.Parse(unit)
		b, _, errB := fe.Parse(unit)
		if (errA == nil) != (errB == nil) || a == nil || b == nil {
			t.Fatalf("inconsistent seed parse for %q", src)
		}
		if a.Fingerprint() != b.Fingerprint() {
			t.Fatalf("seed %q is not deterministic", src)
		}
	}
}
