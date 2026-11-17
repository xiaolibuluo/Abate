/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package cli

import (
	"fmt"

	"abate/ces/schema"
	"abate/kernel/observe"
)

// cesProbe is a deterministic in-simulator observation source derived from
// a CES probe declaration.
type cesProbe struct {
	id    string
	kind  string
	value observe.Value
}

// newCESProbe materializes a deterministic probe from a CES declaration.
func newCESProbe(p *schema.Probe) (observe.Probe, error) {
	value, err := probeValue(p)
	if err != nil {
		return nil, err
	}
	return &cesProbe{id: p.ID, kind: string(p.Kind), value: value}, nil
}

// ID returns the CES probe identifier.
func (p *cesProbe) ID() string { return p.id }

// Kind returns the CES probe kind.
func (p *cesProbe) Kind() string { return p.kind }

// Collect returns the deterministic value; the simulator never reads a clock.
func (p *cesProbe) Collect(_ int64) (observe.Value, string) {
	return p.value, observe.StatusOK
}

// probeValue maps a CES probe spec to its simulated observation value.
func probeValue(p *schema.Probe) (observe.Value, error) {
	switch p.Kind {
	case schema.ProbeHTTP:
		return observe.NewNumber(float64(p.Spec.ExpectStatus)), nil
	case schema.ProbeShell:
		return observe.NewNumber(float64(p.Spec.ExpectExitCode)), nil
	case schema.ProbeMetric, schema.ProbeProm:
		return observe.NewText(p.Spec.Expect), nil
	default:
		return observe.Value{}, fmt.Errorf("unsupported probe kind %q", p.Kind)
	}
}
