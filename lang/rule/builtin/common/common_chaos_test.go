/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package common

import (
	"errors"
	"slices"
	"strings"
	"sync"
	"testing"

	"abate/lang/rule"
)

// chaosTemplate is one well-formed language-local template used as batch input.
func chaosTemplate(seed int) Template {
	return Template{
		Family:   "naming",
		Pattern:  `Ident^chaos` + string(rune('a'+seed%26)),
		Message:  "chaos local template",
		Severity: rule.SeverityInfo,
		Category: rule.CategoryStyle,
		Positive: "Ident:chaos" + string(rune('a'+seed%26)),
		Negative: "Ident:tidy",
	}
}

// TestChaosBatchOrdering proves shuffled, duplicated and empty batches never
// panic and always yield the same deterministic identifier sequence.
func TestChaosBatchOrdering(t *testing.T) {
	t.Parallel()
	base := Generate("PY", nil)
	if len(base) != templateCount {
		t.Fatalf("baseline batch size = %d, want %d", len(base), templateCount)
	}
	local := []Template{chaosTemplate(0), chaosTemplate(1), chaosTemplate(2)}
	batches := [][]Template{
		nil,
		{},
		local,
		slices.Clone(local),
		{local[2], local[0], local[1]},
		{local[0], local[0], local[1], local[1], local[2], local[2]},
	}
	for i, batch := range batches {
		first := Generate("py", batch)
		second := Generate("PY", batch)
		if len(first) != templateCount+uniqueCount(batch) {
			t.Fatalf("batch %d produced %d specs", i, len(first))
		}
		if len(first) != len(second) {
			t.Fatalf("batch %d is not repeatable", i)
		}
		for j := range first {
			if first[j].ID != second[j].ID || first[j].Canonical() != second[j].Canonical() {
				t.Fatalf("batch %d drifted at spec %d", i, j)
			}
		}
		for j := 1; j < len(first); j++ {
			if strings.Compare(string(first[j-1].ID), string(first[j].ID)) > 0 {
				t.Fatalf("batch %d is not id-ordered at %d", i, j)
			}
		}
	}
}

// uniqueCount counts distinct declaration keys in one batch.
func uniqueCount(batch []Template) int {
	seen := make(map[string]struct{}, len(batch))
	for _, item := range batch {
		seen[item.Key()] = struct{}{}
	}
	return len(seen)
}

// TestChaosOversizedBatch proves an oversized batch fails closed without panic.
func TestChaosOversizedBatch(t *testing.T) {
	t.Parallel()
	oversized := make([]Template, 0, 1000)
	for i := 0; i < 1000; i++ {
		item := chaosTemplate(i)
		item.Pattern = `Ident^over` + itoa(i)
		oversized = append(oversized, item)
	}
	if got := Generate("PY", oversized); got != nil {
		t.Fatalf("oversized batch returned %d specs, want nil", len(got))
	}
	if _, err := GenerateChecked("PY", oversized); !errors.Is(err, ErrTooManyTemplates) {
		t.Fatalf("oversized batch error = %v, want ErrTooManyTemplates", err)
	}
}

// itoa formats a small non-negative integer without a float or clock.
func itoa(value int) string {
	const digits = "0123456789"
	if value == 0 {
		return "0"
	}
	buf := [8]byte{}
	pos := len(buf)
	for value > 0 {
		pos--
		buf[pos] = digits[value%10]
		value /= 10
	}
	return string(buf[pos:])
}

// TestChaosInvalidTags proves malformed tags fail closed and never panic.
func TestChaosInvalidTags(t *testing.T) {
	t.Parallel()
	for _, tag := range []string{"py-js", "p y", "!", "c++", "py.1"} {
		if got := Generate(tag, nil); got != nil {
			t.Fatalf("Generate(%q) returned %d specs, want nil", tag, len(got))
		}
		if _, err := GenerateChecked(tag, nil); !errors.Is(err, ErrInvalidTag) {
			t.Fatalf("GenerateChecked(%q) error = %v, want ErrInvalidTag", tag, err)
		}
	}
}

// TestChaosConcurrentRebuild proves concurrent read-only rebuilds agree.
func TestChaosConcurrentRebuild(t *testing.T) {
	t.Parallel()
	base := Fingerprint()
	var group sync.WaitGroup
	digests := make(chan [32]byte, 32)
	for i := 0; i < 32; i++ {
		group.Add(1)
		go func() {
			defer group.Done()
			digests <- Fingerprint()
		}()
	}
	group.Wait()
	close(digests)
	for digest := range digests {
		if digest != base {
			t.Fatal("concurrent rebuild produced a different digest")
		}
	}
}
