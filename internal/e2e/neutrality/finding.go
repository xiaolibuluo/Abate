/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package neutrality

import (
	"go/ast"
	"go/token"
	"path/filepath"
	"strings"
)

type (
	// Finding is one syntax-level zero-exemption violation. File and Line give
	// the exact source position, Rule names the invariant, and Snippet is a
	// stable explanation without any source text copy.
	Finding struct {
		File    string `json:"file"`
		Line    int    `json:"line"`
		Rule    string `json:"rule"`
		Snippet string `json:"snippet"`
	}

	// Result is the complete analysis of one registered source list.
	// ProcessedFiles must equal RegisteredFiles: a file that could not be
	// parsed is a zero-tolerance condition, never a reason to skip it.
	Result struct {
		RegisteredFiles int       `json:"registeredFiles"`
		ProcessedFiles  int       `json:"processedFiles"`
		Findings        []Finding `json:"findings"`
	}
)

// compareFindings orders findings deterministically.
func compareFindings(a, b Finding) int {
	if c := strings.Compare(a.File, b.File); c != 0 {
		return c
	}
	if a.Line != b.Line {
		return a.Line - b.Line
	}
	if c := strings.Compare(a.Rule, b.Rule); c != 0 {
		return c
	}
	return strings.Compare(a.Snippet, b.Snippet)
}

// newFinding renders one violation at a source position.
func newFinding(fset *token.FileSet, at token.Pos, rule,
	snippet string) Finding {
	pos := fset.Position(at)
	return Finding{
		File: pos.Filename, Line: pos.Line, Rule: rule, Snippet: snippet,
	}
}

// provenanceUsed reports whether an expression names a corpus provenance
// field, either as a selector or as a bare identifier.
func provenanceUsed(expr ast.Expr) bool {
	found := false
	ast.Inspect(expr, func(n ast.Node) bool {
		switch v := n.(type) {
		case *ast.SelectorExpr:
			if provenanceFields[v.Sel.Name] {
				found = true
			}
		case *ast.Ident:
			if provenanceFields[v.Name] {
				found = true
			}
		}
		return !found
	})
	return found
}

// corpusExpr reports whether an expression iterates a manifest or a sample.
func corpusExpr(node ast.Node) bool {
	if node == nil {
		return false
	}
	found := false
	ast.Inspect(node, func(n ast.Node) bool {
		switch v := n.(type) {
		case *ast.SelectorExpr:
			if v.Sel.Name == "Entries" {
				found = true
			}
		case *ast.CallExpr:
			if sampleCallName(v.Fun) {
				found = true
			}
		case *ast.Ident:
			if corpusVarName(v.Name) {
				found = true
			}
		}
		return !found
	})
	return found
}

// corpusVarName recognises local identifiers that hold a manifest or sample.
func corpusVarName(name string) bool {
	switch name {
	case "sample", "entries", "sampleEntries", "sortedEntries":
		return true
	default:
		return false
	}
}

// sampleCallName reports a call to Sample, StratifiedSample or Sorted.
func sampleCallName(fun ast.Expr) bool {
	if id, ok := fun.(*ast.Ident); ok {
		return isSampleName(id.Name)
	}
	if sel, ok := fun.(*ast.SelectorExpr); ok {
		return isSampleName(sel.Sel.Name)
	}
	return false
}

// isSampleName reports one of the three deterministic sampler names.
func isSampleName(name string) bool {
	return name == "Sample" || name == "StratifiedSample" || name == "Sorted"
}

// corpusBlock returns the body of a range or for loop over a corpus source.
func corpusBlock(n ast.Node) *ast.BlockStmt {
	switch v := n.(type) {
	case *ast.RangeStmt:
		if corpusExpr(v.X) {
			return v.Body
		}
	case *ast.ForStmt:
		if corpusExpr(v.Cond) || corpusExpr(v.Init) {
			return v.Body
		}
	}
	return nil
}

// appendBranchFindings records continue and break statements of a loop body.
func appendBranchFindings(out *[]Finding, fset *token.FileSet,
	body *ast.BlockStmt, rule string) {
	ast.Inspect(body, func(n ast.Node) bool {
		branch, ok := n.(*ast.BranchStmt)
		if !ok {
			_, isLit := n.(*ast.FuncLit)
			return !isLit
		}
		if branch.Tok == token.CONTINUE || branch.Tok == token.BREAK {
			*out = append(*out, newFinding(fset, branch.Pos(), rule,
				"corpus loop must not skip an entry"))
		}
		return true
	})
}

// appendReturnFindings records provenance-gated early returns of a corpus
// loop. A plain error return is legitimate propagation, not a provenance
// exclusion, so only a return reached through a provenance condition counts.
func appendReturnFindings(out *[]Finding, fset *token.FileSet,
	body *ast.BlockStmt) {
	ast.Inspect(body, func(n ast.Node) bool {
		branch, ok := n.(*ast.IfStmt)
		if !ok {
			_, isLit := n.(*ast.FuncLit)
			return !isLit
		}
		if !provenanceUsed(branch.Cond) {
			return true
		}
		ast.Inspect(branch.Body, func(inner ast.Node) bool {
			if _, isLit := inner.(*ast.FuncLit); isLit {
				return false
			}
			if ret, isReturn := inner.(*ast.ReturnStmt); isReturn {
				*out = append(*out, newFinding(fset, ret.Pos(),
					ruleCorpusExclusion,
					"provenance-gated return excludes corpus entries"))
			}
			return true
		})
		return true
	})
}

// forbiddenCall recognises one forbidden selector call.
func forbiddenCall(call *ast.CallExpr, imports map[string]bool) bool {
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}
	if x, isIdent := sel.X.(*ast.Ident); isIdent &&
		forbiddenSelector(x.Name, sel.Sel.Name) {
		return true
	}
	if _, isNested := sel.X.(*ast.SelectorExpr); isNested &&
		httpCallName(sel.Sel.Name) && imports["net/http"] {
		return true
	}
	return false
}

// forbiddenSelector rejects specific package-level forbidden selectors.
func forbiddenSelector(pkg, name string) bool {
	switch pkg {
	case "time":
		return name == "Now" || name == "Since"
	case "os":
		return name == "Environ"
	case "net":
		return name == "Dial"
	case "http":
		return httpCallName(name)
	default:
		return false
	}
}

// httpCallName reports one of the forbidden network client entry points.
func httpCallName(name string) bool {
	switch name {
	case "Get", "Post", "PostForm", "Do", "NewRequest",
		"NewRequestWithContext", "Head":
		return true
	default:
		return false
	}
}

// importSet records every import path of a parsed file.
func importSet(file *ast.File) map[string]bool {
	out := make(map[string]bool)
	for _, spec := range file.Imports {
		out[strings.Trim(spec.Path.Value, "\"")] = true
	}
	return out
}

// isCommandMain reports a path such as cmd/e2e/main.go.
func isCommandMain(name string) bool {
	slashed := filepath.ToSlash(name)
	return strings.HasPrefix(slashed, "cmd/") &&
		strings.HasSuffix(slashed, "/main.go")
}
