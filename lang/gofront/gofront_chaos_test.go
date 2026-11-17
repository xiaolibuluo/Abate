/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package gofront

import (
	"crypto/sha256"
	"sync"
	"testing"

	"abate/internal/determinism"
	"abate/lang/air"
)

// chaosSourceUnit is the fixed input for all chaos runs.
func chaosSourceUnit() air.SourceUnit {
	return air.SourceUnit{
		Language: "go",
		Files:    []air.SourceFile{air.NewSourceFile("chaos.go", []byte(sampleGo))},
	}
}

// TestChaosEncodeDeterminism100 requires one unique archive hash in 100 runs.
func TestChaosEncodeDeterminism100(t *testing.T) {
	t.Parallel()
	frontend := NewGoFrontend().WithTypeCheck(false)
	fn := func() []byte {
		unit, diagnostics, err := frontend.Parse(chaosSourceUnit())
		if err != nil || len(diagnostics) != 0 {
			t.Fatalf("parse drift: %v %v", err, diagnostics)
		}
		archive, err := air.NewEncoder().Encode(unit)
		if err != nil {
			t.Fatalf("encode: %v", err)
		}
		return archive
	}
	determinism.MustMatch100(t, 0xC0FFEE01, fn)
	unique := make(map[[32]byte]struct{})
	for i := 0; i < 100; i++ {
		unique[sha256.Sum256(fn())] = struct{}{}
	}
	if len(unique) != 1 {
		t.Fatalf("expected exactly one unique SHA-256, got %d", len(unique))
	}
}

// TestChaosHashStateStable checks the canonical JSON projection 100 times.
func TestChaosHashStateStable(t *testing.T) {
	t.Parallel()
	unit := mustParse(t, sampleGo)
	determinism.MustMatch100(t, 0xC0FFEE02, func() []byte {
		sum := determinism.HashState(unit)
		return sum[:]
	})
}

// TestChaosConcurrentParse runs parallel parses under the race detector.
func TestChaosConcurrentParse(t *testing.T) {
	t.Parallel()
	frontend := NewGoFrontend().WithTypeCheck(false)
	reference, diagnostics, err := frontend.Parse(chaosSourceUnit())
	if err != nil || len(diagnostics) != 0 {
		t.Fatalf("reference parse: %v %v", err, diagnostics)
	}
	want := reference.Fingerprint()
	failures := make(chan error, 64)
	var wg sync.WaitGroup
	for worker := 0; worker < 16; worker++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 20; i++ {
				unit, diagnostics, err := frontend.Parse(chaosSourceUnit())
				if err != nil || len(diagnostics) != 0 {
					failures <- air.ErrUnitMismatch
					return
				}
				if unit.Fingerprint() != want {
					failures <- air.ErrUnitMismatch
					return
				}
				if err := VerifyTokenFidelity([]byte(sampleGo), unit); err != nil {
					failures <- err
					return
				}
			}
		}()
	}
	wg.Wait()
	close(failures)
	for err := range failures {
		t.Fatalf("concurrent parse: %v", err)
	}
}

// TestChaosMalformedSourcesNeverPanics feeds deterministic noise to the parser.
func TestChaosMalformedSourcesNeverPanics(t *testing.T) {
	t.Parallel()
	state := uint64(0xD1B54A32D192ED03)
	next := func() uint64 {
		state = state*6364136223846793005 + 1442695040888963407
		return state
	}
	alphabet := []byte("package func {}\n\t()[]\";'0123456789abcXYZ\x00\xff//*")
	frontends := []GoFrontend{
		NewGoFrontend(),
		NewGoFrontend().WithTypeCheck(false),
	}
	for i := 0; i < 600; i++ {
		size := int(next() % 120) //nolint:gosec // the modulo bounds the conversion.
		buf := make([]byte, size)
		for j := range buf {
			buf[j] = alphabet[next()%uint64(len(alphabet))]
		}
		for _, frontend := range frontends {
			unit, diagnostics, err := frontend.Parse(air.SourceUnit{
				Language: "go",
				Files:    []air.SourceFile{air.NewSourceFile("noise.go", buf)},
			})
			if err != nil || unit == nil || unit.Roots[0] == nil {
				t.Fatalf("iteration %d produced no unit: %v", i, err)
			}
			for _, diagnostic := range diagnostics {
				if diagnostic.Code == "" {
					t.Fatalf("iteration %d produced an untagged diagnostic", i)
				}
			}
		}
	}
}

// TestChaosLargeSources checks size boundaries without panicking.
func TestChaosLargeSources(t *testing.T) {
	t.Parallel()
	identifier := "A"
	for i := 0; i < 2000; i++ {
		identifier += "x"
	}
	longIdent := "package p\n\nvar " + identifier + " = 1\n"
	sources := []string{
		"",
		"// only a comment\n",
		"\xef\xbb\xbfpackage p\n",
		longIdent,
		"package p\n\nfunc F() { " + repeatText("_ = 1; ", 2000) + "}\n",
	}
	frontend := NewGoFrontend().WithTypeCheck(false)
	for i, src := range sources {
		unit, _, err := frontend.Parse(air.SourceUnit{
			Language: "go",
			Files:    []air.SourceFile{air.NewSourceFile("x.go", []byte(src))},
		})
		if err != nil || unit == nil || unit.Roots[0] == nil {
			t.Fatalf("boundary source %d failed: %v", i, err)
		}
	}
}

// repeatText repeats a literal a fixed number of times.
func repeatText(text string, count int) string {
	out := make([]byte, 0, len(text)*count)
	for i := 0; i < count; i++ {
		out = append(out, text...)
	}
	return string(out)
}
