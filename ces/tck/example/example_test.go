/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package example

import (
	"strings"
	"testing"

	"abate/ces/tck"
)

// TestExampleDriverConformance runs the full embedded CES-TCK case library
// against the example third-party driver and asserts it is certified
// compliant. This is the concrete demonstration that a downstream project can
// enter the compliance flow: build via the public schema, consume via the
// tck.Driver boundary, and pass every parse/validate/roundtrip case.
func TestExampleDriverConformance(t *testing.T) {
	cases, err := tck.LoadCases()
	if err != nil {
		t.Fatalf("load TCK cases: %v", err)
	}
	if len(cases) == 0 {
		t.Fatal("TCK case library is empty")
	}

	rep, err := tck.Runner{Driver: Driver{}, Name: "example/third-party"}.Run(cases)
	if err != nil {
		t.Fatalf("run TCK: %v", err)
	}

	if rep.Skipped != 0 {
		t.Errorf("example driver skipped %d cases (must support all kinds): %d total", rep.Skipped, rep.Total)
	}
	if rep.Failed != 0 {
		var sb strings.Builder
		for _, d := range rep.Details {
			if !d.Passed {
				sb.WriteString("\n  - ")
				sb.WriteString(d.ID)
				sb.WriteString(": ")
				sb.WriteString(d.Detail)
			}
		}
		t.Errorf("example driver failed %d cases:%s", rep.Failed, sb.String())
	}
	if !rep.Compliant {
		t.Fatalf("example third-party driver NOT compliant: %d/%d passed, %d failed, %d skipped",
			rep.Passed, rep.Total, rep.Failed, rep.Skipped)
	}
	t.Logf("example third-party driver compliant: %d/%d passed", rep.Passed, rep.Total)
}

// TestExampleBuilderRoundtrip proves the producer side: a document assembled by
// Builder marshals, parses and validates without violation, and re-parses
// identically. It exercises the public schema + public pipeline end to end.
func TestExampleBuilderRoundtrip(t *testing.T) {
	b := NewBuilder("example-conformance", "n1").
		WithSteadyState("p1", "ok", 2).
		WithMetricProbe("p1", "up", "1", 1).
		WithMethodExpression("ε")

	doc, err := b.Marshal()
	if err != nil {
		t.Fatalf("marshal built document: %v", err)
	}

	d := Driver{}
	if res, perr := d.Parse(doc, "yaml"); perr != nil || !res.OK {
		t.Fatalf("built document rejected by Parse: ok=%t err=%v rule=%s", res.OK, perr, res.RuleID)
	}

	viols, verr := d.Validate(doc)
	if verr != nil {
		t.Fatalf("validate built document: %v", verr)
	}
	if len(viols) != 0 {
		t.Fatalf("built document has %d unexpected violations: %v", len(viols), viols)
	}

	// Re-parse the same bytes; acceptance must be stable (no divergence).
	if res, perr := d.Parse(doc, "yaml"); perr != nil || !res.OK {
		t.Fatalf("built document diverged on re-parse: ok=%t err=%v", res.OK, perr)
	}
}
