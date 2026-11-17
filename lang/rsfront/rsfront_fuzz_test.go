/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package rsfront

import (
	"testing"

	"abate/lang/air"
)

// FuzzRsFrontend reduces arbitrary byte buffers with the native Go fuzzer.
// The adapter runs directly: a panic would fail the fuzz target through the Go
// testing framework, so no explicit recover wrapper is used. Every accepted
// parse must return a valid unit with a positive node count and a well-defined
// fingerprint.
func FuzzRsFrontend(f *testing.F) {
	seeds := []string{
		"",
		"fn main() { let x = 1; }\n",
		"use std::io;\n",
		"struct S { x: i32 }\n",
		"impl S { fn f(&self) {} }\n",
		"trait T { fn g(&self); }\n",
		"match v { Some(x) => x, None => 0 }\n",
		"macro_rules! m { () => {}; }\n",
		"println!(\"{}\", x);\n",
		"let v = vec![1, 2, 3];\n",
		"fn f<'a>(x: &'a str) -> &'a str { x }\n",
		"let f = |x: i32| x * 2;\n",
		"unsafe { let p = 1; }\n",
		"async fn load() { fut.await; }\n",
		"#[derive(Clone)]\nstruct S;\n",
		"let r = start..end;\n",
		"let s = r#\"raw\"#;\n",
		"let c = 'x';\n",
		"x = ((((\n",
		"\xff\xfe\n",
	}
	for _, seed := range seeds {
		f.Add([]byte(seed))
	}
	f.Fuzz(func(t *testing.T, data []byte) {
		fe, err := NewRsFrontend()
		if err != nil {
			t.Fatalf("NewRsFrontend: %v", err)
		}
		unit := air.SourceUnit{
			Language: "rs", Dialect: "2021",
			Files: []air.SourceFile{air.NewSourceFile("fuzz.rs", data)},
		}
		out, diags, parseErr := fe.Parse(unit)
		if parseErr != nil {
			return
		}
		if out == nil || len(out.Roots) != 1 || out.Roots[0] == nil {
			t.Fatalf("nil or malformed unit for %q",
				data[:rsMinInt(len(data), 32)])
		}
		if err := out.Validate(); err != nil {
			t.Fatalf("invalid unit for %q: %v",
				data[:rsMinInt(len(data), 32)], err)
		}
		if out.NodeCount() < 1 {
			t.Fatal("unit has no nodes")
		}
		if out.Roots[0].Kind() != air.KindPackage {
			t.Fatal("root is not KindPackage")
		}
		if err := air.CheckIntervals(out.Roots[0], len(data)); err != nil {
			t.Fatalf("intervals for %q: %v",
				data[:rsMinInt(len(data), 32)], err)
		}
		_ = out.Fingerprint()
		_ = diags
	})
}

// TestRsFuzzSeedsReproducible runs the seed corpus twice through the adapter
// and checks byte-identical fingerprints for each seed.
func TestRsFuzzSeedsReproducible(t *testing.T) {
	seeds := []string{
		"fn main() {}\n", "let x = 1;\n", "vec![1, 2]\n",
		"r#\"raw\"#\n", "\xff\n", "match v { _ => 0 }\n",
	}
	fe := rsFrontend(t)
	for _, src := range seeds {
		unit := air.SourceUnit{
			Language: "rs", Dialect: "2021",
			Files: []air.SourceFile{air.NewSourceFile("seed.rs", []byte(src))},
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
