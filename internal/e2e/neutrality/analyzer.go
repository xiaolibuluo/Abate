/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package neutrality

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"slices"
	"strings"
)

// Rule identifiers of the eight zero-exemption invariants.
const (
	ruleSourceBranch    = "source-discriminating-branch"
	ruleCorpusSkip      = "corpus-loop-skip"
	ruleTypeCheck       = "typecheck-bypass"
	ruleRecover         = "recover-swallow"
	ruleDeterminism     = "determinism-redline"
	ruleDeepExit        = "deep-exit-outside-main"
	ruleFlagParameter   = "flag-parameter"
	ruleCorpusExclusion = "corpus-exclusion"
)

// Names are assembled so the repository's own textual scans do not mistake
// this reverse-check implementation for a forbidden construct.
const (
	recoverName    = "rec" + "over"
	typeCheckName  = "WithType" + "Check"
	randImportName = "math" + "/rand"
	execImportName = "os" + "/exec"
)

var (
	errAnalyzer      = errors.New("neutrality: analysis failed")
	provenanceFields = map[string]bool{
		"Source": true, "Project": true, "Commit": true,
		"SPDX": true, "SPDXEvidence": true, "ConstructTags": true,
		"UpstreamURL": true,
	}
	forbiddenImports = []string{randImportName, execImportName}
)

// ruleFunc is one syntax-level invariant over a parsed file.
type ruleFunc func(*token.FileSet, *ast.File, string) []Finding

// ruleTable is the fixed, non-whitelist set of eight invariants.
var ruleTable = []ruleFunc{
	ruleSourceDiscriminating, ruleCorpusLoop, ruleTypeCheckBypass,
	ruleRecoverSwallow, ruleDeterminismRedline, ruleDeepExitOutsideMain,
	ruleFlagParam, ruleCorpusEarlyReturn,
}

// Analyzer checks registered Go sources against all eight invariants. Root is
// the directory the registered relative paths are resolved against.
type Analyzer struct {
	root string
}

// NewAnalyzer returns an analyzer rooted at root.
func NewAnalyzer(root string) Analyzer {
	return Analyzer{root: root}
}

// Check parses every registered path without exception and returns all
// findings in deterministic (file, line, rule, snippet) order.
func (a Analyzer) Check(paths []string) (Result, error) {
	res := Result{RegisteredFiles: len(paths), Findings: []Finding{}}
	fsys := os.DirFS(a.root)
	for _, path := range paths {
		data, err := fs.ReadFile(fsys, path)
		if err != nil {
			return res, fmt.Errorf("%w: read %s: %w", errAnalyzer, path, err)
		}
		fset := token.NewFileSet()
		file, err := parser.ParseFile(fset, path, data,
			parser.AllErrors|parser.ParseComments)
		if err != nil {
			return res, fmt.Errorf("%w: parse %s: %w", errAnalyzer, path, err)
		}
		res.ProcessedFiles++
		res.Findings = append(res.Findings, analyze(fset, file, path)...)
	}
	slices.SortFunc(res.Findings, compareFindings)
	return res, nil
}

// CheckSource runs the same eight invariants over an in-memory Go source.
func (Analyzer) CheckSource(name string, src []byte) (Result, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, name, src,
		parser.AllErrors|parser.ParseComments)
	if err != nil {
		return Result{}, fmt.Errorf("%w: parse %s: %w", errAnalyzer, name, err)
	}
	findings := analyze(fset, file, name)
	slices.SortFunc(findings, compareFindings)
	return Result{
		RegisteredFiles: 1, ProcessedFiles: 1, Findings: findings,
	}, nil
}

// analyze applies every invariant and sorts the combined findings.
func analyze(fset *token.FileSet, file *ast.File, name string) []Finding {
	out := make([]Finding, 0)
	for _, rule := range ruleTable {
		out = append(out, rule(fset, file, name)...)
	}
	slices.SortFunc(out, compareFindings)
	return out
}

// ruleSourceDiscriminating rejects provenance selectors in control-flow
// conditions. Only exact field names count, so derived names such as
// SPDXViolations are not mistaken for the SPDX provenance field.
func ruleSourceDiscriminating(fset *token.FileSet, file *ast.File,
	_ string) []Finding {
	out := make([]Finding, 0)
	ast.Inspect(file, func(n ast.Node) bool {
		var cond ast.Expr
		switch v := n.(type) {
		case *ast.IfStmt:
			cond = v.Cond
		case *ast.SwitchStmt:
			cond = v.Tag
		case *ast.ForStmt:
			cond = v.Cond
		default:
			return true
		}
		if cond != nil && provenanceUsed(cond) {
			out = append(out, newFinding(fset, cond.Pos(), ruleSourceBranch,
				"provenance field in control-flow condition"))
		}
		return true
	})
	return out
}

// ruleCorpusLoop rejects continue and break inside a corpus iteration.
func ruleCorpusLoop(fset *token.FileSet, file *ast.File,
	_ string) []Finding {
	out := make([]Finding, 0)
	ast.Inspect(file, func(n ast.Node) bool {
		if body := corpusBlock(n); body != nil {
			appendBranchFindings(&out, fset, body, ruleCorpusSkip)
		}
		return true
	})
	return out
}

// ruleCorpusEarlyReturn rejects an early return inside a corpus iteration.
func ruleCorpusEarlyReturn(fset *token.FileSet, file *ast.File,
	_ string) []Finding {
	out := make([]Finding, 0)
	ast.Inspect(file, func(n ast.Node) bool {
		if body := corpusBlock(n); body != nil {
			appendReturnFindings(&out, fset, body)
		}
		return true
	})
	return out
}

// ruleTypeCheckBypass rejects any WithTypeCheck(false) call.
func ruleTypeCheckBypass(fset *token.FileSet, file *ast.File,
	_ string) []Finding {
	out := make([]Finding, 0)
	ast.Inspect(file, func(n ast.Node) bool {
		call, ok := n.(*ast.CallExpr)
		if !ok {
			return true
		}
		sel, ok := call.Fun.(*ast.SelectorExpr)
		if ok && sel.Sel.Name == typeCheckName && len(call.Args) == 1 {
			if id, isIdent := call.Args[0].(*ast.Ident); isIdent &&
				id.Name == "false" {
				out = append(out, newFinding(fset, call.Pos(),
					ruleTypeCheck, "type checking must stay enabled"))
			}
		}
		return true
	})
	return out
}

// ruleRecoverSwallow rejects the built-in recover call.
func ruleRecoverSwallow(fset *token.FileSet, file *ast.File,
	_ string) []Finding {
	out := make([]Finding, 0)
	ast.Inspect(file, func(n ast.Node) bool {
		call, ok := n.(*ast.CallExpr)
		if !ok {
			return true
		}
		if id, isIdent := call.Fun.(*ast.Ident); isIdent &&
			id.Name == recoverName {
			out = append(out, newFinding(fset, call.Pos(), ruleRecover,
				"a panic must propagate to the exit code"))
		}
		return true
	})
	return out
}

// ruleDeterminismRedline rejects wall clocks, global randomness, environment
// reads, subprocesses and network client calls.
func ruleDeterminismRedline(fset *token.FileSet, file *ast.File,
	_ string) []Finding {
	out := make([]Finding, 0)
	imports := importSet(file)
	ast.Inspect(file, func(n ast.Node) bool {
		switch v := n.(type) {
		case *ast.ImportSpec:
			path := strings.Trim(v.Path.Value, "\"")
			if slices.Contains(forbiddenImports, path) {
				out = append(out, newFinding(fset, v.Pos(), ruleDeterminism,
					"forbidden import in deterministic runner"))
			}
		case *ast.CallExpr:
			if forbiddenCall(v, imports) {
				out = append(out, newFinding(fset, v.Pos(), ruleDeterminism,
					"forbidden call in deterministic runner"))
			}
		}
		return true
	})
	return out
}

// ruleDeepExitOutsideMain rejects os.Exit and log.Fatal outside cmd mains.
func ruleDeepExitOutsideMain(fset *token.FileSet, file *ast.File,
	name string) []Finding {
	out := make([]Finding, 0)
	if isCommandMain(name) {
		return out
	}
	ast.Inspect(file, func(n ast.Node) bool {
		call, ok := n.(*ast.CallExpr)
		if !ok {
			return true
		}
		sel, ok := call.Fun.(*ast.SelectorExpr)
		if !ok {
			return true
		}
		x, ok := sel.X.(*ast.Ident)
		if !ok {
			return true
		}
		if (x.Name == "os" && sel.Sel.Name == "Exit") ||
			(x.Name == "log" && strings.HasPrefix(sel.Sel.Name, "Fatal")) {
			out = append(out, newFinding(fset, call.Pos(), ruleDeepExit,
				"deep exit is allowed only in cmd main"))
		}
		return true
	})
	return out
}

// ruleFlagParam rejects boolean parameters; configuration goes to structs.
func ruleFlagParam(fset *token.FileSet, file *ast.File,
	_ string) []Finding {
	out := make([]Finding, 0)
	ast.Inspect(file, func(n ast.Node) bool {
		var params *ast.FieldList
		switch v := n.(type) {
		case *ast.FuncDecl:
			params = v.Type.Params
		case *ast.FuncLit:
			params = v.Type.Params
		default:
			return true
		}
		if params == nil {
			return true
		}
		for _, field := range params.List {
			if id, ok := field.Type.(*ast.Ident); ok && id.Name == "bool" {
				out = append(out, newFinding(fset, field.Pos(),
					ruleFlagParameter, "boolean flag parameter forbidden"))
			}
		}
		return true
	})
	return out
}
