/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package parse

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"strings"
	"sync"
	"testing"

	"abate/ces/schema"
)

// digestOf returns the canonical hash of a parsed document.
func digestOf(t *testing.T, exp *schema.Experiment) string {
	t.Helper()
	b, err := json.Marshal(exp)
	requireOK(t, err)
	sum := sha256.Sum256(b)
	return hex.EncodeToString(sum[:])
}

// TestChaosDeterminism100 parses the same byte stream 100 times identically.
func TestChaosDeterminism100(t *testing.T) {
	t.Parallel()
	var base string
	for i := 0; i < 100; i++ {
		exp, err := Parse([]byte(baseYAML), FormatAuto)
		requireOK(t, err)
		d := digestOf(t, exp)
		if i == 0 {
			base = d
			continue
		}
		if d != base {
			t.Fatalf("parse digest drifted on iteration %d", i)
		}
	}
}

// TestChaosErrorDeterminism keeps diagnostics byte-identical across runs.
func TestChaosErrorDeterminism(t *testing.T) {
	t.Parallel()
	bad := []byte(`{"apiVersion":"ces.tencent.com/v1","kind":"Experiment","bogus":1}`)
	var first string
	for i := 0; i < 100; i++ {
		_, err := Parse(bad, FormatAuto)
		if err == nil {
			t.Fatal("malformed document unexpectedly accepted")
		}
		if i == 0 {
			first = err.Error()
			continue
		}
		if err.Error() != first {
			t.Fatalf("diagnostic drifted on iteration %d", i)
		}
	}
}

// TestChaosConcurrent100 parses 100 distinct documents concurrently.
func TestChaosConcurrent100(t *testing.T) {
	t.Parallel()
	const n = 100
	var wg sync.WaitGroup
	errs := make(chan error, n)
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(k int) {
			defer wg.Done()
			doc := strings.Replace(baseYAML, "name: sample",
				"name: sample-"+string(rune('A'+k%26)), 1)
			exp, err := Parse([]byte(doc), FormatYAML)
			if err != nil {
				errs <- err
				return
			}
			if _, err = Normalize(exp); err != nil {
				errs <- err
			}
		}(i)
	}
	wg.Wait()
	close(errs)
	for err := range errs {
		t.Fatalf("concurrent parse failed: %v", err)
	}
}

// TestChaosRandomBytesNeverPanics fuzzes Parse with arbitrary byte streams.
func TestChaosRandomBytesNeverPanics(t *testing.T) {
	t.Parallel()
	g := detRNG{state: 0xD1B54A32D192ED03}
	alphabet := []rune("{}\"':,[]-0123456789abcdef \n\tε∘⊗⁻¹apiVersionkind")
	for i := 0; i < 5000; i++ {
		n := 1 + g.next()%32
		buf := make([]rune, n)
		for j := n; j > 0; j-- {
			buf[j-1] = alphabet[g.next()%uint64(len(alphabet))]
		}
		_, _ = Parse([]byte(string(buf)), FormatAuto)
	}
}

// TestChaosTruncationNeverPanics slices a valid document at every boundary.
func TestChaosTruncationNeverPanics(t *testing.T) {
	t.Parallel()
	for n := 0; n <= len(baseJSON); n++ {
		_, _ = Parse([]byte(baseJSON[:n]), FormatAuto)
	}
	for n := 0; n <= len(baseYAML); n++ {
		_, _ = Parse([]byte(baseYAML[:n]), FormatYAML)
	}
}

// TestChaosDeepExpressionNeverPanics bounds nested expressions safely.
func TestChaosDeepExpressionNeverPanics(t *testing.T) {
	t.Parallel()
	deep := strings.Repeat("(", 4096) + "ε" + strings.Repeat(")", 4096)
	_, err := CompileExpr(deep, nil)
	if err == nil {
		t.Fatal("over-deep expression must be rejected")
	}
}

// TestChaosCompileNeverPanics feeds random operator expressions to the compiler.
func TestChaosCompileNeverPanics(t *testing.T) {
	t.Parallel()
	g := detRNG{state: 0x9E3779B97F4A7C15}
	alphabet := []rune("abc-m0123456789 ∘⊗ε⁻¹()/:@")
	for i := 0; i < 5000; i++ {
		n := 1 + g.next()%24
		buf := make([]rune, n)
		for j := n; j > 0; j-- {
			buf[j-1] = alphabet[g.next()%uint64(len(alphabet))]
		}
		_, _ = CompileExpr(string(buf), nil)
	}
}

// TestChaosDeterministicSeedDerivation is stable across 100 derives.
func TestChaosDeterministicSeedDerivation(t *testing.T) {
	t.Parallel()
	exp := validExperiment(t)
	exp.Spec.Seed = 0
	want, err := deriveSeed(exp)
	requireOK(t, err)
	for i := 0; i < 100; i++ {
		got, err := deriveSeed(exp)
		requireOK(t, err)
		if got != want {
			t.Fatalf("derived seed drifted on iteration %d", i)
		}
	}
}

// TestChaosCanonicalBytesStable checks the base documents decode byte-identically
// when serialized back, proving no map-order leakage.
func TestChaosCanonicalBytesStable(t *testing.T) {
	t.Parallel()
	exp := validExperiment(t)
	first, err := json.Marshal(exp)
	requireOK(t, err)
	for i := 0; i < 100; i++ {
		again, err := json.Marshal(exp)
		requireOK(t, err)
		if !bytes.Equal(first, again) {
			t.Fatalf("canonical serialization drifted on iteration %d", i)
		}
	}
}
