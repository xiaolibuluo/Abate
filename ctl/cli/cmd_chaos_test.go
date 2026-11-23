/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/
package cli

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"abate/kernel/injection"
	"abate/kernel/mutation"
	"abate/lang/targetgen"
)

// chaosWorkspace materializes one deterministic Go target module under t's
// temporary directory, so the real mutation backend judges a workspace that
// is byte-identical across runs and machines.
func chaosWorkspace(t *testing.T) string {
	t.Helper()
	target, err := targetgen.Generate(targetgen.Spec{
		Lang: targetgen.LangGo, Module: "example.com/chaostarget", Funcs: 6,
	}, 11)
	if err != nil {
		t.Fatalf("generate target: %v", err)
	}
	dir := filepath.Join(t.TempDir(), "ws")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	for name, data := range target.Files {
		if err := os.WriteFile(filepath.Join(dir, name), data, 0o644); err != nil {
			t.Fatalf("write %s: %v", name, err)
		}
	}
	return dir
}

func TestRunChaosMutationRealKernels(t *testing.T) {
	dir := chaosWorkspace(t)
	summary, report, backend, err := runChaosMutation(context.Background(),
		42, dir, "go", 2, false, defaultChaosMutationPorts())
	if err != nil {
		t.Fatalf("runChaosMutation: %v", err)
	}
	if summary.Operators != 12 {
		t.Errorf("Operators = %d, want 12", summary.Operators)
	}
	if len(summary.Classes) != 12 {
		t.Errorf("len(Classes) = %d, want 12", len(summary.Classes))
	}
	if summary.Frames <= 0 {
		t.Errorf("Frames = %d, want > 0 (injection emits records per fault)", summary.Frames)
	}
	if len(summary.Digest) != 64 {
		t.Errorf("Digest = %q, want 64-hex SHA-256", summary.Digest)
	}
	if backend == nil {
		t.Fatal("backend is nil")
	}
	if len(report.Verdicts) == 0 {
		t.Fatal("no mutants judged")
	}
	if report.TargetFingerprint != backend.Target().Fingerprint() {
		t.Error("report target fingerprint diverges from the real backend target")
	}
}

func TestRunChaosMutationDeterministic(t *testing.T) {
	dir := chaosWorkspace(t)
	summary1, report1, _, err := runChaosMutation(context.Background(),
		42, dir, "go", 2, false, defaultChaosMutationPorts())
	if err != nil {
		t.Fatalf("first run: %v", err)
	}
	summary2, report2, _, err := runChaosMutation(context.Background(),
		42, dir, "go", 2, false, defaultChaosMutationPorts())
	if err != nil {
		t.Fatalf("second run: %v", err)
	}
	if !reflect.DeepEqual(summary1, summary2) {
		t.Errorf("chaos summary diverged: %+v vs %+v", summary1, summary2)
	}
	if report1.Fingerprint != report2.Fingerprint {
		t.Error("report fingerprint diverged across identical seeds")
	}
}

// TestChaosMutationMatchesPureMutation proves the orchestration claim from
// the --chaos help text: injecting and rolling back the chaos faults before
// judging mutants never perturbs the mutation report bytes.
func TestChaosMutationMatchesPureMutation(t *testing.T) {
	dir := chaosWorkspace(t)
	_, chaosReport, _, err := runChaosMutation(context.Background(),
		42, dir, "go", 2, false, defaultChaosMutationPorts())
	if err != nil {
		t.Fatalf("chaos run: %v", err)
	}
	pureReport, _, err := runMutateExperiment(42, dir, "go", 2, false)
	if err != nil {
		t.Fatalf("pure mutation run: %v", err)
	}
	if chaosReport.Fingerprint != pureReport.Fingerprint {
		t.Errorf("chaos fingerprint %x != pure fingerprint %x",
			chaosReport.Fingerprint, pureReport.Fingerprint)
	}
	if len(chaosReport.Verdicts) != len(pureReport.Verdicts) {
		t.Fatalf("verdict count %d != %d", len(chaosReport.Verdicts), len(pureReport.Verdicts))
	}
	for i, v := range chaosReport.Verdicts {
		if v != pureReport.Verdicts[i] {
			t.Errorf("verdict[%d] = %s (chaos) vs %s (pure)", i, v, pureReport.Verdicts[i])
		}
	}
}

func TestRunChaosMutationErrorBranches(t *testing.T) {
	dir := chaosWorkspace(t)
	wantSetup := errors.New("setup boom")
	wantApply := errors.New("apply boom")
	wantEngine := errors.New("engine boom")

	ports := defaultChaosMutationPorts()
	ports.setup = func(uint64) (*injection.Injector, error) { return nil, wantSetup }
	if _, _, _, err := runChaosMutation(context.Background(), 42, dir, "go", 1, false, ports); !errors.Is(err, wantSetup) {
		t.Errorf("setup branch: error = %v, want wrapped %v", err, wantSetup)
	}

	ports = defaultChaosMutationPorts()
	ports.apply = func(context.Context, *injection.Injector) error { return wantApply }
	if _, _, _, err := runChaosMutation(context.Background(), 42, dir, "go", 1, false, ports); !errors.Is(err, wantApply) {
		t.Errorf("apply branch: error = %v, want wrapped %v", err, wantApply)
	}

	ports = defaultChaosMutationPorts()
	ports.engine = func(ctx context.Context, seed uint64, path, lang string, workers int, native bool) (mutation.Report, mutationBackend, error) {
		return mutation.Report{}, nil, wantEngine
	}
	if _, _, _, err := runChaosMutation(context.Background(), 42, dir, "go", 1, false, ports); !errors.Is(err, wantEngine) {
		t.Errorf("engine branch: error = %v, want wrapped %v", err, wantEngine)
	}
}

func TestMutateChaosReportIncludesSummary(t *testing.T) {
	dir := chaosWorkspace(t)
	var stdout, stderr strings.Builder
	code := cmdMutate([]string{dir, "--chaos", "--parallel", "2", "--format", "json"},
		&stdout, &stderr)
	if code != exitOK {
		t.Fatalf("cmdMutate --chaos exit = %d, stderr: %s", code, stderr.String())
	}
	out := stdout.String()
	for _, want := range []string{
		`"chaos"`, `"operators": 12`, `"frames"`, `"digest"`, `"mutants"`,
	} {
		if !strings.Contains(out, want) {
			t.Errorf("chaos report missing %q\noutput:\n%s", want, out)
		}
	}
}

func TestMutateWithoutChaosKeepsReportShape(t *testing.T) {
	dir := chaosWorkspace(t)
	var stdout, stderr strings.Builder
	code := cmdMutate([]string{dir, "--parallel", "2", "--format", "json"},
		&stdout, &stderr)
	if code != exitOK {
		t.Fatalf("cmdMutate exit = %d, stderr: %s", code, stderr.String())
	}
	if strings.Contains(stdout.String(), `"chaos"`) {
		t.Error("chaos-free report must not carry a chaos block")
	}
}
