/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package validate

import "abate/ces/schema"

// probeRules returns the five probe rules in registration order.
func probeRules() []Rule {
	return []Rule{
		newRule(RuleProbe001, checkProbe001),
		newRule(RuleProbe002, checkProbe002),
		newRule(RuleProbe003, checkProbe003),
		newRule(RuleProbe004, checkProbe004),
		newRule(RuleProbe005, checkProbe005),
	}
}

// checkProbe001 enforces unique probe identifiers within the document.
func checkProbe001(exp *schema.Experiment) []Violation {
	seen := make(map[string]bool, len(exp.Spec.Probes))
	var out []Violation
	for i, p := range exp.Spec.Probes {
		if seen[p.ID] {
			out = append(out, errorAt(RuleProbe001,
				path("spec.probes", i)+".id", "duplicate probe id"))
		}
		seen[p.ID] = true
	}
	return out
}

// checkProbe002 requires a positive probe interval.
func checkProbe002(exp *schema.Experiment) []Violation {
	var out []Violation
	for i, p := range exp.Spec.Probes {
		if p.Interval <= 0 {
			out = append(out, errorAt(RuleProbe002,
				path("spec.probes", i)+".interval", "interval must be positive"))
		}
	}
	return out
}

// checkProbe003 rejects external URLs for http probes (simulator iron rule).
func checkProbe003(exp *schema.Experiment) []Violation {
	var out []Violation
	for i, p := range exp.Spec.Probes {
		if p.Kind != schema.ProbeHTTP {
			continue
		}
		if !isSimulatorURL(p.Spec.URL) {
			out = append(out, errorAt(RuleProbe003,
				path("spec.probes", i)+".spec.url",
				"http probe url must target the in-simulator domain"))
		}
	}
	return out
}

// checkProbe004 restricts shell probes to the safe in-simulator command subset.
func checkProbe004(exp *schema.Experiment) []Violation {
	var out []Violation
	for i, p := range exp.Spec.Probes {
		if p.Kind != schema.ProbeShell {
			continue
		}
		if !isSafeShellCommand(p.Spec.Command) {
			out = append(out, errorAt(RuleProbe004,
				path("spec.probes", i)+".spec.command",
				"shell command is outside the safe simulator subset"))
		}
	}
	return out
}

// checkProbe005 validates expectation values: HTTP status 100-599, exit codes
// 0-255, and a non-empty expect for metric/prom probes.
func checkProbe005(exp *schema.Experiment) []Violation {
	var out []Violation
	for i, p := range exp.Spec.Probes {
		at := path("spec.probes", i) + ".spec"
		switch p.Kind {
		case schema.ProbeHTTP:
			if p.Spec.ExpectStatus < 100 || p.Spec.ExpectStatus > 599 {
				out = append(out, errorAt(RuleProbe005, at+".expectStatus",
					"http expectStatus must be within 100-599"))
			}
		case schema.ProbeShell:
			if p.Spec.ExpectExitCode < 0 || p.Spec.ExpectExitCode > 255 {
				out = append(out, errorAt(RuleProbe005, at+".expectExitCode",
					"shell expectExitCode must be within 0-255"))
			}
		case schema.ProbeMetric, schema.ProbeProm:
			if p.Spec.Expect == "" {
				out = append(out, errorAt(RuleProbe005, at+".expect",
					"metric expect must be non-empty"))
			}
		default:
			out = append(out, errorAt(RuleProbe005, path("spec.probes", i)+".kind",
				"unknown probe kind"))
		}
	}
	return out
}
