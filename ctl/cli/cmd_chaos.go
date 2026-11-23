/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/
package cli

import (
	"context"
	"fmt"

	"abate/kernel/injection"
	"abate/kernel/mutation"
)

// chaosSeedSalt decorrelates the chaos world seed from the mutation seed, so
// consuming draws for the fault world never perturbs the mutation stream.
const chaosSeedSalt uint64 = 0x517cc1b727220a95

// ChaosSummary is the deterministic chaos-phase digest of one M⊗C mutate run:
// which fault operators were injected and rolled back, how many kernel frames
// they emitted, and the injector digest proving the virtual world ended in the
// exact state the seed predicts.
type ChaosSummary struct {
	Operators int      `json:"operators"`
	Classes   []string `json:"classes"`
	Frames    int      `json:"frames"`
	Digest    string   `json:"digest"`
}

// chaosMutationPorts groups the L6/L7 entry points of the real-workspace M⊗C
// run as per-call fields, so every error branch stays unit-testable without
// weakening the real error propagation.
type chaosMutationPorts struct {
	setup  func(seed uint64) (*injection.Injector, error)
	apply  func(ctx context.Context, injector *injection.Injector) error
	engine func(ctx context.Context, seed uint64, path, lang string, workers int, native bool) (mutation.Report, mutationBackend, error)
}

// runChaosMutation orchestrates one M⊗C run against a real Go workspace: the
// L6 injection kernel first applies and rolls back every registered fault
// operator inside its fully virtual world (zero syscalls, deterministic
// frames), then the L7 mutation engine judges every scheduled mutant with
// real go build + go test executions. The two phases never share a random
// stream: the chaos world is seeded by seed^chaosSeedSalt, so the mutation
// report stays byte-identical to a chaos-free run over the same workspace.
func runChaosMutation(ctx context.Context, seed uint64, path, lang string, workers int, native bool,
	ports chaosMutationPorts,
) (ChaosSummary, mutation.Report, mutationBackend, error) {
	injector, err := ports.setup(seed ^ chaosSeedSalt)
	if err != nil {
		return ChaosSummary{}, mutation.Report{}, nil, fmt.Errorf("cli: chaos setup: %w", err)
	}
	if err := ports.apply(ctx, injector); err != nil {
		return ChaosSummary{}, mutation.Report{}, nil, fmt.Errorf("cli: chaos apply: %w", err)
	}
	digest := injector.Digest()
	summary := ChaosSummary{
		Operators: len(injector.Registry().All()),
		Classes:   injector.Registry().Classes(),
		Frames:    len(injector.Frames()),
		Digest:    digestHexOfBytes(digest[:]),
	}
	report, backend, err := ports.engine(ctx, seed, path, lang, workers, native)
	if err != nil {
		return ChaosSummary{}, mutation.Report{}, nil, err
	}
	return summary, report, backend, nil
}

// defaultChaosMutationPorts binds the orchestration to the real L6 and L7
// kernels; it is the production wiring used by cmdMutate --chaos.
func defaultChaosMutationPorts() chaosMutationPorts {
	return chaosMutationPorts{
		setup: func(seed uint64) (*injection.Injector, error) {
			world, _ := injection.NewWorld(seed)
			return injection.NewInjector(world)
		},
		apply: func(ctx context.Context, injector *injection.Injector) error {
			for _, op := range injector.Registry().All() {
				if err := injector.Inject(ctx, op); err != nil {
					return fmt.Errorf("inject %s: %w", op.ID(), err)
				}
				if err := injector.Rollback(ctx, op); err != nil {
					return fmt.Errorf("rollback %s: %w", op.ID(), err)
				}
			}
			return nil
		},
		engine: func(ctx context.Context, seed uint64, path, lang string, workers int, native bool) (mutation.Report, mutationBackend, error) {
			return runMutateExperiment(seed, path, lang, workers, native)
		},
	}
}
