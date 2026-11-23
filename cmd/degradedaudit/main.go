/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// degradedaudit enumerates every degraded upstream construct that a real
// full-corpus parse actually observed, per language. It reuses the same
// frontend table, grammar facts and coverage tallies as degradecheck, but
// reports the construct-level worklist instead of the aggregate percentage,
// so remediation can be driven construct by construct.
package main

import (
	"fmt"
	"io/fs"
	"os"
	"sort"

	"abate/internal/corpus"
	"abate/internal/e2e"
	"abate/lang/air"
	fkcov "abate/lang/frontkit/coverage"
	"abate/lang/frontkit/grammar"
)

// manifestFile is the frozen full-corpus registry shared with degradecheck.
const manifestFile = "ci/corpus_manifest.json"

// reportLanguages is the canonical seven-language order.
var reportLanguages = []string{"go", "py", "js", "c", "cs", "rs", "kt"}

// observedTally accumulates the AIR kinds one language's corpus produced.
type observedTally map[string]map[air.NodeKind]int

// run parses every registered corpus entry with its real frontend and
// returns the per-language kind tallies. A read or parse failure aborts the
// audit: a partial worklist would silently understate the remediation set.
func run() (observedTally, error) {
	m, err := corpus.Load(manifestFile)
	if err != nil {
		return nil, err
	}
	ft, err := e2e.NewFrontendTable()
	if err != nil {
		return nil, err
	}
	fsys := os.DirFS(corpus.DefaultCorpusRoot)
	tally := observedTally{}
	for _, lang := range reportLanguages {
		tally[lang] = make(map[air.NodeKind]int)
	}
	for _, entry := range m.Entries {
		if err := observe(fsys, ft, tally, entry); err != nil {
			return nil, err
		}
	}
	return tally, nil
}

// observe folds one corpus file's observed kinds into the language tally.
func observe(fsys fs.FS, ft e2e.FrontendTable, tally observedTally, entry corpus.Entry) error {
	fe, dialect, err := ft.For(entry.Language)
	if err != nil {
		return fmt.Errorf("%s: %v", entry.RelPath, err)
	}
	data, err := fs.ReadFile(fsys, entry.RelPath)
	if err != nil {
		return fmt.Errorf("%s: %v", entry.RelPath, err)
	}
	unit, _, err := fe.Parse(air.SourceUnit{
		Language: entry.Language,
		Dialect:  dialect,
		Files:    []air.SourceFile{air.NewSourceFile(entry.RelPath, data)},
	})
	if err != nil {
		return fmt.Errorf("%s: %v", entry.RelPath, err)
	}
	tally[entry.Language] = fkcov.Merge(tally[entry.Language], fkcov.Observe(unit))
	return nil
}

// degradedRow is one observed degraded construct with its tally.
type degradedRow struct {
	Node  string
	Kind  air.NodeKind
	Count int
	Rule  string
}

// audit prints the observed degraded constructs of one language.
func audit(lang string, tally observedTally) error {
	facts, err := grammar.Load(lang)
	if err != nil {
		// Go carries no embedded grammar facts: nothing to audit.
		if lang == "go" {
			return nil
		}
		return err
	}
	mappings := grammar.Map(facts, grammar.DefaultRules(lang))
	rows := make([]degradedRow, 0)
	for _, mapping := range mappings {
		if mapping.Status != grammar.StatusDegraded {
			continue
		}
		count := tally[lang][mapping.Kind]
		if count == 0 {
			continue
		}
		rows = append(rows, degradedRow{
			Node:  mapping.Node,
			Kind:  mapping.Kind,
			Count: count,
			Rule:  mapping.Rule,
		})
	}
	sort.Slice(rows, func(i, j int) bool {
		if rows[i].Kind != rows[j].Kind {
			return rows[i].Kind < rows[j].Kind
		}
		return rows[i].Node < rows[j].Node
	})
	fmt.Printf("== %s: %d observed degraded constructs ==\n", lang, len(rows))
	for _, row := range rows {
		fmt.Printf("  %-36s kind=%-24s observed=%d rule=%s\n",
			row.Node, row.Kind, row.Count, row.Rule)
	}
	return nil
}

func main() {
	tally, err := run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "degradedaudit: %v\n", err)
		os.Exit(1)
	}
	if len(os.Args) > 1 && os.Args[1] == "xcheck" {
		runXcheck(tally)
		return
	}
	if len(os.Args) > 1 && os.Args[1] == "locate" {
		lang := "js"
		if len(os.Args) > 2 {
			lang = os.Args[2]
		}
		runLocate(lang, 3)
		return
	}
	for _, lang := range reportLanguages {
		if err := audit(lang, tally); err != nil {
			fmt.Fprintf(os.Stderr, "degradedaudit: %v\n", err)
			os.Exit(1)
		}
	}
}
