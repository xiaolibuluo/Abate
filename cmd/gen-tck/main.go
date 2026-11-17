/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Command gen-tck emits the data-driven TCK case library under ces/tck/cases/.
//
// It builds candidate CES documents, runs them through the real parse/normalize/
// validate pipeline, and writes a YAML case whose expected outcome is whatever the
// pipeline actually produces. The case kind (parse/validate/roundtrip) is derived
// from the stage at which the document first fails, so the neutral TCK runner's
// per-kind driver method always reproduces the expected result. This guarantees
// every emitted case is self-consistent and passes the bundled AbateDriver.
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

const base = 1000000000 // 1e9 nanoseconds == 1 second.

func main() {
	out := filepath.Join("ces", "tck", "cases")
	for _, g := range []string{"parse", "validate", "roundtrip"} {
		must(os.MkdirAll(filepath.Join(out, g), 0o750))
	}

	var scenarios []scenario
	scenarios = append(scenarios, parsePositive()...)
	scenarios = append(scenarios, parseNegative()...)
	scenarios = append(scenarios, validateNegative()...)
	scenarios = append(scenarios, validatePositive()...)
	scenarios = append(scenarios, roundtrip()...)

	counters := map[string]*int{}
	get := func(prefix string) int {
		p := counters[prefix]
		if p == nil {
			v := 0
			counters[prefix] = &v
			p = &v
		}
		*p++
		return *p
	}

	written := 0
	for _, s := range scenarios {
		cf := buildCase(s, get)
		dir := filepath.Join(out, cf.kind)
		name := fmt.Sprintf("%s-%s.yaml", cf.kind, cf.id)
		must(os.WriteFile(filepath.Join(dir, name), []byte(cf.body), 0o600))
		written++
	}
	_, _ = fmt.Fprintf(os.Stdout, "gen-tck: wrote %d cases\n", written)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
