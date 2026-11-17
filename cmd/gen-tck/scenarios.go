/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package main

import (
	"fmt"

	"abate/ces/schema"
	"abate/internal/yamlmini"
)

// scenario is one candidate document plus the spec clause it targets and a hint
// used only when the document is fully valid (which family to file it under).
type scenario struct {
	spec string
	e    schema.Experiment
	cat  string // "roundtrip" | "validate-ok" | "parse-ok" (valid docs only)
	raw  string // when set, used verbatim as the document (for malformed YAML)
}

// ---- document builders ----------------------------------------------------

func baseExp() schema.Experiment {
	return schema.Experiment{
		APIVersion: schema.APIVersion,
		Kind:       schema.KindExperiment,
		Metadata:   schema.Metadata{Name: "tck-base"},
		Spec: schema.ExperimentSpec{
			Target: schema.Target{Kind: schema.TargetSimulator, Selector: map[string]string{"node": "n1"}},
			SteadyState: []schema.SteadyState{{
				ProbeRef:    "p1",
				Expectation: schema.Expectation{Type: schema.ExpectValue, Value: "ok"},
				Tolerance:   schema.Tolerance{Type: schema.ToleranceAbsolute, Value: 1},
				HoldFor:     2 * base,
			}},
			Method:   schema.Method{Expression: "ε"},
			Probes:   []schema.Probe{{ID: "p1", Kind: schema.ProbeMetric, Spec: schema.ProbeSpec{Query: "up", Expect: "1"}, Interval: base}},
			Rollback: schema.Rollback{Strategy: schema.StrategyAuto},
			Schedule: schema.Schedule{Continuous: &schema.ContinuousSpec{Duration: base}},
		},
	}
}

func clone(e schema.Experiment) schema.Experiment {
	cp := e
	cp.Metadata.Labels = nil
	cp.Metadata.Annotations = nil
	cp.Spec.Target.Selector = map[string]string{}
	for k, v := range e.Spec.Target.Selector {
		cp.Spec.Target.Selector[k] = v
	}
	cp.Spec.SteadyState = append([]schema.SteadyState(nil), e.Spec.SteadyState...)
	cp.Spec.Probes = append([]schema.Probe(nil), e.Spec.Probes...)
	cp.Spec.Method.Operators = append([]schema.OperatorDef(nil), e.Spec.Method.Operators...)
	return cp
}

// ---- scenario groups ------------------------------------------------------

func parsePositive() []scenario {
	var out []scenario
	for i := 0; i < 16; i++ {
		e := clone(baseExp())
		e.Metadata.Name = fmt.Sprintf("tck-parse-ok-%d", i)
		out = append(out, scenario{spec: "03-ces-spec.md §3.3 (valid envelope parses)", e: e, cat: "parse-ok"})
	}
	return out
}

func parseNegative() []scenario {
	var out []scenario
	// CES-PARSE-001: unknown field rejected by the strict decoder.
	m := map[string]any{}
	b, _ := yamlmini.Marshal(baseExp())
	_ = yamlmini.Unmarshal(b, &m)
	m["unknownField"] = "rejected"
	rawUnknown, _ := yamlmini.Marshal(m)
	out = append(out, scenario{spec: "03-ces-spec.md §3 (unknown field rejected)", raw: string(rawUnknown)})

	// CES-PARSE-002: unsupported apiVersion.
	e := clone(baseExp())
	e.APIVersion = "ces.tencent.com/v2"
	out = append(out, scenario{spec: "03-ces-spec.md §3.3 (version negotiation)", e: e})

	// CES-PARSE-005: malformed YAML.
	out = append(out, scenario{spec: "03-ces-spec.md §3 (malformed document)",
		raw: "apiVersion: ces.tencent.com/v1\n kind: Experiment\n  bad indent\n"})

	// CES-PARSE-006: missing metadata.name.
	e = clone(baseExp())
	e.Metadata.Name = ""
	out = append(out, scenario{spec: "03-ces-spec.md §3.3 (metadata.name required)", e: e})

	// CES-PARSE-007: steady state references an undeclared probe (normalize).
	e = clone(baseExp())
	e.Spec.SteadyState[0].ProbeRef = "ghost"
	out = append(out, scenario{spec: "03-ces-spec.md §3 (steady state references declared probe)", e: e})

	// CES-PARSE-008: two schedule modes set (normalize).
	e = clone(baseExp())
	e.Spec.Schedule = schema.Schedule{
		Continuous: &schema.ContinuousSpec{Duration: base},
		Once:       &schema.OnceSpec{At: base},
	}
	out = append(out, scenario{spec: "03-ces-spec.md §3 (schedule mutual exclusion)", e: e})

	// CES-PARSE-009: expression syntax error.
	e = clone(baseExp())
	e.Spec.Method.Expression = "a ∘"
	out = append(out, scenario{spec: "03-ces-spec.md §3 (expression syntax)", e: e})

	// CES-PARSE-002 negative variety.
	for i := 0; i < 15; i++ {
		e := clone(baseExp())
		e.APIVersion = "ces.tencent.com/v9"
		e.Metadata.Name = fmt.Sprintf("tck-parse-v2neg-%d", i)
		out = append(out, scenario{spec: "03-ces-spec.md §3 (version negotiation negative)", e: e})
	}
	// CES-PARSE-005 negative variety: duplicate probe ids caught at normalize.
	for i := 0; i < 15; i++ {
		e := clone(baseExp())
		e.Spec.Probes = []schema.Probe{
			{ID: "p1", Kind: schema.ProbeMetric, Spec: schema.ProbeSpec{Query: "up", Expect: "1"}, Interval: base},
			{ID: "p1", Kind: schema.ProbeMetric, Spec: schema.ProbeSpec{Query: "up", Expect: "1"}, Interval: base},
		}
		e.Spec.SteadyState[0].ProbeRef = "p1"
		e.Metadata.Name = fmt.Sprintf("tck-parse-dupprobe-%d", i)
		out = append(out, scenario{spec: "03-ces-spec.md §3 (probe id uniqueness)", e: e})
	}
	// CES-PARSE-005 negative variety: non-positive probe interval.
	for i := 0; i < 15; i++ {
		e := clone(baseExp())
		e.Spec.Probes = []schema.Probe{{
			ID: "p1", Kind: schema.ProbeMetric, Spec: schema.ProbeSpec{Query: "up", Expect: "1"}, Interval: 0,
		}}
		e.Metadata.Name = fmt.Sprintf("tck-parse-probeint-%d", i)
		out = append(out, scenario{spec: "03-ces-spec.md §3 (probe interval positive)", e: e})
	}
	return out
}
