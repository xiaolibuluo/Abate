/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package cli

import (
	"strings"
	"testing"

	"abate/ces/schema"
	"abate/ces/validate"
	"abate/kernel/observe"
	"abate/kernel/scheduler"
	"abate/kernel/verify"
)

// TestRunnerTimeBranches covers once, continuous and cron scheduling plus
// default and declared probe intervals.
func TestRunnerTimeBranches(t *testing.T) {
	t.Parallel()
	exp := loadFixture(t, "ok.yaml")
	if got := perturbationTime(exp); got != scheduler.VirtualTime(5*schema.Second) {
		t.Fatalf("once perturbation time = %d", got)
	}
	if got := probeInterval(exp); got != scheduler.VirtualTime(schema.Second) {
		t.Fatalf("declared probe interval = %d", got)
	}
	exp.Spec.Schedule = schema.Schedule{
		Continuous: &schema.ContinuousSpec{Duration: 3 * schema.Second},
	}
	if got := perturbationTime(exp); got != scheduler.VirtualTime(3*schema.Second) {
		t.Fatalf("continuous perturbation time = %d", got)
	}
	exp.Spec.Schedule = schema.Schedule{
		Cron: &schema.CronSpec{Expression: "* * * * *"},
	}
	if got := perturbationTime(exp); got != scheduler.Second {
		t.Fatalf("cron perturbation time = %d", got)
	}
	exp.Spec.Probes = nil
	if got := probeInterval(exp); got != scheduler.Second {
		t.Fatalf("default probe interval = %d", got)
	}
	times := probeSchedule(exp)
	if len(times) != probeSamples {
		t.Fatalf("probe schedule length = %d", len(times))
	}
	if got := buildSteadyStates(exp); len(got) != 1 || got[0].HoldFor != int64(schema.Second) {
		t.Fatalf("steady state projection = %+v", got)
	}
}

// TestRunnerErrorBranches covers unsupported probe kinds and duplicate ids.
func TestRunnerErrorBranches(t *testing.T) {
	t.Parallel()
	unknown := loadFixture(t, "ok.yaml")
	unknown.Spec.Probes[0].Kind = "quantum"
	if _, err := runExperiment(unknown); err == nil {
		t.Fatal("unsupported probe kind must fail the execution chain")
	}
	duplicate := loadFixture(t, "ok.yaml")
	duplicate.Spec.Probes = append(duplicate.Spec.Probes, duplicate.Spec.Probes[0])
	if _, err := buildPipeline(duplicate); err == nil {
		t.Fatal("duplicate probe registration must be rejected")
	}
}

// TestCascadeHandlerBranches covers apply-derived and non-apply events.
func TestCascadeHandlerBranches(t *testing.T) {
	t.Parallel()
	times := []scheduler.VirtualTime{scheduler.Second, 2 * scheduler.Second}
	handler := cascadeHandler(times, scheduler.Second)
	children, err := handler(scheduler.Event{Kind: scheduler.EventApply})
	if err != nil || len(children) != 3 {
		t.Fatalf("apply children = %d, err = %v", len(children), err)
	}
	if children[2].Kind != scheduler.EventRollback {
		t.Fatalf("last child kind = %q", children[2].Kind)
	}
	none, err := handler(scheduler.Event{Kind: scheduler.EventProbe})
	if err != nil || len(none) != 0 {
		t.Fatalf("non-apply children = %d, err = %v", len(none), err)
	}
	initial := buildInitial(loadFixture(t, "ok.yaml"))
	if len(initial) != 1 || len(initial[0].Payload) == 0 {
		t.Fatalf("initial events = %+v", initial)
	}
}

// TestProbeValueBranches covers every probe kind and the unknown-kind error.
func TestProbeValueBranches(t *testing.T) {
	t.Parallel()
	for _, c := range []struct {
		kind schema.ProbeKind
		want string
	}{
		{schema.ProbeHTTP, observe.ValueNumber},
		{schema.ProbeShell, observe.ValueNumber},
		{schema.ProbeMetric, observe.ValueText},
		{schema.ProbeProm, observe.ValueText},
	} {
		decl := &schema.Probe{
			ID: "p", Kind: c.kind,
			Spec: schema.ProbeSpec{Expect: "up", ExpectStatus: 200, ExpectExitCode: 0},
		}
		value, err := probeValue(decl)
		if err != nil || value.Kind != c.want {
			t.Fatalf("probeValue(%q) = %+v, err = %v", c.kind, value, err)
		}
		probe, err := newCESProbe(decl)
		if err != nil {
			t.Fatalf("newCESProbe(%q): %v", c.kind, err)
		}
		if probe.ID() != "p" || probe.Kind() != string(c.kind) {
			t.Fatalf("probe identity = %s/%s", probe.ID(), probe.Kind())
		}
		got, status := probe.Collect(0)
		if status != observe.StatusOK || !got.Equal(value) {
			t.Fatalf("probe Collect = %+v/%s", got, status)
		}
	}
	if _, err := probeValue(&schema.Probe{Kind: "quantum"}); err == nil {
		t.Fatal("unknown probe kind must be rejected")
	}
	if _, err := newCESProbe(&schema.Probe{Kind: "quantum"}); err == nil {
		t.Fatal("newCESProbe must reject an unknown probe kind")
	}
}

// TestReportViewHelpers covers view projection and decision mapping.
func TestReportViewHelpers(t *testing.T) {
	t.Parallel()
	views := violationViews([]validate.Violation{{
		RuleID: "X", FieldPath: "p", Message: "m", Severity: validate.SeverityWarning,
	}})
	if len(views) != 1 || views[0].Severity != "warning" || views[0].RuleID != "X" {
		t.Fatalf("violation views = %+v", views)
	}
	verdicts := verdictViews([]verify.Verdict{{
		ProbeID: "p", Outcome: verify.OutcomeHeld, Reason: "r",
		FirstSeq: 3, Timeline: []uint64{1, 2},
	}})
	if len(verdicts) != 1 || verdicts[0].Timeline[1] != 2 || verdicts[0].FirstSeq != 3 {
		t.Fatalf("verdict views = %+v", verdicts)
	}
	if len(verdictViews(nil)) != 0 || len(violationViews(nil)) != 0 {
		t.Fatal("empty inputs must produce empty views")
	}
	if !steadyViolated([]verify.Verdict{{Outcome: verify.OutcomeViolated}}) {
		t.Fatal("a VIOLATED verdict must be a business failure")
	}
	if steadyViolated(nil) || steadyViolated([]verify.Verdict{{Outcome: verify.OutcomeHeld}}) {
		t.Fatal("HELD and empty verdicts must not be business failures")
	}
	if numericLiteral("42.5") != 42.5 || numericLiteral("abc") != 0 || numericLiteral("") != 0 {
		t.Fatal("numeric literal parsing disagrees with CES canonical numbers")
	}
	console := renderReport(buildRunReport(runOutcome{
		exp:  loadFixture(t, "ok.yaml"),
		opts: runOptions{format: formatConsole},
	}), "xml")
	if !strings.Contains(string(console), "status: VALIDATED") {
		t.Fatalf("console fallback report = %s", console)
	}
}
