/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package frontkit

import (
	"path"
	"strings"
)

// canonicalTags lists the only normative short language tags. Long spellings are
// never used as tags because rule IDs filter on the short first segment.
var canonicalTags = [...]string{"go", "py", "js", "c", "rs", "cs", "kt"}

// languageAlias maps one accepted alias spelling to its canonical short tag.
type languageAlias struct {
	alias     string
	canonical string
}

// languageAliases is the frozen alias matrix. It is an ordered slice, never a
// map that is iterated, so no map traversal order can leak into any result.
var languageAliases = []languageAlias{
	{"golang", "go"},
	{"py3", "py"},
	{"python", "py"},
	{"javascript", "js"},
	{"typescript", "js"},
	{"ts", "js"},
	{"node", "js"},
	{"cpp", "c"},
	{"c++", "c"},
	{"cxx", "c"},
	{"cc", "c"},
	{"rust", "rs"},
	{"csharp", "cs"},
	{"c#", "cs"},
	{"dotnet", "cs"},
	{"kotlin", "kt"},
}

// languageExtension binds one lower-cased file extension to a canonical tag.
type languageExtension struct {
	tag       string
	extension string
}

// languageExtensions is the frozen extension matrix in canonical tag order.
var languageExtensions = []languageExtension{
	{"go", ".go"},
	{"py", ".py"},
	{"py", ".pyi"},
	{"js", ".js"},
	{"js", ".jsx"},
	{"js", ".mjs"},
	{"js", ".cjs"},
	{"js", ".ts"},
	{"js", ".tsx"},
	{"c", ".c"},
	{"c", ".h"},
	{"c", ".cc"},
	{"c", ".cpp"},
	{"c", ".cxx"},
	{"c", ".hpp"},
	{"c", ".hh"},
	{"rs", ".rs"},
	{"cs", ".cs"},
	{"kt", ".kt"},
	{"kt", ".kts"},
}

// CanonicalTag normalizes one language spelling into its canonical short tag.
// Leading and trailing whitespace and ASCII case are ignored. Unknown, empty or
// malformed inputs return false without panicking.
func CanonicalTag(s string) (string, bool) {
	normalized := strings.ToLower(strings.TrimSpace(s))
	if normalized == "" {
		return "", false
	}
	for i := range canonicalTags {
		if canonicalTags[i] == normalized {
			return normalized, true
		}
	}
	for i := range languageAliases {
		if languageAliases[i].alias == normalized {
			return languageAliases[i].canonical, true
		}
	}
	return "", false
}

// ExtensionsFor returns the frozen extension list of a language, in declaration
// order. Aliases are accepted and normalized first. The returned slice is always
// fresh, so callers cannot mutate the package-level table; an unknown tag yields
// an empty slice.
func ExtensionsFor(tag string) []string {
	canonical, ok := CanonicalTag(tag)
	if !ok {
		return []string{}
	}
	out := make([]string, 0, len(languageExtensions))
	for i := range languageExtensions {
		if languageExtensions[i].tag == canonical {
			out = append(out, languageExtensions[i].extension)
		}
	}
	return out
}

// TagForPath maps a virtual source path to its canonical short tag using the
// lower-cased file extension. It never touches the filesystem. Paths without a
// known extension return false without panicking.
func TagForPath(p string) (string, bool) {
	if p == "" {
		return "", false
	}
	ext := strings.ToLower(path.Ext(p))
	if ext == "" {
		return "", false
	}
	for i := range languageExtensions {
		if languageExtensions[i].extension == ext {
			return languageExtensions[i].tag, true
		}
	}
	return "", false
}
