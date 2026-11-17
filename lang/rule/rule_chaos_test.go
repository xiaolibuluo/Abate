/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package rule

import (
	"crypto/sha256"
	"fmt"
	"io"
	"sync"
	"testing"

	"abate/lang/air"
)

// chaosSpecs builds a deterministic rule set without any global state.
func chaosSpecs(count int) []Spec {
	out := make([]Spec, 0, count)
	for i := 0; i < count; i++ {
		out = append(out, Spec{
			ID:        RuleID(fmt.Sprintf("GO-EG-%03d", i+1)),
			Reference: "Effective Go, Names: MixedCaps",
			Severity:  SeverityWarning, Category: CategoryNaming,
			Pattern:  fmt.Sprintf("Ident~token%d", i),
			Message:  "avoid token",
			Mutation: MutationDecl{Family: "M-NAM", Kind: "rename"},
		})
	}
	return out
}

// TestChaosRuleSetHash100 proves 100 repeated loads and checks produce one
// digest, and that concurrent read-only matching is race free.
func TestChaosRuleSetHash100(t *testing.T) {
	t.Parallel()
	specs := chaosSpecs(32)
	unit, _, _ := testTree()
	digest := func() [32]byte {
		registry := NewRegistry()
		if err := registry.Load(specs); err != nil {
			t.Fatalf("load: %v", err)
		}
		hash := sha256.New()
		for _, current := range registry.For("go") {
			_, _ = io.WriteString(hash, current.Spec().Canonical())
			for _, finding := range current.Check(unit) {
				_, _ = io.WriteString(hash, finding.Canonical())
			}
		}
		var sum [32]byte
		copy(sum[:], hash.Sum(nil))
		return sum
	}
	base := digest()
	for i := 0; i < 100; i++ {
		if got := digest(); got != base {
			t.Fatalf("rule-set digest changed at iteration %d", i)
		}
	}
	var group sync.WaitGroup
	digests := make(chan [32]byte, 16)
	for i := 0; i < 16; i++ {
		group.Add(1)
		go func() {
			defer group.Done()
			digests <- digest()
		}()
	}
	group.Wait()
	close(digests)
	for value := range digests {
		if value != base {
			t.Fatal("concurrent read-only matching changed the digest")
		}
	}
}

// TestChaosMatcherReadOnly proves a compiled matcher is reusable in parallel.
func TestChaosMatcherReadOnly(t *testing.T) {
	t.Parallel()
	pattern, err := CompilePattern("(Call $x $.. $y*)")
	if err != nil {
		t.Fatalf("compile: %v", err)
	}
	matcher := NewMatcher(pattern)
	if !matcher.Pattern().Valid() {
		t.Fatal("Matcher.Pattern() lost the compiled pattern")
	}
	unit, _, _ := testTree()
	var group sync.WaitGroup
	for i := 0; i < 24; i++ {
		group.Add(1)
		go func() {
			defer group.Done()
			for _, node := range air.Collect(unit) {
				_ = matcher.Match(node)
			}
		}()
	}
	group.Wait()
}
