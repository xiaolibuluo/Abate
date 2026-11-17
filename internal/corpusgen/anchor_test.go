/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package corpusgen

import (
	"errors"
	"testing"
)

func TestFirstTokenOffsetUsesWordBoundaries(t *testing.T) {
	t.Parallel()
	if got := firstTokenOffset([]byte("abc 1 + 2 == 3"), []string{"1"}); got != 4 {
		t.Fatalf("digit offset = %d", got)
	}
	if got := firstTokenOffset([]byte("step1 x"), []string{"1"}); got != -1 {
		t.Fatalf("digit inside identifier = %d", got)
	}
	if got := firstTokenOffset([]byte("xx == yy"), []string{"=="}); got != 3 {
		t.Fatalf("cmp offset = %d", got)
	}
	if got := firstTokenOffset([]byte("abc"), []string{""}); got != -1 {
		t.Fatal("empty token accepted")
	}
}

func TestHeaderHasToken(t *testing.T) {
	t.Parallel()
	if HeaderHasToken("package synth\n\n") {
		t.Fatal("plain package header rejected")
	}
	for _, bad := range []string{"version 2\n", "a-b\n", "a<b\n", "f()\n"} {
		if !HeaderHasToken(bad) {
			t.Fatalf("header %q accepted", bad)
		}
	}
}

func TestCaptureAnchorPreflight(t *testing.T) {
	t.Parallel()
	content := []byte("package synth\n\nfunc anchor() bool {\n\treturn 1 + 2 == 3 && f()\n}\n")
	start := len("package synth\n\n")
	end := len(content)
	full := CaptureAnchor{Language: "go", Bounds: AnchorBounds{Start: start, End: end}}
	if err := full.Preflight(content); err != nil {
		t.Fatalf("full preflight: %v", err)
	}
	clean := CaptureAnchor{Language: "go", Clean: true,
		Bounds: AnchorBounds{Start: start, End: end}}
	if err := clean.Preflight(content); err != nil {
		t.Fatalf("clean preflight: %v", err)
	}
	bad := CaptureAnchor{Language: "go", Bounds: AnchorBounds{Start: 40, End: 50}}
	if err := bad.Preflight(content); !errors.Is(err, errAnchor) {
		t.Fatalf("out-of-range anchor accepted: %v", err)
	}
	noDigit := CaptureAnchor{Language: "go", Bounds: AnchorBounds{Start: 0, End: end}}
	if err := noDigit.Preflight([]byte("package synth\n")); !errors.Is(err, errAnchor) {
		t.Fatalf("missing digit accepted: %v", err)
	}
}

func TestLabeledRNGIsDeterministic(t *testing.T) {
	t.Parallel()
	first := NewLabeledRNG(20261114, "go:0:0")
	second := NewLabeledRNG(20261114, "go:0:0")
	for i := 0; i < 20; i++ {
		if first.Intn(97) != second.Intn(97) {
			t.Fatalf("stream diverged at %d", i)
		}
	}
	if first.Label() != "go:0:0" {
		t.Fatalf("label = %q", first.Label())
	}
	single := NewLabeledRNG(1, "a")
	if single.Intn(1) != 0 {
		t.Fatal("Intn(1) must be zero")
	}
}
