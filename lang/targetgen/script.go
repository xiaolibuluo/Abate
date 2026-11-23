/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/
package targetgen

import (
	"fmt"

	"abate/sim/rand"
)

// scriptGen is the per-language rendering contract of the non-Go
// generators. Every generated project is one self-contained source file
// whose tests are plain assertion functions invoked from an entry point;
// running the file with the language toolchain is the self-check, so no
// external test framework ever enters the delivery.
type scriptGen interface {
	// fileName returns the project-relative name of the single source file.
	fileName() string
	// header returns the deterministic file header.
	header() string
	// footer returns the entry point invoking every generated test in
	// order; names are the TestName values of the generated functions.
	footer(testNames []string) string
	// build renders one function of the given structure family plus its
	// test body source.
	build(stream *rand.Stream, name, structure, level string) scriptBody
	// buildMixed renders one mixed-structure function combining two
	// families.
	buildMixed(stream *rand.Stream, name, primary, secondary, level string) scriptBody
}

// scriptBody carries the rendered function source and test source of one
// generated target function.
type scriptBody struct {
	source string
	test   string
	pinned bool
}

// generateScript renders a single-file project for one non-Go language,
// deterministically, then packs it with the language manifest if any.
func generateScript(spec Spec, stream *rand.Stream, seed uint64, gen scriptGen) Target {
	source := gen.header()
	tests := ""
	funcs := make([]FunctionInfo, 0, spec.Funcs)
	testNames := make([]string, 0, spec.Funcs)
	pinnedCount := 0
	for i := 0; i < spec.Funcs; i++ {
		name := fmt.Sprintf("target_%02d", i)
		structure := pickStructure(spec, i)
		var body scriptBody
		if spec.Mixed > 0 && stream.Float64() < spec.Mixed {
			secondary := pickStructure(spec, i+3)
			body = gen.buildMixed(stream, name, structure, secondary, spec.Level)
			structure = structure + "+" + secondary
		} else {
			body = gen.build(stream, name, structure, spec.Level)
		}
		source += body.source
		tests += body.test
		if body.pinned {
			pinnedCount++
		}
		testName := "test_" + name
		testNames = append(testNames, testName)
		funcs = append(funcs, FunctionInfo{
			Name:      name,
			Structure: structure,
			TestName:  testName,
		})
	}
	files := map[string][]byte{
		gen.fileName(): []byte(source + tests + gen.footer(testNames)),
	}
	return Target{
		Spec:             spec,
		Seed:             seed,
		Files:            files,
		Funcs:            funcs,
		Tokens:           countGoTokens(files[gen.fileName()]),
		PinnedStructures: pinnedCount,
	}
}
