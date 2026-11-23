/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package main

import (
	"fmt"
	"io/fs"
	"os"
	"sort"

	"abate/internal/corpus"
	"abate/internal/e2e"
	"abate/lang/air"
	"abate/lang/frontkit/grammar"
)

// locateRow pins one observed degraded construct to the corpus file and the
// node position that produced it.
type locateRow struct {
	Node     string
	Kind     air.NodeKind
	File     string
	Position string
}

// locate scans the corpus and reports the first concrete file positions of
// every observed degraded construct of one language, capped for readability.
func locate(lang string, max int) error {
	facts, err := grammar.Load(lang)
	if err != nil {
		return err
	}
	mappings := grammar.Map(facts, grammar.DefaultRules(lang))
	kindToNodes := make(map[air.NodeKind][]string)
	for _, m := range mappings {
		if m.Status == grammar.StatusDegraded {
			kindToNodes[m.Kind] = append(kindToNodes[m.Kind], m.Node)
		}
	}
	if len(kindToNodes) == 0 {
		fmt.Printf("== %s: no degraded mappings ==\n", lang)
		return nil
	}
	m, err := corpus.Load(manifestFile)
	if err != nil {
		return err
	}
	ft, err := e2e.NewFrontendTable()
	if err != nil {
		return err
	}
	fe, dialect, err := ft.For(lang)
	if err != nil {
		return err
	}
	fsys := os.DirFS(corpus.DefaultCorpusRoot)
	found := make(map[air.NodeKind][]locateRow)
	for _, entry := range m.Entries {
		if entry.Language != lang {
			continue
		}
		data, err := fs.ReadFile(fsys, entry.RelPath)
		if err != nil {
			return err
		}
		unit, _, err := fe.Parse(air.SourceUnit{
			Language: lang,
			Dialect:  dialect,
			Files:    []air.SourceFile{air.NewSourceFile(entry.RelPath, data)},
		})
		if err != nil {
			continue
		}
		for _, node := range air.Collect(unit) {
			kind := node.Kind()
			nodes, isDegraded := kindToNodes[kind]
			if !isDegraded {
				continue
			}
			if len(found[kind]) >= max {
				continue
			}
			pos := node.Pos()
			found[kind] = append(found[kind], locateRow{
				Node:     joinNodes(nodes),
				Kind:     kind,
				File:     entry.RelPath,
				Position: fmt.Sprintf("L%d:C%d [%d,%d)", pos.Line, pos.Column, pos.Offset, pos.EndOffset),
			})
		}
	}
	fmt.Printf("== %s degraded construct locations (first %d each) ==\n", lang, max)
	kinds := make([]air.NodeKind, 0, len(found))
	for k := range found {
		kinds = append(kinds, k)
	}
	sort.Slice(kinds, func(i, j int) bool { return kinds[i] < kinds[j] })
	for _, k := range kinds {
		rows := found[k]
		sort.Slice(rows, func(i, j int) bool { return rows[i].File < rows[j].File })
		fmt.Printf("  kind=%s nodes=%s\n", k, rows[0].Node)
		for _, r := range rows {
			fmt.Printf("    %s  %s\n", r.Position, r.File)
		}
	}
	return nil
}

// joinNodes renders the upstream names sharing one degraded kind.
func joinNodes(nodes []string) string {
	out := ""
	for i, n := range nodes {
		if i > 0 {
			out += "|"
		}
		out += n
	}
	return out
}

// runLocate executes the locator for one requested language.
func runLocate(lang string, max int) {
	if err := locate(lang, max); err != nil {
		fmt.Fprintf(os.Stderr, "locate %s: %v\n", lang, err)
		os.Exit(1)
	}
}
