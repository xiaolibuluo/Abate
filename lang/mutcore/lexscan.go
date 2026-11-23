/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/
package mutcore

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"abate/lang/frontkit"
)

// MutableOperators is the shared set of binary operators the operator
// registry rewrites. Languages whose lexer merges compound operators into
// one token (=>, ===, etc.) keep them out of this set.
var MutableOperators = map[string]bool{
	"+": true, "-": true, "*": true, "/": true,
	"<": true, "<=": true, ">": true, ">=": true,
	"==": true, "!=": true,
	"&&": true, "||": true,
}

// BranchKeywords is the shared set of loop-control keywords the branch
// operator rewrites.
var BranchKeywords = map[string]bool{
	"break": true, "continue": true,
}

// LexicalScan walks the workspace tree, lexes every regular file whose
// extension matches one of the given suffixes (dot-prefixed, e.g. ".py"),
// and collects the mutable tokens in deterministic (path, offset) order.
// The classify callback decides the mutability class of each lexical
// token and receives the workspace-relative slash path of the file the
// token belongs to; returning ClassUnknown skips the token. Test-file
// exclusion is the caller's responsibility via the skip predicate on
// workspace-relative paths.
func LexicalScan(root string, exts map[string]bool, skip func(rel string) bool,
	spec frontkit.LexSpec, classify func(path string, tok frontkit.Token) (TokenClass, bool)) ([]Token, error) {
	var tokens []Token
	walk := func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			if skippedDirs[d.Name()] {
				return filepath.SkipDir
			}
			return nil
		}
		if !exts[strings.ToLower(filepath.Ext(d.Name()))] {
			return nil
		}
		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		rel = filepath.ToSlash(rel)
		if skip != nil && skip(rel) {
			return nil
		}
		data, err := os.ReadFile(path) //nolint:gosec // enumerated from the workspace tree.
		if err != nil {
			return fmt.Errorf("%w: %v", ErrNoSourceFiles, err)
		}
		lexer, err := frontkit.NewLexer(spec, data)
		if err != nil {
			return fmt.Errorf("%w: %v", ErrNoSourceFiles, err)
		}
		lexed, err := lexer.Tokenize()
		if err != nil {
			return fmt.Errorf("%w: %v", ErrNoSourceFiles, err)
		}
		for _, lt := range lexed {
			class, ok := classify(rel, lt)
			if !ok || class == ClassUnknown {
				continue
			}
			tokens = append(tokens, Token{
				Path:   rel,
				Offset: lt.Offset,
				End:    lt.EndOffset,
				Line:   lt.Line,
				Column: lt.Column,
				Class:  class,
				Text:   lt.Text,
			})
		}
		return nil
	}
	if err := filepath.WalkDir(root, walk); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrNoSourceFiles, err)
	}
	if len(tokens) == 0 {
		return nil, ErrNoSourceFiles
	}
	sort.Slice(tokens, func(i, j int) bool {
		if tokens[i].Path != tokens[j].Path {
			return tokens[i].Path < tokens[j].Path
		}
		return tokens[i].Offset < tokens[j].Offset
	})
	return tokens, nil
}

// ClassifyLexical maps one frontkit lexical token onto its mutation class
// using the shared operator/number/bool/branch sets. The path argument is
// accepted to match the LexicalScan callback shape and ignored: the
// classification is purely textual, which keeps every frontkit-based
// backend consistent with the lexical operator contract.
func ClassifyLexical(_ string, lt frontkit.Token) (TokenClass, bool) {
	text := lt.Text
	switch {
	case MutableOperators[text]:
		return ClassOperator, true
	case BoolText(text):
		return ClassBool, true
	case BranchKeywords[text]:
		return ClassBranch, true
	case IsMutableNumber(text):
		return ClassNumber, true
	}
	return ClassUnknown, false
}
