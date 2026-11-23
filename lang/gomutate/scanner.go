/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/
package gomutate

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
)

// scanner enumerates the mutable tokens of every non-test Go source file
// under the scan root. It reuses the official Go frontend so discovery
// never disagrees with the compiler that later judges each mutant.
type scanner struct {
	// modRoot is the module root the staged workspace is copied from.
	modRoot string
	// scanRoot is the user-selected subtree discovery is limited to.
	scanRoot string
	fset     *token.FileSet
}

// newScanner builds a scanner walking scanRoot while recording every token
// path relative to modRoot, which keeps token paths directly comparable to
// the module-relative file names of coverprofiles.
func newScanner(modRoot, scanRoot string) *scanner {
	return &scanner{modRoot: modRoot, scanRoot: scanRoot, fset: token.NewFileSet()}
}

// scan parses every non-test .go file below the root and returns the
// discovered tokens in (path, offset) total order. Directory entries are
// visited in lexicographic order so the result never depends on the
// underlying file system enumeration order.
func (s *scanner) scan() ([]Token, error) {
	files, err := s.sourceFiles()
	if err != nil {
		return nil, err
	}
	if len(files) == 0 {
		return nil, ErrNoGoFiles
	}
	tokens := make([]Token, 0, 64)
	for _, path := range files {
		found, err := s.scanFile(path)
		if err != nil {
			return nil, err
		}
		tokens = append(tokens, found...)
	}
	return tokens, nil
}

// sourceFiles lists the non-test .go files below the root, excluding the
// vendor tree, the hidden .git tree and the tool's own staging area.
func (s *scanner) sourceFiles() ([]string, error) {
	if info, err := os.Stat(s.scanRoot); err != nil || !info.IsDir() {
		return nil, ErrInvalidRoot
	}
	files := make([]string, 0, 32)
	walk := func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			switch d.Name() {
			case "vendor", ".git":
				return filepath.SkipDir
			}
			return nil
		}
		name := d.Name()
		if strings.HasSuffix(name, ".go") && !strings.HasSuffix(name, "_test.go") {
			files = append(files, path)
		}
		return nil
	}
	if err := filepath.WalkDir(s.scanRoot, walk); err != nil {
		return nil, err
	}
	slices.Sort(files)
	return files, nil
}

// scanFile parses one file and collects every mutable token it carries.
func (s *scanner) scanFile(path string) ([]Token, error) {
	file, err := parser.ParseFile(s.fset, path, nil, 0)
	if err != nil {
		return nil, err
	}
	rel, err := filepath.Rel(s.modRoot, path)
	if err != nil {
		return nil, err
	}
	rel = filepath.ToSlash(rel)
	tf := s.fset.File(file.Pos())
	out := make([]Token, 0, 16)
	add := func(class, text string, start, end token.Pos) {
		out = append(out, Token{
			Path:   rel,
			Offset: tf.Offset(start),
			End:    tf.Offset(end),
			Line:   tf.Line(start),
			Column: tf.Position(start).Column,
			Class:  class,
			Text:   text,
		})
	}
	ast.Inspect(file, func(n ast.Node) bool {
		switch node := n.(type) {
		case *ast.BinaryExpr:
			add(ClassOperator, node.Op.String(), node.OpPos, node.OpPos+token.Pos(len(node.Op.String())))
		case *ast.BasicLit:
			if node.Kind == token.INT && integerLiteral(node.Value) {
				add(ClassNumber, node.Value, node.ValuePos, node.ValuePos+token.Pos(len(node.Value)))
			}
		case *ast.Ident:
			if node.Name == "true" || node.Name == "false" {
				add(ClassBool, node.Name, node.Pos(), node.End())
			}
		case *ast.BranchStmt:
			if node.Tok == token.BREAK || node.Tok == token.CONTINUE {
				add(ClassBranch, node.Tok.String(), node.Pos(), node.End())
			}
		case *ast.ExprStmt:
			if _, ok := node.X.(*ast.CallExpr); ok {
				add(ClassCall, "call", node.Pos(), node.End())
			}
		}
		return true
	})
	return out, nil
}

// integerLiteral reports whether the literal parses as a non-negative
// decimal integer, which is the only literal family the constant-perturb
// class can rewrite while keeping the mutant compilable. Negative values
// are rejected because the rewritten text replaces the literal without its
// sign, and hex or octal forms are rejected because a decimal replacement
// would silently change the literal's base representation in the report.
func integerLiteral(value string) bool {
	if value == "" {
		return false
	}
	for _, r := range value {
		if r < '0' || r > '9' {
			return false
		}
	}
	if _, err := strconv.ParseUint(value, 10, 63); err != nil {
		return false
	}
	return true
}
