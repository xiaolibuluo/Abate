/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package main

import (
	"fmt"
	"os"
	"sort"

	"abate/lang/air"
	"abate/lang/frontkit/grammar"
)

// xcheck verifies the integrity of the degraded tally by cross-referencing
// every mapping of one language against the observed kind histogram. It
// reports, for each kind, which exact and degraded constructs share it, so a
// degraded observation inflated by an exact sibling becomes visible instead
// of masquerading as a real fidelity gap.
func xcheck(lang string, tally observedTally) error {
	facts, err := grammar.Load(lang)
	if err != nil {
		return err
	}
	mappings := grammar.Map(facts, grammar.DefaultRules(lang))
	type bucket struct {
		kind     air.NodeKind
		exact    []string
		degraded []string
		count    int
	}
	buckets := map[air.NodeKind]*bucket{}
	for _, m := range mappings {
		if m.Kind == air.KindInvalid {
			continue
		}
		b, ok := buckets[m.Kind]
		if !ok {
			b = &bucket{kind: m.Kind}
			buckets[m.Kind] = b
		}
		switch m.Status {
		case grammar.StatusExact:
			b.exact = append(b.exact, m.Node)
		case grammar.StatusDegraded:
			b.degraded = append(b.degraded, m.Node)
		}
	}
	list := make([]*bucket, 0, len(buckets))
	for _, b := range buckets {
		b.count = tally[lang][b.kind]
		if len(b.degraded) > 0 && b.count > 0 {
			list = append(list, b)
		}
	}
	sort.Slice(list, func(i, j int) bool { return list[i].count > list[j].count })
	fmt.Printf("== %s shared-kind analysis ==\n", lang)
	for _, b := range list {
		marker := "REAL"
		if len(b.exact) > 0 {
			marker = "AMBIGUOUS(exact sibling present)"
		}
		fmt.Printf("  kind=%-22s observed=%-8d %s\n", b.kind, b.count, marker)
		if len(b.exact) > 0 {
			fmt.Printf("      exact:    %v\n", b.exact)
		}
		fmt.Printf("      degraded: %v\n", b.degraded)
	}
	return nil
}

// runXcheck executes the cross-check for the six measured languages.
func runXcheck(tally observedTally) {
	for _, lang := range []string{"py", "js", "c", "cs", "rs", "kt"} {
		if err := xcheck(lang, tally); err != nil {
			fmt.Fprintf(os.Stderr, "xcheck %s: %v\n", lang, err)
			os.Exit(1)
		}
	}
}
