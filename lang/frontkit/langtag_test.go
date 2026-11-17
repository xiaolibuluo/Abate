/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package frontkit

import (
	"reflect"
	"strconv"
	"strings"
	"testing"

	"abate/internal/determinism"
)

// canonicalTagCases is the frozen alias matrix exercised by the table tests.
var canonicalTagCases = []struct {
	in   string
	want string
}{
	{"go", "go"}, {"golang", "go"},
	{"py", "py"}, {"py3", "py"}, {"python", "py"},
	{"js", "js"}, {"javascript", "js"}, {"typescript", "js"},
	{"ts", "js"}, {"node", "js"},
	{"c", "c"}, {"cpp", "c"}, {"c++", "c"}, {"cxx", "c"}, {"cc", "c"},
	{"rs", "rs"}, {"rust", "rs"},
	{"cs", "cs"}, {"csharp", "cs"}, {"c#", "cs"}, {"dotnet", "cs"},
	{"kt", "kt"}, {"kotlin", "kt"},
}

// extensionCases is the frozen seven-language extension matrix.
var extensionCases = []struct {
	tag  string
	want []string
}{
	{"go", []string{".go"}},
	{"py", []string{".py", ".pyi"}},
	{"js", []string{".js", ".jsx", ".mjs", ".cjs", ".ts", ".tsx"}},
	{"c", []string{".c", ".h", ".cc", ".cpp", ".cxx", ".hpp", ".hh"}},
	{"rs", []string{".rs"}},
	{"cs", []string{".cs"}},
	{"kt", []string{".kt", ".kts"}},
}

func TestCanonicalTagAliasMatrix(t *testing.T) {
	t.Parallel()
	for _, tc := range canonicalTagCases {
		got, ok := CanonicalTag(tc.in)
		if !ok || got != tc.want {
			t.Fatalf("CanonicalTag(%q) = (%q, %v), want (%q, true)",
				tc.in, got, ok, tc.want)
		}
	}
}

func TestCanonicalTagCaseAndWhitespace(t *testing.T) {
	t.Parallel()
	cases := []struct {
		in   string
		want string
	}{
		{"  PyThOn  ", "py"},
		{"\tGOLANG\n", "go"},
		{"C++", "c"},
		{"C#", "cs"},
		{"  JavaScript ", "js"},
		{" Rust ", "rs"},
		{"KOTLIN", "kt"},
		{"\t js \t", "js"},
	}
	for _, tc := range cases {
		got, ok := CanonicalTag(tc.in)
		if !ok || got != tc.want {
			t.Fatalf("CanonicalTag(%q) = (%q, %v), want (%q, true)",
				tc.in, got, ok, tc.want)
		}
	}
}

func TestCanonicalTagRejectsUnknown(t *testing.T) {
	t.Parallel()
	cases := []string{
		"", "   ", "\t\n", "ruby", "java", "go++", "py 3", "c#extra",
		"not-a-language", "\x00", "typescript2", "gol", "kot",
	}
	for _, in := range cases {
		got, ok := CanonicalTag(in)
		if ok || got != "" {
			t.Fatalf("CanonicalTag(%q) = (%q, %v), want (\"\", false)", in, got, ok)
		}
	}
}

func TestExtensionsForAllLanguages(t *testing.T) {
	t.Parallel()
	for _, tc := range extensionCases {
		got := ExtensionsFor(tc.tag)
		if !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("ExtensionsFor(%q) = %v, want %v", tc.tag, got, tc.want)
		}
	}
}

func TestExtensionsForAliasAndUnknown(t *testing.T) {
	t.Parallel()
	if got := ExtensionsFor("python"); !reflect.DeepEqual(got, extensionCases[1].want) {
		t.Fatalf("ExtensionsFor(\"python\") = %v, want the py table", got)
	}
	if got := ExtensionsFor("c++"); !reflect.DeepEqual(got, extensionCases[3].want) {
		t.Fatalf("ExtensionsFor(\"c++\") = %v, want the c table", got)
	}
	for _, unknown := range []string{"", "ruby", "  ", "\x00"} {
		if got := ExtensionsFor(unknown); len(got) != 0 {
			t.Fatalf("ExtensionsFor(%q) = %v, want an empty slice", unknown, got)
		}
	}
}

func TestExtensionsForDefensiveCopy(t *testing.T) {
	t.Parallel()
	first := ExtensionsFor("js")
	if len(first) == 0 {
		t.Fatal("ExtensionsFor(\"js\") returned no extensions")
	}
	first[0] = ".mutated"
	again := ExtensionsFor("js")
	if again[0] == ".mutated" {
		t.Fatal("ExtensionsFor returned a shared slice; the frozen table was mutated")
	}
}

func TestTagForPathAllExtensions(t *testing.T) {
	t.Parallel()
	for _, tc := range extensionCases {
		for _, ext := range tc.want {
			got, ok := TagForPath("src/pkg/main" + ext)
			if !ok || got != tc.tag {
				t.Fatalf("TagForPath(%q) = (%q, %v), want (%q, true)",
					"src/pkg/main"+ext, got, ok, tc.tag)
			}
			upper := "SRC/PKG/MAIN" + strings.ToUpper(ext)
			gotUpper, okUpper := TagForPath(upper)
			if !okUpper || gotUpper != tc.tag {
				t.Fatalf("TagForPath(%q) = (%q, %v), want (%q, true)",
					upper, gotUpper, okUpper, tc.tag)
			}
		}
	}
}

func TestTagForPathRejectsUnknown(t *testing.T) {
	t.Parallel()
	cases := []string{"", "Makefile", "noext", "a.rb", "a.txt", ".", "a.", "dir/", "a.go.bak"}
	for _, in := range cases {
		got, ok := TagForPath(in)
		if ok || got != "" {
			t.Fatalf("TagForPath(%q) = (%q, %v), want (\"\", false)", in, got, ok)
		}
	}
}

// langTagSnapshot folds every alias, extension and path mapping into one
// canonical byte string for the 100-run determinism assertion.
func langTagSnapshot() []byte {
	parts := make([]string, 0, 128)
	for _, tc := range canonicalTagCases {
		tag, ok := CanonicalTag(tc.in)
		parts = append(parts, tc.in+"="+tag+"="+strconv.FormatBool(ok))
	}
	for _, tc := range extensionCases {
		for _, ext := range ExtensionsFor(tc.tag) {
			got, ok := TagForPath("src/main" + ext)
			parts = append(parts, tc.tag+ext+"="+got+"="+strconv.FormatBool(ok))
		}
	}
	return []byte(strings.Join(parts, "|"))
}

func TestLangTagDeterminism100(t *testing.T) {
	t.Parallel()
	determinism.MustMatch100(t, 42, langTagSnapshot)
}
