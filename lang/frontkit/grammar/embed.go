/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package grammar

import (
	"embed"
	"errors"
	"fmt"
	"io/fs"
	"slices"
)

// ErrUnknownLanguage reports a language tag without embedded grammar data.
var ErrUnknownLanguage = errors.New("grammar: unknown language")

//go:embed data
var embedded embed.FS

// languageData binds one upstream grammar data set to its provenance.
type languageData struct {
	// tag is the Abate short language tag.
	tag string
	// dir is the directory name of the embedded upstream data set.
	dir string
	// version is the upstream grammar package version.
	version string
	// source is the upstream repository the data was taken from.
	source string
}

// languages lists the embedded upstream grammar data sets in ascending tag
// order, so the slice never depends on map iteration.
var languages = []languageData{
	{tag: "c", dir: "c", version: "0.24.1", source: "https://github.com/tree-sitter/tree-sitter-c"},
	{tag: "cs", dir: "csharp", version: "0.23.5", source: "https://github.com/tree-sitter/tree-sitter-c-sharp"},
	{tag: "js", dir: "javascript", version: "0.25.0", source: "https://github.com/tree-sitter/tree-sitter-javascript"},
	{tag: "kt", dir: "kotlin", version: "0.3.8", source: "https://github.com/tree-sitter/tree-sitter-kotlin"},
	{tag: "py", dir: "python", version: "0.25.0", source: "https://github.com/tree-sitter/tree-sitter-python"},
	{tag: "rs", dir: "rust", version: "0.24.0", source: "https://github.com/tree-sitter/tree-sitter-rust"},
}

// SupportedLanguages returns the tags of the embedded upstream grammar data
// sets in ascending order.
func SupportedLanguages() []string {
	out := make([]string, 0, len(languages))
	for _, item := range languages {
		out = append(out, item.tag)
	}
	slices.Sort(out)
	return out
}

// Load returns the upstream grammar fact set of the given language tag.
//
// The data is read from the embedded copy, so the call works offline and
// never touches the filesystem of the caller.
func Load(tag string) (*Grammar, error) {
	item, ok := lookup(tag)
	if !ok {
		return nil, fmt.Errorf("%w: %q", ErrUnknownLanguage, tag)
	}
	return loadFrom(embedded, item)
}

// LoadFromFS returns the upstream grammar fact set stored in an arbitrary
// file system that uses the same layout as the embedded data directory.
//
// It exists so a future upstream revision can be verified from a checkout
// without changing the embedded bytes.
func LoadFromFS(tag string, fsys fs.FS) (*Grammar, error) {
	item, ok := lookup(tag)
	if !ok {
		return nil, fmt.Errorf("%w: %q", ErrUnknownLanguage, tag)
	}
	return loadFrom(fsys, item)
}

// lookup finds the provenance record of one language tag.
func lookup(tag string) (languageData, bool) {
	for _, item := range languages {
		if item.tag == tag {
			return item, true
		}
	}
	return languageData{}, false
}

// loadFrom reads and decodes one grammar data set from a file system.
func loadFrom(fsys fs.FS, item languageData) (*Grammar, error) {
	if fsys == nil {
		return nil, fmt.Errorf("%w: %q", ErrUnknownLanguage, item.tag)
	}
	base := "data/" + item.dir
	nodeTypes, err := fs.ReadFile(fsys, base+"/node-types.json")
	if err != nil {
		return nil, fmt.Errorf("grammar: %q node types: %w", item.tag, err)
	}
	productions, err := fs.ReadFile(fsys, base+"/grammar.json")
	if err != nil {
		return nil, fmt.Errorf("grammar: %q rules: %w", item.tag, err)
	}
	return Parse(item.tag, item.version, item.source, nodeTypes, productions)
}
