/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package sample

import (
	"context"
	"sync"
	"testing"

	"abate/kernel/mutation"
	"abate/plugin/host"
	"abate/plugin/signature"
)

// TestConcurrentOperatorApplyChaos runs 64 independent backends in parallel and
// compares each digest with a serial baseline for the same seed.
func TestConcurrentOperatorApplyChaos(t *testing.T) {
	store, priv := sampleStore(t)
	m := sampleManifest(t, store, priv, "sample/op-chaos", signature.PluginTypeOperator, "clock,rand")
	p := NewOperatorPlugin(m)
	const n = 64
	baselines := make([][32]byte, n)
	for i := 0; i < n; i++ {
		ctx := host.WithBackend(context.Background(), sampleBackend(uint64(i))) //nolint:gosec // i is a small non-negative loop counter.
		out, err := p.Apply(ctx, []byte("concurrent"))
		if err != nil {
			t.Fatalf("baseline %d: %v", i, err)
		}
		baselines[i] = shaOf(out)
	}
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(idx int) {
			defer wg.Done()
			ctx := host.WithBackend(context.Background(), sampleBackend(uint64(idx))) //nolint:gosec // idx is a small non-negative loop counter.
			out, err := p.Apply(ctx, []byte("concurrent"))
			if err != nil {
				t.Errorf("iteration %d: %v", idx, err)
				return
			}
			if shaOf(out) != baselines[idx] {
				t.Errorf("iteration %d diverged from its baseline", idx)
			}
		}(i)
	}
	wg.Wait()
}

// TestConcurrentReporterBuildChaos proves the SARIF reporter has no hidden
// mutable state under concurrent use.
func TestConcurrentReporterBuildChaos(t *testing.T) {
	store, priv := sampleStore(t)
	m := sampleManifest(t, store, priv, "sample/rep-chaos", signature.PluginTypeReporter, "observe")
	p := NewReporterPlugin(m)
	baseline, err := p.Build(sampleFindings(), mutation.Report{Seed: 1})
	if err != nil {
		t.Fatalf("baseline Build: %v", err)
	}
	want := shaOf(baseline)
	const n = 64
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			out, err := p.Build(sampleFindings(), mutation.Report{Seed: 1})
			if err != nil {
				t.Errorf("concurrent Build: %v", err)
				return
			}
			if shaOf(out) != want {
				t.Errorf("concurrent Build diverged")
			}
		}()
	}
	wg.Wait()
}

// TestConcurrentFactoriesChaos constructs all three plugin kinds concurrently.
func TestConcurrentFactoriesChaos(t *testing.T) {
	store, priv := sampleStore(t)
	types := []signature.PluginType{
		signature.PluginTypeOperator, signature.PluginTypeReporter, signature.PluginTypeAdapter,
	}
	factories := []host.Factory{OperatorFactory, ReporterFactory, AdapterFactory}
	const n = 48
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(idx int) {
			defer wg.Done()
			k := idx % len(types)
			m := sampleManifest(t, store, priv, "sample/factory-chaos", types[k], "clock,rand,observe")
			plugin, err := factories[k](m)
			if err != nil || plugin == nil {
				t.Errorf("iteration %d factory: %v", idx, err)
				return
			}
			if plugin.Type() != types[k] {
				t.Errorf("iteration %d produced type %q", idx, plugin.Type())
			}
		}(i)
	}
	wg.Wait()
}
