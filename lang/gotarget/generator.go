/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package gotarget generates deterministic Go mutation-testing target
// modules. Given one seed it materializes a complete, self-contained Go
// module whose functions carry every mutable structure of the seven
// operator classes (arithmetic, boundary, constant, control flow, return
// value, condition negation and call removal), each guarded by real tests
// asserting the pristine behavior. The same seed always yields byte
// identical file bytes, so a generated target can be replayed, fingerprinted
// and fed to `abate mutate` without any drift.
package gotarget

import (
	"errors"
	"fmt"

	"abate/sim/rand"
)

// Spec errors reported before any generation work starts.
var (
	ErrModulePath   = errors.New("gotarget: module path must be non-empty")
	ErrFuncCount    = errors.New("gotarget: function count out of range")
	ErrUnknownLevel = errors.New("gotarget: unknown difficulty level")
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

// Spec is the validated input of one target generation.
type Spec struct {
	// Module is the module path written into go.mod, for example
	// "example.com/target". Only emptiness is rejected here.
	Module string
	// Package is the Go package name shared by every generated file.
	Package string
	// Funcs is the number of target functions to generate.
	Funcs int
	// Level is the difficulty profile steering structure observability.
	Level string
}

// Target is the deterministic result of one generation: the module-relative
// file bytes plus the bookkeeping a reporter needs.
type Target struct {
	Spec   Spec
	Seed   uint64
	Files  map[string][]byte
	Funcs  []FunctionInfo
	Tokens int
}

// FunctionInfo records one generated function: its name, the structure
// family that shapes its body and the name of the test function guarding it.
type FunctionInfo struct {
	Name      string
	Structure string
	TestName  string
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

// Validate checks the spec fields and fills defaults left unset.
func (s Spec) Validate() (Spec, error) {
	if s.Module == "" {
		return s, ErrModulePath
	}
	if s.Package == "" {
		s.Package = "target"
	}
	if !ValidPackageName(s.Package) {
		return s, fmt.Errorf("%w: %q", ErrPackageName, s.Package)
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
	return s, nil
}

// Generate builds the target module for one seed deterministically.
//
// Every random choice (operand constants, loop bounds, edge-case pinning)
// flows through one SplitMix64 stream seeded by seed, in a fixed traversal
// order, so the output bytes depend on the seed and the spec only — never on
// map iteration order, goroutine timing or the clock.
func Generate(spec Spec, seed uint64) (Target, error) {
	spec, err := spec.Validate()
	if err != nil {
		return Target{}, err
	}
	stream := rand.New(seed)
	gen := &generator{spec: spec, stream: stream}
	return gen.build(seed), nil
}

// generator carries the per-run state: the spec, the deterministic stream
// and the accumulated function records.
type generator struct {
	spec   Spec
	stream *rand.Stream
	funcs  []FunctionInfo
}

// build renders the whole module: one source file with every function, one
// test file with every guarding test, plus go.mod.
func (g *generator) build(seed uint64) Target {
	src := &sourceBuilder{pkg: g.spec.Package}
	tst := &testBuilder{pkg: g.spec.Package}
	for i := 0; i < g.spec.Funcs; i++ {
		name := fmt.Sprintf("Target%02d", i)
		structure := g.pickStructure(i)
		body := buildFunction(g.stream, name, structure, g.spec.Level)
		src.addFunction(body.source)
		tst.addFunction(body.test)
		g.funcs = append(g.funcs, FunctionInfo{
			Name:      name,
			Structure: structure,
			TestName:  "Test" + name,
		})
	}
	source := src.render()
	files := map[string][]byte{
		"target.go":      []byte(source),
		"target_test.go": []byte(tst.render()),
		"go.mod":         []byte(renderGoMod(g.spec.Module)),
	}
	return Target{
		Spec:   g.spec,
		Seed:   seed,
		Files:  files,
		Funcs:  g.funcs,
		Tokens: countMutableTokens([]byte(source)),
	}
}

// pickStructure chooses the structure family of the i-th function. The
// seven families are cycled deterministically so every family appears at
// least floor(Funcs/7) times regardless of the level; the level only steers
// how observable each structure is.
func (g *generator) pickStructure(i int) string {
	return structureOrder[i%len(structureOrder)]
}

// structureOrder is the fixed cycle of structure families, matching the
// seven operator classes of kernel/mutation.
var structureOrder = []string{
	"arith",
	"boundary",
	"constant",
	"controlflow",
	"returnvalue",
	"negate",
	"callremove",
}
