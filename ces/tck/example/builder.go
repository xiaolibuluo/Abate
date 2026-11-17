/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package example

import (
	"fmt"

	"abate/ces/schema"
	"abate/internal/yamlmini"
)

// Builder is a small, dependency-light producer that assembles a valid CES
// document using the public schema types. A third party ships its own authoring
// tooling; Builder is the minimal illustration of that surface.
type Builder struct {
	exp schema.Experiment
}

// NewBuilder starts a CES document for the named experiment on a simulator node.
func NewBuilder(name, node string) *Builder {
	b := &Builder{}
	b.exp.APIVersion = schema.Version{}.Current()
	b.exp.Kind = "Experiment"
	b.exp.Metadata.Name = name
	b.exp.Spec.Target = schema.Target{
		Kind:     schema.TargetSimulator,
		Selector: map[string]string{"node": node},
	}
	b.exp.Spec.Method = schema.Method{Expression: "ε"}
	b.exp.Spec.Rollback = schema.Rollback{Strategy: schema.StrategyAuto}
	b.exp.Spec.Schedule = schema.Schedule{
		Continuous: &schema.ContinuousSpec{Duration: schema.VirtualDuration(1_000_000_000)},
	}
	return b
}

// WithSteadyState appends one steady-state hypothesis referencing a probe.
func (b *Builder) WithSteadyState(probeRef, value string, holdSec int64) *Builder {
	b.exp.Spec.SteadyState = append(b.exp.Spec.SteadyState, schema.SteadyState{
		ProbeRef:    probeRef,
		Expectation: schema.Expectation{Type: schema.ExpectValue, Value: value},
		Tolerance:   schema.Tolerance{Type: schema.ToleranceAbsolute, Value: 1},
		HoldFor:     schema.VirtualDuration(holdSec * 1_000_000_000),
	})
	return b
}

// WithMetricProbe declares a metric probe sampled every intervalSec seconds.
func (b *Builder) WithMetricProbe(id, query, expect string, intervalSec int64) *Builder {
	b.exp.Spec.Probes = append(b.exp.Spec.Probes, schema.Probe{
		ID:       id,
		Kind:     schema.ProbeMetric,
		Spec:     schema.ProbeSpec{Query: query, Expect: expect},
		Interval: schema.VirtualDuration(intervalSec * 1_000_000_000),
	})
	return b
}

// WithMethodExpression overrides the perturbation algebra expression.
func (b *Builder) WithMethodExpression(expr string) *Builder {
	b.exp.Spec.Method.Expression = expr
	return b
}

// Build returns the assembled experiment.
func (b *Builder) Build() *schema.Experiment { return &b.exp }

// Marshal serializes the assembled experiment to canonical YAML bytes.
func (b *Builder) Marshal() ([]byte, error) {
	out, err := yamlmini.Marshal(b.Build())
	if err != nil {
		return nil, fmt.Errorf("example: marshal CES document: %w", err)
	}
	return out, nil
}
