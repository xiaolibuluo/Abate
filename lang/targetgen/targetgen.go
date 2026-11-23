/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package targetgen generates deterministic mutation/chaos testing target
// projects for every language frontend of abate. Given one seed it renders
// a complete, self-contained project whose functions carry the mutable
// structures of the seven operator classes plus optional complex structures
// and mixed-structure compositions, each guarded by real tests. The same
// seed always yields byte-identical files, and every delivered project
// passes a real toolchain self-check before it lands on disk.
package targetgen

import (
	"errors"
	"fmt"

	"abate/sim/rand"
)

// Spec errors reported before any generation work starts.
var (
	ErrModulePath     = errors.New("targetgen: module path must be non-empty")
	ErrFuncCount      = errors.New("targetgen: function count out of range")
	ErrUnknownLevel   = errors.New("targetgen: unknown difficulty level")
	ErrUnknownLang    = errors.New("targetgen: unknown language")
	ErrPackageName    = errors.New("targetgen: invalid Go package name")
	ErrMixedRatio     = errors.New("targetgen: mixed ratio out of range")
	ErrDirExists      = errors.New("targetgen: target directory already exists")
	ErrSelfCheck      = errors.New("targetgen: generated project failed self-check")
	ErrToolchainFound = errors.New("targetgen: required toolchain not found in PATH")
)

// Function count bounds of one generated target.
const (
	MinFuncs = 1
	MaxFuncs = 64
)

// Level names accepted by ParseLevel and their mutation-score intent. The
// levels shape the generated source, not the verdict: killing always
// remains the job of the real test suite executed by the mutation backend.
const (
	// LevelEasy generates directly asserted structures with every edge case
	// pinned down, so most mutants are killed outright.
	LevelEasy = "easy"
	// LevelMedium pins the edge cases of about half the structures,
	// yielding a balanced score.
	LevelMedium = "medium"
	// LevelHard leaves most edge behavior unobserved by assertions, so more
	// mutants survive.
	LevelHard = "hard"
)

// Language identifiers of the seven abate frontends.
const (
	LangGo     = "go"
	LangPython = "py"
	LangJS     = "js"
	LangC      = "c"
	LangCSharp = "cs"
	LangRust   = "rs"
	LangKotlin = "kt"
)

// languageOrder is the canonical report order of the seven languages,
// matching the internal/corpus convention.
var languageOrder = []string{
	LangGo, LangPython, LangJS, LangC, LangCSharp, LangRust, LangKotlin,
}

// Languages returns the supported language identifiers in canonical order.
func Languages() []string {
	return append([]string(nil), languageOrder...)
}

// ParseLanguage validates a language identifier.
func ParseLanguage(lang string) (string, error) {
	for _, known := range languageOrder {
		if lang == known {
			return lang, nil
		}
	}
	return "", fmt.Errorf("%w: %q (want one of %v)", ErrUnknownLang, lang, languageOrder)
}

// ParseLevel validates a difficulty level string.
func ParseLevel(level string) (string, error) {
	switch level {
	case LevelEasy, LevelMedium, LevelHard:
		return level, nil
	}
	return "", fmt.Errorf("%w: %q (want %s|%s|%s)",
		ErrUnknownLevel, level, LevelEasy, LevelMedium, LevelHard)
}

// Spec is the validated input of one target generation.
type Spec struct {
	// Lang is the language identifier of the generated project.
	Lang string
	// Module is the module path written into go.mod (Go only); other
	// languages carry no module manifest and ignore the field.
	Module string
	// Package is the Go package name (Go only).
	Package string
	// Funcs is the number of target functions to generate.
	Funcs int
	// Level is the difficulty profile steering structure observability.
	Level string
	// Mixed is the fraction of functions rendered as mixed-structure
	// combos (2-3 structure families in one body), in (0, 1]; zero keeps
	// the one-structure-per-function layout.
	Mixed float64
	// Complex enables the complex structure families (slice, map, error,
	// nested loops, methods, interfaces) in the Go rotation. Non-Go
	// languages ignore the flag: their structure library covers the seven
	// core families only.
	Complex bool
}

// Target is the deterministic result of one generation: the project-relative
// file bytes plus the bookkeeping a reporter needs.
type Target struct {
	Spec   Spec
	Seed   uint64
	Files  map[string][]byte
	Funcs  []FunctionInfo
	Tokens int
	// PinnedStructures counts functions whose edge case is asserted; it
	// feeds the expected-score model of report.go.
	PinnedStructures int
}

// FunctionInfo records one generated function: its name, the structure
// families that shape its body and the name of the test guarding it.
type FunctionInfo struct {
	Name      string
	Structure string
	TestName  string
}

// Validate checks the spec fields and fills defaults left unset.
func (s Spec) Validate() (Spec, error) {
	lang, err := ParseLanguage(s.Lang)
	if err != nil {
		return s, err
	}
	s.Lang = lang
	if s.Lang == LangGo {
		if s.Module == "" {
			return s, ErrModulePath
		}
		if s.Package == "" {
			s.Package = "target"
		}
		if !validPackageName(s.Package) {
			return s, fmt.Errorf("%w: %q", ErrPackageName, s.Package)
		}
	}
	if s.Funcs == 0 {
		s.Funcs = 8
	}
	if s.Funcs < MinFuncs || s.Funcs > MaxFuncs {
		return s, fmt.Errorf("%w: %d (want %d..%d)",
			ErrFuncCount, s.Funcs, MinFuncs, MaxFuncs)
	}
	if s.Level == "" {
		s.Level = LevelMedium
	}
	level, err := ParseLevel(s.Level)
	if err != nil {
		return s, err
	}
	s.Level = level
	if s.Mixed < 0 || s.Mixed > 1 {
		return s, fmt.Errorf("%w: %v (want 0..1)", ErrMixedRatio, s.Mixed)
	}
	return s, nil
}

// Generate builds the target project for one seed deterministically.
//
// Every random choice (structure selection, operands, bounds, edge-case
// pinning) flows through one SplitMix64 stream seeded by seed, in a fixed
// traversal order, so the output bytes depend on the seed and the spec only.
func Generate(spec Spec, seed uint64) (Target, error) {
	spec, err := spec.Validate()
	if err != nil {
		return Target{}, err
	}
	stream := rand.New(seed)
	switch spec.Lang {
	case LangGo:
		return generateGo(spec, stream, seed), nil
	case LangPython:
		return generateScript(spec, stream, seed, pythonGen{}), nil
	case LangJS:
		return generateScript(spec, stream, seed, jsGen{}), nil
	case LangC:
		return generateScript(spec, stream, seed, cGen{}), nil
	case LangCSharp:
		return generateScript(spec, stream, seed, csGen{}), nil
	case LangRust:
		return generateScript(spec, stream, seed, rsGen{}), nil
	case LangKotlin:
		return generateScript(spec, stream, seed, ktGen{}), nil
	}
	// Unreachable: Validate admits only registered languages.
	panic("targetgen: unknown language " + spec.Lang)
}

// structureOrderGo is the fixed cycle of the seven core structure families,
// matching the operator classes of kernel/mutation.
var structureOrderGo = []string{
	"arith",
	"boundary",
	"constant",
	"controlflow",
	"returnvalue",
	"negate",
	"callremove",
}

// structureOrderComplex is the extended cycle interleaving the complex
// structure families (slice, map, error, nested loops, method, interface)
// with the core seven, used when Spec.Complex is set.
var structureOrderComplex = []string{
	"arith",
	"slice",
	"boundary",
	"map",
	"constant",
	"errorpath",
	"controlflow",
	"nested",
	"returnvalue",
	"method",
	"negate",
	"iface",
	"callremove",
}

// structureOrderScript is the family cycle of the non-Go languages: the
// seven core families only.
var structureOrderScript = structureOrderGo

// pickStructure chooses the structure family of the i-th function from the
// active cycle, so every family appears at least floor(Funcs/n) times.
func pickStructure(spec Spec, i int) string {
	if spec.Lang == LangGo && spec.Complex {
		return structureOrderComplex[i%len(structureOrderComplex)]
	}
	return structureOrderScript[i%len(structureOrderScript)]
}

// validPackageName reports whether name is a legal ASCII Go package
// identifier.
func validPackageName(name string) bool {
	if name == "" {
		return false
	}
	for i, r := range name {
		switch {
		case r >= 'a' && r <= 'z':
		case r >= 'A' && r <= 'Z':
		case r == '_':
		case r >= '0' && r <= '9' && i > 0:
		default:
			return false
		}
	}
	return true
}
