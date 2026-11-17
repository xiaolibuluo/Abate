/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package frontkit

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"abate/lang/air"
)

// FrontendBase is the template-method implementation of air.Frontend shared by
// the six language adapters. It fixes the deterministic skeleton: validate,
// sort files, normalize the dialect, tokenize, delegate reduction to
// LanguageSpec.ParseTokens, register files in path order, aggregate bounded
// diagnostics, and attach canonical unit annotations. Language differences live
// only in the injected LanguageSpec; AIR itself is never extended.
type FrontendBase struct {
	spec           LanguageSpec
	maxDiagnostics int
}

// Compile-time proof that the shared base satisfies the seven-language port.
var _ air.Frontend = (*FrontendBase)(nil)

// NewFrontendBase binds one language adapter to the shared parsing skeleton.
func NewFrontendBase(spec LanguageSpec) (*FrontendBase, error) {
	if spec == nil {
		return nil, fmt.Errorf("%w: language spec", ErrNilSpec)
	}
	return &FrontendBase{spec: spec, maxDiagnostics: air.MaxDiagnostics}, nil
}

// WithMaxDiagnostics returns a copy with a positive, bounded diagnostic cap.
func (f *FrontendBase) WithMaxDiagnostics(limit int) *FrontendBase {
	if f == nil {
		return nil
	}
	out := *f
	if limit > 0 && limit <= air.MaxDiagnostics {
		out.maxDiagnostics = limit
	}
	return &out
}

// Parse performs a full parse of every in-memory source buffer in one unit.
// Syntax failures are surfaced as diagnostics and a partial AIR root instead of
// a panic; error is reserved for contract and resource failures.
func (f *FrontendBase) Parse(unit air.SourceUnit) (*air.Unit, []air.Diagnostic, error) {
	if f == nil || f.spec == nil {
		return nil, nil, fmt.Errorf("%w: frontend", ErrNilSpec)
	}
	if err := unit.Validate(); err != nil {
		return nil, nil, err
	}
	tag, ok := CanonicalTag(f.spec.Tag())
	if !ok {
		return nil, nil, fmt.Errorf("%w: %q", ErrUnsupportedLanguage, f.spec.Tag())
	}
	dialect, err := f.spec.NormalizeDialect(unit.Dialect)
	if err != nil {
		return nil, nil, fmt.Errorf("%w: %w", ErrBadDialect, err)
	}
	lexSpec := f.spec.LexSpec()
	if lexSpec == nil {
		return nil, nil, fmt.Errorf("%w: lexical spec", ErrNilSpec)
	}
	out := air.NewUnit(air.NewPackageRef(tag, nil))
	diagnostics := make([]air.Diagnostic, 0)
	for _, src := range unit.SortedFiles() {
		lexer, err := NewLexer(lexSpec, src.Content)
		if err != nil {
			return nil, diagnostics, err
		}
		tokens, err := lexer.Tokenize()
		if err != nil {
			return nil, diagnostics, err
		}
		root, fileDiagnostics, parseErr := f.spec.ParseTokens(tokens, src.Path)
		diagnostics = append(diagnostics, fileDiagnostics...)
		if root == nil || parseErr != nil {
			root = partialRoot(len(src.Content))
			if parseErr != nil {
				diagnostics = append(diagnostics, air.Diagnostic{
					Path: src.Path, Line: 1, Column: 1,
					Severity: air.SeverityError,
					Code:     "LANG-" + strings.ToUpper(tag) + "-E-000",
					Message:  parseErr.Error(),
				})
			}
		}
		root = BindTokenSpellings(tokens, root)
		if err := out.AddFile(air.NewFileRef(src.Path, src.Content), root); err != nil {
			return nil, diagnostics, err
		}
	}
	diagnostics = normalizeDiagnostics(diagnostics, f.maxDiagnostics)
	annotateUnit(out, tag, dialect, len(diagnostics), diagnostics)
	return out, diagnostics, nil
}

// normalizeDiagnostics sorts, deduplicates nothing and truncates to the cap.
// The key is (Path, Line, Column, Code, Message) because air.Diagnostic has no
// byte-offset field; the DiagnosticSink owns the offset-aware ordering.
func normalizeDiagnostics(in []air.Diagnostic, limit int) []air.Diagnostic {
	limit = clampInt(limit, 1, air.MaxDiagnostics)
	out := slices.Clone(in)
	slices.SortStableFunc(out, func(a, b air.Diagnostic) int {
		if order := strings.Compare(a.Path, b.Path); order != 0 {
			return order
		}
		if a.Line != b.Line {
			if a.Line < b.Line {
				return -1
			}
			return 1
		}
		if a.Column != b.Column {
			if a.Column < b.Column {
				return -1
			}
			return 1
		}
		if order := strings.Compare(a.Code, b.Code); order != 0 {
			return order
		}
		return strings.Compare(a.Message, b.Message)
	})
	if len(out) > limit {
		out = out[:limit]
	}
	return out
}

// partialRoot builds a whole-file fallback root honestly marked as partial.
func partialRoot(size int) air.Node {
	start := air.NewPosition(1, 1, 0, size)
	end := air.NewPosition(1, 1, size, size)
	root := air.NewNodeAt(air.KindRegion, start, end, "")
	_ = air.Annotate(root, air.SemPartial, "true")
	return root
}

// annotateUnit writes the fixed, order-independent unit-level metadata.
func annotateUnit(out *air.Unit, tag, dialect string,
	diagnosticCount int, diagnostics []air.Diagnostic) {
	prefix := "lang." + tag + "."
	out.Annotations[prefix+"dialect"] = dialect
	out.Annotations[prefix+"files"] = strconv.Itoa(len(out.Files))
	out.Annotations[prefix+"diagnostics"] = strconv.Itoa(diagnosticCount)
	if diagnosticCount > 0 {
		out.Annotations[prefix+"firstDiagnostic"] = diagnostics[0].String()
	}
}
